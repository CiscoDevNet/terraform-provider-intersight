# CapabilityServerUpgradeSupportMetaAllOf

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

### NewCapabilityServerUpgradeSupportMetaAllOf

`func NewCapabilityServerUpgradeSupportMetaAllOf(classId string, objectType string, ) *CapabilityServerUpgradeSupportMetaAllOf`

NewCapabilityServerUpgradeSupportMetaAllOf instantiates a new CapabilityServerUpgradeSupportMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerUpgradeSupportMetaAllOfWithDefaults

`func NewCapabilityServerUpgradeSupportMetaAllOfWithDefaults() *CapabilityServerUpgradeSupportMetaAllOf`

NewCapabilityServerUpgradeSupportMetaAllOfWithDefaults instantiates a new CapabilityServerUpgradeSupportMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityServerUpgradeSupportMetaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CapabilityServerUpgradeSupportMetaAllOf) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetServerFamily

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetServerFamily() string`

GetServerFamily returns the ServerFamily field if non-nil, zero value otherwise.

### GetServerFamilyOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetServerFamilyOk() (*string, bool)`

GetServerFamilyOk returns a tuple with the ServerFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFamily

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetServerFamily(v string)`

SetServerFamily sets ServerFamily field to given value.

### HasServerFamily

`func (o *CapabilityServerUpgradeSupportMetaAllOf) HasServerFamily() bool`

HasServerFamily returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityServerUpgradeSupportMetaAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityServerUpgradeSupportMetaAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityServerUpgradeSupportMetaAllOf) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityServerUpgradeSupportMetaAllOf) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


