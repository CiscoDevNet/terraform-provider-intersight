package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexSoftwareDistributionComponent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexSoftwareDistributionComponentRead,
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Description: "The bucket name where the files are present, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"component_id": {
				Description: "The HyperFlex Software Distribution Component Identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"component_name": {
				Description: "The HyperFlex Software Distribution Component Name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"file_path": {
				Description: "File location on the cloud storage.",
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
				Description: "The HyperFlex Software Distribution Component Version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexSoftwareDistributionComponent().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexSoftwareDistributionComponentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexSoftwareDistributionComponent{}
	if v, ok := d.GetOk("bucket_name"); ok {
		x := (v.(string))
		o.SetBucketName(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("component_id"); ok {
		x := (v.(string))
		o.SetComponentId(x)
	}
	if v, ok := d.GetOk("component_name"); ok {
		x := (v.(string))
		o.SetComponentName(x)
	}
	if v, ok := d.GetOk("file_path"); ok {
		x := (v.(string))
		o.SetFilePath(x)
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
		return diag.Errorf("json marshal of HyperflexSoftwareDistributionComponent object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareDistributionComponentList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexSoftwareDistributionComponent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexSoftwareDistributionComponentList.GetCount()
	var i int32
	var hyperflexSoftwareDistributionComponentResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareDistributionComponentList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexSoftwareDistributionComponent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexSoftwareDistributionComponentList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexSoftwareDistributionComponent data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["bucket_name"] = (s.GetBucketName())
				temp["class_id"] = (s.GetClassId())
				temp["component_id"] = (s.GetComponentId())
				temp["component_name"] = (s.GetComponentName())
				temp["file_path"] = (s.GetFilePath())
				temp["files_to_download"] = (s.GetFilesToDownload())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["software_distribution_version"] = flattenMapHyperflexSoftwareDistributionVersionRelationship(s.GetSoftwareDistributionVersion(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				hyperflexSoftwareDistributionComponentResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexSoftwareDistributionComponentResults))
	if err := d.Set("results", hyperflexSoftwareDistributionComponentResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexSoftwareDistributionComponentResults[0]["moid"].(string))
	return de
}
