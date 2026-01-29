# RedPacketsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | Total number of items | 
**HasNextPage** | **bool** | Indicates if there are more results | 
**StartCursor** | **string** | Cursor for first item in current page | 
**EndCursor** | **string** | Cursor for last item in current page | 
**Records** | [**[]RedPacketDTO**](RedPacketDTO.md) | Array of red packets | 

## Methods

### NewRedPacketsPage

`func NewRedPacketsPage(total int64, hasNextPage bool, startCursor string, endCursor string, records []RedPacketDTO, ) *RedPacketsPage`

NewRedPacketsPage instantiates a new RedPacketsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedPacketsPageWithDefaults

`func NewRedPacketsPageWithDefaults() *RedPacketsPage`

NewRedPacketsPageWithDefaults instantiates a new RedPacketsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *RedPacketsPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RedPacketsPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RedPacketsPage) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetHasNextPage

`func (o *RedPacketsPage) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *RedPacketsPage) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *RedPacketsPage) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetStartCursor

`func (o *RedPacketsPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *RedPacketsPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *RedPacketsPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.


### GetEndCursor

`func (o *RedPacketsPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *RedPacketsPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *RedPacketsPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.


### GetRecords

`func (o *RedPacketsPage) GetRecords() []RedPacketDTO`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RedPacketsPage) GetRecordsOk() (*[]RedPacketDTO, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RedPacketsPage) SetRecords(v []RedPacketDTO)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


