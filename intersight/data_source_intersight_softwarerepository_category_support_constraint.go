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

func dataSourceSoftwarerepositoryCategorySupportConstraint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryCategorySupportConstraintRead,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
				Computed: true,
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
			"parse_from_image_name": {
				Description: "Fields which tells if the constraint is based on image name parsing.",
				Type:        schema.TypeBool,
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
		},
	}
}

func dataSourceSoftwarerepositoryCategorySupportConstraintRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryCategorySupportConstraint{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("constraint_id"); ok {
		x := (v.(string))
		o.SetConstraintId(x)
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parse_from_image_name"); ok {
		x := (v.(bool))
		o.SetParseFromImageName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryCategorySupportConstraint object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategorySupportConstraintList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching SoftwarerepositoryCategorySupportConstraint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for SoftwarerepositoryCategorySupportConstraint list: %s", err.Error())
	}
	var s = &models.SoftwarerepositoryCategorySupportConstraintList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to SoftwarerepositoryCategorySupportConstraint list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for SoftwarerepositoryCategorySupportConstraint data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for SoftwarerepositoryCategorySupportConstraint data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.SoftwarerepositoryCategorySupportConstraint{}
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
			if err := d.Set("constraint_id", (s.GetConstraintId())); err != nil {
				return diag.Errorf("error occurred while setting property ConstraintId: %s", err.Error())
			}

			if err := d.Set("filtered_models", flattenListSoftwarerepositoryConstraintModels(s.GetFilteredModels(), d)); err != nil {
				return diag.Errorf("error occurred while setting property FilteredModels: %s", err.Error())
			}
			if err := d.Set("mdf_id", (s.GetMdfId())); err != nil {
				return diag.Errorf("error occurred while setting property MdfId: %s", err.Error())
			}
			if err := d.Set("min_version", (s.GetMinVersion())); err != nil {
				return diag.Errorf("error occurred while setting property MinVersion: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("parse_from_image_name", (s.GetParseFromImageName())); err != nil {
				return diag.Errorf("error occurred while setting property ParseFromImageName: %s", err.Error())
			}
			if err := d.Set("supported_models", (s.GetSupportedModels())); err != nil {
				return diag.Errorf("error occurred while setting property SupportedModels: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
