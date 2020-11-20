
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_switch_card"
sidebar_current: "docs-intersight-data-source-equipment-switch-card"
description: |-
Fixed / Removable module on a Fabric Interconnect / Switch.
---

# Data Source: intersight_equipment._switch_card
Fixed / Removable module on a Fabric Interconnect / Switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Detailed description of this switch hardware. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ethernet_switching_mode`:(string) The user configured Ethernet switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fc_switching_mode`:(string) The user configured FC switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_ports`:(int) Number of ports present in this switch hardware. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `out_of_band_ip_address`:(string) Field specifies this Switch's Out-of-band IP address. 
* `out_of_band_ip_gateway`:(string) Field specifies this Switch's default gateway for the out-of-band management interface. 
* `presence`:(string) Presence for this switch hardware. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `slot_id`:(int) Slot identifier of the local Switch slot Interface. 
* `state`:(string) Operational state of the switch hardware. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `vendor`:(string) This field identifies the vendor of the given component. 
