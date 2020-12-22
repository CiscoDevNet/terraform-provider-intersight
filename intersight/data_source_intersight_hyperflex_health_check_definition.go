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

func dataSourceHyperflexHealthCheckDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthCheckDefinitionRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"category": {
				Description: "Category that the health check belongs to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"common_causes": {
				Description: "Static information detailing the common causes for the health check failure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"configuration": {
				Description: "Execution configuration fo the health check script.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_health_check_script_info": {
				Description: "Default version Script info.",
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
						"aggregate_script_name": {
							Description: "Health check aggregate script that runs in the HyperFlex Leader Node. |\nIt aggregates the output of all HyperFlex nodes and provides the health check result.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"hyperflex_version": {
							Description: "HyperFlex Data Platform version running on the target device.",
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
						"script_execute_location": {
							Description: "Location of the health check script's execution on the HyperFlex device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"script_name": {
							Description: "Name of the health check script to be executed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Version of the health check script associated with the health check definition.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"description": {
				Description: "Description of the health check definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"health_check_script_infos": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"aggregate_script_name": {
							Description: "Health check aggregate script that runs in the HyperFlex Leader Node. |\nIt aggregates the output of all HyperFlex nodes and provides the health check result.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"hyperflex_version": {
							Description: "HyperFlex Data Platform version running on the target device.",
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
						"script_execute_location": {
							Description: "Location of the health check script's execution on the HyperFlex device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"script_name": {
							Description: "Name of the health check script to be executed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Version of the health check script associated with the health check definition.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"health_impact": {
				Description: "Static information detailing the health impact of the health check failure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal_name": {
				Description: "Internal name of the health check definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the health check definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"reference": {
				Description: "Static information containing additional reference for the health check.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resolution": {
				Description: "Static information detailing the possible remediation actions that can be taken to remedy the health check failure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"script_execution_mode": {
				Description: "Execution mode of the health script on the HyperFlex cluster.\n* `ON_DEMAND` - Execute the health check on-demand.\n* `SCHEDULED` - Execute the health check on a scheduled interval.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"script_execution_on_compute_nodes": {
				Description: "Indicates if the script needs to be executed on HyperFlex compute nodes. |\nTypically, scripts are only executed on the storage Nodes.",
				Type:        schema.TypeBool,
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
			"target_execution_type": {
				Description: "Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster.\n* `EXECUTE_ON_LEADER_NODE` - Execute the health check script only on the HyperFlex cluster's leader node.\n* `EXECUTE_ON_ALL_NODES` - Execute health check on all nodes and aggregate the results.\n* `EXECUTE_ON_ALL_NODES_AND_AGGREGATE` - Execute the health check on all Nodes and perform custom aggregation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"timeout": {
				Description: "Health check script execution timeout.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"unsupported_hyper_flex_versions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
		},
	}
}

func dataSourceHyperflexHealthCheckDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckDefinition{}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("common_causes"); ok {
		x := (v.(string))
		o.SetCommonCauses(x)
	}
	if v, ok := d.GetOk("configuration"); ok {
		x := (v.(string))
		o.SetConfiguration(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("health_impact"); ok {
		x := (v.(string))
		o.SetHealthImpact(x)
	}
	if v, ok := d.GetOk("internal_name"); ok {
		x := (v.(string))
		o.SetInternalName(x)
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
	if v, ok := d.GetOk("reference"); ok {
		x := (v.(string))
		o.SetReference(x)
	}
	if v, ok := d.GetOk("resolution"); ok {
		x := (v.(string))
		o.SetResolution(x)
	}
	if v, ok := d.GetOk("script_execution_mode"); ok {
		x := (v.(string))
		o.SetScriptExecutionMode(x)
	}
	if v, ok := d.GetOk("script_execution_on_compute_nodes"); ok {
		x := (v.(bool))
		o.SetScriptExecutionOnComputeNodes(x)
	}
	if v, ok := d.GetOk("target_execution_type"); ok {
		x := (v.(string))
		o.SetTargetExecutionType(x)
	}
	if v, ok := d.GetOk("timeout"); ok {
		x := int64(v.(int))
		o.SetTimeout(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHealthCheckDefinition object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckDefinitionList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexHealthCheckDefinition list: %s", err.Error())
	}
	var s = &models.HyperflexHealthCheckDefinitionList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexHealthCheckDefinition list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for HyperflexHealthCheckDefinition did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexHealthCheckDefinition{}
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
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("common_causes", (s.GetCommonCauses())); err != nil {
				return diag.Errorf("error occurred while setting property CommonCauses: %s", err.Error())
			}
			if err := d.Set("configuration", (s.GetConfiguration())); err != nil {
				return diag.Errorf("error occurred while setting property Configuration: %s", err.Error())
			}

			if err := d.Set("default_health_check_script_info", flattenMapHyperflexHealthCheckScriptInfo(s.GetDefaultHealthCheckScriptInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DefaultHealthCheckScriptInfo: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}

			if err := d.Set("health_check_script_infos", flattenListHyperflexHealthCheckScriptInfo(s.GetHealthCheckScriptInfos(), d)); err != nil {
				return diag.Errorf("error occurred while setting property HealthCheckScriptInfos: %s", err.Error())
			}
			if err := d.Set("health_impact", (s.GetHealthImpact())); err != nil {
				return diag.Errorf("error occurred while setting property HealthImpact: %s", err.Error())
			}
			if err := d.Set("internal_name", (s.GetInternalName())); err != nil {
				return diag.Errorf("error occurred while setting property InternalName: %s", err.Error())
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
			if err := d.Set("reference", (s.GetReference())); err != nil {
				return diag.Errorf("error occurred while setting property Reference: %s", err.Error())
			}
			if err := d.Set("resolution", (s.GetResolution())); err != nil {
				return diag.Errorf("error occurred while setting property Resolution: %s", err.Error())
			}
			if err := d.Set("script_execution_mode", (s.GetScriptExecutionMode())); err != nil {
				return diag.Errorf("error occurred while setting property ScriptExecutionMode: %s", err.Error())
			}
			if err := d.Set("script_execution_on_compute_nodes", (s.GetScriptExecutionOnComputeNodes())); err != nil {
				return diag.Errorf("error occurred while setting property ScriptExecutionOnComputeNodes: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("target_execution_type", (s.GetTargetExecutionType())); err != nil {
				return diag.Errorf("error occurred while setting property TargetExecutionType: %s", err.Error())
			}
			if err := d.Set("timeout", (s.GetTimeout())); err != nil {
				return diag.Errorf("error occurred while setting property Timeout: %s", err.Error())
			}
			if err := d.Set("unsupported_hyper_flex_versions", (s.GetUnsupportedHyperFlexVersions())); err != nil {
				return diag.Errorf("error occurred while setting property UnsupportedHyperFlexVersions: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
