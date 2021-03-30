package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexSoftwareDistributionVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexSoftwareDistributionVersionRead,
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
				Description: "The HyperFlex Software Distribution version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexSoftwareDistributionVersion().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexSoftwareDistributionVersionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexSoftwareDistributionVersion{}
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
		return diag.Errorf("json marshal of HyperflexSoftwareDistributionVersion object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareDistributionVersionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexSoftwareDistributionVersion: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexSoftwareDistributionVersionList.GetCount()
	var i int32
	var hyperflexSoftwareDistributionVersionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareDistributionVersionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexSoftwareDistributionVersion: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexSoftwareDistributionVersionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexSoftwareDistributionVersion data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["distribution_components"] = flattenListHyperflexSoftwareDistributionComponentRelationship(s.GetDistributionComponents(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["software_distribution_entry"] = flattenMapHyperflexSoftwareDistributionEntryRelationship(s.GetSoftwareDistributionEntry(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				hyperflexSoftwareDistributionVersionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexSoftwareDistributionVersionResults))
	if err := d.Set("results", hyperflexSoftwareDistributionVersionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexSoftwareDistributionVersionResults[0]["moid"].(string))
	return de
}
