# TokenExtraDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | Pointer to **string** | DTO.TOKEN.EXTRA.COLLECTION_ADDRESS | [optional] 
**EditionNonce** | Pointer to **int64** | DTO.TOKEN.EXTRA.EDITION_NONCE | [optional] 
**Fungible** | Pointer to **bool** | DTO.TOKEN.EXTRA.FUNGIBLE | [optional] 
**IsMutable** | Pointer to **bool** | DTO.TOKEN.EXTRA.IS_MUTABLE | [optional] 
**Key** | Pointer to **string** | DTO.TOKEN.EXTRA.KEY | [optional] 
**IsNative** | Pointer to **bool** | DTO.TOKEN.EXTRA.IS_NATIVE | [optional] 
**PrimarySaleHappened** | Pointer to **bool** | DTO.TOKEN.EXTRA.PRIMARY_SALE_HAPPENED | [optional] 
**LaunchFromProgramAddress** | Pointer to **string** | DTO.TOKEN.EXTRA.LAUNCH_FROM_PROGRAM_ADDRESS | [optional] 
**LaunchFromProtocolFamily** | Pointer to **string** | DTO.TOKEN.EXTRA.LAUNCH_FROM_PROTOCOL_FAMILY | [optional] 
**ProgramAddress** | Pointer to **string** | DTO.TOKEN.EXTRA.PROGRAM_ADDRESS | [optional] 
**MigratedToProgramAddress** | Pointer to **string** | DTO.TOKEN.EXTRA.MIGRATED_TO_PROGRAM_ADDRESS | [optional] 
**MigratedToProtocolFamily** | Pointer to **string** | DTO.TOKEN.EXTRA.MIGRATED_TO_PROTOCOL_FAMILY | [optional] 
**MigratedToPoolAddress** | Pointer to **string** | DTO.TOKEN.EXTRA.MIGRATED_TO_POOL_ADDRESS | [optional] 
**MigratedAt** | Pointer to **int64** | DTO.TOKEN.EXTRA.MIGRATED_AT | [optional] 
**SellerFeeBasisPoints** | Pointer to **int64** | DTO.TOKEN.EXTRA.SELLER_FEE_BASIS_POINTS | [optional] 
**TokenStandard** | Pointer to **string** | DTO.TOKEN.EXTRA.TOKEN_STANDARD | [optional] 
**MintAuthority** | Pointer to **string** | DTO.TOKEN.EXTRA.MINT_AUTHORITY | [optional] 
**FreezeAuthority** | Pointer to **string** | DTO.TOKEN.EXTRA.FREEZE_AUTHORITY | [optional] 
**UpdateAuthority** | Pointer to **string** | DTO.TOKEN.EXTRA.UPDATE_AUTHORITY | [optional] 
**IsVerifiedCollection** | Pointer to **bool** | DTO.TOKEN.EXTRA.IS_VERIFIED_COLLECTION | [optional] 
**IsWrapped** | Pointer to **bool** | DTO.TOKEN.EXTRA.IS_WRAPPED | [optional] 

## Methods

### NewTokenExtraDTO

`func NewTokenExtraDTO() *TokenExtraDTO`

NewTokenExtraDTO instantiates a new TokenExtraDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExtraDTOWithDefaults

`func NewTokenExtraDTOWithDefaults() *TokenExtraDTO`

NewTokenExtraDTOWithDefaults instantiates a new TokenExtraDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *TokenExtraDTO) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *TokenExtraDTO) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *TokenExtraDTO) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.

### HasCollectionAddress

`func (o *TokenExtraDTO) HasCollectionAddress() bool`

HasCollectionAddress returns a boolean if a field has been set.

### GetEditionNonce

`func (o *TokenExtraDTO) GetEditionNonce() int64`

GetEditionNonce returns the EditionNonce field if non-nil, zero value otherwise.

### GetEditionNonceOk

`func (o *TokenExtraDTO) GetEditionNonceOk() (*int64, bool)`

GetEditionNonceOk returns a tuple with the EditionNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionNonce

`func (o *TokenExtraDTO) SetEditionNonce(v int64)`

SetEditionNonce sets EditionNonce field to given value.

### HasEditionNonce

`func (o *TokenExtraDTO) HasEditionNonce() bool`

HasEditionNonce returns a boolean if a field has been set.

### GetFungible

`func (o *TokenExtraDTO) GetFungible() bool`

GetFungible returns the Fungible field if non-nil, zero value otherwise.

### GetFungibleOk

`func (o *TokenExtraDTO) GetFungibleOk() (*bool, bool)`

GetFungibleOk returns a tuple with the Fungible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungible

`func (o *TokenExtraDTO) SetFungible(v bool)`

SetFungible sets Fungible field to given value.

### HasFungible

`func (o *TokenExtraDTO) HasFungible() bool`

HasFungible returns a boolean if a field has been set.

### GetIsMutable

`func (o *TokenExtraDTO) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *TokenExtraDTO) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *TokenExtraDTO) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *TokenExtraDTO) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetKey

`func (o *TokenExtraDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TokenExtraDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TokenExtraDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TokenExtraDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIsNative

`func (o *TokenExtraDTO) GetIsNative() bool`

GetIsNative returns the IsNative field if non-nil, zero value otherwise.

### GetIsNativeOk

`func (o *TokenExtraDTO) GetIsNativeOk() (*bool, bool)`

GetIsNativeOk returns a tuple with the IsNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNative

`func (o *TokenExtraDTO) SetIsNative(v bool)`

SetIsNative sets IsNative field to given value.

### HasIsNative

`func (o *TokenExtraDTO) HasIsNative() bool`

HasIsNative returns a boolean if a field has been set.

### GetPrimarySaleHappened

`func (o *TokenExtraDTO) GetPrimarySaleHappened() bool`

GetPrimarySaleHappened returns the PrimarySaleHappened field if non-nil, zero value otherwise.

### GetPrimarySaleHappenedOk

`func (o *TokenExtraDTO) GetPrimarySaleHappenedOk() (*bool, bool)`

GetPrimarySaleHappenedOk returns a tuple with the PrimarySaleHappened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySaleHappened

`func (o *TokenExtraDTO) SetPrimarySaleHappened(v bool)`

SetPrimarySaleHappened sets PrimarySaleHappened field to given value.

### HasPrimarySaleHappened

`func (o *TokenExtraDTO) HasPrimarySaleHappened() bool`

HasPrimarySaleHappened returns a boolean if a field has been set.

### GetLaunchFromProgramAddress

`func (o *TokenExtraDTO) GetLaunchFromProgramAddress() string`

GetLaunchFromProgramAddress returns the LaunchFromProgramAddress field if non-nil, zero value otherwise.

### GetLaunchFromProgramAddressOk

`func (o *TokenExtraDTO) GetLaunchFromProgramAddressOk() (*string, bool)`

GetLaunchFromProgramAddressOk returns a tuple with the LaunchFromProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchFromProgramAddress

`func (o *TokenExtraDTO) SetLaunchFromProgramAddress(v string)`

SetLaunchFromProgramAddress sets LaunchFromProgramAddress field to given value.

### HasLaunchFromProgramAddress

`func (o *TokenExtraDTO) HasLaunchFromProgramAddress() bool`

HasLaunchFromProgramAddress returns a boolean if a field has been set.

### GetLaunchFromProtocolFamily

`func (o *TokenExtraDTO) GetLaunchFromProtocolFamily() string`

GetLaunchFromProtocolFamily returns the LaunchFromProtocolFamily field if non-nil, zero value otherwise.

### GetLaunchFromProtocolFamilyOk

`func (o *TokenExtraDTO) GetLaunchFromProtocolFamilyOk() (*string, bool)`

GetLaunchFromProtocolFamilyOk returns a tuple with the LaunchFromProtocolFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchFromProtocolFamily

`func (o *TokenExtraDTO) SetLaunchFromProtocolFamily(v string)`

SetLaunchFromProtocolFamily sets LaunchFromProtocolFamily field to given value.

### HasLaunchFromProtocolFamily

`func (o *TokenExtraDTO) HasLaunchFromProtocolFamily() bool`

HasLaunchFromProtocolFamily returns a boolean if a field has been set.

### GetProgramAddress

`func (o *TokenExtraDTO) GetProgramAddress() string`

GetProgramAddress returns the ProgramAddress field if non-nil, zero value otherwise.

### GetProgramAddressOk

`func (o *TokenExtraDTO) GetProgramAddressOk() (*string, bool)`

GetProgramAddressOk returns a tuple with the ProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramAddress

`func (o *TokenExtraDTO) SetProgramAddress(v string)`

SetProgramAddress sets ProgramAddress field to given value.

### HasProgramAddress

`func (o *TokenExtraDTO) HasProgramAddress() bool`

HasProgramAddress returns a boolean if a field has been set.

### GetMigratedToProgramAddress

`func (o *TokenExtraDTO) GetMigratedToProgramAddress() string`

GetMigratedToProgramAddress returns the MigratedToProgramAddress field if non-nil, zero value otherwise.

### GetMigratedToProgramAddressOk

`func (o *TokenExtraDTO) GetMigratedToProgramAddressOk() (*string, bool)`

GetMigratedToProgramAddressOk returns a tuple with the MigratedToProgramAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedToProgramAddress

`func (o *TokenExtraDTO) SetMigratedToProgramAddress(v string)`

SetMigratedToProgramAddress sets MigratedToProgramAddress field to given value.

### HasMigratedToProgramAddress

`func (o *TokenExtraDTO) HasMigratedToProgramAddress() bool`

HasMigratedToProgramAddress returns a boolean if a field has been set.

### GetMigratedToProtocolFamily

`func (o *TokenExtraDTO) GetMigratedToProtocolFamily() string`

GetMigratedToProtocolFamily returns the MigratedToProtocolFamily field if non-nil, zero value otherwise.

### GetMigratedToProtocolFamilyOk

`func (o *TokenExtraDTO) GetMigratedToProtocolFamilyOk() (*string, bool)`

GetMigratedToProtocolFamilyOk returns a tuple with the MigratedToProtocolFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedToProtocolFamily

`func (o *TokenExtraDTO) SetMigratedToProtocolFamily(v string)`

SetMigratedToProtocolFamily sets MigratedToProtocolFamily field to given value.

### HasMigratedToProtocolFamily

`func (o *TokenExtraDTO) HasMigratedToProtocolFamily() bool`

HasMigratedToProtocolFamily returns a boolean if a field has been set.

### GetMigratedToPoolAddress

`func (o *TokenExtraDTO) GetMigratedToPoolAddress() string`

GetMigratedToPoolAddress returns the MigratedToPoolAddress field if non-nil, zero value otherwise.

### GetMigratedToPoolAddressOk

`func (o *TokenExtraDTO) GetMigratedToPoolAddressOk() (*string, bool)`

GetMigratedToPoolAddressOk returns a tuple with the MigratedToPoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedToPoolAddress

`func (o *TokenExtraDTO) SetMigratedToPoolAddress(v string)`

SetMigratedToPoolAddress sets MigratedToPoolAddress field to given value.

### HasMigratedToPoolAddress

`func (o *TokenExtraDTO) HasMigratedToPoolAddress() bool`

HasMigratedToPoolAddress returns a boolean if a field has been set.

### GetMigratedAt

`func (o *TokenExtraDTO) GetMigratedAt() int64`

GetMigratedAt returns the MigratedAt field if non-nil, zero value otherwise.

### GetMigratedAtOk

`func (o *TokenExtraDTO) GetMigratedAtOk() (*int64, bool)`

GetMigratedAtOk returns a tuple with the MigratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedAt

`func (o *TokenExtraDTO) SetMigratedAt(v int64)`

SetMigratedAt sets MigratedAt field to given value.

### HasMigratedAt

`func (o *TokenExtraDTO) HasMigratedAt() bool`

HasMigratedAt returns a boolean if a field has been set.

### GetSellerFeeBasisPoints

`func (o *TokenExtraDTO) GetSellerFeeBasisPoints() int64`

GetSellerFeeBasisPoints returns the SellerFeeBasisPoints field if non-nil, zero value otherwise.

### GetSellerFeeBasisPointsOk

`func (o *TokenExtraDTO) GetSellerFeeBasisPointsOk() (*int64, bool)`

GetSellerFeeBasisPointsOk returns a tuple with the SellerFeeBasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerFeeBasisPoints

`func (o *TokenExtraDTO) SetSellerFeeBasisPoints(v int64)`

SetSellerFeeBasisPoints sets SellerFeeBasisPoints field to given value.

### HasSellerFeeBasisPoints

`func (o *TokenExtraDTO) HasSellerFeeBasisPoints() bool`

HasSellerFeeBasisPoints returns a boolean if a field has been set.

### GetTokenStandard

`func (o *TokenExtraDTO) GetTokenStandard() string`

GetTokenStandard returns the TokenStandard field if non-nil, zero value otherwise.

### GetTokenStandardOk

`func (o *TokenExtraDTO) GetTokenStandardOk() (*string, bool)`

GetTokenStandardOk returns a tuple with the TokenStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandard

`func (o *TokenExtraDTO) SetTokenStandard(v string)`

SetTokenStandard sets TokenStandard field to given value.

### HasTokenStandard

`func (o *TokenExtraDTO) HasTokenStandard() bool`

HasTokenStandard returns a boolean if a field has been set.

### GetMintAuthority

`func (o *TokenExtraDTO) GetMintAuthority() string`

GetMintAuthority returns the MintAuthority field if non-nil, zero value otherwise.

### GetMintAuthorityOk

`func (o *TokenExtraDTO) GetMintAuthorityOk() (*string, bool)`

GetMintAuthorityOk returns a tuple with the MintAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintAuthority

`func (o *TokenExtraDTO) SetMintAuthority(v string)`

SetMintAuthority sets MintAuthority field to given value.

### HasMintAuthority

`func (o *TokenExtraDTO) HasMintAuthority() bool`

HasMintAuthority returns a boolean if a field has been set.

### GetFreezeAuthority

`func (o *TokenExtraDTO) GetFreezeAuthority() string`

GetFreezeAuthority returns the FreezeAuthority field if non-nil, zero value otherwise.

### GetFreezeAuthorityOk

`func (o *TokenExtraDTO) GetFreezeAuthorityOk() (*string, bool)`

GetFreezeAuthorityOk returns a tuple with the FreezeAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeAuthority

`func (o *TokenExtraDTO) SetFreezeAuthority(v string)`

SetFreezeAuthority sets FreezeAuthority field to given value.

### HasFreezeAuthority

`func (o *TokenExtraDTO) HasFreezeAuthority() bool`

HasFreezeAuthority returns a boolean if a field has been set.

### GetUpdateAuthority

`func (o *TokenExtraDTO) GetUpdateAuthority() string`

GetUpdateAuthority returns the UpdateAuthority field if non-nil, zero value otherwise.

### GetUpdateAuthorityOk

`func (o *TokenExtraDTO) GetUpdateAuthorityOk() (*string, bool)`

GetUpdateAuthorityOk returns a tuple with the UpdateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthority

`func (o *TokenExtraDTO) SetUpdateAuthority(v string)`

SetUpdateAuthority sets UpdateAuthority field to given value.

### HasUpdateAuthority

`func (o *TokenExtraDTO) HasUpdateAuthority() bool`

HasUpdateAuthority returns a boolean if a field has been set.

### GetIsVerifiedCollection

`func (o *TokenExtraDTO) GetIsVerifiedCollection() bool`

GetIsVerifiedCollection returns the IsVerifiedCollection field if non-nil, zero value otherwise.

### GetIsVerifiedCollectionOk

`func (o *TokenExtraDTO) GetIsVerifiedCollectionOk() (*bool, bool)`

GetIsVerifiedCollectionOk returns a tuple with the IsVerifiedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerifiedCollection

`func (o *TokenExtraDTO) SetIsVerifiedCollection(v bool)`

SetIsVerifiedCollection sets IsVerifiedCollection field to given value.

### HasIsVerifiedCollection

`func (o *TokenExtraDTO) HasIsVerifiedCollection() bool`

HasIsVerifiedCollection returns a boolean if a field has been set.

### GetIsWrapped

`func (o *TokenExtraDTO) GetIsWrapped() bool`

GetIsWrapped returns the IsWrapped field if non-nil, zero value otherwise.

### GetIsWrappedOk

`func (o *TokenExtraDTO) GetIsWrappedOk() (*bool, bool)`

GetIsWrappedOk returns a tuple with the IsWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWrapped

`func (o *TokenExtraDTO) SetIsWrapped(v bool)`

SetIsWrapped sets IsWrapped field to given value.

### HasIsWrapped

`func (o *TokenExtraDTO) HasIsWrapped() bool`

HasIsWrapped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


