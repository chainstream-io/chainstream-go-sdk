# GainersAndLosersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | GLOBAL.WALLETADDRESS.DESCRIPTION | 
**Pnl** | **float32** | DTO.TRADE.GAINERS_LOSERS.PNL | 
**TradeCount** | **float32** | DTO.TRADE.GAINERS_LOSERS.TRADE_COUNT | 
**Volume** | **float32** | DTO.TRADE.GAINERS_LOSERS.VOLUME | 

## Methods

### NewGainersAndLosersDTO

`func NewGainersAndLosersDTO(address string, pnl float32, tradeCount float32, volume float32, ) *GainersAndLosersDTO`

NewGainersAndLosersDTO instantiates a new GainersAndLosersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGainersAndLosersDTOWithDefaults

`func NewGainersAndLosersDTOWithDefaults() *GainersAndLosersDTO`

NewGainersAndLosersDTOWithDefaults instantiates a new GainersAndLosersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GainersAndLosersDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GainersAndLosersDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GainersAndLosersDTO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPnl

`func (o *GainersAndLosersDTO) GetPnl() float32`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GainersAndLosersDTO) GetPnlOk() (*float32, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GainersAndLosersDTO) SetPnl(v float32)`

SetPnl sets Pnl field to given value.


### GetTradeCount

`func (o *GainersAndLosersDTO) GetTradeCount() float32`

GetTradeCount returns the TradeCount field if non-nil, zero value otherwise.

### GetTradeCountOk

`func (o *GainersAndLosersDTO) GetTradeCountOk() (*float32, bool)`

GetTradeCountOk returns a tuple with the TradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCount

`func (o *GainersAndLosersDTO) SetTradeCount(v float32)`

SetTradeCount sets TradeCount field to given value.


### GetVolume

`func (o *GainersAndLosersDTO) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GainersAndLosersDTO) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GainersAndLosersDTO) SetVolume(v float32)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


