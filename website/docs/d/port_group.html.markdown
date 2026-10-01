---
subcategory: "port"
layout: "intersight"
page_title: "Intersight: intersight_port_group"
description: |-
        Groups (port.Group) are container objects that group multiple physical interfaces under a common switching context. A Fabric Interconnect switch card (and some modular components such as shared I/O modules) can have one or more port groups, typically separated by transport type (Ethernet vs Fibre Channel).
        #### Purpose
        Provide a hierarchical inventory structure for switch interfaces so ports can be organized, discovered, and queried as cohesive sets (e.g., all Ethernet ports vs all FC ports on a given module).
        #### Key Concepts
        - **Port container:** A single Group aggregates multiple ports of a given transport type.
        - **Transport segregation:** The `transport` field distinguishes Ethernet vs Fibre Channel grouping.
        - **Hierarchy anchor:** Acts as the parent for `ether.PhysicalPort` and `fc.PhysicalPort` collections, and for breakout `port.SubGroup` collections where applicable.
        - **Module context:** Inherits permissions from the parent hardware context (e.g., switch card / shared IO module), aligning visibility with the owning device.

---

# Data Source: intersight_port_group
Groups (port.Group) are container objects that group multiple physical interfaces under a common switching context. A Fabric Interconnect switch card (and some modular components such as shared I/O modules) can have one or more port groups, typically separated by transport type (Ethernet vs Fibre Channel).
#### Purpose
Provide a hierarchical inventory structure for switch interfaces so ports can be organized, discovered, and queried as cohesive sets (e.g., all Ethernet ports vs all FC ports on a given module).
#### Key Concepts
- **Port container:** A single Group aggregates multiple ports of a given transport type.
- **Transport segregation:** The `transport` field distinguishes Ethernet vs Fibre Channel grouping.
- **Hierarchy anchor:** Acts as the parent for `ether.PhysicalPort` and `fc.PhysicalPort` collections, and for breakout `port.SubGroup` collections where applicable.
- **Module context:** Inherits permissions from the parent hardware context (e.g., switch card / shared IO module), aligning visibility with the owning device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_port_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `transport`:(string) Type of port group. Values are Eth or Fc. 
 
