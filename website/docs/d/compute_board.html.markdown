---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_board"
description: |-
        The compute.Board object represents the main motherboard of a server. It serves as a central container for the core components of the server, such as CPUs, memory, and storage controllers, and reports on its own operational health.
        #### Purpose
        The primary purpose of the compute.Board object is to model the server's motherboard as a distinct inventory entity. It aggregates the most critical hardware components and provides a focal point for monitoring the board's health, including its power state and any temperature or voltage issues.
        #### Key Concepts
        - **Central Component Hub:** Acts as the parent object for essential hardware like processor.Unit, memory.Array, and storage.Controller, providing a logical grouping for the server's core components.
        - **Health Monitoring:** Reports on its operational status, including power state (operPowerState) and specific health issues (operReason) like temperature or voltage warnings.
        - **Unique Identification:** Identified by a boardId within the server, ensuring it can be uniquely referenced.
        - **Hierarchical Inventory:** Is a child of a compute.RackUnit or compute.Blade, linking it directly to the server it belongs to.

---

# Data Source: intersight_compute_board
The compute.Board object represents the main motherboard of a server. It serves as a central container for the core components of the server, such as CPUs, memory, and storage controllers, and reports on its own operational health.
#### Purpose
The primary purpose of the compute.Board object is to model the server's motherboard as a distinct inventory entity. It aggregates the most critical hardware components and provides a focal point for monitoring the board's health, including its power state and any temperature or voltage issues.
#### Key Concepts
- **Central Component Hub:** Acts as the parent object for essential hardware like processor.Unit, memory.Array, and storage.Controller, providing a logical grouping for the server's core components.
- **Health Monitoring:** Reports on its operational status, including power state (operPowerState) and specific health issues (operReason) like temperature or voltage warnings.
- **Unique Identification:** Identified by a boardId within the server, ensuring it can be uniquely referenced.
- **Hierarchical Inventory:** Is a child of a compute.RackUnit or compute.Blade, linking it directly to the server it belongs to.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_board.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `board_id`:(int) Unique identifier of the mother board present in the server. 
* `cpu_type_controller`:(string) The type of central processing unit on the mother board. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_power_state`:(string) Current power state of the mother board of the server. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
