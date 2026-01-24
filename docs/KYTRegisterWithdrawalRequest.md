# KYTRegisterWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.NETWORK | 
**Asset** | **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ASSET | 
**Address** | **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ADDRESS | 
**AssetAmount** | **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ASSET_AMOUNT | 
**AttemptTimestamp** | **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ATTEMPT_TIMESTAMP | 
**AssetPrice** | Pointer to **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ASSET_PRICE | [optional] 
**AssetDenomination** | Pointer to **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ASSET_DENOMINATION | [optional] 
**AssetId** | Pointer to **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.ASSET_ID | [optional] 
**Memo** | Pointer to **string** | DTO.KYT.KYT_REGISTER_WITHDRAWAL_REQUEST.MEMO | [optional] 

## Methods

### NewKYTRegisterWithdrawalRequest

`func NewKYTRegisterWithdrawalRequest(network string, asset string, address string, assetAmount string, attemptTimestamp string, ) *KYTRegisterWithdrawalRequest`

NewKYTRegisterWithdrawalRequest instantiates a new KYTRegisterWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKYTRegisterWithdrawalRequestWithDefaults

`func NewKYTRegisterWithdrawalRequestWithDefaults() *KYTRegisterWithdrawalRequest`

NewKYTRegisterWithdrawalRequestWithDefaults instantiates a new KYTRegisterWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *KYTRegisterWithdrawalRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KYTRegisterWithdrawalRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KYTRegisterWithdrawalRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAsset

`func (o *KYTRegisterWithdrawalRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *KYTRegisterWithdrawalRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *KYTRegisterWithdrawalRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAddress

`func (o *KYTRegisterWithdrawalRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KYTRegisterWithdrawalRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KYTRegisterWithdrawalRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAssetAmount

`func (o *KYTRegisterWithdrawalRequest) GetAssetAmount() string`

GetAssetAmount returns the AssetAmount field if non-nil, zero value otherwise.

### GetAssetAmountOk

`func (o *KYTRegisterWithdrawalRequest) GetAssetAmountOk() (*string, bool)`

GetAssetAmountOk returns a tuple with the AssetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAmount

`func (o *KYTRegisterWithdrawalRequest) SetAssetAmount(v string)`

SetAssetAmount sets AssetAmount field to given value.


### GetAttemptTimestamp

`func (o *KYTRegisterWithdrawalRequest) GetAttemptTimestamp() string`

GetAttemptTimestamp returns the AttemptTimestamp field if non-nil, zero value otherwise.

### GetAttemptTimestampOk

`func (o *KYTRegisterWithdrawalRequest) GetAttemptTimestampOk() (*string, bool)`

GetAttemptTimestampOk returns a tuple with the AttemptTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptTimestamp

`func (o *KYTRegisterWithdrawalRequest) SetAttemptTimestamp(v string)`

SetAttemptTimestamp sets AttemptTimestamp field to given value.


### GetAssetPrice

`func (o *KYTRegisterWithdrawalRequest) GetAssetPrice() string`

GetAssetPrice returns the AssetPrice field if non-nil, zero value otherwise.

### GetAssetPriceOk

`func (o *KYTRegisterWithdrawalRequest) GetAssetPriceOk() (*string, bool)`

GetAssetPriceOk returns a tuple with the AssetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetPrice

`func (o *KYTRegisterWithdrawalRequest) SetAssetPrice(v string)`

SetAssetPrice sets AssetPrice field to given value.

### HasAssetPrice

`func (o *KYTRegisterWithdrawalRequest) HasAssetPrice() bool`

HasAssetPrice returns a boolean if a field has been set.

### GetAssetDenomination

`func (o *KYTRegisterWithdrawalRequest) GetAssetDenomination() string`

GetAssetDenomination returns the AssetDenomination field if non-nil, zero value otherwise.

### GetAssetDenominationOk

`func (o *KYTRegisterWithdrawalRequest) GetAssetDenominationOk() (*string, bool)`

GetAssetDenominationOk returns a tuple with the AssetDenomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDenomination

`func (o *KYTRegisterWithdrawalRequest) SetAssetDenomination(v string)`

SetAssetDenomination sets AssetDenomination field to given value.

### HasAssetDenomination

`func (o *KYTRegisterWithdrawalRequest) HasAssetDenomination() bool`

HasAssetDenomination returns a boolean if a field has been set.

### GetAssetId

`func (o *KYTRegisterWithdrawalRequest) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *KYTRegisterWithdrawalRequest) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *KYTRegisterWithdrawalRequest) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *KYTRegisterWithdrawalRequest) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetMemo

`func (o *KYTRegisterWithdrawalRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *KYTRegisterWithdrawalRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *KYTRegisterWithdrawalRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *KYTRegisterWithdrawalRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


