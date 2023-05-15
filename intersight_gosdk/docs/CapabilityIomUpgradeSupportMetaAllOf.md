# CapabilityIomUpgradeSupportMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.IomUpgradeSupportMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.IomUpgradeSupportMeta"]
**Description** | Pointer to **string** | Information related to the list of IOMs. Also provides additional information such as hardware name. | [optional] [readonly] 
**DirectUpgrade** | Pointer to **bool** | Indicates if the IOM models have a Device Connector, which in turn allows direct upgrade requests to be sent to the IOM DC. | [optional] [readonly] 
**SeriesId** | Pointer to **string** | Series names of IOMs which will be supported in the firmware operation. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityIomUpgradeSupportMetaAllOf

`func NewCapabilityIomUpgradeSupportMetaAllOf(classId string, objectType string, ) *CapabilityIomUpgradeSupportMetaAllOf`

NewCapabilityIomUpgradeSupportMetaAllOf instantiates a new CapabilityIomUpgradeSupportMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityIomUpgradeSupportMetaAllOfWithDefaults

`func NewCapabilityIomUpgradeSupportMetaAllOfWithDefaults() *CapabilityIomUpgradeSupportMetaAllOf`

NewCapabilityIomUpgradeSupportMetaAllOfWithDefaults instantiates a new CapabilityIomUpgradeSupportMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityIomUpgradeSupportMetaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectUpgrade

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetDirectUpgrade() bool`

GetDirectUpgrade returns the DirectUpgrade field if non-nil, zero value otherwise.

### GetDirectUpgradeOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetDirectUpgradeOk() (*bool, bool)`

GetDirectUpgradeOk returns a tuple with the DirectUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectUpgrade

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetDirectUpgrade(v bool)`

SetDirectUpgrade sets DirectUpgrade field to given value.

### HasDirectUpgrade

`func (o *CapabilityIomUpgradeSupportMetaAllOf) HasDirectUpgrade() bool`

HasDirectUpgrade returns a boolean if a field has been set.

### GetSeriesId

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *CapabilityIomUpgradeSupportMetaAllOf) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityIomUpgradeSupportMetaAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityIomUpgradeSupportMetaAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityIomUpgradeSupportMetaAllOf) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityIomUpgradeSupportMetaAllOf) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


