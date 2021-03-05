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
* `border_vlan_port_count`:(int) The number of available VLAN border ports on a Fabric Interconnect. 
* `compressed_optimization_sets_value`:(int) The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library. 
* `compressed_vlan_port_count`:(string) The number of compressed VLAN ports on a Fabric Interconnect. 
* `compressed_vlan_port_count_value`:(int) The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `total_vlan_port_count`:(int) The total number of VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count`:(string) The number of uncompressed VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count_value`:(int) The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `vlan_port_limit`:(int) The maximum number of VLAN ports allowed on a Fabric Interconnect. 
 