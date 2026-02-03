// WebSocket 订阅测试
//
// 使用方法：
// 1. 运行: go test -v ./tests -run TestWebSocket -timeout 0
// 2. 按 Ctrl+C 手动停止

package tests

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	chainstream "github.com/chainstream-io/chainstream-go-sdk"
	"github.com/chainstream-io/chainstream-go-sdk/stream"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/token"
)

const WS_ACCESS_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleXN0b3JlLUNIQU5HRS1NRSJ9.eyJqdGkiOiJJQWxZMGdNRGJ0ZW5jNnNhT1dheDEiLCJzdWIiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlhdCI6MTc3MDAxMzc0MywiZXhwIjoxNzcwMTAwMTQzLCJjbGllbnRfaWQiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlzcyI6Imh0dHBzOi8vZGV4LmFzaWEuYXV0aC5jaGFpbnN0cmVhbS5pby8iLCJhdWQiOiJodHRwczovL2FwaS5kZXguY2hhaW5zdHJlYW0uaW8ifQ.TVY_FN1MdMogamLvXnYlVoLSXTZpX1b1c3xtUJNB5peUrZCTY_nLB8oOJ-ysBz3qsZhnUpUnX4LKSfuyGXDlfyasJG5c7yrj5zEYUZfkFKxG7PTtNLzXTF-4z0J7VnkbA-VUz1c1z3gaGDf3TFpo_Mfl6Zqf0v_CgDugciJ0ZbJS68gy_EaMXSBHwU7mm_vC2FUFkUfa8qwL3xvEEbOYxIQXJyHmqddnJew6nSyifHEqC_tSYj-o8GDP6PTbfOwqiyB_-T8valMwygFXYMgGstOUnITgNvJE4ciya3yalLuWONoA3LtoDzAq3wABr-0cyIbHVvQJkDT0VjqvJOv89Q"

func TestWebSocket(t *testing.T) {
	fmt.Println("开始测试 WebSocket 连接...")

	// 创建 ChainStreamClient (WebSocket 将在订阅时自动连接)
	client, err := chainstream.NewChainStreamClient(WS_ACCESS_TOKEN, &chainstream.ChainStreamClientOptions{
		StreamURL: "wss://realtime-dex.chainstream.io/connection/websocket",
	})
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	chain := "sol"
	tokenAddress := "So11111111111111111111111111111111111111112" // SOL

	fmt.Printf("订阅 Token K线: %s/%s\n", chain, tokenAddress)
	fmt.Println("持续监听中... (按 Ctrl+C 停止)")

	// 订阅 Token K线 (自动连接 WebSocket)
	messageCount := 0
	unsubscribe := client.Stream.SubscribeTokenCandles(chain, tokenAddress, token.N1s, func(data stream.TokenCandle) {
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

	// 处理中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 一直等待直到收到中断信号
	<-interrupt
	fmt.Println("\n收到中断信号，关闭连接...")
	fmt.Printf("总共收到消息: %d\n", messageCount)
}
