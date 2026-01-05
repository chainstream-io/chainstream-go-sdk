# TokenHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddress** | **string** | DTO.TOKEN_HOLDER.WALLET_ADDRESS | 
**Amount** | **float32** | DTO.TOKEN_HOLDER.AMOUNT | 
**AmountInUsd** | **float32** | DTO.TOKEN_HOLDER.AMOUNT_USD | 
**Percentage** | **string** | DTO.TOKEN_HOLDER.PERCENTAGE | 

## Methods

### NewTokenHolder

`func NewTokenHolder(walletAddress string, amount float32, amountInUsd float32, percentage string, ) *TokenHolder`

NewTokenHolder instantiates a new TokenHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHolderWithDefaults

`func NewTokenHolderWithDefaults() *TokenHolder`

NewTokenHolderWithDefaults instantiates a new TokenHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddress

`func (o *TokenHolder) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *TokenHolder) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *TokenHolder) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetAmount

`func (o *TokenHolder) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenHolder) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenHolder) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAmountInUsd

`func (o *TokenHolder) GetAmountInUsd() float32`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *TokenHolder) GetAmountInUsdOk() (*float32, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *TokenHolder) SetAmountInUsd(v float32)`

SetAmountInUsd sets AmountInUsd field to given value.


### GetPercentage

`func (o *TokenHolder) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *TokenHolder) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *TokenHolder) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


