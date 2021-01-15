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

func dataSourceWorkflowBatchApiExecutor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowBatchApiExecutorRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"batch": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"asset_target_moid": {
							Description: "Asset target defines the remote target endpoints which are either managed by\nIntersight or their service APIs are invoked from Intersight. Generic API executor\nservice Jasmine can invoke these remote APIs via its executors. Asset targets can be\naccessed directly for cloud targets or via an associated Intersight Assist. Prerequisite\nto use asset targets is to persist the target details.\nAsset target MoRef can be given as task input of type TargetDataType. Fusion determines\nand populates the target context with the Assist if associated with. It is set\ninternally at the API level.\nIn case of an associated assist, it is used by Assist to fetch the target details\nand form the API request to send to endpoints. In case of cloud asset targets, Jasmine\nfetched the target details from DB, forms the API request and sends it to the target.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"body": {
							Description: "The optional request body that is sent as part of this API request.\nThe request body can contain a golang template that can be populated with task input\nparameters and previous API output parameters.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"content_type": {
							Description: "Intersight Orchestrator, with the support of response parser specification,\ncan extract the values from API responses and map them to task output parameters.\nThe value extraction is supported for response content types XML and JSON.\nThe type of the content that gets passed as payload and response in this\nAPI.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"description": {
							Description: "A description that task designer can add to individual API requests that explain \nwhat the API call is about.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user friendly label that task designers have given to the batch API request.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "A reference name for this API request within the batch API request.\nThis name shall be used to map the API output parameters to subsequent\nAPI input parameters within a batch API task.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"outcomes": {
							Description: "All the possible outcomes of this API are captured here. Outcomes property\nis a collection property of type workflow.Outcome objects.\nThe outcomes can be mapped to the message to be shown. The outcomes are\nevaluated in the order they are given. At the end of the outcomes list,\nan catchall success/fail outcome can be added with condition as 'true'.\nThis is an optional\nproperty and if not specified the task will be marked as success.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
						"response_spec": {
							Description: "The optional grammar specification for parsing the response to extract the\nrequired values.\nThe specification should have extraction specification specified for\nall the API output parameters.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
						"skip_on_condition": {
							Description: "The skip expression, if provided, allows the batch API executor to skip the\napi execution when the given expression evaluates to true.\nThe expression is given as such a golang template that has to be\nevaluated to a final content true/false. The expression is an optional and in\ncase not provided, the API will always be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"start_delay": {
							Description: "The delay in seconds after which the API needs to be executed.\nBy default, the given API is executed immediately. Specifying a start delay adds to the delay to execution.\nStart Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"timeout": {
							Description: "The duration in seconds by which the API response is expected from the API target.\nIf the end point does not respond for the API request within this timeout\nduration, the task will be marked as failed.",
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
			"constraints": {
				Description: "Enter the constraints on when this task should match against the task definition.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"target_data_type": {
							Description: "List of property constraints that helps to narrow down task implementations based on target device input.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "A detailed description about the batch APIs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"error_response_handler": {
				Description: "A reference to a workflowErrorResponseHandler resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name for the batch API task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"outcomes": {
				Description: "All the possible outcomes of this task are captured here. Outcomes property\nis a collection property of type workflow.Outcome objects.\nThe outcomes can be mapped to the message to be shown. The outcomes are\nevaluated in the order they are given. At the end of the outcomes list,\nan catchall success/fail outcome can be added with condition as 'true'.\nThis is an optional\nproperty and if not specified the task will be marked as success.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"output": {
				Description: "Intersight Orchestrator allows the extraction of required values from API\nresponses using the API response grammar. These extracted values can be mapped\nto task output parameters defined in task definition.\nThe mapping of API output parameters to the task output parameters is provided\nas JSON in this property.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"retry_from_failed_api": {
				Description: "When an execution of a nth API in the Batch fails,\nRetry from falied API flag indicates if the execution should start from the nth API or the first API during task retry.\nBy default the value is set to false.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"skip_on_condition": {
				Description: "The skip expression, if provided, allows the batch API executor to skip the\ntask execution when the given expression evaluates to true.\nThe expression is given as such a golang template that has to be\nevaluated to a final content true/false. The expression is an optional and in\ncase not provided, the API will always be executed.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"task_definition": {
				Description: "A reference to a workflowTaskDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
		},
	}
}

func dataSourceWorkflowBatchApiExecutorRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowBatchApiExecutor{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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
	if v, ok := d.GetOk("retry_from_failed_api"); ok {
		x := (v.(bool))
		o.SetRetryFromFailedApi(x)
	}
	if v, ok := d.GetOk("skip_on_condition"); ok {
		x := (v.(string))
		o.SetSkipOnCondition(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowBatchApiExecutor object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowBatchApiExecutorList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching WorkflowBatchApiExecutor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for WorkflowBatchApiExecutor list: %s", err.Error())
	}
	var s = &models.WorkflowBatchApiExecutorList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to WorkflowBatchApiExecutor list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for WorkflowBatchApiExecutor data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for WorkflowBatchApiExecutor data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.WorkflowBatchApiExecutor{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("batch", flattenListWorkflowApi(s.GetBatch(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Batch: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("constraints", flattenMapWorkflowTaskConstraints(s.GetConstraints(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Constraints: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}

			if err := d.Set("error_response_handler", flattenMapWorkflowErrorResponseHandlerRelationship(s.GetErrorResponseHandler(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ErrorResponseHandler: %s", err.Error())
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
			if err := d.Set("retry_from_failed_api", (s.GetRetryFromFailedApi())); err != nil {
				return diag.Errorf("error occurred while setting property RetryFromFailedApi: %s", err.Error())
			}
			if err := d.Set("skip_on_condition", (s.GetSkipOnCondition())); err != nil {
				return diag.Errorf("error occurred while setting property SkipOnCondition: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("task_definition", flattenMapWorkflowTaskDefinitionRelationship(s.GetTaskDefinition(), d)); err != nil {
				return diag.Errorf("error occurred while setting property TaskDefinition: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
