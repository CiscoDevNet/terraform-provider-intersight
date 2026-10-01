---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_operating_system_vendor"
description: |-
        OperatingSystemVendors are system-owned catalog records representing operating system vendors (for example, vendor families used by the HCL tool). They provide a normalized set of OS vendor names used for OS selection and compatibility lookups.
        #### Purpose
        Provide a canonical list of OS vendors used in HCL compatibility queries and OS-related user workflows.
        #### Key Concepts
        - **Normalized vendor identity:** Centralizes OS vendor naming for consistent matching and filtering.
        - **Foundation for OS catalog:** Used as the parent/reference for OperatingSystem records.
        - **Read-focused catalog:** Primarily consumed by users and tooling to populate selection lists and drive validation inputs.

---

# Data Source: intersight_hcl_operating_system_vendor
OperatingSystemVendors are system-owned catalog records representing operating system vendors (for example, vendor families used by the HCL tool). They provide a normalized set of OS vendor names used for OS selection and compatibility lookups.
#### Purpose
Provide a canonical list of OS vendors used in HCL compatibility queries and OS-related user workflows.
#### Key Concepts
- **Normalized vendor identity:** Centralizes OS vendor naming for consistent matching and filtering.
- **Foundation for OS catalog:** Used as the parent/reference for OperatingSystem records.
- **Read-focused catalog:** Primarily consumed by users and tooling to populate selection lists and drive validation inputs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_operating_system_vendor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the vendor of the operating system. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
