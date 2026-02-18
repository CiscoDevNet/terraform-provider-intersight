---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_device_boot_mode"
description: |-
        The boot.DeviceBootMode object represents the user-configured BIOS boot mode for a server, as defined in its boot policy. It reflects the intended state that should be applied to the server's BIOS.
        #### Purpose
        The primary purpose of this object is to inventory the desired boot mode (e.g., Legacy or Uefi) from a server's configuration policy. This allows for a clear distinction and comparison between the intended configuration (configuredBootMode) and the actual, running state of the server, which is reported by the BIOS.BootMode object.
        #### Key Concepts
        - **Desired State Representation:** Captures the configuredBootMode from a boot policy, representing what the boot mode should be.
        - **Policy Verification:** Facilitates auditing and compliance checks by allowing a direct comparison between the configured and actual boot modes.
        - **Configuration-Centric:** Unlike BIOS.BootMode, which reports the hardware's state, this object reflects the state defined in a management policy.
        - **Server Association:** Directly linked to a compute.Physical object, ensuring the configured boot mode is tied to a specific server.

---

# Data Source: intersight_boot_device_boot_mode
The boot.DeviceBootMode object represents the user-configured BIOS boot mode for a server, as defined in its boot policy. It reflects the intended state that should be applied to the server's BIOS.
#### Purpose
The primary purpose of this object is to inventory the desired boot mode (e.g., "Legacy" or "Uefi") from a server's configuration policy. This allows for a clear distinction and comparison between the intended configuration (configuredBootMode) and the actual, running state of the server, which is reported by the BIOS.BootMode object.
#### Key Concepts
- **Desired State Representation:** Captures the configuredBootMode from a boot policy, representing what the boot mode should be.
- **Policy Verification:** Facilitates auditing and compliance checks by allowing a direct comparison between the configured and actual boot modes.
- **Configuration-Centric:** Unlike BIOS.BootMode, which reports the hardware's state, this object reflects the state defined in a management policy.
- **Server Association:** Directly linked to a compute.Physical object, ensuring the configured boot mode is tied to a specific server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_device_boot_mode.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `configured_boot_mode`:(string) The user desired BIOS boot mode as configured in the boot policy. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
