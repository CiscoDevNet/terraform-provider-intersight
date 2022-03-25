---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_dashboard_memory_details"
description: |-
        Details of Nexus Dashboard's memory.

---

# Data Source: intersight_niatelemetry_nexus_dashboard_memory_details
Details of Nexus Dashboard's memory.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboard_memory_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_name`:(string) Name of the node in Nexus Dashboard cluster. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `memory_capacity`:(int) Memory capacity of a node in Nexus Dashboard. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
