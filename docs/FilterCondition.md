# FilterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | DTO.TOKEN.FILTER.FIELD | [optional] 
**Min** | Pointer to **string** | DTO.TOKEN.FILTER.MIN | [optional] 
**Max** | Pointer to **string** | DTO.TOKEN.FILTER.MAX | [optional] 

## Methods

### NewFilterCondition

`func NewFilterCondition() *FilterCondition`

NewFilterCondition instantiates a new FilterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterConditionWithDefaults

`func NewFilterConditionWithDefaults() *FilterCondition`

NewFilterConditionWithDefaults instantiates a new FilterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FilterCondition) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMin

`func (o *FilterCondition) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *FilterCondition) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *FilterCondition) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *FilterCondition) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *FilterCondition) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FilterCondition) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FilterCondition) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *FilterCondition) HasMax() bool`

HasMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


