---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_rollback_workflow"
description: |-
  Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
---

# Data Source: intersight_workflow_rollback_workflow
Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_rollback_workflow.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action of the rollback workflow such as Create and Start.* `None` - If no action is set, then the default value is set to none for the action field.* `Create` - Create rollback workflow data for the execution of the rollback workflow.* `Start` - Start a new execution of the rollback workflow. 
* `continue_on_task_failure`:(bool) When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the rollback workflow instance (Created, Running, Completed, Failed).* `None` - If no status is set, then the default value is set none for the status field.* `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback.* `Running` - Status of the rollback workflow when it is in-progress.* `Completed` - Status of the rollback workflow after execution is successful.* `Failed` - Status of the rollback workflow after execution results in failure. 
 