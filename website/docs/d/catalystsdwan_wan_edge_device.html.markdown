---
subcategory: "catalystsdwan"
layout: "intersight"
page_title: "Intersight: intersight_catalystsdwan_wan_edge_device"
description: |-
        Details for the Catalyst SDWAN WAN Edge entities.

---

# Data Source: intersight_catalystsdwan_wan_edge_device
Details for the Catalyst SDWAN WAN Edge entities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_catalystsdwan_wan_edge_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bfd_sessions`:(string) Total number of BFD sessions from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `bfd_sessions_up`:(string) Number of BFD sessions in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `config_device_status`:(string) Synchronization status of the WAN Edge Device configuration with Cisco Catalyst SD-WAN Manager. 
* `config_locked`:(string) Indicates if the device configuration is locked. 
* `control_connections`:(string) Total number of control connections from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `control_connections_up`:(string) Number of control connections in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `create_time`:(string) The time when this managed object was created. 
* `device_health`:(string) Health status of the WAN Edge Device. 
* `device_id`:(string) Device ID of the WAN Edge Device. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `device_state`:(string) Device state of the WAN Edge Device. 
* `device_type`:(string) The categorization of the device type. Optional parameter to categorize devices by product type. For example, Meraki device types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, and secureConnect. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hardware_version`:(string) The hardware version of the device. 
* `host_name`:(string) Host name of the WAN Edge Device. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `management_port_hw_addr`:(string) MAC address of the WAN Edge Device management port. 
* `management_port_ip`:(string) IP address of the WAN Edge Device management port. 
* `management_port_subnet`:(string) Subnet of the WAN Edge Device management port. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `organization_name`:(string) The WAN Edge Device organization name. 
* `personality`:(string) Personality of the WAN Edge Device as seen in the Cisco Catalyst SD-WAN Manager (e.g., vEdge, cEdge). 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `reachability`:(string) Reachability of the WAN Edge Device.* `Unknown` - Reachability status of the Catalyst SDWAN device from SDWAN Manager is unknown.* `Reachable` - The Catalyst SDWAN device is reachable from SDWAN Manager.* `Unreachable` - The Catalyst SDWAN device is unreachable from SDWAN Manager.* `Authentication Failed` - Authentication failed when SDWAN Manager attempted to connect to the Catalyst SDWAN device. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `sdwan_user_tags`:(string) User defined tags associated with the WAN Edge Device from Cisco Catalyst SD-WAN Manager. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(int) Site ID of the WAN Edge Device. 
* `site_name`:(string) Site name of the WAN Edge Device. 
* `sp_organization_name`:(string) The WAN Edge Device sp organization name. 
* `system_ip`:(string) System IP of the WAN Edge Device. 
* `tlocs`:(string) Total number of TLOCs from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `tlocs_up`:(string) Number of TLOCs in up state from the WAN Edge Device to Cisco Catalyst SD-WAN Manager. 
* `uptime`:(string) Uptime of the WAN Edge Device. 
* `uuid`:(string) Unique identity of the device. 
* `validity`:(string) Validity of the WAN Edge Device.* `Unknown` - Validity of the Catalyst SDWAN device configuration is unknown.* `Valid` - Configuration on the Catalyst SDWAN device is valid.* `Invalid` - Configuration on the Catalyst SDWAN device is invalid.* `Prestaging` - Configuration on the Catalyst SDWAN device is in pre-staging state.* `Staging` - Configuration on the Catalyst SDWAN device is in staging state.* `Config Init` - Configuration on the Catalyst SDWAN device is in initialization state. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Current running software version of the device. 
 
