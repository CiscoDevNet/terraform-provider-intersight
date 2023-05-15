# StorageHitachiExternalPathGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiExternalPathGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiExternalPathGroup"]
**ExternalParityGroups** | Pointer to [**[]StorageExternalParityGroup**](StorageExternalParityGroup.md) |  | [optional] 
**ExternalPaths** | Pointer to [**[]StorageExternalPath**](StorageExternalPath.md) |  | [optional] 
**ExternalProductId** | Pointer to **string** | Product ID of the external storage system. | [optional] [readonly] 
**ExternalSerialNumber** | Pointer to **string** | Serial number of the external storage system. | [optional] [readonly] 
**Name** | Pointer to **string** | External path group number. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiExternalPathGroup

`func NewStorageHitachiExternalPathGroup(classId string, objectType string, ) *StorageHitachiExternalPathGroup`

NewStorageHitachiExternalPathGroup instantiates a new StorageHitachiExternalPathGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalPathGroupWithDefaults

`func NewStorageHitachiExternalPathGroupWithDefaults() *StorageHitachiExternalPathGroup`

NewStorageHitachiExternalPathGroupWithDefaults instantiates a new StorageHitachiExternalPathGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalPathGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalPathGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalPathGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalPathGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalPathGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalPathGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalParityGroups

`func (o *StorageHitachiExternalPathGroup) GetExternalParityGroups() []StorageExternalParityGroup`

GetExternalParityGroups returns the ExternalParityGroups field if non-nil, zero value otherwise.

### GetExternalParityGroupsOk

`func (o *StorageHitachiExternalPathGroup) GetExternalParityGroupsOk() (*[]StorageExternalParityGroup, bool)`

GetExternalParityGroupsOk returns a tuple with the ExternalParityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroups

`func (o *StorageHitachiExternalPathGroup) SetExternalParityGroups(v []StorageExternalParityGroup)`

SetExternalParityGroups sets ExternalParityGroups field to given value.

### HasExternalParityGroups

`func (o *StorageHitachiExternalPathGroup) HasExternalParityGroups() bool`

HasExternalParityGroups returns a boolean if a field has been set.

### SetExternalParityGroupsNil

`func (o *StorageHitachiExternalPathGroup) SetExternalParityGroupsNil(b bool)`

 SetExternalParityGroupsNil sets the value for ExternalParityGroups to be an explicit nil

### UnsetExternalParityGroups
`func (o *StorageHitachiExternalPathGroup) UnsetExternalParityGroups()`

UnsetExternalParityGroups ensures that no value is present for ExternalParityGroups, not even an explicit nil
### GetExternalPaths

`func (o *StorageHitachiExternalPathGroup) GetExternalPaths() []StorageExternalPath`

GetExternalPaths returns the ExternalPaths field if non-nil, zero value otherwise.

### GetExternalPathsOk

`func (o *StorageHitachiExternalPathGroup) GetExternalPathsOk() (*[]StorageExternalPath, bool)`

GetExternalPathsOk returns a tuple with the ExternalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaths

`func (o *StorageHitachiExternalPathGroup) SetExternalPaths(v []StorageExternalPath)`

SetExternalPaths sets ExternalPaths field to given value.

### HasExternalPaths

`func (o *StorageHitachiExternalPathGroup) HasExternalPaths() bool`

HasExternalPaths returns a boolean if a field has been set.

### SetExternalPathsNil

`func (o *StorageHitachiExternalPathGroup) SetExternalPathsNil(b bool)`

 SetExternalPathsNil sets the value for ExternalPaths to be an explicit nil

### UnsetExternalPaths
`func (o *StorageHitachiExternalPathGroup) UnsetExternalPaths()`

UnsetExternalPaths ensures that no value is present for ExternalPaths, not even an explicit nil
### GetExternalProductId

`func (o *StorageHitachiExternalPathGroup) GetExternalProductId() string`

GetExternalProductId returns the ExternalProductId field if non-nil, zero value otherwise.

### GetExternalProductIdOk

`func (o *StorageHitachiExternalPathGroup) GetExternalProductIdOk() (*string, bool)`

GetExternalProductIdOk returns a tuple with the ExternalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductId

`func (o *StorageHitachiExternalPathGroup) SetExternalProductId(v string)`

SetExternalProductId sets ExternalProductId field to given value.

### HasExternalProductId

`func (o *StorageHitachiExternalPathGroup) HasExternalProductId() bool`

HasExternalProductId returns a boolean if a field has been set.

### GetExternalSerialNumber

`func (o *StorageHitachiExternalPathGroup) GetExternalSerialNumber() string`

GetExternalSerialNumber returns the ExternalSerialNumber field if non-nil, zero value otherwise.

### GetExternalSerialNumberOk

`func (o *StorageHitachiExternalPathGroup) GetExternalSerialNumberOk() (*string, bool)`

GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSerialNumber

`func (o *StorageHitachiExternalPathGroup) SetExternalSerialNumber(v string)`

SetExternalSerialNumber sets ExternalSerialNumber field to given value.

### HasExternalSerialNumber

`func (o *StorageHitachiExternalPathGroup) HasExternalSerialNumber() bool`

HasExternalSerialNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiExternalPathGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiExternalPathGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiExternalPathGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiExternalPathGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalPathGroup) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalPathGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalPathGroup) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalPathGroup) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiExternalPathGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalPathGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalPathGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalPathGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


