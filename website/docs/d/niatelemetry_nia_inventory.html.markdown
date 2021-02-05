---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory"
description: |-
  Inventory object available per device scope. This common object holds a device level information.
---

# Data Source: intersight_niatelemetry_nia_inventory
Inventory object available per device scope. This common object holds a device level information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cpu`:(float) CPU usage of device being inventoried. This determines the percentage of CPU resources used. 
* `crash_reset_logs`:(string) Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system. 
* `customer_device_connector`:(string) Returns the value of the customerDeviceConnector field. 
* `device_discovery`:(string) Returns the value of the deviceDiscovery field. 
* `device_health`:(int) Returns the device health. 
* `device_id`:(string) Returns the value of the deviceId field. 
* `device_name`:(string) Name of device being inventoried. The name the user assigns to the device is inventoried here. 
* `device_type`:(string) Type of device being inventoried. This determines whether the device is a controller, leaf or spine. 
* `device_up_time`:(int) Returns the device up time. 
* `dn`:(string) Dn for the inventories present. 
* `fex_count`:(int) Number of fabric extendors utilized. 
* `infra_wi_node_count`:(int) Number of appliances as physical device that are wired into the cluster. 
* `ip_address`:(string) The IP address of the device being inventoried. 
* `is_virtual_node`:(string) Flag to specify if the node is virtual. 
* `log_in_time`:(string) Last log in time device being inventoried. This determines the last login time on the device. 
* `log_out_time`:(string) Last log out time of device being inventoried. This determines the last logout time on the device. 
* `mac_sec_count`:(int) Number of Macsec configured interfaces on a TOR. 
* `mac_sec_fab_count`:(int) Number of Macsec configured interfaces on a Spine. 
* `macsec_total_count`:(int) Number of total Macsec configured interfaces for all nodes. 
* `memory`:(int) Memory usage of device being inventoried. This determines the percentage of memory resources used. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(string) The ID of the device being inventoried. 
* `nxos_evpn_mac_routes`:(int) Returns the total number of evpn mac routes. 
* `nxos_telnet`:(string) Returns the value of the nxosTelnet field. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `route_prefix_count`:(int) Total nuumber of v4 and v6 routes per node. 
* `route_prefix_v4_count`:(int) Number of v4 routes per node. 
* `route_prefix_v6_count`:(int) Number of v6 routes per node. 
* `serial`:(string) Serial number of device being invetoried. The serial number is unique per device and will be used as the key. 
* `software_download`:(string) Last software downloaded of device being inventoried. This determines if software download API was used. 
* `system_up_time`:(string) The amount of time that the device being inventoried been up. 
* `nr_version`:(string) Software version of device being inventoried. The various software version values for each device are available on cisco.com. 
