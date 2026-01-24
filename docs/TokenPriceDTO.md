# TokenPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | DTO.TOKEN.PRICE.TOKEN_ADDRESS | 
**PriceInUsd** | **string** | DTO.TOKEN.PRICE.PRICE_IN_USD | 
**PriceInNative** | **string** | DTO.TOKEN.PRICE.PRICE_IN_NATIVE | 
**Timestamp** | **int64** | DTO.TOKEN.PRICE.TIMESTAMP | 

## Methods

### NewTokenPriceDTO

`func NewTokenPriceDTO(tokenAddress string, priceInUsd string, priceInNative string, timestamp int64, ) *TokenPriceDTO`

NewTokenPriceDTO instantiates a new TokenPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPriceDTOWithDefaults

`func NewTokenPriceDTOWithDefaults() *TokenPriceDTO`

NewTokenPriceDTOWithDefaults instantiates a new TokenPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *TokenPriceDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenPriceDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenPriceDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetPriceInUsd

`func (o *TokenPriceDTO) GetPriceInUsd() string`

GetPriceInUsd returns the PriceInUsd field if non-nil, zero value otherwise.

### GetPriceInUsdOk

`func (o *TokenPriceDTO) GetPriceInUsdOk() (*string, bool)`

GetPriceInUsdOk returns a tuple with the PriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInUsd

`func (o *TokenPriceDTO) SetPriceInUsd(v string)`

SetPriceInUsd sets PriceInUsd field to given value.


### GetPriceInNative

`func (o *TokenPriceDTO) GetPriceInNative() string`

GetPriceInNative returns the PriceInNative field if non-nil, zero value otherwise.

### GetPriceInNativeOk

`func (o *TokenPriceDTO) GetPriceInNativeOk() (*string, bool)`

GetPriceInNativeOk returns a tuple with the PriceInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInNative

`func (o *TokenPriceDTO) SetPriceInNative(v string)`

SetPriceInNative sets PriceInNative field to given value.


### GetTimestamp

`func (o *TokenPriceDTO) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TokenPriceDTO) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TokenPriceDTO) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


