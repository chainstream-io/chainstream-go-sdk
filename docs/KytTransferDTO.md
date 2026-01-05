# KytTransferDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | DTO.KYT.KYT_TRANSFER.ID | 
**OrgId** | **string** | DTO.KYT.KYT_TRANSFER.ORG_ID | 
**TxHash** | **string** | DTO.KYT.KYT_TRANSFER.TX_HASH | 
**ExternalId** | **string** | DTO.KYT.KYT_TRANSFER.EXTERNAL_ID | 
**CreatedAt** | **time.Time** | DTO.KYT.KYT_TRANSFER.CREATED_AT | 
**UpdatedAt** | **time.Time** | DTO.KYT.KYT_TRANSFER.UPDATED_AT | 

## Methods

### NewKytTransferDTO

`func NewKytTransferDTO(id float32, orgId string, txHash string, externalId string, createdAt time.Time, updatedAt time.Time, ) *KytTransferDTO`

NewKytTransferDTO instantiates a new KytTransferDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytTransferDTOWithDefaults

`func NewKytTransferDTOWithDefaults() *KytTransferDTO`

NewKytTransferDTOWithDefaults instantiates a new KytTransferDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KytTransferDTO) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KytTransferDTO) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KytTransferDTO) SetId(v float32)`

SetId sets Id field to given value.


### GetOrgId

`func (o *KytTransferDTO) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *KytTransferDTO) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *KytTransferDTO) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetTxHash

`func (o *KytTransferDTO) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *KytTransferDTO) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *KytTransferDTO) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetExternalId

`func (o *KytTransferDTO) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *KytTransferDTO) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *KytTransferDTO) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCreatedAt

`func (o *KytTransferDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KytTransferDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KytTransferDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KytTransferDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KytTransferDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KytTransferDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


