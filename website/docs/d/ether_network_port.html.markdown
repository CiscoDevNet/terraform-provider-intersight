---
subcategory: "ether"
layout: "intersight"
page_title: "Intersight: intersight_ether_network_port"
description: |-
        The ether.NetworkPort object represents a network-facing port on a Fabric Extender (FEX) or I/O Module (IOM). These ports, also known as Network Interfaces (NIFs), provide the uplink connection from the FEX/IOM to the parent Fabric Interconnect.
        #### Purpose
        The NetworkPort object is used to model and inventory the physical uplink ports of a FEX or IOM. It captures the port's identity, its operational speed, and its connection status to the parent switch. It is essential for monitoring the health and bandwidth of the connection between the fabric extender and the core fabric.
        #### Key Concepts
        - **Uplink Port Model:** Represents a physical port used for uplinking a FEX or IOM to a Fabric Interconnect.
        - **Connectivity Status:** The operReason property provides specific details on the link's status, such as LinkMissing or LinkMisconnect.

---

# Data Source: intersight_ether_network_port
The ether.NetworkPort object represents a network-facing port on a Fabric Extender (FEX) or I/O Module (IOM). These ports, also known as Network Interfaces (NIFs), provide the uplink connection from the FEX/IOM to the parent Fabric Interconnect.
#### Purpose
The NetworkPort object is used to model and inventory the physical uplink ports of a FEX or IOM. It captures the port's identity, its operational speed, and its connection status to the parent switch. It is essential for monitoring the health and bandwidth of the connection between the fabric extender and the core fabric.
#### Key Concepts
- **Uplink Port Model:** Represents a physical port used for uplinking a FEX or IOM to a Fabric Interconnect.
- **Connectivity Status:** The operReason property provides specific details on the link's status, such as LinkMissing or LinkMisconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ether_network_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `module_id`:(int) Febric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of an Interface. 
* `peer_dn`:(string) Peer DN for network host port of fabric extender. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Network Port operational speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
 
