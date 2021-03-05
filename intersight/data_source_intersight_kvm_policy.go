package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKvmPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKvmPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_local_server_video": {
				Description: "If enabled, displays KVM session on any monitor attached to the server.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_video_encryption": {
				Description: "If enabled, encrypts all video information sent through KVM.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enabled": {
				Description: "State of the vKVM service on the endpoint.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"maximum_sessions": {
				Description: "The maximum number of concurrent KVM sessions allowed.",
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
			"remote_port": {
				Description: "The port used for KVM communication.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKvmPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKvmPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KvmPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enable_local_server_video"); ok {
		x := (v.(bool))
		o.SetEnableLocalServerVideo(x)
	}
	if v, ok := d.GetOk("enable_video_encryption"); ok {
		x := (v.(bool))
		o.SetEnableVideoEncryption(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("maximum_sessions"); ok {
		x := int64(v.(int))
		o.SetMaximumSessions(x)
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
	if v, ok := d.GetOk("remote_port"); ok {
		x := int64(v.(int))
		o.SetRemotePort(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KvmPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KvmApi.GetKvmPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KvmPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KvmPolicyList.GetCount()
	var i int32
	var kvmPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KvmApi.GetKvmPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KvmPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KvmPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KvmPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["enable_local_server_video"] = (s.GetEnableLocalServerVideo())
				temp["enable_video_encryption"] = (s.GetEnableVideoEncryption())
				temp["enabled"] = (s.GetEnabled())
				temp["maximum_sessions"] = (s.GetMaximumSessions())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["remote_port"] = (s.GetRemotePort())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				kvmPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kvmPolicyResults))
	if err := d.Set("results", kvmPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kvmPolicyResults[0]["moid"].(string))
	return de
}
