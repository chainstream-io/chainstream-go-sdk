# \DexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](DexAPI.md#CreateToken) | **Post** /v1/dex/{chain}/create | CONTROLLER.DEX.CREATE.SUMMARY
[**ListDex**](DexAPI.md#ListDex) | **Get** /v1/dex | CONTROLLER.DEX.LIST.SUMMARY
[**Quote**](DexAPI.md#Quote) | **Get** /v1/dex/{chain}/quote | CONTROLLER.DEX.QUOTE.SUMMARY
[**Route**](DexAPI.md#Route) | **Post** /v1/dex/{chain}/route | CONTROLLER.DEX.ROUTE.SUMMARY
[**Swap**](DexAPI.md#Swap) | **Post** /v1/dex/{chain}/swap | CONTROLLER.DEX.SWAP.SUMMARY



## CreateToken

> CreateTokenReply CreateToken(ctx, chain).CreateTokenInput(createTokenInput).Execute()

CONTROLLER.DEX.CREATE.SUMMARY



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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

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

CONTROLLER.DEX.LIST.SUMMARY



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
	limit := float32(8.14) // float32 | DTO.DEX.QUERY.LIMIT (optional) (default to 20)
	dexProgram := "dexProgram_example" // string | DTO.DEX.QUERY.DEX_PROGRAM (optional)

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
 **limit** | **float32** | DTO.DEX.QUERY.LIMIT | [default to 20]
 **dexProgram** | **string** | DTO.DEX.QUERY.DEX_PROGRAM | 

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

CONTROLLER.DEX.QUOTE.SUMMARY



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
	dex := "raydium" // string | DTO.DEX.QUOTE.DEX
	amount := "1000000000" // string | DTO.DEX.QUOTE.AMOUNT
	inputMint := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | DTO.DEX.QUOTE.INPUT_MINT
	outputMint := "So11111111111111111111111111111111111111112" // string | DTO.DEX.QUOTE.OUTPUT_MINT
	exactIn := true // bool | DTO.DEX.QUOTE.EXACT_IN (default to true)
	slippage := float32(10) // float32 | DTO.DEX.QUOTE.SLIPPAGE

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dex** | **string** | DTO.DEX.QUOTE.DEX | 
 **amount** | **string** | DTO.DEX.QUOTE.AMOUNT | 
 **inputMint** | **string** | DTO.DEX.QUOTE.INPUT_MINT | 
 **outputMint** | **string** | DTO.DEX.QUOTE.OUTPUT_MINT | 
 **exactIn** | **bool** | DTO.DEX.QUOTE.EXACT_IN | [default to true]
 **slippage** | **float32** | DTO.DEX.QUOTE.SLIPPAGE | 

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

CONTROLLER.DEX.ROUTE.SUMMARY



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
	swapRouteInput := *openapiclient.NewSwapRouteInput("jupiter", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "1000000000", "ExactIn", float32(5)) // SwapRouteInput | 

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

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

CONTROLLER.DEX.SWAP.SUMMARY



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
	swapInput := *openapiclient.NewSwapInput("raydium", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "1000000000", "ExactIn", float32(10)) // SwapInput | 

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

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

