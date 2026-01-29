# \RedPacketAPI

All URIs are relative to *https://api-dex.chainstream.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimRedpacket**](RedPacketAPI.md#ClaimRedpacket) | **Post** /v1/redpacket/{chain}/claim | RedPacket - Claim
[**CreateRedpacket**](RedPacketAPI.md#CreateRedpacket) | **Post** /v1/redpacket/{chain}/create | RedPacket - Create
[**GetClaims**](RedPacketAPI.md#GetClaims) | **Get** /v1/redpacket/{id}/claims | RedPacket - GetClaims
[**GetClaimsByAddress**](RedPacketAPI.md#GetClaimsByAddress) | **Get** /v1/redpacket/wallet/{address}/claims | RedPacket - GetClaimsByAddress
[**GetRedpacket**](RedPacketAPI.md#GetRedpacket) | **Get** /v1/redpacket/{id} | RedPacket - Get
[**GetRedpackets**](RedPacketAPI.md#GetRedpackets) | **Get** /v1/redpacket | RedPacket - GetRedpackets
[**GetRedpacketsByAddress**](RedPacketAPI.md#GetRedpacketsByAddress) | **Get** /v1/redpacket/wallet/{address}/redpackets | RedPacket - GetRedpacketsByAddress
[**RedpacketSend**](RedPacketAPI.md#RedpacketSend) | **Post** /v1/redpacket/{chain}/send | RedPacket - Send



## ClaimRedpacket

> RedPacketReply ClaimRedpacket(ctx, chain).ClaimRedPacketInput(claimRedPacketInput).Execute()

RedPacket - Claim



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
	claimRedPacketInput := *openapiclient.NewClaimRedPacketInput("sol", "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB") // ClaimRedPacketInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.ClaimRedpacket(context.Background(), chain).ClaimRedPacketInput(claimRedPacketInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.ClaimRedpacket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimRedpacket`: RedPacketReply
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.ClaimRedpacket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimRedpacketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimRedPacketInput** | [**ClaimRedPacketInput**](ClaimRedPacketInput.md) |  | 

### Return type

[**RedPacketReply**](RedPacketReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRedpacket

> CreateRedPacketReply CreateRedpacket(ctx, chain).CreateRedPacketInput(createRedPacketInput).Execute()

RedPacket - Create



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
	createRedPacketInput := *openapiclient.NewCreateRedPacketInput("sol", "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB", int64(10)) // CreateRedPacketInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.CreateRedpacket(context.Background(), chain).CreateRedPacketInput(createRedPacketInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.CreateRedpacket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRedpacket`: CreateRedPacketReply
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.CreateRedpacket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRedpacketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRedPacketInput** | [**CreateRedPacketInput**](CreateRedPacketInput.md) |  | 

### Return type

[**CreateRedPacketReply**](CreateRedPacketReply.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaims

> RedPacketClaimsPage GetClaims(ctx, id).Cursor(cursor).Limit(limit).Direction(direction).Execute()

RedPacket - GetClaims



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
	id := "redpacket123" // string | Redpacket id
	cursor := "cursor_123" // string | Pagination cursor (optional)
	limit := int64(20) // int64 | Number of records per page (optional)
	direction := "desc" // string | Sort direction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.GetClaims(context.Background(), id).Cursor(cursor).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.GetClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClaims`: RedPacketClaimsPage
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.GetClaims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Redpacket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **int64** | Number of records per page | 
 **direction** | **string** | Sort direction | 

### Return type

[**RedPacketClaimsPage**](RedPacketClaimsPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaimsByAddress

> RedPacketClaimsPage GetClaimsByAddress(ctx, address).Cursor(cursor).Limit(limit).Direction(direction).Execute()

RedPacket - GetClaimsByAddress



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
	address := "address_example" // string | 
	cursor := "cursor_123" // string | Pagination cursor (optional)
	limit := int64(20) // int64 | Number of records per page (optional)
	direction := "desc" // string | Sort direction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.GetClaimsByAddress(context.Background(), address).Cursor(cursor).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.GetClaimsByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClaimsByAddress`: RedPacketClaimsPage
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.GetClaimsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **int64** | Number of records per page | 
 **direction** | **string** | Sort direction | 

### Return type

[**RedPacketClaimsPage**](RedPacketClaimsPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedpacket

> RedPacketDTO GetRedpacket(ctx, id).Execute()

RedPacket - Get



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
	id := "redpacket123" // string | Redpacket id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.GetRedpacket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.GetRedpacket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedpacket`: RedPacketDTO
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.GetRedpacket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Redpacket id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedpacketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedPacketDTO**](RedPacketDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedpackets

> RedPacketsPage GetRedpackets(ctx).Cursor(cursor).Limit(limit).Direction(direction).Creator(creator).Chain(chain).Execute()

RedPacket - GetRedpackets



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
	cursor := "cursor_123" // string | Pagination cursor (optional)
	limit := int64(20) // int64 | Number of records per page (optional)
	direction := "desc" // string | Sort direction (optional)
	creator := "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB" // string | Creator wallet address of the red packet (optional)
	chain := "sol" // string | Blockchain network (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.GetRedpackets(context.Background()).Cursor(cursor).Limit(limit).Direction(direction).Creator(creator).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.GetRedpackets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedpackets`: RedPacketsPage
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.GetRedpackets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRedpacketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Pagination cursor | 
 **limit** | **int64** | Number of records per page | 
 **direction** | **string** | Sort direction | 
 **creator** | **string** | Creator wallet address of the red packet | 
 **chain** | **string** | Blockchain network | 

### Return type

[**RedPacketsPage**](RedPacketsPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedpacketsByAddress

> RedPacketsPage GetRedpacketsByAddress(ctx, address).Cursor(cursor).Limit(limit).Direction(direction).Execute()

RedPacket - GetRedpacketsByAddress



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
	address := "address_example" // string | 
	cursor := "cursor_123" // string | Pagination cursor (optional)
	limit := int64(20) // int64 | Number of records per page (optional)
	direction := "desc" // string | Sort direction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.GetRedpacketsByAddress(context.Background(), address).Cursor(cursor).Limit(limit).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.GetRedpacketsByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedpacketsByAddress`: RedPacketsPage
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.GetRedpacketsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedpacketsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Pagination cursor | 
 **limit** | **int64** | Number of records per page | 
 **direction** | **string** | Sort direction | 

### Return type

[**RedPacketsPage**](RedPacketsPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedpacketSend

> RedPacketSendTxResponse RedpacketSend(ctx, chain).RedPacketSendTxInput(redPacketSendTxInput).Execute()

RedPacket - Send



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
	redPacketSendTxInput := *openapiclient.NewRedPacketSendTxInput("AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAEDRgYGpQEDAQIABQcICQoLDA0ODwAAAAAAAAAAAAAQERITFBUWFxgZGhscHR4fAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=") // RedPacketSendTxInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RedPacketAPI.RedpacketSend(context.Background(), chain).RedPacketSendTxInput(redPacketSendTxInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedPacketAPI.RedpacketSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedpacketSend`: RedPacketSendTxResponse
	fmt.Fprintf(os.Stdout, "Response from `RedPacketAPI.RedpacketSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainSymbol**](.md) | A chain name listed in supported networks | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedpacketSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redPacketSendTxInput** | [**RedPacketSendTxInput**](RedPacketSendTxInput.md) |  | 

### Return type

[**RedPacketSendTxResponse**](RedPacketSendTxResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

