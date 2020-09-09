# HyperflexClusterNetworkPolicyRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**Name** | Pointer to **string** | Name of the concrete policy. | [optional] 
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

### NewHyperflexClusterNetworkPolicyRelationship

`func NewHyperflexClusterNetworkPolicyRelationship(classId string, objectType string, ) *HyperflexClusterNetworkPolicyRelationship`

NewHyperflexClusterNetworkPolicyRelationship instantiates a new HyperflexClusterNetworkPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterNetworkPolicyRelationshipWithDefaults

`func NewHyperflexClusterNetworkPolicyRelationshipWithDefaults() *HyperflexClusterNetworkPolicyRelationship`

NewHyperflexClusterNetworkPolicyRelationshipWithDefaults instantiates a new HyperflexClusterNetworkPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexClusterNetworkPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterNetworkPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexClusterNetworkPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexClusterNetworkPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexClusterNetworkPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexClusterNetworkPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexClusterNetworkPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexClusterNetworkPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexClusterNetworkPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexClusterNetworkPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterNetworkPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexClusterNetworkPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexClusterNetworkPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexClusterNetworkPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexClusterNetworkPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexClusterNetworkPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexClusterNetworkPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexClusterNetworkPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexClusterNetworkPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexClusterNetworkPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexClusterNetworkPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexClusterNetworkPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexClusterNetworkPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexClusterNetworkPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexClusterNetworkPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexClusterNetworkPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexClusterNetworkPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexClusterNetworkPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexClusterNetworkPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexClusterNetworkPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexClusterNetworkPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexClusterNetworkPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexClusterNetworkPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexClusterNetworkPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexClusterNetworkPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexClusterNetworkPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexClusterNetworkPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexClusterNetworkPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexClusterNetworkPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexClusterNetworkPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexClusterNetworkPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HyperflexClusterNetworkPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexClusterNetworkPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexClusterNetworkPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJumboFrame

`func (o *HyperflexClusterNetworkPolicyRelationship) GetJumboFrame() bool`

GetJumboFrame returns the JumboFrame field if non-nil, zero value otherwise.

### GetJumboFrameOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetJumboFrameOk() (*bool, bool)`

GetJumboFrameOk returns a tuple with the JumboFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumboFrame

`func (o *HyperflexClusterNetworkPolicyRelationship) SetJumboFrame(v bool)`

SetJumboFrame sets JumboFrame field to given value.

### HasJumboFrame

`func (o *HyperflexClusterNetworkPolicyRelationship) HasJumboFrame() bool`

HasJumboFrame returns a boolean if a field has been set.

### GetKvmIpRange

`func (o *HyperflexClusterNetworkPolicyRelationship) GetKvmIpRange() HyperflexIpAddrRange`

GetKvmIpRange returns the KvmIpRange field if non-nil, zero value otherwise.

### GetKvmIpRangeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool)`

GetKvmIpRangeOk returns a tuple with the KvmIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpRange

`func (o *HyperflexClusterNetworkPolicyRelationship) SetKvmIpRange(v HyperflexIpAddrRange)`

SetKvmIpRange sets KvmIpRange field to given value.

### HasKvmIpRange

`func (o *HyperflexClusterNetworkPolicyRelationship) HasKvmIpRange() bool`

HasKvmIpRange returns a boolean if a field has been set.

### GetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMacPrefixRange() HyperflexMacAddrPrefixRange`

GetMacPrefixRange returns the MacPrefixRange field if non-nil, zero value otherwise.

### GetMacPrefixRangeOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool)`

GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyRelationship) SetMacPrefixRange(v HyperflexMacAddrPrefixRange)`

SetMacPrefixRange sets MacPrefixRange field to given value.

### HasMacPrefixRange

`func (o *HyperflexClusterNetworkPolicyRelationship) HasMacPrefixRange() bool`

HasMacPrefixRange returns a boolean if a field has been set.

### GetMgmtVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMgmtVlan() HyperflexNamedVlan`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetMgmtVlanOk() (*HyperflexNamedVlan, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) SetMgmtVlan(v HyperflexNamedVlan)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### GetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyRelationship) GetUplinkSpeed() string`

GetUplinkSpeed returns the UplinkSpeed field if non-nil, zero value otherwise.

### GetUplinkSpeedOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetUplinkSpeedOk() (*string, bool)`

GetUplinkSpeedOk returns a tuple with the UplinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyRelationship) SetUplinkSpeed(v string)`

SetUplinkSpeed sets UplinkSpeed field to given value.

### HasUplinkSpeed

`func (o *HyperflexClusterNetworkPolicyRelationship) HasUplinkSpeed() bool`

HasUplinkSpeed returns a boolean if a field has been set.

### GetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVmMigrationVlan() HyperflexNamedVlan`

GetVmMigrationVlan returns the VmMigrationVlan field if non-nil, zero value otherwise.

### GetVmMigrationVlanOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVmMigrationVlanOk() (*HyperflexNamedVlan, bool)`

GetVmMigrationVlanOk returns a tuple with the VmMigrationVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) SetVmMigrationVlan(v HyperflexNamedVlan)`

SetVmMigrationVlan sets VmMigrationVlan field to given value.

### HasVmMigrationVlan

`func (o *HyperflexClusterNetworkPolicyRelationship) HasVmMigrationVlan() bool`

HasVmMigrationVlan returns a boolean if a field has been set.

### GetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVmNetworkVlans() []HyperflexNamedVlan`

GetVmNetworkVlans returns the VmNetworkVlans field if non-nil, zero value otherwise.

### GetVmNetworkVlansOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetVmNetworkVlansOk() (*[]HyperflexNamedVlan, bool)`

GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyRelationship) SetVmNetworkVlans(v []HyperflexNamedVlan)`

SetVmNetworkVlans sets VmNetworkVlans field to given value.

### HasVmNetworkVlans

`func (o *HyperflexClusterNetworkPolicyRelationship) HasVmNetworkVlans() bool`

HasVmNetworkVlans returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexClusterNetworkPolicyRelationship) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexClusterNetworkPolicyRelationship) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexClusterNetworkPolicyRelationship) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexClusterNetworkPolicyRelationship) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexClusterNetworkPolicyRelationship) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterNetworkPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterNetworkPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterNetworkPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterNetworkPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


