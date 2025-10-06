# BlueprintGeneratedObjectOperationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectOperationTarget"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectOperationTarget"]
**Name** | Pointer to **string** | The name of the generated object metadata. The objects produced from this metadata are the target of the operation. | [optional] 
**PropertyParameters** | Pointer to **interface{}** | A list of key value pairs where key is the property path and value is the template to derive the value. The target of the pre-generation operation is updated with the resolved values of these parameters. This operation is only applicable if the operation is of type \&quot;Update\&quot;. | [optional] 

## Methods

### NewBlueprintGeneratedObjectOperationTarget

`func NewBlueprintGeneratedObjectOperationTarget(classId string, objectType string, ) *BlueprintGeneratedObjectOperationTarget`

NewBlueprintGeneratedObjectOperationTarget instantiates a new BlueprintGeneratedObjectOperationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectOperationTargetWithDefaults

`func NewBlueprintGeneratedObjectOperationTargetWithDefaults() *BlueprintGeneratedObjectOperationTarget`

NewBlueprintGeneratedObjectOperationTargetWithDefaults instantiates a new BlueprintGeneratedObjectOperationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectOperationTarget) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectOperationTarget) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectOperationTarget) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectOperationTarget) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectOperationTarget) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectOperationTarget) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BlueprintGeneratedObjectOperationTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintGeneratedObjectOperationTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintGeneratedObjectOperationTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintGeneratedObjectOperationTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyParameters

`func (o *BlueprintGeneratedObjectOperationTarget) GetPropertyParameters() interface{}`

GetPropertyParameters returns the PropertyParameters field if non-nil, zero value otherwise.

### GetPropertyParametersOk

`func (o *BlueprintGeneratedObjectOperationTarget) GetPropertyParametersOk() (*interface{}, bool)`

GetPropertyParametersOk returns a tuple with the PropertyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParameters

`func (o *BlueprintGeneratedObjectOperationTarget) SetPropertyParameters(v interface{})`

SetPropertyParameters sets PropertyParameters field to given value.

### HasPropertyParameters

`func (o *BlueprintGeneratedObjectOperationTarget) HasPropertyParameters() bool`

HasPropertyParameters returns a boolean if a field has been set.

### SetPropertyParametersNil

`func (o *BlueprintGeneratedObjectOperationTarget) SetPropertyParametersNil(b bool)`

 SetPropertyParametersNil sets the value for PropertyParameters to be an explicit nil

### UnsetPropertyParameters
`func (o *BlueprintGeneratedObjectOperationTarget) UnsetPropertyParameters()`

UnsetPropertyParameters ensures that no value is present for PropertyParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


