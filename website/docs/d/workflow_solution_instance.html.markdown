---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_solution_instance"
description: |-
  Solution instance is one instance of a solution based on a solution definition.
---

# Data Source: intersight_workflow_solution_instance
Solution instance is one instance of a solution based on a solution definition.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_solution_instance.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for this solution instance. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_status`:(string) Last status of the solution instance which will be reverted when an ongoing solution action instance is aborted.* `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state.* `Okay` - The last action on the solution completed and the solution is in Okay state.* `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the solution instance. Name of the solution instance must be unique within a type of Solution definition. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the solution instance which controls the actions that can be performed on this instance.* `NotCreated` - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state.* `Okay` - The last action on the solution completed and the solution is in Okay state.* `Decommissioned` - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned. 
* `user_id`:(string) The user identifier which indicates the user that started this workflow. 
 