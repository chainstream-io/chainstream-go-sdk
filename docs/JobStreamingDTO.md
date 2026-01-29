# JobStreamingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Streaming job ID | 
**Data** | **map[string]interface{}** | Streaming data | 

## Methods

### NewJobStreamingDTO

`func NewJobStreamingDTO(id float32, data map[string]interface{}, ) *JobStreamingDTO`

NewJobStreamingDTO instantiates a new JobStreamingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStreamingDTOWithDefaults

`func NewJobStreamingDTOWithDefaults() *JobStreamingDTO`

NewJobStreamingDTOWithDefaults instantiates a new JobStreamingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobStreamingDTO) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobStreamingDTO) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobStreamingDTO) SetId(v float32)`

SetId sets Id field to given value.


### GetData

`func (o *JobStreamingDTO) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobStreamingDTO) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobStreamingDTO) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


