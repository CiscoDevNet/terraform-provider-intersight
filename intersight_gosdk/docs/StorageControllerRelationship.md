# StorageControllerRelationship

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
**ControllerFlags** | Pointer to **string** | The flags for the storage controller. | [optional] [readonly] 
**ControllerId** | Pointer to **string** | The Id of the storage controller. | [optional] [readonly] 
**ControllerStatus** | Pointer to **string** | The current status of controller. | [optional] [readonly] 
**ForeignConfigPresent** | Pointer to **bool** | Storage controller has detected disks in foreign config. | [optional] 
**HwRevision** | Pointer to **string** | The hardware revision of controller. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Interface types are Sas, Sata, PCH. | [optional] 
**MaxVolumesSupported** | Pointer to **int64** | Maximum virtual drives that can be created on this Storage Controller. | [optional] 
**OobInterfaceSupported** | Pointer to **string** | The CIMC support for out-of-band configuration of controller. | [optional] [readonly] 
**OperState** | Pointer to **string** | The current operational state of controller. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the storage controller. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The current pci address of controller. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The pci slot name for the controller. | [optional] [readonly] 
**Presence** | Pointer to **string** | Physical Presence State for the Storage Controller. | [optional] [readonly] 
**RaidSupport** | Pointer to **string** | The RAID levels supported by controller. | [optional] [readonly] 
**RebuildRate** | Pointer to **string** | Logical volume or RAID rebuild rate of Storage Controller. | [optional] [readonly] 
**SelfEncryptEnabled** | Pointer to **string** | Storage controller disk self encryption state. | [optional] 
**Type** | Pointer to **string** | Controller types are Raid, FlexFlash. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**DiskGroup** | Pointer to [**[]StorageDiskGroupRelationship**](storage.DiskGroup.Relationship.md) | An array of relationships to storageDiskGroup resources. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDiskExtensions** | Pointer to [**[]StoragePhysicalDiskExtensionRelationship**](storage.PhysicalDiskExtension.Relationship.md) | An array of relationships to storagePhysicalDiskExtension resources. | [optional] [readonly] 
**PhysicalDisks** | Pointer to [**[]StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) | An array of relationships to storagePhysicalDisk resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**VirtualDriveExtensions** | Pointer to [**[]StorageVirtualDriveExtensionRelationship**](storage.VirtualDriveExtension.Relationship.md) | An array of relationships to storageVirtualDriveExtension resources. | [optional] [readonly] 
**VirtualDrives** | Pointer to [**[]StorageVirtualDriveRelationship**](storage.VirtualDrive.Relationship.md) | An array of relationships to storageVirtualDrive resources. | [optional] [readonly] 

## Methods

### NewStorageControllerRelationship

`func NewStorageControllerRelationship(classId string, objectType string, ) *StorageControllerRelationship`

NewStorageControllerRelationship instantiates a new StorageControllerRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageControllerRelationshipWithDefaults

`func NewStorageControllerRelationshipWithDefaults() *StorageControllerRelationship`

NewStorageControllerRelationshipWithDefaults instantiates a new StorageControllerRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageControllerRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageControllerRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageControllerRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageControllerRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageControllerRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageControllerRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageControllerRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageControllerRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageControllerRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageControllerRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageControllerRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageControllerRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageControllerRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageControllerRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageControllerRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageControllerRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageControllerRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageControllerRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageControllerRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageControllerRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageControllerRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageControllerRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageControllerRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageControllerRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageControllerRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageControllerRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageControllerRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageControllerRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageControllerRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageControllerRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageControllerRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageControllerRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageControllerRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageControllerRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageControllerRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageControllerRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageControllerRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageControllerRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageControllerRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageControllerRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageControllerRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageControllerRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageControllerRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageControllerRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageControllerRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageControllerRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageControllerRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageControllerRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageControllerRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageControllerRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageControllerRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageControllerRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageControllerRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageControllerRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageControllerRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageControllerRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageControllerRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageControllerRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageControllerRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageControllerRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageControllerRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageControllerRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageControllerRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageControllerRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageControllerRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageControllerRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageControllerRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageControllerRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageControllerRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageControllerRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageControllerRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageControllerRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageControllerRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageControllerRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageControllerRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageControllerRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageControllerRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageControllerRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageControllerRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageControllerRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageControllerRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageControllerRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageControllerRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageControllerRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageControllerRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageControllerRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageControllerRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageControllerRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageControllerRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageControllerRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageControllerRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageControllerRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetControllerFlags

`func (o *StorageControllerRelationship) GetControllerFlags() string`

GetControllerFlags returns the ControllerFlags field if non-nil, zero value otherwise.

### GetControllerFlagsOk

`func (o *StorageControllerRelationship) GetControllerFlagsOk() (*string, bool)`

GetControllerFlagsOk returns a tuple with the ControllerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFlags

`func (o *StorageControllerRelationship) SetControllerFlags(v string)`

SetControllerFlags sets ControllerFlags field to given value.

### HasControllerFlags

`func (o *StorageControllerRelationship) HasControllerFlags() bool`

HasControllerFlags returns a boolean if a field has been set.

### GetControllerId

`func (o *StorageControllerRelationship) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *StorageControllerRelationship) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *StorageControllerRelationship) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *StorageControllerRelationship) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageControllerRelationship) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageControllerRelationship) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageControllerRelationship) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageControllerRelationship) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetForeignConfigPresent

`func (o *StorageControllerRelationship) GetForeignConfigPresent() bool`

GetForeignConfigPresent returns the ForeignConfigPresent field if non-nil, zero value otherwise.

### GetForeignConfigPresentOk

`func (o *StorageControllerRelationship) GetForeignConfigPresentOk() (*bool, bool)`

GetForeignConfigPresentOk returns a tuple with the ForeignConfigPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignConfigPresent

`func (o *StorageControllerRelationship) SetForeignConfigPresent(v bool)`

SetForeignConfigPresent sets ForeignConfigPresent field to given value.

### HasForeignConfigPresent

`func (o *StorageControllerRelationship) HasForeignConfigPresent() bool`

HasForeignConfigPresent returns a boolean if a field has been set.

### GetHwRevision

`func (o *StorageControllerRelationship) GetHwRevision() string`

GetHwRevision returns the HwRevision field if non-nil, zero value otherwise.

### GetHwRevisionOk

`func (o *StorageControllerRelationship) GetHwRevisionOk() (*string, bool)`

GetHwRevisionOk returns a tuple with the HwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRevision

`func (o *StorageControllerRelationship) SetHwRevision(v string)`

SetHwRevision sets HwRevision field to given value.

### HasHwRevision

`func (o *StorageControllerRelationship) HasHwRevision() bool`

HasHwRevision returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StorageControllerRelationship) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StorageControllerRelationship) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StorageControllerRelationship) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StorageControllerRelationship) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMaxVolumesSupported

`func (o *StorageControllerRelationship) GetMaxVolumesSupported() int64`

GetMaxVolumesSupported returns the MaxVolumesSupported field if non-nil, zero value otherwise.

### GetMaxVolumesSupportedOk

`func (o *StorageControllerRelationship) GetMaxVolumesSupportedOk() (*int64, bool)`

GetMaxVolumesSupportedOk returns a tuple with the MaxVolumesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumesSupported

`func (o *StorageControllerRelationship) SetMaxVolumesSupported(v int64)`

SetMaxVolumesSupported sets MaxVolumesSupported field to given value.

### HasMaxVolumesSupported

`func (o *StorageControllerRelationship) HasMaxVolumesSupported() bool`

HasMaxVolumesSupported returns a boolean if a field has been set.

### GetOobInterfaceSupported

`func (o *StorageControllerRelationship) GetOobInterfaceSupported() string`

GetOobInterfaceSupported returns the OobInterfaceSupported field if non-nil, zero value otherwise.

### GetOobInterfaceSupportedOk

`func (o *StorageControllerRelationship) GetOobInterfaceSupportedOk() (*string, bool)`

GetOobInterfaceSupportedOk returns a tuple with the OobInterfaceSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobInterfaceSupported

`func (o *StorageControllerRelationship) SetOobInterfaceSupported(v string)`

SetOobInterfaceSupported sets OobInterfaceSupported field to given value.

### HasOobInterfaceSupported

`func (o *StorageControllerRelationship) HasOobInterfaceSupported() bool`

HasOobInterfaceSupported returns a boolean if a field has been set.

### GetOperState

`func (o *StorageControllerRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageControllerRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageControllerRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageControllerRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageControllerRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageControllerRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageControllerRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageControllerRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPciAddr

`func (o *StorageControllerRelationship) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *StorageControllerRelationship) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *StorageControllerRelationship) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *StorageControllerRelationship) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPciSlot

`func (o *StorageControllerRelationship) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *StorageControllerRelationship) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *StorageControllerRelationship) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *StorageControllerRelationship) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPresence

`func (o *StorageControllerRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageControllerRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageControllerRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageControllerRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRaidSupport

`func (o *StorageControllerRelationship) GetRaidSupport() string`

GetRaidSupport returns the RaidSupport field if non-nil, zero value otherwise.

### GetRaidSupportOk

`func (o *StorageControllerRelationship) GetRaidSupportOk() (*string, bool)`

GetRaidSupportOk returns a tuple with the RaidSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidSupport

`func (o *StorageControllerRelationship) SetRaidSupport(v string)`

SetRaidSupport sets RaidSupport field to given value.

### HasRaidSupport

`func (o *StorageControllerRelationship) HasRaidSupport() bool`

HasRaidSupport returns a boolean if a field has been set.

### GetRebuildRate

`func (o *StorageControllerRelationship) GetRebuildRate() string`

GetRebuildRate returns the RebuildRate field if non-nil, zero value otherwise.

### GetRebuildRateOk

`func (o *StorageControllerRelationship) GetRebuildRateOk() (*string, bool)`

GetRebuildRateOk returns a tuple with the RebuildRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildRate

`func (o *StorageControllerRelationship) SetRebuildRate(v string)`

SetRebuildRate sets RebuildRate field to given value.

### HasRebuildRate

`func (o *StorageControllerRelationship) HasRebuildRate() bool`

HasRebuildRate returns a boolean if a field has been set.

### GetSelfEncryptEnabled

`func (o *StorageControllerRelationship) GetSelfEncryptEnabled() string`

GetSelfEncryptEnabled returns the SelfEncryptEnabled field if non-nil, zero value otherwise.

### GetSelfEncryptEnabledOk

`func (o *StorageControllerRelationship) GetSelfEncryptEnabledOk() (*string, bool)`

GetSelfEncryptEnabledOk returns a tuple with the SelfEncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfEncryptEnabled

`func (o *StorageControllerRelationship) SetSelfEncryptEnabled(v string)`

SetSelfEncryptEnabled sets SelfEncryptEnabled field to given value.

### HasSelfEncryptEnabled

`func (o *StorageControllerRelationship) HasSelfEncryptEnabled() bool`

HasSelfEncryptEnabled returns a boolean if a field has been set.

### GetType

`func (o *StorageControllerRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageControllerRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageControllerRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageControllerRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComputeBlade

`func (o *StorageControllerRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *StorageControllerRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *StorageControllerRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *StorageControllerRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageControllerRelationship) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageControllerRelationship) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageControllerRelationship) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageControllerRelationship) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageControllerRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageControllerRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageControllerRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageControllerRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageControllerRelationship) GetDiskGroup() []StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageControllerRelationship) GetDiskGroupOk() (*[]StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageControllerRelationship) SetDiskGroup(v []StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageControllerRelationship) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### SetDiskGroupNil

`func (o *StorageControllerRelationship) SetDiskGroupNil(b bool)`

 SetDiskGroupNil sets the value for DiskGroup to be an explicit nil

### UnsetDiskGroup
`func (o *StorageControllerRelationship) UnsetDiskGroup()`

UnsetDiskGroup ensures that no value is present for DiskGroup, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageControllerRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageControllerRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageControllerRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageControllerRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskExtensions

`func (o *StorageControllerRelationship) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StorageControllerRelationship) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StorageControllerRelationship) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StorageControllerRelationship) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StorageControllerRelationship) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StorageControllerRelationship) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetPhysicalDisks

`func (o *StorageControllerRelationship) GetPhysicalDisks() []StoragePhysicalDiskRelationship`

GetPhysicalDisks returns the PhysicalDisks field if non-nil, zero value otherwise.

### GetPhysicalDisksOk

`func (o *StorageControllerRelationship) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool)`

GetPhysicalDisksOk returns a tuple with the PhysicalDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDisks

`func (o *StorageControllerRelationship) SetPhysicalDisks(v []StoragePhysicalDiskRelationship)`

SetPhysicalDisks sets PhysicalDisks field to given value.

### HasPhysicalDisks

`func (o *StorageControllerRelationship) HasPhysicalDisks() bool`

HasPhysicalDisks returns a boolean if a field has been set.

### SetPhysicalDisksNil

`func (o *StorageControllerRelationship) SetPhysicalDisksNil(b bool)`

 SetPhysicalDisksNil sets the value for PhysicalDisks to be an explicit nil

### UnsetPhysicalDisks
`func (o *StorageControllerRelationship) UnsetPhysicalDisks()`

UnsetPhysicalDisks ensures that no value is present for PhysicalDisks, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageControllerRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageControllerRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageControllerRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageControllerRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageControllerRelationship) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageControllerRelationship) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageControllerRelationship) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageControllerRelationship) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageControllerRelationship) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageControllerRelationship) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetVirtualDriveExtensions

`func (o *StorageControllerRelationship) GetVirtualDriveExtensions() []StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtensions returns the VirtualDriveExtensions field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionsOk

`func (o *StorageControllerRelationship) GetVirtualDriveExtensionsOk() (*[]StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionsOk returns a tuple with the VirtualDriveExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtensions

`func (o *StorageControllerRelationship) SetVirtualDriveExtensions(v []StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtensions sets VirtualDriveExtensions field to given value.

### HasVirtualDriveExtensions

`func (o *StorageControllerRelationship) HasVirtualDriveExtensions() bool`

HasVirtualDriveExtensions returns a boolean if a field has been set.

### SetVirtualDriveExtensionsNil

`func (o *StorageControllerRelationship) SetVirtualDriveExtensionsNil(b bool)`

 SetVirtualDriveExtensionsNil sets the value for VirtualDriveExtensions to be an explicit nil

### UnsetVirtualDriveExtensions
`func (o *StorageControllerRelationship) UnsetVirtualDriveExtensions()`

UnsetVirtualDriveExtensions ensures that no value is present for VirtualDriveExtensions, not even an explicit nil
### GetVirtualDrives

`func (o *StorageControllerRelationship) GetVirtualDrives() []StorageVirtualDriveRelationship`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *StorageControllerRelationship) GetVirtualDrivesOk() (*[]StorageVirtualDriveRelationship, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *StorageControllerRelationship) SetVirtualDrives(v []StorageVirtualDriveRelationship)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *StorageControllerRelationship) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.

### SetVirtualDrivesNil

`func (o *StorageControllerRelationship) SetVirtualDrivesNil(b bool)`

 SetVirtualDrivesNil sets the value for VirtualDrives to be an explicit nil

### UnsetVirtualDrives
`func (o *StorageControllerRelationship) UnsetVirtualDrives()`

UnsetVirtualDrives ensures that no value is present for VirtualDrives, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


