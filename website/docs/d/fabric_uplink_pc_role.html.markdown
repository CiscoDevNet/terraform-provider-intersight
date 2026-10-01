---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_uplink_pc_role"
description: |-
        The UplinkPcRole object represents configuration intent for an Ethernet uplink port-channel in a port policy.
        #### Purpose
        UplinkPcRole models an uplink built from multiple member ports. This provides the structure for defining port-channel intent, applying uplink-related policy attachments, and enabling consistent deployment and validation of aggregated uplinks.
        #### Key Concepts
        - **Aggregated uplink intent:** Represents an uplink composed of multiple physical ports.
        - **Policy attachment for port-channels:** Provides an anchor for link aggregation, flow control, link control, VLAN allow-lists, and MACsec associations.
        - **Scale and resiliency modeling:** Enables consistent modeling of redundant/high-bandwidth uplinks.
        - **Policy-scoped identity:** Identifies the port-channel within a port policy context.

---

# Data Source: intersight_fabric_uplink_pc_role
The UplinkPcRole object represents configuration intent for an Ethernet uplink port-channel in a port policy.
#### Purpose
UplinkPcRole models an uplink built from multiple member ports. This provides the structure for defining port-channel intent, applying uplink-related policy attachments, and enabling consistent deployment and validation of aggregated uplinks.
#### Key Concepts
- **Aggregated uplink intent:** Represents an uplink composed of multiple physical ports.
- **Policy attachment for port-channels:** Provides an anchor for link aggregation, flow control, link control, VLAN allow-lists, and MACsec associations.
- **Scale and resiliency modeling:** Enables consistent modeling of redundant/high-bandwidth uplinks.
- **Policy-scoped identity:** Identifies the port-channel within a port policy context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_uplink_pc_role.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_speed`:(string) Admin configured speed for the port.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `50Gbps` - Admin configurable speed 50Gbps.* `100Gbps` - Admin configurable speed 100Gbps.* `400Gbps` - Admin configurable speed 400Gbps.* `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fec`:(string) Forward error correction configuration for Uplink Port Channel member ports.* `Auto` - Forward error correction option 'Auto'. Supported speeds are Auto, 1Gbps, 10Gbps, 25Gbps, 40Gbps and 100 Gbps.* `Cl91` - Forward error correction option 'cl91'. Supported speeds are 25Gbps and 100 Gbps. This FEC setting is not supported on Unified Edge platforms.* `Cl74` - Forward error correction option 'cl74'. Supported speeds are 25Gbps.* `rs-cons16` - Forward error correction option \ rs-cons16\ . Supported speeds are 25Gbps. This FEC setting is not supported on Unified Edge platforms.* `rs-ieee` - Forward error correction option \ rs-ieee\ . Supported speeds are 25Gbps.* `Off` - Turn off forward error correction. Supported speeds are 25Gbps and 100Gbps. Additionally, 1Gbps is supported on Unified Edge platforms. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_label`:(string) The user defined label assigned to the a Port. 
 
