/*
WebSocket Subscription Test

Usage:
1. Run: go test -v -run TestStreamWebSocket -timeout 60s ./test/
2. Press Ctrl+C to stop manually
*/

package chainstream_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	chainstream "github.com/chainstream-io/chainstream-go-sdk"
	"github.com/chainstream-io/chainstream-go-sdk/api"
)

func TestStreamWebSocket(t *testing.T) {
	// Token - replace with a valid token
	token := "YOUR_TOKEN_HERE"

	fmt.Println("Starting WebSocket subscription test...")

	// Create DexClient (WebSocket will auto-connect on subscribe)
	client, err := chainstream.NewDexClient(token, &chainstream.DexAggregatorOptions{
		StreamUrl: "wss://realtime-dex.chainstream.io/connection/websocket",
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	chain := "sol"
	tokenAddress := "So11111111111111111111111111111111111111112" // SOL

	fmt.Printf("Subscribing: %s/%s\n", chain, tokenAddress)
	fmt.Println("Listening... (Press Ctrl+C to stop)")

	// Subscribe to TokenCandles (auto-connects WebSocket)
	messageCount := 0
	unsubscribe := client.Stream.SubscribeTokenCandles(chain, tokenAddress, "1s", func(data api.TokenCandle) {
		messageCount++
		fmt.Printf("[%d] TokenCandle: open=%s, close=%s, high=%s, low=%s, volume=%s, time=%d\n",
			messageCount,
			data.Open,
			data.Close,
			data.High,
			data.Low,
			data.Volume,
			data.Time,
		)
	}, "")
	defer unsubscribe()

	// Set up signal handler for graceful exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal or timeout
	select {
	case <-sigChan:
		fmt.Println("\nReceived interrupt signal, closing connection...")
	case <-time.After(30 * time.Second):
		fmt.Println("\nTest timeout, closing connection...")
	}

	fmt.Printf("Total messages received: %d\n", messageCount)
}
