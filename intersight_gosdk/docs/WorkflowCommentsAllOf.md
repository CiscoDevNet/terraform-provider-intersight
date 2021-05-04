# WorkflowCommentsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.Comments"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.Comments"]
**Description** | Pointer to **string** | Description field provides comment about the template function. | [optional] 
**Examples** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowCommentsAllOf

`func NewWorkflowCommentsAllOf(classId string, objectType string, ) *WorkflowCommentsAllOf`

NewWorkflowCommentsAllOf instantiates a new WorkflowCommentsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCommentsAllOfWithDefaults

`func NewWorkflowCommentsAllOfWithDefaults() *WorkflowCommentsAllOf`

NewWorkflowCommentsAllOfWithDefaults instantiates a new WorkflowCommentsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCommentsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCommentsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCommentsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCommentsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCommentsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCommentsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowCommentsAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCommentsAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCommentsAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCommentsAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExamples

`func (o *WorkflowCommentsAllOf) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *WorkflowCommentsAllOf) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *WorkflowCommentsAllOf) SetExamples(v []string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *WorkflowCommentsAllOf) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### SetExamplesNil

`func (o *WorkflowCommentsAllOf) SetExamplesNil(b bool)`

 SetExamplesNil sets the value for Examples to be an explicit nil

### UnsetExamples
`func (o *WorkflowCommentsAllOf) UnsetExamples()`

UnsetExamples ensures that no value is present for Examples, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


