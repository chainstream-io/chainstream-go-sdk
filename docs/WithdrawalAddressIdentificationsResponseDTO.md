# WithdrawalAddressIdentificationsResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainalysisIdentifications** | [**[]ChainalysisAddressIdentification**](ChainalysisAddressIdentification.md) | DTO.KYT.WITHDRAWAL_ADDRESS_IDENTIFICATIONS_RESPONSE.CHAINALYSIS_IDENTIFICATIONS | 
**CustomAddresses** | **[]string** | DTO.KYT.WITHDRAWAL_ADDRESS_IDENTIFICATIONS_RESPONSE.CUSTOM_ADDRESSES | 

## Methods

### NewWithdrawalAddressIdentificationsResponseDTO

`func NewWithdrawalAddressIdentificationsResponseDTO(chainalysisIdentifications []ChainalysisAddressIdentification, customAddresses []string, ) *WithdrawalAddressIdentificationsResponseDTO`

NewWithdrawalAddressIdentificationsResponseDTO instantiates a new WithdrawalAddressIdentificationsResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalAddressIdentificationsResponseDTOWithDefaults

`func NewWithdrawalAddressIdentificationsResponseDTOWithDefaults() *WithdrawalAddressIdentificationsResponseDTO`

NewWithdrawalAddressIdentificationsResponseDTOWithDefaults instantiates a new WithdrawalAddressIdentificationsResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainalysisIdentifications

`func (o *WithdrawalAddressIdentificationsResponseDTO) GetChainalysisIdentifications() []ChainalysisAddressIdentification`

GetChainalysisIdentifications returns the ChainalysisIdentifications field if non-nil, zero value otherwise.

### GetChainalysisIdentificationsOk

`func (o *WithdrawalAddressIdentificationsResponseDTO) GetChainalysisIdentificationsOk() (*[]ChainalysisAddressIdentification, bool)`

GetChainalysisIdentificationsOk returns a tuple with the ChainalysisIdentifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainalysisIdentifications

`func (o *WithdrawalAddressIdentificationsResponseDTO) SetChainalysisIdentifications(v []ChainalysisAddressIdentification)`

SetChainalysisIdentifications sets ChainalysisIdentifications field to given value.


### GetCustomAddresses

`func (o *WithdrawalAddressIdentificationsResponseDTO) GetCustomAddresses() []string`

GetCustomAddresses returns the CustomAddresses field if non-nil, zero value otherwise.

### GetCustomAddressesOk

`func (o *WithdrawalAddressIdentificationsResponseDTO) GetCustomAddressesOk() (*[]string, bool)`

GetCustomAddressesOk returns a tuple with the CustomAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAddresses

`func (o *WithdrawalAddressIdentificationsResponseDTO) SetCustomAddresses(v []string)`

SetCustomAddresses sets CustomAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


