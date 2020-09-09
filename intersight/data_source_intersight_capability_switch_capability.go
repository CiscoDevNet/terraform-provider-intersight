package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCapabilitySwitchCapability() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCapabilitySwitchCapabilityRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"max_active_span_sessions": {
				Description: "Maximum allowed Traffic Monitoring (SPAN) sessions on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_ethernet_port_channel_members": {
				Description: "Maximum allowed Ethernet Uplink Port-channel members for each Uplink Port-channel on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_ethernet_port_channels": {
				Description: "Maximum allowed Ethernet Uplink Port-channels on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_ethernet_uplink_ports": {
				Description: "Maximum allowed Ethernet Uplink Ports on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_fc_fcoe_port_channels": {
				Description: "Total maximum Fc and Fcoe Port-channels allowed on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_fc_port_channel_members": {
				Description: "Maximum allowed FC Uplink Port-channel members for each FCoE Port-channel on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_fcoe_port_channel_members": {
				Description: "Maximum allowed FCoE Uplink Port-channel members for each FCoE Port-channel on this switch.",
				Type:        schema.TypeInt,
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
			"max_vsans_supported": {
				Description: "Maximum number of Vsans supported on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_active_fans": {
				Description: "Minimum number of fans needed to be active/running on this switch.",
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
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"section": {
				Description: "A reference to a capabilitySection resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"vp_compression_supported": {
				Description: "VP Compression support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
		},
	}
}

func dataSourceCapabilitySwitchCapabilityRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewCapabilitySwitchCapabilityWithDefaults()
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
	if v, ok := d.GetOk("max_active_span_sessions"); ok {
		x := int64(v.(int))
		o.SetMaxActiveSpanSessions(x)
	}
	if v, ok := d.GetOk("max_ethernet_port_channel_members"); ok {
		x := int64(v.(int))
		o.SetMaxEthernetPortChannelMembers(x)
	}
	if v, ok := d.GetOk("max_ethernet_port_channels"); ok {
		x := int64(v.(int))
		o.SetMaxEthernetPortChannels(x)
	}
	if v, ok := d.GetOk("max_ethernet_uplink_ports"); ok {
		x := int64(v.(int))
		o.SetMaxEthernetUplinkPorts(x)
	}
	if v, ok := d.GetOk("max_fc_fcoe_port_channels"); ok {
		x := int64(v.(int))
		o.SetMaxFcFcoePortChannels(x)
	}
	if v, ok := d.GetOk("max_fc_port_channel_members"); ok {
		x := int64(v.(int))
		o.SetMaxFcPortChannelMembers(x)
	}
	if v, ok := d.GetOk("max_fcoe_port_channel_members"); ok {
		x := int64(v.(int))
		o.SetMaxFcoePortChannelMembers(x)
	}
	if v, ok := d.GetOk("max_ports"); ok {
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}
	if v, ok := d.GetOk("max_slots"); ok {
		x := int64(v.(int))
		o.SetMaxSlots(x)
	}
	if v, ok := d.GetOk("max_vsans_supported"); ok {
		x := int64(v.(int))
		o.SetMaxVsansSupported(x)
	}
	if v, ok := d.GetOk("min_active_fans"); ok {
		x := int64(v.(int))
		o.SetMinActiveFans(x)
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
	if v, ok := d.GetOk("vp_compression_supported"); ok {
		x := (v.(bool))
		o.SetVpCompressionSupported(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.CapabilitySwitchCapabilityList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to CapabilitySwitchCapability: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewCapabilitySwitchCapabilityWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("default_fcoe_vlan", (s.GetDefaultFcoeVlan())); err != nil {
				return fmt.Errorf("error occurred while setting property DefaultFcoeVlan: %+v", err)
			}
			if err := d.Set("dynamic_vifs_supported", (s.GetDynamicVifsSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property DynamicVifsSupported: %+v", err)
			}
			if err := d.Set("fan_modules_supported", (s.GetFanModulesSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property FanModulesSupported: %+v", err)
			}

			if err := d.Set("fc_end_host_mode_reserved_vsans", flattenListCapabilityPortRange(s.GetFcEndHostModeReservedVsans(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property FcEndHostModeReservedVsans: %+v", err)
			}
			if err := d.Set("fc_uplink_ports_auto_negotiation_supported", (s.GetFcUplinkPortsAutoNegotiationSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property FcUplinkPortsAutoNegotiationSupported: %+v", err)
			}
			if err := d.Set("locator_beacon_supported", (s.GetLocatorBeaconSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property LocatorBeaconSupported: %+v", err)
			}
			if err := d.Set("max_active_span_sessions", (s.GetMaxActiveSpanSessions())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxActiveSpanSessions: %+v", err)
			}
			if err := d.Set("max_ethernet_port_channel_members", (s.GetMaxEthernetPortChannelMembers())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxEthernetPortChannelMembers: %+v", err)
			}
			if err := d.Set("max_ethernet_port_channels", (s.GetMaxEthernetPortChannels())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxEthernetPortChannels: %+v", err)
			}
			if err := d.Set("max_ethernet_uplink_ports", (s.GetMaxEthernetUplinkPorts())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxEthernetUplinkPorts: %+v", err)
			}
			if err := d.Set("max_fc_fcoe_port_channels", (s.GetMaxFcFcoePortChannels())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxFcFcoePortChannels: %+v", err)
			}
			if err := d.Set("max_fc_port_channel_members", (s.GetMaxFcPortChannelMembers())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxFcPortChannelMembers: %+v", err)
			}
			if err := d.Set("max_fcoe_port_channel_members", (s.GetMaxFcoePortChannelMembers())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxFcoePortChannelMembers: %+v", err)
			}
			if err := d.Set("max_ports", (s.GetMaxPorts())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxPorts: %+v", err)
			}
			if err := d.Set("max_slots", (s.GetMaxSlots())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxSlots: %+v", err)
			}
			if err := d.Set("max_vsans_supported", (s.GetMaxVsansSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property MaxVsansSupported: %+v", err)
			}
			if err := d.Set("min_active_fans", (s.GetMinActiveFans())); err != nil {
				return fmt.Errorf("error occurred while setting property MinActiveFans: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("pid", (s.GetPid())); err != nil {
				return fmt.Errorf("error occurred while setting property Pid: %+v", err)
			}

			if err := d.Set("ports_supporting100g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting100gSpeed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupporting100gSpeed: %+v", err)
			}

			if err := d.Set("ports_supporting10g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting10gSpeed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupporting10gSpeed: %+v", err)
			}

			if err := d.Set("ports_supporting1g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting1gSpeed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupporting1gSpeed: %+v", err)
			}

			if err := d.Set("ports_supporting25g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting25gSpeed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupporting25gSpeed: %+v", err)
			}

			if err := d.Set("ports_supporting40g_speed", flattenListCapabilityPortRange(s.GetPortsSupporting40gSpeed(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupporting40gSpeed: %+v", err)
			}

			if err := d.Set("ports_supporting_breakout", flattenListCapabilityPortRange(s.GetPortsSupportingBreakout(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupportingBreakout: %+v", err)
			}

			if err := d.Set("ports_supporting_fcoe", flattenListCapabilityPortRange(s.GetPortsSupportingFcoe(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupportingFcoe: %+v", err)
			}

			if err := d.Set("ports_supporting_server_role", flattenListCapabilityPortRange(s.GetPortsSupportingServerRole(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property PortsSupportingServerRole: %+v", err)
			}

			if err := d.Set("reserved_vsans", flattenListCapabilityPortRange(s.GetReservedVsans(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property ReservedVsans: %+v", err)
			}

			if err := d.Set("section", flattenMapCapabilitySectionRelationship(s.GetSection(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Section: %+v", err)
			}
			if err := d.Set("sereno_netflow_supported", (s.GetSerenoNetflowSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property SerenoNetflowSupported: %+v", err)
			}
			if err := d.Set("sku", (s.GetSku())); err != nil {
				return fmt.Errorf("error occurred while setting property Sku: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}

			if err := d.Set("unified_ports", flattenListCapabilityPortRange(s.GetUnifiedPorts(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property UnifiedPorts: %+v", err)
			}
			if err := d.Set("unified_rule", (s.GetUnifiedRule())); err != nil {
				return fmt.Errorf("error occurred while setting property UnifiedRule: %+v", err)
			}
			if err := d.Set("vid", (s.GetVid())); err != nil {
				return fmt.Errorf("error occurred while setting property Vid: %+v", err)
			}
			if err := d.Set("vp_compression_supported", (s.GetVpCompressionSupported())); err != nil {
				return fmt.Errorf("error occurred while setting property VpCompressionSupported: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
