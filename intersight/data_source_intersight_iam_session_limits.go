package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamSessionLimits() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamSessionLimitsRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"idle_time_out": {
				Description: "The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"maximum_limit": {
				Description: "The maximum number of sessions allowed in an account. The default value is 128.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
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
			"per_user_limit": {
				Description: "The maximum number of sessions allowed per user. Default value is 32.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"session_time_out": {
				Description: "The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamSessionLimits().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamSessionLimitsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamSessionLimits{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("idle_time_out"); ok {
		x := int64(v.(int))
		o.SetIdleTimeOut(x)
	}
	if v, ok := d.GetOk("maximum_limit"); ok {
		x := int64(v.(int))
		o.SetMaximumLimit(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("per_user_limit"); ok {
		x := int64(v.(int))
		o.SetPerUserLimit(x)
	}
	if v, ok := d.GetOk("session_time_out"); ok {
		x := int64(v.(int))
		o.SetSessionTimeOut(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamSessionLimits object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamSessionLimitsList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamSessionLimits: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamSessionLimitsList.GetCount()
	var i int32
	var iamSessionLimitsResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamSessionLimitsList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamSessionLimits: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamSessionLimitsList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamSessionLimits data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["idle_time_out"] = (s.GetIdleTimeOut())
				temp["maximum_limit"] = (s.GetMaximumLimit())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["per_user_limit"] = (s.GetPerUserLimit())

				temp["permission"] = flattenMapIamPermissionRelationship(s.GetPermission(), d)
				temp["session_time_out"] = (s.GetSessionTimeOut())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iamSessionLimitsResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamSessionLimitsResults))
	if err := d.Set("results", iamSessionLimitsResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamSessionLimitsResults[0]["moid"].(string))
	return de
}
