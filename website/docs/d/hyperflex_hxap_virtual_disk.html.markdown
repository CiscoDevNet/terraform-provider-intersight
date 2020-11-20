
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_virtual_disk"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-virtual-disk"
description: |-
The Virtual disk created on HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex._hxap_virtual_disk
The Virtual disk created on HyperFlex Application Platform compute cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_mode`:(string) Access mode of the virtual disk.* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. 
* `capacity`:(int) Disk capacity represented in bytes. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `mode`:(string) File mode of the disk  example - Filesystem, Block.* `Block` - It is a Block virtual disk.* `Filesystem` - It is a File system virtual disk. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage disk. Name must be unique within a Datastore. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `source_file_path`:(string) Source file path associated with virtual machine disk. 
* `source_virtual_disk`:(string) Source disk name from where the clone is done. 
* `uuid`:(string) UUID of the virtual disk. 
