
---
layout: "intersight"
page_title: "Intersight: intersight_macpool_id_block"
sidebar_current: "docs-intersight-data-source-macpool-id-block"
description: |-
A block of contiguous MAC addresses that are part of a pool.
---

# Data Source: intersight_macpool._id_block
A block of contiguous MAC addresses that are part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
