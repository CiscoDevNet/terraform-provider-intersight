# FabricSwitchControlPolicyAllOf

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

### NewFabricSwitchControlPolicyAllOf

`func NewFabricSwitchControlPolicyAllOf(classId string, objectType string, ) *FabricSwitchControlPolicyAllOf`

NewFabricSwitchControlPolicyAllOf instantiates a new FabricSwitchControlPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchControlPolicyAllOfWithDefaults

`func NewFabricSwitchControlPolicyAllOfWithDefaults() *FabricSwitchControlPolicyAllOf`

NewFabricSwitchControlPolicyAllOfWithDefaults instantiates a new FabricSwitchControlPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchControlPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchControlPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchControlPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchControlPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchControlPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchControlPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAgingSettings

`func (o *FabricSwitchControlPolicyAllOf) GetMacAgingSettings() FabricMacAgingSettings`

GetMacAgingSettings returns the MacAgingSettings field if non-nil, zero value otherwise.

### GetMacAgingSettingsOk

`func (o *FabricSwitchControlPolicyAllOf) GetMacAgingSettingsOk() (*FabricMacAgingSettings, bool)`

GetMacAgingSettingsOk returns a tuple with the MacAgingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAgingSettings

`func (o *FabricSwitchControlPolicyAllOf) SetMacAgingSettings(v FabricMacAgingSettings)`

SetMacAgingSettings sets MacAgingSettings field to given value.

### HasMacAgingSettings

`func (o *FabricSwitchControlPolicyAllOf) HasMacAgingSettings() bool`

HasMacAgingSettings returns a boolean if a field has been set.

### SetMacAgingSettingsNil

`func (o *FabricSwitchControlPolicyAllOf) SetMacAgingSettingsNil(b bool)`

 SetMacAgingSettingsNil sets the value for MacAgingSettings to be an explicit nil

### UnsetMacAgingSettings
`func (o *FabricSwitchControlPolicyAllOf) UnsetMacAgingSettings()`

UnsetMacAgingSettings ensures that no value is present for MacAgingSettings, not even an explicit nil
### GetUdldSettings

`func (o *FabricSwitchControlPolicyAllOf) GetUdldSettings() FabricUdldGlobalSettings`

GetUdldSettings returns the UdldSettings field if non-nil, zero value otherwise.

### GetUdldSettingsOk

`func (o *FabricSwitchControlPolicyAllOf) GetUdldSettingsOk() (*FabricUdldGlobalSettings, bool)`

GetUdldSettingsOk returns a tuple with the UdldSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldSettings

`func (o *FabricSwitchControlPolicyAllOf) SetUdldSettings(v FabricUdldGlobalSettings)`

SetUdldSettings sets UdldSettings field to given value.

### HasUdldSettings

`func (o *FabricSwitchControlPolicyAllOf) HasUdldSettings() bool`

HasUdldSettings returns a boolean if a field has been set.

### SetUdldSettingsNil

`func (o *FabricSwitchControlPolicyAllOf) SetUdldSettingsNil(b bool)`

 SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil

### UnsetUdldSettings
`func (o *FabricSwitchControlPolicyAllOf) UnsetUdldSettings()`

UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
### GetVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicyAllOf) GetVlanPortOptimizationEnabled() bool`

GetVlanPortOptimizationEnabled returns the VlanPortOptimizationEnabled field if non-nil, zero value otherwise.

### GetVlanPortOptimizationEnabledOk

`func (o *FabricSwitchControlPolicyAllOf) GetVlanPortOptimizationEnabledOk() (*bool, bool)`

GetVlanPortOptimizationEnabledOk returns a tuple with the VlanPortOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicyAllOf) SetVlanPortOptimizationEnabled(v bool)`

SetVlanPortOptimizationEnabled sets VlanPortOptimizationEnabled field to given value.

### HasVlanPortOptimizationEnabled

`func (o *FabricSwitchControlPolicyAllOf) HasVlanPortOptimizationEnabled() bool`

HasVlanPortOptimizationEnabled returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchControlPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricSwitchControlPolicyAllOf) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricSwitchControlPolicyAllOf) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricSwitchControlPolicyAllOf) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricSwitchControlPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricSwitchControlPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricSwitchControlPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


