# DexPoolDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramAddress** | Pointer to **string** | DTO.DEX.PROGRAM_ADDRESS | [optional] 
**ProtocolFamily** | Pointer to **string** | DTO.DEX.PROTOCOL_FAMILY | [optional] 
**Image** | Pointer to **string** | DTO.DEX.IMAGE | [optional] 
**Chain** | **string** | DTO.DEX.CHAIN | 
**PoolAddress** | **string** | DTO.DEXPOOL.POOL_ADDRESS | 
**ProtocolName** | Pointer to **string** | DTO.DEXPOOL.PROTOCOL_NAME | [optional] 
**TokenAAddress** | **string** | DTO.DEXPOOL.TOKEN_A | 
**TokenBAddress** | **string** | DTO.DEXPOOL.TOKEN_B | 
**TvlInUsd** | Pointer to **string** | DTO.DEXPOOL.TVL_USD | [optional] 
**TvlInSol** | Pointer to **string** | DTO.DEXPOOL.TVL_SOL | [optional] 

## Methods

### NewDexPoolDTO

`func NewDexPoolDTO(chain string, poolAddress string, tokenAAddress string, tokenBAddress string, ) *DexPoolDTO`

NewDexPoolDTO instantiates a new DexPoolDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexPoolDTOWithDefaults

`func NewDexPoolDTOWithDefaults() *DexPoolDTO`

NewDexPoolDTOWithDefaults instantiates a new DexPoolDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramAddress

`func (o *DexPoolDTO) GetProgramAddress() string`

GetProgramAddress returns the ProgramAddress field if non-nil, zero value otherwise.

### GetProgramAddressOk

`func (o *DexPoolDTO) GetProgramAddressOk() (*string, bool)`

GetProgramAddressOk returns a tuple with the ProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramAddress

`func (o *DexPoolDTO) SetProgramAddress(v string)`

SetProgramAddress sets ProgramAddress field to given value.

### HasProgramAddress

`func (o *DexPoolDTO) HasProgramAddress() bool`

HasProgramAddress returns a boolean if a field has been set.

### GetProtocolFamily

`func (o *DexPoolDTO) GetProtocolFamily() string`

GetProtocolFamily returns the ProtocolFamily field if non-nil, zero value otherwise.

### GetProtocolFamilyOk

`func (o *DexPoolDTO) GetProtocolFamilyOk() (*string, bool)`

GetProtocolFamilyOk returns a tuple with the ProtocolFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFamily

`func (o *DexPoolDTO) SetProtocolFamily(v string)`

SetProtocolFamily sets ProtocolFamily field to given value.

### HasProtocolFamily

`func (o *DexPoolDTO) HasProtocolFamily() bool`

HasProtocolFamily returns a boolean if a field has been set.

### GetImage

`func (o *DexPoolDTO) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DexPoolDTO) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DexPoolDTO) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DexPoolDTO) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetChain

`func (o *DexPoolDTO) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *DexPoolDTO) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *DexPoolDTO) SetChain(v string)`

SetChain sets Chain field to given value.


### GetPoolAddress

`func (o *DexPoolDTO) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *DexPoolDTO) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *DexPoolDTO) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetProtocolName

`func (o *DexPoolDTO) GetProtocolName() string`

GetProtocolName returns the ProtocolName field if non-nil, zero value otherwise.

### GetProtocolNameOk

`func (o *DexPoolDTO) GetProtocolNameOk() (*string, bool)`

GetProtocolNameOk returns a tuple with the ProtocolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolName

`func (o *DexPoolDTO) SetProtocolName(v string)`

SetProtocolName sets ProtocolName field to given value.

### HasProtocolName

`func (o *DexPoolDTO) HasProtocolName() bool`

HasProtocolName returns a boolean if a field has been set.

### GetTokenAAddress

`func (o *DexPoolDTO) GetTokenAAddress() string`

GetTokenAAddress returns the TokenAAddress field if non-nil, zero value otherwise.

### GetTokenAAddressOk

`func (o *DexPoolDTO) GetTokenAAddressOk() (*string, bool)`

GetTokenAAddressOk returns a tuple with the TokenAAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAAddress

`func (o *DexPoolDTO) SetTokenAAddress(v string)`

SetTokenAAddress sets TokenAAddress field to given value.


### GetTokenBAddress

`func (o *DexPoolDTO) GetTokenBAddress() string`

GetTokenBAddress returns the TokenBAddress field if non-nil, zero value otherwise.

### GetTokenBAddressOk

`func (o *DexPoolDTO) GetTokenBAddressOk() (*string, bool)`

GetTokenBAddressOk returns a tuple with the TokenBAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBAddress

`func (o *DexPoolDTO) SetTokenBAddress(v string)`

SetTokenBAddress sets TokenBAddress field to given value.


### GetTvlInUsd

`func (o *DexPoolDTO) GetTvlInUsd() string`

GetTvlInUsd returns the TvlInUsd field if non-nil, zero value otherwise.

### GetTvlInUsdOk

`func (o *DexPoolDTO) GetTvlInUsdOk() (*string, bool)`

GetTvlInUsdOk returns a tuple with the TvlInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInUsd

`func (o *DexPoolDTO) SetTvlInUsd(v string)`

SetTvlInUsd sets TvlInUsd field to given value.

### HasTvlInUsd

`func (o *DexPoolDTO) HasTvlInUsd() bool`

HasTvlInUsd returns a boolean if a field has been set.

### GetTvlInSol

`func (o *DexPoolDTO) GetTvlInSol() string`

GetTvlInSol returns the TvlInSol field if non-nil, zero value otherwise.

### GetTvlInSolOk

`func (o *DexPoolDTO) GetTvlInSolOk() (*string, bool)`

GetTvlInSolOk returns a tuple with the TvlInSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInSol

`func (o *DexPoolDTO) SetTvlInSol(v string)`

SetTvlInSol sets TvlInSol field to given value.

### HasTvlInSol

`func (o *DexPoolDTO) HasTvlInSol() bool`

HasTvlInSol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


