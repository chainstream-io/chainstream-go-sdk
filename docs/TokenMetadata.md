# TokenMetadata

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

## Methods

### NewTokenMetadata

`func NewTokenMetadata(chain string, decimals int64, name string, symbol string, address string, ) *TokenMetadata`

NewTokenMetadata instantiates a new TokenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMetadataWithDefaults

`func NewTokenMetadataWithDefaults() *TokenMetadata`

NewTokenMetadataWithDefaults instantiates a new TokenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *TokenMetadata) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *TokenMetadata) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *TokenMetadata) SetChain(v string)`

SetChain sets Chain field to given value.


### GetDecimals

`func (o *TokenMetadata) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenMetadata) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenMetadata) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.


### GetName

`func (o *TokenMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenMetadata) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenMetadata) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenMetadata) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMetadataAddress

`func (o *TokenMetadata) GetMetadataAddress() string`

GetMetadataAddress returns the MetadataAddress field if non-nil, zero value otherwise.

### GetMetadataAddressOk

`func (o *TokenMetadata) GetMetadataAddressOk() (*string, bool)`

GetMetadataAddressOk returns a tuple with the MetadataAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAddress

`func (o *TokenMetadata) SetMetadataAddress(v string)`

SetMetadataAddress sets MetadataAddress field to given value.

### HasMetadataAddress

`func (o *TokenMetadata) HasMetadataAddress() bool`

HasMetadataAddress returns a boolean if a field has been set.

### GetAddress

`func (o *TokenMetadata) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenMetadata) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenMetadata) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokenCreators

`func (o *TokenMetadata) GetTokenCreators() []TokenCreatorsDTO`

GetTokenCreators returns the TokenCreators field if non-nil, zero value otherwise.

### GetTokenCreatorsOk

`func (o *TokenMetadata) GetTokenCreatorsOk() (*[]TokenCreatorsDTO, bool)`

GetTokenCreatorsOk returns a tuple with the TokenCreators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCreators

`func (o *TokenMetadata) SetTokenCreators(v []TokenCreatorsDTO)`

SetTokenCreators sets TokenCreators field to given value.

### HasTokenCreators

`func (o *TokenMetadata) HasTokenCreators() bool`

HasTokenCreators returns a boolean if a field has been set.

### GetImageUrl

`func (o *TokenMetadata) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *TokenMetadata) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *TokenMetadata) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *TokenMetadata) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetUri

`func (o *TokenMetadata) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TokenMetadata) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TokenMetadata) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TokenMetadata) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetExtra

`func (o *TokenMetadata) GetExtra() TokenExtraDTO`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *TokenMetadata) GetExtraOk() (*TokenExtraDTO, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *TokenMetadata) SetExtra(v TokenExtraDTO)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *TokenMetadata) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetSocialMedias

`func (o *TokenMetadata) GetSocialMedias() TokenSocialMediasDTO`

GetSocialMedias returns the SocialMedias field if non-nil, zero value otherwise.

### GetSocialMediasOk

`func (o *TokenMetadata) GetSocialMediasOk() (*TokenSocialMediasDTO, bool)`

GetSocialMediasOk returns a tuple with the SocialMedias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialMedias

`func (o *TokenMetadata) SetSocialMedias(v TokenSocialMediasDTO)`

SetSocialMedias sets SocialMedias field to given value.

### HasSocialMedias

`func (o *TokenMetadata) HasSocialMedias() bool`

HasSocialMedias returns a boolean if a field has been set.

### GetTokenCreatedAt

`func (o *TokenMetadata) GetTokenCreatedAt() int64`

GetTokenCreatedAt returns the TokenCreatedAt field if non-nil, zero value otherwise.

### GetTokenCreatedAtOk

`func (o *TokenMetadata) GetTokenCreatedAtOk() (*int64, bool)`

GetTokenCreatedAtOk returns a tuple with the TokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCreatedAt

`func (o *TokenMetadata) SetTokenCreatedAt(v int64)`

SetTokenCreatedAt sets TokenCreatedAt field to given value.

### HasTokenCreatedAt

`func (o *TokenMetadata) HasTokenCreatedAt() bool`

HasTokenCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *TokenMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevTotalTokens

`func (o *TokenMetadata) GetDevTotalTokens() string`

GetDevTotalTokens returns the DevTotalTokens field if non-nil, zero value otherwise.

### GetDevTotalTokensOk

`func (o *TokenMetadata) GetDevTotalTokensOk() (*string, bool)`

GetDevTotalTokensOk returns a tuple with the DevTotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTotalTokens

`func (o *TokenMetadata) SetDevTotalTokens(v string)`

SetDevTotalTokens sets DevTotalTokens field to given value.

### HasDevTotalTokens

`func (o *TokenMetadata) HasDevTotalTokens() bool`

HasDevTotalTokens returns a boolean if a field has been set.

### GetDevLastTokenCreatedAt

`func (o *TokenMetadata) GetDevLastTokenCreatedAt() string`

GetDevLastTokenCreatedAt returns the DevLastTokenCreatedAt field if non-nil, zero value otherwise.

### GetDevLastTokenCreatedAtOk

`func (o *TokenMetadata) GetDevLastTokenCreatedAtOk() (*string, bool)`

GetDevLastTokenCreatedAtOk returns a tuple with the DevLastTokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevLastTokenCreatedAt

`func (o *TokenMetadata) SetDevLastTokenCreatedAt(v string)`

SetDevLastTokenCreatedAt sets DevLastTokenCreatedAt field to given value.

### HasDevLastTokenCreatedAt

`func (o *TokenMetadata) HasDevLastTokenCreatedAt() bool`

HasDevLastTokenCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


