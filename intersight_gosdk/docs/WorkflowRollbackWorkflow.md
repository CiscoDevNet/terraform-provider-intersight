# WorkflowRollbackWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.RollbackWorkflow"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.RollbackWorkflow"]
**Action** | Pointer to **string** | The action of the rollback workflow such as Create and Start. * &#x60;None&#x60; - If no action is set, then the default value is set to none for the action field. * &#x60;Create&#x60; - Create rollback workflow data for the execution of the rollback workflow. * &#x60;Start&#x60; - Start a new execution of the rollback workflow. | [optional] [default to "None"]
**ContinueOnTaskFailure** | Pointer to **bool** | When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails. | [optional] [default to true]
**RollbackTasks** | Pointer to [**[]WorkflowRollbackWorkflowTask**](WorkflowRollbackWorkflowTask.md) |  | [optional] 
**SelectedTasks** | Pointer to [**[]WorkflowRollbackWorkflowTask**](WorkflowRollbackWorkflowTask.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the rollback workflow instance (Created, Running, Completed, Failed). * &#x60;None&#x60; - If no status is set, then the default value is set none for the status field. * &#x60;Created&#x60; - Status of the rollback workflow when it identifies the eligible tasks for rollback. * &#x60;Running&#x60; - Status of the rollback workflow when it is in-progress. * &#x60;Completed&#x60; - Status of the rollback workflow after execution is successful. * &#x60;Failed&#x60; - Status of the rollback workflow after execution results in failure. | [optional] [readonly] [default to "None"]
**PrimaryWorkflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**RollbackWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 

## Methods

### NewWorkflowRollbackWorkflow

`func NewWorkflowRollbackWorkflow(classId string, objectType string, ) *WorkflowRollbackWorkflow`

NewWorkflowRollbackWorkflow instantiates a new WorkflowRollbackWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRollbackWorkflowWithDefaults

`func NewWorkflowRollbackWorkflowWithDefaults() *WorkflowRollbackWorkflow`

NewWorkflowRollbackWorkflowWithDefaults instantiates a new WorkflowRollbackWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowRollbackWorkflow) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowRollbackWorkflow) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowRollbackWorkflow) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowRollbackWorkflow) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowRollbackWorkflow) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowRollbackWorkflow) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkflowRollbackWorkflow) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowRollbackWorkflow) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowRollbackWorkflow) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowRollbackWorkflow) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetContinueOnTaskFailure

`func (o *WorkflowRollbackWorkflow) GetContinueOnTaskFailure() bool`

GetContinueOnTaskFailure returns the ContinueOnTaskFailure field if non-nil, zero value otherwise.

### GetContinueOnTaskFailureOk

`func (o *WorkflowRollbackWorkflow) GetContinueOnTaskFailureOk() (*bool, bool)`

GetContinueOnTaskFailureOk returns a tuple with the ContinueOnTaskFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnTaskFailure

`func (o *WorkflowRollbackWorkflow) SetContinueOnTaskFailure(v bool)`

SetContinueOnTaskFailure sets ContinueOnTaskFailure field to given value.

### HasContinueOnTaskFailure

`func (o *WorkflowRollbackWorkflow) HasContinueOnTaskFailure() bool`

HasContinueOnTaskFailure returns a boolean if a field has been set.

### GetRollbackTasks

`func (o *WorkflowRollbackWorkflow) GetRollbackTasks() []WorkflowRollbackWorkflowTask`

GetRollbackTasks returns the RollbackTasks field if non-nil, zero value otherwise.

### GetRollbackTasksOk

`func (o *WorkflowRollbackWorkflow) GetRollbackTasksOk() (*[]WorkflowRollbackWorkflowTask, bool)`

GetRollbackTasksOk returns a tuple with the RollbackTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackTasks

`func (o *WorkflowRollbackWorkflow) SetRollbackTasks(v []WorkflowRollbackWorkflowTask)`

SetRollbackTasks sets RollbackTasks field to given value.

### HasRollbackTasks

`func (o *WorkflowRollbackWorkflow) HasRollbackTasks() bool`

HasRollbackTasks returns a boolean if a field has been set.

### SetRollbackTasksNil

`func (o *WorkflowRollbackWorkflow) SetRollbackTasksNil(b bool)`

 SetRollbackTasksNil sets the value for RollbackTasks to be an explicit nil

### UnsetRollbackTasks
`func (o *WorkflowRollbackWorkflow) UnsetRollbackTasks()`

UnsetRollbackTasks ensures that no value is present for RollbackTasks, not even an explicit nil
### GetSelectedTasks

`func (o *WorkflowRollbackWorkflow) GetSelectedTasks() []WorkflowRollbackWorkflowTask`

GetSelectedTasks returns the SelectedTasks field if non-nil, zero value otherwise.

### GetSelectedTasksOk

`func (o *WorkflowRollbackWorkflow) GetSelectedTasksOk() (*[]WorkflowRollbackWorkflowTask, bool)`

GetSelectedTasksOk returns a tuple with the SelectedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTasks

`func (o *WorkflowRollbackWorkflow) SetSelectedTasks(v []WorkflowRollbackWorkflowTask)`

SetSelectedTasks sets SelectedTasks field to given value.

### HasSelectedTasks

`func (o *WorkflowRollbackWorkflow) HasSelectedTasks() bool`

HasSelectedTasks returns a boolean if a field has been set.

### SetSelectedTasksNil

`func (o *WorkflowRollbackWorkflow) SetSelectedTasksNil(b bool)`

 SetSelectedTasksNil sets the value for SelectedTasks to be an explicit nil

### UnsetSelectedTasks
`func (o *WorkflowRollbackWorkflow) UnsetSelectedTasks()`

UnsetSelectedTasks ensures that no value is present for SelectedTasks, not even an explicit nil
### GetStatus

`func (o *WorkflowRollbackWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRollbackWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRollbackWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRollbackWorkflow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimaryWorkflow

`func (o *WorkflowRollbackWorkflow) GetPrimaryWorkflow() WorkflowWorkflowInfoRelationship`

GetPrimaryWorkflow returns the PrimaryWorkflow field if non-nil, zero value otherwise.

### GetPrimaryWorkflowOk

`func (o *WorkflowRollbackWorkflow) GetPrimaryWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetPrimaryWorkflowOk returns a tuple with the PrimaryWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWorkflow

`func (o *WorkflowRollbackWorkflow) SetPrimaryWorkflow(v WorkflowWorkflowInfoRelationship)`

SetPrimaryWorkflow sets PrimaryWorkflow field to given value.

### HasPrimaryWorkflow

`func (o *WorkflowRollbackWorkflow) HasPrimaryWorkflow() bool`

HasPrimaryWorkflow returns a boolean if a field has been set.

### GetRollbackWorkflows

`func (o *WorkflowRollbackWorkflow) GetRollbackWorkflows() []WorkflowWorkflowInfoRelationship`

GetRollbackWorkflows returns the RollbackWorkflows field if non-nil, zero value otherwise.

### GetRollbackWorkflowsOk

`func (o *WorkflowRollbackWorkflow) GetRollbackWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRollbackWorkflowsOk returns a tuple with the RollbackWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackWorkflows

`func (o *WorkflowRollbackWorkflow) SetRollbackWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRollbackWorkflows sets RollbackWorkflows field to given value.

### HasRollbackWorkflows

`func (o *WorkflowRollbackWorkflow) HasRollbackWorkflows() bool`

HasRollbackWorkflows returns a boolean if a field has been set.

### SetRollbackWorkflowsNil

`func (o *WorkflowRollbackWorkflow) SetRollbackWorkflowsNil(b bool)`

 SetRollbackWorkflowsNil sets the value for RollbackWorkflows to be an explicit nil

### UnsetRollbackWorkflows
`func (o *WorkflowRollbackWorkflow) UnsetRollbackWorkflows()`

UnsetRollbackWorkflows ensures that no value is present for RollbackWorkflows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


