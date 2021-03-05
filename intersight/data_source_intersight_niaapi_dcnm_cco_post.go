package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiaapiDcnmCcoPost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiDcnmCcoPostRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
			"post_date": {
				Description: "The date when this new release notice is posted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"post_type": {
				Description: "The document type of this post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"postid": {
				Description: "Identificator of this inbox post.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"revision": {
				Description: "Revision number of this notice.",
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
					"post_date": {
						Description: "The date when this new release notice is posted.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"post_detail": {
						Description: "Detail of this post including the content and the date it was posted.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
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
								"description": {
									Description: "Description of this new verison release post.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"link": {
									Description: "Link of downloading the release file.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"release_note_link": {
									Description: "The link used to provide a gateway for customer to review the release note.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"release_note_link_title": {
									Description: "The link title used to provide a gateway for customer to review the release note.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"software_download_link": {
									Description: "The link used to provide a gateway for customer to download the release.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"software_download_link_title": {
									Description: "The link title used to provide a gateway for customer to download the release.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"title": {
									Description: "Title of the new verison release post.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"nr_version": {
									Description: "Version number Associate with this Post.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"post_type": {
						Description: "The document type of this post.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"postid": {
						Description: "Identificator of this inbox post.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"revision": {
						Description: "Revision number of this notice.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiaapiDcnmCcoPostRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiDcnmCcoPost{}
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
	if v, ok := d.GetOk("post_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetPostDate(x)
	}
	if v, ok := d.GetOk("post_type"); ok {
		x := (v.(string))
		o.SetPostType(x)
	}
	if v, ok := d.GetOk("postid"); ok {
		x := (v.(string))
		o.SetPostid(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiDcnmCcoPost object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmCcoPostList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiaapiDcnmCcoPost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiaapiDcnmCcoPostList.GetCount()
	var i int32
	var niaapiDcnmCcoPostResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiDcnmCcoPostList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiaapiDcnmCcoPost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiaapiDcnmCcoPostList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiaapiDcnmCcoPost data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["post_date"] = (s.GetPostDate()).String()

				temp["post_detail"] = flattenMapNiaapiNewReleaseDetail(s.GetPostDetail(), d)
				temp["post_type"] = (s.GetPostType())
				temp["postid"] = (s.GetPostid())
				temp["revision"] = (s.GetRevision())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				niaapiDcnmCcoPostResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niaapiDcnmCcoPostResults))
	if err := d.Set("results", niaapiDcnmCcoPostResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niaapiDcnmCcoPostResults[0]["moid"].(string))
	return de
}
