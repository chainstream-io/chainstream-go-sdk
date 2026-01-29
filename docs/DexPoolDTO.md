# DexPoolDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramAddress** | Pointer to **string** | DEX program address | [optional] 
**ProtocolFamily** | Pointer to **string** | DEX protocol family | [optional] 
**Image** | Pointer to **string** | DEX logo image URL | [optional] 
**Chain** | **string** | Blockchain | 
**PoolAddress** | **string** | Pool address | 
**ProtocolName** | Pointer to **string** | DEX protocol name | [optional] 
**TokenAAddress** | **string** | First token in the pool | 
**TokenBAddress** | **string** | Second token in the pool | 
**TvlInUsd** | Pointer to **string** | Total value locked in USD | [optional] 
**TvlInSol** | Pointer to **string** | TVL in SOL | [optional] 
**Type** | Pointer to **float32** | Pool type (0&#x3D;unknown, 1&#x3D;standard, 2&#x3D;concentrated, 3&#x3D;weighted, 4&#x3D;stable) | [optional] 
**Version** | Pointer to **float32** | Pool version (0&#x3D;unknown, 1&#x3D;V1, 2&#x3D;V2, 3&#x3D;V3/CLMM) | [optional] 
**LiquidityModel** | Pointer to **float32** | Liquidity model (0&#x3D;unknown, 1&#x3D;reserve_based, 2&#x3D;position_based, 3&#x3D;virtual, 4&#x3D;shared) | [optional] 
**FeeRate** | Pointer to **string** | Pool fee rate (e.g., 0.003 for 0.3%) | [optional] 
**TickSpacing** | Pointer to **int32** | Tick spacing for concentrated liquidity pools | [optional] 
**TokenCount** | Pointer to **int32** | Number of tokens in pool | [optional] 
**CreatedBlockTimestamp** | Pointer to **string** | Block timestamp when pool was created | [optional] 
**TokenALiquidity** | Pointer to [**DexPoolTokenLiquidity**](DexPoolTokenLiquidity.md) | First token liquidity details | [optional] 
**TokenBLiquidity** | Pointer to [**DexPoolTokenLiquidity**](DexPoolTokenLiquidity.md) | Second token liquidity details | [optional] 

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

### GetType

`func (o *DexPoolDTO) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DexPoolDTO) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DexPoolDTO) SetType(v float32)`

SetType sets Type field to given value.

### HasType

`func (o *DexPoolDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *DexPoolDTO) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DexPoolDTO) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DexPoolDTO) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DexPoolDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLiquidityModel

`func (o *DexPoolDTO) GetLiquidityModel() float32`

GetLiquidityModel returns the LiquidityModel field if non-nil, zero value otherwise.

### GetLiquidityModelOk

`func (o *DexPoolDTO) GetLiquidityModelOk() (*float32, bool)`

GetLiquidityModelOk returns a tuple with the LiquidityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityModel

`func (o *DexPoolDTO) SetLiquidityModel(v float32)`

SetLiquidityModel sets LiquidityModel field to given value.

### HasLiquidityModel

`func (o *DexPoolDTO) HasLiquidityModel() bool`

HasLiquidityModel returns a boolean if a field has been set.

### GetFeeRate

`func (o *DexPoolDTO) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *DexPoolDTO) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *DexPoolDTO) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *DexPoolDTO) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetTickSpacing

`func (o *DexPoolDTO) GetTickSpacing() int32`

GetTickSpacing returns the TickSpacing field if non-nil, zero value otherwise.

### GetTickSpacingOk

`func (o *DexPoolDTO) GetTickSpacingOk() (*int32, bool)`

GetTickSpacingOk returns a tuple with the TickSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSpacing

`func (o *DexPoolDTO) SetTickSpacing(v int32)`

SetTickSpacing sets TickSpacing field to given value.

### HasTickSpacing

`func (o *DexPoolDTO) HasTickSpacing() bool`

HasTickSpacing returns a boolean if a field has been set.

### GetTokenCount

`func (o *DexPoolDTO) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *DexPoolDTO) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *DexPoolDTO) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *DexPoolDTO) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetCreatedBlockTimestamp

`func (o *DexPoolDTO) GetCreatedBlockTimestamp() string`

GetCreatedBlockTimestamp returns the CreatedBlockTimestamp field if non-nil, zero value otherwise.

### GetCreatedBlockTimestampOk

`func (o *DexPoolDTO) GetCreatedBlockTimestampOk() (*string, bool)`

GetCreatedBlockTimestampOk returns a tuple with the CreatedBlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBlockTimestamp

`func (o *DexPoolDTO) SetCreatedBlockTimestamp(v string)`

SetCreatedBlockTimestamp sets CreatedBlockTimestamp field to given value.

### HasCreatedBlockTimestamp

`func (o *DexPoolDTO) HasCreatedBlockTimestamp() bool`

HasCreatedBlockTimestamp returns a boolean if a field has been set.

### GetTokenALiquidity

`func (o *DexPoolDTO) GetTokenALiquidity() DexPoolTokenLiquidity`

GetTokenALiquidity returns the TokenALiquidity field if non-nil, zero value otherwise.

### GetTokenALiquidityOk

`func (o *DexPoolDTO) GetTokenALiquidityOk() (*DexPoolTokenLiquidity, bool)`

GetTokenALiquidityOk returns a tuple with the TokenALiquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenALiquidity

`func (o *DexPoolDTO) SetTokenALiquidity(v DexPoolTokenLiquidity)`

SetTokenALiquidity sets TokenALiquidity field to given value.

### HasTokenALiquidity

`func (o *DexPoolDTO) HasTokenALiquidity() bool`

HasTokenALiquidity returns a boolean if a field has been set.

### GetTokenBLiquidity

`func (o *DexPoolDTO) GetTokenBLiquidity() DexPoolTokenLiquidity`

GetTokenBLiquidity returns the TokenBLiquidity field if non-nil, zero value otherwise.

### GetTokenBLiquidityOk

`func (o *DexPoolDTO) GetTokenBLiquidityOk() (*DexPoolTokenLiquidity, bool)`

GetTokenBLiquidityOk returns a tuple with the TokenBLiquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBLiquidity

`func (o *DexPoolDTO) SetTokenBLiquidity(v DexPoolTokenLiquidity)`

SetTokenBLiquidity sets TokenBLiquidity field to given value.

### HasTokenBLiquidity

`func (o *DexPoolDTO) HasTokenBLiquidity() bool`

HasTokenBLiquidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


