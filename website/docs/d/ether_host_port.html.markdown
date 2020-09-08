
---
layout: "intersight"
page_title: "Intersight: intersight_ether_host_port"
sidebar_current: "docs-intersight-data-source-ether-host-port"
description: |-
Model object contains the details of host port available on IO card or fabric extender.
---

# Data Source: intersight_ether._host_port
Model object contains the details of host port available on IO card or fabric extender.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `mac_address`:(string) Mac Address of a port in the Fabric Interconnect. 
* `mode`:(string) Operating mode of this port. 
* `module_id`:(int) Fabric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_speed`:(string) Current Operational speed for this port. 
* `oper_state`:(string) Operational state of this port (enabled/disabled). 
* `oper_state_qual`:(string) Reason for this port's Operational state. 
* `peer_dn`:(string) PeerDn for ethernet physical port. 
* `port_channel_id`:(int) Port channel id for port channel created on FI switch. 
* `port_id`:(int) Switch physical port identifier. 
* `port_type`:(string) Defines the transport type for this port (ethernet OR fc). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The role assigned to this port. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `speed`:(string) Host Port Speed of IO card or fabric extender. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `transceiver_type`:(string) Transceiver model attached to a port in the Fabric Interconnect. 
