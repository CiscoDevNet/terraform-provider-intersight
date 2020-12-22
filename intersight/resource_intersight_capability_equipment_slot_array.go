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

func resourceCapabilityEquipmentSlotArray() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCapabilityEquipmentSlotArrayCreate,
		ReadContext:   resourceCapabilityEquipmentSlotArrayRead,
		UpdateContext: resourceCapabilityEquipmentSlotArrayUpdate,
		DeleteContext: resourceCapabilityEquipmentSlotArrayDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
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
			"first_index": {
				Description: "First Index information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"height": {
				Description: "Height information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"horizontal_start_offset": {
				Description: "Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_group_separation": {
				Description: "Inline Group Separation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_group_size": {
				Description: "Inline Group Size information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_offset": {
				Description: "Inline Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"location": {
				Description: "Location information for a Switch/Fabric-Interconnect hardware.",
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
			"number_of_slots": {
				Description: "Number of Slots information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orientation": {
				Description: "Orientation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "UCS-FI-6454",
			},
			"selector": {
				Description: "Selector information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sku": {
				Description: "SKU information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slots_per_line": {
				Description: "Slots per line information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeInt,
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
			"transverse_group_separation": {
				Description: "Transverse Group Separation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"transverse_group_size": {
				Description: "Transverse Group Size information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"transverse_offset": {
				Description: "Transverse Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"vertical_start_offset": {
				Description: "Vertical Start Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"width": {
				Description: "Width information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
		},
	}
}

func resourceCapabilityEquipmentSlotArrayCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewCapabilityEquipmentSlotArrayWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("capability.EquipmentSlotArray")

	if v, ok := d.GetOk("first_index"); ok {
		x := v.(float32)
		o.SetFirstIndex(x)
	}

	if v, ok := d.GetOk("height"); ok {
		x := v.(float32)
		o.SetHeight(x)
	}

	if v, ok := d.GetOk("horizontal_start_offset"); ok {
		x := v.(float32)
		o.SetHorizontalStartOffset(x)
	}

	if v, ok := d.GetOk("inline_group_separation"); ok {
		x := v.(float32)
		o.SetInlineGroupSeparation(x)
	}

	if v, ok := d.GetOk("inline_group_size"); ok {
		x := v.(float32)
		o.SetInlineGroupSize(x)
	}

	if v, ok := d.GetOk("inline_offset"); ok {
		x := v.(float32)
		o.SetInlineOffset(x)
	}

	if v, ok := d.GetOk("location"); ok {
		x := (v.(string))
		o.SetLocation(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOk("number_of_slots"); ok {
		x := int64(v.(int))
		o.SetNumberOfSlots(x)
	}

	o.SetObjectType("capability.EquipmentSlotArray")

	if v, ok := d.GetOk("orientation"); ok {
		x := (v.(string))
		o.SetOrientation(x)
	}

	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}

	if v, ok := d.GetOk("selector"); ok {
		x := (v.(string))
		o.SetSelector(x)
	}

	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}

	if v, ok := d.GetOk("slots_per_line"); ok {
		x := int64(v.(int))
		o.SetSlotsPerLine(x)
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

	if v, ok := d.GetOk("transverse_group_separation"); ok {
		x := v.(float32)
		o.SetTransverseGroupSeparation(x)
	}

	if v, ok := d.GetOk("transverse_group_size"); ok {
		x := v.(float32)
		o.SetTransverseGroupSize(x)
	}

	if v, ok := d.GetOk("transverse_offset"); ok {
		x := v.(float32)
		o.SetTransverseOffset(x)
	}

	if v, ok := d.GetOk("vertical_start_offset"); ok {
		x := v.(float32)
		o.SetVerticalStartOffset(x)
	}

	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	if v, ok := d.GetOk("width"); ok {
		x := v.(float32)
		o.SetWidth(x)
	}

	r := conn.ApiClient.CapabilityApi.CreateCapabilityEquipmentSlotArray(conn.ctx).CapabilityEquipmentSlotArray(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating CapabilityEquipmentSlotArray: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceCapabilityEquipmentSlotArrayRead(c, d, meta)
}

func resourceCapabilityEquipmentSlotArrayRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentSlotArrayByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilityEquipmentSlotArray object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching CapabilityEquipmentSlotArray: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("first_index", (s.GetFirstIndex())); err != nil {
		return diag.Errorf("error occurred while setting property FirstIndex in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("height", (s.GetHeight())); err != nil {
		return diag.Errorf("error occurred while setting property Height in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("horizontal_start_offset", (s.GetHorizontalStartOffset())); err != nil {
		return diag.Errorf("error occurred while setting property HorizontalStartOffset in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("inline_group_separation", (s.GetInlineGroupSeparation())); err != nil {
		return diag.Errorf("error occurred while setting property InlineGroupSeparation in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("inline_group_size", (s.GetInlineGroupSize())); err != nil {
		return diag.Errorf("error occurred while setting property InlineGroupSize in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("inline_offset", (s.GetInlineOffset())); err != nil {
		return diag.Errorf("error occurred while setting property InlineOffset in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("location", (s.GetLocation())); err != nil {
		return diag.Errorf("error occurred while setting property Location in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("number_of_slots", (s.GetNumberOfSlots())); err != nil {
		return diag.Errorf("error occurred while setting property NumberOfSlots in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("orientation", (s.GetOrientation())); err != nil {
		return diag.Errorf("error occurred while setting property Orientation in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("pid", (s.GetPid())); err != nil {
		return diag.Errorf("error occurred while setting property Pid in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("selector", (s.GetSelector())); err != nil {
		return diag.Errorf("error occurred while setting property Selector in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("sku", (s.GetSku())); err != nil {
		return diag.Errorf("error occurred while setting property Sku in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("slots_per_line", (s.GetSlotsPerLine())); err != nil {
		return diag.Errorf("error occurred while setting property SlotsPerLine in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("transverse_group_separation", (s.GetTransverseGroupSeparation())); err != nil {
		return diag.Errorf("error occurred while setting property TransverseGroupSeparation in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("transverse_group_size", (s.GetTransverseGroupSize())); err != nil {
		return diag.Errorf("error occurred while setting property TransverseGroupSize in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("transverse_offset", (s.GetTransverseOffset())); err != nil {
		return diag.Errorf("error occurred while setting property TransverseOffset in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("vertical_start_offset", (s.GetVerticalStartOffset())); err != nil {
		return diag.Errorf("error occurred while setting property VerticalStartOffset in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("vid", (s.GetVid())); err != nil {
		return diag.Errorf("error occurred while setting property Vid in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	if err := d.Set("width", (s.GetWidth())); err != nil {
		return diag.Errorf("error occurred while setting property Width in CapabilityEquipmentSlotArray object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceCapabilityEquipmentSlotArrayUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.CapabilityEquipmentSlotArray{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("capability.EquipmentSlotArray")

	if d.HasChange("first_index") {
		v := d.Get("first_index")
		x := v.(float32)
		o.SetFirstIndex(x)
	}

	if d.HasChange("height") {
		v := d.Get("height")
		x := v.(float32)
		o.SetHeight(x)
	}

	if d.HasChange("horizontal_start_offset") {
		v := d.Get("horizontal_start_offset")
		x := v.(float32)
		o.SetHorizontalStartOffset(x)
	}

	if d.HasChange("inline_group_separation") {
		v := d.Get("inline_group_separation")
		x := v.(float32)
		o.SetInlineGroupSeparation(x)
	}

	if d.HasChange("inline_group_size") {
		v := d.Get("inline_group_size")
		x := v.(float32)
		o.SetInlineGroupSize(x)
	}

	if d.HasChange("inline_offset") {
		v := d.Get("inline_offset")
		x := v.(float32)
		o.SetInlineOffset(x)
	}

	if d.HasChange("location") {
		v := d.Get("location")
		x := (v.(string))
		o.SetLocation(x)
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

	if d.HasChange("number_of_slots") {
		v := d.Get("number_of_slots")
		x := int64(v.(int))
		o.SetNumberOfSlots(x)
	}

	o.SetObjectType("capability.EquipmentSlotArray")

	if d.HasChange("orientation") {
		v := d.Get("orientation")
		x := (v.(string))
		o.SetOrientation(x)
	}

	if d.HasChange("pid") {
		v := d.Get("pid")
		x := (v.(string))
		o.SetPid(x)
	}

	if d.HasChange("selector") {
		v := d.Get("selector")
		x := (v.(string))
		o.SetSelector(x)
	}

	if d.HasChange("sku") {
		v := d.Get("sku")
		x := (v.(string))
		o.SetSku(x)
	}

	if d.HasChange("slots_per_line") {
		v := d.Get("slots_per_line")
		x := int64(v.(int))
		o.SetSlotsPerLine(x)
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

	if d.HasChange("transverse_group_separation") {
		v := d.Get("transverse_group_separation")
		x := v.(float32)
		o.SetTransverseGroupSeparation(x)
	}

	if d.HasChange("transverse_group_size") {
		v := d.Get("transverse_group_size")
		x := v.(float32)
		o.SetTransverseGroupSize(x)
	}

	if d.HasChange("transverse_offset") {
		v := d.Get("transverse_offset")
		x := v.(float32)
		o.SetTransverseOffset(x)
	}

	if d.HasChange("vertical_start_offset") {
		v := d.Get("vertical_start_offset")
		x := v.(float32)
		o.SetVerticalStartOffset(x)
	}

	if d.HasChange("vid") {
		v := d.Get("vid")
		x := (v.(string))
		o.SetVid(x)
	}

	if d.HasChange("width") {
		v := d.Get("width")
		x := v.(float32)
		o.SetWidth(x)
	}

	r := conn.ApiClient.CapabilityApi.UpdateCapabilityEquipmentSlotArray(conn.ctx, d.Id()).CapabilityEquipmentSlotArray(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating CapabilityEquipmentSlotArray: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceCapabilityEquipmentSlotArrayRead(c, d, meta)
}

func resourceCapabilityEquipmentSlotArrayDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.CapabilityApi.DeleteCapabilityEquipmentSlotArray(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting CapabilityEquipmentSlotArray object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
