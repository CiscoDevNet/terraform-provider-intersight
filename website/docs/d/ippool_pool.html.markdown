---
subcategory: "ippool"
layout: "intersight"
page_title: "Intersight: intersight_ippool_pool"
description: |-
  Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.
---

# Data Source: intersight_ippool_pool
Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `size`:(int) Total number of identifiers in this pool. 
* `v4_assigned`:(int) Number of IPv4 addresses currently in use. 
* `v4_size`:(int) Number of IPv4 addresses in this pool. 
* `v6_assigned`:(int) Number of IPv6 addresses currently in use. 
* `v6_size`:(int) Number of IPv6 addresses in this pool. 
