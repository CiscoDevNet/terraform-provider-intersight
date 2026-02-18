---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_io_card"
description: |-
        The equipment.IoCard (IOM) object represents an I/O Module, a critical component in a blade chassis that multiplexes and forwards network traffic between the blade servers and the upstream Fabric Interconnects.
        #### Purpose
        The IoCard object is the primary inventory and management entity for an IOM. This provides a comprehensive view of the module's identity, operational state, health, and its connection path to the parent Fabric Interconnect. It also serves as a container for its own sub-components, such as fan modules.
        #### Key Concepts
        - **Chassis Network Hub:** Models the central networking component within a chassis that connects servers to the fabric.
        - **Connectivity and Pathing:** The connectionPath property (e.g., A or B) and its relationship to a network.Element define its role and connection to the fabric.
        - **Health Monitoring:** Reports its overall operState and specific health issues (operReason), such as temperature warnings or low memory conditions.
        - **Hierarchical Inventory:** Acts as a parent container for its own components, like equipment.FanModule, and for the host-facing ports that connect to the servers.

---

# Data Source: intersight_equipment_io_card
The equipment.IoCard (IOM) object represents an I/O Module, a critical component in a blade chassis that multiplexes and forwards network traffic between the blade servers and the upstream Fabric Interconnects.
#### Purpose
The IoCard object is the primary inventory and management entity for an IOM. This provides a comprehensive view of the module's identity, operational state, health, and its connection path to the parent Fabric Interconnect. It also serves as a container for its own sub-components, such as fan modules.
#### Key Concepts
- **Chassis Network Hub:** Models the central networking component within a chassis that connects servers to the fabric.
- **Connectivity and Pathing:** The connectionPath property (e.g., "A" or "B") and its relationship to a network.Element define its role and connection to the fabric.
- **Health Monitoring:** Reports its overall operState and specific health issues (operReason), such as temperature warnings or low memory conditions.
- **Hierarchical Inventory:** Acts as a parent container for its own components, like equipment.FanModule, and for the host-facing ports that connect to the servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_io_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `connection_path`:(string) Switch Id to which the IOM is connected to. The value can be A or B. 
* `connection_status`:(string) Connectivity Status of FEX/IOM to Switch - A or B or AB. 
* `create_time`:(string) The time when this managed object was created. 
* `dc_supported`:(bool) IOM device connector support. 
* `description`:(string) This field is to provide description for the iocard module model. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `module_id`:(int) Module Identifier for the IO module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of IO card or fabric extender. 
* `part_number`:(string) Part Number identifier for the IO module. 
* `pid`:(string) This field identifies the Product ID for the IO module. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `product_name`:(string) This field identifies the Product Name for the iocard module model. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `side`:(string) Location of IOM within a chassis. The value can be left or right. 
* `sku`:(string) This field identifies the Stock Keeping Unit for the IO card module. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) This field identifies the version of the IO card module. 
* `vid`:(string) This field identifies the Vendor ID for the IO card module. 
 
