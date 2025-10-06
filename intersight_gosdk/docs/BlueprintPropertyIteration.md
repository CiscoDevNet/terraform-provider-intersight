# BlueprintPropertyIteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.PropertyIteration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.PropertyIteration"]
**Iteration** | Pointer to **string** | The expression to evaluate for each iteration. This expression can refer to the array input that is being iterated over. The value of this expression will be used to populate the array property in the generated object. | [optional] 
**Name** | Pointer to **string** | The name of the array property in the generated object to populate. | [optional] 

## Methods

### NewBlueprintPropertyIteration

`func NewBlueprintPropertyIteration(classId string, objectType string, ) *BlueprintPropertyIteration`

NewBlueprintPropertyIteration instantiates a new BlueprintPropertyIteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPropertyIterationWithDefaults

`func NewBlueprintPropertyIterationWithDefaults() *BlueprintPropertyIteration`

NewBlueprintPropertyIterationWithDefaults instantiates a new BlueprintPropertyIteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintPropertyIteration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintPropertyIteration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintPropertyIteration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintPropertyIteration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintPropertyIteration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintPropertyIteration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIteration

`func (o *BlueprintPropertyIteration) GetIteration() string`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *BlueprintPropertyIteration) GetIterationOk() (*string, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *BlueprintPropertyIteration) SetIteration(v string)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *BlueprintPropertyIteration) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetName

`func (o *BlueprintPropertyIteration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintPropertyIteration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintPropertyIteration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintPropertyIteration) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


