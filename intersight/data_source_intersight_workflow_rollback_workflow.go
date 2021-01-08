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

func dataSourceWorkflowRollbackWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowRollbackWorkflowRead,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "The action of the rollback workflow such as Create and Start.\n* `None` - If no action is set, then the default value is set to none for the action field.\n* `Create` - Create rollback workflow data for the execution of the rollback workflow.\n* `Start` - Start a new execution of the rollback workflow.",
				Type:        schema.TypeString,
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
			"continue_on_task_failure": {
				Description: "When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"primary_workflow": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"rollback_tasks": {
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
						"description": {
							Description: "Description of the rollback task.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of TaskInfo that needs to be rolled back.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ref_name": {
							Description: "Reference name of TaskInfo that need to be rolled back.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rollback_completed": {
							Description: "Status of the rollback operation for the task.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"rollback_task_name": {
							Description: "Name of TaskInfo that performs the rollback operation.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Description: "Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed.\n* `NotStarted` - Status of rollback task when it is not started rollback.\n* `Not supported` - Status of task when it is not supporting rollback.\n* `Completed` - Status of rollback task once execution is successful.\n* `Failed` - Status of rollback task when it is failed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"task_info_moid": {
							Description: "Moid of TaskInfo that supports rollback operation.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"task_path": {
							Description: "Path of rollback task if it is inside sub-workflow.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"rollback_workflows": {
				Description: "An array of relationships to workflowWorkflowInfo resources.",
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
			"selected_tasks": {
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
						"description": {
							Description: "Description of the rollback task.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of TaskInfo that needs to be rolled back.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"ref_name": {
							Description: "Reference name of TaskInfo that need to be rolled back.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rollback_completed": {
							Description: "Status of the rollback operation for the task.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"rollback_task_name": {
							Description: "Name of TaskInfo that performs the rollback operation.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Description: "Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed.\n* `NotStarted` - Status of rollback task when it is not started rollback.\n* `Not supported` - Status of task when it is not supporting rollback.\n* `Completed` - Status of rollback task once execution is successful.\n* `Failed` - Status of rollback task when it is failed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"task_info_moid": {
							Description: "Moid of TaskInfo that supports rollback operation.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"task_path": {
							Description: "Path of rollback task if it is inside sub-workflow.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"status": {
				Description: "Status of the rollback workflow instance (Created, Running, Completed, Failed).\n* `None` - If no status is set, then the default value is set none for the status field.\n* `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback.\n* `Running` - Status of the rollback workflow when it is in-progress.\n* `Completed` - Status of the rollback workflow after execution is successful.\n* `Failed` - Status of the rollback workflow after execution results in failure.",
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
		},
	}
}

func dataSourceWorkflowRollbackWorkflowRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowRollbackWorkflow{}
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("continue_on_task_failure"); ok {
		x := (v.(bool))
		o.SetContinueOnTaskFailure(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowRollbackWorkflow object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowRollbackWorkflowList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for WorkflowRollbackWorkflow list: %s", err.Error())
	}
	var s = &models.WorkflowRollbackWorkflowList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to WorkflowRollbackWorkflow list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for WorkflowRollbackWorkflow did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.WorkflowRollbackWorkflow{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("action", (s.GetAction())); err != nil {
				return diag.Errorf("error occurred while setting property Action: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("continue_on_task_failure", (s.GetContinueOnTaskFailure())); err != nil {
				return diag.Errorf("error occurred while setting property ContinueOnTaskFailure: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("primary_workflow", flattenMapWorkflowWorkflowInfoRelationship(s.GetPrimaryWorkflow(), d)); err != nil {
				return diag.Errorf("error occurred while setting property PrimaryWorkflow: %s", err.Error())
			}

			if err := d.Set("rollback_tasks", flattenListWorkflowRollbackWorkflowTask(s.GetRollbackTasks(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RollbackTasks: %s", err.Error())
			}

			if err := d.Set("rollback_workflows", flattenListWorkflowWorkflowInfoRelationship(s.GetRollbackWorkflows(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RollbackWorkflows: %s", err.Error())
			}

			if err := d.Set("selected_tasks", flattenListWorkflowRollbackWorkflowTask(s.GetSelectedTasks(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SelectedTasks: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
