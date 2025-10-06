---
subcategory: "scheduler"
layout: "intersight"
page_title: "Intersight: intersight_scheduler_schedule_policy"
description: |-
        Metadata used to create a policy to schedule one-time or repeated tasks.

---

# Data Source: intersight_scheduler_schedule_policy
Metadata used to create a policy to schedule one-time or repeated tasks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_scheduler_schedule_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_block_dates`:(bool) Enable or disable block dates. If set to true, the schedule will not run during the block date interval. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage_count`:(int) The number of profiles, templates and deployments that are using this policy. This is used to determine if the policy can be deleted.If the usageCount is greater than 0, the policy cannot be deleted. 
 
