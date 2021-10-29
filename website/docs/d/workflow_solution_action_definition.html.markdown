---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_solution_action_definition"
description: |-
  Definition to capture the details needed to execute an action on the solution.
---

# Data Source: intersight_workflow_solution_action_definition
Definition to capture the details needed to execute an action on the solution.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_solution_action_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action_type`:(string) Type of actionDefinition which decides on how to trigger the action.* `External` - External actions definition can be triggered by enduser to perform actions on the solution. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again.* `Internal` - Internal action definition is used to trigger periodic actions on the solution instance.* `Repetitive` - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created solution. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for this action which provides information on what are the pre-requisites to use this action on the solution and what features are supported by this action. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a solution definition. 
* `periodicity`:(int) Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 