---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_domain_policy_requirement"
description: |-
        Version Constraint requirement for a domian policy.

---

# Data Source: intersight_capability_domain_policy_requirement
Version Constraint requirement for a domian policy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_domain_policy_requirement.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `min_bundle_version`:(string) Minimum Bundle version from which it is supported. 
* `min_version`:(string) Minimum Version from which policy is supported. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Type of the platform for which version compatibility is specified. Example - 3GFI, 4GFI, etc. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `policy_name`:(string) Policy Name for which version compatibility is specified. Example - snmp.Policy, ldap.Policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
