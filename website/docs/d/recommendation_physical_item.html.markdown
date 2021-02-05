---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_physical_item"
description: |-
  Entity representing the recommended physical device.
---

# Data Source: intersight_recommendation_physical_item
Entity representing the recommended physical device.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(int) Capacity of the physical entity added. 
* `nr_count`:(int) Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation. 
* `is_new`:(bool) If the PhysicalItem is new, this is set to true, else false. 
* `max_count`:(int) Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation. 
* `model`:(string) Model of the recommended physical device which is externally identifiable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the physical device recommended. 
* `source_moid`:(string) Moid of the managed object which represents the existing physical entity. 
* `type`:(string) The type of physical device recommended.* `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk.* `Node` - The Enum value Node represents that the item type recommended is a Storage Node.* `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster. 
* `unit`:(string) Unit of the new capacity.* `TB` - The Enum value TB represents that the measurement unit is in terabytes.* `MB` - The Enum value MB represents that the measurement unit is in megabytes. 
* `uuid`:(string) Uuid of the recommended physical device. 
