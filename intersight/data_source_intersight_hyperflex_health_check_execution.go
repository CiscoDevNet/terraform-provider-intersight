package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexHealthCheckExecution() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthCheckExecutionRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
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
			"uuid": {
				Description: "UUID of an instance of health check execution.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceHyperflexHealthCheckExecutionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckExecution{}
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
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHealthCheckExecution object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckExecutionList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching HyperflexHealthCheckExecution: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexHealthCheckExecution list: %s", err.Error())
	}
	var s = &models.HyperflexHealthCheckExecutionList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexHealthCheckExecution list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexHealthCheckExecution data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexHealthCheckExecution data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexHealthCheckExecution{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("category", (s.GetCategory())); err != nil {
				return diag.Errorf("error occurred while setting property Category: %s", err.Error())
			}
			if err := d.Set("cause", (s.GetCause())); err != nil {
				return diag.Errorf("error occurred while setting property Cause: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("completion_time", (s.GetCompletionTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property CompletionTime: %s", err.Error())
			}

			if err := d.Set("health_check_definition", flattenMapHyperflexHealthCheckDefinitionRelationship(s.GetHealthCheckDefinition(), d)); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckDefinition: %s", err.Error())
			}
			if err := d.Set("health_check_details", (s.GetHealthCheckDetails())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckDetails: %s", err.Error())
			}
			if err := d.Set("health_check_execution_error_details", (s.GetHealthCheckExecutionErrorDetails())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckExecutionErrorDetails: %s", err.Error())
			}
			if err := d.Set("health_check_execution_error_summary", (s.GetHealthCheckExecutionErrorSummary())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckExecutionErrorSummary: %s", err.Error())
			}
			if err := d.Set("health_check_execution_status", (s.GetHealthCheckExecutionStatus())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckExecutionStatus: %s", err.Error())
			}
			if err := d.Set("health_check_result", (s.GetHealthCheckResult())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckResult: %s", err.Error())
			}
			if err := d.Set("health_check_summary", (s.GetHealthCheckSummary())); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckSummary: %s", err.Error())
			}

			if err := d.Set("hx_cluster", flattenMapHyperflexClusterRelationship(s.GetHxCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property HxCluster: %s", err.Error())
			}
			if err := d.Set("hx_device_name", (s.GetHxDeviceName())); err != nil {
				return diag.Errorf("error occurred while setting property HxDeviceName: %s", err.Error())
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
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
