# StorageHyperFlexVolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HyperFlexVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HyperFlexVolume"]
**Capacity** | Pointer to **int64** | Provisioned Capacity of the Storage container in Bytes. | [optional] [readonly] 
**ClientId** | Pointer to **string** | Client ID to which the volume belongs. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | Last modified time as UTC of the volume. | [optional] [readonly] 
**LunUuid** | Pointer to **string** | UUID of Lun associated with the volume. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number of the volume. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the Datastore/Storage Containter. | [optional] [readonly] 
**VolumeAccessMode** | Pointer to **string** | Access Mode of the volume. * &#x60;ReadWriteOnce&#x60; - Read write permisisons to a Virtual disk by a single virtual machine. * &#x60;ReadWriteMany&#x60; - Read write permisisons to a Virtual disk by multiple virtual machines. * &#x60;ReadOnlyMany&#x60; - Read only permisisons to a Virtual disk by multiple virtual machines. * &#x60;&#x60; - Unknown disk access mode. | [optional] [readonly] [default to "ReadWriteOnce"]
**VolumeMode** | Pointer to **string** | Mode of the volume. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [readonly] [default to "Block"]
**VolumeType** | Pointer to **string** | The Type of the volume. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageContainer** | Pointer to [**StorageHyperFlexStorageContainerRelationship**](StorageHyperFlexStorageContainerRelationship.md) |  | [optional] 

## Methods

### NewStorageHyperFlexVolumeAllOf

`func NewStorageHyperFlexVolumeAllOf(classId string, objectType string, ) *StorageHyperFlexVolumeAllOf`

NewStorageHyperFlexVolumeAllOf instantiates a new StorageHyperFlexVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHyperFlexVolumeAllOfWithDefaults

`func NewStorageHyperFlexVolumeAllOfWithDefaults() *StorageHyperFlexVolumeAllOf`

NewStorageHyperFlexVolumeAllOfWithDefaults instantiates a new StorageHyperFlexVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHyperFlexVolumeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHyperFlexVolumeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHyperFlexVolumeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHyperFlexVolumeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHyperFlexVolumeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHyperFlexVolumeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *StorageHyperFlexVolumeAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageHyperFlexVolumeAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageHyperFlexVolumeAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageHyperFlexVolumeAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClientId

`func (o *StorageHyperFlexVolumeAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StorageHyperFlexVolumeAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StorageHyperFlexVolumeAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StorageHyperFlexVolumeAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *StorageHyperFlexVolumeAllOf) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *StorageHyperFlexVolumeAllOf) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *StorageHyperFlexVolumeAllOf) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *StorageHyperFlexVolumeAllOf) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLunUuid

`func (o *StorageHyperFlexVolumeAllOf) GetLunUuid() string`

GetLunUuid returns the LunUuid field if non-nil, zero value otherwise.

### GetLunUuidOk

`func (o *StorageHyperFlexVolumeAllOf) GetLunUuidOk() (*string, bool)`

GetLunUuidOk returns a tuple with the LunUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunUuid

`func (o *StorageHyperFlexVolumeAllOf) SetLunUuid(v string)`

SetLunUuid sets LunUuid field to given value.

### HasLunUuid

`func (o *StorageHyperFlexVolumeAllOf) HasLunUuid() bool`

HasLunUuid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *StorageHyperFlexVolumeAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *StorageHyperFlexVolumeAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *StorageHyperFlexVolumeAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *StorageHyperFlexVolumeAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUuid

`func (o *StorageHyperFlexVolumeAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageHyperFlexVolumeAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageHyperFlexVolumeAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageHyperFlexVolumeAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeAccessMode

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeAccessMode() string`

GetVolumeAccessMode returns the VolumeAccessMode field if non-nil, zero value otherwise.

### GetVolumeAccessModeOk

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeAccessModeOk() (*string, bool)`

GetVolumeAccessModeOk returns a tuple with the VolumeAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAccessMode

`func (o *StorageHyperFlexVolumeAllOf) SetVolumeAccessMode(v string)`

SetVolumeAccessMode sets VolumeAccessMode field to given value.

### HasVolumeAccessMode

`func (o *StorageHyperFlexVolumeAllOf) HasVolumeAccessMode() bool`

HasVolumeAccessMode returns a boolean if a field has been set.

### GetVolumeMode

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *StorageHyperFlexVolumeAllOf) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *StorageHyperFlexVolumeAllOf) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeType

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StorageHyperFlexVolumeAllOf) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StorageHyperFlexVolumeAllOf) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StorageHyperFlexVolumeAllOf) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetCluster

`func (o *StorageHyperFlexVolumeAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *StorageHyperFlexVolumeAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *StorageHyperFlexVolumeAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *StorageHyperFlexVolumeAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHyperFlexVolumeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHyperFlexVolumeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHyperFlexVolumeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHyperFlexVolumeAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageHyperFlexVolumeAllOf) GetStorageContainer() StorageHyperFlexStorageContainerRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageHyperFlexVolumeAllOf) GetStorageContainerOk() (*StorageHyperFlexStorageContainerRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageHyperFlexVolumeAllOf) SetStorageContainer(v StorageHyperFlexStorageContainerRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageHyperFlexVolumeAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


