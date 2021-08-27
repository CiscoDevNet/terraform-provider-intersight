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

func dataSourceNetworkElementSummary() *schema.Resource {
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
		"admin_evac_state": {
			Description: "Administratively configured state of Fabric Evacuation feature, for this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"admin_inband_interface_state": {
			Description: "The administrative state of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"warning": {
						Description: "The count of alarms that have severity type Warning.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
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
		"available_memory": {
			Description: "Available memory (un-used) on this switch platform.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_mo_id": {
			Description: "The database identifier of the registered device of an object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dn": {
			Description: "The Distinguished Name unambiguously identifies an object in the system.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ethernet_mode": {
			Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ethernet_switching_mode": {
			Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
			Type:        schema.TypeString,
			Optional:    true,
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
		},
		"fc_switching_mode": {
			Description: "The user configured FC operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"firmware": {
			Description: "Running firmware information.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_address": {
			Description: "The IP address of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_gateway": {
			Description: "The default gateway of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_mask": {
			Description: "The network mask of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_vlan": {
			Description: "The VLAN ID of the network Element inband management interface.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ipv4_address": {
			Description: "IP version 4 address is saved in this property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"management_mode": {
			Description: "The management mode of the fabric interconnect.\n* `IntersightStandalone` - Intersight Standalone mode of operation.\n* `UCSM` - Unified Computing System Manager mode of operation.\n* `Intersight` - Intersight managed mode of operation.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"model": {
			Description: "This field identifies the model of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "Name of the ElementSummary object is saved in this property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"num_ether_ports": {
			Description: "Total number of Ethernet ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_ether_ports_configured": {
			Description: "Total number of configured Ethernet ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_ether_ports_link_up": {
			Description: "Total number of Ethernet ports which are UP.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_expansion_modules": {
			Description: "Total number of expansion modules.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports": {
			Description: "Total number of FC ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports_configured": {
			Description: "Total number of configured FC ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports_link_up": {
			Description: "Total number of FC ports which are UP.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_evac_state": {
			Description: "Operational state of the Fabric Evacuation feature, for this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"operability": {
			Description: "The switch's current overall operational/health state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_address": {
			Description: "The IP address of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_gateway": {
			Description: "The default gateway of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_mask": {
			Description: "The network mask of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_address": {
			Description: "The IPv4 address of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_gateway": {
			Description: "The default IPv4 gateway of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_mask": {
			Description: "The network mask of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"presence": {
			Description: "This field identifies the presence (equipped) or absence of the given component.",
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
		"revision": {
			Description: "This field identifies the revision of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"rn": {
			Description: "The Relative Name uniquely identifies an object within a given context.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial": {
			Description: "This field identifies the serial of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_object_type": {
			Description: "The source object type of this view MO.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"switch_id": {
			Description: "The Switch Id of the network Element.",
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
		"thermal": {
			Description: "The Thermal status of the fabric interconnect.\n* `unknown` - The default state of the sensor (in case no data is received).\n* `ok` - State of the sensor indicating the sensor's temperature range is okay.\n* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.\n* `upper-critical` - State of the sensor indicating that the temperature is above normal range.\n* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.\n* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.\n* `lower-critical` - State of the sensor indicating that the temperature is below normal range.\n* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"total_memory": {
			Description: "Total available memory on this switch platform.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vendor": {
			Description: "This field identifies the vendor of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nr_version": {
			Description: "Version holds the firmware version related information.",
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
		"admin_evac_state": {
			Description: "Administratively configured state of Fabric Evacuation feature, for this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"admin_inband_interface_state": {
			Description: "The administrative state of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"warning": {
						Description: "The count of alarms that have severity type Warning.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
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
		"available_memory": {
			Description: "Available memory (un-used) on this switch platform.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_mo_id": {
			Description: "The database identifier of the registered device of an object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dn": {
			Description: "The Distinguished Name unambiguously identifies an object in the system.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ethernet_mode": {
			Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ethernet_switching_mode": {
			Description: "The user configured Ethernet operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
			Type:        schema.TypeString,
			Optional:    true,
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
		},
		"fc_switching_mode": {
			Description: "The user configured FC operational mode for this switch (End-Host or Switching).\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"firmware": {
			Description: "Running firmware information.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_address": {
			Description: "The IP address of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_gateway": {
			Description: "The default gateway of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_ip_mask": {
			Description: "The network mask of the network Element inband management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inband_vlan": {
			Description: "The VLAN ID of the network Element inband management interface.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"ipv4_address": {
			Description: "IP version 4 address is saved in this property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"management_mode": {
			Description: "The management mode of the fabric interconnect.\n* `IntersightStandalone` - Intersight Standalone mode of operation.\n* `UCSM` - Unified Computing System Manager mode of operation.\n* `Intersight` - Intersight managed mode of operation.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"model": {
			Description: "This field identifies the model of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "Name of the ElementSummary object is saved in this property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"num_ether_ports": {
			Description: "Total number of Ethernet ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_ether_ports_configured": {
			Description: "Total number of configured Ethernet ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_ether_ports_link_up": {
			Description: "Total number of Ethernet ports which are UP.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_expansion_modules": {
			Description: "Total number of expansion modules.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports": {
			Description: "Total number of FC ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports_configured": {
			Description: "Total number of configured FC ports.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_ports_link_up": {
			Description: "Total number of FC ports which are UP.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_evac_state": {
			Description: "Operational state of the Fabric Evacuation feature, for this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"operability": {
			Description: "The switch's current overall operational/health state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_address": {
			Description: "The IP address of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_gateway": {
			Description: "The default gateway of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ip_mask": {
			Description: "The network mask of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_address": {
			Description: "The IPv4 address of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_gateway": {
			Description: "The default IPv4 gateway of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_ipv4_mask": {
			Description: "The network mask of the network Element out-of-band management interface.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"presence": {
			Description: "This field identifies the presence (equipped) or absence of the given component.",
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
		"revision": {
			Description: "This field identifies the revision of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"rn": {
			Description: "The Relative Name uniquely identifies an object within a given context.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial": {
			Description: "This field identifies the serial of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_object_type": {
			Description: "The source object type of this view MO.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"switch_id": {
			Description: "The Switch Id of the network Element.",
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
		"thermal": {
			Description: "The Thermal status of the fabric interconnect.\n* `unknown` - The default state of the sensor (in case no data is received).\n* `ok` - State of the sensor indicating the sensor's temperature range is okay.\n* `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range.\n* `upper-critical` - State of the sensor indicating that the temperature is above normal range.\n* `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range.\n* `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range.\n* `lower-critical` - State of the sensor indicating that the temperature is below normal range.\n* `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"total_memory": {
			Description: "Total available memory on this switch platform.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vendor": {
			Description: "This field identifies the vendor of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nr_version": {
			Description: "Version holds the firmware version related information.",
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
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceNetworkElementSummaryRead,
		Schema:      model}
}

func dataSourceNetworkElementSummaryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkElementSummary{}
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

	if _, ok := d.GetOk("admin_evac_state"); ok {
		v := d.Get("admin_evac_state")
		x := (v.(string))
		o.SetAdminEvacState(x)
	}

	if _, ok := d.GetOk("admin_inband_interface_state"); ok {
		v := d.Get("admin_inband_interface_state")
		x := (v.(string))
		o.SetAdminInbandInterfaceState(x)
	}

	if _, ok := d.GetOk("alarm_summary"); ok {
		v := d.Get("alarm_summary")
		p := make([]models.ComputeAlarmSummary, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.ComputeAlarmSummary{}
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
			o.SetClassId("compute.AlarmSummary")
			if v, ok := l["critical"]; ok {
				{
					x := int64(v.(int))
					o.SetCritical(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["warning"]; ok {
				{
					x := int64(v.(int))
					o.SetWarning(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAlarmSummary(x)
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

	if _, ok := d.GetOk("available_memory"); ok {
		v := d.Get("available_memory")
		x := (v.(string))
		o.SetAvailableMemory(x)
	}

	if _, ok := d.GetOk("class_id"); ok {
		v := d.Get("class_id")
		x := (v.(string))
		o.SetClassId(x)
	}

	if _, ok := d.GetOk("create_time"); ok {
		v := d.Get("create_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}

	if _, ok := d.GetOk("device_mo_id"); ok {
		v := d.Get("device_mo_id")
		x := (v.(string))
		o.SetDeviceMoId(x)
	}

	if _, ok := d.GetOk("dn"); ok {
		v := d.Get("dn")
		x := (v.(string))
		o.SetDn(x)
	}

	if _, ok := d.GetOk("domain_group_moid"); ok {
		v := d.Get("domain_group_moid")
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if _, ok := d.GetOk("ethernet_mode"); ok {
		v := d.Get("ethernet_mode")
		x := (v.(string))
		o.SetEthernetMode(x)
	}

	if _, ok := d.GetOk("ethernet_switching_mode"); ok {
		v := d.Get("ethernet_switching_mode")
		x := (v.(string))
		o.SetEthernetSwitchingMode(x)
	}

	if _, ok := d.GetOk("fault_summary"); ok {
		v := d.Get("fault_summary")
		x := int64(v.(int))
		o.SetFaultSummary(x)
	}

	if _, ok := d.GetOk("fc_mode"); ok {
		v := d.Get("fc_mode")
		x := (v.(string))
		o.SetFcMode(x)
	}

	if _, ok := d.GetOk("fc_switching_mode"); ok {
		v := d.Get("fc_switching_mode")
		x := (v.(string))
		o.SetFcSwitchingMode(x)
	}

	if _, ok := d.GetOk("firmware"); ok {
		v := d.Get("firmware")
		x := (v.(string))
		o.SetFirmware(x)
	}

	if _, ok := d.GetOk("inband_ip_address"); ok {
		v := d.Get("inband_ip_address")
		x := (v.(string))
		o.SetInbandIpAddress(x)
	}

	if _, ok := d.GetOk("inband_ip_gateway"); ok {
		v := d.Get("inband_ip_gateway")
		x := (v.(string))
		o.SetInbandIpGateway(x)
	}

	if _, ok := d.GetOk("inband_ip_mask"); ok {
		v := d.Get("inband_ip_mask")
		x := (v.(string))
		o.SetInbandIpMask(x)
	}

	if _, ok := d.GetOk("inband_vlan"); ok {
		v := d.Get("inband_vlan")
		x := int64(v.(int))
		o.SetInbandVlan(x)
	}

	if _, ok := d.GetOk("ipv4_address"); ok {
		v := d.Get("ipv4_address")
		x := (v.(string))
		o.SetIpv4Address(x)
	}

	if _, ok := d.GetOk("management_mode"); ok {
		v := d.Get("management_mode")
		x := (v.(string))
		o.SetManagementMode(x)
	}

	if _, ok := d.GetOk("mod_time"); ok {
		v := d.Get("mod_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}

	if _, ok := d.GetOk("model"); ok {
		v := d.Get("model")
		x := (v.(string))
		o.SetModel(x)
	}

	if _, ok := d.GetOk("moid"); ok {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if _, ok := d.GetOk("name"); ok {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if _, ok := d.GetOk("num_ether_ports"); ok {
		v := d.Get("num_ether_ports")
		x := int64(v.(int))
		o.SetNumEtherPorts(x)
	}

	if _, ok := d.GetOk("num_ether_ports_configured"); ok {
		v := d.Get("num_ether_ports_configured")
		x := int64(v.(int))
		o.SetNumEtherPortsConfigured(x)
	}

	if _, ok := d.GetOk("num_ether_ports_link_up"); ok {
		v := d.Get("num_ether_ports_link_up")
		x := int64(v.(int))
		o.SetNumEtherPortsLinkUp(x)
	}

	if _, ok := d.GetOk("num_expansion_modules"); ok {
		v := d.Get("num_expansion_modules")
		x := int64(v.(int))
		o.SetNumExpansionModules(x)
	}

	if _, ok := d.GetOk("num_fc_ports"); ok {
		v := d.Get("num_fc_ports")
		x := int64(v.(int))
		o.SetNumFcPorts(x)
	}

	if _, ok := d.GetOk("num_fc_ports_configured"); ok {
		v := d.Get("num_fc_ports_configured")
		x := int64(v.(int))
		o.SetNumFcPortsConfigured(x)
	}

	if _, ok := d.GetOk("num_fc_ports_link_up"); ok {
		v := d.Get("num_fc_ports_link_up")
		x := int64(v.(int))
		o.SetNumFcPortsLinkUp(x)
	}

	if _, ok := d.GetOk("object_type"); ok {
		v := d.Get("object_type")
		x := (v.(string))
		o.SetObjectType(x)
	}

	if _, ok := d.GetOk("oper_evac_state"); ok {
		v := d.Get("oper_evac_state")
		x := (v.(string))
		o.SetOperEvacState(x)
	}

	if _, ok := d.GetOk("operability"); ok {
		v := d.Get("operability")
		x := (v.(string))
		o.SetOperability(x)
	}

	if _, ok := d.GetOk("out_of_band_ip_address"); ok {
		v := d.Get("out_of_band_ip_address")
		x := (v.(string))
		o.SetOutOfBandIpAddress(x)
	}

	if _, ok := d.GetOk("out_of_band_ip_gateway"); ok {
		v := d.Get("out_of_band_ip_gateway")
		x := (v.(string))
		o.SetOutOfBandIpGateway(x)
	}

	if _, ok := d.GetOk("out_of_band_ip_mask"); ok {
		v := d.Get("out_of_band_ip_mask")
		x := (v.(string))
		o.SetOutOfBandIpMask(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv4_address"); ok {
		v := d.Get("out_of_band_ipv4_address")
		x := (v.(string))
		o.SetOutOfBandIpv4Address(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv4_gateway"); ok {
		v := d.Get("out_of_band_ipv4_gateway")
		x := (v.(string))
		o.SetOutOfBandIpv4Gateway(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv4_mask"); ok {
		v := d.Get("out_of_band_ipv4_mask")
		x := (v.(string))
		o.SetOutOfBandIpv4Mask(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv6_address"); ok {
		v := d.Get("out_of_band_ipv6_address")
		x := (v.(string))
		o.SetOutOfBandIpv6Address(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv6_gateway"); ok {
		v := d.Get("out_of_band_ipv6_gateway")
		x := (v.(string))
		o.SetOutOfBandIpv6Gateway(x)
	}

	if _, ok := d.GetOk("out_of_band_ipv6_prefix"); ok {
		v := d.Get("out_of_band_ipv6_prefix")
		x := (v.(string))
		o.SetOutOfBandIpv6Prefix(x)
	}

	if _, ok := d.GetOk("out_of_band_mac"); ok {
		v := d.Get("out_of_band_mac")
		x := (v.(string))
		o.SetOutOfBandMac(x)
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

	if _, ok := d.GetOk("presence"); ok {
		v := d.Get("presence")
		x := (v.(string))
		o.SetPresence(x)
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

	if _, ok := d.GetOk("revision"); ok {
		v := d.Get("revision")
		x := (v.(string))
		o.SetRevision(x)
	}

	if _, ok := d.GetOk("rn"); ok {
		v := d.Get("rn")
		x := (v.(string))
		o.SetRn(x)
	}

	if _, ok := d.GetOk("serial"); ok {
		v := d.Get("serial")
		x := (v.(string))
		o.SetSerial(x)
	}

	if _, ok := d.GetOk("shared_scope"); ok {
		v := d.Get("shared_scope")
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if _, ok := d.GetOk("source_object_type"); ok {
		v := d.Get("source_object_type")
		x := (v.(string))
		o.SetSourceObjectType(x)
	}

	if _, ok := d.GetOk("switch_id"); ok {
		v := d.Get("switch_id")
		x := (v.(string))
		o.SetSwitchId(x)
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

	if _, ok := d.GetOk("thermal"); ok {
		v := d.Get("thermal")
		x := (v.(string))
		o.SetThermal(x)
	}

	if _, ok := d.GetOk("total_memory"); ok {
		v := d.Get("total_memory")
		x := int64(v.(int))
		o.SetTotalMemory(x)
	}

	if _, ok := d.GetOk("vendor"); ok {
		v := d.Get("vendor")
		x := (v.(string))
		o.SetVendor(x)
	}

	if _, ok := d.GetOk("nr_version"); ok {
		v := d.Get("nr_version")
		x := (v.(string))
		o.SetVersion(x)
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

	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
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
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
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
	if v, ok := d.GetOk("firmware"); ok {
		x := (v.(string))
		o.SetFirmware(x)
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
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.SetIpv4Address(x)
	}
	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("num_ether_ports"); ok {
		x := int64(v.(int))
		o.SetNumEtherPorts(x)
	}
	if v, ok := d.GetOk("num_ether_ports_configured"); ok {
		x := int64(v.(int))
		o.SetNumEtherPortsConfigured(x)
	}
	if v, ok := d.GetOk("num_ether_ports_link_up"); ok {
		x := int64(v.(int))
		o.SetNumEtherPortsLinkUp(x)
	}
	if v, ok := d.GetOk("num_expansion_modules"); ok {
		x := int64(v.(int))
		o.SetNumExpansionModules(x)
	}
	if v, ok := d.GetOk("num_fc_ports"); ok {
		x := int64(v.(int))
		o.SetNumFcPorts(x)
	}
	if v, ok := d.GetOk("num_fc_ports_configured"); ok {
		x := int64(v.(int))
		o.SetNumFcPortsConfigured(x)
	}
	if v, ok := d.GetOk("num_fc_ports_link_up"); ok {
		x := int64(v.(int))
		o.SetNumFcPortsLinkUp(x)
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
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("source_object_type"); ok {
		x := (v.(string))
		o.SetSourceObjectType(x)
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SetSwitchId(x)
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.SetThermal(x)
	}
	if v, ok := d.GetOk("total_memory"); ok {
		x := int64(v.(int))
		o.SetTotalMemory(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NetworkElementSummary object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NetworkApi.GetNetworkElementSummaryList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of NetworkElementSummary: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of NetworkElementSummary: %s", responseErr.Error())
	}
	count := countResponse.NetworkElementSummaryList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for NetworkElementSummary data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var networkElementSummaryResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NetworkApi.GetNetworkElementSummaryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching NetworkElementSummary: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching NetworkElementSummary: %s", responseErr.Error())
		}
		results := resMo.NetworkElementSummaryList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_evac_state"] = (s.GetAdminEvacState())
				temp["admin_inband_interface_state"] = (s.GetAdminInbandInterfaceState())

				temp["alarm_summary"] = flattenMapComputeAlarmSummary(s.GetAlarmSummary(), d)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["available_memory"] = (s.GetAvailableMemory())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["ethernet_mode"] = (s.GetEthernetMode())
				temp["ethernet_switching_mode"] = (s.GetEthernetSwitchingMode())
				temp["fault_summary"] = (s.GetFaultSummary())
				temp["fc_mode"] = (s.GetFcMode())
				temp["fc_switching_mode"] = (s.GetFcSwitchingMode())
				temp["firmware"] = (s.GetFirmware())
				temp["inband_ip_address"] = (s.GetInbandIpAddress())
				temp["inband_ip_gateway"] = (s.GetInbandIpGateway())
				temp["inband_ip_mask"] = (s.GetInbandIpMask())
				temp["inband_vlan"] = (s.GetInbandVlan())
				temp["ipv4_address"] = (s.GetIpv4Address())
				temp["management_mode"] = (s.GetManagementMode())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["num_ether_ports"] = (s.GetNumEtherPorts())
				temp["num_ether_ports_configured"] = (s.GetNumEtherPortsConfigured())
				temp["num_ether_ports_link_up"] = (s.GetNumEtherPortsLinkUp())
				temp["num_expansion_modules"] = (s.GetNumExpansionModules())
				temp["num_fc_ports"] = (s.GetNumFcPorts())
				temp["num_fc_ports_configured"] = (s.GetNumFcPortsConfigured())
				temp["num_fc_ports_link_up"] = (s.GetNumFcPortsLinkUp())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_evac_state"] = (s.GetOperEvacState())
				temp["operability"] = (s.GetOperability())
				temp["out_of_band_ip_address"] = (s.GetOutOfBandIpAddress())
				temp["out_of_band_ip_gateway"] = (s.GetOutOfBandIpGateway())
				temp["out_of_band_ip_mask"] = (s.GetOutOfBandIpMask())
				temp["out_of_band_ipv4_address"] = (s.GetOutOfBandIpv4Address())
				temp["out_of_band_ipv4_gateway"] = (s.GetOutOfBandIpv4Gateway())
				temp["out_of_band_ipv4_mask"] = (s.GetOutOfBandIpv4Mask())
				temp["out_of_band_ipv6_address"] = (s.GetOutOfBandIpv6Address())
				temp["out_of_band_ipv6_gateway"] = (s.GetOutOfBandIpv6Gateway())
				temp["out_of_band_ipv6_prefix"] = (s.GetOutOfBandIpv6Prefix())
				temp["out_of_band_mac"] = (s.GetOutOfBandMac())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["presence"] = (s.GetPresence())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())
				temp["serial"] = (s.GetSerial())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["source_object_type"] = (s.GetSourceObjectType())
				temp["switch_id"] = (s.GetSwitchId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["thermal"] = (s.GetThermal())
				temp["total_memory"] = (s.GetTotalMemory())
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				networkElementSummaryResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(networkElementSummaryResults))
	if err := d.Set("results", networkElementSummaryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(networkElementSummaryResults[0]["moid"].(string))
	return de
}
