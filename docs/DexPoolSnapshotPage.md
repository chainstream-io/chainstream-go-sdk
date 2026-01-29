# DexPoolSnapshotPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DexPoolSnapshotDTO**](DexPoolSnapshotDTO.md) | Array of data items | 
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**HasNext** | **bool** | Indicates if there are more results | [default to false]
**HasPrev** | **bool** | Indicates if there are previous results | [default to false]

## Methods

### NewDexPoolSnapshotPage

`func NewDexPoolSnapshotPage(data []DexPoolSnapshotDTO, hasNext bool, hasPrev bool, ) *DexPoolSnapshotPage`

NewDexPoolSnapshotPage instantiates a new DexPoolSnapshotPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexPoolSnapshotPageWithDefaults

`func NewDexPoolSnapshotPageWithDefaults() *DexPoolSnapshotPage`

NewDexPoolSnapshotPageWithDefaults instantiates a new DexPoolSnapshotPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DexPoolSnapshotPage) GetData() []DexPoolSnapshotDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DexPoolSnapshotPage) GetDataOk() (*[]DexPoolSnapshotDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DexPoolSnapshotPage) SetData(v []DexPoolSnapshotDTO)`

SetData sets Data field to given value.


### GetStartCursor

`func (o *DexPoolSnapshotPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *DexPoolSnapshotPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *DexPoolSnapshotPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *DexPoolSnapshotPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *DexPoolSnapshotPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *DexPoolSnapshotPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *DexPoolSnapshotPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *DexPoolSnapshotPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetHasNext

`func (o *DexPoolSnapshotPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *DexPoolSnapshotPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *DexPoolSnapshotPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *DexPoolSnapshotPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *DexPoolSnapshotPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *DexPoolSnapshotPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


