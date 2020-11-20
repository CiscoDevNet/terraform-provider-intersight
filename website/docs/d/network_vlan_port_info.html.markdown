
---
layout: "intersight"
page_title: "Intersight: intersight_network_vlan_port_info"
sidebar_current: "docs-intersight-data-source-network-vlan-port-info"
description: |-
Vlan Port information of a Fabric Interconnect.
---

# Data Source: intersight_network._vlan_port_info
Vlan Port information of a Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_vlan_port_count`:(int) The number of available VLAN access ports on a Fabric Interconnect. 
* `border_vlan_port_count`:(int) The number of available VLAN border ports on a Fabric Interconnect. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `compressed_optimization_sets_value`:(int) The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library. 
* `compressed_vlan_port_count`:(string) The number of compressed VLAN ports on a Fabric Interconnect. 
* `compressed_vlan_port_count_value`:(int) The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `total_vlan_port_count`:(int) The total number of VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count`:(string) The number of uncompressed VLAN ports on a Fabric Interconnect. 
* `uncompressed_vlan_port_count_value`:(int) The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. 
* `vlan_port_limit`:(int) The maximum number of VLAN ports allowed on a Fabric Interconnect. 
