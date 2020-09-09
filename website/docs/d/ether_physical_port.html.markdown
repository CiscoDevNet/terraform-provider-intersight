
---
layout: "intersight"
page_title: "Intersight: intersight_ether_physical_port"
sidebar_current: "docs-intersight-data-source-ether-physical-port"
description: |-
Physical ethernet port present on a FI.
---

# Data Source: intersight_ether._physical_port
Physical ethernet port present on a FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Administratively configured speed for this port. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port. 
* `aggregate_port_id`:(int) Breakout port member in the Fabric Interconnect. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `license_grace`:(string) The number of days this port's license has been in Grace Period for. 
* `license_state`:(string) The state of the port's licensing. 
* `mac_address`:(string) Mac Address of a port in the Fabric Interconnect. 
* `mode`:(string) Operating mode of this port. 
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
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `transceiver_type`:(string) Transceiver model attached to a port in the Fabric Interconnect. 
