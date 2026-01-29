# EndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Endpoint ID | [optional] 
**Url** | Pointer to **string** | Endpoint URL | [optional] 
**Description** | Pointer to **string** | Endpoint description | [optional] 
**FilterTypes** | Pointer to **[]string** | Filter types | [optional] 
**Channels** | Pointer to **[]string** | Endpoint channels | [optional] 
**Disabled** | Pointer to **bool** | Endpoint disabled status | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Endpoint metadata | [optional] 
**RateLimit** | Pointer to **int64** | Rate limit | [optional] 
**Filter** | Pointer to **string** | Endpoint filter | [optional] 
**CreatedAt** | Pointer to **string** | Created at | [optional] 
**UpdatedAt** | Pointer to **string** | Updated at | [optional] 

## Methods

### NewEndpointResponse

`func NewEndpointResponse() *EndpointResponse`

NewEndpointResponse instantiates a new EndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointResponseWithDefaults

`func NewEndpointResponseWithDefaults() *EndpointResponse`

NewEndpointResponseWithDefaults instantiates a new EndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *EndpointResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EndpointResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EndpointResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EndpointResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *EndpointResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EndpointResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterTypes

`func (o *EndpointResponse) GetFilterTypes() []string`

GetFilterTypes returns the FilterTypes field if non-nil, zero value otherwise.

### GetFilterTypesOk

`func (o *EndpointResponse) GetFilterTypesOk() (*[]string, bool)`

GetFilterTypesOk returns a tuple with the FilterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTypes

`func (o *EndpointResponse) SetFilterTypes(v []string)`

SetFilterTypes sets FilterTypes field to given value.

### HasFilterTypes

`func (o *EndpointResponse) HasFilterTypes() bool`

HasFilterTypes returns a boolean if a field has been set.

### GetChannels

`func (o *EndpointResponse) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *EndpointResponse) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *EndpointResponse) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *EndpointResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDisabled

`func (o *EndpointResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *EndpointResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *EndpointResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *EndpointResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *EndpointResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EndpointResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EndpointResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EndpointResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRateLimit

`func (o *EndpointResponse) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *EndpointResponse) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *EndpointResponse) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *EndpointResponse) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetFilter

`func (o *EndpointResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EndpointResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EndpointResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EndpointResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EndpointResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EndpointResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EndpointResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EndpointResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EndpointResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EndpointResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EndpointResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EndpointResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


