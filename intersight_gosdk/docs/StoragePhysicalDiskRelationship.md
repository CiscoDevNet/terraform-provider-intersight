# StoragePhysicalDiskRelationship

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
**BlockSize** | Pointer to **string** | The block size of the physical disk in bytes. | [optional] [readonly] 
**Bootable** | Pointer to **string** | This field identifies the disk drive as bootable if set to true. | [optional] [readonly] 
**ConfigurationCheckpoint** | Pointer to **string** | The current configuration checkpoint of the physical disk. | [optional] [readonly] 
**ConfigurationState** | Pointer to **string** | The current configuration state of the physical disk. | [optional] [readonly] 
**DiscoveredPath** | Pointer to **string** | The discovered path of the physical disk. | [optional] [readonly] 
**DiskId** | Pointer to **string** | This field identifies the ID assigned to physical disks. | [optional] [readonly] 
**DiskState** | Pointer to **string** | This field identifies the health of the disk. | [optional] [readonly] 
**DriveFirmware** | Pointer to **string** | This field identifies the disk firmware running in the disk. | [optional] 
**DriveState** | Pointer to **string** | The drive state as reported by the controller. | [optional] [readonly] 
**FdeCapable** | Pointer to **string** | Full-Disk Encryption capability parameter of the physical disk. | [optional] 
**HotSpareType** | Pointer to **string** | Type of hotspare configured on the physical disk. | [optional] 
**LinkSpeed** | Pointer to **string** | The speed of the link between the drive and the controller. | [optional] [readonly] 
**LinkState** | Pointer to **string** | The current link state of the physical disk. | [optional] [readonly] 
**NumBlocks** | Pointer to **string** | The number of blocks present on the physical disk. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | Operational power of the physical disk. | [optional] [readonly] 
**OperQualifierReason** | Pointer to **string** | This reason for the operational status of the disk. | [optional] [readonly] 
**Operability** | Pointer to **string** | This field identifies the disk operability of the disk. | [optional] [readonly] 
**PhysicalBlockSize** | Pointer to **string** | The block size of the installed physical disk. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for physicalDisk. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence state of the physical disk. | [optional] [readonly] 
**Protocol** | Pointer to **string** | This field identifies the disk protocol used for communication. | [optional] [readonly] 
**RawSize** | Pointer to **string** | The raw size of the physical disk in MB. | [optional] [readonly] 
**Secured** | Pointer to **string** | This field identifies whether the disk is encrypted. | [optional] 
**Size** | Pointer to **string** | The size of the physical disk in MB. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of the physical disk. | [optional] [readonly] 
**Type** | Pointer to **string** | This field identifies the type of the physical disk. | [optional] [readonly] 
**VariantType** | Pointer to **string** | The variant type of the physical disk. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**LocatorLed** | Pointer to [**EquipmentLocatorLedRelationship**](equipment.LocatorLed.Relationship.md) |  | [optional] 
**PhysicalDiskExtensions** | Pointer to [**[]StoragePhysicalDiskExtensionRelationship**](storage.PhysicalDiskExtension.Relationship.md) | An array of relationships to storagePhysicalDiskExtension resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 
**SasPorts** | Pointer to [**[]StorageSasPortRelationship**](storage.SasPort.Relationship.md) | An array of relationships to storageSasPort resources. | [optional] [readonly] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**StorageEnclosure** | Pointer to [**StorageEnclosureRelationship**](storage.Enclosure.Relationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDiskRelationship

`func NewStoragePhysicalDiskRelationship(classId string, objectType string, ) *StoragePhysicalDiskRelationship`

NewStoragePhysicalDiskRelationship instantiates a new StoragePhysicalDiskRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskRelationshipWithDefaults

`func NewStoragePhysicalDiskRelationshipWithDefaults() *StoragePhysicalDiskRelationship`

NewStoragePhysicalDiskRelationshipWithDefaults instantiates a new StoragePhysicalDiskRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StoragePhysicalDiskRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StoragePhysicalDiskRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StoragePhysicalDiskRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StoragePhysicalDiskRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StoragePhysicalDiskRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePhysicalDiskRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePhysicalDiskRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StoragePhysicalDiskRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StoragePhysicalDiskRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StoragePhysicalDiskRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StoragePhysicalDiskRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StoragePhysicalDiskRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StoragePhysicalDiskRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StoragePhysicalDiskRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StoragePhysicalDiskRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StoragePhysicalDiskRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StoragePhysicalDiskRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StoragePhysicalDiskRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StoragePhysicalDiskRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StoragePhysicalDiskRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StoragePhysicalDiskRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StoragePhysicalDiskRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StoragePhysicalDiskRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StoragePhysicalDiskRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePhysicalDiskRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePhysicalDiskRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StoragePhysicalDiskRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StoragePhysicalDiskRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StoragePhysicalDiskRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StoragePhysicalDiskRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StoragePhysicalDiskRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StoragePhysicalDiskRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StoragePhysicalDiskRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StoragePhysicalDiskRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StoragePhysicalDiskRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoragePhysicalDiskRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoragePhysicalDiskRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoragePhysicalDiskRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StoragePhysicalDiskRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StoragePhysicalDiskRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StoragePhysicalDiskRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StoragePhysicalDiskRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StoragePhysicalDiskRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StoragePhysicalDiskRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StoragePhysicalDiskRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StoragePhysicalDiskRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StoragePhysicalDiskRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StoragePhysicalDiskRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StoragePhysicalDiskRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StoragePhysicalDiskRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StoragePhysicalDiskRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StoragePhysicalDiskRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StoragePhysicalDiskRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StoragePhysicalDiskRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StoragePhysicalDiskRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StoragePhysicalDiskRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StoragePhysicalDiskRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StoragePhysicalDiskRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StoragePhysicalDiskRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StoragePhysicalDiskRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StoragePhysicalDiskRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StoragePhysicalDiskRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StoragePhysicalDiskRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StoragePhysicalDiskRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StoragePhysicalDiskRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StoragePhysicalDiskRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StoragePhysicalDiskRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StoragePhysicalDiskRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StoragePhysicalDiskRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StoragePhysicalDiskRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StoragePhysicalDiskRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StoragePhysicalDiskRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StoragePhysicalDiskRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StoragePhysicalDiskRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StoragePhysicalDiskRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StoragePhysicalDiskRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StoragePhysicalDiskRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StoragePhysicalDiskRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StoragePhysicalDiskRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StoragePhysicalDiskRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StoragePhysicalDiskRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StoragePhysicalDiskRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StoragePhysicalDiskRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StoragePhysicalDiskRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StoragePhysicalDiskRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StoragePhysicalDiskRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StoragePhysicalDiskRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StoragePhysicalDiskRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StoragePhysicalDiskRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StoragePhysicalDiskRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StoragePhysicalDiskRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StoragePhysicalDiskRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetBlockSize

`func (o *StoragePhysicalDiskRelationship) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StoragePhysicalDiskRelationship) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StoragePhysicalDiskRelationship) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StoragePhysicalDiskRelationship) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StoragePhysicalDiskRelationship) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StoragePhysicalDiskRelationship) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StoragePhysicalDiskRelationship) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StoragePhysicalDiskRelationship) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigurationCheckpoint

`func (o *StoragePhysicalDiskRelationship) GetConfigurationCheckpoint() string`

GetConfigurationCheckpoint returns the ConfigurationCheckpoint field if non-nil, zero value otherwise.

### GetConfigurationCheckpointOk

`func (o *StoragePhysicalDiskRelationship) GetConfigurationCheckpointOk() (*string, bool)`

GetConfigurationCheckpointOk returns a tuple with the ConfigurationCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationCheckpoint

`func (o *StoragePhysicalDiskRelationship) SetConfigurationCheckpoint(v string)`

SetConfigurationCheckpoint sets ConfigurationCheckpoint field to given value.

### HasConfigurationCheckpoint

`func (o *StoragePhysicalDiskRelationship) HasConfigurationCheckpoint() bool`

HasConfigurationCheckpoint returns a boolean if a field has been set.

### GetConfigurationState

`func (o *StoragePhysicalDiskRelationship) GetConfigurationState() string`

GetConfigurationState returns the ConfigurationState field if non-nil, zero value otherwise.

### GetConfigurationStateOk

`func (o *StoragePhysicalDiskRelationship) GetConfigurationStateOk() (*string, bool)`

GetConfigurationStateOk returns a tuple with the ConfigurationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationState

`func (o *StoragePhysicalDiskRelationship) SetConfigurationState(v string)`

SetConfigurationState sets ConfigurationState field to given value.

### HasConfigurationState

`func (o *StoragePhysicalDiskRelationship) HasConfigurationState() bool`

HasConfigurationState returns a boolean if a field has been set.

### GetDiscoveredPath

`func (o *StoragePhysicalDiskRelationship) GetDiscoveredPath() string`

GetDiscoveredPath returns the DiscoveredPath field if non-nil, zero value otherwise.

### GetDiscoveredPathOk

`func (o *StoragePhysicalDiskRelationship) GetDiscoveredPathOk() (*string, bool)`

GetDiscoveredPathOk returns a tuple with the DiscoveredPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredPath

`func (o *StoragePhysicalDiskRelationship) SetDiscoveredPath(v string)`

SetDiscoveredPath sets DiscoveredPath field to given value.

### HasDiscoveredPath

`func (o *StoragePhysicalDiskRelationship) HasDiscoveredPath() bool`

HasDiscoveredPath returns a boolean if a field has been set.

### GetDiskId

`func (o *StoragePhysicalDiskRelationship) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StoragePhysicalDiskRelationship) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StoragePhysicalDiskRelationship) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StoragePhysicalDiskRelationship) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetDiskState

`func (o *StoragePhysicalDiskRelationship) GetDiskState() string`

GetDiskState returns the DiskState field if non-nil, zero value otherwise.

### GetDiskStateOk

`func (o *StoragePhysicalDiskRelationship) GetDiskStateOk() (*string, bool)`

GetDiskStateOk returns a tuple with the DiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskState

`func (o *StoragePhysicalDiskRelationship) SetDiskState(v string)`

SetDiskState sets DiskState field to given value.

### HasDiskState

`func (o *StoragePhysicalDiskRelationship) HasDiskState() bool`

HasDiskState returns a boolean if a field has been set.

### GetDriveFirmware

`func (o *StoragePhysicalDiskRelationship) GetDriveFirmware() string`

GetDriveFirmware returns the DriveFirmware field if non-nil, zero value otherwise.

### GetDriveFirmwareOk

`func (o *StoragePhysicalDiskRelationship) GetDriveFirmwareOk() (*string, bool)`

GetDriveFirmwareOk returns a tuple with the DriveFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveFirmware

`func (o *StoragePhysicalDiskRelationship) SetDriveFirmware(v string)`

SetDriveFirmware sets DriveFirmware field to given value.

### HasDriveFirmware

`func (o *StoragePhysicalDiskRelationship) HasDriveFirmware() bool`

HasDriveFirmware returns a boolean if a field has been set.

### GetDriveState

`func (o *StoragePhysicalDiskRelationship) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StoragePhysicalDiskRelationship) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StoragePhysicalDiskRelationship) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StoragePhysicalDiskRelationship) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetFdeCapable

`func (o *StoragePhysicalDiskRelationship) GetFdeCapable() string`

GetFdeCapable returns the FdeCapable field if non-nil, zero value otherwise.

### GetFdeCapableOk

`func (o *StoragePhysicalDiskRelationship) GetFdeCapableOk() (*string, bool)`

GetFdeCapableOk returns a tuple with the FdeCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeCapable

`func (o *StoragePhysicalDiskRelationship) SetFdeCapable(v string)`

SetFdeCapable sets FdeCapable field to given value.

### HasFdeCapable

`func (o *StoragePhysicalDiskRelationship) HasFdeCapable() bool`

HasFdeCapable returns a boolean if a field has been set.

### GetHotSpareType

`func (o *StoragePhysicalDiskRelationship) GetHotSpareType() string`

GetHotSpareType returns the HotSpareType field if non-nil, zero value otherwise.

### GetHotSpareTypeOk

`func (o *StoragePhysicalDiskRelationship) GetHotSpareTypeOk() (*string, bool)`

GetHotSpareTypeOk returns a tuple with the HotSpareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotSpareType

`func (o *StoragePhysicalDiskRelationship) SetHotSpareType(v string)`

SetHotSpareType sets HotSpareType field to given value.

### HasHotSpareType

`func (o *StoragePhysicalDiskRelationship) HasHotSpareType() bool`

HasHotSpareType returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *StoragePhysicalDiskRelationship) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *StoragePhysicalDiskRelationship) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *StoragePhysicalDiskRelationship) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *StoragePhysicalDiskRelationship) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkState

`func (o *StoragePhysicalDiskRelationship) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *StoragePhysicalDiskRelationship) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *StoragePhysicalDiskRelationship) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *StoragePhysicalDiskRelationship) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StoragePhysicalDiskRelationship) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StoragePhysicalDiskRelationship) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StoragePhysicalDiskRelationship) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StoragePhysicalDiskRelationship) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperPowerState

`func (o *StoragePhysicalDiskRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *StoragePhysicalDiskRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *StoragePhysicalDiskRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *StoragePhysicalDiskRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperQualifierReason

`func (o *StoragePhysicalDiskRelationship) GetOperQualifierReason() string`

GetOperQualifierReason returns the OperQualifierReason field if non-nil, zero value otherwise.

### GetOperQualifierReasonOk

`func (o *StoragePhysicalDiskRelationship) GetOperQualifierReasonOk() (*string, bool)`

GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifierReason

`func (o *StoragePhysicalDiskRelationship) SetOperQualifierReason(v string)`

SetOperQualifierReason sets OperQualifierReason field to given value.

### HasOperQualifierReason

`func (o *StoragePhysicalDiskRelationship) HasOperQualifierReason() bool`

HasOperQualifierReason returns a boolean if a field has been set.

### GetOperability

`func (o *StoragePhysicalDiskRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StoragePhysicalDiskRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StoragePhysicalDiskRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StoragePhysicalDiskRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StoragePhysicalDiskRelationship) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StoragePhysicalDiskRelationship) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StoragePhysicalDiskRelationship) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StoragePhysicalDiskRelationship) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPid

`func (o *StoragePhysicalDiskRelationship) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StoragePhysicalDiskRelationship) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StoragePhysicalDiskRelationship) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StoragePhysicalDiskRelationship) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPresence

`func (o *StoragePhysicalDiskRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StoragePhysicalDiskRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StoragePhysicalDiskRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StoragePhysicalDiskRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProtocol

`func (o *StoragePhysicalDiskRelationship) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StoragePhysicalDiskRelationship) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StoragePhysicalDiskRelationship) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StoragePhysicalDiskRelationship) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRawSize

`func (o *StoragePhysicalDiskRelationship) GetRawSize() string`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *StoragePhysicalDiskRelationship) GetRawSizeOk() (*string, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *StoragePhysicalDiskRelationship) SetRawSize(v string)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *StoragePhysicalDiskRelationship) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetSecured

`func (o *StoragePhysicalDiskRelationship) GetSecured() string`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *StoragePhysicalDiskRelationship) GetSecuredOk() (*string, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *StoragePhysicalDiskRelationship) SetSecured(v string)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *StoragePhysicalDiskRelationship) HasSecured() bool`

HasSecured returns a boolean if a field has been set.

### GetSize

`func (o *StoragePhysicalDiskRelationship) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePhysicalDiskRelationship) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePhysicalDiskRelationship) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePhysicalDiskRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetThermal

`func (o *StoragePhysicalDiskRelationship) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *StoragePhysicalDiskRelationship) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *StoragePhysicalDiskRelationship) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *StoragePhysicalDiskRelationship) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetType

`func (o *StoragePhysicalDiskRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoragePhysicalDiskRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoragePhysicalDiskRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoragePhysicalDiskRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariantType

`func (o *StoragePhysicalDiskRelationship) GetVariantType() string`

GetVariantType returns the VariantType field if non-nil, zero value otherwise.

### GetVariantTypeOk

`func (o *StoragePhysicalDiskRelationship) GetVariantTypeOk() (*string, bool)`

GetVariantTypeOk returns a tuple with the VariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantType

`func (o *StoragePhysicalDiskRelationship) SetVariantType(v string)`

SetVariantType sets VariantType field to given value.

### HasVariantType

`func (o *StoragePhysicalDiskRelationship) HasVariantType() bool`

HasVariantType returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetLocatorLed

`func (o *StoragePhysicalDiskRelationship) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *StoragePhysicalDiskRelationship) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *StoragePhysicalDiskRelationship) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *StoragePhysicalDiskRelationship) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetPhysicalDiskExtensions

`func (o *StoragePhysicalDiskRelationship) GetPhysicalDiskExtensions() []StoragePhysicalDiskExtensionRelationship`

GetPhysicalDiskExtensions returns the PhysicalDiskExtensions field if non-nil, zero value otherwise.

### GetPhysicalDiskExtensionsOk

`func (o *StoragePhysicalDiskRelationship) GetPhysicalDiskExtensionsOk() (*[]StoragePhysicalDiskExtensionRelationship, bool)`

GetPhysicalDiskExtensionsOk returns a tuple with the PhysicalDiskExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskExtensions

`func (o *StoragePhysicalDiskRelationship) SetPhysicalDiskExtensions(v []StoragePhysicalDiskExtensionRelationship)`

SetPhysicalDiskExtensions sets PhysicalDiskExtensions field to given value.

### HasPhysicalDiskExtensions

`func (o *StoragePhysicalDiskRelationship) HasPhysicalDiskExtensions() bool`

HasPhysicalDiskExtensions returns a boolean if a field has been set.

### SetPhysicalDiskExtensionsNil

`func (o *StoragePhysicalDiskRelationship) SetPhysicalDiskExtensionsNil(b bool)`

 SetPhysicalDiskExtensionsNil sets the value for PhysicalDiskExtensions to be an explicit nil

### UnsetPhysicalDiskExtensions
`func (o *StoragePhysicalDiskRelationship) UnsetPhysicalDiskExtensions()`

UnsetPhysicalDiskExtensions ensures that no value is present for PhysicalDiskExtensions, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePhysicalDiskRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StoragePhysicalDiskRelationship) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StoragePhysicalDiskRelationship) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StoragePhysicalDiskRelationship) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StoragePhysicalDiskRelationship) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StoragePhysicalDiskRelationship) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StoragePhysicalDiskRelationship) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil
### GetSasPorts

`func (o *StoragePhysicalDiskRelationship) GetSasPorts() []StorageSasPortRelationship`

GetSasPorts returns the SasPorts field if non-nil, zero value otherwise.

### GetSasPortsOk

`func (o *StoragePhysicalDiskRelationship) GetSasPortsOk() (*[]StorageSasPortRelationship, bool)`

GetSasPortsOk returns a tuple with the SasPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPorts

`func (o *StoragePhysicalDiskRelationship) SetSasPorts(v []StorageSasPortRelationship)`

SetSasPorts sets SasPorts field to given value.

### HasSasPorts

`func (o *StoragePhysicalDiskRelationship) HasSasPorts() bool`

HasSasPorts returns a boolean if a field has been set.

### SetSasPortsNil

`func (o *StoragePhysicalDiskRelationship) SetSasPortsNil(b bool)`

 SetSasPortsNil sets the value for SasPorts to be an explicit nil

### UnsetSasPorts
`func (o *StoragePhysicalDiskRelationship) UnsetSasPorts()`

UnsetSasPorts ensures that no value is present for SasPorts, not even an explicit nil
### GetStorageController

`func (o *StoragePhysicalDiskRelationship) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StoragePhysicalDiskRelationship) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StoragePhysicalDiskRelationship) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StoragePhysicalDiskRelationship) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetStorageEnclosure

`func (o *StoragePhysicalDiskRelationship) GetStorageEnclosure() StorageEnclosureRelationship`

GetStorageEnclosure returns the StorageEnclosure field if non-nil, zero value otherwise.

### GetStorageEnclosureOk

`func (o *StoragePhysicalDiskRelationship) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool)`

GetStorageEnclosureOk returns a tuple with the StorageEnclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEnclosure

`func (o *StoragePhysicalDiskRelationship) SetStorageEnclosure(v StorageEnclosureRelationship)`

SetStorageEnclosure sets StorageEnclosure field to given value.

### HasStorageEnclosure

`func (o *StoragePhysicalDiskRelationship) HasStorageEnclosure() bool`

HasStorageEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


