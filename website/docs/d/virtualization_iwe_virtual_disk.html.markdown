---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_iwe_virtual_disk"
description: |-
  The Virtual disk created on Intersight Workload Engine compute cluster.
---

# Data Source: intersight_virtualization_iwe_virtual_disk
The Virtual disk created on Intersight Workload Engine compute cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_iwe_virtual_disk.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_mode`:(string) Access mode of the virtual disk.* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.* `` - Unknown disk access mode. 
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk.* `` - Disk mode is either unknown or not supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Disk size represented in bytes. 
* `source_file_path`:(string) Source file path associated with virtual machine disk. 
* `source_virtual_disk`:(string) Virtual disk used for cloning new disk. 
* `uuid`:(string) UUID of the virtual disk. 
 