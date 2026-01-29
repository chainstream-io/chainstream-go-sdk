# WithdrawalBaseResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | **string** | Asset type | 
**AssetId** | **string** | Asset ID | 
**Network** | **string** | Blockchain network | 
**Address** | **string** | Address | 
**Memo** | **string** | Memo information | 
**AttemptIdentifier** | **string** | Attempt identifier | 
**AssetAmount** | **string** | Asset amount | 
**ExternalId** | **string** | External ID (UUID) | 
**UsdAmount** | **string** | USD amount | 
**UpdatedAt** | **string** | Updated timestamp | 

## Methods

### NewWithdrawalBaseResponseDTO

`func NewWithdrawalBaseResponseDTO(asset string, assetId string, network string, address string, memo string, attemptIdentifier string, assetAmount string, externalId string, usdAmount string, updatedAt string, ) *WithdrawalBaseResponseDTO`

NewWithdrawalBaseResponseDTO instantiates a new WithdrawalBaseResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalBaseResponseDTOWithDefaults

`func NewWithdrawalBaseResponseDTOWithDefaults() *WithdrawalBaseResponseDTO`

NewWithdrawalBaseResponseDTOWithDefaults instantiates a new WithdrawalBaseResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *WithdrawalBaseResponseDTO) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *WithdrawalBaseResponseDTO) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *WithdrawalBaseResponseDTO) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAssetId

`func (o *WithdrawalBaseResponseDTO) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *WithdrawalBaseResponseDTO) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *WithdrawalBaseResponseDTO) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNetwork

`func (o *WithdrawalBaseResponseDTO) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WithdrawalBaseResponseDTO) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WithdrawalBaseResponseDTO) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *WithdrawalBaseResponseDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WithdrawalBaseResponseDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WithdrawalBaseResponseDTO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMemo

`func (o *WithdrawalBaseResponseDTO) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *WithdrawalBaseResponseDTO) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *WithdrawalBaseResponseDTO) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetAttemptIdentifier

`func (o *WithdrawalBaseResponseDTO) GetAttemptIdentifier() string`

GetAttemptIdentifier returns the AttemptIdentifier field if non-nil, zero value otherwise.

### GetAttemptIdentifierOk

`func (o *WithdrawalBaseResponseDTO) GetAttemptIdentifierOk() (*string, bool)`

GetAttemptIdentifierOk returns a tuple with the AttemptIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptIdentifier

`func (o *WithdrawalBaseResponseDTO) SetAttemptIdentifier(v string)`

SetAttemptIdentifier sets AttemptIdentifier field to given value.


### GetAssetAmount

`func (o *WithdrawalBaseResponseDTO) GetAssetAmount() string`

GetAssetAmount returns the AssetAmount field if non-nil, zero value otherwise.

### GetAssetAmountOk

`func (o *WithdrawalBaseResponseDTO) GetAssetAmountOk() (*string, bool)`

GetAssetAmountOk returns a tuple with the AssetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAmount

`func (o *WithdrawalBaseResponseDTO) SetAssetAmount(v string)`

SetAssetAmount sets AssetAmount field to given value.


### GetExternalId

`func (o *WithdrawalBaseResponseDTO) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *WithdrawalBaseResponseDTO) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *WithdrawalBaseResponseDTO) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetUsdAmount

`func (o *WithdrawalBaseResponseDTO) GetUsdAmount() string`

GetUsdAmount returns the UsdAmount field if non-nil, zero value otherwise.

### GetUsdAmountOk

`func (o *WithdrawalBaseResponseDTO) GetUsdAmountOk() (*string, bool)`

GetUsdAmountOk returns a tuple with the UsdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdAmount

`func (o *WithdrawalBaseResponseDTO) SetUsdAmount(v string)`

SetUsdAmount sets UsdAmount field to given value.


### GetUpdatedAt

`func (o *WithdrawalBaseResponseDTO) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WithdrawalBaseResponseDTO) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WithdrawalBaseResponseDTO) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


