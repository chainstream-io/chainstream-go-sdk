# WalletPnlDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Record identifier | 
**Chain** | **string** | Blockchain network | 
**WalletAddress** | **string** | Wallet address | 
**TokenAddress** | **string** | Token contract address | 
**TokenPriceInUsd** | **string** | Current token price in USD | 
**OpenTime** | **int64** | Position opening time | 
**CloseTime** | **int64** | Position closing time | 
**LastTime** | **int64** | Last update time | 
**Balance** | **string** | Current token balance | 
**BuyAmount** | **string** | Total amount of tokens bought | 
**BuyAmountInUsd** | **string** | Total value of tokens bought in USD | 
**Buys** | **string** | Number of buy transactions | 
**SellAmount** | **string** | Total amount of tokens sold | 
**SellAmountInUsd** | **string** | Total value of tokens sold in USD | 
**Sells** | **string** | Number of sell transactions | 
**AverageBuyPriceInUsd** | **string** | Average buying price | 
**AverageSellPriceInUsd** | **string** | Average selling price | 
**UnrealizedProfitInUsd** | **string** | Unrealized profit in USD | 
**UnrealizedProfitRatio** | **string** | Unrealized profit ratio | 
**RealizedProfitInUsd** | **string** | Realized profit in USD | 
**RealizedProfitRatio** | **string** | Realized profit ratio | 
**TotalRealizedProfitInUsd** | **string** | Total realized profit in USD | 
**TotalRealizedProfitRatio** | **string** | Total realized profit ratio | 

## Methods

### NewWalletPnlDTO

`func NewWalletPnlDTO(id int64, chain string, walletAddress string, tokenAddress string, tokenPriceInUsd string, openTime int64, closeTime int64, lastTime int64, balance string, buyAmount string, buyAmountInUsd string, buys string, sellAmount string, sellAmountInUsd string, sells string, averageBuyPriceInUsd string, averageSellPriceInUsd string, unrealizedProfitInUsd string, unrealizedProfitRatio string, realizedProfitInUsd string, realizedProfitRatio string, totalRealizedProfitInUsd string, totalRealizedProfitRatio string, ) *WalletPnlDTO`

NewWalletPnlDTO instantiates a new WalletPnlDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletPnlDTOWithDefaults

`func NewWalletPnlDTOWithDefaults() *WalletPnlDTO`

NewWalletPnlDTOWithDefaults instantiates a new WalletPnlDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletPnlDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletPnlDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletPnlDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetChain

`func (o *WalletPnlDTO) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *WalletPnlDTO) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *WalletPnlDTO) SetChain(v string)`

SetChain sets Chain field to given value.


### GetWalletAddress

`func (o *WalletPnlDTO) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *WalletPnlDTO) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *WalletPnlDTO) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetTokenAddress

`func (o *WalletPnlDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *WalletPnlDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *WalletPnlDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetTokenPriceInUsd

`func (o *WalletPnlDTO) GetTokenPriceInUsd() string`

GetTokenPriceInUsd returns the TokenPriceInUsd field if non-nil, zero value otherwise.

### GetTokenPriceInUsdOk

`func (o *WalletPnlDTO) GetTokenPriceInUsdOk() (*string, bool)`

GetTokenPriceInUsdOk returns a tuple with the TokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPriceInUsd

`func (o *WalletPnlDTO) SetTokenPriceInUsd(v string)`

SetTokenPriceInUsd sets TokenPriceInUsd field to given value.


### GetOpenTime

`func (o *WalletPnlDTO) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *WalletPnlDTO) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *WalletPnlDTO) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.


### GetCloseTime

`func (o *WalletPnlDTO) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *WalletPnlDTO) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *WalletPnlDTO) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.


### GetLastTime

`func (o *WalletPnlDTO) GetLastTime() int64`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *WalletPnlDTO) GetLastTimeOk() (*int64, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *WalletPnlDTO) SetLastTime(v int64)`

SetLastTime sets LastTime field to given value.


### GetBalance

`func (o *WalletPnlDTO) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletPnlDTO) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletPnlDTO) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetBuyAmount

`func (o *WalletPnlDTO) GetBuyAmount() string`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *WalletPnlDTO) GetBuyAmountOk() (*string, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *WalletPnlDTO) SetBuyAmount(v string)`

SetBuyAmount sets BuyAmount field to given value.


### GetBuyAmountInUsd

`func (o *WalletPnlDTO) GetBuyAmountInUsd() string`

GetBuyAmountInUsd returns the BuyAmountInUsd field if non-nil, zero value otherwise.

### GetBuyAmountInUsdOk

`func (o *WalletPnlDTO) GetBuyAmountInUsdOk() (*string, bool)`

GetBuyAmountInUsdOk returns a tuple with the BuyAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInUsd

`func (o *WalletPnlDTO) SetBuyAmountInUsd(v string)`

SetBuyAmountInUsd sets BuyAmountInUsd field to given value.


### GetBuys

`func (o *WalletPnlDTO) GetBuys() string`

GetBuys returns the Buys field if non-nil, zero value otherwise.

### GetBuysOk

`func (o *WalletPnlDTO) GetBuysOk() (*string, bool)`

GetBuysOk returns a tuple with the Buys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuys

`func (o *WalletPnlDTO) SetBuys(v string)`

SetBuys sets Buys field to given value.


### GetSellAmount

`func (o *WalletPnlDTO) GetSellAmount() string`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *WalletPnlDTO) GetSellAmountOk() (*string, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *WalletPnlDTO) SetSellAmount(v string)`

SetSellAmount sets SellAmount field to given value.


### GetSellAmountInUsd

`func (o *WalletPnlDTO) GetSellAmountInUsd() string`

GetSellAmountInUsd returns the SellAmountInUsd field if non-nil, zero value otherwise.

### GetSellAmountInUsdOk

`func (o *WalletPnlDTO) GetSellAmountInUsdOk() (*string, bool)`

GetSellAmountInUsdOk returns a tuple with the SellAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInUsd

`func (o *WalletPnlDTO) SetSellAmountInUsd(v string)`

SetSellAmountInUsd sets SellAmountInUsd field to given value.


### GetSells

`func (o *WalletPnlDTO) GetSells() string`

GetSells returns the Sells field if non-nil, zero value otherwise.

### GetSellsOk

`func (o *WalletPnlDTO) GetSellsOk() (*string, bool)`

GetSellsOk returns a tuple with the Sells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSells

`func (o *WalletPnlDTO) SetSells(v string)`

SetSells sets Sells field to given value.


### GetAverageBuyPriceInUsd

`func (o *WalletPnlDTO) GetAverageBuyPriceInUsd() string`

GetAverageBuyPriceInUsd returns the AverageBuyPriceInUsd field if non-nil, zero value otherwise.

### GetAverageBuyPriceInUsdOk

`func (o *WalletPnlDTO) GetAverageBuyPriceInUsdOk() (*string, bool)`

GetAverageBuyPriceInUsdOk returns a tuple with the AverageBuyPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyPriceInUsd

`func (o *WalletPnlDTO) SetAverageBuyPriceInUsd(v string)`

SetAverageBuyPriceInUsd sets AverageBuyPriceInUsd field to given value.


### GetAverageSellPriceInUsd

`func (o *WalletPnlDTO) GetAverageSellPriceInUsd() string`

GetAverageSellPriceInUsd returns the AverageSellPriceInUsd field if non-nil, zero value otherwise.

### GetAverageSellPriceInUsdOk

`func (o *WalletPnlDTO) GetAverageSellPriceInUsdOk() (*string, bool)`

GetAverageSellPriceInUsdOk returns a tuple with the AverageSellPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSellPriceInUsd

`func (o *WalletPnlDTO) SetAverageSellPriceInUsd(v string)`

SetAverageSellPriceInUsd sets AverageSellPriceInUsd field to given value.


### GetUnrealizedProfitInUsd

`func (o *WalletPnlDTO) GetUnrealizedProfitInUsd() string`

GetUnrealizedProfitInUsd returns the UnrealizedProfitInUsd field if non-nil, zero value otherwise.

### GetUnrealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetUnrealizedProfitInUsdOk() (*string, bool)`

GetUnrealizedProfitInUsdOk returns a tuple with the UnrealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitInUsd

`func (o *WalletPnlDTO) SetUnrealizedProfitInUsd(v string)`

SetUnrealizedProfitInUsd sets UnrealizedProfitInUsd field to given value.


### GetUnrealizedProfitRatio

`func (o *WalletPnlDTO) GetUnrealizedProfitRatio() string`

GetUnrealizedProfitRatio returns the UnrealizedProfitRatio field if non-nil, zero value otherwise.

### GetUnrealizedProfitRatioOk

`func (o *WalletPnlDTO) GetUnrealizedProfitRatioOk() (*string, bool)`

GetUnrealizedProfitRatioOk returns a tuple with the UnrealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitRatio

`func (o *WalletPnlDTO) SetUnrealizedProfitRatio(v string)`

SetUnrealizedProfitRatio sets UnrealizedProfitRatio field to given value.


### GetRealizedProfitInUsd

`func (o *WalletPnlDTO) GetRealizedProfitInUsd() string`

GetRealizedProfitInUsd returns the RealizedProfitInUsd field if non-nil, zero value otherwise.

### GetRealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetRealizedProfitInUsdOk() (*string, bool)`

GetRealizedProfitInUsdOk returns a tuple with the RealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitInUsd

`func (o *WalletPnlDTO) SetRealizedProfitInUsd(v string)`

SetRealizedProfitInUsd sets RealizedProfitInUsd field to given value.


### GetRealizedProfitRatio

`func (o *WalletPnlDTO) GetRealizedProfitRatio() string`

GetRealizedProfitRatio returns the RealizedProfitRatio field if non-nil, zero value otherwise.

### GetRealizedProfitRatioOk

`func (o *WalletPnlDTO) GetRealizedProfitRatioOk() (*string, bool)`

GetRealizedProfitRatioOk returns a tuple with the RealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitRatio

`func (o *WalletPnlDTO) SetRealizedProfitRatio(v string)`

SetRealizedProfitRatio sets RealizedProfitRatio field to given value.


### GetTotalRealizedProfitInUsd

`func (o *WalletPnlDTO) GetTotalRealizedProfitInUsd() string`

GetTotalRealizedProfitInUsd returns the TotalRealizedProfitInUsd field if non-nil, zero value otherwise.

### GetTotalRealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetTotalRealizedProfitInUsdOk() (*string, bool)`

GetTotalRealizedProfitInUsdOk returns a tuple with the TotalRealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitInUsd

`func (o *WalletPnlDTO) SetTotalRealizedProfitInUsd(v string)`

SetTotalRealizedProfitInUsd sets TotalRealizedProfitInUsd field to given value.


### GetTotalRealizedProfitRatio

`func (o *WalletPnlDTO) GetTotalRealizedProfitRatio() string`

GetTotalRealizedProfitRatio returns the TotalRealizedProfitRatio field if non-nil, zero value otherwise.

### GetTotalRealizedProfitRatioOk

`func (o *WalletPnlDTO) GetTotalRealizedProfitRatioOk() (*string, bool)`

GetTotalRealizedProfitRatioOk returns a tuple with the TotalRealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitRatio

`func (o *WalletPnlDTO) SetTotalRealizedProfitRatio(v string)`

SetTotalRealizedProfitRatio sets TotalRealizedProfitRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


