---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_boot_mode"
description: |-
        The bios.BootMode object represents the configured boot mode of a server as reported by its BIOS. This provides a clear distinction between the actual boot mode used by the platform and the desired mode set in a policy.
        #### Purpose
        The main function of this object is to provide an accurate, read-only inventory of the server's current boot mode (e.g., UEFI or Legacy). This allows administrators and automation systems to verify that the server is operating in the expected boot mode and to detect any discrepancies with configured policies.
        #### Key Concepts
        - **Actual State Inventory:** Captures the actualBootMode as reported directly by the server's BIOS, providing a ground-truth value.
        - **Configuration Verification:** Enables comparison between the actual state and the desired state defined in a boot policy.
        - **Read-Only:** Serves as a reporting mechanism rather than a configuration object, ensuring data integrity.
        - **Server-Specific:** Directly associated with a compute.Blade or compute.RackUnit, providing boot mode information for a specific physical server.

---

# Data Source: intersight_bios_boot_mode
The bios.BootMode object represents the configured boot mode of a server as reported by its BIOS. This provides a clear distinction between the actual boot mode used by the platform and the desired mode set in a policy.
#### Purpose
The main function of this object is to provide an accurate, read-only inventory of the server's current boot mode (e.g., UEFI or Legacy). This allows administrators and automation systems to verify that the server is operating in the expected boot mode and to detect any discrepancies with configured policies.
#### Key Concepts
- **Actual State Inventory:** Captures the actualBootMode as reported directly by the server's BIOS, providing a ground-truth value.
- **Configuration Verification:** Enables comparison between the actual state and the desired state defined in a boot policy.
- **Read-Only:** Serves as a reporting mechanism rather than a configuration object, ensuring data integrity.
- **Server-Specific:** Directly associated with a compute.Blade or compute.RackUnit, providing boot mode information for a specific physical server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_boot_mode.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `actual_boot_mode`:(string) The actual BIOS boot mode as reported by the platform BIOS. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
