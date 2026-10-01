---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_server_id_pool"
description: |-
        The ServerIdPool object represents the identifier pool used to allocate server IDs for rack or blade servers within a domain/device-registration scope.
        #### Purpose
        ServerIdPool provides a consistent allocation mechanism for server identifiers to ensure stable, predictable references for servers in inventory and lifecycle operations, including honoring preferred IDs where applicable.
        #### Key Concepts
        - **Stable server identifiers:** Allocates unique IDs to servers for consistent referencing.
        - **Preferred-ID handling:** Integrates policy-driven “preferred IDs” into allocation behavior where supported.
        - **Domain-scoped uniqueness:** Prevents identifier collisions by scoping to the domain/device registration.
        - **Supports lifecycle workflows:** Enables consistent identity across discovery, recommission, and replacement flows.

---

# Data Source: intersight_compute_server_id_pool
The ServerIdPool object represents the identifier pool used to allocate server IDs for rack or blade servers within a domain/device-registration scope.
#### Purpose
ServerIdPool provides a consistent allocation mechanism for server identifiers to ensure stable, predictable references for servers in inventory and lifecycle operations, including honoring preferred IDs where applicable.
#### Key Concepts
- **Stable server identifiers:** Allocates unique IDs to servers for consistent referencing.
- **Preferred-ID handling:** Integrates policy-driven “preferred IDs” into allocation behavior where supported.
- **Domain-scoped uniqueness:** Prevents identifier collisions by scoping to the domain/device registration.
- **Supports lifecycle workflows:** Enables consistent identity across discovery, recommission, and replacement flows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_server_id_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_available_id`:(int) Shows the next available Chassis ID to be allocated. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
