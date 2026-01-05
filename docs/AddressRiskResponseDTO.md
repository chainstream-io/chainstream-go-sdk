# AddressRiskResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.ADDRESS | 
**Risk** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.RISK | 
**RiskReason** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.RISK_REASON | 
**AddressType** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.ADDRESS_TYPE | 
**Cluster** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.CLUSTER | 
**AddressIdentifications** | **[]string** | DTO.KYT.ADDRESS_RISK_RESPONSE.ADDRESS_IDENTIFICATIONS | 
**Exposures** | [**[]AddressExposure**](AddressExposure.md) | DTO.KYT.ADDRESS_RISK_RESPONSE.EXPOSURES | 
**Triggers** | **[]string** | DTO.KYT.ADDRESS_RISK_RESPONSE.TRIGGERS | 
**Status** | **string** | DTO.KYT.ADDRESS_RISK_RESPONSE.STATUS | 

## Methods

### NewAddressRiskResponseDTO

`func NewAddressRiskResponseDTO(address string, risk string, riskReason string, addressType string, cluster string, addressIdentifications []string, exposures []AddressExposure, triggers []string, status string, ) *AddressRiskResponseDTO`

NewAddressRiskResponseDTO instantiates a new AddressRiskResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRiskResponseDTOWithDefaults

`func NewAddressRiskResponseDTOWithDefaults() *AddressRiskResponseDTO`

NewAddressRiskResponseDTOWithDefaults instantiates a new AddressRiskResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressRiskResponseDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressRiskResponseDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressRiskResponseDTO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetRisk

`func (o *AddressRiskResponseDTO) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AddressRiskResponseDTO) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AddressRiskResponseDTO) SetRisk(v string)`

SetRisk sets Risk field to given value.


### GetRiskReason

`func (o *AddressRiskResponseDTO) GetRiskReason() string`

GetRiskReason returns the RiskReason field if non-nil, zero value otherwise.

### GetRiskReasonOk

`func (o *AddressRiskResponseDTO) GetRiskReasonOk() (*string, bool)`

GetRiskReasonOk returns a tuple with the RiskReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskReason

`func (o *AddressRiskResponseDTO) SetRiskReason(v string)`

SetRiskReason sets RiskReason field to given value.


### GetAddressType

`func (o *AddressRiskResponseDTO) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AddressRiskResponseDTO) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AddressRiskResponseDTO) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetCluster

`func (o *AddressRiskResponseDTO) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddressRiskResponseDTO) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddressRiskResponseDTO) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetAddressIdentifications

`func (o *AddressRiskResponseDTO) GetAddressIdentifications() []string`

GetAddressIdentifications returns the AddressIdentifications field if non-nil, zero value otherwise.

### GetAddressIdentificationsOk

`func (o *AddressRiskResponseDTO) GetAddressIdentificationsOk() (*[]string, bool)`

GetAddressIdentificationsOk returns a tuple with the AddressIdentifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressIdentifications

`func (o *AddressRiskResponseDTO) SetAddressIdentifications(v []string)`

SetAddressIdentifications sets AddressIdentifications field to given value.


### GetExposures

`func (o *AddressRiskResponseDTO) GetExposures() []AddressExposure`

GetExposures returns the Exposures field if non-nil, zero value otherwise.

### GetExposuresOk

`func (o *AddressRiskResponseDTO) GetExposuresOk() (*[]AddressExposure, bool)`

GetExposuresOk returns a tuple with the Exposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposures

`func (o *AddressRiskResponseDTO) SetExposures(v []AddressExposure)`

SetExposures sets Exposures field to given value.


### GetTriggers

`func (o *AddressRiskResponseDTO) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *AddressRiskResponseDTO) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *AddressRiskResponseDTO) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.


### GetStatus

`func (o *AddressRiskResponseDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddressRiskResponseDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddressRiskResponseDTO) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


