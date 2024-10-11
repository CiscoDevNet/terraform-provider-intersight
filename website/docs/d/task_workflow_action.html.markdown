---
subcategory: "task"
layout: "intersight"
page_title: "Intersight: intersight_task_workflow_action"
description: |-
        Start a test workflow using this object.

---

# Data Source: intersight_task_workflow_action
Start a test workflow using this object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_task_workflow_action.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action for test workflow.* `start` - Start action for the workflow.* `stop` - Stop action for the workflow. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `input_params`:(string) Json formatted string input parameters to workflow. 
* `is_dynamic`:(bool) When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `retryable`:(bool) When true, the retry operation can be performed on the workflow. 
* `sequence_key`:(string) This key can be set so that the workflow execution can be sequenced with workflows having the same key. An example usecase is say there are diferent workflows which operate on the server, and are executed at the same time on the same server and the sequence key for all the workflows are same, then workflow engine will enforce that the workflow execution are sequenced and started one after the other, based on timestamp of the arrival of the execution requests. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `wait_on_duplicate`:(bool) When true, the workflow will wait for previous one to complete before starting a new one. 
* `workflow_context`:(string) Json formatted string that has the contents of the workflow context used when starting a workflow. 
* `workflow_info_moid`:(string) Returns the workflow info moid of the workflow started by workflow action api. workflowInfoMoid will be an empty sting if an error occurs while starting the workflow. 
 
