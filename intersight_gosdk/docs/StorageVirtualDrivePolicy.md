# StorageVirtualDrivePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.VirtualDrivePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.VirtualDrivePolicy"]
**AccessPolicy** | Pointer to **string** | Access policy that host has on this virtual drive. * &#x60;Default&#x60; - Use platform default access mode. * &#x60;ReadWrite&#x60; - Enables host to perform read-write on the VD. * &#x60;ReadOnly&#x60; - Host can only read from the VD. * &#x60;Blocked&#x60; - Host can neither read nor write to the VD. | [optional] [default to "Default"]
**DriveCache** | Pointer to **string** | Disk cache policy for the virtual drive. * &#x60;Default&#x60; - Use platform default drive cache mode. * &#x60;NoChange&#x60; - Drive cache policy is unchanged. * &#x60;Enable&#x60; - Enables IO caching on the drive. * &#x60;Disable&#x60; - Disables IO caching on the drive. | [optional] [default to "Default"]
**ReadPolicy** | Pointer to **string** | Read ahead mode to be used to read data from this virtual drive. * &#x60;Default&#x60; - Use platform default read ahead mode. * &#x60;ReadAhead&#x60; - Use read ahead mode for the policy. * &#x60;NoReadAhead&#x60; - Do not use read ahead mode for the policy. | [optional] [default to "Default"]
**StripSize** | Pointer to **int32** | Desired strip size - Allowed values are 64KiB, 128KiB, 256KiB, 512KiB, 1024KiB. * &#x60;64&#x60; - Number of bytes in a strip is 64 Kibibytes. * &#x60;128&#x60; - Number of bytes in a strip is 128 Kibibytes. * &#x60;256&#x60; - Number of bytes in a strip is 256 Kibibytes. * &#x60;512&#x60; - Number of bytes in a strip is 512 Kibibytes. * &#x60;1024&#x60; - Number of bytes in a strip is 1024 Kibibytes or 1 Mebibyte. | [optional] [default to 64]
**WritePolicy** | Pointer to **string** | Write mode to be used to write data to this virtual drive. * &#x60;Default&#x60; - Use platform default write mode. * &#x60;WriteThrough&#x60; - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * &#x60;WriteBackGoodBbu&#x60; - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * &#x60;AlwaysWriteBack&#x60; - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. | [optional] [default to "Default"]

## Methods

### NewStorageVirtualDrivePolicy

`func NewStorageVirtualDrivePolicy(classId string, objectType string, ) *StorageVirtualDrivePolicy`

NewStorageVirtualDrivePolicy instantiates a new StorageVirtualDrivePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDrivePolicyWithDefaults

`func NewStorageVirtualDrivePolicyWithDefaults() *StorageVirtualDrivePolicy`

NewStorageVirtualDrivePolicyWithDefaults instantiates a new StorageVirtualDrivePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageVirtualDrivePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDrivePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDrivePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageVirtualDrivePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDrivePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDrivePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessPolicy

`func (o *StorageVirtualDrivePolicy) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageVirtualDrivePolicy) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageVirtualDrivePolicy) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageVirtualDrivePolicy) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetDriveCache

`func (o *StorageVirtualDrivePolicy) GetDriveCache() string`

GetDriveCache returns the DriveCache field if non-nil, zero value otherwise.

### GetDriveCacheOk

`func (o *StorageVirtualDrivePolicy) GetDriveCacheOk() (*string, bool)`

GetDriveCacheOk returns a tuple with the DriveCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCache

`func (o *StorageVirtualDrivePolicy) SetDriveCache(v string)`

SetDriveCache sets DriveCache field to given value.

### HasDriveCache

`func (o *StorageVirtualDrivePolicy) HasDriveCache() bool`

HasDriveCache returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageVirtualDrivePolicy) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageVirtualDrivePolicy) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageVirtualDrivePolicy) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageVirtualDrivePolicy) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetStripSize

`func (o *StorageVirtualDrivePolicy) GetStripSize() int32`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageVirtualDrivePolicy) GetStripSizeOk() (*int32, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageVirtualDrivePolicy) SetStripSize(v int32)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageVirtualDrivePolicy) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetWritePolicy

`func (o *StorageVirtualDrivePolicy) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *StorageVirtualDrivePolicy) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *StorageVirtualDrivePolicy) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *StorageVirtualDrivePolicy) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


