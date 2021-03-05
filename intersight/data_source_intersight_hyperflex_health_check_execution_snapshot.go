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

func dataSourceHyperflexHealthCheckExecutionSnapshot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthCheckExecutionSnapshotRead,
		Schema: map[string]*schema.Schema{
			"category": {
				Description: "Category that the HyperFlex health check Definition belongs to.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cause": {
				Description: "Information detailing the possible cause of the healthcheck failure, if the check fails.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"completion_time": {
				Description: "Health check execution completion time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_details": {
				Description: "Details of the health check execution result.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_execution_error_details": {
				Description: "Error details of a script execution failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_execution_error_summary": {
				Description: "Error summary of a script execution failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_execution_status": {
				Description: "Status of the health check execution.\n* `UNKNOWN` - Indicates that the health heck execution results are unknown.\n* `SUCCEEDED` - Indicates that the health check execution succeeded.\n* `FAILED` - Indicates that the health check execution failed.\n* `TIMED_OUT` - Indicates that the health check execution timed out before completion.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_result": {
				Description: "Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED.\n* `UNKNOWN` - Indicates that the health check results could not be determined.\n* `PASS` - Indicates that the health check passed.\n* `FAIL` - Indicates that the health check failed.\n* `WARN` - Indicates that the health check completed with a warning.\n* `NOT_APPLICABLE` - Indicates that the health check is either unsupported, or not applicable on the Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_check_summary": {
				Description: "A brief summary of health check results.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hx_device_name": {
				Description: "HyperFlex Device Name where the healthcheck is executed.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
					"category": {
						Description: "Category that the HyperFlex health check Definition belongs to.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"cause": {
						Description: "Information detailing the possible cause of the healthcheck failure, if the check fails.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"completion_time": {
						Description: "Health check execution completion time.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_definition": {
						Description: "A reference to a hyperflexHealthCheckDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"health_check_details": {
						Description: "Details of the health check execution result.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_execution_error_details": {
						Description: "Error details of a script execution failure.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_execution_error_summary": {
						Description: "Error summary of a script execution failure.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_execution_status": {
						Description: "Status of the health check execution.\n* `UNKNOWN` - Indicates that the health heck execution results are unknown.\n* `SUCCEEDED` - Indicates that the health check execution succeeded.\n* `FAILED` - Indicates that the health check execution failed.\n* `TIMED_OUT` - Indicates that the health check execution timed out before completion.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_result": {
						Description: "Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED.\n* `UNKNOWN` - Indicates that the health check results could not be determined.\n* `PASS` - Indicates that the health check passed.\n* `FAIL` - Indicates that the health check failed.\n* `WARN` - Indicates that the health check completed with a warning.\n* `NOT_APPLICABLE` - Indicates that the health check is either unsupported, or not applicable on the Cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"health_check_summary": {
						Description: "A brief summary of health check results.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"hx_cluster": {
						Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"hx_device_name": {
						Description: "HyperFlex Device Name where the healthcheck is executed.",
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
					"workflow": {
						Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexHealthCheckExecutionSnapshotRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckExecutionSnapshot{}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("cause"); ok {
		x := (v.(string))
		o.SetCause(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("completion_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCompletionTime(x)
	}
	if v, ok := d.GetOk("health_check_details"); ok {
		x := (v.(string))
		o.SetHealthCheckDetails(x)
	}
	if v, ok := d.GetOk("health_check_execution_error_details"); ok {
		x := (v.(string))
		o.SetHealthCheckExecutionErrorDetails(x)
	}
	if v, ok := d.GetOk("health_check_execution_error_summary"); ok {
		x := (v.(string))
		o.SetHealthCheckExecutionErrorSummary(x)
	}
	if v, ok := d.GetOk("health_check_execution_status"); ok {
		x := (v.(string))
		o.SetHealthCheckExecutionStatus(x)
	}
	if v, ok := d.GetOk("health_check_result"); ok {
		x := (v.(string))
		o.SetHealthCheckResult(x)
	}
	if v, ok := d.GetOk("health_check_summary"); ok {
		x := (v.(string))
		o.SetHealthCheckSummary(x)
	}
	if v, ok := d.GetOk("hx_device_name"); ok {
		x := (v.(string))
		o.SetHxDeviceName(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHealthCheckExecutionSnapshot object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckExecutionSnapshotList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexHealthCheckExecutionSnapshot: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexHealthCheckExecutionSnapshotList.GetCount()
	var i int32
	var hyperflexHealthCheckExecutionSnapshotResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckExecutionSnapshotList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexHealthCheckExecutionSnapshot: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexHealthCheckExecutionSnapshotList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexHealthCheckExecutionSnapshot data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["category"] = (s.GetCategory())
				temp["cause"] = (s.GetCause())
				temp["class_id"] = (s.GetClassId())

				temp["completion_time"] = (s.GetCompletionTime()).String()

				temp["health_check_definition"] = flattenMapHyperflexHealthCheckDefinitionRelationship(s.GetHealthCheckDefinition(), d)
				temp["health_check_details"] = (s.GetHealthCheckDetails())
				temp["health_check_execution_error_details"] = (s.GetHealthCheckExecutionErrorDetails())
				temp["health_check_execution_error_summary"] = (s.GetHealthCheckExecutionErrorSummary())
				temp["health_check_execution_status"] = (s.GetHealthCheckExecutionStatus())
				temp["health_check_result"] = (s.GetHealthCheckResult())
				temp["health_check_summary"] = (s.GetHealthCheckSummary())

				temp["hx_cluster"] = flattenMapHyperflexClusterRelationship(s.GetHxCluster(), d)
				temp["hx_device_name"] = (s.GetHxDeviceName())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["workflow"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflow(), d)
				hyperflexHealthCheckExecutionSnapshotResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexHealthCheckExecutionSnapshotResults))
	if err := d.Set("results", hyperflexHealthCheckExecutionSnapshotResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexHealthCheckExecutionSnapshotResults[0]["moid"].(string))
	return de
}
