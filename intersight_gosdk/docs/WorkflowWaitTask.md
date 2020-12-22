# WorkflowWaitTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WaitTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WaitTask"]
**Prompts** | Pointer to [**[]WorkflowWaitTaskPrompt**](WorkflowWaitTaskPrompt.md) |  | [optional] 

## Methods

### NewWorkflowWaitTask

`func NewWorkflowWaitTask(classId string, objectType string, ) *WorkflowWaitTask`

NewWorkflowWaitTask instantiates a new WorkflowWaitTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWaitTaskWithDefaults

`func NewWorkflowWaitTaskWithDefaults() *WorkflowWaitTask`

NewWorkflowWaitTaskWithDefaults instantiates a new WorkflowWaitTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWaitTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWaitTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWaitTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWaitTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWaitTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWaitTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPrompts

`func (o *WorkflowWaitTask) GetPrompts() []WorkflowWaitTaskPrompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *WorkflowWaitTask) GetPromptsOk() (*[]WorkflowWaitTaskPrompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *WorkflowWaitTask) SetPrompts(v []WorkflowWaitTaskPrompt)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *WorkflowWaitTask) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### SetPromptsNil

`func (o *WorkflowWaitTask) SetPromptsNil(b bool)`

 SetPromptsNil sets the value for Prompts to be an explicit nil

### UnsetPrompts
`func (o *WorkflowWaitTask) UnsetPrompts()`

UnsetPrompts ensures that no value is present for Prompts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


