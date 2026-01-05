# AlertDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertLevel** | **string** | DTO.KYT.ALERT_DETAIL.ALERT_LEVEL | 
**Service** | **string** | DTO.KYT.ALERT_DETAIL.SERVICE | 
**ExternalId** | **string** | DTO.KYT.ALERT_DETAIL.EXTERNAL_ID | 
**AlertAmount** | **float32** | DTO.KYT.ALERT_DETAIL.ALERT_AMOUNT | 
**ExposureType** | **string** | DTO.KYT.ALERT_DETAIL.EXPOSURE_TYPE | 
**CategoryId** | **float32** | DTO.KYT.ALERT_DETAIL.CATEGORY_ID | 
**Memo** | **string** | DTO.KYT.ALERT_DETAIL.MEMO | 

## Methods

### NewAlertDetail

`func NewAlertDetail(alertLevel string, service string, externalId string, alertAmount float32, exposureType string, categoryId float32, memo string, ) *AlertDetail`

NewAlertDetail instantiates a new AlertDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDetailWithDefaults

`func NewAlertDetailWithDefaults() *AlertDetail`

NewAlertDetailWithDefaults instantiates a new AlertDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertLevel

`func (o *AlertDetail) GetAlertLevel() string`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *AlertDetail) GetAlertLevelOk() (*string, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *AlertDetail) SetAlertLevel(v string)`

SetAlertLevel sets AlertLevel field to given value.


### GetService

`func (o *AlertDetail) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AlertDetail) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AlertDetail) SetService(v string)`

SetService sets Service field to given value.


### GetExternalId

`func (o *AlertDetail) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AlertDetail) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AlertDetail) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAlertAmount

`func (o *AlertDetail) GetAlertAmount() float32`

GetAlertAmount returns the AlertAmount field if non-nil, zero value otherwise.

### GetAlertAmountOk

`func (o *AlertDetail) GetAlertAmountOk() (*float32, bool)`

GetAlertAmountOk returns a tuple with the AlertAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertAmount

`func (o *AlertDetail) SetAlertAmount(v float32)`

SetAlertAmount sets AlertAmount field to given value.


### GetExposureType

`func (o *AlertDetail) GetExposureType() string`

GetExposureType returns the ExposureType field if non-nil, zero value otherwise.

### GetExposureTypeOk

`func (o *AlertDetail) GetExposureTypeOk() (*string, bool)`

GetExposureTypeOk returns a tuple with the ExposureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureType

`func (o *AlertDetail) SetExposureType(v string)`

SetExposureType sets ExposureType field to given value.


### GetCategoryId

`func (o *AlertDetail) GetCategoryId() float32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *AlertDetail) GetCategoryIdOk() (*float32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *AlertDetail) SetCategoryId(v float32)`

SetCategoryId sets CategoryId field to given value.


### GetMemo

`func (o *AlertDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AlertDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AlertDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


