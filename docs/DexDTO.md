# DexDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramAddress** | Pointer to **string** | DEX program address | [optional] 
**ProtocolFamily** | Pointer to **string** | DEX protocol family | [optional] 
**Image** | Pointer to **string** | DEX logo image URL | [optional] 
**Chain** | **string** | Blockchain | 

## Methods

### NewDexDTO

`func NewDexDTO(chain string, ) *DexDTO`

NewDexDTO instantiates a new DexDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexDTOWithDefaults

`func NewDexDTOWithDefaults() *DexDTO`

NewDexDTOWithDefaults instantiates a new DexDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramAddress

`func (o *DexDTO) GetProgramAddress() string`

GetProgramAddress returns the ProgramAddress field if non-nil, zero value otherwise.

### GetProgramAddressOk

`func (o *DexDTO) GetProgramAddressOk() (*string, bool)`

GetProgramAddressOk returns a tuple with the ProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramAddress

`func (o *DexDTO) SetProgramAddress(v string)`

SetProgramAddress sets ProgramAddress field to given value.

### HasProgramAddress

`func (o *DexDTO) HasProgramAddress() bool`

HasProgramAddress returns a boolean if a field has been set.

### GetProtocolFamily

`func (o *DexDTO) GetProtocolFamily() string`

GetProtocolFamily returns the ProtocolFamily field if non-nil, zero value otherwise.

### GetProtocolFamilyOk

`func (o *DexDTO) GetProtocolFamilyOk() (*string, bool)`

GetProtocolFamilyOk returns a tuple with the ProtocolFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFamily

`func (o *DexDTO) SetProtocolFamily(v string)`

SetProtocolFamily sets ProtocolFamily field to given value.

### HasProtocolFamily

`func (o *DexDTO) HasProtocolFamily() bool`

HasProtocolFamily returns a boolean if a field has been set.

### GetImage

`func (o *DexDTO) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DexDTO) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DexDTO) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DexDTO) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetChain

`func (o *DexDTO) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *DexDTO) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *DexDTO) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


