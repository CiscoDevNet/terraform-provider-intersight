package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHclDriverImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHclDriverImageRead,
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
			"driver_iso_url": {
				Description: "URL of the driver ISO images.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"management_type": {
				Description: "Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.\n* `UCSM` - The server is managed by UCS Manager.\n* `IMC` - The server is standalone managed by CIMC.",
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
			"server_pid": {
				Description: "Three part ID representing the server model as returned by UCSM/CIMC XML APIs.",
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
		},
	}
}

func dataSourceHclDriverImageRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HclDriverImage{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("driver_iso_url"); ok {
		x := (v.(string))
		o.SetDriverIsoUrl(x)
	}
	if v, ok := d.GetOk("management_type"); ok {
		x := (v.(string))
		o.SetManagementType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("server_pid"); ok {
		x := (v.(string))
		o.SetServerPid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HclDriverImage object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HclApi.GetHclDriverImageList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HclDriverImage: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HclDriverImage list: %s", err.Error())
	}
	var s = &models.HclDriverImageList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HclDriverImage list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for HclDriverImage did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HclDriverImage{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("driver_iso_url", (s.GetDriverIsoUrl())); err != nil {
				return diag.Errorf("error occurred while setting property DriverIsoUrl: %s", err.Error())
			}
			if err := d.Set("management_type", (s.GetManagementType())); err != nil {
				return diag.Errorf("error occurred while setting property ManagementType: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("server_pid", (s.GetServerPid())); err != nil {
				return diag.Errorf("error occurred while setting property ServerPid: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
