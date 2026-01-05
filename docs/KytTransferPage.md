# KytTransferPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | DTO.KYT.STANDARD_PAGE.TOTAL | 
**Page** | **float32** | DTO.KYT.STANDARD_PAGE.PAGE | 
**PageSize** | **float32** | DTO.KYT.STANDARD_PAGE.PAGE_SIZE | 
**TotalPages** | **float32** | DTO.KYT.STANDARD_PAGE.TOTAL_PAGES | 
**Data** | [**[]KytTransferDTO**](KytTransferDTO.md) | DTO.KYT.KYT_TRANSFER_PAGE.DATA | 

## Methods

### NewKytTransferPage

`func NewKytTransferPage(total float32, page float32, pageSize float32, totalPages float32, data []KytTransferDTO, ) *KytTransferPage`

NewKytTransferPage instantiates a new KytTransferPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytTransferPageWithDefaults

`func NewKytTransferPageWithDefaults() *KytTransferPage`

NewKytTransferPageWithDefaults instantiates a new KytTransferPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *KytTransferPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KytTransferPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KytTransferPage) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *KytTransferPage) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *KytTransferPage) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *KytTransferPage) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *KytTransferPage) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *KytTransferPage) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *KytTransferPage) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetTotalPages

`func (o *KytTransferPage) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *KytTransferPage) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *KytTransferPage) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.


### GetData

`func (o *KytTransferPage) GetData() []KytTransferDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KytTransferPage) GetDataOk() (*[]KytTransferDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KytTransferPage) SetData(v []KytTransferDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


