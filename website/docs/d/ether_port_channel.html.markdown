---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_port_channel"
description: |-
        PortChannels are logical Ethernet interfaces on a Fabric Interconnect formed by aggregating multiple PhysicalPorts into a single bundled link. They increase effective bandwidth, provide load balancing, and improve resiliency through redundancy across member links.
        #### Purpose
        Represent aggregated Ethernet connectivity on an FI so consumers can view VLAN/trunking behavior, role and mode, administrative/operational state, and L2/L3 attributes of the port-channel as a single managed interface.
        #### Key Concepts
        - **Link aggregation abstraction**: A PortChannel is a single logical interface backed by multiple physical member ports to improve throughput and availability.
        - **Segmentation/trunking context**: `allowedVlans`, `accessVlan`, and `nativeVlan` capture how VLANs are carried on the aggregated link.
        - **Role-based usage**: `role` indicates intended function (for example uplink vs server-facing), while `mode` reflects how the port-channel operates.
        - **State and diagnostics**: `adminState`, `operState`, and `operStateQual` provide configured enablement, runtime state, and the reason/qualifier for that state.
        - **Performance and MTU**: `operSpeed`, `bandWidth`, and `mtu` summarize the effective characteristics of the aggregated interface.
        - **Identity and addressing**: `portChannelId` uniquely identifies the bundle on the FI; `macAddress`, `ipAddress`, `ipAddressMask`, and `ipv6SubnetCidr` support L2/L3 management and correlation.
        - **Operator-friendly metadata**: `name`, `description`, `status`, and `userLabel` support inventory reporting and operational workflows.

---

# Data Source: intersight_ether_port_channel
PortChannels are logical Ethernet interfaces on a Fabric Interconnect formed by aggregating multiple PhysicalPorts into a single bundled link. They increase effective bandwidth, provide load balancing, and improve resiliency through redundancy across member links.
#### Purpose
Represent aggregated Ethernet connectivity on an FI so consumers can view VLAN/trunking behavior, role and mode, administrative/operational state, and L2/L3 attributes of the port-channel as a single managed interface.
#### Key Concepts
- **Link aggregation abstraction**: A PortChannel is a single logical interface backed by multiple physical member ports to improve throughput and availability.
- **Segmentation/trunking context**: `allowedVlans`, `accessVlan`, and `nativeVlan` capture how VLANs are carried on the aggregated link.
- **Role-based usage**: `role` indicates intended function (for example uplink vs server-facing), while `mode` reflects how the port-channel operates.
- **State and diagnostics**: `adminState`, `operState`, and `operStateQual` provide configured enablement, runtime state, and the reason/qualifier for that state.
- **Performance and MTU**: `operSpeed`, `bandWidth`, and `mtu` summarize the effective characteristics of the aggregated interface.
- **Identity and addressing**: `portChannelId` uniquely identifies the bundle on the FI; `macAddress`, `ipAddress`, `ipAddressMask`, and `ipv6SubnetCidr` support L2/L3 management and correlation.
- **Operator-friendly metadata**: `name`, `description`, `status`, and `userLabel` support inventory reporting and operational workflows.
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
* `oper_vlans`:(string) Operational VLANs on this port. 
* `port_channel_id`:(int) Unique identifier for this port-channel on the FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) This port-channel's configured role (uplink, server, etc.). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Detailed status of this port-channel. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `user_label`:(string) The user defined label assigned to the port channel. 
 
