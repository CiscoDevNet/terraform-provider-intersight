# StorageNetAppAggregateEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppAggregateEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppAggregateEvent"]
**Aggregate** | Pointer to [**StorageNetAppAggregateRelationship**](StorageNetAppAggregateRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppAggregateEventAllOf

`func NewStorageNetAppAggregateEventAllOf(classId string, objectType string, ) *StorageNetAppAggregateEventAllOf`

NewStorageNetAppAggregateEventAllOf instantiates a new StorageNetAppAggregateEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppAggregateEventAllOfWithDefaults

`func NewStorageNetAppAggregateEventAllOfWithDefaults() *StorageNetAppAggregateEventAllOf`

NewStorageNetAppAggregateEventAllOfWithDefaults instantiates a new StorageNetAppAggregateEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppAggregateEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppAggregateEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppAggregateEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppAggregateEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppAggregateEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppAggregateEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregate

`func (o *StorageNetAppAggregateEventAllOf) GetAggregate() StorageNetAppAggregateRelationship`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *StorageNetAppAggregateEventAllOf) GetAggregateOk() (*StorageNetAppAggregateRelationship, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *StorageNetAppAggregateEventAllOf) SetAggregate(v StorageNetAppAggregateRelationship)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *StorageNetAppAggregateEventAllOf) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


