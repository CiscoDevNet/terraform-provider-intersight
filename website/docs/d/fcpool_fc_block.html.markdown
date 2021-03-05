---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_fc_block"
description: |-
  A block of contiguous WWN addresses that are part of a pool.
---

# Data Source: intersight_fcpool_fc_block
A block of contiguous WWN addresses that are part of a pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_fc_block.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
 