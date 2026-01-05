# \WalletAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculatePnl**](WalletAPI.md#CalculatePnl) | **Post** /v1/wallet/{chain}/{walletAddress}/calculate-pnl | CONTROLLER.WALLET.CALCULATE_PNL.SUMMARY
[**GetBalance**](WalletAPI.md#GetBalance) | **Get** /v1/wallet/{chain}/{walletAddress}/balance | CONTROLLER.WALLET.GET_BALANCES.SUMMARY
[**GetPnl**](WalletAPI.md#GetPnl) | **Get** /v1/wallet/{chain}/{walletAddress} | CONTROLLER.WALLET.GET_PNL.SUMMARY
[**GetPnlStats**](WalletAPI.md#GetPnlStats) | **Get** /v1/wallet/{chain}/{walletAddress}/stats | CONTROLLER.WALLET.GET_PNL_STATS.SUMMARY



## CalculatePnl

> BooleanResultDTO CalculatePnl(ctx, chain, walletAddress).CalculatePnlInput(calculatePnlInput).Execute()

CONTROLLER.WALLET.CALCULATE_PNL.SUMMARY



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | GLOBAL.WALLETADDRESS.DESCRIPTION
	calculatePnlInput := *openapiclient.NewCalculatePnlInput() // CalculatePnlInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CalculatePnl(context.Background(), chain, walletAddress).CalculatePnlInput(calculatePnlInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CalculatePnl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculatePnl`: BooleanResultDTO
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CalculatePnl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**walletAddress** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculatePnlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **calculatePnlInput** | [**CalculatePnlInput**](CalculatePnlInput.md) |  | 

### Return type

[**BooleanResultDTO**](BooleanResultDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalance

> WalletBalancesDTO GetBalance(ctx, chain, walletAddress).TokenAddress(tokenAddress).Execute()

CONTROLLER.WALLET.GET_BALANCES.SUMMARY



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | GLOBAL.WALLETADDRESS.DESCRIPTION
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.WALLET.BALANCE.QUERY.TOKEN_ADDRESS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetBalance(context.Background(), chain, walletAddress).TokenAddress(tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: WalletBalancesDTO
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**walletAddress** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenAddress** | **string** | DTO.WALLET.BALANCE.QUERY.TOKEN_ADDRESS | 

### Return type

[**WalletBalancesDTO**](WalletBalancesDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPnl

> []WalletPnlDTO GetPnl(ctx, chain, walletAddress).TokenAddress(tokenAddress).Execute()

CONTROLLER.WALLET.GET_PNL.SUMMARY



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | GLOBAL.WALLETADDRESS.DESCRIPTION
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetPnl(context.Background(), chain, walletAddress).TokenAddress(tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetPnl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPnl`: []WalletPnlDTO
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetPnl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**walletAddress** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPnlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Return type

[**[]WalletPnlDTO**](WalletPnlDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPnlStats

> []WalletPnlDTO GetPnlStats(ctx, chain, walletAddress).Execute()

CONTROLLER.WALLET.GET_PNL_STATS.SUMMARY



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | GLOBAL.WALLETADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetPnlStats(context.Background(), chain, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetPnlStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPnlStats`: []WalletPnlDTO
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetPnlStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**walletAddress** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPnlStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WalletPnlDTO**](WalletPnlDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

