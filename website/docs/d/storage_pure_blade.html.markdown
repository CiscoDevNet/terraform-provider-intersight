---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_blade"
description: |-
        The PureBlade object models a blade hardware component within a PureStorage FlashBlade array. This provides foundational support for hardware inventory and blade-level management.
        #### Purpose
        PureBlade is intended to represent individual hardware blades, aiding administrators in monitoring blade health, status, and capacity. This ensures optimal performance and facilitates maintenance operations within the array.
        #### Key Concepts
        - **Hardware Inventory**: Enables granular tracking of blade components, supporting efficient asset and lifecycle management.
        - **Status Monitoring**: Facilitates real-time status checks and operational insights for each blade.
        - **Integration**: Connects blade data to the broader array and device registration infrastructure.
        - **Access Control**: Enforces privilege sets to ensure only authorized personnel can view or manage blade details.

---

# Data Source: intersight_storage_pure_blade
The PureBlade object models a blade hardware component within a PureStorage FlashBlade array. This provides foundational support for hardware inventory and blade-level management.
#### Purpose
PureBlade is intended to represent individual hardware blades, aiding administrators in monitoring blade health, status, and capacity. This ensures optimal performance and facilitates maintenance operations within the array.
#### Key Concepts
- **Hardware Inventory**: Enables granular tracking of blade components, supporting efficient asset and lifecycle management.
- **Status Monitoring**: Facilitates real-time status checks and operational insights for each blade.
- **Integration**: Connects blade data to the broader array and device registration infrastructure.
- **Access Control**: Enforces privilege sets to ensure only authorized personnel can view or manage blade details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_blade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `blade_id`:(string) Unique identifier of the blade. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the blade within the FlashBlade array. 
* `raw_capacity`:(int) Raw storage capacity of the blade in bytes. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Operational status of the blade. 
 
