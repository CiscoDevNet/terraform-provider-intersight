---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_directory_snapshot"
description: |-
        Directory snapshots are created manually or by adding snapshot policies to managed directories. Each snapshot policy can be re-used for multiple directories.

---

# Data Source: intersight_storage_pure_directory_snapshot
Directory snapshots are created manually or by adding snapshot policies to managed directories. Each snapshot policy can be re-used for multiple directories.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_directory_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_name`:(string) The customizable portion of the client-visible snapshot name. A full snapshot name is constructed in the form of DIR.CLIENT_NAME.SUFFIX where DIR is the full managed directory name, CLIENT_NAME is the client name, and SUFFIX is the suffix. The client-visible snapshot name is CLIENT_NAME.SUFFIX. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) The snapshot creation time, measured in milliseconds since the UNIX epoch. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user-specified name. The name must be locally unique and can be changed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_source`:(string) The directory from which this snapshot was taken. 
 
