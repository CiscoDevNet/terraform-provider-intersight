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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `created_time`:(string) Exact date and time at which snapshot was created. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the snapshot which represents point-in-time copy of volume. 
* `protection_group_name`:(string) Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. 
* `size`:(int) Snapshot size represented in bytes. 
* `nr_source`:(string) Source object on which the snapshot is created. It is the name of the originating volume. 
* `uuid`:(string) UUID of the volume snapshot. 
