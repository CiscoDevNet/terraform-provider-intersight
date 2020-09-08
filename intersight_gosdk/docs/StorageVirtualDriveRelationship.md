# StorageVirtualDriveRelationship

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
**AccessPolicy** | Pointer to **string** | The access policy of the virtual drive. | [optional] [readonly] 
**ActualWriteCachePolicy** | Pointer to **string** | The current write cache policy of the virtual drive. | [optional] [readonly] 
**AvailableSize** | Pointer to **string** | Available storage capacity of the virtual drive. | [optional] [readonly] 
**BlockSize** | Pointer to **string** | Block size of the virtual drive. | [optional] [readonly] 
**Bootable** | Pointer to **string** | The virtual drive bootable state. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | The configuration state of the virtual drive. | [optional] [readonly] 
**ConfiguredWriteCachePolicy** | Pointer to **string** | The requested write cache policy of the virtual drive. | [optional] [readonly] 
**ConnectionProtocol** | Pointer to **string** | The connection protocol of the virtual drive. | [optional] [readonly] 
**DriveCache** | Pointer to **string** | The state of the drive cache of the virtual drive. | [optional] [readonly] 
**DriveSecurity** | Pointer to **string** | The driveSecurity state of the Virtual drive. | [optional] [readonly] 
**DriveState** | Pointer to **string** | The state of the Virtual drive. | [optional] [readonly] 
**IoPolicy** | Pointer to **string** | The Input/Output Policy defined on the Virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Virtual drive. | [optional] [readonly] 
**NumBlocks** | Pointer to **string** | Number of Blocks on the Physical Disk. | [optional] [readonly] 
**OperState** | Pointer to **string** | The current operational state of Virtual drive. | [optional] [readonly] 
**Operability** | Pointer to **string** | The current operability state of Virtual drive. | [optional] [readonly] 
**PhysicalBlockSize** | Pointer to **string** | The block size of the the virtual drive. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence status of the virtual drive. | [optional] [readonly] 
**ReadPolicy** | Pointer to **string** | The read-ahead cache mode of the virtual drive. | [optional] [readonly] 
**SecurityFlags** | Pointer to **string** | The security flags set for this virtual drive. | [optional] [readonly] 
**Size** | Pointer to **string** | The size of the virtual drive in MB. | [optional] [readonly] 
**StripSize** | Pointer to **string** | The strip size is the portion of a stripe that resides on a single drive in the drive group, this is measured in KB. | [optional] [readonly] 
**Type** | Pointer to **string** | The raid type of the virtual drive. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The uuid of the virtual drive. | [optional] [readonly] 
**VendorUuid** | Pointer to **string** | The UUID value of the vendor. | [optional] [readonly] 
**VirtualDriveId** | Pointer to **string** | The identifier for this Virtual drive. | [optional] [readonly] 
**DiskGroup** | Pointer to [**StorageDiskGroupRelationship**](storage.DiskGroup.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PhysicalDiskUsages** | Pointer to [**[]StoragePhysicalDiskUsageRelationship**](storage.PhysicalDiskUsage.Relationship.md) | An array of relationships to storagePhysicalDiskUsage resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**VdMemberEps** | Pointer to [**[]StorageVdMemberEpRelationship**](storage.VdMemberEp.Relationship.md) | An array of relationships to storageVdMemberEp resources. | [optional] [readonly] 
**VirtualDriveExtension** | Pointer to [**StorageVirtualDriveExtensionRelationship**](storage.VirtualDriveExtension.Relationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveRelationship

`func NewStorageVirtualDriveRelationship(classId string, objectType string, ) *StorageVirtualDriveRelationship`

NewStorageVirtualDriveRelationship instantiates a new StorageVirtualDriveRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveRelationshipWithDefaults

`func NewStorageVirtualDriveRelationshipWithDefaults() *StorageVirtualDriveRelationship`

NewStorageVirtualDriveRelationshipWithDefaults instantiates a new StorageVirtualDriveRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageVirtualDriveRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageVirtualDriveRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageVirtualDriveRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageVirtualDriveRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageVirtualDriveRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageVirtualDriveRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageVirtualDriveRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageVirtualDriveRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageVirtualDriveRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageVirtualDriveRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageVirtualDriveRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageVirtualDriveRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageVirtualDriveRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageVirtualDriveRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageVirtualDriveRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageVirtualDriveRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageVirtualDriveRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageVirtualDriveRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageVirtualDriveRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageVirtualDriveRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageVirtualDriveRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageVirtualDriveRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageVirtualDriveRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageVirtualDriveRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageVirtualDriveRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageVirtualDriveRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageVirtualDriveRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageVirtualDriveRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageVirtualDriveRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageVirtualDriveRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageVirtualDriveRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageVirtualDriveRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageVirtualDriveRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageVirtualDriveRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageVirtualDriveRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageVirtualDriveRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageVirtualDriveRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageVirtualDriveRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageVirtualDriveRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageVirtualDriveRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageVirtualDriveRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageVirtualDriveRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageVirtualDriveRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageVirtualDriveRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageVirtualDriveRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageVirtualDriveRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageVirtualDriveRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageVirtualDriveRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageVirtualDriveRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageVirtualDriveRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageVirtualDriveRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageVirtualDriveRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageVirtualDriveRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageVirtualDriveRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageVirtualDriveRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageVirtualDriveRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageVirtualDriveRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageVirtualDriveRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageVirtualDriveRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageVirtualDriveRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageVirtualDriveRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageVirtualDriveRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageVirtualDriveRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageVirtualDriveRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageVirtualDriveRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageVirtualDriveRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageVirtualDriveRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageVirtualDriveRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageVirtualDriveRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageVirtualDriveRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageVirtualDriveRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageVirtualDriveRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageVirtualDriveRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageVirtualDriveRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageVirtualDriveRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageVirtualDriveRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageVirtualDriveRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageVirtualDriveRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageVirtualDriveRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageVirtualDriveRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageVirtualDriveRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageVirtualDriveRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageVirtualDriveRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageVirtualDriveRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageVirtualDriveRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageVirtualDriveRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageVirtualDriveRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageVirtualDriveRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAccessPolicy

`func (o *StorageVirtualDriveRelationship) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageVirtualDriveRelationship) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageVirtualDriveRelationship) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageVirtualDriveRelationship) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetActualWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) GetActualWriteCachePolicy() string`

GetActualWriteCachePolicy returns the ActualWriteCachePolicy field if non-nil, zero value otherwise.

### GetActualWriteCachePolicyOk

`func (o *StorageVirtualDriveRelationship) GetActualWriteCachePolicyOk() (*string, bool)`

GetActualWriteCachePolicyOk returns a tuple with the ActualWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) SetActualWriteCachePolicy(v string)`

SetActualWriteCachePolicy sets ActualWriteCachePolicy field to given value.

### HasActualWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) HasActualWriteCachePolicy() bool`

HasActualWriteCachePolicy returns a boolean if a field has been set.

### GetAvailableSize

`func (o *StorageVirtualDriveRelationship) GetAvailableSize() string`

GetAvailableSize returns the AvailableSize field if non-nil, zero value otherwise.

### GetAvailableSizeOk

`func (o *StorageVirtualDriveRelationship) GetAvailableSizeOk() (*string, bool)`

GetAvailableSizeOk returns a tuple with the AvailableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSize

`func (o *StorageVirtualDriveRelationship) SetAvailableSize(v string)`

SetAvailableSize sets AvailableSize field to given value.

### HasAvailableSize

`func (o *StorageVirtualDriveRelationship) HasAvailableSize() bool`

HasAvailableSize returns a boolean if a field has been set.

### GetBlockSize

`func (o *StorageVirtualDriveRelationship) GetBlockSize() string`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *StorageVirtualDriveRelationship) GetBlockSizeOk() (*string, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *StorageVirtualDriveRelationship) SetBlockSize(v string)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *StorageVirtualDriveRelationship) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootable

`func (o *StorageVirtualDriveRelationship) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDriveRelationship) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDriveRelationship) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDriveRelationship) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetConfigState

`func (o *StorageVirtualDriveRelationship) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *StorageVirtualDriveRelationship) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *StorageVirtualDriveRelationship) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *StorageVirtualDriveRelationship) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) GetConfiguredWriteCachePolicy() string`

GetConfiguredWriteCachePolicy returns the ConfiguredWriteCachePolicy field if non-nil, zero value otherwise.

### GetConfiguredWriteCachePolicyOk

`func (o *StorageVirtualDriveRelationship) GetConfiguredWriteCachePolicyOk() (*string, bool)`

GetConfiguredWriteCachePolicyOk returns a tuple with the ConfiguredWriteCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) SetConfiguredWriteCachePolicy(v string)`

SetConfiguredWriteCachePolicy sets ConfiguredWriteCachePolicy field to given value.

### HasConfiguredWriteCachePolicy

`func (o *StorageVirtualDriveRelationship) HasConfiguredWriteCachePolicy() bool`

HasConfiguredWriteCachePolicy returns a boolean if a field has been set.

### GetConnectionProtocol

`func (o *StorageVirtualDriveRelationship) GetConnectionProtocol() string`

GetConnectionProtocol returns the ConnectionProtocol field if non-nil, zero value otherwise.

### GetConnectionProtocolOk

`func (o *StorageVirtualDriveRelationship) GetConnectionProtocolOk() (*string, bool)`

GetConnectionProtocolOk returns a tuple with the ConnectionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionProtocol

`func (o *StorageVirtualDriveRelationship) SetConnectionProtocol(v string)`

SetConnectionProtocol sets ConnectionProtocol field to given value.

### HasConnectionProtocol

`func (o *StorageVirtualDriveRelationship) HasConnectionProtocol() bool`

HasConnectionProtocol returns a boolean if a field has been set.

### GetDriveCache

`func (o *StorageVirtualDriveRelationship) GetDriveCache() string`

GetDriveCache returns the DriveCache field if non-nil, zero value otherwise.

### GetDriveCacheOk

`func (o *StorageVirtualDriveRelationship) GetDriveCacheOk() (*string, bool)`

GetDriveCacheOk returns a tuple with the DriveCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCache

`func (o *StorageVirtualDriveRelationship) SetDriveCache(v string)`

SetDriveCache sets DriveCache field to given value.

### HasDriveCache

`func (o *StorageVirtualDriveRelationship) HasDriveCache() bool`

HasDriveCache returns a boolean if a field has been set.

### GetDriveSecurity

`func (o *StorageVirtualDriveRelationship) GetDriveSecurity() string`

GetDriveSecurity returns the DriveSecurity field if non-nil, zero value otherwise.

### GetDriveSecurityOk

`func (o *StorageVirtualDriveRelationship) GetDriveSecurityOk() (*string, bool)`

GetDriveSecurityOk returns a tuple with the DriveSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSecurity

`func (o *StorageVirtualDriveRelationship) SetDriveSecurity(v string)`

SetDriveSecurity sets DriveSecurity field to given value.

### HasDriveSecurity

`func (o *StorageVirtualDriveRelationship) HasDriveSecurity() bool`

HasDriveSecurity returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDriveRelationship) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDriveRelationship) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDriveRelationship) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDriveRelationship) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetIoPolicy

`func (o *StorageVirtualDriveRelationship) GetIoPolicy() string`

GetIoPolicy returns the IoPolicy field if non-nil, zero value otherwise.

### GetIoPolicyOk

`func (o *StorageVirtualDriveRelationship) GetIoPolicyOk() (*string, bool)`

GetIoPolicyOk returns a tuple with the IoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoPolicy

`func (o *StorageVirtualDriveRelationship) SetIoPolicy(v string)`

SetIoPolicy sets IoPolicy field to given value.

### HasIoPolicy

`func (o *StorageVirtualDriveRelationship) HasIoPolicy() bool`

HasIoPolicy returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumBlocks

`func (o *StorageVirtualDriveRelationship) GetNumBlocks() string`

GetNumBlocks returns the NumBlocks field if non-nil, zero value otherwise.

### GetNumBlocksOk

`func (o *StorageVirtualDriveRelationship) GetNumBlocksOk() (*string, bool)`

GetNumBlocksOk returns a tuple with the NumBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBlocks

`func (o *StorageVirtualDriveRelationship) SetNumBlocks(v string)`

SetNumBlocks sets NumBlocks field to given value.

### HasNumBlocks

`func (o *StorageVirtualDriveRelationship) HasNumBlocks() bool`

HasNumBlocks returns a boolean if a field has been set.

### GetOperState

`func (o *StorageVirtualDriveRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageVirtualDriveRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageVirtualDriveRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageVirtualDriveRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageVirtualDriveRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageVirtualDriveRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageVirtualDriveRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageVirtualDriveRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPhysicalBlockSize

`func (o *StorageVirtualDriveRelationship) GetPhysicalBlockSize() string`

GetPhysicalBlockSize returns the PhysicalBlockSize field if non-nil, zero value otherwise.

### GetPhysicalBlockSizeOk

`func (o *StorageVirtualDriveRelationship) GetPhysicalBlockSizeOk() (*string, bool)`

GetPhysicalBlockSizeOk returns a tuple with the PhysicalBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBlockSize

`func (o *StorageVirtualDriveRelationship) SetPhysicalBlockSize(v string)`

SetPhysicalBlockSize sets PhysicalBlockSize field to given value.

### HasPhysicalBlockSize

`func (o *StorageVirtualDriveRelationship) HasPhysicalBlockSize() bool`

HasPhysicalBlockSize returns a boolean if a field has been set.

### GetPresence

`func (o *StorageVirtualDriveRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageVirtualDriveRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageVirtualDriveRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageVirtualDriveRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageVirtualDriveRelationship) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageVirtualDriveRelationship) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageVirtualDriveRelationship) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageVirtualDriveRelationship) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetSecurityFlags

`func (o *StorageVirtualDriveRelationship) GetSecurityFlags() string`

GetSecurityFlags returns the SecurityFlags field if non-nil, zero value otherwise.

### GetSecurityFlagsOk

`func (o *StorageVirtualDriveRelationship) GetSecurityFlagsOk() (*string, bool)`

GetSecurityFlagsOk returns a tuple with the SecurityFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFlags

`func (o *StorageVirtualDriveRelationship) SetSecurityFlags(v string)`

SetSecurityFlags sets SecurityFlags field to given value.

### HasSecurityFlags

`func (o *StorageVirtualDriveRelationship) HasSecurityFlags() bool`

HasSecurityFlags returns a boolean if a field has been set.

### GetSize

`func (o *StorageVirtualDriveRelationship) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVirtualDriveRelationship) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVirtualDriveRelationship) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageVirtualDriveRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStripSize

`func (o *StorageVirtualDriveRelationship) GetStripSize() string`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageVirtualDriveRelationship) GetStripSizeOk() (*string, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageVirtualDriveRelationship) SetStripSize(v string)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageVirtualDriveRelationship) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetType

`func (o *StorageVirtualDriveRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageVirtualDriveRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageVirtualDriveRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageVirtualDriveRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDriveRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDriveRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDriveRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDriveRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDriveRelationship) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDriveRelationship) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDriveRelationship) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDriveRelationship) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDriveRelationship) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDriveRelationship) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDriveRelationship) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDriveRelationship) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetDiskGroup

`func (o *StorageVirtualDriveRelationship) GetDiskGroup() StorageDiskGroupRelationship`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *StorageVirtualDriveRelationship) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *StorageVirtualDriveRelationship) SetDiskGroup(v StorageDiskGroupRelationship)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *StorageVirtualDriveRelationship) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDriveRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDriveRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDriveRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDriveRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPhysicalDiskUsages

`func (o *StorageVirtualDriveRelationship) GetPhysicalDiskUsages() []StoragePhysicalDiskUsageRelationship`

GetPhysicalDiskUsages returns the PhysicalDiskUsages field if non-nil, zero value otherwise.

### GetPhysicalDiskUsagesOk

`func (o *StorageVirtualDriveRelationship) GetPhysicalDiskUsagesOk() (*[]StoragePhysicalDiskUsageRelationship, bool)`

GetPhysicalDiskUsagesOk returns a tuple with the PhysicalDiskUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDiskUsages

`func (o *StorageVirtualDriveRelationship) SetPhysicalDiskUsages(v []StoragePhysicalDiskUsageRelationship)`

SetPhysicalDiskUsages sets PhysicalDiskUsages field to given value.

### HasPhysicalDiskUsages

`func (o *StorageVirtualDriveRelationship) HasPhysicalDiskUsages() bool`

HasPhysicalDiskUsages returns a boolean if a field has been set.

### SetPhysicalDiskUsagesNil

`func (o *StorageVirtualDriveRelationship) SetPhysicalDiskUsagesNil(b bool)`

 SetPhysicalDiskUsagesNil sets the value for PhysicalDiskUsages to be an explicit nil

### UnsetPhysicalDiskUsages
`func (o *StorageVirtualDriveRelationship) UnsetPhysicalDiskUsages()`

UnsetPhysicalDiskUsages ensures that no value is present for PhysicalDiskUsages, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageVirtualDriveRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDriveRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDriveRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDriveRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDriveRelationship) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDriveRelationship) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDriveRelationship) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDriveRelationship) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVdMemberEps

`func (o *StorageVirtualDriveRelationship) GetVdMemberEps() []StorageVdMemberEpRelationship`

GetVdMemberEps returns the VdMemberEps field if non-nil, zero value otherwise.

### GetVdMemberEpsOk

`func (o *StorageVirtualDriveRelationship) GetVdMemberEpsOk() (*[]StorageVdMemberEpRelationship, bool)`

GetVdMemberEpsOk returns a tuple with the VdMemberEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdMemberEps

`func (o *StorageVirtualDriveRelationship) SetVdMemberEps(v []StorageVdMemberEpRelationship)`

SetVdMemberEps sets VdMemberEps field to given value.

### HasVdMemberEps

`func (o *StorageVirtualDriveRelationship) HasVdMemberEps() bool`

HasVdMemberEps returns a boolean if a field has been set.

### SetVdMemberEpsNil

`func (o *StorageVirtualDriveRelationship) SetVdMemberEpsNil(b bool)`

 SetVdMemberEpsNil sets the value for VdMemberEps to be an explicit nil

### UnsetVdMemberEps
`func (o *StorageVirtualDriveRelationship) UnsetVdMemberEps()`

UnsetVdMemberEps ensures that no value is present for VdMemberEps, not even an explicit nil
### GetVirtualDriveExtension

`func (o *StorageVirtualDriveRelationship) GetVirtualDriveExtension() StorageVirtualDriveExtensionRelationship`

GetVirtualDriveExtension returns the VirtualDriveExtension field if non-nil, zero value otherwise.

### GetVirtualDriveExtensionOk

`func (o *StorageVirtualDriveRelationship) GetVirtualDriveExtensionOk() (*StorageVirtualDriveExtensionRelationship, bool)`

GetVirtualDriveExtensionOk returns a tuple with the VirtualDriveExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveExtension

`func (o *StorageVirtualDriveRelationship) SetVirtualDriveExtension(v StorageVirtualDriveExtensionRelationship)`

SetVirtualDriveExtension sets VirtualDriveExtension field to given value.

### HasVirtualDriveExtension

`func (o *StorageVirtualDriveRelationship) HasVirtualDriveExtension() bool`

HasVirtualDriveExtension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


