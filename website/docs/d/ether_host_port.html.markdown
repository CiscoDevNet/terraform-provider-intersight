
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `mac_address`:(string) Mac Address of a port in the Fabric Interconnect. 
* `mode`:(string) Operating mode of this port. 
* `module_id`:(int) Fabric extender identifier for this port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
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
