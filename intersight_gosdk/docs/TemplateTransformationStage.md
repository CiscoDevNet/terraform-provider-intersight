# TemplateTransformationStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "template.TransformationStage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "template.TransformationStage"]
**Function** | Pointer to **string** | The function to be executed. | [optional] 
**FunctionArguments** | Pointer to **interface{}** | A collection of arguments for the function being executed. | [optional] 
**Name** | Pointer to **string** | The unique name by which the output of this transformation stage can be accessed in further stages. Only alphanumeric characters are allowed. | [optional] 

## Methods

### NewTemplateTransformationStage

`func NewTemplateTransformationStage(classId string, objectType string, ) *TemplateTransformationStage`

NewTemplateTransformationStage instantiates a new TemplateTransformationStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTransformationStageWithDefaults

`func NewTemplateTransformationStageWithDefaults() *TemplateTransformationStage`

NewTemplateTransformationStageWithDefaults instantiates a new TemplateTransformationStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TemplateTransformationStage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TemplateTransformationStage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TemplateTransformationStage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TemplateTransformationStage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TemplateTransformationStage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TemplateTransformationStage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFunction

`func (o *TemplateTransformationStage) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *TemplateTransformationStage) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *TemplateTransformationStage) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *TemplateTransformationStage) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetFunctionArguments

`func (o *TemplateTransformationStage) GetFunctionArguments() interface{}`

GetFunctionArguments returns the FunctionArguments field if non-nil, zero value otherwise.

### GetFunctionArgumentsOk

`func (o *TemplateTransformationStage) GetFunctionArgumentsOk() (*interface{}, bool)`

GetFunctionArgumentsOk returns a tuple with the FunctionArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionArguments

`func (o *TemplateTransformationStage) SetFunctionArguments(v interface{})`

SetFunctionArguments sets FunctionArguments field to given value.

### HasFunctionArguments

`func (o *TemplateTransformationStage) HasFunctionArguments() bool`

HasFunctionArguments returns a boolean if a field has been set.

### SetFunctionArgumentsNil

`func (o *TemplateTransformationStage) SetFunctionArgumentsNil(b bool)`

 SetFunctionArgumentsNil sets the value for FunctionArguments to be an explicit nil

### UnsetFunctionArguments
`func (o *TemplateTransformationStage) UnsetFunctionArguments()`

UnsetFunctionArguments ensures that no value is present for FunctionArguments, not even an explicit nil
### GetName

`func (o *TemplateTransformationStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateTransformationStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateTransformationStage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateTransformationStage) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


