package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceAppStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceAppStatusRead,
		Schema: map[string]*schema.Schema{
			"app_label": {
				Description: "Unique label to identify the application.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"operational_status": {
				Description: "Operational status of the application.\nOperational status is based on the result of the status\nchecks. If result of any check is Critical, then its\nvalue is Impaired. Otherwise, if result of any check is\nWarning, then its value is AttentionNeeded. If all\nchecks are OK, then its value is Operational.\n* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.\n* `Operational` - Operational status of the Intersight Appliance entity is Operational.\n* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.\n* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ready_count": {
				Description: "Number of replicas ready.  The number of instances of\nthe application currently ready to perform its intended\nfunctions.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"replica_count": {
				Description: "Number of replicas provisioned. The number of instances\nof the application provisioned to run on the Intersight appliance.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"restart_count1_hour": {
				Description: "Number of instance restarts in the last hour.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"restart_count24_hours": {
				Description: "Number of instance restarts in the last 24 hours.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"restart_count5_mins": {
				Description: "Number of instance restarts in the last 5 minutes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"restart_count_total": {
				Description: "Total number of restarts since last deployment.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"running_count": {
				Description: "Number of replicas running. The number of instances of\nthe application currently running.",
				Type:        schema.TypeInt,
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
					"api_statuses": {
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"elapsed_time": {
									Description: "The elapsed time for query in seconds.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"object_name": {
									Description: "Name to identify the API.",
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
								"reason": {
									Description: "Reason to address why the API call failed, if API call was successed, reason would be null.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"response": {
									Description: "Response code of the API call.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"app_label": {
						Description: "Unique label to identify the application.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"group_status": {
						Description: "A reference to a applianceGroupStatus resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"operational_status": {
						Description: "Operational status of the application.\nOperational status is based on the result of the status\nchecks. If result of any check is Critical, then its\nvalue is Impaired. Otherwise, if result of any check is\nWarning, then its value is AttentionNeeded. If all\nchecks are OK, then its value is Operational.\n* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.\n* `Operational` - Operational status of the Intersight Appliance entity is Operational.\n* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.\n* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ready_count": {
						Description: "Number of replicas ready.  The number of instances of\nthe application currently ready to perform its intended\nfunctions.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"replica_count": {
						Description: "Number of replicas provisioned. The number of instances\nof the application provisioned to run on the Intersight appliance.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"restart_count1_hour": {
						Description: "Number of instance restarts in the last hour.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"restart_count24_hours": {
						Description: "Number of instance restarts in the last 24 hours.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"restart_count5_mins": {
						Description: "Number of instance restarts in the last 5 minutes.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"restart_count_total": {
						Description: "Total number of restarts since last deployment.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"running_count": {
						Description: "Number of replicas running. The number of instances of\nthe application currently running.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"status_checks": {
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"code": {
									Description: "Unique identifier of the status check.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"result": {
									Description: "Result of this status check.\n* `OK` - Result of the check is OK.\n* `Warning` - Result of the check is Warning.\n* `Critical` - Result of the check is Critical.\n* `Info` - Result of the check is low.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"system_status": {
						Description: "A reference to a applianceSystemStatus resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceApplianceAppStatusRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceAppStatus{}
	if v, ok := d.GetOk("app_label"); ok {
		x := (v.(string))
		o.SetAppLabel(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("operational_status"); ok {
		x := (v.(string))
		o.SetOperationalStatus(x)
	}
	if v, ok := d.GetOk("ready_count"); ok {
		x := int64(v.(int))
		o.SetReadyCount(x)
	}
	if v, ok := d.GetOk("replica_count"); ok {
		x := int64(v.(int))
		o.SetReplicaCount(x)
	}
	if v, ok := d.GetOk("restart_count1_hour"); ok {
		x := int64(v.(int))
		o.SetRestartCount1Hour(x)
	}
	if v, ok := d.GetOk("restart_count24_hours"); ok {
		x := int64(v.(int))
		o.SetRestartCount24Hours(x)
	}
	if v, ok := d.GetOk("restart_count5_mins"); ok {
		x := int64(v.(int))
		o.SetRestartCount5Mins(x)
	}
	if v, ok := d.GetOk("restart_count_total"); ok {
		x := int64(v.(int))
		o.SetRestartCountTotal(x)
	}
	if v, ok := d.GetOk("running_count"); ok {
		x := int64(v.(int))
		o.SetRunningCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceAppStatus object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceAppStatusList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceAppStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceAppStatusList.GetCount()
	var i int32
	var applianceAppStatusResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceAppStatusList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceAppStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceAppStatusList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceAppStatus data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["api_statuses"] = flattenListApplianceApiStatus(s.GetApiStatuses(), d)
				temp["app_label"] = (s.GetAppLabel())
				temp["class_id"] = (s.GetClassId())

				temp["group_status"] = flattenMapApplianceGroupStatusRelationship(s.GetGroupStatus(), d)
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["operational_status"] = (s.GetOperationalStatus())
				temp["ready_count"] = (s.GetReadyCount())
				temp["replica_count"] = (s.GetReplicaCount())
				temp["restart_count1_hour"] = (s.GetRestartCount1Hour())
				temp["restart_count24_hours"] = (s.GetRestartCount24Hours())
				temp["restart_count5_mins"] = (s.GetRestartCount5Mins())
				temp["restart_count_total"] = (s.GetRestartCountTotal())
				temp["running_count"] = (s.GetRunningCount())

				temp["status_checks"] = flattenListApplianceStatusCheck(s.GetStatusChecks(), d)

				temp["system_status"] = flattenMapApplianceSystemStatusRelationship(s.GetSystemStatus(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				applianceAppStatusResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceAppStatusResults))
	if err := d.Set("results", applianceAppStatusResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceAppStatusResults[0]["moid"].(string))
	return de
}
