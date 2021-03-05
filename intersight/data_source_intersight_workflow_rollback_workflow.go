package intersight

import (
	"context"
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the rollback workflow instance (Created, Running, Completed, Failed).\n* `None` - If no status is set, then the default value is set none for the status field.\n* `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback.\n* `Running` - Status of the rollback workflow when it is in-progress.\n* `Completed` - Status of the rollback workflow after execution is successful.\n* `Failed` - Status of the rollback workflow after execution results in failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceWorkflowRollbackWorkflow().Schema},
				Computed: true,
			}},
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
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowRollbackWorkflowList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.WorkflowRollbackWorkflowList.GetCount()
	var i int32
	var workflowRollbackWorkflowResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowRollbackWorkflowList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowRollbackWorkflow: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.WorkflowRollbackWorkflowList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for WorkflowRollbackWorkflow data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["continue_on_task_failure"] = (s.GetContinueOnTaskFailure())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["primary_workflow"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetPrimaryWorkflow(), d)

				temp["rollback_tasks"] = flattenListWorkflowRollbackWorkflowTask(s.GetRollbackTasks(), d)

				temp["rollback_workflows"] = flattenListWorkflowWorkflowInfoRelationship(s.GetRollbackWorkflows(), d)

				temp["selected_tasks"] = flattenListWorkflowRollbackWorkflowTask(s.GetSelectedTasks(), d)
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				workflowRollbackWorkflowResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowRollbackWorkflowResults))
	if err := d.Set("results", workflowRollbackWorkflowResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowRollbackWorkflowResults[0]["moid"].(string))
	return de
}
