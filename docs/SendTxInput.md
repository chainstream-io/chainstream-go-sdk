# SendTxInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | **string** | Base64 encoded signed transaction | 
**SubmitType** | Pointer to **string** | Transaction submission method | [optional] 
**Options** | Pointer to **map[string]interface{}** | jito | direct | [optional] 

## Methods

### NewSendTxInput

`func NewSendTxInput(signedTx string, ) *SendTxInput`

NewSendTxInput instantiates a new SendTxInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTxInputWithDefaults

`func NewSendTxInputWithDefaults() *SendTxInput`

NewSendTxInputWithDefaults instantiates a new SendTxInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *SendTxInput) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *SendTxInput) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *SendTxInput) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.


### GetSubmitType

`func (o *SendTxInput) GetSubmitType() string`

GetSubmitType returns the SubmitType field if non-nil, zero value otherwise.

### GetSubmitTypeOk

`func (o *SendTxInput) GetSubmitTypeOk() (*string, bool)`

GetSubmitTypeOk returns a tuple with the SubmitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitType

`func (o *SendTxInput) SetSubmitType(v string)`

SetSubmitType sets SubmitType field to given value.

### HasSubmitType

`func (o *SendTxInput) HasSubmitType() bool`

HasSubmitType returns a boolean if a field has been set.

### GetOptions

`func (o *SendTxInput) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SendTxInput) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SendTxInput) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SendTxInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


