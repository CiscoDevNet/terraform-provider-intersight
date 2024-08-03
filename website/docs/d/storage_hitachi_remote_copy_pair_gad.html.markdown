---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_remote_copy_pair_gad"
description: |-
        The global-active device pair entity in Hitachi storage array.

---

# Data Source: intersight_storage_hitachi_remote_copy_pair_gad
The global-active device pair entity in Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_remote_copy_pair_gad.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_local_time`:(string) Time at which the pair was created. The local time of the storage system is displayed. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_ssws`:(bool) Indicates whether the state of the volume on the local storage system is SSWS. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mu_number`:(string) MU (mirror unit) number of the volume. 
* `name`:(string) Object ID of the remote copy pair. 
* `pvol_ldev_id`:(int) LDEV number of primary volume. 
* `pvol_storage_serial`:(string) Serial number of the storage system on the P-VOL. 
* `quorum_disk_id`:(string) ID of the Quorum disk. A value is input only in the case of global-active device. 
* `replication_type`:(string) Pair type of the remote copy pair. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the remote copy pair. 
* `suspended_mode`:(string) Block or Remote instructions when a pair is suspended. 
* `svol_ldev_id`:(int) LDEV number of secondary volume. 
* `svol_storage_serial`:(string) Serial number of the storage system on the S-VOL. 
 
