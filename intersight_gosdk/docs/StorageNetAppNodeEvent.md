# StorageNetAppNodeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNodeEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNodeEvent"]
**ArrayController** | Pointer to [**NullableStorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNodeEvent

`func NewStorageNetAppNodeEvent(classId string, objectType string, ) *StorageNetAppNodeEvent`

NewStorageNetAppNodeEvent instantiates a new StorageNetAppNodeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeEventWithDefaults

`func NewStorageNetAppNodeEventWithDefaults() *StorageNetAppNodeEvent`

NewStorageNetAppNodeEventWithDefaults instantiates a new StorageNetAppNodeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNodeEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNodeEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNodeEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNodeEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNodeEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNodeEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayController

`func (o *StorageNetAppNodeEvent) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNodeEvent) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNodeEvent) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNodeEvent) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### SetArrayControllerNil

`func (o *StorageNetAppNodeEvent) SetArrayControllerNil(b bool)`

 SetArrayControllerNil sets the value for ArrayController to be an explicit nil

### UnsetArrayController
`func (o *StorageNetAppNodeEvent) UnsetArrayController()`

UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


