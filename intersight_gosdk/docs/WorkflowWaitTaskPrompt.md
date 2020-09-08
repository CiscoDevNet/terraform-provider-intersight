# WorkflowWaitTaskPrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description that give more details about what it means to pick this prompt option for the wait task. | [optional] 
**Label** | Pointer to **string** | User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task. | [optional] 
**Name** | Pointer to **string** | Name for the wait prompt. | [optional] 
**TaskStatus** | Pointer to **string** | Task status for the wait task when this prompt option is selected. * &#x60;Scheduled&#x60; - The enum represents the status when task is in scheduled state. * &#x60;InProgress&#x60; - The enum represents the status when task is in-progress state. * &#x60;NoOp&#x60; - The enum represents the status when task is a noop. * &#x60;Timeout&#x60; - The enum represents the status when task has timed out. * &#x60;Completed&#x60; - The enum represents the status when task has completed. * &#x60;Failed&#x60; - The enum represents the status when task has failed. | [optional] [default to "Scheduled"]

## Methods

### NewWorkflowWaitTaskPrompt

`func NewWorkflowWaitTaskPrompt() *WorkflowWaitTaskPrompt`

NewWorkflowWaitTaskPrompt instantiates a new WorkflowWaitTaskPrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWaitTaskPromptWithDefaults

`func NewWorkflowWaitTaskPromptWithDefaults() *WorkflowWaitTaskPrompt`

NewWorkflowWaitTaskPromptWithDefaults instantiates a new WorkflowWaitTaskPrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowWaitTaskPrompt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWaitTaskPrompt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWaitTaskPrompt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWaitTaskPrompt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowWaitTaskPrompt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowWaitTaskPrompt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowWaitTaskPrompt) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowWaitTaskPrompt) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWaitTaskPrompt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWaitTaskPrompt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWaitTaskPrompt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWaitTaskPrompt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaskStatus

`func (o *WorkflowWaitTaskPrompt) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *WorkflowWaitTaskPrompt) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *WorkflowWaitTaskPrompt) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *WorkflowWaitTaskPrompt) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


