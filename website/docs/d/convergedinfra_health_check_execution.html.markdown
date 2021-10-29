---
subcategory: "convergedinfra"
layout: "intersight"
page_title: "Intersight: intersight_convergedinfra_health_check_execution"
description: |-
  Health check execution result for a health check definition on a converged infra pod.
---

# Data Source: intersight_convergedinfra_health_check_execution
Health check execution result for a health check definition on a converged infra pod.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_convergedinfra_health_check_execution.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error`:(string) Error details of a health check execution failure. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `result`:(string) Health check execution result.* `Unknown` - Indicates that the health check results could not be determined.* `Pass` - Indicates that the health check has passed.* `Fail` - Indicates that the health check has failed.* `Warning` - Indicates that the health check completed with a warning.* `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the pod. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the health check execution.* `Unknown` - Indicates that the health heck execution status is unknown. This mostly happens in case where health check could not be performed due to connectivity issues.* `Succeeded` - Indicates that the health check execution has succeeded.* `Failed` - Indicates that the health check execution has failed.* `Timedout` - Indicates that the health check execution timed out before completion. 
* `summary`:(string) A brief summary of health check results. 
 