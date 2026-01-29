# JobDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Job state | 
**Result** | **map[string]interface{}** | Job result | 

## Methods

### NewJobDTO

`func NewJobDTO(state string, result map[string]interface{}, ) *JobDTO`

NewJobDTO instantiates a new JobDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDTOWithDefaults

`func NewJobDTOWithDefaults() *JobDTO`

NewJobDTOWithDefaults instantiates a new JobDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *JobDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JobDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JobDTO) SetState(v string)`

SetState sets State field to given value.


### GetResult

`func (o *JobDTO) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobDTO) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobDTO) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


