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

func dataSourceWorkflowPendingDynamicWorkflowInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowPendingDynamicWorkflowInfoRead,
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
			"input": {
				Description: "The inputs of the workflow.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "A name for the pending dynamic workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pending_services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"src": {
				Description: "The src is workflow owner service.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The current status of the PendingDynamicWorkflowInfo.\n* `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution.\n* `Waiting` - Dynamic workflow is in waiting state and not yet started execution.",
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
			"wait_on_duplicate": {
				Description: "When set to true workflow engine will wait for a duplicate to finish before starting a new one.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"workflow_action_task_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Description: "The action of the Dynamic Workflow.",
							Type:        schema.TypeString,
							Optional:    true,
						},
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"tasks": {
							Description: "The task list that has precedence which dictates how the workflow should be constructed.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
					},
				},
				Computed: true,
			},
			"workflow_ctx": {
				Description: "The workflow's workflow context which contains initiator and target information.",
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
						"initiator_ctx": {
							Description: "Details about initiator of the workflow.",
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
									"initiator_moid": {
										Description: "The moid of the Intersigt managed object that initiated the workflow.",
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
										Computed:    true,
									},
								},
							},
							Computed: true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
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
							Computed: true,
						},
						"workflow_meta_name": {
							Description: "The name of workflowMeta of the workflow running.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"workflow_subtype": {
							Description: "The subtype of the workflow.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"workflow_type": {
							Description: "Type of the workflow being started. This can be any string for client services to distinguish workflow by type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"workflow_info": {
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
			"workflow_key": {
				Description: "This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"workflow_meta": {
				Description: "The metadata of the workflow.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, Optional: true,
			},
		},
	}
}

func dataSourceWorkflowPendingDynamicWorkflowInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowPendingDynamicWorkflowInfo{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.SetSrc(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("wait_on_duplicate"); ok {
		x := (v.(bool))
		o.SetWaitOnDuplicate(x)
	}
	if v, ok := d.GetOk("workflow_key"); ok {
		x := (v.(string))
		o.SetWorkflowKey(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowPendingDynamicWorkflowInfo object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowPendingDynamicWorkflowInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching WorkflowPendingDynamicWorkflowInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for WorkflowPendingDynamicWorkflowInfo list: %s", err.Error())
	}
	var s = &models.WorkflowPendingDynamicWorkflowInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to WorkflowPendingDynamicWorkflowInfo list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for WorkflowPendingDynamicWorkflowInfo did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.WorkflowPendingDynamicWorkflowInfo{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
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
			if err := d.Set("pending_services", (s.GetPendingServices())); err != nil {
				return diag.Errorf("error occurred while setting property PendingServices: %s", err.Error())
			}
			if err := d.Set("src", (s.GetSrc())); err != nil {
				return diag.Errorf("error occurred while setting property Src: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("wait_on_duplicate", (s.GetWaitOnDuplicate())); err != nil {
				return diag.Errorf("error occurred while setting property WaitOnDuplicate: %s", err.Error())
			}

			if err := d.Set("workflow_action_task_lists", flattenListWorkflowDynamicWorkflowActionTaskList(s.GetWorkflowActionTaskLists(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowActionTaskLists: %s", err.Error())
			}

			if err := d.Set("workflow_ctx", flattenMapWorkflowWorkflowCtx(s.GetWorkflowCtx(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowCtx: %s", err.Error())
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowInfo: %s", err.Error())
			}
			if err := d.Set("workflow_key", (s.GetWorkflowKey())); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowKey: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
