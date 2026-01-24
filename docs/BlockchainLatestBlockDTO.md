# BlockchainLatestBlockDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockhash** | **string** | DTO.BLOCKCHAIN.LATEST_BLOCK.BLOCKHASH | 
**LastValidBlockHeight** | **int64** | DTO.BLOCKCHAIN.LATEST_BLOCK.LAST_VALID_BLOCK_HEIGHT | 

## Methods

### NewBlockchainLatestBlockDTO

`func NewBlockchainLatestBlockDTO(blockhash string, lastValidBlockHeight int64, ) *BlockchainLatestBlockDTO`

NewBlockchainLatestBlockDTO instantiates a new BlockchainLatestBlockDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainLatestBlockDTOWithDefaults

`func NewBlockchainLatestBlockDTOWithDefaults() *BlockchainLatestBlockDTO`

NewBlockchainLatestBlockDTOWithDefaults instantiates a new BlockchainLatestBlockDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockhash

`func (o *BlockchainLatestBlockDTO) GetBlockhash() string`

GetBlockhash returns the Blockhash field if non-nil, zero value otherwise.

### GetBlockhashOk

`func (o *BlockchainLatestBlockDTO) GetBlockhashOk() (*string, bool)`

GetBlockhashOk returns a tuple with the Blockhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockhash

`func (o *BlockchainLatestBlockDTO) SetBlockhash(v string)`

SetBlockhash sets Blockhash field to given value.


### GetLastValidBlockHeight

`func (o *BlockchainLatestBlockDTO) GetLastValidBlockHeight() int64`

GetLastValidBlockHeight returns the LastValidBlockHeight field if non-nil, zero value otherwise.

### GetLastValidBlockHeightOk

`func (o *BlockchainLatestBlockDTO) GetLastValidBlockHeightOk() (*int64, bool)`

GetLastValidBlockHeightOk returns a tuple with the LastValidBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidBlockHeight

`func (o *BlockchainLatestBlockDTO) SetLastValidBlockHeight(v int64)`

SetLastValidBlockHeight sets LastValidBlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


