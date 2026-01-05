# EstimateGasLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasLimit** | **string** | Estimated gas limit (hex string) | 
**Chain** | **string** | Chain symbol | 

## Methods

### NewEstimateGasLimitResponse

`func NewEstimateGasLimitResponse(gasLimit string, chain string, ) *EstimateGasLimitResponse`

NewEstimateGasLimitResponse instantiates a new EstimateGasLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateGasLimitResponseWithDefaults

`func NewEstimateGasLimitResponseWithDefaults() *EstimateGasLimitResponse`

NewEstimateGasLimitResponseWithDefaults instantiates a new EstimateGasLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasLimit

`func (o *EstimateGasLimitResponse) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EstimateGasLimitResponse) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EstimateGasLimitResponse) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetChain

`func (o *EstimateGasLimitResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *EstimateGasLimitResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *EstimateGasLimitResponse) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


