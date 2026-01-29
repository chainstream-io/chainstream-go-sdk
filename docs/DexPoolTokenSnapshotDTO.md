# DexPoolTokenSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | Token contract address | 
**TokenDecimals** | **int32** | Token decimals | 
**VaultAmount** | **string** | Raw vault amount (without decimals) | 
**AmountInUsd** | **string** | Liquidity value in USD | 
**AmountInNative** | **string** | Liquidity value in native token (e.g., SOL) | 
**PriceUsd** | **string** | Token price in USD | 
**PriceNative** | **string** | Token price in native token | 

## Methods

### NewDexPoolTokenSnapshotDTO

`func NewDexPoolTokenSnapshotDTO(tokenAddress string, tokenDecimals int32, vaultAmount string, amountInUsd string, amountInNative string, priceUsd string, priceNative string, ) *DexPoolTokenSnapshotDTO`

NewDexPoolTokenSnapshotDTO instantiates a new DexPoolTokenSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexPoolTokenSnapshotDTOWithDefaults

`func NewDexPoolTokenSnapshotDTOWithDefaults() *DexPoolTokenSnapshotDTO`

NewDexPoolTokenSnapshotDTOWithDefaults instantiates a new DexPoolTokenSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *DexPoolTokenSnapshotDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *DexPoolTokenSnapshotDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *DexPoolTokenSnapshotDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetTokenDecimals

`func (o *DexPoolTokenSnapshotDTO) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *DexPoolTokenSnapshotDTO) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *DexPoolTokenSnapshotDTO) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetVaultAmount

`func (o *DexPoolTokenSnapshotDTO) GetVaultAmount() string`

GetVaultAmount returns the VaultAmount field if non-nil, zero value otherwise.

### GetVaultAmountOk

`func (o *DexPoolTokenSnapshotDTO) GetVaultAmountOk() (*string, bool)`

GetVaultAmountOk returns a tuple with the VaultAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAmount

`func (o *DexPoolTokenSnapshotDTO) SetVaultAmount(v string)`

SetVaultAmount sets VaultAmount field to given value.


### GetAmountInUsd

`func (o *DexPoolTokenSnapshotDTO) GetAmountInUsd() string`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *DexPoolTokenSnapshotDTO) GetAmountInUsdOk() (*string, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *DexPoolTokenSnapshotDTO) SetAmountInUsd(v string)`

SetAmountInUsd sets AmountInUsd field to given value.


### GetAmountInNative

`func (o *DexPoolTokenSnapshotDTO) GetAmountInNative() string`

GetAmountInNative returns the AmountInNative field if non-nil, zero value otherwise.

### GetAmountInNativeOk

`func (o *DexPoolTokenSnapshotDTO) GetAmountInNativeOk() (*string, bool)`

GetAmountInNativeOk returns a tuple with the AmountInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInNative

`func (o *DexPoolTokenSnapshotDTO) SetAmountInNative(v string)`

SetAmountInNative sets AmountInNative field to given value.


### GetPriceUsd

`func (o *DexPoolTokenSnapshotDTO) GetPriceUsd() string`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *DexPoolTokenSnapshotDTO) GetPriceUsdOk() (*string, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *DexPoolTokenSnapshotDTO) SetPriceUsd(v string)`

SetPriceUsd sets PriceUsd field to given value.


### GetPriceNative

`func (o *DexPoolTokenSnapshotDTO) GetPriceNative() string`

GetPriceNative returns the PriceNative field if non-nil, zero value otherwise.

### GetPriceNativeOk

`func (o *DexPoolTokenSnapshotDTO) GetPriceNativeOk() (*string, bool)`

GetPriceNativeOk returns a tuple with the PriceNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNative

`func (o *DexPoolTokenSnapshotDTO) SetPriceNative(v string)`

SetPriceNative sets PriceNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


