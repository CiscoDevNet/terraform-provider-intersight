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

func dataSourceFirmwareDistributableMeta() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareDistributableMetaRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bucket_name": {
				Description: "The S3 bucket name where the images are present, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"file_type": {
				Description: "The type of distributable image, example huu, scu, driver, os.\n* `Distributable` - Stores firmware host utility images and fabric images.\n* `DriverDistributable` - Stores driver distributable images.\n* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.\n* `OperatingSystemFile` - Stores operating system iso images.\n* `HyperflexDistributable` - It stores HyperFlex images.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"from_version": {
				Description: "The version from which user can download images from amazon store, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mdfid": {
				Description: "The mdfid of the image provided by cisco.com.",
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
			"software_type_id": {
				Description: "The software type id provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_source": {
				Description: "The image can be downloaded from cisco.com or external cloud store.\n* `Cisco` - External repository hosted on cisco.com.\n* `IntersightCloud` - Repository hosted by the Intersight Cloud.\n* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.\n* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"supported_models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
			"to_version": {
				Description: "The version till which user can download images from amazon store, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceFirmwareDistributableMetaRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareDistributableMeta{}
	if v, ok := d.GetOk("bucket_name"); ok {
		x := (v.(string))
		o.SetBucketName(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("file_type"); ok {
		x := (v.(string))
		o.SetFileType(x)
	}
	if v, ok := d.GetOk("from_version"); ok {
		x := (v.(string))
		o.SetFromVersion(x)
	}
	if v, ok := d.GetOk("mdfid"); ok {
		x := (v.(string))
		o.SetMdfid(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("software_type_id"); ok {
		x := (v.(string))
		o.SetSoftwareTypeId(x)
	}
	if v, ok := d.GetOk("nr_source"); ok {
		x := (v.(string))
		o.SetSource(x)
	}
	if v, ok := d.GetOk("to_version"); ok {
		x := (v.(string))
		o.SetToVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareDistributableMeta object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareDistributableMetaList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching FirmwareDistributableMeta: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for FirmwareDistributableMeta list: %s", err.Error())
	}
	var s = &models.FirmwareDistributableMetaList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to FirmwareDistributableMeta list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for FirmwareDistributableMeta data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for FirmwareDistributableMeta data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.FirmwareDistributableMeta{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("bucket_name", (s.GetBucketName())); err != nil {
				return diag.Errorf("error occurred while setting property BucketName: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("file_type", (s.GetFileType())); err != nil {
				return diag.Errorf("error occurred while setting property FileType: %s", err.Error())
			}
			if err := d.Set("from_version", (s.GetFromVersion())); err != nil {
				return diag.Errorf("error occurred while setting property FromVersion: %s", err.Error())
			}
			if err := d.Set("mdfid", (s.GetMdfid())); err != nil {
				return diag.Errorf("error occurred while setting property Mdfid: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("software_type_id", (s.GetSoftwareTypeId())); err != nil {
				return diag.Errorf("error occurred while setting property SoftwareTypeId: %s", err.Error())
			}
			if err := d.Set("nr_source", (s.GetSource())); err != nil {
				return diag.Errorf("error occurred while setting property Source: %s", err.Error())
			}
			if err := d.Set("supported_models", (s.GetSupportedModels())); err != nil {
				return diag.Errorf("error occurred while setting property SupportedModels: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("to_version", (s.GetToVersion())); err != nil {
				return diag.Errorf("error occurred while setting property ToVersion: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
