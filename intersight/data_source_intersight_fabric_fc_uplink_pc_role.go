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

func dataSourceFabricFcUplinkPcRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricFcUplinkPcRoleRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"admin_speed": {
				Description: "Admin configured speed for the port.\n* `Auto` - Admin configurable speed AUTO ( default ).\n* `8Gbps` - Admin configurable speed 8Gbps.\n* `16Gbps` - Admin configurable speed 16Gbps.\n* `32Gbps` - Admin configurable speed 32Gbps.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fill_pattern": {
				Description: "Fill pattern to differentiate the configs in NPIV.\n* `Idle` - Fc Fill Pattern type Idle.\n* `Arbff` - Fc Fill Pattern type Arbff.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pc_id": {
				Description: "Unique Identifier of the port-channel, local to this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"port_policy": {
				Description: "A reference to a fabricPortPolicy resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"aggregate_port_id": {
							Description: "Breakout port Identifier of the Switch Interface.\nWhen a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.\nWhen a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,\ne.g. the id of the port on the switch.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"port_id": {
							Description: "Port Identifier of the Switch/FEX/Chassis Interface.\nWhen a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,\ne.g. the id of the port on the switch, FEX or chassis.\nWhen a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"slot_id": {
							Description: "Slot Identifier of the Switch/FEX/Chassis Interface.",
							Type:        schema.TypeInt,
							Optional:    true,
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
			"vsan_id": {
				Description: "Virtual San Identifier associated to the FC port.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}

func dataSourceFabricFcUplinkPcRoleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricFcUplinkPcRole{}
	if v, ok := d.GetOk("admin_speed"); ok {
		x := (v.(string))
		o.SetAdminSpeed(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("fill_pattern"); ok {
		x := (v.(string))
		o.SetFillPattern(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pc_id"); ok {
		x := int64(v.(int))
		o.SetPcId(x)
	}
	if v, ok := d.GetOk("vsan_id"); ok {
		x := int64(v.(int))
		o.SetVsanId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricFcUplinkPcRole object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricFcUplinkPcRoleList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching FabricFcUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for FabricFcUplinkPcRole list: %s", err.Error())
	}
	var s = &models.FabricFcUplinkPcRoleList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to FabricFcUplinkPcRole list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for FabricFcUplinkPcRole did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.FabricFcUplinkPcRole{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("admin_speed", (s.GetAdminSpeed())); err != nil {
				return diag.Errorf("error occurred while setting property AdminSpeed: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("fill_pattern", (s.GetFillPattern())); err != nil {
				return diag.Errorf("error occurred while setting property FillPattern: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("pc_id", (s.GetPcId())); err != nil {
				return diag.Errorf("error occurred while setting property PcId: %s", err.Error())
			}

			if err := d.Set("port_policy", flattenMapFabricPortPolicyRelationship(s.GetPortPolicy(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PortPolicy: %s", err.Error())
			}

			if err := d.Set("ports", flattenListFabricPortIdentifier(s.GetPorts(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Ports: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("vsan_id", (s.GetVsanId())); err != nil {
				return diag.Errorf("error occurred while setting property VsanId: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
