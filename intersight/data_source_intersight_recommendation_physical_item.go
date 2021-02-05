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

func dataSourceRecommendationPhysicalItem() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRecommendationPhysicalItemRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"capacity": {
				Description: "Capacity of the physical entity added.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"capacity_runway": {
				Description: "A reference to a recommendationCapacityRunway resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
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
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_count": {
				Description: "Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"is_new": {
				Description: "If the PhysicalItem is new, this is set to true, else false.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"max_count": {
				Description: "Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Description: "Model of the recommended physical device which is externally identifiable.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the physical device recommended.",
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
			"physical_item": {
				Description: "An array of relationships to recommendationPhysicalItem resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
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
			},
			"source_moid": {
				Description: "Moid of the managed object which represents the existing physical entity.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"type": {
				Description: "The type of physical device recommended.\n* `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk.\n* `Node` - The Enum value Node represents that the item type recommended is a Storage Node.\n* `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"unit": {
				Description: "Unit of the new capacity.\n* `TB` - The Enum value TB represents that the measurement unit is in terabytes.\n* `MB` - The Enum value MB represents that the measurement unit is in megabytes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "Uuid of the recommended physical device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceRecommendationPhysicalItemRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.RecommendationPhysicalItem{}
	if v, ok := d.GetOk("capacity"); ok {
		x := int64(v.(int))
		o.SetCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("nr_count"); ok {
		x := int64(v.(int))
		o.SetCount(x)
	}
	if v, ok := d.GetOk("is_new"); ok {
		x := (v.(bool))
		o.SetIsNew(x)
	}
	if v, ok := d.GetOk("max_count"); ok {
		x := int64(v.(int))
		o.SetMaxCount(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
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
	if v, ok := d.GetOk("source_moid"); ok {
		x := (v.(string))
		o.SetSourceMoid(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("unit"); ok {
		x := (v.(string))
		o.SetUnit(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of RecommendationPhysicalItem object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.RecommendationApi.GetRecommendationPhysicalItemList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching RecommendationPhysicalItem: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for RecommendationPhysicalItem list: %s", err.Error())
	}
	var s = &models.RecommendationPhysicalItemList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to RecommendationPhysicalItem list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for RecommendationPhysicalItem data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for RecommendationPhysicalItem data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.RecommendationPhysicalItem{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("capacity", (s.GetCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property Capacity: %s", err.Error())
			}

			if err := d.Set("capacity_runway", flattenMapRecommendationCapacityRunwayRelationship(s.GetCapacityRunway(), d)); err != nil {
				return diag.Errorf("error occurred while setting property CapacityRunway: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("nr_count", (s.GetCount())); err != nil {
				return diag.Errorf("error occurred while setting property Count: %s", err.Error())
			}
			if err := d.Set("is_new", (s.GetIsNew())); err != nil {
				return diag.Errorf("error occurred while setting property IsNew: %s", err.Error())
			}
			if err := d.Set("max_count", (s.GetMaxCount())); err != nil {
				return diag.Errorf("error occurred while setting property MaxCount: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
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

			if err := d.Set("physical_item", flattenListRecommendationPhysicalItemRelationship(s.GetPhysicalItem(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PhysicalItem: %s", err.Error())
			}
			if err := d.Set("source_moid", (s.GetSourceMoid())); err != nil {
				return diag.Errorf("error occurred while setting property SourceMoid: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("unit", (s.GetUnit())); err != nil {
				return diag.Errorf("error occurred while setting property Unit: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
