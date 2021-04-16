# StorageHyperFlexStorageContainerAllOf

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

### NewStorageHyperFlexStorageContainerAllOf

`func NewStorageHyperFlexStorageContainerAllOf(classId string, objectType string, ) *StorageHyperFlexStorageContainerAllOf`

NewStorageHyperFlexStorageContainerAllOf instantiates a new StorageHyperFlexStorageContainerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHyperFlexStorageContainerAllOfWithDefaults

`func NewStorageHyperFlexStorageContainerAllOfWithDefaults() *StorageHyperFlexStorageContainerAllOf`

NewStorageHyperFlexStorageContainerAllOfWithDefaults instantiates a new StorageHyperFlexStorageContainerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHyperFlexStorageContainerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHyperFlexStorageContainerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHyperFlexStorageContainerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHyperFlexStorageContainerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) GetCapacityUtilization() float32`

GetCapacityUtilization returns the CapacityUtilization field if non-nil, zero value otherwise.

### GetCapacityUtilizationOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetCapacityUtilizationOk() (*float32, bool)`

GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) SetCapacityUtilization(v float32)`

SetCapacityUtilization sets CapacityUtilization field to given value.

### HasCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) HasCapacityUtilization() bool`

HasCapacityUtilization returns a boolean if a field has been set.

### GetDataBlockSize

`func (o *StorageHyperFlexStorageContainerAllOf) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *StorageHyperFlexStorageContainerAllOf) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *StorageHyperFlexStorageContainerAllOf) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetInUse

`func (o *StorageHyperFlexStorageContainerAllOf) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *StorageHyperFlexStorageContainerAllOf) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *StorageHyperFlexStorageContainerAllOf) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *StorageHyperFlexStorageContainerAllOf) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *StorageHyperFlexStorageContainerAllOf) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *StorageHyperFlexStorageContainerAllOf) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *StorageHyperFlexStorageContainerAllOf) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *StorageHyperFlexStorageContainerAllOf) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *StorageHyperFlexStorageContainerAllOf) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *StorageHyperFlexStorageContainerAllOf) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *StorageHyperFlexStorageContainerAllOf) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilization() float32`

GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field if non-nil, zero value otherwise.

### GetProvisionedVolumeCapacityUtilizationOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool)`

GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) SetProvisionedVolumeCapacityUtilization(v float32)`

SetProvisionedVolumeCapacityUtilization sets ProvisionedVolumeCapacityUtilization field to given value.

### HasProvisionedVolumeCapacityUtilization

`func (o *StorageHyperFlexStorageContainerAllOf) HasProvisionedVolumeCapacityUtilization() bool`

HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.

### GetType

`func (o *StorageHyperFlexStorageContainerAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageHyperFlexStorageContainerAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageHyperFlexStorageContainerAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainerAllOf) GetUnCompressedUsedBytes() int64`

GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field if non-nil, zero value otherwise.

### GetUnCompressedUsedBytesOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetUnCompressedUsedBytesOk() (*int64, bool)`

GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainerAllOf) SetUnCompressedUsedBytes(v int64)`

SetUnCompressedUsedBytes sets UnCompressedUsedBytes field to given value.

### HasUnCompressedUsedBytes

`func (o *StorageHyperFlexStorageContainerAllOf) HasUnCompressedUsedBytes() bool`

HasUnCompressedUsedBytes returns a boolean if a field has been set.

### GetUuid

`func (o *StorageHyperFlexStorageContainerAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageHyperFlexStorageContainerAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageHyperFlexStorageContainerAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCount

`func (o *StorageHyperFlexStorageContainerAllOf) GetVolumeCount() int64`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetVolumeCountOk() (*int64, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *StorageHyperFlexStorageContainerAllOf) SetVolumeCount(v int64)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *StorageHyperFlexStorageContainerAllOf) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetCluster

`func (o *StorageHyperFlexStorageContainerAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *StorageHyperFlexStorageContainerAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *StorageHyperFlexStorageContainerAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHyperFlexStorageContainerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHyperFlexStorageContainerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHyperFlexStorageContainerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHyperFlexStorageContainerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


