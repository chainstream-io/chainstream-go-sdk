# TradeDetailDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | Blockchain network | 
**BlockHeight** | **int64** | Block height of the trade | 
**BlockTimestamp** | **int64** | Block timestamp | 
**AccountOwnerAddress** | **string** | wallet address | 
**TransactionSignature** | **string** | Transaction signature | 
**TokenName** | **string** | name | 
**TokenSymbol** | **string** | symbol | 
**TokenImageUrl** | **string** | token imageUrl | 
**TokenAddress** | **string** | Token address | 
**PoolAddress** | **string** | Pool address where the trade occurred | 
**DexProgramAddress** | **string** | DEX program address | 
**DexProtocolFamily** | **string** | DEX protocol family | 
**DexImage** | **string** | DEX logo image URL | 
**TokenAmount** | **string** | Token amount | 
**TokenPriceInUsd** | **string** | price in USD | 
**TokenAmountInUsd** | **string** | Token amount in USD | 
**Type** | **string** | Trade type | 
**SideTokenName** | **string** | name | 
**SideTokenSymbol** | **string** | symbol | 
**SideTokenImageUrl** | **string** | token imageUrl | 
**SideTokenAddress** | **string** | Side token address | 
**SideTokenPriceInUsd** | **string** | price in USD | 
**SideTokenAmount** | **string** | Side token amount | 
**SideTokenAmountInUsd** | **string** | Side token amount in USD | 
**Status** | **string** | Trade status | 

## Methods

### NewTradeDetailDTO

`func NewTradeDetailDTO(chain string, blockHeight int64, blockTimestamp int64, accountOwnerAddress string, transactionSignature string, tokenName string, tokenSymbol string, tokenImageUrl string, tokenAddress string, poolAddress string, dexProgramAddress string, dexProtocolFamily string, dexImage string, tokenAmount string, tokenPriceInUsd string, tokenAmountInUsd string, type_ string, sideTokenName string, sideTokenSymbol string, sideTokenImageUrl string, sideTokenAddress string, sideTokenPriceInUsd string, sideTokenAmount string, sideTokenAmountInUsd string, status string, ) *TradeDetailDTO`

NewTradeDetailDTO instantiates a new TradeDetailDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeDetailDTOWithDefaults

`func NewTradeDetailDTOWithDefaults() *TradeDetailDTO`

NewTradeDetailDTOWithDefaults instantiates a new TradeDetailDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *TradeDetailDTO) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *TradeDetailDTO) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *TradeDetailDTO) SetChain(v string)`

SetChain sets Chain field to given value.


### GetBlockHeight

`func (o *TradeDetailDTO) GetBlockHeight() int64`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TradeDetailDTO) GetBlockHeightOk() (*int64, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TradeDetailDTO) SetBlockHeight(v int64)`

SetBlockHeight sets BlockHeight field to given value.


### GetBlockTimestamp

`func (o *TradeDetailDTO) GetBlockTimestamp() int64`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *TradeDetailDTO) GetBlockTimestampOk() (*int64, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *TradeDetailDTO) SetBlockTimestamp(v int64)`

SetBlockTimestamp sets BlockTimestamp field to given value.


### GetAccountOwnerAddress

`func (o *TradeDetailDTO) GetAccountOwnerAddress() string`

GetAccountOwnerAddress returns the AccountOwnerAddress field if non-nil, zero value otherwise.

### GetAccountOwnerAddressOk

`func (o *TradeDetailDTO) GetAccountOwnerAddressOk() (*string, bool)`

GetAccountOwnerAddressOk returns a tuple with the AccountOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerAddress

`func (o *TradeDetailDTO) SetAccountOwnerAddress(v string)`

SetAccountOwnerAddress sets AccountOwnerAddress field to given value.


### GetTransactionSignature

`func (o *TradeDetailDTO) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *TradeDetailDTO) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *TradeDetailDTO) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.


### GetTokenName

`func (o *TradeDetailDTO) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TradeDetailDTO) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TradeDetailDTO) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *TradeDetailDTO) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *TradeDetailDTO) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *TradeDetailDTO) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTokenImageUrl

`func (o *TradeDetailDTO) GetTokenImageUrl() string`

GetTokenImageUrl returns the TokenImageUrl field if non-nil, zero value otherwise.

### GetTokenImageUrlOk

`func (o *TradeDetailDTO) GetTokenImageUrlOk() (*string, bool)`

GetTokenImageUrlOk returns a tuple with the TokenImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenImageUrl

`func (o *TradeDetailDTO) SetTokenImageUrl(v string)`

SetTokenImageUrl sets TokenImageUrl field to given value.


### GetTokenAddress

`func (o *TradeDetailDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TradeDetailDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TradeDetailDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetPoolAddress

`func (o *TradeDetailDTO) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *TradeDetailDTO) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *TradeDetailDTO) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetDexProgramAddress

`func (o *TradeDetailDTO) GetDexProgramAddress() string`

GetDexProgramAddress returns the DexProgramAddress field if non-nil, zero value otherwise.

### GetDexProgramAddressOk

`func (o *TradeDetailDTO) GetDexProgramAddressOk() (*string, bool)`

GetDexProgramAddressOk returns a tuple with the DexProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexProgramAddress

`func (o *TradeDetailDTO) SetDexProgramAddress(v string)`

SetDexProgramAddress sets DexProgramAddress field to given value.


### GetDexProtocolFamily

`func (o *TradeDetailDTO) GetDexProtocolFamily() string`

GetDexProtocolFamily returns the DexProtocolFamily field if non-nil, zero value otherwise.

### GetDexProtocolFamilyOk

`func (o *TradeDetailDTO) GetDexProtocolFamilyOk() (*string, bool)`

GetDexProtocolFamilyOk returns a tuple with the DexProtocolFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexProtocolFamily

`func (o *TradeDetailDTO) SetDexProtocolFamily(v string)`

SetDexProtocolFamily sets DexProtocolFamily field to given value.


### GetDexImage

`func (o *TradeDetailDTO) GetDexImage() string`

GetDexImage returns the DexImage field if non-nil, zero value otherwise.

### GetDexImageOk

`func (o *TradeDetailDTO) GetDexImageOk() (*string, bool)`

GetDexImageOk returns a tuple with the DexImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexImage

`func (o *TradeDetailDTO) SetDexImage(v string)`

SetDexImage sets DexImage field to given value.


### GetTokenAmount

`func (o *TradeDetailDTO) GetTokenAmount() string`

GetTokenAmount returns the TokenAmount field if non-nil, zero value otherwise.

### GetTokenAmountOk

`func (o *TradeDetailDTO) GetTokenAmountOk() (*string, bool)`

GetTokenAmountOk returns a tuple with the TokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmount

`func (o *TradeDetailDTO) SetTokenAmount(v string)`

SetTokenAmount sets TokenAmount field to given value.


### GetTokenPriceInUsd

`func (o *TradeDetailDTO) GetTokenPriceInUsd() string`

GetTokenPriceInUsd returns the TokenPriceInUsd field if non-nil, zero value otherwise.

### GetTokenPriceInUsdOk

`func (o *TradeDetailDTO) GetTokenPriceInUsdOk() (*string, bool)`

GetTokenPriceInUsdOk returns a tuple with the TokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPriceInUsd

`func (o *TradeDetailDTO) SetTokenPriceInUsd(v string)`

SetTokenPriceInUsd sets TokenPriceInUsd field to given value.


### GetTokenAmountInUsd

`func (o *TradeDetailDTO) GetTokenAmountInUsd() string`

GetTokenAmountInUsd returns the TokenAmountInUsd field if non-nil, zero value otherwise.

### GetTokenAmountInUsdOk

`func (o *TradeDetailDTO) GetTokenAmountInUsdOk() (*string, bool)`

GetTokenAmountInUsdOk returns a tuple with the TokenAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmountInUsd

`func (o *TradeDetailDTO) SetTokenAmountInUsd(v string)`

SetTokenAmountInUsd sets TokenAmountInUsd field to given value.


### GetType

`func (o *TradeDetailDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TradeDetailDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TradeDetailDTO) SetType(v string)`

SetType sets Type field to given value.


### GetSideTokenName

`func (o *TradeDetailDTO) GetSideTokenName() string`

GetSideTokenName returns the SideTokenName field if non-nil, zero value otherwise.

### GetSideTokenNameOk

`func (o *TradeDetailDTO) GetSideTokenNameOk() (*string, bool)`

GetSideTokenNameOk returns a tuple with the SideTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenName

`func (o *TradeDetailDTO) SetSideTokenName(v string)`

SetSideTokenName sets SideTokenName field to given value.


### GetSideTokenSymbol

`func (o *TradeDetailDTO) GetSideTokenSymbol() string`

GetSideTokenSymbol returns the SideTokenSymbol field if non-nil, zero value otherwise.

### GetSideTokenSymbolOk

`func (o *TradeDetailDTO) GetSideTokenSymbolOk() (*string, bool)`

GetSideTokenSymbolOk returns a tuple with the SideTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenSymbol

`func (o *TradeDetailDTO) SetSideTokenSymbol(v string)`

SetSideTokenSymbol sets SideTokenSymbol field to given value.


### GetSideTokenImageUrl

`func (o *TradeDetailDTO) GetSideTokenImageUrl() string`

GetSideTokenImageUrl returns the SideTokenImageUrl field if non-nil, zero value otherwise.

### GetSideTokenImageUrlOk

`func (o *TradeDetailDTO) GetSideTokenImageUrlOk() (*string, bool)`

GetSideTokenImageUrlOk returns a tuple with the SideTokenImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenImageUrl

`func (o *TradeDetailDTO) SetSideTokenImageUrl(v string)`

SetSideTokenImageUrl sets SideTokenImageUrl field to given value.


### GetSideTokenAddress

`func (o *TradeDetailDTO) GetSideTokenAddress() string`

GetSideTokenAddress returns the SideTokenAddress field if non-nil, zero value otherwise.

### GetSideTokenAddressOk

`func (o *TradeDetailDTO) GetSideTokenAddressOk() (*string, bool)`

GetSideTokenAddressOk returns a tuple with the SideTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenAddress

`func (o *TradeDetailDTO) SetSideTokenAddress(v string)`

SetSideTokenAddress sets SideTokenAddress field to given value.


### GetSideTokenPriceInUsd

`func (o *TradeDetailDTO) GetSideTokenPriceInUsd() string`

GetSideTokenPriceInUsd returns the SideTokenPriceInUsd field if non-nil, zero value otherwise.

### GetSideTokenPriceInUsdOk

`func (o *TradeDetailDTO) GetSideTokenPriceInUsdOk() (*string, bool)`

GetSideTokenPriceInUsdOk returns a tuple with the SideTokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenPriceInUsd

`func (o *TradeDetailDTO) SetSideTokenPriceInUsd(v string)`

SetSideTokenPriceInUsd sets SideTokenPriceInUsd field to given value.


### GetSideTokenAmount

`func (o *TradeDetailDTO) GetSideTokenAmount() string`

GetSideTokenAmount returns the SideTokenAmount field if non-nil, zero value otherwise.

### GetSideTokenAmountOk

`func (o *TradeDetailDTO) GetSideTokenAmountOk() (*string, bool)`

GetSideTokenAmountOk returns a tuple with the SideTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenAmount

`func (o *TradeDetailDTO) SetSideTokenAmount(v string)`

SetSideTokenAmount sets SideTokenAmount field to given value.


### GetSideTokenAmountInUsd

`func (o *TradeDetailDTO) GetSideTokenAmountInUsd() string`

GetSideTokenAmountInUsd returns the SideTokenAmountInUsd field if non-nil, zero value otherwise.

### GetSideTokenAmountInUsdOk

`func (o *TradeDetailDTO) GetSideTokenAmountInUsdOk() (*string, bool)`

GetSideTokenAmountInUsdOk returns a tuple with the SideTokenAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideTokenAmountInUsd

`func (o *TradeDetailDTO) SetSideTokenAmountInUsd(v string)`

SetSideTokenAmountInUsd sets SideTokenAmountInUsd field to given value.


### GetStatus

`func (o *TradeDetailDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradeDetailDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradeDetailDTO) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


