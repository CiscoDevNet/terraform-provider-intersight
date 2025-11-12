# SoftwarerepositoryBaseDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "softwarerepository.OperatingSystemFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "softwarerepository.OperatingSystemFile"]
**BundleType** | Pointer to **string** | The bundle type of the image, as published on cisco.com. | [optional] [readonly] 
**Guid** | Pointer to **string** | The unique identifier for an image in a Cisco repository. | [optional] [readonly] 
**ImageType** | Pointer to **string** | The type of image which the distributable falls into according to the component it can upgrade. For e.g.; Standalone server, Intersight managed server, Unified Edge server. The field is used in private appliance mode, where image does not have description populated from CCO. | [optional] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**Model** | Pointer to **string** | The endpoint model for which this firmware image is applicable. | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the image. | [optional] [readonly] 
**RecommendedBuild** | Pointer to **string** | The build which is recommended by Cisco. | [optional] 
**ReleaseNotesUrl** | Pointer to **string** | The url for the release notes of this image. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] [readonly] 
**Release** | Pointer to [**NullableSoftwarerepositoryReleaseRelationship**](SoftwarerepositoryReleaseRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryBaseDistributable

`func NewSoftwarerepositoryBaseDistributable(classId string, objectType string, ) *SoftwarerepositoryBaseDistributable`

NewSoftwarerepositoryBaseDistributable instantiates a new SoftwarerepositoryBaseDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryBaseDistributableWithDefaults

`func NewSoftwarerepositoryBaseDistributableWithDefaults() *SoftwarerepositoryBaseDistributable`

NewSoftwarerepositoryBaseDistributableWithDefaults instantiates a new SoftwarerepositoryBaseDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryBaseDistributable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryBaseDistributable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryBaseDistributable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryBaseDistributable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryBaseDistributable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryBaseDistributable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBundleType

`func (o *SoftwarerepositoryBaseDistributable) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *SoftwarerepositoryBaseDistributable) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *SoftwarerepositoryBaseDistributable) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *SoftwarerepositoryBaseDistributable) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetGuid

`func (o *SoftwarerepositoryBaseDistributable) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SoftwarerepositoryBaseDistributable) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SoftwarerepositoryBaseDistributable) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SoftwarerepositoryBaseDistributable) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetImageType

`func (o *SoftwarerepositoryBaseDistributable) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *SoftwarerepositoryBaseDistributable) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *SoftwarerepositoryBaseDistributable) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *SoftwarerepositoryBaseDistributable) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetMdfid

`func (o *SoftwarerepositoryBaseDistributable) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *SoftwarerepositoryBaseDistributable) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *SoftwarerepositoryBaseDistributable) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *SoftwarerepositoryBaseDistributable) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetModel

`func (o *SoftwarerepositoryBaseDistributable) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SoftwarerepositoryBaseDistributable) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SoftwarerepositoryBaseDistributable) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SoftwarerepositoryBaseDistributable) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPlatformType

`func (o *SoftwarerepositoryBaseDistributable) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *SoftwarerepositoryBaseDistributable) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *SoftwarerepositoryBaseDistributable) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *SoftwarerepositoryBaseDistributable) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRecommendedBuild

`func (o *SoftwarerepositoryBaseDistributable) GetRecommendedBuild() string`

GetRecommendedBuild returns the RecommendedBuild field if non-nil, zero value otherwise.

### GetRecommendedBuildOk

`func (o *SoftwarerepositoryBaseDistributable) GetRecommendedBuildOk() (*string, bool)`

GetRecommendedBuildOk returns a tuple with the RecommendedBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedBuild

`func (o *SoftwarerepositoryBaseDistributable) SetRecommendedBuild(v string)`

SetRecommendedBuild sets RecommendedBuild field to given value.

### HasRecommendedBuild

`func (o *SoftwarerepositoryBaseDistributable) HasRecommendedBuild() bool`

HasRecommendedBuild returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *SoftwarerepositoryBaseDistributable) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *SoftwarerepositoryBaseDistributable) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *SoftwarerepositoryBaseDistributable) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.

### HasReleaseNotesUrl

`func (o *SoftwarerepositoryBaseDistributable) HasReleaseNotesUrl() bool`

HasReleaseNotesUrl returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *SoftwarerepositoryBaseDistributable) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *SoftwarerepositoryBaseDistributable) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *SoftwarerepositoryBaseDistributable) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *SoftwarerepositoryBaseDistributable) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetRelease

`func (o *SoftwarerepositoryBaseDistributable) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SoftwarerepositoryBaseDistributable) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SoftwarerepositoryBaseDistributable) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *SoftwarerepositoryBaseDistributable) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *SoftwarerepositoryBaseDistributable) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *SoftwarerepositoryBaseDistributable) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


