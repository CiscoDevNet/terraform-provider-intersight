# HyperflexClusterNetworkPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JumboFrame** | Pointer to **bool** | Enable or disable jumbo frames. | [optional] 
**KvmIpRange** | Pointer to [**HyperflexIpAddrRange**](hyperflex.IpAddrRange.md) |  | [optional] 
**MacPrefixRange** | Pointer to [**HyperflexMacAddrPrefixRange**](hyperflex.MacAddrPrefixRange.md) |  | [optional] 
**MgmtVlan** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**UplinkSpeed** | Pointer to **string** | Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be &#39;1G&#39; or &#39;10G+&#39;. Use &#39;10G+&#39; for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be &#39;default&#39; only. * &#x60;default&#x60; - Current default value set on the hardware platform. * &#x60;1G&#x60; - A link speed of 1 gigabit per second. * &#x60;10G&#x60; - A link speed of 10 gigabits per second or above. | [optional] [default to "default"]
**VmMigrationVlan** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**VmNetworkVlans** | Pointer to [**[]HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterNetworkPolicyAllOf

`func NewHyperflexClusterNetworkPolicyAllOf() *HyperflexClusterNetworkPolicyAllOf`

NewHyperflexClusterNetworkPolicyAllOf instantiates a new HyperflexClusterNetworkPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterNetworkPolicyAllOfWithDefaults

`func NewHyperflexClusterNetworkPolicyAllOfWithDefaults() *HyperflexClusterNetworkPolicyAllOf`

NewHyperflexClusterNetworkPolicyAllOfWithDefaults instantiates a new HyperflexClusterNetworkPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJumboFrame

`func (o *HyperflexClusterNetworkPolicyAllOf) GetJumboFrame() bool`

GetJumboFrame returns the JumboFrame field if non-nil, zero value otherwise.

### GetJumboFrameOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetJumboFrameOk() (*bool, bool)`

GetJumboFrameOk returns a tuple with the JumboFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboFrame

`func (o *HyperflexClusterNetworkPolicyAllOf) SetJumboFrame(v bool)`

SetJumboFrame sets JumboFrame field to given value.

### HasJumboFrame

`func (o *HyperflexClusterNetworkPolicyAllOf) HasJumboFrame() bool`

HasJumboFrame returns a boolean if a field has been set.

### GetKvmIpRange

`func (o *HyperflexClusterNetworkPolicyAllOf) GetKvmIpRange() HyperflexIpAddrRange`

GetKvmIpRange returns the KvmIpRange field if non-nil, zero value otherwise.

### GetKvmIpRangeOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetKvmIpRangeOk returns a tuple with the KvmIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpRange

`func (o *HyperflexClusterNetworkPolicyAllOf) SetKvmIpRange(v HyperflexIpAddrRange)`

SetKvmIpRange sets KvmIpRange field to given value.

### HasKvmIpRange

`func (o *HyperflexClusterNetworkPolicyAllOf) HasKvmIpRange() bool`

HasKvmIpRange returns a boolean if a field has been set.

### GetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyAllOf) GetMacPrefixRange() HyperflexMacAddrPrefixRange`

GetMacPrefixRange returns the MacPrefixRange field if non-nil, zero value otherwise.

### GetMacPrefixRangeOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool)`

GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyAllOf) SetMacPrefixRange(v HyperflexMacAddrPrefixRange)`

SetMacPrefixRange sets MacPrefixRange field to given value.

### HasMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyAllOf) HasMacPrefixRange() bool`

HasMacPrefixRange returns a boolean if a field has been set.

### GetMgmtVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) GetMgmtVlan() HyperflexNamedVlan`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetMgmtVlanOk() (*HyperflexNamedVlan, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) SetMgmtVlan(v HyperflexNamedVlan)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### GetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyAllOf) GetUplinkSpeed() string`

GetUplinkSpeed returns the UplinkSpeed field if non-nil, zero value otherwise.

### GetUplinkSpeedOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetUplinkSpeedOk() (*string, bool)`

GetUplinkSpeedOk returns a tuple with the UplinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyAllOf) SetUplinkSpeed(v string)`

SetUplinkSpeed sets UplinkSpeed field to given value.

### HasUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyAllOf) HasUplinkSpeed() bool`

HasUplinkSpeed returns a boolean if a field has been set.

### GetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) GetVmMigrationVlan() HyperflexNamedVlan`

GetVmMigrationVlan returns the VmMigrationVlan field if non-nil, zero value otherwise.

### GetVmMigrationVlanOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetVmMigrationVlanOk() (*HyperflexNamedVlan, bool)`

GetVmMigrationVlanOk returns a tuple with the VmMigrationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) SetVmMigrationVlan(v HyperflexNamedVlan)`

SetVmMigrationVlan sets VmMigrationVlan field to given value.

### HasVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyAllOf) HasVmMigrationVlan() bool`

HasVmMigrationVlan returns a boolean if a field has been set.

### GetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyAllOf) GetVmNetworkVlans() []HyperflexNamedVlan`

GetVmNetworkVlans returns the VmNetworkVlans field if non-nil, zero value otherwise.

### GetVmNetworkVlansOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetVmNetworkVlansOk() (*[]HyperflexNamedVlan, bool)`

GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyAllOf) SetVmNetworkVlans(v []HyperflexNamedVlan)`

SetVmNetworkVlans sets VmNetworkVlans field to given value.

### HasVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyAllOf) HasVmNetworkVlans() bool`

HasVmNetworkVlans returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexClusterNetworkPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterNetworkPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterNetworkPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterNetworkPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterNetworkPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterNetworkPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


