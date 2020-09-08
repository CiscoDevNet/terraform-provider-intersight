
---
layout: "intersight"
page_title: "Intersight: intersight_network_element"
sidebar_current: "docs-intersight-data-source-network-element"
description: |-
A Unified Computing Systems (UCS) Fabric Interconnect.
---

# Data Source: intersight_network._element
A Unified Computing Systems (UCS) Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_evac_state`:(string) Administratively configured state of Fabric Evacuation feature, for this switch. 
* `admin_inband_interface_state`:(string) The administrative state of the network Element inband management interface. 
* `available_memory`:(string) Available memory (un-used) on this switch platform. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ethernet_mode`:(string) The user configured Ethernet operational mode for this switch (End-Host or Switching). 
* `fault_summary`:(int) The fault summary of the network Element out-of-band management interface. 
* `fc_mode`:(string) The user configured FC operational mode for this switch (End-Host or Switching). 
* `inband_ip_address`:(string) The IP address of the network Element inband management interface. 
* `inband_ip_gateway`:(string) The default gateway of the network Element inband management interface. 
* `inband_ip_mask`:(string) The network mask of the network Element inband management interface. 
* `inband_vlan`:(int) The VLAN ID of the network Element inband management interface. 
* `management_mode`:(string) The management mode of the fabric interconnect.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
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
* `total_memory`:(int) Total available memory on this switch platform. 
* `vendor`:(string) This field identifies the vendor of the given component. 
