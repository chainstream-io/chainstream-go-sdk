# \EndpointAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](EndpointAPI.md#CreateEndpoint) | **Post** /v1/webhook/endpoint | Endpoint - Create
[**DeleteEndpoint**](EndpointAPI.md#DeleteEndpoint) | **Delete** /v1/webhook/endpoint/{id} | Endpoint - Delete
[**GetEndpoint**](EndpointAPI.md#GetEndpoint) | **Get** /v1/webhook/endpoint/{id} | Endpoint - Get
[**GetEndpointSecret**](EndpointAPI.md#GetEndpointSecret) | **Get** /v1/webhook/endpoint/{id}/secret | Endpoint - Get Secret
[**ListEndpoints**](EndpointAPI.md#ListEndpoints) | **Get** /v1/webhook/endpoint | Endpoint - List
[**RotateEndpointSecret**](EndpointAPI.md#RotateEndpointSecret) | **Post** /v1/webhook/endpoint/{id}/secret/rotate | Endpoint - Rotate Secret
[**UpdateEndpoint**](EndpointAPI.md#UpdateEndpoint) | **Patch** /v1/webhook/endpoint | Endpoint - Update



## CreateEndpoint

> EndpointResponse CreateEndpoint(ctx).CreateEndpointInput(createEndpointInput).Execute()

Endpoint - Create



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
	createEndpointInput := *openapiclient.NewCreateEndpointInput() // CreateEndpointInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.CreateEndpoint(context.Background()).CreateEndpointInput(createEndpointInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.CreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoint`: EndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointInput** | [**CreateEndpointInput**](CreateEndpointInput.md) |  | 

### Return type

[**EndpointResponse**](EndpointResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoint

> EndpointOperationResponse DeleteEndpoint(ctx, id).Execute()

Endpoint - Delete



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.DeleteEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.DeleteEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEndpoint`: EndpointOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.DeleteEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointOperationResponse**](EndpointOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoint

> EndpointResponse GetEndpoint(ctx, id).Execute()

Endpoint - Get



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpoint(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoint`: EndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointResponse**](EndpointResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointSecret

> EndpointSecretResponse GetEndpointSecret(ctx, id).Execute()

Endpoint - Get Secret



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.GetEndpointSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpointSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointSecret`: EndpointSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpointSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointSecretResponse**](EndpointSecretResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> EndpointListResponse ListEndpoints(ctx).Limit(limit).Iterator(iterator).Order(order).Execute()

Endpoint - List



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
	limit := int64(789) // int64 | Limit (optional) (default to 100)
	iterator := "iterator_example" // string | Iterator (optional)
	order := "order_example" // string | Order (optional) (default to "ascending")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.ListEndpoints(context.Background()).Limit(limit).Iterator(iterator).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.ListEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEndpoints`: EndpointListResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit | [default to 100]
 **iterator** | **string** | Iterator | 
 **order** | **string** | Order | [default to &quot;ascending&quot;]

### Return type

[**EndpointListResponse**](EndpointListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateEndpointSecret

> EndpointOperationResponse RotateEndpointSecret(ctx, id).Execute()

Endpoint - Rotate Secret



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.RotateEndpointSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.RotateEndpointSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateEndpointSecret`: EndpointOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.RotateEndpointSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateEndpointSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointOperationResponse**](EndpointOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> EndpointResponse UpdateEndpoint(ctx).UpdateEndpointInput(updateEndpointInput).Execute()

Endpoint - Update



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
	updateEndpointInput := *openapiclient.NewUpdateEndpointInput() // UpdateEndpointInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointAPI.UpdateEndpoint(context.Background()).UpdateEndpointInput(updateEndpointInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.UpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoint`: EndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.UpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEndpointInput** | [**UpdateEndpointInput**](UpdateEndpointInput.md) |  | 

### Return type

[**EndpointResponse**](EndpointResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

