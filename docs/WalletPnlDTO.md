# WalletPnlDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | DTO.WALLET.PNL.ID | 
**Chain** | **string** | DTO.WALLET.PNL.CHAIN | 
**WalletAddress** | **string** | DTO.WALLET.PNL.WALLET_ADDRESS | 
**TokenAddress** | **string** | DTO.WALLET.PNL.TOKEN_ADDRESS | 
**TokenPriceInUsd** | **float32** | DTO.WALLET.PNL.TOKEN_PRICE | 
**OpenTime** | **float32** | DTO.WALLET.PNL.OPEN_TIME | 
**CloseTime** | **float32** | DTO.WALLET.PNL.CLOSE_TIME | 
**LastTime** | **float32** | DTO.WALLET.PNL.LAST_TIME | 
**Balance** | **float32** | DTO.WALLET.PNL.BALANCE | 
**BuyAmount** | **float32** | DTO.WALLET.PNL.BUY_AMOUNT | 
**BuyAmountInUsd** | **float32** | DTO.WALLET.PNL.BUY_AMOUNT_USD | 
**Buys** | **float32** | DTO.WALLET.PNL.BUYS | 
**SellAmount** | **float32** | DTO.WALLET.PNL.SELL_AMOUNT | 
**SellAmountInUsd** | **float32** | DTO.WALLET.PNL.SELL_AMOUNT_USD | 
**Sells** | **float32** | DTO.WALLET.PNL.SELLS | 
**AverageBuyPriceInUsd** | **float32** | DTO.WALLET.PNL.AVERAGE_BUY_PRICE | 
**AverageSellPriceInUsd** | **float32** | DTO.WALLET.PNL.AVERAGE_SELL_PRICE | 
**UnrealizedProfitInUsd** | **float32** | DTO.WALLET.PNL.UNREALIZED_PROFIT | 
**UnrealizedProfitRatio** | **float32** | DTO.WALLET.PNL.UNREALIZED_PROFIT_RATIO | 
**RealizedProfitInUsd** | **float32** | DTO.WALLET.PNL.REALIZED_PROFIT | 
**RealizedProfitRatio** | **float32** | DTO.WALLET.PNL.REALIZED_PROFIT_RATIO | 
**TotalRealizedProfitInUsd** | **float32** | DTO.WALLET.PNL.TOTAL_REALIZED_PROFIT | 
**TotalRealizedProfitRatio** | **float32** | DTO.WALLET.PNL.TOTAL_REALIZED_PROFIT_RATIO | 

## Methods

### NewWalletPnlDTO

`func NewWalletPnlDTO(id float32, chain string, walletAddress string, tokenAddress string, tokenPriceInUsd float32, openTime float32, closeTime float32, lastTime float32, balance float32, buyAmount float32, buyAmountInUsd float32, buys float32, sellAmount float32, sellAmountInUsd float32, sells float32, averageBuyPriceInUsd float32, averageSellPriceInUsd float32, unrealizedProfitInUsd float32, unrealizedProfitRatio float32, realizedProfitInUsd float32, realizedProfitRatio float32, totalRealizedProfitInUsd float32, totalRealizedProfitRatio float32, ) *WalletPnlDTO`

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

`func (o *WalletPnlDTO) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletPnlDTO) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletPnlDTO) SetId(v float32)`

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

`func (o *WalletPnlDTO) GetTokenPriceInUsd() float32`

GetTokenPriceInUsd returns the TokenPriceInUsd field if non-nil, zero value otherwise.

### GetTokenPriceInUsdOk

`func (o *WalletPnlDTO) GetTokenPriceInUsdOk() (*float32, bool)`

GetTokenPriceInUsdOk returns a tuple with the TokenPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPriceInUsd

`func (o *WalletPnlDTO) SetTokenPriceInUsd(v float32)`

SetTokenPriceInUsd sets TokenPriceInUsd field to given value.


### GetOpenTime

`func (o *WalletPnlDTO) GetOpenTime() float32`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *WalletPnlDTO) GetOpenTimeOk() (*float32, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *WalletPnlDTO) SetOpenTime(v float32)`

SetOpenTime sets OpenTime field to given value.


### GetCloseTime

`func (o *WalletPnlDTO) GetCloseTime() float32`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *WalletPnlDTO) GetCloseTimeOk() (*float32, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *WalletPnlDTO) SetCloseTime(v float32)`

SetCloseTime sets CloseTime field to given value.


### GetLastTime

`func (o *WalletPnlDTO) GetLastTime() float32`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *WalletPnlDTO) GetLastTimeOk() (*float32, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *WalletPnlDTO) SetLastTime(v float32)`

SetLastTime sets LastTime field to given value.


### GetBalance

`func (o *WalletPnlDTO) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletPnlDTO) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletPnlDTO) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetBuyAmount

`func (o *WalletPnlDTO) GetBuyAmount() float32`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *WalletPnlDTO) GetBuyAmountOk() (*float32, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *WalletPnlDTO) SetBuyAmount(v float32)`

SetBuyAmount sets BuyAmount field to given value.


### GetBuyAmountInUsd

`func (o *WalletPnlDTO) GetBuyAmountInUsd() float32`

GetBuyAmountInUsd returns the BuyAmountInUsd field if non-nil, zero value otherwise.

### GetBuyAmountInUsdOk

`func (o *WalletPnlDTO) GetBuyAmountInUsdOk() (*float32, bool)`

GetBuyAmountInUsdOk returns a tuple with the BuyAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInUsd

`func (o *WalletPnlDTO) SetBuyAmountInUsd(v float32)`

SetBuyAmountInUsd sets BuyAmountInUsd field to given value.


### GetBuys

`func (o *WalletPnlDTO) GetBuys() float32`

GetBuys returns the Buys field if non-nil, zero value otherwise.

### GetBuysOk

`func (o *WalletPnlDTO) GetBuysOk() (*float32, bool)`

GetBuysOk returns a tuple with the Buys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuys

`func (o *WalletPnlDTO) SetBuys(v float32)`

SetBuys sets Buys field to given value.


### GetSellAmount

`func (o *WalletPnlDTO) GetSellAmount() float32`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *WalletPnlDTO) GetSellAmountOk() (*float32, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *WalletPnlDTO) SetSellAmount(v float32)`

SetSellAmount sets SellAmount field to given value.


### GetSellAmountInUsd

`func (o *WalletPnlDTO) GetSellAmountInUsd() float32`

GetSellAmountInUsd returns the SellAmountInUsd field if non-nil, zero value otherwise.

### GetSellAmountInUsdOk

`func (o *WalletPnlDTO) GetSellAmountInUsdOk() (*float32, bool)`

GetSellAmountInUsdOk returns a tuple with the SellAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInUsd

`func (o *WalletPnlDTO) SetSellAmountInUsd(v float32)`

SetSellAmountInUsd sets SellAmountInUsd field to given value.


### GetSells

`func (o *WalletPnlDTO) GetSells() float32`

GetSells returns the Sells field if non-nil, zero value otherwise.

### GetSellsOk

`func (o *WalletPnlDTO) GetSellsOk() (*float32, bool)`

GetSellsOk returns a tuple with the Sells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSells

`func (o *WalletPnlDTO) SetSells(v float32)`

SetSells sets Sells field to given value.


### GetAverageBuyPriceInUsd

`func (o *WalletPnlDTO) GetAverageBuyPriceInUsd() float32`

GetAverageBuyPriceInUsd returns the AverageBuyPriceInUsd field if non-nil, zero value otherwise.

### GetAverageBuyPriceInUsdOk

`func (o *WalletPnlDTO) GetAverageBuyPriceInUsdOk() (*float32, bool)`

GetAverageBuyPriceInUsdOk returns a tuple with the AverageBuyPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyPriceInUsd

`func (o *WalletPnlDTO) SetAverageBuyPriceInUsd(v float32)`

SetAverageBuyPriceInUsd sets AverageBuyPriceInUsd field to given value.


### GetAverageSellPriceInUsd

`func (o *WalletPnlDTO) GetAverageSellPriceInUsd() float32`

GetAverageSellPriceInUsd returns the AverageSellPriceInUsd field if non-nil, zero value otherwise.

### GetAverageSellPriceInUsdOk

`func (o *WalletPnlDTO) GetAverageSellPriceInUsdOk() (*float32, bool)`

GetAverageSellPriceInUsdOk returns a tuple with the AverageSellPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSellPriceInUsd

`func (o *WalletPnlDTO) SetAverageSellPriceInUsd(v float32)`

SetAverageSellPriceInUsd sets AverageSellPriceInUsd field to given value.


### GetUnrealizedProfitInUsd

`func (o *WalletPnlDTO) GetUnrealizedProfitInUsd() float32`

GetUnrealizedProfitInUsd returns the UnrealizedProfitInUsd field if non-nil, zero value otherwise.

### GetUnrealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetUnrealizedProfitInUsdOk() (*float32, bool)`

GetUnrealizedProfitInUsdOk returns a tuple with the UnrealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitInUsd

`func (o *WalletPnlDTO) SetUnrealizedProfitInUsd(v float32)`

SetUnrealizedProfitInUsd sets UnrealizedProfitInUsd field to given value.


### GetUnrealizedProfitRatio

`func (o *WalletPnlDTO) GetUnrealizedProfitRatio() float32`

GetUnrealizedProfitRatio returns the UnrealizedProfitRatio field if non-nil, zero value otherwise.

### GetUnrealizedProfitRatioOk

`func (o *WalletPnlDTO) GetUnrealizedProfitRatioOk() (*float32, bool)`

GetUnrealizedProfitRatioOk returns a tuple with the UnrealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitRatio

`func (o *WalletPnlDTO) SetUnrealizedProfitRatio(v float32)`

SetUnrealizedProfitRatio sets UnrealizedProfitRatio field to given value.


### GetRealizedProfitInUsd

`func (o *WalletPnlDTO) GetRealizedProfitInUsd() float32`

GetRealizedProfitInUsd returns the RealizedProfitInUsd field if non-nil, zero value otherwise.

### GetRealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetRealizedProfitInUsdOk() (*float32, bool)`

GetRealizedProfitInUsdOk returns a tuple with the RealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitInUsd

`func (o *WalletPnlDTO) SetRealizedProfitInUsd(v float32)`

SetRealizedProfitInUsd sets RealizedProfitInUsd field to given value.


### GetRealizedProfitRatio

`func (o *WalletPnlDTO) GetRealizedProfitRatio() float32`

GetRealizedProfitRatio returns the RealizedProfitRatio field if non-nil, zero value otherwise.

### GetRealizedProfitRatioOk

`func (o *WalletPnlDTO) GetRealizedProfitRatioOk() (*float32, bool)`

GetRealizedProfitRatioOk returns a tuple with the RealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitRatio

`func (o *WalletPnlDTO) SetRealizedProfitRatio(v float32)`

SetRealizedProfitRatio sets RealizedProfitRatio field to given value.


### GetTotalRealizedProfitInUsd

`func (o *WalletPnlDTO) GetTotalRealizedProfitInUsd() float32`

GetTotalRealizedProfitInUsd returns the TotalRealizedProfitInUsd field if non-nil, zero value otherwise.

### GetTotalRealizedProfitInUsdOk

`func (o *WalletPnlDTO) GetTotalRealizedProfitInUsdOk() (*float32, bool)`

GetTotalRealizedProfitInUsdOk returns a tuple with the TotalRealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitInUsd

`func (o *WalletPnlDTO) SetTotalRealizedProfitInUsd(v float32)`

SetTotalRealizedProfitInUsd sets TotalRealizedProfitInUsd field to given value.


### GetTotalRealizedProfitRatio

`func (o *WalletPnlDTO) GetTotalRealizedProfitRatio() float32`

GetTotalRealizedProfitRatio returns the TotalRealizedProfitRatio field if non-nil, zero value otherwise.

### GetTotalRealizedProfitRatioOk

`func (o *WalletPnlDTO) GetTotalRealizedProfitRatioOk() (*float32, bool)`

GetTotalRealizedProfitRatioOk returns a tuple with the TotalRealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitRatio

`func (o *WalletPnlDTO) SetTotalRealizedProfitRatio(v float32)`

SetTotalRealizedProfitRatio sets TotalRealizedProfitRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


