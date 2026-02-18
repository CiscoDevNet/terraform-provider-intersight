---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_unit"
description: |-
        The bios.Unit object represents the BIOS (Basic Input/Output System) of a server. It acts as a container for BIOS-related inventory and configuration objects, such as the running firmware version and the system boot order.
        #### Purpose
        The primary purpose of the bios.Unit object is to serve as a central anchor point for all BIOS-related information within the server's inventory. It aggregates critical components of the BIOS configuration, making them accessible for monitoring and management.
        #### Key Concepts
        - **Central BIOS Anchor:** Provides a single, identifiable object representing the server's BIOS.
        - **Firmware Tracking:** Links to firmware.RunningFirmware objects to provide the current, detailed version of the BIOS firmware.
        - **Boot Order Management:** Contains the bios.SystemBootOrder object, which details the actual boot sequence of the server as configured in the BIOS.
        - **Hierarchical Inventory:** Exists within the context of a compute.Blade or compute.RackUnit, tying the BIOS information directly to the physical server.

---

# Data Source: intersight_bios_unit
The bios.Unit object represents the BIOS (Basic Input/Output System) of a server. It acts as a container for BIOS-related inventory and configuration objects, such as the running firmware version and the system boot order.
#### Purpose
The primary purpose of the bios.Unit object is to serve as a central anchor point for all BIOS-related information within the server's inventory. It aggregates critical components of the BIOS configuration, making them accessible for monitoring and management.
#### Key Concepts
- **Central BIOS Anchor:** Provides a single, identifiable object representing the server's BIOS.
- **Firmware Tracking:** Links to firmware.RunningFirmware objects to provide the current, detailed version of the BIOS firmware.
- **Boot Order Management:** Contains the bios.SystemBootOrder object, which details the actual boot sequence of the server as configured in the BIOS.
- **Hierarchical Inventory:** Exists within the context of a compute.Blade or compute.RackUnit, tying the BIOS information directly to the physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_unit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `init_seq`:(string) The initSeq of the equipment. 
* `init_ts`:(string) The initTs of the equipment. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
