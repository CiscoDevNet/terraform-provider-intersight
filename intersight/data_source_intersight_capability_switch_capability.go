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

func dataSourceCapabilitySwitchCapability() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilitySwitchCapabilityRead,
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
			"default_fcoe_vlan": {
				Description: "Default Fcoe VLAN associated with this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"dynamic_vifs_supported": {
				Description: "Dynamic VIFs support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"fan_modules_supported": {
				Description: "Fan Modules support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"fc_end_host_mode_reserved_vsans": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"fc_uplink_ports_auto_negotiation_supported": {
				Description: "Fc Uplink ports auto negotiation speed support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"locator_beacon_supported": {
				Description: "Locator Beacon LED support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_ports": {
				Description: "Maximum allowed physical ports on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_slots": {
				Description: "Maximum allowed physical slots on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "An unique identifer for a capability descriptor.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_limits": {
				Description: "List of network limitations for this switch.",
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
						"max_compressed_port_vlan_count": {
							Description: "Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"max_uncompressed_port_vlan_count": {
							Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_active_traffic_monitoring_sessions": {
							Description: "Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_ethernet_port_channels": {
							Description: "Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_ethernet_uplink_ports": {
							Description: "Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_fc_port_channel_members": {
							Description: "Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_fc_port_channels": {
							Description: "Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_igmp_groups": {
							Description: "Maximum configurable IGMP Groups on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_port_channel_members": {
							Description: "Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_primary_vlan": {
							Description: "Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_secondary_vlan": {
							Description: "Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_secondary_vlan_per_primary": {
							Description: "Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_vifs": {
							Description: "Maximum allowes VIFs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_vlans": {
							Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"minimum_active_fans": {
							Description: "Minimum required fans in 'active' state for this Switch/Fabric-Interconnect.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ports_supporting100g_speed": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting10g_speed": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting1g_speed": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting25g_speed": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting40g_speed": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting_breakout": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting_fcoe": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"ports_supporting_server_role": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"reserved_vsans": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"sereno_netflow_supported": {
				Description: "Sereno Adaptor with Netflow support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"sku": {
				Description: "SKU information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"storage_limits": {
				Description: "List of storage limitations for this switch.",
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
						"maximum_user_zone_count": {
							Description: "Maximum user zones per Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_virtual_fc_interfaces": {
							Description: "Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_virtual_fc_interfaces_per_blade_server": {
							Description: "Maximum configurable Virtual Fibre Channel interfaces per blade.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_vsans": {
							Description: "Maximum configurable VSANs on Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_zone_count": {
							Description: "Zone limit per Switch/Fabric-Interconnect.",
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
			"switching_mode_capabilities": {
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
							Computed:    true,
						},
						"switching_mode": {
							Description: "Switching mode type (endhost, switch) of the switch.\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vp_compression_supported": {
							Description: "VP Compression support on this switch.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"system_limits": {
				Description: "List of system limitations for this switch.",
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
						"maximum_chassis_count": {
							Description: "Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_fex_per_domain": {
							Description: "Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"maximum_servers_per_domain": {
							Description: "Maximum UCS servers per Switch/Fabric-Interconnect.",
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
			"unified_ports": {
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
						"end_port_id": {
							Description: "Ending Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"end_slot_id": {
							Description: "Ending Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"start_port_id": {
							Description: "Starting Port ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"start_slot_id": {
							Description: "Starting Slot ID in this range of ports.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"unified_rule": {
				Description: "The Slider rule for Unified ports on this switch.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceCapabilitySwitchCapabilityRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilitySwitchCapability{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_fcoe_vlan"); ok {
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}
	if v, ok := d.GetOk("dynamic_vifs_supported"); ok {
		x := (v.(bool))
		o.SetDynamicVifsSupported(x)
	}
	if v, ok := d.GetOk("fan_modules_supported"); ok {
		x := (v.(bool))
		o.SetFanModulesSupported(x)
	}
	if v, ok := d.GetOk("fc_uplink_ports_auto_negotiation_supported"); ok {
		x := (v.(bool))
		o.SetFcUplinkPortsAutoNegotiationSupported(x)
	}
	if v, ok := d.GetOk("locator_beacon_supported"); ok {
		x := (v.(bool))
		o.SetLocatorBeaconSupported(x)
	}
	if v, ok := d.GetOk("max_ports"); ok {
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}
	if v, ok := d.GetOk("max_slots"); ok {
		x := int64(v.(int))
		o.SetMaxSlots(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("sereno_netflow_supported"); ok {
		x := (v.(bool))
		o.SetSerenoNetflowSupported(x)
	}
	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}
	if v, ok := d.GetOk("unified_rule"); ok {
		x := (v.(string))
		o.SetUnifiedRule(x)
	}
	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilitySwitchCapability object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for CapabilitySwitchCapability list: %s", err.Error())
	}
	var s = &models.CapabilitySwitchCapabilityList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to CapabilitySwitchCapability list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for CapabilitySwitchCapability data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for CapabilitySwitchCapability data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.CapabilitySwitchCapability{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("default_fcoe_vlan", (s.GetDefaultFcoeVlan())); err != nil {
				return diag.Errorf("error occurred while setting property DefaultFcoeVlan: %s", err.Error())
			}
			if err := d.Set("dynamic_vifs_supported", (s.GetDynamicVifsSupported())); err != nil {
				return diag.Errorf("error occurred while setting property DynamicVifsSupported: %s", err.Error())
			}
			if err := d.Set("fan_modules_supported", (s.GetFanModulesSupported())); err != nil {
				return diag.Errorf("error occurred while setting property FanModulesSupported: %s", err.Error())
			}

			if err := d.Set("fc_end_host_mode_reserved_vsans", flattenListCapabilityPortRange(s.GetFcEndHostModeReservedVsans(), d)); err != nil {
				return diag.Errorf("error occurred while setting property FcEndHostModeReservedVsans: %s", err.Error())
			}
			if err := d.Set("fc_uplink_ports_auto_negotiation_supported", (s.GetFcUplinkPortsAutoNegotiationSupported())); err != nil {
				return diag.Errorf("error occurred while setting property FcUplinkPortsAutoNegotiationSupported: %s", err.Error())
			}
			if err := d.Set("locator_beacon_supported", (s.GetLocatorBeaconSupported())); err != nil {
				return diag.Errorf("error occurred while setting property LocatorBeaconSupported: %s", err.Error())
			}
			if err := d.Set("max_ports", (s.GetMaxPorts())); err != nil {
				return diag.Errorf("error occurred while setting property MaxPorts: %s", err.Error())
			}
			if err := d.Set("max_slots", (s.GetMaxSlots())); err != nil {
				return diag.Errorf("error occurred while setting property MaxSlots: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}

			if err := d.Set("network_limits", flattenMapCapabilitySwitchNetworkLimits(s.GetNetworkLimits(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NetworkLimits: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("pid", (s.GetPid())); err != nil {
				return diag.Errorf("error occurred while setting property Pid: %s", err.Error())
			}

			if err := d.Set("ports_supporting100g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting100gSpeed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupporting100gSpeed: %s", err.Error())
			}

			if err := d.Set("ports_supporting10g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting10gSpeed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupporting10gSpeed: %s", err.Error())
			}

			if err := d.Set("ports_supporting1g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting1gSpeed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupporting1gSpeed: %s", err.Error())
			}

			if err := d.Set("ports_supporting25g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting25gSpeed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupporting25gSpeed: %s", err.Error())
			}

			if err := d.Set("ports_supporting40g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting40gSpeed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupporting40gSpeed: %s", err.Error())
			}

			if err := d.Set("ports_supporting_breakout", flattenListCapabilityPortRange(s.GetPortsSupportingBreakout(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupportingBreakout: %s", err.Error())
			}

			if err := d.Set("ports_supporting_fcoe", flattenListCapabilityPortRange(s.GetPortsSupportingFcoe(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupportingFcoe: %s", err.Error())
			}

			if err := d.Set("ports_supporting_server_role", flattenListCapabilityPortRange(s.GetPortsSupportingServerRole(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortsSupportingServerRole: %s", err.Error())
			}

			if err := d.Set("reserved_vsans", flattenListCapabilityPortRange(s.GetReservedVsans(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ReservedVsans: %s", err.Error())
			}
			if err := d.Set("sereno_netflow_supported", (s.GetSerenoNetflowSupported())); err != nil {
				return diag.Errorf("error occurred while setting property SerenoNetflowSupported: %s", err.Error())
			}
			if err := d.Set("sku", (s.GetSku())); err != nil {
				return diag.Errorf("error occurred while setting property Sku: %s", err.Error())
			}

			if err := d.Set("storage_limits", flattenMapCapabilitySwitchStorageLimits(s.GetStorageLimits(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageLimits: %s", err.Error())
			}

			if err := d.Set("switching_mode_capabilities", flattenListCapabilitySwitchingModeCapability(s.GetSwitchingModeCapabilities(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SwitchingModeCapabilities: %s", err.Error())
			}

			if err := d.Set("system_limits", flattenMapCapabilitySwitchSystemLimits(s.GetSystemLimits(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SystemLimits: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("unified_ports", flattenListCapabilityPortRange(s.GetUnifiedPorts(), d)); err != nil {
				return diag.Errorf("error occurred while setting property UnifiedPorts: %s", err.Error())
			}
			if err := d.Set("unified_rule", (s.GetUnifiedRule())); err != nil {
				return diag.Errorf("error occurred while setting property UnifiedRule: %s", err.Error())
			}
			if err := d.Set("vid", (s.GetVid())); err != nil {
				return diag.Errorf("error occurred while setting property Vid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
