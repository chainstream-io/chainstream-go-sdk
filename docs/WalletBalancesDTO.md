# WalletBalancesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBalancesInUsd** | **string** | Total wallet value | 
**TotalProfitInUsd** | **string** | Total profit in USD | 
**Balances** | [**[]WalletBalanceDetailDTO**](WalletBalanceDetailDTO.md) | Wallet balance details | 

## Methods

### NewWalletBalancesDTO

`func NewWalletBalancesDTO(totalBalancesInUsd string, totalProfitInUsd string, balances []WalletBalanceDetailDTO, ) *WalletBalancesDTO`

NewWalletBalancesDTO instantiates a new WalletBalancesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalancesDTOWithDefaults

`func NewWalletBalancesDTOWithDefaults() *WalletBalancesDTO`

NewWalletBalancesDTOWithDefaults instantiates a new WalletBalancesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBalancesInUsd

`func (o *WalletBalancesDTO) GetTotalBalancesInUsd() string`

GetTotalBalancesInUsd returns the TotalBalancesInUsd field if non-nil, zero value otherwise.

### GetTotalBalancesInUsdOk

`func (o *WalletBalancesDTO) GetTotalBalancesInUsdOk() (*string, bool)`

GetTotalBalancesInUsdOk returns a tuple with the TotalBalancesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalancesInUsd

`func (o *WalletBalancesDTO) SetTotalBalancesInUsd(v string)`

SetTotalBalancesInUsd sets TotalBalancesInUsd field to given value.


### GetTotalProfitInUsd

`func (o *WalletBalancesDTO) GetTotalProfitInUsd() string`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *WalletBalancesDTO) GetTotalProfitInUsdOk() (*string, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *WalletBalancesDTO) SetTotalProfitInUsd(v string)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.


### GetBalances

`func (o *WalletBalancesDTO) GetBalances() []WalletBalanceDetailDTO`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *WalletBalancesDTO) GetBalancesOk() (*[]WalletBalanceDetailDTO, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *WalletBalancesDTO) SetBalances(v []WalletBalanceDetailDTO)`

SetBalances sets Balances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


