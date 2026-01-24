# RedPacketClaimDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PacketId** | **string** | DTO.RED_PACKET.PACKET_ID | 
**Chain** | [**Chain**](Chain.md) | DTO.RED_PACKET.CHAIN | 
**Claimer** | **string** | DTO.RED_PACKET.CLAIMER | 
**Mint** | **string** | DTO.RED_PACKET.MINT | 
**Amount** | **string** | DTO.RED_PACKET.AMOUNT | 
**ClaimedAt** | **int64** | DTO.RED_PACKET.CLAIMED_AT | 
**Creator** | **string** | DTO.RED_PACKET.CREATOR | 
**TxHash** | **string** | DTO.RED_PACKET.TX_HASH | 

## Methods

### NewRedPacketClaimDTO

`func NewRedPacketClaimDTO(packetId string, chain Chain, claimer string, mint string, amount string, claimedAt int64, creator string, txHash string, ) *RedPacketClaimDTO`

NewRedPacketClaimDTO instantiates a new RedPacketClaimDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedPacketClaimDTOWithDefaults

`func NewRedPacketClaimDTOWithDefaults() *RedPacketClaimDTO`

NewRedPacketClaimDTOWithDefaults instantiates a new RedPacketClaimDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPacketId

`func (o *RedPacketClaimDTO) GetPacketId() string`

GetPacketId returns the PacketId field if non-nil, zero value otherwise.

### GetPacketIdOk

`func (o *RedPacketClaimDTO) GetPacketIdOk() (*string, bool)`

GetPacketIdOk returns a tuple with the PacketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketId

`func (o *RedPacketClaimDTO) SetPacketId(v string)`

SetPacketId sets PacketId field to given value.


### GetChain

`func (o *RedPacketClaimDTO) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *RedPacketClaimDTO) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *RedPacketClaimDTO) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetClaimer

`func (o *RedPacketClaimDTO) GetClaimer() string`

GetClaimer returns the Claimer field if non-nil, zero value otherwise.

### GetClaimerOk

`func (o *RedPacketClaimDTO) GetClaimerOk() (*string, bool)`

GetClaimerOk returns a tuple with the Claimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimer

`func (o *RedPacketClaimDTO) SetClaimer(v string)`

SetClaimer sets Claimer field to given value.


### GetMint

`func (o *RedPacketClaimDTO) GetMint() string`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *RedPacketClaimDTO) GetMintOk() (*string, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *RedPacketClaimDTO) SetMint(v string)`

SetMint sets Mint field to given value.


### GetAmount

`func (o *RedPacketClaimDTO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RedPacketClaimDTO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RedPacketClaimDTO) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClaimedAt

`func (o *RedPacketClaimDTO) GetClaimedAt() int64`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *RedPacketClaimDTO) GetClaimedAtOk() (*int64, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *RedPacketClaimDTO) SetClaimedAt(v int64)`

SetClaimedAt sets ClaimedAt field to given value.


### GetCreator

`func (o *RedPacketClaimDTO) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RedPacketClaimDTO) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RedPacketClaimDTO) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetTxHash

`func (o *RedPacketClaimDTO) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *RedPacketClaimDTO) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *RedPacketClaimDTO) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


