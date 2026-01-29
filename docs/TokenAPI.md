# \TokenAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCandles**](TokenAPI.md#GetCandles) | **Get** /v1/token/{chain}/{tokenAddress}/candles | Token - Candles
[**GetCreation**](TokenAPI.md#GetCreation) | **Get** /v1/token/{chain}/{tokenAddress}/creation | Token - Creation
[**GetDevTokens**](TokenAPI.md#GetDevTokens) | **Get** /v1/token/{chain}/dev/{devAddress}/tokens | Token - Get Dev Tokens
[**GetHolders**](TokenAPI.md#GetHolders) | **Get** /v1/token/{chain}/{tokenAddress}/holders | Token - Holders
[**GetHoldersMulti**](TokenAPI.md#GetHoldersMulti) | **Get** /v1/token/{chain}/{tokenAddress}/holders/multi | Token - Holders (Multi)
[**GetMarketData**](TokenAPI.md#GetMarketData) | **Get** /v1/token/{chain}/{tokenAddress}/marketData | Token - Market Data
[**GetMarketDataMulti**](TokenAPI.md#GetMarketDataMulti) | **Get** /v1/token/{chain}/marketData/multi | Token - Market Data (Multi)
[**GetMetadata**](TokenAPI.md#GetMetadata) | **Get** /v1/token/{chain}/{tokenAddress}/metadata | Token - Metadata
[**GetMetadataMulti**](TokenAPI.md#GetMetadataMulti) | **Get** /v1/token/{chain}/metadata/multi | Token - Metadata (Multi)
[**GetMintAndBurn**](TokenAPI.md#GetMintAndBurn) | **Get** /v1/token/{chain}/{tokenAddress}/mintAndBurn | Token - Mint and Burn
[**GetPools**](TokenAPI.md#GetPools) | **Get** /v1/token/{chain}/{tokenAddress}/pools | Token - Liquidity
[**GetPriceByTime**](TokenAPI.md#GetPriceByTime) | **Get** /v1/token/{chain}/{tokenAddress}/price | Token - Price by Time
[**GetPrices**](TokenAPI.md#GetPrices) | **Get** /v1/token/{chain}/{tokenAddress}/prices | Token - Prices
[**GetSecurity**](TokenAPI.md#GetSecurity) | **Get** /v1/token/{chain}/{tokenAddress}/security | Token - Security
[**GetStats**](TokenAPI.md#GetStats) | **Get** /v1/token/{chain}/{tokenAddress}/stats | Token - Stats
[**GetStatsMulti**](TokenAPI.md#GetStatsMulti) | **Get** /v1/token/{chain}/stats/multi | Token - Stats (Multi)
[**GetToken**](TokenAPI.md#GetToken) | **Get** /v1/token/{chain}/{tokenAddress} | Token - Detail
[**GetTokenLiquiditySnapshots**](TokenAPI.md#GetTokenLiquiditySnapshots) | **Get** /v1/token/{chain}/{tokenAddress}/liquiditySnapshots | Token - Liquidity Snapshots
[**GetTokenTraders**](TokenAPI.md#GetTokenTraders) | **Get** /v1/token/{chain}/{tokenAddress}/traders/{tag} | Token - Get Token Traders
[**GetTokens**](TokenAPI.md#GetTokens) | **Get** /v1/token/{chain}/multi | Token - Detail (Multi)
[**GetTopHolders**](TokenAPI.md#GetTopHolders) | **Get** /v1/token/{chain}/{tokenAddress}/topHolders | Token - Top Holders
[**ListToken**](TokenAPI.md#ListToken) | **Get** /v1/token/{chain}/list | Token - List (Filtered)
[**Search**](TokenAPI.md#Search) | **Get** /v1/token/search | Token - Search



## GetCandles

> []Candle GetCandles(ctx, chain, tokenAddress).Resolution(resolution).From(from).To(to).Limit(limit).Execute()

Token - Candles



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	resolution := openapiclient.Resolution("1s") // Resolution | Time resolution for candle data. Note: 1s resolution data is only kept for the last 1 hour; 15s and 30s resolution data is kept for the last 6 hours; 1m resolution data is kept for the last 12 hours; data for other resolutions is stored permanently
	from := int64(1741647950000) // int64 | Start timestamp (Unix epoch in milliseconds) (optional)
	to := int64(1741700970000) // int64 | End timestamp (Unix epoch in milliseconds) (optional)
	limit := int64(100) // int64 | Number of results per page (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCandlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resolution** | [**Resolution**](Resolution.md) | Time resolution for candle data. Note: 1s resolution data is only kept for the last 1 hour; 15s and 30s resolution data is kept for the last 6 hours; 1m resolution data is kept for the last 12 hours; data for other resolutions is stored permanently | 
 **from** | **int64** | Start timestamp (Unix epoch in milliseconds) | 
 **to** | **int64** | End timestamp (Unix epoch in milliseconds) | 
 **limit** | **int64** | Number of results per page | 

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

Token - Creation



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - Get Dev Tokens



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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
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

Token - Holders



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := float32(8.14) // float32 | Number of results per page (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]

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

Token - Holders (Multi)



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	walletAddresses := "MJKqp326RZCHnAAbew9MDdui3iCKWco7fsK9sVuZTX2,2RH6rUTPBJ9rUDPpuV9b8z1YL56k1tYU6Uk5ZoaEFFSK" // string | A list of wallet addresses in string separated by commas (,)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldersMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **walletAddresses** | **string** | A list of wallet addresses in string separated by commas (,) | 

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

Token - Market Data



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - Market Data (Multi)



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | A list of token addresses in string separated by commas (,)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketDataMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | A list of token addresses in string separated by commas (,) | 

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

Token - Metadata



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - Metadata (Multi)



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | A list of token addresses in string separated by commas (,)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | A list of token addresses in string separated by commas (,) | 

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

Token - Mint and Burn



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := float32(8.14) // float32 | Number of results per page (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")
	type_ := "all" // string | Type of operation to filter (all, mint, or burn) (optional) (default to "all")

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintAndBurnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]
 **type_** | **string** | Type of operation to filter (all, mint, or burn) | [default to &quot;all&quot;]

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

> []DexPoolDTO GetPools(ctx, chain, tokenAddress).SortBy(sortBy).SortDirection(sortDirection).MinTvlInSol(minTvlInSol).MaxTvlInSol(maxTvlInSol).MinTvlInUsd(minTvlInUsd).MaxTvlInUsd(maxTvlInUsd).Execute()

Token - Liquidity



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	sortBy := "tvlInUsd" // string | Sort field (tvlInSol or tvlInUsd) (optional)
	sortDirection := "sortDirection_example" // string | Sort direction (asc or desc) (optional) (default to "desc")
	minTvlInSol := "100" // string | Minimum TVL in SOL (optional)
	maxTvlInSol := "10000" // string | Maximum TVL in SOL (optional)
	minTvlInUsd := "1000" // string | Minimum TVL in USD (optional)
	maxTvlInUsd := "100000" // string | Maximum TVL in USD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetPools(context.Background(), chain, tokenAddress).SortBy(sortBy).SortDirection(sortDirection).MinTvlInSol(minTvlInSol).MaxTvlInSol(maxTvlInSol).MinTvlInUsd(minTvlInUsd).MaxTvlInUsd(maxTvlInUsd).Execute()
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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | Sort field (tvlInSol or tvlInUsd) | 
 **sortDirection** | **string** | Sort direction (asc or desc) | [default to &quot;desc&quot;]
 **minTvlInSol** | **string** | Minimum TVL in SOL | 
 **maxTvlInSol** | **string** | Maximum TVL in SOL | 
 **minTvlInUsd** | **string** | Minimum TVL in USD | 
 **maxTvlInUsd** | **string** | Maximum TVL in USD | 

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

Token - Price by Time



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	timestamp := int64(1754055151) // int64 | Timestamp for price query (Unix epoch in seconds)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timestamp** | **int64** | Timestamp for price query (Unix epoch in seconds) | 

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

Token - Prices



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	cursor := "cursor_example" // string | Pagination cursor (optional)
	limit := float32(8.14) // float32 | Number of results per page (optional) (default to 20)
	direction := "direction_example" // string | Pagination direction (next or prev) (optional) (default to "next")

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]

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

Token - Security



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - Stats



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - Stats (Multi)



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | A list of token addresses in string separated by commas (,)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsMultiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | A list of token addresses in string separated by commas (,) | 

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

Token - Detail



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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


## GetTokenLiquiditySnapshots

> TokenLiquiditySnapshotPage GetTokenLiquiditySnapshots(ctx, chain, tokenAddress).Time(time).Cursor(cursor).Limit(limit).Execute()

Token - Liquidity Snapshots



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token
	time := int64(1705312800) // int64 | Target Unix timestamp (seconds) for snapshot query. Returns the nearest snapshot before or at this time. (optional)
	cursor := "eyJpZCI6MTAwfQ==" // string | Pagination cursor (optional)
	limit := int32(20) // int32 | Number of results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.GetTokenLiquiditySnapshots(context.Background(), chain, tokenAddress).Time(time).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokenLiquiditySnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenLiquiditySnapshots`: TokenLiquiditySnapshotPage
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokenLiquiditySnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenLiquiditySnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **int64** | Target Unix timestamp (seconds) for snapshot query. Returns the nearest snapshot before or at this time. | 
 **cursor** | **string** | Pagination cursor | 
 **limit** | **int32** | Number of results per page | [default to 20]

### Return type

[**TokenLiquiditySnapshotPage**](TokenLiquiditySnapshotPage.md)

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

Token - Get Token Traders



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
	tokenAddress := "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | An address of a token
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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 
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

Token - Detail (Multi)



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
	tokenAddresses := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN,EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // string | A list of token addresses in string separated by commas (,)
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	filterBy := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenAddresses** | **string** | A list of token addresses in string separated by commas (,) | 
 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **filterBy** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 

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

Token - Top Holders



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
	tokenAddress := "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN" // string | An address of a token

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**tokenAddress** | **string** | An address of a token | 

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

Token - List (Filtered)



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
	sort := "sort_example" // string | Sort direction (optional) (default to "desc")
	sortBy := "h24VolumeInUsd" // string | Sort by field (optional)
	minH24VolumeInUsd := "minH24VolumeInUsd_example" // string | Minimum 24h volume in USD (optional)
	maxH24VolumeInUsd := "maxH24VolumeInUsd_example" // string | Maximum 24h volume in USD (optional)
	minH24PriceChangeRatio := "minH24PriceChangeRatio_example" // string | Minimum 24h price change ratio (optional)
	maxH24PriceChangeRatio := "maxH24PriceChangeRatio_example" // string | Maximum 24h price change ratio (optional)
	minH24Buys := "minH24Buys_example" // string | Minimum 24h buys (optional)
	maxH24Buys := "maxH24Buys_example" // string | Maximum 24h buys (optional)
	minH24Sells := "minH24Sells_example" // string | Minimum 24h sells (optional)
	maxH24Sells := "maxH24Sells_example" // string | Maximum 24h sells (optional)
	minH24Trades := "minH24Trades_example" // string | Minimum 24h trades (optional)
	maxH24Trades := "maxH24Trades_example" // string | Maximum 24h trades (optional)
	minH24Buyers := "minH24Buyers_example" // string | Minimum 24h buyers (optional)
	maxH24Buyers := "maxH24Buyers_example" // string | Maximum 24h buyers (optional)
	minH24Sellers := "minH24Sellers_example" // string | Minimum 24h sellers (optional)
	maxH24Sellers := "maxH24Sellers_example" // string | Maximum 24h sellers (optional)
	minH24BuyVolumeInUsd := "minH24BuyVolumeInUsd_example" // string | Minimum 24h buy volume in USD (optional)
	maxH24BuyVolumeInUsd := "maxH24BuyVolumeInUsd_example" // string | Maximum 24h buy volume in USD (optional)
	minH24SellVolumeInUsd := "minH24SellVolumeInUsd_example" // string | Minimum 24h sell volume in USD (optional)
	maxH24SellVolumeInUsd := "maxH24SellVolumeInUsd_example" // string | Maximum 24h sell volume in USD (optional)
	minH4VolumeInUsd := "minH4VolumeInUsd_example" // string | Minimum 4h volume in USD (optional)
	maxH4VolumeInUsd := "maxH4VolumeInUsd_example" // string | Maximum 4h volume in USD (optional)
	minH4PriceChangeRatio := "minH4PriceChangeRatio_example" // string | Minimum 4h price change ratio (optional)
	maxH4PriceChangeRatio := "maxH4PriceChangeRatio_example" // string | Maximum 4h price change ratio (optional)
	minH4Buys := "minH4Buys_example" // string | Minimum 4h buys (optional)
	maxH4Buys := "maxH4Buys_example" // string | Maximum 4h buys (optional)
	minH4Sells := "minH4Sells_example" // string | Minimum 4h sells (optional)
	maxH4Sells := "maxH4Sells_example" // string | Maximum 4h sells (optional)
	minH4Trades := "minH4Trades_example" // string | Minimum 4h trades (optional)
	maxH4Trades := "maxH4Trades_example" // string | Maximum 4h trades (optional)
	minH4Buyers := "minH4Buyers_example" // string | Minimum 4h buyers (optional)
	maxH4Buyers := "maxH4Buyers_example" // string | Maximum 4h buyers (optional)
	minH4Sellers := "minH4Sellers_example" // string | Minimum 4h sellers (optional)
	maxH4Sellers := "maxH4Sellers_example" // string | Maximum 4h sellers (optional)
	minH4BuyVolumeInUsd := "minH4BuyVolumeInUsd_example" // string | Minimum 4h buy volume in USD (optional)
	maxH4BuyVolumeInUsd := "maxH4BuyVolumeInUsd_example" // string | Maximum 4h buy volume in USD (optional)
	minH4SellVolumeInUsd := "minH4SellVolumeInUsd_example" // string | Minimum 4h sell volume in USD (optional)
	maxH4SellVolumeInUsd := "maxH4SellVolumeInUsd_example" // string | Maximum 4h sell volume in USD (optional)
	minH1VolumeInUsd := "minH1VolumeInUsd_example" // string | Minimum 1h volume in USD (optional)
	maxH1VolumeInUsd := "maxH1VolumeInUsd_example" // string | Maximum 1h volume in USD (optional)
	minH1PriceChangeRatio := "minH1PriceChangeRatio_example" // string | Minimum 1h price change ratio (optional)
	maxH1PriceChangeRatio := "maxH1PriceChangeRatio_example" // string | Maximum 1h price change ratio (optional)
	minH1Buys := "minH1Buys_example" // string | Minimum 1h buys (optional)
	maxH1Buys := "maxH1Buys_example" // string | Maximum 1h buys (optional)
	minH1Sells := "minH1Sells_example" // string | Minimum 1h sells (optional)
	maxH1Sells := "maxH1Sells_example" // string | Maximum 1h sells (optional)
	minH1Trades := "minH1Trades_example" // string | Minimum 1h trades (optional)
	maxH1Trades := "maxH1Trades_example" // string | Maximum 1h trades (optional)
	minH1Buyers := "minH1Buyers_example" // string | Minimum 1h buyers (optional)
	maxH1Buyers := "maxH1Buyers_example" // string | Maximum 1h buyers (optional)
	minH1Sellers := "minH1Sellers_example" // string | Minimum 1h sellers (optional)
	maxH1Sellers := "maxH1Sellers_example" // string | Maximum 1h sellers (optional)
	minH1BuyVolumeInUsd := "minH1BuyVolumeInUsd_example" // string | Minimum 1h buy volume in USD (optional)
	maxH1BuyVolumeInUsd := "maxH1BuyVolumeInUsd_example" // string | Maximum 1h buy volume in USD (optional)
	minH1SellVolumeInUsd := "minH1SellVolumeInUsd_example" // string | Minimum 1h sell volume in USD (optional)
	maxH1SellVolumeInUsd := "maxH1SellVolumeInUsd_example" // string | Maximum 1h sell volume in USD (optional)
	minM30VolumeInUsd := "minM30VolumeInUsd_example" // string | Minimum 30m volume in USD (optional)
	maxM30VolumeInUsd := "maxM30VolumeInUsd_example" // string | Maximum 30m volume in USD (optional)
	minM30PriceChangeRatio := "minM30PriceChangeRatio_example" // string | Minimum 30m price change ratio (optional)
	maxM30PriceChangeRatio := "maxM30PriceChangeRatio_example" // string | Maximum 30m price change ratio (optional)
	minM30Buys := "minM30Buys_example" // string | Minimum 30m buys (optional)
	maxM30Buys := "maxM30Buys_example" // string | Maximum 30m buys (optional)
	minM30Sells := "minM30Sells_example" // string | Minimum 30m sells (optional)
	maxM30Sells := "maxM30Sells_example" // string | Maximum 30m sells (optional)
	minM30Trades := "minM30Trades_example" // string | Minimum 30m trades (optional)
	maxM30Trades := "maxM30Trades_example" // string | Maximum 30m trades (optional)
	minM30Buyers := "minM30Buyers_example" // string | Minimum 30m buyers (optional)
	maxM30Buyers := "maxM30Buyers_example" // string | Maximum 30m buyers (optional)
	minM30Sellers := "minM30Sellers_example" // string | Minimum 30m sellers (optional)
	maxM30Sellers := "maxM30Sellers_example" // string | Maximum 30m sellers (optional)
	minM30BuyVolumeInUsd := "minM30BuyVolumeInUsd_example" // string | Minimum 30m buy volume in USD (optional)
	maxM30BuyVolumeInUsd := "maxM30BuyVolumeInUsd_example" // string | Maximum 30m buy volume in USD (optional)
	minM30SellVolumeInUsd := "minM30SellVolumeInUsd_example" // string | Minimum 30m sell volume in USD (optional)
	maxM30SellVolumeInUsd := "maxM30SellVolumeInUsd_example" // string | Maximum 30m sell volume in USD (optional)
	minM15VolumeInUsd := "minM15VolumeInUsd_example" // string | Minimum 15m volume in USD (optional)
	maxM15VolumeInUsd := "maxM15VolumeInUsd_example" // string | Maximum 15m volume in USD (optional)
	minM15PriceChangeRatio := "minM15PriceChangeRatio_example" // string | Minimum 15m price change ratio (optional)
	maxM15PriceChangeRatio := "maxM15PriceChangeRatio_example" // string | Maximum 15m price change ratio (optional)
	minM15Buys := "minM15Buys_example" // string | Minimum 15m buys (optional)
	maxM15Buys := "maxM15Buys_example" // string | Maximum 15m buys (optional)
	minM15Sells := "minM15Sells_example" // string | Minimum 15m sells (optional)
	maxM15Sells := "maxM15Sells_example" // string | Maximum 15m sells (optional)
	minM15Trades := "minM15Trades_example" // string | Minimum 15m trades (optional)
	maxM15Trades := "maxM15Trades_example" // string | Maximum 15m trades (optional)
	minM15Buyers := "minM15Buyers_example" // string | Minimum 15m buyers (optional)
	maxM15Buyers := "maxM15Buyers_example" // string | Maximum 15m buyers (optional)
	minM15Sellers := "minM15Sellers_example" // string | Minimum 15m sellers (optional)
	maxM15Sellers := "maxM15Sellers_example" // string | Maximum 15m sellers (optional)
	minM15BuyVolumeInUsd := "minM15BuyVolumeInUsd_example" // string | Minimum 15m buy volume in USD (optional)
	maxM15BuyVolumeInUsd := "maxM15BuyVolumeInUsd_example" // string | Maximum 15m buy volume in USD (optional)
	minM15SellVolumeInUsd := "minM15SellVolumeInUsd_example" // string | Minimum 15m sell volume in USD (optional)
	maxM15SellVolumeInUsd := "maxM15SellVolumeInUsd_example" // string | Maximum 15m sell volume in USD (optional)
	minM5VolumeInUsd := "minM5VolumeInUsd_example" // string | Minimum 5m volume in USD (optional)
	maxM5VolumeInUsd := "maxM5VolumeInUsd_example" // string | Maximum 5m volume in USD (optional)
	minM5PriceChangeRatio := "minM5PriceChangeRatio_example" // string | Minimum 5m price change ratio (optional)
	maxM5PriceChangeRatio := "maxM5PriceChangeRatio_example" // string | Maximum 5m price change ratio (optional)
	minM5Buys := "minM5Buys_example" // string | Minimum 5m buys (optional)
	maxM5Buys := "maxM5Buys_example" // string | Maximum 5m buys (optional)
	minM5Sells := "minM5Sells_example" // string | Minimum 5m sells (optional)
	maxM5Sells := "maxM5Sells_example" // string | Maximum 5m sells (optional)
	minM5Trades := "minM5Trades_example" // string | Minimum 5m trades (optional)
	maxM5Trades := "maxM5Trades_example" // string | Maximum 5m trades (optional)
	minM5Buyers := "minM5Buyers_example" // string | Minimum 5m buyers (optional)
	maxM5Buyers := "maxM5Buyers_example" // string | Maximum 5m buyers (optional)
	minM5Sellers := "minM5Sellers_example" // string | Minimum 5m sellers (optional)
	maxM5Sellers := "maxM5Sellers_example" // string | Maximum 5m sellers (optional)
	minM5BuyVolumeInUsd := "minM5BuyVolumeInUsd_example" // string | Minimum 5m buy volume in USD (optional)
	maxM5BuyVolumeInUsd := "maxM5BuyVolumeInUsd_example" // string | Maximum 5m buy volume in USD (optional)
	minM5SellVolumeInUsd := "minM5SellVolumeInUsd_example" // string | Minimum 5m sell volume in USD (optional)
	maxM5SellVolumeInUsd := "maxM5SellVolumeInUsd_example" // string | Maximum 5m sell volume in USD (optional)
	minM1VolumeInUsd := "minM1VolumeInUsd_example" // string | Minimum 1m volume in USD (optional)
	maxM1VolumeInUsd := "maxM1VolumeInUsd_example" // string | Maximum 1m volume in USD (optional)
	minM1PriceChangeRatio := "minM1PriceChangeRatio_example" // string | Minimum 1m price change ratio (optional)
	maxM1PriceChangeRatio := "maxM1PriceChangeRatio_example" // string | Maximum 1m price change ratio (optional)
	minM1Buys := "minM1Buys_example" // string | Minimum 1m buys (optional)
	maxM1Buys := "maxM1Buys_example" // string | Maximum 1m buys (optional)
	minM1Sells := "minM1Sells_example" // string | Minimum 1m sells (optional)
	maxM1Sells := "maxM1Sells_example" // string | Maximum 1m sells (optional)
	minM1Trades := "minM1Trades_example" // string | Minimum 1m trades (optional)
	maxM1Trades := "maxM1Trades_example" // string | Maximum 1m trades (optional)
	minM1Buyers := "minM1Buyers_example" // string | Minimum 1m buyers (optional)
	maxM1Buyers := "maxM1Buyers_example" // string | Maximum 1m buyers (optional)
	minM1Sellers := "minM1Sellers_example" // string | Minimum 1m sellers (optional)
	maxM1Sellers := "maxM1Sellers_example" // string | Maximum 1m sellers (optional)
	minM1BuyVolumeInUsd := "minM1BuyVolumeInUsd_example" // string | Minimum 1m buy volume in USD (optional)
	maxM1BuyVolumeInUsd := "maxM1BuyVolumeInUsd_example" // string | Maximum 1m buy volume in USD (optional)
	minM1SellVolumeInUsd := "minM1SellVolumeInUsd_example" // string | Minimum 1m sell volume in USD (optional)
	maxM1SellVolumeInUsd := "maxM1SellVolumeInUsd_example" // string | Maximum 1m sell volume in USD (optional)

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
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **float32** | Number of results per page | [default to 20]
 **direction** | **string** | Pagination direction (next or prev) | [default to &quot;next&quot;]
 **sort** | **string** | Sort direction | [default to &quot;desc&quot;]
 **sortBy** | **string** | Sort by field | 
 **minH24VolumeInUsd** | **string** | Minimum 24h volume in USD | 
 **maxH24VolumeInUsd** | **string** | Maximum 24h volume in USD | 
 **minH24PriceChangeRatio** | **string** | Minimum 24h price change ratio | 
 **maxH24PriceChangeRatio** | **string** | Maximum 24h price change ratio | 
 **minH24Buys** | **string** | Minimum 24h buys | 
 **maxH24Buys** | **string** | Maximum 24h buys | 
 **minH24Sells** | **string** | Minimum 24h sells | 
 **maxH24Sells** | **string** | Maximum 24h sells | 
 **minH24Trades** | **string** | Minimum 24h trades | 
 **maxH24Trades** | **string** | Maximum 24h trades | 
 **minH24Buyers** | **string** | Minimum 24h buyers | 
 **maxH24Buyers** | **string** | Maximum 24h buyers | 
 **minH24Sellers** | **string** | Minimum 24h sellers | 
 **maxH24Sellers** | **string** | Maximum 24h sellers | 
 **minH24BuyVolumeInUsd** | **string** | Minimum 24h buy volume in USD | 
 **maxH24BuyVolumeInUsd** | **string** | Maximum 24h buy volume in USD | 
 **minH24SellVolumeInUsd** | **string** | Minimum 24h sell volume in USD | 
 **maxH24SellVolumeInUsd** | **string** | Maximum 24h sell volume in USD | 
 **minH4VolumeInUsd** | **string** | Minimum 4h volume in USD | 
 **maxH4VolumeInUsd** | **string** | Maximum 4h volume in USD | 
 **minH4PriceChangeRatio** | **string** | Minimum 4h price change ratio | 
 **maxH4PriceChangeRatio** | **string** | Maximum 4h price change ratio | 
 **minH4Buys** | **string** | Minimum 4h buys | 
 **maxH4Buys** | **string** | Maximum 4h buys | 
 **minH4Sells** | **string** | Minimum 4h sells | 
 **maxH4Sells** | **string** | Maximum 4h sells | 
 **minH4Trades** | **string** | Minimum 4h trades | 
 **maxH4Trades** | **string** | Maximum 4h trades | 
 **minH4Buyers** | **string** | Minimum 4h buyers | 
 **maxH4Buyers** | **string** | Maximum 4h buyers | 
 **minH4Sellers** | **string** | Minimum 4h sellers | 
 **maxH4Sellers** | **string** | Maximum 4h sellers | 
 **minH4BuyVolumeInUsd** | **string** | Minimum 4h buy volume in USD | 
 **maxH4BuyVolumeInUsd** | **string** | Maximum 4h buy volume in USD | 
 **minH4SellVolumeInUsd** | **string** | Minimum 4h sell volume in USD | 
 **maxH4SellVolumeInUsd** | **string** | Maximum 4h sell volume in USD | 
 **minH1VolumeInUsd** | **string** | Minimum 1h volume in USD | 
 **maxH1VolumeInUsd** | **string** | Maximum 1h volume in USD | 
 **minH1PriceChangeRatio** | **string** | Minimum 1h price change ratio | 
 **maxH1PriceChangeRatio** | **string** | Maximum 1h price change ratio | 
 **minH1Buys** | **string** | Minimum 1h buys | 
 **maxH1Buys** | **string** | Maximum 1h buys | 
 **minH1Sells** | **string** | Minimum 1h sells | 
 **maxH1Sells** | **string** | Maximum 1h sells | 
 **minH1Trades** | **string** | Minimum 1h trades | 
 **maxH1Trades** | **string** | Maximum 1h trades | 
 **minH1Buyers** | **string** | Minimum 1h buyers | 
 **maxH1Buyers** | **string** | Maximum 1h buyers | 
 **minH1Sellers** | **string** | Minimum 1h sellers | 
 **maxH1Sellers** | **string** | Maximum 1h sellers | 
 **minH1BuyVolumeInUsd** | **string** | Minimum 1h buy volume in USD | 
 **maxH1BuyVolumeInUsd** | **string** | Maximum 1h buy volume in USD | 
 **minH1SellVolumeInUsd** | **string** | Minimum 1h sell volume in USD | 
 **maxH1SellVolumeInUsd** | **string** | Maximum 1h sell volume in USD | 
 **minM30VolumeInUsd** | **string** | Minimum 30m volume in USD | 
 **maxM30VolumeInUsd** | **string** | Maximum 30m volume in USD | 
 **minM30PriceChangeRatio** | **string** | Minimum 30m price change ratio | 
 **maxM30PriceChangeRatio** | **string** | Maximum 30m price change ratio | 
 **minM30Buys** | **string** | Minimum 30m buys | 
 **maxM30Buys** | **string** | Maximum 30m buys | 
 **minM30Sells** | **string** | Minimum 30m sells | 
 **maxM30Sells** | **string** | Maximum 30m sells | 
 **minM30Trades** | **string** | Minimum 30m trades | 
 **maxM30Trades** | **string** | Maximum 30m trades | 
 **minM30Buyers** | **string** | Minimum 30m buyers | 
 **maxM30Buyers** | **string** | Maximum 30m buyers | 
 **minM30Sellers** | **string** | Minimum 30m sellers | 
 **maxM30Sellers** | **string** | Maximum 30m sellers | 
 **minM30BuyVolumeInUsd** | **string** | Minimum 30m buy volume in USD | 
 **maxM30BuyVolumeInUsd** | **string** | Maximum 30m buy volume in USD | 
 **minM30SellVolumeInUsd** | **string** | Minimum 30m sell volume in USD | 
 **maxM30SellVolumeInUsd** | **string** | Maximum 30m sell volume in USD | 
 **minM15VolumeInUsd** | **string** | Minimum 15m volume in USD | 
 **maxM15VolumeInUsd** | **string** | Maximum 15m volume in USD | 
 **minM15PriceChangeRatio** | **string** | Minimum 15m price change ratio | 
 **maxM15PriceChangeRatio** | **string** | Maximum 15m price change ratio | 
 **minM15Buys** | **string** | Minimum 15m buys | 
 **maxM15Buys** | **string** | Maximum 15m buys | 
 **minM15Sells** | **string** | Minimum 15m sells | 
 **maxM15Sells** | **string** | Maximum 15m sells | 
 **minM15Trades** | **string** | Minimum 15m trades | 
 **maxM15Trades** | **string** | Maximum 15m trades | 
 **minM15Buyers** | **string** | Minimum 15m buyers | 
 **maxM15Buyers** | **string** | Maximum 15m buyers | 
 **minM15Sellers** | **string** | Minimum 15m sellers | 
 **maxM15Sellers** | **string** | Maximum 15m sellers | 
 **minM15BuyVolumeInUsd** | **string** | Minimum 15m buy volume in USD | 
 **maxM15BuyVolumeInUsd** | **string** | Maximum 15m buy volume in USD | 
 **minM15SellVolumeInUsd** | **string** | Minimum 15m sell volume in USD | 
 **maxM15SellVolumeInUsd** | **string** | Maximum 15m sell volume in USD | 
 **minM5VolumeInUsd** | **string** | Minimum 5m volume in USD | 
 **maxM5VolumeInUsd** | **string** | Maximum 5m volume in USD | 
 **minM5PriceChangeRatio** | **string** | Minimum 5m price change ratio | 
 **maxM5PriceChangeRatio** | **string** | Maximum 5m price change ratio | 
 **minM5Buys** | **string** | Minimum 5m buys | 
 **maxM5Buys** | **string** | Maximum 5m buys | 
 **minM5Sells** | **string** | Minimum 5m sells | 
 **maxM5Sells** | **string** | Maximum 5m sells | 
 **minM5Trades** | **string** | Minimum 5m trades | 
 **maxM5Trades** | **string** | Maximum 5m trades | 
 **minM5Buyers** | **string** | Minimum 5m buyers | 
 **maxM5Buyers** | **string** | Maximum 5m buyers | 
 **minM5Sellers** | **string** | Minimum 5m sellers | 
 **maxM5Sellers** | **string** | Maximum 5m sellers | 
 **minM5BuyVolumeInUsd** | **string** | Minimum 5m buy volume in USD | 
 **maxM5BuyVolumeInUsd** | **string** | Maximum 5m buy volume in USD | 
 **minM5SellVolumeInUsd** | **string** | Minimum 5m sell volume in USD | 
 **maxM5SellVolumeInUsd** | **string** | Maximum 5m sell volume in USD | 
 **minM1VolumeInUsd** | **string** | Minimum 1m volume in USD | 
 **maxM1VolumeInUsd** | **string** | Maximum 1m volume in USD | 
 **minM1PriceChangeRatio** | **string** | Minimum 1m price change ratio | 
 **maxM1PriceChangeRatio** | **string** | Maximum 1m price change ratio | 
 **minM1Buys** | **string** | Minimum 1m buys | 
 **maxM1Buys** | **string** | Maximum 1m buys | 
 **minM1Sells** | **string** | Minimum 1m sells | 
 **maxM1Sells** | **string** | Maximum 1m sells | 
 **minM1Trades** | **string** | Minimum 1m trades | 
 **maxM1Trades** | **string** | Maximum 1m trades | 
 **minM1Buyers** | **string** | Minimum 1m buyers | 
 **maxM1Buyers** | **string** | Maximum 1m buyers | 
 **minM1Sellers** | **string** | Minimum 1m sellers | 
 **maxM1Sellers** | **string** | Maximum 1m sellers | 
 **minM1BuyVolumeInUsd** | **string** | Minimum 1m buy volume in USD | 
 **maxM1BuyVolumeInUsd** | **string** | Maximum 1m buy volume in USD | 
 **minM1SellVolumeInUsd** | **string** | Minimum 1m sell volume in USD | 
 **maxM1SellVolumeInUsd** | **string** | Maximum 1m sell volume in USD | 

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

Token - Search



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
	chains := []string{"Inner_example"} // []string | Chain filter (optional)
	q := "USDC" // string | Search query string for token name, symbol or address (optional)
	limit := int64(789) // int64 | Number of results per page (optional) (default to 20)
	sort := "sort_example" // string | Sort direction (optional) (default to "desc")
	protocols := []string{"Inner_example"} // []string | Protocol filter, for supported protocols please check the dexName values returned by DEX - List (optional)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	sortBy := "priceInUsd" // string | Field to sort by (optional)

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
 **chains** | **[]string** | Chain filter | 
 **q** | **string** | Search query string for token name, symbol or address | 
 **limit** | **int64** | Number of results per page | [default to 20]
 **sort** | **string** | Sort direction | [default to &quot;desc&quot;]
 **protocols** | **[]string** | Protocol filter, for supported protocols please check the dexName values returned by DEX - List | 
 **cursor** | **string** | Pagination cursor | 
 **sortBy** | **string** | Field to sort by | 

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

