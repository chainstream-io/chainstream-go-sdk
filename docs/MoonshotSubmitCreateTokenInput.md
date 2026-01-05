# MoonshotSubmitCreateTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | **string** | DTO.DEX.MOONSHOT.SUBMIT.CREATE_TOKEN.SIGNED_TX | 
**Extra** | Pointer to **map[string]interface{}** | DTO.DEX.MOONSHOT.SUBMIT.CREATE_TOKEN.EXTRA | [optional] 

## Methods

### NewMoonshotSubmitCreateTokenInput

`func NewMoonshotSubmitCreateTokenInput(signedTx string, ) *MoonshotSubmitCreateTokenInput`

NewMoonshotSubmitCreateTokenInput instantiates a new MoonshotSubmitCreateTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoonshotSubmitCreateTokenInputWithDefaults

`func NewMoonshotSubmitCreateTokenInputWithDefaults() *MoonshotSubmitCreateTokenInput`

NewMoonshotSubmitCreateTokenInputWithDefaults instantiates a new MoonshotSubmitCreateTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *MoonshotSubmitCreateTokenInput) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *MoonshotSubmitCreateTokenInput) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *MoonshotSubmitCreateTokenInput) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.


### GetExtra

`func (o *MoonshotSubmitCreateTokenInput) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *MoonshotSubmitCreateTokenInput) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *MoonshotSubmitCreateTokenInput) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *MoonshotSubmitCreateTokenInput) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


