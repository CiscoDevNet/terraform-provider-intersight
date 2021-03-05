package intersight

import (
	"context"
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"sku": {
				Description: "SKU information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCapabilityEquipmentPhysicalDef().Schema},
				Computed: true,
			}},
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
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentPhysicalDefList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CapabilityEquipmentPhysicalDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CapabilityEquipmentPhysicalDefList.GetCount()
	var i int32
	var capabilityEquipmentPhysicalDefResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentPhysicalDefList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilityEquipmentPhysicalDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CapabilityEquipmentPhysicalDefList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CapabilityEquipmentPhysicalDef data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["depth"] = (s.GetDepth())
				temp["height"] = (s.GetHeight())
				temp["max_power"] = (s.GetMaxPower())
				temp["min_power"] = (s.GetMinPower())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["nominal_power"] = (s.GetNominalPower())
				temp["object_type"] = (s.GetObjectType())
				temp["pid"] = (s.GetPid())
				temp["sku"] = (s.GetSku())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vid"] = (s.GetVid())
				temp["weight"] = (s.GetWeight())
				temp["width"] = (s.GetWidth())
				capabilityEquipmentPhysicalDefResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilityEquipmentPhysicalDefResults))
	if err := d.Set("results", capabilityEquipmentPhysicalDefResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilityEquipmentPhysicalDefResults[0]["moid"].(string))
	return de
}
