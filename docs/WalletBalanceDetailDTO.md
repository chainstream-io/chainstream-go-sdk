# WalletBalanceDetailDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenAddress** | **string** | Token contract address | 
**Amount** | **string** | Token amount | 
**ValueInUsd** | **string** | Token amount in USD | 
**Name** | **string** | Token name | 
**Symbol** | **string** | Token symbol | 
**ImageUrl** | **string** | Token logo iamgeUrl | 
**PriceInSol** | **string** | Price in SOL | 
**PriceInUsd** | **string** | Price in USD | 
**PriceChangeRatioInUsd24h** | **string** | 24h price change ratio | 
**UnrealizedProfitInUsd** | **string** | Unrealized profit in USD | 
**UnrealizedProfitRatio** | **string** | Unrealized profit ratio | 
**OpenTime** | **int64** | Position opening time | 
**CloseTime** | **int64** | Position closing time | 
**Buys** | **string** | Number of buy transactions | 
**BuyAmount** | **string** | Total amount of tokens bought | 
**BuyAmountInUsd** | **string** | Total value of tokens bought in USD | 
**Sells** | **string** | Number of sell transactions | 
**SellAmount** | **string** | Total amount of tokens sold | 
**SellAmountInUsd** | **string** | Total value of tokens sold in USD | 
**AverageBuyPriceInUsd** | **string** | Average buying price | 
**AverageSellPriceInUsd** | **string** | Average selling price | 
**RealizedProfitInUsd** | **string** | Realized profit in USD | 
**RealizedProfitRatio** | **string** | Realized profit ratio | 
**TotalRealizedProfitInUsd** | **string** | Total realized profit in USD | 
**TotalRealizedProfitRatio** | **string** | Total realized profit ratio | 

## Methods

### NewWalletBalanceDetailDTO

`func NewWalletBalanceDetailDTO(tokenAddress string, amount string, valueInUsd string, name string, symbol string, imageUrl string, priceInSol string, priceInUsd string, priceChangeRatioInUsd24h string, unrealizedProfitInUsd string, unrealizedProfitRatio string, openTime int64, closeTime int64, buys string, buyAmount string, buyAmountInUsd string, sells string, sellAmount string, sellAmountInUsd string, averageBuyPriceInUsd string, averageSellPriceInUsd string, realizedProfitInUsd string, realizedProfitRatio string, totalRealizedProfitInUsd string, totalRealizedProfitRatio string, ) *WalletBalanceDetailDTO`

NewWalletBalanceDetailDTO instantiates a new WalletBalanceDetailDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceDetailDTOWithDefaults

`func NewWalletBalanceDetailDTOWithDefaults() *WalletBalanceDetailDTO`

NewWalletBalanceDetailDTOWithDefaults instantiates a new WalletBalanceDetailDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenAddress

`func (o *WalletBalanceDetailDTO) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *WalletBalanceDetailDTO) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *WalletBalanceDetailDTO) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetAmount

`func (o *WalletBalanceDetailDTO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletBalanceDetailDTO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletBalanceDetailDTO) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetValueInUsd

`func (o *WalletBalanceDetailDTO) GetValueInUsd() string`

GetValueInUsd returns the ValueInUsd field if non-nil, zero value otherwise.

### GetValueInUsdOk

`func (o *WalletBalanceDetailDTO) GetValueInUsdOk() (*string, bool)`

GetValueInUsdOk returns a tuple with the ValueInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInUsd

`func (o *WalletBalanceDetailDTO) SetValueInUsd(v string)`

SetValueInUsd sets ValueInUsd field to given value.


### GetName

`func (o *WalletBalanceDetailDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletBalanceDetailDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletBalanceDetailDTO) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *WalletBalanceDetailDTO) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *WalletBalanceDetailDTO) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *WalletBalanceDetailDTO) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetImageUrl

`func (o *WalletBalanceDetailDTO) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WalletBalanceDetailDTO) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WalletBalanceDetailDTO) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetPriceInSol

`func (o *WalletBalanceDetailDTO) GetPriceInSol() string`

GetPriceInSol returns the PriceInSol field if non-nil, zero value otherwise.

### GetPriceInSolOk

`func (o *WalletBalanceDetailDTO) GetPriceInSolOk() (*string, bool)`

GetPriceInSolOk returns a tuple with the PriceInSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInSol

`func (o *WalletBalanceDetailDTO) SetPriceInSol(v string)`

SetPriceInSol sets PriceInSol field to given value.


### GetPriceInUsd

`func (o *WalletBalanceDetailDTO) GetPriceInUsd() string`

GetPriceInUsd returns the PriceInUsd field if non-nil, zero value otherwise.

### GetPriceInUsdOk

`func (o *WalletBalanceDetailDTO) GetPriceInUsdOk() (*string, bool)`

GetPriceInUsdOk returns a tuple with the PriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInUsd

`func (o *WalletBalanceDetailDTO) SetPriceInUsd(v string)`

SetPriceInUsd sets PriceInUsd field to given value.


### GetPriceChangeRatioInUsd24h

`func (o *WalletBalanceDetailDTO) GetPriceChangeRatioInUsd24h() string`

GetPriceChangeRatioInUsd24h returns the PriceChangeRatioInUsd24h field if non-nil, zero value otherwise.

### GetPriceChangeRatioInUsd24hOk

`func (o *WalletBalanceDetailDTO) GetPriceChangeRatioInUsd24hOk() (*string, bool)`

GetPriceChangeRatioInUsd24hOk returns a tuple with the PriceChangeRatioInUsd24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangeRatioInUsd24h

`func (o *WalletBalanceDetailDTO) SetPriceChangeRatioInUsd24h(v string)`

SetPriceChangeRatioInUsd24h sets PriceChangeRatioInUsd24h field to given value.


### GetUnrealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) GetUnrealizedProfitInUsd() string`

GetUnrealizedProfitInUsd returns the UnrealizedProfitInUsd field if non-nil, zero value otherwise.

### GetUnrealizedProfitInUsdOk

`func (o *WalletBalanceDetailDTO) GetUnrealizedProfitInUsdOk() (*string, bool)`

GetUnrealizedProfitInUsdOk returns a tuple with the UnrealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) SetUnrealizedProfitInUsd(v string)`

SetUnrealizedProfitInUsd sets UnrealizedProfitInUsd field to given value.


### GetUnrealizedProfitRatio

`func (o *WalletBalanceDetailDTO) GetUnrealizedProfitRatio() string`

GetUnrealizedProfitRatio returns the UnrealizedProfitRatio field if non-nil, zero value otherwise.

### GetUnrealizedProfitRatioOk

`func (o *WalletBalanceDetailDTO) GetUnrealizedProfitRatioOk() (*string, bool)`

GetUnrealizedProfitRatioOk returns a tuple with the UnrealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfitRatio

`func (o *WalletBalanceDetailDTO) SetUnrealizedProfitRatio(v string)`

SetUnrealizedProfitRatio sets UnrealizedProfitRatio field to given value.


### GetOpenTime

`func (o *WalletBalanceDetailDTO) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *WalletBalanceDetailDTO) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *WalletBalanceDetailDTO) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.


### GetCloseTime

`func (o *WalletBalanceDetailDTO) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *WalletBalanceDetailDTO) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *WalletBalanceDetailDTO) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.


### GetBuys

`func (o *WalletBalanceDetailDTO) GetBuys() string`

GetBuys returns the Buys field if non-nil, zero value otherwise.

### GetBuysOk

`func (o *WalletBalanceDetailDTO) GetBuysOk() (*string, bool)`

GetBuysOk returns a tuple with the Buys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuys

`func (o *WalletBalanceDetailDTO) SetBuys(v string)`

SetBuys sets Buys field to given value.


### GetBuyAmount

`func (o *WalletBalanceDetailDTO) GetBuyAmount() string`

GetBuyAmount returns the BuyAmount field if non-nil, zero value otherwise.

### GetBuyAmountOk

`func (o *WalletBalanceDetailDTO) GetBuyAmountOk() (*string, bool)`

GetBuyAmountOk returns a tuple with the BuyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmount

`func (o *WalletBalanceDetailDTO) SetBuyAmount(v string)`

SetBuyAmount sets BuyAmount field to given value.


### GetBuyAmountInUsd

`func (o *WalletBalanceDetailDTO) GetBuyAmountInUsd() string`

GetBuyAmountInUsd returns the BuyAmountInUsd field if non-nil, zero value otherwise.

### GetBuyAmountInUsdOk

`func (o *WalletBalanceDetailDTO) GetBuyAmountInUsdOk() (*string, bool)`

GetBuyAmountInUsdOk returns a tuple with the BuyAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyAmountInUsd

`func (o *WalletBalanceDetailDTO) SetBuyAmountInUsd(v string)`

SetBuyAmountInUsd sets BuyAmountInUsd field to given value.


### GetSells

`func (o *WalletBalanceDetailDTO) GetSells() string`

GetSells returns the Sells field if non-nil, zero value otherwise.

### GetSellsOk

`func (o *WalletBalanceDetailDTO) GetSellsOk() (*string, bool)`

GetSellsOk returns a tuple with the Sells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSells

`func (o *WalletBalanceDetailDTO) SetSells(v string)`

SetSells sets Sells field to given value.


### GetSellAmount

`func (o *WalletBalanceDetailDTO) GetSellAmount() string`

GetSellAmount returns the SellAmount field if non-nil, zero value otherwise.

### GetSellAmountOk

`func (o *WalletBalanceDetailDTO) GetSellAmountOk() (*string, bool)`

GetSellAmountOk returns a tuple with the SellAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmount

`func (o *WalletBalanceDetailDTO) SetSellAmount(v string)`

SetSellAmount sets SellAmount field to given value.


### GetSellAmountInUsd

`func (o *WalletBalanceDetailDTO) GetSellAmountInUsd() string`

GetSellAmountInUsd returns the SellAmountInUsd field if non-nil, zero value otherwise.

### GetSellAmountInUsdOk

`func (o *WalletBalanceDetailDTO) GetSellAmountInUsdOk() (*string, bool)`

GetSellAmountInUsdOk returns a tuple with the SellAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAmountInUsd

`func (o *WalletBalanceDetailDTO) SetSellAmountInUsd(v string)`

SetSellAmountInUsd sets SellAmountInUsd field to given value.


### GetAverageBuyPriceInUsd

`func (o *WalletBalanceDetailDTO) GetAverageBuyPriceInUsd() string`

GetAverageBuyPriceInUsd returns the AverageBuyPriceInUsd field if non-nil, zero value otherwise.

### GetAverageBuyPriceInUsdOk

`func (o *WalletBalanceDetailDTO) GetAverageBuyPriceInUsdOk() (*string, bool)`

GetAverageBuyPriceInUsdOk returns a tuple with the AverageBuyPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBuyPriceInUsd

`func (o *WalletBalanceDetailDTO) SetAverageBuyPriceInUsd(v string)`

SetAverageBuyPriceInUsd sets AverageBuyPriceInUsd field to given value.


### GetAverageSellPriceInUsd

`func (o *WalletBalanceDetailDTO) GetAverageSellPriceInUsd() string`

GetAverageSellPriceInUsd returns the AverageSellPriceInUsd field if non-nil, zero value otherwise.

### GetAverageSellPriceInUsdOk

`func (o *WalletBalanceDetailDTO) GetAverageSellPriceInUsdOk() (*string, bool)`

GetAverageSellPriceInUsdOk returns a tuple with the AverageSellPriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSellPriceInUsd

`func (o *WalletBalanceDetailDTO) SetAverageSellPriceInUsd(v string)`

SetAverageSellPriceInUsd sets AverageSellPriceInUsd field to given value.


### GetRealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) GetRealizedProfitInUsd() string`

GetRealizedProfitInUsd returns the RealizedProfitInUsd field if non-nil, zero value otherwise.

### GetRealizedProfitInUsdOk

`func (o *WalletBalanceDetailDTO) GetRealizedProfitInUsdOk() (*string, bool)`

GetRealizedProfitInUsdOk returns a tuple with the RealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) SetRealizedProfitInUsd(v string)`

SetRealizedProfitInUsd sets RealizedProfitInUsd field to given value.


### GetRealizedProfitRatio

`func (o *WalletBalanceDetailDTO) GetRealizedProfitRatio() string`

GetRealizedProfitRatio returns the RealizedProfitRatio field if non-nil, zero value otherwise.

### GetRealizedProfitRatioOk

`func (o *WalletBalanceDetailDTO) GetRealizedProfitRatioOk() (*string, bool)`

GetRealizedProfitRatioOk returns a tuple with the RealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfitRatio

`func (o *WalletBalanceDetailDTO) SetRealizedProfitRatio(v string)`

SetRealizedProfitRatio sets RealizedProfitRatio field to given value.


### GetTotalRealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) GetTotalRealizedProfitInUsd() string`

GetTotalRealizedProfitInUsd returns the TotalRealizedProfitInUsd field if non-nil, zero value otherwise.

### GetTotalRealizedProfitInUsdOk

`func (o *WalletBalanceDetailDTO) GetTotalRealizedProfitInUsdOk() (*string, bool)`

GetTotalRealizedProfitInUsdOk returns a tuple with the TotalRealizedProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitInUsd

`func (o *WalletBalanceDetailDTO) SetTotalRealizedProfitInUsd(v string)`

SetTotalRealizedProfitInUsd sets TotalRealizedProfitInUsd field to given value.


### GetTotalRealizedProfitRatio

`func (o *WalletBalanceDetailDTO) GetTotalRealizedProfitRatio() string`

GetTotalRealizedProfitRatio returns the TotalRealizedProfitRatio field if non-nil, zero value otherwise.

### GetTotalRealizedProfitRatioOk

`func (o *WalletBalanceDetailDTO) GetTotalRealizedProfitRatioOk() (*string, bool)`

GetTotalRealizedProfitRatioOk returns a tuple with the TotalRealizedProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRealizedProfitRatio

`func (o *WalletBalanceDetailDTO) SetTotalRealizedProfitRatio(v string)`

SetTotalRealizedProfitRatio sets TotalRealizedProfitRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


