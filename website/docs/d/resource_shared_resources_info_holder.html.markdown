---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_shared_resources_info_holder"
description: |-
        For each MO, the shared config resource relationships from a permission resource are maintained by SharedResourcesInfoHolder. For example, let orgCommon is shared with orgRegular. A user having Server Administrator privilege on orgRegular and a minimum of READ access on sol policy API for orgCommon will be able to create a server profile in orgRegular by referencing an sol policy of orgCommon.

---

# Data Source: intersight_resource_shared_resources_info_holder
For each MO, the shared config resource relationships from a permission resource are maintained by SharedResourcesInfoHolder. For example, let orgCommon is shared with orgRegular. A user having Server Administrator privilege on orgRegular and a minimum of READ access on sol policy API for orgCommon will be able to create a server profile in orgRegular by referencing an sol policy of orgCommon.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_shared_resources_info_holder.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
