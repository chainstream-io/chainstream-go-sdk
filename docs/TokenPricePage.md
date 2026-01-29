# TokenPricePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Indicates if there are more results | [optional] [default to false]
**HasPrev** | Pointer to **bool** | Indicates if there are previous results | [optional] [default to false]
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**Total** | Pointer to **float32** | Total number of items | [optional] 
**Data** | [**[]TokenPriceDTO**](TokenPriceDTO.md) | Array of token price data | 

## Methods

### NewTokenPricePage

`func NewTokenPricePage(data []TokenPriceDTO, ) *TokenPricePage`

NewTokenPricePage instantiates a new TokenPricePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPricePageWithDefaults

`func NewTokenPricePageWithDefaults() *TokenPricePage`

NewTokenPricePageWithDefaults instantiates a new TokenPricePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *TokenPricePage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TokenPricePage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TokenPricePage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TokenPricePage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *TokenPricePage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TokenPricePage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TokenPricePage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *TokenPricePage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *TokenPricePage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TokenPricePage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TokenPricePage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TokenPricePage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TokenPricePage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TokenPricePage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TokenPricePage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TokenPricePage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *TokenPricePage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokenPricePage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokenPricePage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TokenPricePage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *TokenPricePage) GetData() []TokenPriceDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenPricePage) GetDataOk() (*[]TokenPriceDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenPricePage) SetData(v []TokenPriceDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


