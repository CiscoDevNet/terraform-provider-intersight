---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_port_channel"
description: |-
        Model contains the details of the ethernet port-channels configured on the FI.

---

# Data Source: intersight_ether_port_channel
Model contains the details of the ethernet port-channels configured on the FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ether_port_channel.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_vlan`:(string) Access VLANs for this port-channel, on this FI. 
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port-channel. 
* `allowed_vlans`:(string) Allowed VLANs on this port-channel, on this FI. 
* `band_width`:(string) Bandwidth of this port-channel. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of this port-channel. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ip_address`:(string) IP address of this port-channel. 
* `ip_address_mask`:(int) IP address mask of this port-channel. 
* `ipv6_subnet_cidr`:(string) IPv6 subnet in CIDR notation of this port-channel. Ex. 2000::/8. 
* `mac_address`:(string) MAC address of this port-channel. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) Operating mode of this port-channel. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit of this port-channel. 
* `name`:(string) Name of the port channel. 
* `native_vlan`:(string) Native VLAN for this port-channel, on this FI. 
* `oper_speed`:(string) Operational speed of this port-channel. 
* `oper_state`:(string) Operational state of this port-channel. 
* `oper_state_qual`:(string) Reason for this port-channel's Operational state. 
* `port_channel_id`:(int) Unique identifier for this port-channel on the FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) This port-channel's configured role (uplink, server, etc.). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Detailed status of this port-channel. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
 
