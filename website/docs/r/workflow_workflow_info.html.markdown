---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_info"
description: |-
        Contains information for a workflow which is an execution instance of the workflow definition given in the relationship. The workflow definition will provide the schema of the inputs taken to start the workflow execution and the schema of the outputs generated at the end of successful workflow execution. The sequence of tasks to be executed is also provided in the workflow definition. For a workflow to successfully start execution the following properties must be provided- Name, AssociatedObject that carries the relationship to Organization under which the workflow must be executed, WorkflowDefinition, and Inputs with all the required data in order to start workflow execution.

---

# Resource: intersight_workflow_workflow_info
Contains information for a workflow which is an execution instance of the workflow definition given in the relationship. The workflow definition will provide the schema of the inputs taken to start the workflow execution and the schema of the outputs generated at the end of successful workflow execution. The sequence of tasks to be executed is also provided in the workflow definition. For a workflow to successfully start execution the following properties must be provided- Name, AssociatedObject that carries the relationship to Organization under which the workflow must be executed, WorkflowDefinition, and Inputs with all the required data in order to start workflow execution.
## Usage Example
### Resource Creation

```hcl
resource "intersight_workflow_workflow_info" "workflow_workflow_info1" {
  name         = "workflow_workflow_info1"
  action       = "Create"
  properties {
    object_type = "workflow.WorkflowInfoProperties"
    retryable   = false
  }
  success_workflow_cleanup_duration = 2160
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  workflow_definition {
    object_type = "workflow.WorkflowDefinition"
    moid        = var.workflow_workflow_definition
  }
}

variable "workflow_workflow_definition" {
  type        = string
  description = "<moid workflow workflow definition>"
}
```

## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) The action of the workflow such as start, cancel, retry, pause.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Rerun` - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action.* `Retry` - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `RetryFromTask` - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_object`:(HashMap) - A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cleanup_time`:(string)(ReadOnly) The time when the workflow info will be removed from the database. When WorkflowInfo is created, cleanup time will be set to 181 days. As the workflow progresses through different states the cleanup time can be updated. A cleanup time of 0 means the workflow is not scheduled for cleanup. An active workflow that continues to schedule & run tasks can run for any amount of time and there is no upper bound for such workflows. Workflows that are not actively running, say in Paused or Waiting states will be removed after 181 days. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `email`:(string)(ReadOnly) The email address of the user who started this workflow. In the case of LDAP users, this field can hold either a username or an email. 
* `end_time`:(string)(ReadOnly) The time when the workflow reached a final state. 
* `failed_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days. 
* `input`:(JSON as string) All the given inputs for the workflow. The schema for the inputs is defined in the InputDefinition section of the WorkflowDefinition. The InputDefinition will provide a list of input fields to be accepted, the associated datatype of the inputs and any additional constraints on the inputs. For more information please refer to InputDefinition property in the the the WorkflowDefinition resource. The inputs for a workflow are provided as a collection of key-value pairs, where key is the name of the input and value is any valid JSON data which conforms to the datatype of the input as specified in the InputDefinition. When the input passed into a workflow does not match the datatype or the constraints specified in the workflow definition, it will not be accepted. For example, if the InputDefinition specified that workflow must accept a string name 'key' and the value passed for key must adhere to a regex pattern. If Workflow was started with input where 'key' is not a string matching the regex pattern, an error will be generated and workflow will not start execution. During workflow definition design, the input passed into the workflow will be referred using the format 'workflow.input.<inputName>'. If the input is referred directly in a mapping it will be in the format '${workflow.input.<inputName>}' or inside a template mapping in the format '{{.global.workflow.input.<inputName>}}'. 
* `inst_id`:(string)(ReadOnly) A workflow instance Id which is the unique identified for the workflow execution. 
* `internal`:(bool)(ReadOnly) Denotes that an Intersight service started this workflow as internal and hence will not be shown in Intersight User Interface. Typically these are internal system maintenance workflows which are triggered by Intersight services. 
* `last_action`:(string)(ReadOnly) The last action that was issued on the workflow is saved in this field.* `None` - No action is set, this is the default value for action field.* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.* `Start` - Start a new execution of the workflow.* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system.* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.* `Rerun` - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action.* `Retry` - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.* `RetryFromTask` - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt.* `Cancel` - Cancel the workflow that is in running or waiting state. 
* `message`:(Array)
This complex property has following sub-properties:
  + `message`:(string)(ReadOnly) An i18n message that can be translated into multiple languages to support internationalization. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `severity`:(string)(ReadOnly) The severity of the Task or Workflow message warning/error/info etc.* `Info` - The enum represents the log level to be used to convey info message.* `Warning` - The enum represents the log level to be used to convey warning message.* `Debug` - The enum represents the log level to be used to convey debug message.* `Error` - The enum represents the log level to be used to convey error message. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the workflow execution instance. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `output`:(JSON as string)(ReadOnly) All the generated outputs for the workflow. The schema for the outputs are defined in the OutputDefinition section of the WorkflowDefinition. The OutputDefinition will provide a list of output fields that could be generated after workflow execution is completed and the associated datatype of the outputs. For more information please refer to OutputDefinition property in WorkflowDefinition resource. The output for the workflow is generated as a collection of key-value pairs, where key is the name of the output and value is any valid JSON data which conforms to the datatype of output as specified in the OutputDefinition. During workflow definition design, if a workflow is included as a sub-workflow inside a parent workflow then the outputs generated by the sub-workflow can be used in the workflow design. For example, if workflow was included into parent workflow as 'SubWorkflowSample1', then that output can be referred as 'SubWorkflowSample1.output.<outputName>'. In the output is referred directly in a mapping it will be in the format '${SubWorkflowSample1.output.<outputName>}' or inside a template mapping will be in the format '{{SubWorkflowSample1.output.<outputName>}}'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `parent_task_info`:(HashMap) -(ReadOnly) A reference to a workflowTaskInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pause_reason`:(string)(ReadOnly) Denotes the reason workflow is in paused status.* `None` - Pause reason is none, which indicates there is no reason for the pause state.* `TaskWithWarning` - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings.* `SystemMaintenance` - Pause reason indicates the workflow is in this state based on actions of system admin for maintenance. 
* `permission`:(HashMap) - A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `progress`:(float)(ReadOnly) This field indicates percentage of workflow task completion based on the total number of tasks in the workflow. The total number of tasks in the workflow is calculated based on the longest path the workflow execution can take. So progress is calculated based on the percentage of tasks that completed out of the total number of tasks that could be executed. Progress is not a representation of the time taken to complete the workflow. A task is considered as completed if the task status is either \ NO_OP\  or \ COMPLETED\ . If the task status is \ SKIP_TO_FAIL\ , the workflow will be terminated and the progress of the workflow will be set to 100. 
* `properties`:(HashMap) -(ReadOnly) Type to capture all the properties for the workflow info passed on from workflow definition. 
This complex property has following sub-properties:
  + `cancelable`:(HashMap) -(ReadOnly) Holds the parameters and conditions for a workflow to be cancelable. 
This complex property has following sub-properties:
    + `cancelable_states`:
                (Array of schema.TypeString) -
    + `enabled`:(bool) When true the workflow can be cancelled. The action can be further restricted by the mode and cancelableStates properties. 
    + `mode`:(string) Mode controls how the workflow can be canceled.* `ApiOnly` - The workflow can only be canceled via API call.* `All` - The workflow can be canceled from API or from the user interface. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `retryable`:(bool)(ReadOnly) When true, this workflow can be retried within 2 weeks from the last failure. 
  + `rollback_action`:(string)(ReadOnly) Status of rollback for this workflow instance. The rollback action can be enabled, disabled or completed.* `Disabled` - Status of the rollback action when workflow is disabled for rollback.* `Enabled` - Status of the rollback action when workflow is enabled for rollback.* `Completed` - Status of the rollback action once workflow completes the rollback for all eligible tasks. 
  + `rollback_on_cancel`:(bool)(ReadOnly) When set to true, the changes are automatically rolled back if the workflow execution is cancelled. 
  + `rollback_on_failure`:(bool)(ReadOnly) When set to true, the changes are automatically rolled back if the workflow fails to execute. 
* `retry_from_task_name`:(string) This field is required when RetryFromTask action is issued for a workflow that is in a 'final' state. The workflow will be retried from the specified task. This field must specify a task name which is the unique name of the task within the workflow. The task name must be one of the tasks that were completed or failed in the previous run. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src`:(string)(ReadOnly) The source service that started the workflow execution and hence represents the owning service for this workflow. 
* `start_time`:(string)(ReadOnly) The time when the workflow was started for execution. 
* `status`:(string)(ReadOnly) A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). The \ status\  field has been deprecated and is now replaced with the \ workflowStatus\  field. 
* `success_workflow_cleanup_duration`:(int) The duration in hours after which the workflow info for successful workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `task_info_update`:(HashMap) - Used to update a TaskInfo instance in the WorkflowInfo, it is used as a way to update status and provide user inputs for a WaitTask. 
This complex property has following sub-properties:
  + `input`:(JSON as string) Inputs for the specified TaskInfo. Inputs must only be provided for tasks which has included an input definition and the inputs must match the constraints specified in the input definition. 
  + `name`:(string) Name of the task being updated and this name must match the task instance name included inside the workflow definition. This name is also captured in the RefName property of the TaskInfo object for the task. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string) New status of the task being updated, only Failed and Completed statuses are supported, Completed is the default value in case no status is provided.* `Scheduled` - The enum represents the status when task is in scheduled state.* `InProgress` - The enum represents the status when task is in-progress state.* `NoOp` - The enum represents the status when task is a noop.* `Timeout` - The enum represents the status when task has timed out.* `Completed` - The enum represents the status when task has completed.* `Failed` - The enum represents the status when task has failed. 
* `task_infos`:(Array)(ReadOnly) An array of relationships to workflowTaskInfo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `trace_id`:(string)(ReadOnly) The trace id to keep track of workflow execution. 
* `type`:(string)(ReadOnly) A type of the workflow (serverconfig, ansible_monitoring). 
* `user_action_required`:(bool)(ReadOnly) Property will be set when a user action is required on the workflow. This can be because the workflow is waiting for a wait task to be updated, workflow is paused or workflow launched by a configuration object has failed and needs to be retried in order to complete successfully. 
* `user_id`:(string)(ReadOnly) The user identifier which indicates the user that started this workflow. 
* `variable`:(JSON as string)(ReadOnly) All the generated variables for the workflow. During workflow execution, the variables will be updated as per the variableParameters specified after each task execution. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `wait_reason`:(string)(ReadOnly) Denotes the reason workflow is in waiting status.* `None` - Wait reason is none, which indicates there is no reason for the waiting state.* `GatherTasks` - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks.* `Duplicate` - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow.* `RateLimit` - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold.* `WaitTask` - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update.* `PendingRetryFailed` - Wait reason when the workflow is pending a RetryFailed action.* `WaitingToStart` - Workflow is waiting to start on workflow engine. 
* `workflow_ctx`:(HashMap) - The workflow context which contains initiator and target information. 
This complex property has following sub-properties:
  + `initiator_ctx`:(HashMap) - Details about initiator of the workflow. Any Intersight object resource can be set as the initiator of the workflow. For workflows executed by an Intersight service, an applicable service object will be set as the initiator. For example, during server profile deployment workflow, the server profile object will be set as the initiator by the system. For user created workflows, this field is optional and for workflows executed from Intersight workflow execution page, the workflow definition object will be set as the Initiator. 
This complex property has following sub-properties:
    + `initiator_moid`:(string) The moid of the Intersight managed object that initiated the workflow. 
    + `initiator_name`:(string) Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow. 
    + `initiator_type`:(string) Type of Intersight managed object that initiated the workflow. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_ctx_list`:(Array)
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `target_moid`:(string)(ReadOnly) Moid of the target Intersight managed object. 
    + `target_name`:(string)(ReadOnly) Name of the target instance. 
    + `target_type`:(string)(ReadOnly) Object type of the target Intersight managed object. 
  + `workflow_subtype`:(string)(ReadOnly) The subtype of the dynamic workflow. For example - Intersight services offer the following subtypes [Validate, Deploy, Import] for dynamic workflow of type serverconfig. This field is not applicable for user created workflows. 
  + `workflow_type`:(string)(ReadOnly) Intersight services set the type of dynamic workflow that need to be built and executed. This field is not applicable for user created workflows. WorkflowType set as ServerConfig states that a dynamic workflow is executing tasks related to server configuration. 
* `workflow_definition`:(HashMap) - A reference to a workflowWorkflowDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `workflow_status`:(string)(ReadOnly) The current state of the workflow execution instance. A draft workflow execution will be in NotStarted state and when \ Start\  action is issued then the workflow will move into Waiting state until the first task of the workflow is scheduled at which time it will move into InProgress state. When execution reaches a final state it move to either Completed, Failed or Terminated state. For more details look at the description for each state.* `NotStarted` - Initially all the workflow instances are at \ NotStarted\  state. A workflow can be drafted in this state by issuing Create action. When a workflow is in this state the inputs can be updated until the workflow is started.* `InProgress` - A workflow execution moves into \ InProgress\  state when the first task of the workflow is scheduled for execution and continues to remain in that state as long as there are tasks executing or yet to be scheduled for execution.* `Waiting` - Workflow can go to waiting state due to execution of wait task present in the workflow or the workflow has not started yet either due to duplicate workflow is running or due to workflow throttling. Once Workflow engine picks up the workflow for execution, it will move to in progress state.* `Completed` - A workflow execution moves into Completed state when the execution path of the workflow has reached the Success node in the workflow design and there are no more tasks to be executed. Completed is the final state for the workflow execution instance and no further actions are allowed on this workflow instance.* `Failed` - A workflow execution moves into a Failed state when the execution path of the workflow has reached the Failed node in the workflow design and there are no more tasks to be scheduled. A Failed node can be reached when the last executed task has failed or timed out and there are no further retries available for the task. Also as per the workflow design, the last executed task did not specify an OnFailure task to be executed and hence by default, the execution will reach the Failed node. Actions like \ Rerun\ , \ RetryFailed\  and \ RetryFromTask\  can be issued on failed workflow instances. Please refer to the \ Action\  description for more details.* `Terminated` - A workflow execution moves to Terminated state when user issues a \ Cancel\  action or due to internal errors caused during workflow execution. e.g. - Task input transformation has failed. Terminated is a final state of the workflow, no further action are allowed on this workflow instance.* `Canceled` - A workflow execution moves to Canceled state when a user issues a \ Cancel\  action. Cancel is not a final state, the workflow engine will issue cancel to all the running tasks and then move the workflow to the \ Terminated\  state.* `Paused` - A workflow execution moves to Paused state when user issues a \ Pause\  action. When in paused state the current running task will complete its execution but no further tasks will be scheduled until the workflow is resumed. A paused workflow is resumed when the user issues a \ Resume\  action. Paused workflows can be canceled by user. 


## Import
`intersight_workflow_workflow_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_workflow_info.example 1234567890987654321abcde
``` 
