# FirmwareBaseDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**BundleType** | Pointer to **string** | The bundle type of the image, as published on cisco.com. | [optional] [readonly] 
**ComponentMeta** | Pointer to [**[]FirmwareComponentMeta**](FirmwareComponentMeta.md) |  | [optional] 
**Guid** | Pointer to **string** | The unique identifier for an image in a Cisco repository. | [optional] [readonly] 
**ImageType** | Pointer to **string** | The type of image which the distributable falls into according to the component it can upgrade. For e.g.; Standalone server, Intersight managed server, UCS Managed Fabric Interconnect. The field is used in private appliance mode, where image does not have description populated from CCO. | [optional] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**Model** | Pointer to **string** | The endpoint model for which this firmware image is applicable. | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the image. | [optional] [readonly] 
**RecommendedBuild** | Pointer to **string** | The build which is recommended by Cisco. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The url for the release notes of this image. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] [default to "Cisco"]
**DistributableMetas** | Pointer to [**[]FirmwareDistributableMetaRelationship**](FirmwareDistributableMetaRelationship.md) | An array of relationships to firmwareDistributableMeta resources. | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareBaseDistributable

`func NewFirmwareBaseDistributable(classId string, objectType string, ) *FirmwareBaseDistributable`

NewFirmwareBaseDistributable instantiates a new FirmwareBaseDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaseDistributableWithDefaults

`func NewFirmwareBaseDistributableWithDefaults() *FirmwareBaseDistributable`

NewFirmwareBaseDistributableWithDefaults instantiates a new FirmwareBaseDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareBaseDistributable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareBaseDistributable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareBaseDistributable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareBaseDistributable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareBaseDistributable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareBaseDistributable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBundleType

`func (o *FirmwareBaseDistributable) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *FirmwareBaseDistributable) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *FirmwareBaseDistributable) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *FirmwareBaseDistributable) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetComponentMeta

`func (o *FirmwareBaseDistributable) GetComponentMeta() []FirmwareComponentMeta`

GetComponentMeta returns the ComponentMeta field if non-nil, zero value otherwise.

### GetComponentMetaOk

`func (o *FirmwareBaseDistributable) GetComponentMetaOk() (*[]FirmwareComponentMeta, bool)`

GetComponentMetaOk returns a tuple with the ComponentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMeta

`func (o *FirmwareBaseDistributable) SetComponentMeta(v []FirmwareComponentMeta)`

SetComponentMeta sets ComponentMeta field to given value.

### HasComponentMeta

`func (o *FirmwareBaseDistributable) HasComponentMeta() bool`

HasComponentMeta returns a boolean if a field has been set.

### SetComponentMetaNil

`func (o *FirmwareBaseDistributable) SetComponentMetaNil(b bool)`

 SetComponentMetaNil sets the value for ComponentMeta to be an explicit nil

### UnsetComponentMeta
`func (o *FirmwareBaseDistributable) UnsetComponentMeta()`

UnsetComponentMeta ensures that no value is present for ComponentMeta, not even an explicit nil
### GetGuid

`func (o *FirmwareBaseDistributable) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *FirmwareBaseDistributable) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *FirmwareBaseDistributable) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *FirmwareBaseDistributable) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetImageType

`func (o *FirmwareBaseDistributable) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *FirmwareBaseDistributable) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *FirmwareBaseDistributable) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *FirmwareBaseDistributable) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetMdfid

`func (o *FirmwareBaseDistributable) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *FirmwareBaseDistributable) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *FirmwareBaseDistributable) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *FirmwareBaseDistributable) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetModel

`func (o *FirmwareBaseDistributable) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FirmwareBaseDistributable) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FirmwareBaseDistributable) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FirmwareBaseDistributable) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPlatformType

`func (o *FirmwareBaseDistributable) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *FirmwareBaseDistributable) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *FirmwareBaseDistributable) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *FirmwareBaseDistributable) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRecommendedBuild

`func (o *FirmwareBaseDistributable) GetRecommendedBuild() string`

GetRecommendedBuild returns the RecommendedBuild field if non-nil, zero value otherwise.

### GetRecommendedBuildOk

`func (o *FirmwareBaseDistributable) GetRecommendedBuildOk() (*string, bool)`

GetRecommendedBuildOk returns a tuple with the RecommendedBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedBuild

`func (o *FirmwareBaseDistributable) SetRecommendedBuild(v string)`

SetRecommendedBuild sets RecommendedBuild field to given value.

### HasRecommendedBuild

`func (o *FirmwareBaseDistributable) HasRecommendedBuild() bool`

HasRecommendedBuild returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *FirmwareBaseDistributable) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *FirmwareBaseDistributable) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *FirmwareBaseDistributable) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *FirmwareBaseDistributable) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *FirmwareBaseDistributable) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *FirmwareBaseDistributable) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *FirmwareBaseDistributable) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *FirmwareBaseDistributable) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *FirmwareBaseDistributable) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *FirmwareBaseDistributable) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *FirmwareBaseDistributable) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *FirmwareBaseDistributable) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *FirmwareBaseDistributable) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *FirmwareBaseDistributable) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil
### GetVendor

`func (o *FirmwareBaseDistributable) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareBaseDistributable) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareBaseDistributable) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *FirmwareBaseDistributable) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDistributableMetas

`func (o *FirmwareBaseDistributable) GetDistributableMetas() []FirmwareDistributableMetaRelationship`

GetDistributableMetas returns the DistributableMetas field if non-nil, zero value otherwise.

### GetDistributableMetasOk

`func (o *FirmwareBaseDistributable) GetDistributableMetasOk() (*[]FirmwareDistributableMetaRelationship, bool)`

GetDistributableMetasOk returns a tuple with the DistributableMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributableMetas

`func (o *FirmwareBaseDistributable) SetDistributableMetas(v []FirmwareDistributableMetaRelationship)`

SetDistributableMetas sets DistributableMetas field to given value.

### HasDistributableMetas

`func (o *FirmwareBaseDistributable) HasDistributableMetas() bool`

HasDistributableMetas returns a boolean if a field has been set.

### SetDistributableMetasNil

`func (o *FirmwareBaseDistributable) SetDistributableMetasNil(b bool)`

 SetDistributableMetasNil sets the value for DistributableMetas to be an explicit nil

### UnsetDistributableMetas
`func (o *FirmwareBaseDistributable) UnsetDistributableMetas()`

UnsetDistributableMetas ensures that no value is present for DistributableMetas, not even an explicit nil
### GetRelease

`func (o *FirmwareBaseDistributable) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareBaseDistributable) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareBaseDistributable) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareBaseDistributable) HasRelease() bool`

HasRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


