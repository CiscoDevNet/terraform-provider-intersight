
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_volume_snapshot"
sidebar_current: "docs-intersight-data-source-storage-pure-volume-snapshot"
description: |-
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
---

# Data Source: intersight_storage._pure_volume_snapshot
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `created_time`:(string) Exact date and time at which snapshot was created. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the snapshot which represents point-in-time copy of volume. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `protection_group_name`:(string) Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. 
* `serial`:(string) Unique serial number of the snapshot allocated by the storage array. 
* `size`:(int) Snapshot size represented in bytes. 
* `source`:(string) Source object on which the snapshot is created. It is the name of the originating volume. 
