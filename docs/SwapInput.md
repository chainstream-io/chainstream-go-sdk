# SwapInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** | DEX identifier for the trade | 
**UserAddress** | **string** | Public key of the wallet initiating the transaction | 
**PriorityFee** | Pointer to **string** | Priority fee in SOL to increase transaction processing speed | [optional] 
**PoolAddress** | Pointer to **string** | DEX pool address | [optional] 
**Amount** | **string** | Amount to swap. Use \&quot;auto\&quot; for full balance or percentage like \&quot;50%\&quot; | 
**SwapMode** | **string** | Swap direction mode | 
**Slippage** | **int64** | Slippage tolerance percentage | [default to 10]
**InputMint** | Pointer to **string** | Input Mint, the base token address | [optional] 
**OutputMint** | Pointer to **string** | Ouput Mint, the quote token address | [optional] 

## Methods

### NewSwapInput

`func NewSwapInput(dex string, userAddress string, amount string, swapMode string, slippage int64, ) *SwapInput`

NewSwapInput instantiates a new SwapInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapInputWithDefaults

`func NewSwapInputWithDefaults() *SwapInput`

NewSwapInputWithDefaults instantiates a new SwapInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *SwapInput) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *SwapInput) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *SwapInput) SetDex(v string)`

SetDex sets Dex field to given value.


### GetUserAddress

`func (o *SwapInput) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *SwapInput) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *SwapInput) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetPriorityFee

`func (o *SwapInput) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *SwapInput) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *SwapInput) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.

### HasPriorityFee

`func (o *SwapInput) HasPriorityFee() bool`

HasPriorityFee returns a boolean if a field has been set.

### GetPoolAddress

`func (o *SwapInput) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *SwapInput) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *SwapInput) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.

### HasPoolAddress

`func (o *SwapInput) HasPoolAddress() bool`

HasPoolAddress returns a boolean if a field has been set.

### GetAmount

`func (o *SwapInput) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SwapInput) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SwapInput) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSwapMode

`func (o *SwapInput) GetSwapMode() string`

GetSwapMode returns the SwapMode field if non-nil, zero value otherwise.

### GetSwapModeOk

`func (o *SwapInput) GetSwapModeOk() (*string, bool)`

GetSwapModeOk returns a tuple with the SwapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapMode

`func (o *SwapInput) SetSwapMode(v string)`

SetSwapMode sets SwapMode field to given value.


### GetSlippage

`func (o *SwapInput) GetSlippage() int64`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *SwapInput) GetSlippageOk() (*int64, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *SwapInput) SetSlippage(v int64)`

SetSlippage sets Slippage field to given value.


### GetInputMint

`func (o *SwapInput) GetInputMint() string`

GetInputMint returns the InputMint field if non-nil, zero value otherwise.

### GetInputMintOk

`func (o *SwapInput) GetInputMintOk() (*string, bool)`

GetInputMintOk returns a tuple with the InputMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMint

`func (o *SwapInput) SetInputMint(v string)`

SetInputMint sets InputMint field to given value.

### HasInputMint

`func (o *SwapInput) HasInputMint() bool`

HasInputMint returns a boolean if a field has been set.

### GetOutputMint

`func (o *SwapInput) GetOutputMint() string`

GetOutputMint returns the OutputMint field if non-nil, zero value otherwise.

### GetOutputMintOk

`func (o *SwapInput) GetOutputMintOk() (*string, bool)`

GetOutputMintOk returns a tuple with the OutputMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMint

`func (o *SwapInput) SetOutputMint(v string)`

SetOutputMint sets OutputMint field to given value.

### HasOutputMint

`func (o *SwapInput) HasOutputMint() bool`

HasOutputMint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


