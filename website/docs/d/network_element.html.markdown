---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_element"
description: |-
  A Unified Computing Systems (UCS) Fabric Interconnect.
---

# Data Source: intersight_network_element
A Unified Computing Systems (UCS) Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_element.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_evac_state`:(string) Administratively configured state of Fabric Evacuation feature, for this switch. 
* `admin_inband_interface_state`:(string) The administrative state of the network Element inband management interface. 
* `available_memory`:(string) Available memory (un-used) on this switch platform. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ethernet_mode`:(string) The user configured Ethernet operational mode for this switch (End-Host or Switching). 
* `ethernet_switching_mode`:(string) The user configured Ethernet operational mode for this switch (End-Host or Switching).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fault_summary`:(int) The fault summary of the network Element out-of-band management interface. 
* `fc_mode`:(string) The user configured FC operational mode for this switch (End-Host or Switching). 
* `fc_switching_mode`:(string) The user configured FC operational mode for this switch (End-Host or Switching).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `inband_ip_address`:(string) The IP address of the network Element inband management interface. 
* `inband_ip_gateway`:(string) The default gateway of the network Element inband management interface. 
* `inband_ip_mask`:(string) The network mask of the network Element inband management interface. 
* `inband_vlan`:(int) The VLAN ID of the network Element inband management interface. 
* `management_mode`:(string) The management mode of the fabric interconnect.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_evac_state`:(string) Operational state of the Fabric Evacuation feature, for this switch. 
* `operability`:(string) The switch's current overall operational/health state. 
* `out_of_band_ip_address`:(string) The IP address of the network Element out-of-band management interface. 
* `out_of_band_ip_gateway`:(string) The default gateway of the network Element out-of-band management interface. 
* `out_of_band_ip_mask`:(string) The network mask of the network Element out-of-band management interface. 
* `out_of_band_ipv4_address`:(string) The IPv4 address of the network Element out-of-band management interface. 
* `out_of_band_ipv4_gateway`:(string) The default IPv4 gateway of the network Element out-of-band management interface. 
* `out_of_band_ipv4_mask`:(string) The network mask of the network Element out-of-band management interface. 
* `out_of_band_ipv6_address`:(string) The IPv6 address of the network Element out-of-band management interface. 
* `out_of_band_ipv6_gateway`:(string) The default IPv6 gateway of the network Element out-of-band management interface. 
* `out_of_band_ipv6_prefix`:(string) The network mask of the network Element out-of-band management interface. 
* `out_of_band_mac`:(string) The MAC address of the network Element out-of-band management interface. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `switch_id`:(string) The Switch Id of the network Element. 
* `thermal`:(string) The Thermal status of the fabric interconnect.* `unknown` - The default state of the sensor (in case no data is received).* `ok` - State of the sensor indicating the sensor's temperature range is okay.* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.* `upper-critical` - State of the sensor indicating that the temperature is above normal range.* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.* `lower-critical` - State of the sensor indicating that the temperature is below normal range.* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range. 
* `total_memory`:(int) Total available memory on this switch platform. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 