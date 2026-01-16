---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_management_access_policy"
description: |-
        Displays a list of policies that define management access rules and privileges.

---

# Data Source: intersight_storage_pure_management_access_policy
Displays a list of policies that define management access rules and privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_management_access_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregation_strategy`:(string) The strategy used to combine the effects of multiple rules in this policy (e.g., 'deny-override'). 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Returns a value of true if the management access policy is currently active and enforced. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the management access policy. 
* `policy_type`:(string) The type of policy, typically 'management-access'. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
