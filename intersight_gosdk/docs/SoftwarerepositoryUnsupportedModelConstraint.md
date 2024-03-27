# SoftwarerepositoryUnsupportedModelConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.UnsupportedModelConstraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.UnsupportedModelConstraint"]
**MaxVersion** | Pointer to **string** | UCSM version above which the server model is unsupported. | [optional] 
**Name** | Pointer to **string** | Name of the platform constraint. | [optional] 
**PlatformRegex** | Pointer to **string** | Regular expression of the image name. | [optional] 
**UnsupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryUnsupportedModelConstraint

`func NewSoftwarerepositoryUnsupportedModelConstraint(classId string, objectType string, ) *SoftwarerepositoryUnsupportedModelConstraint`

NewSoftwarerepositoryUnsupportedModelConstraint instantiates a new SoftwarerepositoryUnsupportedModelConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryUnsupportedModelConstraintWithDefaults

`func NewSoftwarerepositoryUnsupportedModelConstraintWithDefaults() *SoftwarerepositoryUnsupportedModelConstraint`

NewSoftwarerepositoryUnsupportedModelConstraintWithDefaults instantiates a new SoftwarerepositoryUnsupportedModelConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxVersion

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *SoftwarerepositoryUnsupportedModelConstraint) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetName

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarerepositoryUnsupportedModelConstraint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformRegex

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetPlatformRegex() string`

GetPlatformRegex returns the PlatformRegex field if non-nil, zero value otherwise.

### GetPlatformRegexOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetPlatformRegexOk() (*string, bool)`

GetPlatformRegexOk returns a tuple with the PlatformRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformRegex

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetPlatformRegex(v string)`

SetPlatformRegex sets PlatformRegex field to given value.

### HasPlatformRegex

`func (o *SoftwarerepositoryUnsupportedModelConstraint) HasPlatformRegex() bool`

HasPlatformRegex returns a boolean if a field has been set.

### GetUnsupportedModels

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetUnsupportedModels() []string`

GetUnsupportedModels returns the UnsupportedModels field if non-nil, zero value otherwise.

### GetUnsupportedModelsOk

`func (o *SoftwarerepositoryUnsupportedModelConstraint) GetUnsupportedModelsOk() (*[]string, bool)`

GetUnsupportedModelsOk returns a tuple with the UnsupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedModels

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetUnsupportedModels(v []string)`

SetUnsupportedModels sets UnsupportedModels field to given value.

### HasUnsupportedModels

`func (o *SoftwarerepositoryUnsupportedModelConstraint) HasUnsupportedModels() bool`

HasUnsupportedModels returns a boolean if a field has been set.

### SetUnsupportedModelsNil

`func (o *SoftwarerepositoryUnsupportedModelConstraint) SetUnsupportedModelsNil(b bool)`

 SetUnsupportedModelsNil sets the value for UnsupportedModels to be an explicit nil

### UnsetUnsupportedModels
`func (o *SoftwarerepositoryUnsupportedModelConstraint) UnsetUnsupportedModels()`

UnsetUnsupportedModels ensures that no value is present for UnsupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


