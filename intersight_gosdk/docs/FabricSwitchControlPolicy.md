# FabricSwitchControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchControlPolicy"]
**AesPrimaryKey** | Pointer to **string** | Encrypts MACsec keys in type-6 format. If a MACsec key is already provided in a type-6 format, the primary key decrypts it. | [optional] 
**EthernetSwitchingMode** | Pointer to **string** | Enable or Disable Ethernet End Host Switching Mode. * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [default to "end-host"]
**FabricPcVhbaReset** | Pointer to **string** | When enabled, a Registered State Change Notification (RSCN) is sent to the VIC adapter when any member port within the fabric port-channel goes down and vHBA would reset to restore the connection immediately. When disabled (default), vHBA reset is done only when all the members of a fabric port-channel are down. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**FcSwitchingMode** | Pointer to **string** | Enable or Disable FC End Host Switching Mode. * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [default to "end-host"]
**IsAesPrimaryKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;aesPrimaryKey&#39; property has been set. | [optional] [readonly] [default to false]
**MacAgingSettings** | Pointer to [**NullableFabricMacAgingSettings**](FabricMacAgingSettings.md) |  | [optional] 
**ReservedVlanStartId** | Pointer to **int64** | The starting ID for VLANs reserved for internal use within the Fabric Interconnect. This VLAN ID is the starting ID of a contiguous block of 128 VLANs that cannot be configured for user data.  This range of VLANs cannot be configured in VLAN policy. If this property is not configured, VLAN range 3915 - 4042 is reserved for internal use by default. | [optional] [default to 3915]
**UdldSettings** | Pointer to [**NullableFabricUdldGlobalSettings**](FabricUdldGlobalSettings.md) |  | [optional] 
**VlanPortOptimizationEnabled** | Pointer to **bool** | To enable or disable the VLAN port count optimization. This feature will always be enabled for Cisco UCS Fabric Interconnect 9108 100G. | [optional] [default to false]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricBaseSwitchProfileRelationship**](FabricBaseSwitchProfileRelationship.md) | An array of relationships to fabricBaseSwitchProfile resources. | [optional] 

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


### GetAesPrimaryKey

`func (o *FabricSwitchControlPolicy) GetAesPrimaryKey() string`

GetAesPrimaryKey returns the AesPrimaryKey field if non-nil, zero value otherwise.

### GetAesPrimaryKeyOk

`func (o *FabricSwitchControlPolicy) GetAesPrimaryKeyOk() (*string, bool)`

GetAesPrimaryKeyOk returns a tuple with the AesPrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAesPrimaryKey

`func (o *FabricSwitchControlPolicy) SetAesPrimaryKey(v string)`

SetAesPrimaryKey sets AesPrimaryKey field to given value.

### HasAesPrimaryKey

`func (o *FabricSwitchControlPolicy) HasAesPrimaryKey() bool`

HasAesPrimaryKey returns a boolean if a field has been set.

### GetEthernetSwitchingMode

`func (o *FabricSwitchControlPolicy) GetEthernetSwitchingMode() string`

GetEthernetSwitchingMode returns the EthernetSwitchingMode field if non-nil, zero value otherwise.

### GetEthernetSwitchingModeOk

`func (o *FabricSwitchControlPolicy) GetEthernetSwitchingModeOk() (*string, bool)`

GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetSwitchingMode

`func (o *FabricSwitchControlPolicy) SetEthernetSwitchingMode(v string)`

SetEthernetSwitchingMode sets EthernetSwitchingMode field to given value.

### HasEthernetSwitchingMode

`func (o *FabricSwitchControlPolicy) HasEthernetSwitchingMode() bool`

HasEthernetSwitchingMode returns a boolean if a field has been set.

### GetFabricPcVhbaReset

`func (o *FabricSwitchControlPolicy) GetFabricPcVhbaReset() string`

GetFabricPcVhbaReset returns the FabricPcVhbaReset field if non-nil, zero value otherwise.

### GetFabricPcVhbaResetOk

`func (o *FabricSwitchControlPolicy) GetFabricPcVhbaResetOk() (*string, bool)`

GetFabricPcVhbaResetOk returns a tuple with the FabricPcVhbaReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricPcVhbaReset

`func (o *FabricSwitchControlPolicy) SetFabricPcVhbaReset(v string)`

SetFabricPcVhbaReset sets FabricPcVhbaReset field to given value.

### HasFabricPcVhbaReset

`func (o *FabricSwitchControlPolicy) HasFabricPcVhbaReset() bool`

HasFabricPcVhbaReset returns a boolean if a field has been set.

### GetFcSwitchingMode

`func (o *FabricSwitchControlPolicy) GetFcSwitchingMode() string`

GetFcSwitchingMode returns the FcSwitchingMode field if non-nil, zero value otherwise.

### GetFcSwitchingModeOk

`func (o *FabricSwitchControlPolicy) GetFcSwitchingModeOk() (*string, bool)`

GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcSwitchingMode

`func (o *FabricSwitchControlPolicy) SetFcSwitchingMode(v string)`

SetFcSwitchingMode sets FcSwitchingMode field to given value.

### HasFcSwitchingMode

`func (o *FabricSwitchControlPolicy) HasFcSwitchingMode() bool`

HasFcSwitchingMode returns a boolean if a field has been set.

### GetIsAesPrimaryKeySet

`func (o *FabricSwitchControlPolicy) GetIsAesPrimaryKeySet() bool`

GetIsAesPrimaryKeySet returns the IsAesPrimaryKeySet field if non-nil, zero value otherwise.

### GetIsAesPrimaryKeySetOk

`func (o *FabricSwitchControlPolicy) GetIsAesPrimaryKeySetOk() (*bool, bool)`

GetIsAesPrimaryKeySetOk returns a tuple with the IsAesPrimaryKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAesPrimaryKeySet

`func (o *FabricSwitchControlPolicy) SetIsAesPrimaryKeySet(v bool)`

SetIsAesPrimaryKeySet sets IsAesPrimaryKeySet field to given value.

### HasIsAesPrimaryKeySet

`func (o *FabricSwitchControlPolicy) HasIsAesPrimaryKeySet() bool`

HasIsAesPrimaryKeySet returns a boolean if a field has been set.

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
### GetReservedVlanStartId

`func (o *FabricSwitchControlPolicy) GetReservedVlanStartId() int64`

GetReservedVlanStartId returns the ReservedVlanStartId field if non-nil, zero value otherwise.

### GetReservedVlanStartIdOk

`func (o *FabricSwitchControlPolicy) GetReservedVlanStartIdOk() (*int64, bool)`

GetReservedVlanStartIdOk returns a tuple with the ReservedVlanStartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlanStartId

`func (o *FabricSwitchControlPolicy) SetReservedVlanStartId(v int64)`

SetReservedVlanStartId sets ReservedVlanStartId field to given value.

### HasReservedVlanStartId

`func (o *FabricSwitchControlPolicy) HasReservedVlanStartId() bool`

HasReservedVlanStartId returns a boolean if a field has been set.

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

### SetOrganizationNil

`func (o *FabricSwitchControlPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FabricSwitchControlPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *FabricSwitchControlPolicy) GetProfiles() []FabricBaseSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricSwitchControlPolicy) GetProfilesOk() (*[]FabricBaseSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricSwitchControlPolicy) SetProfiles(v []FabricBaseSwitchProfileRelationship)`

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


