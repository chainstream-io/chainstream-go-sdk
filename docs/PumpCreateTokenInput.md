# PumpCreateTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** | DTO.DEX.IDENTIFIER | 
**UserAddress** | **string** | DTO.DEX.WALLET | 
**PriorityFee** | Pointer to **string** | DTO.DEX.BASE.PRIORITY_FEE | [optional] 
**Twitter** | Pointer to **string** | DTO.DEX.PUMPFUN.MINT.TWITTER | [optional] 
**Telegram** | Pointer to **string** | DTO.DEX.PUMPFUN.MINT.TELEGRAM | [optional] 
**Website** | Pointer to **string** | DTO.DEX.PUMPFUN.MINT.WEBSITE | [optional] 
**Name** | **string** | DTO.DEX.PUMPFUN.MINT.NAME | 
**Symbol** | **string** | DTO.DEX.PUMPFUN.MINT.SYMBOL | 
**MigrationDex** | **string** | DTO.DEX.PUMPFUN.MINT.MIGRATION_DEX | 
**Image** | **string** | DTO.DEX.PUMPFUN.MINT.IMAGE | 
**MintAddress** | **string** | DTO.DEX.PUMPFUN.MINT.MINT_ADDRESS | 
**Description** | **string** | DTO.DEX.PUMPFUN.MINT.DESCRIPTION | 

## Methods

### NewPumpCreateTokenInput

`func NewPumpCreateTokenInput(dex string, userAddress string, name string, symbol string, migrationDex string, image string, mintAddress string, description string, ) *PumpCreateTokenInput`

NewPumpCreateTokenInput instantiates a new PumpCreateTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPumpCreateTokenInputWithDefaults

`func NewPumpCreateTokenInputWithDefaults() *PumpCreateTokenInput`

NewPumpCreateTokenInputWithDefaults instantiates a new PumpCreateTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *PumpCreateTokenInput) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *PumpCreateTokenInput) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *PumpCreateTokenInput) SetDex(v string)`

SetDex sets Dex field to given value.


### GetUserAddress

`func (o *PumpCreateTokenInput) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *PumpCreateTokenInput) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *PumpCreateTokenInput) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetPriorityFee

`func (o *PumpCreateTokenInput) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *PumpCreateTokenInput) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *PumpCreateTokenInput) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.

### HasPriorityFee

`func (o *PumpCreateTokenInput) HasPriorityFee() bool`

HasPriorityFee returns a boolean if a field has been set.

### GetTwitter

`func (o *PumpCreateTokenInput) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *PumpCreateTokenInput) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *PumpCreateTokenInput) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *PumpCreateTokenInput) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetTelegram

`func (o *PumpCreateTokenInput) GetTelegram() string`

GetTelegram returns the Telegram field if non-nil, zero value otherwise.

### GetTelegramOk

`func (o *PumpCreateTokenInput) GetTelegramOk() (*string, bool)`

GetTelegramOk returns a tuple with the Telegram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegram

`func (o *PumpCreateTokenInput) SetTelegram(v string)`

SetTelegram sets Telegram field to given value.

### HasTelegram

`func (o *PumpCreateTokenInput) HasTelegram() bool`

HasTelegram returns a boolean if a field has been set.

### GetWebsite

`func (o *PumpCreateTokenInput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *PumpCreateTokenInput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *PumpCreateTokenInput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *PumpCreateTokenInput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetName

`func (o *PumpCreateTokenInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PumpCreateTokenInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PumpCreateTokenInput) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *PumpCreateTokenInput) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PumpCreateTokenInput) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PumpCreateTokenInput) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMigrationDex

`func (o *PumpCreateTokenInput) GetMigrationDex() string`

GetMigrationDex returns the MigrationDex field if non-nil, zero value otherwise.

### GetMigrationDexOk

`func (o *PumpCreateTokenInput) GetMigrationDexOk() (*string, bool)`

GetMigrationDexOk returns a tuple with the MigrationDex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationDex

`func (o *PumpCreateTokenInput) SetMigrationDex(v string)`

SetMigrationDex sets MigrationDex field to given value.


### GetImage

`func (o *PumpCreateTokenInput) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PumpCreateTokenInput) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PumpCreateTokenInput) SetImage(v string)`

SetImage sets Image field to given value.


### GetMintAddress

`func (o *PumpCreateTokenInput) GetMintAddress() string`

GetMintAddress returns the MintAddress field if non-nil, zero value otherwise.

### GetMintAddressOk

`func (o *PumpCreateTokenInput) GetMintAddressOk() (*string, bool)`

GetMintAddressOk returns a tuple with the MintAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAddress

`func (o *PumpCreateTokenInput) SetMintAddress(v string)`

SetMintAddress sets MintAddress field to given value.


### GetDescription

`func (o *PumpCreateTokenInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PumpCreateTokenInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PumpCreateTokenInput) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


