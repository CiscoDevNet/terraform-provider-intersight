---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_info"
description: |-
  Contains information for a workflow execution which is a runtime instance of workflow.
---

# Resource: intersight_workflow_workflow_info
Contains information for a workflow execution which is a runtime instance of workflow.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `action`:(string) The action of the workflow such as start, cancel, retry, pause.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Retry` - Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `associated_object`:(HashMap) - A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cleanup_time`:(string)(Computed) The time when the workflow info will be removed from database. 
* `email`:(string)(Computed) The email address of the user who started this workflow. 
* `end_time`:(string)(Computed) The time when the workflow reached a final state. 
* `failed_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. 
* `input`: All the given inputs for the workflow. 
* `inst_id`:(string)(Computed) A workflow instance Id which is the unique identified for the workflow execution. 
* `internal`:(bool) Denotes if this workflow is internal and should be hidden from user view of running workflows. 
* `last_action`:(string)(Computed) The last action that was issued on the workflow is saved in this field.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Retry` - Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `message`:(Array)
This complex property has following sub-properties:
  + `message`:(string) An i18n message that can be translated in multiple languages to support internationalization. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `severity`:(string) The severity of the Task or Workflow message warning/error/info etc.* `Info` - The enum represents the log level to be used to convey info message.* `Warning` - The enum represents the log level to be used to convey warning message.* `Debug` - The enum represents the log level to be used to convey debug message.* `Error` - The enum represents the log level to be used to convey error message. 
* `meta_version`:(int) Version of the workflow metadata for which this workflow execution was started. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the workflow execution instance. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `output`:(Computed) All the generated outputs for the workflow. 
* `parent_task_info`:(HashMap) -(Computed) A reference to a workflowTaskInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pause_reason`:(string) Denotes the reason workflow is in paused status.* `None` - Pause reason is none, which indicates there is no reason for the pause state.* `TaskWithWarning` - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings. 
* `pending_dynamic_workflow_info`:(HashMap) -(Computed) A reference to a workflowPendingDynamicWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission`:(HashMap) - A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `progress`:(float)(Computed) This field indicates percentage of workflow task execution. 
* `properties`:(HashMap) -(Computed) Type to capture all the properties for the workflow info passed on from workflow definition. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `retryable`:(bool) When true, this workflow can be retried if has not been modified for more than a period of 2 weeks. 
  + `rollback_action`:(string)(Computed) Status of rollback for this workflow instance. The rollback action of the workflow can be enabled, disabled, completed.* `Disabled` - Status of the rollback action when workflow is disabled for rollback.* `Enabled` - Status of the rollback action when workflow is enabled for rollback.* `Completed` - Status of the rollback action once workflow completes the rollback for all eligiable tasks. 
* `retry_from_task_name`:(string) This field is applicable when Retry action is issued for a workflow which is in a final state. When this field is not specified then the workflow will retry from the start of the workflow. When this field is specified then the workflow will be retried from the specified task. The field should carry the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration. 
* `src`:(string)(Computed) The source microservice name which is the owner for this workflow. 
* `start_time`:(string)(Computed) The time when the workflow was started for execution. 
* `status`:(string)(Computed) A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). 
* `success_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for successful workflow will be removed from database. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `task_infos`:(Array)(Computed) An array of relationships to workflowTaskInfo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `trace_id`:(string)(Computed) The trace id to keep track of workflow execution. 
* `type`:(string)(Computed) A type of the workflow (serverconfig, ansible_monitoring). 
* `user_action_required`:(bool)(Computed) Property will be set when an user action is required on the workflow. This can be because the workflow is waiting for a wait task to be updated, workflow is paused or workflow launched by a configuration object has failed and needs to be retried in order to complete successfully. 
* `user_id`:(string)(Computed) The user identifier which indicates the user that started this workflow. 
* `wait_reason`:(string) Denotes the reason workflow is in waiting status.* `None` - Wait reason is none, which indicates there is no reason for the waiting state.* `GatherTasks` - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks.* `Duplicate` - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow.* `RateLimit` - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold.* `WaitTask` - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update.* `PendingRetryFailed` - Wait reason when the workflow is pending a RetryFailed action. 
* `workflow_ctx`:(HashMap) - The workflow context which contains initiator and target information. 
This complex property has following sub-properties:
  + `initiator_ctx`:(HashMap) - Details about initiator of the workflow. 
This complex property has following sub-properties:
    + `initiator_moid`:(string) The moid of the Intersigt managed object that initiated the workflow. 
    + `initiator_name`:(string) Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow. 
    + `initiator_type`:(string) Type of Intersight managed object that initiated the workflow. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_ctx_list`:(Array)
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `target_moid`:(string) Moid of the target Intersight managed object. 
    + `target_name`:(string) Name of the target instance. 
    + `target_type`:(string) Object type of the target Intersight managed object. 
  + `workflow_meta_name`:(string) The name of workflowMeta of the workflow running. 
  + `workflow_subtype`:(string) The subtype of the workflow. 
  + `workflow_type`:(string) Type of the workflow being started. This can be any string for client services to distinguish workflow by type. 
* `workflow_definition`:(HashMap) - A reference to a workflowWorkflowDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `workflow_meta_type`:(string) The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.* `SystemDefined` - System defined workflow definition.* `UserDefined` - User defined workflow definition.* `Dynamic` - Dynamically defined workflow definition. 
* `workflow_task_count`:(int)(Computed) Total number of workflow tasks in this workflow. 
* `workflow_worker_task_count`:(int)(Computed) Total number of worker tasks in this workflow. This count doesn't include the control tasks in the workflow. 


## Import
`intersight_workflow_workflow_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_workflow_info.example 1234567890987654321abcde
``` 