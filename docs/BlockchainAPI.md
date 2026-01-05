# \BlockchainAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestBlock**](BlockchainAPI.md#GetLatestBlock) | **Get** /v1/blockchain/{chain}/latest_block | CONTROLLER.GET_LATEST_BLOCK.SUMMARY
[**GetSupportedBlockchains**](BlockchainAPI.md#GetSupportedBlockchains) | **Get** /v1/blockchain | CONTROLLER.GET_SUPPORTED_BLOCKCHAINS.SUMMARY



## GetLatestBlock

> BlockchainLatestBlockDTO GetLatestBlock(ctx, chain).Execute()

CONTROLLER.GET_LATEST_BLOCK.SUMMARY



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetLatestBlock(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetLatestBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestBlock`: BlockchainLatestBlockDTO
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetLatestBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | GLOBAL.CHAIN.DESCRIPTION | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainLatestBlockDTO**](BlockchainLatestBlockDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedBlockchains

> []BlockchainDTO GetSupportedBlockchains(ctx).Execute()

CONTROLLER.GET_SUPPORTED_BLOCKCHAINS.SUMMARY



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetSupportedBlockchains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetSupportedBlockchains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedBlockchains`: []BlockchainDTO
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetSupportedBlockchains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedBlockchainsRequest struct via the builder pattern


### Return type

[**[]BlockchainDTO**](BlockchainDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

