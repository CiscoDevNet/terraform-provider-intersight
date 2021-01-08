package intersight

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkflowRollbackWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowRollbackWorkflowCreate,
		ReadContext:   resourceWorkflowRollbackWorkflowRead,
		UpdateContext: resourceWorkflowRollbackWorkflowUpdate,
		DeleteContext: resourceWorkflowRollbackWorkflowDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "The action of the rollback workflow such as Create and Start.\n* `None` - If no action is set, then the default value is set to none for the action field.\n* `Create` - Create rollback workflow data for the execution of the rollback workflow.\n* `Start` - Start a new execution of the rollback workflow.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
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
				Computed:    true,
			},
			"continue_on_task_failure": {
				Description: "When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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

func resourceWorkflowRollbackWorkflowCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewWorkflowRollbackWorkflowWithDefaults()
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("workflow.RollbackWorkflow")

	if v, ok := d.GetOkExists("continue_on_task_failure"); ok {
		x := v.(bool)
		o.SetContinueOnTaskFailure(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("workflow.RollbackWorkflow")

	if v, ok := d.GetOk("primary_workflow"); ok {
		p := make([]models.WorkflowWorkflowInfoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowWorkflowInfoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrimaryWorkflow(x)
		}
	}

	if v, ok := d.GetOk("rollback_tasks"); ok {
		x := make([]models.WorkflowRollbackWorkflowTask, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowRollbackWorkflowTaskWithDefaults()
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
			o.SetClassId("workflow.RollbackWorkflowTask")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
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
			if v, ok := l["ref_name"]; ok {
				{
					x := (v.(string))
					o.SetRefName(x)
				}
			}
			if v, ok := l["rollback_completed"]; ok {
				{
					x := (v.(bool))
					o.SetRollbackCompleted(x)
				}
			}
			if v, ok := l["rollback_task_name"]; ok {
				{
					x := (v.(string))
					o.SetRollbackTaskName(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			if v, ok := l["task_info_moid"]; ok {
				{
					x := (v.(string))
					o.SetTaskInfoMoid(x)
				}
			}
			if v, ok := l["task_path"]; ok {
				{
					x := (v.(string))
					o.SetTaskPath(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetRollbackTasks(x)
		}
	}

	if v, ok := d.GetOk("rollback_workflows"); ok {
		x := make([]models.WorkflowWorkflowInfoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
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
			x = append(x, models.MoMoRefAsWorkflowWorkflowInfoRelationship(o))
		}
		if len(x) > 0 {
			o.SetRollbackWorkflows(x)
		}
	}

	if v, ok := d.GetOk("selected_tasks"); ok {
		x := make([]models.WorkflowRollbackWorkflowTask, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowRollbackWorkflowTaskWithDefaults()
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
			o.SetClassId("workflow.RollbackWorkflowTask")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
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
			if v, ok := l["ref_name"]; ok {
				{
					x := (v.(string))
					o.SetRefName(x)
				}
			}
			if v, ok := l["rollback_completed"]; ok {
				{
					x := (v.(bool))
					o.SetRollbackCompleted(x)
				}
			}
			if v, ok := l["rollback_task_name"]; ok {
				{
					x := (v.(string))
					o.SetRollbackTaskName(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			if v, ok := l["task_info_moid"]; ok {
				{
					x := (v.(string))
					o.SetTaskInfoMoid(x)
				}
			}
			if v, ok := l["task_path"]; ok {
				{
					x := (v.(string))
					o.SetTaskPath(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetSelectedTasks(x)
		}
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
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

	r := conn.ApiClient.WorkflowApi.CreateWorkflowRollbackWorkflow(conn.ctx).WorkflowRollbackWorkflow(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceWorkflowRollbackWorkflowRead(c, d, meta)
}

func resourceWorkflowRollbackWorkflowRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.WorkflowApi.GetWorkflowRollbackWorkflowByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "WorkflowRollbackWorkflow object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("action", (s.GetAction())); err != nil {
		return diag.Errorf("error occurred while setting property Action in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("continue_on_task_failure", (s.GetContinueOnTaskFailure())); err != nil {
		return diag.Errorf("error occurred while setting property ContinueOnTaskFailure in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("primary_workflow", flattenMapWorkflowWorkflowInfoRelationship(s.GetPrimaryWorkflow(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PrimaryWorkflow in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("rollback_tasks", flattenListWorkflowRollbackWorkflowTask(s.GetRollbackTasks(), d)); err != nil {
		return diag.Errorf("error occurred while setting property RollbackTasks in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("rollback_workflows", flattenListWorkflowWorkflowInfoRelationship(s.GetRollbackWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property RollbackWorkflows in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("selected_tasks", flattenListWorkflowRollbackWorkflowTask(s.GetSelectedTasks(), d)); err != nil {
		return diag.Errorf("error occurred while setting property SelectedTasks in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("status", (s.GetStatus())); err != nil {
		return diag.Errorf("error occurred while setting property Status in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in WorkflowRollbackWorkflow object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceWorkflowRollbackWorkflowUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.WorkflowRollbackWorkflow{}
	if d.HasChange("action") {
		v := d.Get("action")
		x := (v.(string))
		o.SetAction(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("workflow.RollbackWorkflow")

	if d.HasChange("continue_on_task_failure") {
		v := d.Get("continue_on_task_failure")
		x := (v.(bool))
		o.SetContinueOnTaskFailure(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("workflow.RollbackWorkflow")

	if d.HasChange("primary_workflow") {
		v := d.Get("primary_workflow")
		p := make([]models.WorkflowWorkflowInfoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowWorkflowInfoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrimaryWorkflow(x)
		}
	}

	if d.HasChange("rollback_tasks") {
		v := d.Get("rollback_tasks")
		x := make([]models.WorkflowRollbackWorkflowTask, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowRollbackWorkflowTask{}
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
			o.SetClassId("workflow.RollbackWorkflowTask")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
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
			if v, ok := l["ref_name"]; ok {
				{
					x := (v.(string))
					o.SetRefName(x)
				}
			}
			if v, ok := l["rollback_completed"]; ok {
				{
					x := (v.(bool))
					o.SetRollbackCompleted(x)
				}
			}
			if v, ok := l["rollback_task_name"]; ok {
				{
					x := (v.(string))
					o.SetRollbackTaskName(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			if v, ok := l["task_info_moid"]; ok {
				{
					x := (v.(string))
					o.SetTaskInfoMoid(x)
				}
			}
			if v, ok := l["task_path"]; ok {
				{
					x := (v.(string))
					o.SetTaskPath(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetRollbackTasks(x)
		}
	}

	if d.HasChange("rollback_workflows") {
		v := d.Get("rollback_workflows")
		x := make([]models.WorkflowWorkflowInfoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
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
			x = append(x, models.MoMoRefAsWorkflowWorkflowInfoRelationship(o))
		}
		if len(x) > 0 {
			o.SetRollbackWorkflows(x)
		}
	}

	if d.HasChange("selected_tasks") {
		v := d.Get("selected_tasks")
		x := make([]models.WorkflowRollbackWorkflowTask, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowRollbackWorkflowTask{}
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
			o.SetClassId("workflow.RollbackWorkflowTask")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
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
			if v, ok := l["ref_name"]; ok {
				{
					x := (v.(string))
					o.SetRefName(x)
				}
			}
			if v, ok := l["rollback_completed"]; ok {
				{
					x := (v.(bool))
					o.SetRollbackCompleted(x)
				}
			}
			if v, ok := l["rollback_task_name"]; ok {
				{
					x := (v.(string))
					o.SetRollbackTaskName(x)
				}
			}
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			if v, ok := l["task_info_moid"]; ok {
				{
					x := (v.(string))
					o.SetTaskInfoMoid(x)
				}
			}
			if v, ok := l["task_path"]; ok {
				{
					x := (v.(string))
					o.SetTaskPath(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetSelectedTasks(x)
		}
	}

	if d.HasChange("status") {
		v := d.Get("status")
		x := (v.(string))
		o.SetStatus(x)
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

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowRollbackWorkflow(conn.ctx, d.Id()).WorkflowRollbackWorkflow(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceWorkflowRollbackWorkflowRead(c, d, meta)
}

func resourceWorkflowRollbackWorkflowDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.WorkflowApi.DeleteWorkflowRollbackWorkflow(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting WorkflowRollbackWorkflow object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
