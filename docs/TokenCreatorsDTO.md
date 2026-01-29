# TokenCreatorsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Creator wallet address | [optional] 
**Share** | Pointer to **int64** | Creator share percentage (0-100) | [optional] 
**IsVerified** | Pointer to **bool** | Whether the creator is verified | [optional] 

## Methods

### NewTokenCreatorsDTO

`func NewTokenCreatorsDTO() *TokenCreatorsDTO`

NewTokenCreatorsDTO instantiates a new TokenCreatorsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreatorsDTOWithDefaults

`func NewTokenCreatorsDTOWithDefaults() *TokenCreatorsDTO`

NewTokenCreatorsDTOWithDefaults instantiates a new TokenCreatorsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenCreatorsDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenCreatorsDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenCreatorsDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TokenCreatorsDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetShare

`func (o *TokenCreatorsDTO) GetShare() int64`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *TokenCreatorsDTO) GetShareOk() (*int64, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *TokenCreatorsDTO) SetShare(v int64)`

SetShare sets Share field to given value.

### HasShare

`func (o *TokenCreatorsDTO) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetIsVerified

`func (o *TokenCreatorsDTO) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *TokenCreatorsDTO) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *TokenCreatorsDTO) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *TokenCreatorsDTO) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


