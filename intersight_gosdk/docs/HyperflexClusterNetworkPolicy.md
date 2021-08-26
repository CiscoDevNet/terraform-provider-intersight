# HyperflexClusterNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterNetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterNetworkPolicy"]
**JumboFrame** | Pointer to **bool** | Enable or disable jumbo frames. | [optional] 
**KvmIpRange** | Pointer to [**NullableHyperflexIpAddrRange**](HyperflexIpAddrRange.md) |  | [optional] 
**MacPrefixRange** | Pointer to [**NullableHyperflexMacAddrPrefixRange**](HyperflexMacAddrPrefixRange.md) |  | [optional] 
**MgmtVlan** | Pointer to [**NullableHyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**UplinkSpeed** | Pointer to **string** | Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be &#39;1G&#39; or &#39;10G+&#39;. Use &#39;10G+&#39; for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be &#39;default&#39; only. * &#x60;default&#x60; - Current default value set on the hardware platform. * &#x60;1G&#x60; - A link speed of 1 gigabit per second. * &#x60;10G&#x60; - A link speed of 10 gigabits per second or above. | [optional] [default to "default"]
**VmMigrationVlan** | Pointer to [**NullableHyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**VmNetworkVlans** | Pointer to [**[]HyperflexNamedVlan**](HyperflexNamedVlan.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterNetworkPolicy

`func NewHyperflexClusterNetworkPolicy(classId string, objectType string, ) *HyperflexClusterNetworkPolicy`

NewHyperflexClusterNetworkPolicy instantiates a new HyperflexClusterNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterNetworkPolicyWithDefaults

`func NewHyperflexClusterNetworkPolicyWithDefaults() *HyperflexClusterNetworkPolicy`

NewHyperflexClusterNetworkPolicyWithDefaults instantiates a new HyperflexClusterNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterNetworkPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterNetworkPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterNetworkPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterNetworkPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterNetworkPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterNetworkPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetJumboFrame

`func (o *HyperflexClusterNetworkPolicy) GetJumboFrame() bool`

GetJumboFrame returns the JumboFrame field if non-nil, zero value otherwise.

### GetJumboFrameOk

`func (o *HyperflexClusterNetworkPolicy) GetJumboFrameOk() (*bool, bool)`

GetJumboFrameOk returns a tuple with the JumboFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboFrame

`func (o *HyperflexClusterNetworkPolicy) SetJumboFrame(v bool)`

SetJumboFrame sets JumboFrame field to given value.

### HasJumboFrame

`func (o *HyperflexClusterNetworkPolicy) HasJumboFrame() bool`

HasJumboFrame returns a boolean if a field has been set.

### GetKvmIpRange

`func (o *HyperflexClusterNetworkPolicy) GetKvmIpRange() HyperflexIpAddrRange`

GetKvmIpRange returns the KvmIpRange field if non-nil, zero value otherwise.

### GetKvmIpRangeOk

`func (o *HyperflexClusterNetworkPolicy) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetKvmIpRangeOk returns a tuple with the KvmIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpRange

`func (o *HyperflexClusterNetworkPolicy) SetKvmIpRange(v HyperflexIpAddrRange)`

SetKvmIpRange sets KvmIpRange field to given value.

### HasKvmIpRange

`func (o *HyperflexClusterNetworkPolicy) HasKvmIpRange() bool`

HasKvmIpRange returns a boolean if a field has been set.

### SetKvmIpRangeNil

`func (o *HyperflexClusterNetworkPolicy) SetKvmIpRangeNil(b bool)`

 SetKvmIpRangeNil sets the value for KvmIpRange to be an explicit nil

### UnsetKvmIpRange
`func (o *HyperflexClusterNetworkPolicy) UnsetKvmIpRange()`

UnsetKvmIpRange ensures that no value is present for KvmIpRange, not even an explicit nil
### GetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRange() HyperflexMacAddrPrefixRange`

GetMacPrefixRange returns the MacPrefixRange field if non-nil, zero value otherwise.

### GetMacPrefixRangeOk

`func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool)`

GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRange(v HyperflexMacAddrPrefixRange)`

SetMacPrefixRange sets MacPrefixRange field to given value.

### HasMacPrefixRange

`func (o *HyperflexClusterNetworkPolicy) HasMacPrefixRange() bool`

HasMacPrefixRange returns a boolean if a field has been set.

### SetMacPrefixRangeNil

`func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRangeNil(b bool)`

 SetMacPrefixRangeNil sets the value for MacPrefixRange to be an explicit nil

### UnsetMacPrefixRange
`func (o *HyperflexClusterNetworkPolicy) UnsetMacPrefixRange()`

UnsetMacPrefixRange ensures that no value is present for MacPrefixRange, not even an explicit nil
### GetMgmtVlan

`func (o *HyperflexClusterNetworkPolicy) GetMgmtVlan() HyperflexNamedVlan`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *HyperflexClusterNetworkPolicy) GetMgmtVlanOk() (*HyperflexNamedVlan, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *HyperflexClusterNetworkPolicy) SetMgmtVlan(v HyperflexNamedVlan)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *HyperflexClusterNetworkPolicy) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### SetMgmtVlanNil

`func (o *HyperflexClusterNetworkPolicy) SetMgmtVlanNil(b bool)`

 SetMgmtVlanNil sets the value for MgmtVlan to be an explicit nil

### UnsetMgmtVlan
`func (o *HyperflexClusterNetworkPolicy) UnsetMgmtVlan()`

UnsetMgmtVlan ensures that no value is present for MgmtVlan, not even an explicit nil
### GetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeed() string`

GetUplinkSpeed returns the UplinkSpeed field if non-nil, zero value otherwise.

### GetUplinkSpeedOk

`func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeedOk() (*string, bool)`

GetUplinkSpeedOk returns a tuple with the UplinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicy) SetUplinkSpeed(v string)`

SetUplinkSpeed sets UplinkSpeed field to given value.

### HasUplinkSpeed

`func (o *HyperflexClusterNetworkPolicy) HasUplinkSpeed() bool`

HasUplinkSpeed returns a boolean if a field has been set.

### GetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlan() HyperflexNamedVlan`

GetVmMigrationVlan returns the VmMigrationVlan field if non-nil, zero value otherwise.

### GetVmMigrationVlanOk

`func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlanOk() (*HyperflexNamedVlan, bool)`

GetVmMigrationVlanOk returns a tuple with the VmMigrationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlan(v HyperflexNamedVlan)`

SetVmMigrationVlan sets VmMigrationVlan field to given value.

### HasVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicy) HasVmMigrationVlan() bool`

HasVmMigrationVlan returns a boolean if a field has been set.

### SetVmMigrationVlanNil

`func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlanNil(b bool)`

 SetVmMigrationVlanNil sets the value for VmMigrationVlan to be an explicit nil

### UnsetVmMigrationVlan
`func (o *HyperflexClusterNetworkPolicy) UnsetVmMigrationVlan()`

UnsetVmMigrationVlan ensures that no value is present for VmMigrationVlan, not even an explicit nil
### GetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlans() []HyperflexNamedVlan`

GetVmNetworkVlans returns the VmNetworkVlans field if non-nil, zero value otherwise.

### GetVmNetworkVlansOk

`func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlansOk() (*[]HyperflexNamedVlan, bool)`

GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicy) SetVmNetworkVlans(v []HyperflexNamedVlan)`

SetVmNetworkVlans sets VmNetworkVlans field to given value.

### HasVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicy) HasVmNetworkVlans() bool`

HasVmNetworkVlans returns a boolean if a field has been set.

### SetVmNetworkVlansNil

`func (o *HyperflexClusterNetworkPolicy) SetVmNetworkVlansNil(b bool)`

 SetVmNetworkVlansNil sets the value for VmNetworkVlans to be an explicit nil

### UnsetVmNetworkVlans
`func (o *HyperflexClusterNetworkPolicy) UnsetVmNetworkVlans()`

UnsetVmNetworkVlans ensures that no value is present for VmNetworkVlans, not even an explicit nil
### GetClusterProfiles

`func (o *HyperflexClusterNetworkPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterNetworkPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterNetworkPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterNetworkPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterNetworkPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterNetworkPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterNetworkPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


