# SchedulerTaskSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.TaskSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.TaskSchedule"]
**Action** | Pointer to **string** | The action of the scheduled task such as suspend or resume. * &#x60;None&#x60; - No action is set (default). * &#x60;Suspend&#x60; - Suspend a scheduled task indefinitely. * &#x60;Resume&#x60; - Resume a suspended scheduled task. * &#x60;SuspendTill&#x60; - Suspend the scheduled task until a specified end-date. | [optional] [default to "None"]
**Description** | Pointer to **string** | A description to describe the schedule for easier identification. | [optional] 
**LastAction** | Pointer to **string** | The last action for the scheduled task is saved in this field. Set to none if there was no action. * &#x60;None&#x60; - No action is set (default). * &#x60;Suspend&#x60; - Suspend a scheduled task indefinitely. * &#x60;Resume&#x60; - Resume a suspended scheduled task. * &#x60;SuspendTill&#x60; - Suspend the scheduled task until a specified end-date. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | A schedule name for easier identification (not required to be unique). | [optional] 
**ScheduleParams** | Pointer to [**SchedulerBaseScheduleParams**](SchedulerBaseScheduleParams.md) |  | [optional] 
**Status** | Pointer to [**NullableSchedulerTaskScheduleStatus**](SchedulerTaskScheduleStatus.md) |  | [optional] 
**SuspendEndTime** | Pointer to **time.Time** | Suspend a task until an end date. this applies only to the action suspendTill. | [optional] 
**TaskRequest** | Pointer to [**SchedulerRestStimTaskRequest**](SchedulerRestStimTaskRequest.md) |  | [optional] 
**Type** | Pointer to **string** | An Enum describing the type of scheduler to use. * &#x60;None&#x60; - No value was set for the schedule type (Enum value None). * &#x60;OneTime&#x60; - Define a one-time task execution time that will not automatically repeat. * &#x60;Recurring&#x60; - Specify a recurring task cadence based on a predefined pattern, such as daily, weekly, monthly, yearly, or every &lt;interval&gt; pattern. This option is not currently supported. | [optional] [default to "None"]
**AssociatedObject** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**WorkflowDefinition** | Pointer to [**NullableWorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewSchedulerTaskSchedule

`func NewSchedulerTaskSchedule(classId string, objectType string, ) *SchedulerTaskSchedule`

NewSchedulerTaskSchedule instantiates a new SchedulerTaskSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskScheduleWithDefaults

`func NewSchedulerTaskScheduleWithDefaults() *SchedulerTaskSchedule`

NewSchedulerTaskScheduleWithDefaults instantiates a new SchedulerTaskSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *SchedulerTaskSchedule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SchedulerTaskSchedule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SchedulerTaskSchedule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SchedulerTaskSchedule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *SchedulerTaskSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchedulerTaskSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchedulerTaskSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchedulerTaskSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastAction

`func (o *SchedulerTaskSchedule) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *SchedulerTaskSchedule) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *SchedulerTaskSchedule) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *SchedulerTaskSchedule) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetName

`func (o *SchedulerTaskSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulerTaskSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleParams

`func (o *SchedulerTaskSchedule) GetScheduleParams() SchedulerBaseScheduleParams`

GetScheduleParams returns the ScheduleParams field if non-nil, zero value otherwise.

### GetScheduleParamsOk

`func (o *SchedulerTaskSchedule) GetScheduleParamsOk() (*SchedulerBaseScheduleParams, bool)`

GetScheduleParamsOk returns a tuple with the ScheduleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleParams

`func (o *SchedulerTaskSchedule) SetScheduleParams(v SchedulerBaseScheduleParams)`

SetScheduleParams sets ScheduleParams field to given value.

### HasScheduleParams

`func (o *SchedulerTaskSchedule) HasScheduleParams() bool`

HasScheduleParams returns a boolean if a field has been set.

### GetStatus

`func (o *SchedulerTaskSchedule) GetStatus() SchedulerTaskScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskSchedule) GetStatusOk() (*SchedulerTaskScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskSchedule) SetStatus(v SchedulerTaskScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskSchedule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskSchedule) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskSchedule) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSuspendEndTime

`func (o *SchedulerTaskSchedule) GetSuspendEndTime() time.Time`

GetSuspendEndTime returns the SuspendEndTime field if non-nil, zero value otherwise.

### GetSuspendEndTimeOk

`func (o *SchedulerTaskSchedule) GetSuspendEndTimeOk() (*time.Time, bool)`

GetSuspendEndTimeOk returns a tuple with the SuspendEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendEndTime

`func (o *SchedulerTaskSchedule) SetSuspendEndTime(v time.Time)`

SetSuspendEndTime sets SuspendEndTime field to given value.

### HasSuspendEndTime

`func (o *SchedulerTaskSchedule) HasSuspendEndTime() bool`

HasSuspendEndTime returns a boolean if a field has been set.

### GetTaskRequest

`func (o *SchedulerTaskSchedule) GetTaskRequest() SchedulerRestStimTaskRequest`

GetTaskRequest returns the TaskRequest field if non-nil, zero value otherwise.

### GetTaskRequestOk

`func (o *SchedulerTaskSchedule) GetTaskRequestOk() (*SchedulerRestStimTaskRequest, bool)`

GetTaskRequestOk returns a tuple with the TaskRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRequest

`func (o *SchedulerTaskSchedule) SetTaskRequest(v SchedulerRestStimTaskRequest)`

SetTaskRequest sets TaskRequest field to given value.

### HasTaskRequest

`func (o *SchedulerTaskSchedule) HasTaskRequest() bool`

HasTaskRequest returns a boolean if a field has been set.

### GetType

`func (o *SchedulerTaskSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerTaskSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerTaskSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchedulerTaskSchedule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *SchedulerTaskSchedule) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *SchedulerTaskSchedule) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *SchedulerTaskSchedule) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *SchedulerTaskSchedule) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### SetAssociatedObjectNil

`func (o *SchedulerTaskSchedule) SetAssociatedObjectNil(b bool)`

 SetAssociatedObjectNil sets the value for AssociatedObject to be an explicit nil

### UnsetAssociatedObject
`func (o *SchedulerTaskSchedule) UnsetAssociatedObject()`

UnsetAssociatedObject ensures that no value is present for AssociatedObject, not even an explicit nil
### GetWorkflowDefinition

`func (o *SchedulerTaskSchedule) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *SchedulerTaskSchedule) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *SchedulerTaskSchedule) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *SchedulerTaskSchedule) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.

### SetWorkflowDefinitionNil

`func (o *SchedulerTaskSchedule) SetWorkflowDefinitionNil(b bool)`

 SetWorkflowDefinitionNil sets the value for WorkflowDefinition to be an explicit nil

### UnsetWorkflowDefinition
`func (o *SchedulerTaskSchedule) UnsetWorkflowDefinition()`

UnsetWorkflowDefinition ensures that no value is present for WorkflowDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


