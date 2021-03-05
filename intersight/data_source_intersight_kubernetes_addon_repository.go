package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAddonRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAddonRepositoryRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"insecure_skip_verification": {
				Description: "Allow connecting to http registries or https registries which do not have certificate signed by a well known CA.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"is_token_set": {
				Description: "Indicates whether the value of the 'token' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the addon repository or registry.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"repository_url": {
				Description: "URL for the repository where the addon is hosted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"username": {
				Description: "Username to authenticate to the addon registry.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAddonRepository().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAddonRepositoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAddonRepository{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("insecure_skip_verification"); ok {
		x := (v.(bool))
		o.SetInsecureSkipVerification(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("is_token_set"); ok {
		x := (v.(bool))
		o.SetIsTokenSet(x)
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
	if v, ok := d.GetOk("repository_url"); ok {
		x := (v.(string))
		o.SetRepositoryUrl(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAddonRepository object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonRepositoryList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAddonRepository: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAddonRepositoryList.GetCount()
	var i int32
	var kubernetesAddonRepositoryResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonRepositoryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAddonRepository: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAddonRepositoryList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAddonRepository data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ca_cert"] = flattenMapX509Certificate(s.GetCaCert(), d)

				temp["catalog"] = flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["insecure_skip_verification"] = (s.GetInsecureSkipVerification())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["is_token_set"] = (s.GetIsTokenSet())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["repository_url"] = (s.GetRepositoryUrl())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["username"] = (s.GetUsername())
				kubernetesAddonRepositoryResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAddonRepositoryResults))
	if err := d.Set("results", kubernetesAddonRepositoryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAddonRepositoryResults[0]["moid"].(string))
	return de
}
