# SwapReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerializedTx** | **string** | DTO.DEX.SWAP.SERIALIZED_TX | 
**ElapsedTime** | **float32** | DTO.DEX.SWAP.ELAPSED_TIME | 

## Methods

### NewSwapReply

`func NewSwapReply(serializedTx string, elapsedTime float32, ) *SwapReply`

NewSwapReply instantiates a new SwapReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapReplyWithDefaults

`func NewSwapReplyWithDefaults() *SwapReply`

NewSwapReplyWithDefaults instantiates a new SwapReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerializedTx

`func (o *SwapReply) GetSerializedTx() string`

GetSerializedTx returns the SerializedTx field if non-nil, zero value otherwise.

### GetSerializedTxOk

`func (o *SwapReply) GetSerializedTxOk() (*string, bool)`

GetSerializedTxOk returns a tuple with the SerializedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedTx

`func (o *SwapReply) SetSerializedTx(v string)`

SetSerializedTx sets SerializedTx field to given value.


### GetElapsedTime

`func (o *SwapReply) GetElapsedTime() float32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SwapReply) GetElapsedTimeOk() (*float32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SwapReply) SetElapsedTime(v float32)`

SetElapsedTime sets ElapsedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


