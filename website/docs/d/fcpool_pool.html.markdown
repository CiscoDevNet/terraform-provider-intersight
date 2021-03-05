---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_pool"
description: |-
  Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
---

# Data Source: intersight_fcpool_pool
Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_pool.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `pool_purpose`:(string) Purpose of this WWN pool. 
* `size`:(int) Total number of identifiers in this pool. 
 