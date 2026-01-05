# TransferBaseResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.UPDATED_AT | 
**Asset** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.ASSET | 
**AssetId** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.ASSET_ID | 
**Network** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.NETWORK | 
**TransferReference** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.TRANSFER_REFERENCE | 
**Memo** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.MEMO | 
**Tx** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.TX | 
**Idx** | **float32** | DTO.KYT.TRANSFER_BASE_RESPONSE.IDX | 
**UsdAmount** | **float32** | DTO.KYT.TRANSFER_BASE_RESPONSE.USD_AMOUNT | 
**AssetAmount** | **float32** | DTO.KYT.TRANSFER_BASE_RESPONSE.ASSET_AMOUNT | 
**Timestamp** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.TIMESTAMP | 
**OutputAddress** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.OUTPUT_ADDRESS | 
**ExternalId** | **string** | DTO.KYT.TRANSFER_BASE_RESPONSE.EXTERNAL_ID | 

## Methods

### NewTransferBaseResponseDTO

`func NewTransferBaseResponseDTO(updatedAt string, asset string, assetId string, network string, transferReference string, memo string, tx string, idx float32, usdAmount float32, assetAmount float32, timestamp string, outputAddress string, externalId string, ) *TransferBaseResponseDTO`

NewTransferBaseResponseDTO instantiates a new TransferBaseResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferBaseResponseDTOWithDefaults

`func NewTransferBaseResponseDTOWithDefaults() *TransferBaseResponseDTO`

NewTransferBaseResponseDTOWithDefaults instantiates a new TransferBaseResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *TransferBaseResponseDTO) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransferBaseResponseDTO) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransferBaseResponseDTO) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAsset

`func (o *TransferBaseResponseDTO) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TransferBaseResponseDTO) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TransferBaseResponseDTO) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAssetId

`func (o *TransferBaseResponseDTO) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransferBaseResponseDTO) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransferBaseResponseDTO) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNetwork

`func (o *TransferBaseResponseDTO) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransferBaseResponseDTO) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransferBaseResponseDTO) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransferReference

`func (o *TransferBaseResponseDTO) GetTransferReference() string`

GetTransferReference returns the TransferReference field if non-nil, zero value otherwise.

### GetTransferReferenceOk

`func (o *TransferBaseResponseDTO) GetTransferReferenceOk() (*string, bool)`

GetTransferReferenceOk returns a tuple with the TransferReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferReference

`func (o *TransferBaseResponseDTO) SetTransferReference(v string)`

SetTransferReference sets TransferReference field to given value.


### GetMemo

`func (o *TransferBaseResponseDTO) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransferBaseResponseDTO) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransferBaseResponseDTO) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetTx

`func (o *TransferBaseResponseDTO) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *TransferBaseResponseDTO) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *TransferBaseResponseDTO) SetTx(v string)`

SetTx sets Tx field to given value.


### GetIdx

`func (o *TransferBaseResponseDTO) GetIdx() float32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *TransferBaseResponseDTO) GetIdxOk() (*float32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *TransferBaseResponseDTO) SetIdx(v float32)`

SetIdx sets Idx field to given value.


### GetUsdAmount

`func (o *TransferBaseResponseDTO) GetUsdAmount() float32`

GetUsdAmount returns the UsdAmount field if non-nil, zero value otherwise.

### GetUsdAmountOk

`func (o *TransferBaseResponseDTO) GetUsdAmountOk() (*float32, bool)`

GetUsdAmountOk returns a tuple with the UsdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdAmount

`func (o *TransferBaseResponseDTO) SetUsdAmount(v float32)`

SetUsdAmount sets UsdAmount field to given value.


### GetAssetAmount

`func (o *TransferBaseResponseDTO) GetAssetAmount() float32`

GetAssetAmount returns the AssetAmount field if non-nil, zero value otherwise.

### GetAssetAmountOk

`func (o *TransferBaseResponseDTO) GetAssetAmountOk() (*float32, bool)`

GetAssetAmountOk returns a tuple with the AssetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetAmount

`func (o *TransferBaseResponseDTO) SetAssetAmount(v float32)`

SetAssetAmount sets AssetAmount field to given value.


### GetTimestamp

`func (o *TransferBaseResponseDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransferBaseResponseDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransferBaseResponseDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetOutputAddress

`func (o *TransferBaseResponseDTO) GetOutputAddress() string`

GetOutputAddress returns the OutputAddress field if non-nil, zero value otherwise.

### GetOutputAddressOk

`func (o *TransferBaseResponseDTO) GetOutputAddressOk() (*string, bool)`

GetOutputAddressOk returns a tuple with the OutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputAddress

`func (o *TransferBaseResponseDTO) SetOutputAddress(v string)`

SetOutputAddress sets OutputAddress field to given value.


### GetExternalId

`func (o *TransferBaseResponseDTO) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TransferBaseResponseDTO) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TransferBaseResponseDTO) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


