package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricVsan() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricVsanRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_zoning": {
				Description: "Enables or Disables the default zoning state.\n* `Enabled` - Admin configured Enabled State.\n* `Disabled` - Admin configured Disabled State.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fc_zone_sharing_mode": {
				Description: "Logical grouping mode for fc ports.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fcoe_vlan": {
				Description: "FCOE Vlan associated to the VSAN configuration.",
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
				Description: "User given name for the VSAN configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vsan_id": {
				Description: "Virtual San Identifier in the switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricVsan().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricVsanRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricVsan{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_zoning"); ok {
		x := (v.(string))
		o.SetDefaultZoning(x)
	}
	if v, ok := d.GetOk("fc_zone_sharing_mode"); ok {
		x := (v.(string))
		o.SetFcZoneSharingMode(x)
	}
	if v, ok := d.GetOk("fcoe_vlan"); ok {
		x := int64(v.(int))
		o.SetFcoeVlan(x)
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
	if v, ok := d.GetOk("vsan_id"); ok {
		x := int64(v.(int))
		o.SetVsanId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricVsan object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricVsanList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricVsan: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricVsanList.GetCount()
	var i int32
	var fabricVsanResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricVsanList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricVsan: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricVsanList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricVsan data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["default_zoning"] = (s.GetDefaultZoning())

				temp["fc_network_policy"] = flattenMapFabricFcNetworkPolicyRelationship(s.GetFcNetworkPolicy(), d)
				temp["fc_zone_sharing_mode"] = (s.GetFcZoneSharingMode())
				temp["fcoe_vlan"] = (s.GetFcoeVlan())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vsan_id"] = (s.GetVsanId())
				fabricVsanResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricVsanResults))
	if err := d.Set("results", fabricVsanResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricVsanResults[0]["moid"].(string))
	return de
}
