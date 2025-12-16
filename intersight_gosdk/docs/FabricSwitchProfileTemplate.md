# FabricSwitchProfileTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchProfileTemplate"]
**EnableOverride** | Pointer to **bool** | When enabled, the configuration of the derived instances may override the template configuration. | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**UpdateStatus** | Pointer to **string** | The template sync status with all derived objects. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**UsageCount** | Pointer to **int64** | The number of objects derived from a Template MO instance. | [optional] [readonly] 
**SwitchClusterProfileTemplate** | Pointer to [**NullableFabricSwitchClusterProfileTemplateRelationship**](FabricSwitchClusterProfileTemplateRelationship.md) |  | [optional] 

## Methods

### NewFabricSwitchProfileTemplate

`func NewFabricSwitchProfileTemplate(classId string, objectType string, ) *FabricSwitchProfileTemplate`

NewFabricSwitchProfileTemplate instantiates a new FabricSwitchProfileTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchProfileTemplateWithDefaults

`func NewFabricSwitchProfileTemplateWithDefaults() *FabricSwitchProfileTemplate`

NewFabricSwitchProfileTemplateWithDefaults instantiates a new FabricSwitchProfileTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchProfileTemplate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchProfileTemplate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchProfileTemplate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchProfileTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchProfileTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchProfileTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableOverride

`func (o *FabricSwitchProfileTemplate) GetEnableOverride() bool`

GetEnableOverride returns the EnableOverride field if non-nil, zero value otherwise.

### GetEnableOverrideOk

`func (o *FabricSwitchProfileTemplate) GetEnableOverrideOk() (*bool, bool)`

GetEnableOverrideOk returns a tuple with the EnableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOverride

`func (o *FabricSwitchProfileTemplate) SetEnableOverride(v bool)`

SetEnableOverride sets EnableOverride field to given value.

### HasEnableOverride

`func (o *FabricSwitchProfileTemplate) HasEnableOverride() bool`

HasEnableOverride returns a boolean if a field has been set.

### GetTemplateActions

`func (o *FabricSwitchProfileTemplate) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *FabricSwitchProfileTemplate) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *FabricSwitchProfileTemplate) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *FabricSwitchProfileTemplate) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *FabricSwitchProfileTemplate) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *FabricSwitchProfileTemplate) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetUpdateStatus

`func (o *FabricSwitchProfileTemplate) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *FabricSwitchProfileTemplate) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *FabricSwitchProfileTemplate) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *FabricSwitchProfileTemplate) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetUsageCount

`func (o *FabricSwitchProfileTemplate) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *FabricSwitchProfileTemplate) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *FabricSwitchProfileTemplate) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *FabricSwitchProfileTemplate) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplate) GetSwitchClusterProfileTemplate() FabricSwitchClusterProfileTemplateRelationship`

GetSwitchClusterProfileTemplate returns the SwitchClusterProfileTemplate field if non-nil, zero value otherwise.

### GetSwitchClusterProfileTemplateOk

`func (o *FabricSwitchProfileTemplate) GetSwitchClusterProfileTemplateOk() (*FabricSwitchClusterProfileTemplateRelationship, bool)`

GetSwitchClusterProfileTemplateOk returns a tuple with the SwitchClusterProfileTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplate) SetSwitchClusterProfileTemplate(v FabricSwitchClusterProfileTemplateRelationship)`

SetSwitchClusterProfileTemplate sets SwitchClusterProfileTemplate field to given value.

### HasSwitchClusterProfileTemplate

`func (o *FabricSwitchProfileTemplate) HasSwitchClusterProfileTemplate() bool`

HasSwitchClusterProfileTemplate returns a boolean if a field has been set.

### SetSwitchClusterProfileTemplateNil

`func (o *FabricSwitchProfileTemplate) SetSwitchClusterProfileTemplateNil(b bool)`

 SetSwitchClusterProfileTemplateNil sets the value for SwitchClusterProfileTemplate to be an explicit nil

### UnsetSwitchClusterProfileTemplate
`func (o *FabricSwitchProfileTemplate) UnsetSwitchClusterProfileTemplate()`

UnsetSwitchClusterProfileTemplate ensures that no value is present for SwitchClusterProfileTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


