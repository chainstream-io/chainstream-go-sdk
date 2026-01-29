# \DexAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](DexAPI.md#CreateToken) | **Post** /v1/dex/{chain}/create | Dex - Create Token
[**ListDex**](DexAPI.md#ListDex) | **Get** /v1/dex | Dex - List
[**Quote**](DexAPI.md#Quote) | **Get** /v1/dex/{chain}/quote | Dex - Get Quote
[**Route**](DexAPI.md#Route) | **Post** /v1/dex/{chain}/route | Dex - Route
[**Swap**](DexAPI.md#Swap) | **Post** /v1/dex/{chain}/swap | Dex - Swap



## CreateToken

> CreateTokenReply CreateToken(ctx, chain).CreateTokenInput(createTokenInput).Execute()

Dex - Create Token



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
	createTokenInput := *openapiclient.NewCreateTokenInput("raydium", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "Candy Token", "CANDY") // CreateTokenInput | Token creation parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.CreateToken(context.Background(), chain).CreateTokenInput(createTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.CreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateToken`: CreateTokenReply
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTokenInput** | [**CreateTokenInput**](CreateTokenInput.md) | Token creation parameters | 

### Return type

[**CreateTokenReply**](CreateTokenReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDex

> DexPage ListDex(ctx).Chains(chains).Limit(limit).DexProgram(dexProgram).Execute()

Dex - List



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
	chains := []string{"Inner_example"} // []string |  (optional)
	limit := int64(789) // int64 | Number of results per page (optional) (default to 20)
	dexProgram := "dexProgram_example" // string | dex program address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.ListDex(context.Background()).Chains(chains).Limit(limit).DexProgram(dexProgram).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.ListDex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDex`: DexPage
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.ListDex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chains** | **[]string** |  | 
 **limit** | **int64** | Number of results per page | [default to 20]
 **dexProgram** | **string** | dex program address | 

### Return type

[**DexPage**](DexPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Quote

> QuoteResponse Quote(ctx, chain).Dex(dex).Amount(amount).InputMint(inputMint).OutputMint(outputMint).ExactIn(exactIn).Slippage(slippage).Execute()

Dex - Get Quote



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
	dex := "raydium" // string | DEX protocol type
	amount := "1000000000" // string | Trading amount
	inputMint := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | Input token address
	outputMint := "So11111111111111111111111111111111111111112" // string | Output token address
	exactIn := true // bool | Exact input mode (default to true)
	slippage := int64(10) // int64 | Slippage tolerance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.Quote(context.Background(), chain).Dex(dex).Amount(amount).InputMint(inputMint).OutputMint(outputMint).ExactIn(exactIn).Slippage(slippage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.Quote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Quote`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.Quote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dex** | **string** | DEX protocol type | 
 **amount** | **string** | Trading amount | 
 **inputMint** | **string** | Input token address | 
 **outputMint** | **string** | Output token address | 
 **exactIn** | **bool** | Exact input mode | [default to true]
 **slippage** | **int64** | Slippage tolerance | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Route

> SwapRouteResponse Route(ctx, chain).SwapRouteInput(swapRouteInput).Execute()

Dex - Route



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
	swapRouteInput := *openapiclient.NewSwapRouteInput("jupiter", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "1000000000", "ExactIn", int64(5)) // SwapRouteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.Route(context.Background(), chain).SwapRouteInput(swapRouteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.Route``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Route`: SwapRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.Route`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **swapRouteInput** | [**SwapRouteInput**](SwapRouteInput.md) |  | 

### Return type

[**SwapRouteResponse**](SwapRouteResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Swap

> SwapReply Swap(ctx, chain).SwapInput(swapInput).Execute()

Dex - Swap



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
	swapInput := *openapiclient.NewSwapInput("raydium", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "1000000000", "ExactIn", int64(10)) // SwapInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexAPI.Swap(context.Background(), chain).SwapInput(swapInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexAPI.Swap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Swap`: SwapReply
	fmt.Fprintf(os.Stdout, "Response from `DexAPI.Swap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **swapInput** | [**SwapInput**](SwapInput.md) |  | 

### Return type

[**SwapReply**](SwapReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

