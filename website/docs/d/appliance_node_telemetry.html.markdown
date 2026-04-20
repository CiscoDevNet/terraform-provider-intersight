---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_node_telemetry"
description: |-
        NodeTelemetry represents the operational status of an Assist node,
        including CPU, memory, and disk usage metrics collected from Prometheus.

---

# Data Source: intersight_appliance_node_telemetry
NodeTelemetry represents the operational status of an Assist node,
including CPU, memory, and disk usage metrics collected from Prometheus.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_node_telemetry.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cpu_usage`:(string) CpuUsage indicates the operational status of CPU utilization on the Assist node.Values are Optimal, Warning, or Critical based on configured thresholds.* `Optimal` - Resource usage is within normal operating parameters.* `Warning` - Resource usage has exceeded warning thresholds and attention may be needed.* `Critical` - Resource usage has exceeded critical thresholds and immediate attention is required. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `instance`:(string) The instance identifier from which the metrics were collected. 
* `memory_usage`:(string) MemoryUsage indicates the operational status of memory utilization on the Assist node.Values are Optimal, Warning, or Critical based on configured thresholds.* `Optimal` - Resource usage is within normal operating parameters.* `Warning` - Resource usage has exceeded warning thresholds and attention may be needed.* `Critical` - Resource usage has exceeded critical thresholds and immediate attention is required. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
