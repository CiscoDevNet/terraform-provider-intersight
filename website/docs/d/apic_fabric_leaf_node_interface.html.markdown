---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_fabric_leaf_node_interface"
description: |-
        FabricLeafNodeInterface object has been deprecated. The FabricLeafNodeInterface object in Cisco APIC defines the interfaces associated with fabric leaf nodes, focusing on connectivity and communication.
        #### Purpose
        Fabric leaf node interfaces provide the necessary ports and connections for network devices, enabling seamless communication and data exchange.
        #### Key Concepts
        - **Connectivity Points:** Represents the interfaces that connect network devices to the fabric leaf nodes.
        - **Communication Management:** Supports the configuration and management of network interfaces for optimal data flow.
        - **Policy Application:** Ensures that network policies are applied consistently across leaf node interfaces.

---

# Data Source: intersight_apic_fabric_leaf_node_interface
FabricLeafNodeInterface object has been deprecated. The FabricLeafNodeInterface object in Cisco APIC defines the interfaces associated with fabric leaf nodes, focusing on connectivity and communication.
#### Purpose
Fabric leaf node interfaces provide the necessary ports and connections for network devices, enabling seamless communication and data exchange.
#### Key Concepts
- **Connectivity Points:** Represents the interfaces that connect network devices to the fabric leaf nodes.
- **Communication Management:** Supports the configuration and management of network interfaces for optimal data flow.
- **Policy Application:** Ensures that network policies are applied consistently across leaf node interfaces.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_fabric_leaf_node_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_leaf_node_dn`:(string) Fabric leaf node Distinguished Name (DN). 
* `fabric_leaf_node_id`:(string) Fabric leaf node identification number. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Object name in Cisco Application Policy Infrastructure Controller (APIC). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
