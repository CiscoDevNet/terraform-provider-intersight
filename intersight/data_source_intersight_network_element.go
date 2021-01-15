package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkElement() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkElementRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"admin_evac_state": {
				Description: "Administratively configured state of Fabric Evacuation feature, for this switch.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"admin_inband_interface_state": {
				Description: "The administrative state of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"alarm_summary": {
				Description: "The summary of alarm counts based on the severity types (Critical or Warning).",
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
						"critical": {
							Description: "The count of alarms that have severity type Critical.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"warning": {
							Description: "The count of alarms that have severity type Warning.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"available_memory": {
				Description: "Available memory (un-used) on this switch platform.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cards": {
				Description: "An array of relationships to equipmentSwitchCard resources.",
				Type:        schema.TypeList,
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ethernet_mode": {
				Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ethernet_switching_mode": {
				Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fanmodules": {
				Description: "An array of relationships to equipmentFanModule resources.",
				Type:        schema.TypeList,
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
			"fault_summary": {
				Description: "The fault summary of the network Element out-of-band management interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fc_mode": {
				Description: "The user configured FC operational mode for this switch (End-Host or Switching).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fc_switching_mode": {
				Description: "The user configured FC operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_address": {
				Description: "The IP address of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_gateway": {
				Description: "The default gateway of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_ip_mask": {
				Description: "The network mask of the network Element inband management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"inband_vlan": {
				Description: "The VLAN ID of the network Element inband management interface.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"inventory_device_info": {
				Description: "A reference to a inventoryDeviceInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"management_controller": {
				Description: "A reference to a managementController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"management_entity": {
				Description: "A reference to a managementEntity resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"management_mode": {
				Description: "The management mode of the fabric interconnect.\n* `IntersightStandalone` - Intersight Standalone mode of operation.\n* `UCSM` - Unified Computing System Manager mode of operation.\n* `Intersight` - Intersight managed mode of operation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "This field identifies the model of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"network_fc_zone_info": {
				Description: "A reference to a networkFcZoneInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"network_vlan_port_info": {
				Description: "A reference to a networkVlanPortInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_evac_state": {
				Description: "Operational state of the Fabric Evacuation feature, for this switch.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "The switch's current overall operational/health state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_address": {
				Description: "The IP address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_gateway": {
				Description: "The default gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ip_mask": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_address": {
				Description: "The IPv4 address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_gateway": {
				Description: "The default IPv4 gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv4_mask": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"out_of_band_ipv6_address": {
				Description: "The IPv6 address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_ipv6_gateway": {
				Description: "The default IPv6 gateway of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_ipv6_prefix": {
				Description: "The network mask of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_mac": {
				Description: "The MAC address of the network Element out-of-band management interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port_mac_bindings": {
				Description: "An array of relationships to portMacBinding resources.",
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
			"psus": {
				Description: "An array of relationships to equipmentPsu resources.",
				Type:        schema.TypeList,
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
			"revision": {
				Description: "This field identifies the revision of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_items": {
				Description: "An array of relationships to storageItem resources.",
				Type:        schema.TypeList,
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
			"switch_id": {
				Description: "The Switch Id of the network Element.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"top_system": {
				Description: "A reference to a topSystem resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"total_memory": {
				Description: "Total available memory on this switch platform.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"ucsm_running_firmware": {
				Description: "A reference to a firmwareRunningFirmware resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceNetworkElementRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkElement{}
	if v, ok := d.GetOk("admin_evac_state"); ok {
		x := (v.(string))
		o.SetAdminEvacState(x)
	}
	if v, ok := d.GetOk("admin_inband_interface_state"); ok {
		x := (v.(string))
		o.SetAdminInbandInterfaceState(x)
	}
	if v, ok := d.GetOk("available_memory"); ok {
		x := (v.(string))
		o.SetAvailableMemory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("ethernet_mode"); ok {
		x := (v.(string))
		o.SetEthernetMode(x)
	}
	if v, ok := d.GetOk("ethernet_switching_mode"); ok {
		x := (v.(string))
		o.SetEthernetSwitchingMode(x)
	}
	if v, ok := d.GetOk("fault_summary"); ok {
		x := int64(v.(int))
		o.SetFaultSummary(x)
	}
	if v, ok := d.GetOk("fc_mode"); ok {
		x := (v.(string))
		o.SetFcMode(x)
	}
	if v, ok := d.GetOk("fc_switching_mode"); ok {
		x := (v.(string))
		o.SetFcSwitchingMode(x)
	}
	if v, ok := d.GetOk("inband_ip_address"); ok {
		x := (v.(string))
		o.SetInbandIpAddress(x)
	}
	if v, ok := d.GetOk("inband_ip_gateway"); ok {
		x := (v.(string))
		o.SetInbandIpGateway(x)
	}
	if v, ok := d.GetOk("inband_ip_mask"); ok {
		x := (v.(string))
		o.SetInbandIpMask(x)
	}
	if v, ok := d.GetOk("inband_vlan"); ok {
		x := int64(v.(int))
		o.SetInbandVlan(x)
	}
	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_evac_state"); ok {
		x := (v.(string))
		o.SetOperEvacState(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpAddress(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpGateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ip_mask"); ok {
		x := (v.(string))
		o.SetOutOfBandIpMask(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Address(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Gateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv4_mask"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv4Mask(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_address"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Address(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_gateway"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Gateway(x)
	}
	if v, ok := d.GetOk("out_of_band_ipv6_prefix"); ok {
		x := (v.(string))
		o.SetOutOfBandIpv6Prefix(x)
	}
	if v, ok := d.GetOk("out_of_band_mac"); ok {
		x := (v.(string))
		o.SetOutOfBandMac(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SetSwitchId(x)
	}
	if v, ok := d.GetOk("total_memory"); ok {
		x := int64(v.(int))
		o.SetTotalMemory(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NetworkElement object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NetworkApi.GetNetworkElementList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching NetworkElement: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NetworkElement list: %s", err.Error())
	}
	var s = &models.NetworkElementList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NetworkElement list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NetworkElement data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NetworkElement data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NetworkElement{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("admin_evac_state", (s.GetAdminEvacState())); err != nil {
				return diag.Errorf("error occurred while setting property AdminEvacState: %s", err.Error())
			}
			if err := d.Set("admin_inband_interface_state", (s.GetAdminInbandInterfaceState())); err != nil {
				return diag.Errorf("error occurred while setting property AdminInbandInterfaceState: %s", err.Error())
			}

			if err := d.Set("alarm_summary", flattenMapComputeAlarmSummary(s.GetAlarmSummary(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AlarmSummary: %s", err.Error())
			}
			if err := d.Set("available_memory", (s.GetAvailableMemory())); err != nil {
				return diag.Errorf("error occurred while setting property AvailableMemory: %s", err.Error())
			}

			if err := d.Set("cards", flattenListEquipmentSwitchCardRelationship(s.GetCards(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cards: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceMoId: %s", err.Error())
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return diag.Errorf("error occurred while setting property Dn: %s", err.Error())
			}
			if err := d.Set("ethernet_mode", (s.GetEthernetMode())); err != nil {
				return diag.Errorf("error occurred while setting property EthernetMode: %s", err.Error())
			}
			if err := d.Set("ethernet_switching_mode", (s.GetEthernetSwitchingMode())); err != nil {
				return diag.Errorf("error occurred while setting property EthernetSwitchingMode: %s", err.Error())
			}

			if err := d.Set("fanmodules", flattenListEquipmentFanModuleRelationship(s.GetFanmodules(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Fanmodules: %s", err.Error())
			}
			if err := d.Set("fault_summary", (s.GetFaultSummary())); err != nil {
				return diag.Errorf("error occurred while setting property FaultSummary: %s", err.Error())
			}
			if err := d.Set("fc_mode", (s.GetFcMode())); err != nil {
				return diag.Errorf("error occurred while setting property FcMode: %s", err.Error())
			}
			if err := d.Set("fc_switching_mode", (s.GetFcSwitchingMode())); err != nil {
				return diag.Errorf("error occurred while setting property FcSwitchingMode: %s", err.Error())
			}
			if err := d.Set("inband_ip_address", (s.GetInbandIpAddress())); err != nil {
				return diag.Errorf("error occurred while setting property InbandIpAddress: %s", err.Error())
			}
			if err := d.Set("inband_ip_gateway", (s.GetInbandIpGateway())); err != nil {
				return diag.Errorf("error occurred while setting property InbandIpGateway: %s", err.Error())
			}
			if err := d.Set("inband_ip_mask", (s.GetInbandIpMask())); err != nil {
				return diag.Errorf("error occurred while setting property InbandIpMask: %s", err.Error())
			}
			if err := d.Set("inband_vlan", (s.GetInbandVlan())); err != nil {
				return diag.Errorf("error occurred while setting property InbandVlan: %s", err.Error())
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InventoryDeviceInfo: %s", err.Error())
			}

			if err := d.Set("management_controller", flattenMapManagementControllerRelationship(s.GetManagementController(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ManagementController: %s", err.Error())
			}

			if err := d.Set("management_entity", flattenMapManagementEntityRelationship(s.GetManagementEntity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ManagementEntity: %s", err.Error())
			}
			if err := d.Set("management_mode", (s.GetManagementMode())); err != nil {
				return diag.Errorf("error occurred while setting property ManagementMode: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}

			if err := d.Set("network_fc_zone_info", flattenMapNetworkFcZoneInfoRelationship(s.GetNetworkFcZoneInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NetworkFcZoneInfo: %s", err.Error())
			}

			if err := d.Set("network_vlan_port_info", flattenMapNetworkVlanPortInfoRelationship(s.GetNetworkVlanPortInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NetworkVlanPortInfo: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("oper_evac_state", (s.GetOperEvacState())); err != nil {
				return diag.Errorf("error occurred while setting property OperEvacState: %s", err.Error())
			}
			if err := d.Set("operability", (s.GetOperability())); err != nil {
				return diag.Errorf("error occurred while setting property Operability: %s", err.Error())
			}
			if err := d.Set("out_of_band_ip_address", (s.GetOutOfBandIpAddress())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpAddress: %s", err.Error())
			}
			if err := d.Set("out_of_band_ip_gateway", (s.GetOutOfBandIpGateway())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpGateway: %s", err.Error())
			}
			if err := d.Set("out_of_band_ip_mask", (s.GetOutOfBandIpMask())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpMask: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv4_address", (s.GetOutOfBandIpv4Address())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv4Address: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv4_gateway", (s.GetOutOfBandIpv4Gateway())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv4Gateway: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv4_mask", (s.GetOutOfBandIpv4Mask())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv4Mask: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv6_address", (s.GetOutOfBandIpv6Address())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv6Address: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv6_gateway", (s.GetOutOfBandIpv6Gateway())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv6Gateway: %s", err.Error())
			}
			if err := d.Set("out_of_band_ipv6_prefix", (s.GetOutOfBandIpv6Prefix())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandIpv6Prefix: %s", err.Error())
			}
			if err := d.Set("out_of_band_mac", (s.GetOutOfBandMac())); err != nil {
				return diag.Errorf("error occurred while setting property OutOfBandMac: %s", err.Error())
			}

			if err := d.Set("port_mac_bindings", flattenListPortMacBindingRelationship(s.GetPortMacBindings(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortMacBindings: %s", err.Error())
			}

			if err := d.Set("psus", flattenListEquipmentPsuRelationship(s.GetPsus(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Psus: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return diag.Errorf("error occurred while setting property Revision: %s", err.Error())
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return diag.Errorf("error occurred while setting property Rn: %s", err.Error())
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return diag.Errorf("error occurred while setting property Serial: %s", err.Error())
			}

			if err := d.Set("storage_items", flattenListStorageItemRelationship(s.GetStorageItems(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageItems: %s", err.Error())
			}
			if err := d.Set("switch_id", (s.GetSwitchId())); err != nil {
				return diag.Errorf("error occurred while setting property SwitchId: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("top_system", flattenMapTopSystemRelationship(s.GetTopSystem(), d)); err != nil {
				return diag.Errorf("error occurred while setting property TopSystem: %s", err.Error())
			}
			if err := d.Set("total_memory", (s.GetTotalMemory())); err != nil {
				return diag.Errorf("error occurred while setting property TotalMemory: %s", err.Error())
			}

			if err := d.Set("ucsm_running_firmware", flattenMapFirmwareRunningFirmwareRelationship(s.GetUcsmRunningFirmware(), d)); err != nil {
				return diag.Errorf("error occurred while setting property UcsmRunningFirmware: %s", err.Error())
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return diag.Errorf("error occurred while setting property Vendor: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
