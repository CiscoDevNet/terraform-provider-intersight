package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricPortMode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricPortModeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"custom_mode": {
				Description: "Custom Port Mode specified for the port range.\n* `FibreChannel` - Fibre Channel Port Types.\n* `BreakoutEthernet10G` - Breakout Ethernet 10G Port Type.\n* `BreakoutEthernet25G` - Breakout Ethernet 25G Port Type.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port_id_end": {
				Description: "Ending range of the Port Identifier.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"port_id_start": {
				Description: "Starting range of the Port Identifier.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"slot_id": {
				Description: "Slot Identifier of the switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricPortMode().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricPortModeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricPortMode{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("custom_mode"); ok {
		x := (v.(string))
		o.SetCustomMode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("port_id_end"); ok {
		x := int64(v.(int))
		o.SetPortIdEnd(x)
	}
	if v, ok := d.GetOk("port_id_start"); ok {
		x := int64(v.(int))
		o.SetPortIdStart(x)
	}
	if v, ok := d.GetOk("slot_id"); ok {
		x := int64(v.(int))
		o.SetSlotId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricPortMode object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricPortModeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricPortMode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricPortModeList.GetCount()
	var i int32
	var fabricPortModeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricPortModeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricPortMode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricPortModeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricPortMode data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["custom_mode"] = (s.GetCustomMode())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["port_id_end"] = (s.GetPortIdEnd())
				temp["port_id_start"] = (s.GetPortIdStart())

				temp["port_policy"] = flattenMapFabricPortPolicyRelationship(s.GetPortPolicy(), d)
				temp["slot_id"] = (s.GetSlotId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				fabricPortModeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricPortModeResults))
	if err := d.Set("results", fabricPortModeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricPortModeResults[0]["moid"].(string))
	return de
}
