---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_controller_drive"
description: |-
        The Local Storage present in a server.

---

# Data Source: intersight_storage_controller_drive
The Local Storage present in a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_controller_drive.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_type`:(string) The type of Local Storage like FlexMMC, USB Drive.* `Unknown` - Storage partition type is not known.* `FlexMMC` - The FlexMMC type of storage local drive.* `USB` - The USB type of storage drive. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of local storage drive. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the local Storage. 
* `partition_count`:(int) Total Partition count in local storage. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `storage_id`:(string) The Id of the local Storage. 
* `type`:(string) The type of storage like internal or external.* `Unknown` - Not any of the known Storage Types.* `Internal` - The internal storage type.* `External` - The external storage type. 
 
