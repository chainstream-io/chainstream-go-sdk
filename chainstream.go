package chainstream

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/blockchain"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/dex"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/dexpool"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/ipfs"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/job"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/kyt"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/ranking"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/redpacket"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/token"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/trade"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/transaction"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/wallet"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/watchlist"
	"github.com/chainstream-io/chainstream-go-sdk/v2/openapi/webhook"
	"github.com/chainstream-io/chainstream-go-sdk/v2/stream"
)

// LIB_VERSION is the version of the ChainStream Go SDK
const LIB_VERSION = "2.0.9"

// DefaultServerURL is the default ChainStream API server URL.
const DefaultServerURL = "https://api.chainstream.io"

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
	// ApiKey for X-API-KEY header authentication.
	// When set, requests use X-API-KEY header instead of Authorization: Bearer.
	// Use this OR accessToken OR walletSigner, not multiple.
	ApiKey string
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
	Ipfs        *ipfs.ClientWithResponses
	Job         *job.ClientWithResponses
	Kyt         *kyt.ClientWithResponses
	Ranking     *ranking.ClientWithResponses
	RedPacket   *redpacket.ClientWithResponses
	Token       *token.ClientWithResponses
	Trade       *trade.ClientWithResponses
	Transaction *transaction.ClientWithResponses
	Wallet      *wallet.ClientWithResponses
	Watchlist   *watchlist.ClientWithResponses
	Webhook     *webhook.ClientWithResponses

	// WebSocket streaming API
	Stream *stream.StreamApi
}

// NewChainStreamClient creates a new ChainStream client with all API modules.
func NewChainStreamClient(accessToken string, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient(accessToken, nil, nil, options)
}

// NewChainStreamClientWithTokenProvider creates a ChainStream client with token provider
func NewChainStreamClientWithTokenProvider(tokenProvider TokenProvider, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient("", tokenProvider, nil, options)
}

// NewChainStreamClientWithApiKey creates a ChainStream client using X-API-KEY header authentication.
func NewChainStreamClientWithApiKey(apiKey string, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	if options == nil {
		options = &ChainStreamClientOptions{}
	}
	options.ApiKey = apiKey
	return createChainStreamClient("", nil, nil, options)
}

// NewChainStreamClientWithWalletSigner creates a ChainStream client using x402 wallet signature auth.
// Each API request is signed with the wallet's private key instead of using a Bearer token.
func NewChainStreamClientWithWalletSigner(signer WalletSigner, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
	return createChainStreamClient("", nil, signer, options)
}

// createChainStreamClient is a helper function to create a ChainStream client with common setup
func createChainStreamClient(accessToken string, tokenProvider TokenProvider, walletSigner WalletSigner, options *ChainStreamClientOptions) (*ChainStreamClient, error) {
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
		blockchain.WithRequestEditorFn(authEditorFn[blockchain.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		blockchain.WithRequestEditorFn(userAgentFn[blockchain.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain client: %w", err)
	}

	// Create dex client
	client.Dex, err = dex.NewClientWithResponses(serverURL,
		dex.WithRequestEditorFn(authEditorFn[dex.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		dex.WithRequestEditorFn(userAgentFn[dex.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create dex client: %w", err)
	}

	// Create dexpool client
	client.DexPool, err = dexpool.NewClientWithResponses(serverURL,
		dexpool.WithRequestEditorFn(authEditorFn[dexpool.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		dexpool.WithRequestEditorFn(userAgentFn[dexpool.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create dexpool client: %w", err)
	}

	// Create ipfs client
	client.Ipfs, err = ipfs.NewClientWithResponses(serverURL,
		ipfs.WithRequestEditorFn(authEditorFn[ipfs.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		ipfs.WithRequestEditorFn(userAgentFn[ipfs.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ipfs client: %w", err)
	}

	// Create job client
	client.Job, err = job.NewClientWithResponses(serverURL,
		job.WithRequestEditorFn(authEditorFn[job.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		job.WithRequestEditorFn(userAgentFn[job.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create job client: %w", err)
	}

	// Create kyt client
	client.Kyt, err = kyt.NewClientWithResponses(serverURL,
		kyt.WithRequestEditorFn(authEditorFn[kyt.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		kyt.WithRequestEditorFn(userAgentFn[kyt.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create kyt client: %w", err)
	}

	// Create ranking client
	client.Ranking, err = ranking.NewClientWithResponses(serverURL,
		ranking.WithRequestEditorFn(authEditorFn[ranking.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		ranking.WithRequestEditorFn(userAgentFn[ranking.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ranking client: %w", err)
	}

	// Create redpacket client
	client.RedPacket, err = redpacket.NewClientWithResponses(serverURL,
		redpacket.WithRequestEditorFn(authEditorFn[redpacket.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		redpacket.WithRequestEditorFn(userAgentFn[redpacket.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create redpacket client: %w", err)
	}

	// Create token client
	client.Token, err = token.NewClientWithResponses(serverURL,
		token.WithRequestEditorFn(authEditorFn[token.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		token.WithRequestEditorFn(userAgentFn[token.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Create trade client
	client.Trade, err = trade.NewClientWithResponses(serverURL,
		trade.WithRequestEditorFn(authEditorFn[trade.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		trade.WithRequestEditorFn(userAgentFn[trade.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trade client: %w", err)
	}

	// Create transaction client
	client.Transaction, err = transaction.NewClientWithResponses(serverURL,
		transaction.WithRequestEditorFn(authEditorFn[transaction.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		transaction.WithRequestEditorFn(userAgentFn[transaction.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction client: %w", err)
	}

	// Create wallet client
	client.Wallet, err = wallet.NewClientWithResponses(serverURL,
		wallet.WithRequestEditorFn(authEditorFn[wallet.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		wallet.WithRequestEditorFn(userAgentFn[wallet.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet client: %w", err)
	}

	// Create watchlist client
	client.Watchlist, err = watchlist.NewClientWithResponses(serverURL,
		watchlist.WithRequestEditorFn(authEditorFn[watchlist.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		watchlist.WithRequestEditorFn(userAgentFn[watchlist.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create watchlist client: %w", err)
	}

	// Create webhook client
	client.Webhook, err = webhook.NewClientWithResponses(serverURL,
		webhook.WithRequestEditorFn(authEditorFn[webhook.RequestEditorFn](authToken, walletSigner, options.ApiKey)),
		webhook.WithRequestEditorFn(userAgentFn[webhook.RequestEditorFn]()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook client: %w", err)
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

// authEditorFn returns a request editor that uses wallet signer, API key, or bearer token.
// Priority: walletSigner > apiKey > accessToken.
func authEditorFn[T ~func(context.Context, *http.Request) error](accessToken string, walletSigner WalletSigner, apiKey string) T {
	if walletSigner != nil {
		return WalletAuthHeaderFn[T](walletSigner)
	}
	if apiKey != "" {
		return apiKeyHeaderFn[T](apiKey)
	}
	return bearerHeaderFn[T](accessToken)
}

// apiKeyHeaderFn returns a request editor function that adds the X-API-KEY header.
func apiKeyHeaderFn[T ~func(context.Context, *http.Request) error](apiKey string) T {
	return T(func(_ context.Context, req *http.Request) error {
		req.Header.Set("X-API-KEY", apiKey)
		req.Header.Set("Content-Type", "application/json")
		return nil
	})
}

// bearerHeaderFn returns a request editor function that adds the Authorization: Bearer header.
func bearerHeaderFn[T ~func(context.Context, *http.Request) error](accessToken string) T {
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
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := c.WaitForJobWithContext(ctx, jobId)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("job timed out after %v", timeout)
		}
		return nil, err
	}
	return result, nil
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
	sseUrl := fmt.Sprintf("%s/v2/job/%s/streaming", c.requestCtx.BaseUrl, jobId)

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

				status, _ := data["status"].(string)
				switch status {
				case "pending":
					continue
				case "error":
					return nil, fmt.Errorf("job error: %s", extractErrorMessage(data))
				case "failed":
					return nil, fmt.Errorf("job failed: %s", extractErrorMessage(data))
				case "completed":
					if success, ok := data["success"].(bool); ok && !success {
						return nil, fmt.Errorf("transaction failed: %s", extractErrorMessage(data))
					}
					return data, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("job context cancelled")
}

func extractErrorMessage(data map[string]interface{}) string {
	if msg, ok := data["error"].(string); ok && msg != "" {
		return msg
	}
	if msg, ok := data["message"].(string); ok && msg != "" {
		return msg
	}
	return "unknown error"
}

// Close closes the client connection
func (c *ChainStreamClient) Close() error {
	if c.Stream != nil {
		return c.Stream.Disconnect()
	}
	return nil
}
