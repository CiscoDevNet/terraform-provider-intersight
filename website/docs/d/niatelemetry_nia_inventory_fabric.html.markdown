---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_inventory_fabric"
description: |-
        Inventory Object available for Fabric-scoped attributes.

---

# Data Source: intersight_niatelemetry_nia_inventory_fabric
Inventory Object available for Fabric-scoped attributes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_inventory_fabric.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `anycast_gw_mac`:(string) Returns the aycast gateway mac. 
* `bgp_established_interface_count`:(int) Counts the number of BGP interfaces that are in established state. 
* `bgw_interface_up_count`:(int) Count number of active interfaces on border gateways. 
* `border_gateway_spine_count`:(int) Count number of border gateway spines in the fabric inventory. 
* `border_leaf_count`:(int) Count number of border leafs in the fabric inventory. 
* `cloudsec_autoconfig`:(bool) Cloudsec autoconfig details on the fabric. 
* `create_time`:(string) The time when this managed object was created. 
* `dci_subnet_range`:(string) Returns the dci subnet range. 
* `dci_subnet_target_mask`:(string) Returns the dci subnet target mask. 
* `dcnmtracker_enabled`:(bool) Returns the value of the dcnmtrackerEnabled field. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ebgp_evpn_link_up_count`:(int) Count number of ebgp evpn active interfaces. 
* `fabric_id`:(string) Uniquely identifies a fabric. 
* `fabric_name`:(string) Returns the value of the Name of a fabric. 
* `fabric_parent`:(string) Parent of the fabric on DCNM. 
* `fabric_technology`:(string) Fabric Technology details on the fabric. 
* `feature_ptp`:(string) PTP feature details on the fabric. 
* `is_bgw_present`:(bool) Checks if border gateway is present in the fabric inventory. 
* `is_enable_nxapi_http`:(bool) Check if NXAPI HTTP is enable or not on the fabric. 
* `is_enable_real_time_backup`:(bool) Check if real time backup is enable or not on the fabric. 
* `is_ngoam_enabled`:(bool) Returns if ngoam is enabled. 
* `is_scheduled_back_up_enabled`:(bool) Returns if the scheduled backup is enabled. 
* `leaf_count`:(int) Returns total number of leafs in the fabric. 
* `link_state_routing`:(string) Link state routing details on the fabric. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `network_deployment_count`:(int) No of networks deployed on a fabric. 
* `ntp_server_ip_list`:(string) NTP server IP List on the fabric. 
* `nxos_vni_bw_sites_count`:(int) Returns the count of vnis between sites. 
* `nxos_vrf_bw_sites_count`:(int) Returns the count of vrfs between sites. 
* `nxos_vrf_count`:(int) Returns the value of the nxosVrfCount field. 
* `replication_mode`:(string) Replication mode details on the fabric. 
* `rp_mode`:(string) RP Mode details on the fabric. 
* `serial`:(string) Serial number of device being inventoried. The serial number is unique per device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `software_image`:(string) Software image details on the fabric. 
* `spine_count`:(int) Returns total number of spines in the fabric. 
* `syslog_server_ip_list`:(string) Syslog server IP list on DCNM. 
* `syslog_sev`:(string) Syslog sev details on the fabric. 
* `template_name`:(string) Template name of the fabric on DCNM. 
* `vlan_vni_mappings`:(string) VLAN to VNI mappings configured in the DCNM. 
* `vni_ip_count`:(int) Count number of IP addresses configured in the DCNM networks. 
* `vrf_deployment_count`:(int) No of vrfs deployed on a fabric. 
 
