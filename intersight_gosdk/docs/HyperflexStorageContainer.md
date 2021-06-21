# HyperflexStorageContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.StorageContainer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.StorageContainer"]
**DataBlockSize** | Pointer to **int64** | Storage container data block size in bytes. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Indicates whether the storage container has volumes. | [optional] [readonly] 
**Kind** | Pointer to **string** | Indicates whether the storage container was user-created, or system-created. * &#x60;UNKNOWN&#x60; - The storage container creator is unknown. * &#x60;USER_CREATED&#x60; - The storage container was created by a user action. * &#x60;INTERNAL&#x60; - The storage container was created by the system. | [optional] [readonly] [default to "UNKNOWN"]
**LastAccessTime** | Pointer to **time.Time** | Storage container&#39;s last access time. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | Storage container&#39;s last modified time. | [optional] [readonly] 
**MountStatus** | Pointer to **string** | Storage container mount status. Applicable only for NFS type. * &#x60;NOT_APPLICABLE&#x60; - The HyperFlex storage container mount status is not applicable. * &#x60;NORMAL&#x60; - The HyperFlex storage container mount status is normal. * &#x60;ALERT&#x60; - The HyperFlex storage container mount status is alert. * &#x60;FAILED&#x60; - The HyperFlex storage container mount status is failed. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**MountSummary** | Pointer to **string** | Storage container mount summary. Applicable only for NFS type. * &#x60;NOT_APPLICABLE&#x60; - The mount summary is not applicable for this HyperFlex storage container. * &#x60;MOUNTED&#x60; - The HyperFlex storage container is mounted. * &#x60;UNMOUNTED&#x60; - The HyperFlex storage container is unmounted. * &#x60;MOUNT_FAILURE&#x60; - The HyperFlex storage container mount summary is failure. * &#x60;UNMOUNT_FAILURE&#x60; - The HyperFlex storage container unmount summary is failure. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**ProvisionedCapacity** | Pointer to **int64** | Provisioned capacity of the storage container in bytes. | [optional] [readonly] 
**ProvisionedVolumeCapacityUtilization** | Pointer to **float32** | Provisioned capacity utilization of all volumes associated with the storage container. | [optional] [readonly] 
**Type** | Pointer to **string** | Storage container type (SMB/NFS/iSCSI). * &#x60;NFS&#x60; - Storage container created/accesed through NFS protocol. * &#x60;SMB&#x60; - Storage container created/accessed through SMB protocol. * &#x60;iSCSI&#x60; - Storage container created/accessed through iSCSI protocol. | [optional] [readonly] [default to "NFS"]
**UnCompressedUsedBytes** | Pointer to **int64** | Uncompressed bytes on storage container. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the datastore/storage container. | [optional] [readonly] 
**VolumeCount** | Pointer to **int64** | Number of volumes associated with the storage container. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]HyperflexVolumeRelationship**](HyperflexVolumeRelationship.md) | An array of relationships to hyperflexVolume resources. | [optional] [readonly] 

## Methods

### NewHyperflexStorageContainer

`func NewHyperflexStorageContainer(classId string, objectType string, ) *HyperflexStorageContainer`

NewHyperflexStorageContainer instantiates a new HyperflexStorageContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStorageContainerWithDefaults

`func NewHyperflexStorageContainerWithDefaults() *HyperflexStorageContainer`

NewHyperflexStorageContainerWithDefaults instantiates a new HyperflexStorageContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStorageContainer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStorageContainer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStorageContainer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStorageContainer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStorageContainer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStorageContainer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataBlockSize

`func (o *HyperflexStorageContainer) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *HyperflexStorageContainer) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *HyperflexStorageContainer) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *HyperflexStorageContainer) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetInUse

`func (o *HyperflexStorageContainer) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *HyperflexStorageContainer) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *HyperflexStorageContainer) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *HyperflexStorageContainer) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetKind

`func (o *HyperflexStorageContainer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HyperflexStorageContainer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HyperflexStorageContainer) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HyperflexStorageContainer) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *HyperflexStorageContainer) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *HyperflexStorageContainer) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *HyperflexStorageContainer) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *HyperflexStorageContainer) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *HyperflexStorageContainer) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HyperflexStorageContainer) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HyperflexStorageContainer) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *HyperflexStorageContainer) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMountStatus

`func (o *HyperflexStorageContainer) GetMountStatus() string`

GetMountStatus returns the MountStatus field if non-nil, zero value otherwise.

### GetMountStatusOk

`func (o *HyperflexStorageContainer) GetMountStatusOk() (*string, bool)`

GetMountStatusOk returns a tuple with the MountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountStatus

`func (o *HyperflexStorageContainer) SetMountStatus(v string)`

SetMountStatus sets MountStatus field to given value.

### HasMountStatus

`func (o *HyperflexStorageContainer) HasMountStatus() bool`

HasMountStatus returns a boolean if a field has been set.

### GetMountSummary

`func (o *HyperflexStorageContainer) GetMountSummary() string`

GetMountSummary returns the MountSummary field if non-nil, zero value otherwise.

### GetMountSummaryOk

`func (o *HyperflexStorageContainer) GetMountSummaryOk() (*string, bool)`

GetMountSummaryOk returns a tuple with the MountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountSummary

`func (o *HyperflexStorageContainer) SetMountSummary(v string)`

SetMountSummary sets MountSummary field to given value.

### HasMountSummary

`func (o *HyperflexStorageContainer) HasMountSummary() bool`

HasMountSummary returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *HyperflexStorageContainer) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *HyperflexStorageContainer) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *HyperflexStorageContainer) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *HyperflexStorageContainer) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainer) GetProvisionedVolumeCapacityUtilization() float32`

GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field if non-nil, zero value otherwise.

### GetProvisionedVolumeCapacityUtilizationOk

`func (o *HyperflexStorageContainer) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool)`

GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainer) SetProvisionedVolumeCapacityUtilization(v float32)`

SetProvisionedVolumeCapacityUtilization sets ProvisionedVolumeCapacityUtilization field to given value.

### HasProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainer) HasProvisionedVolumeCapacityUtilization() bool`

HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.

### GetType

`func (o *HyperflexStorageContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexStorageContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexStorageContainer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexStorageContainer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnCompressedUsedBytes

`func (o *HyperflexStorageContainer) GetUnCompressedUsedBytes() int64`

GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field if non-nil, zero value otherwise.

### GetUnCompressedUsedBytesOk

`func (o *HyperflexStorageContainer) GetUnCompressedUsedBytesOk() (*int64, bool)`

GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCompressedUsedBytes

`func (o *HyperflexStorageContainer) SetUnCompressedUsedBytes(v int64)`

SetUnCompressedUsedBytes sets UnCompressedUsedBytes field to given value.

### HasUnCompressedUsedBytes

`func (o *HyperflexStorageContainer) HasUnCompressedUsedBytes() bool`

HasUnCompressedUsedBytes returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexStorageContainer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexStorageContainer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexStorageContainer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexStorageContainer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCount

`func (o *HyperflexStorageContainer) GetVolumeCount() int64`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *HyperflexStorageContainer) GetVolumeCountOk() (*int64, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *HyperflexStorageContainer) SetVolumeCount(v int64)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *HyperflexStorageContainer) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexStorageContainer) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexStorageContainer) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexStorageContainer) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexStorageContainer) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetVolumes

`func (o *HyperflexStorageContainer) GetVolumes() []HyperflexVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *HyperflexStorageContainer) GetVolumesOk() (*[]HyperflexVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *HyperflexStorageContainer) SetVolumes(v []HyperflexVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *HyperflexStorageContainer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *HyperflexStorageContainer) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *HyperflexStorageContainer) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


