package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSoftwarerepositoryCategorySupportConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceSoftwarerepositoryCategorySupportConstraintCreate,
		Read:   resourceSoftwarerepositoryCategorySupportConstraintRead,
		Update: resourceSoftwarerepositoryCategorySupportConstraintUpdate,
		Delete: resourceSoftwarerepositoryCategorySupportConstraintDelete,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"constraint_id": {
				Description: "Identifier for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filtered_models": {
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"min_version": {
							Description: "Minimum version of the image.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Name of the contraint, used to identify constriant type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"platform_regex": {
							Description: "Regular expression of the image name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"supported_models": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
				ForceNew:    true,
			},
			"name": {
				Description: "An unique identifer for a capability descriptor.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parse_from_image_name": {
				Description: "Fields which tells if the constraint is based on image name parsing.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"section": {
				Description: "A reference to a capabilitySection resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
		},
	}
}

func resourceSoftwarerepositoryCategorySupportConstraintCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSoftwarerepositoryCategorySupportConstraintWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("softwarerepository.CategorySupportConstraint")

	if v, ok := d.GetOk("constraint_id"); ok {
		x := (v.(string))
		o.SetConstraintId(x)
	}

	if v, ok := d.GetOk("filtered_models"); ok {
		x := make([]models.SoftwarerepositoryConstraintModels, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewSoftwarerepositoryConstraintModelsWithDefaults()
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
			o.SetClassId("softwarerepository.ConstraintModels")
			if v, ok := l["min_version"]; ok {
				{
					x := (v.(string))
					o.SetMinVersion(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["platform_regex"]; ok {
				{
					x := (v.(string))
					o.SetPlatformRegex(x)
				}
			}
			if v, ok := l["supported_models"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetSupportedModels(x)
					}
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetFilteredModels(x)
		}
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

	o.SetObjectType("softwarerepository.CategorySupportConstraint")

	if v, ok := d.GetOkExists("parse_from_image_name"); ok {
		x := v.(bool)
		o.SetParseFromImageName(x)
	}

	if v, ok := d.GetOk("section"); ok {
		p := make([]models.CapabilitySectionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsCapabilitySectionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSection(x)
		}
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

	r := conn.ApiClient.SoftwarerepositoryApi.CreateSoftwarerepositoryCategorySupportConstraint(conn.ctx).SoftwarerepositoryCategorySupportConstraint(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceSoftwarerepositoryCategorySupportConstraintRead(d, meta)
}

func resourceSoftwarerepositoryCategorySupportConstraintRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategorySupportConstraintByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("constraint_id", (s.GetConstraintId())); err != nil {
		return fmt.Errorf("error occurred while setting property ConstraintId: %+v", err)
	}

	if err := d.Set("filtered_models", flattenListSoftwarerepositoryConstraintModels(s.GetFilteredModels(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property FilteredModels: %+v", err)
	}

	if err := d.Set("mdf_id", (s.GetMdfId())); err != nil {
		return fmt.Errorf("error occurred while setting property MdfId: %+v", err)
	}

	if err := d.Set("min_version", (s.GetMinVersion())); err != nil {
		return fmt.Errorf("error occurred while setting property MinVersion: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return fmt.Errorf("error occurred while setting property Name: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("parse_from_image_name", (s.GetParseFromImageName())); err != nil {
		return fmt.Errorf("error occurred while setting property ParseFromImageName: %+v", err)
	}

	if err := d.Set("section", flattenMapCapabilitySectionRelationship(s.GetSection(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Section: %+v", err)
	}

	if err := d.Set("supported_models", (s.GetSupportedModels())); err != nil {
		return fmt.Errorf("error occurred while setting property SupportedModels: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceSoftwarerepositoryCategorySupportConstraintUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewSoftwarerepositoryCategorySupportConstraintWithDefaults()
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("softwarerepository.CategorySupportConstraint")

	if d.HasChange("constraint_id") {
		v := d.Get("constraint_id")
		x := (v.(string))
		o.SetConstraintId(x)
	}

	if d.HasChange("filtered_models") {
		v := d.Get("filtered_models")
		x := make([]models.SoftwarerepositoryConstraintModels, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewSoftwarerepositoryConstraintModelsWithDefaults()
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
			o.SetClassId("softwarerepository.ConstraintModels")
			if v, ok := l["min_version"]; ok {
				{
					x := (v.(string))
					o.SetMinVersion(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["platform_regex"]; ok {
				{
					x := (v.(string))
					o.SetPlatformRegex(x)
				}
			}
			if v, ok := l["supported_models"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						x = append(x, y.Index(i).Interface().(string))
					}
					if len(x) > 0 {
						o.SetSupportedModels(x)
					}
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetFilteredModels(x)
		}
	}

	if d.HasChange("mdf_id") {
		v := d.Get("mdf_id")
		x := (v.(string))
		o.SetMdfId(x)
	}

	if d.HasChange("min_version") {
		v := d.Get("min_version")
		x := (v.(string))
		o.SetMinVersion(x)
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

	o.SetObjectType("softwarerepository.CategorySupportConstraint")

	if d.HasChange("parse_from_image_name") {
		v := d.Get("parse_from_image_name")
		x := (v.(bool))
		o.SetParseFromImageName(x)
	}

	if d.HasChange("section") {
		v := d.Get("section")
		p := make([]models.CapabilitySectionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsCapabilitySectionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSection(x)
		}
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

	if d.HasChange("tags") {
		v := d.Get("tags")
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

	r := conn.ApiClient.SoftwarerepositoryApi.UpdateSoftwarerepositoryCategorySupportConstraint(conn.ctx, d.Id()).SoftwarerepositoryCategorySupportConstraint(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceSoftwarerepositoryCategorySupportConstraintRead(d, meta)
}

func resourceSoftwarerepositoryCategorySupportConstraintDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	p := conn.ApiClient.SoftwarerepositoryApi.DeleteSoftwarerepositoryCategorySupportConstraint(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
