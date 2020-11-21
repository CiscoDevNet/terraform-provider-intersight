# StorageHyperFlexVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HyperFlexVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HyperFlexVolume"]
**Capacity** | Pointer to **int64** | Provisioned Capacity of the Storage container in bytes. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | Serial number of the volume. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the Datastore/Storage Container. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageContainer** | Pointer to [**StorageHyperFlexStorageContainerRelationship**](storage.HyperFlexStorageContainer.Relationship.md) |  | [optional] 

## Methods

### NewStorageHyperFlexVolume

`func NewStorageHyperFlexVolume(classId string, objectType string, ) *StorageHyperFlexVolume`

NewStorageHyperFlexVolume instantiates a new StorageHyperFlexVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHyperFlexVolumeWithDefaults

`func NewStorageHyperFlexVolumeWithDefaults() *StorageHyperFlexVolume`

NewStorageHyperFlexVolumeWithDefaults instantiates a new StorageHyperFlexVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHyperFlexVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHyperFlexVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHyperFlexVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHyperFlexVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHyperFlexVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHyperFlexVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *StorageHyperFlexVolume) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *StorageHyperFlexVolume) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *StorageHyperFlexVolume) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *StorageHyperFlexVolume) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetSerialNumber

`func (o *StorageHyperFlexVolume) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *StorageHyperFlexVolume) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *StorageHyperFlexVolume) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *StorageHyperFlexVolume) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUuid

`func (o *StorageHyperFlexVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageHyperFlexVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageHyperFlexVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageHyperFlexVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCluster

`func (o *StorageHyperFlexVolume) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *StorageHyperFlexVolume) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *StorageHyperFlexVolume) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *StorageHyperFlexVolume) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHyperFlexVolume) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHyperFlexVolume) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHyperFlexVolume) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHyperFlexVolume) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageHyperFlexVolume) GetStorageContainer() StorageHyperFlexStorageContainerRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageHyperFlexVolume) GetStorageContainerOk() (*StorageHyperFlexStorageContainerRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageHyperFlexVolume) SetStorageContainer(v StorageHyperFlexStorageContainerRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageHyperFlexVolume) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


