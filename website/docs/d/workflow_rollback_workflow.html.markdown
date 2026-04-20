---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_rollback_workflow"
description: |-
        The RollbackWorkflow object addresses the management and execution of rollback operations within the workflow system. This provides a structured approach to reverting tasks and sub-workflows workflows, ensuring system stability and reliability.
        #### Purpose
        RollbackWorkflow is designed to manage rollback operations within workflows, ensuring controlled and secure recovery processes. It enables the reversal of tasks and sub-workflows to support error correction and system restoration. Additionally, it maintains the relationship between the primary workflow execution and its associated rollback workflows, providing structured management and oversight of rollback activities.
        #### Key Concepts
        - **Eligible tasks:** Holds a list of task executions which are eligible for rollback.
        - **Rollback Strategies:** Implements strategies for rollback, when failure occurs.
        - **Status Monitoring:** Tracks the status of rollback operations, providing insights into execution progress and outcomes.

---

# Data Source: intersight_workflow_rollback_workflow
The RollbackWorkflow object addresses the management and execution of rollback operations within the workflow system. This provides a structured approach to reverting tasks and sub-workflows workflows, ensuring system stability and reliability.  
#### Purpose  
RollbackWorkflow is designed to manage rollback operations within workflows, ensuring controlled and secure recovery processes. It enables the reversal of tasks and sub-workflows to support error correction and system restoration. Additionally, it maintains the relationship between the primary workflow execution and its associated rollback workflows, providing structured management and oversight of rollback activities.
#### Key Concepts  
- **Eligible tasks:** Holds a list of task executions which are eligible for rollback. 
- **Rollback Strategies:** Implements strategies for rollback, when failure occurs. 
- **Status Monitoring:** Tracks the status of rollback operations, providing insights into execution progress and outcomes.
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
 
