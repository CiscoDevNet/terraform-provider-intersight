---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_group_status"
description: |-
  Status of a group of applications.
---

# Data Source: intersight_appliance_group_status
Status of a group of applications.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_group_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `group_name`:(string) The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `overall_status`:(string) The overall API status from this group's applications. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 