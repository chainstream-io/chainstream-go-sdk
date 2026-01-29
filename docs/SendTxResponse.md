# SendTxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** | Transaction signature/hash | 
**ElapsedTime** | **int64** | Transaction processing time in milliseconds | 
**JobId** | **string** | job id | 

## Methods

### NewSendTxResponse

`func NewSendTxResponse(signature string, elapsedTime int64, jobId string, ) *SendTxResponse`

NewSendTxResponse instantiates a new SendTxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTxResponseWithDefaults

`func NewSendTxResponseWithDefaults() *SendTxResponse`

NewSendTxResponseWithDefaults instantiates a new SendTxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *SendTxResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SendTxResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SendTxResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetElapsedTime

`func (o *SendTxResponse) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SendTxResponse) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SendTxResponse) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.


### GetJobId

`func (o *SendTxResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SendTxResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SendTxResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


