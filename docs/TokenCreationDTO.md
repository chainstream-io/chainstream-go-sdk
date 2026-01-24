# TokenCreationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | DTO.TOKEN.CREATION.TOKEN_ADDRESS | 
**BlockHeight** | **int64** | DTO.TOKEN.CREATION.BLOCK_HEIGHT | 
**BlockSlot** | **int64** | DTO.TOKEN.CREATION.BLOCK_SLOT | 
**BlockHash** | **string** | DTO.TOKEN.CREATION.BLOCK_HASH | 
**BlockTimestamp** | Pointer to **int64** | DTO.TOKEN.CREATION.BLOCK_TIMESTAMP | [optional] 
**TransactionSignature** | **string** | DTO.TOKEN.CREATION.TRANSACTION_SIGNATURE | 
**Type** | **string** | DTO.TOKEN.CREATION.TYPE | 

## Methods

### NewTokenCreationDTO

`func NewTokenCreationDTO(tokenAddress string, blockHeight int64, blockSlot int64, blockHash string, transactionSignature string, type_ string, ) *TokenCreationDTO`

NewTokenCreationDTO instantiates a new TokenCreationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreationDTOWithDefaults

`func NewTokenCreationDTOWithDefaults() *TokenCreationDTO`

NewTokenCreationDTOWithDefaults instantiates a new TokenCreationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *TokenCreationDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenCreationDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenCreationDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetBlockHeight

`func (o *TokenCreationDTO) GetBlockHeight() int64`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TokenCreationDTO) GetBlockHeightOk() (*int64, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TokenCreationDTO) SetBlockHeight(v int64)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockSlot

`func (o *TokenCreationDTO) GetBlockSlot() int64`

GetBlockSlot returns the BlockSlot field if non-nil, zero value otherwise.

### GetBlockSlotOk

`func (o *TokenCreationDTO) GetBlockSlotOk() (*int64, bool)`

GetBlockSlotOk returns a tuple with the BlockSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSlot

`func (o *TokenCreationDTO) SetBlockSlot(v int64)`

SetBlockSlot sets BlockSlot field to given value.


### GetBlockHash

`func (o *TokenCreationDTO) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TokenCreationDTO) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TokenCreationDTO) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockTimestamp

`func (o *TokenCreationDTO) GetBlockTimestamp() int64`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *TokenCreationDTO) GetBlockTimestampOk() (*int64, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *TokenCreationDTO) SetBlockTimestamp(v int64)`

SetBlockTimestamp sets BlockTimestamp field to given value.

### HasBlockTimestamp

`func (o *TokenCreationDTO) HasBlockTimestamp() bool`

HasBlockTimestamp returns a boolean if a field has been set.

### GetTransactionSignature

`func (o *TokenCreationDTO) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *TokenCreationDTO) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *TokenCreationDTO) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.


### GetType

`func (o *TokenCreationDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenCreationDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenCreationDTO) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


