# RedPacketClaimsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | DTO.PAGE.TOTAL | 
**HasNextPage** | **bool** | DTO.PAGE.HAS_NEXT | 
**StartCursor** | **string** | DTO.PAGE.START_CURSOR | 
**EndCursor** | **string** | DTO.PAGE.END_CURSOR | 
**Records** | [**[]RedPacketClaimDTO**](RedPacketClaimDTO.md) | Array of claim records | 

## Methods

### NewRedPacketClaimsPage

`func NewRedPacketClaimsPage(total int64, hasNextPage bool, startCursor string, endCursor string, records []RedPacketClaimDTO, ) *RedPacketClaimsPage`

NewRedPacketClaimsPage instantiates a new RedPacketClaimsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedPacketClaimsPageWithDefaults

`func NewRedPacketClaimsPageWithDefaults() *RedPacketClaimsPage`

NewRedPacketClaimsPageWithDefaults instantiates a new RedPacketClaimsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *RedPacketClaimsPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RedPacketClaimsPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RedPacketClaimsPage) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetHasNextPage

`func (o *RedPacketClaimsPage) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *RedPacketClaimsPage) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *RedPacketClaimsPage) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetStartCursor

`func (o *RedPacketClaimsPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *RedPacketClaimsPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *RedPacketClaimsPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.


### GetEndCursor

`func (o *RedPacketClaimsPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *RedPacketClaimsPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *RedPacketClaimsPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.


### GetRecords

`func (o *RedPacketClaimsPage) GetRecords() []RedPacketClaimDTO`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RedPacketClaimsPage) GetRecordsOk() (*[]RedPacketClaimDTO, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RedPacketClaimsPage) SetRecords(v []RedPacketClaimDTO)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


