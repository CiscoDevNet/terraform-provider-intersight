package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHyperflexSoftwareDistributionComponent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHyperflexSoftwareDistributionComponentCreate,
		ReadContext:   resourceHyperflexSoftwareDistributionComponentRead,
		UpdateContext: resourceHyperflexSoftwareDistributionComponentUpdate,
		DeleteContext: resourceHyperflexSoftwareDistributionComponentDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bucket_name": {
				Description: "The bucket name where the files are present, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"files_to_download": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"software_distribution_version": {
				Description: "A reference to a hyperflexSoftwareDistributionVersion resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"nr_version": {
				Description: "The HyperFlex Software Distribution Component Version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceHyperflexSoftwareDistributionComponentCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewHyperflexSoftwareDistributionComponentWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("bucket_name"); ok {
		x := (v.(string))
		o.SetBucketName(x)
	}

	o.SetClassId("hyperflex.SoftwareDistributionComponent")

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

	if v, ok := d.GetOk("files_to_download"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetFilesToDownload(x)
		}
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("hyperflex.SoftwareDistributionComponent")

	if v, ok := d.GetOk("software_distribution_version"); ok {
		p := make([]models.HyperflexSoftwareDistributionVersionRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexSoftwareDistributionVersionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSoftwareDistributionVersion(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.HyperflexApi.CreateHyperflexSoftwareDistributionComponent(conn.ctx).HyperflexSoftwareDistributionComponent(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating HyperflexSoftwareDistributionComponent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceHyperflexSoftwareDistributionComponentRead(c, d, meta)...)
}

func resourceHyperflexSoftwareDistributionComponentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareDistributionComponentByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexSoftwareDistributionComponent object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching HyperflexSoftwareDistributionComponent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("bucket_name", (s.GetBucketName())); err != nil {
		return diag.Errorf("error occurred while setting property BucketName in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("component_id", (s.GetComponentId())); err != nil {
		return diag.Errorf("error occurred while setting property ComponentId in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("component_name", (s.GetComponentName())); err != nil {
		return diag.Errorf("error occurred while setting property ComponentName in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("file_path", (s.GetFilePath())); err != nil {
		return diag.Errorf("error occurred while setting property FilePath in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("files_to_download", (s.GetFilesToDownload())); err != nil {
		return diag.Errorf("error occurred while setting property FilesToDownload in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("software_distribution_version", flattenMapHyperflexSoftwareDistributionVersionRelationship(s.GetSoftwareDistributionVersion(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SoftwareDistributionVersion in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return diag.Errorf("error occurred while setting property Version in HyperflexSoftwareDistributionComponent object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHyperflexSoftwareDistributionComponentUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexSoftwareDistributionComponent{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("bucket_name") {
		v := d.Get("bucket_name")
		x := (v.(string))
		o.SetBucketName(x)
	}

	o.SetClassId("hyperflex.SoftwareDistributionComponent")

	if d.HasChange("component_id") {
		v := d.Get("component_id")
		x := (v.(string))
		o.SetComponentId(x)
	}

	if d.HasChange("component_name") {
		v := d.Get("component_name")
		x := (v.(string))
		o.SetComponentName(x)
	}

	if d.HasChange("file_path") {
		v := d.Get("file_path")
		x := (v.(string))
		o.SetFilePath(x)
	}

	if d.HasChange("files_to_download") {
		v := d.Get("files_to_download")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetFilesToDownload(x)
		}
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("hyperflex.SoftwareDistributionComponent")

	if d.HasChange("software_distribution_version") {
		v := d.Get("software_distribution_version")
		p := make([]models.HyperflexSoftwareDistributionVersionRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexSoftwareDistributionVersionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSoftwareDistributionVersion(x)
		}
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("nr_version") {
		v := d.Get("nr_version")
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexSoftwareDistributionComponent(conn.ctx, d.Id()).HyperflexSoftwareDistributionComponent(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating HyperflexSoftwareDistributionComponent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceHyperflexSoftwareDistributionComponentRead(c, d, meta)...)
}

func resourceHyperflexSoftwareDistributionComponentDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HyperflexApi.DeleteHyperflexSoftwareDistributionComponent(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting HyperflexSoftwareDistributionComponent object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
