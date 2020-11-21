
---
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_block"
sidebar_current: "docs-intersight-data-source-uuidpool-block"
description: |-
A block of contiguous UUID addresses that are part of a pool.
---

# Data Source: intersight_uuidpool._block
A block of contiguous UUID addresses that are part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
