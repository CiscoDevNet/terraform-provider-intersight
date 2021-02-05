---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_lease"
description: |-
  Lease represents a single IQN address that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_iqnpool_lease
Lease represents a single IQN address that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `iqn_address`:(string) IQN address allocated for pool-based allocation \ prefix+suffix+number\ . 
* `moid`:(string) The unique identifier of this Managed Object instance. 
