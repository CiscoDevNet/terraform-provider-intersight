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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `controller_ip_address`:(string) Returns controller's IP address details. 
* `cpu`:(float) CPU usage of device being inventoried. This determines the percentage of CPU resources used. 
* `crash_reset_logs`:(string) Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system. 
* `create_time`:(string) The time when this managed object was created. 
* `customer_device_connector`:(string) Returns the value of the customerDeviceConnector field. 
* `dcnm_license_state`:(string) Returns the License state of the device. 
* `device_discovery`:(string) Returns the value of the deviceDiscovery field. 
* `device_health`:(int) Returns the device health. 
* `device_id`:(string) Returns the value of the deviceId field. 
* `device_name`:(string) Name of device being inventoried. The name the user assigns to the device is inventoried here. 
* `device_type`:(string) Type of device being inventoried. This determines whether the device is a controller, leaf or spine. 
* `device_up_time`:(int) Returns the device up time. 
* `dn`:(string) Dn for the inventories present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_name`:(string) Name of the fabric of the device being inventoried. 
* `fex_count`:(int) Number of fabric extendors utilized. 
* `infra_wi_node_count`:(int) Number of appliances as physical device that are wired into the cluster. 
* `ip_address`:(string) The IP address of the device being inventoried. 
* `is_virtual_node`:(string) Flag to specify if the node is virtual. 
* `last_reboot_time`:(string) Returns the last reboot Time of the device. 
* `last_reset_reason`:(string) Returns the last reset reason of the device. 
* `license_type`:(string) Returns the License type of the device. 
* `log_in_time`:(string) Last log in time device being inventoried. This determines the last login time on the device. 
* `log_out_time`:(string) Last log out time of device being inventoried. This determines the last logout time on the device. 
* `mac_sec_count`:(int) Number of Macsec configured interfaces on a TOR. 
* `mac_sec_fab_count`:(int) Number of Macsec configured interfaces on a Spine. 
* `macsec_total_count`:(int) Number of total Macsec configured interfaces for all nodes. 
* `memory`:(int) Memory usage of device being inventoried. This determines the percentage of memory resources used. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nexus_cloud_membership_status`:(bool) Returns if site has been onboarded to nexus cloud or not. 
* `node_id`:(string) The ID of the device being inventoried. 
* `nxos_dci_interface_status`:(string) Returns the status of dci interface configured. 
* `nxos_nve_interface_status`:(string) Returns the value of the nxosNveInterface field. 
* `nxos_ospf_neighbors`:(int) Total number of ospf neighbors per switch in DCNM. 
* `nxos_pim_neighbors`:(string) Total number of pim neighbors per switch in DCNM. 
* `nxos_telnet`:(string) Returns the value of the nxosTelnet field. 
* `nxos_total_routes`:(int) Total number of routes configured in the DCNM. 
* `record_type`:(string) Type of record DCNM / APIC / SE / Nexus Switch. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `route_prefix_count`:(int) Total nuumber of v4 and v6 routes per node. 
* `route_prefix_v4_count`:(int) Number of v4 routes per node. 
* `route_prefix_v6_count`:(int) Number of v6 routes per node. 
* `serial`:(string) Serial number of device being invetoried. The serial number is unique per device and will be used as the key. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `siteuuid`:(string) Returns the uuid of the Nexus Cloud site associated to the inventory object. 
* `smart_account_id`:(int) Returns the value of the smartAccountId/CustomerId field. 
* `software_download`:(string) Last software downloaded of device being inventoried. This determines if software download API was used. 
* `system_up_time`:(string) The amount of time that the device being inventoried been up. 
* `nr_version`:(string) Software version of device being inventoried. The various software version values for each device are available on cisco.com. 
 
