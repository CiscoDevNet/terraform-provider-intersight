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

func dataSourceForecastInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceForecastInstanceRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"alt_model": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat}},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_id": {
				Description: "The Moid of the Intersight managed device instance for which regression model is derived.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"forecast_def": {
				Description: "A reference to a forecastDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"full_cap_days": {
				Description: "The number of days remaining before the device reaches its full functional capacity.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"metric_name": {
				Description: "The name of the metric for which regression model is generated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"min_days_for_forecast": {
				Description: "The minimum number of days the HyperFlex cluster should be up for computing forecast.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Description: "Predictive model generated from the historical data of a device or source.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accuracy": {
							Description: "The standard error of the estimate is a measure of the accuracy of predictions from predective modeling.",
							Type:        schema.TypeFloat,
							Optional:    true,
						},
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
						"model_data": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeFloat}},
						"model_type": {
							Description: "Model type indicating type of predictive model used for computing forecast.\n* `Linear` - The Enum value Linear represents that the predictive model type used for forecast computation is linear regression.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"threshold_days": {
				Description: "The number of days remaining before the device reaches the specified threshold for the metric as defined in definition.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceForecastInstanceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ForecastInstance{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("full_cap_days"); ok {
		x := int64(v.(int))
		o.SetFullCapDays(x)
	}
	if v, ok := d.GetOk("metric_name"); ok {
		x := (v.(string))
		o.SetMetricName(x)
	}
	if v, ok := d.GetOk("min_days_for_forecast"); ok {
		x := int64(v.(int))
		o.SetMinDaysForForecast(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("threshold_days"); ok {
		x := int64(v.(int))
		o.SetThresholdDays(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ForecastInstance object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.ForecastApi.GetForecastInstanceList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching ForecastInstance: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for ForecastInstance list: %s", err.Error())
	}
	var s = &models.ForecastInstanceList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to ForecastInstance list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for ForecastInstance did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.ForecastInstance{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("alt_model", (s.GetAltModel())); err != nil {
				return diag.Errorf("error occurred while setting property AltModel: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("device_id", (s.GetDeviceId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceId: %s", err.Error())
			}

			if err := d.Set("forecast_def", flattenMapForecastDefinitionRelationship(s.GetForecastDef(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ForecastDef: %s", err.Error())
			}
			if err := d.Set("full_cap_days", (s.GetFullCapDays())); err != nil {
				return diag.Errorf("error occurred while setting property FullCapDays: %s", err.Error())
			}
			if err := d.Set("metric_name", (s.GetMetricName())); err != nil {
				return diag.Errorf("error occurred while setting property MetricName: %s", err.Error())
			}
			if err := d.Set("min_days_for_forecast", (s.GetMinDaysForForecast())); err != nil {
				return diag.Errorf("error occurred while setting property MinDaysForForecast: %s", err.Error())
			}

			if err := d.Set("model", flattenMapForecastModel(s.GetModel(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("threshold_days", (s.GetThresholdDays())); err != nil {
				return diag.Errorf("error occurred while setting property ThresholdDays: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
