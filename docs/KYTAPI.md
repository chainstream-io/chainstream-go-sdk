# \KYTAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressRisk**](KYTAPI.md#GetAddressRisk) | **Get** /v1/kyt/addresses/{address}/risk | CONTROLLER.KYT.GET_ADDRESS_RISK.SUMMARY
[**GetTransferAlerts**](KYTAPI.md#GetTransferAlerts) | **Get** /v1/kyt/transfers/{transferId}/alerts | CONTROLLER.KYT.GET_TRANSFER_ALERTS.SUMMARY
[**GetTransferDirectExposure**](KYTAPI.md#GetTransferDirectExposure) | **Get** /v1/kyt/transfers/{transferId}/exposures/direct | CONTROLLER.KYT.GET_TRANSFER_DIRECT_EXPOSURE.SUMMARY
[**GetTransferNetworkIdentifications**](KYTAPI.md#GetTransferNetworkIdentifications) | **Get** /v1/kyt/transfers/{transferId}/network-identifications | CONTROLLER.KYT.GET_TRANSFER_NETWORK_IDENTIFICATIONS.SUMMARY
[**GetTransferSummary**](KYTAPI.md#GetTransferSummary) | **Get** /v1/kyt/transfers/{transferId}/summary | CONTROLLER.KYT.GET_TRANSFER_SUMMARY.SUMMARY
[**GetWithdrawalAddressIdentifications**](KYTAPI.md#GetWithdrawalAddressIdentifications) | **Get** /v1/kyt/withdrawal/{withdrawalId}/address-identifications | CONTROLLER.KYT.GET_WITHDRAWAL_ADDRESS_IDENTIFICATIONS.SUMMARY
[**GetWithdrawalAlerts**](KYTAPI.md#GetWithdrawalAlerts) | **Get** /v1/kyt/withdrawal/{withdrawalId}/alerts | CONTROLLER.KYT.GET_WITHDRAWAL_ALERTS.SUMMARY
[**GetWithdrawalDirectExposure**](KYTAPI.md#GetWithdrawalDirectExposure) | **Get** /v1/kyt/withdrawal/{withdrawalId}/exposures/direct | CONTROLLER.KYT.GET_WITHDRAWAL_DIRECT_EXPOSURE.SUMMARY
[**GetWithdrawalFraudAssessment**](KYTAPI.md#GetWithdrawalFraudAssessment) | **Get** /v1/kyt/withdrawal/{withdrawalId}/fraud-assessment | CONTROLLER.KYT.GET_WITHDRAWAL_FRAUD_ASSESSMENT.SUMMARY
[**GetWithdrawalNetworkIdentifications**](KYTAPI.md#GetWithdrawalNetworkIdentifications) | **Get** /v1/kyt/withdrawal/{withdrawalId}/network-identifications | CONTROLLER.KYT.GET_WITHDRAWAL_NETWORK_IDENTIFICATIONS.SUMMARY
[**GetWithdrawalSummary**](KYTAPI.md#GetWithdrawalSummary) | **Get** /v1/kyt/withdrawal/{withdrawalId}/summary | CONTROLLER.KYT.GET_WITHDRAWAL_SUMMARY.SUMMARY
[**RegisterAddress**](KYTAPI.md#RegisterAddress) | **Post** /v1/kyt/address | CONTROLLER.KYT.REGISTER_ADDRESS.SUMMARY
[**RegisterTransfer**](KYTAPI.md#RegisterTransfer) | **Post** /v1/kyt/transfer | CONTROLLER.KYT.REGISTER_TRANSFER.SUMMARY
[**RegisterWithdrawal**](KYTAPI.md#RegisterWithdrawal) | **Post** /v1/kyt/withdrawal | CONTROLLER.KYT.REGISTER_WITHDRAWAL.SUMMARY



## GetAddressRisk

> AddressRiskResponseDTO GetAddressRisk(ctx, address).Execute()

CONTROLLER.KYT.GET_ADDRESS_RISK.SUMMARY



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
	address := "0x0038AC785dfB6C82b2c9A7B3B6854e08a10cb9f1" // string | CONTROLLER.KYT.PARAM.ADDRESS

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetAddressRisk(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetAddressRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressRisk`: AddressRiskResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetAddressRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | CONTROLLER.KYT.PARAM.ADDRESS | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressRiskResponseDTO**](AddressRiskResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferAlerts

> TransferAlertsResponseDTO GetTransferAlerts(ctx, transferId).Execute()

CONTROLLER.KYT.GET_TRANSFER_ALERTS.SUMMARY



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
	transferId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.TRANSFER_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetTransferAlerts(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetTransferAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransferAlerts`: TransferAlertsResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetTransferAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | CONTROLLER.KYT.PARAM.TRANSFER_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferAlertsResponseDTO**](TransferAlertsResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferDirectExposure

> TransferDirectExposureResponseDTO GetTransferDirectExposure(ctx, transferId).Execute()

CONTROLLER.KYT.GET_TRANSFER_DIRECT_EXPOSURE.SUMMARY



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
	transferId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.TRANSFER_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetTransferDirectExposure(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetTransferDirectExposure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransferDirectExposure`: TransferDirectExposureResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetTransferDirectExposure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | CONTROLLER.KYT.PARAM.TRANSFER_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferDirectExposureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferDirectExposureResponseDTO**](TransferDirectExposureResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferNetworkIdentifications

> TransferNetworkIdentificationsResponseDTO GetTransferNetworkIdentifications(ctx, transferId).Execute()

CONTROLLER.KYT.GET_TRANSFER_NETWORK_IDENTIFICATIONS.SUMMARY



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
	transferId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.TRANSFER_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetTransferNetworkIdentifications(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetTransferNetworkIdentifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransferNetworkIdentifications`: TransferNetworkIdentificationsResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetTransferNetworkIdentifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | CONTROLLER.KYT.PARAM.TRANSFER_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferNetworkIdentificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferNetworkIdentificationsResponseDTO**](TransferNetworkIdentificationsResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferSummary

> TransferBaseResponseDTO GetTransferSummary(ctx, transferId).Execute()

CONTROLLER.KYT.GET_TRANSFER_SUMMARY.SUMMARY



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
	transferId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.TRANSFER_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetTransferSummary(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetTransferSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransferSummary`: TransferBaseResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetTransferSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | CONTROLLER.KYT.PARAM.TRANSFER_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferBaseResponseDTO**](TransferBaseResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalAddressIdentifications

> WithdrawalAddressIdentificationsResponseDTO GetWithdrawalAddressIdentifications(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_ADDRESS_IDENTIFICATIONS.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalAddressIdentifications(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalAddressIdentifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalAddressIdentifications`: WithdrawalAddressIdentificationsResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalAddressIdentifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalAddressIdentificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WithdrawalAddressIdentificationsResponseDTO**](WithdrawalAddressIdentificationsResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalAlerts

> TransferAlertsResponseDTO GetWithdrawalAlerts(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_ALERTS.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalAlerts(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalAlerts`: TransferAlertsResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferAlertsResponseDTO**](TransferAlertsResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalDirectExposure

> TransferDirectExposureResponseDTO GetWithdrawalDirectExposure(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_DIRECT_EXPOSURE.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalDirectExposure(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalDirectExposure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalDirectExposure`: TransferDirectExposureResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalDirectExposure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalDirectExposureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferDirectExposureResponseDTO**](TransferDirectExposureResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalFraudAssessment

> WithdrawalFraudAssessmentResponseDTO GetWithdrawalFraudAssessment(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_FRAUD_ASSESSMENT.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalFraudAssessment(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalFraudAssessment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalFraudAssessment`: WithdrawalFraudAssessmentResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalFraudAssessment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalFraudAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WithdrawalFraudAssessmentResponseDTO**](WithdrawalFraudAssessmentResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalNetworkIdentifications

> TransferNetworkIdentificationsResponseDTO GetWithdrawalNetworkIdentifications(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_NETWORK_IDENTIFICATIONS.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalNetworkIdentifications(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalNetworkIdentifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalNetworkIdentifications`: TransferNetworkIdentificationsResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalNetworkIdentifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalNetworkIdentificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferNetworkIdentificationsResponseDTO**](TransferNetworkIdentificationsResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalSummary

> WithdrawalBaseResponseDTO GetWithdrawalSummary(ctx, withdrawalId).Execute()

CONTROLLER.KYT.GET_WITHDRAWAL_SUMMARY.SUMMARY



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
	withdrawalId := "123e4567-e89b-12d3-a456-426614174000" // string | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.GetWithdrawalSummary(context.Background(), withdrawalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.GetWithdrawalSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWithdrawalSummary`: WithdrawalBaseResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.GetWithdrawalSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** | CONTROLLER.KYT.PARAM.WITHDRAWAL_ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WithdrawalBaseResponseDTO**](WithdrawalBaseResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAddress

> RegisterAddressResponseDTO RegisterAddress(ctx).RegisterAddressRequest(registerAddressRequest).Execute()

CONTROLLER.KYT.REGISTER_ADDRESS.SUMMARY



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
	registerAddressRequest := *openapiclient.NewRegisterAddressRequest("0x0038AC785dfB6C82b2c9A7B3B6854e08a10cb9f1") // RegisterAddressRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.RegisterAddress(context.Background()).RegisterAddressRequest(registerAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.RegisterAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAddress`: RegisterAddressResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.RegisterAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerAddressRequest** | [**RegisterAddressRequest**](RegisterAddressRequest.md) |  | 

### Return type

[**RegisterAddressResponseDTO**](RegisterAddressResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTransfer

> TransferBaseResponseDTO RegisterTransfer(ctx).KYTRegisterTransferRequest(kYTRegisterTransferRequest).Execute()

CONTROLLER.KYT.REGISTER_TRANSFER.SUMMARY



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
	kYTRegisterTransferRequest := *openapiclient.NewKYTRegisterTransferRequest("Solana", "SOL", "39z5QAprVrzaFzfHu1JHPgBf9dSqYdNYhH31d3PEd4hWiWL1LML7qCct5MHGxaRAgjjj1nC3XUyLwtzGQmYqUk4y:7xfh4GGMbM3bjpWsTsmBTCfSyuvXFuqnC89fShAqk5Nj", "sent") // KYTRegisterTransferRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.RegisterTransfer(context.Background()).KYTRegisterTransferRequest(kYTRegisterTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.RegisterTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterTransfer`: TransferBaseResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.RegisterTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kYTRegisterTransferRequest** | [**KYTRegisterTransferRequest**](KYTRegisterTransferRequest.md) |  | 

### Return type

[**TransferBaseResponseDTO**](TransferBaseResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterWithdrawal

> WithdrawalBaseResponseDTO RegisterWithdrawal(ctx).KYTRegisterWithdrawalRequest(kYTRegisterWithdrawalRequest).Execute()

CONTROLLER.KYT.REGISTER_WITHDRAWAL.SUMMARY



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
	kYTRegisterWithdrawalRequest := *openapiclient.NewKYTRegisterWithdrawalRequest("Solana", "SOL", "D1Mc6j9xQWgR1o1Z7yU5nVVXFQiAYx7FG9AW1aVfwrUM", float32(5), "2026-01-04T17:25:40.008307") // KYTRegisterWithdrawalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KYTAPI.RegisterWithdrawal(context.Background()).KYTRegisterWithdrawalRequest(kYTRegisterWithdrawalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KYTAPI.RegisterWithdrawal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterWithdrawal`: WithdrawalBaseResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `KYTAPI.RegisterWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kYTRegisterWithdrawalRequest** | [**KYTRegisterWithdrawalRequest**](KYTRegisterWithdrawalRequest.md) |  | 

### Return type

[**WithdrawalBaseResponseDTO**](WithdrawalBaseResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

