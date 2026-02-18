---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_io_expander"
description: |-
        The equipment.IoExpander object represents an I/O expander card, which serves as an extension for servers, particularly within a Colusa-type chassis architecture.
        #### Purpose
        The primary purpose of this object is to inventory and report the status of an I/O expander card. These cards provide additional I/O capabilities or connectivity options for the servers they are attached to. The object captures the expander's identity and its operational state.
        #### Key Concepts
        - **I/O Extension:** Models a card that expands the I/O capabilities of a server.
        - **State Reporting:** The operState property provides the current operational status of the expander card.
        - **Hardware Inventory:** Captures the base identity of the card, including its model, vendor, and serial number.
        - **Server Component:** It is a child of a compute.Blade, linking it directly to the server it expands.

---

# Data Source: intersight_equipment_io_expander
The equipment.IoExpander object represents an I/O expander card, which serves as an extension for servers, particularly within a Colusa-type chassis architecture.
#### Purpose
The primary purpose of this object is to inventory and report the status of an I/O expander card. These cards provide additional I/O capabilities or connectivity options for the servers they are attached to. The object captures the expander's identity and its operational state.
#### Key Concepts
- **I/O Extension:** Models a card that expands the I/O capabilities of a server.
- **State Reporting:** The operState property provides the current operational status of the expander card.
- **Hardware Inventory:** Captures the base identity of the card, including its model, vendor, and serial number.
- **Server Component:** It is a child of a compute.Blade, linking it directly to the server it expands.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_io_expander.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Identifies the operational state of I/O expander. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
