package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexNodeProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexNodeProfileRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hxdp_data_ip": {
				Description: "IP address for storage data network (Controller VM interface).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hxdp_mgmt_ip": {
				Description: "IP address for HyperFlex management network.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_control_ip": {
				Description: "IP address for hypervisor control such as VM migration or pod management.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_data_ip": {
				Description: "IP address for storage data network (Hypervisor interface).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_mgmt_ip": {
				Description: "IP address for Hypervisor management network.",
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
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexNodeProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexNodeProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexNodeProfile{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("hxdp_data_ip"); ok {
		x := (v.(string))
		o.SetHxdpDataIp(x)
	}
	if v, ok := d.GetOk("hxdp_mgmt_ip"); ok {
		x := (v.(string))
		o.SetHxdpMgmtIp(x)
	}
	if v, ok := d.GetOk("hypervisor_control_ip"); ok {
		x := (v.(string))
		o.SetHypervisorControlIp(x)
	}
	if v, ok := d.GetOk("hypervisor_data_ip"); ok {
		x := (v.(string))
		o.SetHypervisorDataIp(x)
	}
	if v, ok := d.GetOk("hypervisor_mgmt_ip"); ok {
		x := (v.(string))
		o.SetHypervisorMgmtIp(x)
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
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexNodeProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexNodeProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexNodeProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexNodeProfileList.GetCount()
	var i int32
	var hyperflexNodeProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexNodeProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexNodeProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexNodeProfileList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexNodeProfile data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["assigned_server"] = flattenMapComputePhysicalRelationship(s.GetAssignedServer(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profile"] = flattenMapHyperflexClusterProfileRelationship(s.GetClusterProfile(), d)
				temp["description"] = (s.GetDescription())
				temp["hxdp_data_ip"] = (s.GetHxdpDataIp())
				temp["hxdp_mgmt_ip"] = (s.GetHxdpMgmtIp())
				temp["hypervisor_control_ip"] = (s.GetHypervisorControlIp())
				temp["hypervisor_data_ip"] = (s.GetHypervisorDataIp())
				temp["hypervisor_mgmt_ip"] = (s.GetHypervisorMgmtIp())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				hyperflexNodeProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexNodeProfileResults))
	if err := d.Set("results", hyperflexNodeProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexNodeProfileResults[0]["moid"].(string))
	return de
}
