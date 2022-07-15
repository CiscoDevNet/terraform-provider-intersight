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
* `allocated_budget`:(int) Sets the Allocated Power Budget of the Chassis (in Watts). This field is only supported for Cisco UCS X series Chassis. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dynamic_rebalancing`:(string) Sets the Dynamic Power Rebalancing mode of the Chassis. If enabled, this mode allows the chassis to dynamically reallocate the power between servers depending on their power usage. This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `extended_power_capacity`:(string) Sets the Extended Power Capacity of the Chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies.  This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `power_priority`:(string) Sets the Power Priority of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS X series servers.* `Low` - Set the Power Priority to Low.* `Medium` - Set the Power Priority to Medium.* `High` - Set the Power Priority to High. 
* `power_profiling`:(string) Sets the Power Profiling of the Server. If Enabled, this field allows the power manager to run power profiling  utility to determine the power needs of the server.  This field is only supported for Cisco UCS X series servers.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `power_restore_state`:(string) Sets the Power Restore State of the Server. In the absence of Intersight connectivity, the chassis will use this policy  to recover the host power after a power loss event.  This field is only supported for Cisco UCS X series servers.* `AlwaysOff` - Set the Power Restore Mode to Off.* `AlwaysOn` - Set the Power Restore Mode to On.* `LastState` - Set the Power Restore Mode to LastState. 
* `power_save_mode`:(string) Sets the Power Save mode of the Chassis. If the requested power budget is less than available power\u00a0capacity,  the additional PSUs not required to comply with redundancy policy are placed in Power Save mode. This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `redundancy_mode`:(string) Sets the Power Redundancy Mode of the Chassis.  Redundancy Mode determines the number of PSUs the chassis keeps as redundant.  N+2 mode is only supported for Cisco UCS X series Chassis.* `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis.* `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained.* `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy.* `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
