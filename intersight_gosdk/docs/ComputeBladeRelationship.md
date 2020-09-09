# ComputeBladeRelationship

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
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**AdminPowerState** | Pointer to **string** | The desired power state of the server. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**ComputeAlarmSummary**](compute.AlarmSummary.md) |  | [optional] 
**AssetTag** | Pointer to **string** | The user defined asset tag assigned to the server. | [optional] [readonly] 
**AvailableMemory** | Pointer to **int64** | The amount of memory available on the server. | [optional] [readonly] 
**BiosPostComplete** | Pointer to **bool** | The BIOS POST completion status of the server. | [optional] 
**FaultSummary** | Pointer to **int64** | The fault summary for the server. | [optional] 
**KvmIpAddresses** | Pointer to [**[]ComputeIpAddress**](compute.IpAddress.md) |  | [optional] 
**ManagementMode** | Pointer to **string** | The management mode of the server. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [default to "IntersightStandalone"]
**MemorySpeed** | Pointer to **string** | The maximum memory speed in MHz available on the server. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | Management address of the server. | [optional] 
**NumAdaptors** | Pointer to **int64** | The total number of network adapters present on the server. | [optional] [readonly] 
**NumCpuCores** | Pointer to **int64** | The total number of CPU cores present on the server. | [optional] [readonly] 
**NumCpuCoresEnabled** | Pointer to **int64** | The total number of CPU cores enabled on the server. | [optional] [readonly] 
**NumCpus** | Pointer to **int64** | The total number of CPUs present on the server. | [optional] [readonly] 
**NumEthHostInterfaces** | Pointer to **int64** | The total number of vNICs which are visible to a host on the server. | [optional] [readonly] 
**NumFcHostInterfaces** | Pointer to **int64** | The total number of vHBAs which are visible to a host on the server. | [optional] [readonly] 
**NumThreads** | Pointer to **int64** | The total number of threads the server is capable of handling. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The actual power state of the server. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the server. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability of the server. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the registered device - whether managed by UCSM or operating in standalone mode. | [optional] 
**Presence** | Pointer to **string** | Indicates if a server is present in a slot and is applicable for blade servers. | [optional] [readonly] 
**ServiceProfile** | Pointer to **string** | The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. | [optional] [readonly] 
**TotalMemory** | Pointer to **int64** | The total memory available on the server. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the server. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The universally unique identity of the server. | [optional] [readonly] 
**MgmtIdentity** | Pointer to [**EquipmentPhysicalIdentityRelationship**](equipment.PhysicalIdentity.Relationship.md) |  | [optional] 
**ChassisId** | Pointer to **string** | The id of the chassis that the blade is located in. | [optional] [readonly] 
**ScaledMode** | Pointer to **string** | The mode of the server that determines it is scaled. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | The slot number in the chassis that the blade is located in. | [optional] [readonly] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](adapter.Unit.Relationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosUnits** | Pointer to [**[]BiosUnitRelationship**](bios.Unit.Relationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**Bmc** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**Board** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**EquipmentIoExpanders** | Pointer to [**[]EquipmentIoExpanderRelationship**](equipment.IoExpander.Relationship.md) | An array of relationships to equipmentIoExpander resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](inventory.GenericInventoryHolder.Relationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](pci.Device.Relationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**TopSystemRelationship**](top.System.Relationship.md) |  | [optional] 

## Methods

### NewComputeBladeRelationship

`func NewComputeBladeRelationship(classId string, objectType string, ) *ComputeBladeRelationship`

NewComputeBladeRelationship instantiates a new ComputeBladeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBladeRelationshipWithDefaults

`func NewComputeBladeRelationshipWithDefaults() *ComputeBladeRelationship`

NewComputeBladeRelationshipWithDefaults instantiates a new ComputeBladeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ComputeBladeRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ComputeBladeRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ComputeBladeRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ComputeBladeRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ComputeBladeRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBladeRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBladeRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ComputeBladeRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ComputeBladeRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ComputeBladeRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ComputeBladeRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ComputeBladeRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ComputeBladeRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ComputeBladeRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ComputeBladeRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ComputeBladeRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ComputeBladeRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ComputeBladeRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ComputeBladeRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ComputeBladeRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ComputeBladeRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ComputeBladeRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ComputeBladeRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ComputeBladeRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBladeRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBladeRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ComputeBladeRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ComputeBladeRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ComputeBladeRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ComputeBladeRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ComputeBladeRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ComputeBladeRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ComputeBladeRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ComputeBladeRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ComputeBladeRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComputeBladeRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComputeBladeRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ComputeBladeRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ComputeBladeRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ComputeBladeRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ComputeBladeRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ComputeBladeRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ComputeBladeRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ComputeBladeRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ComputeBladeRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ComputeBladeRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ComputeBladeRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ComputeBladeRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ComputeBladeRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ComputeBladeRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ComputeBladeRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ComputeBladeRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ComputeBladeRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ComputeBladeRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ComputeBladeRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ComputeBladeRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ComputeBladeRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ComputeBladeRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ComputeBladeRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ComputeBladeRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ComputeBladeRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ComputeBladeRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ComputeBladeRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ComputeBladeRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ComputeBladeRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputeBladeRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputeBladeRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputeBladeRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ComputeBladeRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputeBladeRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputeBladeRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputeBladeRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ComputeBladeRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ComputeBladeRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ComputeBladeRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ComputeBladeRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *ComputeBladeRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeBladeRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeBladeRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeBladeRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ComputeBladeRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputeBladeRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputeBladeRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputeBladeRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *ComputeBladeRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ComputeBladeRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ComputeBladeRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ComputeBladeRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *ComputeBladeRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputeBladeRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputeBladeRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputeBladeRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdminPowerState

`func (o *ComputeBladeRelationship) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputeBladeRelationship) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputeBladeRelationship) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputeBladeRelationship) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputeBladeRelationship) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputeBladeRelationship) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputeBladeRelationship) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputeBladeRelationship) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputeBladeRelationship) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputeBladeRelationship) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputeBladeRelationship) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputeBladeRelationship) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputeBladeRelationship) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputeBladeRelationship) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputeBladeRelationship) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputeBladeRelationship) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputeBladeRelationship) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputeBladeRelationship) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputeBladeRelationship) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputeBladeRelationship) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputeBladeRelationship) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputeBladeRelationship) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputeBladeRelationship) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputeBladeRelationship) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputeBladeRelationship) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputeBladeRelationship) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputeBladeRelationship) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputeBladeRelationship) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputeBladeRelationship) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputeBladeRelationship) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputeBladeRelationship) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputeBladeRelationship) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputeBladeRelationship) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputeBladeRelationship) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputeBladeRelationship) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputeBladeRelationship) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputeBladeRelationship) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputeBladeRelationship) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputeBladeRelationship) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputeBladeRelationship) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputeBladeRelationship) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputeBladeRelationship) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputeBladeRelationship) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputeBladeRelationship) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputeBladeRelationship) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputeBladeRelationship) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputeBladeRelationship) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputeBladeRelationship) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputeBladeRelationship) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputeBladeRelationship) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputeBladeRelationship) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputeBladeRelationship) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputeBladeRelationship) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputeBladeRelationship) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputeBladeRelationship) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputeBladeRelationship) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputeBladeRelationship) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputeBladeRelationship) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputeBladeRelationship) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputeBladeRelationship) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputeBladeRelationship) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputeBladeRelationship) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputeBladeRelationship) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputeBladeRelationship) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputeBladeRelationship) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputeBladeRelationship) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputeBladeRelationship) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputeBladeRelationship) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputeBladeRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputeBladeRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputeBladeRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputeBladeRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *ComputeBladeRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputeBladeRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputeBladeRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputeBladeRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputeBladeRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputeBladeRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputeBladeRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputeBladeRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputeBladeRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputeBladeRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputeBladeRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputeBladeRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPresence

`func (o *ComputeBladeRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputeBladeRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputeBladeRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputeBladeRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputeBladeRelationship) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputeBladeRelationship) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputeBladeRelationship) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputeBladeRelationship) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputeBladeRelationship) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputeBladeRelationship) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputeBladeRelationship) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputeBladeRelationship) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputeBladeRelationship) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputeBladeRelationship) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputeBladeRelationship) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputeBladeRelationship) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputeBladeRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputeBladeRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputeBladeRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputeBladeRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMgmtIdentity

`func (o *ComputeBladeRelationship) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship`

GetMgmtIdentity returns the MgmtIdentity field if non-nil, zero value otherwise.

### GetMgmtIdentityOk

`func (o *ComputeBladeRelationship) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetMgmtIdentityOk returns a tuple with the MgmtIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIdentity

`func (o *ComputeBladeRelationship) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship)`

SetMgmtIdentity sets MgmtIdentity field to given value.

### HasMgmtIdentity

`func (o *ComputeBladeRelationship) HasMgmtIdentity() bool`

HasMgmtIdentity returns a boolean if a field has been set.

### GetChassisId

`func (o *ComputeBladeRelationship) GetChassisId() string`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *ComputeBladeRelationship) GetChassisIdOk() (*string, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *ComputeBladeRelationship) SetChassisId(v string)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *ComputeBladeRelationship) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetScaledMode

`func (o *ComputeBladeRelationship) GetScaledMode() string`

GetScaledMode returns the ScaledMode field if non-nil, zero value otherwise.

### GetScaledModeOk

`func (o *ComputeBladeRelationship) GetScaledModeOk() (*string, bool)`

GetScaledModeOk returns a tuple with the ScaledMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaledMode

`func (o *ComputeBladeRelationship) SetScaledMode(v string)`

SetScaledMode sets ScaledMode field to given value.

### HasScaledMode

`func (o *ComputeBladeRelationship) HasScaledMode() bool`

HasScaledMode returns a boolean if a field has been set.

### GetSlotId

`func (o *ComputeBladeRelationship) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *ComputeBladeRelationship) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *ComputeBladeRelationship) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *ComputeBladeRelationship) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeBladeRelationship) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeBladeRelationship) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeBladeRelationship) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeBladeRelationship) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeBladeRelationship) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeBladeRelationship) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosUnits

`func (o *ComputeBladeRelationship) GetBiosUnits() []BiosUnitRelationship`

GetBiosUnits returns the BiosUnits field if non-nil, zero value otherwise.

### GetBiosUnitsOk

`func (o *ComputeBladeRelationship) GetBiosUnitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosUnitsOk returns a tuple with the BiosUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnits

`func (o *ComputeBladeRelationship) SetBiosUnits(v []BiosUnitRelationship)`

SetBiosUnits sets BiosUnits field to given value.

### HasBiosUnits

`func (o *ComputeBladeRelationship) HasBiosUnits() bool`

HasBiosUnits returns a boolean if a field has been set.

### SetBiosUnitsNil

`func (o *ComputeBladeRelationship) SetBiosUnitsNil(b bool)`

 SetBiosUnitsNil sets the value for BiosUnits to be an explicit nil

### UnsetBiosUnits
`func (o *ComputeBladeRelationship) UnsetBiosUnits()`

UnsetBiosUnits ensures that no value is present for BiosUnits, not even an explicit nil
### GetBmc

`func (o *ComputeBladeRelationship) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeBladeRelationship) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeBladeRelationship) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeBladeRelationship) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### GetBoard

`func (o *ComputeBladeRelationship) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeBladeRelationship) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeBladeRelationship) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeBladeRelationship) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *ComputeBladeRelationship) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *ComputeBladeRelationship) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *ComputeBladeRelationship) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *ComputeBladeRelationship) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetEquipmentIoExpanders

`func (o *ComputeBladeRelationship) GetEquipmentIoExpanders() []EquipmentIoExpanderRelationship`

GetEquipmentIoExpanders returns the EquipmentIoExpanders field if non-nil, zero value otherwise.

### GetEquipmentIoExpandersOk

`func (o *ComputeBladeRelationship) GetEquipmentIoExpandersOk() (*[]EquipmentIoExpanderRelationship, bool)`

GetEquipmentIoExpandersOk returns a tuple with the EquipmentIoExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIoExpanders

`func (o *ComputeBladeRelationship) SetEquipmentIoExpanders(v []EquipmentIoExpanderRelationship)`

SetEquipmentIoExpanders sets EquipmentIoExpanders field to given value.

### HasEquipmentIoExpanders

`func (o *ComputeBladeRelationship) HasEquipmentIoExpanders() bool`

HasEquipmentIoExpanders returns a boolean if a field has been set.

### SetEquipmentIoExpandersNil

`func (o *ComputeBladeRelationship) SetEquipmentIoExpandersNil(b bool)`

 SetEquipmentIoExpandersNil sets the value for EquipmentIoExpanders to be an explicit nil

### UnsetEquipmentIoExpanders
`func (o *ComputeBladeRelationship) UnsetEquipmentIoExpanders()`

UnsetEquipmentIoExpanders ensures that no value is present for EquipmentIoExpanders, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeBladeRelationship) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeBladeRelationship) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeBladeRelationship) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeBladeRelationship) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeBladeRelationship) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeBladeRelationship) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeBladeRelationship) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeBladeRelationship) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeBladeRelationship) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeBladeRelationship) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeBladeRelationship) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeBladeRelationship) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeBladeRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeBladeRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeBladeRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeBladeRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *ComputeBladeRelationship) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeBladeRelationship) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeBladeRelationship) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeBladeRelationship) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeBladeRelationship) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeBladeRelationship) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeBladeRelationship) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeBladeRelationship) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeBladeRelationship) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeBladeRelationship) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeBladeRelationship) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeBladeRelationship) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeBladeRelationship) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeBladeRelationship) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeBladeRelationship) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeBladeRelationship) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetProcessors

`func (o *ComputeBladeRelationship) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeBladeRelationship) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeBladeRelationship) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeBladeRelationship) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeBladeRelationship) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeBladeRelationship) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetRegisteredDevice

`func (o *ComputeBladeRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeBladeRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeBladeRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeBladeRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ComputeBladeRelationship) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeBladeRelationship) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeBladeRelationship) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeBladeRelationship) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeBladeRelationship) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeBladeRelationship) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeBladeRelationship) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeBladeRelationship) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeBladeRelationship) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeBladeRelationship) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeBladeRelationship) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeBladeRelationship) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeBladeRelationship) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeBladeRelationship) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeBladeRelationship) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeBladeRelationship) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


