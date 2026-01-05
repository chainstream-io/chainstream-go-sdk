# \DexPoolAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDexpool**](DexPoolAPI.md#GetDexpool) | **Get** /v1/dexpools/{chain}/{poolAddress} | CONTROLLER.DEXPOOL.GET.SUMMARY



## GetDexpool

> DexPoolDTO GetDexpool(ctx, chain, poolAddress).Execute()

CONTROLLER.DEXPOOL.GET.SUMMARY



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
	poolAddress := "58oQChx4yWmvKdwLLZzBi4ChoCc2fqCUWBkwMihLYQo2" // string | GLOBAL.POOLADDRESS.DESCRIPTION

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
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 
**poolAddress** | **string** | GLOBAL.POOLADDRESS.DESCRIPTION | 

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

