package intersight

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFabricFcUplinkPcRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFabricFcUplinkPcRoleCreate,
		ReadContext:   resourceFabricFcUplinkPcRoleRead,
		UpdateContext: resourceFabricFcUplinkPcRoleUpdate,
		DeleteContext: resourceFabricFcUplinkPcRoleDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
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
				Default:     "Auto",
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fill_pattern": {
				Description: "Fill pattern to differentiate the configs in NPIV.\n* `Idle` - Fc Fill Pattern type Idle.\n* `Arbff` - Fc Fill Pattern type Arbff.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Idle",
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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

func resourceFabricFcUplinkPcRoleCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewFabricFcUplinkPcRoleWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("admin_speed"); ok {
		x := (v.(string))
		o.SetAdminSpeed(x)
	}

	o.SetClassId("fabric.FcUplinkPcRole")

	if v, ok := d.GetOk("fill_pattern"); ok {
		x := (v.(string))
		o.SetFillPattern(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("fabric.FcUplinkPcRole")

	if v, ok := d.GetOk("pc_id"); ok {
		x := int64(v.(int))
		o.SetPcId(x)
	}

	if v, ok := d.GetOk("port_policy"); ok {
		p := make([]models.FabricPortPolicyRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
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
			p = append(p, models.MoMoRefAsFabricPortPolicyRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPortPolicy(x)
		}
	}

	if v, ok := d.GetOk("ports"); ok {
		x := make([]models.FabricPortIdentifier, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewFabricPortIdentifierWithDefaults()
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
			if v, ok := l["aggregate_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetAggregatePortId(x)
				}
			}
			o.SetClassId("fabric.PortIdentifier")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetPortId(x)
				}
			}
			if v, ok := l["slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPorts(x)
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

	if v, ok := d.GetOk("vsan_id"); ok {
		x := int64(v.(int))
		o.SetVsanId(x)
	}

	r := conn.ApiClient.FabricApi.CreateFabricFcUplinkPcRole(conn.ctx).FabricFcUplinkPcRole(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating FabricFcUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceFabricFcUplinkPcRoleRead(c, d, meta)
}

func resourceFabricFcUplinkPcRoleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.FabricApi.GetFabricFcUplinkPcRoleByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "FabricFcUplinkPcRole object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching FabricFcUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("admin_speed", (s.GetAdminSpeed())); err != nil {
		return diag.Errorf("error occurred while setting property AdminSpeed in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("fill_pattern", (s.GetFillPattern())); err != nil {
		return diag.Errorf("error occurred while setting property FillPattern in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("pc_id", (s.GetPcId())); err != nil {
		return diag.Errorf("error occurred while setting property PcId in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("port_policy", flattenMapFabricPortPolicyRelationship(s.GetPortPolicy(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PortPolicy in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("ports", flattenListFabricPortIdentifier(s.GetPorts(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ports in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in FabricFcUplinkPcRole object: %s", err.Error())
	}

	if err := d.Set("vsan_id", (s.GetVsanId())); err != nil {
		return diag.Errorf("error occurred while setting property VsanId in FabricFcUplinkPcRole object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceFabricFcUplinkPcRoleUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.FabricFcUplinkPcRole{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("admin_speed") {
		v := d.Get("admin_speed")
		x := (v.(string))
		o.SetAdminSpeed(x)
	}

	o.SetClassId("fabric.FcUplinkPcRole")

	if d.HasChange("fill_pattern") {
		v := d.Get("fill_pattern")
		x := (v.(string))
		o.SetFillPattern(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("fabric.FcUplinkPcRole")

	if d.HasChange("pc_id") {
		v := d.Get("pc_id")
		x := int64(v.(int))
		o.SetPcId(x)
	}

	if d.HasChange("port_policy") {
		v := d.Get("port_policy")
		p := make([]models.FabricPortPolicyRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsFabricPortPolicyRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPortPolicy(x)
		}
	}

	if d.HasChange("ports") {
		v := d.Get("ports")
		x := make([]models.FabricPortIdentifier, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.FabricPortIdentifier{}
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
			if v, ok := l["aggregate_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetAggregatePortId(x)
				}
			}
			o.SetClassId("fabric.PortIdentifier")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetPortId(x)
				}
			}
			if v, ok := l["slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetSlotId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPorts(x)
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
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("vsan_id") {
		v := d.Get("vsan_id")
		x := int64(v.(int))
		o.SetVsanId(x)
	}

	r := conn.ApiClient.FabricApi.UpdateFabricFcUplinkPcRole(conn.ctx, d.Id()).FabricFcUplinkPcRole(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating FabricFcUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceFabricFcUplinkPcRoleRead(c, d, meta)
}

func resourceFabricFcUplinkPcRoleDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.FabricApi.DeleteFabricFcUplinkPcRole(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting FabricFcUplinkPcRole object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
