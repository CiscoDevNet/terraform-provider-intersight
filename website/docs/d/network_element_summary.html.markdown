---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_element_summary"
description: |-
        View MO which aggregates information pertaining to a network element from mutiple MOs.

---

# Data Source: intersight_network_element_summary
View MO which aggregates information pertaining to a network element from mutiple MOs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_element_summary.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_evac_state`:(string) Administratively configured state of Fabric Evacuation feature, for this switch.* `` - Evacuation state of the switch is unknown.* `enabled` - Evacuation state of the switch is enabled.* `disabled` - Evacuation state of the switch is disabled.* `applying` - Evacuation state of the switch when evacuation is in progress.* `on` - Evacuation state of the switch is enabled.* `off` - Evacuation state of the switch is disabled.* `N/A` - Evacuation state of the switch is not applicable. 
* `admin_inband_interface_state`:(string) The administrative state of the network Element inband management interface. 
* `available_memory`:(string) Available memory (un-used) on this switch platform. 
* `bundle_version`:(string) Running firmware bundle information. 
* `chassis`:(string) Chassis IP of the switch. 
* `conf_mod_ts`:(string) Configuration modified timestamp of the switch. 
* `conf_mod_ts_backup`:(string) Configuration modified backup timestamp of the switch. 
* `connection_status`:(string) Connection status of the switch. 
* `create_time`:(string) The time when this managed object was created. 
* `default_domain`:(string) The default domain name configured on the switch. 
* `device_mo_id`:(string) The MoId of the registered device that coresponds to the server. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ethernet_mode`:(string) The user configured Ethernet operational mode for this switch (End-Host or Switching). 
* `ethernet_switching_mode`:(string) The user configured Ethernet operational mode for this switch (End-Host or Switching).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `fault_summary`:(int) The fault summary for the server. 
* `fc_mode`:(string) The user configured FC operational mode for this switch (End-Host or Switching). 
* `fc_switching_mode`:(string) The user configured FC operational mode for this switch (End-Host or Switching).* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
* `firmware`:(string) Running firmware information. 
* `firmware_version`:(string) Running firmware information. 
* `fpga_upgrade_needed`:(bool) The flag to check vulnerability with secure boot technology. 
* `inband_ip_address`:(string) The IP address of the network Element inband management interface. 
* `inband_ip_gateway`:(string) The default gateway of the network Element inband management interface. 
* `inband_ip_mask`:(string) The network mask of the network Element inband management interface. 
* `inband_vlan`:(int) The VLAN ID of the network Element inband management interface. 
* `inter_cluster_link_state`:(string) The intercluster link state of the switch.* `Unknown` - The operational state of the link is not known.* `Up` - The operational state of the link is up.* `Down` - The operational state of the link is down.* `Degraded` - The link is operational but degraded. This state is applicable to port channels when any one of the member links is down. 
* `ipv4_address`:(string) IP version 4 address is saved in this property. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `management_mode`:(string) The management mode of the fabric interconnect.* `IntersightStandalone` - Intersight Standalone mode of operation.* `UCSM` - Unified Computing System Manager mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The admin configured name of the system. 
* `num_ether_ports`:(int) Total number of Ethernet ports. 
* `num_ether_ports_configured`:(int) Total number of configured Ethernet ports. 
* `num_ether_ports_link_up`:(int) Total number of Ethernet ports which are UP. 
* `num_expansion_modules`:(int) Total number of expansion modules. 
* `num_fc_ports`:(int) Total number of FC ports. 
* `num_fc_ports_configured`:(int) Total number of configured FC ports. 
* `num_fc_ports_link_up`:(int) Total number of FC ports which are UP. 
* `oper_evac_state`:(string) Operational state of the Fabric Evacuation feature, for this switch.* `` - Evacuation state of the switch is unknown.* `enabled` - Evacuation state of the switch is enabled.* `disabled` - Evacuation state of the switch is disabled.* `applying` - Evacuation state of the switch when evacuation is in progress.* `on` - Evacuation state of the switch is enabled.* `off` - Evacuation state of the switch is disabled.* `N/A` - Evacuation state of the switch is not applicable. 
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
* `part_number`:(string) Part number of the switch. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `reserved_vlan_start_id`:(int) The reserved VLAN start ID of the Network Element. A block of 128 VLANs are reserved for internal use and cannot be used for carrying network traffic. 
* `revision`:(string) The product revision of the switch. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_object_type`:(string) Stores the source object type used to fill the properties of this object. 
* `status`:(string) The status of the switch. 
* `switch_id`:(string) The Switch Id of the network Element. 
* `switch_profile_name`:(string) The name of switch profile associated with the switch. 
* `switch_type`:(string) The Switch type that the network element is a part of.* `FabricInterconnect` - The default Switch type of UCSM and IMM mode devices.* `NexusDevice` - Switch type of Nexus devices.* `MDSDevice` - Switch type of Nexus MDS devices. 
* `system_up_time`:(string) System up time of the switch. 
* `thermal`:(string) The Thermal status of the fabric interconnect.* `unknown` - The default state of the sensor (in case no data is received).* `ok` - State of the sensor indicating the sensor's temperature range is okay.* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.* `upper-critical` - State of the sensor indicating that the temperature is above normal range.* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.* `lower-critical` - State of the sensor indicating that the temperature is below normal range.* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range. 
* `total_memory`:(int) Total available memory on this switch platform. 
* `user_label`:(string) The user defined label assigned to the server. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Firmware version of the switch. 
 
