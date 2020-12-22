---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_pool"
description: |-
  Pool represents a collection of MAC addresses that can be allocated to VNICs of a server profile.
---

# Data Source: intersight_macpool_pool
Pool represents a collection of MAC addresses that can be allocated to VNICs of a server profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `size`:(int) Total number of identifiers in this pool. 
