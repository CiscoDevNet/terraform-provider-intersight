---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_remote_replication"
description: |-
        A remote copy pair entity in Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_remote_replication
A remote copy pair entity in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_remote_replication.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `consistency_group_id`:(string) Consistency group ID. If no consistency group consists, information is not input. 
* `copy_pace`:(string) Copy speed. Number for the size of tracks to be copied. 
* `create_time`:(string) The time when this managed object was created. 
* `delta_status`:(string) Status of the 3DC multi-target configuration that uses delta resync. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fence_level`:(string) Fence level. Whether the P-VOL can be written to when the pair is split due to error. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mu_number`:(string) MU (mirror unit) number of the volume. 
* `name`:(string) Object ID of the remote copy pair. 
* `path_group_id`:(string) Path group ID of the RCU. 
* `pvol_io_mode`:(string) I-O mode of the P-VOL. Information is input only in the case of global-active device. 
* `pvol_journal_id`:(string) Journal ID of the P-VOL. A value is input only in the case of Universal Replicator. 
* `pvol_ldev_id`:(int) LDEV number of primary volume. 
* `pvol_storage_serial`:(string) Serial number of the storage system on the P-VOL. 
* `quorum_disk_id`:(string) ID of the Quorum disk. A value is input only in the case of global-active device. 
* `replication_type`:(string) Pair type of the remote copy pair. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the remote copy pair. 
* `svol_io_mode`:(string) I-O mode of the S-VOL. Information is input only in the case of global-active device. 
* `svol_journal_id`:(string) Journal ID of the S-VOL. A value is input only in the case of Universal Replicator. 
* `svol_ldev_id`:(int) LDEV number of secondary volume. 
* `svol_storage_serial`:(string) Serial number of the storage system on the S-VOL. 
 
