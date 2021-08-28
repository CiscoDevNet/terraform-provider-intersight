package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryNiaFeatureUsage() *schema.Resource {
	var subSchema = map[string]*schema.Schema{"aaa_ldap_provider_count": {
		Description: "Returns the total number of AAA Ldap Providers.",
		Type:        schema.TypeInt,
		Optional:    true,
	},
		"aaa_radius_provider_count": {
			Description: "Returns the total number of AAA Radius Providers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"aaa_tacacs_provider_count": {
			Description: "Returns the total number of AAA Tacacs Providers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"account_moid": {
			Description: "The Account ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"apic_count": {
			Description: "Number of APIC controllers. This determines the value of controllers for the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_is_telnet_enabled": {
			Description: "Returns if telnet is enabled on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"apic_ntp_count": {
			Description: "Count of NTP servers configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_snmp_community_count": {
			Description: "Number of SNMP communities configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_sys_log_grp_count": {
			Description: "Number of logging groups configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_sys_log_src_count": {
			Description: "Number of logging sources configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"app_center_count": {
			Description: "ACI APPs feature usage scale.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ave": {
			Description: "AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"bd_count": {
			Description: "Number of BDs. This determines the total number of Broadcast Domains across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"callhome_smart_group_count": {
			Description: "Number of call home smart monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cloud_sec_peer_count": {
			Description: "Number of Cloudsec SA peers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"comp_hv_count": {
			Description: "Number of compute hypervisors on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"config_exportp_count": {
			Description: "Number of system backup configure export policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"config_job_count": {
			Description: "Number of system backup configure jobs on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"consistency_checker_app": {
			Description: "Consistency checker application usage. This determines if the fabric has Consistency checker application installed.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"contract_count": {
			Description: "Number of contracts. This determines the total number of Contracts configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dns_count": {
			Description: "DNS feature usage. This determines the total number of DNS configurations across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"eigrp_count": {
			Description: "Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"epg_count": {
			Description: "Number of End Point Groups. This determines the total number of End Point Groups across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_module_count": {
			Description: "Returns the total number of fabric module slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_setupp_count": {
			Description: "Number of Multi-Pods per fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_count": {
			Description: "Total number of FCoE N-Port for DOM, VSAn, and VLAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_dom_count": {
			Description: "Number of FCoE N-Port DOM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_vlan_count": {
			Description: "Number of FCoE N-Port VLAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_vsan_count": {
			Description: "Number of FCoE N-Port VSAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fv_sla_def_count": {
			Description: "Number of Internet Protocol Service Level Agreements Monitoring policy objects for object tracking.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"hsrp_count": {
			Description: "Hsrp feature usage. This determines the total number of HSRP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ibgp_count": {
			Description: "Ibgp feature usage. This determines the total number of BGP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"igmp_access_list_count": {
			Description: "IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"igmp_snoop": {
			Description: "IGMP Snooping feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ip_epg_count": {
			Description: "Number of IP based End Point Groups. This determines the total number of IP End Point Groups across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"is_bgp_route_reflectors_feature_used": {
			Description: "BGP route reflector usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_bridge_domains_feature_used": {
			Description: "Brodge domains feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_common_local_user_name": {
			Description: "Returns value of isCommonLocalUserName field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_contracts_feature_used": {
			Description: "Contracts feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_epg_feature_used": {
			Description: "EPG feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_filters_feature_used": {
			Description: "Filters feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_http_configured": {
			Description: "Returns if HTTP is configured.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_https_configured": {
			Description: "Returns if HTTPS is configured.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ntp_feature_used": {
			Description: "NTP feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ptp_feature_used": {
			Description: "Ptp feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_synce_feature_used": {
			Description: "Synce feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_tech_support_collected": {
			Description: "Status of techsupport collection.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_tenants_feature_used": {
			Description: "Tenants feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_vrfs_feature_used": {
			Description: "VRF feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"isis_count": {
			Description: "Isis feature usage. This determines the total number of ISIS sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"l2_multicast": {
			Description: "L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"leaf_count": {
			Description: "Number of Leafs. This determines the total number of Leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"local_username_count": {
			Description: "Returns count of local users.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_block_duration": {
			Description: "Returns login block duration value.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_max_failed_attempts": {
			Description: "Returns the maximum failed attempts on login.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_max_failed_attempts_window": {
			Description: "Returns the maximum failed attempt windows on login.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"maintenance_mode_count": {
			Description: "Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"management_over_v6_count": {
			Description: "Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"microsoft_useg_vmm_ep_pd_count": {
			Description: "Number of Microsoft microsegmentation VmmEpPD objects. Ensures that Microsoft was configured.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"net_flow_count": {
			Description: "Number of Netflow monitor policies.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nir": {
			Description: "NIR application usage. This determines if the fabric has NIR application installed.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"open_stack": {
			Description: "Open stack feature usage.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"opflex_kubernetes_count": {
			Description: "Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ospf_count": {
			Description: "Ospf feature usage. This determines the total number of OSPF sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"password_history_count": {
			Description: "Returns count of passwords.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"password_strength_check": {
			Description: "Returns if the password is strong or not.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"password_strength_profile_count": {
			Description: "Returns the number of password strength profile.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"poe_count": {
			Description: "POE feature usage. This determines the total number of POE configurations across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"port_security_count": {
			Description: "Number of objects with Port Security enabled. Non-Zero value indicates the object as enabled.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qin_vni_tunnel_count": {
			Description: "QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qos_cong_count": {
			Description: "Number of Quality Of Service congestion class.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qos_pfc_pol_count": {
			Description: "Number of Quality Of Service class.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"realm_count": {
			Description: "Returns the value of count of realms.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"record_type": {
			Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"record_version": {
			Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"registered_device": {
			Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"remote_leaf_count": {
			Description: "Number of remote Leafs. This determines the total number of remote leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"scvmm_count": {
			Description: "SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"shared_l3_out_count": {
			Description: "SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"site_name": {
			Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. There will be a feature usage object per site in Multi site scenario. In multi-site scenario the site name is available in all the requests being made.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"smart_call_home": {
			Description: "Smart callhome feature usage. This determines if this feature is being enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"smart_license": {
			Description: "Details of smart license.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"active_mode": {
						Description: "Indicate the mode smart license is curerntly running.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"auth_status": {
						Description: "Authorization status of the smart license.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"license_udi": {
						Description: "License Udi of the smart license.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"smart_account": {
						Description: "Smart licensing account name in CSSM and is retrieved from CSSM after regsitration.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"snapshot_count": {
			Description: "Returns count of snapshots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp": {
			Description: "SNMP feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snmp_community_access_count": {
			Description: "Returns count of SNMP Community Access.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_group_count": {
			Description: "Number of SNMP monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_trap_count": {
			Description: "Returns count of SNMP trap.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_v3_count": {
			Description: "Returns count of SNMP V3 on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_count": {
			Description: "Number of Span Sources and Destinations.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_dst_count": {
			Description: "Number of Span Destinations with valid state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_src_count": {
			Description: "Number of Span Sources with valid state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"spine_count": {
			Description: "Number of Spines. This determines the total number of spine switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ssh_over_v6_count": {
			Description: "Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ssh_v2_count": {
			Description: "Returns count of ssh V2 on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"supervisor_module_count": {
			Description: "Returns the total number of supervisor module slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_group_count": {
			Description: "Number of syslog monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_over_v6_count": {
			Description: "Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"system_controller_count": {
			Description: "Returns the total number of system controller slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tacacs_group_count": {
			Description: "Number of tacacs monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tenant_count": {
			Description: "Number of tenants. This determines the total number of tenants configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tier_two_leaf_count": {
			Description: "Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"total_critical_faults": {
			Description: "Returns the total number of critical faults.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"twamp": {
			Description: "TWAMP feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"useg": {
			Description: "VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_ware_vds_count": {
			Description: "Number of objects with VmWare vCenter 6.5 support. Checks the controller revision value.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_ctrlrp_count": {
			Description: "Number of Virtual Machine Monitor controller policy objects for VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_domp_count": {
			Description: "Number of Virtual Machine Monitor domain policy model objects for VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_ep_pd_count": {
			Description: "Microsegmentation Distributed Virtual Switch feature usage. Gets the number of objects associated to VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vnsm_dev_count": {
			Description: "Number of objects with L4-L7 Device Package Import enabled. Checks for the vendor and the model.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpod_count": {
			Description: "Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"webtoken_timeout_seconds": {
			Description: "Timeout for web token in seconds.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
	var model = map[string]*schema.Schema{"aaa_ldap_provider_count": {
		Description: "Returns the total number of AAA Ldap Providers.",
		Type:        schema.TypeInt,
		Optional:    true,
	},
		"aaa_radius_provider_count": {
			Description: "Returns the total number of AAA Radius Providers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"aaa_tacacs_provider_count": {
			Description: "Returns the total number of AAA Tacacs Providers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"account_moid": {
			Description: "The Account ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"apic_count": {
			Description: "Number of APIC controllers. This determines the value of controllers for the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_is_telnet_enabled": {
			Description: "Returns if telnet is enabled on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"apic_ntp_count": {
			Description: "Count of NTP servers configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_snmp_community_count": {
			Description: "Number of SNMP communities configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_sys_log_grp_count": {
			Description: "Number of logging groups configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"apic_sys_log_src_count": {
			Description: "Number of logging sources configured on APIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"app_center_count": {
			Description: "ACI APPs feature usage scale.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ave": {
			Description: "AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"bd_count": {
			Description: "Number of BDs. This determines the total number of Broadcast Domains across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"callhome_smart_group_count": {
			Description: "Number of call home smart monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cloud_sec_peer_count": {
			Description: "Number of Cloudsec SA peers.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"comp_hv_count": {
			Description: "Number of compute hypervisors on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"config_exportp_count": {
			Description: "Number of system backup configure export policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"config_job_count": {
			Description: "Number of system backup configure jobs on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"consistency_checker_app": {
			Description: "Consistency checker application usage. This determines if the fabric has Consistency checker application installed.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"contract_count": {
			Description: "Number of contracts. This determines the total number of Contracts configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dns_count": {
			Description: "DNS feature usage. This determines the total number of DNS configurations across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"eigrp_count": {
			Description: "Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"epg_count": {
			Description: "Number of End Point Groups. This determines the total number of End Point Groups across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_module_count": {
			Description: "Returns the total number of fabric module slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_setupp_count": {
			Description: "Number of Multi-Pods per fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_count": {
			Description: "Total number of FCoE N-Port for DOM, VSAn, and VLAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_dom_count": {
			Description: "Number of FCoE N-Port DOM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_vlan_count": {
			Description: "Number of FCoE N-Port VLAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fcoe_nport_vsan_count": {
			Description: "Number of FCoE N-Port VSAN.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fv_sla_def_count": {
			Description: "Number of Internet Protocol Service Level Agreements Monitoring policy objects for object tracking.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"hsrp_count": {
			Description: "Hsrp feature usage. This determines the total number of HSRP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ibgp_count": {
			Description: "Ibgp feature usage. This determines the total number of BGP sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"igmp_access_list_count": {
			Description: "IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"igmp_snoop": {
			Description: "IGMP Snooping feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ip_epg_count": {
			Description: "Number of IP based End Point Groups. This determines the total number of IP End Point Groups across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"is_bgp_route_reflectors_feature_used": {
			Description: "BGP route reflector usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_bridge_domains_feature_used": {
			Description: "Brodge domains feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_common_local_user_name": {
			Description: "Returns value of isCommonLocalUserName field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_contracts_feature_used": {
			Description: "Contracts feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_epg_feature_used": {
			Description: "EPG feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_filters_feature_used": {
			Description: "Filters feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_http_configured": {
			Description: "Returns if HTTP is configured.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_https_configured": {
			Description: "Returns if HTTPS is configured.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ntp_feature_used": {
			Description: "NTP feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ptp_feature_used": {
			Description: "Ptp feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_synce_feature_used": {
			Description: "Synce feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_tech_support_collected": {
			Description: "Status of techsupport collection.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_tenants_feature_used": {
			Description: "Tenants feature usage on APIC.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_vrfs_feature_used": {
			Description: "VRF feature usage on APIC controller.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"isis_count": {
			Description: "Isis feature usage. This determines the total number of ISIS sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"l2_multicast": {
			Description: "L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"leaf_count": {
			Description: "Number of Leafs. This determines the total number of Leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"local_username_count": {
			Description: "Returns count of local users.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_block_duration": {
			Description: "Returns login block duration value.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_max_failed_attempts": {
			Description: "Returns the maximum failed attempts on login.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"login_max_failed_attempts_window": {
			Description: "Returns the maximum failed attempt windows on login.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"maintenance_mode_count": {
			Description: "Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"management_over_v6_count": {
			Description: "Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"microsoft_useg_vmm_ep_pd_count": {
			Description: "Number of Microsoft microsegmentation VmmEpPD objects. Ensures that Microsoft was configured.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"net_flow_count": {
			Description: "Number of Netflow monitor policies.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nir": {
			Description: "NIR application usage. This determines if the fabric has NIR application installed.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"open_stack": {
			Description: "Open stack feature usage.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"opflex_kubernetes_count": {
			Description: "Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ospf_count": {
			Description: "Ospf feature usage. This determines the total number of OSPF sessions across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"password_history_count": {
			Description: "Returns count of passwords.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"password_strength_check": {
			Description: "Returns if the password is strong or not.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"password_strength_profile_count": {
			Description: "Returns the number of password strength profile.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"poe_count": {
			Description: "POE feature usage. This determines the total number of POE configurations across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"port_security_count": {
			Description: "Number of objects with Port Security enabled. Non-Zero value indicates the object as enabled.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qin_vni_tunnel_count": {
			Description: "QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qos_cong_count": {
			Description: "Number of Quality Of Service congestion class.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"qos_pfc_pol_count": {
			Description: "Number of Quality Of Service class.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"realm_count": {
			Description: "Returns the value of count of realms.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"record_type": {
			Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"record_version": {
			Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"registered_device": {
			Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"remote_leaf_count": {
			Description: "Number of remote Leafs. This determines the total number of remote leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"scvmm_count": {
			Description: "SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"shared_l3_out_count": {
			Description: "SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"site_name": {
			Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. There will be a feature usage object per site in Multi site scenario. In multi-site scenario the site name is available in all the requests being made.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"smart_call_home": {
			Description: "Smart callhome feature usage. This determines if this feature is being enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"smart_license": {
			Description: "Details of smart license.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"active_mode": {
						Description: "Indicate the mode smart license is curerntly running.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"auth_status": {
						Description: "Authorization status of the smart license.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"license_udi": {
						Description: "License Udi of the smart license.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"smart_account": {
						Description: "Smart licensing account name in CSSM and is retrieved from CSSM after regsitration.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"snapshot_count": {
			Description: "Returns count of snapshots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp": {
			Description: "SNMP feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snmp_community_access_count": {
			Description: "Returns count of SNMP Community Access.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_group_count": {
			Description: "Number of SNMP monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_trap_count": {
			Description: "Returns count of SNMP trap.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"snmp_v3_count": {
			Description: "Returns count of SNMP V3 on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_count": {
			Description: "Number of Span Sources and Destinations.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_dst_count": {
			Description: "Number of Span Destinations with valid state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"span_src_count": {
			Description: "Number of Span Sources with valid state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"spine_count": {
			Description: "Number of Spines. This determines the total number of spine switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ssh_over_v6_count": {
			Description: "Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ssh_v2_count": {
			Description: "Returns count of ssh V2 on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"supervisor_module_count": {
			Description: "Returns the total number of supervisor module slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_group_count": {
			Description: "Number of syslog monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_over_v6_count": {
			Description: "Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"system_controller_count": {
			Description: "Returns the total number of system controller slots.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tacacs_group_count": {
			Description: "Number of tacacs monitoring policies on the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tenant_count": {
			Description: "Number of tenants. This determines the total number of tenants configured across the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tier_two_leaf_count": {
			Description: "Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"total_critical_faults": {
			Description: "Returns the total number of critical faults.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"twamp": {
			Description: "TWAMP feature usage. This determines if this feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"useg": {
			Description: "VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_ware_vds_count": {
			Description: "Number of objects with VmWare vCenter 6.5 support. Checks the controller revision value.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_ctrlrp_count": {
			Description: "Number of Virtual Machine Monitor controller policy objects for VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_domp_count": {
			Description: "Number of Virtual Machine Monitor domain policy model objects for VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmm_ep_pd_count": {
			Description: "Microsegmentation Distributed Virtual Switch feature usage. Gets the number of objects associated to VMware vCenter.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vnsm_dev_count": {
			Description: "Number of objects with L4-L7 Device Package Import enabled. Checks for the vendor and the model.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpod_count": {
			Description: "Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"webtoken_timeout_seconds": {
			Description: "Timeout for web token in seconds.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaFeatureUsageRead,
		Schema:      model}
}

func dataSourceNiatelemetryNiaFeatureUsageRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaFeatureUsage{}
	if _, ok := d.GetOk("aaa_ldap_provider_count"); ok {
		v := d.Get("aaa_ldap_provider_count")
		x := int64(v.(int))
		o.SetAaaLdapProviderCount(x)
	}

	if _, ok := d.GetOk("aaa_radius_provider_count"); ok {
		v := d.Get("aaa_radius_provider_count")
		x := int64(v.(int))
		o.SetAaaRadiusProviderCount(x)
	}

	if _, ok := d.GetOk("aaa_tacacs_provider_count"); ok {
		v := d.Get("aaa_tacacs_provider_count")
		x := int64(v.(int))
		o.SetAaaTacacsProviderCount(x)
	}

	if _, ok := d.GetOk("account_moid"); ok {
		v := d.Get("account_moid")
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if _, ok := d.GetOk("additional_properties"); ok {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if _, ok := d.GetOk("ancestors"); ok {
		v := d.Get("ancestors")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if _, ok := d.GetOk("apic_count"); ok {
		v := d.Get("apic_count")
		x := int64(v.(int))
		o.SetApicCount(x)
	}

	o.SetApicIsTelnetEnabled(d.Get("apic_is_telnet_enabled").(bool))

	if _, ok := d.GetOk("apic_ntp_count"); ok {
		v := d.Get("apic_ntp_count")
		x := int64(v.(int))
		o.SetApicNtpCount(x)
	}

	if _, ok := d.GetOk("apic_snmp_community_count"); ok {
		v := d.Get("apic_snmp_community_count")
		x := int64(v.(int))
		o.SetApicSnmpCommunityCount(x)
	}

	if _, ok := d.GetOk("apic_sys_log_grp_count"); ok {
		v := d.Get("apic_sys_log_grp_count")
		x := int64(v.(int))
		o.SetApicSysLogGrpCount(x)
	}

	if _, ok := d.GetOk("apic_sys_log_src_count"); ok {
		v := d.Get("apic_sys_log_src_count")
		x := int64(v.(int))
		o.SetApicSysLogSrcCount(x)
	}

	if _, ok := d.GetOk("app_center_count"); ok {
		v := d.Get("app_center_count")
		x := int64(v.(int))
		o.SetAppCenterCount(x)
	}

	if _, ok := d.GetOk("ave"); ok {
		v := d.Get("ave")
		x := (v.(string))
		o.SetAve(x)
	}

	if _, ok := d.GetOk("bd_count"); ok {
		v := d.Get("bd_count")
		x := int64(v.(int))
		o.SetBdCount(x)
	}

	if _, ok := d.GetOk("callhome_smart_group_count"); ok {
		v := d.Get("callhome_smart_group_count")
		x := int64(v.(int))
		o.SetCallhomeSmartGroupCount(x)
	}

	if _, ok := d.GetOk("class_id"); ok {
		v := d.Get("class_id")
		x := (v.(string))
		o.SetClassId(x)
	}

	if _, ok := d.GetOk("cloud_sec_peer_count"); ok {
		v := d.Get("cloud_sec_peer_count")
		x := int64(v.(int))
		o.SetCloudSecPeerCount(x)
	}

	if _, ok := d.GetOk("comp_hv_count"); ok {
		v := d.Get("comp_hv_count")
		x := int64(v.(int))
		o.SetCompHvCount(x)
	}

	if _, ok := d.GetOk("config_exportp_count"); ok {
		v := d.Get("config_exportp_count")
		x := int64(v.(int))
		o.SetConfigExportpCount(x)
	}

	if _, ok := d.GetOk("config_job_count"); ok {
		v := d.Get("config_job_count")
		x := int64(v.(int))
		o.SetConfigJobCount(x)
	}

	if _, ok := d.GetOk("consistency_checker_app"); ok {
		v := d.Get("consistency_checker_app")
		x := (v.(string))
		o.SetConsistencyCheckerApp(x)
	}

	if _, ok := d.GetOk("contract_count"); ok {
		v := d.Get("contract_count")
		x := int64(v.(int))
		o.SetContractCount(x)
	}

	if _, ok := d.GetOk("create_time"); ok {
		v := d.Get("create_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}

	if _, ok := d.GetOk("dns_count"); ok {
		v := d.Get("dns_count")
		x := int64(v.(int))
		o.SetDnsCount(x)
	}

	if _, ok := d.GetOk("domain_group_moid"); ok {
		v := d.Get("domain_group_moid")
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if _, ok := d.GetOk("eigrp_count"); ok {
		v := d.Get("eigrp_count")
		x := int64(v.(int))
		o.SetEigrpCount(x)
	}

	if _, ok := d.GetOk("epg_count"); ok {
		v := d.Get("epg_count")
		x := int64(v.(int))
		o.SetEpgCount(x)
	}

	if _, ok := d.GetOk("fabric_module_count"); ok {
		v := d.Get("fabric_module_count")
		x := int64(v.(int))
		o.SetFabricModuleCount(x)
	}

	if _, ok := d.GetOk("fabric_setupp_count"); ok {
		v := d.Get("fabric_setupp_count")
		x := int64(v.(int))
		o.SetFabricSetuppCount(x)
	}

	if _, ok := d.GetOk("fcoe_nport_count"); ok {
		v := d.Get("fcoe_nport_count")
		x := int64(v.(int))
		o.SetFcoeNportCount(x)
	}

	if _, ok := d.GetOk("fcoe_nport_dom_count"); ok {
		v := d.Get("fcoe_nport_dom_count")
		x := int64(v.(int))
		o.SetFcoeNportDomCount(x)
	}

	if _, ok := d.GetOk("fcoe_nport_vlan_count"); ok {
		v := d.Get("fcoe_nport_vlan_count")
		x := int64(v.(int))
		o.SetFcoeNportVlanCount(x)
	}

	if _, ok := d.GetOk("fcoe_nport_vsan_count"); ok {
		v := d.Get("fcoe_nport_vsan_count")
		x := int64(v.(int))
		o.SetFcoeNportVsanCount(x)
	}

	if _, ok := d.GetOk("fv_sla_def_count"); ok {
		v := d.Get("fv_sla_def_count")
		x := int64(v.(int))
		o.SetFvSlaDefCount(x)
	}

	if _, ok := d.GetOk("hsrp_count"); ok {
		v := d.Get("hsrp_count")
		x := int64(v.(int))
		o.SetHsrpCount(x)
	}

	if _, ok := d.GetOk("ibgp_count"); ok {
		v := d.Get("ibgp_count")
		x := int64(v.(int))
		o.SetIbgpCount(x)
	}

	if _, ok := d.GetOk("igmp_access_list_count"); ok {
		v := d.Get("igmp_access_list_count")
		x := int64(v.(int))
		o.SetIgmpAccessListCount(x)
	}

	if _, ok := d.GetOk("igmp_snoop"); ok {
		v := d.Get("igmp_snoop")
		x := (v.(string))
		o.SetIgmpSnoop(x)
	}

	if _, ok := d.GetOk("ip_epg_count"); ok {
		v := d.Get("ip_epg_count")
		x := int64(v.(int))
		o.SetIpEpgCount(x)
	}

	o.SetIsBgpRouteReflectorsFeatureUsed(d.Get("is_bgp_route_reflectors_feature_used").(bool))

	o.SetIsBridgeDomainsFeatureUsed(d.Get("is_bridge_domains_feature_used").(bool))

	o.SetIsCommonLocalUserName(d.Get("is_common_local_user_name").(bool))

	o.SetIsContractsFeatureUsed(d.Get("is_contracts_feature_used").(bool))

	o.SetIsEpgFeatureUsed(d.Get("is_epg_feature_used").(bool))

	o.SetIsFiltersFeatureUsed(d.Get("is_filters_feature_used").(bool))

	o.SetIsHttpConfigured(d.Get("is_http_configured").(bool))

	o.SetIsHttpsConfigured(d.Get("is_https_configured").(bool))

	o.SetIsNtpFeatureUsed(d.Get("is_ntp_feature_used").(bool))

	o.SetIsPtpFeatureUsed(d.Get("is_ptp_feature_used").(bool))

	o.SetIsSynceFeatureUsed(d.Get("is_synce_feature_used").(bool))

	if _, ok := d.GetOk("is_tech_support_collected"); ok {
		v := d.Get("is_tech_support_collected")
		x := (v.(string))
		o.SetIsTechSupportCollected(x)
	}

	o.SetIsTenantsFeatureUsed(d.Get("is_tenants_feature_used").(bool))

	o.SetIsVrfsFeatureUsed(d.Get("is_vrfs_feature_used").(bool))

	if _, ok := d.GetOk("isis_count"); ok {
		v := d.Get("isis_count")
		x := int64(v.(int))
		o.SetIsisCount(x)
	}

	if _, ok := d.GetOk("l2_multicast"); ok {
		v := d.Get("l2_multicast")
		x := (v.(string))
		o.SetL2Multicast(x)
	}

	if _, ok := d.GetOk("leaf_count"); ok {
		v := d.Get("leaf_count")
		x := int64(v.(int))
		o.SetLeafCount(x)
	}

	if _, ok := d.GetOk("local_username_count"); ok {
		v := d.Get("local_username_count")
		x := int64(v.(int))
		o.SetLocalUsernameCount(x)
	}

	if _, ok := d.GetOk("login_block_duration"); ok {
		v := d.Get("login_block_duration")
		x := int64(v.(int))
		o.SetLoginBlockDuration(x)
	}

	if _, ok := d.GetOk("login_max_failed_attempts"); ok {
		v := d.Get("login_max_failed_attempts")
		x := int64(v.(int))
		o.SetLoginMaxFailedAttempts(x)
	}

	if _, ok := d.GetOk("login_max_failed_attempts_window"); ok {
		v := d.Get("login_max_failed_attempts_window")
		x := int64(v.(int))
		o.SetLoginMaxFailedAttemptsWindow(x)
	}

	if _, ok := d.GetOk("maintenance_mode_count"); ok {
		v := d.Get("maintenance_mode_count")
		x := int64(v.(int))
		o.SetMaintenanceModeCount(x)
	}

	if _, ok := d.GetOk("management_over_v6_count"); ok {
		v := d.Get("management_over_v6_count")
		x := int64(v.(int))
		o.SetManagementOverV6Count(x)
	}

	if _, ok := d.GetOk("microsoft_useg_vmm_ep_pd_count"); ok {
		v := d.Get("microsoft_useg_vmm_ep_pd_count")
		x := int64(v.(int))
		o.SetMicrosoftUsegVmmEpPdCount(x)
	}

	if _, ok := d.GetOk("mod_time"); ok {
		v := d.Get("mod_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}

	if _, ok := d.GetOk("moid"); ok {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if _, ok := d.GetOk("net_flow_count"); ok {
		v := d.Get("net_flow_count")
		x := int64(v.(int))
		o.SetNetFlowCount(x)
	}

	if _, ok := d.GetOk("nir"); ok {
		v := d.Get("nir")
		x := (v.(string))
		o.SetNir(x)
	}

	if _, ok := d.GetOk("object_type"); ok {
		v := d.Get("object_type")
		x := (v.(string))
		o.SetObjectType(x)
	}

	if _, ok := d.GetOk("open_stack"); ok {
		v := d.Get("open_stack")
		x := (v.(string))
		o.SetOpenStack(x)
	}

	if _, ok := d.GetOk("opflex_kubernetes_count"); ok {
		v := d.Get("opflex_kubernetes_count")
		x := int64(v.(int))
		o.SetOpflexKubernetesCount(x)
	}

	if _, ok := d.GetOk("ospf_count"); ok {
		v := d.Get("ospf_count")
		x := int64(v.(int))
		o.SetOspfCount(x)
	}

	if _, ok := d.GetOk("owners"); ok {
		v := d.Get("owners")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.SetOwners(x)
	}

	if _, ok := d.GetOk("parent"); ok {
		v := d.Get("parent")
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if _, ok := d.GetOk("password_history_count"); ok {
		v := d.Get("password_history_count")
		x := int64(v.(int))
		o.SetPasswordHistoryCount(x)
	}

	if _, ok := d.GetOk("password_strength_check"); ok {
		v := d.Get("password_strength_check")
		x := (v.(string))
		o.SetPasswordStrengthCheck(x)
	}

	if _, ok := d.GetOk("password_strength_profile_count"); ok {
		v := d.Get("password_strength_profile_count")
		x := int64(v.(int))
		o.SetPasswordStrengthProfileCount(x)
	}

	if _, ok := d.GetOk("permission_resources"); ok {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if _, ok := d.GetOk("poe_count"); ok {
		v := d.Get("poe_count")
		x := int64(v.(int))
		o.SetPoeCount(x)
	}

	if _, ok := d.GetOk("port_security_count"); ok {
		v := d.Get("port_security_count")
		x := int64(v.(int))
		o.SetPortSecurityCount(x)
	}

	if _, ok := d.GetOk("qin_vni_tunnel_count"); ok {
		v := d.Get("qin_vni_tunnel_count")
		x := int64(v.(int))
		o.SetQinVniTunnelCount(x)
	}

	if _, ok := d.GetOk("qos_cong_count"); ok {
		v := d.Get("qos_cong_count")
		x := int64(v.(int))
		o.SetQosCongCount(x)
	}

	if _, ok := d.GetOk("qos_pfc_pol_count"); ok {
		v := d.Get("qos_pfc_pol_count")
		x := int64(v.(int))
		o.SetQosPfcPolCount(x)
	}

	if _, ok := d.GetOk("realm_count"); ok {
		v := d.Get("realm_count")
		x := int64(v.(int))
		o.SetRealmCount(x)
	}

	if _, ok := d.GetOk("record_type"); ok {
		v := d.Get("record_type")
		x := (v.(string))
		o.SetRecordType(x)
	}

	if _, ok := d.GetOk("record_version"); ok {
		v := d.Get("record_version")
		x := (v.(string))
		o.SetRecordVersion(x)
	}

	if _, ok := d.GetOk("registered_device"); ok {
		v := d.Get("registered_device")
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if _, ok := d.GetOk("remote_leaf_count"); ok {
		v := d.Get("remote_leaf_count")
		x := int64(v.(int))
		o.SetRemoteLeafCount(x)
	}

	if _, ok := d.GetOk("scvmm_count"); ok {
		v := d.Get("scvmm_count")
		x := int64(v.(int))
		o.SetScvmmCount(x)
	}

	if _, ok := d.GetOk("shared_l3_out_count"); ok {
		v := d.Get("shared_l3_out_count")
		x := int64(v.(int))
		o.SetSharedL3OutCount(x)
	}

	if _, ok := d.GetOk("shared_scope"); ok {
		v := d.Get("shared_scope")
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if _, ok := d.GetOk("site_name"); ok {
		v := d.Get("site_name")
		x := (v.(string))
		o.SetSiteName(x)
	}

	if _, ok := d.GetOk("smart_call_home"); ok {
		v := d.Get("smart_call_home")
		x := (v.(string))
		o.SetSmartCallHome(x)
	}

	if _, ok := d.GetOk("smart_license"); ok {
		v := d.Get("smart_license")
		p := make([]models.NiatelemetrySmartLicense, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetrySmartLicense{}
			if v, ok := l["active_mode"]; ok {
				{
					x := (v.(string))
					o.SetActiveMode(x)
				}
			}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["auth_status"]; ok {
				{
					x := (v.(string))
					o.SetAuthStatus(x)
				}
			}
			o.SetClassId("niatelemetry.SmartLicense")
			if v, ok := l["license_udi"]; ok {
				{
					x := (v.(string))
					o.SetLicenseUdi(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["smart_account"]; ok {
				{
					x := (v.(string))
					o.SetSmartAccount(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSmartLicense(x)
		}
	}

	if _, ok := d.GetOk("snapshot_count"); ok {
		v := d.Get("snapshot_count")
		x := int64(v.(int))
		o.SetSnapshotCount(x)
	}

	if _, ok := d.GetOk("snmp"); ok {
		v := d.Get("snmp")
		x := (v.(string))
		o.SetSnmp(x)
	}

	if _, ok := d.GetOk("snmp_community_access_count"); ok {
		v := d.Get("snmp_community_access_count")
		x := int64(v.(int))
		o.SetSnmpCommunityAccessCount(x)
	}

	if _, ok := d.GetOk("snmp_group_count"); ok {
		v := d.Get("snmp_group_count")
		x := int64(v.(int))
		o.SetSnmpGroupCount(x)
	}

	if _, ok := d.GetOk("snmp_trap_count"); ok {
		v := d.Get("snmp_trap_count")
		x := int64(v.(int))
		o.SetSnmpTrapCount(x)
	}

	if _, ok := d.GetOk("snmp_v3_count"); ok {
		v := d.Get("snmp_v3_count")
		x := int64(v.(int))
		o.SetSnmpV3Count(x)
	}

	if _, ok := d.GetOk("span_count"); ok {
		v := d.Get("span_count")
		x := int64(v.(int))
		o.SetSpanCount(x)
	}

	if _, ok := d.GetOk("span_dst_count"); ok {
		v := d.Get("span_dst_count")
		x := int64(v.(int))
		o.SetSpanDstCount(x)
	}

	if _, ok := d.GetOk("span_src_count"); ok {
		v := d.Get("span_src_count")
		x := int64(v.(int))
		o.SetSpanSrcCount(x)
	}

	if _, ok := d.GetOk("spine_count"); ok {
		v := d.Get("spine_count")
		x := int64(v.(int))
		o.SetSpineCount(x)
	}

	if _, ok := d.GetOk("ssh_over_v6_count"); ok {
		v := d.Get("ssh_over_v6_count")
		x := int64(v.(int))
		o.SetSshOverV6Count(x)
	}

	if _, ok := d.GetOk("ssh_v2_count"); ok {
		v := d.Get("ssh_v2_count")
		x := int64(v.(int))
		o.SetSshV2Count(x)
	}

	if _, ok := d.GetOk("supervisor_module_count"); ok {
		v := d.Get("supervisor_module_count")
		x := int64(v.(int))
		o.SetSupervisorModuleCount(x)
	}

	if _, ok := d.GetOk("syslog_group_count"); ok {
		v := d.Get("syslog_group_count")
		x := int64(v.(int))
		o.SetSyslogGroupCount(x)
	}

	if _, ok := d.GetOk("syslog_over_v6_count"); ok {
		v := d.Get("syslog_over_v6_count")
		x := int64(v.(int))
		o.SetSyslogOverV6Count(x)
	}

	if _, ok := d.GetOk("system_controller_count"); ok {
		v := d.Get("system_controller_count")
		x := int64(v.(int))
		o.SetSystemControllerCount(x)
	}

	if _, ok := d.GetOk("tacacs_group_count"); ok {
		v := d.Get("tacacs_group_count")
		x := int64(v.(int))
		o.SetTacacsGroupCount(x)
	}

	if _, ok := d.GetOk("tags"); ok {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if _, ok := d.GetOk("tenant_count"); ok {
		v := d.Get("tenant_count")
		x := int64(v.(int))
		o.SetTenantCount(x)
	}

	if _, ok := d.GetOk("tier_two_leaf_count"); ok {
		v := d.Get("tier_two_leaf_count")
		x := int64(v.(int))
		o.SetTierTwoLeafCount(x)
	}

	if _, ok := d.GetOk("total_critical_faults"); ok {
		v := d.Get("total_critical_faults")
		x := int64(v.(int))
		o.SetTotalCriticalFaults(x)
	}

	if _, ok := d.GetOk("twamp"); ok {
		v := d.Get("twamp")
		x := (v.(string))
		o.SetTwamp(x)
	}

	if _, ok := d.GetOk("useg"); ok {
		v := d.Get("useg")
		x := (v.(string))
		o.SetUseg(x)
	}

	if _, ok := d.GetOk("version_context"); ok {
		v := d.Get("version_context")
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if _, ok := d.GetOk("vm_ware_vds_count"); ok {
		v := d.Get("vm_ware_vds_count")
		x := int64(v.(int))
		o.SetVmWareVdsCount(x)
	}

	if _, ok := d.GetOk("vmm_ctrlrp_count"); ok {
		v := d.Get("vmm_ctrlrp_count")
		x := int64(v.(int))
		o.SetVmmCtrlrpCount(x)
	}

	if _, ok := d.GetOk("vmm_domp_count"); ok {
		v := d.Get("vmm_domp_count")
		x := int64(v.(int))
		o.SetVmmDompCount(x)
	}

	if _, ok := d.GetOk("vmm_ep_pd_count"); ok {
		v := d.Get("vmm_ep_pd_count")
		x := int64(v.(int))
		o.SetVmmEpPdCount(x)
	}

	if _, ok := d.GetOk("vnsm_dev_count"); ok {
		v := d.Get("vnsm_dev_count")
		x := int64(v.(int))
		o.SetVnsmDevCount(x)
	}

	if _, ok := d.GetOk("vpod_count"); ok {
		v := d.Get("vpod_count")
		x := int64(v.(int))
		o.SetVpodCount(x)
	}

	if _, ok := d.GetOk("webtoken_timeout_seconds"); ok {
		v := d.Get("webtoken_timeout_seconds")
		x := int64(v.(int))
		o.SetWebtokenTimeoutSeconds(x)
	}

	if v, ok := d.GetOk("aaa_ldap_provider_count"); ok {
		x := int64(v.(int))
		o.SetAaaLdapProviderCount(x)
	}
	if v, ok := d.GetOk("aaa_radius_provider_count"); ok {
		x := int64(v.(int))
		o.SetAaaRadiusProviderCount(x)
	}
	if v, ok := d.GetOk("aaa_tacacs_provider_count"); ok {
		x := int64(v.(int))
		o.SetAaaTacacsProviderCount(x)
	}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("apic_count"); ok {
		x := int64(v.(int))
		o.SetApicCount(x)
	}
	if v, ok := d.GetOk("apic_is_telnet_enabled"); ok {
		x := (v.(bool))
		o.SetApicIsTelnetEnabled(x)
	}
	if v, ok := d.GetOk("apic_ntp_count"); ok {
		x := int64(v.(int))
		o.SetApicNtpCount(x)
	}
	if v, ok := d.GetOk("apic_snmp_community_count"); ok {
		x := int64(v.(int))
		o.SetApicSnmpCommunityCount(x)
	}
	if v, ok := d.GetOk("apic_sys_log_grp_count"); ok {
		x := int64(v.(int))
		o.SetApicSysLogGrpCount(x)
	}
	if v, ok := d.GetOk("apic_sys_log_src_count"); ok {
		x := int64(v.(int))
		o.SetApicSysLogSrcCount(x)
	}
	if v, ok := d.GetOk("app_center_count"); ok {
		x := int64(v.(int))
		o.SetAppCenterCount(x)
	}
	if v, ok := d.GetOk("ave"); ok {
		x := (v.(string))
		o.SetAve(x)
	}
	if v, ok := d.GetOk("bd_count"); ok {
		x := int64(v.(int))
		o.SetBdCount(x)
	}
	if v, ok := d.GetOk("callhome_smart_group_count"); ok {
		x := int64(v.(int))
		o.SetCallhomeSmartGroupCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cloud_sec_peer_count"); ok {
		x := int64(v.(int))
		o.SetCloudSecPeerCount(x)
	}
	if v, ok := d.GetOk("comp_hv_count"); ok {
		x := int64(v.(int))
		o.SetCompHvCount(x)
	}
	if v, ok := d.GetOk("config_exportp_count"); ok {
		x := int64(v.(int))
		o.SetConfigExportpCount(x)
	}
	if v, ok := d.GetOk("config_job_count"); ok {
		x := int64(v.(int))
		o.SetConfigJobCount(x)
	}
	if v, ok := d.GetOk("consistency_checker_app"); ok {
		x := (v.(string))
		o.SetConsistencyCheckerApp(x)
	}
	if v, ok := d.GetOk("contract_count"); ok {
		x := int64(v.(int))
		o.SetContractCount(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("dns_count"); ok {
		x := int64(v.(int))
		o.SetDnsCount(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("eigrp_count"); ok {
		x := int64(v.(int))
		o.SetEigrpCount(x)
	}
	if v, ok := d.GetOk("epg_count"); ok {
		x := int64(v.(int))
		o.SetEpgCount(x)
	}
	if v, ok := d.GetOk("fabric_module_count"); ok {
		x := int64(v.(int))
		o.SetFabricModuleCount(x)
	}
	if v, ok := d.GetOk("fabric_setupp_count"); ok {
		x := int64(v.(int))
		o.SetFabricSetuppCount(x)
	}
	if v, ok := d.GetOk("fcoe_nport_count"); ok {
		x := int64(v.(int))
		o.SetFcoeNportCount(x)
	}
	if v, ok := d.GetOk("fcoe_nport_dom_count"); ok {
		x := int64(v.(int))
		o.SetFcoeNportDomCount(x)
	}
	if v, ok := d.GetOk("fcoe_nport_vlan_count"); ok {
		x := int64(v.(int))
		o.SetFcoeNportVlanCount(x)
	}
	if v, ok := d.GetOk("fcoe_nport_vsan_count"); ok {
		x := int64(v.(int))
		o.SetFcoeNportVsanCount(x)
	}
	if v, ok := d.GetOk("fv_sla_def_count"); ok {
		x := int64(v.(int))
		o.SetFvSlaDefCount(x)
	}
	if v, ok := d.GetOk("hsrp_count"); ok {
		x := int64(v.(int))
		o.SetHsrpCount(x)
	}
	if v, ok := d.GetOk("ibgp_count"); ok {
		x := int64(v.(int))
		o.SetIbgpCount(x)
	}
	if v, ok := d.GetOk("igmp_access_list_count"); ok {
		x := int64(v.(int))
		o.SetIgmpAccessListCount(x)
	}
	if v, ok := d.GetOk("igmp_snoop"); ok {
		x := (v.(string))
		o.SetIgmpSnoop(x)
	}
	if v, ok := d.GetOk("ip_epg_count"); ok {
		x := int64(v.(int))
		o.SetIpEpgCount(x)
	}
	if v, ok := d.GetOk("is_bgp_route_reflectors_feature_used"); ok {
		x := (v.(bool))
		o.SetIsBgpRouteReflectorsFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_bridge_domains_feature_used"); ok {
		x := (v.(bool))
		o.SetIsBridgeDomainsFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_common_local_user_name"); ok {
		x := (v.(bool))
		o.SetIsCommonLocalUserName(x)
	}
	if v, ok := d.GetOk("is_contracts_feature_used"); ok {
		x := (v.(bool))
		o.SetIsContractsFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_epg_feature_used"); ok {
		x := (v.(bool))
		o.SetIsEpgFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_filters_feature_used"); ok {
		x := (v.(bool))
		o.SetIsFiltersFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_http_configured"); ok {
		x := (v.(bool))
		o.SetIsHttpConfigured(x)
	}
	if v, ok := d.GetOk("is_https_configured"); ok {
		x := (v.(bool))
		o.SetIsHttpsConfigured(x)
	}
	if v, ok := d.GetOk("is_ntp_feature_used"); ok {
		x := (v.(bool))
		o.SetIsNtpFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_ptp_feature_used"); ok {
		x := (v.(bool))
		o.SetIsPtpFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_synce_feature_used"); ok {
		x := (v.(bool))
		o.SetIsSynceFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_tech_support_collected"); ok {
		x := (v.(string))
		o.SetIsTechSupportCollected(x)
	}
	if v, ok := d.GetOk("is_tenants_feature_used"); ok {
		x := (v.(bool))
		o.SetIsTenantsFeatureUsed(x)
	}
	if v, ok := d.GetOk("is_vrfs_feature_used"); ok {
		x := (v.(bool))
		o.SetIsVrfsFeatureUsed(x)
	}
	if v, ok := d.GetOk("isis_count"); ok {
		x := int64(v.(int))
		o.SetIsisCount(x)
	}
	if v, ok := d.GetOk("l2_multicast"); ok {
		x := (v.(string))
		o.SetL2Multicast(x)
	}
	if v, ok := d.GetOk("leaf_count"); ok {
		x := int64(v.(int))
		o.SetLeafCount(x)
	}
	if v, ok := d.GetOk("local_username_count"); ok {
		x := int64(v.(int))
		o.SetLocalUsernameCount(x)
	}
	if v, ok := d.GetOk("login_block_duration"); ok {
		x := int64(v.(int))
		o.SetLoginBlockDuration(x)
	}
	if v, ok := d.GetOk("login_max_failed_attempts"); ok {
		x := int64(v.(int))
		o.SetLoginMaxFailedAttempts(x)
	}
	if v, ok := d.GetOk("login_max_failed_attempts_window"); ok {
		x := int64(v.(int))
		o.SetLoginMaxFailedAttemptsWindow(x)
	}
	if v, ok := d.GetOk("maintenance_mode_count"); ok {
		x := int64(v.(int))
		o.SetMaintenanceModeCount(x)
	}
	if v, ok := d.GetOk("management_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetManagementOverV6Count(x)
	}
	if v, ok := d.GetOk("microsoft_useg_vmm_ep_pd_count"); ok {
		x := int64(v.(int))
		o.SetMicrosoftUsegVmmEpPdCount(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("net_flow_count"); ok {
		x := int64(v.(int))
		o.SetNetFlowCount(x)
	}
	if v, ok := d.GetOk("nir"); ok {
		x := (v.(string))
		o.SetNir(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("open_stack"); ok {
		x := (v.(string))
		o.SetOpenStack(x)
	}
	if v, ok := d.GetOk("opflex_kubernetes_count"); ok {
		x := int64(v.(int))
		o.SetOpflexKubernetesCount(x)
	}
	if v, ok := d.GetOk("ospf_count"); ok {
		x := int64(v.(int))
		o.SetOspfCount(x)
	}
	if v, ok := d.GetOk("password_history_count"); ok {
		x := int64(v.(int))
		o.SetPasswordHistoryCount(x)
	}
	if v, ok := d.GetOk("password_strength_check"); ok {
		x := (v.(string))
		o.SetPasswordStrengthCheck(x)
	}
	if v, ok := d.GetOk("password_strength_profile_count"); ok {
		x := int64(v.(int))
		o.SetPasswordStrengthProfileCount(x)
	}
	if v, ok := d.GetOk("poe_count"); ok {
		x := int64(v.(int))
		o.SetPoeCount(x)
	}
	if v, ok := d.GetOk("port_security_count"); ok {
		x := int64(v.(int))
		o.SetPortSecurityCount(x)
	}
	if v, ok := d.GetOk("qin_vni_tunnel_count"); ok {
		x := int64(v.(int))
		o.SetQinVniTunnelCount(x)
	}
	if v, ok := d.GetOk("qos_cong_count"); ok {
		x := int64(v.(int))
		o.SetQosCongCount(x)
	}
	if v, ok := d.GetOk("qos_pfc_pol_count"); ok {
		x := int64(v.(int))
		o.SetQosPfcPolCount(x)
	}
	if v, ok := d.GetOk("realm_count"); ok {
		x := int64(v.(int))
		o.SetRealmCount(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("remote_leaf_count"); ok {
		x := int64(v.(int))
		o.SetRemoteLeafCount(x)
	}
	if v, ok := d.GetOk("scvmm_count"); ok {
		x := int64(v.(int))
		o.SetScvmmCount(x)
	}
	if v, ok := d.GetOk("shared_l3_out_count"); ok {
		x := int64(v.(int))
		o.SetSharedL3OutCount(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("smart_call_home"); ok {
		x := (v.(string))
		o.SetSmartCallHome(x)
	}
	if v, ok := d.GetOk("snapshot_count"); ok {
		x := int64(v.(int))
		o.SetSnapshotCount(x)
	}
	if v, ok := d.GetOk("snmp"); ok {
		x := (v.(string))
		o.SetSnmp(x)
	}
	if v, ok := d.GetOk("snmp_community_access_count"); ok {
		x := int64(v.(int))
		o.SetSnmpCommunityAccessCount(x)
	}
	if v, ok := d.GetOk("snmp_group_count"); ok {
		x := int64(v.(int))
		o.SetSnmpGroupCount(x)
	}
	if v, ok := d.GetOk("snmp_trap_count"); ok {
		x := int64(v.(int))
		o.SetSnmpTrapCount(x)
	}
	if v, ok := d.GetOk("snmp_v3_count"); ok {
		x := int64(v.(int))
		o.SetSnmpV3Count(x)
	}
	if v, ok := d.GetOk("span_count"); ok {
		x := int64(v.(int))
		o.SetSpanCount(x)
	}
	if v, ok := d.GetOk("span_dst_count"); ok {
		x := int64(v.(int))
		o.SetSpanDstCount(x)
	}
	if v, ok := d.GetOk("span_src_count"); ok {
		x := int64(v.(int))
		o.SetSpanSrcCount(x)
	}
	if v, ok := d.GetOk("spine_count"); ok {
		x := int64(v.(int))
		o.SetSpineCount(x)
	}
	if v, ok := d.GetOk("ssh_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetSshOverV6Count(x)
	}
	if v, ok := d.GetOk("ssh_v2_count"); ok {
		x := int64(v.(int))
		o.SetSshV2Count(x)
	}
	if v, ok := d.GetOk("supervisor_module_count"); ok {
		x := int64(v.(int))
		o.SetSupervisorModuleCount(x)
	}
	if v, ok := d.GetOk("syslog_group_count"); ok {
		x := int64(v.(int))
		o.SetSyslogGroupCount(x)
	}
	if v, ok := d.GetOk("syslog_over_v6_count"); ok {
		x := int64(v.(int))
		o.SetSyslogOverV6Count(x)
	}
	if v, ok := d.GetOk("system_controller_count"); ok {
		x := int64(v.(int))
		o.SetSystemControllerCount(x)
	}
	if v, ok := d.GetOk("tacacs_group_count"); ok {
		x := int64(v.(int))
		o.SetTacacsGroupCount(x)
	}
	if v, ok := d.GetOk("tenant_count"); ok {
		x := int64(v.(int))
		o.SetTenantCount(x)
	}
	if v, ok := d.GetOk("tier_two_leaf_count"); ok {
		x := int64(v.(int))
		o.SetTierTwoLeafCount(x)
	}
	if v, ok := d.GetOk("total_critical_faults"); ok {
		x := int64(v.(int))
		o.SetTotalCriticalFaults(x)
	}
	if v, ok := d.GetOk("twamp"); ok {
		x := (v.(string))
		o.SetTwamp(x)
	}
	if v, ok := d.GetOk("useg"); ok {
		x := (v.(string))
		o.SetUseg(x)
	}
	if v, ok := d.GetOk("vm_ware_vds_count"); ok {
		x := int64(v.(int))
		o.SetVmWareVdsCount(x)
	}
	if v, ok := d.GetOk("vmm_ctrlrp_count"); ok {
		x := int64(v.(int))
		o.SetVmmCtrlrpCount(x)
	}
	if v, ok := d.GetOk("vmm_domp_count"); ok {
		x := int64(v.(int))
		o.SetVmmDompCount(x)
	}
	if v, ok := d.GetOk("vmm_ep_pd_count"); ok {
		x := int64(v.(int))
		o.SetVmmEpPdCount(x)
	}
	if v, ok := d.GetOk("vnsm_dev_count"); ok {
		x := int64(v.(int))
		o.SetVnsmDevCount(x)
	}
	if v, ok := d.GetOk("vpod_count"); ok {
		x := int64(v.(int))
		o.SetVpodCount(x)
	}
	if v, ok := d.GetOk("webtoken_timeout_seconds"); ok {
		x := int64(v.(int))
		o.SetWebtokenTimeoutSeconds(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaFeatureUsage object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaFeatureUsageList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of NiatelemetryNiaFeatureUsage: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaFeatureUsage: %s", responseErr.Error())
	}
	count := countResponse.NiatelemetryNiaFeatureUsageList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for NiatelemetryNiaFeatureUsage data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var niatelemetryNiaFeatureUsageResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaFeatureUsageList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching NiatelemetryNiaFeatureUsage: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching NiatelemetryNiaFeatureUsage: %s", responseErr.Error())
		}
		results := resMo.NiatelemetryNiaFeatureUsageList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["aaa_ldap_provider_count"] = (s.GetAaaLdapProviderCount())
				temp["aaa_radius_provider_count"] = (s.GetAaaRadiusProviderCount())
				temp["aaa_tacacs_provider_count"] = (s.GetAaaTacacsProviderCount())
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["apic_count"] = (s.GetApicCount())
				temp["apic_is_telnet_enabled"] = (s.GetApicIsTelnetEnabled())
				temp["apic_ntp_count"] = (s.GetApicNtpCount())
				temp["apic_snmp_community_count"] = (s.GetApicSnmpCommunityCount())
				temp["apic_sys_log_grp_count"] = (s.GetApicSysLogGrpCount())
				temp["apic_sys_log_src_count"] = (s.GetApicSysLogSrcCount())
				temp["app_center_count"] = (s.GetAppCenterCount())
				temp["ave"] = (s.GetAve())
				temp["bd_count"] = (s.GetBdCount())
				temp["callhome_smart_group_count"] = (s.GetCallhomeSmartGroupCount())
				temp["class_id"] = (s.GetClassId())
				temp["cloud_sec_peer_count"] = (s.GetCloudSecPeerCount())
				temp["comp_hv_count"] = (s.GetCompHvCount())
				temp["config_exportp_count"] = (s.GetConfigExportpCount())
				temp["config_job_count"] = (s.GetConfigJobCount())
				temp["consistency_checker_app"] = (s.GetConsistencyCheckerApp())
				temp["contract_count"] = (s.GetContractCount())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["dns_count"] = (s.GetDnsCount())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["eigrp_count"] = (s.GetEigrpCount())
				temp["epg_count"] = (s.GetEpgCount())
				temp["fabric_module_count"] = (s.GetFabricModuleCount())
				temp["fabric_setupp_count"] = (s.GetFabricSetuppCount())
				temp["fcoe_nport_count"] = (s.GetFcoeNportCount())
				temp["fcoe_nport_dom_count"] = (s.GetFcoeNportDomCount())
				temp["fcoe_nport_vlan_count"] = (s.GetFcoeNportVlanCount())
				temp["fcoe_nport_vsan_count"] = (s.GetFcoeNportVsanCount())
				temp["fv_sla_def_count"] = (s.GetFvSlaDefCount())
				temp["hsrp_count"] = (s.GetHsrpCount())
				temp["ibgp_count"] = (s.GetIbgpCount())
				temp["igmp_access_list_count"] = (s.GetIgmpAccessListCount())
				temp["igmp_snoop"] = (s.GetIgmpSnoop())
				temp["ip_epg_count"] = (s.GetIpEpgCount())
				temp["is_bgp_route_reflectors_feature_used"] = (s.GetIsBgpRouteReflectorsFeatureUsed())
				temp["is_bridge_domains_feature_used"] = (s.GetIsBridgeDomainsFeatureUsed())
				temp["is_common_local_user_name"] = (s.GetIsCommonLocalUserName())
				temp["is_contracts_feature_used"] = (s.GetIsContractsFeatureUsed())
				temp["is_epg_feature_used"] = (s.GetIsEpgFeatureUsed())
				temp["is_filters_feature_used"] = (s.GetIsFiltersFeatureUsed())
				temp["is_http_configured"] = (s.GetIsHttpConfigured())
				temp["is_https_configured"] = (s.GetIsHttpsConfigured())
				temp["is_ntp_feature_used"] = (s.GetIsNtpFeatureUsed())
				temp["is_ptp_feature_used"] = (s.GetIsPtpFeatureUsed())
				temp["is_synce_feature_used"] = (s.GetIsSynceFeatureUsed())
				temp["is_tech_support_collected"] = (s.GetIsTechSupportCollected())
				temp["is_tenants_feature_used"] = (s.GetIsTenantsFeatureUsed())
				temp["is_vrfs_feature_used"] = (s.GetIsVrfsFeatureUsed())
				temp["isis_count"] = (s.GetIsisCount())
				temp["l2_multicast"] = (s.GetL2Multicast())
				temp["leaf_count"] = (s.GetLeafCount())
				temp["local_username_count"] = (s.GetLocalUsernameCount())
				temp["login_block_duration"] = (s.GetLoginBlockDuration())
				temp["login_max_failed_attempts"] = (s.GetLoginMaxFailedAttempts())
				temp["login_max_failed_attempts_window"] = (s.GetLoginMaxFailedAttemptsWindow())
				temp["maintenance_mode_count"] = (s.GetMaintenanceModeCount())
				temp["management_over_v6_count"] = (s.GetManagementOverV6Count())
				temp["microsoft_useg_vmm_ep_pd_count"] = (s.GetMicrosoftUsegVmmEpPdCount())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["net_flow_count"] = (s.GetNetFlowCount())
				temp["nir"] = (s.GetNir())
				temp["object_type"] = (s.GetObjectType())
				temp["open_stack"] = (s.GetOpenStack())
				temp["opflex_kubernetes_count"] = (s.GetOpflexKubernetesCount())
				temp["ospf_count"] = (s.GetOspfCount())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)
				temp["password_history_count"] = (s.GetPasswordHistoryCount())
				temp["password_strength_check"] = (s.GetPasswordStrengthCheck())
				temp["password_strength_profile_count"] = (s.GetPasswordStrengthProfileCount())

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["poe_count"] = (s.GetPoeCount())
				temp["port_security_count"] = (s.GetPortSecurityCount())
				temp["qin_vni_tunnel_count"] = (s.GetQinVniTunnelCount())
				temp["qos_cong_count"] = (s.GetQosCongCount())
				temp["qos_pfc_pol_count"] = (s.GetQosPfcPolCount())
				temp["realm_count"] = (s.GetRealmCount())
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["remote_leaf_count"] = (s.GetRemoteLeafCount())
				temp["scvmm_count"] = (s.GetScvmmCount())
				temp["shared_l3_out_count"] = (s.GetSharedL3OutCount())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["site_name"] = (s.GetSiteName())
				temp["smart_call_home"] = (s.GetSmartCallHome())

				temp["smart_license"] = flattenMapNiatelemetrySmartLicense(s.GetSmartLicense(), d)
				temp["snapshot_count"] = (s.GetSnapshotCount())
				temp["snmp"] = (s.GetSnmp())
				temp["snmp_community_access_count"] = (s.GetSnmpCommunityAccessCount())
				temp["snmp_group_count"] = (s.GetSnmpGroupCount())
				temp["snmp_trap_count"] = (s.GetSnmpTrapCount())
				temp["snmp_v3_count"] = (s.GetSnmpV3Count())
				temp["span_count"] = (s.GetSpanCount())
				temp["span_dst_count"] = (s.GetSpanDstCount())
				temp["span_src_count"] = (s.GetSpanSrcCount())
				temp["spine_count"] = (s.GetSpineCount())
				temp["ssh_over_v6_count"] = (s.GetSshOverV6Count())
				temp["ssh_v2_count"] = (s.GetSshV2Count())
				temp["supervisor_module_count"] = (s.GetSupervisorModuleCount())
				temp["syslog_group_count"] = (s.GetSyslogGroupCount())
				temp["syslog_over_v6_count"] = (s.GetSyslogOverV6Count())
				temp["system_controller_count"] = (s.GetSystemControllerCount())
				temp["tacacs_group_count"] = (s.GetTacacsGroupCount())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["tenant_count"] = (s.GetTenantCount())
				temp["tier_two_leaf_count"] = (s.GetTierTwoLeafCount())
				temp["total_critical_faults"] = (s.GetTotalCriticalFaults())
				temp["twamp"] = (s.GetTwamp())
				temp["useg"] = (s.GetUseg())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vm_ware_vds_count"] = (s.GetVmWareVdsCount())
				temp["vmm_ctrlrp_count"] = (s.GetVmmCtrlrpCount())
				temp["vmm_domp_count"] = (s.GetVmmDompCount())
				temp["vmm_ep_pd_count"] = (s.GetVmmEpPdCount())
				temp["vnsm_dev_count"] = (s.GetVnsmDevCount())
				temp["vpod_count"] = (s.GetVpodCount())
				temp["webtoken_timeout_seconds"] = (s.GetWebtokenTimeoutSeconds())
				niatelemetryNiaFeatureUsageResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaFeatureUsageResults))
	if err := d.Set("results", niatelemetryNiaFeatureUsageResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaFeatureUsageResults[0]["moid"].(string))
	return de
}
