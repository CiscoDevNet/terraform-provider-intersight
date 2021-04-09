---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_vlan_port_info"
description: |-
  Vlan Port information of a Fabric Interconnect.
---

# Data Source: intersight_network_vlan_port_info
Vlan Port information of a Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_vlan_port_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_vlan_port_count`:(int) The number of available VLAN access ports on a Fabric Interconnect. 
* `account_moid`:(string) The Account ID for this managed object. 
* `border_vlan_port_count`:(int) The number of available VLAN border ports on a Fabric Interconnect. 
* `compressed_optimization_sets_value`:(int) The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library. 
* `compressed_vlan_port_count`:(string) The number of compressed VLAN ports on a Fabric Interconnect. 
* `compressed_vlan_port_count_value`:(int) The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_vlan_port_count`:(int) The total number of VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count`:(string) The number of uncompressed VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count_value`:(int) The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `vlan_port_limit`:(int) The maximum number of VLAN ports allowed on a Fabric Interconnect. 
 