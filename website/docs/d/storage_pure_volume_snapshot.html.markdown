
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `created_time`:(string) Exact date and time at which snapshot was created. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the snapshot which represents point-in-time copy of volume. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `protection_group_name`:(string) Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume. 
* `serial`:(string) Unique serial number of the snapshot allocated by the storage array. 
* `size`:(int) Snapshot size represented in bytes. 
* `source`:(string) Source object on which the snapshot is created. It is the name of the originating volume. 
