---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume_snapshot"
description: |-
  Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
---

# Data Source: intersight_storage_pure_volume_snapshot
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string) Exact date and time at which snapshot was created. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the snapshot which represents point-in-time copy of volume. 
* `protection_group_name`:(string) Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. 
* `serial`:(string) Unique serial number of the snapshot allocated by the storage array. 
* `size`:(int) Snapshot size represented in bytes. 
* `nr_source`:(string) Source object on which the snapshot is created. It is the name of the originating volume. 
