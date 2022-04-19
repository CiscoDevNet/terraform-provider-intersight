package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCapabilitySwitchCapability() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCapabilitySwitchCapabilityCreate,
		ReadContext:   resourceCapabilitySwitchCapabilityRead,
		UpdateContext: resourceCapabilitySwitchCapabilityUpdate,
		DeleteContext: resourceCapabilitySwitchCapabilityDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CustomizeTagDiff,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ancestors": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "capability.SwitchCapability",
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"default_fcoe_vlan": {
				Description: "Default Fcoe VLAN associated with this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
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
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "capability.SwitchNetworkLimits",
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
							Default:     "capability.SwitchNetworkLimits",
						},
					},
				},
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "capability.SwitchCapability",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `UCS-FI-6536` - The standard 5th generation UCS Fabric Interconnect with 36 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "UCS-FI-6454",
			},
			"ports_supporting100g_speed": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting10g_speed": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting1g_speed": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting25g_speed": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting40g_speed": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting_breakout": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting_fcoe": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"ports_supporting_server_role": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"reserved_vsans": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"sereno_netflow_supported": {
				Description: "Sereno Adaptor with Netflow support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
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
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "capability.SwitchStorageLimits",
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
							Default:     "capability.SwitchStorageLimits",
						},
					},
				},
			},
			"switching_mode_capabilities": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.SwitchingModeCapability",
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "capability.SwitchingModeCapability",
						},
						"switching_mode": {
							Description: "Switching mode type (endhost, switch) of the switch.\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "end-host",
						},
						"vp_compression_supported": {
							Description: "VP Compression support on this switch.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
			"system_limits": {
				Description: "List of system limitations for this switch.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "capability.SwitchSystemLimits",
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
							Default:     "capability.SwitchSystemLimits",
						},
					},
				},
			},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.PortRange",
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
							Default:     "capability.PortRange",
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
			},
			"unified_rule": {
				Description: "The Slider rule for Unified ports on this switch.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"timestamp": {
							Description: "The time this versioned Managed Object was created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceCapabilitySwitchCapabilityCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewCapabilitySwitchCapabilityWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("capability.SwitchCapability")

	if v, ok := d.GetOkExists("default_fcoe_vlan"); ok {
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}

	if v, ok := d.GetOkExists("dynamic_vifs_supported"); ok {
		x := (v.(bool))
		o.SetDynamicVifsSupported(x)
	}

	if v, ok := d.GetOkExists("fan_modules_supported"); ok {
		x := (v.(bool))
		o.SetFanModulesSupported(x)
	}

	if v, ok := d.GetOk("fc_end_host_mode_reserved_vsans"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetFcEndHostModeReservedVsans(x)
		}
	}

	if v, ok := d.GetOkExists("fc_uplink_ports_auto_negotiation_supported"); ok {
		x := (v.(bool))
		o.SetFcUplinkPortsAutoNegotiationSupported(x)
	}

	if v, ok := d.GetOkExists("locator_beacon_supported"); ok {
		x := (v.(bool))
		o.SetLocatorBeaconSupported(x)
	}

	if v, ok := d.GetOkExists("max_ports"); ok {
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}

	if v, ok := d.GetOkExists("max_slots"); ok {
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

	if v, ok := d.GetOk("network_limits"); ok {
		p := make([]models.CapabilitySwitchNetworkLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewCapabilitySwitchNetworkLimitsWithDefaults()
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
			if v, ok := l["max_compressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxCompressedPortVlanCount(x)
				}
			}
			if v, ok := l["max_uncompressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxUncompressedPortVlanCount(x)
				}
			}
			if v, ok := l["maximum_active_traffic_monitoring_sessions"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumActiveTrafficMonitoringSessions(x)
				}
			}
			if v, ok := l["maximum_ethernet_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetPortChannels(x)
				}
			}
			if v, ok := l["maximum_ethernet_uplink_ports"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetUplinkPorts(x)
				}
			}
			if v, ok := l["maximum_fc_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_fc_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannels(x)
				}
			}
			if v, ok := l["maximum_igmp_groups"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumIgmpGroups(x)
				}
			}
			if v, ok := l["maximum_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_primary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPrimaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan_per_primary"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlanPerPrimary(x)
				}
			}
			if v, ok := l["maximum_vifs"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVifs(x)
				}
			}
			if v, ok := l["maximum_vlans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVlans(x)
				}
			}
			if v, ok := l["minimum_active_fans"]; ok {
				{
					x := int64(v.(int))
					o.SetMinimumActiveFans(x)
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
			o.SetNetworkLimits(x)
		}
	}

	o.SetObjectType("capability.SwitchCapability")

	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}

	if v, ok := d.GetOk("ports_supporting100g_speed"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupporting100gSpeed(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting10g_speed"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupporting10gSpeed(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting1g_speed"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupporting1gSpeed(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting25g_speed"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupporting25gSpeed(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting40g_speed"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupporting40gSpeed(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting_breakout"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupportingBreakout(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting_fcoe"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupportingFcoe(x)
		}
	}

	if v, ok := d.GetOk("ports_supporting_server_role"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPortsSupportingServerRole(x)
		}
	}

	if v, ok := d.GetOk("reserved_vsans"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetReservedVsans(x)
		}
	}

	if v, ok := d.GetOkExists("sereno_netflow_supported"); ok {
		x := (v.(bool))
		o.SetSerenoNetflowSupported(x)
	}

	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}

	if v, ok := d.GetOk("storage_limits"); ok {
		p := make([]models.CapabilitySwitchStorageLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewCapabilitySwitchStorageLimitsWithDefaults()
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
			if v, ok := l["maximum_user_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumUserZoneCount(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfaces(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces_per_blade_server"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfacesPerBladeServer(x)
				}
			}
			if v, ok := l["maximum_vsans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVsans(x)
				}
			}
			if v, ok := l["maximum_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumZoneCount(x)
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
			o.SetStorageLimits(x)
		}
	}

	if v, ok := d.GetOk("switching_mode_capabilities"); ok {
		x := make([]models.CapabilitySwitchingModeCapability, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilitySwitchingModeCapabilityWithDefaults()
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
			o.SetClassId("capability.SwitchingModeCapability")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["switching_mode"]; ok {
				{
					x := (v.(string))
					o.SetSwitchingMode(x)
				}
			}
			if v, ok := l["vp_compression_supported"]; ok {
				{
					x := (v.(bool))
					o.SetVpCompressionSupported(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetSwitchingModeCapabilities(x)
		}
	}

	if v, ok := d.GetOk("system_limits"); ok {
		p := make([]models.CapabilitySwitchSystemLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewCapabilitySwitchSystemLimitsWithDefaults()
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
			if v, ok := l["maximum_chassis_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumChassisCount(x)
				}
			}
			if v, ok := l["maximum_fex_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFexPerDomain(x)
				}
			}
			if v, ok := l["maximum_servers_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumServersPerDomain(x)
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
			o.SetSystemLimits(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
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
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("unified_ports"); ok {
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityPortRangeWithDefaults()
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetUnifiedPorts(x)
		}
	}

	if v, ok := d.GetOk("unified_rule"); ok {
		x := (v.(string))
		o.SetUnifiedRule(x)
	}

	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	r := conn.ApiClient.CapabilityApi.CreateCapabilitySwitchCapability(conn.ctx).CapabilitySwitchCapability(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating CapabilitySwitchCapability: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceCapabilitySwitchCapabilityRead(c, d, meta)...)
}

func resourceCapabilitySwitchCapabilityRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	r := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilitySwitchCapability object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("default_fcoe_vlan", (s.GetDefaultFcoeVlan())); err != nil {
		return diag.Errorf("error occurred while setting property DefaultFcoeVlan in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("dynamic_vifs_supported", (s.GetDynamicVifsSupported())); err != nil {
		return diag.Errorf("error occurred while setting property DynamicVifsSupported in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("fan_modules_supported", (s.GetFanModulesSupported())); err != nil {
		return diag.Errorf("error occurred while setting property FanModulesSupported in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("fc_end_host_mode_reserved_vsans", flattenListCapabilityPortRange(s.GetFcEndHostModeReservedVsans(), d)); err != nil {
		return diag.Errorf("error occurred while setting property FcEndHostModeReservedVsans in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("fc_uplink_ports_auto_negotiation_supported", (s.GetFcUplinkPortsAutoNegotiationSupported())); err != nil {
		return diag.Errorf("error occurred while setting property FcUplinkPortsAutoNegotiationSupported in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("locator_beacon_supported", (s.GetLocatorBeaconSupported())); err != nil {
		return diag.Errorf("error occurred while setting property LocatorBeaconSupported in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("max_ports", (s.GetMaxPorts())); err != nil {
		return diag.Errorf("error occurred while setting property MaxPorts in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("max_slots", (s.GetMaxSlots())); err != nil {
		return diag.Errorf("error occurred while setting property MaxSlots in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("network_limits", flattenMapCapabilitySwitchNetworkLimits(s.GetNetworkLimits(), d)); err != nil {
		return diag.Errorf("error occurred while setting property NetworkLimits in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("pid", (s.GetPid())); err != nil {
		return diag.Errorf("error occurred while setting property Pid in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting100g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting100gSpeed(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupporting100gSpeed in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting10g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting10gSpeed(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupporting10gSpeed in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting1g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting1gSpeed(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupporting1gSpeed in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting25g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting25gSpeed(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupporting25gSpeed in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting40g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting40gSpeed(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupporting40gSpeed in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting_breakout", flattenListCapabilityPortRange(s.GetPortsSupportingBreakout(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupportingBreakout in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting_fcoe", flattenListCapabilityPortRange(s.GetPortsSupportingFcoe(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupportingFcoe in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("ports_supporting_server_role", flattenListCapabilityPortRange(s.GetPortsSupportingServerRole(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortsSupportingServerRole in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("reserved_vsans", flattenListCapabilityPortRange(s.GetReservedVsans(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ReservedVsans in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("sereno_netflow_supported", (s.GetSerenoNetflowSupported())); err != nil {
		return diag.Errorf("error occurred while setting property SerenoNetflowSupported in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("sku", (s.GetSku())); err != nil {
		return diag.Errorf("error occurred while setting property Sku in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("storage_limits", flattenMapCapabilitySwitchStorageLimits(s.GetStorageLimits(), d)); err != nil {
		return diag.Errorf("error occurred while setting property StorageLimits in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("switching_mode_capabilities", flattenListCapabilitySwitchingModeCapability(s.GetSwitchingModeCapabilities(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SwitchingModeCapabilities in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("system_limits", flattenMapCapabilitySwitchSystemLimits(s.GetSystemLimits(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SystemLimits in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("unified_ports", flattenListCapabilityPortRange(s.GetUnifiedPorts(), d)); err != nil {
		return diag.Errorf("error occurred while setting property UnifiedPorts in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("unified_rule", (s.GetUnifiedRule())); err != nil {
		return diag.Errorf("error occurred while setting property UnifiedRule in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in CapabilitySwitchCapability object: %s", err.Error())
	}

	if err := d.Set("vid", (s.GetVid())); err != nil {
		return diag.Errorf("error occurred while setting property Vid in CapabilitySwitchCapability object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceCapabilitySwitchCapabilityUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilitySwitchCapability{}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("capability.SwitchCapability")

	if d.HasChange("default_fcoe_vlan") {
		v := d.Get("default_fcoe_vlan")
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}

	if d.HasChange("dynamic_vifs_supported") {
		v := d.Get("dynamic_vifs_supported")
		x := (v.(bool))
		o.SetDynamicVifsSupported(x)
	}

	if d.HasChange("fan_modules_supported") {
		v := d.Get("fan_modules_supported")
		x := (v.(bool))
		o.SetFanModulesSupported(x)
	}

	if d.HasChange("fc_end_host_mode_reserved_vsans") {
		v := d.Get("fc_end_host_mode_reserved_vsans")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetFcEndHostModeReservedVsans(x)
	}

	if d.HasChange("fc_uplink_ports_auto_negotiation_supported") {
		v := d.Get("fc_uplink_ports_auto_negotiation_supported")
		x := (v.(bool))
		o.SetFcUplinkPortsAutoNegotiationSupported(x)
	}

	if d.HasChange("locator_beacon_supported") {
		v := d.Get("locator_beacon_supported")
		x := (v.(bool))
		o.SetLocatorBeaconSupported(x)
	}

	if d.HasChange("max_ports") {
		v := d.Get("max_ports")
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}

	if d.HasChange("max_slots") {
		v := d.Get("max_slots")
		x := int64(v.(int))
		o.SetMaxSlots(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if d.HasChange("network_limits") {
		v := d.Get("network_limits")
		p := make([]models.CapabilitySwitchNetworkLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchNetworkLimits{}
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
			if v, ok := l["max_compressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxCompressedPortVlanCount(x)
				}
			}
			if v, ok := l["max_uncompressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxUncompressedPortVlanCount(x)
				}
			}
			if v, ok := l["maximum_active_traffic_monitoring_sessions"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumActiveTrafficMonitoringSessions(x)
				}
			}
			if v, ok := l["maximum_ethernet_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetPortChannels(x)
				}
			}
			if v, ok := l["maximum_ethernet_uplink_ports"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetUplinkPorts(x)
				}
			}
			if v, ok := l["maximum_fc_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_fc_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannels(x)
				}
			}
			if v, ok := l["maximum_igmp_groups"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumIgmpGroups(x)
				}
			}
			if v, ok := l["maximum_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_primary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPrimaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan_per_primary"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlanPerPrimary(x)
				}
			}
			if v, ok := l["maximum_vifs"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVifs(x)
				}
			}
			if v, ok := l["maximum_vlans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVlans(x)
				}
			}
			if v, ok := l["minimum_active_fans"]; ok {
				{
					x := int64(v.(int))
					o.SetMinimumActiveFans(x)
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
			o.SetNetworkLimits(x)
		}
	}

	o.SetObjectType("capability.SwitchCapability")

	if d.HasChange("pid") {
		v := d.Get("pid")
		x := (v.(string))
		o.SetPid(x)
	}

	if d.HasChange("ports_supporting100g_speed") {
		v := d.Get("ports_supporting100g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting100gSpeed(x)
	}

	if d.HasChange("ports_supporting10g_speed") {
		v := d.Get("ports_supporting10g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting10gSpeed(x)
	}

	if d.HasChange("ports_supporting1g_speed") {
		v := d.Get("ports_supporting1g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting1gSpeed(x)
	}

	if d.HasChange("ports_supporting25g_speed") {
		v := d.Get("ports_supporting25g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting25gSpeed(x)
	}

	if d.HasChange("ports_supporting40g_speed") {
		v := d.Get("ports_supporting40g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting40gSpeed(x)
	}

	if d.HasChange("ports_supporting_breakout") {
		v := d.Get("ports_supporting_breakout")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingBreakout(x)
	}

	if d.HasChange("ports_supporting_fcoe") {
		v := d.Get("ports_supporting_fcoe")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingFcoe(x)
	}

	if d.HasChange("ports_supporting_server_role") {
		v := d.Get("ports_supporting_server_role")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingServerRole(x)
	}

	if d.HasChange("reserved_vsans") {
		v := d.Get("reserved_vsans")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetReservedVsans(x)
	}

	if d.HasChange("sereno_netflow_supported") {
		v := d.Get("sereno_netflow_supported")
		x := (v.(bool))
		o.SetSerenoNetflowSupported(x)
	}

	if d.HasChange("sku") {
		v := d.Get("sku")
		x := (v.(string))
		o.SetSku(x)
	}

	if d.HasChange("storage_limits") {
		v := d.Get("storage_limits")
		p := make([]models.CapabilitySwitchStorageLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchStorageLimits{}
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
			if v, ok := l["maximum_user_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumUserZoneCount(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfaces(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces_per_blade_server"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfacesPerBladeServer(x)
				}
			}
			if v, ok := l["maximum_vsans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVsans(x)
				}
			}
			if v, ok := l["maximum_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumZoneCount(x)
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
			o.SetStorageLimits(x)
		}
	}

	if d.HasChange("switching_mode_capabilities") {
		v := d.Get("switching_mode_capabilities")
		x := make([]models.CapabilitySwitchingModeCapability, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilitySwitchingModeCapability{}
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
			o.SetClassId("capability.SwitchingModeCapability")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["switching_mode"]; ok {
				{
					x := (v.(string))
					o.SetSwitchingMode(x)
				}
			}
			if v, ok := l["vp_compression_supported"]; ok {
				{
					x := (v.(bool))
					o.SetVpCompressionSupported(x)
				}
			}
			x = append(x, *o)
		}
		o.SetSwitchingModeCapabilities(x)
	}

	if d.HasChange("system_limits") {
		v := d.Get("system_limits")
		p := make([]models.CapabilitySwitchSystemLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchSystemLimits{}
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
			if v, ok := l["maximum_chassis_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumChassisCount(x)
				}
			}
			if v, ok := l["maximum_fex_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFexPerDomain(x)
				}
			}
			if v, ok := l["maximum_servers_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumServersPerDomain(x)
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
			o.SetSystemLimits(x)
		}
	}

	if d.HasChange("tags") {
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

	if d.HasChange("unified_ports") {
		v := d.Get("unified_ports")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
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
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetUnifiedPorts(x)
	}

	if d.HasChange("unified_rule") {
		v := d.Get("unified_rule")
		x := (v.(string))
		o.SetUnifiedRule(x)
	}

	if d.HasChange("vid") {
		v := d.Get("vid")
		x := (v.(string))
		o.SetVid(x)
	}

	r := conn.ApiClient.CapabilityApi.UpdateCapabilitySwitchCapability(conn.ctx, d.Id()).CapabilitySwitchCapability(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating CapabilitySwitchCapability: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceCapabilitySwitchCapabilityRead(c, d, meta)...)
}

func resourceCapabilitySwitchCapabilityDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.CapabilityApi.DeleteCapabilitySwitchCapability(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilitySwitchCapabilityDelete: CapabilitySwitchCapability object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting CapabilitySwitchCapability object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting CapabilitySwitchCapability object: %s", deleteErr.Error())
	}
	return de
}
