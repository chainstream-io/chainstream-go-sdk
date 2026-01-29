# BlockchainDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Blockchain symbol | 
**Name** | **string** | Blockchain name | 
**ExplorerUrl** | **string** | Blockchain explorer URL | 
**ChainId** | **int64** | Blockchain chain ID | 

## Methods

### NewBlockchainDTO

`func NewBlockchainDTO(symbol string, name string, explorerUrl string, chainId int64, ) *BlockchainDTO`

NewBlockchainDTO instantiates a new BlockchainDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainDTOWithDefaults

`func NewBlockchainDTOWithDefaults() *BlockchainDTO`

NewBlockchainDTOWithDefaults instantiates a new BlockchainDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *BlockchainDTO) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BlockchainDTO) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BlockchainDTO) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *BlockchainDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlockchainDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlockchainDTO) SetName(v string)`

SetName sets Name field to given value.


### GetExplorerUrl

`func (o *BlockchainDTO) GetExplorerUrl() string`

GetExplorerUrl returns the ExplorerUrl field if non-nil, zero value otherwise.

### GetExplorerUrlOk

`func (o *BlockchainDTO) GetExplorerUrlOk() (*string, bool)`

GetExplorerUrlOk returns a tuple with the ExplorerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerUrl

`func (o *BlockchainDTO) SetExplorerUrl(v string)`

SetExplorerUrl sets ExplorerUrl field to given value.


### GetChainId

`func (o *BlockchainDTO) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *BlockchainDTO) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *BlockchainDTO) SetChainId(v int64)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


