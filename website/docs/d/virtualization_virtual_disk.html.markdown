---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_virtual_disk"
description: |-
  Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
---

# Data Source: intersight_virtualization_virtual_disk
Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_virtual_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(string) Disk capacity to be provided with units example - 10Gi. 
* `create_time`:(string) The time when this managed object was created. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk.* `` - Disk mode is either unknown or not supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_certs`:(string) Base64 encoded CA certificates of the https source to check against. 
* `source_disk_to_clone`:(string) Source disk from which the content is copied. 
* `source_file_path`:(string) Image path used to import on the created disk. 
 