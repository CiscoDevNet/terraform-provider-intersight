
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `compressed_vlan_port_count`:(string) The number of compressed VLAN ports on a Fabric Interconnect. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `uncompressed_vlan_port_count`:(string) The number of uncompressed VLAN ports on a Fabric Interconnect. 
* `vlan_port_limit`:(int) The maximum number of VLAN ports allowed on a Fabric Interconnect. 
