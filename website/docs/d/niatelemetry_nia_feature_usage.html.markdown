---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_feature_usage"
description: |-
        Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.

---

# Data Source: intersight_niatelemetry_nia_feature_usage
Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nia_feature_usage.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aaa_ldap_provider_count`:(int) Returns the total number of AAA Ldap Providers. 
* `aaa_radius_provider_count`:(int) Returns the total number of AAA Radius Providers. 
* `aaa_tacacs_provider_count`:(int) Returns the total number of AAA Tacacs Providers. 
* `account_moid`:(string) The Account ID for this managed object. 
* `apic_cluster_health`:(string) Cluster health for the APIC controller. 
* `apic_count`:(int) Number of APIC controllers. This determines the value of controllers for the fabric. 
* `apic_is_telnet_enabled`:(bool) Returns if telnet is enabled on APIC. 
* `apic_ntp_count`:(int) Count of NTP servers configured on APIC. 
* `apic_snmp_community_count`:(int) Number of SNMP communities configured on APIC. 
* `apic_sys_log_grp_count`:(int) Number of logging groups configured on APIC. 
* `apic_sys_log_src_count`:(int) Number of logging sources configured on APIC. 
* `app_center_count`:(int) ACI APPs feature usage scale. 
* `ave`:(string) AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled. 
* `bd_count`:(int) Number of BDs. This determines the total number of Broadcast Domains across the fabric. 
* `callhome_smart_group_count`:(int) Number of call home smart monitoring policies on the fabric. 
* `cloud_sec_peer_count`:(int) Number of Cloudsec SA peers. 
* `comp_hv_count`:(int) Number of compute hypervisors on the fabric. 
* `config_exportp_count`:(int) Number of system backup configure export policies on the fabric. 
* `config_job_count`:(int) Number of system backup configure jobs on the fabric. 
* `consistency_checker_app`:(string) Consistency checker application usage. This determines if the fabric has Consistency checker application installed. 
* `contract_count`:(int) Number of contracts. This determines the total number of Contracts configured across the fabric. 
* `create_time`:(string) The time when this managed object was created. 
* `dns_count`:(int) DNS feature usage. This determines the total number of DNS configurations across the fabric. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `eigrp_count`:(int) Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric. 
* `epg_count`:(int) Number of End Point Groups. This determines the total number of End Point Groups across the fabric. 
* `fabric_module_count`:(int) Returns the total number of fabric module slots. 
* `fabric_setupp_count`:(int) Number of Multi-Pods per fabric. 
* `fcoe_nport_count`:(int) Total number of FCoE N-Port for DOM, VSAn, and VLAN. 
* `fcoe_nport_dom_count`:(int) Number of FCoE N-Port DOM. 
* `fcoe_nport_vlan_count`:(int) Number of FCoE N-Port VLAN. 
* `fcoe_nport_vsan_count`:(int) Number of FCoE N-Port VSAN. 
* `fv_sla_def_count`:(int) Number of Internet Protocol Service Level Agreements Monitoring policy objects for object tracking. 
* `hsrp_count`:(int) Hsrp feature usage. This determines the total number of HSRP sessions across the fabric. 
* `ibgp_count`:(int) Ibgp feature usage. This determines the total number of BGP sessions across the fabric. 
* `igmp_access_list_count`:(int) IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric. 
* `igmp_snoop`:(string) IGMP Snooping feature usage. This determines if this feature is enabled or disabled. 
* `ip_epg_count`:(int) Number of IP based End Point Groups. This determines the total number of IP End Point Groups across the fabric. 
* `is_bgp_route_reflectors_feature_used`:(bool) BGP route reflector usage on APIC. 
* `is_bridge_domains_feature_used`:(bool) Brodge domains feature usage on APIC controller. 
* `is_common_local_user_name`:(bool) Returns value of isCommonLocalUserName field. 
* `is_contracts_feature_used`:(bool) Contracts feature usage on APIC controller. 
* `is_epg_feature_used`:(bool) EPG feature usage on APIC controller. 
* `is_filters_feature_used`:(bool) Filters feature usage on APIC. 
* `is_http_configured`:(bool) Returns if HTTP is configured. 
* `is_https_configured`:(bool) Returns if HTTPS is configured. 
* `is_ntp_feature_used`:(bool) NTP feature usage on APIC controller. 
* `is_ptp_feature_used`:(bool) Ptp feature usage on APIC. 
* `is_synce_feature_used`:(bool) Synce feature usage on APIC. 
* `is_tech_support_collected`:(string) Status of techsupport collection. 
* `is_tenants_feature_used`:(bool) Tenants feature usage on APIC. 
* `is_vrfs_feature_used`:(bool) VRF feature usage on APIC controller. 
* `isis_count`:(int) Isis feature usage. This determines the total number of ISIS sessions across the fabric. 
* `l2_multicast`:(string) L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric. 
* `latency_ptp_mode`:(string) Returns the Latency ptp mode for the controller. 
* `leaf_count`:(int) Number of Leafs. This determines the total number of Leaf switches in the fabric. 
* `local_username_count`:(int) Returns count of local users. 
* `login_block_duration`:(int) Returns login block duration value. 
* `login_max_failed_attempts`:(int) Returns the maximum failed attempts on login. 
* `login_max_failed_attempts_window`:(int) Returns the maximum failed attempt windows on login. 
* `maintenance_mode_count`:(int) Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode. 
* `management_over_v6_count`:(int) Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `microsoft_useg_vmm_ep_pd_count`:(int) Number of Microsoft microsegmentation VmmEpPD objects. Ensures that Microsoft was configured. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `net_flow_count`:(int) Number of Netflow monitor policies. 
* `nir`:(string) NIR application usage. This determines if the fabric has NIR application installed. 
* `open_stack`:(string) Open stack feature usage. 
* `opflex_kubernetes_count`:(int) Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes. 
* `ospf_count`:(int) Ospf feature usage. This determines the total number of OSPF sessions across the fabric. 
* `password_history_count`:(int) Returns count of passwords. 
* `password_strength_check`:(string) Returns if the password is strong or not. 
* `password_strength_profile_count`:(int) Returns the number of password strength profile. 
* `poe_count`:(int) POE feature usage. This determines the total number of POE configurations across the fabric. 
* `port_security_count`:(int) Number of objects with Port Security enabled. Non-Zero value indicates the object as enabled. 
* `qin_vni_tunnel_count`:(int) QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it. 
* `qos_cong_count`:(int) Number of Quality Of Service congestion class. 
* `qos_pfc_pol_count`:(int) Number of Quality Of Service class. 
* `realm_count`:(int) Returns the value of count of realms. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `remote_leaf_count`:(int) Number of remote Leafs. This determines the total number of remote leaf switches in the fabric. 
* `scvmm_count`:(int) SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric. 
* `shared_l3_out_count`:(int) SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. There will be a feature usage object per site in Multi site scenario. In multi-site scenario the site name is available in all the requests being made. 
* `smart_call_home`:(string) Smart callhome feature usage. This determines if this feature is being enabled or disabled. 
* `snapshot_count`:(int) Returns count of snapshots. 
* `snmp`:(string) SNMP feature usage. This determines if this feature is enabled or disabled. 
* `snmp_community_access_count`:(int) Returns count of SNMP Community Access. 
* `snmp_group_count`:(int) Number of SNMP monitoring policies on the fabric. 
* `snmp_trap_count`:(int) Returns count of SNMP trap. 
* `snmp_v3_count`:(int) Returns count of SNMP V3 on the fabric. 
* `span_count`:(int) Number of Span Sources and Destinations. 
* `span_dst_count`:(int) Number of Span Destinations with valid state. 
* `span_src_count`:(int) Number of Span Sources with valid state. 
* `spine_count`:(int) Number of Spines. This determines the total number of spine switches in the fabric. 
* `ssh_over_v6_count`:(int) Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `ssh_v2_count`:(int) Returns count of ssh V2 on the fabric. 
* `supervisor_module_count`:(int) Returns the total number of supervisor module slots. 
* `syslog_group_count`:(int) Number of syslog monitoring policies on the fabric. 
* `syslog_over_v6_count`:(int) Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `system_controller_count`:(int) Returns the total number of system controller slots. 
* `tacacs_group_count`:(int) Number of tacacs monitoring policies on the fabric. 
* `tenant_count`:(int) Number of tenants. This determines the total number of tenants configured across the fabric. 
* `tier_two_leaf_count`:(int) Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric. 
* `total_critical_faults`:(int) Returns the total number of critical faults. 
* `twamp`:(string) TWAMP feature usage. This determines if this feature is enabled or disabled. 
* `useg`:(string) VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled. 
* `vm_ware_vds_count`:(int) Number of objects with VmWare vCenter 6.5 support. Checks the controller revision value. 
* `vmm_ctrlrp_count`:(int) Number of Virtual Machine Monitor controller policy objects for VMware vCenter. 
* `vmm_domp_count`:(int) Number of Virtual Machine Monitor domain policy model objects for VMware vCenter. 
* `vmm_ep_pd_count`:(int) Microsegmentation Distributed Virtual Switch feature usage. Gets the number of objects associated to VMware vCenter. 
* `vnsm_dev_count`:(int) Number of objects with L4-L7 Device Package Import enabled. Checks for the vendor and the model. 
* `vpod_count`:(int) Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics. 
* `webtoken_timeout_seconds`:(int) Timeout for web token in seconds. 
 
