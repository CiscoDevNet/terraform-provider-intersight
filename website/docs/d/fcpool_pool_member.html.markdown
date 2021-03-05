---
subcategory: "fcpool"
layout: "intersight"
page_title: "Intersight: intersight_fcpool_pool_member"
description: |-
  PoolMember represents a single WWN ID that is part of a pool.
---

# Data Source: intersight_fcpool_pool_member
PoolMember represents a single WWN ID that is part of a pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fcpool_pool_member.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `wwn_id`:(string) WWN ID of this pool member. 
 