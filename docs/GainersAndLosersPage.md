# GainersAndLosersPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | DTO.PAGE.HAS_NEXT | [optional] [default to false]
**HasPrev** | Pointer to **bool** | DTO.PAGE.HAS_PREV | [optional] [default to false]
**StartCursor** | Pointer to **string** | DTO.PAGE.START_CURSOR | [optional] 
**EndCursor** | Pointer to **string** | DTO.PAGE.END_CURSOR | [optional] 
**Total** | Pointer to **float32** | DTO.PAGE.TOTAL | [optional] 
**Data** | [**[]GainersAndLosersDTO**](GainersAndLosersDTO.md) | DTO.TRADE.GAINERS_LOSERS.PAGE.DATA | 

## Methods

### NewGainersAndLosersPage

`func NewGainersAndLosersPage(data []GainersAndLosersDTO, ) *GainersAndLosersPage`

NewGainersAndLosersPage instantiates a new GainersAndLosersPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGainersAndLosersPageWithDefaults

`func NewGainersAndLosersPageWithDefaults() *GainersAndLosersPage`

NewGainersAndLosersPageWithDefaults instantiates a new GainersAndLosersPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *GainersAndLosersPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *GainersAndLosersPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *GainersAndLosersPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *GainersAndLosersPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *GainersAndLosersPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *GainersAndLosersPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *GainersAndLosersPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *GainersAndLosersPage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *GainersAndLosersPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *GainersAndLosersPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *GainersAndLosersPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *GainersAndLosersPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *GainersAndLosersPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *GainersAndLosersPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *GainersAndLosersPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *GainersAndLosersPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *GainersAndLosersPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GainersAndLosersPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GainersAndLosersPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GainersAndLosersPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *GainersAndLosersPage) GetData() []GainersAndLosersDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GainersAndLosersPage) GetDataOk() (*[]GainersAndLosersDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GainersAndLosersPage) SetData(v []GainersAndLosersDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


