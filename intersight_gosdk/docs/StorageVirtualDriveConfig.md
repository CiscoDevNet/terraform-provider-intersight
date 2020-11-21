# StorageVirtualDriveConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.VirtualDriveConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.VirtualDriveConfig"]
**AccessPolicy** | Pointer to **string** | Access policy that host has on this virtual drive. * &#x60;Default&#x60; - Use platform default access mode. * &#x60;ReadWrite&#x60; - Enables host to perform read-write on the VD. * &#x60;ReadOnly&#x60; - Host can only read from the VD. * &#x60;Blocked&#x60; - Host can neither read nor write to the VD. | [optional] [default to "Default"]
**BootDrive** | Pointer to **bool** | The flag enables the use of this virtual drive as a boot drive. | [optional] 
**DiskGroupName** | Pointer to **string** | Disk group policy that has the disk group in which this virtual drive needs to be created. | [optional] [readonly] 
**DiskGroupPolicy** | Pointer to **string** | Disk group policy that has the disk group in which this virtual drive needs to be created. | [optional] 
**DriveCache** | Pointer to **string** | Drive Cache property expect disk cache policy. * &#x60;Default&#x60; - Use platform default drive cache mode. * &#x60;NoChange&#x60; - Drive cache policy is unchanged. * &#x60;Enable&#x60; - Enables IO caching on the drive. * &#x60;Disable&#x60; - Disables IO caching on the drive. | [optional] [default to "Default"]
**ExpandToAvailable** | Pointer to **bool** | The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored. | [optional] 
**IoPolicy** | Pointer to **string** | Desired IO mode - direct IO or cached IO. * &#x60;Default&#x60; - Use platform default IO mode. * &#x60;Direct&#x60; - Use direct IO for writing directly into the disk. * &#x60;Cached&#x60; - Use cached IO for writing into cache and then to disk. | [optional] [default to "Default"]
**Name** | Pointer to **string** | The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. | [optional] 
**ReadPolicy** | Pointer to **string** | Read ahead mode to be used to read data from this virtual drive. * &#x60;Default&#x60; - Use platform default read ahead mode. * &#x60;ReadAhead&#x60; - Use read ahead mode for the policy. * &#x60;NoReadAhead&#x60; - Do not use read ahead mode for the policy. | [optional] [default to "Default"]
**Size** | Pointer to **int64** | Virtual drive size in MB. Size is mandatory field unless the &#39;Expand to Available&#39; option is enabled. | [optional] 
**StripSize** | Pointer to **string** | The strip size is the portion of a stripe that resides on a single drive in the drive group. The stripe consists of the data segments that the RAID controller writes across multiple drives, not including parity drives. * &#x60;Default&#x60; - Use platform default strip size for a virtual drive. * &#x60;32k&#x60; - Enables a strip size of 32k for a virtual drive. * &#x60;64k&#x60; - Enables a strip size of 64k for a virtual drive. * &#x60;128k&#x60; - Enables a strip size of 128k for a virtual drive. * &#x60;256k&#x60; - Enables a strip size of 256k for a virtual drive. * &#x60;512k&#x60; - Enables a strip size of 512k for a virtual drive. * &#x60;1024k&#x60; - Enables a strip size of 1024k for a virtual drive. | [optional] [default to "Default"]
**Vdid** | Pointer to **string** | Unique Id of the Virtual Drive to be used to identify Virtual Drive when name is empty. | [optional] [readonly] 
**WritePolicy** | Pointer to **string** | Write mode to be used to write data to this virtual drive. * &#x60;Default&#x60; - Use platform default write mode. * &#x60;WriteThrough&#x60; - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * &#x60;WriteBackGoodBbu&#x60; - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * &#x60;AlwaysWriteBack&#x60; - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. | [optional] [default to "Default"]

## Methods

### NewStorageVirtualDriveConfig

`func NewStorageVirtualDriveConfig(classId string, objectType string, ) *StorageVirtualDriveConfig`

NewStorageVirtualDriveConfig instantiates a new StorageVirtualDriveConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveConfigWithDefaults

`func NewStorageVirtualDriveConfigWithDefaults() *StorageVirtualDriveConfig`

NewStorageVirtualDriveConfigWithDefaults instantiates a new StorageVirtualDriveConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageVirtualDriveConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageVirtualDriveConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessPolicy

`func (o *StorageVirtualDriveConfig) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageVirtualDriveConfig) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageVirtualDriveConfig) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageVirtualDriveConfig) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetBootDrive

`func (o *StorageVirtualDriveConfig) GetBootDrive() bool`

GetBootDrive returns the BootDrive field if non-nil, zero value otherwise.

### GetBootDriveOk

`func (o *StorageVirtualDriveConfig) GetBootDriveOk() (*bool, bool)`

GetBootDriveOk returns a tuple with the BootDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDrive

`func (o *StorageVirtualDriveConfig) SetBootDrive(v bool)`

SetBootDrive sets BootDrive field to given value.

### HasBootDrive

`func (o *StorageVirtualDriveConfig) HasBootDrive() bool`

HasBootDrive returns a boolean if a field has been set.

### GetDiskGroupName

`func (o *StorageVirtualDriveConfig) GetDiskGroupName() string`

GetDiskGroupName returns the DiskGroupName field if non-nil, zero value otherwise.

### GetDiskGroupNameOk

`func (o *StorageVirtualDriveConfig) GetDiskGroupNameOk() (*string, bool)`

GetDiskGroupNameOk returns a tuple with the DiskGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroupName

`func (o *StorageVirtualDriveConfig) SetDiskGroupName(v string)`

SetDiskGroupName sets DiskGroupName field to given value.

### HasDiskGroupName

`func (o *StorageVirtualDriveConfig) HasDiskGroupName() bool`

HasDiskGroupName returns a boolean if a field has been set.

### GetDiskGroupPolicy

`func (o *StorageVirtualDriveConfig) GetDiskGroupPolicy() string`

GetDiskGroupPolicy returns the DiskGroupPolicy field if non-nil, zero value otherwise.

### GetDiskGroupPolicyOk

`func (o *StorageVirtualDriveConfig) GetDiskGroupPolicyOk() (*string, bool)`

GetDiskGroupPolicyOk returns a tuple with the DiskGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroupPolicy

`func (o *StorageVirtualDriveConfig) SetDiskGroupPolicy(v string)`

SetDiskGroupPolicy sets DiskGroupPolicy field to given value.

### HasDiskGroupPolicy

`func (o *StorageVirtualDriveConfig) HasDiskGroupPolicy() bool`

HasDiskGroupPolicy returns a boolean if a field has been set.

### GetDriveCache

`func (o *StorageVirtualDriveConfig) GetDriveCache() string`

GetDriveCache returns the DriveCache field if non-nil, zero value otherwise.

### GetDriveCacheOk

`func (o *StorageVirtualDriveConfig) GetDriveCacheOk() (*string, bool)`

GetDriveCacheOk returns a tuple with the DriveCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCache

`func (o *StorageVirtualDriveConfig) SetDriveCache(v string)`

SetDriveCache sets DriveCache field to given value.

### HasDriveCache

`func (o *StorageVirtualDriveConfig) HasDriveCache() bool`

HasDriveCache returns a boolean if a field has been set.

### GetExpandToAvailable

`func (o *StorageVirtualDriveConfig) GetExpandToAvailable() bool`

GetExpandToAvailable returns the ExpandToAvailable field if non-nil, zero value otherwise.

### GetExpandToAvailableOk

`func (o *StorageVirtualDriveConfig) GetExpandToAvailableOk() (*bool, bool)`

GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandToAvailable

`func (o *StorageVirtualDriveConfig) SetExpandToAvailable(v bool)`

SetExpandToAvailable sets ExpandToAvailable field to given value.

### HasExpandToAvailable

`func (o *StorageVirtualDriveConfig) HasExpandToAvailable() bool`

HasExpandToAvailable returns a boolean if a field has been set.

### GetIoPolicy

`func (o *StorageVirtualDriveConfig) GetIoPolicy() string`

GetIoPolicy returns the IoPolicy field if non-nil, zero value otherwise.

### GetIoPolicyOk

`func (o *StorageVirtualDriveConfig) GetIoPolicyOk() (*string, bool)`

GetIoPolicyOk returns a tuple with the IoPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoPolicy

`func (o *StorageVirtualDriveConfig) SetIoPolicy(v string)`

SetIoPolicy sets IoPolicy field to given value.

### HasIoPolicy

`func (o *StorageVirtualDriveConfig) HasIoPolicy() bool`

HasIoPolicy returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageVirtualDriveConfig) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageVirtualDriveConfig) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageVirtualDriveConfig) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageVirtualDriveConfig) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetSize

`func (o *StorageVirtualDriveConfig) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVirtualDriveConfig) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVirtualDriveConfig) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageVirtualDriveConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStripSize

`func (o *StorageVirtualDriveConfig) GetStripSize() string`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageVirtualDriveConfig) GetStripSizeOk() (*string, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageVirtualDriveConfig) SetStripSize(v string)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageVirtualDriveConfig) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetVdid

`func (o *StorageVirtualDriveConfig) GetVdid() string`

GetVdid returns the Vdid field if non-nil, zero value otherwise.

### GetVdidOk

`func (o *StorageVirtualDriveConfig) GetVdidOk() (*string, bool)`

GetVdidOk returns a tuple with the Vdid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdid

`func (o *StorageVirtualDriveConfig) SetVdid(v string)`

SetVdid sets Vdid field to given value.

### HasVdid

`func (o *StorageVirtualDriveConfig) HasVdid() bool`

HasVdid returns a boolean if a field has been set.

### GetWritePolicy

`func (o *StorageVirtualDriveConfig) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *StorageVirtualDriveConfig) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *StorageVirtualDriveConfig) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *StorageVirtualDriveConfig) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


