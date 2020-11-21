package intersight

import (
	"encoding/json"
	"fmt"
	"log"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkflowBatchApiExecutor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowBatchApiExecutorCreate,
		Read:   resourceWorkflowBatchApiExecutorRead,
		Update: resourceWorkflowBatchApiExecutorUpdate,
		Delete: resourceWorkflowBatchApiExecutorDelete,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name for the batch API task.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
		},
	}
}

func resourceWorkflowBatchApiExecutorCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowBatchApiExecutorWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("batch"); ok {
		x := make([]models.WorkflowApi, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowApiWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["asset_target_moid"]; ok {
				{
					x := (v.(string))
					o.SetAssetTargetMoid(x)
				}
			}
			if v, ok := l["body"]; ok {
				{
					x := (v.(string))
					o.SetBody(x)
				}
			}
			o.SetClassId("workflow.Api")
			if v, ok := l["content_type"]; ok {
				{
					x := (v.(string))
					o.SetContentType(x)
				}
			}
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["label"]; ok {
				{
					x := (v.(string))
					o.SetLabel(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["outcomes"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetOutcomes(x)
				}
			}
			if v, ok := l["response_spec"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetResponseSpec(x)
				}
			}
			if v, ok := l["skip_on_condition"]; ok {
				{
					x := (v.(string))
					o.SetSkipOnCondition(x)
				}
			}
			if v, ok := l["start_delay"]; ok {
				{
					x := int64(v.(int))
					o.SetStartDelay(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetBatch(x)
		}
	}

	o.SetClassId("workflow.BatchApiExecutor")

	if v, ok := d.GetOk("constraints"); ok {
		p := make([]models.WorkflowTaskConstraints, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewWorkflowTaskConstraintsWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("workflow.TaskConstraints")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["target_data_type"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetTargetDataType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetConstraints(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("error_response_handler"); ok {
		p := make([]models.WorkflowErrorResponseHandlerRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsWorkflowErrorResponseHandlerRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetErrorResponseHandler(x)
		}
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workflow.BatchApiExecutor")

	if v, ok := d.GetOk("outcomes"); ok {
		x := v.(map[string]interface{})
		o.SetOutcomes(x)
	}

	if v, ok := d.GetOk("output"); ok {
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if v, ok := d.GetOkExists("retry_from_failed_api"); ok {
		x := v.(bool)
		o.SetRetryFromFailedApi(x)
	}

	if v, ok := d.GetOk("skip_on_condition"); ok {
		x := (v.(string))
		o.SetSkipOnCondition(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("task_definition"); ok {
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsWorkflowTaskDefinitionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetTaskDefinition(x)
		}
	}

	r := conn.ApiClient.WorkflowApi.CreateWorkflowBatchApiExecutor(conn.ctx).WorkflowBatchApiExecutor(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.WorkflowApi.GetWorkflowBatchApiExecutorByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("batch", flattenListWorkflowApi(s.GetBatch(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Batch: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("constraints", flattenMapWorkflowTaskConstraints(s.GetConstraints(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Constraints: %+v", err)
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return fmt.Errorf("error occurred while setting property Description: %+v", err)
	}

	if err := d.Set("error_response_handler", flattenMapWorkflowErrorResponseHandlerRelationship(s.GetErrorResponseHandler(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property ErrorResponseHandler: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return fmt.Errorf("error occurred while setting property Name: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("outcomes", (s.GetOutcomes())); err != nil {
		return fmt.Errorf("error occurred while setting property Outcomes: %+v", err)
	}

	if err := d.Set("output", (s.GetOutput())); err != nil {
		return fmt.Errorf("error occurred while setting property Output: %+v", err)
	}

	if err := d.Set("retry_from_failed_api", (s.GetRetryFromFailedApi())); err != nil {
		return fmt.Errorf("error occurred while setting property RetryFromFailedApi: %+v", err)
	}

	if err := d.Set("skip_on_condition", (s.GetSkipOnCondition())); err != nil {
		return fmt.Errorf("error occurred while setting property SkipOnCondition: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	if err := d.Set("task_definition", flattenMapWorkflowTaskDefinitionRelationship(s.GetTaskDefinition(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property TaskDefinition: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceWorkflowBatchApiExecutorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.WorkflowBatchApiExecutor{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("batch") {
		v := d.Get("batch")
		x := make([]models.WorkflowApi, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowApi{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["asset_target_moid"]; ok {
				{
					x := (v.(string))
					o.SetAssetTargetMoid(x)
				}
			}
			if v, ok := l["body"]; ok {
				{
					x := (v.(string))
					o.SetBody(x)
				}
			}
			o.SetClassId("workflow.Api")
			if v, ok := l["content_type"]; ok {
				{
					x := (v.(string))
					o.SetContentType(x)
				}
			}
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["label"]; ok {
				{
					x := (v.(string))
					o.SetLabel(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["outcomes"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetOutcomes(x)
				}
			}
			if v, ok := l["response_spec"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetResponseSpec(x)
				}
			}
			if v, ok := l["skip_on_condition"]; ok {
				{
					x := (v.(string))
					o.SetSkipOnCondition(x)
				}
			}
			if v, ok := l["start_delay"]; ok {
				{
					x := int64(v.(int))
					o.SetStartDelay(x)
				}
			}
			if v, ok := l["timeout"]; ok {
				{
					x := int64(v.(int))
					o.SetTimeout(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetBatch(x)
		}
	}

	o.SetClassId("workflow.BatchApiExecutor")

	if d.HasChange("constraints") {
		v := d.Get("constraints")
		p := make([]models.WorkflowTaskConstraints, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.WorkflowTaskConstraints{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("workflow.TaskConstraints")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["target_data_type"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetTargetDataType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetConstraints(x)
		}
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("error_response_handler") {
		v := d.Get("error_response_handler")
		p := make([]models.WorkflowErrorResponseHandlerRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsWorkflowErrorResponseHandlerRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetErrorResponseHandler(x)
		}
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workflow.BatchApiExecutor")

	if d.HasChange("outcomes") {
		v := d.Get("outcomes")
		x := v.(map[string]interface{})
		o.SetOutcomes(x)
	}

	if d.HasChange("output") {
		v := d.Get("output")
		x := v.(map[string]interface{})
		o.SetOutput(x)
	}

	if d.HasChange("retry_from_failed_api") {
		v := d.Get("retry_from_failed_api")
		x := (v.(bool))
		o.SetRetryFromFailedApi(x)
	}

	if d.HasChange("skip_on_condition") {
		v := d.Get("skip_on_condition")
		x := (v.(string))
		o.SetSkipOnCondition(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("task_definition") {
		v := d.Get("task_definition")
		p := make([]models.WorkflowTaskDefinitionRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsWorkflowTaskDefinitionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetTaskDefinition(x)
		}
	}

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowBatchApiExecutor(conn.ctx, d.Id()).WorkflowBatchApiExecutor(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowBatchApiExecutorRead(d, meta)
}

func resourceWorkflowBatchApiExecutorDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	p := conn.ApiClient.WorkflowApi.DeleteWorkflowBatchApiExecutor(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
