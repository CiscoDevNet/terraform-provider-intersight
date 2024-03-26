---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_health_check_execution_snapshot"
description: |-
        Health check execution snapshot of the HyperFlex cluster.

---

# Data Source: intersight_hyperflex_cluster_health_check_execution_snapshot
Health check execution snapshot of the HyperFlex cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_health_check_execution_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `execution_context`:(string) The execution context of the HyperFlex health checks.* `UNKNOWN` - The current context of HyperFlex health check execution is unknown.* `WORKFLOW` - The HyperFlex health check execution is initiated through an orchestration workflow.* `SCHEDULED` - The HyperFlex health check execution is through a scheduled run. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timestamp`:(string) Timestamp of the last health check execution on the HyperFlex cluster. 
 
