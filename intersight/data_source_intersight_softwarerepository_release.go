package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSoftwarerepositoryRelease() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryReleaseRead,
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
			"release_notes_url": {
				Description: "The URL for the release notes of this image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware.\n* `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches.\n* `ComputeSystem` - The images in a release that correspond to servers.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "Cisco provided release version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSoftwarerepositoryRelease().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSoftwarerepositoryReleaseRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryRelease{}
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
	if v, ok := d.GetOk("release_notes_url"); ok {
		x := (v.(string))
		o.SetReleaseNotesUrl(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryRelease object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryReleaseList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SoftwarerepositoryRelease: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SoftwarerepositoryReleaseList.GetCount()
	var i int32
	var softwarerepositoryReleaseResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryReleaseList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SoftwarerepositoryRelease: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SoftwarerepositoryReleaseList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SoftwarerepositoryRelease data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["catalog"] = flattenMapSoftwarerepositoryCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["release_notes_url"] = (s.GetReleaseNotesUrl())
				temp["supported_models"] = (s.GetSupportedModels())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["nr_version"] = (s.GetVersion())
				softwarerepositoryReleaseResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(softwarerepositoryReleaseResults))
	if err := d.Set("results", softwarerepositoryReleaseResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(softwarerepositoryReleaseResults[0]["moid"].(string))
	return de
}
