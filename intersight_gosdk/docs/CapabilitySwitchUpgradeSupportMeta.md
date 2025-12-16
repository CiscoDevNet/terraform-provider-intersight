# CapabilitySwitchUpgradeSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchUpgradeSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchUpgradeSupportMeta"]
**Description** | Pointer to **string** | Information related to the list of Component models. | [optional] [readonly] 
**MaxVersion** | Pointer to **string** | Maximum Fabric Interconnect version for the component to be supported. | [optional] [readonly] 
**MinVersion** | Pointer to **string** | Minimum Fabric Interconnect version for the component to be supported. | [optional] [readonly] 
**SeriesId** | Pointer to **string** | Series names of Component. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilitySwitchUpgradeSupportMeta

`func NewCapabilitySwitchUpgradeSupportMeta(classId string, objectType string, ) *CapabilitySwitchUpgradeSupportMeta`

NewCapabilitySwitchUpgradeSupportMeta instantiates a new CapabilitySwitchUpgradeSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchUpgradeSupportMetaWithDefaults

`func NewCapabilitySwitchUpgradeSupportMetaWithDefaults() *CapabilitySwitchUpgradeSupportMeta`

NewCapabilitySwitchUpgradeSupportMetaWithDefaults instantiates a new CapabilitySwitchUpgradeSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchUpgradeSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchUpgradeSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchUpgradeSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchUpgradeSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilitySwitchUpgradeSupportMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilitySwitchUpgradeSupportMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilitySwitchUpgradeSupportMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetMinVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *CapabilitySwitchUpgradeSupportMeta) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetSeriesId

`func (o *CapabilitySwitchUpgradeSupportMeta) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CapabilitySwitchUpgradeSupportMeta) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CapabilitySwitchUpgradeSupportMeta) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilitySwitchUpgradeSupportMeta) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilitySwitchUpgradeSupportMeta) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilitySwitchUpgradeSupportMeta) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilitySwitchUpgradeSupportMeta) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilitySwitchUpgradeSupportMeta) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilitySwitchUpgradeSupportMeta) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


