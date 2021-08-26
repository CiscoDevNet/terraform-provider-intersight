---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_app_details"
description: |-
  Details of apps installed on Nexus Dashboard.
---

# Data Source: intersight_niatelemetry_app_details
Details of apps installed on Nexus Dashboard.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_app_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `app_name`:(string) Names of apps running on ND. 
* `app_status`:(string) Status of apps running on ND. 
* `app_version`:(string) Versions of apps running on ND. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nexus_dashboard`:(string) Clustername on which apps are running on ND. 
* `numberof_sites_onboarded`:(int) Number of sites on which particular app installed on ND. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of apps running on ND. 
 