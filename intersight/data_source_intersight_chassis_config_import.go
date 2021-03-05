package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChassisConfigImport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChassisConfigImportRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the imported profile.",
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
			"policy_prefix": {
				Description: "Policy prefix for the policies of the imported chassis profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"profile_name": {
				Description: "Profile name for the imported chassis profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceChassisConfigImport().Schema},
				Computed: true,
			}},
	}
}

func dataSourceChassisConfigImportRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ChassisConfigImport{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("policy_prefix"); ok {
		x := (v.(string))
		o.SetPolicyPrefix(x)
	}
	if v, ok := d.GetOk("profile_name"); ok {
		x := (v.(string))
		o.SetProfileName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ChassisConfigImport object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ChassisApi.GetChassisConfigImportList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ChassisConfigImport: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ChassisConfigImportList.GetCount()
	var i int32
	var chassisConfigImportResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ChassisApi.GetChassisConfigImportList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ChassisConfigImport: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ChassisConfigImportList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ChassisConfigImport data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["chassis"] = flattenMapEquipmentChassisRelationship(s.GetChassis(), d)

				temp["chassis_profile"] = flattenMapChassisProfileRelationship(s.GetChassisProfile(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["policy_prefix"] = (s.GetPolicyPrefix())
				temp["policy_types"] = (s.GetPolicyTypes())
				temp["profile_name"] = (s.GetProfileName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				chassisConfigImportResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(chassisConfigImportResults))
	if err := d.Set("results", chassisConfigImportResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(chassisConfigImportResults[0]["moid"].(string))
	return de
}
