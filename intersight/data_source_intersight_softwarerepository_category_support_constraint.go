package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSoftwarerepositoryCategorySupportConstraint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryCategorySupportConstraintRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"constraint_id": {
				Description: "Identifier for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mdf_id": {
				Description: "Cisco software repository image category identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"min_version": {
				Description: "Minimum image version from where the models can be supported.",
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
				Description: "An unique identifer for a capability descriptor.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parse_from_image_name": {
				Description: "Fields which tells if the constraint is based on image name parsing.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSoftwarerepositoryCategorySupportConstraint().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSoftwarerepositoryCategorySupportConstraintRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryCategorySupportConstraint{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("constraint_id"); ok {
		x := (v.(string))
		o.SetConstraintId(x)
	}
	if v, ok := d.GetOk("mdf_id"); ok {
		x := (v.(string))
		o.SetMdfId(x)
	}
	if v, ok := d.GetOk("min_version"); ok {
		x := (v.(string))
		o.SetMinVersion(x)
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
	if v, ok := d.GetOk("parse_from_image_name"); ok {
		x := (v.(bool))
		o.SetParseFromImageName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryCategorySupportConstraint object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategorySupportConstraintList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SoftwarerepositoryCategorySupportConstraint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SoftwarerepositoryCategorySupportConstraintList.GetCount()
	var i int32
	var softwarerepositoryCategorySupportConstraintResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategorySupportConstraintList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SoftwarerepositoryCategorySupportConstraint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SoftwarerepositoryCategorySupportConstraintList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SoftwarerepositoryCategorySupportConstraint data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["constraint_id"] = (s.GetConstraintId())

				temp["filtered_models"] = flattenListSoftwarerepositoryConstraintModels(s.GetFilteredModels(), d)
				temp["mdf_id"] = (s.GetMdfId())
				temp["min_version"] = (s.GetMinVersion())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["parse_from_image_name"] = (s.GetParseFromImageName())
				temp["supported_models"] = (s.GetSupportedModels())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				softwarerepositoryCategorySupportConstraintResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(softwarerepositoryCategorySupportConstraintResults))
	if err := d.Set("results", softwarerepositoryCategorySupportConstraintResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(softwarerepositoryCategorySupportConstraintResults[0]["moid"].(string))
	return de
}
