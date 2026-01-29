# DexPoolSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolAddress** | **string** | Pool address | 
**SnapshotTime** | **int64** | Actual snapshot Unix timestamp (seconds) | 
**TvlInUsd** | **string** | Total value locked in USD | 
**TvlInNative** | **string** | TVL in SOL | 
**TokenA** | [**DexPoolTokenSnapshotDTO**](DexPoolTokenSnapshotDTO.md) | First token snapshot data | 
**TokenB** | [**DexPoolTokenSnapshotDTO**](DexPoolTokenSnapshotDTO.md) | Second token snapshot data | 
**BlockHeight** | **int64** | Block height at snapshot time | 
**BlockSlot** | Pointer to **int64** | Block slot at snapshot time (Solana only) | [optional] 
**BlockTimestamp** | **int64** | Block Unix timestamp at snapshot time (seconds) | 

## Methods

### NewDexPoolSnapshotDTO

`func NewDexPoolSnapshotDTO(poolAddress string, snapshotTime int64, tvlInUsd string, tvlInNative string, tokenA DexPoolTokenSnapshotDTO, tokenB DexPoolTokenSnapshotDTO, blockHeight int64, blockTimestamp int64, ) *DexPoolSnapshotDTO`

NewDexPoolSnapshotDTO instantiates a new DexPoolSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexPoolSnapshotDTOWithDefaults

`func NewDexPoolSnapshotDTOWithDefaults() *DexPoolSnapshotDTO`

NewDexPoolSnapshotDTOWithDefaults instantiates a new DexPoolSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolAddress

`func (o *DexPoolSnapshotDTO) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *DexPoolSnapshotDTO) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *DexPoolSnapshotDTO) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetSnapshotTime

`func (o *DexPoolSnapshotDTO) GetSnapshotTime() int64`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *DexPoolSnapshotDTO) GetSnapshotTimeOk() (*int64, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *DexPoolSnapshotDTO) SetSnapshotTime(v int64)`

SetSnapshotTime sets SnapshotTime field to given value.


### GetTvlInUsd

`func (o *DexPoolSnapshotDTO) GetTvlInUsd() string`

GetTvlInUsd returns the TvlInUsd field if non-nil, zero value otherwise.

### GetTvlInUsdOk

`func (o *DexPoolSnapshotDTO) GetTvlInUsdOk() (*string, bool)`

GetTvlInUsdOk returns a tuple with the TvlInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInUsd

`func (o *DexPoolSnapshotDTO) SetTvlInUsd(v string)`

SetTvlInUsd sets TvlInUsd field to given value.


### GetTvlInNative

`func (o *DexPoolSnapshotDTO) GetTvlInNative() string`

GetTvlInNative returns the TvlInNative field if non-nil, zero value otherwise.

### GetTvlInNativeOk

`func (o *DexPoolSnapshotDTO) GetTvlInNativeOk() (*string, bool)`

GetTvlInNativeOk returns a tuple with the TvlInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInNative

`func (o *DexPoolSnapshotDTO) SetTvlInNative(v string)`

SetTvlInNative sets TvlInNative field to given value.


### GetTokenA

`func (o *DexPoolSnapshotDTO) GetTokenA() DexPoolTokenSnapshotDTO`

GetTokenA returns the TokenA field if non-nil, zero value otherwise.

### GetTokenAOk

`func (o *DexPoolSnapshotDTO) GetTokenAOk() (*DexPoolTokenSnapshotDTO, bool)`

GetTokenAOk returns a tuple with the TokenA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenA

`func (o *DexPoolSnapshotDTO) SetTokenA(v DexPoolTokenSnapshotDTO)`

SetTokenA sets TokenA field to given value.


### GetTokenB

`func (o *DexPoolSnapshotDTO) GetTokenB() DexPoolTokenSnapshotDTO`

GetTokenB returns the TokenB field if non-nil, zero value otherwise.

### GetTokenBOk

`func (o *DexPoolSnapshotDTO) GetTokenBOk() (*DexPoolTokenSnapshotDTO, bool)`

GetTokenBOk returns a tuple with the TokenB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenB

`func (o *DexPoolSnapshotDTO) SetTokenB(v DexPoolTokenSnapshotDTO)`

SetTokenB sets TokenB field to given value.


### GetBlockHeight

`func (o *DexPoolSnapshotDTO) GetBlockHeight() int64`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *DexPoolSnapshotDTO) GetBlockHeightOk() (*int64, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *DexPoolSnapshotDTO) SetBlockHeight(v int64)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockSlot

`func (o *DexPoolSnapshotDTO) GetBlockSlot() int64`

GetBlockSlot returns the BlockSlot field if non-nil, zero value otherwise.

### GetBlockSlotOk

`func (o *DexPoolSnapshotDTO) GetBlockSlotOk() (*int64, bool)`

GetBlockSlotOk returns a tuple with the BlockSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSlot

`func (o *DexPoolSnapshotDTO) SetBlockSlot(v int64)`

SetBlockSlot sets BlockSlot field to given value.

### HasBlockSlot

`func (o *DexPoolSnapshotDTO) HasBlockSlot() bool`

HasBlockSlot returns a boolean if a field has been set.

### GetBlockTimestamp

`func (o *DexPoolSnapshotDTO) GetBlockTimestamp() int64`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *DexPoolSnapshotDTO) GetBlockTimestampOk() (*int64, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *DexPoolSnapshotDTO) SetBlockTimestamp(v int64)`

SetBlockTimestamp sets BlockTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


