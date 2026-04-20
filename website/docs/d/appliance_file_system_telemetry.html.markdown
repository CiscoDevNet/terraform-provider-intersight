---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_system_telemetry"
description: |-
        Calculated operational status of a file system on a node.

---

# Data Source: intersight_appliance_file_system_telemetry
Calculated operational status of a file system on a node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_file_system_telemetry.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `disk_order`:(string) Symbolic order identifier of the physical disk from /dev/disk/by-order (e.g., disk1, disk2).This provides a stable, human-readable identifier for the disk that persists across reboots,unlike device names which may change. Currently, the relationship between filesystems to disksis one-to-one and this is unlikely to change in the future. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `instance`:(string) The instance identifier from which the metrics were collected. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mountpoint`:(string) Mount point of this file system. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage`:(string) Operational status of filesystem utilization.Values are Optimal, Warning, or Critical based on configured thresholds.* `Optimal` - Resource usage is within normal operating parameters.* `Warning` - Resource usage has exceeded warning thresholds and attention may be needed.* `Critical` - Resource usage has exceeded critical thresholds and immediate attention is required. 
 
