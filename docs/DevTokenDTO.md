# DevTokenDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | DTO.DEV_TOKEN.ADDRESS | 
**Metadata** | [**TokenMetadata**](TokenMetadata.md) | DTO.DEV_TOKEN.METADATA | 
**MarketData** | [**TokenMarketData**](TokenMarketData.md) | DTO.DEV_TOKEN.MARKET_DATA | 
**Stats** | [**TokenStat**](TokenStat.md) | DTO.DEV_TOKEN.STATS | 

## Methods

### NewDevTokenDTO

`func NewDevTokenDTO(address string, metadata TokenMetadata, marketData TokenMarketData, stats TokenStat, ) *DevTokenDTO`

NewDevTokenDTO instantiates a new DevTokenDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevTokenDTOWithDefaults

`func NewDevTokenDTOWithDefaults() *DevTokenDTO`

NewDevTokenDTOWithDefaults instantiates a new DevTokenDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DevTokenDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevTokenDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevTokenDTO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMetadata

`func (o *DevTokenDTO) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DevTokenDTO) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DevTokenDTO) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.


### GetMarketData

`func (o *DevTokenDTO) GetMarketData() TokenMarketData`

GetMarketData returns the MarketData field if non-nil, zero value otherwise.

### GetMarketDataOk

`func (o *DevTokenDTO) GetMarketDataOk() (*TokenMarketData, bool)`

GetMarketDataOk returns a tuple with the MarketData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketData

`func (o *DevTokenDTO) SetMarketData(v TokenMarketData)`

SetMarketData sets MarketData field to given value.


### GetStats

`func (o *DevTokenDTO) GetStats() TokenStat`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DevTokenDTO) GetStatsOk() (*TokenStat, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DevTokenDTO) SetStats(v TokenStat)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


