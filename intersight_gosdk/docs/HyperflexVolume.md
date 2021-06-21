# HyperflexVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Volume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Volume"]
**Capacity** | Pointer to **int64** | Provisioned Capacity of the volume in bytes. | [optional] [readonly] 
**ClientId** | Pointer to **string** | Client (tenant) ID to which the volume belongs. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | Last modified time as UTC of the volume. | [optional] [readonly] 
**LunUuid** | Pointer to **string** | UUID of LUN associated with the volume. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number of the volume. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The unique identifier for this volume. | [optional] [readonly] 
**VolumeAccessMode** | Pointer to **string** | Access Mode of the volume. * &#x60;ReadWriteOnce&#x60; - Read write permisisons to a Virtual disk by a single virtual machine. * &#x60;ReadWriteMany&#x60; - Read write permisisons to a Virtual disk by multiple virtual machines. * &#x60;ReadOnlyMany&#x60; - Read only permisisons to a Virtual disk by multiple virtual machines. * &#x60;&#x60; - Unknown disk access mode. | [optional] [readonly] [default to "ReadWriteOnce"]
**VolumeMode** | Pointer to **string** | The mode of the HyperFlex volume. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [readonly] [default to "Block"]
**VolumeType** | Pointer to **string** | The type of the HyperFlex volume. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**StorageContainer** | Pointer to [**HyperflexStorageContainerRelationship**](hyperflex.StorageContainer.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexVolume

`func NewHyperflexVolume(classId string, objectType string, ) *HyperflexVolume`

NewHyperflexVolume instantiates a new HyperflexVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVolumeWithDefaults

`func NewHyperflexVolumeWithDefaults() *HyperflexVolume`

NewHyperflexVolumeWithDefaults instantiates a new HyperflexVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *HyperflexVolume) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HyperflexVolume) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HyperflexVolume) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HyperflexVolume) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClientId

`func (o *HyperflexVolume) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *HyperflexVolume) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *HyperflexVolume) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *HyperflexVolume) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *HyperflexVolume) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HyperflexVolume) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HyperflexVolume) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *HyperflexVolume) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLunUuid

`func (o *HyperflexVolume) GetLunUuid() string`

GetLunUuid returns the LunUuid field if non-nil, zero value otherwise.

### GetLunUuidOk

`func (o *HyperflexVolume) GetLunUuidOk() (*string, bool)`

GetLunUuidOk returns a tuple with the LunUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunUuid

`func (o *HyperflexVolume) SetLunUuid(v string)`

SetLunUuid sets LunUuid field to given value.

### HasLunUuid

`func (o *HyperflexVolume) HasLunUuid() bool`

HasLunUuid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HyperflexVolume) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HyperflexVolume) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HyperflexVolume) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HyperflexVolume) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeAccessMode

`func (o *HyperflexVolume) GetVolumeAccessMode() string`

GetVolumeAccessMode returns the VolumeAccessMode field if non-nil, zero value otherwise.

### GetVolumeAccessModeOk

`func (o *HyperflexVolume) GetVolumeAccessModeOk() (*string, bool)`

GetVolumeAccessModeOk returns a tuple with the VolumeAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAccessMode

`func (o *HyperflexVolume) SetVolumeAccessMode(v string)`

SetVolumeAccessMode sets VolumeAccessMode field to given value.

### HasVolumeAccessMode

`func (o *HyperflexVolume) HasVolumeAccessMode() bool`

HasVolumeAccessMode returns a boolean if a field has been set.

### GetVolumeMode

`func (o *HyperflexVolume) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *HyperflexVolume) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *HyperflexVolume) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *HyperflexVolume) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeType

`func (o *HyperflexVolume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *HyperflexVolume) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *HyperflexVolume) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *HyperflexVolume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexVolume) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexVolume) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexVolume) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexVolume) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetStorageContainer

`func (o *HyperflexVolume) GetStorageContainer() HyperflexStorageContainerRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *HyperflexVolume) GetStorageContainerOk() (*HyperflexStorageContainerRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *HyperflexVolume) SetStorageContainer(v HyperflexStorageContainerRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *HyperflexVolume) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


