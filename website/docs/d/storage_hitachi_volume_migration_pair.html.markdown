---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_volume_migration_pair"
description: |-
        The HitachiVolumeMigrationPair object defines a copy pair used for volume migration within the Hitachi storage array, focusing on migration operations and pair management.
        #### Purpose
        HitachiVolumeMigrationPair serves as a pivotal element for managing volume migration pairs, supporting operations related to copy modes and pair statuses.
        #### Key Concepts
        - **Migration Operations:** Manages copy modes, LDEV numbers, and status information for volume migration pairs.
        - **Privilege Sets:** Ensures secure access and management of migration pair settings with defined privilege sets.
        - **Licensing:** Operates under specified entitlements, supporting authorized migration operations.

---

# Data Source: intersight_storage_hitachi_volume_migration_pair
The HitachiVolumeMigrationPair object defines a copy pair used for volume migration within the Hitachi storage array, focusing on migration operations and pair management.
#### Purpose
HitachiVolumeMigrationPair serves as a pivotal element for managing volume migration pairs, supporting operations related to copy modes and pair statuses.
#### Key Concepts
- **Migration Operations:** Manages copy modes, LDEV numbers, and status information for volume migration pairs.
- **Privilege Sets:** Ensures secure access and management of migration pair settings with defined privilege sets.
- **Licensing:** Operates under specified entitlements, supporting authorized migration operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_volume_migration_pair.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `copy_mode`:(string) Copy mode. NotSynchronized or VolumeMigration is stored. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `local_clone_copypair_id`:(string) Object ID of the pair. The following informations of pair are output, separated by commas. <copy group name>, <device group name for the P-VOL (source volume)>, <device group name for the S-VOL (target volume)>, <name of the pair>. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pvol_ldev_id`:(int) LDEV number of the P-VOL (source volume) with a decimal (base 10) number. 
* `pvol_status`:(string) Pair volume status of the P-VOL. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svol_ldev_id`:(int) LDEV number of the S-VOL (target volume) with a decimal (base 10) number. 
* `svol_status`:(string) Pair volume status of the S-VOL. 
 
