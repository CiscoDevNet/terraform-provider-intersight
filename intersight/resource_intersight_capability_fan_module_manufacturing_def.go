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

func resourceCapabilityFanModuleManufacturingDef() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCapabilityFanModuleManufacturingDefCreate,
		ReadContext:   resourceCapabilityFanModuleManufacturingDefRead,
		UpdateContext: resourceCapabilityFanModuleManufacturingDefUpdate,
		DeleteContext: resourceCapabilityFanModuleManufacturingDefDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"caption": {
				Description: "Caption for a fan module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description for a fan module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product Identifier for a fan module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_name": {
				Description: "Product Name for Fan Module Unit.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sku": {
				Description: "SKU information for a fan module.",
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
			"vid": {
				Description: "VID information for a fan module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceCapabilityFanModuleManufacturingDefCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewCapabilityFanModuleManufacturingDefWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("caption"); ok {
		x := (v.(string))
		o.SetCaption(x)
	}

	o.SetClassId("capability.FanModuleManufacturingDef")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("capability.FanModuleManufacturingDef")

	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}

	if v, ok := d.GetOk("product_name"); ok {
		x := (v.(string))
		o.SetProductName(x)
	}

	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
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

	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	r := conn.ApiClient.CapabilityApi.CreateCapabilityFanModuleManufacturingDef(conn.ctx).CapabilityFanModuleManufacturingDef(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating CapabilityFanModuleManufacturingDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceCapabilityFanModuleManufacturingDefRead(c, d, meta)
}

func resourceCapabilityFanModuleManufacturingDefRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.CapabilityApi.GetCapabilityFanModuleManufacturingDefByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilityFanModuleManufacturingDef object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching CapabilityFanModuleManufacturingDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("caption", (s.GetCaption())); err != nil {
		return diag.Errorf("error occurred while setting property Caption in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("pid", (s.GetPid())); err != nil {
		return diag.Errorf("error occurred while setting property Pid in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("product_name", (s.GetProductName())); err != nil {
		return diag.Errorf("error occurred while setting property ProductName in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("sku", (s.GetSku())); err != nil {
		return diag.Errorf("error occurred while setting property Sku in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	if err := d.Set("vid", (s.GetVid())); err != nil {
		return diag.Errorf("error occurred while setting property Vid in CapabilityFanModuleManufacturingDef object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceCapabilityFanModuleManufacturingDefUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.CapabilityFanModuleManufacturingDef{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("caption") {
		v := d.Get("caption")
		x := (v.(string))
		o.SetCaption(x)
	}

	o.SetClassId("capability.FanModuleManufacturingDef")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
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

	o.SetObjectType("capability.FanModuleManufacturingDef")

	if d.HasChange("pid") {
		v := d.Get("pid")
		x := (v.(string))
		o.SetPid(x)
	}

	if d.HasChange("product_name") {
		v := d.Get("product_name")
		x := (v.(string))
		o.SetProductName(x)
	}

	if d.HasChange("sku") {
		v := d.Get("sku")
		x := (v.(string))
		o.SetSku(x)
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

	if d.HasChange("vid") {
		v := d.Get("vid")
		x := (v.(string))
		o.SetVid(x)
	}

	r := conn.ApiClient.CapabilityApi.UpdateCapabilityFanModuleManufacturingDef(conn.ctx, d.Id()).CapabilityFanModuleManufacturingDef(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating CapabilityFanModuleManufacturingDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceCapabilityFanModuleManufacturingDefRead(c, d, meta)
}

func resourceCapabilityFanModuleManufacturingDefDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.CapabilityApi.DeleteCapabilityFanModuleManufacturingDef(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting CapabilityFanModuleManufacturingDef object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
