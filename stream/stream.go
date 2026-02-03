package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chainstream-io/chainstream-go-sdk/openapi/token"
	"github.com/chainstream-io/centrifuge-go"
)

// RequestContext represents the request context for WebSocket connections
type RequestContext struct {
	BaseUrl       string
	StreamUrl     string
	AccessToken   string
	TokenProvider func() (string, error)
}

// StreamApi handles WebSocket streaming functionality using Centrifuge
type StreamApi struct {
	client       *centrifuge.Client
	listenersMap map[string][]StreamCallback[interface{}]
	mutex        sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
	requestCtx   *RequestContext
	connected    bool
	connectMutex sync.Mutex
}

// StreamCallback represents a callback function for stream data
type StreamCallback[T any] func(data T)

// Unsubscribe represents an unsubscribe function
type Unsubscribe func()

// NewStreamApi creates a new StreamApi instance
func NewStreamApi(requestCtx *RequestContext) *StreamApi {
	ctx, cancel := context.WithCancel(context.Background())
	return &StreamApi{
		listenersMap: make(map[string][]StreamCallback[interface{}]),
		ctx:          ctx,
		cancel:       cancel,
		requestCtx:   requestCtx,
	}
}

// buildWsUrl builds WebSocket URL with token as query parameter
func (s *StreamApi) buildWsUrl(endpoint string, token string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return endpoint
	}
	q := u.Query()
	q.Set("token", token)
	u.RawQuery = q.Encode()
	return u.String()
}

// Connect establishes WebSocket connection to the server using Centrifuge.
// This is automatically called when you use subscribe methods if not already connected.
// You can also call this method manually for explicit control.
func (s *StreamApi) Connect() error {
	s.connectMutex.Lock()
	defer s.connectMutex.Unlock()

	if s.connected {
		return nil
	}

	return s.connectInternal()
}

// Disconnect closes the WebSocket connection
func (s *StreamApi) Disconnect() error {
	s.connectMutex.Lock()
	defer s.connectMutex.Unlock()

	s.cancel()
	if s.client != nil {
		s.client.Close()
	}
	s.connected = false
	return nil
}

// IsConnected returns whether the WebSocket is connected
func (s *StreamApi) IsConnected() bool {
	s.connectMutex.Lock()
	defer s.connectMutex.Unlock()
	return s.connected
}

// ensureConnected ensures WebSocket is connected, auto-connect if not
func (s *StreamApi) ensureConnected() error {
	s.connectMutex.Lock()
	defer s.connectMutex.Unlock()

	if s.connected {
		return nil
	}

	// Need to connect
	return s.connectInternal()
}

// connectInternal is the internal connect logic (must be called with connectMutex held)
func (s *StreamApi) connectInternal() error {
	// Get access token
	var token string
	var err error
	if s.requestCtx.TokenProvider != nil {
		token, err = s.requestCtx.TokenProvider()
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
	} else {
		token = s.requestCtx.AccessToken
	}

	// Build WebSocket URL with token
	wsUrl := s.buildWsUrl(s.requestCtx.StreamUrl, token)

	// Create Centrifuge client configuration
	config := centrifuge.Config{
		Token: token,
		GetToken: func(ctx centrifuge.ConnectionTokenEvent) (string, error) {
			if s.requestCtx.TokenProvider != nil {
				return s.requestCtx.TokenProvider()
			}
			return s.requestCtx.AccessToken, nil
		},
		ReadTimeout:      30 * time.Second,
		WriteTimeout:     30 * time.Second,
		HandshakeTimeout: 30 * time.Second,
	}

	// Create Centrifuge client with token in URL
	client := centrifuge.NewJsonClient(wsUrl, config)

	// Set up event handlers
	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		log.Println("[streaming] connected")
	})

	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		log.Printf("[streaming] disconnected: %v", e.Reason)
	})

	client.OnError(func(e centrifuge.ErrorEvent) {
		log.Printf("[streaming] error: %v", e.Error)
	})

	s.client = client

	// Connect to server
	if err := client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	s.connected = true
	return nil
}

// Subscribe subscribes to a channel using Centrifuge protocol
func (s *StreamApi) Subscribe(channel string, callback StreamCallback[interface{}], filter, methodName string) Unsubscribe {
	// Ensure connected before subscribing
	if err := s.ensureConnected(); err != nil {
		log.Printf("[streaming] error connecting: %v", err)
		return func() {}
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Process filter expression
	processedFilter := filter
	if filter != "" && methodName != "" {
		processedFilter = ReplaceFilterFields(filter, methodName)
	}

	// Create subscription options
	opts := centrifuge.SubscriptionConfig{
		Delta: "fossil",
	}
	if processedFilter != "" {
		opts.Filter = processedFilter
	}

	// Create subscription
	sub, err := s.client.NewSubscription(channel, opts)
	if err != nil {
		log.Printf("[streaming] error creating subscription: %v", err)
		return func() {}
	}

	// Set up subscription event handlers
	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		log.Printf("[streaming] subscribed to channel: %s", channel)
	})

	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		log.Printf("[streaming] unsubscribed from channel: %s", channel)
	})

	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Printf("[streaming] received publication on channel %s, data length: %d", channel, len(e.Data))

		// Parse JSON data from bytes
		var jsonData interface{}
		if err := json.Unmarshal(e.Data, &jsonData); err != nil {
			log.Printf("[streaming] error parsing JSON: %v, raw data: %s", err, string(e.Data))
			return
		}

		log.Printf("[streaming] parsed data: %v", jsonData)

		// Dispatch message to listeners
		s.dispatchMessage(channel, jsonData)
	})

	// Subscribe to the channel
	if err := sub.Subscribe(); err != nil {
		log.Printf("[streaming] error subscribing: %v", err)
		return func() {}
	}

	// Add listener
	if listeners, exists := s.listenersMap[channel]; exists {
		s.listenersMap[channel] = append(listeners, callback)
	} else {
		s.listenersMap[channel] = []StreamCallback[interface{}]{callback}
	}

	// Return unsubscribe function
	return func() {
		s.Unsubscribe(channel, callback)
	}
}

// Unsubscribe unsubscribes from a channel
func (s *StreamApi) Unsubscribe(channel string, callback StreamCallback[interface{}]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	listeners, exists := s.listenersMap[channel]
	if !exists {
		return
	}

	// Remove listener
	for i, listener := range listeners {
		if &listener == &callback {
			s.listenersMap[channel] = append(listeners[:i], listeners[i+1:]...)
			break
		}
	}

	// If no listeners left, unsubscribe
	if len(s.listenersMap[channel]) == 0 {
		delete(s.listenersMap, channel)

		// Get subscription and unsubscribe
		if sub, exists := s.client.GetSubscription(channel); exists && sub != nil {
			sub.Unsubscribe()
		}

		log.Printf("[streaming] unsubscribed from channel: %s", channel)
	}
}

// dispatchMessage distributes messages to listeners
func (s *StreamApi) dispatchMessage(channel string, data interface{}) {
	s.mutex.RLock()
	listeners, exists := s.listenersMap[channel]
	listenerCount := len(listeners)
	s.mutex.RUnlock()

	log.Printf("[streaming] dispatchMessage for channel %s, listeners exist: %v, count: %d", channel, exists, listenerCount)

	if !exists {
		return
	}

	for i, listener := range listeners {
		log.Printf("[streaming] calling listener %d for channel %s", i, channel)
		go func(callback StreamCallback[interface{}]) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[streaming] panic in callback: %v", r)
				}
			}()
			callback(data)
		}(listener)
	}
}

// formatScientificNotation formats scientific notation numbers
func (s *StreamApi) formatScientificNotation(value interface{}) string {
	if value == nil {
		return "0"
	}

	strValue := fmt.Sprintf("%v", value)
	if strings.Contains(strValue, "e-") || strings.Contains(strValue, "E-") {
		// Try to parse as number
		if num, err := strconv.ParseFloat(strValue, 64); err == nil {
			return fmt.Sprintf("%.20f", num)
		}
	}

	return strValue
}

// SubscribeTokenCandles subscribes to token candle data
func (s *StreamApi) SubscribeTokenCandles(chain, tokenAddress string, resolution token.Resolution, callback StreamCallback[TokenCandle], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-candle:%s_%s_%s", chain, tokenAddress, resolution)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		candle := TokenCandle{
			Open:       s.formatScientificNotation(dataMap["o"]),
			Close:      s.formatScientificNotation(dataMap["c"]),
			High:       s.formatScientificNotation(dataMap["h"]),
			Low:        s.formatScientificNotation(dataMap["l"]),
			Volume:     s.formatScientificNotation(dataMap["v"]),
			Resolution: getString(dataMap, "r"),
			Time:       getInt64(dataMap, "t"),
			Number:     getInt(dataMap, "n"),
		}

		callback(candle)
	}, filter, "subscribeTokenCandles")
}

// SubscribeTokenStats subscribes to token statistics
func (s *StreamApi) SubscribeTokenStats(chain, tokenAddress string, callback StreamCallback[TokenStat], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-stats:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		stat := TokenStat{
			Address:   getString(dataMap, "a"),
			Timestamp: getInt64(dataMap, "t"),
		}

		// Set 1-minute data
		if val, ok := dataMap["b1m"]; ok {
			if v, ok := val.(float64); ok {
				stat.Buys1m = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["s1m"]; ok {
			if v, ok := val.(float64); ok {
				stat.Sells1m = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["be1m"]; ok {
			if v, ok := val.(float64); ok {
				stat.Buyers1m = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["se1m"]; ok {
			if v, ok := val.(float64); ok {
				stat.Sellers1m = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["bviu1m"]; ok {
			stat.BuyVolumeInUsd1m = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["sviu1m"]; ok {
			stat.SellVolumeInUsd1m = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["p1m"]; ok {
			stat.Price1m = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["oiu1m"]; ok {
			stat.OpenInUsd1m = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["ciu1m"]; ok {
			stat.CloseInUsd1m = PtrString(s.formatScientificNotation(val))
		}

		// Set current price
		if val, ok := dataMap["p"]; ok {
			stat.Price = PtrString(s.formatScientificNotation(val))
		}

		callback(stat)
	}, filter, "subscribeTokenStats")
}

// SubscribeTokenHolders subscribes to token holder data
func (s *StreamApi) SubscribeTokenHolders(chain, tokenAddress string, callback StreamCallback[TokenHolder], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-holding:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		holder := TokenHolder{
			TokenAddress: getString(dataMap, "a"),
			Timestamp:    getInt64(dataMap, "ts"),
		}

		if val, ok := dataMap["h"]; ok {
			if v, ok := val.(float64); ok {
				holder.Holders = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["t100a"]; ok {
			holder.Top100Amount = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["t10a"]; ok {
			holder.Top10Amount = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["t100h"]; ok {
			if v, ok := val.(float64); ok {
				holder.Top100Holders = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["t10h"]; ok {
			if v, ok := val.(float64); ok {
				holder.Top10Holders = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["t100r"]; ok {
			holder.Top100Ratio = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["t10r"]; ok {
			holder.Top10Ratio = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["ch"]; ok {
			if v, ok := val.(float64); ok {
				holder.CreatorsHolders = PtrInt(int(v))
			}
		}
		if val, ok := dataMap["ca"]; ok {
			holder.CreatorsAmount = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["cr"]; ok {
			holder.CreatorsRatio = PtrString(s.formatScientificNotation(val))
		}

		callback(holder)
	}, filter, "subscribeTokenHolders")
}

// SubscribeNewToken subscribes to new token events
func (s *StreamApi) SubscribeNewToken(chain string, callback StreamCallback[NewToken], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-new-token:%s", chain)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		token := NewToken{
			TokenAddress: getString(dataMap, "a"),
			Name:         getString(dataMap, "n"),
			Symbol:       getString(dataMap, "s"),
			CreatedAtMs:  getInt64(dataMap, "cts"),
		}

		if val, ok := dataMap["dec"]; ok {
			if v, ok := val.(float64); ok {
				token.Decimals = PtrInt(int(v))
			}
		}

		if val, ok := dataMap["lf"]; ok {
			if lfMap, ok := val.(map[string]interface{}); ok {
				token.LaunchFrom = &DexProtocol{}
				if pa, ok := lfMap["pa"]; ok {
					token.LaunchFrom.ProgramAddress = PtrString(pa.(string))
				}
				if pf, ok := lfMap["pf"]; ok {
					token.LaunchFrom.ProtocolFamily = PtrString(pf.(string))
				}
				if pn, ok := lfMap["pn"]; ok {
					token.LaunchFrom.ProtocolName = PtrString(pn.(string))
				}
			}
		}

		callback(token)
	}, filter, "subscribeNewToken")
}

// SubscribeWalletBalance subscribes to wallet balance data
func (s *StreamApi) SubscribeWalletBalance(chain, walletAddress string, callback StreamCallback[WalletBalance], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-wallet-balance:%s_%s", chain, walletAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		balance := WalletBalance{
			WalletAddress:   getString(dataMap, "a"),
			TokenAddress:    getString(dataMap, "ta"),
			TokenPriceInUsd: s.formatScientificNotation(dataMap["tpiu"]),
			Balance:         s.formatScientificNotation(dataMap["b"]),
			Timestamp:       getInt64(dataMap, "t"),
		}

		callback(balance)
	}, filter, "subscribeWalletBalance")
}

// SubscribeWalletPnl subscribes to wallet PnL data
func (s *StreamApi) SubscribeWalletPnl(chain, walletAddress string, callback StreamCallback[WalletTokenPnl], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-wallet-token-pnl:%s_%s", chain, walletAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		pnl := WalletTokenPnl{
			WalletAddress:            getString(dataMap, "a"),
			TokenAddress:             getString(dataMap, "ta"),
			TokenPriceInUsd:          s.formatScientificNotation(dataMap["tpiu"]),
			Timestamp:                getInt64(dataMap, "t"),
			OpenTime:                 getInt64(dataMap, "ot"),
			LastTime:                 getInt64(dataMap, "lt"),
			CloseTime:                getInt64(dataMap, "ct"),
			BuyAmount:                s.formatScientificNotation(dataMap["ba"]),
			BuyAmountInUsd:           s.formatScientificNotation(dataMap["baiu"]),
			BuyCount:                 getInt(dataMap, "bs"),
			BuyCount30d:              getInt(dataMap, "bs30d"),
			BuyCount7d:               getInt(dataMap, "bs7d"),
			SellAmount:               s.formatScientificNotation(dataMap["sa"]),
			SellAmountInUsd:          s.formatScientificNotation(dataMap["saiu"]),
			SellCount:                getInt(dataMap, "ss"),
			SellCount30d:             getInt(dataMap, "ss30d"),
			SellCount7d:              getInt(dataMap, "ss7d"),
			HeldDurationTimestamp:    getInt64(dataMap, "hdts"),
			AverageBuyPriceInUsd:     s.formatScientificNotation(dataMap["abpiu"]),
			AverageSellPriceInUsd:    s.formatScientificNotation(dataMap["aspiu"]),
			UnrealizedProfitInUsd:    s.formatScientificNotation(dataMap["upiu"]),
			UnrealizedProfitRatio:    s.formatScientificNotation(dataMap["upr"]),
			RealizedProfitInUsd:      s.formatScientificNotation(dataMap["rpiu"]),
			RealizedProfitRatio:      s.formatScientificNotation(dataMap["rpr"]),
			TotalRealizedProfitInUsd: s.formatScientificNotation(dataMap["trpiu"]),
			TotalRealizedProfitRatio: s.formatScientificNotation(dataMap["trr"]),
		}

		callback(pnl)
	}, filter, "subscribeWalletPnl")
}

// SubscribeTokenTrade subscribes to token trade data
func (s *StreamApi) SubscribeTokenTrade(chain, tokenAddress string, callback StreamCallback[TradeActivity], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-trade:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		trade := TradeActivity{
			TokenAddress:      getString(dataMap, "a"),
			Timestamp:         getInt64(dataMap, "t"),
			Kind:              getString(dataMap, "k"),
			BuyAmount:         s.formatScientificNotation(dataMap["ba"]),
			BuyAmountInUsd:    s.formatScientificNotation(dataMap["baiu"]),
			BuyTokenAddress:   getString(dataMap, "btma"),
			BuyTokenName:      getString(dataMap, "btn"),
			BuyTokenSymbol:    getString(dataMap, "bts"),
			BuyWalletAddress:  getString(dataMap, "bwa"),
			SellAmount:        s.formatScientificNotation(dataMap["sa"]),
			SellAmountInUsd:   s.formatScientificNotation(dataMap["saiu"]),
			SellTokenAddress:  getString(dataMap, "stma"),
			SellTokenName:     getString(dataMap, "stn"),
			SellTokenSymbol:   getString(dataMap, "sts"),
			SellWalletAddress: getString(dataMap, "swa"),
			TxHash:            getString(dataMap, "h"),
		}

		callback(trade)
	}, filter, "subscribeTokenTrades")
}

// SubscribeDexPoolBalance subscribes to DEX pool balance data
func (s *StreamApi) SubscribeDexPoolBalance(chain, poolAddress string, callback StreamCallback[DexPoolBalance]) Unsubscribe {
	channel := fmt.Sprintf("dex-pool-balance:%s_%s", chain, poolAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		balance := DexPoolBalance{
			PoolAddress:          getString(dataMap, "a"),
			TokenAAddress:        getString(dataMap, "taa"),
			TokenALiquidityInUsd: s.formatScientificNotation(dataMap["taliu"]),
			TokenBAddress:        getString(dataMap, "tba"),
			TokenBLiquidityInUsd: s.formatScientificNotation(dataMap["tbliu"]),
		}

		callback(balance)
	}, "", "subscribeDexPoolBalance")
}

// SubscribeNewTokensMetadata subscribes to new tokens metadata
func (s *StreamApi) SubscribeNewTokensMetadata(chain string, callback StreamCallback[[]TokenMetadata]) Unsubscribe {
	channel := fmt.Sprintf("dex-new-tokens-metadata:%s", chain)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenMetadata
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			metadata := TokenMetadata{
				TokenAddress: getString(itemMap, "a"),
			}
			if val, ok := itemMap["n"]; ok {
				metadata.Name = PtrString(val.(string))
			}
			if val, ok := itemMap["s"]; ok {
				metadata.Symbol = PtrString(val.(string))
			}
			if val, ok := itemMap["iu"]; ok {
				metadata.ImageUrl = PtrString(val.(string))
			}
			if val, ok := itemMap["de"]; ok {
				metadata.Description = PtrString(val.(string))
			}
			if val, ok := itemMap["cts"]; ok {
				cts := int64(val.(float64))
				metadata.CreatedAtMs = &cts
			}
			if sm, ok := itemMap["sm"].(map[string]interface{}); ok {
				metadata.SocialMedia = &SocialMedia{}
				if v, ok := sm["tw"]; ok {
					metadata.SocialMedia.Twitter = PtrString(v.(string))
				}
				if v, ok := sm["tg"]; ok {
					metadata.SocialMedia.Telegram = PtrString(v.(string))
				}
				if v, ok := sm["w"]; ok {
					metadata.SocialMedia.Website = PtrString(v.(string))
				}
			}

			result = append(result, metadata)
		}

		callback(result)
	}, "", "subscribeNewTokensMetadata")
}

// SubscribeTokenSupply subscribes to token supply data
func (s *StreamApi) SubscribeTokenSupply(chain, tokenAddress string, callback StreamCallback[TokenSupply], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-supply:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		supply := TokenSupply{
			TokenAddress: getString(dataMap, "a"),
			Timestamp:    getInt64(dataMap, "ts"),
		}
		if val, ok := dataMap["s"]; ok {
			supply.Supply = PtrString(s.formatScientificNotation(val))
		}
		if val, ok := dataMap["mc"]; ok {
			supply.MarketCapInUsd = PtrString(s.formatScientificNotation(val))
		}

		callback(supply)
	}, filter, "subscribeTokenSupply")
}

// SubscribeTokenLiquidity subscribes to token liquidity data
func (s *StreamApi) SubscribeTokenLiquidity(chain, tokenAddress string, callback StreamCallback[TokenLiquidity], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-general-stat-num:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		liquidity := TokenLiquidity{
			TokenAddress: getString(dataMap, "a"),
			MetricType:   MetricType(getString(dataMap, "t")),
			Value:        s.formatScientificNotation(dataMap["v"]),
			Timestamp:    getInt64(dataMap, "ts"),
		}

		callback(liquidity)
	}, filter, "subscribeTokenLiquidity")
}

// SubscribeTokenMaxLiquidity subscribes to token max liquidity data
// Pushes the max liquidity info of a token in a single pool
// Channel: dex-token-liquidity:{chain}_{token_address}
func (s *StreamApi) SubscribeTokenMaxLiquidity(chain, tokenAddress string, callback StreamCallback[TokenMaxLiquidity], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-liquidity:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		liquidity := TokenMaxLiquidity{
			TokenAddress:      getString(dataMap, "a"),
			PoolAddress:       getString(dataMap, "p"),
			LiquidityInUsd:    s.formatScientificNotation(dataMap["liu"]),
			LiquidityInNative: s.formatScientificNotation(dataMap["lin"]),
			Timestamp:         getInt64(dataMap, "ts"),
		}

		callback(liquidity)
	}, filter, "subscribeTokenMaxLiquidity")
}

// SubscribeTokenTotalLiquidity subscribes to token total liquidity data
// Pushes the total liquidity info of a token across all pools
// Channel: dex-token-total-liquidity:{chain}_{token_address}
func (s *StreamApi) SubscribeTokenTotalLiquidity(chain, tokenAddress string, callback StreamCallback[TokenTotalLiquidity], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-token-total-liquidity:%s_%s", chain, tokenAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		liquidity := TokenTotalLiquidity{
			TokenAddress:      getString(dataMap, "a"),
			LiquidityInUsd:    s.formatScientificNotation(dataMap["liu"]),
			LiquidityInNative: s.formatScientificNotation(dataMap["lin"]),
			PoolCount:         getInt(dataMap, "pc"),
			Timestamp:         getInt64(dataMap, "ts"),
		}

		callback(liquidity)
	}, filter, "subscribeTokenTotalLiquidity")
}

// SubscribeRankingTokensLiquidity subscribes to ranking tokens liquidity data
func (s *StreamApi) SubscribeRankingTokensLiquidity(chain string, channelType ChannelType, callback StreamCallback[[]TokenLiquidity]) Unsubscribe {
	channel := fmt.Sprintf("dex-ranking-token-general_stat_num-list:%s_%s", chain, channelType)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenLiquidity
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			liquidity := TokenLiquidity{
				TokenAddress: getString(itemMap, "a"),
				MetricType:   MetricType(getString(itemMap, "t")),
				Value:        s.formatScientificNotation(itemMap["v"]),
				Timestamp:    getInt64(itemMap, "ts"),
			}
			result = append(result, liquidity)
		}

		callback(result)
	}, "", "subscribeRankingTokensLiquidity")
}

// SubscribeRankingTokensStats subscribes to ranking tokens stats
func (s *StreamApi) SubscribeRankingTokensStats(chain string, channelType ChannelType, callback StreamCallback[[]TokenStat]) Unsubscribe {
	channel := fmt.Sprintf("dex-ranking-token-stats-list:%s_%s", chain, channelType)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenStat
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			stat := TokenStat{
				Address:   getString(itemMap, "a"),
				Timestamp: getInt64(itemMap, "t"),
			}

			// 1m data
			if val, ok := itemMap["b1m"]; ok {
				stat.Buys1m = PtrInt(int(val.(float64)))
			}
			if val, ok := itemMap["s1m"]; ok {
				stat.Sells1m = PtrInt(int(val.(float64)))
			}
			if val, ok := itemMap["bviu1m"]; ok {
				stat.BuyVolumeInUsd1m = PtrString(s.formatScientificNotation(val))
			}
			if val, ok := itemMap["sviu1m"]; ok {
				stat.SellVolumeInUsd1m = PtrString(s.formatScientificNotation(val))
			}
			if val, ok := itemMap["p"]; ok {
				stat.Price = PtrString(s.formatScientificNotation(val))
			}

			result = append(result, stat)
		}

		callback(result)
	}, "", "subscribeRankingTokensStats")
}

// SubscribeRankingTokensHolders subscribes to ranking tokens holders
func (s *StreamApi) SubscribeRankingTokensHolders(chain string, channelType ChannelType, callback StreamCallback[[]TokenHolder]) Unsubscribe {
	channel := fmt.Sprintf("dex-ranking-token-holding-list:%s_%s", chain, channelType)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenHolder
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			holder := TokenHolder{
				TokenAddress: getString(itemMap, "a"),
				Timestamp:    getInt64(itemMap, "ts"),
			}
			if val, ok := itemMap["h"]; ok {
				holder.Holders = PtrInt(int(val.(float64)))
			}
			if val, ok := itemMap["t100a"]; ok {
				holder.Top100Amount = PtrString(s.formatScientificNotation(val))
			}
			if val, ok := itemMap["t10a"]; ok {
				holder.Top10Amount = PtrString(s.formatScientificNotation(val))
			}

			result = append(result, holder)
		}

		callback(result)
	}, "", "subscribeRankingTokensHolders")
}

// SubscribeRankingTokensSupply subscribes to ranking tokens supply
func (s *StreamApi) SubscribeRankingTokensSupply(chain string, channelType ChannelType, callback StreamCallback[[]TokenSupply]) Unsubscribe {
	channel := fmt.Sprintf("dex-ranking-token-supply-list:%s_%s", chain, channelType)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenSupply
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			supply := TokenSupply{
				TokenAddress: getString(itemMap, "a"),
				Timestamp:    getInt64(itemMap, "ts"),
			}
			if val, ok := itemMap["s"]; ok {
				supply.Supply = PtrString(s.formatScientificNotation(val))
			}
			if val, ok := itemMap["mc"]; ok {
				supply.MarketCapInUsd = PtrString(s.formatScientificNotation(val))
			}

			result = append(result, supply)
		}

		callback(result)
	}, "", "subscribeRankingTokensSupply")
}

// SubscribeRankingTokensBondingCurve subscribes to ranking tokens bonding curve
func (s *StreamApi) SubscribeRankingTokensBondingCurve(chain string, callback StreamCallback[[]TokenBondingCurve]) Unsubscribe {
	channel := fmt.Sprintf("dex-ranking-token-bounding-curve-list:%s_new", chain)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []TokenBondingCurve
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			bc := TokenBondingCurve{}
			if val, ok := itemMap["a"]; ok {
				bc.TokenAddress = PtrString(val.(string))
			}
			if val, ok := itemMap["pr"]; ok {
				bc.ProgressRatio = PtrString(s.formatScientificNotation(val))
			}

			result = append(result, bc)
		}

		callback(result)
	}, "", "subscribeRankingTokensBondingCurve")
}

// SubscribeWalletPnlList subscribes to wallet PnL list data
func (s *StreamApi) SubscribeWalletPnlList(chain, walletAddress string, callback StreamCallback[[]WalletPnl]) Unsubscribe {
	channel := fmt.Sprintf("dex-wallet-pnl-list:%s_%s", chain, walletAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []WalletPnl
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			pnl := WalletPnl{
				WalletAddress:        getString(itemMap, "a"),
				Buys:                 getInt(itemMap, "bs"),
				BuyAmount:            s.formatScientificNotation(itemMap["ba"]),
				BuyAmountInUsd:       s.formatScientificNotation(itemMap["baiu"]),
				AverageBuyPriceInUsd: s.formatScientificNotation(itemMap["abpiu"]),
				SellAmount:           s.formatScientificNotation(itemMap["sa"]),
				SellAmountInUsd:      s.formatScientificNotation(itemMap["saiu"]),
				Sells:                getInt(itemMap, "ss"),
				Wins:                 getInt(itemMap, "ws"),
				WinRatio:             s.formatScientificNotation(itemMap["wr"]),
				PnlInUsd:             s.formatScientificNotation(itemMap["piu"]),
				AveragePnlInUsd:      s.formatScientificNotation(itemMap["apiu"]),
				PnlRatio:             s.formatScientificNotation(itemMap["pr"]),
				ProfitableDays:       getInt(itemMap, "pd"),
				LosingDays:           getInt(itemMap, "ld"),
				Tokens:               getInt(itemMap, "ts"),
				Resolution:           getString(itemMap, "r"),
			}

			result = append(result, pnl)
		}

		callback(result)
	}, "", "subscribeWalletPnlList")
}

// SubscribeWalletTrade subscribes to wallet trade data
func (s *StreamApi) SubscribeWalletTrade(chain, walletAddress string, callback StreamCallback[TradeActivity], filter string) Unsubscribe {
	channel := fmt.Sprintf("dex-wallet-trade:%s_%s", chain, walletAddress)
	return s.Subscribe(channel, func(data interface{}) {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return
		}

		trade := TradeActivity{
			TokenAddress:      getString(dataMap, "a"),
			Timestamp:         getInt64(dataMap, "t"),
			Kind:              getString(dataMap, "k"),
			BuyAmount:         s.formatScientificNotation(dataMap["ba"]),
			BuyAmountInUsd:    s.formatScientificNotation(dataMap["baiu"]),
			BuyTokenAddress:   getString(dataMap, "btma"),
			BuyTokenName:      getString(dataMap, "btn"),
			BuyTokenSymbol:    getString(dataMap, "bts"),
			BuyWalletAddress:  getString(dataMap, "bwa"),
			SellAmount:        s.formatScientificNotation(dataMap["sa"]),
			SellAmountInUsd:   s.formatScientificNotation(dataMap["saiu"]),
			SellTokenAddress:  getString(dataMap, "stma"),
			SellTokenName:     getString(dataMap, "stn"),
			SellTokenSymbol:   getString(dataMap, "sts"),
			SellWalletAddress: getString(dataMap, "swa"),
			TxHash:            getString(dataMap, "h"),
		}

		callback(trade)
	}, filter, "subscribeWalletTrade")
}

// SubscribeRankingTokensList subscribes to ranking tokens list
func (s *StreamApi) SubscribeRankingTokensList(chain string, rankingType RankingType, dex *Dex, callback StreamCallback[[]RankingTokenList]) Unsubscribe {
	var channel string
	if dex != nil {
		channel = fmt.Sprintf("dex-ranking-list:%s_%s_%s", chain, rankingType, *dex)
	} else {
		channel = fmt.Sprintf("dex-ranking-list:%s_%s", chain, rankingType)
	}

	return s.Subscribe(channel, func(data interface{}) {
		dataArr, ok := data.([]interface{})
		if !ok {
			return
		}

		var result []RankingTokenList
		for _, item := range dataArr {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			rankingItem := RankingTokenList{}

			// Parse metadata (t)
			if t, ok := itemMap["t"].(map[string]interface{}); ok {
				rankingItem.Metadata = &TokenMetadata{
					TokenAddress: getString(t, "a"),
				}
				if val, ok := t["n"]; ok {
					rankingItem.Metadata.Name = PtrString(val.(string))
				}
				if val, ok := t["s"]; ok {
					rankingItem.Metadata.Symbol = PtrString(val.(string))
				}
				if val, ok := t["iu"]; ok {
					rankingItem.Metadata.ImageUrl = PtrString(val.(string))
				}
				if val, ok := t["cts"]; ok {
					cts := int64(val.(float64))
					rankingItem.Metadata.CreatedAtMs = &cts
				}
			}

			// Parse bonding curve (bc)
			if bc, ok := itemMap["bc"].(map[string]interface{}); ok {
				rankingItem.BondingCurve = &TokenBondingCurve{}
				if val, ok := bc["pr"]; ok {
					rankingItem.BondingCurve.ProgressRatio = PtrString(s.formatScientificNotation(val))
				}
			}

			// Parse supply (s)
			if sup, ok := itemMap["s"].(map[string]interface{}); ok {
				rankingItem.Supply = &TokenSupply{
					TokenAddress: getString(sup, "a"),
				}
				if val, ok := sup["s"]; ok {
					rankingItem.Supply.Supply = PtrString(s.formatScientificNotation(val))
				}
				if val, ok := sup["mc"]; ok {
					rankingItem.Supply.MarketCapInUsd = PtrString(s.formatScientificNotation(val))
				}
			}

			// Parse stat (ts)
			if ts, ok := itemMap["ts"].(map[string]interface{}); ok {
				rankingItem.Stat = &TokenStat{
					Address: getString(ts, "a"),
				}
				if val, ok := ts["p"]; ok {
					rankingItem.Stat.Price = PtrString(s.formatScientificNotation(val))
				}
			}

			result = append(result, rankingItem)
		}

		callback(result)
	}, "", "subscribeRankingTokensList")
}

// Helper functions
func PtrInt(v int) *int {
	return &v
}

func PtrString(v string) *string {
	return &v
}

func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}

func getInt(m map[string]interface{}, key string) int {
	if val, ok := m[key]; ok {
		if f, ok := val.(float64); ok {
			return int(f)
		}
	}
	return 0
}

func getInt64(m map[string]interface{}, key string) int64 {
	if val, ok := m[key]; ok {
		if f, ok := val.(float64); ok {
			return int64(f)
		}
	}
	return 0
}
