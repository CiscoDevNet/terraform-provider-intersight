package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHyperflexHealthCheckDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHyperflexHealthCheckDefinitionCreate,
		ReadContext:   resourceHyperflexHealthCheckDefinitionRead,
		UpdateContext: resourceHyperflexHealthCheckDefinitionUpdate,
		DeleteContext: resourceHyperflexHealthCheckDefinitionDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"hyperflex_version": {
							Description: "HyperFlex Data Platform version running on the target device.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"hyperflex_version": {
							Description: "HyperFlex Data Platform version running on the target device.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the health check definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
				Default:     "ON_DEMAND",
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
				Default:     "EXECUTE_ON_LEADER_NODE",
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

func resourceHyperflexHealthCheckDefinitionCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHyperflexHealthCheckDefinitionWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}

	o.SetClassId("hyperflex.HealthCheckDefinition")

	if v, ok := d.GetOk("common_causes"); ok {
		x := (v.(string))
		o.SetCommonCauses(x)
	}

	if v, ok := d.GetOk("configuration"); ok {
		x := (v.(string))
		o.SetConfiguration(x)
	}

	if v, ok := d.GetOk("default_health_check_script_info"); ok {
		p := make([]models.HyperflexHealthCheckScriptInfo, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewHyperflexHealthCheckScriptInfoWithDefaults()
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
			if v, ok := l["aggregate_script_name"]; ok {
				{
					x := (v.(string))
					o.SetAggregateScriptName(x)
				}
			}
			o.SetClassId("hyperflex.HealthCheckScriptInfo")
			if v, ok := l["hyperflex_version"]; ok {
				{
					x := (v.(string))
					o.SetHyperflexVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["script_execute_location"]; ok {
				{
					x := (v.(string))
					o.SetScriptExecuteLocation(x)
				}
			}
			if v, ok := l["script_name"]; ok {
				{
					x := (v.(string))
					o.SetScriptName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetDefaultHealthCheckScriptInfo(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("health_check_script_infos"); ok {
		x := make([]models.HyperflexHealthCheckScriptInfo, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewHyperflexHealthCheckScriptInfoWithDefaults()
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
			if v, ok := l["aggregate_script_name"]; ok {
				{
					x := (v.(string))
					o.SetAggregateScriptName(x)
				}
			}
			o.SetClassId("hyperflex.HealthCheckScriptInfo")
			if v, ok := l["hyperflex_version"]; ok {
				{
					x := (v.(string))
					o.SetHyperflexVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["script_execute_location"]; ok {
				{
					x := (v.(string))
					o.SetScriptExecuteLocation(x)
				}
			}
			if v, ok := l["script_name"]; ok {
				{
					x := (v.(string))
					o.SetScriptName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetHealthCheckScriptInfos(x)
		}
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

	o.SetObjectType("hyperflex.HealthCheckDefinition")

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

	if v, ok := d.GetOkExists("script_execution_on_compute_nodes"); ok {
		x := v.(bool)
		o.SetScriptExecutionOnComputeNodes(x)
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

	if v, ok := d.GetOk("target_execution_type"); ok {
		x := (v.(string))
		o.SetTargetExecutionType(x)
	}

	if v, ok := d.GetOk("timeout"); ok {
		x := int64(v.(int))
		o.SetTimeout(x)
	}

	if v, ok := d.GetOk("unsupported_hyper_flex_versions"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetUnsupportedHyperFlexVersions(x)
		}
	}

	r := conn.ApiClient.HyperflexApi.CreateHyperflexHealthCheckDefinition(conn.ctx).HyperflexHealthCheckDefinition(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceHyperflexHealthCheckDefinitionRead(c, d, meta)
}

func resourceHyperflexHealthCheckDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckDefinitionByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexHealthCheckDefinition object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("category", (s.GetCategory())); err != nil {
		return diag.Errorf("error occurred while setting property Category in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("common_causes", (s.GetCommonCauses())); err != nil {
		return diag.Errorf("error occurred while setting property CommonCauses in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("configuration", (s.GetConfiguration())); err != nil {
		return diag.Errorf("error occurred while setting property Configuration in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("default_health_check_script_info", flattenMapHyperflexHealthCheckScriptInfo(s.GetDefaultHealthCheckScriptInfo(), d)); err != nil {
		return diag.Errorf("error occurred while setting property DefaultHealthCheckScriptInfo in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("health_check_script_infos", flattenListHyperflexHealthCheckScriptInfo(s.GetHealthCheckScriptInfos(), d)); err != nil {
		return diag.Errorf("error occurred while setting property HealthCheckScriptInfos in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("health_impact", (s.GetHealthImpact())); err != nil {
		return diag.Errorf("error occurred while setting property HealthImpact in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("internal_name", (s.GetInternalName())); err != nil {
		return diag.Errorf("error occurred while setting property InternalName in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("reference", (s.GetReference())); err != nil {
		return diag.Errorf("error occurred while setting property Reference in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("resolution", (s.GetResolution())); err != nil {
		return diag.Errorf("error occurred while setting property Resolution in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("script_execution_mode", (s.GetScriptExecutionMode())); err != nil {
		return diag.Errorf("error occurred while setting property ScriptExecutionMode in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("script_execution_on_compute_nodes", (s.GetScriptExecutionOnComputeNodes())); err != nil {
		return diag.Errorf("error occurred while setting property ScriptExecutionOnComputeNodes in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("target_execution_type", (s.GetTargetExecutionType())); err != nil {
		return diag.Errorf("error occurred while setting property TargetExecutionType in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("timeout", (s.GetTimeout())); err != nil {
		return diag.Errorf("error occurred while setting property Timeout in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	if err := d.Set("unsupported_hyper_flex_versions", (s.GetUnsupportedHyperFlexVersions())); err != nil {
		return diag.Errorf("error occurred while setting property UnsupportedHyperFlexVersions in HyperflexHealthCheckDefinition object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHyperflexHealthCheckDefinitionUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.HyperflexHealthCheckDefinition{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("category") {
		v := d.Get("category")
		x := (v.(string))
		o.SetCategory(x)
	}

	o.SetClassId("hyperflex.HealthCheckDefinition")

	if d.HasChange("common_causes") {
		v := d.Get("common_causes")
		x := (v.(string))
		o.SetCommonCauses(x)
	}

	if d.HasChange("configuration") {
		v := d.Get("configuration")
		x := (v.(string))
		o.SetConfiguration(x)
	}

	if d.HasChange("default_health_check_script_info") {
		v := d.Get("default_health_check_script_info")
		p := make([]models.HyperflexHealthCheckScriptInfo, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexHealthCheckScriptInfo{}
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
			if v, ok := l["aggregate_script_name"]; ok {
				{
					x := (v.(string))
					o.SetAggregateScriptName(x)
				}
			}
			o.SetClassId("hyperflex.HealthCheckScriptInfo")
			if v, ok := l["hyperflex_version"]; ok {
				{
					x := (v.(string))
					o.SetHyperflexVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["script_execute_location"]; ok {
				{
					x := (v.(string))
					o.SetScriptExecuteLocation(x)
				}
			}
			if v, ok := l["script_name"]; ok {
				{
					x := (v.(string))
					o.SetScriptName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetDefaultHealthCheckScriptInfo(x)
		}
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("health_check_script_infos") {
		v := d.Get("health_check_script_infos")
		x := make([]models.HyperflexHealthCheckScriptInfo, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.HyperflexHealthCheckScriptInfo{}
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
			if v, ok := l["aggregate_script_name"]; ok {
				{
					x := (v.(string))
					o.SetAggregateScriptName(x)
				}
			}
			o.SetClassId("hyperflex.HealthCheckScriptInfo")
			if v, ok := l["hyperflex_version"]; ok {
				{
					x := (v.(string))
					o.SetHyperflexVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["script_execute_location"]; ok {
				{
					x := (v.(string))
					o.SetScriptExecuteLocation(x)
				}
			}
			if v, ok := l["script_name"]; ok {
				{
					x := (v.(string))
					o.SetScriptName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetHealthCheckScriptInfos(x)
		}
	}

	if d.HasChange("health_impact") {
		v := d.Get("health_impact")
		x := (v.(string))
		o.SetHealthImpact(x)
	}

	if d.HasChange("internal_name") {
		v := d.Get("internal_name")
		x := (v.(string))
		o.SetInternalName(x)
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

	o.SetObjectType("hyperflex.HealthCheckDefinition")

	if d.HasChange("reference") {
		v := d.Get("reference")
		x := (v.(string))
		o.SetReference(x)
	}

	if d.HasChange("resolution") {
		v := d.Get("resolution")
		x := (v.(string))
		o.SetResolution(x)
	}

	if d.HasChange("script_execution_mode") {
		v := d.Get("script_execution_mode")
		x := (v.(string))
		o.SetScriptExecutionMode(x)
	}

	if d.HasChange("script_execution_on_compute_nodes") {
		v := d.Get("script_execution_on_compute_nodes")
		x := (v.(bool))
		o.SetScriptExecutionOnComputeNodes(x)
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

	if d.HasChange("target_execution_type") {
		v := d.Get("target_execution_type")
		x := (v.(string))
		o.SetTargetExecutionType(x)
	}

	if d.HasChange("timeout") {
		v := d.Get("timeout")
		x := int64(v.(int))
		o.SetTimeout(x)
	}

	if d.HasChange("unsupported_hyper_flex_versions") {
		v := d.Get("unsupported_hyper_flex_versions")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetUnsupportedHyperFlexVersions(x)
		}
	}

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexHealthCheckDefinition(conn.ctx, d.Id()).HyperflexHealthCheckDefinition(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceHyperflexHealthCheckDefinitionRead(c, d, meta)
}

func resourceHyperflexHealthCheckDefinitionDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HyperflexApi.DeleteHyperflexHealthCheckDefinition(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting HyperflexHealthCheckDefinition object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
