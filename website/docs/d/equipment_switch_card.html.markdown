---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_switch_card"
description: |-
  Fixed / Removable module on a Fabric Interconnect / Switch.
---

# Data Source: intersight_equipment_switch_card
Fixed / Removable module on a Fabric Interconnect / Switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_switch_card.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Detailed description of this switch hardware. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ethernet_switching_mode`:(string) The user configured Ethernet switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fc_switching_mode`:(string) The user configured FC switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_ports`:(int) Number of ports present in this switch hardware. 
* `out_of_band_ip_address`:(string) Field specifies this Switch's Out-of-band IP address. 
* `out_of_band_ip_gateway`:(string) Field specifies this Switch's default gateway for the out-of-band management interface. 
* `presence`:(string) Presence for this switch hardware. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `slot_id`:(int) Slot identifier of the local Switch slot Interface. 
* `state`:(string) Operational state of the switch hardware. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `thermal`:(string) The Thermal status of the fabric interconnect.* `unknown` - The default state of the sensor (in case no data is received).* `ok` - State of the sensor indicating the sensor's temperature range is okay.* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.* `upper-critical` - State of the sensor indicating that the temperature is above normal range.* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.* `lower-critical` - State of the sensor indicating that the temperature is below normal range.* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 