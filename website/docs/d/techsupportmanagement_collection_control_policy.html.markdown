---
subcategory: "techsupportmanagement"
layout: "intersight"
page_title: "Intersight: intersight_techsupportmanagement_collection_control_policy"
description: |-
  Policy to control techsupport collection for a specific account.
---

# Data Source: intersight_techsupportmanagement_collection_control_policy
Policy to control techsupport collection for a specific account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_techsupportmanagement_collection_control_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_type`:(string) Deployment type defines whether the policy is associated with a SaaS or Appliance account.* `None` - Service deployment type None.* `SaaS` - Service deployment type SaaS.* `Appliance` - Service deployment type Appliance. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tech_support_collection`:(string) Enable or Disable techsupport collection for a specific account.* `Enable` - Enable techsupport collection.* `Disable` - Disable techsupport collection. 
 