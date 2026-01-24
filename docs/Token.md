# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | DTO.TOKEN.METADATA.CHAIN | 
**Decimals** | **int64** | DTO.TOKEN.METADATA.DECIMALS | 
**Name** | **string** | DTO.TOKEN.METADATA.NAME | 
**Symbol** | **string** | DTO.TOKEN.METADATA.SYMBOL | 
**MetadataAddress** | Pointer to **string** | DTO.TOKEN.METADATA.METADATA_ADDRESS | [optional] 
**Address** | **string** | DTO.TOKEN.METADATA.ADDRESS | 
**TokenCreators** | Pointer to [**[]TokenCreatorsDTO**](TokenCreatorsDTO.md) | DTO.TOKEN.METADATA.TOKEN_CREATORS | [optional] 
**ImageUrl** | Pointer to **string** | DTO.TOKEN.METADATA.IMAGE_URL | [optional] 
**Uri** | Pointer to **string** | DTO.TOKEN.METADATA.URI | [optional] 
**Extra** | Pointer to [**TokenExtraDTO**](TokenExtraDTO.md) | DTO.TOKEN.METADATA.EXTRA | [optional] 
**SocialMedias** | Pointer to [**TokenSocialMediasDTO**](TokenSocialMediasDTO.md) | DTO.TOKEN.METADATA.SOCIAL_MEDIAS | [optional] 
**TokenCreatedAt** | Pointer to **int64** | DTO.TOKEN.METADATA.TOKEN_CREATED_AT | [optional] 
**Description** | Pointer to **string** | DTO.TOKEN.METADATA.DESCRIPTION | [optional] 
**DevTotalTokens** | Pointer to **string** | DTO.TOKEN.METADATA.DEV_TOTAL_TOKENS | [optional] 
**DevLastTokenCreatedAt** | Pointer to **string** | DTO.TOKEN.METADATA.DEV_LAST_TOKEN_CREATED_AT | [optional] 
**Market** | Pointer to **string** | DTO.TOKEN.MARKET_ID | [optional] 
**Extension** | Pointer to **map[string]interface{}** | DTO.TOKEN.EXTENSION | [optional] 
**Stats** | Pointer to [**TokenStat**](TokenStat.md) | DTO.TOKEN.STATS | [optional] 
**Liquidity** | Pointer to [**[]DexPoolDTO**](DexPoolDTO.md) | DTO.TOKEN.LIQUIDITY | [optional] 
**MarketData** | [**TokenMarketData**](TokenMarketData.md) | DTO.TOKEN.MARKET_CAP_INFO | 

## Methods

### NewToken

`func NewToken(chain string, decimals int64, name string, symbol string, address string, marketData TokenMarketData, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *Token) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Token) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Token) SetChain(v string)`

SetChain sets Chain field to given value.


### GetDecimals

`func (o *Token) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Token) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Token) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.


### GetName

`func (o *Token) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Token) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Token) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *Token) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Token) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Token) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMetadataAddress

`func (o *Token) GetMetadataAddress() string`

GetMetadataAddress returns the MetadataAddress field if non-nil, zero value otherwise.

### GetMetadataAddressOk

`func (o *Token) GetMetadataAddressOk() (*string, bool)`

GetMetadataAddressOk returns a tuple with the MetadataAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAddress

`func (o *Token) SetMetadataAddress(v string)`

SetMetadataAddress sets MetadataAddress field to given value.

### HasMetadataAddress

`func (o *Token) HasMetadataAddress() bool`

HasMetadataAddress returns a boolean if a field has been set.

### GetAddress

`func (o *Token) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Token) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Token) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokenCreators

`func (o *Token) GetTokenCreators() []TokenCreatorsDTO`

GetTokenCreators returns the TokenCreators field if non-nil, zero value otherwise.

### GetTokenCreatorsOk

`func (o *Token) GetTokenCreatorsOk() (*[]TokenCreatorsDTO, bool)`

GetTokenCreatorsOk returns a tuple with the TokenCreators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCreators

`func (o *Token) SetTokenCreators(v []TokenCreatorsDTO)`

SetTokenCreators sets TokenCreators field to given value.

### HasTokenCreators

`func (o *Token) HasTokenCreators() bool`

HasTokenCreators returns a boolean if a field has been set.

### GetImageUrl

`func (o *Token) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Token) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Token) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Token) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetUri

`func (o *Token) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Token) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Token) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Token) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetExtra

`func (o *Token) GetExtra() TokenExtraDTO`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Token) GetExtraOk() (*TokenExtraDTO, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Token) SetExtra(v TokenExtraDTO)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Token) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetSocialMedias

`func (o *Token) GetSocialMedias() TokenSocialMediasDTO`

GetSocialMedias returns the SocialMedias field if non-nil, zero value otherwise.

### GetSocialMediasOk

`func (o *Token) GetSocialMediasOk() (*TokenSocialMediasDTO, bool)`

GetSocialMediasOk returns a tuple with the SocialMedias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialMedias

`func (o *Token) SetSocialMedias(v TokenSocialMediasDTO)`

SetSocialMedias sets SocialMedias field to given value.

### HasSocialMedias

`func (o *Token) HasSocialMedias() bool`

HasSocialMedias returns a boolean if a field has been set.

### GetTokenCreatedAt

`func (o *Token) GetTokenCreatedAt() int64`

GetTokenCreatedAt returns the TokenCreatedAt field if non-nil, zero value otherwise.

### GetTokenCreatedAtOk

`func (o *Token) GetTokenCreatedAtOk() (*int64, bool)`

GetTokenCreatedAtOk returns a tuple with the TokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCreatedAt

`func (o *Token) SetTokenCreatedAt(v int64)`

SetTokenCreatedAt sets TokenCreatedAt field to given value.

### HasTokenCreatedAt

`func (o *Token) HasTokenCreatedAt() bool`

HasTokenCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Token) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Token) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Token) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Token) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevTotalTokens

`func (o *Token) GetDevTotalTokens() string`

GetDevTotalTokens returns the DevTotalTokens field if non-nil, zero value otherwise.

### GetDevTotalTokensOk

`func (o *Token) GetDevTotalTokensOk() (*string, bool)`

GetDevTotalTokensOk returns a tuple with the DevTotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTotalTokens

`func (o *Token) SetDevTotalTokens(v string)`

SetDevTotalTokens sets DevTotalTokens field to given value.

### HasDevTotalTokens

`func (o *Token) HasDevTotalTokens() bool`

HasDevTotalTokens returns a boolean if a field has been set.

### GetDevLastTokenCreatedAt

`func (o *Token) GetDevLastTokenCreatedAt() string`

GetDevLastTokenCreatedAt returns the DevLastTokenCreatedAt field if non-nil, zero value otherwise.

### GetDevLastTokenCreatedAtOk

`func (o *Token) GetDevLastTokenCreatedAtOk() (*string, bool)`

GetDevLastTokenCreatedAtOk returns a tuple with the DevLastTokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevLastTokenCreatedAt

`func (o *Token) SetDevLastTokenCreatedAt(v string)`

SetDevLastTokenCreatedAt sets DevLastTokenCreatedAt field to given value.

### HasDevLastTokenCreatedAt

`func (o *Token) HasDevLastTokenCreatedAt() bool`

HasDevLastTokenCreatedAt returns a boolean if a field has been set.

### GetMarket

`func (o *Token) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Token) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Token) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Token) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetExtension

`func (o *Token) GetExtension() map[string]interface{}`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Token) GetExtensionOk() (*map[string]interface{}, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Token) SetExtension(v map[string]interface{})`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Token) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetStats

`func (o *Token) GetStats() TokenStat`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Token) GetStatsOk() (*TokenStat, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Token) SetStats(v TokenStat)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Token) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetLiquidity

`func (o *Token) GetLiquidity() []DexPoolDTO`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *Token) GetLiquidityOk() (*[]DexPoolDTO, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *Token) SetLiquidity(v []DexPoolDTO)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *Token) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetMarketData

`func (o *Token) GetMarketData() TokenMarketData`

GetMarketData returns the MarketData field if non-nil, zero value otherwise.

### GetMarketDataOk

`func (o *Token) GetMarketDataOk() (*TokenMarketData, bool)`

GetMarketDataOk returns a tuple with the MarketData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketData

`func (o *Token) SetMarketData(v TokenMarketData)`

SetMarketData sets MarketData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


