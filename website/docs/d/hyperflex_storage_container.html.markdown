---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_storage_container"
description: |-
  A storage container (datastore) entity.
---

# Data Source: intersight_hyperflex_storage_container
A storage container (datastore) entity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_storage_container.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Storage container's creation time. 
* `data_block_size`:(int) Storage container data block size in bytes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `in_use`:(bool) Indicates whether the storage container has volumes. 
* `kind`:(string) Indicates whether the storage container was user-created, or system-created.* `UNKNOWN` - The storage container creator is unknown.* `USER_CREATED` - The storage container was created by a user action.* `INTERNAL` - The storage container was created by the system. 
* `last_access_time`:(string) Storage container's last access time. 
* `last_modified_time`:(string) Storage container's last modified time. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mount_status`:(string) Storage container mount status. Applicable only for NFS type.* `NOT_APPLICABLE` - The HyperFlex storage container mount status is not applicable.* `NORMAL` - The HyperFlex storage container mount status is normal.* `ALERT` - The HyperFlex storage container mount status is alert.* `FAILED` - The HyperFlex storage container mount status is failed. 
* `mount_summary`:(string) Storage container mount summary. Applicable only for NFS type.* `NOT_APPLICABLE` - The mount summary is not applicable for this HyperFlex storage container.* `MOUNTED` - The HyperFlex storage container is mounted.* `UNMOUNTED` - The HyperFlex storage container is unmounted.* `MOUNT_FAILURE` - The HyperFlex storage container mount summary is failure.* `UNMOUNT_FAILURE` - The HyperFlex storage container unmount summary is failure. 
* `name`:(string) Name of the storage container. 
* `provisioned_capacity`:(int) Provisioned capacity of the storage container in bytes. 
* `provisioned_volume_capacity_utilization`:(float) Provisioned capacity utilization of all volumes associated with the storage container. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Storage container type (SMB/NFS/iSCSI).* `NFS` - Storage container created/accesed through NFS protocol.* `SMB` - Storage container created/accessed through SMB protocol.* `iSCSI` - Storage container created/accessed through iSCSI protocol. 
* `un_compressed_used_bytes`:(int) Uncompressed bytes on storage container. 
* `uuid`:(string) UUID of the datastore/storage container. 
* `volume_count`:(int) Number of volumes associated with the storage container. 
 