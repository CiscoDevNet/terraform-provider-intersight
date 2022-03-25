---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_capacity_runway"
description: |-
        Entity representing the new capacity runway based on recommendations.

---

# Data Source: intersight_recommendation_capacity_runway
Entity representing the new capacity runway based on recommendations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_capacity_runway.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `additional_capacity`:(int) Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_updated_time`:(string) The time when the recommendation was last updated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the recommendation. 
* `period`:(int) Number of months in future for which recommendation is provided for. 
* `requirement_met`:(bool) Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false. 
* `runway`:(int) This represents the new runway, that is the number of days remaining before the cluster's storage utilization reaches the recommended capacity limit after the recommended hardware is added. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_capacity`:(int) Total capacity of the cluster after the recommended hardware is added. 
* `unit`:(string) Unit for the new capacity.* `TB` - The Enum value TB represents that the measurement unit is in terabytes.* `MB` - The Enum value MB represents that the measurement unit is in megabytes. 
 
