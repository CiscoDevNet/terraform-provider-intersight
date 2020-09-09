# AdapterHostEthInterfaceRelationship

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
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of an Interface. | [optional] 
**AcknowledgedPeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 
**PeerInterface** | Pointer to [**EtherPhysicalPortBaseRelationship**](ether.PhysicalPortBase.Relationship.md) |  | [optional] 
**AdminState** | Pointer to **string** | Admin state of the Host Ethernet Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Ethernet Interface. | [optional] [readonly] 
**HostEthInterfaceId** | Pointer to **int64** | Unique Identifier for an Host Ethernet Interface within the adapter object. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Type of External Ethernet Interface. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Ethernet Interface. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability status of Host Ethernet Channel Interface. | [optional] [readonly] 
**OriginalMacAddress** | Pointer to **string** | The factory default Mac address of the Host Ethernet Interface. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The PCI address of the Host Ethernet Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | The distinguished name of the peer endpoint connected to the Host Ethernet interface. | [optional] [readonly] 
**VirtualizationPreference** | Pointer to **string** | Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not. | [optional] [readonly] 
**VnicDn** | Pointer to **string** | The Virtual Ethernet Interface DN connected to the Host Ethernet Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterHostEthInterfaceRelationship

`func NewAdapterHostEthInterfaceRelationship(classId string, objectType string, ) *AdapterHostEthInterfaceRelationship`

NewAdapterHostEthInterfaceRelationship instantiates a new AdapterHostEthInterfaceRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostEthInterfaceRelationshipWithDefaults

`func NewAdapterHostEthInterfaceRelationshipWithDefaults() *AdapterHostEthInterfaceRelationship`

NewAdapterHostEthInterfaceRelationshipWithDefaults instantiates a new AdapterHostEthInterfaceRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *AdapterHostEthInterfaceRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AdapterHostEthInterfaceRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AdapterHostEthInterfaceRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AdapterHostEthInterfaceRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *AdapterHostEthInterfaceRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostEthInterfaceRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostEthInterfaceRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *AdapterHostEthInterfaceRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AdapterHostEthInterfaceRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AdapterHostEthInterfaceRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AdapterHostEthInterfaceRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *AdapterHostEthInterfaceRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *AdapterHostEthInterfaceRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *AdapterHostEthInterfaceRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *AdapterHostEthInterfaceRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *AdapterHostEthInterfaceRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *AdapterHostEthInterfaceRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *AdapterHostEthInterfaceRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *AdapterHostEthInterfaceRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *AdapterHostEthInterfaceRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *AdapterHostEthInterfaceRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *AdapterHostEthInterfaceRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *AdapterHostEthInterfaceRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *AdapterHostEthInterfaceRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostEthInterfaceRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostEthInterfaceRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *AdapterHostEthInterfaceRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AdapterHostEthInterfaceRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AdapterHostEthInterfaceRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AdapterHostEthInterfaceRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *AdapterHostEthInterfaceRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *AdapterHostEthInterfaceRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *AdapterHostEthInterfaceRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *AdapterHostEthInterfaceRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *AdapterHostEthInterfaceRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdapterHostEthInterfaceRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdapterHostEthInterfaceRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdapterHostEthInterfaceRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *AdapterHostEthInterfaceRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *AdapterHostEthInterfaceRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *AdapterHostEthInterfaceRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *AdapterHostEthInterfaceRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *AdapterHostEthInterfaceRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *AdapterHostEthInterfaceRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *AdapterHostEthInterfaceRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *AdapterHostEthInterfaceRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *AdapterHostEthInterfaceRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *AdapterHostEthInterfaceRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *AdapterHostEthInterfaceRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AdapterHostEthInterfaceRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AdapterHostEthInterfaceRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AdapterHostEthInterfaceRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *AdapterHostEthInterfaceRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *AdapterHostEthInterfaceRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *AdapterHostEthInterfaceRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *AdapterHostEthInterfaceRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *AdapterHostEthInterfaceRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *AdapterHostEthInterfaceRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *AdapterHostEthInterfaceRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *AdapterHostEthInterfaceRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *AdapterHostEthInterfaceRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *AdapterHostEthInterfaceRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *AdapterHostEthInterfaceRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *AdapterHostEthInterfaceRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *AdapterHostEthInterfaceRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *AdapterHostEthInterfaceRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *AdapterHostEthInterfaceRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *AdapterHostEthInterfaceRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *AdapterHostEthInterfaceRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *AdapterHostEthInterfaceRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *AdapterHostEthInterfaceRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *AdapterHostEthInterfaceRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *AdapterHostEthInterfaceRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *AdapterHostEthInterfaceRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *AdapterHostEthInterfaceRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *AdapterHostEthInterfaceRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterHostEthInterfaceRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostEthInterfaceRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostEthInterfaceRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostEthInterfaceRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetAcknowledgedPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) GetAcknowledgedPeerInterface() EtherPhysicalPortBaseRelationship`

GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field if non-nil, zero value otherwise.

### GetAcknowledgedPeerInterfaceOk

`func (o *AdapterHostEthInterfaceRelationship) GetAcknowledgedPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) SetAcknowledgedPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetAcknowledgedPeerInterface sets AcknowledgedPeerInterface field to given value.

### HasAcknowledgedPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) HasAcknowledgedPeerInterface() bool`

HasAcknowledgedPeerInterface returns a boolean if a field has been set.

### GetPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) GetPeerInterface() EtherPhysicalPortBaseRelationship`

GetPeerInterface returns the PeerInterface field if non-nil, zero value otherwise.

### GetPeerInterfaceOk

`func (o *AdapterHostEthInterfaceRelationship) GetPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool)`

GetPeerInterfaceOk returns a tuple with the PeerInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) SetPeerInterface(v EtherPhysicalPortBaseRelationship)`

SetPeerInterface sets PeerInterface field to given value.

### HasPeerInterface

`func (o *AdapterHostEthInterfaceRelationship) HasPeerInterface() bool`

HasPeerInterface returns a boolean if a field has been set.

### GetAdminState

`func (o *AdapterHostEthInterfaceRelationship) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostEthInterfaceRelationship) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostEthInterfaceRelationship) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostEthInterfaceRelationship) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostEthInterfaceRelationship) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostEthInterfaceRelationship) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostEthInterfaceRelationship) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostEthInterfaceRelationship) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostEthInterfaceId

`func (o *AdapterHostEthInterfaceRelationship) GetHostEthInterfaceId() int64`

GetHostEthInterfaceId returns the HostEthInterfaceId field if non-nil, zero value otherwise.

### GetHostEthInterfaceIdOk

`func (o *AdapterHostEthInterfaceRelationship) GetHostEthInterfaceIdOk() (*int64, bool)`

GetHostEthInterfaceIdOk returns a tuple with the HostEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthInterfaceId

`func (o *AdapterHostEthInterfaceRelationship) SetHostEthInterfaceId(v int64)`

SetHostEthInterfaceId sets HostEthInterfaceId field to given value.

### HasHostEthInterfaceId

`func (o *AdapterHostEthInterfaceRelationship) HasHostEthInterfaceId() bool`

HasHostEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterHostEthInterfaceRelationship) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterHostEthInterfaceRelationship) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterHostEthInterfaceRelationship) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterHostEthInterfaceRelationship) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterHostEthInterfaceRelationship) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterHostEthInterfaceRelationship) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterHostEthInterfaceRelationship) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterHostEthInterfaceRelationship) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostEthInterfaceRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostEthInterfaceRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostEthInterfaceRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostEthInterfaceRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostEthInterfaceRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostEthInterfaceRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostEthInterfaceRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostEthInterfaceRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalMacAddress

`func (o *AdapterHostEthInterfaceRelationship) GetOriginalMacAddress() string`

GetOriginalMacAddress returns the OriginalMacAddress field if non-nil, zero value otherwise.

### GetOriginalMacAddressOk

`func (o *AdapterHostEthInterfaceRelationship) GetOriginalMacAddressOk() (*string, bool)`

GetOriginalMacAddressOk returns a tuple with the OriginalMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMacAddress

`func (o *AdapterHostEthInterfaceRelationship) SetOriginalMacAddress(v string)`

SetOriginalMacAddress sets OriginalMacAddress field to given value.

### HasOriginalMacAddress

`func (o *AdapterHostEthInterfaceRelationship) HasOriginalMacAddress() bool`

HasOriginalMacAddress returns a boolean if a field has been set.

### GetPciAddr

`func (o *AdapterHostEthInterfaceRelationship) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *AdapterHostEthInterfaceRelationship) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *AdapterHostEthInterfaceRelationship) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *AdapterHostEthInterfaceRelationship) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostEthInterfaceRelationship) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostEthInterfaceRelationship) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostEthInterfaceRelationship) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostEthInterfaceRelationship) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetVirtualizationPreference

`func (o *AdapterHostEthInterfaceRelationship) GetVirtualizationPreference() string`

GetVirtualizationPreference returns the VirtualizationPreference field if non-nil, zero value otherwise.

### GetVirtualizationPreferenceOk

`func (o *AdapterHostEthInterfaceRelationship) GetVirtualizationPreferenceOk() (*string, bool)`

GetVirtualizationPreferenceOk returns a tuple with the VirtualizationPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationPreference

`func (o *AdapterHostEthInterfaceRelationship) SetVirtualizationPreference(v string)`

SetVirtualizationPreference sets VirtualizationPreference field to given value.

### HasVirtualizationPreference

`func (o *AdapterHostEthInterfaceRelationship) HasVirtualizationPreference() bool`

HasVirtualizationPreference returns a boolean if a field has been set.

### GetVnicDn

`func (o *AdapterHostEthInterfaceRelationship) GetVnicDn() string`

GetVnicDn returns the VnicDn field if non-nil, zero value otherwise.

### GetVnicDnOk

`func (o *AdapterHostEthInterfaceRelationship) GetVnicDnOk() (*string, bool)`

GetVnicDnOk returns a tuple with the VnicDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicDn

`func (o *AdapterHostEthInterfaceRelationship) SetVnicDn(v string)`

SetVnicDn sets VnicDn field to given value.

### HasVnicDn

`func (o *AdapterHostEthInterfaceRelationship) HasVnicDn() bool`

HasVnicDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostEthInterfaceRelationship) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostEthInterfaceRelationship) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostEthInterfaceRelationship) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostEthInterfaceRelationship) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostEthInterfaceRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostEthInterfaceRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostEthInterfaceRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostEthInterfaceRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostEthInterfaceRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostEthInterfaceRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


