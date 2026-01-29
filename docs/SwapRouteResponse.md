# SwapRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | [**SwapRouteInput**](SwapRouteInput.md) | Original swap request parameters | 
**SerializedTx** | **string** | Base64 encoded transaction | 
**RouteInfo** | **map[string]interface{}** | Detailed routing information | 
**ElapsedTime** | **int64** | Time taken to process the request in milliseconds | 

## Methods

### NewSwapRouteResponse

`func NewSwapRouteResponse(args SwapRouteInput, serializedTx string, routeInfo map[string]interface{}, elapsedTime int64, ) *SwapRouteResponse`

NewSwapRouteResponse instantiates a new SwapRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapRouteResponseWithDefaults

`func NewSwapRouteResponseWithDefaults() *SwapRouteResponse`

NewSwapRouteResponseWithDefaults instantiates a new SwapRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *SwapRouteResponse) GetArgs() SwapRouteInput`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *SwapRouteResponse) GetArgsOk() (*SwapRouteInput, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *SwapRouteResponse) SetArgs(v SwapRouteInput)`

SetArgs sets Args field to given value.


### GetSerializedTx

`func (o *SwapRouteResponse) GetSerializedTx() string`

GetSerializedTx returns the SerializedTx field if non-nil, zero value otherwise.

### GetSerializedTxOk

`func (o *SwapRouteResponse) GetSerializedTxOk() (*string, bool)`

GetSerializedTxOk returns a tuple with the SerializedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedTx

`func (o *SwapRouteResponse) SetSerializedTx(v string)`

SetSerializedTx sets SerializedTx field to given value.


### GetRouteInfo

`func (o *SwapRouteResponse) GetRouteInfo() map[string]interface{}`

GetRouteInfo returns the RouteInfo field if non-nil, zero value otherwise.

### GetRouteInfoOk

`func (o *SwapRouteResponse) GetRouteInfoOk() (*map[string]interface{}, bool)`

GetRouteInfoOk returns a tuple with the RouteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteInfo

`func (o *SwapRouteResponse) SetRouteInfo(v map[string]interface{})`

SetRouteInfo sets RouteInfo field to given value.


### GetElapsedTime

`func (o *SwapRouteResponse) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SwapRouteResponse) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SwapRouteResponse) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


