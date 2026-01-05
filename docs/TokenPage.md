# TokenPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | DTO.PAGE.HAS_NEXT | [optional] [default to false]
**HasPrev** | Pointer to **bool** | DTO.PAGE.HAS_PREV | [optional] [default to false]
**StartCursor** | Pointer to **string** | DTO.PAGE.START_CURSOR | [optional] 
**EndCursor** | Pointer to **string** | DTO.PAGE.END_CURSOR | [optional] 
**Total** | Pointer to **float32** | DTO.PAGE.TOTAL | [optional] 
**CountsByProtocols** | Pointer to **map[string]interface{}** | DTO.TOKEN.PAGE.COUNTS_BY_PROTOCOLS | [optional] 
**Data** | [**[]Token**](Token.md) | DTO.TOKEN.PAGE.DATA | 

## Methods

### NewTokenPage

`func NewTokenPage(data []Token, ) *TokenPage`

NewTokenPage instantiates a new TokenPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPageWithDefaults

`func NewTokenPageWithDefaults() *TokenPage`

NewTokenPageWithDefaults instantiates a new TokenPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *TokenPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TokenPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TokenPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TokenPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrev

`func (o *TokenPage) GetHasPrev() bool`

GetHasPrev returns the HasPrev field if non-nil, zero value otherwise.

### GetHasPrevOk

`func (o *TokenPage) GetHasPrevOk() (*bool, bool)`

GetHasPrevOk returns a tuple with the HasPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrev

`func (o *TokenPage) SetHasPrev(v bool)`

SetHasPrev sets HasPrev field to given value.

### HasHasPrev

`func (o *TokenPage) HasHasPrev() bool`

HasHasPrev returns a boolean if a field has been set.

### GetStartCursor

`func (o *TokenPage) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *TokenPage) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *TokenPage) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *TokenPage) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetEndCursor

`func (o *TokenPage) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *TokenPage) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *TokenPage) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *TokenPage) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetTotal

`func (o *TokenPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokenPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokenPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TokenPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCountsByProtocols

`func (o *TokenPage) GetCountsByProtocols() map[string]interface{}`

GetCountsByProtocols returns the CountsByProtocols field if non-nil, zero value otherwise.

### GetCountsByProtocolsOk

`func (o *TokenPage) GetCountsByProtocolsOk() (*map[string]interface{}, bool)`

GetCountsByProtocolsOk returns a tuple with the CountsByProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByProtocols

`func (o *TokenPage) SetCountsByProtocols(v map[string]interface{})`

SetCountsByProtocols sets CountsByProtocols field to given value.

### HasCountsByProtocols

`func (o *TokenPage) HasCountsByProtocols() bool`

HasCountsByProtocols returns a boolean if a field has been set.

### GetData

`func (o *TokenPage) GetData() []Token`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenPage) GetDataOk() (*[]Token, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenPage) SetData(v []Token)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


