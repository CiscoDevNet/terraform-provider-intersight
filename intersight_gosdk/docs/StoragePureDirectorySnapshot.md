# StoragePureDirectorySnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureDirectorySnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureDirectorySnapshot"]
**ClientName** | Pointer to **string** | The customizable portion of the client-visible snapshot name. A full snapshot name is constructed in the form of DIR.CLIENT_NAME.SUFFIX where DIR is the full managed directory name, CLIENT_NAME is the client name, and SUFFIX is the suffix. The client-visible snapshot name is CLIENT_NAME.SUFFIX. | [optional] [readonly] 
**CreatedTime** | Pointer to **time.Time** | The snapshot creation time, measured in milliseconds since the UNIX epoch. | [optional] [readonly] 
**Name** | Pointer to **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [readonly] 
**Source** | Pointer to **string** | The directory from which this snapshot was taken. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Directory** | Pointer to [**NullableStoragePureDirectoryRelationship**](StoragePureDirectoryRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureDirectorySnapshot

`func NewStoragePureDirectorySnapshot(classId string, objectType string, ) *StoragePureDirectorySnapshot`

NewStoragePureDirectorySnapshot instantiates a new StoragePureDirectorySnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureDirectorySnapshotWithDefaults

`func NewStoragePureDirectorySnapshotWithDefaults() *StoragePureDirectorySnapshot`

NewStoragePureDirectorySnapshotWithDefaults instantiates a new StoragePureDirectorySnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureDirectorySnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureDirectorySnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureDirectorySnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureDirectorySnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureDirectorySnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureDirectorySnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClientName

`func (o *StoragePureDirectorySnapshot) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *StoragePureDirectorySnapshot) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *StoragePureDirectorySnapshot) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *StoragePureDirectorySnapshot) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *StoragePureDirectorySnapshot) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StoragePureDirectorySnapshot) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StoragePureDirectorySnapshot) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StoragePureDirectorySnapshot) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *StoragePureDirectorySnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureDirectorySnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureDirectorySnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureDirectorySnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureDirectorySnapshot) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureDirectorySnapshot) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureDirectorySnapshot) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureDirectorySnapshot) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureDirectorySnapshot) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureDirectorySnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureDirectorySnapshot) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureDirectorySnapshot) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureDirectorySnapshot) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureDirectorySnapshot) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetDirectory

`func (o *StoragePureDirectorySnapshot) GetDirectory() StoragePureDirectoryRelationship`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *StoragePureDirectorySnapshot) GetDirectoryOk() (*StoragePureDirectoryRelationship, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *StoragePureDirectorySnapshot) SetDirectory(v StoragePureDirectoryRelationship)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *StoragePureDirectorySnapshot) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### SetDirectoryNil

`func (o *StoragePureDirectorySnapshot) SetDirectoryNil(b bool)`

 SetDirectoryNil sets the value for Directory to be an explicit nil

### UnsetDirectory
`func (o *StoragePureDirectorySnapshot) UnsetDirectory()`

UnsetDirectory ensures that no value is present for Directory, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureDirectorySnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureDirectorySnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureDirectorySnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureDirectorySnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureDirectorySnapshot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureDirectorySnapshot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


