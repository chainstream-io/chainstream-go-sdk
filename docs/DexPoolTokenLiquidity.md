# DexPoolTokenLiquidity

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

### NewDexPoolTokenLiquidity

`func NewDexPoolTokenLiquidity(tokenAddress string, tokenDecimals int32, vaultAmount string, amountInUsd string, amountInNative string, priceUsd string, priceNative string, ) *DexPoolTokenLiquidity`

NewDexPoolTokenLiquidity instantiates a new DexPoolTokenLiquidity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexPoolTokenLiquidityWithDefaults

`func NewDexPoolTokenLiquidityWithDefaults() *DexPoolTokenLiquidity`

NewDexPoolTokenLiquidityWithDefaults instantiates a new DexPoolTokenLiquidity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *DexPoolTokenLiquidity) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *DexPoolTokenLiquidity) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *DexPoolTokenLiquidity) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetTokenDecimals

`func (o *DexPoolTokenLiquidity) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *DexPoolTokenLiquidity) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *DexPoolTokenLiquidity) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetVaultAmount

`func (o *DexPoolTokenLiquidity) GetVaultAmount() string`

GetVaultAmount returns the VaultAmount field if non-nil, zero value otherwise.

### GetVaultAmountOk

`func (o *DexPoolTokenLiquidity) GetVaultAmountOk() (*string, bool)`

GetVaultAmountOk returns a tuple with the VaultAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAmount

`func (o *DexPoolTokenLiquidity) SetVaultAmount(v string)`

SetVaultAmount sets VaultAmount field to given value.


### GetAmountInUsd

`func (o *DexPoolTokenLiquidity) GetAmountInUsd() string`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *DexPoolTokenLiquidity) GetAmountInUsdOk() (*string, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *DexPoolTokenLiquidity) SetAmountInUsd(v string)`

SetAmountInUsd sets AmountInUsd field to given value.


### GetAmountInNative

`func (o *DexPoolTokenLiquidity) GetAmountInNative() string`

GetAmountInNative returns the AmountInNative field if non-nil, zero value otherwise.

### GetAmountInNativeOk

`func (o *DexPoolTokenLiquidity) GetAmountInNativeOk() (*string, bool)`

GetAmountInNativeOk returns a tuple with the AmountInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInNative

`func (o *DexPoolTokenLiquidity) SetAmountInNative(v string)`

SetAmountInNative sets AmountInNative field to given value.


### GetPriceUsd

`func (o *DexPoolTokenLiquidity) GetPriceUsd() string`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *DexPoolTokenLiquidity) GetPriceUsdOk() (*string, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *DexPoolTokenLiquidity) SetPriceUsd(v string)`

SetPriceUsd sets PriceUsd field to given value.


### GetPriceNative

`func (o *DexPoolTokenLiquidity) GetPriceNative() string`

GetPriceNative returns the PriceNative field if non-nil, zero value otherwise.

### GetPriceNativeOk

`func (o *DexPoolTokenLiquidity) GetPriceNativeOk() (*string, bool)`

GetPriceNativeOk returns a tuple with the PriceNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNative

`func (o *DexPoolTokenLiquidity) SetPriceNative(v string)`

SetPriceNative sets PriceNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


