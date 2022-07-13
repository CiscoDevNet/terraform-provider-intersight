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

func dataSourceNiatelemetryNiaInventoryFabric() *schema.Resource {
	var subSchema = map[string]*schema.Schema{"account_moid": {
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
		"anycast_gw_mac": {
			Description: "Returns the aycast gateway mac.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"bgp_established_interface_count": {
			Description: "Counts the number of BGP interfaces that are in established state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"bgw_count": {
			Description: "Returns number of bgw switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"bgw_interface_up_count": {
			Description: "Count number of active interfaces on border gateways.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"border_gateway_spine_count": {
			Description: "Count number of border gateway spines in the fabric inventory.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"border_leaf_count": {
			Description: "Count number of border leafs in the fabric inventory.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cloudsec_autoconfig": {
			Description: "Cloudsec autoconfig details on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dci_subnet_range": {
			Description: "Returns the dci subnet range.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dci_subnet_target_mask": {
			Description: "Returns the dci subnet target mask.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dcnmtracker_enabled": {
			Description: "Returns the value of the dcnmtrackerEnabled field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ebgp_evpn_link_up_count": {
			Description: "Count number of ebgp evpn active interfaces.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_id": {
			Description: "Uniquely identifies a fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_name": {
			Description: "Returns the value of the Name of a fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_parent": {
			Description: "Parent of the fabric on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_technology": {
			Description: "Fabric Technology details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_type": {
			Description: "Fabric type information string.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"feature_ptp": {
			Description: "PTP feature details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_bgw_present": {
			Description: "Checks if border gateway is present in the fabric inventory.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_enable_nxapi_http": {
			Description: "Check if NXAPI HTTP is enabled or not on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_enable_real_time_backup": {
			Description: "Check if real time backup is enabled or not on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ngoam_enabled": {
			Description: "Returns if ngoam is enabled.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_scheduled_back_up_enabled": {
			Description: "Returns if the scheduled backup is enabled.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_trm_enabled": {
			Description: "Is TRM enabled for the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"leaf_count": {
			Description: "Returns total number of leafs in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"link_state_routing": {
			Description: "Link state routing details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"link_type": {
			Description: "Fabric oper status information.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"logical_links": {
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
					"db_id": {
						Description: "Return value of dbId attribute.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"is_present": {
						Description: "Return value of isPresent attribute.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"link_addr1": {
						Description: "Return value of linkAddr1 attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_addr2": {
						Description: "Return value of linkAddr2 attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_state": {
						Description: "Return value of linkState attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_type": {
						Description: "Return value of linkType attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"uptime": {
						Description: "Return value of uptime attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		"network_deployment_count": {
			Description: "No of networks deployed on a fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"network_deployment_status": {
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
					"id": {
						Description: "Returns the id of network/vrf.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Returns the name of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Returns the deployment status of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ntp_server_ip_list": {
			Description: "NTP server IP List on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_vni_bw_sites_count": {
			Description: "Returns the count of vnis between sites.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vrf_bw_sites_count": {
			Description: "Returns the count of vrfs between sites.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vrf_count": {
			Description: "Returns the value of the nxosVrfCount field.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_status": {
			Description: "Fabric oper status information.",
			Type:        schema.TypeString,
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
		"replication_mode": {
			Description: "Replication mode details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"rp_mode": {
			Description: "RP Mode details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial": {
			Description: "Serial number of device being inventoried. The serial number is unique per device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"site_name": {
			Description: "Name of fabric domain of the controller.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"software_image": {
			Description: "Software image details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"spine_count": {
			Description: "Returns total number of spines in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_server_ip_list": {
			Description: "Syslog server IP list on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"syslog_sev": {
			Description: "Syslog sev details on the fabric.",
			Type:        schema.TypeString,
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
		"template_name": {
			Description: "Template name of the fabric on DCNM.",
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
		"vlan_vni_mappings": {
			Description: "VLAN to VNI mappings configured in the DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vni_ip_count": {
			Description: "Count number of IP addresses configured in the DCNM networks.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpc_details": {
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
					"is_vpc_configured": {
						Description: "Returns boolean if VPC is configured on switch or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"peer_switch_db_id": {
						Description: "Returns peer switch id if VPC configured.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"switch_db_id": {
						Description: "Returns the switch id of the switch.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"vrf_deployment_count": {
			Description: "No of vrfs deployed on a fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vrf_deployment_status": {
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
					"id": {
						Description: "Returns the id of network/vrf.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Returns the name of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Returns the deployment status of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"xsite_network_count": {
			Description: "Returns deployed network count for bgw/bgws switches in the MSD fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"xsite_vrf_count": {
			Description: "Returns deployed vrf count for bgw/bgws switches in the MSD fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
	var model = map[string]*schema.Schema{"account_moid": {
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
		"anycast_gw_mac": {
			Description: "Returns the aycast gateway mac.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"bgp_established_interface_count": {
			Description: "Counts the number of BGP interfaces that are in established state.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"bgw_count": {
			Description: "Returns number of bgw switches in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"bgw_interface_up_count": {
			Description: "Count number of active interfaces on border gateways.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"border_gateway_spine_count": {
			Description: "Count number of border gateway spines in the fabric inventory.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"border_leaf_count": {
			Description: "Count number of border leafs in the fabric inventory.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cloudsec_autoconfig": {
			Description: "Cloudsec autoconfig details on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dci_subnet_range": {
			Description: "Returns the dci subnet range.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dci_subnet_target_mask": {
			Description: "Returns the dci subnet target mask.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dcnmtracker_enabled": {
			Description: "Returns the value of the dcnmtrackerEnabled field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ebgp_evpn_link_up_count": {
			Description: "Count number of ebgp evpn active interfaces.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"fabric_id": {
			Description: "Uniquely identifies a fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_name": {
			Description: "Returns the value of the Name of a fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_parent": {
			Description: "Parent of the fabric on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_technology": {
			Description: "Fabric Technology details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_type": {
			Description: "Fabric type information string.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"feature_ptp": {
			Description: "PTP feature details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_bgw_present": {
			Description: "Checks if border gateway is present in the fabric inventory.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_enable_nxapi_http": {
			Description: "Check if NXAPI HTTP is enabled or not on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_enable_real_time_backup": {
			Description: "Check if real time backup is enabled or not on the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_ngoam_enabled": {
			Description: "Returns if ngoam is enabled.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_scheduled_back_up_enabled": {
			Description: "Returns if the scheduled backup is enabled.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_trm_enabled": {
			Description: "Is TRM enabled for the fabric.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"leaf_count": {
			Description: "Returns total number of leafs in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"link_state_routing": {
			Description: "Link state routing details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"link_type": {
			Description: "Fabric oper status information.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"logical_links": {
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
					"db_id": {
						Description: "Return value of dbId attribute.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"is_present": {
						Description: "Return value of isPresent attribute.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"link_addr1": {
						Description: "Return value of linkAddr1 attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_addr2": {
						Description: "Return value of linkAddr2 attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_state": {
						Description: "Return value of linkState attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"link_type": {
						Description: "Return value of linkType attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"uptime": {
						Description: "Return value of uptime attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		"network_deployment_count": {
			Description: "No of networks deployed on a fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"network_deployment_status": {
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
					"id": {
						Description: "Returns the id of network/vrf.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Returns the name of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Returns the deployment status of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ntp_server_ip_list": {
			Description: "NTP server IP List on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_vni_bw_sites_count": {
			Description: "Returns the count of vnis between sites.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vrf_bw_sites_count": {
			Description: "Returns the count of vrfs between sites.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vrf_count": {
			Description: "Returns the value of the nxosVrfCount field.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_status": {
			Description: "Fabric oper status information.",
			Type:        schema.TypeString,
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
		"replication_mode": {
			Description: "Replication mode details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"rp_mode": {
			Description: "RP Mode details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial": {
			Description: "Serial number of device being inventoried. The serial number is unique per device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"site_name": {
			Description: "Name of fabric domain of the controller.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"software_image": {
			Description: "Software image details on the fabric.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"spine_count": {
			Description: "Returns total number of spines in the fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"syslog_server_ip_list": {
			Description: "Syslog server IP list on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"syslog_sev": {
			Description: "Syslog sev details on the fabric.",
			Type:        schema.TypeString,
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
		"template_name": {
			Description: "Template name of the fabric on DCNM.",
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
		"vlan_vni_mappings": {
			Description: "VLAN to VNI mappings configured in the DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vni_ip_count": {
			Description: "Count number of IP addresses configured in the DCNM networks.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpc_details": {
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
					"is_vpc_configured": {
						Description: "Returns boolean if VPC is configured on switch or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"peer_switch_db_id": {
						Description: "Returns peer switch id if VPC configured.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"switch_db_id": {
						Description: "Returns the switch id of the switch.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"vrf_deployment_count": {
			Description: "No of vrfs deployed on a fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vrf_deployment_status": {
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
					"id": {
						Description: "Returns the id of network/vrf.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Returns the name of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Returns the deployment status of network/vrf.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"xsite_network_count": {
			Description: "Returns deployed network count for bgw/bgws switches in the MSD fabric.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"xsite_vrf_count": {
			Description: "Returns deployed vrf count for bgw/bgws switches in the MSD fabric.",
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
		ReadContext: dataSourceNiatelemetryNiaInventoryFabricRead,
		Schema:      model}
}

func dataSourceNiatelemetryNiaInventoryFabricRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventoryFabric{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
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

	if v, ok := d.GetOk("anycast_gw_mac"); ok {
		x := (v.(string))
		o.SetAnycastGwMac(x)
	}

	if v, ok := d.GetOkExists("bgp_established_interface_count"); ok {
		x := int64(v.(int))
		o.SetBgpEstablishedInterfaceCount(x)
	}

	if v, ok := d.GetOkExists("bgw_count"); ok {
		x := int64(v.(int))
		o.SetBgwCount(x)
	}

	if v, ok := d.GetOkExists("bgw_interface_up_count"); ok {
		x := int64(v.(int))
		o.SetBgwInterfaceUpCount(x)
	}

	if v, ok := d.GetOkExists("border_gateway_spine_count"); ok {
		x := int64(v.(int))
		o.SetBorderGatewaySpineCount(x)
	}

	if v, ok := d.GetOkExists("border_leaf_count"); ok {
		x := int64(v.(int))
		o.SetBorderLeafCount(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOkExists("cloudsec_autoconfig"); ok {
		x := (v.(bool))
		o.SetCloudsecAutoconfig(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("dci_subnet_range"); ok {
		x := (v.(string))
		o.SetDciSubnetRange(x)
	}

	if v, ok := d.GetOk("dci_subnet_target_mask"); ok {
		x := (v.(string))
		o.SetDciSubnetTargetMask(x)
	}

	if v, ok := d.GetOkExists("dcnmtracker_enabled"); ok {
		x := (v.(bool))
		o.SetDcnmtrackerEnabled(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOkExists("ebgp_evpn_link_up_count"); ok {
		x := int64(v.(int))
		o.SetEbgpEvpnLinkUpCount(x)
	}

	if v, ok := d.GetOk("fabric_id"); ok {
		x := (v.(string))
		o.SetFabricId(x)
	}

	if v, ok := d.GetOk("fabric_name"); ok {
		x := (v.(string))
		o.SetFabricName(x)
	}

	if v, ok := d.GetOk("fabric_parent"); ok {
		x := (v.(string))
		o.SetFabricParent(x)
	}

	if v, ok := d.GetOk("fabric_technology"); ok {
		x := (v.(string))
		o.SetFabricTechnology(x)
	}

	if v, ok := d.GetOk("fabric_type"); ok {
		x := (v.(string))
		o.SetFabricType(x)
	}

	if v, ok := d.GetOk("feature_ptp"); ok {
		x := (v.(string))
		o.SetFeaturePtp(x)
	}

	if v, ok := d.GetOkExists("is_bgw_present"); ok {
		x := (v.(bool))
		o.SetIsBgwPresent(x)
	}

	if v, ok := d.GetOkExists("is_enable_nxapi_http"); ok {
		x := (v.(bool))
		o.SetIsEnableNxapiHttp(x)
	}

	if v, ok := d.GetOkExists("is_enable_real_time_backup"); ok {
		x := (v.(bool))
		o.SetIsEnableRealTimeBackup(x)
	}

	if v, ok := d.GetOkExists("is_ngoam_enabled"); ok {
		x := (v.(bool))
		o.SetIsNgoamEnabled(x)
	}

	if v, ok := d.GetOkExists("is_scheduled_back_up_enabled"); ok {
		x := (v.(bool))
		o.SetIsScheduledBackUpEnabled(x)
	}

	if v, ok := d.GetOkExists("is_trm_enabled"); ok {
		x := (v.(bool))
		o.SetIsTrmEnabled(x)
	}

	if v, ok := d.GetOkExists("leaf_count"); ok {
		x := int64(v.(int))
		o.SetLeafCount(x)
	}

	if v, ok := d.GetOk("link_state_routing"); ok {
		x := (v.(string))
		o.SetLinkStateRouting(x)
	}

	if v, ok := d.GetOk("link_type"); ok {
		x := (v.(string))
		o.SetLinkType(x)
	}

	if v, ok := d.GetOk("logical_links"); ok {
		x := make([]models.NiatelemetryLogicalLink, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryLogicalLink{}
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
			o.SetClassId("niatelemetry.LogicalLink")
			if v, ok := l["db_id"]; ok {
				{
					x := int64(v.(int))
					o.SetDbId(x)
				}
			}
			if v, ok := l["is_present"]; ok {
				{
					x := (v.(bool))
					o.SetIsPresent(x)
				}
			}
			if v, ok := l["link_addr1"]; ok {
				{
					x := (v.(string))
					o.SetLinkAddr1(x)
				}
			}
			if v, ok := l["link_addr2"]; ok {
				{
					x := (v.(string))
					o.SetLinkAddr2(x)
				}
			}
			if v, ok := l["link_state"]; ok {
				{
					x := (v.(string))
					o.SetLinkState(x)
				}
			}
			if v, ok := l["link_type"]; ok {
				{
					x := (v.(string))
					o.SetLinkType(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["uptime"]; ok {
				{
					x := (v.(string))
					o.SetUptime(x)
				}
			}
			x = append(x, *o)
		}
		o.SetLogicalLinks(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOkExists("network_deployment_count"); ok {
		x := int64(v.(int))
		o.SetNetworkDeploymentCount(x)
	}

	if v, ok := d.GetOk("network_deployment_status"); ok {
		x := make([]models.NiatelemetryDeploymentStatus, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryDeploymentStatus{}
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
			o.SetClassId("niatelemetry.DeploymentStatus")
			if v, ok := l["id"]; ok {
				{
					x := int64(v.(int))
					o.SetId(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			x = append(x, *o)
		}
		o.SetNetworkDeploymentStatus(x)
	}

	if v, ok := d.GetOk("ntp_server_ip_list"); ok {
		x := (v.(string))
		o.SetNtpServerIpList(x)
	}

	if v, ok := d.GetOkExists("nxos_vni_bw_sites_count"); ok {
		x := int64(v.(int))
		o.SetNxosVniBwSitesCount(x)
	}

	if v, ok := d.GetOkExists("nxos_vrf_bw_sites_count"); ok {
		x := int64(v.(int))
		o.SetNxosVrfBwSitesCount(x)
	}

	if v, ok := d.GetOkExists("nxos_vrf_count"); ok {
		x := int64(v.(int))
		o.SetNxosVrfCount(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("oper_status"); ok {
		x := (v.(string))
		o.SetOperStatus(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
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

	if v, ok := d.GetOk("permission_resources"); ok {
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

	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}

	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
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

	if v, ok := d.GetOk("replication_mode"); ok {
		x := (v.(string))
		o.SetReplicationMode(x)
	}

	if v, ok := d.GetOk("rp_mode"); ok {
		x := (v.(string))
		o.SetRpMode(x)
	}

	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}

	if v, ok := d.GetOk("software_image"); ok {
		x := (v.(string))
		o.SetSoftwareImage(x)
	}

	if v, ok := d.GetOkExists("spine_count"); ok {
		x := int64(v.(int))
		o.SetSpineCount(x)
	}

	if v, ok := d.GetOk("syslog_server_ip_list"); ok {
		x := (v.(string))
		o.SetSyslogServerIpList(x)
	}

	if v, ok := d.GetOk("syslog_sev"); ok {
		x := (v.(string))
		o.SetSyslogSev(x)
	}

	if v, ok := d.GetOk("tags"); ok {
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

	if v, ok := d.GetOk("template_name"); ok {
		x := (v.(string))
		o.SetTemplateName(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
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

	if v, ok := d.GetOk("vlan_vni_mappings"); ok {
		x := (v.(string))
		o.SetVlanVniMappings(x)
	}

	if v, ok := d.GetOkExists("vni_ip_count"); ok {
		x := int64(v.(int))
		o.SetVniIpCount(x)
	}

	if v, ok := d.GetOk("vpc_details"); ok {
		x := make([]models.NiatelemetryVpcDetails, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryVpcDetails{}
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
			o.SetClassId("niatelemetry.VpcDetails")
			if v, ok := l["is_vpc_configured"]; ok {
				{
					x := (v.(bool))
					o.SetIsVpcConfigured(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["peer_switch_db_id"]; ok {
				{
					x := int64(v.(int))
					o.SetPeerSwitchDbId(x)
				}
			}
			if v, ok := l["switch_db_id"]; ok {
				{
					x := int64(v.(int))
					o.SetSwitchDbId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetVpcDetails(x)
	}

	if v, ok := d.GetOkExists("vrf_deployment_count"); ok {
		x := int64(v.(int))
		o.SetVrfDeploymentCount(x)
	}

	if v, ok := d.GetOk("vrf_deployment_status"); ok {
		x := make([]models.NiatelemetryDeploymentStatus, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryDeploymentStatus{}
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
			o.SetClassId("niatelemetry.DeploymentStatus")
			if v, ok := l["id"]; ok {
				{
					x := int64(v.(int))
					o.SetId(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			x = append(x, *o)
		}
		o.SetVrfDeploymentStatus(x)
	}

	if v, ok := d.GetOkExists("xsite_network_count"); ok {
		x := int64(v.(int))
		o.SetXsiteNetworkCount(x)
	}

	if v, ok := d.GetOkExists("xsite_vrf_count"); ok {
		x := int64(v.(int))
		o.SetXsiteVrfCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventoryFabric object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryFabricList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryFabric: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryFabric: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for NiatelemetryNiaInventoryFabric data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var niatelemetryNiaInventoryFabricResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryFabricList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryFabric: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryFabric: %s", responseErr.Error())
		}
		results := resMo.NiatelemetryNiaInventoryFabricList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["anycast_gw_mac"] = (s.GetAnycastGwMac())
				temp["bgp_established_interface_count"] = (s.GetBgpEstablishedInterfaceCount())
				temp["bgw_count"] = (s.GetBgwCount())
				temp["bgw_interface_up_count"] = (s.GetBgwInterfaceUpCount())
				temp["border_gateway_spine_count"] = (s.GetBorderGatewaySpineCount())
				temp["border_leaf_count"] = (s.GetBorderLeafCount())
				temp["class_id"] = (s.GetClassId())
				temp["cloudsec_autoconfig"] = (s.GetCloudsecAutoconfig())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["dci_subnet_range"] = (s.GetDciSubnetRange())
				temp["dci_subnet_target_mask"] = (s.GetDciSubnetTargetMask())
				temp["dcnmtracker_enabled"] = (s.GetDcnmtrackerEnabled())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["ebgp_evpn_link_up_count"] = (s.GetEbgpEvpnLinkUpCount())
				temp["fabric_id"] = (s.GetFabricId())
				temp["fabric_name"] = (s.GetFabricName())
				temp["fabric_parent"] = (s.GetFabricParent())
				temp["fabric_technology"] = (s.GetFabricTechnology())
				temp["fabric_type"] = (s.GetFabricType())
				temp["feature_ptp"] = (s.GetFeaturePtp())
				temp["is_bgw_present"] = (s.GetIsBgwPresent())
				temp["is_enable_nxapi_http"] = (s.GetIsEnableNxapiHttp())
				temp["is_enable_real_time_backup"] = (s.GetIsEnableRealTimeBackup())
				temp["is_ngoam_enabled"] = (s.GetIsNgoamEnabled())
				temp["is_scheduled_back_up_enabled"] = (s.GetIsScheduledBackUpEnabled())
				temp["is_trm_enabled"] = (s.GetIsTrmEnabled())
				temp["leaf_count"] = (s.GetLeafCount())
				temp["link_state_routing"] = (s.GetLinkStateRouting())
				temp["link_type"] = (s.GetLinkType())

				temp["logical_links"] = flattenListNiatelemetryLogicalLink(s.GetLogicalLinks(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["network_deployment_count"] = (s.GetNetworkDeploymentCount())

				temp["network_deployment_status"] = flattenListNiatelemetryDeploymentStatus(s.GetNetworkDeploymentStatus(), d)
				temp["ntp_server_ip_list"] = (s.GetNtpServerIpList())
				temp["nxos_vni_bw_sites_count"] = (s.GetNxosVniBwSitesCount())
				temp["nxos_vrf_bw_sites_count"] = (s.GetNxosVrfBwSitesCount())
				temp["nxos_vrf_count"] = (s.GetNxosVrfCount())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_status"] = (s.GetOperStatus())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["replication_mode"] = (s.GetReplicationMode())
				temp["rp_mode"] = (s.GetRpMode())
				temp["serial"] = (s.GetSerial())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["site_name"] = (s.GetSiteName())
				temp["software_image"] = (s.GetSoftwareImage())
				temp["spine_count"] = (s.GetSpineCount())
				temp["syslog_server_ip_list"] = (s.GetSyslogServerIpList())
				temp["syslog_sev"] = (s.GetSyslogSev())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["template_name"] = (s.GetTemplateName())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vlan_vni_mappings"] = (s.GetVlanVniMappings())
				temp["vni_ip_count"] = (s.GetVniIpCount())

				temp["vpc_details"] = flattenListNiatelemetryVpcDetails(s.GetVpcDetails(), d)
				temp["vrf_deployment_count"] = (s.GetVrfDeploymentCount())

				temp["vrf_deployment_status"] = flattenListNiatelemetryDeploymentStatus(s.GetVrfDeploymentStatus(), d)
				temp["xsite_network_count"] = (s.GetXsiteNetworkCount())
				temp["xsite_vrf_count"] = (s.GetXsiteVrfCount())
				niatelemetryNiaInventoryFabricResults = append(niatelemetryNiaInventoryFabricResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaInventoryFabricResults))
	if err := d.Set("results", niatelemetryNiaInventoryFabricResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaInventoryFabricResults[0]["moid"].(string))
	return de
}
