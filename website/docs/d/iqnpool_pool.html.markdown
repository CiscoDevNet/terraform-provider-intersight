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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iqnpool_pool.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assigned`:(int) Number of IDs that are currently assigned (in use). 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `prefix`:(string) The prefix for any IQN blocks created for this pool. IQN Prefix must have the following format \ iqn.yyyy-mm.naming-authority\ , where naming-authority is usually the reverse syntax of the Internet domain name of the naming authority. 
* `reserved`:(int) Number of IDs that are currently reserved (and not in use). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) Total number of identifiers in this pool. 
 
