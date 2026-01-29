# \TransactionAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateGasLimit**](TransactionAPI.md#EstimateGasLimit) | **Post** /v1/transaction/{chain}/estimate-gas-limit | Transaction - Estimate Gas Limit
[**GetGasPrice**](TransactionAPI.md#GetGasPrice) | **Get** /v1/transaction/{chain}/gas-price | Transaction - Get Gas Price
[**Send**](TransactionAPI.md#Send) | **Post** /v1/transaction/{chain}/send | Transaction - Send



## EstimateGasLimit

> EstimateGasLimitResponse EstimateGasLimit(ctx, chain).EstimateGasLimitInput(estimateGasLimitInput).Execute()

Transaction - Estimate Gas Limit



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
	chain := "chain_example" // string | A chain name listed in supported networks
	estimateGasLimitInput := *openapiclient.NewEstimateGasLimitInput("0x742d35Cc6634C0532925a3b8D8C9C4b8a8d4C8b8", "0x742d35Cc6634C0532925a3b8D8C9C4b8a8d4C8b8", "0xa9059cbb000000000000000000000000742d35cc6634c0532925a3b8d8c9c4b8a8d4c8b80000000000000000000000000000000000000000000000000de0b6b3a7640000") // EstimateGasLimitInput | Transaction parameters for gas estimation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.EstimateGasLimit(context.Background(), chain).EstimateGasLimitInput(estimateGasLimitInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.EstimateGasLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateGasLimit`: EstimateGasLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.EstimateGasLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateGasLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **estimateGasLimitInput** | [**EstimateGasLimitInput**](EstimateGasLimitInput.md) | Transaction parameters for gas estimation | 

### Return type

[**EstimateGasLimitResponse**](EstimateGasLimitResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGasPrice

> GasPriceResponse GetGasPrice(ctx, chain).Execute()

Transaction - Get Gas Price



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
	chain := "chain_example" // string | A chain name listed in supported networks

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.GetGasPrice(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.GetGasPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGasPrice`: GasPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.GetGasPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGasPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GasPriceResponse**](GasPriceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> SendTxResponse Send(ctx, chain).SendTxInput(sendTxInput).Execute()

Transaction - Send



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
	chain := "{"SOLANA":"sol","BASE":"base","BSC":"bsc","POLYGON":"polygon","ARBITRUM":"arbitrum","OPTIMISM":"optimism","AVALANCHE":"avalanche","ETHEREUM":"eth","ZKSYNC":"zksync","SUI":"sui"}" // string | A chain name listed in supported networks
	sendTxInput := *openapiclient.NewSendTxInput("AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEDRgYGpQEDAQIABQcICQoLDA0ODwAAAAAAAAAAAAAQERITFBUWFxgZGhscHR4fAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=") // SendTxInput | Transaction parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.Send(context.Background(), chain).SendTxInput(sendTxInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.Send``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Send`: SendTxResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.Send`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTxInput** | [**SendTxInput**](SendTxInput.md) | Transaction parameters | 

### Return type

[**SendTxResponse**](SendTxResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

