# KytAddressPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | DTO.KYT.STANDARD_PAGE.TOTAL | 
**Page** | **float32** | DTO.KYT.STANDARD_PAGE.PAGE | 
**PageSize** | **float32** | DTO.KYT.STANDARD_PAGE.PAGE_SIZE | 
**TotalPages** | **float32** | DTO.KYT.STANDARD_PAGE.TOTAL_PAGES | 
**Data** | [**[]KytAddressDTO**](KytAddressDTO.md) | DTO.KYT.KYT_ADDRESS_PAGE.DATA | 

## Methods

### NewKytAddressPage

`func NewKytAddressPage(total float32, page float32, pageSize float32, totalPages float32, data []KytAddressDTO, ) *KytAddressPage`

NewKytAddressPage instantiates a new KytAddressPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytAddressPageWithDefaults

`func NewKytAddressPageWithDefaults() *KytAddressPage`

NewKytAddressPageWithDefaults instantiates a new KytAddressPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *KytAddressPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KytAddressPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KytAddressPage) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *KytAddressPage) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *KytAddressPage) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *KytAddressPage) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *KytAddressPage) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *KytAddressPage) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *KytAddressPage) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetTotalPages

`func (o *KytAddressPage) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *KytAddressPage) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *KytAddressPage) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.


### GetData

`func (o *KytAddressPage) GetData() []KytAddressDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KytAddressPage) GetDataOk() (*[]KytAddressDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KytAddressPage) SetData(v []KytAddressDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


