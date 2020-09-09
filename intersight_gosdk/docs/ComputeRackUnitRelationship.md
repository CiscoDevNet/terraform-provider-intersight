# ComputeRackUnitRelationship

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
**ConnectionStatus** | Pointer to **string** | Connectivity Status of RackUnit to Switch - A or B or AB. | [optional] [readonly] 
**ServerId** | Pointer to **int64** | RackUnit ID that uniquely identifies the server. | [optional] [readonly] 
**TopologyScanStatus** | Pointer to **string** | To maintain the Topology workflow run status. | [optional] 
**Adapters** | Pointer to [**[]AdapterUnitRelationship**](adapter.Unit.Relationship.md) | An array of relationships to adapterUnit resources. | [optional] [readonly] 
**BiosBootmode** | Pointer to [**BiosBootModeRelationship**](bios.BootMode.Relationship.md) |  | [optional] 
**Biosunits** | Pointer to [**[]BiosUnitRelationship**](bios.Unit.Relationship.md) | An array of relationships to biosUnit resources. | [optional] [readonly] 
**Bmc** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**Board** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**BootDeviceBootmode** | Pointer to [**BootDeviceBootModeRelationship**](boot.DeviceBootMode.Relationship.md) |  | [optional] 
**Fanmodules** | Pointer to [**[]EquipmentFanModuleRelationship**](equipment.FanModule.Relationship.md) | An array of relationships to equipmentFanModule resources. | [optional] [readonly] 
**GenericInventoryHolders** | Pointer to [**[]InventoryGenericInventoryHolderRelationship**](inventory.GenericInventoryHolder.Relationship.md) | An array of relationships to inventoryGenericInventoryHolder resources. | [optional] [readonly] 
**GraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](graphics.Card.Relationship.md) | An array of relationships to graphicsCard resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**MemoryArrays** | Pointer to [**[]MemoryArrayRelationship**](memory.Array.Relationship.md) | An array of relationships to memoryArray resources. | [optional] 
**PciDevices** | Pointer to [**[]PciDeviceRelationship**](pci.Device.Relationship.md) | An array of relationships to pciDevice resources. | [optional] [readonly] 
**Processors** | Pointer to [**[]ProcessorUnitRelationship**](processor.Unit.Relationship.md) | An array of relationships to processorUnit resources. | [optional] 
**Psus** | Pointer to [**[]EquipmentPsuRelationship**](equipment.Psu.Relationship.md) | An array of relationships to equipmentPsu resources. | [optional] [readonly] 
**RackEnclosureSlot** | Pointer to [**EquipmentRackEnclosureSlotRelationship**](equipment.RackEnclosureSlot.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**SasExpanders** | Pointer to [**[]StorageSasExpanderRelationship**](storage.SasExpander.Relationship.md) | An array of relationships to storageSasExpander resources. | [optional] 
**StorageControllers** | Pointer to [**[]StorageControllerRelationship**](storage.Controller.Relationship.md) | An array of relationships to storageController resources. | [optional] 
**StorageEnclosures** | Pointer to [**[]StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) | An array of relationships to storageEnclosure resources. | [optional] [readonly] 
**TopSystem** | Pointer to [**TopSystemRelationship**](top.System.Relationship.md) |  | [optional] 

## Methods

### NewComputeRackUnitRelationship

`func NewComputeRackUnitRelationship(classId string, objectType string, ) *ComputeRackUnitRelationship`

NewComputeRackUnitRelationship instantiates a new ComputeRackUnitRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRackUnitRelationshipWithDefaults

`func NewComputeRackUnitRelationshipWithDefaults() *ComputeRackUnitRelationship`

NewComputeRackUnitRelationshipWithDefaults instantiates a new ComputeRackUnitRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ComputeRackUnitRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ComputeRackUnitRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ComputeRackUnitRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ComputeRackUnitRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ComputeRackUnitRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeRackUnitRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeRackUnitRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ComputeRackUnitRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ComputeRackUnitRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ComputeRackUnitRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ComputeRackUnitRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ComputeRackUnitRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ComputeRackUnitRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ComputeRackUnitRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ComputeRackUnitRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ComputeRackUnitRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ComputeRackUnitRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ComputeRackUnitRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ComputeRackUnitRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ComputeRackUnitRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ComputeRackUnitRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ComputeRackUnitRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ComputeRackUnitRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ComputeRackUnitRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeRackUnitRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeRackUnitRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ComputeRackUnitRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ComputeRackUnitRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ComputeRackUnitRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ComputeRackUnitRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ComputeRackUnitRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ComputeRackUnitRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ComputeRackUnitRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ComputeRackUnitRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ComputeRackUnitRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComputeRackUnitRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComputeRackUnitRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ComputeRackUnitRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ComputeRackUnitRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ComputeRackUnitRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ComputeRackUnitRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ComputeRackUnitRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ComputeRackUnitRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ComputeRackUnitRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ComputeRackUnitRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ComputeRackUnitRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ComputeRackUnitRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ComputeRackUnitRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ComputeRackUnitRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ComputeRackUnitRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ComputeRackUnitRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ComputeRackUnitRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ComputeRackUnitRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ComputeRackUnitRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ComputeRackUnitRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ComputeRackUnitRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ComputeRackUnitRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ComputeRackUnitRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ComputeRackUnitRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ComputeRackUnitRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ComputeRackUnitRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ComputeRackUnitRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ComputeRackUnitRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ComputeRackUnitRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ComputeRackUnitRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputeRackUnitRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputeRackUnitRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputeRackUnitRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ComputeRackUnitRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputeRackUnitRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputeRackUnitRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputeRackUnitRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ComputeRackUnitRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ComputeRackUnitRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ComputeRackUnitRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ComputeRackUnitRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *ComputeRackUnitRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputeRackUnitRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputeRackUnitRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputeRackUnitRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ComputeRackUnitRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputeRackUnitRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputeRackUnitRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputeRackUnitRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *ComputeRackUnitRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ComputeRackUnitRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ComputeRackUnitRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ComputeRackUnitRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *ComputeRackUnitRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputeRackUnitRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputeRackUnitRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputeRackUnitRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdminPowerState

`func (o *ComputeRackUnitRelationship) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputeRackUnitRelationship) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputeRackUnitRelationship) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputeRackUnitRelationship) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputeRackUnitRelationship) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputeRackUnitRelationship) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputeRackUnitRelationship) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputeRackUnitRelationship) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputeRackUnitRelationship) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputeRackUnitRelationship) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputeRackUnitRelationship) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputeRackUnitRelationship) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputeRackUnitRelationship) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputeRackUnitRelationship) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputeRackUnitRelationship) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputeRackUnitRelationship) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputeRackUnitRelationship) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputeRackUnitRelationship) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputeRackUnitRelationship) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputeRackUnitRelationship) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputeRackUnitRelationship) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputeRackUnitRelationship) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputeRackUnitRelationship) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputeRackUnitRelationship) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputeRackUnitRelationship) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputeRackUnitRelationship) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputeRackUnitRelationship) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputeRackUnitRelationship) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputeRackUnitRelationship) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputeRackUnitRelationship) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputeRackUnitRelationship) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputeRackUnitRelationship) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputeRackUnitRelationship) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputeRackUnitRelationship) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputeRackUnitRelationship) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputeRackUnitRelationship) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputeRackUnitRelationship) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputeRackUnitRelationship) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputeRackUnitRelationship) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputeRackUnitRelationship) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputeRackUnitRelationship) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputeRackUnitRelationship) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputeRackUnitRelationship) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputeRackUnitRelationship) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputeRackUnitRelationship) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputeRackUnitRelationship) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputeRackUnitRelationship) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputeRackUnitRelationship) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputeRackUnitRelationship) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputeRackUnitRelationship) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputeRackUnitRelationship) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputeRackUnitRelationship) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputeRackUnitRelationship) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputeRackUnitRelationship) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputeRackUnitRelationship) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputeRackUnitRelationship) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputeRackUnitRelationship) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputeRackUnitRelationship) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputeRackUnitRelationship) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputeRackUnitRelationship) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputeRackUnitRelationship) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputeRackUnitRelationship) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputeRackUnitRelationship) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputeRackUnitRelationship) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputeRackUnitRelationship) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputeRackUnitRelationship) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputeRackUnitRelationship) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputeRackUnitRelationship) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputeRackUnitRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputeRackUnitRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputeRackUnitRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputeRackUnitRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *ComputeRackUnitRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputeRackUnitRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputeRackUnitRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputeRackUnitRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputeRackUnitRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputeRackUnitRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputeRackUnitRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputeRackUnitRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputeRackUnitRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputeRackUnitRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputeRackUnitRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputeRackUnitRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPresence

`func (o *ComputeRackUnitRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputeRackUnitRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputeRackUnitRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputeRackUnitRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputeRackUnitRelationship) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputeRackUnitRelationship) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputeRackUnitRelationship) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputeRackUnitRelationship) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputeRackUnitRelationship) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputeRackUnitRelationship) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputeRackUnitRelationship) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputeRackUnitRelationship) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputeRackUnitRelationship) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputeRackUnitRelationship) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputeRackUnitRelationship) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputeRackUnitRelationship) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputeRackUnitRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputeRackUnitRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputeRackUnitRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputeRackUnitRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMgmtIdentity

`func (o *ComputeRackUnitRelationship) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship`

GetMgmtIdentity returns the MgmtIdentity field if non-nil, zero value otherwise.

### GetMgmtIdentityOk

`func (o *ComputeRackUnitRelationship) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetMgmtIdentityOk returns a tuple with the MgmtIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIdentity

`func (o *ComputeRackUnitRelationship) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship)`

SetMgmtIdentity sets MgmtIdentity field to given value.

### HasMgmtIdentity

`func (o *ComputeRackUnitRelationship) HasMgmtIdentity() bool`

HasMgmtIdentity returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *ComputeRackUnitRelationship) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ComputeRackUnitRelationship) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ComputeRackUnitRelationship) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ComputeRackUnitRelationship) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetServerId

`func (o *ComputeRackUnitRelationship) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ComputeRackUnitRelationship) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ComputeRackUnitRelationship) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ComputeRackUnitRelationship) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetTopologyScanStatus

`func (o *ComputeRackUnitRelationship) GetTopologyScanStatus() string`

GetTopologyScanStatus returns the TopologyScanStatus field if non-nil, zero value otherwise.

### GetTopologyScanStatusOk

`func (o *ComputeRackUnitRelationship) GetTopologyScanStatusOk() (*string, bool)`

GetTopologyScanStatusOk returns a tuple with the TopologyScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyScanStatus

`func (o *ComputeRackUnitRelationship) SetTopologyScanStatus(v string)`

SetTopologyScanStatus sets TopologyScanStatus field to given value.

### HasTopologyScanStatus

`func (o *ComputeRackUnitRelationship) HasTopologyScanStatus() bool`

HasTopologyScanStatus returns a boolean if a field has been set.

### GetAdapters

`func (o *ComputeRackUnitRelationship) GetAdapters() []AdapterUnitRelationship`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *ComputeRackUnitRelationship) GetAdaptersOk() (*[]AdapterUnitRelationship, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *ComputeRackUnitRelationship) SetAdapters(v []AdapterUnitRelationship)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *ComputeRackUnitRelationship) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### SetAdaptersNil

`func (o *ComputeRackUnitRelationship) SetAdaptersNil(b bool)`

 SetAdaptersNil sets the value for Adapters to be an explicit nil

### UnsetAdapters
`func (o *ComputeRackUnitRelationship) UnsetAdapters()`

UnsetAdapters ensures that no value is present for Adapters, not even an explicit nil
### GetBiosBootmode

`func (o *ComputeRackUnitRelationship) GetBiosBootmode() BiosBootModeRelationship`

GetBiosBootmode returns the BiosBootmode field if non-nil, zero value otherwise.

### GetBiosBootmodeOk

`func (o *ComputeRackUnitRelationship) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool)`

GetBiosBootmodeOk returns a tuple with the BiosBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosBootmode

`func (o *ComputeRackUnitRelationship) SetBiosBootmode(v BiosBootModeRelationship)`

SetBiosBootmode sets BiosBootmode field to given value.

### HasBiosBootmode

`func (o *ComputeRackUnitRelationship) HasBiosBootmode() bool`

HasBiosBootmode returns a boolean if a field has been set.

### GetBiosunits

`func (o *ComputeRackUnitRelationship) GetBiosunits() []BiosUnitRelationship`

GetBiosunits returns the Biosunits field if non-nil, zero value otherwise.

### GetBiosunitsOk

`func (o *ComputeRackUnitRelationship) GetBiosunitsOk() (*[]BiosUnitRelationship, bool)`

GetBiosunitsOk returns a tuple with the Biosunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosunits

`func (o *ComputeRackUnitRelationship) SetBiosunits(v []BiosUnitRelationship)`

SetBiosunits sets Biosunits field to given value.

### HasBiosunits

`func (o *ComputeRackUnitRelationship) HasBiosunits() bool`

HasBiosunits returns a boolean if a field has been set.

### SetBiosunitsNil

`func (o *ComputeRackUnitRelationship) SetBiosunitsNil(b bool)`

 SetBiosunitsNil sets the value for Biosunits to be an explicit nil

### UnsetBiosunits
`func (o *ComputeRackUnitRelationship) UnsetBiosunits()`

UnsetBiosunits ensures that no value is present for Biosunits, not even an explicit nil
### GetBmc

`func (o *ComputeRackUnitRelationship) GetBmc() ManagementControllerRelationship`

GetBmc returns the Bmc field if non-nil, zero value otherwise.

### GetBmcOk

`func (o *ComputeRackUnitRelationship) GetBmcOk() (*ManagementControllerRelationship, bool)`

GetBmcOk returns a tuple with the Bmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmc

`func (o *ComputeRackUnitRelationship) SetBmc(v ManagementControllerRelationship)`

SetBmc sets Bmc field to given value.

### HasBmc

`func (o *ComputeRackUnitRelationship) HasBmc() bool`

HasBmc returns a boolean if a field has been set.

### GetBoard

`func (o *ComputeRackUnitRelationship) GetBoard() ComputeBoardRelationship`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *ComputeRackUnitRelationship) GetBoardOk() (*ComputeBoardRelationship, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *ComputeRackUnitRelationship) SetBoard(v ComputeBoardRelationship)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *ComputeRackUnitRelationship) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetBootDeviceBootmode

`func (o *ComputeRackUnitRelationship) GetBootDeviceBootmode() BootDeviceBootModeRelationship`

GetBootDeviceBootmode returns the BootDeviceBootmode field if non-nil, zero value otherwise.

### GetBootDeviceBootmodeOk

`func (o *ComputeRackUnitRelationship) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool)`

GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceBootmode

`func (o *ComputeRackUnitRelationship) SetBootDeviceBootmode(v BootDeviceBootModeRelationship)`

SetBootDeviceBootmode sets BootDeviceBootmode field to given value.

### HasBootDeviceBootmode

`func (o *ComputeRackUnitRelationship) HasBootDeviceBootmode() bool`

HasBootDeviceBootmode returns a boolean if a field has been set.

### GetFanmodules

`func (o *ComputeRackUnitRelationship) GetFanmodules() []EquipmentFanModuleRelationship`

GetFanmodules returns the Fanmodules field if non-nil, zero value otherwise.

### GetFanmodulesOk

`func (o *ComputeRackUnitRelationship) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool)`

GetFanmodulesOk returns a tuple with the Fanmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanmodules

`func (o *ComputeRackUnitRelationship) SetFanmodules(v []EquipmentFanModuleRelationship)`

SetFanmodules sets Fanmodules field to given value.

### HasFanmodules

`func (o *ComputeRackUnitRelationship) HasFanmodules() bool`

HasFanmodules returns a boolean if a field has been set.

### SetFanmodulesNil

`func (o *ComputeRackUnitRelationship) SetFanmodulesNil(b bool)`

 SetFanmodulesNil sets the value for Fanmodules to be an explicit nil

### UnsetFanmodules
`func (o *ComputeRackUnitRelationship) UnsetFanmodules()`

UnsetFanmodules ensures that no value is present for Fanmodules, not even an explicit nil
### GetGenericInventoryHolders

`func (o *ComputeRackUnitRelationship) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship`

GetGenericInventoryHolders returns the GenericInventoryHolders field if non-nil, zero value otherwise.

### GetGenericInventoryHoldersOk

`func (o *ComputeRackUnitRelationship) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool)`

GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericInventoryHolders

`func (o *ComputeRackUnitRelationship) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship)`

SetGenericInventoryHolders sets GenericInventoryHolders field to given value.

### HasGenericInventoryHolders

`func (o *ComputeRackUnitRelationship) HasGenericInventoryHolders() bool`

HasGenericInventoryHolders returns a boolean if a field has been set.

### SetGenericInventoryHoldersNil

`func (o *ComputeRackUnitRelationship) SetGenericInventoryHoldersNil(b bool)`

 SetGenericInventoryHoldersNil sets the value for GenericInventoryHolders to be an explicit nil

### UnsetGenericInventoryHolders
`func (o *ComputeRackUnitRelationship) UnsetGenericInventoryHolders()`

UnsetGenericInventoryHolders ensures that no value is present for GenericInventoryHolders, not even an explicit nil
### GetGraphicsCards

`func (o *ComputeRackUnitRelationship) GetGraphicsCards() []GraphicsCardRelationship`

GetGraphicsCards returns the GraphicsCards field if non-nil, zero value otherwise.

### GetGraphicsCardsOk

`func (o *ComputeRackUnitRelationship) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetGraphicsCardsOk returns a tuple with the GraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCards

`func (o *ComputeRackUnitRelationship) SetGraphicsCards(v []GraphicsCardRelationship)`

SetGraphicsCards sets GraphicsCards field to given value.

### HasGraphicsCards

`func (o *ComputeRackUnitRelationship) HasGraphicsCards() bool`

HasGraphicsCards returns a boolean if a field has been set.

### SetGraphicsCardsNil

`func (o *ComputeRackUnitRelationship) SetGraphicsCardsNil(b bool)`

 SetGraphicsCardsNil sets the value for GraphicsCards to be an explicit nil

### UnsetGraphicsCards
`func (o *ComputeRackUnitRelationship) UnsetGraphicsCards()`

UnsetGraphicsCards ensures that no value is present for GraphicsCards, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *ComputeRackUnitRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeRackUnitRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeRackUnitRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeRackUnitRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *ComputeRackUnitRelationship) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *ComputeRackUnitRelationship) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *ComputeRackUnitRelationship) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *ComputeRackUnitRelationship) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetMemoryArrays

`func (o *ComputeRackUnitRelationship) GetMemoryArrays() []MemoryArrayRelationship`

GetMemoryArrays returns the MemoryArrays field if non-nil, zero value otherwise.

### GetMemoryArraysOk

`func (o *ComputeRackUnitRelationship) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool)`

GetMemoryArraysOk returns a tuple with the MemoryArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArrays

`func (o *ComputeRackUnitRelationship) SetMemoryArrays(v []MemoryArrayRelationship)`

SetMemoryArrays sets MemoryArrays field to given value.

### HasMemoryArrays

`func (o *ComputeRackUnitRelationship) HasMemoryArrays() bool`

HasMemoryArrays returns a boolean if a field has been set.

### SetMemoryArraysNil

`func (o *ComputeRackUnitRelationship) SetMemoryArraysNil(b bool)`

 SetMemoryArraysNil sets the value for MemoryArrays to be an explicit nil

### UnsetMemoryArrays
`func (o *ComputeRackUnitRelationship) UnsetMemoryArrays()`

UnsetMemoryArrays ensures that no value is present for MemoryArrays, not even an explicit nil
### GetPciDevices

`func (o *ComputeRackUnitRelationship) GetPciDevices() []PciDeviceRelationship`

GetPciDevices returns the PciDevices field if non-nil, zero value otherwise.

### GetPciDevicesOk

`func (o *ComputeRackUnitRelationship) GetPciDevicesOk() (*[]PciDeviceRelationship, bool)`

GetPciDevicesOk returns a tuple with the PciDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciDevices

`func (o *ComputeRackUnitRelationship) SetPciDevices(v []PciDeviceRelationship)`

SetPciDevices sets PciDevices field to given value.

### HasPciDevices

`func (o *ComputeRackUnitRelationship) HasPciDevices() bool`

HasPciDevices returns a boolean if a field has been set.

### SetPciDevicesNil

`func (o *ComputeRackUnitRelationship) SetPciDevicesNil(b bool)`

 SetPciDevicesNil sets the value for PciDevices to be an explicit nil

### UnsetPciDevices
`func (o *ComputeRackUnitRelationship) UnsetPciDevices()`

UnsetPciDevices ensures that no value is present for PciDevices, not even an explicit nil
### GetProcessors

`func (o *ComputeRackUnitRelationship) GetProcessors() []ProcessorUnitRelationship`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ComputeRackUnitRelationship) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ComputeRackUnitRelationship) SetProcessors(v []ProcessorUnitRelationship)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *ComputeRackUnitRelationship) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### SetProcessorsNil

`func (o *ComputeRackUnitRelationship) SetProcessorsNil(b bool)`

 SetProcessorsNil sets the value for Processors to be an explicit nil

### UnsetProcessors
`func (o *ComputeRackUnitRelationship) UnsetProcessors()`

UnsetProcessors ensures that no value is present for Processors, not even an explicit nil
### GetPsus

`func (o *ComputeRackUnitRelationship) GetPsus() []EquipmentPsuRelationship`

GetPsus returns the Psus field if non-nil, zero value otherwise.

### GetPsusOk

`func (o *ComputeRackUnitRelationship) GetPsusOk() (*[]EquipmentPsuRelationship, bool)`

GetPsusOk returns a tuple with the Psus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsus

`func (o *ComputeRackUnitRelationship) SetPsus(v []EquipmentPsuRelationship)`

SetPsus sets Psus field to given value.

### HasPsus

`func (o *ComputeRackUnitRelationship) HasPsus() bool`

HasPsus returns a boolean if a field has been set.

### SetPsusNil

`func (o *ComputeRackUnitRelationship) SetPsusNil(b bool)`

 SetPsusNil sets the value for Psus to be an explicit nil

### UnsetPsus
`func (o *ComputeRackUnitRelationship) UnsetPsus()`

UnsetPsus ensures that no value is present for Psus, not even an explicit nil
### GetRackEnclosureSlot

`func (o *ComputeRackUnitRelationship) GetRackEnclosureSlot() EquipmentRackEnclosureSlotRelationship`

GetRackEnclosureSlot returns the RackEnclosureSlot field if non-nil, zero value otherwise.

### GetRackEnclosureSlotOk

`func (o *ComputeRackUnitRelationship) GetRackEnclosureSlotOk() (*EquipmentRackEnclosureSlotRelationship, bool)`

GetRackEnclosureSlotOk returns a tuple with the RackEnclosureSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackEnclosureSlot

`func (o *ComputeRackUnitRelationship) SetRackEnclosureSlot(v EquipmentRackEnclosureSlotRelationship)`

SetRackEnclosureSlot sets RackEnclosureSlot field to given value.

### HasRackEnclosureSlot

`func (o *ComputeRackUnitRelationship) HasRackEnclosureSlot() bool`

HasRackEnclosureSlot returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ComputeRackUnitRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeRackUnitRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeRackUnitRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeRackUnitRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSasExpanders

`func (o *ComputeRackUnitRelationship) GetSasExpanders() []StorageSasExpanderRelationship`

GetSasExpanders returns the SasExpanders field if non-nil, zero value otherwise.

### GetSasExpandersOk

`func (o *ComputeRackUnitRelationship) GetSasExpandersOk() (*[]StorageSasExpanderRelationship, bool)`

GetSasExpandersOk returns a tuple with the SasExpanders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasExpanders

`func (o *ComputeRackUnitRelationship) SetSasExpanders(v []StorageSasExpanderRelationship)`

SetSasExpanders sets SasExpanders field to given value.

### HasSasExpanders

`func (o *ComputeRackUnitRelationship) HasSasExpanders() bool`

HasSasExpanders returns a boolean if a field has been set.

### SetSasExpandersNil

`func (o *ComputeRackUnitRelationship) SetSasExpandersNil(b bool)`

 SetSasExpandersNil sets the value for SasExpanders to be an explicit nil

### UnsetSasExpanders
`func (o *ComputeRackUnitRelationship) UnsetSasExpanders()`

UnsetSasExpanders ensures that no value is present for SasExpanders, not even an explicit nil
### GetStorageControllers

`func (o *ComputeRackUnitRelationship) GetStorageControllers() []StorageControllerRelationship`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ComputeRackUnitRelationship) GetStorageControllersOk() (*[]StorageControllerRelationship, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ComputeRackUnitRelationship) SetStorageControllers(v []StorageControllerRelationship)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ComputeRackUnitRelationship) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *ComputeRackUnitRelationship) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *ComputeRackUnitRelationship) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetStorageEnclosures

`func (o *ComputeRackUnitRelationship) GetStorageEnclosures() []StorageEnclosureRelationship`

GetStorageEnclosures returns the StorageEnclosures field if non-nil, zero value otherwise.

### GetStorageEnclosuresOk

`func (o *ComputeRackUnitRelationship) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool)`

GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosures

`func (o *ComputeRackUnitRelationship) SetStorageEnclosures(v []StorageEnclosureRelationship)`

SetStorageEnclosures sets StorageEnclosures field to given value.

### HasStorageEnclosures

`func (o *ComputeRackUnitRelationship) HasStorageEnclosures() bool`

HasStorageEnclosures returns a boolean if a field has been set.

### SetStorageEnclosuresNil

`func (o *ComputeRackUnitRelationship) SetStorageEnclosuresNil(b bool)`

 SetStorageEnclosuresNil sets the value for StorageEnclosures to be an explicit nil

### UnsetStorageEnclosures
`func (o *ComputeRackUnitRelationship) UnsetStorageEnclosures()`

UnsetStorageEnclosures ensures that no value is present for StorageEnclosures, not even an explicit nil
### GetTopSystem

`func (o *ComputeRackUnitRelationship) GetTopSystem() TopSystemRelationship`

GetTopSystem returns the TopSystem field if non-nil, zero value otherwise.

### GetTopSystemOk

`func (o *ComputeRackUnitRelationship) GetTopSystemOk() (*TopSystemRelationship, bool)`

GetTopSystemOk returns a tuple with the TopSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSystem

`func (o *ComputeRackUnitRelationship) SetTopSystem(v TopSystemRelationship)`

SetTopSystem sets TopSystem field to given value.

### HasTopSystem

`func (o *ComputeRackUnitRelationship) HasTopSystem() bool`

HasTopSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


