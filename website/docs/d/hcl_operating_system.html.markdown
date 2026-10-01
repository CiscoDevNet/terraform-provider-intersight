---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_operating_system"
description: |-
        OperatingSystems store operating system details used by the platform for OS-related workflows such as OS installation and compatibility selection. Each OperatingSystem captures the OS version information and links to the vendor that distributes the operating system.
        #### Purpose
        Provide a system-owned catalog of operating system versions and their associated vendors so users and automation can reference supported OS options during server management and OS installation workflows.
        #### Key Concepts
        - **System-managed OS catalog**: `owner: system` indicates the platform maintains the canonical OS entries that consumers read for selection and validation.
        - **Version representation**: `version` captures the operating system version identifier used in UI/workflows.
        - **Vendor association**: `vendor` links each OS entry to an `OperatingSystemVendor`, enabling grouping and filtering by distributor.
        - **Lifecycle coupling to vendor**: `onpeerdelete: cascade` ensures OS entries are cleaned up if the associated vendor entry is removed.
        - **Licensed read access**: READ requires the **Essentials** entitlement and is available to server and OS-install related roles for operational use.

---

# Data Source: intersight_hcl_operating_system
OperatingSystems store operating system details used by the platform for OS-related workflows such as OS installation and compatibility selection. Each OperatingSystem captures the OS version information and links to the vendor that distributes the operating system.
#### Purpose
Provide a system-owned catalog of operating system versions and their associated vendors so users and automation can reference supported OS options during server management and OS installation workflows.
#### Key Concepts
- **System-managed OS catalog**: `owner: system` indicates the platform maintains the canonical OS entries that consumers read for selection and validation.
- **Version representation**: `version` captures the operating system version identifier used in UI/workflows.
- **Vendor association**: `vendor` links each OS entry to an `OperatingSystemVendor`, enabling grouping and filtering by distributor.
- **Lifecycle coupling to vendor**: `onpeerdelete: cascade` ensures OS entries are cleaned up if the associated vendor entry is removed.
- **Licensed read access**: READ requires the **Essentials** entitlement and is available to server and OS-install related roles for operational use.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_operating_system.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) Version of the Operating System. 
 
