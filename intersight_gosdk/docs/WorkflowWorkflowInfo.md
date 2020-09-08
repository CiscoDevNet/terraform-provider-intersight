# WorkflowWorkflowInfo

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

### NewWorkflowWorkflowInfo

`func NewWorkflowWorkflowInfo() *WorkflowWorkflowInfo`

NewWorkflowWorkflowInfo instantiates a new WorkflowWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoWithDefaults

`func NewWorkflowWorkflowInfoWithDefaults() *WorkflowWorkflowInfo`

NewWorkflowWorkflowInfoWithDefaults instantiates a new WorkflowWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowWorkflowInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowWorkflowInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowWorkflowInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowWorkflowInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanupTime

`func (o *WorkflowWorkflowInfo) GetCleanupTime() time.Time`

GetCleanupTime returns the CleanupTime field if non-nil, zero value otherwise.

### GetCleanupTimeOk

`func (o *WorkflowWorkflowInfo) GetCleanupTimeOk() (*time.Time, bool)`

GetCleanupTimeOk returns a tuple with the CleanupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupTime

`func (o *WorkflowWorkflowInfo) SetCleanupTime(v time.Time)`

SetCleanupTime sets CleanupTime field to given value.

### HasCleanupTime

`func (o *WorkflowWorkflowInfo) HasCleanupTime() bool`

HasCleanupTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) GetFailedWorkflowCleanupDuration() int64`

GetFailedWorkflowCleanupDuration returns the FailedWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetFailedWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfo) GetFailedWorkflowCleanupDurationOk() (*int64, bool)`

GetFailedWorkflowCleanupDurationOk returns a tuple with the FailedWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) SetFailedWorkflowCleanupDuration(v int64)`

SetFailedWorkflowCleanupDuration sets FailedWorkflowCleanupDuration field to given value.

### HasFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) HasFailedWorkflowCleanupDuration() bool`

HasFailedWorkflowCleanupDuration returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowInfo) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowInfo) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowInfo) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowInfo) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowWorkflowInfo) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowWorkflowInfo) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetInstId

`func (o *WorkflowWorkflowInfo) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *WorkflowWorkflowInfo) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *WorkflowWorkflowInfo) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *WorkflowWorkflowInfo) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowWorkflowInfo) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowWorkflowInfo) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowWorkflowInfo) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowWorkflowInfo) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLastAction

`func (o *WorkflowWorkflowInfo) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowWorkflowInfo) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowWorkflowInfo) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowWorkflowInfo) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowWorkflowInfo) GetMessage() []WorkflowMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowWorkflowInfo) GetMessageOk() (*[]WorkflowMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowWorkflowInfo) SetMessage(v []WorkflowMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowWorkflowInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetaVersion

`func (o *WorkflowWorkflowInfo) GetMetaVersion() int64`

GetMetaVersion returns the MetaVersion field if non-nil, zero value otherwise.

### GetMetaVersionOk

`func (o *WorkflowWorkflowInfo) GetMetaVersionOk() (*int64, bool)`

GetMetaVersionOk returns a tuple with the MetaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaVersion

`func (o *WorkflowWorkflowInfo) SetMetaVersion(v int64)`

SetMetaVersion sets MetaVersion field to given value.

### HasMetaVersion

`func (o *WorkflowWorkflowInfo) HasMetaVersion() bool`

HasMetaVersion returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowInfo) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowInfo) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowInfo) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowInfo) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowWorkflowInfo) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowWorkflowInfo) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPauseReason

`func (o *WorkflowWorkflowInfo) GetPauseReason() string`

GetPauseReason returns the PauseReason field if non-nil, zero value otherwise.

### GetPauseReasonOk

`func (o *WorkflowWorkflowInfo) GetPauseReasonOk() (*string, bool)`

GetPauseReasonOk returns a tuple with the PauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReason

`func (o *WorkflowWorkflowInfo) SetPauseReason(v string)`

SetPauseReason sets PauseReason field to given value.

### HasPauseReason

`func (o *WorkflowWorkflowInfo) HasPauseReason() bool`

HasPauseReason returns a boolean if a field has been set.

### GetProgress

`func (o *WorkflowWorkflowInfo) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *WorkflowWorkflowInfo) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *WorkflowWorkflowInfo) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *WorkflowWorkflowInfo) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowWorkflowInfo) GetProperties() WorkflowWorkflowInfoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowInfo) GetPropertiesOk() (*WorkflowWorkflowInfoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowInfo) SetProperties(v WorkflowWorkflowInfoProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRetryFromTaskName

`func (o *WorkflowWorkflowInfo) GetRetryFromTaskName() string`

GetRetryFromTaskName returns the RetryFromTaskName field if non-nil, zero value otherwise.

### GetRetryFromTaskNameOk

`func (o *WorkflowWorkflowInfo) GetRetryFromTaskNameOk() (*string, bool)`

GetRetryFromTaskNameOk returns a tuple with the RetryFromTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromTaskName

`func (o *WorkflowWorkflowInfo) SetRetryFromTaskName(v string)`

SetRetryFromTaskName sets RetryFromTaskName field to given value.

### HasRetryFromTaskName

`func (o *WorkflowWorkflowInfo) HasRetryFromTaskName() bool`

HasRetryFromTaskName returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowInfo) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowInfo) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowInfo) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowInfo) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) GetSuccessWorkflowCleanupDuration() int64`

GetSuccessWorkflowCleanupDuration returns the SuccessWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetSuccessWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfo) GetSuccessWorkflowCleanupDurationOk() (*int64, bool)`

GetSuccessWorkflowCleanupDurationOk returns a tuple with the SuccessWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) SetSuccessWorkflowCleanupDuration(v int64)`

SetSuccessWorkflowCleanupDuration sets SuccessWorkflowCleanupDuration field to given value.

### HasSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfo) HasSuccessWorkflowCleanupDuration() bool`

HasSuccessWorkflowCleanupDuration returns a boolean if a field has been set.

### GetTraceId

`func (o *WorkflowWorkflowInfo) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *WorkflowWorkflowInfo) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *WorkflowWorkflowInfo) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *WorkflowWorkflowInfo) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowWorkflowInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowWorkflowInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowWorkflowInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowWorkflowInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowWorkflowInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitReason

`func (o *WorkflowWorkflowInfo) GetWaitReason() string`

GetWaitReason returns the WaitReason field if non-nil, zero value otherwise.

### GetWaitReasonOk

`func (o *WorkflowWorkflowInfo) GetWaitReasonOk() (*string, bool)`

GetWaitReasonOk returns a tuple with the WaitReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitReason

`func (o *WorkflowWorkflowInfo) SetWaitReason(v string)`

SetWaitReason sets WaitReason field to given value.

### HasWaitReason

`func (o *WorkflowWorkflowInfo) HasWaitReason() bool`

HasWaitReason returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowWorkflowInfo) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowWorkflowInfo) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowWorkflowInfo) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowWorkflowInfo) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### GetWorkflowMetaType

`func (o *WorkflowWorkflowInfo) GetWorkflowMetaType() string`

GetWorkflowMetaType returns the WorkflowMetaType field if non-nil, zero value otherwise.

### GetWorkflowMetaTypeOk

`func (o *WorkflowWorkflowInfo) GetWorkflowMetaTypeOk() (*string, bool)`

GetWorkflowMetaTypeOk returns a tuple with the WorkflowMetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetaType

`func (o *WorkflowWorkflowInfo) SetWorkflowMetaType(v string)`

SetWorkflowMetaType sets WorkflowMetaType field to given value.

### HasWorkflowMetaType

`func (o *WorkflowWorkflowInfo) HasWorkflowMetaType() bool`

HasWorkflowMetaType returns a boolean if a field has been set.

### GetWorkflowTaskCount

`func (o *WorkflowWorkflowInfo) GetWorkflowTaskCount() int64`

GetWorkflowTaskCount returns the WorkflowTaskCount field if non-nil, zero value otherwise.

### GetWorkflowTaskCountOk

`func (o *WorkflowWorkflowInfo) GetWorkflowTaskCountOk() (*int64, bool)`

GetWorkflowTaskCountOk returns a tuple with the WorkflowTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskCount

`func (o *WorkflowWorkflowInfo) SetWorkflowTaskCount(v int64)`

SetWorkflowTaskCount sets WorkflowTaskCount field to given value.

### HasWorkflowTaskCount

`func (o *WorkflowWorkflowInfo) HasWorkflowTaskCount() bool`

HasWorkflowTaskCount returns a boolean if a field has been set.

### GetVar0ClusterProfile

`func (o *WorkflowWorkflowInfo) GetVar0ClusterProfile() HyperflexClusterProfileRelationship`

GetVar0ClusterProfile returns the Var0ClusterProfile field if non-nil, zero value otherwise.

### GetVar0ClusterProfileOk

`func (o *WorkflowWorkflowInfo) GetVar0ClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetVar0ClusterProfileOk returns a tuple with the Var0ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0ClusterProfile

`func (o *WorkflowWorkflowInfo) SetVar0ClusterProfile(v HyperflexClusterProfileRelationship)`

SetVar0ClusterProfile sets Var0ClusterProfile field to given value.

### HasVar0ClusterProfile

`func (o *WorkflowWorkflowInfo) HasVar0ClusterProfile() bool`

HasVar0ClusterProfile returns a boolean if a field has been set.

### GetVar1SwitchProfile

`func (o *WorkflowWorkflowInfo) GetVar1SwitchProfile() FabricSwitchProfileRelationship`

GetVar1SwitchProfile returns the Var1SwitchProfile field if non-nil, zero value otherwise.

### GetVar1SwitchProfileOk

`func (o *WorkflowWorkflowInfo) GetVar1SwitchProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetVar1SwitchProfileOk returns a tuple with the Var1SwitchProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1SwitchProfile

`func (o *WorkflowWorkflowInfo) SetVar1SwitchProfile(v FabricSwitchProfileRelationship)`

SetVar1SwitchProfile sets Var1SwitchProfile field to given value.

### HasVar1SwitchProfile

`func (o *WorkflowWorkflowInfo) HasVar1SwitchProfile() bool`

HasVar1SwitchProfile returns a boolean if a field has been set.

### GetAccount

`func (o *WorkflowWorkflowInfo) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WorkflowWorkflowInfo) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WorkflowWorkflowInfo) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WorkflowWorkflowInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *WorkflowWorkflowInfo) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *WorkflowWorkflowInfo) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *WorkflowWorkflowInfo) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *WorkflowWorkflowInfo) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowWorkflowInfo) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowInfo) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowInfo) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetParentTaskInfo

`func (o *WorkflowWorkflowInfo) GetParentTaskInfo() WorkflowTaskInfoRelationship`

GetParentTaskInfo returns the ParentTaskInfo field if non-nil, zero value otherwise.

### GetParentTaskInfoOk

`func (o *WorkflowWorkflowInfo) GetParentTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetParentTaskInfoOk returns a tuple with the ParentTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskInfo

`func (o *WorkflowWorkflowInfo) SetParentTaskInfo(v WorkflowTaskInfoRelationship)`

SetParentTaskInfo sets ParentTaskInfo field to given value.

### HasParentTaskInfo

`func (o *WorkflowWorkflowInfo) HasParentTaskInfo() bool`

HasParentTaskInfo returns a boolean if a field has been set.

### GetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfo) GetPendingDynamicWorkflowInfo() WorkflowPendingDynamicWorkflowInfoRelationship`

GetPendingDynamicWorkflowInfo returns the PendingDynamicWorkflowInfo field if non-nil, zero value otherwise.

### GetPendingDynamicWorkflowInfoOk

`func (o *WorkflowWorkflowInfo) GetPendingDynamicWorkflowInfoOk() (*WorkflowPendingDynamicWorkflowInfoRelationship, bool)`

GetPendingDynamicWorkflowInfoOk returns a tuple with the PendingDynamicWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfo) SetPendingDynamicWorkflowInfo(v WorkflowPendingDynamicWorkflowInfoRelationship)`

SetPendingDynamicWorkflowInfo sets PendingDynamicWorkflowInfo field to given value.

### HasPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfo) HasPendingDynamicWorkflowInfo() bool`

HasPendingDynamicWorkflowInfo returns a boolean if a field has been set.

### GetPermission

`func (o *WorkflowWorkflowInfo) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *WorkflowWorkflowInfo) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *WorkflowWorkflowInfo) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *WorkflowWorkflowInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTaskInfos

`func (o *WorkflowWorkflowInfo) GetTaskInfos() []WorkflowTaskInfoRelationship`

GetTaskInfos returns the TaskInfos field if non-nil, zero value otherwise.

### GetTaskInfosOk

`func (o *WorkflowWorkflowInfo) GetTaskInfosOk() (*[]WorkflowTaskInfoRelationship, bool)`

GetTaskInfosOk returns a tuple with the TaskInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfos

`func (o *WorkflowWorkflowInfo) SetTaskInfos(v []WorkflowTaskInfoRelationship)`

SetTaskInfos sets TaskInfos field to given value.

### HasTaskInfos

`func (o *WorkflowWorkflowInfo) HasTaskInfos() bool`

HasTaskInfos returns a boolean if a field has been set.

### SetTaskInfosNil

`func (o *WorkflowWorkflowInfo) SetTaskInfosNil(b bool)`

 SetTaskInfosNil sets the value for TaskInfos to be an explicit nil

### UnsetTaskInfos
`func (o *WorkflowWorkflowInfo) UnsetTaskInfos()`

UnsetTaskInfos ensures that no value is present for TaskInfos, not even an explicit nil
### GetWorkflowDefinition

`func (o *WorkflowWorkflowInfo) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowWorkflowInfo) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowWorkflowInfo) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowWorkflowInfo) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


