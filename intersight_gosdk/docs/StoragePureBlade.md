# StoragePureBlade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureBlade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureBlade"]
**BladeId** | Pointer to **string** | Unique identifier of the blade. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the blade within the FlashBlade array. | [optional] [readonly] 
**RawCapacity** | Pointer to **int64** | Raw storage capacity of the blade in bytes. | [optional] [readonly] 
**Status** | Pointer to **string** | Operational status of the blade. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureFlashBladeSystemRelationship**](StoragePureFlashBladeSystemRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureBlade

`func NewStoragePureBlade(classId string, objectType string, ) *StoragePureBlade`

NewStoragePureBlade instantiates a new StoragePureBlade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureBladeWithDefaults

`func NewStoragePureBladeWithDefaults() *StoragePureBlade`

NewStoragePureBladeWithDefaults instantiates a new StoragePureBlade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureBlade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureBlade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureBlade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureBlade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureBlade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureBlade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBladeId

`func (o *StoragePureBlade) GetBladeId() string`

GetBladeId returns the BladeId field if non-nil, zero value otherwise.

### GetBladeIdOk

`func (o *StoragePureBlade) GetBladeIdOk() (*string, bool)`

GetBladeIdOk returns a tuple with the BladeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBladeId

`func (o *StoragePureBlade) SetBladeId(v string)`

SetBladeId sets BladeId field to given value.

### HasBladeId

`func (o *StoragePureBlade) HasBladeId() bool`

HasBladeId returns a boolean if a field has been set.

### GetName

`func (o *StoragePureBlade) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureBlade) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureBlade) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureBlade) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawCapacity

`func (o *StoragePureBlade) GetRawCapacity() int64`

GetRawCapacity returns the RawCapacity field if non-nil, zero value otherwise.

### GetRawCapacityOk

`func (o *StoragePureBlade) GetRawCapacityOk() (*int64, bool)`

GetRawCapacityOk returns a tuple with the RawCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCapacity

`func (o *StoragePureBlade) SetRawCapacity(v int64)`

SetRawCapacity sets RawCapacity field to given value.

### HasRawCapacity

`func (o *StoragePureBlade) HasRawCapacity() bool`

HasRawCapacity returns a boolean if a field has been set.

### GetStatus

`func (o *StoragePureBlade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StoragePureBlade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StoragePureBlade) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StoragePureBlade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureBlade) GetArray() StoragePureFlashBladeSystemRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureBlade) GetArrayOk() (*StoragePureFlashBladeSystemRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureBlade) SetArray(v StoragePureFlashBladeSystemRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureBlade) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureBlade) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureBlade) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureBlade) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureBlade) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureBlade) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureBlade) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureBlade) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureBlade) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


