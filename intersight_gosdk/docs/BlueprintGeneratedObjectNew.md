# BlueprintGeneratedObjectNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectNew"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectNew"]
**Collection** | Pointer to **bool** | The flag to indicate that a collection of objects will be created based on an array input. | [optional] [default to false]
**CreateObjectType** | Pointer to **string** | The type of the object being created new. | [optional] 

## Methods

### NewBlueprintGeneratedObjectNew

`func NewBlueprintGeneratedObjectNew(classId string, objectType string, ) *BlueprintGeneratedObjectNew`

NewBlueprintGeneratedObjectNew instantiates a new BlueprintGeneratedObjectNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectNewWithDefaults

`func NewBlueprintGeneratedObjectNewWithDefaults() *BlueprintGeneratedObjectNew`

NewBlueprintGeneratedObjectNewWithDefaults instantiates a new BlueprintGeneratedObjectNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectNew) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectNew) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectNew) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectNew) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectNew) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectNew) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollection

`func (o *BlueprintGeneratedObjectNew) GetCollection() bool`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BlueprintGeneratedObjectNew) GetCollectionOk() (*bool, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BlueprintGeneratedObjectNew) SetCollection(v bool)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *BlueprintGeneratedObjectNew) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCreateObjectType

`func (o *BlueprintGeneratedObjectNew) GetCreateObjectType() string`

GetCreateObjectType returns the CreateObjectType field if non-nil, zero value otherwise.

### GetCreateObjectTypeOk

`func (o *BlueprintGeneratedObjectNew) GetCreateObjectTypeOk() (*string, bool)`

GetCreateObjectTypeOk returns a tuple with the CreateObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObjectType

`func (o *BlueprintGeneratedObjectNew) SetCreateObjectType(v string)`

SetCreateObjectType sets CreateObjectType field to given value.

### HasCreateObjectType

`func (o *BlueprintGeneratedObjectNew) HasCreateObjectType() bool`

HasCreateObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


