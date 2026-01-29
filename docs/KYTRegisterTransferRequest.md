# KYTRegisterTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Transfer network | 
**Asset** | **string** | Asset type | 
**TransferReference** | **string** | Transfer hash/transaction ID | 
**Direction** | **string** | Transfer direction | 

## Methods

### NewKYTRegisterTransferRequest

`func NewKYTRegisterTransferRequest(network string, asset string, transferReference string, direction string, ) *KYTRegisterTransferRequest`

NewKYTRegisterTransferRequest instantiates a new KYTRegisterTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKYTRegisterTransferRequestWithDefaults

`func NewKYTRegisterTransferRequestWithDefaults() *KYTRegisterTransferRequest`

NewKYTRegisterTransferRequestWithDefaults instantiates a new KYTRegisterTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *KYTRegisterTransferRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KYTRegisterTransferRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KYTRegisterTransferRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAsset

`func (o *KYTRegisterTransferRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *KYTRegisterTransferRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *KYTRegisterTransferRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetTransferReference

`func (o *KYTRegisterTransferRequest) GetTransferReference() string`

GetTransferReference returns the TransferReference field if non-nil, zero value otherwise.

### GetTransferReferenceOk

`func (o *KYTRegisterTransferRequest) GetTransferReferenceOk() (*string, bool)`

GetTransferReferenceOk returns a tuple with the TransferReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferReference

`func (o *KYTRegisterTransferRequest) SetTransferReference(v string)`

SetTransferReference sets TransferReference field to given value.


### GetDirection

`func (o *KYTRegisterTransferRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KYTRegisterTransferRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KYTRegisterTransferRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


