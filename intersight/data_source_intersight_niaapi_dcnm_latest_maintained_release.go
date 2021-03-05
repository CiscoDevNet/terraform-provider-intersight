package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiaapiDcnmLatestMaintainedRelease() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiDcnmLatestMaintainedReleaseRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"maintained_release": {
				Description: "Lastest maintained release.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"software_release": {
				Description: "Software release version string.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_tag": {
				Description: "Long lived version or short lived version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maintained_release": {
						Description: "Lastest maintained release.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"software_release": {
						Description: "Software release version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"version_tag": {
						Description: "Long lived version or short lived version.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiaapiDcnmLatestMaintainedReleaseRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiDcnmLatestMaintainedRelease{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("maintained_release"); ok {
		x := (v.(string))
		o.SetMaintainedRelease(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("software_release"); ok {
		x := (v.(string))
		o.SetSoftwareRelease(x)
	}
	if v, ok := d.GetOk("version_tag"); ok {
		x := (v.(string))
		o.SetVersionTag(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiDcnmLatestMaintainedRelease object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiaapiDcnmLatestMaintainedRelease: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiaapiDcnmLatestMaintainedReleaseList.GetCount()
	var i int32
	var niaapiDcnmLatestMaintainedReleaseResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmLatestMaintainedReleaseList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiaapiDcnmLatestMaintainedRelease: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiaapiDcnmLatestMaintainedReleaseList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiaapiDcnmLatestMaintainedRelease data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["maintained_release"] = (s.GetMaintainedRelease())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["software_release"] = (s.GetSoftwareRelease())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["version_tag"] = (s.GetVersionTag())
				niaapiDcnmLatestMaintainedReleaseResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niaapiDcnmLatestMaintainedReleaseResults))
	if err := d.Set("results", niaapiDcnmLatestMaintainedReleaseResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niaapiDcnmLatestMaintainedReleaseResults[0]["moid"].(string))
	return de
}
