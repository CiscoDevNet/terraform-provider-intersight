---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_info"
description: |-
  Task instance which represents the run time instance of a task within a workflow.
---

# Data Source: intersight_workflow_task_info
Task instance which represents the run time instance of a task within a workflow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_task_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The task description and this is the description that was added when the task was included into the workflow. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time stamp when the task reached a final state. 
* `failure_reason`:(string) Description of the reason why the task failed. 
* `inst_id`:(string) The instance ID of the task running in the workflow engine. 
* `internal`:(bool) Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow. 
* `label`:(string) User friendly short label to describe this task instance in the workflow. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Task definition name which specifies the task type. 
* `ref_name`:(string) The task reference name to ensure its unique inside a workflow. 
* `retry_count`:(int) A counter for number of retries. 
* `rollback_disabled`:(bool) The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. 
* `running_inst_id`:(string) The instance ID of the task that is currently being executed. When retrying a workflow with failed tasks, the task in workflow engine will have a new instance ID, but the task may still be in-progress. In this case, the task instId reflects the instance ID in the workflow engine, while runningInstId reflects the instance ID of the instance that is currently being executed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) The time stamp when the task started execution. 
* `status`:(string) The status of the task and this will specify if the task is running or has reached a final state. 
 