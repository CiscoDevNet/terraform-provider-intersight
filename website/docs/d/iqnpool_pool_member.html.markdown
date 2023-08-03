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
To access the ith object of the results obtained, use `data.intersight_iqnpool_pool_member.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `assigned`:(bool) Boolean to represent whether the ID is in use. 
* `assigned_by_another`:(bool) Boolean to represent whether the ID is used either statically or by another pool. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `iqn_address`:(string) IQN Address of this pool member. It is constructed as <prefix>:<suffix>:<number>. 
* `iqn_number`:(int) Number of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>. 
* `iqn_prefix`:(string) Prefix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>. 
* `iqn_suffix`:(string) Suffix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reserved`:(bool) Boolean to represent whether the ID is reserved. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
