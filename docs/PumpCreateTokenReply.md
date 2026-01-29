# PumpCreateTokenReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerializedTx** | **string** | Base64 encoded transaction | 
**Extra** | Pointer to **map[string]interface{}** | Additional metadata about the created token | [optional] 

## Methods

### NewPumpCreateTokenReply

`func NewPumpCreateTokenReply(serializedTx string, ) *PumpCreateTokenReply`

NewPumpCreateTokenReply instantiates a new PumpCreateTokenReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPumpCreateTokenReplyWithDefaults

`func NewPumpCreateTokenReplyWithDefaults() *PumpCreateTokenReply`

NewPumpCreateTokenReplyWithDefaults instantiates a new PumpCreateTokenReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerializedTx

`func (o *PumpCreateTokenReply) GetSerializedTx() string`

GetSerializedTx returns the SerializedTx field if non-nil, zero value otherwise.

### GetSerializedTxOk

`func (o *PumpCreateTokenReply) GetSerializedTxOk() (*string, bool)`

GetSerializedTxOk returns a tuple with the SerializedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedTx

`func (o *PumpCreateTokenReply) SetSerializedTx(v string)`

SetSerializedTx sets SerializedTx field to given value.


### GetExtra

`func (o *PumpCreateTokenReply) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *PumpCreateTokenReply) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *PumpCreateTokenReply) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *PumpCreateTokenReply) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


