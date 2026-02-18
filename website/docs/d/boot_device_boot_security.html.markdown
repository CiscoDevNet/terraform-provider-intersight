---
subcategory: "boot"
layout: "intersight"
page_title: "Intersight: intersight_boot_device_boot_security"
description: |-
        The boot.DeviceBootSecurity object represents the user-configured state of UEFI Secure Boot for a server, as defined in its boot policy.
        #### Purpose
        This object's main function is to inventory the desired state of Secure Boot (Enabled or Disabled) from a server's applied policy. This provides a clear record of the intended security posture for the server's boot process. This allows for auditing and verification against the actual Secure Boot state reported by the BIOS.
        #### Key Concepts
        - **Desired Security State:** The secureBoot property captures the intended configuration from a boot policy.
        - **Compliance and Auditing:** Enables administrators to verify that servers are compliant with security policies requiring Secure Boot to be enabled.
        - **Configuration-Centric View:** Reflects the policy-defined setting, which can be compared against the actual hardware state for consistency.
        - **Server-Specific:** Associated with a compute.Physical object, linking the configured Secure Boot state to a specific server.

---

# Data Source: intersight_boot_device_boot_security
The boot.DeviceBootSecurity object represents the user-configured state of UEFI Secure Boot for a server, as defined in its boot policy.
#### Purpose
This object's main function is to inventory the desired state of Secure Boot ("Enabled" or "Disabled") from a server's applied policy. This provides a clear record of the intended security posture for the server's boot process. This allows for auditing and verification against the actual Secure Boot state reported by the BIOS.
#### Key Concepts
- **Desired Security State:** The secureBoot property captures the intended configuration from a boot policy.
- **Compliance and Auditing:** Enables administrators to verify that servers are compliant with security policies requiring Secure Boot to be enabled.
- **Configuration-Centric View:** Reflects the policy-defined setting, which can be compared against the actual hardware state for consistency.
- **Server-Specific:** Associated with a compute.Physical object, linking the configured Secure Boot state to a specific server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_boot_device_boot_security.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `secure_boot`:(string) The user desired BIOS secure boot as configured in the boot policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
