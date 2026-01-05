# TopTradersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | DTO.TRADE.TOP_TRADERS.TOKEN_ADDRESS | 
**Owner** | **string** | DTO.TRADE.TOP_TRADERS.OWNER | 
**Tags** | **[]string** | DTO.TRADE.TOP_TRADERS.TAGS | 
**Type** | **string** | DTO.TRADE.TOP_TRADERS.TYPE | 
**Volume** | **float32** | DTO.TRADE.TOP_TRADERS.VOLUME | 
**Trade** | **float32** | DTO.TRADE.TOP_TRADERS.TRADE | 
**TradeBuy** | **float32** | DTO.TRADE.TOP_TRADERS.TRADE_BUY | 
**TradeSell** | **float32** | DTO.TRADE.TOP_TRADERS.TRADE_SELL | 
**VolumeBuy** | **float32** | DTO.TRADE.TOP_TRADERS.VOLUME_BUY | 
**VolumeSell** | **float32** | DTO.TRADE.TOP_TRADERS.VOLUME_SELL | 
**IsScaledUiToken** | **bool** | DTO.TRADE.TOP_TRADERS.IS_SCALED_UI_TOKEN | 
**Multiplier** | Pointer to **string** | DTO.TRADE.TOP_TRADERS.MULTIPLIER | [optional] 

## Methods

### NewTopTradersDTO

`func NewTopTradersDTO(tokenAddress string, owner string, tags []string, type_ string, volume float32, trade float32, tradeBuy float32, tradeSell float32, volumeBuy float32, volumeSell float32, isScaledUiToken bool, ) *TopTradersDTO`

NewTopTradersDTO instantiates a new TopTradersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopTradersDTOWithDefaults

`func NewTopTradersDTOWithDefaults() *TopTradersDTO`

NewTopTradersDTOWithDefaults instantiates a new TopTradersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *TopTradersDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TopTradersDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TopTradersDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetOwner

`func (o *TopTradersDTO) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TopTradersDTO) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TopTradersDTO) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTags

`func (o *TopTradersDTO) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TopTradersDTO) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TopTradersDTO) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *TopTradersDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopTradersDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopTradersDTO) SetType(v string)`

SetType sets Type field to given value.


### GetVolume

`func (o *TopTradersDTO) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TopTradersDTO) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TopTradersDTO) SetVolume(v float32)`

SetVolume sets Volume field to given value.


### GetTrade

`func (o *TopTradersDTO) GetTrade() float32`

GetTrade returns the Trade field if non-nil, zero value otherwise.

### GetTradeOk

`func (o *TopTradersDTO) GetTradeOk() (*float32, bool)`

GetTradeOk returns a tuple with the Trade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrade

`func (o *TopTradersDTO) SetTrade(v float32)`

SetTrade sets Trade field to given value.


### GetTradeBuy

`func (o *TopTradersDTO) GetTradeBuy() float32`

GetTradeBuy returns the TradeBuy field if non-nil, zero value otherwise.

### GetTradeBuyOk

`func (o *TopTradersDTO) GetTradeBuyOk() (*float32, bool)`

GetTradeBuyOk returns a tuple with the TradeBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeBuy

`func (o *TopTradersDTO) SetTradeBuy(v float32)`

SetTradeBuy sets TradeBuy field to given value.


### GetTradeSell

`func (o *TopTradersDTO) GetTradeSell() float32`

GetTradeSell returns the TradeSell field if non-nil, zero value otherwise.

### GetTradeSellOk

`func (o *TopTradersDTO) GetTradeSellOk() (*float32, bool)`

GetTradeSellOk returns a tuple with the TradeSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSell

`func (o *TopTradersDTO) SetTradeSell(v float32)`

SetTradeSell sets TradeSell field to given value.


### GetVolumeBuy

`func (o *TopTradersDTO) GetVolumeBuy() float32`

GetVolumeBuy returns the VolumeBuy field if non-nil, zero value otherwise.

### GetVolumeBuyOk

`func (o *TopTradersDTO) GetVolumeBuyOk() (*float32, bool)`

GetVolumeBuyOk returns a tuple with the VolumeBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBuy

`func (o *TopTradersDTO) SetVolumeBuy(v float32)`

SetVolumeBuy sets VolumeBuy field to given value.


### GetVolumeSell

`func (o *TopTradersDTO) GetVolumeSell() float32`

GetVolumeSell returns the VolumeSell field if non-nil, zero value otherwise.

### GetVolumeSellOk

`func (o *TopTradersDTO) GetVolumeSellOk() (*float32, bool)`

GetVolumeSellOk returns a tuple with the VolumeSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSell

`func (o *TopTradersDTO) SetVolumeSell(v float32)`

SetVolumeSell sets VolumeSell field to given value.


### GetIsScaledUiToken

`func (o *TopTradersDTO) GetIsScaledUiToken() bool`

GetIsScaledUiToken returns the IsScaledUiToken field if non-nil, zero value otherwise.

### GetIsScaledUiTokenOk

`func (o *TopTradersDTO) GetIsScaledUiTokenOk() (*bool, bool)`

GetIsScaledUiTokenOk returns a tuple with the IsScaledUiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScaledUiToken

`func (o *TopTradersDTO) SetIsScaledUiToken(v bool)`

SetIsScaledUiToken sets IsScaledUiToken field to given value.


### GetMultiplier

`func (o *TopTradersDTO) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *TopTradersDTO) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *TopTradersDTO) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *TopTradersDTO) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


