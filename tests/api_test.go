// API 测试
//
// 使用方法：
// 运行: go test -v ./tests -run TestGetToken

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	chainstream "github.com/chainstream-io/chainstream-go-sdk"
	"github.com/chainstream-io/chainstream-go-sdk/openapi/token"
)

const ACCESS_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImtleXN0b3JlLUNIQU5HRS1NRSJ9.eyJqdGkiOiJJQWxZMGdNRGJ0ZW5jNnNhT1dheDEiLCJzdWIiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlhdCI6MTc3MDAxMzc0MywiZXhwIjoxNzcwMTAwMTQzLCJjbGllbnRfaWQiOiJIR2hWbmpiSWlheDFIcDNUakdUd083WU9FUkJURXRwaSIsImlzcyI6Imh0dHBzOi8vZGV4LmFzaWEuYXV0aC5jaGFpbnN0cmVhbS5pby8iLCJhdWQiOiJodHRwczovL2FwaS5kZXguY2hhaW5zdHJlYW0uaW8ifQ.TVY_FN1MdMogamLvXnYlVoLSXTZpX1b1c3xtUJNB5peUrZCTY_nLB8oOJ-ysBz3qsZhnUpUnX4LKSfuyGXDlfyasJG5c7yrj5zEYUZfkFKxG7PTtNLzXTF-4z0J7VnkbA-VUz1c1z3gaGDf3TFpo_Mfl6Zqf0v_CgDugciJ0ZbJS68gy_EaMXSBHwU7mm_vC2FUFkUfa8qwL3xvEEbOYxIQXJyHmqddnJew6nSyifHEqC_tSYj-o8GDP6PTbfOwqiyB_-T8valMwygFXYMgGstOUnITgNvJE4ciya3yalLuWONoA3LtoDzAq3wABr-0cyIbHVvQJkDT0VjqvJOv89Q"

func TestGetToken(t *testing.T) {
	fmt.Println("开始测试 GetToken API...")

	// 初始化 SDK 客户端
	client, err := chainstream.NewChainStreamClient(ACCESS_TOKEN, &chainstream.ChainStreamClientOptions{})
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	chain := token.Sol
	tokenAddress := "So11111111111111111111111111111111111111112" // SOL

	fmt.Printf("查询: %s/%s\n", chain, tokenAddress)

	// 通过 SDK 调用 API
	resp, err := client.Token.GetToken(context.Background(), chain, tokenAddress)
	if err != nil {
		t.Fatalf("API 调用失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("解析 JSON 失败: %v", err)
	}

	fmt.Println("\n返回结果:")
	prettyJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(prettyJSON))
}
