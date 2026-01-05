# \DefiSolanaPumpfunAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PumpfunCreateToken**](DefiSolanaPumpfunAPI.md#PumpfunCreateToken) | **Post** /v1/solana/pumpfun/create | CONTROLLER.PUMPFUN.CREATE_TOKEN.SUMMARY



## PumpfunCreateToken

> PumpCreateTokenReply PumpfunCreateToken(ctx).PumpCreateTokenInput(pumpCreateTokenInput).Execute()

CONTROLLER.PUMPFUN.CREATE_TOKEN.SUMMARY



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
	pumpCreateTokenInput := *openapiclient.NewPumpCreateTokenInput("pumpfun", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "My Token", "MTK", "RAYDIUM", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...", "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", "A revolutionary new token on Solana") // PumpCreateTokenInput | Required parameters for minting token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefiSolanaPumpfunAPI.PumpfunCreateToken(context.Background()).PumpCreateTokenInput(pumpCreateTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefiSolanaPumpfunAPI.PumpfunCreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PumpfunCreateToken`: PumpCreateTokenReply
	fmt.Fprintf(os.Stdout, "Response from `DefiSolanaPumpfunAPI.PumpfunCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPumpfunCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pumpCreateTokenInput** | [**PumpCreateTokenInput**](PumpCreateTokenInput.md) | Required parameters for minting token | 

### Return type

[**PumpCreateTokenReply**](PumpCreateTokenReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

