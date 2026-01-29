# \RankingAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinalStretchTokens**](RankingAPI.md#GetFinalStretchTokens) | **Get** /v1/ranking/{chain}/finalStretch | Ranking - FinalStretch Tokens
[**GetHotTokens**](RankingAPI.md#GetHotTokens) | **Get** /v1/ranking/{chain}/hotTokens/{duration} | Ranking - Hot Tokens
[**GetMigratedTokens**](RankingAPI.md#GetMigratedTokens) | **Get** /v1/ranking/{chain}/migrated | Ranking - Migrated Tokens
[**GetNewTokens**](RankingAPI.md#GetNewTokens) | **Get** /v1/ranking/{chain}/newTokens | Ranking - New Tokens
[**GetStocksTokens**](RankingAPI.md#GetStocksTokens) | **Get** /v1/ranking/{chain}/stocks | Ranking - Stocks Tokens



## GetFinalStretchTokens

> []Token GetFinalStretchTokens(ctx, chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

Ranking - FinalStretch Tokens



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
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)
	tag := "tag_example" // string | Token tag filter for specific protocols or categories (optional)
	filters := []string{"Inner_example"} // []string | Filters (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | Launchpad Platform (optional)
	searchKeywords := []string{"Inner_example"} // []string | Search Keywords (optional)
	excludeKeywords := []string{"Inner_example"} // []string | Exclude Keywords (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingAPI.GetFinalStretchTokens(context.Background(), chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingAPI.GetFinalStretchTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinalStretchTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `RankingAPI.GetFinalStretchTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalStretchTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 
 **tag** | **string** | Token tag filter for specific protocols or categories | 
 **filters** | **[]string** | Filters | 
 **launchpadPlatform** | **[]string** | Launchpad Platform | 
 **searchKeywords** | **[]string** | Search Keywords | 
 **excludeKeywords** | **[]string** | Exclude Keywords | 

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


## GetHotTokens

> []Token GetHotTokens(ctx, chain, duration).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

Ranking - Hot Tokens



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
	duration := "duration_example" // string | Duration of the ranking
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)
	tag := "tag_example" // string | Token tag filter for specific protocols or categories (optional)
	filters := []string{"Inner_example"} // []string | Filters (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | Launchpad Platform (optional)
	searchKeywords := []string{"Inner_example"} // []string | Search Keywords (optional)
	excludeKeywords := []string{"Inner_example"} // []string | Exclude Keywords (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingAPI.GetHotTokens(context.Background(), chain, duration).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingAPI.GetHotTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHotTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `RankingAPI.GetHotTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**duration** | **string** | Duration of the ranking | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 
 **tag** | **string** | Token tag filter for specific protocols or categories | 
 **filters** | **[]string** | Filters | 
 **launchpadPlatform** | **[]string** | Launchpad Platform | 
 **searchKeywords** | **[]string** | Search Keywords | 
 **excludeKeywords** | **[]string** | Exclude Keywords | 

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


## GetMigratedTokens

> []Token GetMigratedTokens(ctx, chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

Ranking - Migrated Tokens



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
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)
	tag := "tag_example" // string | Token tag filter for specific protocols or categories (optional)
	filters := []string{"Inner_example"} // []string | Filters (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | Launchpad Platform (optional)
	searchKeywords := []string{"Inner_example"} // []string | Search Keywords (optional)
	excludeKeywords := []string{"Inner_example"} // []string | Exclude Keywords (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingAPI.GetMigratedTokens(context.Background(), chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingAPI.GetMigratedTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMigratedTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `RankingAPI.GetMigratedTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigratedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 
 **tag** | **string** | Token tag filter for specific protocols or categories | 
 **filters** | **[]string** | Filters | 
 **launchpadPlatform** | **[]string** | Launchpad Platform | 
 **searchKeywords** | **[]string** | Search Keywords | 
 **excludeKeywords** | **[]string** | Exclude Keywords | 

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


## GetNewTokens

> []Token GetNewTokens(ctx, chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

Ranking - New Tokens



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
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)
	tag := "tag_example" // string | Token tag filter for specific protocols or categories (optional)
	filters := []string{"Inner_example"} // []string | Filters (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | Launchpad Platform (optional)
	searchKeywords := []string{"Inner_example"} // []string | Search Keywords (optional)
	excludeKeywords := []string{"Inner_example"} // []string | Exclude Keywords (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingAPI.GetNewTokens(context.Background(), chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingAPI.GetNewTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `RankingAPI.GetNewTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 
 **tag** | **string** | Token tag filter for specific protocols or categories | 
 **filters** | **[]string** | Filters | 
 **launchpadPlatform** | **[]string** | Launchpad Platform | 
 **searchKeywords** | **[]string** | Search Keywords | 
 **excludeKeywords** | **[]string** | Exclude Keywords | 

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


## GetStocksTokens

> []Token GetStocksTokens(ctx, chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

Ranking - Stocks Tokens



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
	sortBy := "marketData.marketCapInUsd" // string | Sort field (optional)
	sortDirection := "sortDirection_example" // string | Sort Direction (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | Filter field (optional)
	tag := "tag_example" // string | Token tag filter for specific protocols or categories (optional)
	filters := []string{"Inner_example"} // []string | Filters (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | Launchpad Platform (optional)
	searchKeywords := []string{"Inner_example"} // []string | Search Keywords (optional)
	excludeKeywords := []string{"Inner_example"} // []string | Exclude Keywords (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingAPI.GetStocksTokens(context.Background(), chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingAPI.GetStocksTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStocksTokens`: []Token
	fmt.Fprintf(os.Stdout, "Response from `RankingAPI.GetStocksTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Sort field | 
 **sortDirection** | **string** | Sort Direction | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | Filter field | 
 **tag** | **string** | Token tag filter for specific protocols or categories | 
 **filters** | **[]string** | Filters | 
 **launchpadPlatform** | **[]string** | Launchpad Platform | 
 **searchKeywords** | **[]string** | Search Keywords | 
 **excludeKeywords** | **[]string** | Exclude Keywords | 

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

