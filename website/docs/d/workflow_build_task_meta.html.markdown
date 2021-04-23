---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta"
description: |-
  Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
---

# Data Source: intersight_workflow_build_task_meta
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_build_task_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name for the BuildTaskMeta instance used to created a dynamic workflow. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src`:(string) Microservice owner for the task in this workflow. 
* `task_type`:(string) The type of the task within this workflow. 
* `workflow_type`:(string) The type for the dynamic workflow. 
 