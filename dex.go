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

	"github.com/chainstream-io/chainstream-go-sdk/api"
	"github.com/chainstream-io/chainstream-go-sdk/openapi"
)

const (
	LIB_VERSION = "0.1.17"
)

// TokenProvider represents a token provider interface
type TokenProvider interface {
	GetToken() (string, error)
}

// ChainStreamClientOptions represents ChainStream client options
type ChainStreamClientOptions struct {
	Debug     bool   `json:"debug,omitempty"`
	ServerUrl string `json:"serverUrl,omitempty"`
	StreamUrl string `json:"streamUrl,omitempty"`
	// AutoConnectWebSocket controls whether to automatically connect to WebSocket on initialization.
	// Default: false
	// If set to true, WebSocket will connect automatically.
	// If false or not set, connection will happen automatically when you use subscribe methods.
	AutoConnectWebSocket bool `json:"autoConnectWebSocket,omitempty"`
}

// ChainStreamClient represents the main ChainStream client
type ChainStreamClient struct {
	requestCtx    *api.DexRequestContext
	configuration *openapi.Configuration
	client        *openapi.APIClient
	stream        *api.StreamApi

	// API services
	Dex         *openapi.DexAPIService
	DexPool     *openapi.DexPoolAPIService
	Token       *openapi.TokenAPIService
	Wallet      *openapi.WalletAPIService
	Trade       *openapi.TradeAPIService
	Ranking     *openapi.RankingAPIService
	Transaction *openapi.TransactionAPIService
	Moonshot    *openapi.DefiSolanaMoonshotAPIService
	Pumpfun     *openapi.DefiSolanaPumpfunAPIService
	Stream      *api.StreamApi
	RedPacket   *openapi.RedPacketAPIService
	Ipfs        *openapi.IpfsAPIService
	Blockchain  *openapi.BlockchainAPIService
	Watchlist   *openapi.WatchlistAPIService
	KYT         *openapi.KYTAPIService
	Endpoint    *openapi.EndpointAPIService
	Jobs        *openapi.JobsAPIService
}

// NewChainStreamClient creates a new ChainStream client
// createChainStreamClient is a helper function to create a ChainStream client with common setup
func createChainStreamClient(accessToken string, tokenProvider TokenProvider, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	if options == nil {
		options = &ChainStreamClientOptions{}
	}

	// Set default values
	serverUrl := options.ServerUrl
	if serverUrl == "" {
		serverUrl = "https://api-dex.chainstream.io"
	}

	streamUrl := options.StreamUrl
	if streamUrl == "" {
		streamUrl = "wss://realtime-dex.chainstream.io/connection/websocket"
	}

	// Get token (either from string or provider)
	var token string
	var err error
	if tokenProvider != nil {
		token, err = tokenProvider.GetToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get initial token: %w", err)
		}
	} else {
		token = accessToken
	}

	// Create request context
	requestCtx := &api.DexRequestContext{
		BaseUrl:     serverUrl,
		StreamUrl:   streamUrl,
		AccessToken: token,
	}
	if tokenProvider != nil {
		requestCtx.TokenProvider = tokenProvider.GetToken
	}

	// Create configuration
	config := openapi.NewConfiguration()
	config.Servers = []openapi.ServerConfiguration{
		{
			URL:         serverUrl,
			Description: "Production server",
		},
	}
	config.UserAgent = fmt.Sprintf("dex/%s/go", LIB_VERSION)
	config.Debug = options.Debug

	// Set authentication
	config.DefaultHeader = map[string]string{
		"Authorization": "Bearer " + token,
		"User-Agent":    fmt.Sprintf("dex/%s/go", LIB_VERSION),
	}

	// Create API client
	client := openapi.NewAPIClient(config)

	// Create stream API
	stream := api.NewStreamApi(requestCtx)

	// Create ChainStream client
	csClient := &ChainStreamClient{
		requestCtx:    requestCtx,
		configuration: config,
		client:        client,
		stream:        stream,

		// Set API services
		Dex:         client.DexAPI,
		DexPool:     client.DexPoolAPI,
		Token:       client.TokenAPI,
		Wallet:      client.WalletAPI,
		Trade:       client.TradeAPI,
		Ranking:     client.RankingAPI,
		Transaction: client.TransactionAPI,
		Moonshot:    client.DefiSolanaMoonshotAPI,
		Pumpfun:     client.DefiSolanaPumpfunAPI,
		Stream:      stream,
		RedPacket:   client.RedPacketAPI,
		Ipfs:        client.IpfsAPI,
		Blockchain:  client.BlockchainAPI,
		Watchlist:   client.WatchlistAPI,
		KYT:         client.KYTAPI,
		Endpoint:    client.EndpointAPI,
		Jobs:        client.JobsAPI,
	}

	// Only auto-connect WebSocket if explicitly enabled
	if options.AutoConnectWebSocket {
		if err := stream.Connect(); err != nil {
			return nil, fmt.Errorf("failed to connect to stream service: %w", err)
		}
	}

	return csClient, nil
}

func NewChainStreamClient(accessToken string, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient(accessToken, nil, options)
}

// NewChainStreamClientWithTokenProvider creates a ChainStream client with token provider
func NewChainStreamClientWithTokenProvider(tokenProvider TokenProvider, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient("", tokenProvider, options)
}

// WaitForJob waits for job completion
func (c *ChainStreamClient) WaitForJob(jobId string, timeout time.Duration) (map[string]interface{}, error) {
	// Get access token
	var token string
	var err error
	if c.requestCtx.TokenProvider != nil {
		token, err = c.requestCtx.TokenProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	} else {
		token = c.requestCtx.AccessToken
	}

	// Build SSE URL
	sseUrl := fmt.Sprintf("%s/jobs/%s/streaming", c.requestCtx.BaseUrl, jobId)

	// Create HTTP request
	req, err := http.NewRequest("GET", sseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	// Send request
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
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
	var token string
	var err error
	if c.requestCtx.TokenProvider != nil {
		token, err = c.requestCtx.TokenProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	} else {
		token = c.requestCtx.AccessToken
	}

	// Build SSE URL
	sseUrl := fmt.Sprintf("%s/jobs/%s/streaming", c.requestCtx.BaseUrl, jobId)

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", sseUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
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
	if c.stream != nil {
		return c.stream.Disconnect()
	}
	return nil
}
