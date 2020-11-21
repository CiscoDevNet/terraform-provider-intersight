# SoftwarerepositoryConstraintModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.ConstraintModels"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.ConstraintModels"]
**MinVersion** | Pointer to **string** | Minimum version of the image. | [optional] 
**Name** | Pointer to **string** | Name of the contraint, used to identify constriant type. | [optional] 
**PlatformRegex** | Pointer to **string** | Regular expression of the image name. | [optional] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryConstraintModels

`func NewSoftwarerepositoryConstraintModels(classId string, objectType string, ) *SoftwarerepositoryConstraintModels`

NewSoftwarerepositoryConstraintModels instantiates a new SoftwarerepositoryConstraintModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryConstraintModelsWithDefaults

`func NewSoftwarerepositoryConstraintModelsWithDefaults() *SoftwarerepositoryConstraintModels`

NewSoftwarerepositoryConstraintModelsWithDefaults instantiates a new SoftwarerepositoryConstraintModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryConstraintModels) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryConstraintModels) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryConstraintModels) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryConstraintModels) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryConstraintModels) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryConstraintModels) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMinVersion

`func (o *SoftwarerepositoryConstraintModels) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *SoftwarerepositoryConstraintModels) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *SoftwarerepositoryConstraintModels) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *SoftwarerepositoryConstraintModels) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryConstraintModels) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryConstraintModels) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryConstraintModels) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryConstraintModels) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformRegex

`func (o *SoftwarerepositoryConstraintModels) GetPlatformRegex() string`

GetPlatformRegex returns the PlatformRegex field if non-nil, zero value otherwise.

### GetPlatformRegexOk

`func (o *SoftwarerepositoryConstraintModels) GetPlatformRegexOk() (*string, bool)`

GetPlatformRegexOk returns a tuple with the PlatformRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformRegex

`func (o *SoftwarerepositoryConstraintModels) SetPlatformRegex(v string)`

SetPlatformRegex sets PlatformRegex field to given value.

### HasPlatformRegex

`func (o *SoftwarerepositoryConstraintModels) HasPlatformRegex() bool`

HasPlatformRegex returns a boolean if a field has been set.

### GetSupportedModels

`func (o *SoftwarerepositoryConstraintModels) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *SoftwarerepositoryConstraintModels) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *SoftwarerepositoryConstraintModels) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *SoftwarerepositoryConstraintModels) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *SoftwarerepositoryConstraintModels) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *SoftwarerepositoryConstraintModels) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


