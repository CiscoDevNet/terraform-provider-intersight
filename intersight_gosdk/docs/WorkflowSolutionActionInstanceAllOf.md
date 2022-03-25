# WorkflowSolutionActionInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionActionInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionActionInstance"]
**Action** | Pointer to **string** | Name of the action that needs to be performed on the solution instance. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Validate&#x60; - Validation the action instance inputs and run the validation workflows. * &#x60;Start&#x60; - Start a new execution of the action instance. * &#x60;Retry&#x60; - Retry the solution action instance from the beginning. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed from the failure point. * &#x60;Cancel&#x60; - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * &#x60;Stop&#x60; - Stop the action instance which is in progress and didn&#39;t complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. | [optional] [default to "None"]
**EndTime** | Pointer to **time.Time** | The time when the action was stopped or completed execution last time. | [optional] [readonly] 
**Input** | Pointer to **interface{}** | Inputs for a solution action and the format is specified by input definition of the solution action definition. | [optional] 
**LastAction** | Pointer to **string** | The last action that was issued on the action definition workflows is saved in this property. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Validate&#x60; - Validation the action instance inputs and run the validation workflows. * &#x60;Start&#x60; - Start a new execution of the action instance. * &#x60;Retry&#x60; - Retry the solution action instance from the beginning. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed from the failure point. * &#x60;Cancel&#x60; - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * &#x60;Stop&#x60; - Stop the action instance which is in progress and didn&#39;t complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | Name for the action instance is created in the system by appending name of the solution instance to the name of the action definition. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | The time when the action was started for execution last time. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the solution action instance. * &#x60;NotStarted&#x60; - Solution action is not yet started and it is in a draft mode. A solution action instance can be deleted in this state. * &#x60;Validating&#x60; - A validate action has been triggered on the action and until it completes the start action cannot be issued. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The action on the solution failed and can be retried. * &#x60;Completed&#x60; - The action on the solution completed successfully. * &#x60;Stopping&#x60; - The stop action is running on the action instance. | [optional] [readonly] [default to "NotStarted"]
**UpgradedMoid** | Pointer to **string** | Stores the upgraded Moid for help during future lookups. | [optional] [readonly] 
**ActionWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**SolutionActionDefinition** | Pointer to [**WorkflowSolutionActionDefinitionRelationship**](WorkflowSolutionActionDefinitionRelationship.md) |  | [optional] 
**SolutionDefinition** | Pointer to [**WorkflowSolutionDefinitionRelationship**](WorkflowSolutionDefinitionRelationship.md) |  | [optional] 
**SolutionInstance** | Pointer to [**WorkflowSolutionInstanceRelationship**](WorkflowSolutionInstanceRelationship.md) |  | [optional] 
**StopWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**ValidationWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionActionInstanceAllOf

`func NewWorkflowSolutionActionInstanceAllOf(classId string, objectType string, ) *WorkflowSolutionActionInstanceAllOf`

NewWorkflowSolutionActionInstanceAllOf instantiates a new WorkflowSolutionActionInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionActionInstanceAllOfWithDefaults

`func NewWorkflowSolutionActionInstanceAllOfWithDefaults() *WorkflowSolutionActionInstanceAllOf`

NewWorkflowSolutionActionInstanceAllOfWithDefaults instantiates a new WorkflowSolutionActionInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionActionInstanceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionActionInstanceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionActionInstanceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionActionInstanceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkflowSolutionActionInstanceAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowSolutionActionInstanceAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowSolutionActionInstanceAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowSolutionActionInstanceAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowSolutionActionInstanceAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowSolutionActionInstanceAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowSolutionActionInstanceAllOf) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowSolutionActionInstanceAllOf) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowSolutionActionInstanceAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowSolutionActionInstanceAllOf) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowSolutionActionInstanceAllOf) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLastAction

`func (o *WorkflowSolutionActionInstanceAllOf) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowSolutionActionInstanceAllOf) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowSolutionActionInstanceAllOf) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSolutionActionInstanceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionActionInstanceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionActionInstanceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowSolutionActionInstanceAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowSolutionActionInstanceAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowSolutionActionInstanceAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowSolutionActionInstanceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowSolutionActionInstanceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowSolutionActionInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradedMoid

`func (o *WorkflowSolutionActionInstanceAllOf) GetUpgradedMoid() string`

GetUpgradedMoid returns the UpgradedMoid field if non-nil, zero value otherwise.

### GetUpgradedMoidOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetUpgradedMoidOk() (*string, bool)`

GetUpgradedMoidOk returns a tuple with the UpgradedMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedMoid

`func (o *WorkflowSolutionActionInstanceAllOf) SetUpgradedMoid(v string)`

SetUpgradedMoid sets UpgradedMoid field to given value.

### HasUpgradedMoid

`func (o *WorkflowSolutionActionInstanceAllOf) HasUpgradedMoid() bool`

HasUpgradedMoid returns a boolean if a field has been set.

### GetActionWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) GetActionWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetActionWorkflowInfo returns the ActionWorkflowInfo field if non-nil, zero value otherwise.

### GetActionWorkflowInfoOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetActionWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetActionWorkflowInfoOk returns a tuple with the ActionWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) SetActionWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetActionWorkflowInfo sets ActionWorkflowInfo field to given value.

### HasActionWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) HasActionWorkflowInfo() bool`

HasActionWorkflowInfo returns a boolean if a field has been set.

### GetSolutionActionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionActionDefinition() WorkflowSolutionActionDefinitionRelationship`

GetSolutionActionDefinition returns the SolutionActionDefinition field if non-nil, zero value otherwise.

### GetSolutionActionDefinitionOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionActionDefinitionOk() (*WorkflowSolutionActionDefinitionRelationship, bool)`

GetSolutionActionDefinitionOk returns a tuple with the SolutionActionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionActionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) SetSolutionActionDefinition(v WorkflowSolutionActionDefinitionRelationship)`

SetSolutionActionDefinition sets SolutionActionDefinition field to given value.

### HasSolutionActionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) HasSolutionActionDefinition() bool`

HasSolutionActionDefinition returns a boolean if a field has been set.

### GetSolutionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship`

GetSolutionDefinition returns the SolutionDefinition field if non-nil, zero value otherwise.

### GetSolutionDefinitionOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool)`

GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship)`

SetSolutionDefinition sets SolutionDefinition field to given value.

### HasSolutionDefinition

`func (o *WorkflowSolutionActionInstanceAllOf) HasSolutionDefinition() bool`

HasSolutionDefinition returns a boolean if a field has been set.

### GetSolutionInstance

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionInstance() WorkflowSolutionInstanceRelationship`

GetSolutionInstance returns the SolutionInstance field if non-nil, zero value otherwise.

### GetSolutionInstanceOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool)`

GetSolutionInstanceOk returns a tuple with the SolutionInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionInstance

`func (o *WorkflowSolutionActionInstanceAllOf) SetSolutionInstance(v WorkflowSolutionInstanceRelationship)`

SetSolutionInstance sets SolutionInstance field to given value.

### HasSolutionInstance

`func (o *WorkflowSolutionActionInstanceAllOf) HasSolutionInstance() bool`

HasSolutionInstance returns a boolean if a field has been set.

### GetStopWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) GetStopWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetStopWorkflowInfo returns the StopWorkflowInfo field if non-nil, zero value otherwise.

### GetStopWorkflowInfoOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetStopWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetStopWorkflowInfoOk returns a tuple with the StopWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) SetStopWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetStopWorkflowInfo sets StopWorkflowInfo field to given value.

### HasStopWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) HasStopWorkflowInfo() bool`

HasStopWorkflowInfo returns a boolean if a field has been set.

### GetValidationWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) GetValidationWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetValidationWorkflowInfo returns the ValidationWorkflowInfo field if non-nil, zero value otherwise.

### GetValidationWorkflowInfoOk

`func (o *WorkflowSolutionActionInstanceAllOf) GetValidationWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetValidationWorkflowInfoOk returns a tuple with the ValidationWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) SetValidationWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetValidationWorkflowInfo sets ValidationWorkflowInfo field to given value.

### HasValidationWorkflowInfo

`func (o *WorkflowSolutionActionInstanceAllOf) HasValidationWorkflowInfo() bool`

HasValidationWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


