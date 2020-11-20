
---
layout: "intersight"
page_title: "Intersight: intersight_fc_physical_port"
sidebar_current: "docs-intersight-data-source-fc-physical-port"
description: |-
Physical fibre channel port present on a FI.
---

# Data Source: intersight_fc._physical_port
Physical fibre channel port present on a FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Administrator configured Speed applied on the port. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this port. 
* `b2b_credit`:(int) Buffer to Buffer credits of FC port. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `max_speed`:(string) Maximum Speed with which the port operates. 
* `mode`:(string) Mode information N_proxy, F or E associated to the Fibre Channel port. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `oper_speed`:(string) Operational Speed with which the port operates. 
* `oper_state`:(string) Operational state of this port (enabled/disabled). 
* `oper_state_qual`:(string) Reason for this port's Operational state. 
* `peer_dn`:(string) PeerDn for fibre channel physical port. 
* `port_channel_id`:(int) Port channel id of FC port channel created on FI switch. 
* `port_id`:(int) Switch physical port identifier. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) The role assigned to this port. 
* `slot_id`:(int) Switch expansion slot module identifier. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `transceiver_type`:(string) Transceiver type of a Fibre Channel port. 
* `vsan`:(int) Virtual San that is associated to the port. 
* `wwn`:(string) World Wide Name of a Fibre Channel port. 
