# BlueprintGeneratedObjectSourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectSourceReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectSourceReference"]
**SourceParameter** | Pointer to **string** | When the source object is of type reference, then the source parameter will be the parameter that refers to the source object. | [optional] 

## Methods

### NewBlueprintGeneratedObjectSourceReference

`func NewBlueprintGeneratedObjectSourceReference(classId string, objectType string, ) *BlueprintGeneratedObjectSourceReference`

NewBlueprintGeneratedObjectSourceReference instantiates a new BlueprintGeneratedObjectSourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectSourceReferenceWithDefaults

`func NewBlueprintGeneratedObjectSourceReferenceWithDefaults() *BlueprintGeneratedObjectSourceReference`

NewBlueprintGeneratedObjectSourceReferenceWithDefaults instantiates a new BlueprintGeneratedObjectSourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectSourceReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectSourceReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectSourceReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectSourceReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectSourceReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectSourceReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceParameter

`func (o *BlueprintGeneratedObjectSourceReference) GetSourceParameter() string`

GetSourceParameter returns the SourceParameter field if non-nil, zero value otherwise.

### GetSourceParameterOk

`func (o *BlueprintGeneratedObjectSourceReference) GetSourceParameterOk() (*string, bool)`

GetSourceParameterOk returns a tuple with the SourceParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameter

`func (o *BlueprintGeneratedObjectSourceReference) SetSourceParameter(v string)`

SetSourceParameter sets SourceParameter field to given value.

### HasSourceParameter

`func (o *BlueprintGeneratedObjectSourceReference) HasSourceParameter() bool`

HasSourceParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


