# BlueprintBlueprintInheritanceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.BlueprintInheritanceReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.BlueprintInheritanceReference"]
**Name** | Pointer to **string** | The name of the referenced abstract blueprint. | [optional] 

## Methods

### NewBlueprintBlueprintInheritanceReference

`func NewBlueprintBlueprintInheritanceReference(classId string, objectType string, ) *BlueprintBlueprintInheritanceReference`

NewBlueprintBlueprintInheritanceReference instantiates a new BlueprintBlueprintInheritanceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBlueprintInheritanceReferenceWithDefaults

`func NewBlueprintBlueprintInheritanceReferenceWithDefaults() *BlueprintBlueprintInheritanceReference`

NewBlueprintBlueprintInheritanceReferenceWithDefaults instantiates a new BlueprintBlueprintInheritanceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintBlueprintInheritanceReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintBlueprintInheritanceReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintBlueprintInheritanceReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintBlueprintInheritanceReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintBlueprintInheritanceReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintBlueprintInheritanceReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BlueprintBlueprintInheritanceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintBlueprintInheritanceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintBlueprintInheritanceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintBlueprintInheritanceReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


