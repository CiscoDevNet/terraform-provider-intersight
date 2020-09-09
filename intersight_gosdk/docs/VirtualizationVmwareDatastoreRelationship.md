# VirtualizationVmwareDatastoreRelationship

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
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Capacity** | Pointer to [**VirtualizationStorageCapacity**](virtualization.StorageCapacity.md) |  | [optional] 
**HostCount** | Pointer to **int64** | Number of hosts attached to or supported-by this datastore. | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference). | [optional] 
**Name** | Pointer to **string** | Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations. | [optional] 
**Type** | Pointer to **string** | A string indicating the type of the datastore (VMFS, NFS, etc). * &#x60;Unknown&#x60; - The nature of the file system is unknown. * &#x60;VMFS&#x60; - It is a Virtual Machine Filesystem. * &#x60;NFS&#x60; - It is a Network File System. * &#x60;vSAN&#x60; - It is a virtual Storage Area Network file system. * &#x60;VirtualVolume&#x60; - A Virtual Volume datastore represents a storage container in a hypervisor server. | [optional] [default to "Unknown"]
**VmCount** | Pointer to **int64** | Number of virtual machines relying on (using) this datastore. | [optional] 
**Accessible** | Pointer to **bool** | Shows if this datastore is accessible. | [optional] 
**MaintenanceMode** | Pointer to **bool** | Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode. | [optional] 
**MultipleHostAccess** | Pointer to **bool** | Indicates if this datastore is connected to multiple hosts. | [optional] 
**Status** | Pointer to **string** | Datastore health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**ThinProvisioningSupported** | Pointer to **bool** | Indicates if this datastore supports thin provisioning for files. | [optional] 
**UnCommitted** | Pointer to **int64** | Space uncommitted in this datastore in bytes. | [optional] 
**Url** | Pointer to **string** | The URL to access this datastore (example - &#39;ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/&#39;). | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) | An array of relationships to virtualizationVmwareHost resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareDatastoreRelationship

`func NewVirtualizationVmwareDatastoreRelationship(classId string, objectType string, ) *VirtualizationVmwareDatastoreRelationship`

NewVirtualizationVmwareDatastoreRelationship instantiates a new VirtualizationVmwareDatastoreRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDatastoreRelationshipWithDefaults

`func NewVirtualizationVmwareDatastoreRelationshipWithDefaults() *VirtualizationVmwareDatastoreRelationship`

NewVirtualizationVmwareDatastoreRelationshipWithDefaults instantiates a new VirtualizationVmwareDatastoreRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VirtualizationVmwareDatastoreRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VirtualizationVmwareDatastoreRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VirtualizationVmwareDatastoreRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VirtualizationVmwareDatastoreRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDatastoreRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VirtualizationVmwareDatastoreRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VirtualizationVmwareDatastoreRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VirtualizationVmwareDatastoreRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VirtualizationVmwareDatastoreRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VirtualizationVmwareDatastoreRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VirtualizationVmwareDatastoreRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VirtualizationVmwareDatastoreRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VirtualizationVmwareDatastoreRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VirtualizationVmwareDatastoreRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VirtualizationVmwareDatastoreRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VirtualizationVmwareDatastoreRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VirtualizationVmwareDatastoreRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VirtualizationVmwareDatastoreRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDatastoreRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VirtualizationVmwareDatastoreRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VirtualizationVmwareDatastoreRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VirtualizationVmwareDatastoreRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VirtualizationVmwareDatastoreRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VirtualizationVmwareDatastoreRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VirtualizationVmwareDatastoreRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VirtualizationVmwareDatastoreRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualizationVmwareDatastoreRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualizationVmwareDatastoreRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VirtualizationVmwareDatastoreRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VirtualizationVmwareDatastoreRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VirtualizationVmwareDatastoreRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VirtualizationVmwareDatastoreRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VirtualizationVmwareDatastoreRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VirtualizationVmwareDatastoreRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VirtualizationVmwareDatastoreRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VirtualizationVmwareDatastoreRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VirtualizationVmwareDatastoreRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VirtualizationVmwareDatastoreRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VirtualizationVmwareDatastoreRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VirtualizationVmwareDatastoreRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VirtualizationVmwareDatastoreRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VirtualizationVmwareDatastoreRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VirtualizationVmwareDatastoreRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VirtualizationVmwareDatastoreRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VirtualizationVmwareDatastoreRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VirtualizationVmwareDatastoreRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VirtualizationVmwareDatastoreRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VirtualizationVmwareDatastoreRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VirtualizationVmwareDatastoreRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetRegisteredDevice

`func (o *VirtualizationVmwareDatastoreRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVmwareDatastoreRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVmwareDatastoreRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationVmwareDatastoreRelationship) GetCapacity() VirtualizationStorageCapacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVmwareDatastoreRelationship) SetCapacity(v VirtualizationStorageCapacity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVmwareDatastoreRelationship) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationVmwareDatastoreRelationship) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationVmwareDatastoreRelationship) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationVmwareDatastoreRelationship) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationVmwareDatastoreRelationship) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationVmwareDatastoreRelationship) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationVmwareDatastoreRelationship) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmwareDatastoreRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmwareDatastoreRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmwareDatastoreRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVmwareDatastoreRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVmwareDatastoreRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVmwareDatastoreRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationVmwareDatastoreRelationship) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationVmwareDatastoreRelationship) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationVmwareDatastoreRelationship) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetAccessible

`func (o *VirtualizationVmwareDatastoreRelationship) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *VirtualizationVmwareDatastoreRelationship) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *VirtualizationVmwareDatastoreRelationship) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationVmwareDatastoreRelationship) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationVmwareDatastoreRelationship) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationVmwareDatastoreRelationship) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreRelationship) GetMultipleHostAccess() bool`

GetMultipleHostAccess returns the MultipleHostAccess field if non-nil, zero value otherwise.

### GetMultipleHostAccessOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetMultipleHostAccessOk() (*bool, bool)`

GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreRelationship) SetMultipleHostAccess(v bool)`

SetMultipleHostAccess sets MultipleHostAccess field to given value.

### HasMultipleHostAccess

`func (o *VirtualizationVmwareDatastoreRelationship) HasMultipleHostAccess() bool`

HasMultipleHostAccess returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVmwareDatastoreRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVmwareDatastoreRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVmwareDatastoreRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreRelationship) GetThinProvisioningSupported() bool`

GetThinProvisioningSupported returns the ThinProvisioningSupported field if non-nil, zero value otherwise.

### GetThinProvisioningSupportedOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetThinProvisioningSupportedOk() (*bool, bool)`

GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreRelationship) SetThinProvisioningSupported(v bool)`

SetThinProvisioningSupported sets ThinProvisioningSupported field to given value.

### HasThinProvisioningSupported

`func (o *VirtualizationVmwareDatastoreRelationship) HasThinProvisioningSupported() bool`

HasThinProvisioningSupported returns a boolean if a field has been set.

### GetUnCommitted

`func (o *VirtualizationVmwareDatastoreRelationship) GetUnCommitted() int64`

GetUnCommitted returns the UnCommitted field if non-nil, zero value otherwise.

### GetUnCommittedOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetUnCommittedOk() (*int64, bool)`

GetUnCommittedOk returns a tuple with the UnCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCommitted

`func (o *VirtualizationVmwareDatastoreRelationship) SetUnCommitted(v int64)`

SetUnCommitted sets UnCommitted field to given value.

### HasUnCommitted

`func (o *VirtualizationVmwareDatastoreRelationship) HasUnCommitted() bool`

HasUnCommitted returns a boolean if a field has been set.

### GetUrl

`func (o *VirtualizationVmwareDatastoreRelationship) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualizationVmwareDatastoreRelationship) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VirtualizationVmwareDatastoreRelationship) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareDatastoreRelationship) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareDatastoreRelationship) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareDatastoreRelationship) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDatastoreRelationship) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDatastoreRelationship) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDatastoreRelationship) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationVmwareDatastoreRelationship) GetHosts() []VirtualizationVmwareHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationVmwareDatastoreRelationship) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationVmwareDatastoreRelationship) SetHosts(v []VirtualizationVmwareHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationVmwareDatastoreRelationship) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationVmwareDatastoreRelationship) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationVmwareDatastoreRelationship) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


