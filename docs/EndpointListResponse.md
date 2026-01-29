# EndpointListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EndpointResponse**](EndpointResponse.md) | Page data | [optional] 
**Done** | Pointer to **bool** | Page done | [optional] 
**Iterator** | Pointer to **string** | Page iterator | [optional] 
**PrevIterator** | Pointer to **string** | Previous page iterator | [optional] 

## Methods

### NewEndpointListResponse

`func NewEndpointListResponse() *EndpointListResponse`

NewEndpointListResponse instantiates a new EndpointListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointListResponseWithDefaults

`func NewEndpointListResponseWithDefaults() *EndpointListResponse`

NewEndpointListResponseWithDefaults instantiates a new EndpointListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EndpointListResponse) GetData() []EndpointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndpointListResponse) GetDataOk() (*[]EndpointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndpointListResponse) SetData(v []EndpointResponse)`

SetData sets Data field to given value.

### HasData

`func (o *EndpointListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDone

`func (o *EndpointListResponse) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *EndpointListResponse) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *EndpointListResponse) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *EndpointListResponse) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetIterator

`func (o *EndpointListResponse) GetIterator() string`

GetIterator returns the Iterator field if non-nil, zero value otherwise.

### GetIteratorOk

`func (o *EndpointListResponse) GetIteratorOk() (*string, bool)`

GetIteratorOk returns a tuple with the Iterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterator

`func (o *EndpointListResponse) SetIterator(v string)`

SetIterator sets Iterator field to given value.

### HasIterator

`func (o *EndpointListResponse) HasIterator() bool`

HasIterator returns a boolean if a field has been set.

### GetPrevIterator

`func (o *EndpointListResponse) GetPrevIterator() string`

GetPrevIterator returns the PrevIterator field if non-nil, zero value otherwise.

### GetPrevIteratorOk

`func (o *EndpointListResponse) GetPrevIteratorOk() (*string, bool)`

GetPrevIteratorOk returns a tuple with the PrevIterator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevIterator

`func (o *EndpointListResponse) SetPrevIterator(v string)`

SetPrevIterator sets PrevIterator field to given value.

### HasPrevIterator

`func (o *EndpointListResponse) HasPrevIterator() bool`

HasPrevIterator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


