
---
layout: "intersight"
page_title: "Intersight: intersight_storage_physical_disk_usage"
sidebar_current: "docs-intersight-data-source-storage-physical-disk-usage"
description: |-
Has usage map between physical disks and virtual drives.
---

# Data Source: intersight_storage._physical_disk_usage
Has usage map between physical disks and virtual drives.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_blocks`:(string) The number of blocks that are a part of the virtual drive. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `physical_drive`:(string) The physical disk for which the usage is reported. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `span`:(string) The span of the physical disk. 
* `starting_block`:(string) The starting block id of the virtual drive within the physical drive. 
* `state`:(string) The current state of the physical disk usage. 
* `virtual_drive`:(string) The virtual drive corresponding to the physical disk. 
