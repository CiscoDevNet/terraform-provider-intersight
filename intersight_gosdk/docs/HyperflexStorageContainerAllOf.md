# HyperflexStorageContainerAllOf

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

### NewHyperflexStorageContainerAllOf

`func NewHyperflexStorageContainerAllOf(classId string, objectType string, ) *HyperflexStorageContainerAllOf`

NewHyperflexStorageContainerAllOf instantiates a new HyperflexStorageContainerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStorageContainerAllOfWithDefaults

`func NewHyperflexStorageContainerAllOfWithDefaults() *HyperflexStorageContainerAllOf`

NewHyperflexStorageContainerAllOfWithDefaults instantiates a new HyperflexStorageContainerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStorageContainerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStorageContainerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStorageContainerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStorageContainerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStorageContainerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStorageContainerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataBlockSize

`func (o *HyperflexStorageContainerAllOf) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *HyperflexStorageContainerAllOf) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *HyperflexStorageContainerAllOf) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *HyperflexStorageContainerAllOf) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetInUse

`func (o *HyperflexStorageContainerAllOf) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *HyperflexStorageContainerAllOf) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *HyperflexStorageContainerAllOf) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *HyperflexStorageContainerAllOf) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetKind

`func (o *HyperflexStorageContainerAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HyperflexStorageContainerAllOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HyperflexStorageContainerAllOf) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HyperflexStorageContainerAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *HyperflexStorageContainerAllOf) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *HyperflexStorageContainerAllOf) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *HyperflexStorageContainerAllOf) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *HyperflexStorageContainerAllOf) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *HyperflexStorageContainerAllOf) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HyperflexStorageContainerAllOf) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HyperflexStorageContainerAllOf) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *HyperflexStorageContainerAllOf) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMountStatus

`func (o *HyperflexStorageContainerAllOf) GetMountStatus() string`

GetMountStatus returns the MountStatus field if non-nil, zero value otherwise.

### GetMountStatusOk

`func (o *HyperflexStorageContainerAllOf) GetMountStatusOk() (*string, bool)`

GetMountStatusOk returns a tuple with the MountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountStatus

`func (o *HyperflexStorageContainerAllOf) SetMountStatus(v string)`

SetMountStatus sets MountStatus field to given value.

### HasMountStatus

`func (o *HyperflexStorageContainerAllOf) HasMountStatus() bool`

HasMountStatus returns a boolean if a field has been set.

### GetMountSummary

`func (o *HyperflexStorageContainerAllOf) GetMountSummary() string`

GetMountSummary returns the MountSummary field if non-nil, zero value otherwise.

### GetMountSummaryOk

`func (o *HyperflexStorageContainerAllOf) GetMountSummaryOk() (*string, bool)`

GetMountSummaryOk returns a tuple with the MountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountSummary

`func (o *HyperflexStorageContainerAllOf) SetMountSummary(v string)`

SetMountSummary sets MountSummary field to given value.

### HasMountSummary

`func (o *HyperflexStorageContainerAllOf) HasMountSummary() bool`

HasMountSummary returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *HyperflexStorageContainerAllOf) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *HyperflexStorageContainerAllOf) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *HyperflexStorageContainerAllOf) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *HyperflexStorageContainerAllOf) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilization() float32`

GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field if non-nil, zero value otherwise.

### GetProvisionedVolumeCapacityUtilizationOk

`func (o *HyperflexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool)`

GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainerAllOf) SetProvisionedVolumeCapacityUtilization(v float32)`

SetProvisionedVolumeCapacityUtilization sets ProvisionedVolumeCapacityUtilization field to given value.

### HasProvisionedVolumeCapacityUtilization

`func (o *HyperflexStorageContainerAllOf) HasProvisionedVolumeCapacityUtilization() bool`

HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.

### GetType

`func (o *HyperflexStorageContainerAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexStorageContainerAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexStorageContainerAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexStorageContainerAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnCompressedUsedBytes

`func (o *HyperflexStorageContainerAllOf) GetUnCompressedUsedBytes() int64`

GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field if non-nil, zero value otherwise.

### GetUnCompressedUsedBytesOk

`func (o *HyperflexStorageContainerAllOf) GetUnCompressedUsedBytesOk() (*int64, bool)`

GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCompressedUsedBytes

`func (o *HyperflexStorageContainerAllOf) SetUnCompressedUsedBytes(v int64)`

SetUnCompressedUsedBytes sets UnCompressedUsedBytes field to given value.

### HasUnCompressedUsedBytes

`func (o *HyperflexStorageContainerAllOf) HasUnCompressedUsedBytes() bool`

HasUnCompressedUsedBytes returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexStorageContainerAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexStorageContainerAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexStorageContainerAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexStorageContainerAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCount

`func (o *HyperflexStorageContainerAllOf) GetVolumeCount() int64`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *HyperflexStorageContainerAllOf) GetVolumeCountOk() (*int64, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *HyperflexStorageContainerAllOf) SetVolumeCount(v int64)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *HyperflexStorageContainerAllOf) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexStorageContainerAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexStorageContainerAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexStorageContainerAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexStorageContainerAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetVolumes

`func (o *HyperflexStorageContainerAllOf) GetVolumes() []HyperflexVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *HyperflexStorageContainerAllOf) GetVolumesOk() (*[]HyperflexVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *HyperflexStorageContainerAllOf) SetVolumes(v []HyperflexVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *HyperflexStorageContainerAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *HyperflexStorageContainerAllOf) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *HyperflexStorageContainerAllOf) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


