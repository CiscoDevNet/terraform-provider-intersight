# StoragePureProtectionGroupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureProtectionGroupSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureProtectionGroupSnapshot"]
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureProtectionGroupSnapshot

`func NewStoragePureProtectionGroupSnapshot(classId string, objectType string, ) *StoragePureProtectionGroupSnapshot`

NewStoragePureProtectionGroupSnapshot instantiates a new StoragePureProtectionGroupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupSnapshotWithDefaults

`func NewStoragePureProtectionGroupSnapshotWithDefaults() *StoragePureProtectionGroupSnapshot`

NewStoragePureProtectionGroupSnapshotWithDefaults instantiates a new StoragePureProtectionGroupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureProtectionGroupSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureProtectionGroupSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureProtectionGroupSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureProtectionGroupSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureProtectionGroupSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureProtectionGroupSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArray

`func (o *StoragePureProtectionGroupSnapshot) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureProtectionGroupSnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureProtectionGroupSnapshot) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureProtectionGroupSnapshot) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureProtectionGroupSnapshot) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureProtectionGroupSnapshot) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureProtectionGroupSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureProtectionGroupSnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


