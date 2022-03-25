# WorkflowServiceItemActionInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemActionInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemActionInstance"]
**Action** | Pointer to **string** | Name of the action that needs to be performed on the service item instance. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Validate&#x60; - Validate the action instance inputs and run the validation workflows. * &#x60;Start&#x60; - Start a new execution of the action instance. * &#x60;Retry&#x60; - Retry the service item action instance from the beginning. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed from the failure point. * &#x60;Cancel&#x60; - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * &#x60;Stop&#x60; - Stop the action instance which is in progress and didn&#39;t complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. | [optional] [default to "None"]
**EndTime** | Pointer to **time.Time** | The time when the action was stopped or completed execution last time. | [optional] [readonly] 
**Input** | Pointer to **interface{}** | Inputs for a service item action and the format is specified by input definition of the service item action definition. | [optional] 
**LastAction** | Pointer to **string** | The last action that was issued on the action definition workflows is saved in this property. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Validate&#x60; - Validate the action instance inputs and run the validation workflows. * &#x60;Start&#x60; - Start a new execution of the action instance. * &#x60;Retry&#x60; - Retry the service item action instance from the beginning. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed from the failure point. * &#x60;Cancel&#x60; - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * &#x60;Stop&#x60; - Stop the action instance which is in progress and didn&#39;t complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | Name for the action instance is created in the system by appending name of the service item instance to the name of the action definition. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | The time when the action was started for execution last time. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the service item action instance. * &#x60;NotStarted&#x60; - An action on the service item is not yet started and it is in a draft mode. A service item action instance can be deleted in this state. * &#x60;Validating&#x60; - A validate action has been triggered on the action and until it completes the start action cannot be issued. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The action on the service item instance failed and can be retried. * &#x60;Completed&#x60; - The action on the service item instance completed successfully. * &#x60;Stopping&#x60; - The stop action is running on the action instance. | [optional] [readonly] [default to "NotStarted"]
**ActionWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**ServiceItemActionDefinition** | Pointer to [**WorkflowServiceItemActionDefinitionRelationship**](WorkflowServiceItemActionDefinitionRelationship.md) |  | [optional] 
**ServiceItemDefinition** | Pointer to [**WorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 
**StopWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**ValidationWorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemActionInstanceAllOf

`func NewWorkflowServiceItemActionInstanceAllOf(classId string, objectType string, ) *WorkflowServiceItemActionInstanceAllOf`

NewWorkflowServiceItemActionInstanceAllOf instantiates a new WorkflowServiceItemActionInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemActionInstanceAllOfWithDefaults

`func NewWorkflowServiceItemActionInstanceAllOfWithDefaults() *WorkflowServiceItemActionInstanceAllOf`

NewWorkflowServiceItemActionInstanceAllOfWithDefaults instantiates a new WorkflowServiceItemActionInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemActionInstanceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemActionInstanceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemActionInstanceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemActionInstanceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkflowServiceItemActionInstanceAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowServiceItemActionInstanceAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowServiceItemActionInstanceAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowServiceItemActionInstanceAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowServiceItemActionInstanceAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowServiceItemActionInstanceAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowServiceItemActionInstanceAllOf) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowServiceItemActionInstanceAllOf) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowServiceItemActionInstanceAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowServiceItemActionInstanceAllOf) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowServiceItemActionInstanceAllOf) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLastAction

`func (o *WorkflowServiceItemActionInstanceAllOf) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowServiceItemActionInstanceAllOf) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowServiceItemActionInstanceAllOf) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemActionInstanceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemActionInstanceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemActionInstanceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowServiceItemActionInstanceAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowServiceItemActionInstanceAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowServiceItemActionInstanceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowServiceItemActionInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActionWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) GetActionWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetActionWorkflowInfo returns the ActionWorkflowInfo field if non-nil, zero value otherwise.

### GetActionWorkflowInfoOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetActionWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetActionWorkflowInfoOk returns a tuple with the ActionWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) SetActionWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetActionWorkflowInfo sets ActionWorkflowInfo field to given value.

### HasActionWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) HasActionWorkflowInfo() bool`

HasActionWorkflowInfo returns a boolean if a field has been set.

### GetServiceItemActionDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemActionDefinition() WorkflowServiceItemActionDefinitionRelationship`

GetServiceItemActionDefinition returns the ServiceItemActionDefinition field if non-nil, zero value otherwise.

### GetServiceItemActionDefinitionOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemActionDefinitionOk() (*WorkflowServiceItemActionDefinitionRelationship, bool)`

GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemActionDefinition(v WorkflowServiceItemActionDefinitionRelationship)`

SetServiceItemActionDefinition sets ServiceItemActionDefinition field to given value.

### HasServiceItemActionDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemActionDefinition() bool`

HasServiceItemActionDefinition returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.

### GetStopWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStopWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetStopWorkflowInfo returns the StopWorkflowInfo field if non-nil, zero value otherwise.

### GetStopWorkflowInfoOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetStopWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetStopWorkflowInfoOk returns a tuple with the StopWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) SetStopWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetStopWorkflowInfo sets StopWorkflowInfo field to given value.

### HasStopWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) HasStopWorkflowInfo() bool`

HasStopWorkflowInfo returns a boolean if a field has been set.

### GetValidationWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) GetValidationWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetValidationWorkflowInfo returns the ValidationWorkflowInfo field if non-nil, zero value otherwise.

### GetValidationWorkflowInfoOk

`func (o *WorkflowServiceItemActionInstanceAllOf) GetValidationWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetValidationWorkflowInfoOk returns a tuple with the ValidationWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) SetValidationWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetValidationWorkflowInfo sets ValidationWorkflowInfo field to given value.

### HasValidationWorkflowInfo

`func (o *WorkflowServiceItemActionInstanceAllOf) HasValidationWorkflowInfo() bool`

HasValidationWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


