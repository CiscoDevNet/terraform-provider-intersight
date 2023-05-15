---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_inc_license_count"
description: |-
        Customer operation object to request reservation code.

---

# Data Source: intersight_license_inc_license_count
Customer operation object to request reservation code.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_inc_license_count.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `premier100_gfx_count`:(int) The total number of devices claimed in the premier 100G fixed tier Intersight Nexus Cloud. 
* `premier10_gfx_count`:(int) The total number of devices claimed in the premier 10G fixed tier Intersight Nexus Cloud. 
* `premier1_gfx_count`:(int) The total number of devices claimed in the premier 1G fixed tier Intersight Nexus Cloud. 
* `premier_centralized_mod8_slot_count`:(int) The total number of devices claimed in the CentralizedMod8Slot premier tier Intersight Nexus Cloud. 
* `premier_d2_ops_fixed_count`:(int) The total number of devices claimed in the D2Ops Fixed premier tier Intersight Nexus Cloud. 
* `premier_d2_ops_mod_count`:(int) The total number of devices claimed in the D2Ops modular premier tier Intersight Nexus Cloud. 
* `premier_distributed_mod8_slot_count`:(int) The total number of devices claimed in the DistributedMod8Slot premier tier Intersight Nexus Cloud. 
* `premier_mod4_slot_count`:(int) The total number of devices claimed in the modular 4 slot premier tier Intersight Nexus Cloud. 
* `premier_mod8_slot_count`:(int) The total number of devices claimed in the modular 8 slot premier tier Intersight Nexus Cloud. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
