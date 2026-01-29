# TokenHolderPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Indicates if there are more results | [optional] [default to false]
**HasPrev** | Pointer to **bool** | Indicates if there are previous results | [optional] [default to false]
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**Total** | Pointer to **float32** | Total number of items | [optional] 
**Data** | [**[]TokenHolder**](TokenHolder.md) | Array of token holder data | 

## Methods

### NewTokenHolderPage

`func NewTokenHolderPage(data []TokenHolder, ) *TokenHolderPage`

NewTokenHolderPage instantiates a new TokenHolderPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHolderPageWithDefaults

`func NewTokenHolderPageWithDefaults() *TokenHolderPage`

NewTokenHolderPageWithDefaults instantiates a new TokenHolderPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *TokenHolderPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TokenHolderPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TokenHolderPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TokenHolderPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *TokenHolderPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TokenHolderPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TokenHolderPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *TokenHolderPage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *TokenHolderPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TokenHolderPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TokenHolderPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TokenHolderPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TokenHolderPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TokenHolderPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TokenHolderPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TokenHolderPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *TokenHolderPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokenHolderPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokenHolderPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TokenHolderPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *TokenHolderPage) GetData() []TokenHolder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenHolderPage) GetDataOk() (*[]TokenHolder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenHolderPage) SetData(v []TokenHolder)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


