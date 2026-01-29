# \JobsAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](JobsAPI.md#Get) | **Get** /jobs/{id} | Job - Get
[**Streaming**](JobsAPI.md#Streaming) | **Get** /jobs/{id}/streaming | Job - Streaming



## Get

> JobDTO Get(ctx, id).Execute()

Job - Get



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
	id := "12345" // string | Job identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.Get(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: JobDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobDTO**](JobDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Streaming

> JobStreamingDTO Streaming(ctx, id).Execute()

Job - Streaming



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
	id := "12345" // string | Job identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.Streaming(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.Streaming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Streaming`: JobStreamingDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.Streaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobStreamingDTO**](JobStreamingDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

