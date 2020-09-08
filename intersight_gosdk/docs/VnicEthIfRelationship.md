# VnicEthIfRelationship

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
**Cdn** | Pointer to [**VnicCdn**](vnic.Cdn.md) |  | [optional] 
**FailoverEnabled** | Pointer to **bool** | Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address that is assigned to the vnic based on the MAC pool that has been assigned to the LAN Connectivity Policy. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**Placement** | Pointer to [**VnicPlacementSettings**](vnic.PlacementSettings.md) |  | [optional] 
**StandbyVifId** | Pointer to **int64** | The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. | [optional] [readonly] 
**UsnicSettings** | Pointer to [**VnicUsnicSettings**](vnic.UsnicSettings.md) |  | [optional] 
**VifId** | Pointer to **int64** | The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. | [optional] [readonly] 
**VmqSettings** | Pointer to [**VnicVmqSettings**](vnic.VmqSettings.md) |  | [optional] 
**EthAdapterPolicy** | Pointer to [**VnicEthAdapterPolicyRelationship**](vnic.EthAdapterPolicy.Relationship.md) |  | [optional] 
**EthNetworkPolicy** | Pointer to [**VnicEthNetworkPolicyRelationship**](vnic.EthNetworkPolicy.Relationship.md) |  | [optional] 
**EthQosPolicy** | Pointer to [**VnicEthQosPolicyRelationship**](vnic.EthQosPolicy.Relationship.md) |  | [optional] 
**FabricEthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](fabric.EthNetworkControlPolicy.Relationship.md) |  | [optional] 
**FabricEthNetworkGroupPolicy** | Pointer to [**[]FabricEthNetworkGroupPolicyRelationship**](fabric.EthNetworkGroupPolicy.Relationship.md) | An array of relationships to fabricEthNetworkGroupPolicy resources. | [optional] 
**LanConnectivityPolicy** | Pointer to [**VnicLanConnectivityPolicyRelationship**](vnic.LanConnectivityPolicy.Relationship.md) |  | [optional] 
**LcpVnic** | Pointer to [**VnicEthIfRelationship**](vnic.EthIf.Relationship.md) |  | [optional] 
**MacLease** | Pointer to [**MacpoolLeaseRelationship**](macpool.Lease.Relationship.md) |  | [optional] 
**MacPool** | Pointer to [**MacpoolPoolRelationship**](macpool.Pool.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) |  | [optional] 
**SpVnics** | Pointer to [**[]VnicEthIfRelationship**](vnic.EthIf.Relationship.md) | An array of relationships to vnicEthIf resources. | [optional] 

## Methods

### NewVnicEthIfRelationship

`func NewVnicEthIfRelationship(classId string, objectType string, ) *VnicEthIfRelationship`

NewVnicEthIfRelationship instantiates a new VnicEthIfRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthIfRelationshipWithDefaults

`func NewVnicEthIfRelationshipWithDefaults() *VnicEthIfRelationship`

NewVnicEthIfRelationshipWithDefaults instantiates a new VnicEthIfRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VnicEthIfRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VnicEthIfRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VnicEthIfRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VnicEthIfRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VnicEthIfRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthIfRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthIfRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VnicEthIfRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VnicEthIfRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VnicEthIfRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VnicEthIfRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VnicEthIfRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VnicEthIfRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VnicEthIfRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VnicEthIfRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VnicEthIfRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VnicEthIfRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VnicEthIfRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VnicEthIfRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VnicEthIfRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VnicEthIfRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VnicEthIfRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VnicEthIfRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VnicEthIfRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthIfRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthIfRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VnicEthIfRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VnicEthIfRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VnicEthIfRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VnicEthIfRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VnicEthIfRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VnicEthIfRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VnicEthIfRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VnicEthIfRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VnicEthIfRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VnicEthIfRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VnicEthIfRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VnicEthIfRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VnicEthIfRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VnicEthIfRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VnicEthIfRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VnicEthIfRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VnicEthIfRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VnicEthIfRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VnicEthIfRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VnicEthIfRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VnicEthIfRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VnicEthIfRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VnicEthIfRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VnicEthIfRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VnicEthIfRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VnicEthIfRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VnicEthIfRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VnicEthIfRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VnicEthIfRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VnicEthIfRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VnicEthIfRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VnicEthIfRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VnicEthIfRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VnicEthIfRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VnicEthIfRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VnicEthIfRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VnicEthIfRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VnicEthIfRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetCdn

`func (o *VnicEthIfRelationship) GetCdn() VnicCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *VnicEthIfRelationship) GetCdnOk() (*VnicCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *VnicEthIfRelationship) SetCdn(v VnicCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *VnicEthIfRelationship) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetFailoverEnabled

`func (o *VnicEthIfRelationship) GetFailoverEnabled() bool`

GetFailoverEnabled returns the FailoverEnabled field if non-nil, zero value otherwise.

### GetFailoverEnabledOk

`func (o *VnicEthIfRelationship) GetFailoverEnabledOk() (*bool, bool)`

GetFailoverEnabledOk returns a tuple with the FailoverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverEnabled

`func (o *VnicEthIfRelationship) SetFailoverEnabled(v bool)`

SetFailoverEnabled sets FailoverEnabled field to given value.

### HasFailoverEnabled

`func (o *VnicEthIfRelationship) HasFailoverEnabled() bool`

HasFailoverEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *VnicEthIfRelationship) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VnicEthIfRelationship) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VnicEthIfRelationship) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VnicEthIfRelationship) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *VnicEthIfRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthIfRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthIfRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthIfRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicEthIfRelationship) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicEthIfRelationship) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicEthIfRelationship) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicEthIfRelationship) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicEthIfRelationship) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicEthIfRelationship) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicEthIfRelationship) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicEthIfRelationship) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetStandbyVifId

`func (o *VnicEthIfRelationship) GetStandbyVifId() int64`

GetStandbyVifId returns the StandbyVifId field if non-nil, zero value otherwise.

### GetStandbyVifIdOk

`func (o *VnicEthIfRelationship) GetStandbyVifIdOk() (*int64, bool)`

GetStandbyVifIdOk returns a tuple with the StandbyVifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyVifId

`func (o *VnicEthIfRelationship) SetStandbyVifId(v int64)`

SetStandbyVifId sets StandbyVifId field to given value.

### HasStandbyVifId

`func (o *VnicEthIfRelationship) HasStandbyVifId() bool`

HasStandbyVifId returns a boolean if a field has been set.

### GetUsnicSettings

`func (o *VnicEthIfRelationship) GetUsnicSettings() VnicUsnicSettings`

GetUsnicSettings returns the UsnicSettings field if non-nil, zero value otherwise.

### GetUsnicSettingsOk

`func (o *VnicEthIfRelationship) GetUsnicSettingsOk() (*VnicUsnicSettings, bool)`

GetUsnicSettingsOk returns a tuple with the UsnicSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnicSettings

`func (o *VnicEthIfRelationship) SetUsnicSettings(v VnicUsnicSettings)`

SetUsnicSettings sets UsnicSettings field to given value.

### HasUsnicSettings

`func (o *VnicEthIfRelationship) HasUsnicSettings() bool`

HasUsnicSettings returns a boolean if a field has been set.

### GetVifId

`func (o *VnicEthIfRelationship) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicEthIfRelationship) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicEthIfRelationship) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicEthIfRelationship) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetVmqSettings

`func (o *VnicEthIfRelationship) GetVmqSettings() VnicVmqSettings`

GetVmqSettings returns the VmqSettings field if non-nil, zero value otherwise.

### GetVmqSettingsOk

`func (o *VnicEthIfRelationship) GetVmqSettingsOk() (*VnicVmqSettings, bool)`

GetVmqSettingsOk returns a tuple with the VmqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmqSettings

`func (o *VnicEthIfRelationship) SetVmqSettings(v VnicVmqSettings)`

SetVmqSettings sets VmqSettings field to given value.

### HasVmqSettings

`func (o *VnicEthIfRelationship) HasVmqSettings() bool`

HasVmqSettings returns a boolean if a field has been set.

### GetEthAdapterPolicy

`func (o *VnicEthIfRelationship) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship`

GetEthAdapterPolicy returns the EthAdapterPolicy field if non-nil, zero value otherwise.

### GetEthAdapterPolicyOk

`func (o *VnicEthIfRelationship) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool)`

GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthAdapterPolicy

`func (o *VnicEthIfRelationship) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship)`

SetEthAdapterPolicy sets EthAdapterPolicy field to given value.

### HasEthAdapterPolicy

`func (o *VnicEthIfRelationship) HasEthAdapterPolicy() bool`

HasEthAdapterPolicy returns a boolean if a field has been set.

### GetEthNetworkPolicy

`func (o *VnicEthIfRelationship) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship`

GetEthNetworkPolicy returns the EthNetworkPolicy field if non-nil, zero value otherwise.

### GetEthNetworkPolicyOk

`func (o *VnicEthIfRelationship) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool)`

GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkPolicy

`func (o *VnicEthIfRelationship) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship)`

SetEthNetworkPolicy sets EthNetworkPolicy field to given value.

### HasEthNetworkPolicy

`func (o *VnicEthIfRelationship) HasEthNetworkPolicy() bool`

HasEthNetworkPolicy returns a boolean if a field has been set.

### GetEthQosPolicy

`func (o *VnicEthIfRelationship) GetEthQosPolicy() VnicEthQosPolicyRelationship`

GetEthQosPolicy returns the EthQosPolicy field if non-nil, zero value otherwise.

### GetEthQosPolicyOk

`func (o *VnicEthIfRelationship) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool)`

GetEthQosPolicyOk returns a tuple with the EthQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthQosPolicy

`func (o *VnicEthIfRelationship) SetEthQosPolicy(v VnicEthQosPolicyRelationship)`

SetEthQosPolicy sets EthQosPolicy field to given value.

### HasEthQosPolicy

`func (o *VnicEthIfRelationship) HasEthQosPolicy() bool`

HasEthQosPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkControlPolicy

`func (o *VnicEthIfRelationship) GetFabricEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetFabricEthNetworkControlPolicy returns the FabricEthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkControlPolicyOk

`func (o *VnicEthIfRelationship) GetFabricEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetFabricEthNetworkControlPolicyOk returns a tuple with the FabricEthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkControlPolicy

`func (o *VnicEthIfRelationship) SetFabricEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetFabricEthNetworkControlPolicy sets FabricEthNetworkControlPolicy field to given value.

### HasFabricEthNetworkControlPolicy

`func (o *VnicEthIfRelationship) HasFabricEthNetworkControlPolicy() bool`

HasFabricEthNetworkControlPolicy returns a boolean if a field has been set.

### GetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfRelationship) GetFabricEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship`

GetFabricEthNetworkGroupPolicy returns the FabricEthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetFabricEthNetworkGroupPolicyOk

`func (o *VnicEthIfRelationship) GetFabricEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool)`

GetFabricEthNetworkGroupPolicyOk returns a tuple with the FabricEthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricEthNetworkGroupPolicy

`func (o *VnicEthIfRelationship) SetFabricEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship)`

SetFabricEthNetworkGroupPolicy sets FabricEthNetworkGroupPolicy field to given value.

### HasFabricEthNetworkGroupPolicy

`func (o *VnicEthIfRelationship) HasFabricEthNetworkGroupPolicy() bool`

HasFabricEthNetworkGroupPolicy returns a boolean if a field has been set.

### SetFabricEthNetworkGroupPolicyNil

`func (o *VnicEthIfRelationship) SetFabricEthNetworkGroupPolicyNil(b bool)`

 SetFabricEthNetworkGroupPolicyNil sets the value for FabricEthNetworkGroupPolicy to be an explicit nil

### UnsetFabricEthNetworkGroupPolicy
`func (o *VnicEthIfRelationship) UnsetFabricEthNetworkGroupPolicy()`

UnsetFabricEthNetworkGroupPolicy ensures that no value is present for FabricEthNetworkGroupPolicy, not even an explicit nil
### GetLanConnectivityPolicy

`func (o *VnicEthIfRelationship) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship`

GetLanConnectivityPolicy returns the LanConnectivityPolicy field if non-nil, zero value otherwise.

### GetLanConnectivityPolicyOk

`func (o *VnicEthIfRelationship) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool)`

GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanConnectivityPolicy

`func (o *VnicEthIfRelationship) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship)`

SetLanConnectivityPolicy sets LanConnectivityPolicy field to given value.

### HasLanConnectivityPolicy

`func (o *VnicEthIfRelationship) HasLanConnectivityPolicy() bool`

HasLanConnectivityPolicy returns a boolean if a field has been set.

### GetLcpVnic

`func (o *VnicEthIfRelationship) GetLcpVnic() VnicEthIfRelationship`

GetLcpVnic returns the LcpVnic field if non-nil, zero value otherwise.

### GetLcpVnicOk

`func (o *VnicEthIfRelationship) GetLcpVnicOk() (*VnicEthIfRelationship, bool)`

GetLcpVnicOk returns a tuple with the LcpVnic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcpVnic

`func (o *VnicEthIfRelationship) SetLcpVnic(v VnicEthIfRelationship)`

SetLcpVnic sets LcpVnic field to given value.

### HasLcpVnic

`func (o *VnicEthIfRelationship) HasLcpVnic() bool`

HasLcpVnic returns a boolean if a field has been set.

### GetMacLease

`func (o *VnicEthIfRelationship) GetMacLease() MacpoolLeaseRelationship`

GetMacLease returns the MacLease field if non-nil, zero value otherwise.

### GetMacLeaseOk

`func (o *VnicEthIfRelationship) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool)`

GetMacLeaseOk returns a tuple with the MacLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLease

`func (o *VnicEthIfRelationship) SetMacLease(v MacpoolLeaseRelationship)`

SetMacLease sets MacLease field to given value.

### HasMacLease

`func (o *VnicEthIfRelationship) HasMacLease() bool`

HasMacLease returns a boolean if a field has been set.

### GetMacPool

`func (o *VnicEthIfRelationship) GetMacPool() MacpoolPoolRelationship`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *VnicEthIfRelationship) GetMacPoolOk() (*MacpoolPoolRelationship, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *VnicEthIfRelationship) SetMacPool(v MacpoolPoolRelationship)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *VnicEthIfRelationship) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetProfile

`func (o *VnicEthIfRelationship) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicEthIfRelationship) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicEthIfRelationship) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicEthIfRelationship) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSpVnics

`func (o *VnicEthIfRelationship) GetSpVnics() []VnicEthIfRelationship`

GetSpVnics returns the SpVnics field if non-nil, zero value otherwise.

### GetSpVnicsOk

`func (o *VnicEthIfRelationship) GetSpVnicsOk() (*[]VnicEthIfRelationship, bool)`

GetSpVnicsOk returns a tuple with the SpVnics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVnics

`func (o *VnicEthIfRelationship) SetSpVnics(v []VnicEthIfRelationship)`

SetSpVnics sets SpVnics field to given value.

### HasSpVnics

`func (o *VnicEthIfRelationship) HasSpVnics() bool`

HasSpVnics returns a boolean if a field has been set.

### SetSpVnicsNil

`func (o *VnicEthIfRelationship) SetSpVnicsNil(b bool)`

 SetSpVnicsNil sets the value for SpVnics to be an explicit nil

### UnsetSpVnics
`func (o *VnicEthIfRelationship) UnsetSpVnics()`

UnsetSpVnics ensures that no value is present for SpVnics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


