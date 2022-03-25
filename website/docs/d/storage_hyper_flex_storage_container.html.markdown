---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hyper_flex_storage_container"
description: |-
        A Storage Container (Datastore) entity.

---

# Data Source: intersight_storage_hyper_flex_storage_container
A Storage Container (Datastore) entity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hyper_flex_storage_container.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Storage container's creation time. 
* `data_block_size`:(int) Storage Container data block size 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `in_use`:(bool) Indicates whether the Storage Container has Volumes. 
* `last_access_time`:(string) Storage container's last access time. 
* `last_modified_time`:(string) Storage container's last modified time. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `provisioned_capacity`:(int) Provisioned Capacity of the Storage container. 
* `provisioned_volume_capacity_utilization`:(float) Provisioned Capacity Utilization of All Volumes associated with the Storage Container. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Storage Container type (SMB/NFS/iSCSI).* `NFS` - Storage container created/accesed through NFS protocol.* `SMB` - Storage container created/accessed through SMB protocol.* `iSCSI` - Storage container created/accessed through iSCSI protocol. 
* `un_compressed_used_bytes`:(int) Uncompressed bytes on Storage Container. 
* `uuid`:(string) UUID of the Datastore/Storage Containter. 
* `volume_count`:(int) Number of Volumes associated with the Storage Container. 
 
