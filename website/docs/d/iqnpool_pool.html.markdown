---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_pool"
description: |-
  Pool represents a collection of iSCSI Qualified Names (IQNs) for use as initiator identifiers by iSCSI vNICs.
---

# Data Source: intersight_iqnpool_pool
Pool represents a collection of iSCSI Qualified Names (IQNs) for use as initiator identifiers by iSCSI vNICs.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `prefix`:(string) The prefix for IQN blocks created for this pool. 
* `size`:(int) Total number of identifiers in this pool. 
