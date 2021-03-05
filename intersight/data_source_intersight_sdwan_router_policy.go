package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdwanRouterPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSdwanRouterPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployment_size": {
				Description: "Scale of the SD-WAN router virtual machine deployment.\n* `Typical` - Typical deployment configuration with 4 vCPUs and 4GB RAM.\n* `Minimal` - Minimal deployment configuration with 2 vCPUs and 4GB RAM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
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
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wan_count": {
				Description: "Number of WAN connections across the SD-WAN site.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"wan_termination_type": {
				Description: "Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers.\n* `Single` - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes.\n* `Dual` - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSdwanRouterPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSdwanRouterPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SdwanRouterPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("deployment_size"); ok {
		x := (v.(string))
		o.SetDeploymentSize(x)
	}
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("wan_count"); ok {
		x := int64(v.(int))
		o.SetWanCount(x)
	}
	if v, ok := d.GetOk("wan_termination_type"); ok {
		x := (v.(string))
		o.SetWanTerminationType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SdwanRouterPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SdwanApi.GetSdwanRouterPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SdwanRouterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SdwanRouterPolicyList.GetCount()
	var i int32
	var sdwanRouterPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SdwanApi.GetSdwanRouterPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SdwanRouterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SdwanRouterPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SdwanRouterPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["deployment_size"] = (s.GetDeploymentSize())
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListSdwanProfileRelationship(s.GetProfiles(), d)

				temp["solution_image"] = flattenMapSoftwareSolutionDistributableRelationship(s.GetSolutionImage(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["wan_count"] = (s.GetWanCount())
				temp["wan_termination_type"] = (s.GetWanTerminationType())
				sdwanRouterPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(sdwanRouterPolicyResults))
	if err := d.Set("results", sdwanRouterPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(sdwanRouterPolicyResults[0]["moid"].(string))
	return de
}
