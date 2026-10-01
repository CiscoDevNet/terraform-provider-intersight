---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume_group"
description: |-
        PureVolumeGroups represent groups of volumes that can be managed as a single entity on a Pure Storage array. They simplify administration by allowing related volumes to be organized together and governed consistently, rather than managed one by one.
        #### Purpose
        Provides a read-only inventory view of Pure volume groups so administrators can identify grouped volumes and correlate them to the owning array and device registration for streamlined storage management.
        #### Key Concepts
        - **Administrative grouping construct**: A volume group is a logical collection of volumes treated as one management unit.
        - **Consistent policy/application context**: Grouping supports simpler, more uniform handling of storage resources across related volumes.
        - **Stable group identity**: Identified by `name`, which is also indexed for efficient lookup.
        - **Array and device correlation**:
        - `array` links the group to the owning `PureArray` and cascades on array deletion.
        - `registeredDevice` ties the group to the Intersight connection for the storage array.
        - **Read-only, licensed inventory**: Available via READ and requires the **Advantage** entitlement.
        - **Indexed for scale**: Additional indexes on `RegisteredDevice` and `Array` support efficient filtering/reporting across storage environments.

---

# Data Source: intersight_storage_pure_volume_group
PureVolumeGroups represent groups of volumes that can be managed as a single entity on a Pure Storage array. They simplify administration by allowing related volumes to be organized together and governed consistently, rather than managed one by one.
#### Purpose
Provides a read-only inventory view of Pure volume groups so administrators can identify grouped volumes and correlate them to the owning array and device registration for streamlined storage management.
#### Key Concepts
- **Administrative grouping construct**: A volume group is a logical collection of volumes treated as one management unit.
- **Consistent policy/application context**: Grouping supports simpler, more uniform handling of storage resources across related volumes.
- **Stable group identity**: Identified by `name`, which is also indexed for efficient lookup.
- **Array and device correlation**:
  - `array` links the group to the owning `PureArray` and cascades on array deletion.
  - `registeredDevice` ties the group to the Intersight connection for the storage array.
- **Read-only, licensed inventory**: Available via READ and requires the **Advantage** entitlement.
- **Indexed for scale**: Additional indexes on `RegisteredDevice` and `Array` support efficient filtering/reporting across storage environments.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_volume_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Volume Group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
