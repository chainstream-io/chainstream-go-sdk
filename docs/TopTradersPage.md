# TopTradersPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Indicates if there are more results | [optional] [default to false]
**HasPrev** | Pointer to **bool** | Indicates if there are previous results | [optional] [default to false]
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**Total** | Pointer to **float32** | Total number of items | [optional] 
**Data** | [**[]TopTradersDTO**](TopTradersDTO.md) | Array of top traders data | 

## Methods

### NewTopTradersPage

`func NewTopTradersPage(data []TopTradersDTO, ) *TopTradersPage`

NewTopTradersPage instantiates a new TopTradersPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopTradersPageWithDefaults

`func NewTopTradersPageWithDefaults() *TopTradersPage`

NewTopTradersPageWithDefaults instantiates a new TopTradersPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *TopTradersPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TopTradersPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TopTradersPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TopTradersPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *TopTradersPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TopTradersPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TopTradersPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *TopTradersPage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *TopTradersPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TopTradersPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TopTradersPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TopTradersPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TopTradersPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TopTradersPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TopTradersPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TopTradersPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *TopTradersPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TopTradersPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TopTradersPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TopTradersPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *TopTradersPage) GetData() []TopTradersDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TopTradersPage) GetDataOk() (*[]TopTradersDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TopTradersPage) SetData(v []TopTradersDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


