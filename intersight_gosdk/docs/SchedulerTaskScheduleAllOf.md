# SchedulerTaskScheduleAllOf

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
**Type** | Pointer to **string** | An Enum describing the type of scheduler to use. * &#x60;None&#x60; - No value was set for the schedule type (Enum value None). * &#x60;OneTime&#x60; - Define a one-time task execution time that will not automatically repeat. * &#x60;Recurring&#x60; - Specify a recurring task cadence based on a predefined pattern, such as daily, weekly, monthly, yearly, or every &lt;interval&gt; pattern. | [optional] [default to "None"]
**AssociatedObject** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**WorkflowDefinition** | Pointer to [**WorkflowWorkflowDefinitionRelationship**](WorkflowWorkflowDefinitionRelationship.md) |  | [optional] 

## Methods

### NewSchedulerTaskScheduleAllOf

`func NewSchedulerTaskScheduleAllOf(classId string, objectType string, ) *SchedulerTaskScheduleAllOf`

NewSchedulerTaskScheduleAllOf instantiates a new SchedulerTaskScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskScheduleAllOfWithDefaults

`func NewSchedulerTaskScheduleAllOfWithDefaults() *SchedulerTaskScheduleAllOf`

NewSchedulerTaskScheduleAllOfWithDefaults instantiates a new SchedulerTaskScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *SchedulerTaskScheduleAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SchedulerTaskScheduleAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SchedulerTaskScheduleAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SchedulerTaskScheduleAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *SchedulerTaskScheduleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchedulerTaskScheduleAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchedulerTaskScheduleAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchedulerTaskScheduleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastAction

`func (o *SchedulerTaskScheduleAllOf) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *SchedulerTaskScheduleAllOf) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *SchedulerTaskScheduleAllOf) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *SchedulerTaskScheduleAllOf) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetName

`func (o *SchedulerTaskScheduleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchedulerTaskScheduleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchedulerTaskScheduleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchedulerTaskScheduleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleParams

`func (o *SchedulerTaskScheduleAllOf) GetScheduleParams() SchedulerBaseScheduleParams`

GetScheduleParams returns the ScheduleParams field if non-nil, zero value otherwise.

### GetScheduleParamsOk

`func (o *SchedulerTaskScheduleAllOf) GetScheduleParamsOk() (*SchedulerBaseScheduleParams, bool)`

GetScheduleParamsOk returns a tuple with the ScheduleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleParams

`func (o *SchedulerTaskScheduleAllOf) SetScheduleParams(v SchedulerBaseScheduleParams)`

SetScheduleParams sets ScheduleParams field to given value.

### HasScheduleParams

`func (o *SchedulerTaskScheduleAllOf) HasScheduleParams() bool`

HasScheduleParams returns a boolean if a field has been set.

### GetStatus

`func (o *SchedulerTaskScheduleAllOf) GetStatus() SchedulerTaskScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskScheduleAllOf) GetStatusOk() (*SchedulerTaskScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskScheduleAllOf) SetStatus(v SchedulerTaskScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskScheduleAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SchedulerTaskScheduleAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SchedulerTaskScheduleAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSuspendEndTime

`func (o *SchedulerTaskScheduleAllOf) GetSuspendEndTime() time.Time`

GetSuspendEndTime returns the SuspendEndTime field if non-nil, zero value otherwise.

### GetSuspendEndTimeOk

`func (o *SchedulerTaskScheduleAllOf) GetSuspendEndTimeOk() (*time.Time, bool)`

GetSuspendEndTimeOk returns a tuple with the SuspendEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendEndTime

`func (o *SchedulerTaskScheduleAllOf) SetSuspendEndTime(v time.Time)`

SetSuspendEndTime sets SuspendEndTime field to given value.

### HasSuspendEndTime

`func (o *SchedulerTaskScheduleAllOf) HasSuspendEndTime() bool`

HasSuspendEndTime returns a boolean if a field has been set.

### GetTaskRequest

`func (o *SchedulerTaskScheduleAllOf) GetTaskRequest() SchedulerRestStimTaskRequest`

GetTaskRequest returns the TaskRequest field if non-nil, zero value otherwise.

### GetTaskRequestOk

`func (o *SchedulerTaskScheduleAllOf) GetTaskRequestOk() (*SchedulerRestStimTaskRequest, bool)`

GetTaskRequestOk returns a tuple with the TaskRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRequest

`func (o *SchedulerTaskScheduleAllOf) SetTaskRequest(v SchedulerRestStimTaskRequest)`

SetTaskRequest sets TaskRequest field to given value.

### HasTaskRequest

`func (o *SchedulerTaskScheduleAllOf) HasTaskRequest() bool`

HasTaskRequest returns a boolean if a field has been set.

### GetType

`func (o *SchedulerTaskScheduleAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchedulerTaskScheduleAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchedulerTaskScheduleAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchedulerTaskScheduleAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *SchedulerTaskScheduleAllOf) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *SchedulerTaskScheduleAllOf) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *SchedulerTaskScheduleAllOf) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *SchedulerTaskScheduleAllOf) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *SchedulerTaskScheduleAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *SchedulerTaskScheduleAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *SchedulerTaskScheduleAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *SchedulerTaskScheduleAllOf) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


