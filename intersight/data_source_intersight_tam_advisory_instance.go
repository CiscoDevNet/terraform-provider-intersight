package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTamAdvisoryInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTamAdvisoryInstanceRead,
		Schema: map[string]*schema.Schema{
			"affected_object_moid": {
				Description: "Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"affected_object_type": {
				Description: "Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_state_change_time": {
				Description: "Timestamp when a state change was observed on this advisory instnace.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_verified_time": {
				Description: "Timestamp when this advisory was last evaluated.",
				Type:        schema.TypeString,
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
			"state": {
				Description: "Current state of the advisory instance (Active/Cleared/Unknown etc.).\n* `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object.\n* `active` - Advisory instance is currently active and applicable for the affected managed object.\n* `cleared` - Advisory instance is no longer applicable for the affected managed object.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceTamAdvisoryInstance().Schema},
				Computed: true,
			}},
	}
}

func dataSourceTamAdvisoryInstanceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TamAdvisoryInstance{}
	if v, ok := d.GetOk("affected_object_moid"); ok {
		x := (v.(string))
		o.SetAffectedObjectMoid(x)
	}
	if v, ok := d.GetOk("affected_object_type"); ok {
		x := (v.(string))
		o.SetAffectedObjectType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("last_state_change_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastStateChangeTime(x)
	}
	if v, ok := d.GetOk("last_verified_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastVerifiedTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of TamAdvisoryInstance object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.TamApi.GetTamAdvisoryInstanceList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of TamAdvisoryInstance: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.TamAdvisoryInstanceList.GetCount()
	var i int32
	var tamAdvisoryInstanceResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.TamApi.GetTamAdvisoryInstanceList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching TamAdvisoryInstance: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.TamAdvisoryInstanceList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for TamAdvisoryInstance data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["advisory"] = flattenMapTamBaseAdvisoryRelationship(s.GetAdvisory(), d)
				temp["affected_object_moid"] = (s.GetAffectedObjectMoid())
				temp["affected_object_type"] = (s.GetAffectedObjectType())
				temp["class_id"] = (s.GetClassId())

				temp["device_registration"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDeviceRegistration(), d)

				temp["last_state_change_time"] = (s.GetLastStateChangeTime()).String()

				temp["last_verified_time"] = (s.GetLastVerifiedTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["state"] = (s.GetState())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				tamAdvisoryInstanceResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(tamAdvisoryInstanceResults))
	if err := d.Set("results", tamAdvisoryInstanceResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(tamAdvisoryInstanceResults[0]["moid"].(string))
	return de
}
