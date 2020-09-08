
---
layout: "intersight"
page_title: "Intersight: intersight_ether_network_port"
sidebar_current: "docs-intersight-data-source-ether-network-port"
description: |-
Model contains the details of the ethernet port connected to the FI side.
---

# Data Source: intersight_ether._network_port
Model contains the details of the ethernet port connected to the FI side.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `module_id`:(int) Febric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Operational state of an Interface. 
* `peer_dn`:(string) Peer DN for network host port of fabric extender. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Network Port Speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
