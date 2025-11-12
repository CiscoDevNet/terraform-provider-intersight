# BlueprintGeneratedObjectSourceClone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectSourceClone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectSourceClone"]
**CloneType** | Pointer to **string** | The type of the source object. * &#x60;Clone&#x60; - The generated object will cloned from the given source object. * &#x60;DeepClone&#x60; - The generated object will deep cloned from the given source object. This means the object referenced in the source object will also be cloned. | [optional] [default to "Clone"]

## Methods

### NewBlueprintGeneratedObjectSourceClone

`func NewBlueprintGeneratedObjectSourceClone(classId string, objectType string, ) *BlueprintGeneratedObjectSourceClone`

NewBlueprintGeneratedObjectSourceClone instantiates a new BlueprintGeneratedObjectSourceClone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectSourceCloneWithDefaults

`func NewBlueprintGeneratedObjectSourceCloneWithDefaults() *BlueprintGeneratedObjectSourceClone`

NewBlueprintGeneratedObjectSourceCloneWithDefaults instantiates a new BlueprintGeneratedObjectSourceClone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectSourceClone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectSourceClone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectSourceClone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectSourceClone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectSourceClone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectSourceClone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloneType

`func (o *BlueprintGeneratedObjectSourceClone) GetCloneType() string`

GetCloneType returns the CloneType field if non-nil, zero value otherwise.

### GetCloneTypeOk

`func (o *BlueprintGeneratedObjectSourceClone) GetCloneTypeOk() (*string, bool)`

GetCloneTypeOk returns a tuple with the CloneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneType

`func (o *BlueprintGeneratedObjectSourceClone) SetCloneType(v string)`

SetCloneType sets CloneType field to given value.

### HasCloneType

`func (o *BlueprintGeneratedObjectSourceClone) HasCloneType() bool`

HasCloneType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


