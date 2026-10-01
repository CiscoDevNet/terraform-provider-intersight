---
subcategory: "port"
layout: "intersight"
page_title: "Intersight: intersight_port_sub_group"
description: |-
        SubGroups (port.SubGroup) represent break-out port groupings under a parent port group. A SubGroup corresponds to an aggregate port (for example, a breakout of port 1/9) and holds the individual breakout member ports created from that aggregate.
        #### Purpose
        Model breakout topology so users and automation can understand how an aggregate port is subdivided into multiple physical member ports and how those member ports map to Ethernet, FC, or host-port inventory.
        #### Key Concepts
        - **Breakout representation:** A SubGroup corresponds to one aggregate port and contains the resulting breakout member ports.
        - **Aggregate port identity:** The `aggregatePortId` identifies the parent/aggregate port being broken out.
        - **Transport awareness:** `transport` indicates whether the subgroup is Ethernet or Fibre Channel.
        - **Mixed port membership:** Can contain collections of Ethernet physical ports, FC physical ports, and (for fabric extenders/IOMs) breakout-generated host ports.
        - **Hierarchical correlation:** Provides the structured intermediate layer between a port group and the final per-port objects created by breakout.

---

# Data Source: intersight_port_sub_group
SubGroups (port.SubGroup) represent break-out port groupings under a parent port group. A SubGroup corresponds to an aggregate port (for example, a breakout of port 1/9) and holds the individual breakout member ports created from that aggregate.
#### Purpose
Model breakout topology so users and automation can understand how an aggregate port is subdivided into multiple physical member ports and how those member ports map to Ethernet, FC, or host-port inventory.
#### Key Concepts
- **Breakout representation:** A SubGroup corresponds to one aggregate port and contains the resulting breakout member ports.
- **Aggregate port identity:** The `aggregatePortId` identifies the parent/aggregate port being broken out.
- **Transport awareness:** `transport` indicates whether the subgroup is Ethernet or Fibre Channel.
- **Mixed port membership:** Can contain collections of Ethernet physical ports, FC physical ports, and (for fabric extenders/IOMs) breakout-generated host ports.
- **Hierarchical correlation:** Provides the structured intermediate layer between a port group and the final per-port objects created by breakout.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_port_sub_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Breakout port member in the Fabric Interconnect. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `transport`:(string) Type of port sub-group. Values are Eth or Fc. 
 
