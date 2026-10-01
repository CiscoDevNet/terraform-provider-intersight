---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_chassis_id_pool"
description: |-
        The ChassisIdPool object represents the identifier pool used to allocate chassis IDs for newly discovered chassis within a device registration/domain context.
        #### Purpose
        ChassisIdPool provides a controlled mechanism for allocating chassis identifiers, enabling consistent addressing and correlation for chassis entities discovered and managed within a domain. It also supports honoring preferred IDs defined through higher-level policy intent.
        #### Key Concepts
        - **Deterministic ID allocation:** Ensures chassis identifiers are assigned in a controlled, conflict-free manner.
        - **Preferred-ID integration:** Supports propagation of user-preferred IDs (where defined) into the allocation pool behavior.
        - **Domain scoping:** Tied to a specific device registration context to prevent cross-domain collisions.
        - **Lifecycle-aligned existence:** Exists as long as the corresponding device registration/domain context exists.

---

# Data Source: intersight_equipment_chassis_id_pool
The ChassisIdPool object represents the identifier pool used to allocate chassis IDs for newly discovered chassis within a device registration/domain context.
#### Purpose
ChassisIdPool provides a controlled mechanism for allocating chassis identifiers, enabling consistent addressing and correlation for chassis entities discovered and managed within a domain. It also supports honoring preferred IDs defined through higher-level policy intent.
#### Key Concepts
- **Deterministic ID allocation:** Ensures chassis identifiers are assigned in a controlled, conflict-free manner.
- **Preferred-ID integration:** Supports propagation of user-preferred IDs (where defined) into the allocation pool behavior.
- **Domain scoping:** Tied to a specific device registration context to prevent cross-domain collisions.
- **Lifecycle-aligned existence:** Exists as long as the corresponding device registration/domain context exists.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_chassis_id_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_available_id`:(int) Shows the next available Chassis ID to be allocated. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
