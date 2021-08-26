# StoragePureVolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureVolume"]
**Created** | Pointer to **time.Time** | Creation time of the volume. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number of the volume. | [optional] [readonly] 
**Source** | Pointer to **string** | Source from which the volume is created. Applicable only if the volume is cloned from other volume or snapshot. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](StoragePureProtectionGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureVolumeAllOf

`func NewStoragePureVolumeAllOf(classId string, objectType string, ) *StoragePureVolumeAllOf`

NewStoragePureVolumeAllOf instantiates a new StoragePureVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureVolumeAllOfWithDefaults

`func NewStoragePureVolumeAllOfWithDefaults() *StoragePureVolumeAllOf`

NewStoragePureVolumeAllOfWithDefaults instantiates a new StoragePureVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureVolumeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureVolumeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureVolumeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureVolumeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureVolumeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureVolumeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreated

`func (o *StoragePureVolumeAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoragePureVolumeAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoragePureVolumeAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoragePureVolumeAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSerial

`func (o *StoragePureVolumeAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StoragePureVolumeAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StoragePureVolumeAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StoragePureVolumeAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureVolumeAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureVolumeAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureVolumeAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureVolumeAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureVolumeAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureVolumeAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureVolumeAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureVolumeAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureVolumeAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureVolumeAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureVolumeAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureVolumeAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureVolumeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureVolumeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureVolumeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureVolumeAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


