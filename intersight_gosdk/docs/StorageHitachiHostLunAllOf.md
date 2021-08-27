# StorageHitachiHostLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiHostLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiHostLun"]
**PortId** | Pointer to **string** | Port ID of the Hitachi host lun. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**Host** | Pointer to [**StorageHitachiHostRelationship**](StorageHitachiHostRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volume** | Pointer to [**StorageHitachiVolumeRelationship**](StorageHitachiVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiHostLunAllOf

`func NewStorageHitachiHostLunAllOf(classId string, objectType string, ) *StorageHitachiHostLunAllOf`

NewStorageHitachiHostLunAllOf instantiates a new StorageHitachiHostLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiHostLunAllOfWithDefaults

`func NewStorageHitachiHostLunAllOfWithDefaults() *StorageHitachiHostLunAllOf`

NewStorageHitachiHostLunAllOfWithDefaults instantiates a new StorageHitachiHostLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiHostLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiHostLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiHostLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiHostLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiHostLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiHostLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPortId

`func (o *StorageHitachiHostLunAllOf) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageHitachiHostLunAllOf) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageHitachiHostLunAllOf) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageHitachiHostLunAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiHostLunAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiHostLunAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiHostLunAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiHostLunAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHost

`func (o *StorageHitachiHostLunAllOf) GetHost() StorageHitachiHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageHitachiHostLunAllOf) GetHostOk() (*StorageHitachiHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageHitachiHostLunAllOf) SetHost(v StorageHitachiHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageHitachiHostLunAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiHostLunAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiHostLunAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiHostLunAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiHostLunAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVolume

`func (o *StorageHitachiHostLunAllOf) GetVolume() StorageHitachiVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageHitachiHostLunAllOf) GetVolumeOk() (*StorageHitachiVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageHitachiHostLunAllOf) SetVolume(v StorageHitachiVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageHitachiHostLunAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


