
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `created_time`:(string) Protection group snapshot creation time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Protection group snapshot name which represents point-in-time copy of all members in protection group. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `size`:(int) Snapshot size represented in bytes. It is a cumulative size of all snapshots in a set. 
* `source`:(string) Source protection group name on which the snapshot is created. 
