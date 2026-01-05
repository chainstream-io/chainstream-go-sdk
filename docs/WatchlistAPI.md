# \WatchlistAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchlistAdd**](WatchlistAPI.md#WatchlistAdd) | **Post** /v1/watchlist/{chain}/{walletAddress} | CONTROLLER.WATCHLIST.ADD.SUMMARY



## WatchlistAdd

> BooleanResultDTO WatchlistAdd(ctx, chain, walletAddress).Execute()

CONTROLLER.WATCHLIST.ADD.SUMMARY



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
	walletAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.WALLETADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchlistAPI.WatchlistAdd(context.Background(), chain, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistAPI.WatchlistAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchlistAdd`: BooleanResultDTO
	fmt.Fprintf(os.Stdout, "Response from `WatchlistAPI.WatchlistAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**walletAddress** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchlistAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BooleanResultDTO**](BooleanResultDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

