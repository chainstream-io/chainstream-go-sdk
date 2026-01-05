# CreateTokenReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerializedTx** | **string** | DTO.DEX.CREATE_TOKEN.SERIALIZED_TX | 
**MintAddress** | **string** | DTO.DEX.CREATE_TOKEN.MINT_ADDRESS | 

## Methods

### NewCreateTokenReply

`func NewCreateTokenReply(serializedTx string, mintAddress string, ) *CreateTokenReply`

NewCreateTokenReply instantiates a new CreateTokenReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenReplyWithDefaults

`func NewCreateTokenReplyWithDefaults() *CreateTokenReply`

NewCreateTokenReplyWithDefaults instantiates a new CreateTokenReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerializedTx

`func (o *CreateTokenReply) GetSerializedTx() string`

GetSerializedTx returns the SerializedTx field if non-nil, zero value otherwise.

### GetSerializedTxOk

`func (o *CreateTokenReply) GetSerializedTxOk() (*string, bool)`

GetSerializedTxOk returns a tuple with the SerializedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedTx

`func (o *CreateTokenReply) SetSerializedTx(v string)`

SetSerializedTx sets SerializedTx field to given value.


### GetMintAddress

`func (o *CreateTokenReply) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *CreateTokenReply) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *CreateTokenReply) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


