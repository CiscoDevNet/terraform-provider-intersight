
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `module_id`:(int) Febric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `oper_state`:(string) Operational state of an Interface. 
* `peer_dn`:(string) Peer DN for network host port of fabric extender. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Network Port Speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
