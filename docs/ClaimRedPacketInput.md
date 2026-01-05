# ClaimRedPacketInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | DTO.RED_PACKET.CHAIN | 
**PacketId** | Pointer to **string** | DTO.RED_PACKET.PACKET_ID | [optional] 
**ShareId** | Pointer to **string** | DTO.RED_PACKET.SHARE_ID | [optional] 
**Password** | Pointer to **string** | DTO.RED_PACKET.PASSWORD | [optional] 
**Claimer** | **string** | DTO.RED_PACKET.CLAIMER | 

## Methods

### NewClaimRedPacketInput

`func NewClaimRedPacketInput(chain string, claimer string, ) *ClaimRedPacketInput`

NewClaimRedPacketInput instantiates a new ClaimRedPacketInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimRedPacketInputWithDefaults

`func NewClaimRedPacketInputWithDefaults() *ClaimRedPacketInput`

NewClaimRedPacketInputWithDefaults instantiates a new ClaimRedPacketInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ClaimRedPacketInput) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ClaimRedPacketInput) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ClaimRedPacketInput) SetChain(v string)`

SetChain sets Chain field to given value.


### GetPacketId

`func (o *ClaimRedPacketInput) GetPacketId() string`

GetPacketId returns the PacketId field if non-nil, zero value otherwise.

### GetPacketIdOk

`func (o *ClaimRedPacketInput) GetPacketIdOk() (*string, bool)`

GetPacketIdOk returns a tuple with the PacketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketId

`func (o *ClaimRedPacketInput) SetPacketId(v string)`

SetPacketId sets PacketId field to given value.

### HasPacketId

`func (o *ClaimRedPacketInput) HasPacketId() bool`

HasPacketId returns a boolean if a field has been set.

### GetShareId

`func (o *ClaimRedPacketInput) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *ClaimRedPacketInput) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *ClaimRedPacketInput) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *ClaimRedPacketInput) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### GetPassword

`func (o *ClaimRedPacketInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClaimRedPacketInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClaimRedPacketInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClaimRedPacketInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetClaimer

`func (o *ClaimRedPacketInput) GetClaimer() string`

GetClaimer returns the Claimer field if non-nil, zero value otherwise.

### GetClaimerOk

`func (o *ClaimRedPacketInput) GetClaimerOk() (*string, bool)`

GetClaimerOk returns a tuple with the Claimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimer

`func (o *ClaimRedPacketInput) SetClaimer(v string)`

SetClaimer sets Claimer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


