---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_rollback_workflow"
description: |-
  Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
---

# Resource: intersight_workflow_rollback_workflow
Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
## Argument Reference
The following arguments are supported:
* `action`:(string) The action of the rollback workflow such as Create and Start.* `None` - If no action is set, then the default value is set to none for the action field.* `Create` - Create rollback workflow data for the execution of the rollback workflow.* `Start` - Start a new execution of the rollback workflow. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `continue_on_task_failure`:(bool) When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `primary_workflow`:(Array with Maximum of one item) - A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `rollback_tasks`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `description`:(string)(Computed) Description of the rollback task. 
  + `name`:(string) Name of TaskInfo that needs to be rolled back. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ref_name`:(string) Reference name of TaskInfo that need to be rolled back. 
  + `rollback_completed`:(bool)(Computed) Status of the rollback operation for the task. 
  + `rollback_task_name`:(string)(Computed) Name of TaskInfo that performs the rollback operation. 
  + `status`:(string)(Computed) Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed.* `NotStarted` - Status of rollback task when it is not started rollback.* `Not supported` - Status of task when it is not supporting rollback.* `Completed` - Status of rollback task once execution is successful.* `Failed` - Status of rollback task when it is failed. 
  + `task_info_moid`:(string) Moid of TaskInfo that supports rollback operation. 
  + `task_path`:(string)(Computed) Path of rollback task if it is inside sub-workflow. 
* `rollback_workflows`:(Array)(Computed) An array of relationships to workflowWorkflowInfo resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `selected_tasks`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `description`:(string)(Computed) Description of the rollback task. 
  + `name`:(string) Name of TaskInfo that needs to be rolled back. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ref_name`:(string) Reference name of TaskInfo that need to be rolled back. 
  + `rollback_completed`:(bool)(Computed) Status of the rollback operation for the task. 
  + `rollback_task_name`:(string)(Computed) Name of TaskInfo that performs the rollback operation. 
  + `status`:(string)(Computed) Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed.* `NotStarted` - Status of rollback task when it is not started rollback.* `Not supported` - Status of task when it is not supporting rollback.* `Completed` - Status of rollback task once execution is successful.* `Failed` - Status of rollback task when it is failed. 
  + `task_info_moid`:(string) Moid of TaskInfo that supports rollback operation. 
  + `task_path`:(string)(Computed) Path of rollback task if it is inside sub-workflow. 
* `status`:(string)(Computed) Status of the rollback workflow instance (Created, Running, Completed, Failed).* `None` - If no status is set, then the default value is set none for the status field.* `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback.* `Running` - Status of the rollback workflow when it is in-progress.* `Completed` - Status of the rollback workflow after execution is successful.* `Failed` - Status of the rollback workflow after execution results in failure. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_workflow_rollback_workflow` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_workflow_rollback_workflow.example 1234567890987654321abcde
```