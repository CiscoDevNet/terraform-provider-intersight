# VirtualizationVmwareHostRelationship

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
**CpuInfo** | Pointer to [**VirtualizationCpuInfo**](virtualization.CpuInfo.md) |  | [optional] 
**HardwareInfo** | Pointer to [**InfraHardwareInfo**](infra.HardwareInfo.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | Identifies the broad type of the underlying hypervisor. * &#x60;Unknown&#x60; - The type of the hypervisor is unknown. * &#x60;ESXi&#x60; - A Vmware ESXi hypervisor of any version. * &#x60;HXAP&#x60; - A Cisco HyperFlex Application Platform hypervisor. | [optional] [default to "Unknown"]
**Identity** | Pointer to **string** | The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). | [optional] 
**MaintenanceMode** | Pointer to **bool** | Is this host in maintenance mode. Set to true or false. | [optional] 
**MemoryCapacity** | Pointer to [**VirtualizationMemoryCapacity**](virtualization.MemoryCapacity.md) |  | [optional] 
**Model** | Pointer to **string** | Commercial model information about this hardware. | [optional] 
**Name** | Pointer to **string** | Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. | [optional] 
**ProcessorCapacity** | Pointer to [**VirtualizationComputeCapacity**](virtualization.ComputeCapacity.md) |  | [optional] 
**ProductInfo** | Pointer to [**VirtualizationProductInfo**](virtualization.ProductInfo.md) |  | [optional] 
**Serial** | Pointer to **string** | Serial number of this host (internally generated). | [optional] 
**Status** | Pointer to **string** | Host health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**UpTime** | Pointer to **string** | The uptime of the host, stored as Duration (from w3c). | [optional] 
**Uuid** | Pointer to **string** | Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. | [optional] 
**Vendor** | Pointer to **string** | Commercial vendor details of this hardware. | [optional] 
**BootTime** | Pointer to [**time.Time**](time.Time.md) | The time when this host booted up. | [optional] 
**ConnectionState** | Pointer to **string** | Indicates if the host is connected to the vCenter. Values are connected, not connected. | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. | [optional] [default to "Unknown"]
**NetworkAdapterCount** | Pointer to **int64** | The count of all network adapters attached to this host. | [optional] 
**ResourceConsumed** | Pointer to [**VirtualizationVmwareResourceConsumption**](virtualization.VmwareResourceConsumption.md) |  | [optional] 
**StorageAdapterCount** | Pointer to **int64** | The count of all storage adapters attached to this host. | [optional] 
**VcenterHostId** | Pointer to **string** | The identity of this host within vCenter (optional). | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Datastores** | Pointer to [**[]VirtualizationVmwareDatastoreRelationship**](virtualization.VmwareDatastore.Relationship.md) | An array of relationships to virtualizationVmwareDatastore resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareHostRelationship

`func NewVirtualizationVmwareHostRelationship(classId string, objectType string, ) *VirtualizationVmwareHostRelationship`

NewVirtualizationVmwareHostRelationship instantiates a new VirtualizationVmwareHostRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareHostRelationshipWithDefaults

`func NewVirtualizationVmwareHostRelationshipWithDefaults() *VirtualizationVmwareHostRelationship`

NewVirtualizationVmwareHostRelationshipWithDefaults instantiates a new VirtualizationVmwareHostRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VirtualizationVmwareHostRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VirtualizationVmwareHostRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VirtualizationVmwareHostRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VirtualizationVmwareHostRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VirtualizationVmwareHostRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareHostRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareHostRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VirtualizationVmwareHostRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VirtualizationVmwareHostRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VirtualizationVmwareHostRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VirtualizationVmwareHostRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VirtualizationVmwareHostRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VirtualizationVmwareHostRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VirtualizationVmwareHostRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VirtualizationVmwareHostRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VirtualizationVmwareHostRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VirtualizationVmwareHostRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VirtualizationVmwareHostRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VirtualizationVmwareHostRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VirtualizationVmwareHostRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VirtualizationVmwareHostRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VirtualizationVmwareHostRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VirtualizationVmwareHostRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VirtualizationVmwareHostRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareHostRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareHostRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VirtualizationVmwareHostRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VirtualizationVmwareHostRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VirtualizationVmwareHostRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VirtualizationVmwareHostRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VirtualizationVmwareHostRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VirtualizationVmwareHostRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VirtualizationVmwareHostRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VirtualizationVmwareHostRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VirtualizationVmwareHostRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualizationVmwareHostRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualizationVmwareHostRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualizationVmwareHostRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VirtualizationVmwareHostRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VirtualizationVmwareHostRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VirtualizationVmwareHostRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VirtualizationVmwareHostRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VirtualizationVmwareHostRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VirtualizationVmwareHostRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VirtualizationVmwareHostRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VirtualizationVmwareHostRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VirtualizationVmwareHostRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VirtualizationVmwareHostRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VirtualizationVmwareHostRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VirtualizationVmwareHostRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VirtualizationVmwareHostRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VirtualizationVmwareHostRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VirtualizationVmwareHostRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VirtualizationVmwareHostRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VirtualizationVmwareHostRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VirtualizationVmwareHostRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VirtualizationVmwareHostRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VirtualizationVmwareHostRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VirtualizationVmwareHostRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VirtualizationVmwareHostRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VirtualizationVmwareHostRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VirtualizationVmwareHostRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VirtualizationVmwareHostRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VirtualizationVmwareHostRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetRegisteredDevice

`func (o *VirtualizationVmwareHostRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVmwareHostRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVmwareHostRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVmwareHostRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetCpuInfo

`func (o *VirtualizationVmwareHostRelationship) GetCpuInfo() VirtualizationCpuInfo`

GetCpuInfo returns the CpuInfo field if non-nil, zero value otherwise.

### GetCpuInfoOk

`func (o *VirtualizationVmwareHostRelationship) GetCpuInfoOk() (*VirtualizationCpuInfo, bool)`

GetCpuInfoOk returns a tuple with the CpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuInfo

`func (o *VirtualizationVmwareHostRelationship) SetCpuInfo(v VirtualizationCpuInfo)`

SetCpuInfo sets CpuInfo field to given value.

### HasCpuInfo

`func (o *VirtualizationVmwareHostRelationship) HasCpuInfo() bool`

HasCpuInfo returns a boolean if a field has been set.

### GetHardwareInfo

`func (o *VirtualizationVmwareHostRelationship) GetHardwareInfo() InfraHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *VirtualizationVmwareHostRelationship) GetHardwareInfoOk() (*InfraHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *VirtualizationVmwareHostRelationship) SetHardwareInfo(v InfraHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *VirtualizationVmwareHostRelationship) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VirtualizationVmwareHostRelationship) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VirtualizationVmwareHostRelationship) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VirtualizationVmwareHostRelationship) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VirtualizationVmwareHostRelationship) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationVmwareHostRelationship) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationVmwareHostRelationship) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationVmwareHostRelationship) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationVmwareHostRelationship) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaintenanceMode

`func (o *VirtualizationVmwareHostRelationship) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualizationVmwareHostRelationship) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualizationVmwareHostRelationship) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualizationVmwareHostRelationship) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *VirtualizationVmwareHostRelationship) GetMemoryCapacity() VirtualizationMemoryCapacity`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VirtualizationVmwareHostRelationship) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VirtualizationVmwareHostRelationship) SetMemoryCapacity(v VirtualizationMemoryCapacity)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *VirtualizationVmwareHostRelationship) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetModel

`func (o *VirtualizationVmwareHostRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VirtualizationVmwareHostRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VirtualizationVmwareHostRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VirtualizationVmwareHostRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmwareHostRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmwareHostRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmwareHostRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmwareHostRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessorCapacity

`func (o *VirtualizationVmwareHostRelationship) GetProcessorCapacity() VirtualizationComputeCapacity`

GetProcessorCapacity returns the ProcessorCapacity field if non-nil, zero value otherwise.

### GetProcessorCapacityOk

`func (o *VirtualizationVmwareHostRelationship) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool)`

GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCapacity

`func (o *VirtualizationVmwareHostRelationship) SetProcessorCapacity(v VirtualizationComputeCapacity)`

SetProcessorCapacity sets ProcessorCapacity field to given value.

### HasProcessorCapacity

`func (o *VirtualizationVmwareHostRelationship) HasProcessorCapacity() bool`

HasProcessorCapacity returns a boolean if a field has been set.

### GetProductInfo

`func (o *VirtualizationVmwareHostRelationship) GetProductInfo() VirtualizationProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *VirtualizationVmwareHostRelationship) GetProductInfoOk() (*VirtualizationProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *VirtualizationVmwareHostRelationship) SetProductInfo(v VirtualizationProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *VirtualizationVmwareHostRelationship) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualizationVmwareHostRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationVmwareHostRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationVmwareHostRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationVmwareHostRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVmwareHostRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVmwareHostRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVmwareHostRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVmwareHostRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpTime

`func (o *VirtualizationVmwareHostRelationship) GetUpTime() string`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *VirtualizationVmwareHostRelationship) GetUpTimeOk() (*string, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *VirtualizationVmwareHostRelationship) SetUpTime(v string)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *VirtualizationVmwareHostRelationship) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationVmwareHostRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVmwareHostRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVmwareHostRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVmwareHostRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationVmwareHostRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationVmwareHostRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationVmwareHostRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationVmwareHostRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetBootTime

`func (o *VirtualizationVmwareHostRelationship) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationVmwareHostRelationship) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationVmwareHostRelationship) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationVmwareHostRelationship) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetConnectionState

`func (o *VirtualizationVmwareHostRelationship) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VirtualizationVmwareHostRelationship) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VirtualizationVmwareHostRelationship) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *VirtualizationVmwareHostRelationship) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetHwPowerState

`func (o *VirtualizationVmwareHostRelationship) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *VirtualizationVmwareHostRelationship) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *VirtualizationVmwareHostRelationship) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *VirtualizationVmwareHostRelationship) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetNetworkAdapterCount

`func (o *VirtualizationVmwareHostRelationship) GetNetworkAdapterCount() int64`

GetNetworkAdapterCount returns the NetworkAdapterCount field if non-nil, zero value otherwise.

### GetNetworkAdapterCountOk

`func (o *VirtualizationVmwareHostRelationship) GetNetworkAdapterCountOk() (*int64, bool)`

GetNetworkAdapterCountOk returns a tuple with the NetworkAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterCount

`func (o *VirtualizationVmwareHostRelationship) SetNetworkAdapterCount(v int64)`

SetNetworkAdapterCount sets NetworkAdapterCount field to given value.

### HasNetworkAdapterCount

`func (o *VirtualizationVmwareHostRelationship) HasNetworkAdapterCount() bool`

HasNetworkAdapterCount returns a boolean if a field has been set.

### GetResourceConsumed

`func (o *VirtualizationVmwareHostRelationship) GetResourceConsumed() VirtualizationVmwareResourceConsumption`

GetResourceConsumed returns the ResourceConsumed field if non-nil, zero value otherwise.

### GetResourceConsumedOk

`func (o *VirtualizationVmwareHostRelationship) GetResourceConsumedOk() (*VirtualizationVmwareResourceConsumption, bool)`

GetResourceConsumedOk returns a tuple with the ResourceConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConsumed

`func (o *VirtualizationVmwareHostRelationship) SetResourceConsumed(v VirtualizationVmwareResourceConsumption)`

SetResourceConsumed sets ResourceConsumed field to given value.

### HasResourceConsumed

`func (o *VirtualizationVmwareHostRelationship) HasResourceConsumed() bool`

HasResourceConsumed returns a boolean if a field has been set.

### GetStorageAdapterCount

`func (o *VirtualizationVmwareHostRelationship) GetStorageAdapterCount() int64`

GetStorageAdapterCount returns the StorageAdapterCount field if non-nil, zero value otherwise.

### GetStorageAdapterCountOk

`func (o *VirtualizationVmwareHostRelationship) GetStorageAdapterCountOk() (*int64, bool)`

GetStorageAdapterCountOk returns a tuple with the StorageAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAdapterCount

`func (o *VirtualizationVmwareHostRelationship) SetStorageAdapterCount(v int64)`

SetStorageAdapterCount sets StorageAdapterCount field to given value.

### HasStorageAdapterCount

`func (o *VirtualizationVmwareHostRelationship) HasStorageAdapterCount() bool`

HasStorageAdapterCount returns a boolean if a field has been set.

### GetVcenterHostId

`func (o *VirtualizationVmwareHostRelationship) GetVcenterHostId() string`

GetVcenterHostId returns the VcenterHostId field if non-nil, zero value otherwise.

### GetVcenterHostIdOk

`func (o *VirtualizationVmwareHostRelationship) GetVcenterHostIdOk() (*string, bool)`

GetVcenterHostIdOk returns a tuple with the VcenterHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterHostId

`func (o *VirtualizationVmwareHostRelationship) SetVcenterHostId(v string)`

SetVcenterHostId sets VcenterHostId field to given value.

### HasVcenterHostId

`func (o *VirtualizationVmwareHostRelationship) HasVcenterHostId() bool`

HasVcenterHostId returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareHostRelationship) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareHostRelationship) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareHostRelationship) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareHostRelationship) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareHostRelationship) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareHostRelationship) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareHostRelationship) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareHostRelationship) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatastores

`func (o *VirtualizationVmwareHostRelationship) GetDatastores() []VirtualizationVmwareDatastoreRelationship`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualizationVmwareHostRelationship) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualizationVmwareHostRelationship) SetDatastores(v []VirtualizationVmwareDatastoreRelationship)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VirtualizationVmwareHostRelationship) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *VirtualizationVmwareHostRelationship) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *VirtualizationVmwareHostRelationship) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


