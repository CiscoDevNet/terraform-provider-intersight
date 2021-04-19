# StorageHyperFlexStorageContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HyperFlexStorageContainer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HyperFlexStorageContainer"]
**CapacityUtilization** | Pointer to **float32** | Capacity Utilization of Storage Container. | [optional] [readonly] 
**DataBlockSize** | Pointer to **int64** | Storage Container data block size | [optional] [readonly] 
**InUse** | Pointer to **bool** | Indicates whether the Storage Container has Volumes. | [optional] [readonly] 
**LastAccessTime** | Pointer to **time.Time** | Storage container&#39;s last access time. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | Storage container&#39;s last modified time. | [optional] [readonly] 
**ProvisionedCapacity** | Pointer to **int64** | Provisioned Capacity of the Storage container. | [optional] [readonly] 
**ProvisionedVolumeCapacityUtilization** | Pointer to **float32** | Provisioned Capacity Utilization of All Volumes associated with the Storage Container. | [optional] [readonly] 
**Type** | Pointer to **string** | Storage Container type (SMB/NFS/iSCSI). * &#x60;NFS&#x60; - Storage container created/accesed through NFS protocol. * &#x60;SMB&#x60; - Storage container created/accessed through SMB protocol. * &#x60;iSCSI&#x60; - Storage container created/accessed through iSCSI protocol. | [optional] [readonly] [default to "NFS"]
**UnCompressedUsedBytes** | Pointer to **int64** | Uncompressed bytes on Storage Container. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the Datastore/Storage Containter. | [optional] [readonly] 
**VolumeCount** | Pointer to **int64** | Number of Volumes associated with the Storage Container. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageHyperFlexStorageContainer

`func NewStorageHyperFlexStorageContainer(classId string, objectType string, ) *StorageHyperFlexStorageContainer`

NewStorageHyperFlexStorageContainer instantiates a new StorageHyperFlexStorageContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHyperFlexStorageContainerWithDefaults

`func NewStorageHyperFlexStorageContainerWithDefaults() *StorageHyperFlexStorageContainer`

NewStorageHyperFlexStorageContainerWithDefaults instantiates a new StorageHyperFlexStorageContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHyperFlexStorageContainer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHyperFlexStorageContainer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHyperFlexStorageContainer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHyperFlexStorageContainer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHyperFlexStorageContainer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHyperFlexStorageContainer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) GetCapacityUtilization() float32`

GetCapacityUtilization returns the CapacityUtilization field if non-nil, zero value otherwise.

### GetCapacityUtilizationOk

`func (o *StorageHyperFlexStorageContainer) GetCapacityUtilizationOk() (*float32, bool)`

GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) SetCapacityUtilization(v float32)`

SetCapacityUtilization sets CapacityUtilization field to given value.

### HasCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) HasCapacityUtilization() bool`

HasCapacityUtilization returns a boolean if a field has been set.

### GetDataBlockSize

`func (o *StorageHyperFlexStorageContainer) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *StorageHyperFlexStorageContainer) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *StorageHyperFlexStorageContainer) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *StorageHyperFlexStorageContainer) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetInUse

`func (o *StorageHyperFlexStorageContainer) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *StorageHyperFlexStorageContainer) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *StorageHyperFlexStorageContainer) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *StorageHyperFlexStorageContainer) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *StorageHyperFlexStorageContainer) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *StorageHyperFlexStorageContainer) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *StorageHyperFlexStorageContainer) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *StorageHyperFlexStorageContainer) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *StorageHyperFlexStorageContainer) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *StorageHyperFlexStorageContainer) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *StorageHyperFlexStorageContainer) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *StorageHyperFlexStorageContainer) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *StorageHyperFlexStorageContainer) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *StorageHyperFlexStorageContainer) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *StorageHyperFlexStorageContainer) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *StorageHyperFlexStorageContainer) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) GetProvisionedVolumeCapacityUtilization() float32`

GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field if non-nil, zero value otherwise.

### GetProvisionedVolumeCapacityUtilizationOk

`func (o *StorageHyperFlexStorageContainer) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool)`

GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) SetProvisionedVolumeCapacityUtilization(v float32)`

SetProvisionedVolumeCapacityUtilization sets ProvisionedVolumeCapacityUtilization field to given value.

### HasProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainer) HasProvisionedVolumeCapacityUtilization() bool`

HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.

### GetType

`func (o *StorageHyperFlexStorageContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageHyperFlexStorageContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageHyperFlexStorageContainer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageHyperFlexStorageContainer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainer) GetUnCompressedUsedBytes() int64`

GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field if non-nil, zero value otherwise.

### GetUnCompressedUsedBytesOk

`func (o *StorageHyperFlexStorageContainer) GetUnCompressedUsedBytesOk() (*int64, bool)`

GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainer) SetUnCompressedUsedBytes(v int64)`

SetUnCompressedUsedBytes sets UnCompressedUsedBytes field to given value.

### HasUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainer) HasUnCompressedUsedBytes() bool`

HasUnCompressedUsedBytes returns a boolean if a field has been set.

### GetUuid

`func (o *StorageHyperFlexStorageContainer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageHyperFlexStorageContainer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageHyperFlexStorageContainer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageHyperFlexStorageContainer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCount

`func (o *StorageHyperFlexStorageContainer) GetVolumeCount() int64`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *StorageHyperFlexStorageContainer) GetVolumeCountOk() (*int64, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *StorageHyperFlexStorageContainer) SetVolumeCount(v int64)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *StorageHyperFlexStorageContainer) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetCluster

`func (o *StorageHyperFlexStorageContainer) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *StorageHyperFlexStorageContainer) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *StorageHyperFlexStorageContainer) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *StorageHyperFlexStorageContainer) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHyperFlexStorageContainer) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHyperFlexStorageContainer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHyperFlexStorageContainer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHyperFlexStorageContainer) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


