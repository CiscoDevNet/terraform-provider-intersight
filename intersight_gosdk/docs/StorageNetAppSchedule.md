# StorageNetAppSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSchedule"]
**Name** | Pointer to **string** | The name of the schedule. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the schedule (cron or interval). | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally unique identifier of the schedule. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppSchedule

`func NewStorageNetAppSchedule(classId string, objectType string, ) *StorageNetAppSchedule`

NewStorageNetAppSchedule instantiates a new StorageNetAppSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppScheduleWithDefaults

`func NewStorageNetAppScheduleWithDefaults() *StorageNetAppSchedule`

NewStorageNetAppScheduleWithDefaults instantiates a new StorageNetAppSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageNetAppSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppSchedule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppSchedule) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppSchedule) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppSchedule) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppSchedule) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppSchedule) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppSchedule) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppSchedule) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppSchedule) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


