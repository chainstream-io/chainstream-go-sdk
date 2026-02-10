package chainstream

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/chainstream-io/chainstream-go-sdk/stream"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/blockchain"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/defi_moonshot"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/defi_pumpfun"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/dex"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/dexpool"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/endpoint"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/ipfs"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/jobs"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/kyt"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/ranking"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/redpacket"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/token"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/trade"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/transaction"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/wallet"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/watchlist"
)

// LIB_VERSION is the version of the ChainStream Go SDK
const LIB_VERSION = "0.1.6"

// DefaultServerURL is the default ChainStream API server URL.
const DefaultServerURL = "https://api-dex.chainstream.io"

// DefaultStreamURL is the default ChainStream WebSocket URL.
const DefaultStreamURL = "wss://realtime-dex.chainstream.io/connection/websocket"

// TokenProvider represents a token provider interface
type TokenProvider interface {
	GetToken() (string, error)
}

// ChainStreamClientOptions contains configuration options for the ChainStream client.
type ChainStreamClientOptions struct {
	// ServerURL is the base URL of the ChainStream API server.
	ServerURL string
	// StreamURL is the WebSocket URL for streaming data.
	StreamURL string
	// Debugging enables debug logging when true.
	Debugging bool
	// AutoConnectWebSocket controls whether to automatically connect to WebSocket on initialization.
	// Default: false
	// If set to true, WebSocket will connect automatically.
	// If false or not set, connection will happen automatically when you use subscribe methods.
	AutoConnectWebSocket bool
}

// ChainStreamClient is the main ChainStream client that provides access to all API modules.
type ChainStreamClient struct {
	requestCtx *stream.RequestContext

	// API module clients
	Blockchain  *blockchain.ClientWithResponses
	Dex         *dex.ClientWithResponses
	DexPool     *dexpool.ClientWithResponses
	Endpoint    *endpoint.ClientWithResponses
	Ipfs        *ipfs.ClientWithResponses
	Jobs        *jobs.ClientWithResponses
	Kyt         *kyt.ClientWithResponses
	Ranking     *ranking.ClientWithResponses
	RedPacket   *redpacket.ClientWithResponses
	Token       *token.ClientWithResponses
	Trade       *trade.ClientWithResponses
	Transaction *transaction.ClientWithResponses
	Wallet      *wallet.ClientWithResponses
	Watchlist   *watchlist.ClientWithResponses
	Moonshot    *defi_moonshot.ClientWithResponses
	Pumpfun     *defi_pumpfun.ClientWithResponses

	// WebSocket streaming API
	Stream *stream.StreamApi
}

// NewChainStreamClient creates a new ChainStream client with all API modules.
func NewChainStreamClient(accessToken string, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient(accessToken, nil, options)
}

// NewChainStreamClientWithTokenProvider creates a ChainStream client with token provider
func NewChainStreamClientWithTokenProvider(tokenProvider TokenProvider, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient("", tokenProvider, options)
}

// createChainStreamClient is a helper function to create a ChainStream client with common setup
func createChainStreamClient(accessToken string, tokenProvider TokenProvider, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	if options == nil {
		options = &ChainStreamClientOptions{}
	}

	serverURL := options.ServerURL
	if serverURL == "" {
		serverURL = DefaultServerURL
	}

	streamURL := options.StreamURL
	if streamURL == "" {
		streamURL = DefaultStreamURL
	}

	// Get access token (either from string or provider)
	var authToken string
	var err error
	if tokenProvider != nil {
		authToken, err = tokenProvider.GetToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get initial token: %w", err)
		}
	} else {
		authToken = accessToken
	}

	// Create request context
	requestCtx := &stream.RequestContext{
		BaseUrl:     serverURL,
		StreamUrl:   streamURL,
		AccessToken: authToken,
	}
	if tokenProvider != nil {
		requestCtx.TokenProvider = tokenProvider.GetToken
	}

	client := &ChainStreamClient{requestCtx: requestCtx}

	// Create blockchain client
	client.Blockchain, err = blockchain.NewClientWithResponses(serverURL,
		blockchain.WithRequestEditorFn(authHeaderFn[blockchain.RequestEditorFn](authToken)),
		blockchain.WithRequestEditorFn(userAgentFn[blockchain.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain client: %w", err)
	}

	// Create dex client
	client.Dex, err = dex.NewClientWithResponses(serverURL,
		dex.WithRequestEditorFn(authHeaderFn[dex.RequestEditorFn](authToken)),
		dex.WithRequestEditorFn(userAgentFn[dex.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create dex client: %w", err)
	}

	// Create dexpool client
	client.DexPool, err = dexpool.NewClientWithResponses(serverURL,
		dexpool.WithRequestEditorFn(authHeaderFn[dexpool.RequestEditorFn](authToken)),
		dexpool.WithRequestEditorFn(userAgentFn[dexpool.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create dexpool client: %w", err)
	}

	// Create endpoint client
	client.Endpoint, err = endpoint.NewClientWithResponses(serverURL,
		endpoint.WithRequestEditorFn(authHeaderFn[endpoint.RequestEditorFn](authToken)),
		endpoint.WithRequestEditorFn(userAgentFn[endpoint.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint client: %w", err)
	}

	// Create ipfs client
	client.Ipfs, err = ipfs.NewClientWithResponses(serverURL,
		ipfs.WithRequestEditorFn(authHeaderFn[ipfs.RequestEditorFn](authToken)),
		ipfs.WithRequestEditorFn(userAgentFn[ipfs.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ipfs client: %w", err)
	}

	// Create jobs client
	client.Jobs, err = jobs.NewClientWithResponses(serverURL,
		jobs.WithRequestEditorFn(authHeaderFn[jobs.RequestEditorFn](authToken)),
		jobs.WithRequestEditorFn(userAgentFn[jobs.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create jobs client: %w", err)
	}

	// Create kyt client
	client.Kyt, err = kyt.NewClientWithResponses(serverURL,
		kyt.WithRequestEditorFn(authHeaderFn[kyt.RequestEditorFn](authToken)),
		kyt.WithRequestEditorFn(userAgentFn[kyt.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create kyt client: %w", err)
	}

	// Create ranking client
	client.Ranking, err = ranking.NewClientWithResponses(serverURL,
		ranking.WithRequestEditorFn(authHeaderFn[ranking.RequestEditorFn](authToken)),
		ranking.WithRequestEditorFn(userAgentFn[ranking.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ranking client: %w", err)
	}

	// Create redpacket client
	client.RedPacket, err = redpacket.NewClientWithResponses(serverURL,
		redpacket.WithRequestEditorFn(authHeaderFn[redpacket.RequestEditorFn](authToken)),
		redpacket.WithRequestEditorFn(userAgentFn[redpacket.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create redpacket client: %w", err)
	}

	// Create token client
	client.Token, err = token.NewClientWithResponses(serverURL,
		token.WithRequestEditorFn(authHeaderFn[token.RequestEditorFn](authToken)),
		token.WithRequestEditorFn(userAgentFn[token.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Create trade client
	client.Trade, err = trade.NewClientWithResponses(serverURL,
		trade.WithRequestEditorFn(authHeaderFn[trade.RequestEditorFn](authToken)),
		trade.WithRequestEditorFn(userAgentFn[trade.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trade client: %w", err)
	}

	// Create transaction client
	client.Transaction, err = transaction.NewClientWithResponses(serverURL,
		transaction.WithRequestEditorFn(authHeaderFn[transaction.RequestEditorFn](authToken)),
		transaction.WithRequestEditorFn(userAgentFn[transaction.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction client: %w", err)
	}

	// Create wallet client
	client.Wallet, err = wallet.NewClientWithResponses(serverURL,
		wallet.WithRequestEditorFn(authHeaderFn[wallet.RequestEditorFn](authToken)),
		wallet.WithRequestEditorFn(userAgentFn[wallet.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet client: %w", err)
	}

	// Create watchlist client
	client.Watchlist, err = watchlist.NewClientWithResponses(serverURL,
		watchlist.WithRequestEditorFn(authHeaderFn[watchlist.RequestEditorFn](authToken)),
		watchlist.WithRequestEditorFn(userAgentFn[watchlist.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create watchlist client: %w", err)
	}

	// Create moonshot client
	client.Moonshot, err = defi_moonshot.NewClientWithResponses(serverURL,
		defi_moonshot.WithRequestEditorFn(authHeaderFn[defi_moonshot.RequestEditorFn](authToken)),
		defi_moonshot.WithRequestEditorFn(userAgentFn[defi_moonshot.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create moonshot client: %w", err)
	}

	// Create pumpfun client
	client.Pumpfun, err = defi_pumpfun.NewClientWithResponses(serverURL,
		defi_pumpfun.WithRequestEditorFn(authHeaderFn[defi_pumpfun.RequestEditorFn](authToken)),
		defi_pumpfun.WithRequestEditorFn(userAgentFn[defi_pumpfun.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create pumpfun client: %w", err)
	}

	// Create stream API
	client.Stream = stream.NewStreamApi(requestCtx)

	// Only auto-connect WebSocket if explicitly enabled
	if options.AutoConnectWebSocket {
		if err := client.Stream.Connect(); err != nil {
			return nil, fmt.Errorf("failed to connect to stream service: %w", err)
		}
	}

	return client, nil
}

// authHeaderFn returns a request editor function that adds the Authorization header.
func authHeaderFn[T ~func(context.Context, *http.Request) error](accessToken string) T {
	return T(func(_ context.Context, req *http.Request) error {
		if accessToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}
		req.Header.Set("Content-Type", "application/json")
		return nil
	})
}

// userAgentFn returns a request editor function that adds the User-Agent header.
func userAgentFn[T ~func(context.Context, *http.Request) error]() T {
	return T(func(_ context.Context, req *http.Request) error {
		req.Header.Set("User-Agent", fmt.Sprintf("chainstream-sdk/%s/go", LIB_VERSION))
		return nil
	})
}

// WaitForJob waits for job completion
func (c *ChainStreamClient) WaitForJob(jobId string, timeout time.Duration) (map[string]interface{}, error) {
	// Get access token
	var authToken string
	var err error
	if c.requestCtx.TokenProvider != nil {
		authToken, err = c.requestCtx.TokenProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	} else {
		authToken = c.requestCtx.AccessToken
	}

	// Build SSE URL
	sseUrl := fmt.Sprintf("%s/jobs/%s/streaming", c.requestCtx.BaseUrl, jobId)

	// Create HTTP request
	req, err := http.NewRequest("GET", sseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	// Send request
	httpClient := &http.Client{Timeout: timeout}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read SSE stream
	decoder := json.NewDecoder(resp.Body)
	for {
		var data map[string]interface{}
		if err := decoder.Decode(&data); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		if status, ok := data["status"].(string); ok {
			if status == "error" {
				message := "unknown error"
				if msg, ok := data["message"].(string); ok {
					message = msg
				}
				return nil, fmt.Errorf("job error: %s", message)
			} else if status == "completed" {
				return data, nil
			}
		}
	}

	return nil, fmt.Errorf("job timed out after %v", timeout)
}

// WaitForJobWithContext waits for job completion with context
func (c *ChainStreamClient) WaitForJobWithContext(ctx context.Context, jobId string) (map[string]interface{}, error) {
	// Get access token
	var authToken string
	var err error
	if c.requestCtx.TokenProvider != nil {
		authToken, err = c.requestCtx.TokenProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	} else {
		authToken = c.requestCtx.AccessToken
	}

	// Build SSE URL
	sseUrl := fmt.Sprintf("%s/jobs/%s/streaming", c.requestCtx.BaseUrl, jobId)

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", sseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	// Send request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read SSE stream
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			line := scanner.Text()

			if strings.HasPrefix(line, "data: ") {
				jsonData := strings.TrimPrefix(line, "data: ")
				if jsonData == "" {
					continue
				}

				var data map[string]interface{}
				if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
					return nil, fmt.Errorf("failed to decode response: %w", err)
				}

				if status, ok := data["status"].(string); ok {
					if status == "error" {
						message := "unknown error"
						if msg, ok := data["message"].(string); ok {
							message = msg
						}
						return nil, fmt.Errorf("job error: %s", message)
					} else if status == "completed" {
						return data, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("job context cancelled")
}

// Close closes the client connection
func (c *ChainStreamClient) Close() error {
	if c.Stream != nil {
		return c.Stream.Disconnect()
	}
	return nil
}
