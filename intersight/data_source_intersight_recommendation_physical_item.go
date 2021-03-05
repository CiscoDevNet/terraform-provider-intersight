package intersight

import (
	"context"
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
			"capacity": {
				Description: "Capacity of the physical entity added.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
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
			"source_moid": {
				Description: "Moid of the managed object which represents the existing physical entity.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
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
				}},
				Computed: true,
			}},
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
	countResponse, _, responseErr := conn.ApiClient.RecommendationApi.GetRecommendationPhysicalItemList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of RecommendationPhysicalItem: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.RecommendationPhysicalItemList.GetCount()
	var i int32
	var recommendationPhysicalItemResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.RecommendationApi.GetRecommendationPhysicalItemList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching RecommendationPhysicalItem: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.RecommendationPhysicalItemList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for RecommendationPhysicalItem data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["capacity"] = (s.GetCapacity())

				temp["capacity_runway"] = flattenMapRecommendationCapacityRunwayRelationship(s.GetCapacityRunway(), d)
				temp["class_id"] = (s.GetClassId())
				temp["nr_count"] = (s.GetCount())
				temp["is_new"] = (s.GetIsNew())
				temp["max_count"] = (s.GetMaxCount())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["physical_item"] = flattenListRecommendationPhysicalItemRelationship(s.GetPhysicalItem(), d)
				temp["source_moid"] = (s.GetSourceMoid())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["unit"] = (s.GetUnit())
				temp["uuid"] = (s.GetUuid())
				recommendationPhysicalItemResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(recommendationPhysicalItemResults))
	if err := d.Set("results", recommendationPhysicalItemResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(recommendationPhysicalItemResults[0]["moid"].(string))
	return de
}
