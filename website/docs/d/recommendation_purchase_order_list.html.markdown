---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_purchase_order_list"
description: |-
        Entity representing the list of items in csv format for the current recommendation. This can be used to generate an estimate by uploading it to Cisco Commerce Workspace.

---

# Data Source: intersight_recommendation_purchase_order_list
Entity representing the list of items in csv format for the current recommendation. This can be used to generate an estimate by uploading it to Cisco Commerce Workspace.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_purchase_order_list.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `item_list`:(string) The comma seperated list of items for the current recommendation. This can be used to generate an estimate by uploading it to Cisco Commerce Workspace. 
* `last_updated_time`:(string) The time when the recommendation was last updated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the recommendation. 
* `requirement_met`:(bool) Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
