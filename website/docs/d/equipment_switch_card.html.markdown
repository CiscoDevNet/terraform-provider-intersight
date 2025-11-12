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
To access the ith object of the results obtained, use `data.intersight_equipment_switch_card.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `base_mac`:(string) The starting (base) MAC address of the switch hardware like \ d0-e0-42-87-39-00\ . 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed description of this switch hardware. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_mac`:(string) The ending MAC address of the switch hardware like \ d0-e0-42-87-39-5f\ . 
* `ethernet_switching_mode`:(string) The user configured Ethernet switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fc_switching_mode`:(string) The user configured FC switching mode for this switch (End-Host or Switch).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `hw_version`:(string) The Hardware version of the switch hardware. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `jumbo_frame_enabled`:(bool) Jumbo Frame configuration for the switch. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the switch card like Line Card-1. 
* `num_ports`:(int) Number of ports present in this switch hardware. 
* `out_of_band_ip_address`:(string) Field specifies this Switch's Out-of-band IP address. 
* `out_of_band_ip_gateway`:(string) Field specifies this Switch's default gateway for the out-of-band management interface. 
* `out_of_band_ip_mask`:(string) Field specifies the Netmask for this Switch's Out-of-band IP address. 
* `out_of_band_mac`:(string) The MAC address of the Switch's out-of-band management interface. 
* `part_number`:(string) The part number of the switch hardware. 
* `power_state`:(string) Power state of the switch hardware.* `unknown` - The power state of the switch hardware is unknown.* `off` - The power state of the switch hardware is off.* `on` - The power state of the switch hardware is on.* `deny` - The power state of the switch hardware is deny.* `multi-boot-fail` - The power state of the switch hardware is multi-boot-fail. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot identifier of the SwitchCard within the Switch. 
* `state`:(string) Operational state of the switch hardware. 
* `status`:(string) The connection status of the switch hardware like up/down.* `Down` - Connection status of the switch card is down.* `Up` - Connection status of the switch card is up. 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `thermal`:(string) The Thermal status of the fabric interconnect.* `unknown` - The default state of the sensor (in case no data is received).* `ok` - State of the sensor indicating the sensor's temperature range is okay.* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.* `upper-critical` - State of the sensor indicating that the temperature is above normal range.* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.* `lower-critical` - State of the sensor indicating that the temperature is below normal range.* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range. 
* `type`:(string) Type of the switch card based on the capability like 4 Gbps or 2 Gbps type etc. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
