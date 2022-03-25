---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_flex_util_virtual_drive"
description: |-
        Storage Flex Util Virtual Drive.

---

# Data Source: intersight_storage_flex_util_virtual_drive
Storage Flex Util Virtual Drive.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_flex_util_virtual_drive.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `drive_status`:(string) Status of the Flex Util virtual drive. 
* `drive_type`:(string) Type of virtual drive managed by flex util controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `partition_id`:(string) Disk Partition Id of virtual drive managed by flex util controller. 
* `partition_name`:(string) Partition name of the Flex Util virtual drive. 
* `resident_image`:(string) The resident image on the flex util virtual Drive. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(string) Size of the Flex Util virtual drive. 
* `virtual_drive`:(string) Virtual drive on the Flex Util controller. 
 
