# WorkflowWorkflowInfoRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
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

### NewWorkflowWorkflowInfoRelationship

`func NewWorkflowWorkflowInfoRelationship(classId string, objectType string, ) *WorkflowWorkflowInfoRelationship`

NewWorkflowWorkflowInfoRelationship instantiates a new WorkflowWorkflowInfoRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowInfoRelationshipWithDefaults

`func NewWorkflowWorkflowInfoRelationshipWithDefaults() *WorkflowWorkflowInfoRelationship`

NewWorkflowWorkflowInfoRelationshipWithDefaults instantiates a new WorkflowWorkflowInfoRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *WorkflowWorkflowInfoRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *WorkflowWorkflowInfoRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *WorkflowWorkflowInfoRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *WorkflowWorkflowInfoRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *WorkflowWorkflowInfoRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowInfoRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowInfoRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *WorkflowWorkflowInfoRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowWorkflowInfoRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowWorkflowInfoRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowWorkflowInfoRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *WorkflowWorkflowInfoRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *WorkflowWorkflowInfoRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *WorkflowWorkflowInfoRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *WorkflowWorkflowInfoRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *WorkflowWorkflowInfoRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *WorkflowWorkflowInfoRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *WorkflowWorkflowInfoRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *WorkflowWorkflowInfoRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *WorkflowWorkflowInfoRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *WorkflowWorkflowInfoRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *WorkflowWorkflowInfoRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *WorkflowWorkflowInfoRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowWorkflowInfoRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowInfoRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowInfoRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *WorkflowWorkflowInfoRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *WorkflowWorkflowInfoRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *WorkflowWorkflowInfoRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *WorkflowWorkflowInfoRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *WorkflowWorkflowInfoRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *WorkflowWorkflowInfoRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *WorkflowWorkflowInfoRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *WorkflowWorkflowInfoRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowWorkflowInfoRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowWorkflowInfoRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowWorkflowInfoRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowWorkflowInfoRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *WorkflowWorkflowInfoRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *WorkflowWorkflowInfoRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *WorkflowWorkflowInfoRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *WorkflowWorkflowInfoRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *WorkflowWorkflowInfoRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *WorkflowWorkflowInfoRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *WorkflowWorkflowInfoRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *WorkflowWorkflowInfoRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *WorkflowWorkflowInfoRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *WorkflowWorkflowInfoRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *WorkflowWorkflowInfoRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WorkflowWorkflowInfoRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WorkflowWorkflowInfoRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WorkflowWorkflowInfoRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *WorkflowWorkflowInfoRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *WorkflowWorkflowInfoRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *WorkflowWorkflowInfoRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *WorkflowWorkflowInfoRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *WorkflowWorkflowInfoRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *WorkflowWorkflowInfoRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *WorkflowWorkflowInfoRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *WorkflowWorkflowInfoRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *WorkflowWorkflowInfoRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *WorkflowWorkflowInfoRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *WorkflowWorkflowInfoRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *WorkflowWorkflowInfoRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAction

`func (o *WorkflowWorkflowInfoRelationship) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowWorkflowInfoRelationship) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowWorkflowInfoRelationship) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowWorkflowInfoRelationship) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCleanupTime

`func (o *WorkflowWorkflowInfoRelationship) GetCleanupTime() time.Time`

GetCleanupTime returns the CleanupTime field if non-nil, zero value otherwise.

### GetCleanupTimeOk

`func (o *WorkflowWorkflowInfoRelationship) GetCleanupTimeOk() (*time.Time, bool)`

GetCleanupTimeOk returns a tuple with the CleanupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupTime

`func (o *WorkflowWorkflowInfoRelationship) SetCleanupTime(v time.Time)`

SetCleanupTime sets CleanupTime field to given value.

### HasCleanupTime

`func (o *WorkflowWorkflowInfoRelationship) HasCleanupTime() bool`

HasCleanupTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowInfoRelationship) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowInfoRelationship) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowInfoRelationship) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowInfoRelationship) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) GetFailedWorkflowCleanupDuration() int64`

GetFailedWorkflowCleanupDuration returns the FailedWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetFailedWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfoRelationship) GetFailedWorkflowCleanupDurationOk() (*int64, bool)`

GetFailedWorkflowCleanupDurationOk returns a tuple with the FailedWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) SetFailedWorkflowCleanupDuration(v int64)`

SetFailedWorkflowCleanupDuration sets FailedWorkflowCleanupDuration field to given value.

### HasFailedWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) HasFailedWorkflowCleanupDuration() bool`

HasFailedWorkflowCleanupDuration returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowInfoRelationship) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowInfoRelationship) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowInfoRelationship) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowInfoRelationship) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowWorkflowInfoRelationship) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowWorkflowInfoRelationship) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetInstId

`func (o *WorkflowWorkflowInfoRelationship) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *WorkflowWorkflowInfoRelationship) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *WorkflowWorkflowInfoRelationship) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *WorkflowWorkflowInfoRelationship) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInternal

`func (o *WorkflowWorkflowInfoRelationship) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowWorkflowInfoRelationship) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowWorkflowInfoRelationship) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowWorkflowInfoRelationship) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLastAction

`func (o *WorkflowWorkflowInfoRelationship) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *WorkflowWorkflowInfoRelationship) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *WorkflowWorkflowInfoRelationship) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *WorkflowWorkflowInfoRelationship) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowWorkflowInfoRelationship) GetMessage() []WorkflowMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowWorkflowInfoRelationship) GetMessageOk() (*[]WorkflowMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowWorkflowInfoRelationship) SetMessage(v []WorkflowMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowWorkflowInfoRelationship) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetaVersion

`func (o *WorkflowWorkflowInfoRelationship) GetMetaVersion() int64`

GetMetaVersion returns the MetaVersion field if non-nil, zero value otherwise.

### GetMetaVersionOk

`func (o *WorkflowWorkflowInfoRelationship) GetMetaVersionOk() (*int64, bool)`

GetMetaVersionOk returns a tuple with the MetaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaVersion

`func (o *WorkflowWorkflowInfoRelationship) SetMetaVersion(v int64)`

SetMetaVersion sets MetaVersion field to given value.

### HasMetaVersion

`func (o *WorkflowWorkflowInfoRelationship) HasMetaVersion() bool`

HasMetaVersion returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowInfoRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowInfoRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowInfoRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowInfoRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowInfoRelationship) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowInfoRelationship) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowInfoRelationship) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowInfoRelationship) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowWorkflowInfoRelationship) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowWorkflowInfoRelationship) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPauseReason

`func (o *WorkflowWorkflowInfoRelationship) GetPauseReason() string`

GetPauseReason returns the PauseReason field if non-nil, zero value otherwise.

### GetPauseReasonOk

`func (o *WorkflowWorkflowInfoRelationship) GetPauseReasonOk() (*string, bool)`

GetPauseReasonOk returns a tuple with the PauseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseReason

`func (o *WorkflowWorkflowInfoRelationship) SetPauseReason(v string)`

SetPauseReason sets PauseReason field to given value.

### HasPauseReason

`func (o *WorkflowWorkflowInfoRelationship) HasPauseReason() bool`

HasPauseReason returns a boolean if a field has been set.

### GetProgress

`func (o *WorkflowWorkflowInfoRelationship) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *WorkflowWorkflowInfoRelationship) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *WorkflowWorkflowInfoRelationship) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *WorkflowWorkflowInfoRelationship) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowWorkflowInfoRelationship) GetProperties() WorkflowWorkflowInfoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowInfoRelationship) GetPropertiesOk() (*WorkflowWorkflowInfoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowInfoRelationship) SetProperties(v WorkflowWorkflowInfoProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowInfoRelationship) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRetryFromTaskName

`func (o *WorkflowWorkflowInfoRelationship) GetRetryFromTaskName() string`

GetRetryFromTaskName returns the RetryFromTaskName field if non-nil, zero value otherwise.

### GetRetryFromTaskNameOk

`func (o *WorkflowWorkflowInfoRelationship) GetRetryFromTaskNameOk() (*string, bool)`

GetRetryFromTaskNameOk returns a tuple with the RetryFromTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFromTaskName

`func (o *WorkflowWorkflowInfoRelationship) SetRetryFromTaskName(v string)`

SetRetryFromTaskName sets RetryFromTaskName field to given value.

### HasRetryFromTaskName

`func (o *WorkflowWorkflowInfoRelationship) HasRetryFromTaskName() bool`

HasRetryFromTaskName returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowWorkflowInfoRelationship) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowWorkflowInfoRelationship) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowWorkflowInfoRelationship) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowWorkflowInfoRelationship) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowInfoRelationship) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowInfoRelationship) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowInfoRelationship) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowInfoRelationship) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowInfoRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowInfoRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowInfoRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowInfoRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) GetSuccessWorkflowCleanupDuration() int64`

GetSuccessWorkflowCleanupDuration returns the SuccessWorkflowCleanupDuration field if non-nil, zero value otherwise.

### GetSuccessWorkflowCleanupDurationOk

`func (o *WorkflowWorkflowInfoRelationship) GetSuccessWorkflowCleanupDurationOk() (*int64, bool)`

GetSuccessWorkflowCleanupDurationOk returns a tuple with the SuccessWorkflowCleanupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) SetSuccessWorkflowCleanupDuration(v int64)`

SetSuccessWorkflowCleanupDuration sets SuccessWorkflowCleanupDuration field to given value.

### HasSuccessWorkflowCleanupDuration

`func (o *WorkflowWorkflowInfoRelationship) HasSuccessWorkflowCleanupDuration() bool`

HasSuccessWorkflowCleanupDuration returns a boolean if a field has been set.

### GetTraceId

`func (o *WorkflowWorkflowInfoRelationship) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *WorkflowWorkflowInfoRelationship) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *WorkflowWorkflowInfoRelationship) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *WorkflowWorkflowInfoRelationship) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowWorkflowInfoRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowWorkflowInfoRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowWorkflowInfoRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowWorkflowInfoRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowWorkflowInfoRelationship) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowWorkflowInfoRelationship) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowWorkflowInfoRelationship) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowWorkflowInfoRelationship) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWaitReason

`func (o *WorkflowWorkflowInfoRelationship) GetWaitReason() string`

GetWaitReason returns the WaitReason field if non-nil, zero value otherwise.

### GetWaitReasonOk

`func (o *WorkflowWorkflowInfoRelationship) GetWaitReasonOk() (*string, bool)`

GetWaitReasonOk returns a tuple with the WaitReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitReason

`func (o *WorkflowWorkflowInfoRelationship) SetWaitReason(v string)`

SetWaitReason sets WaitReason field to given value.

### HasWaitReason

`func (o *WorkflowWorkflowInfoRelationship) HasWaitReason() bool`

HasWaitReason returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowWorkflowInfoRelationship) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowWorkflowInfoRelationship) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### GetWorkflowMetaType

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowMetaType() string`

GetWorkflowMetaType returns the WorkflowMetaType field if non-nil, zero value otherwise.

### GetWorkflowMetaTypeOk

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowMetaTypeOk() (*string, bool)`

GetWorkflowMetaTypeOk returns a tuple with the WorkflowMetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetaType

`func (o *WorkflowWorkflowInfoRelationship) SetWorkflowMetaType(v string)`

SetWorkflowMetaType sets WorkflowMetaType field to given value.

### HasWorkflowMetaType

`func (o *WorkflowWorkflowInfoRelationship) HasWorkflowMetaType() bool`

HasWorkflowMetaType returns a boolean if a field has been set.

### GetWorkflowTaskCount

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowTaskCount() int64`

GetWorkflowTaskCount returns the WorkflowTaskCount field if non-nil, zero value otherwise.

### GetWorkflowTaskCountOk

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowTaskCountOk() (*int64, bool)`

GetWorkflowTaskCountOk returns a tuple with the WorkflowTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskCount

`func (o *WorkflowWorkflowInfoRelationship) SetWorkflowTaskCount(v int64)`

SetWorkflowTaskCount sets WorkflowTaskCount field to given value.

### HasWorkflowTaskCount

`func (o *WorkflowWorkflowInfoRelationship) HasWorkflowTaskCount() bool`

HasWorkflowTaskCount returns a boolean if a field has been set.

### GetVar0ClusterProfile

`func (o *WorkflowWorkflowInfoRelationship) GetVar0ClusterProfile() HyperflexClusterProfileRelationship`

GetVar0ClusterProfile returns the Var0ClusterProfile field if non-nil, zero value otherwise.

### GetVar0ClusterProfileOk

`func (o *WorkflowWorkflowInfoRelationship) GetVar0ClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetVar0ClusterProfileOk returns a tuple with the Var0ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0ClusterProfile

`func (o *WorkflowWorkflowInfoRelationship) SetVar0ClusterProfile(v HyperflexClusterProfileRelationship)`

SetVar0ClusterProfile sets Var0ClusterProfile field to given value.

### HasVar0ClusterProfile

`func (o *WorkflowWorkflowInfoRelationship) HasVar0ClusterProfile() bool`

HasVar0ClusterProfile returns a boolean if a field has been set.

### GetVar1SwitchProfile

`func (o *WorkflowWorkflowInfoRelationship) GetVar1SwitchProfile() FabricSwitchProfileRelationship`

GetVar1SwitchProfile returns the Var1SwitchProfile field if non-nil, zero value otherwise.

### GetVar1SwitchProfileOk

`func (o *WorkflowWorkflowInfoRelationship) GetVar1SwitchProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetVar1SwitchProfileOk returns a tuple with the Var1SwitchProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1SwitchProfile

`func (o *WorkflowWorkflowInfoRelationship) SetVar1SwitchProfile(v FabricSwitchProfileRelationship)`

SetVar1SwitchProfile sets Var1SwitchProfile field to given value.

### HasVar1SwitchProfile

`func (o *WorkflowWorkflowInfoRelationship) HasVar1SwitchProfile() bool`

HasVar1SwitchProfile returns a boolean if a field has been set.

### GetAccount

`func (o *WorkflowWorkflowInfoRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WorkflowWorkflowInfoRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WorkflowWorkflowInfoRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WorkflowWorkflowInfoRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *WorkflowWorkflowInfoRelationship) GetAssociatedObject() MoBaseMoRelationship`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *WorkflowWorkflowInfoRelationship) GetAssociatedObjectOk() (*MoBaseMoRelationship, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *WorkflowWorkflowInfoRelationship) SetAssociatedObject(v MoBaseMoRelationship)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *WorkflowWorkflowInfoRelationship) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowWorkflowInfoRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowWorkflowInfoRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowWorkflowInfoRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowWorkflowInfoRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetParentTaskInfo

`func (o *WorkflowWorkflowInfoRelationship) GetParentTaskInfo() WorkflowTaskInfoRelationship`

GetParentTaskInfo returns the ParentTaskInfo field if non-nil, zero value otherwise.

### GetParentTaskInfoOk

`func (o *WorkflowWorkflowInfoRelationship) GetParentTaskInfoOk() (*WorkflowTaskInfoRelationship, bool)`

GetParentTaskInfoOk returns a tuple with the ParentTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskInfo

`func (o *WorkflowWorkflowInfoRelationship) SetParentTaskInfo(v WorkflowTaskInfoRelationship)`

SetParentTaskInfo sets ParentTaskInfo field to given value.

### HasParentTaskInfo

`func (o *WorkflowWorkflowInfoRelationship) HasParentTaskInfo() bool`

HasParentTaskInfo returns a boolean if a field has been set.

### GetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoRelationship) GetPendingDynamicWorkflowInfo() WorkflowPendingDynamicWorkflowInfoRelationship`

GetPendingDynamicWorkflowInfo returns the PendingDynamicWorkflowInfo field if non-nil, zero value otherwise.

### GetPendingDynamicWorkflowInfoOk

`func (o *WorkflowWorkflowInfoRelationship) GetPendingDynamicWorkflowInfoOk() (*WorkflowPendingDynamicWorkflowInfoRelationship, bool)`

GetPendingDynamicWorkflowInfoOk returns a tuple with the PendingDynamicWorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoRelationship) SetPendingDynamicWorkflowInfo(v WorkflowPendingDynamicWorkflowInfoRelationship)`

SetPendingDynamicWorkflowInfo sets PendingDynamicWorkflowInfo field to given value.

### HasPendingDynamicWorkflowInfo

`func (o *WorkflowWorkflowInfoRelationship) HasPendingDynamicWorkflowInfo() bool`

HasPendingDynamicWorkflowInfo returns a boolean if a field has been set.

### GetPermission

`func (o *WorkflowWorkflowInfoRelationship) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *WorkflowWorkflowInfoRelationship) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *WorkflowWorkflowInfoRelationship) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *WorkflowWorkflowInfoRelationship) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTaskInfos

`func (o *WorkflowWorkflowInfoRelationship) GetTaskInfos() []WorkflowTaskInfoRelationship`

GetTaskInfos returns the TaskInfos field if non-nil, zero value otherwise.

### GetTaskInfosOk

`func (o *WorkflowWorkflowInfoRelationship) GetTaskInfosOk() (*[]WorkflowTaskInfoRelationship, bool)`

GetTaskInfosOk returns a tuple with the TaskInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInfos

`func (o *WorkflowWorkflowInfoRelationship) SetTaskInfos(v []WorkflowTaskInfoRelationship)`

SetTaskInfos sets TaskInfos field to given value.

### HasTaskInfos

`func (o *WorkflowWorkflowInfoRelationship) HasTaskInfos() bool`

HasTaskInfos returns a boolean if a field has been set.

### SetTaskInfosNil

`func (o *WorkflowWorkflowInfoRelationship) SetTaskInfosNil(b bool)`

 SetTaskInfosNil sets the value for TaskInfos to be an explicit nil

### UnsetTaskInfos
`func (o *WorkflowWorkflowInfoRelationship) UnsetTaskInfos()`

UnsetTaskInfos ensures that no value is present for TaskInfos, not even an explicit nil
### GetWorkflowDefinition

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowWorkflowInfoRelationship) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowWorkflowInfoRelationship) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowWorkflowInfoRelationship) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


