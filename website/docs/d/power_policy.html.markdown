---
subcategory: "power"
layout: "intersight"
page_title: "Intersight: intersight_power_policy"
description: |-
  Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.
---

# Data Source: intersight_power_policy
Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_power_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocated_budget`:(int) Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dynamic_rebalancing`:(string) Sets the Dynamic Power Rebalancing of the System. This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `power_priority`:(string) Sets the Power Priority of the System. This field is only supported for Cisco UCS X series servers.* `Low` - Set the Power Priority to Low.* `Medium` - Set the Power Priority to Medium.* `High` - Set the Power Priority to High. 
* `power_profiling`:(string) Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `power_restore_state`:(string) Sets the Power Restore State of the Server. This field is only supported for Cisco UCS X series servers.* `AlwaysOff` - Set the Power Restore Mode to Off.* `AlwaysOn` - Set the Power Restore Mode to On.* `LastState` - Set the Power Restore Mode to LastState. 
* `power_save_mode`:(string) Sets the Power Save mode of the System. This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `redundancy_mode`:(string) Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis.* `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis.* `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained.* `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy.* `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 