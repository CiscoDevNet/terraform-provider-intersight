---
subcategory: "power"
layout: "intersight"
page_title: "Intersight: intersight_power_control_state"
description: |-
  Managed object used to track chassis power capping information.
---

# Data Source: intersight_power_control_state
Managed object used to track chassis power capping information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_power_control_state.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocated_power`:(int) This field identifies the allocated power on the chassis in Watts. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `grid_max_power`:(int) This field identifies the available power when PSUs are in grid mode in Watts. 
* `max_required_power`:(int) This field identifies the maximum power required by the endpoint in Watts. 
* `min_required_power`:(int) This field identifies the minimum power required by the endpoint in Watts. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `n1_max_power`:(int) This field identifies the available power when PSUs are in N+1 mode in Watts. 
* `n2_max_power`:(int) This field identifies the available power when PSUs are in N+2 mode in Watts. 
* `non_redundant_max_power`:(int) This field identifies the available power when PSUs are in non-redundant mode in Watts. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 