---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_debug_log"
description: |-
        Debug information captured during task execution, when the enable debug flag is set at the workflow definition level.

---

# Data Source: intersight_workflow_task_debug_log
Debug information captured during task execution, when the enable debug flag is set at the workflow definition level.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_task_debug_log.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `retry_count`:(int) A counter for number of retries. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `task_inst_id`:(string) The unique identifier for task instance. 
 
