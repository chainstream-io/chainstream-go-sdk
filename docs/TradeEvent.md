# TradeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maker** | **string** | Maker address | 
**BaseAmount** | **string** | Base token amount | 
**QuoteAmount** | **string** | Quote token amount | 
**QuoteSymbol** | **string** | Quote token symbol | 
**QuoteAddress** | **string** | Quote token address | 
**AmountInUsd** | **string** | Amount in USD | 
**Timestamp** | **int64** | Trade timestamp (Unix epoch in seconds) | 
**Event** | [**TradeType**](TradeType.md) | Trade event type | 
**TxHash** | **string** | Transaction hash | 
**PriceInUsd** | **string** | Price in USD | 
**Id** | **string** | Trade ID | 
**BuyCostUsd** | **string** | Buy cost in USD | 
**TokenAddress** | **string** | Token address | 

## Methods

### NewTradeEvent

`func NewTradeEvent(maker string, baseAmount string, quoteAmount string, quoteSymbol string, quoteAddress string, amountInUsd string, timestamp int64, event TradeType, txHash string, priceInUsd string, id string, buyCostUsd string, tokenAddress string, ) *TradeEvent`

NewTradeEvent instantiates a new TradeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeEventWithDefaults

`func NewTradeEventWithDefaults() *TradeEvent`

NewTradeEventWithDefaults instantiates a new TradeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaker

`func (o *TradeEvent) GetMaker() string`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *TradeEvent) GetMakerOk() (*string, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *TradeEvent) SetMaker(v string)`

SetMaker sets Maker field to given value.


### GetBaseAmount

`func (o *TradeEvent) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *TradeEvent) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *TradeEvent) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.


### GetQuoteAmount

`func (o *TradeEvent) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *TradeEvent) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *TradeEvent) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.


### GetQuoteSymbol

`func (o *TradeEvent) GetQuoteSymbol() string`

GetQuoteSymbol returns the QuoteSymbol field if non-nil, zero value otherwise.

### GetQuoteSymbolOk

`func (o *TradeEvent) GetQuoteSymbolOk() (*string, bool)`

GetQuoteSymbolOk returns a tuple with the QuoteSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSymbol

`func (o *TradeEvent) SetQuoteSymbol(v string)`

SetQuoteSymbol sets QuoteSymbol field to given value.


### GetQuoteAddress

`func (o *TradeEvent) GetQuoteAddress() string`

GetQuoteAddress returns the QuoteAddress field if non-nil, zero value otherwise.

### GetQuoteAddressOk

`func (o *TradeEvent) GetQuoteAddressOk() (*string, bool)`

GetQuoteAddressOk returns a tuple with the QuoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAddress

`func (o *TradeEvent) SetQuoteAddress(v string)`

SetQuoteAddress sets QuoteAddress field to given value.


### GetAmountInUsd

`func (o *TradeEvent) GetAmountInUsd() string`

GetAmountInUsd returns the AmountInUsd field if non-nil, zero value otherwise.

### GetAmountInUsdOk

`func (o *TradeEvent) GetAmountInUsdOk() (*string, bool)`

GetAmountInUsdOk returns a tuple with the AmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInUsd

`func (o *TradeEvent) SetAmountInUsd(v string)`

SetAmountInUsd sets AmountInUsd field to given value.


### GetTimestamp

`func (o *TradeEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TradeEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TradeEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetEvent

`func (o *TradeEvent) GetEvent() TradeType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TradeEvent) GetEventOk() (*TradeType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TradeEvent) SetEvent(v TradeType)`

SetEvent sets Event field to given value.


### GetTxHash

`func (o *TradeEvent) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TradeEvent) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TradeEvent) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetPriceInUsd

`func (o *TradeEvent) GetPriceInUsd() string`

GetPriceInUsd returns the PriceInUsd field if non-nil, zero value otherwise.

### GetPriceInUsdOk

`func (o *TradeEvent) GetPriceInUsdOk() (*string, bool)`

GetPriceInUsdOk returns a tuple with the PriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInUsd

`func (o *TradeEvent) SetPriceInUsd(v string)`

SetPriceInUsd sets PriceInUsd field to given value.


### GetId

`func (o *TradeEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradeEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradeEvent) SetId(v string)`

SetId sets Id field to given value.


### GetBuyCostUsd

`func (o *TradeEvent) GetBuyCostUsd() string`

GetBuyCostUsd returns the BuyCostUsd field if non-nil, zero value otherwise.

### GetBuyCostUsdOk

`func (o *TradeEvent) GetBuyCostUsdOk() (*string, bool)`

GetBuyCostUsdOk returns a tuple with the BuyCostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyCostUsd

`func (o *TradeEvent) SetBuyCostUsd(v string)`

SetBuyCostUsd sets BuyCostUsd field to given value.


### GetTokenAddress

`func (o *TradeEvent) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TradeEvent) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TradeEvent) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


