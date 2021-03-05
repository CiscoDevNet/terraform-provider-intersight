---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_shadow_block"
description: |-
  A block of Contiguous IP addresses that are part of a shadow pool.
---

# Data Source: intersight_ippool_shadow_block
A block of Contiguous IP addresses that are part of a shadow pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_shadow_block.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `ip_type`:(string) Type of this IP addresses blocks.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
 