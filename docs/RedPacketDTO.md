# RedPacketDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | [**Chain**](Chain.md) | GLOBAL.CHAIN.DESCRIPTION | 
**Id** | **string** | DTO.RED_PACKET.ID | 
**ShareId** | **string** | DTO.RED_PACKET.SHARE_ID | 
**TxHash** | **string** | DTO.RED_PACKET.TX_HASH | 
**Creator** | **string** | DTO.RED_PACKET.CREATOR | 
**Mint** | **string** | DTO.RED_PACKET.MINT | 
**TotalAmount** | **string** | DTO.RED_PACKET.TOTAL_AMOUNT | 
**Memo** | **string** | DTO.RED_PACKET.MEMO | 
**MaxClaims** | **float32** | DTO.RED_PACKET.MAX_CLAIMS | 
**ClaimAuthority** | **string** | DTO.RED_PACKET.CLAIM_AUTHORITY | 
**Expired** | **bool** | DTO.RED_PACKET.WITHDRAWED | 
**Expiration** | **float32** | DTO.RED_PACKET.EXPIRATION | 
**CreatedAt** | **float32** | DTO.RED_PACKET.CREATED_AT | 
**ExpiredAt** | **float32** | DTO.RED_PACKET.EXPIRES_AT | 
**ClaimedCount** | **float32** | DTO.RED_PACKET.CLAIMED_COUNT | 
**ClaimedAmount** | **string** | DTO.RED_PACKET.CLAIMED_AMOUNT | 
**RefundedAmount** | **string** | DTO.RED_PACKET.REFUND_AMOUNT | 

## Methods

### NewRedPacketDTO

`func NewRedPacketDTO(chain Chain, id string, shareId string, txHash string, creator string, mint string, totalAmount string, memo string, maxClaims float32, claimAuthority string, expired bool, expiration float32, createdAt float32, expiredAt float32, claimedCount float32, claimedAmount string, refundedAmount string, ) *RedPacketDTO`

NewRedPacketDTO instantiates a new RedPacketDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedPacketDTOWithDefaults

`func NewRedPacketDTOWithDefaults() *RedPacketDTO`

NewRedPacketDTOWithDefaults instantiates a new RedPacketDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *RedPacketDTO) GetChain() Chain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *RedPacketDTO) GetChainOk() (*Chain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *RedPacketDTO) SetChain(v Chain)`

SetChain sets Chain field to given value.


### GetId

`func (o *RedPacketDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedPacketDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedPacketDTO) SetId(v string)`

SetId sets Id field to given value.


### GetShareId

`func (o *RedPacketDTO) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *RedPacketDTO) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *RedPacketDTO) SetShareId(v string)`

SetShareId sets ShareId field to given value.


### GetTxHash

`func (o *RedPacketDTO) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *RedPacketDTO) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *RedPacketDTO) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetCreator

`func (o *RedPacketDTO) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RedPacketDTO) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RedPacketDTO) SetCreator(v string)`

SetCreator sets Creator field to given value.


### GetMint

`func (o *RedPacketDTO) GetMint() string`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *RedPacketDTO) GetMintOk() (*string, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *RedPacketDTO) SetMint(v string)`

SetMint sets Mint field to given value.


### GetTotalAmount

`func (o *RedPacketDTO) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *RedPacketDTO) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *RedPacketDTO) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetMemo

`func (o *RedPacketDTO) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *RedPacketDTO) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *RedPacketDTO) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetMaxClaims

`func (o *RedPacketDTO) GetMaxClaims() float32`

GetMaxClaims returns the MaxClaims field if non-nil, zero value otherwise.

### GetMaxClaimsOk

`func (o *RedPacketDTO) GetMaxClaimsOk() (*float32, bool)`

GetMaxClaimsOk returns a tuple with the MaxClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClaims

`func (o *RedPacketDTO) SetMaxClaims(v float32)`

SetMaxClaims sets MaxClaims field to given value.


### GetClaimAuthority

`func (o *RedPacketDTO) GetClaimAuthority() string`

GetClaimAuthority returns the ClaimAuthority field if non-nil, zero value otherwise.

### GetClaimAuthorityOk

`func (o *RedPacketDTO) GetClaimAuthorityOk() (*string, bool)`

GetClaimAuthorityOk returns a tuple with the ClaimAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimAuthority

`func (o *RedPacketDTO) SetClaimAuthority(v string)`

SetClaimAuthority sets ClaimAuthority field to given value.


### GetExpired

`func (o *RedPacketDTO) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *RedPacketDTO) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *RedPacketDTO) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetExpiration

`func (o *RedPacketDTO) GetExpiration() float32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *RedPacketDTO) GetExpirationOk() (*float32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *RedPacketDTO) SetExpiration(v float32)`

SetExpiration sets Expiration field to given value.


### GetCreatedAt

`func (o *RedPacketDTO) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RedPacketDTO) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RedPacketDTO) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiredAt

`func (o *RedPacketDTO) GetExpiredAt() float32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *RedPacketDTO) GetExpiredAtOk() (*float32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *RedPacketDTO) SetExpiredAt(v float32)`

SetExpiredAt sets ExpiredAt field to given value.


### GetClaimedCount

`func (o *RedPacketDTO) GetClaimedCount() float32`

GetClaimedCount returns the ClaimedCount field if non-nil, zero value otherwise.

### GetClaimedCountOk

`func (o *RedPacketDTO) GetClaimedCountOk() (*float32, bool)`

GetClaimedCountOk returns a tuple with the ClaimedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedCount

`func (o *RedPacketDTO) SetClaimedCount(v float32)`

SetClaimedCount sets ClaimedCount field to given value.


### GetClaimedAmount

`func (o *RedPacketDTO) GetClaimedAmount() string`

GetClaimedAmount returns the ClaimedAmount field if non-nil, zero value otherwise.

### GetClaimedAmountOk

`func (o *RedPacketDTO) GetClaimedAmountOk() (*string, bool)`

GetClaimedAmountOk returns a tuple with the ClaimedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAmount

`func (o *RedPacketDTO) SetClaimedAmount(v string)`

SetClaimedAmount sets ClaimedAmount field to given value.


### GetRefundedAmount

`func (o *RedPacketDTO) GetRefundedAmount() string`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *RedPacketDTO) GetRefundedAmountOk() (*string, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *RedPacketDTO) SetRefundedAmount(v string)`

SetRefundedAmount sets RefundedAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


