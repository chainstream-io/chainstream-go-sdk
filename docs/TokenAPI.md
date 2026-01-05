# \TokenAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCandles**](TokenAPI.md#GetCandles) | **Get** /v1/token/{chain}/{tokenAddress}/candles | CONTROLLER.TOKEN.GET_CANDLES.SUMMARY
[**GetCreation**](TokenAPI.md#GetCreation) | **Get** /v1/token/{chain}/{tokenAddress}/creation | CONTROLLER.TOKEN.GET_CREATION.SUMMARY
[**GetDevTokens**](TokenAPI.md#GetDevTokens) | **Get** /v1/token/{chain}/dev/{devAddress}/tokens | CONTROLLER.TOKEN.GET_DEV_TOKENS.SUMMARY
[**GetHolders**](TokenAPI.md#GetHolders) | **Get** /v1/token/{chain}/{tokenAddress}/holders | CONTROLLER.TOKEN.GET_HOLDERS.SUMMARY
[**GetHoldersMulti**](TokenAPI.md#GetHoldersMulti) | **Get** /v1/token/{chain}/{tokenAddress}/holders/multi | CONTROLLER.TOKEN.GET_HOLDERS_MULTI.SUMMARY
[**GetMarketData**](TokenAPI.md#GetMarketData) | **Get** /v1/token/{chain}/{tokenAddress}/marketData | CONTROLLER.TOKEN.GET_MARKET_CAP.SUMMARY
[**GetMarketDataMulti**](TokenAPI.md#GetMarketDataMulti) | **Get** /v1/token/{chain}/marketData/multi | CONTROLLER.TOKEN.GET_MARKET_DATA_MULTI.SUMMARY
[**GetMetadata**](TokenAPI.md#GetMetadata) | **Get** /v1/token/{chain}/{tokenAddress}/metadata | CONTROLLER.TOKEN.GET_METADATA.SUMMARY
[**GetMetadataMulti**](TokenAPI.md#GetMetadataMulti) | **Get** /v1/token/{chain}/metadata/multi | CONTROLLER.TOKEN.GET_METADATA_MULTI.SUMMARY
[**GetMintAndBurn**](TokenAPI.md#GetMintAndBurn) | **Get** /v1/token/{chain}/{tokenAddress}/mintAndBurn | CONTROLLER.TOKEN.GET_MINT_AND_BURN.SUMMARY
[**GetPools**](TokenAPI.md#GetPools) | **Get** /v1/token/{chain}/{tokenAddress}/pools | CONTROLLER.TOKEN.GET_POOLS.SUMMARY
[**GetPriceByTime**](TokenAPI.md#GetPriceByTime) | **Get** /v1/token/{chain}/{tokenAddress}/price | CONTROLLER.TOKEN.GET_PRICE_BY_TIME.SUMMARY
[**GetPrices**](TokenAPI.md#GetPrices) | **Get** /v1/token/{chain}/{tokenAddress}/prices | CONTROLLER.TOKEN.GET_PRICES.SUMMARY
[**GetSecurity**](TokenAPI.md#GetSecurity) | **Get** /v1/token/{chain}/{tokenAddress}/security | CONTROLLER.TOKEN.GET_SECURITY.SUMMARY
[**GetStats**](TokenAPI.md#GetStats) | **Get** /v1/token/{chain}/{tokenAddress}/stats | CONTROLLER.TOKEN.GET_STATS.SUMMARY
[**GetStatsMulti**](TokenAPI.md#GetStatsMulti) | **Get** /v1/token/{chain}/stats/multi | CONTROLLER.TOKEN.GET_STATS_MULTI.SUMMARY
[**GetToken**](TokenAPI.md#GetToken) | **Get** /v1/token/{chain}/{tokenAddress} | CONTROLLER.TOKEN.GET.SUMMARY
[**GetTokenTraders**](TokenAPI.md#GetTokenTraders) | **Get** /v1/token/{chain}/{tokenAddress}/traders/{tag} | CONTROLLER.TOKEN.GET_TOKEN_TRADERS.SUMMARY
[**GetTokens**](TokenAPI.md#GetTokens) | **Get** /v1/token/{chain}/multi | CONTROLLER.TOKEN.GET_TOKENS.SUMMARY
[**GetTopHolders**](TokenAPI.md#GetTopHolders) | **Get** /v1/token/{chain}/{tokenAddress}/topHolders | CONTROLLER.TOKEN.GET_TOP_HOLDERS.SUMMARY
[**ListToken**](TokenAPI.md#ListToken) | **Get** /v1/token/{chain}/list | CONTROLLER.TOKEN.GET_TOKEN_LIST.SUMMARY
[**Search**](TokenAPI.md#Search) | **Get** /v1/token/search | CONTROLLER.TOKEN.SEARCH.SUMMARY



## GetCandles

> []Candle GetCandles(ctx, chain, tokenAddress).Resolution(resolution).From(from).To(to).Limit(limit).Execute()

CONTROLLER.TOKEN.GET_CANDLES.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	resolution := openapiclient.Resolution("1s") // Resolution | DTO.CANDLE.RESOLUTION
	from := float32(1741647950000) // float32 | DTO.CANDLE.FROM (optional)
	to := float32(1741700970000) // float32 | DTO.CANDLE.TO (optional)
	limit := float32(100) // float32 | DTO.CANDLE.LIMIT (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetCandles(context.Background(), chain, tokenAddress).Resolution(resolution).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetCandles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCandles`: []Candle
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetCandles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolution** | [**Resolution**](Resolution.md) | DTO.CANDLE.RESOLUTION | 
 **from** | **float32** | DTO.CANDLE.FROM | 
 **to** | **float32** | DTO.CANDLE.TO | 
 **limit** | **float32** | DTO.CANDLE.LIMIT | 

### Return type

[**[]Candle**](Candle.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreation

> TokenCreationDTO GetCreation(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_CREATION.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetCreation(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetCreation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreation`: TokenCreationDTO
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetCreation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenCreationDTO**](TokenCreationDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevTokens

> []DevTokenDTO GetDevTokens(ctx, chain, devAddress).Execute()

CONTROLLER.TOKEN.GET_DEV_TOKENS.SUMMARY



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
	devAddress := "3NUHqNG2Big2WxbYnaBzSNmxwE4NrxgckJuoiyGxu3Am" // string | Developer address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetDevTokens(context.Background(), chain, devAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetDevTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevTokens`: []DevTokenDTO
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetDevTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**devAddress** | **string** | Developer address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DevTokenDTO**](DevTokenDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHolders

> TokenHolderPage GetHolders(ctx, chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Execute()

CONTROLLER.TOKEN.GET_HOLDERS.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR.DESCRIPTION (optional)
	limit := float32(8.14) // float32 | DTO.PAGE.LIMIT (optional) (default to 20)
	direction := "direction_example" // string | DTO.PAGE.DIRECTION (optional) (default to "next")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetHolders(context.Background(), chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHolders`: TokenHolderPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]

### Return type

[**TokenHolderPage**](TokenHolderPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHoldersMulti

> []TokenHolder GetHoldersMulti(ctx, chain, tokenAddress).WalletAddresses(walletAddresses).Execute()

CONTROLLER.TOKEN.GET_HOLDERS_MULTI.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	walletAddresses := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2,2RH6rUTPBJ9rUDPpuV9b8z1YL56k1tYU6Uk5ZoaEFFSK" // string | GLOBAL.WALLETADDRESSES.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetHoldersMulti(context.Background(), chain, tokenAddress).WalletAddresses(walletAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetHoldersMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHoldersMulti`: []TokenHolder
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetHoldersMulti`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldersMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **walletAddresses** | **string** | GLOBAL.WALLETADDRESSES.DESCRIPTION | 

### Return type

[**[]TokenHolder**](TokenHolder.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketData

> TokenMarketData GetMarketData(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_MARKET_CAP.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetMarketData(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetMarketData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketData`: TokenMarketData
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetMarketData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenMarketData**](TokenMarketData.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketDataMulti

> map[string]TokenMarketData GetMarketDataMulti(ctx, chain).TokenAddresses(tokenAddresses).Execute()

CONTROLLER.TOKEN.GET_MARKET_DATA_MULTI.SUMMARY



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | GLOBAL.TOKENADDRESSES.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetMarketDataMulti(context.Background(), chain).TokenAddresses(tokenAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetMarketDataMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketDataMulti`: map[string]TokenMarketData
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetMarketDataMulti`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketDataMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | GLOBAL.TOKENADDRESSES.DESCRIPTION | 

### Return type

[**map[string]TokenMarketData**](TokenMarketData.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadata

> TokenMetadata GetMetadata(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_METADATA.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetMetadata(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadata`: TokenMetadata
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenMetadata**](TokenMetadata.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataMulti

> map[string]TokenMetadata GetMetadataMulti(ctx, chain).TokenAddresses(tokenAddresses).Execute()

CONTROLLER.TOKEN.GET_METADATA_MULTI.SUMMARY



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | GLOBAL.TOKENADDRESSES.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetMetadataMulti(context.Background(), chain).TokenAddresses(tokenAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetMetadataMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadataMulti`: map[string]TokenMetadata
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetMetadataMulti`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | GLOBAL.TOKENADDRESSES.DESCRIPTION | 

### Return type

[**map[string]TokenMetadata**](TokenMetadata.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMintAndBurn

> TokenCreationPage GetMintAndBurn(ctx, chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Type_(type_).Execute()

CONTROLLER.TOKEN.GET_MINT_AND_BURN.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR.DESCRIPTION (optional)
	limit := float32(8.14) // float32 | DTO.PAGE.LIMIT (optional) (default to 20)
	direction := "direction_example" // string | DTO.PAGE.DIRECTION (optional) (default to "next")
	type_ := "all" // string | DTO.TOKEN.MINT_AND_BURN.TYPE (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetMintAndBurn(context.Background(), chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetMintAndBurn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMintAndBurn`: TokenCreationPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetMintAndBurn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintAndBurnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]
 **type_** | **string** | DTO.TOKEN.MINT_AND_BURN.TYPE | [default to &quot;all&quot;]

### Return type

[**TokenCreationPage**](TokenCreationPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPools

> []DexPoolDTO GetPools(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_POOLS.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetPools(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPools`: []DexPoolDTO
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DexPoolDTO**](DexPoolDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceByTime

> TokenPriceDTO GetPriceByTime(ctx, chain, tokenAddress).Timestamp(timestamp).Execute()

CONTROLLER.TOKEN.GET_PRICE_BY_TIME.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	timestamp := float32(1754055151) // float32 | DTO.TOKEN.PRICE.QUERY.TIMESTAMP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetPriceByTime(context.Background(), chain, tokenAddress).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetPriceByTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceByTime`: TokenPriceDTO
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetPriceByTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timestamp** | **float32** | DTO.TOKEN.PRICE.QUERY.TIMESTAMP | 

### Return type

[**TokenPriceDTO**](TokenPriceDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrices

> TokenPricePage GetPrices(ctx, chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Execute()

CONTROLLER.TOKEN.GET_PRICES.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR.DESCRIPTION (optional)
	limit := float32(8.14) // float32 | DTO.PAGE.LIMIT (optional) (default to 20)
	direction := "direction_example" // string | DTO.PAGE.DIRECTION (optional) (default to "next")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetPrices(context.Background(), chain, tokenAddress).Cursor(cursor).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrices`: TokenPricePage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]

### Return type

[**TokenPricePage**](TokenPricePage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurity

> map[string]interface{} GetSecurity(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_SECURITY.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetSecurity(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStats

> TokenStat GetStats(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_STATS.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetStats(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStats`: TokenStat
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenStat**](TokenStat.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatsMulti

> map[string]TokenStat GetStatsMulti(ctx, chain).TokenAddresses(tokenAddresses).Execute()

CONTROLLER.TOKEN.GET_STATS_MULTI.SUMMARY



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | GLOBAL.TOKENADDRESSES.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetStatsMulti(context.Background(), chain).TokenAddresses(tokenAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetStatsMulti``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatsMulti`: map[string]TokenStat
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetStatsMulti`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | GLOBAL.TOKENADDRESSES.DESCRIPTION | 

### Return type

[**map[string]TokenStat**](TokenStat.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> Token GetToken(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetToken(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: Token
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Token**](Token.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenTraders

> []TokenTrader GetTokenTraders(ctx, chain, tokenAddress, tag).Execute()

CONTROLLER.TOKEN.GET_TOKEN_TRADERS.SUMMARY



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
	tokenAddress := "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | GLOBAL.TOKENADDRESS.DESCRIPTION
	tag := openapiclient.TokenTraderTag("fresh") // TokenTraderTag | Token trader tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTokenTraders(context.Background(), chain, tokenAddress, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokenTraders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenTraders`: []TokenTrader
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokenTraders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 
**tag** | [**TokenTraderTag**](.md) | Token trader tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenTradersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]TokenTrader**](TokenTrader.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> []Token GetTokens(ctx, chain).TokenAddresses(tokenAddresses).SortBy(sortBy).SortDirection(sortDirection).FilterBy(filterBy).Execute()

CONTROLLER.TOKEN.GET_TOKENS.SUMMARY



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | GLOBAL.TOKENADDRESSES.DESCRIPTION
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	filterBy := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTokens(context.Background(), chain).TokenAddresses(tokenAddresses).SortBy(sortBy).SortDirection(sortDirection).FilterBy(filterBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | GLOBAL.TOKENADDRESSES.DESCRIPTION | 
 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **filterBy** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 

### Return type

[**[]Token**](Token.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopHolders

> TokenHolderPage GetTopHolders(ctx, chain, tokenAddress).Execute()

CONTROLLER.TOKEN.GET_TOP_HOLDERS.SUMMARY



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | GLOBAL.TOKENADDRESS.DESCRIPTION

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTopHolders(context.Background(), chain, tokenAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTopHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopHolders`: TokenHolderPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTopHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**tokenAddress** | **string** | GLOBAL.TOKENADDRESS.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenHolderPage**](TokenHolderPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListToken

> TokenListPage ListToken(ctx, chain).Cursor(cursor).Limit(limit).Direction(direction).Sort(sort).SortBy(sortBy).MinH24VolumeInUsd(minH24VolumeInUsd).MaxH24VolumeInUsd(maxH24VolumeInUsd).MinH24PriceChangeRatio(minH24PriceChangeRatio).MaxH24PriceChangeRatio(maxH24PriceChangeRatio).MinH24Buys(minH24Buys).MaxH24Buys(maxH24Buys).MinH24Sells(minH24Sells).MaxH24Sells(maxH24Sells).MinH24Trades(minH24Trades).MaxH24Trades(maxH24Trades).MinH24Buyers(minH24Buyers).MaxH24Buyers(maxH24Buyers).MinH24Sellers(minH24Sellers).MaxH24Sellers(maxH24Sellers).MinH24BuyVolumeInUsd(minH24BuyVolumeInUsd).MaxH24BuyVolumeInUsd(maxH24BuyVolumeInUsd).MinH24SellVolumeInUsd(minH24SellVolumeInUsd).MaxH24SellVolumeInUsd(maxH24SellVolumeInUsd).MinH4VolumeInUsd(minH4VolumeInUsd).MaxH4VolumeInUsd(maxH4VolumeInUsd).MinH4PriceChangeRatio(minH4PriceChangeRatio).MaxH4PriceChangeRatio(maxH4PriceChangeRatio).MinH4Buys(minH4Buys).MaxH4Buys(maxH4Buys).MinH4Sells(minH4Sells).MaxH4Sells(maxH4Sells).MinH4Trades(minH4Trades).MaxH4Trades(maxH4Trades).MinH4Buyers(minH4Buyers).MaxH4Buyers(maxH4Buyers).MinH4Sellers(minH4Sellers).MaxH4Sellers(maxH4Sellers).MinH4BuyVolumeInUsd(minH4BuyVolumeInUsd).MaxH4BuyVolumeInUsd(maxH4BuyVolumeInUsd).MinH4SellVolumeInUsd(minH4SellVolumeInUsd).MaxH4SellVolumeInUsd(maxH4SellVolumeInUsd).MinH1VolumeInUsd(minH1VolumeInUsd).MaxH1VolumeInUsd(maxH1VolumeInUsd).MinH1PriceChangeRatio(minH1PriceChangeRatio).MaxH1PriceChangeRatio(maxH1PriceChangeRatio).MinH1Buys(minH1Buys).MaxH1Buys(maxH1Buys).MinH1Sells(minH1Sells).MaxH1Sells(maxH1Sells).MinH1Trades(minH1Trades).MaxH1Trades(maxH1Trades).MinH1Buyers(minH1Buyers).MaxH1Buyers(maxH1Buyers).MinH1Sellers(minH1Sellers).MaxH1Sellers(maxH1Sellers).MinH1BuyVolumeInUsd(minH1BuyVolumeInUsd).MaxH1BuyVolumeInUsd(maxH1BuyVolumeInUsd).MinH1SellVolumeInUsd(minH1SellVolumeInUsd).MaxH1SellVolumeInUsd(maxH1SellVolumeInUsd).MinM30VolumeInUsd(minM30VolumeInUsd).MaxM30VolumeInUsd(maxM30VolumeInUsd).MinM30PriceChangeRatio(minM30PriceChangeRatio).MaxM30PriceChangeRatio(maxM30PriceChangeRatio).MinM30Buys(minM30Buys).MaxM30Buys(maxM30Buys).MinM30Sells(minM30Sells).MaxM30Sells(maxM30Sells).MinM30Trades(minM30Trades).MaxM30Trades(maxM30Trades).MinM30Buyers(minM30Buyers).MaxM30Buyers(maxM30Buyers).MinM30Sellers(minM30Sellers).MaxM30Sellers(maxM30Sellers).MinM30BuyVolumeInUsd(minM30BuyVolumeInUsd).MaxM30BuyVolumeInUsd(maxM30BuyVolumeInUsd).MinM30SellVolumeInUsd(minM30SellVolumeInUsd).MaxM30SellVolumeInUsd(maxM30SellVolumeInUsd).MinM15VolumeInUsd(minM15VolumeInUsd).MaxM15VolumeInUsd(maxM15VolumeInUsd).MinM15PriceChangeRatio(minM15PriceChangeRatio).MaxM15PriceChangeRatio(maxM15PriceChangeRatio).MinM15Buys(minM15Buys).MaxM15Buys(maxM15Buys).MinM15Sells(minM15Sells).MaxM15Sells(maxM15Sells).MinM15Trades(minM15Trades).MaxM15Trades(maxM15Trades).MinM15Buyers(minM15Buyers).MaxM15Buyers(maxM15Buyers).MinM15Sellers(minM15Sellers).MaxM15Sellers(maxM15Sellers).MinM15BuyVolumeInUsd(minM15BuyVolumeInUsd).MaxM15BuyVolumeInUsd(maxM15BuyVolumeInUsd).MinM15SellVolumeInUsd(minM15SellVolumeInUsd).MaxM15SellVolumeInUsd(maxM15SellVolumeInUsd).MinM5VolumeInUsd(minM5VolumeInUsd).MaxM5VolumeInUsd(maxM5VolumeInUsd).MinM5PriceChangeRatio(minM5PriceChangeRatio).MaxM5PriceChangeRatio(maxM5PriceChangeRatio).MinM5Buys(minM5Buys).MaxM5Buys(maxM5Buys).MinM5Sells(minM5Sells).MaxM5Sells(maxM5Sells).MinM5Trades(minM5Trades).MaxM5Trades(maxM5Trades).MinM5Buyers(minM5Buyers).MaxM5Buyers(maxM5Buyers).MinM5Sellers(minM5Sellers).MaxM5Sellers(maxM5Sellers).MinM5BuyVolumeInUsd(minM5BuyVolumeInUsd).MaxM5BuyVolumeInUsd(maxM5BuyVolumeInUsd).MinM5SellVolumeInUsd(minM5SellVolumeInUsd).MaxM5SellVolumeInUsd(maxM5SellVolumeInUsd).MinM1VolumeInUsd(minM1VolumeInUsd).MaxM1VolumeInUsd(maxM1VolumeInUsd).MinM1PriceChangeRatio(minM1PriceChangeRatio).MaxM1PriceChangeRatio(maxM1PriceChangeRatio).MinM1Buys(minM1Buys).MaxM1Buys(maxM1Buys).MinM1Sells(minM1Sells).MaxM1Sells(maxM1Sells).MinM1Trades(minM1Trades).MaxM1Trades(maxM1Trades).MinM1Buyers(minM1Buyers).MaxM1Buyers(maxM1Buyers).MinM1Sellers(minM1Sellers).MaxM1Sellers(maxM1Sellers).MinM1BuyVolumeInUsd(minM1BuyVolumeInUsd).MaxM1BuyVolumeInUsd(maxM1BuyVolumeInUsd).MinM1SellVolumeInUsd(minM1SellVolumeInUsd).MaxM1SellVolumeInUsd(maxM1SellVolumeInUsd).Execute()

CONTROLLER.TOKEN.GET_TOKEN_LIST.SUMMARY



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
	sort := "sort_example" // string | DTO.TOKEN.SEARCH.SORT_DIRECTION (optional) (default to "desc")
	sortBy := "h24VolumeInUsd" // string | DTO.TOKEN.LIST.QUERY.SORT_BY (optional)
	minH24VolumeInUsd := "minH24VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_VOLUME_IN_USD (optional)
	maxH24VolumeInUsd := "maxH24VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_VOLUME_IN_USD (optional)
	minH24PriceChangeRatio := "minH24PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_PRICE_CHANGE_RATIO (optional)
	maxH24PriceChangeRatio := "maxH24PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_PRICE_CHANGE_RATIO (optional)
	minH24Buys := "minH24Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_BUYS (optional)
	maxH24Buys := "maxH24Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_BUYS (optional)
	minH24Sells := "minH24Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_SELLS (optional)
	maxH24Sells := "maxH24Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_SELLS (optional)
	minH24Trades := "minH24Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_TRADES (optional)
	maxH24Trades := "maxH24Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_TRADES (optional)
	minH24Buyers := "minH24Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_BUYERS (optional)
	maxH24Buyers := "maxH24Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_BUYERS (optional)
	minH24Sellers := "minH24Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_SELLERS (optional)
	maxH24Sellers := "maxH24Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_SELLERS (optional)
	minH24BuyVolumeInUsd := "minH24BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_BUY_VOLUME_IN_USD (optional)
	maxH24BuyVolumeInUsd := "maxH24BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_BUY_VOLUME_IN_USD (optional)
	minH24SellVolumeInUsd := "minH24SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H24_SELL_VOLUME_IN_USD (optional)
	maxH24SellVolumeInUsd := "maxH24SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H24_SELL_VOLUME_IN_USD (optional)
	minH4VolumeInUsd := "minH4VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_VOLUME_IN_USD (optional)
	maxH4VolumeInUsd := "maxH4VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_VOLUME_IN_USD (optional)
	minH4PriceChangeRatio := "minH4PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_PRICE_CHANGE_RATIO (optional)
	maxH4PriceChangeRatio := "maxH4PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_PRICE_CHANGE_RATIO (optional)
	minH4Buys := "minH4Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_BUYS (optional)
	maxH4Buys := "maxH4Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_BUYS (optional)
	minH4Sells := "minH4Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_SELLS (optional)
	maxH4Sells := "maxH4Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_SELLS (optional)
	minH4Trades := "minH4Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_TRADES (optional)
	maxH4Trades := "maxH4Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_TRADES (optional)
	minH4Buyers := "minH4Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_BUYERS (optional)
	maxH4Buyers := "maxH4Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_BUYERS (optional)
	minH4Sellers := "minH4Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_SELLERS (optional)
	maxH4Sellers := "maxH4Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_SELLERS (optional)
	minH4BuyVolumeInUsd := "minH4BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_BUY_VOLUME_IN_USD (optional)
	maxH4BuyVolumeInUsd := "maxH4BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_BUY_VOLUME_IN_USD (optional)
	minH4SellVolumeInUsd := "minH4SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H4_SELL_VOLUME_IN_USD (optional)
	maxH4SellVolumeInUsd := "maxH4SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H4_SELL_VOLUME_IN_USD (optional)
	minH1VolumeInUsd := "minH1VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_VOLUME_IN_USD (optional)
	maxH1VolumeInUsd := "maxH1VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_VOLUME_IN_USD (optional)
	minH1PriceChangeRatio := "minH1PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_PRICE_CHANGE_RATIO (optional)
	maxH1PriceChangeRatio := "maxH1PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_PRICE_CHANGE_RATIO (optional)
	minH1Buys := "minH1Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_BUYS (optional)
	maxH1Buys := "maxH1Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_BUYS (optional)
	minH1Sells := "minH1Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_SELLS (optional)
	maxH1Sells := "maxH1Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_SELLS (optional)
	minH1Trades := "minH1Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_TRADES (optional)
	maxH1Trades := "maxH1Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_TRADES (optional)
	minH1Buyers := "minH1Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_BUYERS (optional)
	maxH1Buyers := "maxH1Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_BUYERS (optional)
	minH1Sellers := "minH1Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_SELLERS (optional)
	maxH1Sellers := "maxH1Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_SELLERS (optional)
	minH1BuyVolumeInUsd := "minH1BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_BUY_VOLUME_IN_USD (optional)
	maxH1BuyVolumeInUsd := "maxH1BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_BUY_VOLUME_IN_USD (optional)
	minH1SellVolumeInUsd := "minH1SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_H1_SELL_VOLUME_IN_USD (optional)
	maxH1SellVolumeInUsd := "maxH1SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_H1_SELL_VOLUME_IN_USD (optional)
	minM30VolumeInUsd := "minM30VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_VOLUME_IN_USD (optional)
	maxM30VolumeInUsd := "maxM30VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_VOLUME_IN_USD (optional)
	minM30PriceChangeRatio := "minM30PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_PRICE_CHANGE_RATIO (optional)
	maxM30PriceChangeRatio := "maxM30PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_PRICE_CHANGE_RATIO (optional)
	minM30Buys := "minM30Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_BUYS (optional)
	maxM30Buys := "maxM30Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_BUYS (optional)
	minM30Sells := "minM30Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_SELLS (optional)
	maxM30Sells := "maxM30Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_SELLS (optional)
	minM30Trades := "minM30Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_TRADES (optional)
	maxM30Trades := "maxM30Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_TRADES (optional)
	minM30Buyers := "minM30Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_BUYERS (optional)
	maxM30Buyers := "maxM30Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_BUYERS (optional)
	minM30Sellers := "minM30Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_SELLERS (optional)
	maxM30Sellers := "maxM30Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_SELLERS (optional)
	minM30BuyVolumeInUsd := "minM30BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_BUY_VOLUME_IN_USD (optional)
	maxM30BuyVolumeInUsd := "maxM30BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_BUY_VOLUME_IN_USD (optional)
	minM30SellVolumeInUsd := "minM30SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M30_SELL_VOLUME_IN_USD (optional)
	maxM30SellVolumeInUsd := "maxM30SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M30_SELL_VOLUME_IN_USD (optional)
	minM15VolumeInUsd := "minM15VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_VOLUME_IN_USD (optional)
	maxM15VolumeInUsd := "maxM15VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_VOLUME_IN_USD (optional)
	minM15PriceChangeRatio := "minM15PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_PRICE_CHANGE_RATIO (optional)
	maxM15PriceChangeRatio := "maxM15PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_PRICE_CHANGE_RATIO (optional)
	minM15Buys := "minM15Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_BUYS (optional)
	maxM15Buys := "maxM15Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_BUYS (optional)
	minM15Sells := "minM15Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_SELLS (optional)
	maxM15Sells := "maxM15Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_SELLS (optional)
	minM15Trades := "minM15Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_TRADES (optional)
	maxM15Trades := "maxM15Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_TRADES (optional)
	minM15Buyers := "minM15Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_BUYERS (optional)
	maxM15Buyers := "maxM15Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_BUYERS (optional)
	minM15Sellers := "minM15Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_SELLERS (optional)
	maxM15Sellers := "maxM15Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_SELLERS (optional)
	minM15BuyVolumeInUsd := "minM15BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_BUY_VOLUME_IN_USD (optional)
	maxM15BuyVolumeInUsd := "maxM15BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_BUY_VOLUME_IN_USD (optional)
	minM15SellVolumeInUsd := "minM15SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M15_SELL_VOLUME_IN_USD (optional)
	maxM15SellVolumeInUsd := "maxM15SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M15_SELL_VOLUME_IN_USD (optional)
	minM5VolumeInUsd := "minM5VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_VOLUME_IN_USD (optional)
	maxM5VolumeInUsd := "maxM5VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_VOLUME_IN_USD (optional)
	minM5PriceChangeRatio := "minM5PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_PRICE_CHANGE_RATIO (optional)
	maxM5PriceChangeRatio := "maxM5PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_PRICE_CHANGE_RATIO (optional)
	minM5Buys := "minM5Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_BUYS (optional)
	maxM5Buys := "maxM5Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_BUYS (optional)
	minM5Sells := "minM5Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_SELLS (optional)
	maxM5Sells := "maxM5Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_SELLS (optional)
	minM5Trades := "minM5Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_TRADES (optional)
	maxM5Trades := "maxM5Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_TRADES (optional)
	minM5Buyers := "minM5Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_BUYERS (optional)
	maxM5Buyers := "maxM5Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_BUYERS (optional)
	minM5Sellers := "minM5Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_SELLERS (optional)
	maxM5Sellers := "maxM5Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_SELLERS (optional)
	minM5BuyVolumeInUsd := "minM5BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_BUY_VOLUME_IN_USD (optional)
	maxM5BuyVolumeInUsd := "maxM5BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_BUY_VOLUME_IN_USD (optional)
	minM5SellVolumeInUsd := "minM5SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M5_SELL_VOLUME_IN_USD (optional)
	maxM5SellVolumeInUsd := "maxM5SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M5_SELL_VOLUME_IN_USD (optional)
	minM1VolumeInUsd := "minM1VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_VOLUME_IN_USD (optional)
	maxM1VolumeInUsd := "maxM1VolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_VOLUME_IN_USD (optional)
	minM1PriceChangeRatio := "minM1PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_PRICE_CHANGE_RATIO (optional)
	maxM1PriceChangeRatio := "maxM1PriceChangeRatio_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_PRICE_CHANGE_RATIO (optional)
	minM1Buys := "minM1Buys_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_BUYS (optional)
	maxM1Buys := "maxM1Buys_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_BUYS (optional)
	minM1Sells := "minM1Sells_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_SELLS (optional)
	maxM1Sells := "maxM1Sells_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_SELLS (optional)
	minM1Trades := "minM1Trades_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_TRADES (optional)
	maxM1Trades := "maxM1Trades_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_TRADES (optional)
	minM1Buyers := "minM1Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_BUYERS (optional)
	maxM1Buyers := "maxM1Buyers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_BUYERS (optional)
	minM1Sellers := "minM1Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_SELLERS (optional)
	maxM1Sellers := "maxM1Sellers_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_SELLERS (optional)
	minM1BuyVolumeInUsd := "minM1BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_BUY_VOLUME_IN_USD (optional)
	maxM1BuyVolumeInUsd := "maxM1BuyVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_BUY_VOLUME_IN_USD (optional)
	minM1SellVolumeInUsd := "minM1SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MIN_M1_SELL_VOLUME_IN_USD (optional)
	maxM1SellVolumeInUsd := "maxM1SellVolumeInUsd_example" // string | DTO.TOKEN.LIST.QUERY.MAX_M1_SELL_VOLUME_IN_USD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.ListToken(context.Background(), chain).Cursor(cursor).Limit(limit).Direction(direction).Sort(sort).SortBy(sortBy).MinH24VolumeInUsd(minH24VolumeInUsd).MaxH24VolumeInUsd(maxH24VolumeInUsd).MinH24PriceChangeRatio(minH24PriceChangeRatio).MaxH24PriceChangeRatio(maxH24PriceChangeRatio).MinH24Buys(minH24Buys).MaxH24Buys(maxH24Buys).MinH24Sells(minH24Sells).MaxH24Sells(maxH24Sells).MinH24Trades(minH24Trades).MaxH24Trades(maxH24Trades).MinH24Buyers(minH24Buyers).MaxH24Buyers(maxH24Buyers).MinH24Sellers(minH24Sellers).MaxH24Sellers(maxH24Sellers).MinH24BuyVolumeInUsd(minH24BuyVolumeInUsd).MaxH24BuyVolumeInUsd(maxH24BuyVolumeInUsd).MinH24SellVolumeInUsd(minH24SellVolumeInUsd).MaxH24SellVolumeInUsd(maxH24SellVolumeInUsd).MinH4VolumeInUsd(minH4VolumeInUsd).MaxH4VolumeInUsd(maxH4VolumeInUsd).MinH4PriceChangeRatio(minH4PriceChangeRatio).MaxH4PriceChangeRatio(maxH4PriceChangeRatio).MinH4Buys(minH4Buys).MaxH4Buys(maxH4Buys).MinH4Sells(minH4Sells).MaxH4Sells(maxH4Sells).MinH4Trades(minH4Trades).MaxH4Trades(maxH4Trades).MinH4Buyers(minH4Buyers).MaxH4Buyers(maxH4Buyers).MinH4Sellers(minH4Sellers).MaxH4Sellers(maxH4Sellers).MinH4BuyVolumeInUsd(minH4BuyVolumeInUsd).MaxH4BuyVolumeInUsd(maxH4BuyVolumeInUsd).MinH4SellVolumeInUsd(minH4SellVolumeInUsd).MaxH4SellVolumeInUsd(maxH4SellVolumeInUsd).MinH1VolumeInUsd(minH1VolumeInUsd).MaxH1VolumeInUsd(maxH1VolumeInUsd).MinH1PriceChangeRatio(minH1PriceChangeRatio).MaxH1PriceChangeRatio(maxH1PriceChangeRatio).MinH1Buys(minH1Buys).MaxH1Buys(maxH1Buys).MinH1Sells(minH1Sells).MaxH1Sells(maxH1Sells).MinH1Trades(minH1Trades).MaxH1Trades(maxH1Trades).MinH1Buyers(minH1Buyers).MaxH1Buyers(maxH1Buyers).MinH1Sellers(minH1Sellers).MaxH1Sellers(maxH1Sellers).MinH1BuyVolumeInUsd(minH1BuyVolumeInUsd).MaxH1BuyVolumeInUsd(maxH1BuyVolumeInUsd).MinH1SellVolumeInUsd(minH1SellVolumeInUsd).MaxH1SellVolumeInUsd(maxH1SellVolumeInUsd).MinM30VolumeInUsd(minM30VolumeInUsd).MaxM30VolumeInUsd(maxM30VolumeInUsd).MinM30PriceChangeRatio(minM30PriceChangeRatio).MaxM30PriceChangeRatio(maxM30PriceChangeRatio).MinM30Buys(minM30Buys).MaxM30Buys(maxM30Buys).MinM30Sells(minM30Sells).MaxM30Sells(maxM30Sells).MinM30Trades(minM30Trades).MaxM30Trades(maxM30Trades).MinM30Buyers(minM30Buyers).MaxM30Buyers(maxM30Buyers).MinM30Sellers(minM30Sellers).MaxM30Sellers(maxM30Sellers).MinM30BuyVolumeInUsd(minM30BuyVolumeInUsd).MaxM30BuyVolumeInUsd(maxM30BuyVolumeInUsd).MinM30SellVolumeInUsd(minM30SellVolumeInUsd).MaxM30SellVolumeInUsd(maxM30SellVolumeInUsd).MinM15VolumeInUsd(minM15VolumeInUsd).MaxM15VolumeInUsd(maxM15VolumeInUsd).MinM15PriceChangeRatio(minM15PriceChangeRatio).MaxM15PriceChangeRatio(maxM15PriceChangeRatio).MinM15Buys(minM15Buys).MaxM15Buys(maxM15Buys).MinM15Sells(minM15Sells).MaxM15Sells(maxM15Sells).MinM15Trades(minM15Trades).MaxM15Trades(maxM15Trades).MinM15Buyers(minM15Buyers).MaxM15Buyers(maxM15Buyers).MinM15Sellers(minM15Sellers).MaxM15Sellers(maxM15Sellers).MinM15BuyVolumeInUsd(minM15BuyVolumeInUsd).MaxM15BuyVolumeInUsd(maxM15BuyVolumeInUsd).MinM15SellVolumeInUsd(minM15SellVolumeInUsd).MaxM15SellVolumeInUsd(maxM15SellVolumeInUsd).MinM5VolumeInUsd(minM5VolumeInUsd).MaxM5VolumeInUsd(maxM5VolumeInUsd).MinM5PriceChangeRatio(minM5PriceChangeRatio).MaxM5PriceChangeRatio(maxM5PriceChangeRatio).MinM5Buys(minM5Buys).MaxM5Buys(maxM5Buys).MinM5Sells(minM5Sells).MaxM5Sells(maxM5Sells).MinM5Trades(minM5Trades).MaxM5Trades(maxM5Trades).MinM5Buyers(minM5Buyers).MaxM5Buyers(maxM5Buyers).MinM5Sellers(minM5Sellers).MaxM5Sellers(maxM5Sellers).MinM5BuyVolumeInUsd(minM5BuyVolumeInUsd).MaxM5BuyVolumeInUsd(maxM5BuyVolumeInUsd).MinM5SellVolumeInUsd(minM5SellVolumeInUsd).MaxM5SellVolumeInUsd(maxM5SellVolumeInUsd).MinM1VolumeInUsd(minM1VolumeInUsd).MaxM1VolumeInUsd(maxM1VolumeInUsd).MinM1PriceChangeRatio(minM1PriceChangeRatio).MaxM1PriceChangeRatio(maxM1PriceChangeRatio).MinM1Buys(minM1Buys).MaxM1Buys(maxM1Buys).MinM1Sells(minM1Sells).MaxM1Sells(maxM1Sells).MinM1Trades(minM1Trades).MaxM1Trades(maxM1Trades).MinM1Buyers(minM1Buyers).MaxM1Buyers(maxM1Buyers).MinM1Sellers(minM1Sellers).MaxM1Sellers(maxM1Sellers).MinM1BuyVolumeInUsd(minM1BuyVolumeInUsd).MaxM1BuyVolumeInUsd(maxM1BuyVolumeInUsd).MinM1SellVolumeInUsd(minM1SellVolumeInUsd).MaxM1SellVolumeInUsd(maxM1SellVolumeInUsd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.ListToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListToken`: TokenListPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.ListToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | DTO.PAGE.CURSOR.DESCRIPTION | 
 **limit** | **float32** | DTO.PAGE.LIMIT | [default to 20]
 **direction** | **string** | DTO.PAGE.DIRECTION | [default to &quot;next&quot;]
 **sort** | **string** | DTO.TOKEN.SEARCH.SORT_DIRECTION | [default to &quot;desc&quot;]
 **sortBy** | **string** | DTO.TOKEN.LIST.QUERY.SORT_BY | 
 **minH24VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_VOLUME_IN_USD | 
 **maxH24VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_VOLUME_IN_USD | 
 **minH24PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_PRICE_CHANGE_RATIO | 
 **maxH24PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_PRICE_CHANGE_RATIO | 
 **minH24Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_BUYS | 
 **maxH24Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_BUYS | 
 **minH24Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_SELLS | 
 **maxH24Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_SELLS | 
 **minH24Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_TRADES | 
 **maxH24Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_TRADES | 
 **minH24Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_BUYERS | 
 **maxH24Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_BUYERS | 
 **minH24Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_SELLERS | 
 **maxH24Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_SELLERS | 
 **minH24BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_BUY_VOLUME_IN_USD | 
 **maxH24BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_BUY_VOLUME_IN_USD | 
 **minH24SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H24_SELL_VOLUME_IN_USD | 
 **maxH24SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H24_SELL_VOLUME_IN_USD | 
 **minH4VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_VOLUME_IN_USD | 
 **maxH4VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_VOLUME_IN_USD | 
 **minH4PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_PRICE_CHANGE_RATIO | 
 **maxH4PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_PRICE_CHANGE_RATIO | 
 **minH4Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_BUYS | 
 **maxH4Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_BUYS | 
 **minH4Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_SELLS | 
 **maxH4Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_SELLS | 
 **minH4Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_TRADES | 
 **maxH4Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_TRADES | 
 **minH4Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_BUYERS | 
 **maxH4Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_BUYERS | 
 **minH4Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_SELLERS | 
 **maxH4Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_SELLERS | 
 **minH4BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_BUY_VOLUME_IN_USD | 
 **maxH4BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_BUY_VOLUME_IN_USD | 
 **minH4SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H4_SELL_VOLUME_IN_USD | 
 **maxH4SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H4_SELL_VOLUME_IN_USD | 
 **minH1VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_VOLUME_IN_USD | 
 **maxH1VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_VOLUME_IN_USD | 
 **minH1PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_PRICE_CHANGE_RATIO | 
 **maxH1PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_PRICE_CHANGE_RATIO | 
 **minH1Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_BUYS | 
 **maxH1Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_BUYS | 
 **minH1Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_SELLS | 
 **maxH1Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_SELLS | 
 **minH1Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_TRADES | 
 **maxH1Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_TRADES | 
 **minH1Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_BUYERS | 
 **maxH1Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_BUYERS | 
 **minH1Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_SELLERS | 
 **maxH1Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_SELLERS | 
 **minH1BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_BUY_VOLUME_IN_USD | 
 **maxH1BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_BUY_VOLUME_IN_USD | 
 **minH1SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_H1_SELL_VOLUME_IN_USD | 
 **maxH1SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_H1_SELL_VOLUME_IN_USD | 
 **minM30VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_VOLUME_IN_USD | 
 **maxM30VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_VOLUME_IN_USD | 
 **minM30PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_PRICE_CHANGE_RATIO | 
 **maxM30PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_PRICE_CHANGE_RATIO | 
 **minM30Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_BUYS | 
 **maxM30Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_BUYS | 
 **minM30Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_SELLS | 
 **maxM30Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_SELLS | 
 **minM30Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_TRADES | 
 **maxM30Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_TRADES | 
 **minM30Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_BUYERS | 
 **maxM30Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_BUYERS | 
 **minM30Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_SELLERS | 
 **maxM30Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_SELLERS | 
 **minM30BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_BUY_VOLUME_IN_USD | 
 **maxM30BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_BUY_VOLUME_IN_USD | 
 **minM30SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M30_SELL_VOLUME_IN_USD | 
 **maxM30SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M30_SELL_VOLUME_IN_USD | 
 **minM15VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_VOLUME_IN_USD | 
 **maxM15VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_VOLUME_IN_USD | 
 **minM15PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_PRICE_CHANGE_RATIO | 
 **maxM15PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_PRICE_CHANGE_RATIO | 
 **minM15Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_BUYS | 
 **maxM15Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_BUYS | 
 **minM15Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_SELLS | 
 **maxM15Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_SELLS | 
 **minM15Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_TRADES | 
 **maxM15Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_TRADES | 
 **minM15Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_BUYERS | 
 **maxM15Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_BUYERS | 
 **minM15Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_SELLERS | 
 **maxM15Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_SELLERS | 
 **minM15BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_BUY_VOLUME_IN_USD | 
 **maxM15BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_BUY_VOLUME_IN_USD | 
 **minM15SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M15_SELL_VOLUME_IN_USD | 
 **maxM15SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M15_SELL_VOLUME_IN_USD | 
 **minM5VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_VOLUME_IN_USD | 
 **maxM5VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_VOLUME_IN_USD | 
 **minM5PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_PRICE_CHANGE_RATIO | 
 **maxM5PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_PRICE_CHANGE_RATIO | 
 **minM5Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_BUYS | 
 **maxM5Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_BUYS | 
 **minM5Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_SELLS | 
 **maxM5Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_SELLS | 
 **minM5Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_TRADES | 
 **maxM5Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_TRADES | 
 **minM5Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_BUYERS | 
 **maxM5Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_BUYERS | 
 **minM5Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_SELLERS | 
 **maxM5Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_SELLERS | 
 **minM5BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_BUY_VOLUME_IN_USD | 
 **maxM5BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_BUY_VOLUME_IN_USD | 
 **minM5SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M5_SELL_VOLUME_IN_USD | 
 **maxM5SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M5_SELL_VOLUME_IN_USD | 
 **minM1VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_VOLUME_IN_USD | 
 **maxM1VolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_VOLUME_IN_USD | 
 **minM1PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_PRICE_CHANGE_RATIO | 
 **maxM1PriceChangeRatio** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_PRICE_CHANGE_RATIO | 
 **minM1Buys** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_BUYS | 
 **maxM1Buys** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_BUYS | 
 **minM1Sells** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_SELLS | 
 **maxM1Sells** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_SELLS | 
 **minM1Trades** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_TRADES | 
 **maxM1Trades** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_TRADES | 
 **minM1Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_BUYERS | 
 **maxM1Buyers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_BUYERS | 
 **minM1Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_SELLERS | 
 **maxM1Sellers** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_SELLERS | 
 **minM1BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_BUY_VOLUME_IN_USD | 
 **maxM1BuyVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_BUY_VOLUME_IN_USD | 
 **minM1SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MIN_M1_SELL_VOLUME_IN_USD | 
 **maxM1SellVolumeInUsd** | **string** | DTO.TOKEN.LIST.QUERY.MAX_M1_SELL_VOLUME_IN_USD | 

### Return type

[**TokenListPage**](TokenListPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> TokenPage Search(ctx).Chains(chains).Q(q).Limit(limit).Sort(sort).Protocols(protocols).Cursor(cursor).SortBy(sortBy).Execute()

CONTROLLER.TOKEN.SEARCH.SUMMARY



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
	chains := []string{"Inner_example"} // []string | DTO.TOKEN.SEARCH.CHAINS (optional)
	q := "USDC" // string | DTO.TOKEN.SEARCH.QUERY (optional)
	limit := float32(8.14) // float32 | DTO.TOKEN.SEARCH.LIMIT (optional) (default to 20)
	sort := "sort_example" // string | DTO.TOKEN.SEARCH.SORT_DIRECTION (optional) (default to "desc")
	protocols := []string{"Inner_example"} // []string | DTO.TOKEN.SEARCH.PROTOCOLS (optional)
	cursor := "cursor_example" // string | DTO.PAGE.CURSOR (optional)
	sortBy := "priceInUsd" // string | DTO.TOKEN.SEARCH.SORT_BY (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.Search(context.Background()).Chains(chains).Q(q).Limit(limit).Sort(sort).Protocols(protocols).Cursor(cursor).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: TokenPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chains** | **[]string** | DTO.TOKEN.SEARCH.CHAINS | 
 **q** | **string** | DTO.TOKEN.SEARCH.QUERY | 
 **limit** | **float32** | DTO.TOKEN.SEARCH.LIMIT | [default to 20]
 **sort** | **string** | DTO.TOKEN.SEARCH.SORT_DIRECTION | [default to &quot;desc&quot;]
 **protocols** | **[]string** | DTO.TOKEN.SEARCH.PROTOCOLS | 
 **cursor** | **string** | DTO.PAGE.CURSOR | 
 **sortBy** | **string** | DTO.TOKEN.SEARCH.SORT_BY | 

### Return type

[**TokenPage**](TokenPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

