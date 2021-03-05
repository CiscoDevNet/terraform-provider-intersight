---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_block_lease"
description: |-
  BlockLease represents an IP address that is allocated from a pool to a specific entity like server profile.
---

# Data Source: intersight_ippool_block_lease
BlockLease represents an IP address that is allocated from a pool to a specific entity like server profile.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ippool_block_lease.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `nr_count`:(int) Count of number of leases requested. 
* `ip_type`:(string) Type of the IP address requested.* `IPv4` - IP V4 address type requested.* `IPv6` - IP V6 address type requested. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 