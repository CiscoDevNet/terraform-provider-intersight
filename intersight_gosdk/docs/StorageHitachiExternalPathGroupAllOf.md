# StorageHitachiExternalPathGroupAllOf

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

### NewStorageHitachiExternalPathGroupAllOf

`func NewStorageHitachiExternalPathGroupAllOf(classId string, objectType string, ) *StorageHitachiExternalPathGroupAllOf`

NewStorageHitachiExternalPathGroupAllOf instantiates a new StorageHitachiExternalPathGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalPathGroupAllOfWithDefaults

`func NewStorageHitachiExternalPathGroupAllOfWithDefaults() *StorageHitachiExternalPathGroupAllOf`

NewStorageHitachiExternalPathGroupAllOfWithDefaults instantiates a new StorageHitachiExternalPathGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalPathGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalPathGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalPathGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalPathGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalParityGroups

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalParityGroups() []StorageExternalParityGroup`

GetExternalParityGroups returns the ExternalParityGroups field if non-nil, zero value otherwise.

### GetExternalParityGroupsOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalParityGroupsOk() (*[]StorageExternalParityGroup, bool)`

GetExternalParityGroupsOk returns a tuple with the ExternalParityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroups

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalParityGroups(v []StorageExternalParityGroup)`

SetExternalParityGroups sets ExternalParityGroups field to given value.

### HasExternalParityGroups

`func (o *StorageHitachiExternalPathGroupAllOf) HasExternalParityGroups() bool`

HasExternalParityGroups returns a boolean if a field has been set.

### SetExternalParityGroupsNil

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalParityGroupsNil(b bool)`

 SetExternalParityGroupsNil sets the value for ExternalParityGroups to be an explicit nil

### UnsetExternalParityGroups
`func (o *StorageHitachiExternalPathGroupAllOf) UnsetExternalParityGroups()`

UnsetExternalParityGroups ensures that no value is present for ExternalParityGroups, not even an explicit nil
### GetExternalPaths

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalPaths() []StorageExternalPath`

GetExternalPaths returns the ExternalPaths field if non-nil, zero value otherwise.

### GetExternalPathsOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalPathsOk() (*[]StorageExternalPath, bool)`

GetExternalPathsOk returns a tuple with the ExternalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaths

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalPaths(v []StorageExternalPath)`

SetExternalPaths sets ExternalPaths field to given value.

### HasExternalPaths

`func (o *StorageHitachiExternalPathGroupAllOf) HasExternalPaths() bool`

HasExternalPaths returns a boolean if a field has been set.

### SetExternalPathsNil

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalPathsNil(b bool)`

 SetExternalPathsNil sets the value for ExternalPaths to be an explicit nil

### UnsetExternalPaths
`func (o *StorageHitachiExternalPathGroupAllOf) UnsetExternalPaths()`

UnsetExternalPaths ensures that no value is present for ExternalPaths, not even an explicit nil
### GetExternalProductId

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalProductId() string`

GetExternalProductId returns the ExternalProductId field if non-nil, zero value otherwise.

### GetExternalProductIdOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalProductIdOk() (*string, bool)`

GetExternalProductIdOk returns a tuple with the ExternalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductId

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalProductId(v string)`

SetExternalProductId sets ExternalProductId field to given value.

### HasExternalProductId

`func (o *StorageHitachiExternalPathGroupAllOf) HasExternalProductId() bool`

HasExternalProductId returns a boolean if a field has been set.

### GetExternalSerialNumber

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalSerialNumber() string`

GetExternalSerialNumber returns the ExternalSerialNumber field if non-nil, zero value otherwise.

### GetExternalSerialNumberOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetExternalSerialNumberOk() (*string, bool)`

GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSerialNumber

`func (o *StorageHitachiExternalPathGroupAllOf) SetExternalSerialNumber(v string)`

SetExternalSerialNumber sets ExternalSerialNumber field to given value.

### HasExternalSerialNumber

`func (o *StorageHitachiExternalPathGroupAllOf) HasExternalSerialNumber() bool`

HasExternalSerialNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiExternalPathGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiExternalPathGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiExternalPathGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalPathGroupAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalPathGroupAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalPathGroupAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiExternalPathGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalPathGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalPathGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalPathGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


