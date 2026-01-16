---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group"
description: |-
        The PureProtectionGroup object is a critical element for managing data protection within a PureStorage FlashArray, facilitating volume protection through snapshots.
        #### Purpose
        PureProtectionGroup enables the protection of volumes by associating them with protection groups, ensuring data integrity and availability through snapshot creation.
        #### Key Concepts
        - **Snapshot Management:** PureProtectionGroup orchestrates the creation and management of snapshots, providing a robust mechanism for data protection.
        - **Volume Association:** It supports direct or indirect volume association, allowing flexible protection strategies within the FlashArray.
        - **Security:** Utilizes privilege sets to ensure secure access and management of protection groups.

---

# Data Source: intersight_storage_pure_protection_group
The PureProtectionGroup object is a critical element for managing data protection within a PureStorage FlashArray, facilitating volume protection through snapshots.  
#### Purpose  
PureProtectionGroup enables the protection of volumes by associating them with protection groups, ensuring data integrity and availability through snapshot creation.  
#### Key Concepts
- **Snapshot Management:** PureProtectionGroup orchestrates the creation and management of snapshots, providing a robust mechanism for data protection. 
- **Volume Association:** It supports direct or indirect volume association, allowing flexible protection strategies within the FlashArray. 
- **Security:** Utilizes privilege sets to ensure secure access and management of protection groups.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_protection_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the protection Group. 
* `pod_name`:(string) A pod representing a collection of protection groups and volumes is created on one array and stretched to another array, resulting in fully synchronized writes between the two arrays. 
* `prefix`:(string) Prefix used for all generated snapshots from the protection group. 
* `realm_name`:(string) A realm is the core multi-tenancy component on a Pure Flash Array, providing a self-contained, virtual storage environment with dedicated policies and quotas for secure data isolation and predictable performance. 
* `replication_enabled`:(bool) Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Overall size of all snapshots in the protection group, represented in bytes. 
* `snapshot_enabled`:(bool) Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. 
* `nr_source`:(string) Name of PureStorage array name on which the protection group is created. 
 
