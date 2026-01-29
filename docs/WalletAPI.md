# \WalletAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculatePnl**](WalletAPI.md#CalculatePnl) | **Post** /v1/wallet/{chain}/{walletAddress}/calculate-pnl | Wallet - Calculate PNL
[**GetBalance**](WalletAPI.md#GetBalance) | **Get** /v1/wallet/{chain}/{walletAddress}/balance | Wallet - Balances
[**GetPnl**](WalletAPI.md#GetPnl) | **Get** /v1/wallet/{chain}/{walletAddress} | Wallet - PNL
[**GetPnlStats**](WalletAPI.md#GetPnlStats) | **Get** /v1/wallet/{chain}/{walletAddress}/stats | Wallet - PNL Stats



## CalculatePnl

> BooleanResultDTO CalculatePnl(ctx, chain, walletAddress).CalculatePnlInput(calculatePnlInput).Execute()

Wallet - Calculate PNL



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | An address of a wallet
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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**walletAddress** | **string** | An address of a wallet | 

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

Wallet - Balances



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | An address of a wallet
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Specific token address to filter wallet balances (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**walletAddress** | **string** | An address of a wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenAddress** | **string** | Specific token address to filter wallet balances | 

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

Wallet - PNL



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | An address of a wallet
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**walletAddress** | **string** | An address of a wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPnlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenAddress** | **string** | An address of a token | 

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

Wallet - PNL Stats



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
	walletAddress := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2" // string | An address of a wallet

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**walletAddress** | **string** | An address of a wallet | 

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

