---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_virtual_disk"
description: |-
  The Virtual disk created on HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex_hxap_virtual_disk
The Virtual disk created on HyperFlex Application Platform compute cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_virtual_disk.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_mode`:(string) Access mode of the virtual disk.* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `size`:(int) Disk size represented in bytes. 
* `source_file_path`:(string) Source file path associated with virtual machine disk. 
* `source_virtual_disk`:(string) Source disk name from where the clone is done. 
* `uuid`:(string) UUID of the virtual disk. 
 