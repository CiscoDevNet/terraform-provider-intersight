package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCapabilityEquipmentSlotArray() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCapabilityEquipmentSlotArrayRead,
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

func dataSourceCapabilityEquipmentSlotArrayRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.CapabilityEquipmentSlotArray{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
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

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentSlotArrayList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.CapabilityEquipmentSlotArrayList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to CapabilityEquipmentSlotArray: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.CapabilityEquipmentSlotArray{}
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
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
