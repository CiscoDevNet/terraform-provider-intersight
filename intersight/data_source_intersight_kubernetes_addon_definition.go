package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAddonDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAddonDefinitionRead,
		Schema: map[string]*schema.Schema{
			"chart_url": {
				Description: "Description of the addon component.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_install_strategy": {
				Description: "Default installation strategy for the release.\n* `InstallOnly` - Only install in green field. No action in case of failure or removal.\n* `NoAction` - No install action performed.\n* `Always` - Attempt install if chart is not already installed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_namespace": {
				Description: "Default namespace to install the release.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_upgrade_strategy": {
				Description: "Default upgrade strategy for the release.\n* `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure.\n* `NoAction` - This choice enables No upgrades to be performed.\n* `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure.\n* `AlwaysReinstall` - Always remove older release and reinstall.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the addon component.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"digest": {
				Description: "Digest used to verify the integrity of an addon.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"icon_url": {
				Description: "Icon used to represent the addon in UI.",
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
				Description: "Name of an addon component.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "Version of the addon component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAddonDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAddonDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAddonDefinition{}
	if v, ok := d.GetOk("chart_url"); ok {
		x := (v.(string))
		o.SetChartUrl(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_install_strategy"); ok {
		x := (v.(string))
		o.SetDefaultInstallStrategy(x)
	}
	if v, ok := d.GetOk("default_namespace"); ok {
		x := (v.(string))
		o.SetDefaultNamespace(x)
	}
	if v, ok := d.GetOk("default_upgrade_strategy"); ok {
		x := (v.(string))
		o.SetDefaultUpgradeStrategy(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("digest"); ok {
		x := (v.(string))
		o.SetDigest(x)
	}
	if v, ok := d.GetOk("icon_url"); ok {
		x := (v.(string))
		o.SetIconUrl(x)
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
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAddonDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAddonDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAddonDefinitionList.GetCount()
	var i int32
	var kubernetesAddonDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAddonDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAddonDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAddonDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAddonDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["catalog"] = flattenMapKubernetesCatalogRelationship(s.GetCatalog(), d)
				temp["chart_url"] = (s.GetChartUrl())
				temp["class_id"] = (s.GetClassId())
				temp["default_install_strategy"] = (s.GetDefaultInstallStrategy())
				temp["default_namespace"] = (s.GetDefaultNamespace())
				temp["default_upgrade_strategy"] = (s.GetDefaultUpgradeStrategy())
				temp["description"] = (s.GetDescription())
				temp["digest"] = (s.GetDigest())
				temp["icon_url"] = (s.GetIconUrl())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				kubernetesAddonDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAddonDefinitionResults))
	if err := d.Set("results", kubernetesAddonDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAddonDefinitionResults[0]["moid"].(string))
	return de
}
