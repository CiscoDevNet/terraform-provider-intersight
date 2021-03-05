package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAddon() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAddonRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"install_strategy": {
				Description: "Addon install strategy to determine whether an addon is installed if not present.\n* `InstallOnly` - Only install in green field. No action in case of failure or removal.\n* `NoAction` - No install action performed.\n* `Always` - Attempt install if chart is not already installed.",
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
				Description: "Name of addon to be installed on a Kubernetes cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"overrides": {
				Description: "Properties that can be overridden for an addon.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"upgrade_strategy": {
				Description: "Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade.\n* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.\n* `NoAction` - This choice enables No upgrades to be performed.\n* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.\n* `AlwaysReinstall` - Always remove older release and reinstall.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAddon().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAddonRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAddon{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("install_strategy"); ok {
		x := (v.(string))
		o.SetInstallStrategy(x)
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
	if v, ok := d.GetOk("overrides"); ok {
		x := (v.(string))
		o.SetOverrides(x)
	}
	if v, ok := d.GetOk("upgrade_strategy"); ok {
		x := (v.(string))
		o.SetUpgradeStrategy(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAddon object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAddon: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAddonList.GetCount()
	var i int32
	var kubernetesAddonResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAddon: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAddonList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAddon data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["addon_definition"] = flattenMapKubernetesAddonDefinitionRelationship(s.GetAddonDefinition(), d)
				temp["class_id"] = (s.GetClassId())
				temp["install_strategy"] = (s.GetInstallStrategy())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["override_sets"] = flattenListKubernetesKeyValue(s.GetOverrideSets(), d)
				temp["overrides"] = (s.GetOverrides())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["upgrade_strategy"] = (s.GetUpgradeStrategy())
				kubernetesAddonResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAddonResults))
	if err := d.Set("results", kubernetesAddonResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAddonResults[0]["moid"].(string))
	return de
}
