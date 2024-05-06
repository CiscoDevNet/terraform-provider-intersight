---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_control_policy"
description: |-
        A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.

---

# Data Source: intersight_fabric_switch_control_policy
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_control_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ethernet_switching_mode`:(string) Enable or Disable Ethernet End Host Switching Mode.* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fabric_pc_vhba_reset`:(string) When enabled, a Registered State Change Notification (RSCN) is sent to the VIC adapter when any member port within the fabric port-channel goes down and vHBA would reset to restore the connection immediately. When disabled (default), vHBA reset is done only when all the members of a fabric port-channel are down.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `fc_switching_mode`:(string) Enable or Disable FC End Host Switching Mode.* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `reserved_vlan_start_id`:(int) The starting ID for VLANs reserved for internal use within the Fabric Interconnect. This VLAN ID is the starting ID ofa contiguous block of 128 VLANs that cannot be configured for user data.  This range of VLANs cannot be configured inVLAN policy.If this property is not configured, VLAN range 3915 - 4042 is reserved for internal use by default. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vlan_port_optimization_enabled`:(bool) To enable or disable the VLAN port count optimization. This feature will always be enabled for Cisco UCS Fabric Interconnect 9108 100G. 
 
