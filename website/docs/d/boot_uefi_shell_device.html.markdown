---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_uefi_shell_device"
description: |-
        The boot.UefiShellDevice object represents the UEFI (Unified Extensible Firmware Interface) Shell as a configurable boot option within a server's boot policy.
        #### Purpose
        The UefiShellDevice object allows administrators to include the UEFI Shell as a target in the server's boot sequence. The UEFI Shell provides a command-line environment for scripting, diagnostics, and manually launching UEFI applications, which can be a valuable tool for advanced troubleshooting and system maintenance before an operating system loads.
        #### Key Concepts
        - **Diagnostic Environment:** Defines the built-in UEFI Shell as a bootable target.
        - **Configurable Priority:** The order property determines its position in the boot sequence, allowing it to be prioritized for maintenance tasks.
        - **Enable/Disable Control:** The state property allows the UEFI Shell to be easily enabled or disabled as a boot option within a policy.
        - **Boot Policy Element:** Functions as a component within a boot policy that is associated with a compute.Physical server object.

---

# Data Source: intersight_boot_uefi_shell_device
The boot.UefiShellDevice object represents the UEFI (Unified Extensible Firmware Interface) Shell as a configurable boot option within a server's boot policy.
#### Purpose
The UefiShellDevice object allows administrators to include the UEFI Shell as a target in the server's boot sequence. The UEFI Shell provides a command-line environment for scripting, diagnostics, and manually launching UEFI applications, which can be a valuable tool for advanced troubleshooting and system maintenance before an operating system loads.
#### Key Concepts
- **Diagnostic Environment:** Defines the built-in UEFI Shell as a bootable target.
- **Configurable Priority:** The order property determines its position in the boot sequence, allowing it to be prioritized for maintenance tasks.
- **Enable/Disable Control:** The state property allows the UEFI Shell to be easily enabled or disabled as a boot option within a policy.
- **Boot Policy Element:** Functions as a component within a boot policy that is associated with a compute.Physical server object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_uefi_shell_device.<custom_name>.results[i].<propertyname>`.
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
* `name`:(string) The name of the boot device configured in the boot policy. 
* `order`:(int) The order of the boot device configured in the boot policy. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the boot device configured in the boot policy. 
* `type`:(string) The type of the boot device configured in the boot policy. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
