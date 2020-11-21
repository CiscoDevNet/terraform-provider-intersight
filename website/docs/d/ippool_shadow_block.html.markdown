
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_shadow_block"
sidebar_current: "docs-intersight-data-source-ippool-shadow-block"
description: |-
A block of Contiguous IP addresses that are part of a shadow pool.
---

# Data Source: intersight_ippool._shadow_block
A block of Contiguous IP addresses that are part of a shadow pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `ip_type`:(string) Type of this IP addresses blocks.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
