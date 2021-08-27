# StorageBaseArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**StorageUtilization** | Pointer to [**NullableStorageBaseCapacity**](StorageBaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseArray

`func NewStorageBaseArray(classId string, objectType string, ) *StorageBaseArray`

NewStorageBaseArray instantiates a new StorageBaseArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayWithDefaults

`func NewStorageBaseArrayWithDefaults() *StorageBaseArray`

NewStorageBaseArrayWithDefaults instantiates a new StorageBaseArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseArray) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseArray) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseArray) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseArray) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseArray) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseArray) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStorageUtilization

`func (o *StorageBaseArray) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseArray) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseArray) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseArray) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageBaseArray) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageBaseArray) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


