---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_shared_io_module"
description: |-
        The equipment.SharedIoModule object represents a shared I/O module (SIOM), which is an adapter-like component typically found in a System I/O Controller (SIOC) for Cisco S-Series storage servers. This provides the data path from the server to the Fabric Interconnects.
        #### Purpose
        
        The SharedIoModule object is used to inventory and manage the SIOM. It captures the module's identity, its operational state (operState), and its connectivity status (reachability) to the A and B sides of the fabric. It acts as a container for the network ports (port.Group objects).
        #### Key Concepts
        - **Storage Server Connectivity:** Models the specific I/O adapter used in S-Series servers.
        - **Redundancy and Reachability:** The reachability property indicates its connection status to both Fabric Interconnects, which is key for monitoring high availability.
        - **Hierarchical Inventory:** Serves as a parent for port.Group objects and is itself a child of an equipment.SystemIoController.
        - **Configuration State:** The configState property tracks its configuration status within the system.

---

# Data Source: intersight_equipment_shared_io_module
The equipment.SharedIoModule object represents a shared I/O module (SIOM), which is an adapter-like component typically found in a System I/O Controller (SIOC) for Cisco S-Series storage servers. This provides the data path from the server to the Fabric Interconnects.
#### Purpose
     
The SharedIoModule object is used to inventory and manage the SIOM. It captures the module's identity, its operational state (operState), and its connectivity status (reachability) to the A and B sides of the fabric. It acts as a container for the network ports (port.Group objects).
#### Key Concepts
- **Storage Server Connectivity:** Models the specific I/O adapter used in S-Series servers.
- **Redundancy and Reachability:** The reachability property indicates its connection status to both Fabric Interconnects, which is key for monitoring high availability.
- **Hierarchical Inventory:** Serves as a parent for port.Group objects and is itself a child of an equipment.SystemIoController.
- **Configuration State:** The configState property tracks its configuration status within the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_shared_io_module.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_state`:(string) This field identifies the configuration state for this SIOM Unit. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `discovery`:(string) This field identifies the discovery state of SIOM. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mac_of_shared_iom_aside`:(string) This field identifies the MAC of IOM-A side. 
* `mac_of_shared_iom_bside`:(string) This field identifies the MAC of IOM-B side. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) This field identifies the SIOM operational state. 
* `part_number`:(string) This field identifies the Part Number for this SIOM Unit. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `reachability`:(string) This field identifies the reachability to FI-A and B side. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usr_lbl`:(string) User label configured for the SIOM. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vid`:(string) This field identifies the vendor id for this SIOM Unit. 
 
