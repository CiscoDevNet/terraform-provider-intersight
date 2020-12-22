# WorkflowWaitTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WaitTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WaitTask"]
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


