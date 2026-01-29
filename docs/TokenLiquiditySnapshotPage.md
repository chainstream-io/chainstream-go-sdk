# TokenLiquiditySnapshotPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenLiquiditySnapshotDTO**](TokenLiquiditySnapshotDTO.md) | Array of data items | 
**StartCursor** | Pointer to **string** | Cursor for first item in current page | [optional] 
**EndCursor** | Pointer to **string** | Cursor for last item in current page | [optional] 
**HasNext** | **bool** | Indicates if there are more results | [default to false]
**HasPrev** | **bool** | Indicates if there are previous results | [default to false]

## Methods

### NewTokenLiquiditySnapshotPage

`func NewTokenLiquiditySnapshotPage(data []TokenLiquiditySnapshotDTO, hasNext bool, hasPrev bool, ) *TokenLiquiditySnapshotPage`

NewTokenLiquiditySnapshotPage instantiates a new TokenLiquiditySnapshotPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLiquiditySnapshotPageWithDefaults

`func NewTokenLiquiditySnapshotPageWithDefaults() *TokenLiquiditySnapshotPage`

NewTokenLiquiditySnapshotPageWithDefaults instantiates a new TokenLiquiditySnapshotPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenLiquiditySnapshotPage) GetData() []TokenLiquiditySnapshotDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenLiquiditySnapshotPage) GetDataOk() (*[]TokenLiquiditySnapshotDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenLiquiditySnapshotPage) SetData(v []TokenLiquiditySnapshotDTO)`

SetData sets Data field to given value.


### GetStartCursor

`func (o *TokenLiquiditySnapshotPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TokenLiquiditySnapshotPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TokenLiquiditySnapshotPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TokenLiquiditySnapshotPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TokenLiquiditySnapshotPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TokenLiquiditySnapshotPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TokenLiquiditySnapshotPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TokenLiquiditySnapshotPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetHasNext

`func (o *TokenLiquiditySnapshotPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TokenLiquiditySnapshotPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TokenLiquiditySnapshotPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrev

`func (o *TokenLiquiditySnapshotPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TokenLiquiditySnapshotPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TokenLiquiditySnapshotPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


