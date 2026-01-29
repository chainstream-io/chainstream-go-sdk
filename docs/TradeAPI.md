# \TradeAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivities**](TradeAPI.md#GetActivities) | **Get** /v1/trade/{chain}/activities | Activity - List
[**GetTopTraders**](TradeAPI.md#GetTopTraders) | **Get** /v1/trade/{chain}/top-traders | Trade - Top Traders
[**GetTrades**](TradeAPI.md#GetTrades) | **Get** /v1/trade/{chain} | Trade - List



## GetActivities

> TradePage GetActivities(ctx, chain).Cursor(cursor).Limit(limit).Direction(direction).TokenAddress(tokenAddress).WalletAddress(walletAddress).PoolAddress(poolAddress).BeforeTimestamp(beforeTimestamp).AfterTimestamp(afterTimestamp).BeforeBlockHeight(beforeBlockHeight).AfterBlockHeight(afterBlockHeight).TransactionsSignature(transactionsSignature).Type_(type_).Execute()

Activity - List



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
	chain := openapiclient.ChainSymbol("sol") // ChainSymbol | A chain name listed in supported networks
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := float32(8.14) // float32 | Number of results per page (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Token address to query trades (optional)
	walletAddress := "3xd4LGVWtYXLBspR6X5JWbW49NXmEehfPtX6Kqx98b4w" // string | Wallet address to query trades (optional)
	poolAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Pool address to filter trades (optional)
	beforeTimestamp := int64(1717334400000) // int64 | Start timestamp for filtering trades (Unix epoch in seconds) (optional)
	afterTimestamp := int64(1717334400000) // int64 | End timestamp for filtering trades (Unix epoch in seconds) (optional)
	beforeBlockHeight := int64(332417228) // int64 | Filter trades before this block height (optional)
	afterBlockHeight := int64(332417228) // int64 | Filter trades after this block height (optional)
	transactionsSignature := "37XpPt9Ak6JiE1V3sftJDtdUsvR9FVFRqkZmoT3dp4BTD9pgyTWn1XgHH6R7NjuJ4pBMAgj8JvZtxQrf4s6NTC5F,3QF8Fn4ReoEjQhfZHvJy8ykodBJRZktcP21j1bQ8aM6uFXQV1CuqUoPDLNGJpkUC6bLhghxWcf54VYzRaPM66GwH" // string | Transaction signature/hash (optional)
	type_ := "BUY" // string | Activity type (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]
 **tokenAddress** | **string** | Token address to query trades | 
 **walletAddress** | **string** | Wallet address to query trades | 
 **poolAddress** | **string** | Pool address to filter trades | 
 **beforeTimestamp** | **int64** | Start timestamp for filtering trades (Unix epoch in seconds) | 
 **afterTimestamp** | **int64** | End timestamp for filtering trades (Unix epoch in seconds) | 
 **beforeBlockHeight** | **int64** | Filter trades before this block height | 
 **afterBlockHeight** | **int64** | Filter trades after this block height | 
 **transactionsSignature** | **string** | Transaction signature/hash | 
 **type_** | **string** | Activity type | 

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


## GetTopTraders

> TopTradersPage GetTopTraders(ctx, chain).TokenAddress(tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).TimeFrame(timeFrame).SortType(sortType).SortBy(sortBy).Execute()

Trade - Top Traders



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
	chain := openapiclient.ChainSymbol("sol") // ChainSymbol | A chain name listed in supported networks
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Token address to query trades
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := int64(789) // int64 | Maximum number of top traders to return (max 10) (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")
	timeFrame := "24h" // string | Time frame for filtering trades (optional) (default to "24h")
	sortType := "desc" // string | Sort type for trade results (optional) (default to "desc")
	sortBy := "tradeAmount" // string | Field to sort trades by (optional) (default to "tradeAmount")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTopTraders(context.Background(), chain).TokenAddress(tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).TimeFrame(timeFrame).SortType(sortType).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTopTraders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopTraders`: TopTradersPage
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTopTraders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopTradersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddress** | **string** | Token address to query trades | 
 **cursor** | **string** | Pagination cursor | 
 **limit** | **int64** | Maximum number of top traders to return (max 10) | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]
 **timeFrame** | **string** | Time frame for filtering trades | [default to &quot;24h&quot;]
 **sortType** | **string** | Sort type for trade results | [default to &quot;desc&quot;]
 **sortBy** | **string** | Field to sort trades by | [default to &quot;tradeAmount&quot;]

### Return type

[**TopTradersPage**](TopTradersPage.md)

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

Trade - List



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
	chain := openapiclient.ChainSymbol("sol") // ChainSymbol | A chain name listed in supported networks
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := float32(8.14) // float32 | Number of results per page (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Token address to query trades (optional)
	walletAddress := "3xd4LGVWtYXLBspR6X5JWbW49NXmEehfPtX6Kqx98b4w" // string | Wallet address to query trades (optional)
	poolAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Pool address to filter trades (optional)
	beforeTimestamp := int64(1717334400000) // int64 | Start timestamp for filtering trades (Unix epoch in seconds) (optional)
	afterTimestamp := int64(1717334400000) // int64 | End timestamp for filtering trades (Unix epoch in seconds) (optional)
	beforeBlockHeight := int64(332417228) // int64 | Filter trades before this block height (optional)
	afterBlockHeight := int64(332417228) // int64 | Filter trades after this block height (optional)
	transactionsSignature := "37XpPt9Ak6JiE1V3sftJDtdUsvR9FVFRqkZmoT3dp4BTD9pgyTWn1XgHH6R7NjuJ4pBMAgj8JvZtxQrf4s6NTC5F,3QF8Fn4ReoEjQhfZHvJy8ykodBJRZktcP21j1bQ8aM6uFXQV1CuqUoPDLNGJpkUC6bLhghxWcf54VYzRaPM66GwH" // string | Transaction signature/hash (optional)
	type_ := "BUY" // string | Trade type (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]
 **tokenAddress** | **string** | Token address to query trades | 
 **walletAddress** | **string** | Wallet address to query trades | 
 **poolAddress** | **string** | Pool address to filter trades | 
 **beforeTimestamp** | **int64** | Start timestamp for filtering trades (Unix epoch in seconds) | 
 **afterTimestamp** | **int64** | End timestamp for filtering trades (Unix epoch in seconds) | 
 **beforeBlockHeight** | **int64** | Filter trades before this block height | 
 **afterBlockHeight** | **int64** | Filter trades after this block height | 
 **transactionsSignature** | **string** | Transaction signature/hash | 
 **type_** | **string** | Trade type | 

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

