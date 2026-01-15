package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chainstream-io/centrifuge-go"
)

const (
	LIB_VERSION = "0.0.10"
)

// DexRequestContext represents the request context for WebSocket connections
type DexRequestContext struct {
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
	requestCtx   *DexRequestContext
}

// StreamCallback represents a callback function for stream data
type StreamCallback[T any] func(data T)

// Unsubscribe represents an unsubscribe function
type Unsubscribe func()

// NewStreamApi creates a new StreamApi instance
func NewStreamApi(requestCtx *DexRequestContext) *StreamApi {
	ctx, cancel := context.WithCancel(context.Background())
	return &StreamApi{
		listenersMap: make(map[string][]StreamCallback[interface{}]),
		ctx:          ctx,
		cancel:       cancel,
		requestCtx:   requestCtx,
	}
}

// Connect establishes WebSocket connection to the server using Centrifuge
func (s *StreamApi) Connect() error {
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

	// Create HTTP headers with Authorization for WebSocket handshake
	headers := make(http.Header)
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	headers.Set("User-Agent", fmt.Sprintf("chainstream-io/sdk/go/%s", LIB_VERSION))

	// Create Centrifuge client configuration
	config := centrifuge.Config{
		Token:  token,
		Header: headers,
		GetToken: func(ctx centrifuge.ConnectionTokenEvent) (string, error) {
			var newToken string
			var err error
			if s.requestCtx.TokenProvider != nil {
				newToken, err = s.requestCtx.TokenProvider()
			} else {
				newToken = s.requestCtx.AccessToken
				err = nil
			}
			if err != nil {
				return "", err
			}
			// Update HTTP headers with new token for reconnection
			log.Println("[streaming] getToken called, updating headers...")
			newHeaders := make(http.Header)
			newHeaders.Set("Authorization", fmt.Sprintf("Bearer %s", newToken))
			newHeaders.Set("User-Agent", fmt.Sprintf("chainstream-io/sdk/go/%s", LIB_VERSION))
			s.client.SetHttpHeaders(newHeaders)
			return newToken, nil
		},
		ReadTimeout:      30 * time.Second,
		WriteTimeout:     30 * time.Second,
		HandshakeTimeout: 30 * time.Second,
	}

	// Create Centrifuge client
	client := centrifuge.NewJsonClient(s.requestCtx.StreamUrl, config)

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

	return nil
}

// Disconnect closes the WebSocket connection
func (s *StreamApi) Disconnect() error {
	s.cancel()
	if s.client != nil {
		s.client.Close()
	}
	return nil
}

// Subscribe subscribes to a channel using Centrifuge protocol
func (s *StreamApi) Subscribe(channel string, callback StreamCallback[interface{}], filter, methodName string) Unsubscribe {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Process filter expression
	processedFilter := filter
	if filter != "" && methodName != "" {
		processedFilter = ReplaceFilterFields(filter, methodName)
		fmt.Printf("🔍 Original filter: %s\n", filter)
		fmt.Printf("🔍 Processed filter: %s\n", processedFilter)
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
		// Parse JSON data from bytes
		var jsonData interface{}
		if err := json.Unmarshal(e.Data, &jsonData); err != nil {
			log.Printf("[streaming] error parsing JSON: %v", err)
			return
		}

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
	s.mutex.RUnlock()

	if !exists {
		return
	}

	for _, listener := range listeners {
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
func (s *StreamApi) SubscribeTokenCandles(chain, tokenAddress string, resolution string, callback StreamCallback[TokenCandle], filter string) Unsubscribe {
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
			Resolution: dataMap["r"].(string),
			Time:       int64(dataMap["t"].(float64)),
			Number:     int(dataMap["n"].(float64)),
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
			Address:   dataMap["a"].(string),
			Timestamp: int64(dataMap["t"].(float64)),
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
			TokenAddress: dataMap["a"].(string),
			Timestamp:    int64(dataMap["ts"].(float64)),
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
			TokenAddress: dataMap["a"].(string),
			Name:         dataMap["n"].(string),
			Symbol:       dataMap["s"].(string),
			CreatedAtMs:  int64(dataMap["cts"].(float64)),
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
			WalletAddress:   dataMap["a"].(string),
			TokenAddress:    dataMap["ta"].(string),
			TokenPriceInUsd: s.formatScientificNotation(dataMap["tpiu"]),
			Balance:         s.formatScientificNotation(dataMap["b"]),
			Timestamp:       int64(dataMap["t"].(float64)),
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
			WalletAddress:            dataMap["a"].(string),
			TokenAddress:             dataMap["ta"].(string),
			TokenPriceInUsd:          s.formatScientificNotation(dataMap["tpiu"]),
			Timestamp:                int64(dataMap["t"].(float64)),
			OpenTime:                 int64(dataMap["ot"].(float64)),
			LastTime:                 int64(dataMap["lt"].(float64)),
			CloseTime:                int64(dataMap["ct"].(float64)),
			BuyAmount:                s.formatScientificNotation(dataMap["ba"]),
			BuyAmountInUsd:           s.formatScientificNotation(dataMap["baiu"]),
			BuyCount:                 int(dataMap["bs"].(float64)),
			BuyCount30d:              int(dataMap["bs30d"].(float64)),
			BuyCount7d:               int(dataMap["bs7d"].(float64)),
			SellAmount:               s.formatScientificNotation(dataMap["sa"]),
			SellAmountInUsd:          s.formatScientificNotation(dataMap["saiu"]),
			SellCount:                int(dataMap["ss"].(float64)),
			SellCount30d:             int(dataMap["ss30d"].(float64)),
			SellCount7d:              int(dataMap["ss7d"].(float64)),
			HeldDurationTimestamp:    int64(dataMap["hdts"].(float64)),
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
			TokenAddress:      dataMap["a"].(string),
			Timestamp:         int64(dataMap["t"].(float64)),
			Kind:              dataMap["k"].(string),
			BuyAmount:         s.formatScientificNotation(dataMap["ba"]),
			BuyAmountInUsd:    s.formatScientificNotation(dataMap["baiu"]),
			BuyTokenAddress:   dataMap["btma"].(string),
			BuyTokenName:      dataMap["btn"].(string),
			BuyTokenSymbol:    dataMap["bts"].(string),
			BuyWalletAddress:  dataMap["bwa"].(string),
			SellAmount:        s.formatScientificNotation(dataMap["sa"]),
			SellAmountInUsd:   s.formatScientificNotation(dataMap["saiu"]),
			SellTokenAddress:  dataMap["stma"].(string),
			SellTokenName:     dataMap["stn"].(string),
			SellTokenSymbol:   dataMap["sts"].(string),
			SellWalletAddress: dataMap["swa"].(string),
			TxHash:            dataMap["h"].(string),
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
			PoolAddress:          dataMap["a"].(string),
			TokenAAddress:        dataMap["taa"].(string),
			TokenALiquidityInUsd: s.formatScientificNotation(dataMap["taliu"]),
			TokenBAddress:        dataMap["tba"].(string),
			TokenBLiquidityInUsd: s.formatScientificNotation(dataMap["tbliu"]),
		}

		callback(balance)
	}, "", "subscribeDexPoolBalance")
}

// Helper functions
func PtrInt(v int) *int {
	return &v
}

func PtrString(v string) *string {
	return &v
}
