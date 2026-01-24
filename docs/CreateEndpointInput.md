# CreateEndpointInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** | DTO.ENDPOINT.CHANNELS | [optional] 
**Description** | Pointer to **string** | DTO.ENDPOINT.DESCRIPTION | [optional] 
**Disabled** | Pointer to **bool** | DTO.ENDPOINT.DISABLED | [optional] [default to false]
**FilterTypes** | Pointer to **[]string** | DTO.ENDPOINT.FILTER_TYPES | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | DTO.ENDPOINT.METADATA | [optional] 
**RateLimit** | Pointer to **int64** | DTO.ENDPOINT.RATE_LIMIT | [optional] 
**Url** | Pointer to **string** | DTO.ENDPOINT.URL | [optional] 
**Filter** | Pointer to **string** | DTO.ENDPOINT.FILTER | [optional] 

## Methods

### NewCreateEndpointInput

`func NewCreateEndpointInput() *CreateEndpointInput`

NewCreateEndpointInput instantiates a new CreateEndpointInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointInputWithDefaults

`func NewCreateEndpointInputWithDefaults() *CreateEndpointInput`

NewCreateEndpointInputWithDefaults instantiates a new CreateEndpointInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *CreateEndpointInput) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CreateEndpointInput) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CreateEndpointInput) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CreateEndpointInput) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDescription

`func (o *CreateEndpointInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEndpointInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEndpointInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEndpointInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *CreateEndpointInput) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateEndpointInput) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateEndpointInput) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateEndpointInput) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFilterTypes

`func (o *CreateEndpointInput) GetFilterTypes() []string`

GetFilterTypes returns the FilterTypes field if non-nil, zero value otherwise.

### GetFilterTypesOk

`func (o *CreateEndpointInput) GetFilterTypesOk() (*[]string, bool)`

GetFilterTypesOk returns a tuple with the FilterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTypes

`func (o *CreateEndpointInput) SetFilterTypes(v []string)`

SetFilterTypes sets FilterTypes field to given value.

### HasFilterTypes

`func (o *CreateEndpointInput) HasFilterTypes() bool`

HasFilterTypes returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateEndpointInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateEndpointInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateEndpointInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateEndpointInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRateLimit

`func (o *CreateEndpointInput) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *CreateEndpointInput) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *CreateEndpointInput) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *CreateEndpointInput) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetUrl

`func (o *CreateEndpointInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateEndpointInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateEndpointInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateEndpointInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilter

`func (o *CreateEndpointInput) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateEndpointInput) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateEndpointInput) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateEndpointInput) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


