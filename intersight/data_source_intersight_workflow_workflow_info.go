package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getWorkflowWorkflowInfoSchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account": {
		Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				},
				"object_type": {
					Description: "The fully-qualified name of the remote type referred by this relationship.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"selector": {
					Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
					Type:        schema.TypeString,
					Optional:    true,
				},
			},
		},
	},
		"account_moid": {
			Description: "The Account ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"action": {
			Description: "The action of the workflow such as start, cancel, retry, pause.\n* `None` - No action is set, this is the default value for action field.\n* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.\n* `Start` - Start a new execution of the workflow.\n* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system.\n* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.\n* `Rerun` - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action.\n* `Retry` - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.\n* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.\n* `RetryFromTask` - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt.\n* `Cancel` - Cancel the workflow that is in running or waiting state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"associated_object": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cleanup_time": {
			Description: "The time when the workflow info will be removed from the database. When WorkflowInfo is created, cleanup time will be set to 181 days. As the workflow progresses through different states the cleanup time can be updated. A cleanup time of 0 means the workflow is not scheduled for cleanup. An active workflow that continues to schedule & run tasks can run for any amount of time and there is no upper bound for such workflows. Workflows that are not actively running, say in Paused or Waiting states will be removed after 181 days.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"email": {
			Description: "The email address of the user who started this workflow. In the case of LDAP users, this field can hold either a username or an email.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"end_time": {
			Description: "The time when the workflow reached a final state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"failed_workflow_cleanup_duration": {
			Description: "The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"input": {
			Description: "All the given inputs for the workflow. The schema for the inputs is defined in the InputDefinition section of the WorkflowDefinition. The InputDefinition will provide a list of input fields to be accepted, the associated datatype of the inputs and any additional constraints on the inputs. For more information please refer to InputDefinition property in the the the WorkflowDefinition resource. The inputs for a workflow are provided as a collection of key-value pairs, where key is the name of the input and value is any valid JSON data which conforms to the datatype of the input as specified in the InputDefinition. When the input passed into a workflow does not match the datatype or the constraints specified in the workflow definition, it will not be accepted. For example, if the InputDefinition specified that workflow must accept a string name 'key' and the value passed for key must adhere to a regex pattern. If Workflow was started with input where 'key' is not a string matching the regex pattern, an error will be generated and workflow will not start execution. During workflow definition design, the input passed into the workflow will be referred using the format 'workflow.input.<inputName>'. If the input is referred directly in a mapping it will be in the format '${workflow.input.<inputName>}' or inside a template mapping in the format '{{.global.workflow.input.<inputName>}}'.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"inst_id": {
			Description: "A workflow instance Id which is the unique identified for the workflow execution.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"internal": {
			Description: "Denotes that an Intersight service started this workflow as internal and hence will not be shown in Intersight User Interface. Typically these are internal system maintenance workflows which are triggered by Intersight services.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"last_action": {
			Description: "The last action that was issued on the workflow is saved in this field.\n* `None` - No action is set, this is the default value for action field.\n* `Create` - Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.\n* `Start` - Start a new execution of the workflow.\n* `Pause` - Pause the workflow, this can only be issued on workflows that are in running state. A workflow can be paused for a maximum of 180 days, after 180 days the workflow will be terminated by the system.\n* `Resume` - Resume the workflow which was previously paused through pause action on the workflow.\n* `Rerun` - Rerun the workflow that has previously reached a failed state. The workflow is run from the beginning using inputs from previous execution. Completed and currently running workflows cannot be rerun. Workflows do not have to be marked for retry to use this action.\n* `Retry` - This action has been deprecated. Please use RetryFailed, Rerun or RetryFromTask action. Retry the workflow that has previously reached a final state and has the retryable property set to true. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted from the first task.  The task name in retryFromTaskName must be one of the tasks that completed or failed in the previous run. It is not possible to retry a workflow from a task which wasn't run in the previous iteration.\n* `RetryFailed` - Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.\n* `RetryFromTask` - Retry the workflow that has previously reached a failed state and has the retryable property set to true. A running or waiting workflow cannot be retried. RetryFromTaskName must be passed along with this action, and the workflow will be started from that specific task. The task name in RetryFromTaskName must be one of the tasks that was executed in the previous attempt. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt.\n* `Cancel` - Cancel the workflow that is in running or waiting state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"message": {
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
					"message": {
						Description: "An i18n message that can be translated into multiple languages to support internationalization.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"severity": {
						Description: "The severity of the Task or Workflow message warning/error/info etc.\n* `Info` - The enum represents the log level to be used to convey info message.\n* `Warning` - The enum represents the log level to be used to convey warning message.\n* `Debug` - The enum represents the log level to be used to convey debug message.\n* `Error` - The enum represents the log level to be used to convey error message.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "A name of the workflow execution instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"organization": {
			Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"output": {
			Description: "All the generated outputs for the workflow. The schema for the outputs are defined in the OutputDefinition section of the WorkflowDefinition. The OutputDefinition will provide a list of output fields that could be generated after workflow execution is completed and the associated datatype of the outputs. For more information please refer to OutputDefinition property in WorkflowDefinition resource. The output for the workflow is generated as a collection of key-value pairs, where key is the name of the output and value is any valid JSON data which conforms to the datatype of output as specified in the OutputDefinition. During workflow definition design, if a workflow is included as a sub-workflow inside a parent workflow then the outputs generated by the sub-workflow can be used in the workflow design. For example, if workflow was included into parent workflow as 'SubWorkflowSample1', then that output can be referred as 'SubWorkflowSample1.output.<outputName>'. In the output is referred directly in a mapping it will be in the format '${SubWorkflowSample1.output.<outputName>}' or inside a template mapping will be in the format '{{SubWorkflowSample1.output.<outputName>}}'.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"parent_task_info": {
			Description: "A reference to a workflowTaskInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"pause_reason": {
			Description: "Denotes the reason workflow is in paused status.\n* `None` - Pause reason is none, which indicates there is no reason for the pause state.\n* `TaskWithWarning` - Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings.\n* `SystemMaintenance` - Pause reason indicates the workflow is in this state based on actions of system admin for maintenance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"permission": {
			Description: "A reference to a iamPermission resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"progress": {
			Description: "This field indicates percentage of workflow task completion based on the total number of tasks in the workflow. The total number of tasks in the workflow is calculated based on the longest path the workflow execution can take. So progress is calculated based on the percentage of tasks that completed out of the total number of tasks that could be executed. Progress is not a representation of the time taken to complete the workflow. A task is considered as completed if the task status is either \"NO_OP\" or \"COMPLETED\". If the task status is \"SKIP_TO_FAIL\", the workflow will be terminated and the progress of the workflow will be set to 100.",
			Type:        schema.TypeFloat,
			Optional:    true,
		},
		"properties": {
			Description: "Type to capture all the properties for the workflow info passed on from workflow definition.",
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
					"cancelable": {
						Description: "Holds the parameters and conditions for a workflow to be cancelable.",
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
								"cancelable_states": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Schema{
										Type: schema.TypeString}},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"enabled": {
									Description: "When true the workflow can be cancelled. The action can be further restricted by the mode and cancelableStates properties.",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"mode": {
									Description: "Mode controls how the workflow can be canceled.\n* `ApiOnly` - The workflow can only be canceled via API call.\n* `All` - The workflow can be canceled from API or from the user interface.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"retryable": {
						Description: "When true, this workflow can be retried within 2 weeks from the last failure.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"rollback_action": {
						Description: "Status of rollback for this workflow instance. The rollback action can be enabled, disabled or completed.\n* `Disabled` - Status of the rollback action when workflow is disabled for rollback.\n* `Enabled` - Status of the rollback action when workflow is enabled for rollback.\n* `Completed` - Status of the rollback action once workflow completes the rollback for all eligible tasks.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"rollback_on_cancel": {
						Description: "When set to true, the changes are automatically rolled back if the workflow execution is cancelled.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"rollback_on_failure": {
						Description: "When set to true, the changes are automatically rolled back if the workflow fails to execute.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"retry_from_task_name": {
			Description: "This field is required when RetryFromTask action is issued for a workflow that is in a 'final' state. The workflow will be retried from the specified task. This field must specify a task name which is the unique name of the task within the workflow. The task name must be one of the tasks that were completed or failed in the previous run. It is not possible to retry a workflow from a task that wasn't run in the previous execution attempt.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"src": {
			Description: "The source service that started the workflow execution and hence represents the owning service for this workflow.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"start_time": {
			Description: "The time when the workflow was started for execution.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status": {
			Description: "A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED). The \"status\" field has been deprecated and is now replaced with the \"workflowStatus\" field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"success_workflow_cleanup_duration": {
			Description: "The duration in hours after which the workflow info for successful workflow will be removed from database. The minimum is 1 hour, maximum is 365 days and default is 90 days.",
			Type:        schema.TypeInt,
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
		"task_info_update": {
			Description: "Used to update a TaskInfo instance in the WorkflowInfo, it is used as a way to update status and provide user inputs for a WaitTask.",
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
					"input": {
						Description: "Inputs for the specified TaskInfo. Inputs must only be provided for tasks which has included an input definition and the inputs must match the constraints specified in the input definition.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the task being updated and this name must match the task instance name included inside the workflow definition. This name is also captured in the RefName property of the TaskInfo object for the task.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "New status of the task being updated, only Failed and Completed statuses are supported, Completed is the default value in case no status is provided.\n* `Scheduled` - The enum represents the status when task is in scheduled state.\n* `InProgress` - The enum represents the status when task is in-progress state.\n* `NoOp` - The enum represents the status when task is a noop.\n* `Timeout` - The enum represents the status when task has timed out.\n* `Completed` - The enum represents the status when task has completed.\n* `Failed` - The enum represents the status when task has failed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"task_infos": {
			Description: "An array of relationships to workflowTaskInfo resources.",
			Type:        schema.TypeList,
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"trace_id": {
			Description: "The trace id to keep track of workflow execution.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"type": {
			Description: "A type of the workflow (serverconfig, ansible_monitoring).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"user_action_required": {
			Description: "Property will be set when a user action is required on the workflow. This can be because the workflow is waiting for a wait task to be updated, workflow is paused or workflow launched by a configuration object has failed and needs to be retried in order to complete successfully.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"user_id": {
			Description: "The user identifier which indicates the user that started this workflow.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"variable": {
			Description: "All the generated variables for the workflow. During workflow execution, the variables will be updated as per the variableParameters specified after each task execution.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
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
					"interested_mos": {
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
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"marked_for_deletion": {
						Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
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
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"wait_reason": {
			Description: "Denotes the reason workflow is in waiting status.\n* `None` - Wait reason is none, which indicates there is no reason for the waiting state.\n* `GatherTasks` - Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks.\n* `Duplicate` - Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow.\n* `RateLimit` - Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold.\n* `WaitTask` - Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update.\n* `PendingRetryFailed` - Wait reason when the workflow is pending a RetryFailed action.\n* `WaitingToStart` - Workflow is waiting to start on workflow engine.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"workflow_ctx": {
			Description: "The workflow context which contains initiator and target information.",
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
					"initiator_ctx": {
						Description: "Details about initiator of the workflow. Any Intersight object resource can be set as the initiator of the workflow. For workflows executed by an Intersight service, an applicable service object will be set as the initiator. For example, during server profile deployment workflow, the server profile object will be set as the initiator by the system. For user created workflows, this field is optional and for workflows executed from Intersight workflow execution page, the workflow definition object will be set as the Initiator.",
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
								"initiator_moid": {
									Description: "The moid of the Intersight managed object that initiated the workflow.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"initiator_name": {
									Description: "Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"initiator_type": {
									Description: "Type of Intersight managed object that initiated the workflow.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"target_ctx_list": {
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
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"target_moid": {
									Description: "Moid of the target Intersight managed object.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"target_name": {
									Description: "Name of the target instance.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"target_type": {
									Description: "Object type of the target Intersight managed object.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"workflow_subtype": {
						Description: "The subtype of the dynamic workflow. For example - Intersight services offer the following subtypes [Validate, Deploy, Import] for dynamic workflow of type serverconfig. This field is not applicable for user created workflows.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"workflow_type": {
						Description: "Intersight services set the type of dynamic workflow that need to be built and executed. This field is not applicable for user created workflows. WorkflowType set as ServerConfig states that a dynamic workflow is executing tasks related to server configuration.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"workflow_definition": {
			Description: "A reference to a workflowWorkflowDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"workflow_status": {
			Description: "The current state of the workflow execution instance. A draft workflow execution will be in NotStarted state and when \"Start\" action is issued then the workflow will move into Waiting state until the first task of the workflow is scheduled at which time it will move into InProgress state. When execution reaches a final state it move to either Completed, Failed or Terminated state. For more details look at the description for each state.\n* `NotStarted` - Initially all the workflow instances are at \"NotStarted\" state. A workflow can be drafted in this state by issuing Create action. When a workflow is in this state the inputs can be updated until the workflow is started.\n* `InProgress` - A workflow execution moves into \"InProgress\" state when the first task of the workflow is scheduled for execution and continues to remain in that state as long as there are tasks executing or yet to be scheduled for execution.\n* `Waiting` - Workflow can go to waiting state due to execution of wait task present in the workflow or the workflow has not started yet either due to duplicate workflow is running or due to workflow throttling. Once Workflow engine picks up the workflow for execution, it will move to in progress state.\n* `Completed` - A workflow execution moves into Completed state when the execution path of the workflow has reached the Success node in the workflow design and there are no more tasks to be executed. Completed is the final state for the workflow execution instance and no further actions are allowed on this workflow instance.\n* `Failed` - A workflow execution moves into a Failed state when the execution path of the workflow has reached the Failed node in the workflow design and there are no more tasks to be scheduled. A Failed node can be reached when the last executed task has failed or timed out and there are no further retries available for the task. Also as per the workflow design, the last executed task did not specify an OnFailure task to be executed and hence by default, the execution will reach the Failed node. Actions like \"Rerun\", \"RetryFailed\" and \"RetryFromTask\" can be issued on failed workflow instances. Please refer to the \"Action\" description for more details.\n* `Terminated` - A workflow execution moves to Terminated state when user issues a \"Cancel\" action or due to internal errors caused during workflow execution. e.g. - Task input transformation has failed. Terminated is a final state of the workflow, no further action are allowed on this workflow instance.\n* `Canceled` - A workflow execution moves to Canceled state when a user issues a \"Cancel\" action. Cancel is not a final state, the workflow engine will issue cancel to all the running tasks and then move the workflow to the \"Terminated\" state.\n* `Paused` - A workflow execution moves to Paused state when user issues a \"Pause\" action. When in paused state the current running task will complete its execution but no further tasks will be scheduled until the workflow is resumed. A paused workflow is resumed when the user issues a \"Resume\" action. Paused workflows can be canceled by user.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
	return schemaMap
}

func dataSourceWorkflowWorkflowInfo() *schema.Resource {
	var subSchema = getWorkflowWorkflowInfoSchema()
	var model = getWorkflowWorkflowInfoSchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceWorkflowWorkflowInfoRead,
		Schema:      model}
}

func dataSourceWorkflowWorkflowInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowWorkflowInfo{}
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

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

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("associated_object"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAssociatedObject(x)
		}
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("cleanup_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCleanupTime(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("email"); ok {
		x := (v.(string))
		o.SetEmail(x)
	}

	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetEndTime(x)
	}

	if v, ok := d.GetOkExists("failed_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.SetFailedWorkflowCleanupDuration(x)
	}

	if v, ok := d.GetOk("input"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			x2 := x1.(map[string]interface{})
			o.SetInput(x2)
		}
	}

	if v, ok := d.GetOk("inst_id"); ok {
		x := (v.(string))
		o.SetInstId(x)
	}

	if v, ok := d.GetOkExists("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
	}

	if v, ok := d.GetOk("last_action"); ok {
		x := (v.(string))
		o.SetLastAction(x)
	}

	if v, ok := d.GetOk("message"); ok {
		x := make([]models.WorkflowMessage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowMessage{}
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
			o.SetClassId("workflow.Message")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetMessage(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
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

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("output"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			x2 := x1.(map[string]interface{})
			o.SetOutput(x2)
		}
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("parent_task_info"); ok {
		p := make([]models.WorkflowTaskInfoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowTaskInfoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParentTaskInfo(x)
		}
	}

	if v, ok := d.GetOk("pause_reason"); ok {
		x := (v.(string))
		o.SetPauseReason(x)
	}

	if v, ok := d.GetOk("permission"); ok {
		p := make([]models.IamPermissionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsIamPermissionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPermission(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("progress"); ok {
		x := float32(v.(float64))
		o.SetProgress(x)
	}

	if v, ok := d.GetOk("properties"); ok {
		p := make([]models.WorkflowWorkflowInfoProperties, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.WorkflowWorkflowInfoProperties{}
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
			o.SetClassId("workflow.WorkflowInfoProperties")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetProperties(x)
		}
	}

	if v, ok := d.GetOk("retry_from_task_name"); ok {
		x := (v.(string))
		o.SetRetryFromTaskName(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.SetSrc(x)
	}

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetStartTime(x)
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	if v, ok := d.GetOkExists("success_workflow_cleanup_duration"); ok {
		x := int64(v.(int))
		o.SetSuccessWorkflowCleanupDuration(x)
	}

	if v, ok := d.GetOk("tags"); ok {
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
		o.SetTags(x)
	}

	if v, ok := d.GetOk("task_info_update"); ok {
		p := make([]models.WorkflowTaskInfoUpdate, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.WorkflowTaskInfoUpdate{}
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
			o.SetClassId("workflow.TaskInfoUpdate")
			if v, ok := l["input"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInput(x2)
					}
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
			if v, ok := l["status"]; ok {
				{
					x := (v.(string))
					o.SetStatus(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetTaskInfoUpdate(x)
		}
	}

	if v, ok := d.GetOk("task_infos"); ok {
		x := make([]models.WorkflowTaskInfoRelationship, 0)
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
			x = append(x, models.MoMoRefAsWorkflowTaskInfoRelationship(o))
		}
		o.SetTaskInfos(x)
	}

	if v, ok := d.GetOk("trace_id"); ok {
		x := (v.(string))
		o.SetTraceId(x)
	}

	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	if v, ok := d.GetOkExists("user_action_required"); ok {
		x := (v.(bool))
		o.SetUserActionRequired(x)
	}

	if v, ok := d.GetOk("user_id"); ok {
		x := (v.(string))
		o.SetUserId(x)
	}

	if v, ok := d.GetOk("variable"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			x2 := x1.(map[string]interface{})
			o.SetVariable(x2)
		}
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
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
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
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
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if v, ok := d.GetOk("wait_reason"); ok {
		x := (v.(string))
		o.SetWaitReason(x)
	}

	if v, ok := d.GetOk("workflow_ctx"); ok {
		p := make([]models.WorkflowWorkflowCtx, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.WorkflowWorkflowCtx{}
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
			o.SetClassId("workflow.WorkflowCtx")
			if v, ok := l["initiator_ctx"]; ok {
				{
					p := make([]models.WorkflowInitiatorContext, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewWorkflowInitiatorContextWithDefaults()
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
						o.SetClassId("workflow.InitiatorContext")
						if v, ok := l["initiator_moid"]; ok {
							{
								x := (v.(string))
								o.SetInitiatorMoid(x)
							}
						}
						if v, ok := l["initiator_name"]; ok {
							{
								x := (v.(string))
								o.SetInitiatorName(x)
							}
						}
						if v, ok := l["initiator_type"]; ok {
							{
								x := (v.(string))
								o.SetInitiatorType(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetInitiatorCtx(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["target_ctx_list"]; ok {
				{
					x := make([]models.WorkflowTargetContext, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowTargetContextWithDefaults()
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
						o.SetClassId("workflow.TargetContext")
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetTargetCtxList(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetWorkflowCtx(x)
		}
	}

	if v, ok := d.GetOk("workflow_definition"); ok {
		p := make([]models.WorkflowWorkflowDefinitionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowWorkflowDefinitionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetWorkflowDefinition(x)
		}
	}

	if v, ok := d.GetOk("workflow_status"); ok {
		x := (v.(string))
		o.SetWorkflowStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowWorkflowInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowInfoList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of WorkflowWorkflowInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of WorkflowWorkflowInfo: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for WorkflowWorkflowInfo data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var workflowWorkflowInfoResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching WorkflowWorkflowInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching WorkflowWorkflowInfo: %s", responseErr.Error())
		}
		results := resMo.WorkflowWorkflowInfoList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["account_moid"] = (s.GetAccountMoid())
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["associated_object"] = flattenMapMoBaseMoRelationship(s.GetAssociatedObject(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cleanup_time"] = (s.GetCleanupTime()).String()

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["email"] = (s.GetEmail())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["failed_workflow_cleanup_duration"] = (s.GetFailedWorkflowCleanupDuration())
				temp["input"] = flattenAdditionalProperties(s.GetInput())
				temp["inst_id"] = (s.GetInstId())
				temp["internal"] = (s.GetInternal())
				temp["last_action"] = (s.GetLastAction())

				temp["message"] = flattenListWorkflowMessage(s.GetMessage(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["output"] = flattenAdditionalProperties(s.GetOutput())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["parent_task_info"] = flattenMapWorkflowTaskInfoRelationship(s.GetParentTaskInfo(), d)
				temp["pause_reason"] = (s.GetPauseReason())

				temp["permission"] = flattenMapIamPermissionRelationship(s.GetPermission(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["progress"] = (s.GetProgress())

				temp["properties"] = flattenMapWorkflowWorkflowInfoProperties(s.GetProperties(), d)
				temp["retry_from_task_name"] = (s.GetRetryFromTaskName())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["src"] = (s.GetSrc())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())
				temp["success_workflow_cleanup_duration"] = (s.GetSuccessWorkflowCleanupDuration())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["task_infos"] = flattenListWorkflowTaskInfoRelationship(s.GetTaskInfos(), d)
				temp["trace_id"] = (s.GetTraceId())
				temp["type"] = (s.GetType())
				temp["user_action_required"] = (s.GetUserActionRequired())
				temp["user_id"] = (s.GetUserId())
				temp["variable"] = flattenAdditionalProperties(s.GetVariable())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["wait_reason"] = (s.GetWaitReason())

				temp["workflow_ctx"] = flattenMapWorkflowWorkflowCtx(s.GetWorkflowCtx(), d)

				temp["workflow_definition"] = flattenMapWorkflowWorkflowDefinitionRelationship(s.GetWorkflowDefinition(), d)
				temp["workflow_status"] = (s.GetWorkflowStatus())
				workflowWorkflowInfoResults = append(workflowWorkflowInfoResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(workflowWorkflowInfoResults))
	if err := d.Set("results", workflowWorkflowInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowWorkflowInfoResults[0]["moid"].(string))
	return de
}
