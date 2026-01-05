# GasPriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** | Current gas price in wei (hex string) | 
**Chain** | **string** | Chain symbol | 

## Methods

### NewGasPriceResponse

`func NewGasPriceResponse(gasPrice string, chain string, ) *GasPriceResponse`

NewGasPriceResponse instantiates a new GasPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasPriceResponseWithDefaults

`func NewGasPriceResponseWithDefaults() *GasPriceResponse`

NewGasPriceResponseWithDefaults instantiates a new GasPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *GasPriceResponse) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GasPriceResponse) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GasPriceResponse) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetChain

`func (o *GasPriceResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GasPriceResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GasPriceResponse) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


