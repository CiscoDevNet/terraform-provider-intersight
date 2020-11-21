package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiApicFieldNotice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiApicFieldNoticeRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
		},
	}
}

func dataSourceNiaapiApicFieldNoticeRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.NiaapiApicFieldNotice{}
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiaapiApi.GetNiaapiApicFieldNoticeList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiaapiApicFieldNoticeList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiaapiApicFieldNotice: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiaapiApicFieldNotice{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("bugid", (s.GetBugid())); err != nil {
				return fmt.Errorf("error occurred while setting property Bugid: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("field_notice_desc", (s.GetFieldNoticeDesc())); err != nil {
				return fmt.Errorf("error occurred while setting property FieldNoticeDesc: %+v", err)
			}
			if err := d.Set("field_notice_id", (s.GetFieldNoticeId())); err != nil {
				return fmt.Errorf("error occurred while setting property FieldNoticeId: %+v", err)
			}
			if err := d.Set("field_notice_url", (s.GetFieldNoticeUrl())); err != nil {
				return fmt.Errorf("error occurred while setting property FieldNoticeUrl: %+v", err)
			}
			if err := d.Set("headline", (s.GetHeadline())); err != nil {
				return fmt.Errorf("error occurred while setting property Headline: %+v", err)
			}
			if err := d.Set("hwpid", (s.GetHwpid())); err != nil {
				return fmt.Errorf("error occurred while setting property Hwpid: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("revision_info", flattenListNiaapiRevisionInfo(s.GetRevisionInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RevisionInfo: %+v", err)
			}
			if err := d.Set("sw_release", (s.GetSwRelease())); err != nil {
				return fmt.Errorf("error occurred while setting property SwRelease: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("workaround_url", (s.GetWorkaroundUrl())); err != nil {
				return fmt.Errorf("error occurred while setting property WorkaroundUrl: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
