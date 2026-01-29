# MoonshotCreateTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** | DEX identifier for the trade | 
**UserAddress** | **string** | Public key of the wallet initiating the transaction | 
**PriorityFee** | Pointer to **string** | Priority fee in SOL to increase transaction processing speed | [optional] 
**Name** | **string** | Name of the token | 
**Symbol** | **string** | Token symbol/ticker | 
**MigrationDex** | **string** | Target DEX for token migration | 
**Icon** | **string** | Token icon URL (Base64 or HTTPS) | 
**Description** | **string** | Token description | 
**Links** | [**[]Link**](Link.md) | Social and website links | 
**Banner** | Pointer to **string** | Banner image URL (defaults to icon if not provided) | [optional] 
**TokenAmount** | **string** | Total token supply amount (in base units) | 

## Methods

### NewMoonshotCreateTokenInput

`func NewMoonshotCreateTokenInput(dex string, userAddress string, name string, symbol string, migrationDex string, icon string, description string, links []Link, tokenAmount string, ) *MoonshotCreateTokenInput`

NewMoonshotCreateTokenInput instantiates a new MoonshotCreateTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoonshotCreateTokenInputWithDefaults

`func NewMoonshotCreateTokenInputWithDefaults() *MoonshotCreateTokenInput`

NewMoonshotCreateTokenInputWithDefaults instantiates a new MoonshotCreateTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *MoonshotCreateTokenInput) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *MoonshotCreateTokenInput) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *MoonshotCreateTokenInput) SetDex(v string)`

SetDex sets Dex field to given value.


### GetUserAddress

`func (o *MoonshotCreateTokenInput) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *MoonshotCreateTokenInput) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *MoonshotCreateTokenInput) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetPriorityFee

`func (o *MoonshotCreateTokenInput) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *MoonshotCreateTokenInput) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *MoonshotCreateTokenInput) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.

### HasPriorityFee

`func (o *MoonshotCreateTokenInput) HasPriorityFee() bool`

HasPriorityFee returns a boolean if a field has been set.

### GetName

`func (o *MoonshotCreateTokenInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoonshotCreateTokenInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoonshotCreateTokenInput) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *MoonshotCreateTokenInput) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MoonshotCreateTokenInput) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MoonshotCreateTokenInput) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMigrationDex

`func (o *MoonshotCreateTokenInput) GetMigrationDex() string`

GetMigrationDex returns the MigrationDex field if non-nil, zero value otherwise.

### GetMigrationDexOk

`func (o *MoonshotCreateTokenInput) GetMigrationDexOk() (*string, bool)`

GetMigrationDexOk returns a tuple with the MigrationDex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationDex

`func (o *MoonshotCreateTokenInput) SetMigrationDex(v string)`

SetMigrationDex sets MigrationDex field to given value.


### GetIcon

`func (o *MoonshotCreateTokenInput) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MoonshotCreateTokenInput) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MoonshotCreateTokenInput) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetDescription

`func (o *MoonshotCreateTokenInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MoonshotCreateTokenInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MoonshotCreateTokenInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLinks

`func (o *MoonshotCreateTokenInput) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MoonshotCreateTokenInput) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MoonshotCreateTokenInput) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetBanner

`func (o *MoonshotCreateTokenInput) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *MoonshotCreateTokenInput) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *MoonshotCreateTokenInput) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *MoonshotCreateTokenInput) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetTokenAmount

`func (o *MoonshotCreateTokenInput) GetTokenAmount() string`

GetTokenAmount returns the TokenAmount field if non-nil, zero value otherwise.

### GetTokenAmountOk

`func (o *MoonshotCreateTokenInput) GetTokenAmountOk() (*string, bool)`

GetTokenAmountOk returns a tuple with the TokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmount

`func (o *MoonshotCreateTokenInput) SetTokenAmount(v string)`

SetTokenAmount sets TokenAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


