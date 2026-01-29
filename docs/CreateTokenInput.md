# CreateTokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** | DEX identifier for the trade | 
**UserAddress** | **string** | Public key of the wallet initiating the transaction | 
**PriorityFee** | Pointer to **string** | Priority fee in SOL to increase transaction processing speed | [optional] 
**Name** | **string** | Name of the token to be created | 
**Symbol** | **string** | Token symbol/ticker | 
**Uri** | Pointer to **string** | URI for token metadata (usually points to image or JSON) | [optional] 
**Image** | Pointer to **string** | Token image URL (Base64 or HTTPS) | [optional] 
**Extra** | Pointer to **map[string]interface{}** | Additional metadata about the created token | [optional] 

## Methods

### NewCreateTokenInput

`func NewCreateTokenInput(dex string, userAddress string, name string, symbol string, ) *CreateTokenInput`

NewCreateTokenInput instantiates a new CreateTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenInputWithDefaults

`func NewCreateTokenInputWithDefaults() *CreateTokenInput`

NewCreateTokenInputWithDefaults instantiates a new CreateTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *CreateTokenInput) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *CreateTokenInput) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *CreateTokenInput) SetDex(v string)`

SetDex sets Dex field to given value.


### GetUserAddress

`func (o *CreateTokenInput) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *CreateTokenInput) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *CreateTokenInput) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetPriorityFee

`func (o *CreateTokenInput) GetPriorityFee() string`

GetPriorityFee returns the PriorityFee field if non-nil, zero value otherwise.

### GetPriorityFeeOk

`func (o *CreateTokenInput) GetPriorityFeeOk() (*string, bool)`

GetPriorityFeeOk returns a tuple with the PriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFee

`func (o *CreateTokenInput) SetPriorityFee(v string)`

SetPriorityFee sets PriorityFee field to given value.

### HasPriorityFee

`func (o *CreateTokenInput) HasPriorityFee() bool`

HasPriorityFee returns a boolean if a field has been set.

### GetName

`func (o *CreateTokenInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTokenInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTokenInput) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *CreateTokenInput) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CreateTokenInput) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CreateTokenInput) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetUri

`func (o *CreateTokenInput) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CreateTokenInput) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CreateTokenInput) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CreateTokenInput) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetImage

`func (o *CreateTokenInput) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateTokenInput) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateTokenInput) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateTokenInput) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetExtra

`func (o *CreateTokenInput) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateTokenInput) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateTokenInput) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CreateTokenInput) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


