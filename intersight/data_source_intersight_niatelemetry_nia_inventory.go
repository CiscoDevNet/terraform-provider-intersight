package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryNiaInventory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"customer_device_connector": {
				Description: "Returns the value of the customerDeviceConnector field.",
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
			"dn": {
				Description: "Dn for the inventories present.",
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"node_id": {
				Description: "The ID of the device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nxos_dci_interface_status": {
				Description: "Returns the status of dci interface configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nxos_nve_interface_status": {
				Description: "Returns the value of the nxosNveInterface field.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"site_name": {
				Description: "Name of fabric domain of the controller.",
				Type:        schema.TypeString,
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
			"nr_version": {
				Description: "Software version of device being inventoried. The various software version values for each device are available on cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					"customer_device_connector": {
						Description: "Returns the value of the customerDeviceConnector field.",
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
									Computed:    true,
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
						Computed: true,
					},
					"dn": {
						Description: "Dn for the inventories present.",
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
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
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
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"node_id": {
						Description: "The ID of the device being inventoried.",
						Type:        schema.TypeString,
						Optional:    true,
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
									Computed:    true,
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
						Computed: true,
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
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"nxos_nve_interface_status": {
						Description: "Returns the value of the nxosNveInterface field.",
						Type:        schema.TypeString,
						Optional:    true,
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
									Computed:    true,
								},
							},
						},
						Computed: true,
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
									Computed:    true,
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
						Computed: true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
						Computed:    true,
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
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
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
					"site_name": {
						Description: "Name of fabric domain of the controller.",
						Type:        schema.TypeString,
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryNiaInventoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventory{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cpu"); ok {
		x := v.(float32)
		o.SetCpu(x)
	}
	if v, ok := d.GetOk("crash_reset_logs"); ok {
		x := (v.(string))
		o.SetCrashResetLogs(x)
	}
	if v, ok := d.GetOk("customer_device_connector"); ok {
		x := (v.(string))
		o.SetCustomerDeviceConnector(x)
	}
	if v, ok := d.GetOk("device_discovery"); ok {
		x := (v.(string))
		o.SetDeviceDiscovery(x)
	}
	if v, ok := d.GetOk("device_health"); ok {
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
	if v, ok := d.GetOk("device_up_time"); ok {
		x := int64(v.(int))
		o.SetDeviceUpTime(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fex_count"); ok {
		x := int64(v.(int))
		o.SetFexCount(x)
	}
	if v, ok := d.GetOk("infra_wi_node_count"); ok {
		x := int64(v.(int))
		o.SetInfraWiNodeCount(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("is_virtual_node"); ok {
		x := (v.(string))
		o.SetIsVirtualNode(x)
	}
	if v, ok := d.GetOk("log_in_time"); ok {
		x := (v.(string))
		o.SetLogInTime(x)
	}
	if v, ok := d.GetOk("log_out_time"); ok {
		x := (v.(string))
		o.SetLogOutTime(x)
	}
	if v, ok := d.GetOk("mac_sec_count"); ok {
		x := int64(v.(int))
		o.SetMacSecCount(x)
	}
	if v, ok := d.GetOk("mac_sec_fab_count"); ok {
		x := int64(v.(int))
		o.SetMacSecFabCount(x)
	}
	if v, ok := d.GetOk("macsec_total_count"); ok {
		x := int64(v.(int))
		o.SetMacsecTotalCount(x)
	}
	if v, ok := d.GetOk("memory"); ok {
		x := int64(v.(int))
		o.SetMemory(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("node_id"); ok {
		x := (v.(string))
		o.SetNodeId(x)
	}
	if v, ok := d.GetOk("nxos_dci_interface_status"); ok {
		x := (v.(string))
		o.SetNxosDciInterfaceStatus(x)
	}
	if v, ok := d.GetOk("nxos_nve_interface_status"); ok {
		x := (v.(string))
		o.SetNxosNveInterfaceStatus(x)
	}
	if v, ok := d.GetOk("nxos_ospf_neighbors"); ok {
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
	if v, ok := d.GetOk("nxos_total_routes"); ok {
		x := int64(v.(int))
		o.SetNxosTotalRoutes(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("route_prefix_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixCount(x)
	}
	if v, ok := d.GetOk("route_prefix_v4_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV4Count(x)
	}
	if v, ok := d.GetOk("route_prefix_v6_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV6Count(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("software_download"); ok {
		x := (v.(string))
		o.SetSoftwareDownload(x)
	}
	if v, ok := d.GetOk("system_up_time"); ok {
		x := (v.(string))
		o.SetSystemUpTime(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventory object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryNiaInventoryList.GetCount()
	var i int32
	var niatelemetryNiaInventoryResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryNiaInventoryList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryNiaInventory data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["cpu"] = (s.GetCpu())
				temp["crash_reset_logs"] = (s.GetCrashResetLogs())
				temp["customer_device_connector"] = (s.GetCustomerDeviceConnector())
				temp["device_discovery"] = (s.GetDeviceDiscovery())
				temp["device_health"] = (s.GetDeviceHealth())
				temp["device_id"] = (s.GetDeviceId())
				temp["device_name"] = (s.GetDeviceName())
				temp["device_type"] = (s.GetDeviceType())
				temp["device_up_time"] = (s.GetDeviceUpTime())

				temp["disk"] = flattenMapNiatelemetryDiskinfo(s.GetDisk(), d)
				temp["dn"] = (s.GetDn())
				temp["fex_count"] = (s.GetFexCount())
				temp["infra_wi_node_count"] = (s.GetInfraWiNodeCount())
				temp["ip_address"] = (s.GetIpAddress())
				temp["is_virtual_node"] = (s.GetIsVirtualNode())

				temp["license_state"] = flattenMapNiatelemetryNiaLicenseStateRelationship(s.GetLicenseState(), d)
				temp["log_in_time"] = (s.GetLogInTime())
				temp["log_out_time"] = (s.GetLogOutTime())
				temp["mac_sec_count"] = (s.GetMacSecCount())
				temp["mac_sec_fab_count"] = (s.GetMacSecFabCount())
				temp["macsec_total_count"] = (s.GetMacsecTotalCount())
				temp["memory"] = (s.GetMemory())
				temp["moid"] = (s.GetMoid())
				temp["node_id"] = (s.GetNodeId())

				temp["nxos_bgp_mvpn"] = flattenMapNiatelemetryNxosBgpMvpn(s.GetNxosBgpMvpn(), d)
				temp["nxos_dci_interface_status"] = (s.GetNxosDciInterfaceStatus())

				temp["nxos_interface_brief"] = flattenMapNiatelemetryInterface(s.GetNxosInterfaceBrief(), d)
				temp["nxos_nve_interface_status"] = (s.GetNxosNveInterfaceStatus())

				temp["nxos_nve_vni"] = flattenMapNiatelemetryNveVni(s.GetNxosNveVni(), d)
				temp["nxos_ospf_neighbors"] = (s.GetNxosOspfNeighbors())
				temp["nxos_pim_neighbors"] = (s.GetNxosPimNeighbors())
				temp["nxos_telnet"] = (s.GetNxosTelnet())
				temp["nxos_total_routes"] = (s.GetNxosTotalRoutes())

				temp["nxos_vtp"] = flattenMapNiatelemetryNxosVtp(s.GetNxosVtp(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["route_prefix_count"] = (s.GetRoutePrefixCount())
				temp["route_prefix_v4_count"] = (s.GetRoutePrefixV4Count())
				temp["route_prefix_v6_count"] = (s.GetRoutePrefixV6Count())
				temp["serial"] = (s.GetSerial())
				temp["site_name"] = (s.GetSiteName())
				temp["software_download"] = (s.GetSoftwareDownload())
				temp["system_up_time"] = (s.GetSystemUpTime())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
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
