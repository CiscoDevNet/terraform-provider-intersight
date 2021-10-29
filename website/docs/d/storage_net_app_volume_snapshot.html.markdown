---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_volume_snapshot"
description: |-
  NetApp Volume Snapshot is a read-only image of a traditional or FlexVol volume, or an aggregate, that captures the state of the file system at a point in time.
---

# Data Source: intersight_storage_net_app_volume_snapshot
NetApp Volume Snapshot is a read-only image of a traditional or FlexVol volume, or an aggregate, that captures the state of the file system at a point in time.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_volume_snapshot.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Exact date and time at which snapshot was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the snapshot which represents point-in-time copy of volume. 
* `protection_group_name`:(string) Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Snapshot size represented in bytes. 
* `nr_source`:(string) Source object on which the snapshot is created. It is the name of the originating volume. 
* `uuid`:(string) Universally unique identifier of the volume snapshot. 
 