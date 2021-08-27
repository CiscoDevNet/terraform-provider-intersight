# FabricSwitchControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchControlPolicy"]
**MacAgingSettings** | Pointer to [**NullableFabricMacAgingSettings**](FabricMacAgingSettings.md) |  | [optional] 
**UdldSettings** | Pointer to [**NullableFabricUdldGlobalSettings**](FabricUdldGlobalSettings.md) |  | [optional] 
**VlanPortOptimizationEnabled** | Pointer to **bool** | To enable or disable the VLAN port count optimization. | [optional] [default to false]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchControlPolicy

`func NewFabricSwitchControlPolicy(classId string, objectType string, ) *FabricSwitchControlPolicy`

NewFabricSwitchControlPolicy instantiates a new FabricSwitchControlPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchControlPolicyWithDefaults

`func NewFabricSwitchControlPolicyWithDefaults() *FabricSwitchControlPolicy`

NewFabricSwitchControlPolicyWithDefaults instantiates a new FabricSwitchControlPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchControlPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchControlPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchControlPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchControlPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchControlPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchControlPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAgingSettings

`func (o *FabricSwitchControlPolicy) GetMacAgingSettings() FabricMacAgingSettings`

GetMacAgingSettings returns the MacAgingSettings field if non-nil, zero value otherwise.

### GetMacAgingSettingsOk

`func (o *FabricSwitchControlPolicy) GetMacAgingSettingsOk() (*FabricMacAgingSettings, bool)`

GetMacAgingSettingsOk returns a tuple with the MacAgingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAgingSettings

`func (o *FabricSwitchControlPolicy) SetMacAgingSettings(v FabricMacAgingSettings)`

SetMacAgingSettings sets MacAgingSettings field to given value.

### HasMacAgingSettings

`func (o *FabricSwitchControlPolicy) HasMacAgingSettings() bool`

HasMacAgingSettings returns a boolean if a field has been set.

### SetMacAgingSettingsNil

`func (o *FabricSwitchControlPolicy) SetMacAgingSettingsNil(b bool)`

 SetMacAgingSettingsNil sets the value for MacAgingSettings to be an explicit nil

### UnsetMacAgingSettings
`func (o *FabricSwitchControlPolicy) UnsetMacAgingSettings()`

UnsetMacAgingSettings ensures that no value is present for MacAgingSettings, not even an explicit nil
### GetUdldSettings

`func (o *FabricSwitchControlPolicy) GetUdldSettings() FabricUdldGlobalSettings`

GetUdldSettings returns the UdldSettings field if non-nil, zero value otherwise.

### GetUdldSettingsOk

`func (o *FabricSwitchControlPolicy) GetUdldSettingsOk() (*FabricUdldGlobalSettings, bool)`

GetUdldSettingsOk returns a tuple with the UdldSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldSettings

`func (o *FabricSwitchControlPolicy) SetUdldSettings(v FabricUdldGlobalSettings)`

SetUdldSettings sets UdldSettings field to given value.

### HasUdldSettings

`func (o *FabricSwitchControlPolicy) HasUdldSettings() bool`

HasUdldSettings returns a boolean if a field has been set.

### SetUdldSettingsNil

`func (o *FabricSwitchControlPolicy) SetUdldSettingsNil(b bool)`

 SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil

### UnsetUdldSettings
`func (o *FabricSwitchControlPolicy) UnsetUdldSettings()`

UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
### GetVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicy) GetVlanPortOptimizationEnabled() bool`

GetVlanPortOptimizationEnabled returns the VlanPortOptimizationEnabled field if non-nil, zero value otherwise.

### GetVlanPortOptimizationEnabledOk

`func (o *FabricSwitchControlPolicy) GetVlanPortOptimizationEnabledOk() (*bool, bool)`

GetVlanPortOptimizationEnabledOk returns a tuple with the VlanPortOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicy) SetVlanPortOptimizationEnabled(v bool)`

SetVlanPortOptimizationEnabled sets VlanPortOptimizationEnabled field to given value.

### HasVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicy) HasVlanPortOptimizationEnabled() bool`

HasVlanPortOptimizationEnabled returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchControlPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchControlPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchControlPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricSwitchControlPolicy) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricSwitchControlPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricSwitchControlPolicy) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricSwitchControlPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricSwitchControlPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricSwitchControlPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


