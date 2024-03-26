---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_schedule_policy"
description: |-
        Continuous health check schedule policy of a HyperFlex cluster.

---

# Data Source: intersight_hyperflex_health_check_schedule_policy
Continuous health check schedule policy of a HyperFlex cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_health_check_schedule_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_scheduled_on`:(string) The date and time when this HealthCheck Policy was last enabled. 
* `last_unscheduled_on`:(string) The date and time when this HealthCheck Policy was last disabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_expected_execution`:(string) The date and time when the next health check execution is expected. 
* `policy_enabled`:(bool) Indicates whether HealthCheck schedule policy is enabled on the HyperFlex cluster. 
* `schedule_interval`:(int) The frequency at which the health checks are run on the HyperFlex cluster.* `86400` - Execute the health check every 24 hours.* `43200` - Execute the health check every 12 hours.* `21600` - Execute the health check every 6 hours.* `10800` - Execute the health check every 3 hours.* `3600` - Execute the health check every 1 hours.* `300` - Execute the health check every 5 minutes.* `0` - Disable the continuous health check. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) The unique identifier of the health check policy. 
 
