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

	"github.com/chainstream-io/chainstream-go-sdk/api"
)

func TestStreamWebSocket(t *testing.T) {
	// Token - replace with a valid token
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleXN0b3JlLUNIQU5HRS1NRSJ9.eyJqdGkiOiJhQjFweXFXR0hYRkVYLUlmNm5jd18iLCJzdWIiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlhdCI6MTc2ODg0MDY1NiwiZXhwIjoxNzY4OTI3MDU2LCJjbGllbnRfaWQiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlzcyI6Imh0dHBzOi8vZGV4LmFzaWEuYXV0aC5jaGFpbnN0cmVhbS5pby8iLCJhdWQiOiJodHRwczovL2FwaS5kZXguY2hhaW5zdHJlYW0uaW8ifQ.GNiv_oRPmhnWmi6w6bHGr3-zyw9fy8sf2kCnFBNKVBoQb5iw_DZIDH20NpghpotDx9s_ZGgfAN_wUsP6Fffg-M5xufSVym4961QcAfRRsJ_YyR_rxqYhLMrGrfCANIRv7qpzhxKB2OteDuSrLNfjBNQI_n-TnHJHOp_Xra3r-gTaCMEpxpl7nRTViILQUIzoLyrNvilBbKgFCaldmt6mtBZKSd9gSB5A5EOn8nDrV6LUo-0n9b7nhgq5zl6pXOXm9LYO3gjlXw5KNAR0K11g8_seNt6oMRPK1To7nr7Fm0G_6Hte0yzLftJCaCBe5HjNV0S6HBc1mSGSry1NAXFCig"

	fmt.Println("Starting WebSocket subscription test...")

	// Create DexClient (WebSocket will auto-connect on subscribe)
	client, err := openapi.NewDexClient(token, &openapi.DexAggregatorOptions{
		StreamUrl: "wss://realtime-dex.openapi.io/connection/websocket",
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
