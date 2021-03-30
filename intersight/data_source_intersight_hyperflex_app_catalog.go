package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexAppCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexAppCatalogRead,
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
			"nr_version": {
				Description: "The catalog version used in HyperFlex cluster configuration service.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexAppCatalog().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexAppCatalogRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexAppCatalog{}
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
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexAppCatalog object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexAppCatalogList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexAppCatalog: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexAppCatalogList.GetCount()
	var i int32
	var hyperflexAppCatalogResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexAppCatalogList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexAppCatalog: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexAppCatalogList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexAppCatalog data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["feature_limit_external"] = flattenMapHyperflexFeatureLimitExternalRelationship(s.GetFeatureLimitExternal(), d)

				temp["feature_limit_internal"] = flattenMapHyperflexFeatureLimitInternalRelationship(s.GetFeatureLimitInternal(), d)

				temp["hxdp_versions"] = flattenListHyperflexHxdpVersionRelationship(s.GetHxdpVersions(), d)

				temp["hyperflex_capability_infos"] = flattenListHyperflexCapabilityInfoRelationship(s.GetHyperflexCapabilityInfos(), d)

				temp["hyperflex_software_compatibility_infos"] = flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(s.GetHyperflexSoftwareCompatibilityInfos(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["server_firmware_version"] = flattenMapHyperflexServerFirmwareVersionRelationship(s.GetServerFirmwareVersion(), d)

				temp["server_model"] = flattenMapHyperflexServerModelRelationship(s.GetServerModel(), d)

				temp["software_distributions"] = flattenListHyperflexSoftwareDistributionEntryRelationship(s.GetSoftwareDistributions(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				hyperflexAppCatalogResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexAppCatalogResults))
	if err := d.Set("results", hyperflexAppCatalogResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexAppCatalogResults[0]["moid"].(string))
	return de
}
