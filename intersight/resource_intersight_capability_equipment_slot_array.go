package intersight

import (
	"encoding/json"
	"fmt"
	"log"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCapabilityEquipmentSlotArray() *schema.Resource {
	return &schema.Resource{
		Create: resourceCapabilityEquipmentSlotArrayCreate,
		Read:   resourceCapabilityEquipmentSlotArrayRead,
		Update: resourceCapabilityEquipmentSlotArrayUpdate,
		Delete: resourceCapabilityEquipmentSlotArrayDelete,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
				Description: "Product Identifier for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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

func resourceCapabilityEquipmentSlotArrayCreate(d *schema.ResourceData, meta interface{}) error {
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

	if v, ok := d.GetOk("section"); ok {
		p := make([]models.CapabilitySectionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsCapabilitySectionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSection(x)
		}
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
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceCapabilityEquipmentSlotArrayRead(d, meta)
}

func resourceCapabilityEquipmentSlotArrayRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentSlotArrayByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("first_index", (s.GetFirstIndex())); err != nil {
		return fmt.Errorf("error occurred while setting property FirstIndex: %+v", err)
	}

	if err := d.Set("height", (s.GetHeight())); err != nil {
		return fmt.Errorf("error occurred while setting property Height: %+v", err)
	}

	if err := d.Set("horizontal_start_offset", (s.GetHorizontalStartOffset())); err != nil {
		return fmt.Errorf("error occurred while setting property HorizontalStartOffset: %+v", err)
	}

	if err := d.Set("inline_group_separation", (s.GetInlineGroupSeparation())); err != nil {
		return fmt.Errorf("error occurred while setting property InlineGroupSeparation: %+v", err)
	}

	if err := d.Set("inline_group_size", (s.GetInlineGroupSize())); err != nil {
		return fmt.Errorf("error occurred while setting property InlineGroupSize: %+v", err)
	}

	if err := d.Set("inline_offset", (s.GetInlineOffset())); err != nil {
		return fmt.Errorf("error occurred while setting property InlineOffset: %+v", err)
	}

	if err := d.Set("location", (s.GetLocation())); err != nil {
		return fmt.Errorf("error occurred while setting property Location: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return fmt.Errorf("error occurred while setting property Name: %+v", err)
	}

	if err := d.Set("number_of_slots", (s.GetNumberOfSlots())); err != nil {
		return fmt.Errorf("error occurred while setting property NumberOfSlots: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("orientation", (s.GetOrientation())); err != nil {
		return fmt.Errorf("error occurred while setting property Orientation: %+v", err)
	}

	if err := d.Set("pid", (s.GetPid())); err != nil {
		return fmt.Errorf("error occurred while setting property Pid: %+v", err)
	}

	if err := d.Set("section", flattenMapCapabilitySectionRelationship(s.GetSection(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Section: %+v", err)
	}

	if err := d.Set("selector", (s.GetSelector())); err != nil {
		return fmt.Errorf("error occurred while setting property Selector: %+v", err)
	}

	if err := d.Set("sku", (s.GetSku())); err != nil {
		return fmt.Errorf("error occurred while setting property Sku: %+v", err)
	}

	if err := d.Set("slots_per_line", (s.GetSlotsPerLine())); err != nil {
		return fmt.Errorf("error occurred while setting property SlotsPerLine: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	if err := d.Set("transverse_group_separation", (s.GetTransverseGroupSeparation())); err != nil {
		return fmt.Errorf("error occurred while setting property TransverseGroupSeparation: %+v", err)
	}

	if err := d.Set("transverse_group_size", (s.GetTransverseGroupSize())); err != nil {
		return fmt.Errorf("error occurred while setting property TransverseGroupSize: %+v", err)
	}

	if err := d.Set("transverse_offset", (s.GetTransverseOffset())); err != nil {
		return fmt.Errorf("error occurred while setting property TransverseOffset: %+v", err)
	}

	if err := d.Set("vertical_start_offset", (s.GetVerticalStartOffset())); err != nil {
		return fmt.Errorf("error occurred while setting property VerticalStartOffset: %+v", err)
	}

	if err := d.Set("vid", (s.GetVid())); err != nil {
		return fmt.Errorf("error occurred while setting property Vid: %+v", err)
	}

	if err := d.Set("width", (s.GetWidth())); err != nil {
		return fmt.Errorf("error occurred while setting property Width: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceCapabilityEquipmentSlotArrayUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewCapabilityEquipmentSlotArrayWithDefaults()
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

	if d.HasChange("section") {
		v := d.Get("section")
		p := make([]models.CapabilitySectionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsCapabilitySectionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSection(x)
		}
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
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceCapabilityEquipmentSlotArrayRead(d, meta)
}

func resourceCapabilityEquipmentSlotArrayDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	p := conn.ApiClient.CapabilityApi.DeleteCapabilityEquipmentSlotArray(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
