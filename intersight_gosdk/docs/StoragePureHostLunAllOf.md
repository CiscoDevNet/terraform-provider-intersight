# StoragePureHostLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureHostLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureHostLun"]
**HostGroupName** | Pointer to **string** | Name of the host group associated with LUN. | [optional] [readonly] 
**Shared** | Pointer to **bool** | Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Host** | Pointer to [**StoragePureHostRelationship**](StoragePureHostRelationship.md) |  | [optional] 
**HostGroup** | Pointer to [**StoragePureHostGroupRelationship**](StoragePureHostGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Volume** | Pointer to [**StoragePureVolumeRelationship**](StoragePureVolumeRelationship.md) |  | [optional] 

## Methods

### NewStoragePureHostLunAllOf

`func NewStoragePureHostLunAllOf(classId string, objectType string, ) *StoragePureHostLunAllOf`

NewStoragePureHostLunAllOf instantiates a new StoragePureHostLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureHostLunAllOfWithDefaults

`func NewStoragePureHostLunAllOfWithDefaults() *StoragePureHostLunAllOf`

NewStoragePureHostLunAllOfWithDefaults instantiates a new StoragePureHostLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureHostLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureHostLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureHostLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureHostLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureHostLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureHostLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostGroupName

`func (o *StoragePureHostLunAllOf) GetHostGroupName() string`

GetHostGroupName returns the HostGroupName field if non-nil, zero value otherwise.

### GetHostGroupNameOk

`func (o *StoragePureHostLunAllOf) GetHostGroupNameOk() (*string, bool)`

GetHostGroupNameOk returns a tuple with the HostGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupName

`func (o *StoragePureHostLunAllOf) SetHostGroupName(v string)`

SetHostGroupName sets HostGroupName field to given value.

### HasHostGroupName

`func (o *StoragePureHostLunAllOf) HasHostGroupName() bool`

HasHostGroupName returns a boolean if a field has been set.

### GetShared

`func (o *StoragePureHostLunAllOf) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StoragePureHostLunAllOf) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StoragePureHostLunAllOf) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StoragePureHostLunAllOf) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureHostLunAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureHostLunAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureHostLunAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureHostLunAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHost

`func (o *StoragePureHostLunAllOf) GetHost() StoragePureHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StoragePureHostLunAllOf) GetHostOk() (*StoragePureHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StoragePureHostLunAllOf) SetHost(v StoragePureHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StoragePureHostLunAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostGroup

`func (o *StoragePureHostLunAllOf) GetHostGroup() StoragePureHostGroupRelationship`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *StoragePureHostLunAllOf) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *StoragePureHostLunAllOf) SetHostGroup(v StoragePureHostGroupRelationship)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *StoragePureHostLunAllOf) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureHostLunAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureHostLunAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureHostLunAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureHostLunAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVolume

`func (o *StoragePureHostLunAllOf) GetVolume() StoragePureVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StoragePureHostLunAllOf) GetVolumeOk() (*StoragePureVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StoragePureHostLunAllOf) SetVolume(v StoragePureVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StoragePureHostLunAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


