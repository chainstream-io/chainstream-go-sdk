# CreateRedPacketInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | Blockchain network | 
**Creator** | **string** | Creator wallet address of the red packet | 
**Mint** | **string** | Token mint address for the red packet | 
**MaxClaims** | **int64** | Maximum number of recipients | 
**TotalAmount** | Pointer to **string** | Total amount of tokens to be distributed | [optional] 
**FixedAmount** | Pointer to **string** | Fixed amount per claim (for equal distribution) | [optional] 
**Memo** | Pointer to **string** | Optional memo message for the red packet | [optional] 
**Password** | Pointer to **string** | Optional password to claim the red packet | [optional] 
**ClaimAuthority** | Pointer to **string** | Optional authority address that can claim the red packet | [optional] 

## Methods

### NewCreateRedPacketInput

`func NewCreateRedPacketInput(chain string, creator string, mint string, maxClaims int64, ) *CreateRedPacketInput`

NewCreateRedPacketInput instantiates a new CreateRedPacketInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRedPacketInputWithDefaults

`func NewCreateRedPacketInputWithDefaults() *CreateRedPacketInput`

NewCreateRedPacketInputWithDefaults instantiates a new CreateRedPacketInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *CreateRedPacketInput) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateRedPacketInput) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateRedPacketInput) SetChain(v string)`

SetChain sets Chain field to given value.


### GetCreator

`func (o *CreateRedPacketInput) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateRedPacketInput) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateRedPacketInput) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetMint

`func (o *CreateRedPacketInput) GetMint() string`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *CreateRedPacketInput) GetMintOk() (*string, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *CreateRedPacketInput) SetMint(v string)`

SetMint sets Mint field to given value.


### GetMaxClaims

`func (o *CreateRedPacketInput) GetMaxClaims() int64`

GetMaxClaims returns the MaxClaims field if non-nil, zero value otherwise.

### GetMaxClaimsOk

`func (o *CreateRedPacketInput) GetMaxClaimsOk() (*int64, bool)`

GetMaxClaimsOk returns a tuple with the MaxClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClaims

`func (o *CreateRedPacketInput) SetMaxClaims(v int64)`

SetMaxClaims sets MaxClaims field to given value.


### GetTotalAmount

`func (o *CreateRedPacketInput) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CreateRedPacketInput) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CreateRedPacketInput) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CreateRedPacketInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetFixedAmount

`func (o *CreateRedPacketInput) GetFixedAmount() string`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *CreateRedPacketInput) GetFixedAmountOk() (*string, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *CreateRedPacketInput) SetFixedAmount(v string)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *CreateRedPacketInput) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetMemo

`func (o *CreateRedPacketInput) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreateRedPacketInput) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreateRedPacketInput) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreateRedPacketInput) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPassword

`func (o *CreateRedPacketInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateRedPacketInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateRedPacketInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateRedPacketInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetClaimAuthority

`func (o *CreateRedPacketInput) GetClaimAuthority() string`

GetClaimAuthority returns the ClaimAuthority field if non-nil, zero value otherwise.

### GetClaimAuthorityOk

`func (o *CreateRedPacketInput) GetClaimAuthorityOk() (*string, bool)`

GetClaimAuthorityOk returns a tuple with the ClaimAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimAuthority

`func (o *CreateRedPacketInput) SetClaimAuthority(v string)`

SetClaimAuthority sets ClaimAuthority field to given value.

### HasClaimAuthority

`func (o *CreateRedPacketInput) HasClaimAuthority() bool`

HasClaimAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


