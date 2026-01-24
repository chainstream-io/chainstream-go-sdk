# \TradeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivities**](TradeAPI.md#GetActivities) | **Get** /v1/trade/{chain}/activities | CONTROLLER.TRADE.GET.TOKEN_ACTIVITIES.SUMMARY
[**GetTrades**](TradeAPI.md#GetTrades) | **Get** /v1/trade/{chain} | CONTROLLER.TRADE.GET.TOKEN.SUMMARY



## GetActivities

> TradePage GetActivities(ctx, chain).Cursor(cursor).Limit(limit).Direction(direction).TokenAddress(tokenAddress).WalletAddress(walletAddress).PoolAddress(poolAddress).BeforeTimestamp(beforeTimestamp).AfterTimestamp(afterTimestamp).BeforeBlockHeight(beforeBlockHeight).AfterBlockHeight(afterBlockHeight).TransactionsSignature(transactionsSignature).Type_(type_).Execute()

CONTROLLER.TRADE.GET.TOKEN_ACTIVITIES.SUMMARY



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chainstream-io/chainstream-go-sdk"
)

func main() {
	chain := openapiclient.ChainSymbol("sol") // ChainSymbol | GLOBAL.CHAIN.DESCRIPTION
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR.DESCRIPTION (optional)
	limit := float32(8.14) // float32 | DTO.PAGE.LIMIT (optional) (default to 20)
	direction := "direction_example" // string | DTO.PAGE.DIRECTION (optional) (default to "next")
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.TRADE.QUERY.TOKEN_ADDRESS (optional)
	walletAddress := "3xd4LGVWtYXLBspR6X5JWbW49NXmEehfPtX6Kqx98b4w" // string | DTO.TRADE.QUERY.WALLET_ADDRESS (optional)
	poolAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.TRADE.QUERY.POOL_ADDRESS (optional)
	beforeTimestamp := int64(1717334400000) // int64 | DTO.TRADE.QUERY.BEFORE_TIMESTAMP (optional)
	afterTimestamp := int64(1717334400000) // int64 | DTO.TRADE.QUERY.AFTER_TIMESTAMP (optional)
	beforeBlockHeight := int64(332417228) // int64 | DTO.TRADE.QUERY.BEFORE_BLOCK_HEIGHT (optional)
	afterBlockHeight := int64(332417228) // int64 | DTO.TRADE.QUERY.AFTER_BLOCK_HEIGHT (optional)
	transactionsSignature := "37XpPt9Ak6JiE1V3sftJDtdUsvR9FVFRqkZmoT3dp4BTD9pgyTWn1XgHH6R7NjuJ4pBMAgj8JvZtxQrf4s6NTC5F,3QF8Fn4ReoEjQhfZHvJy8ykodBJRZktcP21j1bQ8aM6uFXQV1CuqUoPDLNGJpkUC6bLhghxWcf54VYzRaPM66GwH" // string | DTO.TRADE.QUERY.TRANSACTIONS_SIGNATURE (optional)
	type_ := "BUY" // string | DTO.TRADE.QUERY.ACTIVITIES_TYPE (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetActivities(context.Background(), chain).Cursor(cursor).Limit(limit).Direction(direction).TokenAddress(tokenAddress).WalletAddress(walletAddress).PoolAddress(poolAddress).BeforeTimestamp(beforeTimestamp).AfterTimestamp(afterTimestamp).BeforeBlockHeight(beforeBlockHeight).AfterBlockHeight(afterBlockHeight).TransactionsSignature(transactionsSignature).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivities`: TradePage
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]
 **tokenAddress** | **string** | DTO.TRADE.QUERY.TOKEN_ADDRESS | 
 **walletAddress** | **string** | DTO.TRADE.QUERY.WALLET_ADDRESS | 
 **poolAddress** | **string** | DTO.TRADE.QUERY.POOL_ADDRESS | 
 **beforeTimestamp** | **int64** | DTO.TRADE.QUERY.BEFORE_TIMESTAMP | 
 **afterTimestamp** | **int64** | DTO.TRADE.QUERY.AFTER_TIMESTAMP | 
 **beforeBlockHeight** | **int64** | DTO.TRADE.QUERY.BEFORE_BLOCK_HEIGHT | 
 **afterBlockHeight** | **int64** | DTO.TRADE.QUERY.AFTER_BLOCK_HEIGHT | 
 **transactionsSignature** | **string** | DTO.TRADE.QUERY.TRANSACTIONS_SIGNATURE | 
 **type_** | **string** | DTO.TRADE.QUERY.ACTIVITIES_TYPE | 

### Return type

[**TradePage**](TradePage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrades

> TradePage GetTrades(ctx, chain).Cursor(cursor).Limit(limit).Direction(direction).TokenAddress(tokenAddress).WalletAddress(walletAddress).PoolAddress(poolAddress).BeforeTimestamp(beforeTimestamp).AfterTimestamp(afterTimestamp).BeforeBlockHeight(beforeBlockHeight).AfterBlockHeight(afterBlockHeight).TransactionsSignature(transactionsSignature).Type_(type_).Execute()

CONTROLLER.TRADE.GET.TOKEN.SUMMARY



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/chainstream-io/chainstream-go-sdk"
)

func main() {
	chain := openapiclient.ChainSymbol("sol") // ChainSymbol | GLOBAL.CHAIN.DESCRIPTION
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR.DESCRIPTION (optional)
	limit := float32(8.14) // float32 | DTO.PAGE.LIMIT (optional) (default to 20)
	direction := "direction_example" // string | DTO.PAGE.DIRECTION (optional) (default to "next")
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.TRADE.QUERY.TOKEN_ADDRESS (optional)
	walletAddress := "3xd4LGVWtYXLBspR6X5JWbW49NXmEehfPtX6Kqx98b4w" // string | DTO.TRADE.QUERY.WALLET_ADDRESS (optional)
	poolAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.TRADE.QUERY.POOL_ADDRESS (optional)
	beforeTimestamp := int64(1717334400000) // int64 | DTO.TRADE.QUERY.BEFORE_TIMESTAMP (optional)
	afterTimestamp := int64(1717334400000) // int64 | DTO.TRADE.QUERY.AFTER_TIMESTAMP (optional)
	beforeBlockHeight := int64(332417228) // int64 | DTO.TRADE.QUERY.BEFORE_BLOCK_HEIGHT (optional)
	afterBlockHeight := int64(332417228) // int64 | DTO.TRADE.QUERY.AFTER_BLOCK_HEIGHT (optional)
	transactionsSignature := "37XpPt9Ak6JiE1V3sftJDtdUsvR9FVFRqkZmoT3dp4BTD9pgyTWn1XgHH6R7NjuJ4pBMAgj8JvZtxQrf4s6NTC5F,3QF8Fn4ReoEjQhfZHvJy8ykodBJRZktcP21j1bQ8aM6uFXQV1CuqUoPDLNGJpkUC6bLhghxWcf54VYzRaPM66GwH" // string | DTO.TRADE.QUERY.TRANSACTIONS_SIGNATURE (optional)
	type_ := "BUY" // string | DTO.TRADE.QUERY.TRADES_TYPE (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTrades(context.Background(), chain).Cursor(cursor).Limit(limit).Direction(direction).TokenAddress(tokenAddress).WalletAddress(walletAddress).PoolAddress(poolAddress).BeforeTimestamp(beforeTimestamp).AfterTimestamp(afterTimestamp).BeforeBlockHeight(beforeBlockHeight).AfterBlockHeight(afterBlockHeight).TransactionsSignature(transactionsSignature).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrades`: TradePage
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]
 **tokenAddress** | **string** | DTO.TRADE.QUERY.TOKEN_ADDRESS | 
 **walletAddress** | **string** | DTO.TRADE.QUERY.WALLET_ADDRESS | 
 **poolAddress** | **string** | DTO.TRADE.QUERY.POOL_ADDRESS | 
 **beforeTimestamp** | **int64** | DTO.TRADE.QUERY.BEFORE_TIMESTAMP | 
 **afterTimestamp** | **int64** | DTO.TRADE.QUERY.AFTER_TIMESTAMP | 
 **beforeBlockHeight** | **int64** | DTO.TRADE.QUERY.BEFORE_BLOCK_HEIGHT | 
 **afterBlockHeight** | **int64** | DTO.TRADE.QUERY.AFTER_BLOCK_HEIGHT | 
 **transactionsSignature** | **string** | DTO.TRADE.QUERY.TRANSACTIONS_SIGNATURE | 
 **type_** | **string** | DTO.TRADE.QUERY.TRADES_TYPE | 

### Return type

[**TradePage**](TradePage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

