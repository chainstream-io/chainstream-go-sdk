# SwapRouteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** | DTO.DEX.IDENTIFIER | 
**UserAddress** | **string** | DTO.DEX.WALLET | 
**PriorityFee** | Pointer to **string** | DTO.DEX.BASE.PRIORITY_FEE | [optional] 
**Amount** | **string** | DTO.DEX.SWAP.AMOUNT | 
**SwapMode** | **string** | DTO.DEX.SWAP.MODE | 
**Slippage** | **int64** | DTO.DEX.SWAP.SLIPPAGE | [default to 5]
**InputMint** | Pointer to **string** | DTO.DEX.SWAP.INPUT_MINT | [optional] 
**OutputMint** | Pointer to **string** | DTO.DEX.SWAP.OUTPUT_MINT | [optional] 
**RecipientAddress** | Pointer to **string** | DTO.DEX.SWAP.RECIPIENT_ADDRESS | [optional] 
**Permit** | Pointer to **string** | DTO.DEX.SWAP.PERMIT | [optional] 
**Deadline** | Pointer to **int64** | DTO.DEX.SWAP.DEADLINE | [optional] 
**TipFee** | Pointer to **string** | DTO.DEX.BASE.TIP_FEE | [optional] 
**IsAntiMev** | Pointer to **bool** | DTO.DEX.BASE.IS_ANTI_MEV | [optional] [default to false]

## Methods

### NewSwapRouteInput

`func NewSwapRouteInput(dex string, userAddress string, amount string, swapMode string, slippage int64, ) *SwapRouteInput`

NewSwapRouteInput instantiates a new SwapRouteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapRouteInputWithDefaults

`func NewSwapRouteInputWithDefaults() *SwapRouteInput`

NewSwapRouteInputWithDefaults instantiates a new SwapRouteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *SwapRouteInput) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *SwapRouteInput) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *SwapRouteInput) SetDex(v string)`

SetDex sets Dex field to given value.


### GetUserAddress

`func (o *SwapRouteInput) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *SwapRouteInput) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *SwapRouteInput) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetPriorityFee

`func (o *SwapRouteInput) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *SwapRouteInput) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *SwapRouteInput) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.

### HasPriorityFee

`func (o *SwapRouteInput) HasPriorityFee() bool`

HasPriorityFee returns a boolean if a field has been set.

### GetAmount

`func (o *SwapRouteInput) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SwapRouteInput) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SwapRouteInput) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSwapMode

`func (o *SwapRouteInput) GetSwapMode() string`

GetSwapMode returns the SwapMode field if non-nil, zero value otherwise.

### GetSwapModeOk

`func (o *SwapRouteInput) GetSwapModeOk() (*string, bool)`

GetSwapModeOk returns a tuple with the SwapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapMode

`func (o *SwapRouteInput) SetSwapMode(v string)`

SetSwapMode sets SwapMode field to given value.


### GetSlippage

`func (o *SwapRouteInput) GetSlippage() int64`

GetSlippage returns the Slippage field if non-nil, zero value otherwise.

### GetSlippageOk

`func (o *SwapRouteInput) GetSlippageOk() (*int64, bool)`

GetSlippageOk returns a tuple with the Slippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippage

`func (o *SwapRouteInput) SetSlippage(v int64)`

SetSlippage sets Slippage field to given value.


### GetInputMint

`func (o *SwapRouteInput) GetInputMint() string`

GetInputMint returns the InputMint field if non-nil, zero value otherwise.

### GetInputMintOk

`func (o *SwapRouteInput) GetInputMintOk() (*string, bool)`

GetInputMintOk returns a tuple with the InputMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMint

`func (o *SwapRouteInput) SetInputMint(v string)`

SetInputMint sets InputMint field to given value.

### HasInputMint

`func (o *SwapRouteInput) HasInputMint() bool`

HasInputMint returns a boolean if a field has been set.

### GetOutputMint

`func (o *SwapRouteInput) GetOutputMint() string`

GetOutputMint returns the OutputMint field if non-nil, zero value otherwise.

### GetOutputMintOk

`func (o *SwapRouteInput) GetOutputMintOk() (*string, bool)`

GetOutputMintOk returns a tuple with the OutputMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMint

`func (o *SwapRouteInput) SetOutputMint(v string)`

SetOutputMint sets OutputMint field to given value.

### HasOutputMint

`func (o *SwapRouteInput) HasOutputMint() bool`

HasOutputMint returns a boolean if a field has been set.

### GetRecipientAddress

`func (o *SwapRouteInput) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *SwapRouteInput) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *SwapRouteInput) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.

### HasRecipientAddress

`func (o *SwapRouteInput) HasRecipientAddress() bool`

HasRecipientAddress returns a boolean if a field has been set.

### GetPermit

`func (o *SwapRouteInput) GetPermit() string`

GetPermit returns the Permit field if non-nil, zero value otherwise.

### GetPermitOk

`func (o *SwapRouteInput) GetPermitOk() (*string, bool)`

GetPermitOk returns a tuple with the Permit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermit

`func (o *SwapRouteInput) SetPermit(v string)`

SetPermit sets Permit field to given value.

### HasPermit

`func (o *SwapRouteInput) HasPermit() bool`

HasPermit returns a boolean if a field has been set.

### GetDeadline

`func (o *SwapRouteInput) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *SwapRouteInput) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *SwapRouteInput) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *SwapRouteInput) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetTipFee

`func (o *SwapRouteInput) GetTipFee() string`

GetTipFee returns the TipFee field if non-nil, zero value otherwise.

### GetTipFeeOk

`func (o *SwapRouteInput) GetTipFeeOk() (*string, bool)`

GetTipFeeOk returns a tuple with the TipFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipFee

`func (o *SwapRouteInput) SetTipFee(v string)`

SetTipFee sets TipFee field to given value.

### HasTipFee

`func (o *SwapRouteInput) HasTipFee() bool`

HasTipFee returns a boolean if a field has been set.

### GetIsAntiMev

`func (o *SwapRouteInput) GetIsAntiMev() bool`

GetIsAntiMev returns the IsAntiMev field if non-nil, zero value otherwise.

### GetIsAntiMevOk

`func (o *SwapRouteInput) GetIsAntiMevOk() (*bool, bool)`

GetIsAntiMevOk returns a tuple with the IsAntiMev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAntiMev

`func (o *SwapRouteInput) SetIsAntiMev(v bool)`

SetIsAntiMev sets IsAntiMev field to given value.

### HasIsAntiMev

`func (o *SwapRouteInput) HasIsAntiMev() bool`

HasIsAntiMev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


