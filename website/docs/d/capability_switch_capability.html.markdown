---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_capability"
description: |-
  Type to represent additional switch specific capabilities.
---

# Data Source: intersight_capability_switch_capability
Type to represent additional switch specific capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_switch_capability.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `default_fcoe_vlan`:(int) Default Fcoe VLAN associated with this switch. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dynamic_vifs_supported`:(bool) Dynamic VIFs support on this switch. 
* `fan_modules_supported`:(bool) Fan Modules support on this switch. 
* `fc_uplink_ports_auto_negotiation_supported`:(bool) Fc Uplink ports auto negotiation speed support on this switch. 
* `locator_beacon_supported`:(bool) Locator Beacon LED support on this switch. 
* `max_ports`:(int) Maximum allowed physical ports on this switch. 
* `max_slots`:(int) Maximum allowed physical slots on this switch. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `sereno_netflow_supported`:(bool) Sereno Adaptor with Netflow support on this switch. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `unified_rule`:(string) The Slider rule for Unified ports on this switch. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
 