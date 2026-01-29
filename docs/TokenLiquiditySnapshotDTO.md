# TokenLiquiditySnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotTime** | **int64** | Actual snapshot Unix timestamp (seconds) | 
**MaxLiquidityPoolAddress** | **string** | Pool address with maximum liquidity | 
**MaxLiquidityInUsd** | **string** | Maximum single pool liquidity in USD | 
**MaxLiquidityInNative** | **string** | Maximum single pool liquidity in native token | 
**TotalLiquidityInUsd** | **string** | Total liquidity across all pools in USD | 
**TotalLiquidityInNative** | **string** | Total liquidity across all pools in native token | 
**PoolCount** | **int32** | Number of pools for this token | 
**PriceUsd** | **string** | Token price in USD at snapshot time | 
**PriceNative** | **string** | Token price in native token at snapshot time | 
**CalculatedAt** | **int64** | Unix timestamp when liquidity was calculated (seconds) | 

## Methods

### NewTokenLiquiditySnapshotDTO

`func NewTokenLiquiditySnapshotDTO(snapshotTime int64, maxLiquidityPoolAddress string, maxLiquidityInUsd string, maxLiquidityInNative string, totalLiquidityInUsd string, totalLiquidityInNative string, poolCount int32, priceUsd string, priceNative string, calculatedAt int64, ) *TokenLiquiditySnapshotDTO`

NewTokenLiquiditySnapshotDTO instantiates a new TokenLiquiditySnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLiquiditySnapshotDTOWithDefaults

`func NewTokenLiquiditySnapshotDTOWithDefaults() *TokenLiquiditySnapshotDTO`

NewTokenLiquiditySnapshotDTOWithDefaults instantiates a new TokenLiquiditySnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotTime

`func (o *TokenLiquiditySnapshotDTO) GetSnapshotTime() int64`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *TokenLiquiditySnapshotDTO) GetSnapshotTimeOk() (*int64, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *TokenLiquiditySnapshotDTO) SetSnapshotTime(v int64)`

SetSnapshotTime sets SnapshotTime field to given value.


### GetMaxLiquidityPoolAddress

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityPoolAddress() string`

GetMaxLiquidityPoolAddress returns the MaxLiquidityPoolAddress field if non-nil, zero value otherwise.

### GetMaxLiquidityPoolAddressOk

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityPoolAddressOk() (*string, bool)`

GetMaxLiquidityPoolAddressOk returns a tuple with the MaxLiquidityPoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLiquidityPoolAddress

`func (o *TokenLiquiditySnapshotDTO) SetMaxLiquidityPoolAddress(v string)`

SetMaxLiquidityPoolAddress sets MaxLiquidityPoolAddress field to given value.


### GetMaxLiquidityInUsd

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityInUsd() string`

GetMaxLiquidityInUsd returns the MaxLiquidityInUsd field if non-nil, zero value otherwise.

### GetMaxLiquidityInUsdOk

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityInUsdOk() (*string, bool)`

GetMaxLiquidityInUsdOk returns a tuple with the MaxLiquidityInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLiquidityInUsd

`func (o *TokenLiquiditySnapshotDTO) SetMaxLiquidityInUsd(v string)`

SetMaxLiquidityInUsd sets MaxLiquidityInUsd field to given value.


### GetMaxLiquidityInNative

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityInNative() string`

GetMaxLiquidityInNative returns the MaxLiquidityInNative field if non-nil, zero value otherwise.

### GetMaxLiquidityInNativeOk

`func (o *TokenLiquiditySnapshotDTO) GetMaxLiquidityInNativeOk() (*string, bool)`

GetMaxLiquidityInNativeOk returns a tuple with the MaxLiquidityInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLiquidityInNative

`func (o *TokenLiquiditySnapshotDTO) SetMaxLiquidityInNative(v string)`

SetMaxLiquidityInNative sets MaxLiquidityInNative field to given value.


### GetTotalLiquidityInUsd

`func (o *TokenLiquiditySnapshotDTO) GetTotalLiquidityInUsd() string`

GetTotalLiquidityInUsd returns the TotalLiquidityInUsd field if non-nil, zero value otherwise.

### GetTotalLiquidityInUsdOk

`func (o *TokenLiquiditySnapshotDTO) GetTotalLiquidityInUsdOk() (*string, bool)`

GetTotalLiquidityInUsdOk returns a tuple with the TotalLiquidityInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiquidityInUsd

`func (o *TokenLiquiditySnapshotDTO) SetTotalLiquidityInUsd(v string)`

SetTotalLiquidityInUsd sets TotalLiquidityInUsd field to given value.


### GetTotalLiquidityInNative

`func (o *TokenLiquiditySnapshotDTO) GetTotalLiquidityInNative() string`

GetTotalLiquidityInNative returns the TotalLiquidityInNative field if non-nil, zero value otherwise.

### GetTotalLiquidityInNativeOk

`func (o *TokenLiquiditySnapshotDTO) GetTotalLiquidityInNativeOk() (*string, bool)`

GetTotalLiquidityInNativeOk returns a tuple with the TotalLiquidityInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiquidityInNative

`func (o *TokenLiquiditySnapshotDTO) SetTotalLiquidityInNative(v string)`

SetTotalLiquidityInNative sets TotalLiquidityInNative field to given value.


### GetPoolCount

`func (o *TokenLiquiditySnapshotDTO) GetPoolCount() int32`

GetPoolCount returns the PoolCount field if non-nil, zero value otherwise.

### GetPoolCountOk

`func (o *TokenLiquiditySnapshotDTO) GetPoolCountOk() (*int32, bool)`

GetPoolCountOk returns a tuple with the PoolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolCount

`func (o *TokenLiquiditySnapshotDTO) SetPoolCount(v int32)`

SetPoolCount sets PoolCount field to given value.


### GetPriceUsd

`func (o *TokenLiquiditySnapshotDTO) GetPriceUsd() string`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *TokenLiquiditySnapshotDTO) GetPriceUsdOk() (*string, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *TokenLiquiditySnapshotDTO) SetPriceUsd(v string)`

SetPriceUsd sets PriceUsd field to given value.


### GetPriceNative

`func (o *TokenLiquiditySnapshotDTO) GetPriceNative() string`

GetPriceNative returns the PriceNative field if non-nil, zero value otherwise.

### GetPriceNativeOk

`func (o *TokenLiquiditySnapshotDTO) GetPriceNativeOk() (*string, bool)`

GetPriceNativeOk returns a tuple with the PriceNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNative

`func (o *TokenLiquiditySnapshotDTO) SetPriceNative(v string)`

SetPriceNative sets PriceNative field to given value.


### GetCalculatedAt

`func (o *TokenLiquiditySnapshotDTO) GetCalculatedAt() int64`

GetCalculatedAt returns the CalculatedAt field if non-nil, zero value otherwise.

### GetCalculatedAtOk

`func (o *TokenLiquiditySnapshotDTO) GetCalculatedAtOk() (*int64, bool)`

GetCalculatedAtOk returns a tuple with the CalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedAt

`func (o *TokenLiquiditySnapshotDTO) SetCalculatedAt(v int64)`

SetCalculatedAt sets CalculatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


