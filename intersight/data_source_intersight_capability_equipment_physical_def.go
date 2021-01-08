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

func dataSourceCapabilityEquipmentPhysicalDef() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilityEquipmentPhysicalDefRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"depth": {
				Description: "Depth information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"height": {
				Description: "Height information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_power": {
				Description: "Max Power information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_power": {
				Description: "Min Power information for a Switch/Fabric-Interconnect.",
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
			"nominal_power": {
				Description: "Nominal Power information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
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
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"weight": {
				Description: "Weight information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"width": {
				Description: "Width information for a Switch/Fabric-Interconnect.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
		},
	}
}

func dataSourceCapabilityEquipmentPhysicalDefRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilityEquipmentPhysicalDef{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("depth"); ok {
		x := v.(float32)
		o.SetDepth(x)
	}
	if v, ok := d.GetOk("height"); ok {
		x := v.(float32)
		o.SetHeight(x)
	}
	if v, ok := d.GetOk("max_power"); ok {
		x := int64(v.(int))
		o.SetMaxPower(x)
	}
	if v, ok := d.GetOk("min_power"); ok {
		x := int64(v.(int))
		o.SetMinPower(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("nominal_power"); ok {
		x := int64(v.(int))
		o.SetNominalPower(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}
	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}
	if v, ok := d.GetOk("weight"); ok {
		x := v.(float32)
		o.SetWeight(x)
	}
	if v, ok := d.GetOk("width"); ok {
		x := v.(float32)
		o.SetWidth(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilityEquipmentPhysicalDef object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentPhysicalDefList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching CapabilityEquipmentPhysicalDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for CapabilityEquipmentPhysicalDef list: %s", err.Error())
	}
	var s = &models.CapabilityEquipmentPhysicalDefList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to CapabilityEquipmentPhysicalDef list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for CapabilityEquipmentPhysicalDef did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.CapabilityEquipmentPhysicalDef{}
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
			if err := d.Set("depth", (s.GetDepth())); err != nil {
				return diag.Errorf("error occurred while setting property Depth: %s", err.Error())
			}
			if err := d.Set("height", (s.GetHeight())); err != nil {
				return diag.Errorf("error occurred while setting property Height: %s", err.Error())
			}
			if err := d.Set("max_power", (s.GetMaxPower())); err != nil {
				return diag.Errorf("error occurred while setting property MaxPower: %s", err.Error())
			}
			if err := d.Set("min_power", (s.GetMinPower())); err != nil {
				return diag.Errorf("error occurred while setting property MinPower: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("nominal_power", (s.GetNominalPower())); err != nil {
				return diag.Errorf("error occurred while setting property NominalPower: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("pid", (s.GetPid())); err != nil {
				return diag.Errorf("error occurred while setting property Pid: %s", err.Error())
			}
			if err := d.Set("sku", (s.GetSku())); err != nil {
				return diag.Errorf("error occurred while setting property Sku: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("vid", (s.GetVid())); err != nil {
				return diag.Errorf("error occurred while setting property Vid: %s", err.Error())
			}
			if err := d.Set("weight", (s.GetWeight())); err != nil {
				return diag.Errorf("error occurred while setting property Weight: %s", err.Error())
			}
			if err := d.Set("width", (s.GetWidth())); err != nil {
				return diag.Errorf("error occurred while setting property Width: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
