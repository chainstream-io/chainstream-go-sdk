# UpdateEndpointInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** | DTO.ENDPOINT.ENDPOINT_ID | [optional] 
**Channels** | Pointer to **[]string** | DTO.ENDPOINT.CHANNELS | [optional] 
**Description** | Pointer to **string** | DTO.ENDPOINT.DESCRIPTION | [optional] 
**Disabled** | Pointer to **bool** | DTO.ENDPOINT.DISABLED | [optional] [default to false]
**FilterTypes** | Pointer to **[]string** | DTO.ENDPOINT.FILTER_TYPES | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | DTO.ENDPOINT.METADATA | [optional] 
**RateLimit** | Pointer to **int64** | DTO.ENDPOINT.RATE_LIMIT | [optional] 
**Url** | Pointer to **string** | DTO.ENDPOINT.URL | [optional] 
**Filter** | Pointer to **string** | DTO.ENDPOINT.FILTER | [optional] 

## Methods

### NewUpdateEndpointInput

`func NewUpdateEndpointInput() *UpdateEndpointInput`

NewUpdateEndpointInput instantiates a new UpdateEndpointInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEndpointInputWithDefaults

`func NewUpdateEndpointInputWithDefaults() *UpdateEndpointInput`

NewUpdateEndpointInputWithDefaults instantiates a new UpdateEndpointInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *UpdateEndpointInput) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *UpdateEndpointInput) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *UpdateEndpointInput) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *UpdateEndpointInput) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetChannels

`func (o *UpdateEndpointInput) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UpdateEndpointInput) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UpdateEndpointInput) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *UpdateEndpointInput) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEndpointInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEndpointInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEndpointInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEndpointInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *UpdateEndpointInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateEndpointInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateEndpointInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateEndpointInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFilterTypes

`func (o *UpdateEndpointInput) GetFilterTypes() []string`

GetFilterTypes returns the FilterTypes field if non-nil, zero value otherwise.

### GetFilterTypesOk

`func (o *UpdateEndpointInput) GetFilterTypesOk() (*[]string, bool)`

GetFilterTypesOk returns a tuple with the FilterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTypes

`func (o *UpdateEndpointInput) SetFilterTypes(v []string)`

SetFilterTypes sets FilterTypes field to given value.

### HasFilterTypes

`func (o *UpdateEndpointInput) HasFilterTypes() bool`

HasFilterTypes returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateEndpointInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateEndpointInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateEndpointInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateEndpointInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRateLimit

`func (o *UpdateEndpointInput) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *UpdateEndpointInput) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *UpdateEndpointInput) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *UpdateEndpointInput) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateEndpointInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateEndpointInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateEndpointInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateEndpointInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateEndpointInput) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateEndpointInput) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateEndpointInput) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateEndpointInput) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


