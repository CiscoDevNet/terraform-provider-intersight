# WorkflowWorkflowInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action of the workflow such as start, cancel, retry, pause. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Create&#x60; - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow. * &#x60;Start&#x60; - Start a new execution of the workflow. * &#x60;Pause&#x60; - Pause the workflow, this can only be issued on workflows that are in running state. * &#x60;Resume&#x60; - Resume the workflow which was previously paused through pause action on the workflow. * &#x60;Retry&#x60; - Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn&#39;t run in the previous iteration. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task. * &#x60;Cancel&#x60; - Cancel the workflow that is in running or waiting state. | [optional] [default to "None"]
**CleanupTime** | Pointer to [**time.Time**](time.Time.md) | The time when the workflow info will be removed from database. | [optional] [readonly] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the workflow reached a final state. | [optional] [readonly] 
**FailedWorkflowCleanupDuration** | Pointer to **int64** | The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. | [optional] 
**Input** | Pointer to **interface{}** | All the given inputs for the workflow. | [optional] 
**InstId** | Pointer to **string** | A workflow instance Id which is the unique identified for the workflow execution. | [optional] [readonly] 
**Internal** | Pointer to **bool** | Denotes if this workflow is internal and should be hidden from user view of running workflows. | [optional] 
**LastAction** | Pointer to **string** | The last action that was issued on the workflow is saved in this field. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Create&#x60; - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow. * &#x60;Start&#x60; - Start a new execution of the workflow. * &#x60;Pause&#x60; - Pause the workflow, this can only be issued on workflows that are in running state. * &#x60;Resume&#x60; - Resume the workflow which was previously paused through pause action on the workflow. * &#x60;Retry&#x60; - Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn&#39;t run in the previous iteration. * &#x60;RetryFailed&#x60; - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task. * &#x60;Cancel&#x60; - Cancel the workflow that is in running or waiting state. | [optional] [readonly] [default to "None"]
**Message** | Pointer to [**[]WorkflowMessage**](workflow.Message.md) |  | [optional] 
**MetaVersion** | Pointer to **int64** | Version of the workflow metadata for which this workflow execution was started. | [optional] 
**Name** | Pointer to **string** | A name of the workflow execution instance. | [optional] 
**Output** | Pointer to **interface{}** | All the generated outputs for the workflow. | [optional] [readonly] 
**PauseReason** | Pointer to **string** | Denotes the reason workflow is in paused status. * &#x60;None&#x60; - Pause reason is none, which indicates there is no reason for the pause state. * &#x60;TaskWithWarning&#x60; - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings. | [optional] [default to "None"]
**Progress** | Pointer to **float32** | This field indicates percentage of workflow task execution. | [optional] [readonly] 
**Properties** | Pointer to [**WorkflowWorkflowInfoProperties**](workflow.WorkflowInfoProperties.md) |  | [optional] 
**RetryFromTaskName** | Pointer to **string** | This field is applicable when Retry action is issued for a workflow which is in a final state. When this field is not specified then the workflow will retry from the start of the workflow. When this field is specified then the workflow will be retried from the specified task. The field should carry the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn&#39;t run in the previous iteration. | [optional] 
**Src** | Pointer to **string** | The source microservice name which is the owner for this workflow. | [optional] [readonly] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time when the workflow was started for execution. | [optional] [readonly] 
**Status** | Pointer to **string** | A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). | [optional] [readonly] 
**SuccessWorkflowCleanupDuration** | Pointer to **int64** | The duration in hours after which the workflow info for successful workflow will be removed from database. | [optional] 
**TraceId** | Pointer to **string** | The trace id to keep track of workflow execution. | [optional] [readonly] 
**Type** | Pointer to **string** | A type of the workflow (serverconfig, ansible_monitoring). | [optional] [readonly] 
**UserId** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**WaitReason** | Pointer to **string** | Denotes the reason workflow is in waiting status. * &#x60;None&#x60; - Wait reason is none, which indicates there is no reason for the waiting state. * &#x60;GatherTasks&#x60; - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks. * &#x60;Duplicate&#x60; - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow. * &#x60;RateLimit&#x60; - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold. * &#x60;WaitTask&#x60; - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update. * &#x60;PendingRetryFailed&#x60; - Wait reason when the workflow is pending a RetryFailed action. | [optional] [default to "None"]
**WorkflowCtx** | Pointer to [**WorkflowWorkflowCtx**](workflow.WorkflowCtx.md) |  | [optional] 
**WorkflowMetaType** | Pointer to **string** | The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance. * &#x60;SystemDefined&#x60; - System defined workflow definition. * &#x60;UserDefined&#x60; - User defined workflow definition. * &#x60;Dynamic&#x60; - Dynamically defined workflow definition. | [optional] [default to "SystemDefined"]
**WorkflowTaskCount** | Pointer to **int64** | Total number of workflow tasks in this workflow. | [optional] [readonly] 
**Var0ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) |  | [optional] 
**Var1SwitchProfile** | Pointer to [**FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**AssociatedObject** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ParentTaskInfo** | Pointer to [**WorkflowTaskInfoRelationship**](workflow.TaskInfo.Relationship.md) |  | [optional] 
**PendingDynamicWorkflowInfo** | Pointer to [**WorkflowPendingDynamicWorkflowInfoRelationship**](workflow.PendingDynamicWorkflowInfo.Relationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**TaskInfos** | Pointer to [**[]WorkflowTaskInfoRelationship**](workflow.TaskInfo.Relationship.md) | An array of relationships to workflowTaskInfo resources. | [optional] [readonly] 
**WorkflowDefinition** | Pointer to [**WorkflowWorkflowDefinitionRelationship**](workflow.WorkflowDefinition.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowInfoAllOf

`func NewWorkflowWorkflowInfoAllOf() *WorkflowWorkflowInfoAllOf`

NewWorkflowWorkflowInfoAllOf instantiates a new WorkflowWorkflowInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoAllOfWithDefaults

`func NewWorkflowWorkflowInfoAllOfWithDefaults() *WorkflowWorkflowInfoAllOf`

NewWorkflowWorkflowInfoAllOfWithDefaults instantiates a new WorkflowWorkflowInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowWorkflowInfoAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowWorkflowInfoAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowWorkflowInfoAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowWorkflowInfoAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanupTime

`func (o *WorkflowWorkflowInfoAllOf) GetCleanupTime() time.Time`

GetCleanupTime returns the CleanupTime field if non-nil, zero value otherwise.

### GetCleanupTimeOk

`func (o *WorkflowWorkflowInfoAllOf) GetCleanupTimeOk() (*time.Time, bool)`

GetCleanupTimeOk returns a tuple with the CleanupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupTime

`func (o *WorkflowWorkflowInfoAllOf) SetCleanupTime(v time.Time)`

SetCleanupTime sets CleanupTime field to given value.

### HasCleanupTime

`func (o *WorkflowWorkflowInfoAllOf) HasCleanupTime() bool`

HasCleanupTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowInfoAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowInfoAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowInfoAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowInfoAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) GetFailedWorkflowCleanupDuration() int64`

GetFailedWorkflowCleanupDuration returns the FailedWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetFailedWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfoAllOf) GetFailedWorkflowCleanupDurationOk() (*int64, bool)`

GetFailedWorkflowCleanupDurationOk returns a tuple with the FailedWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) SetFailedWorkflowCleanupDuration(v int64)`

SetFailedWorkflowCleanupDuration sets FailedWorkflowCleanupDuration field to given value.

### HasFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) HasFailedWorkflowCleanupDuration() bool`

HasFailedWorkflowCleanupDuration returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowInfoAllOf) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowInfoAllOf) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowInfoAllOf) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowInfoAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowWorkflowInfoAllOf) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowWorkflowInfoAllOf) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetInstId

`func (o *WorkflowWorkflowInfoAllOf) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *WorkflowWorkflowInfoAllOf) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *WorkflowWorkflowInfoAllOf) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *WorkflowWorkflowInfoAllOf) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowWorkflowInfoAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowWorkflowInfoAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowWorkflowInfoAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowWorkflowInfoAllOf) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLastAction

`func (o *WorkflowWorkflowInfoAllOf) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowWorkflowInfoAllOf) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowWorkflowInfoAllOf) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowWorkflowInfoAllOf) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowWorkflowInfoAllOf) GetMessage() []WorkflowMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowWorkflowInfoAllOf) GetMessageOk() (*[]WorkflowMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowWorkflowInfoAllOf) SetMessage(v []WorkflowMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowWorkflowInfoAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetaVersion

`func (o *WorkflowWorkflowInfoAllOf) GetMetaVersion() int64`

GetMetaVersion returns the MetaVersion field if non-nil, zero value otherwise.

### GetMetaVersionOk

`func (o *WorkflowWorkflowInfoAllOf) GetMetaVersionOk() (*int64, bool)`

GetMetaVersionOk returns a tuple with the MetaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaVersion

`func (o *WorkflowWorkflowInfoAllOf) SetMetaVersion(v int64)`

SetMetaVersion sets MetaVersion field to given value.

### HasMetaVersion

`func (o *WorkflowWorkflowInfoAllOf) HasMetaVersion() bool`

HasMetaVersion returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowInfoAllOf) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowInfoAllOf) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowInfoAllOf) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowInfoAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowWorkflowInfoAllOf) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowWorkflowInfoAllOf) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPauseReason

`func (o *WorkflowWorkflowInfoAllOf) GetPauseReason() string`

GetPauseReason returns the PauseReason field if non-nil, zero value otherwise.

### GetPauseReasonOk

`func (o *WorkflowWorkflowInfoAllOf) GetPauseReasonOk() (*string, bool)`

GetPauseReasonOk returns a tuple with the PauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReason

`func (o *WorkflowWorkflowInfoAllOf) SetPauseReason(v string)`

SetPauseReason sets PauseReason field to given value.

### HasPauseReason

`func (o *WorkflowWorkflowInfoAllOf) HasPauseReason() bool`

HasPauseReason returns a boolean if a field has been set.

### GetProgress

`func (o *WorkflowWorkflowInfoAllOf) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *WorkflowWorkflowInfoAllOf) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *WorkflowWorkflowInfoAllOf) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *WorkflowWorkflowInfoAllOf) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowWorkflowInfoAllOf) GetProperties() WorkflowWorkflowInfoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowInfoAllOf) GetPropertiesOk() (*WorkflowWorkflowInfoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowInfoAllOf) SetProperties(v WorkflowWorkflowInfoProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowInfoAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRetryFromTaskName

`func (o *WorkflowWorkflowInfoAllOf) GetRetryFromTaskName() string`

GetRetryFromTaskName returns the RetryFromTaskName field if non-nil, zero value otherwise.

### GetRetryFromTaskNameOk

`func (o *WorkflowWorkflowInfoAllOf) GetRetryFromTaskNameOk() (*string, bool)`

GetRetryFromTaskNameOk returns a tuple with the RetryFromTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromTaskName

`func (o *WorkflowWorkflowInfoAllOf) SetRetryFromTaskName(v string)`

SetRetryFromTaskName sets RetryFromTaskName field to given value.

### HasRetryFromTaskName

`func (o *WorkflowWorkflowInfoAllOf) HasRetryFromTaskName() bool`

HasRetryFromTaskName returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowInfoAllOf) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowInfoAllOf) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowInfoAllOf) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowInfoAllOf) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowInfoAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowInfoAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowInfoAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowInfoAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) GetSuccessWorkflowCleanupDuration() int64`

GetSuccessWorkflowCleanupDuration returns the SuccessWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetSuccessWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfoAllOf) GetSuccessWorkflowCleanupDurationOk() (*int64, bool)`

GetSuccessWorkflowCleanupDurationOk returns a tuple with the SuccessWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) SetSuccessWorkflowCleanupDuration(v int64)`

SetSuccessWorkflowCleanupDuration sets SuccessWorkflowCleanupDuration field to given value.

### HasSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoAllOf) HasSuccessWorkflowCleanupDuration() bool`

HasSuccessWorkflowCleanupDuration returns a boolean if a field has been set.

### GetTraceId

`func (o *WorkflowWorkflowInfoAllOf) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *WorkflowWorkflowInfoAllOf) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *WorkflowWorkflowInfoAllOf) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *WorkflowWorkflowInfoAllOf) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowWorkflowInfoAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowInfoAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowInfoAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowInfoAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowWorkflowInfoAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowWorkflowInfoAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowWorkflowInfoAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowWorkflowInfoAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitReason

`func (o *WorkflowWorkflowInfoAllOf) GetWaitReason() string`

GetWaitReason returns the WaitReason field if non-nil, zero value otherwise.

### GetWaitReasonOk

`func (o *WorkflowWorkflowInfoAllOf) GetWaitReasonOk() (*string, bool)`

GetWaitReasonOk returns a tuple with the WaitReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitReason

`func (o *WorkflowWorkflowInfoAllOf) SetWaitReason(v string)`

SetWaitReason sets WaitReason field to given value.

### HasWaitReason

`func (o *WorkflowWorkflowInfoAllOf) HasWaitReason() bool`

HasWaitReason returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowWorkflowInfoAllOf) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowWorkflowInfoAllOf) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### GetWorkflowMetaType

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowMetaType() string`

GetWorkflowMetaType returns the WorkflowMetaType field if non-nil, zero value otherwise.

### GetWorkflowMetaTypeOk

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowMetaTypeOk() (*string, bool)`

GetWorkflowMetaTypeOk returns a tuple with the WorkflowMetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetaType

`func (o *WorkflowWorkflowInfoAllOf) SetWorkflowMetaType(v string)`

SetWorkflowMetaType sets WorkflowMetaType field to given value.

### HasWorkflowMetaType

`func (o *WorkflowWorkflowInfoAllOf) HasWorkflowMetaType() bool`

HasWorkflowMetaType returns a boolean if a field has been set.

### GetWorkflowTaskCount

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowTaskCount() int64`

GetWorkflowTaskCount returns the WorkflowTaskCount field if non-nil, zero value otherwise.

### GetWorkflowTaskCountOk

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowTaskCountOk() (*int64, bool)`

GetWorkflowTaskCountOk returns a tuple with the WorkflowTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskCount

`func (o *WorkflowWorkflowInfoAllOf) SetWorkflowTaskCount(v int64)`

SetWorkflowTaskCount sets WorkflowTaskCount field to given value.

### HasWorkflowTaskCount

`func (o *WorkflowWorkflowInfoAllOf) HasWorkflowTaskCount() bool`

HasWorkflowTaskCount returns a boolean if a field has been set.

### GetVar0ClusterProfile

`func (o *WorkflowWorkflowInfoAllOf) GetVar0ClusterProfile() HyperflexClusterProfileRelationship`

GetVar0ClusterProfile returns the Var0ClusterProfile field if non-nil, zero value otherwise.

### GetVar0ClusterProfileOk

`func (o *WorkflowWorkflowInfoAllOf) GetVar0ClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetVar0ClusterProfileOk returns a tuple with the Var0ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0ClusterProfile

`func (o *WorkflowWorkflowInfoAllOf) SetVar0ClusterProfile(v HyperflexClusterProfileRelationship)`

SetVar0ClusterProfile sets Var0ClusterProfile field to given value.

### HasVar0ClusterProfile

`func (o *WorkflowWorkflowInfoAllOf) HasVar0ClusterProfile() bool`

HasVar0ClusterProfile returns a boolean if a field has been set.

### GetVar1SwitchProfile

`func (o *WorkflowWorkflowInfoAllOf) GetVar1SwitchProfile() FabricSwitchProfileRelationship`

GetVar1SwitchProfile returns the Var1SwitchProfile field if non-nil, zero value otherwise.

### GetVar1SwitchProfileOk

`func (o *WorkflowWorkflowInfoAllOf) GetVar1SwitchProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetVar1SwitchProfileOk returns a tuple with the Var1SwitchProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1SwitchProfile

`func (o *WorkflowWorkflowInfoAllOf) SetVar1SwitchProfile(v FabricSwitchProfileRelationship)`

SetVar1SwitchProfile sets Var1SwitchProfile field to given value.

### HasVar1SwitchProfile

`func (o *WorkflowWorkflowInfoAllOf) HasVar1SwitchProfile() bool`

HasVar1SwitchProfile returns a boolean if a field has been set.

### GetAccount

`func (o *WorkflowWorkflowInfoAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WorkflowWorkflowInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WorkflowWorkflowInfoAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WorkflowWorkflowInfoAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *WorkflowWorkflowInfoAllOf) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *WorkflowWorkflowInfoAllOf) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *WorkflowWorkflowInfoAllOf) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *WorkflowWorkflowInfoAllOf) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowWorkflowInfoAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowInfoAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowInfoAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowInfoAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetParentTaskInfo

`func (o *WorkflowWorkflowInfoAllOf) GetParentTaskInfo() WorkflowTaskInfoRelationship`

GetParentTaskInfo returns the ParentTaskInfo field if non-nil, zero value otherwise.

### GetParentTaskInfoOk

`func (o *WorkflowWorkflowInfoAllOf) GetParentTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetParentTaskInfoOk returns a tuple with the ParentTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskInfo

`func (o *WorkflowWorkflowInfoAllOf) SetParentTaskInfo(v WorkflowTaskInfoRelationship)`

SetParentTaskInfo sets ParentTaskInfo field to given value.

### HasParentTaskInfo

`func (o *WorkflowWorkflowInfoAllOf) HasParentTaskInfo() bool`

HasParentTaskInfo returns a boolean if a field has been set.

### GetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoAllOf) GetPendingDynamicWorkflowInfo() WorkflowPendingDynamicWorkflowInfoRelationship`

GetPendingDynamicWorkflowInfo returns the PendingDynamicWorkflowInfo field if non-nil, zero value otherwise.

### GetPendingDynamicWorkflowInfoOk

`func (o *WorkflowWorkflowInfoAllOf) GetPendingDynamicWorkflowInfoOk() (*WorkflowPendingDynamicWorkflowInfoRelationship, bool)`

GetPendingDynamicWorkflowInfoOk returns a tuple with the PendingDynamicWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoAllOf) SetPendingDynamicWorkflowInfo(v WorkflowPendingDynamicWorkflowInfoRelationship)`

SetPendingDynamicWorkflowInfo sets PendingDynamicWorkflowInfo field to given value.

### HasPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoAllOf) HasPendingDynamicWorkflowInfo() bool`

HasPendingDynamicWorkflowInfo returns a boolean if a field has been set.

### GetPermission

`func (o *WorkflowWorkflowInfoAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *WorkflowWorkflowInfoAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *WorkflowWorkflowInfoAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *WorkflowWorkflowInfoAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTaskInfos

`func (o *WorkflowWorkflowInfoAllOf) GetTaskInfos() []WorkflowTaskInfoRelationship`

GetTaskInfos returns the TaskInfos field if non-nil, zero value otherwise.

### GetTaskInfosOk

`func (o *WorkflowWorkflowInfoAllOf) GetTaskInfosOk() (*[]WorkflowTaskInfoRelationship, bool)`

GetTaskInfosOk returns a tuple with the TaskInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfos

`func (o *WorkflowWorkflowInfoAllOf) SetTaskInfos(v []WorkflowTaskInfoRelationship)`

SetTaskInfos sets TaskInfos field to given value.

### HasTaskInfos

`func (o *WorkflowWorkflowInfoAllOf) HasTaskInfos() bool`

HasTaskInfos returns a boolean if a field has been set.

### SetTaskInfosNil

`func (o *WorkflowWorkflowInfoAllOf) SetTaskInfosNil(b bool)`

 SetTaskInfosNil sets the value for TaskInfos to be an explicit nil

### UnsetTaskInfos
`func (o *WorkflowWorkflowInfoAllOf) UnsetTaskInfos()`

UnsetTaskInfos ensures that no value is present for TaskInfos, not even an explicit nil
### GetWorkflowDefinition

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowWorkflowInfoAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowWorkflowInfoAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowWorkflowInfoAllOf) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


