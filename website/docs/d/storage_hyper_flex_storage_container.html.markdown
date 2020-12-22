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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string) Storage container's creation time. 
* `last_modified_time`:(string) Storage container's last modified time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `provisioned_capacity`:(int) Provisioned Capacity of the Storage container in bytes. 
* `type`:(string) Storage Container type (SMB/NFS/iSCSI).* `NFS` - Storage container created/accesed through NFS protocol.* `SMB` - Storage container created/accessed through SMB protocol.* `iSCSI` - Storage container created/accessed through iSCSI protocol. 
* `un_compressed_used_bytes`:(int) Uncompressed bytes on Storage Container. 
* `uuid`:(string) Uuid of the Datastore/Storage Container. 
