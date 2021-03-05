package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiaapiDcnmFieldNotice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiDcnmFieldNoticeRead,
		Schema: map[string]*schema.Schema{
			"bugid": {
				Description: "Bug Id associated with this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_desc": {
				Description: "Field notice Description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_id": {
				Description: "Fieldnotice Id of this notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field_notice_url": {
				Description: "Field notice URL link to the notice webpage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"headline": {
				Description: "The headline of this field notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hwpid": {
				Description: "Hardware PID for affected models.",
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
			"sw_release": {
				Description: "Software Release number for affected versions.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"workaround_url": {
				Description: "URL of workaround of this notice.",
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
					"bugid": {
						Description: "Bug Id associated with this notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"field_notice_desc": {
						Description: "Field notice Description.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"field_notice_id": {
						Description: "Fieldnotice Id of this notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"field_notice_url": {
						Description: "Field notice URL link to the notice webpage.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"headline": {
						Description: "The headline of this field notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hwpid": {
						Description: "Hardware PID for affected models.",
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
					"revision_info": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"date_published": {
									Description: "The date the revision is made.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"revision_comment": {
									Description: "The changes made in this revision.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"revision_no": {
									Description: "The Revision No. of this revision.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"sw_release": {
						Description: "Software Release number for affected versions.",
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
					"workaround_url": {
						Description: "URL of workaround of this notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiaapiDcnmFieldNoticeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiDcnmFieldNotice{}
	if v, ok := d.GetOk("bugid"); ok {
		x := (v.(string))
		o.SetBugid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("field_notice_desc"); ok {
		x := (v.(string))
		o.SetFieldNoticeDesc(x)
	}
	if v, ok := d.GetOk("field_notice_id"); ok {
		x := (v.(string))
		o.SetFieldNoticeId(x)
	}
	if v, ok := d.GetOk("field_notice_url"); ok {
		x := (v.(string))
		o.SetFieldNoticeUrl(x)
	}
	if v, ok := d.GetOk("headline"); ok {
		x := (v.(string))
		o.SetHeadline(x)
	}
	if v, ok := d.GetOk("hwpid"); ok {
		x := (v.(string))
		o.SetHwpid(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("sw_release"); ok {
		x := (v.(string))
		o.SetSwRelease(x)
	}
	if v, ok := d.GetOk("workaround_url"); ok {
		x := (v.(string))
		o.SetWorkaroundUrl(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiDcnmFieldNotice object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmFieldNoticeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiaapiDcnmFieldNotice: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiaapiDcnmFieldNoticeList.GetCount()
	var i int32
	var niaapiDcnmFieldNoticeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmFieldNoticeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiaapiDcnmFieldNotice: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiaapiDcnmFieldNoticeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiaapiDcnmFieldNotice data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["bugid"] = (s.GetBugid())
				temp["class_id"] = (s.GetClassId())
				temp["field_notice_desc"] = (s.GetFieldNoticeDesc())
				temp["field_notice_id"] = (s.GetFieldNoticeId())
				temp["field_notice_url"] = (s.GetFieldNoticeUrl())
				temp["headline"] = (s.GetHeadline())
				temp["hwpid"] = (s.GetHwpid())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["revision_info"] = flattenListNiaapiRevisionInfo(s.GetRevisionInfo(), d)
				temp["sw_release"] = (s.GetSwRelease())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["workaround_url"] = (s.GetWorkaroundUrl())
				niaapiDcnmFieldNoticeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niaapiDcnmFieldNoticeResults))
	if err := d.Set("results", niaapiDcnmFieldNoticeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niaapiDcnmFieldNoticeResults[0]["moid"].(string))
	return de
}
