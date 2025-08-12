---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_server_power_parameters"
description: |-
        Managed object used to track server power parameters information.

---

# Data Source: intersight_compute_server_power_parameters
Managed object used to track server power parameters information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_server_power_parameters.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `power_allocation`:(int) This field identifies the maximum power that has been allocated to the blade by CMC in Watts. Power budget for the chassis is configured by the power policy. That budget is then divided among the blades in the chassis by CMC. 
* `power_priority`:(string) Power Priority level of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS B series and X series servers.* `Unknown` - Power allocation priority of server is either unknown or not supported on CMC firmware version.* `Low` - Power allocation priority of server is low.* `Medium` - Power allocation priority of server is medium.* `High` - Power allocation priority of server is high. 
* `power_profiling`:(string) Status of Power Profiling setting of the Server. If Enabled, this field allows the power manager to run power profiling utility to determine the power needs of the server. This field is only supported for Cisco UCS X series servers.* `Unknown` - Power Profiling state is either unknown for the server or not supported on BMC firmware version.* `Enabled` - Power Profiling is enabled for the server.* `Disabled` - Power Profiling is either disabled for the server or not supported on BMC firmware version. 
* `power_restore`:(string) Value of the power restore policy for the server. In the absence of Intersight connectivity, the server will use this state to recover the host power after a power loss event.* `Unknown` - Power restore state for the server is either Unknown or not supported on BMC firmware version.* `Always On` - Power restore state for server is set to Always On.* `Always Off` - Power restore state for server is set to Always Off.* `Last State` - Power restore state for server is set to Last State. 
* `processor_package_power_limit`:(string) Processor Package Power Limit (PPL) of a server. PPL refers to the amount of power that a CPU can draw from the power supply. The Processor Package Power Limit (PPL) feature is currently available exclusively on Cisco UCS C225/C245 M8 servers.* `Unknown` - Processor package power limit is either unknown for the server or not supported on BMC firmware version.* `Default` - Processor package power limit is default for the server.* `Maximum` - Processor package power limit is maximum for the server.* `Minimum` - Processor package power limit is minimum for the server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
