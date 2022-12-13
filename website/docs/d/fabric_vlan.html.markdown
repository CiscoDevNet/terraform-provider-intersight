---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vlan"
description: |-
        Configuration object for Virtual LAN.

---

# Data Source: intersight_fabric_vlan
Configuration object for Virtual LAN.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_vlan.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_allow_on_uplinks`:(bool) Enable to automatically allow this VLAN on all uplinks. Disable must be specified for Disjoint Layer 2 VLAN configuration. Default VLAN-1 cannot be configured as Disjoint Layer 2 VLAN. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_native`:(bool) Used to define whether this VLAN is to be classified as 'native' for traffic in this FI. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The 'name' used to identify this VLAN. 
* `primary_vlan_id`:(int) The Primary VLAN ID of the VLAN, if the sharing type of the VLAN is Isolated or Community. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sharing_type`:(string) The sharing type of this VLAN.* `None` - This represents a regular VLAN.* `Primary` - This represents a primary VLAN.* `Isolated` - This represents an isolated VLAN.* `Community` - This represents a community VLAN. 
* `vlan_id`:(int) The identifier for this Virtual LAN. 
 
