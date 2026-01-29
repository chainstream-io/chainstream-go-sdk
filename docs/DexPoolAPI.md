# \DexPoolAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDexpool**](DexPoolAPI.md#GetDexpool) | **Get** /v1/dexpools/{chain}/{poolAddress} | DexPool - Detail
[**GetDexpoolSnapshots**](DexPoolAPI.md#GetDexpoolSnapshots) | **Get** /v1/dexpools/{chain}/{poolAddress}/snapshots | DexPool - Liquidity Snapshots



## GetDexpool

> DexPoolDTO GetDexpool(ctx, chain, poolAddress).Execute()

DexPool - Detail



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
	poolAddress := "58oQChx4yWmvKdwLLZzBi4ChoCc2fqCUWBkwMihLYQo2" // string | A pool address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexPoolAPI.GetDexpool(context.Background(), chain, poolAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexPoolAPI.GetDexpool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDexpool`: DexPoolDTO
	fmt.Fprintf(os.Stdout, "Response from `DexPoolAPI.GetDexpool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**poolAddress** | **string** | A pool address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDexpoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DexPoolDTO**](DexPoolDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDexpoolSnapshots

> DexPoolSnapshotPage GetDexpoolSnapshots(ctx, chain, poolAddress).Time(time).Cursor(cursor).Limit(limit).Execute()

DexPool - Liquidity Snapshots



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
	poolAddress := "58oQChx4yWmvKdwLLZzBi4ChoCc2fqCUWBkwMihLYQo2" // string | A pool address
	time := int64(1705312800) // int64 | Target Unix timestamp (seconds) for snapshot query. Returns the nearest snapshot before or at this time. (optional)
	cursor := "eyJpZCI6MTAwfQ==" // string | Pagination cursor (optional)
	limit := int32(20) // int32 | Number of results per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DexPoolAPI.GetDexpoolSnapshots(context.Background(), chain, poolAddress).Time(time).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DexPoolAPI.GetDexpoolSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDexpoolSnapshots`: DexPoolSnapshotPage
	fmt.Fprintf(os.Stdout, "Response from `DexPoolAPI.GetDexpoolSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 
**poolAddress** | **string** | A pool address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDexpoolSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **time** | **int64** | Target Unix timestamp (seconds) for snapshot query. Returns the nearest snapshot before or at this time. | 
 **cursor** | **string** | Pagination cursor | 
 **limit** | **int32** | Number of results per page | [default to 20]

### Return type

[**DexPoolSnapshotPage**](DexPoolSnapshotPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

