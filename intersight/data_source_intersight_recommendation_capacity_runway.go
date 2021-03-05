package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRecommendationCapacityRunway() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRecommendationCapacityRunwayRead,
		Schema: map[string]*schema.Schema{
			"additional_capacity": {
				Description: "Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_updated_time": {
				Description: "The time when the recommendation was last updated.",
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
				Description: "The name of the recommendation.",
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
			"period": {
				Description: "Number of months in future for which recommendation is provided for.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"requirement_met": {
				Description: "Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"runway": {
				Description: "This represents the new runway, that is the number of days remaining before the cluster's storage utilization reaches the recommended capacity limit after the recommended hardware is added.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"total_capacity": {
				Description: "Total capacity of the cluster after the recommended hardware is added.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"unit": {
				Description: "Unit for the new capacity.\n* `TB` - The Enum value TB represents that the measurement unit is in terabytes.\n* `MB` - The Enum value MB represents that the measurement unit is in megabytes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_capacity": {
					Description: "Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster.",
					Type:        schema.TypeInt,
					Optional:    true,
					Computed:    true,
				},
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
					"forecast_instance": {
						Description: "A reference to a forecastInstance resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"last_updated_time": {
						Description: "The time when the recommendation was last updated.",
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
						Description: "The name of the recommendation.",
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
					"period": {
						Description: "Number of months in future for which recommendation is provided for.",
						Type:        schema.TypeInt,
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
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"requirement_met": {
						Description: "Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"runway": {
						Description: "This represents the new runway, that is the number of days remaining before the cluster's storage utilization reaches the recommended capacity limit after the recommended hardware is added.",
						Type:        schema.TypeInt,
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
					"total_capacity": {
						Description: "Total capacity of the cluster after the recommended hardware is added.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"unit": {
						Description: "Unit for the new capacity.\n* `TB` - The Enum value TB represents that the measurement unit is in terabytes.\n* `MB` - The Enum value MB represents that the measurement unit is in megabytes.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceRecommendationCapacityRunwayRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.RecommendationCapacityRunway{}
	if v, ok := d.GetOk("additional_capacity"); ok {
		x := int64(v.(int))
		o.SetAdditionalCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("last_updated_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastUpdatedTime(x)
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
	if v, ok := d.GetOk("period"); ok {
		x := int64(v.(int))
		o.SetPeriod(x)
	}
	if v, ok := d.GetOk("requirement_met"); ok {
		x := (v.(bool))
		o.SetRequirementMet(x)
	}
	if v, ok := d.GetOk("runway"); ok {
		x := int64(v.(int))
		o.SetRunway(x)
	}
	if v, ok := d.GetOk("total_capacity"); ok {
		x := int64(v.(int))
		o.SetTotalCapacity(x)
	}
	if v, ok := d.GetOk("unit"); ok {
		x := (v.(string))
		o.SetUnit(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of RecommendationCapacityRunway object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.RecommendationApi.GetRecommendationCapacityRunwayList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of RecommendationCapacityRunway: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.RecommendationCapacityRunwayList.GetCount()
	var i int32
	var recommendationCapacityRunwayResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.RecommendationApi.GetRecommendationCapacityRunwayList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching RecommendationCapacityRunway: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.RecommendationCapacityRunwayList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for RecommendationCapacityRunway data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_capacity"] = (s.GetAdditionalCapacity())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["forecast_instance"] = flattenMapForecastInstanceRelationship(s.GetForecastInstance(), d)

				temp["last_updated_time"] = (s.GetLastUpdatedTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["period"] = (s.GetPeriod())

				temp["physical_item"] = flattenListRecommendationPhysicalItemRelationship(s.GetPhysicalItem(), d)

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["requirement_met"] = (s.GetRequirementMet())
				temp["runway"] = (s.GetRunway())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["total_capacity"] = (s.GetTotalCapacity())
				temp["unit"] = (s.GetUnit())
				recommendationCapacityRunwayResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(recommendationCapacityRunwayResults))
	if err := d.Set("results", recommendationCapacityRunwayResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(recommendationCapacityRunwayResults[0]["moid"].(string))
	return de
}
