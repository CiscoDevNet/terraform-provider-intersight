
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_blocks`:(string) The number of blocks that are a part of the virtual drive. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `physical_drive`:(string) The physical disk for which the usage is reported. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `span`:(string) The span of the physical disk. 
* `starting_block`:(string) The starting block id of the virtual drive within the physical drive. 
* `state`:(string) The current state of the physical disk usage. 
* `virtual_drive`:(string) The virtual drive corresponding to the physical disk. 
