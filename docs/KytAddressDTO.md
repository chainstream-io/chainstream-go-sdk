# KytAddressDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | DTO.KYT.KYT_ADDRESS.ID | 
**OrgId** | **string** | DTO.KYT.KYT_ADDRESS.ORG_ID | 
**Address** | **string** | DTO.KYT.KYT_ADDRESS.ADDRESS | 
**CreatedAt** | **time.Time** | DTO.KYT.KYT_ADDRESS.CREATED_AT | 
**UpdatedAt** | **time.Time** | DTO.KYT.KYT_ADDRESS.UPDATED_AT | 

## Methods

### NewKytAddressDTO

`func NewKytAddressDTO(id float32, orgId string, address string, createdAt time.Time, updatedAt time.Time, ) *KytAddressDTO`

NewKytAddressDTO instantiates a new KytAddressDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytAddressDTOWithDefaults

`func NewKytAddressDTOWithDefaults() *KytAddressDTO`

NewKytAddressDTOWithDefaults instantiates a new KytAddressDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KytAddressDTO) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KytAddressDTO) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KytAddressDTO) SetId(v float32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *KytAddressDTO) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *KytAddressDTO) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *KytAddressDTO) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAddress

`func (o *KytAddressDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KytAddressDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KytAddressDTO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCreatedAt

`func (o *KytAddressDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KytAddressDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KytAddressDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KytAddressDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KytAddressDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KytAddressDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


