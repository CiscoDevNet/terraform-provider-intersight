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

func dataSourceNiatelemetryNiaInventory() *schema.Resource {
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
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"controller_ip_address": {
			Description: "Returns controller's IP address details.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu": {
			Description: "CPU usage of device being inventoried. This determines the percentage of CPU resources used.",
			Type:        schema.TypeFloat,
			Optional:    true,
		},
		"crash_reset_logs": {
			Description: "Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"customer_device_connector": {
			Description: "Returns the value of the customerDeviceConnector field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dcnm_license_state": {
			Description: "Returns the License state of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_discovery": {
			Description: "Returns the value of the deviceDiscovery field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_health": {
			Description: "Returns the device health.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"device_id": {
			Description: "Returns the value of the deviceId field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_name": {
			Description: "Name of device being inventoried. The name the user assigns to the device is inventoried here.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_type": {
			Description: "Type of device being inventoried. This determines whether the device is a controller, leaf or spine.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_up_time": {
			Description: "Returns the device up time.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"disk": {
			Description: "Disk Usage of device being inventoried. This determines the amount of disk usage.",
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
					"free": {
						Description: "The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"total": {
						Description: "The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"used": {
						Description: "The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"dn": {
			Description: "Dn for the inventories present.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_name": {
			Description: "Name of the fabric of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fex_count": {
			Description: "Number of fabric extendors utilized.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"infra_wi_node_count": {
			Description: "Number of appliances as physical device that are wired into the cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"interface": {
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
					"name": {
						Description: "Return value of name of the port.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_state": {
						Description: "Return value of operState attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"xcvr_present": {
						Description: "Return whether sfp is present or not.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ip_address": {
			Description: "The IP address of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_virtual_node": {
			Description: "Flag to specify if the node is virtual.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"last_reboot_time": {
			Description: "Returns the last reboot Time of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"last_reset_reason": {
			Description: "Returns the last reset reason of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"license_state": {
			Description: "A reference to a niatelemetryNiaLicenseState resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"license_type": {
			Description: "Returns the License type of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"log_in_time": {
			Description: "Last log in time device being inventoried. This determines the last login time on the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"log_out_time": {
			Description: "Last log out time of device being inventoried. This determines the last logout time on the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_sec_count": {
			Description: "Number of Macsec configured interfaces on a TOR.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mac_sec_fab_count": {
			Description: "Number of Macsec configured interfaces on a Spine.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"macsec_total_count": {
			Description: "Number of total Macsec configured interfaces for all nodes.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"memory": {
			Description: "Memory usage of device being inventoried. This determines the percentage of memory resources used.",
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
		"node_id": {
			Description: "The ID of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_bgp_evpn": {
			Description: "Returns Nxos EVPN BGP details.",
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
					"nxos_evpn_mac_count": {
						Description: "Returns the EVPN mac count.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"total_networks": {
						Description: "Returns the BGP EVPN total networks.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_paths": {
						Description: "Returns the BGP EVPN total paths.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_bgp_mvpn": {
			Description: "Returns the value of the nxosTrmConfigs field.",
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
					"capable_peers": {
						Description: "Return count of BGP MVPN table capable peers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configured_peers": {
						Description: "Return count of BGP MVPN configured peers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"memory_used": {
						Description: "Return value of BGP MVPN memory used.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_cluster_lists": {
						Description: "Return value of BGP MVPN cluster list.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_communities": {
						Description: "Return count of BGP MVPN communities.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"table_version": {
						Description: "Return value of BGP MVPN table version.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_networks": {
						Description: "Return count of BGP MVPN networks.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_paths": {
						Description: "Return count of BGP MVPN paths.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_bootflash_details": {
			Description: "Returns the value of the nxosBootflashDetails field.",
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
					"fw_rev": {
						Description: "Return firmware revision in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"model_type": {
						Description: "Return model type in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"serial": {
						Description: "Return serial id in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_dci_interface_status": {
			Description: "Returns the status of dci interface configured.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_interface_brief": {
			Description: "Returns the value of the nxosInterfaceBrief field.",
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
					"interface_down_count": {
						Description: "Return value of number of interafces down.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"interface_up_count": {
						Description: "Return value of number of interafces up.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_nve_interface_status": {
			Description: "Returns the value of the nxosNveInterface field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_nve_packet_counters": {
			Description: "Returns the value of the nxosNvePacketCounters field.",
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
					"mcast_inpkts": {
						Description: "Return mcast in packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"mcast_outbytes": {
						Description: "Return mcast outbytes count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ucast_inpkts": {
						Description: "Return ucast in packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ucast_outpkts": {
						Description: "Return ucast out packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_nve_vni": {
			Description: "Returns the value of the nxosNveVni field.",
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
					"cp_vni_count": {
						Description: "Return value of cp vni count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"cp_vni_down": {
						Description: "Return value of cp vni down count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"cp_vni_up": {
						Description: "Return value of cp vni up count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_count": {
						Description: "Return value of dp vni count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_down": {
						Description: "Return value of cp vni down count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_up": {
						Description: "Return value of cp vni up count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_ospf_neighbors": {
			Description: "Total number of ospf neighbors per switch in DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_pim_neighbors": {
			Description: "Total number of pim neighbors per switch in DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_telnet": {
			Description: "Returns the value of the nxosTelnet field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_total_routes": {
			Description: "Total number of routes configured in the DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vtp": {
			Description: "Returns the value of the nxosVtpStatus field.",
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
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_mode": {
						Description: "Returns the status of operational mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"pruning_mode": {
						Description: "Returns the status pruning mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"running_version": {
						Description: "Returns the running version of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"trap_enabled": {
						Description: "Returns the status of trap in vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"v2_mode": {
						Description: "Returns the status of v2 mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Returns version of vtp running.",
						Type:        schema.TypeInt,
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
		"route_prefix_count": {
			Description: "Total nuumber of v4 and v6 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"route_prefix_v4_count": {
			Description: "Number of v4 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"route_prefix_v6_count": {
			Description: "Number of v6 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"serial": {
			Description: "Serial number of device being invetoried. The serial number is unique per device and will be used as the key.",
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
		"smart_account_id": {
			Description: "Returns the value of the smartAccountId/CustomerId field.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"software_download": {
			Description: "Last software downloaded of device being inventoried. This determines if software download API was used.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"system_up_time": {
			Description: "The amount of time that the device being inventoried been up.",
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
		"nr_version": {
			Description: "Software version of device being inventoried. The various software version values for each device are available on cisco.com.",
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
		"vni_status": {
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
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni": {
						Description: "Returns the vni id of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni_state": {
						Description: "Returns the vni state of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni_type": {
						Description: "Returns the vni type of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"controller_ip_address": {
			Description: "Returns controller's IP address details.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu": {
			Description: "CPU usage of device being inventoried. This determines the percentage of CPU resources used.",
			Type:        schema.TypeFloat,
			Optional:    true,
		},
		"crash_reset_logs": {
			Description: "Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"customer_device_connector": {
			Description: "Returns the value of the customerDeviceConnector field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dcnm_license_state": {
			Description: "Returns the License state of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_discovery": {
			Description: "Returns the value of the deviceDiscovery field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_health": {
			Description: "Returns the device health.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"device_id": {
			Description: "Returns the value of the deviceId field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_name": {
			Description: "Name of device being inventoried. The name the user assigns to the device is inventoried here.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_type": {
			Description: "Type of device being inventoried. This determines whether the device is a controller, leaf or spine.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_up_time": {
			Description: "Returns the device up time.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"disk": {
			Description: "Disk Usage of device being inventoried. This determines the amount of disk usage.",
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
					"free": {
						Description: "The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"total": {
						Description: "The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"used": {
						Description: "The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"dn": {
			Description: "Dn for the inventories present.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fabric_name": {
			Description: "Name of the fabric of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"fex_count": {
			Description: "Number of fabric extendors utilized.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"infra_wi_node_count": {
			Description: "Number of appliances as physical device that are wired into the cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"interface": {
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
					"name": {
						Description: "Return value of name of the port.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_state": {
						Description: "Return value of operState attribute.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"xcvr_present": {
						Description: "Return whether sfp is present or not.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ip_address": {
			Description: "The IP address of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_virtual_node": {
			Description: "Flag to specify if the node is virtual.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"last_reboot_time": {
			Description: "Returns the last reboot Time of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"last_reset_reason": {
			Description: "Returns the last reset reason of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"license_state": {
			Description: "A reference to a niatelemetryNiaLicenseState resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"license_type": {
			Description: "Returns the License type of the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"log_in_time": {
			Description: "Last log in time device being inventoried. This determines the last login time on the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"log_out_time": {
			Description: "Last log out time of device being inventoried. This determines the last logout time on the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_sec_count": {
			Description: "Number of Macsec configured interfaces on a TOR.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mac_sec_fab_count": {
			Description: "Number of Macsec configured interfaces on a Spine.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"macsec_total_count": {
			Description: "Number of total Macsec configured interfaces for all nodes.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"memory": {
			Description: "Memory usage of device being inventoried. This determines the percentage of memory resources used.",
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
		"node_id": {
			Description: "The ID of the device being inventoried.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_bgp_evpn": {
			Description: "Returns Nxos EVPN BGP details.",
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
					"nxos_evpn_mac_count": {
						Description: "Returns the EVPN mac count.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"total_networks": {
						Description: "Returns the BGP EVPN total networks.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_paths": {
						Description: "Returns the BGP EVPN total paths.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_bgp_mvpn": {
			Description: "Returns the value of the nxosTrmConfigs field.",
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
					"capable_peers": {
						Description: "Return count of BGP MVPN table capable peers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configured_peers": {
						Description: "Return count of BGP MVPN configured peers.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"memory_used": {
						Description: "Return value of BGP MVPN memory used.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_cluster_lists": {
						Description: "Return value of BGP MVPN cluster list.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_communities": {
						Description: "Return count of BGP MVPN communities.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"table_version": {
						Description: "Return value of BGP MVPN table version.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_networks": {
						Description: "Return count of BGP MVPN networks.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"total_paths": {
						Description: "Return count of BGP MVPN paths.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_bootflash_details": {
			Description: "Returns the value of the nxosBootflashDetails field.",
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
					"fw_rev": {
						Description: "Return firmware revision in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"model_type": {
						Description: "Return model type in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"serial": {
						Description: "Return serial id in boot flash details.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_dci_interface_status": {
			Description: "Returns the status of dci interface configured.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_interface_brief": {
			Description: "Returns the value of the nxosInterfaceBrief field.",
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
					"interface_down_count": {
						Description: "Return value of number of interafces down.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"interface_up_count": {
						Description: "Return value of number of interafces up.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_nve_interface_status": {
			Description: "Returns the value of the nxosNveInterface field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_nve_packet_counters": {
			Description: "Returns the value of the nxosNvePacketCounters field.",
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
					"mcast_inpkts": {
						Description: "Return mcast in packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"mcast_outbytes": {
						Description: "Return mcast outbytes count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ucast_inpkts": {
						Description: "Return ucast in packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"ucast_outpkts": {
						Description: "Return ucast out packet count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"nxos_nve_vni": {
			Description: "Returns the value of the nxosNveVni field.",
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
					"cp_vni_count": {
						Description: "Return value of cp vni count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"cp_vni_down": {
						Description: "Return value of cp vni down count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"cp_vni_up": {
						Description: "Return value of cp vni up count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_count": {
						Description: "Return value of dp vni count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_down": {
						Description: "Return value of cp vni down count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"dp_vni_up": {
						Description: "Return value of cp vni up count.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nxos_ospf_neighbors": {
			Description: "Total number of ospf neighbors per switch in DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_pim_neighbors": {
			Description: "Total number of pim neighbors per switch in DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_telnet": {
			Description: "Returns the value of the nxosTelnet field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nxos_total_routes": {
			Description: "Total number of routes configured in the DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_vtp": {
			Description: "Returns the value of the nxosVtpStatus field.",
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
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"oper_mode": {
						Description: "Returns the status of operational mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"pruning_mode": {
						Description: "Returns the status pruning mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"running_version": {
						Description: "Returns the running version of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"trap_enabled": {
						Description: "Returns the status of trap in vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"v2_mode": {
						Description: "Returns the status of v2 mode of vtp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Returns version of vtp running.",
						Type:        schema.TypeInt,
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
		"route_prefix_count": {
			Description: "Total nuumber of v4 and v6 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"route_prefix_v4_count": {
			Description: "Number of v4 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"route_prefix_v6_count": {
			Description: "Number of v6 routes per node.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"serial": {
			Description: "Serial number of device being invetoried. The serial number is unique per device and will be used as the key.",
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
		"smart_account_id": {
			Description: "Returns the value of the smartAccountId/CustomerId field.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"software_download": {
			Description: "Last software downloaded of device being inventoried. This determines if software download API was used.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"system_up_time": {
			Description: "The amount of time that the device being inventoried been up.",
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
		"nr_version": {
			Description: "Software version of device being inventoried. The various software version values for each device are available on cisco.com.",
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
		"vni_status": {
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
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni": {
						Description: "Returns the vni id of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni_state": {
						Description: "Returns the vni state of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vni_type": {
						Description: "Returns the vni type of the vni.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryRead,
		Schema:      model}
}

func dataSourceNiatelemetryNiaInventoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventory{}
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

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("controller_ip_address"); ok {
		x := (v.(string))
		o.SetControllerIpAddress(x)
	}

	if v, ok := d.GetOk("cpu"); ok {
		x := float32(v.(float64))
		o.SetCpu(x)
	}

	if v, ok := d.GetOk("crash_reset_logs"); ok {
		x := (v.(string))
		o.SetCrashResetLogs(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("customer_device_connector"); ok {
		x := (v.(string))
		o.SetCustomerDeviceConnector(x)
	}

	if v, ok := d.GetOk("dcnm_license_state"); ok {
		x := (v.(string))
		o.SetDcnmLicenseState(x)
	}

	if v, ok := d.GetOk("device_discovery"); ok {
		x := (v.(string))
		o.SetDeviceDiscovery(x)
	}

	if v, ok := d.GetOkExists("device_health"); ok {
		x := int64(v.(int))
		o.SetDeviceHealth(x)
	}

	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}

	if v, ok := d.GetOk("device_name"); ok {
		x := (v.(string))
		o.SetDeviceName(x)
	}

	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
	}

	if v, ok := d.GetOkExists("device_up_time"); ok {
		x := int64(v.(int))
		o.SetDeviceUpTime(x)
	}

	if v, ok := d.GetOk("disk"); ok {
		p := make([]models.NiatelemetryDiskinfo, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryDiskinfo{}
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
			o.SetClassId("")
			if v, ok := l["free"]; ok {
				{
					x := int64(v.(int))
					o.SetFree(x)
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
			if v, ok := l["total"]; ok {
				{
					x := int64(v.(int))
					o.SetTotal(x)
				}
			}
			if v, ok := l["used"]; ok {
				{
					x := int64(v.(int))
					o.SetUsed(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetDisk(x)
		}
	}

	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("fabric_name"); ok {
		x := (v.(string))
		o.SetFabricName(x)
	}

	if v, ok := d.GetOkExists("fex_count"); ok {
		x := int64(v.(int))
		o.SetFexCount(x)
	}

	if v, ok := d.GetOkExists("infra_wi_node_count"); ok {
		x := int64(v.(int))
		o.SetInfraWiNodeCount(x)
	}

	if v, ok := d.GetOk("interface"); ok {
		x := make([]models.NiatelemetryInterfaceElement, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryInterfaceElement{}
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
			o.SetClassId("niatelemetry.InterfaceElement")
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
			if v, ok := l["oper_state"]; ok {
				{
					x := (v.(string))
					o.SetOperState(x)
				}
			}
			if v, ok := l["xcvr_present"]; ok {
				{
					x := (v.(string))
					o.SetXcvrPresent(x)
				}
			}
			x = append(x, *o)
		}
		o.SetInterface(x)
	}

	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}

	if v, ok := d.GetOk("is_virtual_node"); ok {
		x := (v.(string))
		o.SetIsVirtualNode(x)
	}

	if v, ok := d.GetOk("last_reboot_time"); ok {
		x := (v.(string))
		o.SetLastRebootTime(x)
	}

	if v, ok := d.GetOk("last_reset_reason"); ok {
		x := (v.(string))
		o.SetLastResetReason(x)
	}

	if v, ok := d.GetOk("license_state"); ok {
		p := make([]models.NiatelemetryNiaLicenseStateRelationship, 0, 1)
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
			o.SetClassId("")
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
			p = append(p, models.MoMoRefAsNiatelemetryNiaLicenseStateRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetLicenseState(x)
		}
	}

	if v, ok := d.GetOk("license_type"); ok {
		x := (v.(string))
		o.SetLicenseType(x)
	}

	if v, ok := d.GetOk("log_in_time"); ok {
		x := (v.(string))
		o.SetLogInTime(x)
	}

	if v, ok := d.GetOk("log_out_time"); ok {
		x := (v.(string))
		o.SetLogOutTime(x)
	}

	if v, ok := d.GetOkExists("mac_sec_count"); ok {
		x := int64(v.(int))
		o.SetMacSecCount(x)
	}

	if v, ok := d.GetOkExists("mac_sec_fab_count"); ok {
		x := int64(v.(int))
		o.SetMacSecFabCount(x)
	}

	if v, ok := d.GetOkExists("macsec_total_count"); ok {
		x := int64(v.(int))
		o.SetMacsecTotalCount(x)
	}

	if v, ok := d.GetOkExists("memory"); ok {
		x := int64(v.(int))
		o.SetMemory(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("node_id"); ok {
		x := (v.(string))
		o.SetNodeId(x)
	}

	if v, ok := d.GetOk("nxos_bgp_evpn"); ok {
		p := make([]models.NiatelemetryNxosBgpEvpn, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNxosBgpEvpn{}
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
			o.SetClassId("")
			if v, ok := l["nxos_evpn_mac_count"]; ok {
				{
					x := (v.(string))
					o.SetNxosEvpnMacCount(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["total_networks"]; ok {
				{
					x := int64(v.(int))
					o.SetTotalNetworks(x)
				}
			}
			if v, ok := l["total_paths"]; ok {
				{
					x := int64(v.(int))
					o.SetTotalPaths(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNxosBgpEvpn(x)
		}
	}

	if v, ok := d.GetOk("nxos_bgp_mvpn"); ok {
		p := make([]models.NiatelemetryNxosBgpMvpn, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNxosBgpMvpn{}
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
			if v, ok := l["capable_peers"]; ok {
				{
					x := int64(v.(int))
					o.SetCapablePeers(x)
				}
			}
			o.SetClassId("")
			if v, ok := l["configured_peers"]; ok {
				{
					x := int64(v.(int))
					o.SetConfiguredPeers(x)
				}
			}
			if v, ok := l["memory_used"]; ok {
				{
					x := int64(v.(int))
					o.SetMemoryUsed(x)
				}
			}
			if v, ok := l["number_of_cluster_lists"]; ok {
				{
					x := int64(v.(int))
					o.SetNumberOfClusterLists(x)
				}
			}
			if v, ok := l["number_of_communities"]; ok {
				{
					x := int64(v.(int))
					o.SetNumberOfCommunities(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["table_version"]; ok {
				{
					x := int64(v.(int))
					o.SetTableVersion(x)
				}
			}
			if v, ok := l["total_networks"]; ok {
				{
					x := int64(v.(int))
					o.SetTotalNetworks(x)
				}
			}
			if v, ok := l["total_paths"]; ok {
				{
					x := int64(v.(int))
					o.SetTotalPaths(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNxosBgpMvpn(x)
		}
	}

	if v, ok := d.GetOk("nxos_bootflash_details"); ok {
		p := make([]models.NiatelemetryBootflashDetails, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryBootflashDetails{}
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
			o.SetClassId("")
			if v, ok := l["fw_rev"]; ok {
				{
					x := (v.(string))
					o.SetFwRev(x)
				}
			}
			if v, ok := l["model_type"]; ok {
				{
					x := (v.(string))
					o.SetModelType(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["serial"]; ok {
				{
					x := (v.(string))
					o.SetSerial(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNxosBootflashDetails(x)
		}
	}

	if v, ok := d.GetOk("nxos_dci_interface_status"); ok {
		x := (v.(string))
		o.SetNxosDciInterfaceStatus(x)
	}

	if v, ok := d.GetOk("nxos_interface_brief"); ok {
		p := make([]models.NiatelemetryInterface, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryInterface{}
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
			o.SetClassId("")
			if v, ok := l["interface_down_count"]; ok {
				{
					x := int64(v.(int))
					o.SetInterfaceDownCount(x)
				}
			}
			if v, ok := l["interface_up_count"]; ok {
				{
					x := int64(v.(int))
					o.SetInterfaceUpCount(x)
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
			o.SetNxosInterfaceBrief(x)
		}
	}

	if v, ok := d.GetOk("nxos_nve_interface_status"); ok {
		x := (v.(string))
		o.SetNxosNveInterfaceStatus(x)
	}

	if v, ok := d.GetOk("nxos_nve_packet_counters"); ok {
		p := make([]models.NiatelemetryNvePacketCounters, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNvePacketCounters{}
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
			o.SetClassId("")
			if v, ok := l["mcast_inpkts"]; ok {
				{
					x := int64(v.(int))
					o.SetMcastInpkts(x)
				}
			}
			if v, ok := l["mcast_outbytes"]; ok {
				{
					x := int64(v.(int))
					o.SetMcastOutbytes(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["ucast_inpkts"]; ok {
				{
					x := int64(v.(int))
					o.SetUcastInpkts(x)
				}
			}
			if v, ok := l["ucast_outpkts"]; ok {
				{
					x := int64(v.(int))
					o.SetUcastOutpkts(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNxosNvePacketCounters(x)
		}
	}

	if v, ok := d.GetOk("nxos_nve_vni"); ok {
		p := make([]models.NiatelemetryNveVni, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNveVni{}
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
			o.SetClassId("")
			if v, ok := l["cp_vni_count"]; ok {
				{
					x := int64(v.(int))
					o.SetCpVniCount(x)
				}
			}
			if v, ok := l["cp_vni_down"]; ok {
				{
					x := int64(v.(int))
					o.SetCpVniDown(x)
				}
			}
			if v, ok := l["cp_vni_up"]; ok {
				{
					x := int64(v.(int))
					o.SetCpVniUp(x)
				}
			}
			if v, ok := l["dp_vni_count"]; ok {
				{
					x := int64(v.(int))
					o.SetDpVniCount(x)
				}
			}
			if v, ok := l["dp_vni_down"]; ok {
				{
					x := int64(v.(int))
					o.SetDpVniDown(x)
				}
			}
			if v, ok := l["dp_vni_up"]; ok {
				{
					x := int64(v.(int))
					o.SetDpVniUp(x)
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
			o.SetNxosNveVni(x)
		}
	}

	if v, ok := d.GetOkExists("nxos_ospf_neighbors"); ok {
		x := int64(v.(int))
		o.SetNxosOspfNeighbors(x)
	}

	if v, ok := d.GetOk("nxos_pim_neighbors"); ok {
		x := (v.(string))
		o.SetNxosPimNeighbors(x)
	}

	if v, ok := d.GetOk("nxos_telnet"); ok {
		x := (v.(string))
		o.SetNxosTelnet(x)
	}

	if v, ok := d.GetOkExists("nxos_total_routes"); ok {
		x := int64(v.(int))
		o.SetNxosTotalRoutes(x)
	}

	if v, ok := d.GetOk("nxos_vtp"); ok {
		p := make([]models.NiatelemetryNxosVtp, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNxosVtp{}
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
			o.SetClassId("")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["oper_mode"]; ok {
				{
					x := (v.(string))
					o.SetOperMode(x)
				}
			}
			if v, ok := l["pruning_mode"]; ok {
				{
					x := (v.(string))
					o.SetPruningMode(x)
				}
			}
			if v, ok := l["running_version"]; ok {
				{
					x := (v.(string))
					o.SetRunningVersion(x)
				}
			}
			if v, ok := l["trap_enabled"]; ok {
				{
					x := (v.(string))
					o.SetTrapEnabled(x)
				}
			}
			if v, ok := l["v2_mode"]; ok {
				{
					x := (v.(string))
					o.SetV2Mode(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNxosVtp(x)
		}
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
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
			o.SetClassId("")
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
			o.SetClassId("")
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

	if v, ok := d.GetOkExists("route_prefix_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixCount(x)
	}

	if v, ok := d.GetOkExists("route_prefix_v4_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV4Count(x)
	}

	if v, ok := d.GetOkExists("route_prefix_v6_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV6Count(x)
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

	if v, ok := d.GetOkExists("smart_account_id"); ok {
		x := int64(v.(int))
		o.SetSmartAccountId(x)
	}

	if v, ok := d.GetOk("software_download"); ok {
		x := (v.(string))
		o.SetSoftwareDownload(x)
	}

	if v, ok := d.GetOk("system_up_time"); ok {
		x := (v.(string))
		o.SetSystemUpTime(x)
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

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
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
			o.SetClassId("")
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

	if v, ok := d.GetOk("vni_status"); ok {
		x := make([]models.NiatelemetryVniStatus, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryVniStatus{}
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
			o.SetClassId("niatelemetry.VniStatus")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["vni"]; ok {
				{
					x := (v.(string))
					o.SetVni(x)
				}
			}
			if v, ok := l["vni_state"]; ok {
				{
					x := (v.(string))
					o.SetVniState(x)
				}
			}
			if v, ok := l["vni_type"]; ok {
				{
					x := (v.(string))
					o.SetVniType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetVniStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventory object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventory: %s", responseErr.Error())
	}
	count := countResponse.NiatelemetryNiaInventoryList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for NiatelemetryNiaInventory data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var niatelemetryNiaInventoryResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching NiatelemetryNiaInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventory: %s", responseErr.Error())
		}
		results := resMo.NiatelemetryNiaInventoryList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())
				temp["controller_ip_address"] = (s.GetControllerIpAddress())
				temp["cpu"] = (s.GetCpu())
				temp["crash_reset_logs"] = (s.GetCrashResetLogs())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["customer_device_connector"] = (s.GetCustomerDeviceConnector())
				temp["dcnm_license_state"] = (s.GetDcnmLicenseState())
				temp["device_discovery"] = (s.GetDeviceDiscovery())
				temp["device_health"] = (s.GetDeviceHealth())
				temp["device_id"] = (s.GetDeviceId())
				temp["device_name"] = (s.GetDeviceName())
				temp["device_type"] = (s.GetDeviceType())
				temp["device_up_time"] = (s.GetDeviceUpTime())

				temp["disk"] = flattenMapNiatelemetryDiskinfo(s.GetDisk(), d)
				temp["dn"] = (s.GetDn())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["fabric_name"] = (s.GetFabricName())
				temp["fex_count"] = (s.GetFexCount())
				temp["infra_wi_node_count"] = (s.GetInfraWiNodeCount())

				temp["interface"] = flattenListNiatelemetryInterfaceElement(s.GetInterface(), d)
				temp["ip_address"] = (s.GetIpAddress())
				temp["is_virtual_node"] = (s.GetIsVirtualNode())
				temp["last_reboot_time"] = (s.GetLastRebootTime())
				temp["last_reset_reason"] = (s.GetLastResetReason())

				temp["license_state"] = flattenMapNiatelemetryNiaLicenseStateRelationship(s.GetLicenseState(), d)
				temp["license_type"] = (s.GetLicenseType())
				temp["log_in_time"] = (s.GetLogInTime())
				temp["log_out_time"] = (s.GetLogOutTime())
				temp["mac_sec_count"] = (s.GetMacSecCount())
				temp["mac_sec_fab_count"] = (s.GetMacSecFabCount())
				temp["macsec_total_count"] = (s.GetMacsecTotalCount())
				temp["memory"] = (s.GetMemory())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["node_id"] = (s.GetNodeId())

				temp["nxos_bgp_evpn"] = flattenMapNiatelemetryNxosBgpEvpn(s.GetNxosBgpEvpn(), d)

				temp["nxos_bgp_mvpn"] = flattenMapNiatelemetryNxosBgpMvpn(s.GetNxosBgpMvpn(), d)

				temp["nxos_bootflash_details"] = flattenMapNiatelemetryBootflashDetails(s.GetNxosBootflashDetails(), d)
				temp["nxos_dci_interface_status"] = (s.GetNxosDciInterfaceStatus())

				temp["nxos_interface_brief"] = flattenMapNiatelemetryInterface(s.GetNxosInterfaceBrief(), d)
				temp["nxos_nve_interface_status"] = (s.GetNxosNveInterfaceStatus())

				temp["nxos_nve_packet_counters"] = flattenMapNiatelemetryNvePacketCounters(s.GetNxosNvePacketCounters(), d)

				temp["nxos_nve_vni"] = flattenMapNiatelemetryNveVni(s.GetNxosNveVni(), d)
				temp["nxos_ospf_neighbors"] = (s.GetNxosOspfNeighbors())
				temp["nxos_pim_neighbors"] = (s.GetNxosPimNeighbors())
				temp["nxos_telnet"] = (s.GetNxosTelnet())
				temp["nxos_total_routes"] = (s.GetNxosTotalRoutes())

				temp["nxos_vtp"] = flattenMapNiatelemetryNxosVtp(s.GetNxosVtp(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["route_prefix_count"] = (s.GetRoutePrefixCount())
				temp["route_prefix_v4_count"] = (s.GetRoutePrefixV4Count())
				temp["route_prefix_v6_count"] = (s.GetRoutePrefixV6Count())
				temp["serial"] = (s.GetSerial())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["site_name"] = (s.GetSiteName())
				temp["smart_account_id"] = (s.GetSmartAccountId())
				temp["software_download"] = (s.GetSoftwareDownload())
				temp["system_up_time"] = (s.GetSystemUpTime())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["vni_status"] = flattenListNiatelemetryVniStatus(s.GetVniStatus(), d)
				niatelemetryNiaInventoryResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaInventoryResults))
	if err := d.Set("results", niatelemetryNiaInventoryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaInventoryResults[0]["moid"].(string))
	return de
}
