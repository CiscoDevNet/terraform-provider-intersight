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

func resourceSoftwarerepositoryCategoryMapper() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSoftwarerepositoryCategoryMapperCreate,
		ReadContext:   resourceSoftwarerepositoryCategoryMapperRead,
		UpdateContext: resourceSoftwarerepositoryCategoryMapperUpdate,
		DeleteContext: resourceSoftwarerepositoryCategoryMapperDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"category": {
				Description: "The category of the model series.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"file_type": {
				Description: "The type of distributable image, example huu, scu, driver, os.\n* `Distributable` - Stores firmware host utility images and fabric images.\n* `DriverDistributable` - Stores driver distributable images.\n* `ServerConfigurationUtilityDistributable` - Stores server configuration utility images.\n* `OperatingSystemFile` - Stores operating system iso images.\n* `HyperflexDistributable` - It stores HyperFlex images.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Distributable",
			},
			"mdf_id": {
				Description: "Cisco software repository image category identifier.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
			"regex_pattern": {
				Description: "The regex that all images of this category follow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_source": {
				Description: "The image can be downloaded from cisco.com or external cloud store.\n* `Cisco` - External repository hosted on cisco.com.\n* `IntersightCloud` - Repository hosted by the Intersight Cloud.\n* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.\n* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Cisco",
			},
			"supported_models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"sw_id": {
				Description: "The software type id provided by cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag_types": {
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
			"nr_version": {
				Description: "The version from which user can download images from amazon store, if source is external cloud store.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceSoftwarerepositoryCategoryMapperCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewSoftwarerepositoryCategoryMapperWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}

	o.SetClassId("softwarerepository.CategoryMapper")

	if v, ok := d.GetOk("file_type"); ok {
		x := (v.(string))
		o.SetFileType(x)
	}

	if v, ok := d.GetOk("mdf_id"); ok {
		x := (v.(string))
		o.SetMdfId(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("softwarerepository.CategoryMapper")

	if v, ok := d.GetOk("regex_pattern"); ok {
		x := (v.(string))
		o.SetRegexPattern(x)
	}

	if v, ok := d.GetOk("nr_source"); ok {
		x := (v.(string))
		o.SetSource(x)
	}

	if v, ok := d.GetOk("supported_models"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetSupportedModels(x)
		}
	}

	if v, ok := d.GetOk("sw_id"); ok {
		x := (v.(string))
		o.SetSwId(x)
	}

	if v, ok := d.GetOk("tag_types"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetTagTypes(x)
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

	r := conn.ApiClient.SoftwarerepositoryApi.CreateSoftwarerepositoryCategoryMapper(conn.ctx).SoftwarerepositoryCategoryMapper(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating SoftwarerepositoryCategoryMapper: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceSoftwarerepositoryCategoryMapperRead(c, d, meta)...)
}

func resourceSoftwarerepositoryCategoryMapperRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategoryMapperByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "SoftwarerepositoryCategoryMapper object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching SoftwarerepositoryCategoryMapper: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("category", (s.GetCategory())); err != nil {
		return diag.Errorf("error occurred while setting property Category in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("file_type", (s.GetFileType())); err != nil {
		return diag.Errorf("error occurred while setting property FileType in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("mdf_id", (s.GetMdfId())); err != nil {
		return diag.Errorf("error occurred while setting property MdfId in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("regex_pattern", (s.GetRegexPattern())); err != nil {
		return diag.Errorf("error occurred while setting property RegexPattern in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("nr_source", (s.GetSource())); err != nil {
		return diag.Errorf("error occurred while setting property Source in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("supported_models", (s.GetSupportedModels())); err != nil {
		return diag.Errorf("error occurred while setting property SupportedModels in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("sw_id", (s.GetSwId())); err != nil {
		return diag.Errorf("error occurred while setting property SwId in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("tag_types", (s.GetTagTypes())); err != nil {
		return diag.Errorf("error occurred while setting property TagTypes in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return diag.Errorf("error occurred while setting property Version in SoftwarerepositoryCategoryMapper object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceSoftwarerepositoryCategoryMapperUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryCategoryMapper{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("category") {
		v := d.Get("category")
		x := (v.(string))
		o.SetCategory(x)
	}

	o.SetClassId("softwarerepository.CategoryMapper")

	if d.HasChange("file_type") {
		v := d.Get("file_type")
		x := (v.(string))
		o.SetFileType(x)
	}

	if d.HasChange("mdf_id") {
		v := d.Get("mdf_id")
		x := (v.(string))
		o.SetMdfId(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("softwarerepository.CategoryMapper")

	if d.HasChange("regex_pattern") {
		v := d.Get("regex_pattern")
		x := (v.(string))
		o.SetRegexPattern(x)
	}

	if d.HasChange("nr_source") {
		v := d.Get("nr_source")
		x := (v.(string))
		o.SetSource(x)
	}

	if d.HasChange("supported_models") {
		v := d.Get("supported_models")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetSupportedModels(x)
		}
	}

	if d.HasChange("sw_id") {
		v := d.Get("sw_id")
		x := (v.(string))
		o.SetSwId(x)
	}

	if d.HasChange("tag_types") {
		v := d.Get("tag_types")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetTagTypes(x)
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

	r := conn.ApiClient.SoftwarerepositoryApi.UpdateSoftwarerepositoryCategoryMapper(conn.ctx, d.Id()).SoftwarerepositoryCategoryMapper(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating SoftwarerepositoryCategoryMapper: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceSoftwarerepositoryCategoryMapperRead(c, d, meta)...)
}

func resourceSoftwarerepositoryCategoryMapperDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.SoftwarerepositoryApi.DeleteSoftwarerepositoryCategoryMapper(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting SoftwarerepositoryCategoryMapper object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
