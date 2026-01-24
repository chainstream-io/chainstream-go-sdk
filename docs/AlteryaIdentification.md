# AlteryaIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.OBJECT | 
**Address** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.ADDRESS | 
**Chain** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.CHAIN | 
**RiskLevel** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.RISK_LEVEL | 
**Reason** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.REASON | 
**Labels** | **[]string** | DTO.KYT.ALTERYA_IDENTIFICATION.LABELS | 
**OnChainActivity** | [**OnChainActivity**](OnChainActivity.md) | DTO.KYT.ALTERYA_IDENTIFICATION.ON_CHAIN_ACTIVITY | 
**RiskScore** | **string** | DTO.KYT.ALTERYA_IDENTIFICATION.RISK_SCORE | 

## Methods

### NewAlteryaIdentification

`func NewAlteryaIdentification(object string, address string, chain string, riskLevel string, reason string, labels []string, onChainActivity OnChainActivity, riskScore string, ) *AlteryaIdentification`

NewAlteryaIdentification instantiates a new AlteryaIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlteryaIdentificationWithDefaults

`func NewAlteryaIdentificationWithDefaults() *AlteryaIdentification`

NewAlteryaIdentificationWithDefaults instantiates a new AlteryaIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *AlteryaIdentification) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AlteryaIdentification) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AlteryaIdentification) SetObject(v string)`

SetObject sets Object field to given value.


### GetAddress

`func (o *AlteryaIdentification) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AlteryaIdentification) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AlteryaIdentification) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChain

`func (o *AlteryaIdentification) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *AlteryaIdentification) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *AlteryaIdentification) SetChain(v string)`

SetChain sets Chain field to given value.


### GetRiskLevel

`func (o *AlteryaIdentification) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *AlteryaIdentification) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *AlteryaIdentification) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.


### GetReason

`func (o *AlteryaIdentification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlteryaIdentification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlteryaIdentification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetLabels

`func (o *AlteryaIdentification) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlteryaIdentification) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlteryaIdentification) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetOnChainActivity

`func (o *AlteryaIdentification) GetOnChainActivity() OnChainActivity`

GetOnChainActivity returns the OnChainActivity field if non-nil, zero value otherwise.

### GetOnChainActivityOk

`func (o *AlteryaIdentification) GetOnChainActivityOk() (*OnChainActivity, bool)`

GetOnChainActivityOk returns a tuple with the OnChainActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnChainActivity

`func (o *AlteryaIdentification) SetOnChainActivity(v OnChainActivity)`

SetOnChainActivity sets OnChainActivity field to given value.


### GetRiskScore

`func (o *AlteryaIdentification) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AlteryaIdentification) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AlteryaIdentification) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


