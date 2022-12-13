---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_info"
description: |-
        Contains information for a workflow execution which is a runtime instance of workflow.

---

# Data Source: intersight_workflow_workflow_info
Contains information for a workflow execution which is a runtime instance of workflow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_workflow_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action of the workflow such as start, cancel, retry, pause.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Retry` - Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `cleanup_time`:(string) The time when the workflow info will be removed from database. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email`:(string) The email address of the user who started this workflow. 
* `end_time`:(string) The time when the workflow reached a final state. 
* `failed_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. 
* `inst_id`:(string) A workflow instance Id which is the unique identified for the workflow execution. 
* `internal`:(bool) Denotes if this workflow is internal and should be hidden from user view of running workflows. 
* `last_action`:(string) The last action that was issued on the workflow is saved in this field.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Retry` - Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `meta_version`:(int) Version of the workflow metadata for which this workflow execution was started. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the workflow execution instance. 
* `pause_reason`:(string) Denotes the reason workflow is in paused status.* `None` - Pause reason is none, which indicates there is no reason for the pause state.* `TaskWithWarning` - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings.* `SystemMaintenance` - Pause reason indicates the workflow is in this state based on actions of system admin for maintenance. 
* `progress`:(float) The progress of a workflow is calculated based on the total number of tasks in the workflow and the number of tasks completed. A task is considered as completed if the task status is either \ NO_OP\  or \ COMPLETED\ . If the task status is \ SKIP_TO_FAIL\ , the workflow will be terminated and the progress of the workflow will be set to 100. 
* `retry_from_task_name`:(string) This field is applicable when Retry action is issued for a workflow which is in 'final' state. When this field is not specified, the workflow will be retried from the start i.e., the first task. When this field is specified then the workflow will be retried from the specified task. This field should specify the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src`:(string) The source microservice name which is the owner of this workflow. 
* `start_time`:(string) The time when the workflow was started for execution. 
* `status`:(string) A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). 
* `success_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for successful workflow will be removed from database. 
* `trace_id`:(string) The trace id to keep track of workflow execution. 
* `type`:(string) A type of the workflow (serverconfig, ansible_monitoring). 
* `user_action_required`:(bool) Property will be set when a user action is required on the workflow. This can be because the workflow is waiting for a wait task to be updated, workflow is paused or workflow launched by a configuration object has failed and needs to be retried in order to complete successfully. 
* `user_id`:(string) The user identifier which indicates the user that started this workflow. 
* `wait_reason`:(string) Denotes the reason workflow is in waiting status.* `None` - Wait reason is none, which indicates there is no reason for the waiting state.* `GatherTasks` - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks.* `Duplicate` - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow.* `RateLimit` - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold.* `WaitTask` - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update.* `PendingRetryFailed` - Wait reason when the workflow is pending a RetryFailed action.* `WaitingToStart` - Workflow is waiting to start on workflow engine. 
* `workflow_meta_type`:(string) The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.* `SystemDefined` - System defined workflow definition.* `UserDefined` - User defined workflow definition.* `Dynamic` - Dynamically defined workflow definition. 
 
