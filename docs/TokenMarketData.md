# TokenMarketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSupply** | **string** | DTO.TOKEN.MARKET_CAP.TOTAL_SUPPLY | 
**MarketCapInSol** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.MARKET_CAP_IN_SOL | [optional] 
**MarketCapInUsd** | **string** | DTO.TOKEN.MARKET_CAP.MARKET_CAP_IN_USD | 
**Top10TotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP10_TOTAL_HOLDINGS | [optional] 
**Top10HoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP10_HOLDINGS_RATIO | [optional] 
**Top100TotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP100_TOTAL_HOLDINGS | [optional] 
**Top100HoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP100_HOLDINGS_RATIO | [optional] 
**Holders** | **float32** | DTO.TOKEN.MARKET_CAP.HOLDERS | 
**PriceInSol** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.PRICE_IN_SOL | [optional] 
**PriceInUsd** | **string** | DTO.TOKEN.MARKET_CAP.PRICE_IN_USD | 
**TvlInSol** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.LIQUIDITY_IN_SOL | [optional] 
**TvlInUsd** | **string** | DTO.TOKEN.MARKET_CAP.LIQUIDITY_IN_USD | 
**CompletionRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.COMPLETION_RATIO | [optional] 
**Top50TotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP50_TOTAL_HOLDINGS | [optional] 
**Top50HoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.TOP50_HOLDINGS_RATIO | [optional] 
**BluechipTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BLUECHIP_TOTAL_HOLDERS | [optional] 
**BluechipTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BLUECHIP_TOTAL_HOLDINGS | [optional] 
**BluechipHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BLUECHIP_HOLDINGS_RATIO | [optional] 
**KolTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.KOL_TOTAL_HOLDERS | [optional] 
**KolTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.KOL_TOTAL_HOLDINGS | [optional] 
**KolHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.KOL_HOLDINGS_RATIO | [optional] 
**SniperTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SNIPER_TOTAL_HOLDERS | [optional] 
**SniperTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SNIPER_TOTAL_HOLDINGS | [optional] 
**SniperHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SNIPER_HOLDINGS_RATIO | [optional] 
**ProTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.PRO_TOTAL_HOLDERS | [optional] 
**ProTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.PRO_TOTAL_HOLDINGS | [optional] 
**ProHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.PRO_HOLDINGS_RATIO | [optional] 
**InsiderTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.INSIDER_TOTAL_HOLDERS | [optional] 
**InsiderTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.INSIDER_TOTAL_HOLDINGS | [optional] 
**InsiderHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.INSIDER_HOLDINGS_RATIO | [optional] 
**SandwishTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SANDWISH_TOTAL_HOLDERS | [optional] 
**SandwishTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SANDWISH_TOTAL_HOLDINGS | [optional] 
**SandwishHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.SANDWISH_HOLDINGS_RATIO | [optional] 
**FreshTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.FRESH_TOTAL_HOLDERS | [optional] 
**FreshTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.FRESH_TOTAL_HOLDINGS | [optional] 
**FreshHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.FRESH_HOLDINGS_RATIO | [optional] 
**BundleTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BUNDLE_TOTAL_HOLDERS | [optional] 
**BundleTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BUNDLE_TOTAL_HOLDINGS | [optional] 
**BundleHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.BUNDLE_HOLDINGS_RATIO | [optional] 
**DevTotalHolders** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.DEV_TOTAL_HOLDERS | [optional] 
**DevTotalHoldings** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.DEV_TOTAL_HOLDINGS | [optional] 
**DevHoldingsRatio** | Pointer to **string** | DTO.TOKEN.MARKET_CAP.DEV_HOLDINGS_RATIO | [optional] 

## Methods

### NewTokenMarketData

`func NewTokenMarketData(totalSupply string, marketCapInUsd string, holders float32, priceInUsd string, tvlInUsd string, ) *TokenMarketData`

NewTokenMarketData instantiates a new TokenMarketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMarketDataWithDefaults

`func NewTokenMarketDataWithDefaults() *TokenMarketData`

NewTokenMarketDataWithDefaults instantiates a new TokenMarketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSupply

`func (o *TokenMarketData) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenMarketData) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenMarketData) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.


### GetMarketCapInSol

`func (o *TokenMarketData) GetMarketCapInSol() string`

GetMarketCapInSol returns the MarketCapInSol field if non-nil, zero value otherwise.

### GetMarketCapInSolOk

`func (o *TokenMarketData) GetMarketCapInSolOk() (*string, bool)`

GetMarketCapInSolOk returns a tuple with the MarketCapInSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapInSol

`func (o *TokenMarketData) SetMarketCapInSol(v string)`

SetMarketCapInSol sets MarketCapInSol field to given value.

### HasMarketCapInSol

`func (o *TokenMarketData) HasMarketCapInSol() bool`

HasMarketCapInSol returns a boolean if a field has been set.

### GetMarketCapInUsd

`func (o *TokenMarketData) GetMarketCapInUsd() string`

GetMarketCapInUsd returns the MarketCapInUsd field if non-nil, zero value otherwise.

### GetMarketCapInUsdOk

`func (o *TokenMarketData) GetMarketCapInUsdOk() (*string, bool)`

GetMarketCapInUsdOk returns a tuple with the MarketCapInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapInUsd

`func (o *TokenMarketData) SetMarketCapInUsd(v string)`

SetMarketCapInUsd sets MarketCapInUsd field to given value.


### GetTop10TotalHoldings

`func (o *TokenMarketData) GetTop10TotalHoldings() string`

GetTop10TotalHoldings returns the Top10TotalHoldings field if non-nil, zero value otherwise.

### GetTop10TotalHoldingsOk

`func (o *TokenMarketData) GetTop10TotalHoldingsOk() (*string, bool)`

GetTop10TotalHoldingsOk returns a tuple with the Top10TotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10TotalHoldings

`func (o *TokenMarketData) SetTop10TotalHoldings(v string)`

SetTop10TotalHoldings sets Top10TotalHoldings field to given value.

### HasTop10TotalHoldings

`func (o *TokenMarketData) HasTop10TotalHoldings() bool`

HasTop10TotalHoldings returns a boolean if a field has been set.

### GetTop10HoldingsRatio

`func (o *TokenMarketData) GetTop10HoldingsRatio() string`

GetTop10HoldingsRatio returns the Top10HoldingsRatio field if non-nil, zero value otherwise.

### GetTop10HoldingsRatioOk

`func (o *TokenMarketData) GetTop10HoldingsRatioOk() (*string, bool)`

GetTop10HoldingsRatioOk returns a tuple with the Top10HoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10HoldingsRatio

`func (o *TokenMarketData) SetTop10HoldingsRatio(v string)`

SetTop10HoldingsRatio sets Top10HoldingsRatio field to given value.

### HasTop10HoldingsRatio

`func (o *TokenMarketData) HasTop10HoldingsRatio() bool`

HasTop10HoldingsRatio returns a boolean if a field has been set.

### GetTop100TotalHoldings

`func (o *TokenMarketData) GetTop100TotalHoldings() string`

GetTop100TotalHoldings returns the Top100TotalHoldings field if non-nil, zero value otherwise.

### GetTop100TotalHoldingsOk

`func (o *TokenMarketData) GetTop100TotalHoldingsOk() (*string, bool)`

GetTop100TotalHoldingsOk returns a tuple with the Top100TotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop100TotalHoldings

`func (o *TokenMarketData) SetTop100TotalHoldings(v string)`

SetTop100TotalHoldings sets Top100TotalHoldings field to given value.

### HasTop100TotalHoldings

`func (o *TokenMarketData) HasTop100TotalHoldings() bool`

HasTop100TotalHoldings returns a boolean if a field has been set.

### GetTop100HoldingsRatio

`func (o *TokenMarketData) GetTop100HoldingsRatio() string`

GetTop100HoldingsRatio returns the Top100HoldingsRatio field if non-nil, zero value otherwise.

### GetTop100HoldingsRatioOk

`func (o *TokenMarketData) GetTop100HoldingsRatioOk() (*string, bool)`

GetTop100HoldingsRatioOk returns a tuple with the Top100HoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop100HoldingsRatio

`func (o *TokenMarketData) SetTop100HoldingsRatio(v string)`

SetTop100HoldingsRatio sets Top100HoldingsRatio field to given value.

### HasTop100HoldingsRatio

`func (o *TokenMarketData) HasTop100HoldingsRatio() bool`

HasTop100HoldingsRatio returns a boolean if a field has been set.

### GetHolders

`func (o *TokenMarketData) GetHolders() float32`

GetHolders returns the Holders field if non-nil, zero value otherwise.

### GetHoldersOk

`func (o *TokenMarketData) GetHoldersOk() (*float32, bool)`

GetHoldersOk returns a tuple with the Holders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolders

`func (o *TokenMarketData) SetHolders(v float32)`

SetHolders sets Holders field to given value.


### GetPriceInSol

`func (o *TokenMarketData) GetPriceInSol() string`

GetPriceInSol returns the PriceInSol field if non-nil, zero value otherwise.

### GetPriceInSolOk

`func (o *TokenMarketData) GetPriceInSolOk() (*string, bool)`

GetPriceInSolOk returns a tuple with the PriceInSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInSol

`func (o *TokenMarketData) SetPriceInSol(v string)`

SetPriceInSol sets PriceInSol field to given value.

### HasPriceInSol

`func (o *TokenMarketData) HasPriceInSol() bool`

HasPriceInSol returns a boolean if a field has been set.

### GetPriceInUsd

`func (o *TokenMarketData) GetPriceInUsd() string`

GetPriceInUsd returns the PriceInUsd field if non-nil, zero value otherwise.

### GetPriceInUsdOk

`func (o *TokenMarketData) GetPriceInUsdOk() (*string, bool)`

GetPriceInUsdOk returns a tuple with the PriceInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInUsd

`func (o *TokenMarketData) SetPriceInUsd(v string)`

SetPriceInUsd sets PriceInUsd field to given value.


### GetTvlInSol

`func (o *TokenMarketData) GetTvlInSol() string`

GetTvlInSol returns the TvlInSol field if non-nil, zero value otherwise.

### GetTvlInSolOk

`func (o *TokenMarketData) GetTvlInSolOk() (*string, bool)`

GetTvlInSolOk returns a tuple with the TvlInSol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInSol

`func (o *TokenMarketData) SetTvlInSol(v string)`

SetTvlInSol sets TvlInSol field to given value.

### HasTvlInSol

`func (o *TokenMarketData) HasTvlInSol() bool`

HasTvlInSol returns a boolean if a field has been set.

### GetTvlInUsd

`func (o *TokenMarketData) GetTvlInUsd() string`

GetTvlInUsd returns the TvlInUsd field if non-nil, zero value otherwise.

### GetTvlInUsdOk

`func (o *TokenMarketData) GetTvlInUsdOk() (*string, bool)`

GetTvlInUsdOk returns a tuple with the TvlInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvlInUsd

`func (o *TokenMarketData) SetTvlInUsd(v string)`

SetTvlInUsd sets TvlInUsd field to given value.


### GetCompletionRatio

`func (o *TokenMarketData) GetCompletionRatio() string`

GetCompletionRatio returns the CompletionRatio field if non-nil, zero value otherwise.

### GetCompletionRatioOk

`func (o *TokenMarketData) GetCompletionRatioOk() (*string, bool)`

GetCompletionRatioOk returns a tuple with the CompletionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionRatio

`func (o *TokenMarketData) SetCompletionRatio(v string)`

SetCompletionRatio sets CompletionRatio field to given value.

### HasCompletionRatio

`func (o *TokenMarketData) HasCompletionRatio() bool`

HasCompletionRatio returns a boolean if a field has been set.

### GetTop50TotalHoldings

`func (o *TokenMarketData) GetTop50TotalHoldings() string`

GetTop50TotalHoldings returns the Top50TotalHoldings field if non-nil, zero value otherwise.

### GetTop50TotalHoldingsOk

`func (o *TokenMarketData) GetTop50TotalHoldingsOk() (*string, bool)`

GetTop50TotalHoldingsOk returns a tuple with the Top50TotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop50TotalHoldings

`func (o *TokenMarketData) SetTop50TotalHoldings(v string)`

SetTop50TotalHoldings sets Top50TotalHoldings field to given value.

### HasTop50TotalHoldings

`func (o *TokenMarketData) HasTop50TotalHoldings() bool`

HasTop50TotalHoldings returns a boolean if a field has been set.

### GetTop50HoldingsRatio

`func (o *TokenMarketData) GetTop50HoldingsRatio() string`

GetTop50HoldingsRatio returns the Top50HoldingsRatio field if non-nil, zero value otherwise.

### GetTop50HoldingsRatioOk

`func (o *TokenMarketData) GetTop50HoldingsRatioOk() (*string, bool)`

GetTop50HoldingsRatioOk returns a tuple with the Top50HoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop50HoldingsRatio

`func (o *TokenMarketData) SetTop50HoldingsRatio(v string)`

SetTop50HoldingsRatio sets Top50HoldingsRatio field to given value.

### HasTop50HoldingsRatio

`func (o *TokenMarketData) HasTop50HoldingsRatio() bool`

HasTop50HoldingsRatio returns a boolean if a field has been set.

### GetBluechipTotalHolders

`func (o *TokenMarketData) GetBluechipTotalHolders() string`

GetBluechipTotalHolders returns the BluechipTotalHolders field if non-nil, zero value otherwise.

### GetBluechipTotalHoldersOk

`func (o *TokenMarketData) GetBluechipTotalHoldersOk() (*string, bool)`

GetBluechipTotalHoldersOk returns a tuple with the BluechipTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluechipTotalHolders

`func (o *TokenMarketData) SetBluechipTotalHolders(v string)`

SetBluechipTotalHolders sets BluechipTotalHolders field to given value.

### HasBluechipTotalHolders

`func (o *TokenMarketData) HasBluechipTotalHolders() bool`

HasBluechipTotalHolders returns a boolean if a field has been set.

### GetBluechipTotalHoldings

`func (o *TokenMarketData) GetBluechipTotalHoldings() string`

GetBluechipTotalHoldings returns the BluechipTotalHoldings field if non-nil, zero value otherwise.

### GetBluechipTotalHoldingsOk

`func (o *TokenMarketData) GetBluechipTotalHoldingsOk() (*string, bool)`

GetBluechipTotalHoldingsOk returns a tuple with the BluechipTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluechipTotalHoldings

`func (o *TokenMarketData) SetBluechipTotalHoldings(v string)`

SetBluechipTotalHoldings sets BluechipTotalHoldings field to given value.

### HasBluechipTotalHoldings

`func (o *TokenMarketData) HasBluechipTotalHoldings() bool`

HasBluechipTotalHoldings returns a boolean if a field has been set.

### GetBluechipHoldingsRatio

`func (o *TokenMarketData) GetBluechipHoldingsRatio() string`

GetBluechipHoldingsRatio returns the BluechipHoldingsRatio field if non-nil, zero value otherwise.

### GetBluechipHoldingsRatioOk

`func (o *TokenMarketData) GetBluechipHoldingsRatioOk() (*string, bool)`

GetBluechipHoldingsRatioOk returns a tuple with the BluechipHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluechipHoldingsRatio

`func (o *TokenMarketData) SetBluechipHoldingsRatio(v string)`

SetBluechipHoldingsRatio sets BluechipHoldingsRatio field to given value.

### HasBluechipHoldingsRatio

`func (o *TokenMarketData) HasBluechipHoldingsRatio() bool`

HasBluechipHoldingsRatio returns a boolean if a field has been set.

### GetKolTotalHolders

`func (o *TokenMarketData) GetKolTotalHolders() string`

GetKolTotalHolders returns the KolTotalHolders field if non-nil, zero value otherwise.

### GetKolTotalHoldersOk

`func (o *TokenMarketData) GetKolTotalHoldersOk() (*string, bool)`

GetKolTotalHoldersOk returns a tuple with the KolTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKolTotalHolders

`func (o *TokenMarketData) SetKolTotalHolders(v string)`

SetKolTotalHolders sets KolTotalHolders field to given value.

### HasKolTotalHolders

`func (o *TokenMarketData) HasKolTotalHolders() bool`

HasKolTotalHolders returns a boolean if a field has been set.

### GetKolTotalHoldings

`func (o *TokenMarketData) GetKolTotalHoldings() string`

GetKolTotalHoldings returns the KolTotalHoldings field if non-nil, zero value otherwise.

### GetKolTotalHoldingsOk

`func (o *TokenMarketData) GetKolTotalHoldingsOk() (*string, bool)`

GetKolTotalHoldingsOk returns a tuple with the KolTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKolTotalHoldings

`func (o *TokenMarketData) SetKolTotalHoldings(v string)`

SetKolTotalHoldings sets KolTotalHoldings field to given value.

### HasKolTotalHoldings

`func (o *TokenMarketData) HasKolTotalHoldings() bool`

HasKolTotalHoldings returns a boolean if a field has been set.

### GetKolHoldingsRatio

`func (o *TokenMarketData) GetKolHoldingsRatio() string`

GetKolHoldingsRatio returns the KolHoldingsRatio field if non-nil, zero value otherwise.

### GetKolHoldingsRatioOk

`func (o *TokenMarketData) GetKolHoldingsRatioOk() (*string, bool)`

GetKolHoldingsRatioOk returns a tuple with the KolHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKolHoldingsRatio

`func (o *TokenMarketData) SetKolHoldingsRatio(v string)`

SetKolHoldingsRatio sets KolHoldingsRatio field to given value.

### HasKolHoldingsRatio

`func (o *TokenMarketData) HasKolHoldingsRatio() bool`

HasKolHoldingsRatio returns a boolean if a field has been set.

### GetSniperTotalHolders

`func (o *TokenMarketData) GetSniperTotalHolders() string`

GetSniperTotalHolders returns the SniperTotalHolders field if non-nil, zero value otherwise.

### GetSniperTotalHoldersOk

`func (o *TokenMarketData) GetSniperTotalHoldersOk() (*string, bool)`

GetSniperTotalHoldersOk returns a tuple with the SniperTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniperTotalHolders

`func (o *TokenMarketData) SetSniperTotalHolders(v string)`

SetSniperTotalHolders sets SniperTotalHolders field to given value.

### HasSniperTotalHolders

`func (o *TokenMarketData) HasSniperTotalHolders() bool`

HasSniperTotalHolders returns a boolean if a field has been set.

### GetSniperTotalHoldings

`func (o *TokenMarketData) GetSniperTotalHoldings() string`

GetSniperTotalHoldings returns the SniperTotalHoldings field if non-nil, zero value otherwise.

### GetSniperTotalHoldingsOk

`func (o *TokenMarketData) GetSniperTotalHoldingsOk() (*string, bool)`

GetSniperTotalHoldingsOk returns a tuple with the SniperTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniperTotalHoldings

`func (o *TokenMarketData) SetSniperTotalHoldings(v string)`

SetSniperTotalHoldings sets SniperTotalHoldings field to given value.

### HasSniperTotalHoldings

`func (o *TokenMarketData) HasSniperTotalHoldings() bool`

HasSniperTotalHoldings returns a boolean if a field has been set.

### GetSniperHoldingsRatio

`func (o *TokenMarketData) GetSniperHoldingsRatio() string`

GetSniperHoldingsRatio returns the SniperHoldingsRatio field if non-nil, zero value otherwise.

### GetSniperHoldingsRatioOk

`func (o *TokenMarketData) GetSniperHoldingsRatioOk() (*string, bool)`

GetSniperHoldingsRatioOk returns a tuple with the SniperHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniperHoldingsRatio

`func (o *TokenMarketData) SetSniperHoldingsRatio(v string)`

SetSniperHoldingsRatio sets SniperHoldingsRatio field to given value.

### HasSniperHoldingsRatio

`func (o *TokenMarketData) HasSniperHoldingsRatio() bool`

HasSniperHoldingsRatio returns a boolean if a field has been set.

### GetProTotalHolders

`func (o *TokenMarketData) GetProTotalHolders() string`

GetProTotalHolders returns the ProTotalHolders field if non-nil, zero value otherwise.

### GetProTotalHoldersOk

`func (o *TokenMarketData) GetProTotalHoldersOk() (*string, bool)`

GetProTotalHoldersOk returns a tuple with the ProTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProTotalHolders

`func (o *TokenMarketData) SetProTotalHolders(v string)`

SetProTotalHolders sets ProTotalHolders field to given value.

### HasProTotalHolders

`func (o *TokenMarketData) HasProTotalHolders() bool`

HasProTotalHolders returns a boolean if a field has been set.

### GetProTotalHoldings

`func (o *TokenMarketData) GetProTotalHoldings() string`

GetProTotalHoldings returns the ProTotalHoldings field if non-nil, zero value otherwise.

### GetProTotalHoldingsOk

`func (o *TokenMarketData) GetProTotalHoldingsOk() (*string, bool)`

GetProTotalHoldingsOk returns a tuple with the ProTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProTotalHoldings

`func (o *TokenMarketData) SetProTotalHoldings(v string)`

SetProTotalHoldings sets ProTotalHoldings field to given value.

### HasProTotalHoldings

`func (o *TokenMarketData) HasProTotalHoldings() bool`

HasProTotalHoldings returns a boolean if a field has been set.

### GetProHoldingsRatio

`func (o *TokenMarketData) GetProHoldingsRatio() string`

GetProHoldingsRatio returns the ProHoldingsRatio field if non-nil, zero value otherwise.

### GetProHoldingsRatioOk

`func (o *TokenMarketData) GetProHoldingsRatioOk() (*string, bool)`

GetProHoldingsRatioOk returns a tuple with the ProHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProHoldingsRatio

`func (o *TokenMarketData) SetProHoldingsRatio(v string)`

SetProHoldingsRatio sets ProHoldingsRatio field to given value.

### HasProHoldingsRatio

`func (o *TokenMarketData) HasProHoldingsRatio() bool`

HasProHoldingsRatio returns a boolean if a field has been set.

### GetInsiderTotalHolders

`func (o *TokenMarketData) GetInsiderTotalHolders() string`

GetInsiderTotalHolders returns the InsiderTotalHolders field if non-nil, zero value otherwise.

### GetInsiderTotalHoldersOk

`func (o *TokenMarketData) GetInsiderTotalHoldersOk() (*string, bool)`

GetInsiderTotalHoldersOk returns a tuple with the InsiderTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderTotalHolders

`func (o *TokenMarketData) SetInsiderTotalHolders(v string)`

SetInsiderTotalHolders sets InsiderTotalHolders field to given value.

### HasInsiderTotalHolders

`func (o *TokenMarketData) HasInsiderTotalHolders() bool`

HasInsiderTotalHolders returns a boolean if a field has been set.

### GetInsiderTotalHoldings

`func (o *TokenMarketData) GetInsiderTotalHoldings() string`

GetInsiderTotalHoldings returns the InsiderTotalHoldings field if non-nil, zero value otherwise.

### GetInsiderTotalHoldingsOk

`func (o *TokenMarketData) GetInsiderTotalHoldingsOk() (*string, bool)`

GetInsiderTotalHoldingsOk returns a tuple with the InsiderTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderTotalHoldings

`func (o *TokenMarketData) SetInsiderTotalHoldings(v string)`

SetInsiderTotalHoldings sets InsiderTotalHoldings field to given value.

### HasInsiderTotalHoldings

`func (o *TokenMarketData) HasInsiderTotalHoldings() bool`

HasInsiderTotalHoldings returns a boolean if a field has been set.

### GetInsiderHoldingsRatio

`func (o *TokenMarketData) GetInsiderHoldingsRatio() string`

GetInsiderHoldingsRatio returns the InsiderHoldingsRatio field if non-nil, zero value otherwise.

### GetInsiderHoldingsRatioOk

`func (o *TokenMarketData) GetInsiderHoldingsRatioOk() (*string, bool)`

GetInsiderHoldingsRatioOk returns a tuple with the InsiderHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderHoldingsRatio

`func (o *TokenMarketData) SetInsiderHoldingsRatio(v string)`

SetInsiderHoldingsRatio sets InsiderHoldingsRatio field to given value.

### HasInsiderHoldingsRatio

`func (o *TokenMarketData) HasInsiderHoldingsRatio() bool`

HasInsiderHoldingsRatio returns a boolean if a field has been set.

### GetSandwishTotalHolders

`func (o *TokenMarketData) GetSandwishTotalHolders() string`

GetSandwishTotalHolders returns the SandwishTotalHolders field if non-nil, zero value otherwise.

### GetSandwishTotalHoldersOk

`func (o *TokenMarketData) GetSandwishTotalHoldersOk() (*string, bool)`

GetSandwishTotalHoldersOk returns a tuple with the SandwishTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandwishTotalHolders

`func (o *TokenMarketData) SetSandwishTotalHolders(v string)`

SetSandwishTotalHolders sets SandwishTotalHolders field to given value.

### HasSandwishTotalHolders

`func (o *TokenMarketData) HasSandwishTotalHolders() bool`

HasSandwishTotalHolders returns a boolean if a field has been set.

### GetSandwishTotalHoldings

`func (o *TokenMarketData) GetSandwishTotalHoldings() string`

GetSandwishTotalHoldings returns the SandwishTotalHoldings field if non-nil, zero value otherwise.

### GetSandwishTotalHoldingsOk

`func (o *TokenMarketData) GetSandwishTotalHoldingsOk() (*string, bool)`

GetSandwishTotalHoldingsOk returns a tuple with the SandwishTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandwishTotalHoldings

`func (o *TokenMarketData) SetSandwishTotalHoldings(v string)`

SetSandwishTotalHoldings sets SandwishTotalHoldings field to given value.

### HasSandwishTotalHoldings

`func (o *TokenMarketData) HasSandwishTotalHoldings() bool`

HasSandwishTotalHoldings returns a boolean if a field has been set.

### GetSandwishHoldingsRatio

`func (o *TokenMarketData) GetSandwishHoldingsRatio() string`

GetSandwishHoldingsRatio returns the SandwishHoldingsRatio field if non-nil, zero value otherwise.

### GetSandwishHoldingsRatioOk

`func (o *TokenMarketData) GetSandwishHoldingsRatioOk() (*string, bool)`

GetSandwishHoldingsRatioOk returns a tuple with the SandwishHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandwishHoldingsRatio

`func (o *TokenMarketData) SetSandwishHoldingsRatio(v string)`

SetSandwishHoldingsRatio sets SandwishHoldingsRatio field to given value.

### HasSandwishHoldingsRatio

`func (o *TokenMarketData) HasSandwishHoldingsRatio() bool`

HasSandwishHoldingsRatio returns a boolean if a field has been set.

### GetFreshTotalHolders

`func (o *TokenMarketData) GetFreshTotalHolders() string`

GetFreshTotalHolders returns the FreshTotalHolders field if non-nil, zero value otherwise.

### GetFreshTotalHoldersOk

`func (o *TokenMarketData) GetFreshTotalHoldersOk() (*string, bool)`

GetFreshTotalHoldersOk returns a tuple with the FreshTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshTotalHolders

`func (o *TokenMarketData) SetFreshTotalHolders(v string)`

SetFreshTotalHolders sets FreshTotalHolders field to given value.

### HasFreshTotalHolders

`func (o *TokenMarketData) HasFreshTotalHolders() bool`

HasFreshTotalHolders returns a boolean if a field has been set.

### GetFreshTotalHoldings

`func (o *TokenMarketData) GetFreshTotalHoldings() string`

GetFreshTotalHoldings returns the FreshTotalHoldings field if non-nil, zero value otherwise.

### GetFreshTotalHoldingsOk

`func (o *TokenMarketData) GetFreshTotalHoldingsOk() (*string, bool)`

GetFreshTotalHoldingsOk returns a tuple with the FreshTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshTotalHoldings

`func (o *TokenMarketData) SetFreshTotalHoldings(v string)`

SetFreshTotalHoldings sets FreshTotalHoldings field to given value.

### HasFreshTotalHoldings

`func (o *TokenMarketData) HasFreshTotalHoldings() bool`

HasFreshTotalHoldings returns a boolean if a field has been set.

### GetFreshHoldingsRatio

`func (o *TokenMarketData) GetFreshHoldingsRatio() string`

GetFreshHoldingsRatio returns the FreshHoldingsRatio field if non-nil, zero value otherwise.

### GetFreshHoldingsRatioOk

`func (o *TokenMarketData) GetFreshHoldingsRatioOk() (*string, bool)`

GetFreshHoldingsRatioOk returns a tuple with the FreshHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshHoldingsRatio

`func (o *TokenMarketData) SetFreshHoldingsRatio(v string)`

SetFreshHoldingsRatio sets FreshHoldingsRatio field to given value.

### HasFreshHoldingsRatio

`func (o *TokenMarketData) HasFreshHoldingsRatio() bool`

HasFreshHoldingsRatio returns a boolean if a field has been set.

### GetBundleTotalHolders

`func (o *TokenMarketData) GetBundleTotalHolders() string`

GetBundleTotalHolders returns the BundleTotalHolders field if non-nil, zero value otherwise.

### GetBundleTotalHoldersOk

`func (o *TokenMarketData) GetBundleTotalHoldersOk() (*string, bool)`

GetBundleTotalHoldersOk returns a tuple with the BundleTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleTotalHolders

`func (o *TokenMarketData) SetBundleTotalHolders(v string)`

SetBundleTotalHolders sets BundleTotalHolders field to given value.

### HasBundleTotalHolders

`func (o *TokenMarketData) HasBundleTotalHolders() bool`

HasBundleTotalHolders returns a boolean if a field has been set.

### GetBundleTotalHoldings

`func (o *TokenMarketData) GetBundleTotalHoldings() string`

GetBundleTotalHoldings returns the BundleTotalHoldings field if non-nil, zero value otherwise.

### GetBundleTotalHoldingsOk

`func (o *TokenMarketData) GetBundleTotalHoldingsOk() (*string, bool)`

GetBundleTotalHoldingsOk returns a tuple with the BundleTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleTotalHoldings

`func (o *TokenMarketData) SetBundleTotalHoldings(v string)`

SetBundleTotalHoldings sets BundleTotalHoldings field to given value.

### HasBundleTotalHoldings

`func (o *TokenMarketData) HasBundleTotalHoldings() bool`

HasBundleTotalHoldings returns a boolean if a field has been set.

### GetBundleHoldingsRatio

`func (o *TokenMarketData) GetBundleHoldingsRatio() string`

GetBundleHoldingsRatio returns the BundleHoldingsRatio field if non-nil, zero value otherwise.

### GetBundleHoldingsRatioOk

`func (o *TokenMarketData) GetBundleHoldingsRatioOk() (*string, bool)`

GetBundleHoldingsRatioOk returns a tuple with the BundleHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleHoldingsRatio

`func (o *TokenMarketData) SetBundleHoldingsRatio(v string)`

SetBundleHoldingsRatio sets BundleHoldingsRatio field to given value.

### HasBundleHoldingsRatio

`func (o *TokenMarketData) HasBundleHoldingsRatio() bool`

HasBundleHoldingsRatio returns a boolean if a field has been set.

### GetDevTotalHolders

`func (o *TokenMarketData) GetDevTotalHolders() string`

GetDevTotalHolders returns the DevTotalHolders field if non-nil, zero value otherwise.

### GetDevTotalHoldersOk

`func (o *TokenMarketData) GetDevTotalHoldersOk() (*string, bool)`

GetDevTotalHoldersOk returns a tuple with the DevTotalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTotalHolders

`func (o *TokenMarketData) SetDevTotalHolders(v string)`

SetDevTotalHolders sets DevTotalHolders field to given value.

### HasDevTotalHolders

`func (o *TokenMarketData) HasDevTotalHolders() bool`

HasDevTotalHolders returns a boolean if a field has been set.

### GetDevTotalHoldings

`func (o *TokenMarketData) GetDevTotalHoldings() string`

GetDevTotalHoldings returns the DevTotalHoldings field if non-nil, zero value otherwise.

### GetDevTotalHoldingsOk

`func (o *TokenMarketData) GetDevTotalHoldingsOk() (*string, bool)`

GetDevTotalHoldingsOk returns a tuple with the DevTotalHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTotalHoldings

`func (o *TokenMarketData) SetDevTotalHoldings(v string)`

SetDevTotalHoldings sets DevTotalHoldings field to given value.

### HasDevTotalHoldings

`func (o *TokenMarketData) HasDevTotalHoldings() bool`

HasDevTotalHoldings returns a boolean if a field has been set.

### GetDevHoldingsRatio

`func (o *TokenMarketData) GetDevHoldingsRatio() string`

GetDevHoldingsRatio returns the DevHoldingsRatio field if non-nil, zero value otherwise.

### GetDevHoldingsRatioOk

`func (o *TokenMarketData) GetDevHoldingsRatioOk() (*string, bool)`

GetDevHoldingsRatioOk returns a tuple with the DevHoldingsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevHoldingsRatio

`func (o *TokenMarketData) SetDevHoldingsRatio(v string)`

SetDevHoldingsRatio sets DevHoldingsRatio field to given value.

### HasDevHoldingsRatio

`func (o *TokenMarketData) HasDevHoldingsRatio() bool`

HasDevHoldingsRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


