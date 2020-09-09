# VnicFcIfRelationship

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
**Name** | Pointer to **string** | Name of the virtual fibre channel interface. | [optional] 
**Order** | Pointer to **int64** | The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**PersistentBindings** | Pointer to **bool** | Enables retention of LUN ID associations in memory until they are manually cleared. | [optional] 
**Placement** | Pointer to [**VnicPlacementSettings**](vnic.PlacementSettings.md) |  | [optional] 
**Type** | Pointer to **string** | VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * &#x60;fc-initiator&#x60; - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-initiator&#x60; - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * &#x60;fc-nvme-target&#x60; - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * &#x60;fc-target&#x60; - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver. | [optional] [default to "fc-initiator"]
**VifId** | Pointer to **int64** | This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The WWPN address that is assigned to the vhba based on the wwn pool that has been assigned to the SAN Connectivity Policy. | [optional] [readonly] 
**FcAdapterPolicy** | Pointer to [**VnicFcAdapterPolicyRelationship**](vnic.FcAdapterPolicy.Relationship.md) |  | [optional] 
**FcNetworkPolicy** | Pointer to [**VnicFcNetworkPolicyRelationship**](vnic.FcNetworkPolicy.Relationship.md) |  | [optional] 
**FcQosPolicy** | Pointer to [**VnicFcQosPolicyRelationship**](vnic.FcQosPolicy.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) |  | [optional] 
**SanConnectivityPolicy** | Pointer to [**VnicSanConnectivityPolicyRelationship**](vnic.SanConnectivityPolicy.Relationship.md) |  | [optional] 
**ScpVhba** | Pointer to [**VnicFcIfRelationship**](vnic.FcIf.Relationship.md) |  | [optional] 
**SpVhbas** | Pointer to [**[]VnicFcIfRelationship**](vnic.FcIf.Relationship.md) | An array of relationships to vnicFcIf resources. | [optional] 
**WwpnLease** | Pointer to [**FcpoolLeaseRelationship**](fcpool.Lease.Relationship.md) |  | [optional] 
**WwpnPool** | Pointer to [**FcpoolPoolRelationship**](fcpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewVnicFcIfRelationship

`func NewVnicFcIfRelationship(classId string, objectType string, ) *VnicFcIfRelationship`

NewVnicFcIfRelationship instantiates a new VnicFcIfRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcIfRelationshipWithDefaults

`func NewVnicFcIfRelationshipWithDefaults() *VnicFcIfRelationship`

NewVnicFcIfRelationshipWithDefaults instantiates a new VnicFcIfRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VnicFcIfRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VnicFcIfRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VnicFcIfRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VnicFcIfRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VnicFcIfRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcIfRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcIfRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VnicFcIfRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VnicFcIfRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VnicFcIfRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VnicFcIfRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VnicFcIfRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VnicFcIfRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VnicFcIfRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VnicFcIfRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VnicFcIfRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VnicFcIfRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VnicFcIfRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VnicFcIfRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VnicFcIfRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VnicFcIfRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VnicFcIfRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VnicFcIfRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VnicFcIfRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcIfRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcIfRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VnicFcIfRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VnicFcIfRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VnicFcIfRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VnicFcIfRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VnicFcIfRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VnicFcIfRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VnicFcIfRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VnicFcIfRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VnicFcIfRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VnicFcIfRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VnicFcIfRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VnicFcIfRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VnicFcIfRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VnicFcIfRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VnicFcIfRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VnicFcIfRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VnicFcIfRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VnicFcIfRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VnicFcIfRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VnicFcIfRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VnicFcIfRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VnicFcIfRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VnicFcIfRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VnicFcIfRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VnicFcIfRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VnicFcIfRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VnicFcIfRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VnicFcIfRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VnicFcIfRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VnicFcIfRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VnicFcIfRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VnicFcIfRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VnicFcIfRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VnicFcIfRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VnicFcIfRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VnicFcIfRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VnicFcIfRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VnicFcIfRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetName

`func (o *VnicFcIfRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcIfRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcIfRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcIfRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VnicFcIfRelationship) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VnicFcIfRelationship) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VnicFcIfRelationship) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VnicFcIfRelationship) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPersistentBindings

`func (o *VnicFcIfRelationship) GetPersistentBindings() bool`

GetPersistentBindings returns the PersistentBindings field if non-nil, zero value otherwise.

### GetPersistentBindingsOk

`func (o *VnicFcIfRelationship) GetPersistentBindingsOk() (*bool, bool)`

GetPersistentBindingsOk returns a tuple with the PersistentBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBindings

`func (o *VnicFcIfRelationship) SetPersistentBindings(v bool)`

SetPersistentBindings sets PersistentBindings field to given value.

### HasPersistentBindings

`func (o *VnicFcIfRelationship) HasPersistentBindings() bool`

HasPersistentBindings returns a boolean if a field has been set.

### GetPlacement

`func (o *VnicFcIfRelationship) GetPlacement() VnicPlacementSettings`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VnicFcIfRelationship) GetPlacementOk() (*VnicPlacementSettings, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VnicFcIfRelationship) SetPlacement(v VnicPlacementSettings)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VnicFcIfRelationship) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetType

`func (o *VnicFcIfRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnicFcIfRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnicFcIfRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VnicFcIfRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVifId

`func (o *VnicFcIfRelationship) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *VnicFcIfRelationship) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *VnicFcIfRelationship) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *VnicFcIfRelationship) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

### GetWwpn

`func (o *VnicFcIfRelationship) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *VnicFcIfRelationship) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *VnicFcIfRelationship) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *VnicFcIfRelationship) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetFcAdapterPolicy

`func (o *VnicFcIfRelationship) GetFcAdapterPolicy() VnicFcAdapterPolicyRelationship`

GetFcAdapterPolicy returns the FcAdapterPolicy field if non-nil, zero value otherwise.

### GetFcAdapterPolicyOk

`func (o *VnicFcIfRelationship) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyRelationship, bool)`

GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcAdapterPolicy

`func (o *VnicFcIfRelationship) SetFcAdapterPolicy(v VnicFcAdapterPolicyRelationship)`

SetFcAdapterPolicy sets FcAdapterPolicy field to given value.

### HasFcAdapterPolicy

`func (o *VnicFcIfRelationship) HasFcAdapterPolicy() bool`

HasFcAdapterPolicy returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *VnicFcIfRelationship) GetFcNetworkPolicy() VnicFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *VnicFcIfRelationship) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *VnicFcIfRelationship) SetFcNetworkPolicy(v VnicFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *VnicFcIfRelationship) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.

### GetFcQosPolicy

`func (o *VnicFcIfRelationship) GetFcQosPolicy() VnicFcQosPolicyRelationship`

GetFcQosPolicy returns the FcQosPolicy field if non-nil, zero value otherwise.

### GetFcQosPolicyOk

`func (o *VnicFcIfRelationship) GetFcQosPolicyOk() (*VnicFcQosPolicyRelationship, bool)`

GetFcQosPolicyOk returns a tuple with the FcQosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcQosPolicy

`func (o *VnicFcIfRelationship) SetFcQosPolicy(v VnicFcQosPolicyRelationship)`

SetFcQosPolicy sets FcQosPolicy field to given value.

### HasFcQosPolicy

`func (o *VnicFcIfRelationship) HasFcQosPolicy() bool`

HasFcQosPolicy returns a boolean if a field has been set.

### GetProfile

`func (o *VnicFcIfRelationship) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicFcIfRelationship) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicFcIfRelationship) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicFcIfRelationship) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSanConnectivityPolicy

`func (o *VnicFcIfRelationship) GetSanConnectivityPolicy() VnicSanConnectivityPolicyRelationship`

GetSanConnectivityPolicy returns the SanConnectivityPolicy field if non-nil, zero value otherwise.

### GetSanConnectivityPolicyOk

`func (o *VnicFcIfRelationship) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyRelationship, bool)`

GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanConnectivityPolicy

`func (o *VnicFcIfRelationship) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyRelationship)`

SetSanConnectivityPolicy sets SanConnectivityPolicy field to given value.

### HasSanConnectivityPolicy

`func (o *VnicFcIfRelationship) HasSanConnectivityPolicy() bool`

HasSanConnectivityPolicy returns a boolean if a field has been set.

### GetScpVhba

`func (o *VnicFcIfRelationship) GetScpVhba() VnicFcIfRelationship`

GetScpVhba returns the ScpVhba field if non-nil, zero value otherwise.

### GetScpVhbaOk

`func (o *VnicFcIfRelationship) GetScpVhbaOk() (*VnicFcIfRelationship, bool)`

GetScpVhbaOk returns a tuple with the ScpVhba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpVhba

`func (o *VnicFcIfRelationship) SetScpVhba(v VnicFcIfRelationship)`

SetScpVhba sets ScpVhba field to given value.

### HasScpVhba

`func (o *VnicFcIfRelationship) HasScpVhba() bool`

HasScpVhba returns a boolean if a field has been set.

### GetSpVhbas

`func (o *VnicFcIfRelationship) GetSpVhbas() []VnicFcIfRelationship`

GetSpVhbas returns the SpVhbas field if non-nil, zero value otherwise.

### GetSpVhbasOk

`func (o *VnicFcIfRelationship) GetSpVhbasOk() (*[]VnicFcIfRelationship, bool)`

GetSpVhbasOk returns a tuple with the SpVhbas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVhbas

`func (o *VnicFcIfRelationship) SetSpVhbas(v []VnicFcIfRelationship)`

SetSpVhbas sets SpVhbas field to given value.

### HasSpVhbas

`func (o *VnicFcIfRelationship) HasSpVhbas() bool`

HasSpVhbas returns a boolean if a field has been set.

### SetSpVhbasNil

`func (o *VnicFcIfRelationship) SetSpVhbasNil(b bool)`

 SetSpVhbasNil sets the value for SpVhbas to be an explicit nil

### UnsetSpVhbas
`func (o *VnicFcIfRelationship) UnsetSpVhbas()`

UnsetSpVhbas ensures that no value is present for SpVhbas, not even an explicit nil
### GetWwpnLease

`func (o *VnicFcIfRelationship) GetWwpnLease() FcpoolLeaseRelationship`

GetWwpnLease returns the WwpnLease field if non-nil, zero value otherwise.

### GetWwpnLeaseOk

`func (o *VnicFcIfRelationship) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwpnLeaseOk returns a tuple with the WwpnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnLease

`func (o *VnicFcIfRelationship) SetWwpnLease(v FcpoolLeaseRelationship)`

SetWwpnLease sets WwpnLease field to given value.

### HasWwpnLease

`func (o *VnicFcIfRelationship) HasWwpnLease() bool`

HasWwpnLease returns a boolean if a field has been set.

### GetWwpnPool

`func (o *VnicFcIfRelationship) GetWwpnPool() FcpoolPoolRelationship`

GetWwpnPool returns the WwpnPool field if non-nil, zero value otherwise.

### GetWwpnPoolOk

`func (o *VnicFcIfRelationship) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool)`

GetWwpnPoolOk returns a tuple with the WwpnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpnPool

`func (o *VnicFcIfRelationship) SetWwpnPool(v FcpoolPoolRelationship)`

SetWwpnPool sets WwpnPool field to given value.

### HasWwpnPool

`func (o *VnicFcIfRelationship) HasWwpnPool() bool`

HasWwpnPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


