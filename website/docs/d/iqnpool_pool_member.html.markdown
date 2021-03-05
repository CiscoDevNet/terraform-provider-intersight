---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_pool_member"
description: |-
  PoolMember represents a single IQN address that is part of a pool.
---

# Data Source: intersight_iqnpool_pool_member
PoolMember represents a single IQN address that is part of a pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iqnpool_pool_member.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `iqn_address`:(string) IQN Address of this pool member. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 