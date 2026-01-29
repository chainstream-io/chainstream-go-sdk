# TopTradersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | Token address for top traders query | 
**WalletAddress** | **string** | Wallet address of the trader | 
**TradeCount** | **int64** | Total number of trades | 
**TradeAmount** | **string** | Total trade amount | 
**BuyCount** | **int64** | Number of buy trades | 
**BuyAmount** | **string** | Total buy amount | 
**BuyAmountInUsd** | **string** | Total buy amount in USD | 
**BuyAmountInNative** | **string** | Total buy amount in native token | 
**SellCount** | **int64** | Number of sell trades | 
**SellAmount** | **string** | Total sell amount | 
**SellAmountInUsd** | **string** | Total sell amount in USD | 
**SellAmountInNative** | **string** | Total sell amount in native token | 

## Methods

### NewTopTradersDTO

`func NewTopTradersDTO(tokenAddress string, walletAddress string, tradeCount int64, tradeAmount string, buyCount int64, buyAmount string, buyAmountInUsd string, buyAmountInNative string, sellCount int64, sellAmount string, sellAmountInUsd string, sellAmountInNative string, ) *TopTradersDTO`

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


### GetWalletAddress

`func (o *TopTradersDTO) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *TopTradersDTO) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *TopTradersDTO) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetTradeCount

`func (o *TopTradersDTO) GetTradeCount() int64`

GetTradeCount returns the TradeCount field if non-nil, zero value otherwise.

### GetTradeCountOk

`func (o *TopTradersDTO) GetTradeCountOk() (*int64, bool)`

GetTradeCountOk returns a tuple with the TradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCount

`func (o *TopTradersDTO) SetTradeCount(v int64)`

SetTradeCount sets TradeCount field to given value.


### GetTradeAmount

`func (o *TopTradersDTO) GetTradeAmount() string`

GetTradeAmount returns the TradeAmount field if non-nil, zero value otherwise.

### GetTradeAmountOk

`func (o *TopTradersDTO) GetTradeAmountOk() (*string, bool)`

GetTradeAmountOk returns a tuple with the TradeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmount

`func (o *TopTradersDTO) SetTradeAmount(v string)`

SetTradeAmount sets TradeAmount field to given value.


### GetBuyCount

`func (o *TopTradersDTO) GetBuyCount() int64`

GetBuyCount returns the BuyCount field if non-nil, zero value otherwise.

### GetBuyCountOk

`func (o *TopTradersDTO) GetBuyCountOk() (*int64, bool)`

GetBuyCountOk returns a tuple with the BuyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyCount

`func (o *TopTradersDTO) SetBuyCount(v int64)`

SetBuyCount sets BuyCount field to given value.


### GetBuyAmount

`func (o *TopTradersDTO) GetBuyAmount() string`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *TopTradersDTO) GetBuyAmountOk() (*string, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *TopTradersDTO) SetBuyAmount(v string)`

SetBuyAmount sets BuyAmount field to given value.


### GetBuyAmountInUsd

`func (o *TopTradersDTO) GetBuyAmountInUsd() string`

GetBuyAmountInUsd returns the BuyAmountInUsd field if non-nil, zero value otherwise.

### GetBuyAmountInUsdOk

`func (o *TopTradersDTO) GetBuyAmountInUsdOk() (*string, bool)`

GetBuyAmountInUsdOk returns a tuple with the BuyAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInUsd

`func (o *TopTradersDTO) SetBuyAmountInUsd(v string)`

SetBuyAmountInUsd sets BuyAmountInUsd field to given value.


### GetBuyAmountInNative

`func (o *TopTradersDTO) GetBuyAmountInNative() string`

GetBuyAmountInNative returns the BuyAmountInNative field if non-nil, zero value otherwise.

### GetBuyAmountInNativeOk

`func (o *TopTradersDTO) GetBuyAmountInNativeOk() (*string, bool)`

GetBuyAmountInNativeOk returns a tuple with the BuyAmountInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInNative

`func (o *TopTradersDTO) SetBuyAmountInNative(v string)`

SetBuyAmountInNative sets BuyAmountInNative field to given value.


### GetSellCount

`func (o *TopTradersDTO) GetSellCount() int64`

GetSellCount returns the SellCount field if non-nil, zero value otherwise.

### GetSellCountOk

`func (o *TopTradersDTO) GetSellCountOk() (*int64, bool)`

GetSellCountOk returns a tuple with the SellCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellCount

`func (o *TopTradersDTO) SetSellCount(v int64)`

SetSellCount sets SellCount field to given value.


### GetSellAmount

`func (o *TopTradersDTO) GetSellAmount() string`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *TopTradersDTO) GetSellAmountOk() (*string, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *TopTradersDTO) SetSellAmount(v string)`

SetSellAmount sets SellAmount field to given value.


### GetSellAmountInUsd

`func (o *TopTradersDTO) GetSellAmountInUsd() string`

GetSellAmountInUsd returns the SellAmountInUsd field if non-nil, zero value otherwise.

### GetSellAmountInUsdOk

`func (o *TopTradersDTO) GetSellAmountInUsdOk() (*string, bool)`

GetSellAmountInUsdOk returns a tuple with the SellAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInUsd

`func (o *TopTradersDTO) SetSellAmountInUsd(v string)`

SetSellAmountInUsd sets SellAmountInUsd field to given value.


### GetSellAmountInNative

`func (o *TopTradersDTO) GetSellAmountInNative() string`

GetSellAmountInNative returns the SellAmountInNative field if non-nil, zero value otherwise.

### GetSellAmountInNativeOk

`func (o *TopTradersDTO) GetSellAmountInNativeOk() (*string, bool)`

GetSellAmountInNativeOk returns a tuple with the SellAmountInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInNative

`func (o *TopTradersDTO) SetSellAmountInNative(v string)`

SetSellAmountInNative sets SellAmountInNative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


