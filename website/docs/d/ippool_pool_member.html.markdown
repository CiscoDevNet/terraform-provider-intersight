---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_pool_member"
description: |-
  PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
---

# Data Source: intersight_ippool_pool_member
PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_pool_member.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `ip_v4_address`:(string) IPv4 Address of this pool member. 
* `ip_v6_address`:(string) IPv6 Address of this pool member. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 