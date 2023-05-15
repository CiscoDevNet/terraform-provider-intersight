# CapabilityServerUpgradeSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerUpgradeSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerUpgradeSupportMeta"]
**Description** | Pointer to **string** | Verbose description regarding this group of server. | [optional] [readonly] 
**Platform** | Pointer to **string** | The target platform for which the mapping is defined. | [optional] [readonly] 
**ServerFamily** | Pointer to **string** | Classification of a set of server models. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityServerUpgradeSupportMeta

`func NewCapabilityServerUpgradeSupportMeta(classId string, objectType string, ) *CapabilityServerUpgradeSupportMeta`

NewCapabilityServerUpgradeSupportMeta instantiates a new CapabilityServerUpgradeSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerUpgradeSupportMetaWithDefaults

`func NewCapabilityServerUpgradeSupportMetaWithDefaults() *CapabilityServerUpgradeSupportMeta`

NewCapabilityServerUpgradeSupportMetaWithDefaults instantiates a new CapabilityServerUpgradeSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerUpgradeSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerUpgradeSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerUpgradeSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerUpgradeSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerUpgradeSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerUpgradeSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityServerUpgradeSupportMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityServerUpgradeSupportMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityServerUpgradeSupportMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityServerUpgradeSupportMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *CapabilityServerUpgradeSupportMeta) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CapabilityServerUpgradeSupportMeta) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CapabilityServerUpgradeSupportMeta) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CapabilityServerUpgradeSupportMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetServerFamily

`func (o *CapabilityServerUpgradeSupportMeta) GetServerFamily() string`

GetServerFamily returns the ServerFamily field if non-nil, zero value otherwise.

### GetServerFamilyOk

`func (o *CapabilityServerUpgradeSupportMeta) GetServerFamilyOk() (*string, bool)`

GetServerFamilyOk returns a tuple with the ServerFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFamily

`func (o *CapabilityServerUpgradeSupportMeta) SetServerFamily(v string)`

SetServerFamily sets ServerFamily field to given value.

### HasServerFamily

`func (o *CapabilityServerUpgradeSupportMeta) HasServerFamily() bool`

HasServerFamily returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityServerUpgradeSupportMeta) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityServerUpgradeSupportMeta) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityServerUpgradeSupportMeta) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityServerUpgradeSupportMeta) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityServerUpgradeSupportMeta) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityServerUpgradeSupportMeta) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


