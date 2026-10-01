---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_uplink_role"
description: |-
        The FcUplinkRole object represents configuration intent for a Fibre Channel uplink port in a port policy.
        #### Purpose
        FcUplinkRole defines a port’s role as an FC uplink and expresses FC-uplink-specific configuration intent such as VSAN association and FC port behavior, enabling consistent SAN uplink configuration.
        #### Key Concepts
        - **FC uplink intent:** Declares a port is used for SAN uplink connectivity.
        - **VSAN association:** Represents SAN segmentation alignment for uplink ports.
        - **Policy-driven SAN configuration:** Enables consistent FC uplink deployment across profiles.
        - **Validation readiness:** Intended to be validated against FC policy and platform constraints.

---

# Data Source: intersight_fabric_fc_uplink_role
The FcUplinkRole object represents configuration intent for a Fibre Channel uplink port in a port policy.
#### Purpose
FcUplinkRole defines a port’s role as an FC uplink and expresses FC-uplink-specific configuration intent such as VSAN association and FC port behavior, enabling consistent SAN uplink configuration.
#### Key Concepts
- **FC uplink intent:** Declares a port is used for SAN uplink connectivity.
- **VSAN association:** Represents SAN segmentation alignment for uplink ports.
- **Policy-driven SAN configuration:** Enables consistent FC uplink deployment across profiles.
- **Validation readiness:** Intended to be validated against FC policy and platform constraints.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_fc_uplink_role.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_speed`:(string) Admin configured speed for the port.* `16Gbps` - Admin configurable speed 16Gbps.* `8Gbps` - Admin configurable speed 8Gbps.* `32Gbps` - Admin configurable speed 32Gbps.* `64Gbps` - Admin configurable speed 64Gbps.* `Auto` - Admin configurable speed AUTO ( default ). 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fill_pattern`:(string) Fill pattern to differentiate the configs in NPIV.* `Idle` - Fc Fill Pattern type Idle.* `Arbff` - Fc Fill Pattern type Arbff. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
* `user_label`:(string) The user defined label assigned to a Port. 
* `vsan_id`:(int) Virtual San Identifier associated to the FC port. 
 
