---
subcategory: "openapi"
layout: "intersight"
page_title: "Intersight: intersight_openapi_task_generation_result"
description: |-
        Provides information about the status of the tasks created for a TaskGenerationRequest.

---

# Data Source: intersight_openapi_task_generation_result
Provides information about the status of the tasks created for a TaskGenerationRequest.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_openapi_task_generation_result.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `failure_reason`:(string) An error message for when task creation fails. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `task_display_name`:(string) The display label of the task that is created. 
* `task_name`:(string) The name of the task that is created. 
* `task_status`:(string) Denotes the status of the task creation.* `none` - Indicates the default status* `InProgress` - Indicates that operation is in progress* `Completed` - Indicates that the operation is complete* `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details. 
* `task_version`:(int) The version number of the created tasks. 
 
