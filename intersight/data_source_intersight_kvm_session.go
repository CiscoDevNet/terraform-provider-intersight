package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKvmSession() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKvmSessionRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"one_time_password": {
				Description: "Temporary one-time password for KVM access.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sso_supported": {
				Description: "Indicates if KVM SSO is supported on the server.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Description: "Username used for KVM access.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKvmSession().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKvmSessionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KvmSession{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("one_time_password"); ok {
		x := (v.(string))
		o.SetOneTimePassword(x)
	}
	if v, ok := d.GetOk("sso_supported"); ok {
		x := (v.(bool))
		o.SetSsoSupported(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KvmSession object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KvmApi.GetKvmSessionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KvmSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KvmSessionList.GetCount()
	var i int32
	var kvmSessionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KvmApi.GetKvmSessionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KvmSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KvmSessionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KvmSession data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["one_time_password"] = (s.GetOneTimePassword())

				temp["server"] = flattenMapComputePhysicalRelationship(s.GetServer(), d)

				temp["session"] = flattenMapIamSessionRelationship(s.GetSession(), d)
				temp["sso_supported"] = (s.GetSsoSupported())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tunnel"] = flattenMapKvmTunnelRelationship(s.GetTunnel(), d)
				temp["username"] = (s.GetUsername())
				kvmSessionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kvmSessionResults))
	if err := d.Set("results", kvmSessionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kvmSessionResults[0]["moid"].(string))
	return de
}
