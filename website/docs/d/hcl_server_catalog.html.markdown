---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_server_catalog"
description: |-
        ServerCatalogs are system-owned catalog entries that map a server PID (three-part server identifier) to the set of supported processor families for that server model. This supports HCL-driven validation and guided selection of compatible CPU families.
        #### Purpose
        Expose server model → supported processor family mappings so users and automation can construct valid server/CPU combinations for HCL compatibility evaluation.
        #### Key Concepts
        - **Server model identifier:** Uses `serverPid` as the canonical lookup key for the server model.
        - **Processor family mapping:** Lists supported processor families associated with the server PID.
        - **Compatibility guidance:** Enables front-end selection and backend validation to avoid unsupported server/CPU combinations.
        - **System catalog semantics:** Maintained as system-owned reference data and exposed read-only to authorized roles.

---

# Data Source: intersight_hcl_server_catalog
ServerCatalogs are system-owned catalog entries that map a server PID (three-part server identifier) to the set of supported processor families for that server model. This supports HCL-driven validation and guided selection of compatible CPU families.
#### Purpose
Expose server model → supported processor family mappings so users and automation can construct valid server/CPU combinations for HCL compatibility evaluation.
#### Key Concepts
- **Server model identifier:** Uses `serverPid` as the canonical lookup key for the server model.
- **Processor family mapping:** Lists supported processor families associated with the server PID.
- **Compatibility guidance:** Enables front-end selection and backend validation to avoid unsupported server/CPU combinations.
- **System catalog semantics:** Maintained as system-owned reference data and exposed read-only to authorized roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_server_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
