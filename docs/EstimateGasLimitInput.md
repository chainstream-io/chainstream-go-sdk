# EstimateGasLimitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | From address | 
**To** | **string** | To address | 
**Data** | **string** | Transaction data (hex) | 
**Value** | Pointer to **string** | Value to send (in wei, hex string) | [optional] 

## Methods

### NewEstimateGasLimitInput

`func NewEstimateGasLimitInput(from string, to string, data string, ) *EstimateGasLimitInput`

NewEstimateGasLimitInput instantiates a new EstimateGasLimitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateGasLimitInputWithDefaults

`func NewEstimateGasLimitInputWithDefaults() *EstimateGasLimitInput`

NewEstimateGasLimitInputWithDefaults instantiates a new EstimateGasLimitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *EstimateGasLimitInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EstimateGasLimitInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EstimateGasLimitInput) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *EstimateGasLimitInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EstimateGasLimitInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EstimateGasLimitInput) SetTo(v string)`

SetTo sets To field to given value.


### GetData

`func (o *EstimateGasLimitInput) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EstimateGasLimitInput) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EstimateGasLimitInput) SetData(v string)`

SetData sets Data field to given value.


### GetValue

`func (o *EstimateGasLimitInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EstimateGasLimitInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EstimateGasLimitInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EstimateGasLimitInput) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


