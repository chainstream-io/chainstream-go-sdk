# TokenTrader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | DTO.TOKEN_TRADER.ADDRESS | 
**TransactionSignature** | Pointer to **string** | DTO.TOKEN_TRADER.TRANSACTION_SIGNATURE | [optional] 
**BlockHash** | Pointer to **string** | DTO.TOKEN_TRADER.BLOCK_HASH | [optional] 
**BlockHeight** | Pointer to **int64** | DTO.TOKEN_TRADER.BLOCK_HEIGHT | [optional] 
**BlockSlot** | Pointer to **int64** | DTO.TOKEN_TRADER.BLOCK_SLOT | [optional] 
**BlockTimestamp** | Pointer to **time.Time** | DTO.TOKEN_TRADER.BLOCK_TIMESTAMP | [optional] 
**OnchainCreatedAt** | Pointer to **map[string]interface{}** | DTO.TOKEN_TRADER.ONCHAIN_CREATED_AT | [optional] 
**TradeCount** | Pointer to **int64** | DTO.TOKEN_TRADER.TRADE_COUNT | [optional] 
**TradeAmountInNative** | Pointer to **string** | DTO.TOKEN_TRADER.TRADE_AMOUNT_IN_NATIVE | [optional] 
**TradeAmountInUsd** | Pointer to **string** | DTO.TOKEN_TRADER.TRADE_AMOUNT_IN_USD | [optional] 
**PercentileRankTradeCount** | Pointer to **int64** | DTO.TOKEN_TRADER.PERCENTILE_RANK_TRADE_COUNT | [optional] 
**PercentileRankTradeAmountInUsd** | Pointer to **string** | DTO.TOKEN_TRADER.PERCENTILE_RANK_TRADE_AMOUNT_IN_USD | [optional] 
**RankTradeAmountInUsd** | Pointer to **string** | DTO.TOKEN_TRADER.RANK_TRADE_AMOUNT_IN_USD | [optional] 

## Methods

### NewTokenTrader

`func NewTokenTrader(address string, ) *TokenTrader`

NewTokenTrader instantiates a new TokenTrader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTraderWithDefaults

`func NewTokenTraderWithDefaults() *TokenTrader`

NewTokenTraderWithDefaults instantiates a new TokenTrader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenTrader) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenTrader) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenTrader) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTransactionSignature

`func (o *TokenTrader) GetTransactionSignature() string`

GetTransactionSignature returns the TransactionSignature field if non-nil, zero value otherwise.

### GetTransactionSignatureOk

`func (o *TokenTrader) GetTransactionSignatureOk() (*string, bool)`

GetTransactionSignatureOk returns a tuple with the TransactionSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSignature

`func (o *TokenTrader) SetTransactionSignature(v string)`

SetTransactionSignature sets TransactionSignature field to given value.

### HasTransactionSignature

`func (o *TokenTrader) HasTransactionSignature() bool`

HasTransactionSignature returns a boolean if a field has been set.

### GetBlockHash

`func (o *TokenTrader) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TokenTrader) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TokenTrader) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TokenTrader) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockHeight

`func (o *TokenTrader) GetBlockHeight() int64`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *TokenTrader) GetBlockHeightOk() (*int64, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *TokenTrader) SetBlockHeight(v int64)`

SetBlockHeight sets BlockHeight field to given value.

### HasBlockHeight

`func (o *TokenTrader) HasBlockHeight() bool`

HasBlockHeight returns a boolean if a field has been set.

### GetBlockSlot

`func (o *TokenTrader) GetBlockSlot() int64`

GetBlockSlot returns the BlockSlot field if non-nil, zero value otherwise.

### GetBlockSlotOk

`func (o *TokenTrader) GetBlockSlotOk() (*int64, bool)`

GetBlockSlotOk returns a tuple with the BlockSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSlot

`func (o *TokenTrader) SetBlockSlot(v int64)`

SetBlockSlot sets BlockSlot field to given value.

### HasBlockSlot

`func (o *TokenTrader) HasBlockSlot() bool`

HasBlockSlot returns a boolean if a field has been set.

### GetBlockTimestamp

`func (o *TokenTrader) GetBlockTimestamp() time.Time`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *TokenTrader) GetBlockTimestampOk() (*time.Time, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *TokenTrader) SetBlockTimestamp(v time.Time)`

SetBlockTimestamp sets BlockTimestamp field to given value.

### HasBlockTimestamp

`func (o *TokenTrader) HasBlockTimestamp() bool`

HasBlockTimestamp returns a boolean if a field has been set.

### GetOnchainCreatedAt

`func (o *TokenTrader) GetOnchainCreatedAt() map[string]interface{}`

GetOnchainCreatedAt returns the OnchainCreatedAt field if non-nil, zero value otherwise.

### GetOnchainCreatedAtOk

`func (o *TokenTrader) GetOnchainCreatedAtOk() (*map[string]interface{}, bool)`

GetOnchainCreatedAtOk returns a tuple with the OnchainCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnchainCreatedAt

`func (o *TokenTrader) SetOnchainCreatedAt(v map[string]interface{})`

SetOnchainCreatedAt sets OnchainCreatedAt field to given value.

### HasOnchainCreatedAt

`func (o *TokenTrader) HasOnchainCreatedAt() bool`

HasOnchainCreatedAt returns a boolean if a field has been set.

### GetTradeCount

`func (o *TokenTrader) GetTradeCount() int64`

GetTradeCount returns the TradeCount field if non-nil, zero value otherwise.

### GetTradeCountOk

`func (o *TokenTrader) GetTradeCountOk() (*int64, bool)`

GetTradeCountOk returns a tuple with the TradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCount

`func (o *TokenTrader) SetTradeCount(v int64)`

SetTradeCount sets TradeCount field to given value.

### HasTradeCount

`func (o *TokenTrader) HasTradeCount() bool`

HasTradeCount returns a boolean if a field has been set.

### GetTradeAmountInNative

`func (o *TokenTrader) GetTradeAmountInNative() string`

GetTradeAmountInNative returns the TradeAmountInNative field if non-nil, zero value otherwise.

### GetTradeAmountInNativeOk

`func (o *TokenTrader) GetTradeAmountInNativeOk() (*string, bool)`

GetTradeAmountInNativeOk returns a tuple with the TradeAmountInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmountInNative

`func (o *TokenTrader) SetTradeAmountInNative(v string)`

SetTradeAmountInNative sets TradeAmountInNative field to given value.

### HasTradeAmountInNative

`func (o *TokenTrader) HasTradeAmountInNative() bool`

HasTradeAmountInNative returns a boolean if a field has been set.

### GetTradeAmountInUsd

`func (o *TokenTrader) GetTradeAmountInUsd() string`

GetTradeAmountInUsd returns the TradeAmountInUsd field if non-nil, zero value otherwise.

### GetTradeAmountInUsdOk

`func (o *TokenTrader) GetTradeAmountInUsdOk() (*string, bool)`

GetTradeAmountInUsdOk returns a tuple with the TradeAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmountInUsd

`func (o *TokenTrader) SetTradeAmountInUsd(v string)`

SetTradeAmountInUsd sets TradeAmountInUsd field to given value.

### HasTradeAmountInUsd

`func (o *TokenTrader) HasTradeAmountInUsd() bool`

HasTradeAmountInUsd returns a boolean if a field has been set.

### GetPercentileRankTradeCount

`func (o *TokenTrader) GetPercentileRankTradeCount() int64`

GetPercentileRankTradeCount returns the PercentileRankTradeCount field if non-nil, zero value otherwise.

### GetPercentileRankTradeCountOk

`func (o *TokenTrader) GetPercentileRankTradeCountOk() (*int64, bool)`

GetPercentileRankTradeCountOk returns a tuple with the PercentileRankTradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileRankTradeCount

`func (o *TokenTrader) SetPercentileRankTradeCount(v int64)`

SetPercentileRankTradeCount sets PercentileRankTradeCount field to given value.

### HasPercentileRankTradeCount

`func (o *TokenTrader) HasPercentileRankTradeCount() bool`

HasPercentileRankTradeCount returns a boolean if a field has been set.

### GetPercentileRankTradeAmountInUsd

`func (o *TokenTrader) GetPercentileRankTradeAmountInUsd() string`

GetPercentileRankTradeAmountInUsd returns the PercentileRankTradeAmountInUsd field if non-nil, zero value otherwise.

### GetPercentileRankTradeAmountInUsdOk

`func (o *TokenTrader) GetPercentileRankTradeAmountInUsdOk() (*string, bool)`

GetPercentileRankTradeAmountInUsdOk returns a tuple with the PercentileRankTradeAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileRankTradeAmountInUsd

`func (o *TokenTrader) SetPercentileRankTradeAmountInUsd(v string)`

SetPercentileRankTradeAmountInUsd sets PercentileRankTradeAmountInUsd field to given value.

### HasPercentileRankTradeAmountInUsd

`func (o *TokenTrader) HasPercentileRankTradeAmountInUsd() bool`

HasPercentileRankTradeAmountInUsd returns a boolean if a field has been set.

### GetRankTradeAmountInUsd

`func (o *TokenTrader) GetRankTradeAmountInUsd() string`

GetRankTradeAmountInUsd returns the RankTradeAmountInUsd field if non-nil, zero value otherwise.

### GetRankTradeAmountInUsdOk

`func (o *TokenTrader) GetRankTradeAmountInUsdOk() (*string, bool)`

GetRankTradeAmountInUsdOk returns a tuple with the RankTradeAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankTradeAmountInUsd

`func (o *TokenTrader) SetRankTradeAmountInUsd(v string)`

SetRankTradeAmountInUsd sets RankTradeAmountInUsd field to given value.

### HasRankTradeAmountInUsd

`func (o *TokenTrader) HasRankTradeAmountInUsd() bool`

HasRankTradeAmountInUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


