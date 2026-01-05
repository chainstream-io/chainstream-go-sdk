# CreateRedPacketReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxSerialize** | **string** | DTO.RED_PACKET.SERIALIZED_TX | 
**ShareId** | **string** | DTO.RED_PACKET.SHARE_ID | 

## Methods

### NewCreateRedPacketReply

`func NewCreateRedPacketReply(txSerialize string, shareId string, ) *CreateRedPacketReply`

NewCreateRedPacketReply instantiates a new CreateRedPacketReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRedPacketReplyWithDefaults

`func NewCreateRedPacketReplyWithDefaults() *CreateRedPacketReply`

NewCreateRedPacketReplyWithDefaults instantiates a new CreateRedPacketReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxSerialize

`func (o *CreateRedPacketReply) GetTxSerialize() string`

GetTxSerialize returns the TxSerialize field if non-nil, zero value otherwise.

### GetTxSerializeOk

`func (o *CreateRedPacketReply) GetTxSerializeOk() (*string, bool)`

GetTxSerializeOk returns a tuple with the TxSerialize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSerialize

`func (o *CreateRedPacketReply) SetTxSerialize(v string)`

SetTxSerialize sets TxSerialize field to given value.


### GetShareId

`func (o *CreateRedPacketReply) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *CreateRedPacketReply) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *CreateRedPacketReply) SetShareId(v string)`

SetShareId sets ShareId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


