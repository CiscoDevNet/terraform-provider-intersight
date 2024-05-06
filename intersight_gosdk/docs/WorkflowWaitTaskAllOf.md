# WorkflowWaitTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WaitTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WaitTask"]
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Prompts** | Pointer to [**[]WorkflowWaitTaskPrompt**](WorkflowWaitTaskPrompt.md) |  | [optional] 

## Methods

### NewWorkflowWaitTaskAllOf

`func NewWorkflowWaitTaskAllOf(classId string, objectType string, ) *WorkflowWaitTaskAllOf`

NewWorkflowWaitTaskAllOf instantiates a new WorkflowWaitTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWaitTaskAllOfWithDefaults

`func NewWorkflowWaitTaskAllOfWithDefaults() *WorkflowWaitTaskAllOf`

NewWorkflowWaitTaskAllOfWithDefaults instantiates a new WorkflowWaitTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWaitTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWaitTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWaitTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWaitTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWaitTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWaitTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInputDefinition

`func (o *WorkflowWaitTaskAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowWaitTaskAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowWaitTaskAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowWaitTaskAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowWaitTaskAllOf) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowWaitTaskAllOf) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetOutputDefinition

`func (o *WorkflowWaitTaskAllOf) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowWaitTaskAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowWaitTaskAllOf) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowWaitTaskAllOf) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowWaitTaskAllOf) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowWaitTaskAllOf) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetPrompts

`func (o *WorkflowWaitTaskAllOf) GetPrompts() []WorkflowWaitTaskPrompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *WorkflowWaitTaskAllOf) GetPromptsOk() (*[]WorkflowWaitTaskPrompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *WorkflowWaitTaskAllOf) SetPrompts(v []WorkflowWaitTaskPrompt)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *WorkflowWaitTaskAllOf) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### SetPromptsNil

`func (o *WorkflowWaitTaskAllOf) SetPromptsNil(b bool)`

 SetPromptsNil sets the value for Prompts to be an explicit nil

### UnsetPrompts
`func (o *WorkflowWaitTaskAllOf) UnsetPrompts()`

UnsetPrompts ensures that no value is present for Prompts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


