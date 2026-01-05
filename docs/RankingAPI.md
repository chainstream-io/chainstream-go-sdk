# \RankingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinalStretchTokens**](RankingAPI.md#GetFinalStretchTokens) | **Get** /v1/ranking/{chain}/finalStretch | CONTROLLER.RANKING.FINAL_STRETCH_TOKENS.SUMMARY
[**GetHotTokens**](RankingAPI.md#GetHotTokens) | **Get** /v1/ranking/{chain}/hotTokens/{duration} | CONTROLLER.RANKING.HOT_TOKENS.SUMMARY
[**GetMigratedTokens**](RankingAPI.md#GetMigratedTokens) | **Get** /v1/ranking/{chain}/migrated | CONTROLLER.RANKING.MIGRATED_TOKENS.SUMMARY
[**GetNewTokens**](RankingAPI.md#GetNewTokens) | **Get** /v1/ranking/{chain}/newTokens | CONTROLLER.RANKING.NEW_TOKENS.SUMMARY
[**GetStocksTokens**](RankingAPI.md#GetStocksTokens) | **Get** /v1/ranking/{chain}/stocks | CONTROLLER.RANKING.STOCKS_TOKENS.SUMMARY



## GetFinalStretchTokens

> []Token GetFinalStretchTokens(ctx, chain).SortBy(sortBy).SortDirection(sortDirection).RangeFilters(rangeFilters).Tag(tag).Filters(filters).LaunchpadPlatform(launchpadPlatform).SearchKeywords(searchKeywords).ExcludeKeywords(excludeKeywords).Execute()

CONTROLLER.RANKING.FINAL_STRETCH_TOKENS.SUMMARY



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
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)
	tag := "tag_example" // string | DTO.TOKEN.REQUEST.TAG (optional)
	filters := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.FILTERS (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM (optional)
	searchKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.SEARCH_KEYWORDS (optional)
	excludeKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS (optional)

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalStretchTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 
 **tag** | **string** | DTO.TOKEN.REQUEST.TAG | 
 **filters** | **[]string** | DTO.TOKEN.FILTER.FILTERS | 
 **launchpadPlatform** | **[]string** | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM | 
 **searchKeywords** | **[]string** | DTO.TOKEN.FILTER.SEARCH_KEYWORDS | 
 **excludeKeywords** | **[]string** | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS | 

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

CONTROLLER.RANKING.HOT_TOKENS.SUMMARY



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
	duration := "duration_example" // string | CONTROLLER.RANKING.HOT_TOKENS.DURATION.DESCRIPTION
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)
	tag := "tag_example" // string | DTO.TOKEN.REQUEST.TAG (optional)
	filters := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.FILTERS (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM (optional)
	searchKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.SEARCH_KEYWORDS (optional)
	excludeKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS (optional)

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**duration** | **string** | CONTROLLER.RANKING.HOT_TOKENS.DURATION.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHotTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 
 **tag** | **string** | DTO.TOKEN.REQUEST.TAG | 
 **filters** | **[]string** | DTO.TOKEN.FILTER.FILTERS | 
 **launchpadPlatform** | **[]string** | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM | 
 **searchKeywords** | **[]string** | DTO.TOKEN.FILTER.SEARCH_KEYWORDS | 
 **excludeKeywords** | **[]string** | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS | 

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

CONTROLLER.RANKING.MIGRATED_TOKENS.SUMMARY



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
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)
	tag := "tag_example" // string | DTO.TOKEN.REQUEST.TAG (optional)
	filters := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.FILTERS (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM (optional)
	searchKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.SEARCH_KEYWORDS (optional)
	excludeKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS (optional)

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigratedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 
 **tag** | **string** | DTO.TOKEN.REQUEST.TAG | 
 **filters** | **[]string** | DTO.TOKEN.FILTER.FILTERS | 
 **launchpadPlatform** | **[]string** | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM | 
 **searchKeywords** | **[]string** | DTO.TOKEN.FILTER.SEARCH_KEYWORDS | 
 **excludeKeywords** | **[]string** | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS | 

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

CONTROLLER.RANKING.NEW_TOKENS.SUMMARY



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
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)
	tag := "tag_example" // string | DTO.TOKEN.REQUEST.TAG (optional)
	filters := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.FILTERS (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM (optional)
	searchKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.SEARCH_KEYWORDS (optional)
	excludeKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS (optional)

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 
 **tag** | **string** | DTO.TOKEN.REQUEST.TAG | 
 **filters** | **[]string** | DTO.TOKEN.FILTER.FILTERS | 
 **launchpadPlatform** | **[]string** | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM | 
 **searchKeywords** | **[]string** | DTO.TOKEN.FILTER.SEARCH_KEYWORDS | 
 **excludeKeywords** | **[]string** | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS | 

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

CONTROLLER.RANKING.STOCKS_TOKENS.SUMMARY



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
	sortBy := "marketData.marketCapInUsd" // string | DTO.TOKEN.REQUEST.SORT_BY (optional)
	sortDirection := "sortDirection_example" // string | DTO.TOKEN.REQUEST.SORT_DIRECTION (optional) (default to "DESC")
	rangeFilters := []openapiclient.FilterCondition{*openapiclient.NewFilterCondition()} // []FilterCondition | DTO.TOKEN.REQUEST.FILTER_BY (optional)
	tag := "tag_example" // string | DTO.TOKEN.REQUEST.TAG (optional)
	filters := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.FILTERS (optional)
	launchpadPlatform := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM (optional)
	searchKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.SEARCH_KEYWORDS (optional)
	excludeKeywords := []string{"Inner_example"} // []string | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS (optional)

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | DTO.TOKEN.REQUEST.SORT_BY | 
 **sortDirection** | **string** | DTO.TOKEN.REQUEST.SORT_DIRECTION | [default to &quot;DESC&quot;]
 **rangeFilters** | [**[]FilterCondition**](FilterCondition.md) | DTO.TOKEN.REQUEST.FILTER_BY | 
 **tag** | **string** | DTO.TOKEN.REQUEST.TAG | 
 **filters** | **[]string** | DTO.TOKEN.FILTER.FILTERS | 
 **launchpadPlatform** | **[]string** | DTO.TOKEN.FILTER.LAUNCHPAD_PLATFORM | 
 **searchKeywords** | **[]string** | DTO.TOKEN.FILTER.SEARCH_KEYWORDS | 
 **excludeKeywords** | **[]string** | DTO.TOKEN.FILTER.EXCLUDE_KEYWORDS | 

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

