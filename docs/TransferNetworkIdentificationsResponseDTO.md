# TransferNetworkIdentificationsResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | DTO.KYT.TRANSFER_NETWORK_IDENTIFICATIONS_RESPONSE.COUNT | 
**NetworkIdentificationOrgs** | [**[]NetworkIdentificationOrg**](NetworkIdentificationOrg.md) | DTO.KYT.TRANSFER_NETWORK_IDENTIFICATIONS_RESPONSE.NETWORK_IDENTIFICATION_ORGS | 

## Methods

### NewTransferNetworkIdentificationsResponseDTO

`func NewTransferNetworkIdentificationsResponseDTO(count int64, networkIdentificationOrgs []NetworkIdentificationOrg, ) *TransferNetworkIdentificationsResponseDTO`

NewTransferNetworkIdentificationsResponseDTO instantiates a new TransferNetworkIdentificationsResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferNetworkIdentificationsResponseDTOWithDefaults

`func NewTransferNetworkIdentificationsResponseDTOWithDefaults() *TransferNetworkIdentificationsResponseDTO`

NewTransferNetworkIdentificationsResponseDTOWithDefaults instantiates a new TransferNetworkIdentificationsResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransferNetworkIdentificationsResponseDTO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransferNetworkIdentificationsResponseDTO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransferNetworkIdentificationsResponseDTO) SetCount(v int64)`

SetCount sets Count field to given value.


### GetNetworkIdentificationOrgs

`func (o *TransferNetworkIdentificationsResponseDTO) GetNetworkIdentificationOrgs() []NetworkIdentificationOrg`

GetNetworkIdentificationOrgs returns the NetworkIdentificationOrgs field if non-nil, zero value otherwise.

### GetNetworkIdentificationOrgsOk

`func (o *TransferNetworkIdentificationsResponseDTO) GetNetworkIdentificationOrgsOk() (*[]NetworkIdentificationOrg, bool)`

GetNetworkIdentificationOrgsOk returns a tuple with the NetworkIdentificationOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdentificationOrgs

`func (o *TransferNetworkIdentificationsResponseDTO) SetNetworkIdentificationOrgs(v []NetworkIdentificationOrg)`

SetNetworkIdentificationOrgs sets NetworkIdentificationOrgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


