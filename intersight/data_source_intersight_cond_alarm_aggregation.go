package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCondAlarmAggregation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCondAlarmAggregationRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"critical_alarms_count": {
				Description: "Count of all alarms with severity Critical, irrespective of acknowledgement status.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"health": {
				Description: "Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"info_alarms_count": {
				Description: "Count of all alarms with severity Info, irrespective of acknowledgement status.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"mo_type": {
				Description: "Managed object type. For example, FI managed object type will be network.Element.",
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
			"warning_alarms_count": {
				Description: "Count of all alarms with severity Warning, irrespective of acknowledgement status.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"alarm_aggregation_source": {
						Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"alarm_summary": {
						Description: "The summary of alarm counts based on the severity types (Critical or Warning).",
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
								"critical": {
									Description: "The count of alarms that have severity type Critical.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"warning": {
									Description: "The count of alarms that have severity type Warning.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"critical_alarms_count": {
						Description: "Count of all alarms with severity Critical, irrespective of acknowledgement status.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"health": {
						Description: "Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"info_alarms_count": {
						Description: "Count of all alarms with severity Info, irrespective of acknowledgement status.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"mo_type": {
						Description: "Managed object type. For example, FI managed object type will be network.Element.",
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
					"warning_alarms_count": {
						Description: "Count of all alarms with severity Warning, irrespective of acknowledgement status.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceCondAlarmAggregationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CondAlarmAggregation{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("critical_alarms_count"); ok {
		x := int64(v.(int))
		o.SetCriticalAlarmsCount(x)
	}
	if v, ok := d.GetOk("health"); ok {
		x := (v.(string))
		o.SetHealth(x)
	}
	if v, ok := d.GetOk("info_alarms_count"); ok {
		x := int64(v.(int))
		o.SetInfoAlarmsCount(x)
	}
	if v, ok := d.GetOk("mo_type"); ok {
		x := (v.(string))
		o.SetMoType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("warning_alarms_count"); ok {
		x := int64(v.(int))
		o.SetWarningAlarmsCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CondAlarmAggregation object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CondApi.GetCondAlarmAggregationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CondAlarmAggregation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CondAlarmAggregationList.GetCount()
	var i int32
	var condAlarmAggregationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CondApi.GetCondAlarmAggregationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CondAlarmAggregation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CondAlarmAggregationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CondAlarmAggregation data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["alarm_aggregation_source"] = flattenMapMoBaseMoRelationship(s.GetAlarmAggregationSource(), d)

				temp["alarm_summary"] = flattenMapCondAlarmSummary(s.GetAlarmSummary(), d)
				temp["class_id"] = (s.GetClassId())
				temp["critical_alarms_count"] = (s.GetCriticalAlarmsCount())
				temp["health"] = (s.GetHealth())
				temp["info_alarms_count"] = (s.GetInfoAlarmsCount())
				temp["mo_type"] = (s.GetMoType())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["warning_alarms_count"] = (s.GetWarningAlarmsCount())
				condAlarmAggregationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(condAlarmAggregationResults))
	if err := d.Set("results", condAlarmAggregationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(condAlarmAggregationResults[0]["moid"].(string))
	return de
}
