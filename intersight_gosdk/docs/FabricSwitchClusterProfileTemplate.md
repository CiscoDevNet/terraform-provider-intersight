# FabricSwitchClusterProfileTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchClusterProfileTemplate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchClusterProfileTemplate"]
**EnableOverride** | Pointer to **bool** | When enabled, the configuration of the derived instances may override the template configuration. | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**UpdateStatus** | Pointer to **string** | The template sync status with all derived objects. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**Usage** | Pointer to **int64** | The count of switch cluster profiles derived from the template. | [optional] [readonly] [default to 0]
**UsageCount** | Pointer to **int64** | The number of objects derived from a Template MO instance. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SwitchProfileTemplates** | Pointer to [**[]FabricSwitchProfileTemplateRelationship**](FabricSwitchProfileTemplateRelationship.md) | An array of relationships to fabricSwitchProfileTemplate resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfileTemplate

`func NewFabricSwitchClusterProfileTemplate(classId string, objectType string, ) *FabricSwitchClusterProfileTemplate`

NewFabricSwitchClusterProfileTemplate instantiates a new FabricSwitchClusterProfileTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileTemplateWithDefaults

`func NewFabricSwitchClusterProfileTemplateWithDefaults() *FabricSwitchClusterProfileTemplate`

NewFabricSwitchClusterProfileTemplateWithDefaults instantiates a new FabricSwitchClusterProfileTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchClusterProfileTemplate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfileTemplate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfileTemplate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchClusterProfileTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfileTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfileTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableOverride

`func (o *FabricSwitchClusterProfileTemplate) GetEnableOverride() bool`

GetEnableOverride returns the EnableOverride field if non-nil, zero value otherwise.

### GetEnableOverrideOk

`func (o *FabricSwitchClusterProfileTemplate) GetEnableOverrideOk() (*bool, bool)`

GetEnableOverrideOk returns a tuple with the EnableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOverride

`func (o *FabricSwitchClusterProfileTemplate) SetEnableOverride(v bool)`

SetEnableOverride sets EnableOverride field to given value.

### HasEnableOverride

`func (o *FabricSwitchClusterProfileTemplate) HasEnableOverride() bool`

HasEnableOverride returns a boolean if a field has been set.

### GetTemplateActions

`func (o *FabricSwitchClusterProfileTemplate) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *FabricSwitchClusterProfileTemplate) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *FabricSwitchClusterProfileTemplate) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *FabricSwitchClusterProfileTemplate) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *FabricSwitchClusterProfileTemplate) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *FabricSwitchClusterProfileTemplate) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetUpdateStatus

`func (o *FabricSwitchClusterProfileTemplate) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *FabricSwitchClusterProfileTemplate) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *FabricSwitchClusterProfileTemplate) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *FabricSwitchClusterProfileTemplate) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetUsage

`func (o *FabricSwitchClusterProfileTemplate) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FabricSwitchClusterProfileTemplate) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FabricSwitchClusterProfileTemplate) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FabricSwitchClusterProfileTemplate) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsageCount

`func (o *FabricSwitchClusterProfileTemplate) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *FabricSwitchClusterProfileTemplate) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *FabricSwitchClusterProfileTemplate) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *FabricSwitchClusterProfileTemplate) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfileTemplate) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfileTemplate) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfileTemplate) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfileTemplate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FabricSwitchClusterProfileTemplate) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FabricSwitchClusterProfileTemplate) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplate) GetSwitchProfileTemplates() []FabricSwitchProfileTemplateRelationship`

GetSwitchProfileTemplates returns the SwitchProfileTemplates field if non-nil, zero value otherwise.

### GetSwitchProfileTemplatesOk

`func (o *FabricSwitchClusterProfileTemplate) GetSwitchProfileTemplatesOk() (*[]FabricSwitchProfileTemplateRelationship, bool)`

GetSwitchProfileTemplatesOk returns a tuple with the SwitchProfileTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplate) SetSwitchProfileTemplates(v []FabricSwitchProfileTemplateRelationship)`

SetSwitchProfileTemplates sets SwitchProfileTemplates field to given value.

### HasSwitchProfileTemplates

`func (o *FabricSwitchClusterProfileTemplate) HasSwitchProfileTemplates() bool`

HasSwitchProfileTemplates returns a boolean if a field has been set.

### SetSwitchProfileTemplatesNil

`func (o *FabricSwitchClusterProfileTemplate) SetSwitchProfileTemplatesNil(b bool)`

 SetSwitchProfileTemplatesNil sets the value for SwitchProfileTemplates to be an explicit nil

### UnsetSwitchProfileTemplates
`func (o *FabricSwitchClusterProfileTemplate) UnsetSwitchProfileTemplates()`

UnsetSwitchProfileTemplates ensures that no value is present for SwitchProfileTemplates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


