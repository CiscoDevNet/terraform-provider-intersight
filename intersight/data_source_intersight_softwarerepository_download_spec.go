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

func dataSourceSoftwarerepositoryDownloadSpec() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryDownloadSpecRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"auth_token": {
				Description: "The OAuth2 token that will be used during image download by the endpoint to authenticate with file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"certificate": {
				Description: "The certificate of file server that will be used by the endpoint to validate the server before starting the file download.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"file": {
				Description: "A reference to a softwarerepositoryFile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"filename": {
				Description: "The name of the firmware image.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"md5sum": {
				Description: "MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download.",
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
			"size": {
				Description: "The size (in bytes) of the firmware image.",
				Type:        schema.TypeInt,
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
			"url": {
				Description: "The URL of this file in file server. The endpoint uses this URL to download the file from the file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSoftwarerepositoryDownloadSpecRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryDownloadSpec{}
	if v, ok := d.GetOk("auth_token"); ok {
		x := (v.(string))
		o.SetAuthToken(x)
	}
	if v, ok := d.GetOk("certificate"); ok {
		x := (v.(string))
		o.SetCertificate(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}
	if v, ok := d.GetOk("md5sum"); ok {
		x := (v.(string))
		o.SetMd5sum(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("url"); ok {
		x := (v.(string))
		o.SetUrl(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryDownloadSpec object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryDownloadSpecList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching SoftwarerepositoryDownloadSpec: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for SoftwarerepositoryDownloadSpec list: %s", err.Error())
	}
	var s = &models.SoftwarerepositoryDownloadSpecList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to SoftwarerepositoryDownloadSpec list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for SoftwarerepositoryDownloadSpec did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.SoftwarerepositoryDownloadSpec{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("auth_token", (s.GetAuthToken())); err != nil {
				return diag.Errorf("error occurred while setting property AuthToken: %s", err.Error())
			}
			if err := d.Set("certificate", (s.GetCertificate())); err != nil {
				return diag.Errorf("error occurred while setting property Certificate: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("file", flattenMapSoftwarerepositoryFileRelationship(s.GetFile(), d)); err != nil {
				return diag.Errorf("error occurred while setting property File: %s", err.Error())
			}
			if err := d.Set("filename", (s.GetFilename())); err != nil {
				return diag.Errorf("error occurred while setting property Filename: %s", err.Error())
			}
			if err := d.Set("md5sum", (s.GetMd5sum())); err != nil {
				return diag.Errorf("error occurred while setting property Md5sum: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return diag.Errorf("error occurred while setting property Size: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("url", (s.GetUrl())); err != nil {
				return diag.Errorf("error occurred while setting property Url: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
