
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group_snapshot"
sidebar_current: "docs-intersight-data-source-storage-pure-protection-group-snapshot"
description: |-
Protection group snapshot entity in Pure protection group.
---

# Data Source: intersight_storage._pure_protection_group_snapshot
Protection group snapshot entity in Pure protection group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `created_time`:(string) Protection group snapshot creation time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Protection group snapshot name which represents point-in-time copy of all members in protection group. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `size`:(int) Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. 
* `source`:(string) Source protection group name on which the snapshot is created. 
