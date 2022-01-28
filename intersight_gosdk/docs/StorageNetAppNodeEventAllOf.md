# StorageNetAppNodeEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNodeEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNodeEvent"]
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNodeEventAllOf

`func NewStorageNetAppNodeEventAllOf(classId string, objectType string, ) *StorageNetAppNodeEventAllOf`

NewStorageNetAppNodeEventAllOf instantiates a new StorageNetAppNodeEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNodeEventAllOfWithDefaults

`func NewStorageNetAppNodeEventAllOfWithDefaults() *StorageNetAppNodeEventAllOf`

NewStorageNetAppNodeEventAllOfWithDefaults instantiates a new StorageNetAppNodeEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNodeEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNodeEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNodeEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNodeEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNodeEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNodeEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayController

`func (o *StorageNetAppNodeEventAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNodeEventAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNodeEventAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNodeEventAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


