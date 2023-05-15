# CapabilityChassisUpgradeSupportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ChassisUpgradeSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ChassisUpgradeSupportMeta"]
**AdaptersUpgradedViaHsu** | Pointer to **bool** | If enabled, indicates that adapters in servers within this chassis would be upgraded by HSU. | [optional] [readonly] 
**Description** | Pointer to **string** | Verbose description regarding this group of chassis. | [optional] [readonly] 
**SeriesId** | Pointer to **string** | Classification of a set of chassis models. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityChassisUpgradeSupportMeta

`func NewCapabilityChassisUpgradeSupportMeta(classId string, objectType string, ) *CapabilityChassisUpgradeSupportMeta`

NewCapabilityChassisUpgradeSupportMeta instantiates a new CapabilityChassisUpgradeSupportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityChassisUpgradeSupportMetaWithDefaults

`func NewCapabilityChassisUpgradeSupportMetaWithDefaults() *CapabilityChassisUpgradeSupportMeta`

NewCapabilityChassisUpgradeSupportMetaWithDefaults instantiates a new CapabilityChassisUpgradeSupportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityChassisUpgradeSupportMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityChassisUpgradeSupportMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityChassisUpgradeSupportMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityChassisUpgradeSupportMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdaptersUpgradedViaHsu

`func (o *CapabilityChassisUpgradeSupportMeta) GetAdaptersUpgradedViaHsu() bool`

GetAdaptersUpgradedViaHsu returns the AdaptersUpgradedViaHsu field if non-nil, zero value otherwise.

### GetAdaptersUpgradedViaHsuOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetAdaptersUpgradedViaHsuOk() (*bool, bool)`

GetAdaptersUpgradedViaHsuOk returns a tuple with the AdaptersUpgradedViaHsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptersUpgradedViaHsu

`func (o *CapabilityChassisUpgradeSupportMeta) SetAdaptersUpgradedViaHsu(v bool)`

SetAdaptersUpgradedViaHsu sets AdaptersUpgradedViaHsu field to given value.

### HasAdaptersUpgradedViaHsu

`func (o *CapabilityChassisUpgradeSupportMeta) HasAdaptersUpgradedViaHsu() bool`

HasAdaptersUpgradedViaHsu returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityChassisUpgradeSupportMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityChassisUpgradeSupportMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityChassisUpgradeSupportMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeriesId

`func (o *CapabilityChassisUpgradeSupportMeta) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CapabilityChassisUpgradeSupportMeta) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CapabilityChassisUpgradeSupportMeta) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityChassisUpgradeSupportMeta) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityChassisUpgradeSupportMeta) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityChassisUpgradeSupportMeta) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityChassisUpgradeSupportMeta) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityChassisUpgradeSupportMeta) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityChassisUpgradeSupportMeta) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


