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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_physical_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `capacity`:(int) Capacity of the physical entity added. 
* `nr_count`:(int) Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_new`:(bool) If the PhysicalItem is new, this is set to true, else false. 
* `max_count`:(int) Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the recommended physical device which is externally identifiable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the physical device recommended. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_moid`:(string) Moid of the managed object which represents the existing physical entity. 
* `type`:(string) The type of physical device recommended.* `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk.* `Node` - The Enum value Node represents that the item type recommended is a Storage Node.* `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster. 
* `unit`:(string) Unit of the new capacity.* `TB` - The Enum value TB represents that the measurement unit is in terabytes.* `MB` - The Enum value MB represents that the measurement unit is in megabytes. 
* `uuid`:(string) Uuid of the recommended physical device. 
 
