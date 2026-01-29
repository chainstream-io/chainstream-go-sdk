# TokenListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Indicates if there are more results | [optional] [default to false]
**HasPrev** | Pointer to **bool** | Indicates if there are previous results | [optional] [default to false]
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**Total** | Pointer to **float32** | Total number of items | [optional] 
**Data** | [**[]Token**](Token.md) | Array of token data | 

## Methods

### NewTokenListPage

`func NewTokenListPage(data []Token, ) *TokenListPage`

NewTokenListPage instantiates a new TokenListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListPageWithDefaults

`func NewTokenListPageWithDefaults() *TokenListPage`

NewTokenListPageWithDefaults instantiates a new TokenListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *TokenListPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TokenListPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TokenListPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TokenListPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *TokenListPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TokenListPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TokenListPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *TokenListPage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *TokenListPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TokenListPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TokenListPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TokenListPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TokenListPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TokenListPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TokenListPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TokenListPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *TokenListPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokenListPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokenListPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TokenListPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *TokenListPage) GetData() []Token`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenListPage) GetDataOk() (*[]Token, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenListPage) SetData(v []Token)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


