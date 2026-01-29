# \DefiSolanaMoonshotAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoonshotCreateToken**](DefiSolanaMoonshotAPI.md#MoonshotCreateToken) | **Post** /v1/solana/moonshot/create | Moonshot - Create Token
[**MoonshotSubmitCreateToken**](DefiSolanaMoonshotAPI.md#MoonshotSubmitCreateToken) | **Post** /v1/solana/moonshot/submitCreateToken | Moonshot - Submit Create Token



## MoonshotCreateToken

> MoonshotCreateTokenReply MoonshotCreateToken(ctx).MoonshotCreateTokenInput(moonshotCreateTokenInput).Execute()

Moonshot - Create Token



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
	moonshotCreateTokenInput := *openapiclient.NewMoonshotCreateTokenInput("moonshot", "oQPnhXAbLbMuKHESaGrbXT17CyvWCpLyERSJA9HCYd7", "My Token", "MTK", "RAYDIUM", "https://assets.mytoken.com/icon.png", "A revolutionary new token for the Solana ecosystem", []openapiclient.Link{*openapiclient.NewLink("https://twitter.com/tokenproject", "Twitter")}, "1000000000") // MoonshotCreateTokenInput | Token creation parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefiSolanaMoonshotAPI.MoonshotCreateToken(context.Background()).MoonshotCreateTokenInput(moonshotCreateTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefiSolanaMoonshotAPI.MoonshotCreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoonshotCreateToken`: MoonshotCreateTokenReply
	fmt.Fprintf(os.Stdout, "Response from `DefiSolanaMoonshotAPI.MoonshotCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoonshotCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moonshotCreateTokenInput** | [**MoonshotCreateTokenInput**](MoonshotCreateTokenInput.md) | Token creation parameters | 

### Return type

[**MoonshotCreateTokenReply**](MoonshotCreateTokenReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoonshotSubmitCreateToken

> MoonshotSubmitCreateToken200Response MoonshotSubmitCreateToken(ctx).MoonshotSubmitCreateTokenInput(moonshotSubmitCreateTokenInput).Execute()

Moonshot - Submit Create Token



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
	moonshotSubmitCreateTokenInput := *openapiclient.NewMoonshotSubmitCreateTokenInput("AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEDRgYGpQEDAQIABQcICQoLDA0ODwAAAAAAAAAAAAAQERITFBUWFxgZGhscHR4fAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=") // MoonshotSubmitCreateTokenInput | Signed transaction and token details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefiSolanaMoonshotAPI.MoonshotSubmitCreateToken(context.Background()).MoonshotSubmitCreateTokenInput(moonshotSubmitCreateTokenInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefiSolanaMoonshotAPI.MoonshotSubmitCreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoonshotSubmitCreateToken`: MoonshotSubmitCreateToken200Response
	fmt.Fprintf(os.Stdout, "Response from `DefiSolanaMoonshotAPI.MoonshotSubmitCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoonshotSubmitCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moonshotSubmitCreateTokenInput** | [**MoonshotSubmitCreateTokenInput**](MoonshotSubmitCreateTokenInput.md) | Signed transaction and token details | 

### Return type

[**MoonshotSubmitCreateToken200Response**](MoonshotSubmitCreateToken200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

