# Candle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Open** | **string** | DTO.CANDLE.OPEN | 
**Close** | **string** | DTO.CANDLE.CLOSE | 
**High** | **string** | DTO.CANDLE.HIGH | 
**Low** | **string** | DTO.CANDLE.LOW | 
**Volume** | **string** | DTO.CANDLE.VOLUME | 
**Resolution** | [**Resolution**](Resolution.md) | DTO.CANDLE.RESOLUTION | 
**Time** | **string** | DTO.CANDLE.TIME | 

## Methods

### NewCandle

`func NewCandle(open string, close string, high string, low string, volume string, resolution Resolution, time string, ) *Candle`

NewCandle instantiates a new Candle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandleWithDefaults

`func NewCandleWithDefaults() *Candle`

NewCandleWithDefaults instantiates a new Candle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpen

`func (o *Candle) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Candle) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Candle) SetOpen(v string)`

SetOpen sets Open field to given value.


### GetClose

`func (o *Candle) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Candle) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Candle) SetClose(v string)`

SetClose sets Close field to given value.


### GetHigh

`func (o *Candle) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *Candle) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *Candle) SetHigh(v string)`

SetHigh sets High field to given value.


### GetLow

`func (o *Candle) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *Candle) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *Candle) SetLow(v string)`

SetLow sets Low field to given value.


### GetVolume

`func (o *Candle) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Candle) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Candle) SetVolume(v string)`

SetVolume sets Volume field to given value.


### GetResolution

`func (o *Candle) GetResolution() Resolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Candle) GetResolutionOk() (*Resolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Candle) SetResolution(v Resolution)`

SetResolution sets Resolution field to given value.


### GetTime

`func (o *Candle) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Candle) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Candle) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


