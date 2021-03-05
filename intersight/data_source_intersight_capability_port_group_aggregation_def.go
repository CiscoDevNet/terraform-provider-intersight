package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCapabilityPortGroupAggregationDef() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilityPortGroupAggregationDefRead,
		Schema: map[string]*schema.Schema{
			"aggregation_cap": {
				Description: "Aggregation capability for port group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hw40_g_port_group_cap": {
				Description: "Indicates support for 40G port group capability.",
				Type:        schema.TypeBool,
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pgtype": {
				Description: "The type of port group for which this capability is defined.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCapabilityPortGroupAggregationDef().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCapabilityPortGroupAggregationDefRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilityPortGroupAggregationDef{}
	if v, ok := d.GetOk("aggregation_cap"); ok {
		x := (v.(string))
		o.SetAggregationCap(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("hw40_g_port_group_cap"); ok {
		x := (v.(bool))
		o.SetHw40GPortGroupCap(x)
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
	if v, ok := d.GetOk("pgtype"); ok {
		x := (v.(string))
		o.SetPgtype(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilityPortGroupAggregationDef object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityPortGroupAggregationDefList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CapabilityPortGroupAggregationDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CapabilityPortGroupAggregationDefList.GetCount()
	var i int32
	var capabilityPortGroupAggregationDefResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityPortGroupAggregationDefList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilityPortGroupAggregationDef: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CapabilityPortGroupAggregationDefList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CapabilityPortGroupAggregationDef data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["aggregation_cap"] = (s.GetAggregationCap())
				temp["class_id"] = (s.GetClassId())
				temp["hw40_g_port_group_cap"] = (s.GetHw40GPortGroupCap())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["pgtype"] = (s.GetPgtype())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				capabilityPortGroupAggregationDefResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilityPortGroupAggregationDefResults))
	if err := d.Set("results", capabilityPortGroupAggregationDefResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilityPortGroupAggregationDefResults[0]["moid"].(string))
	return de
}
