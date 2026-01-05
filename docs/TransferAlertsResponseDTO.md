# TransferAlertsResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | [**[]AlertDetail**](AlertDetail.md) | DTO.KYT.TRANSFER_ALERTS_RESPONSE.ALERTS | 

## Methods

### NewTransferAlertsResponseDTO

`func NewTransferAlertsResponseDTO(alerts []AlertDetail, ) *TransferAlertsResponseDTO`

NewTransferAlertsResponseDTO instantiates a new TransferAlertsResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAlertsResponseDTOWithDefaults

`func NewTransferAlertsResponseDTOWithDefaults() *TransferAlertsResponseDTO`

NewTransferAlertsResponseDTOWithDefaults instantiates a new TransferAlertsResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *TransferAlertsResponseDTO) GetAlerts() []AlertDetail`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *TransferAlertsResponseDTO) GetAlertsOk() (*[]AlertDetail, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *TransferAlertsResponseDTO) SetAlerts(v []AlertDetail)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


