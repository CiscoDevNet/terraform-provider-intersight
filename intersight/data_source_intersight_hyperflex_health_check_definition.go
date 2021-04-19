package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexHealthCheckDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthCheckDefinitionRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the health check definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"mod_time": {
				Description: "The time when this managed object was last modified.",
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
			},
			"script_execution_on_compute_nodes": {
				Description: "Indicates if the script needs to be executed on HyperFlex compute nodes. |\nTypically, scripts are only executed on the storage Nodes.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexHealthCheckDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexHealthCheckDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckDefinition{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
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
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("health_impact"); ok {
		x := (v.(string))
		o.SetHealthImpact(x)
	}
	if v, ok := d.GetOk("internal_name"); ok {
		x := (v.(string))
		o.SetInternalName(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexHealthCheckDefinition: %s", responseErr.Error())
	}
	count := countResponse.HyperflexHealthCheckDefinitionList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexHealthCheckDefinition data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexHealthCheckDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexHealthCheckDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexHealthCheckDefinition: %s", responseErr.Error())
		}
		results := resMo.HyperflexHealthCheckDefinitionList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["category"] = (s.GetCategory())
				temp["class_id"] = (s.GetClassId())
				temp["common_causes"] = (s.GetCommonCauses())
				temp["configuration"] = (s.GetConfiguration())

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["default_health_check_script_info"] = flattenMapHyperflexHealthCheckScriptInfo(s.GetDefaultHealthCheckScriptInfo(), d)
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["health_check_script_infos"] = flattenListHyperflexHealthCheckScriptInfo(s.GetHealthCheckScriptInfos(), d)
				temp["health_impact"] = (s.GetHealthImpact())
				temp["internal_name"] = (s.GetInternalName())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["reference"] = (s.GetReference())
				temp["resolution"] = (s.GetResolution())
				temp["script_execution_mode"] = (s.GetScriptExecutionMode())
				temp["script_execution_on_compute_nodes"] = (s.GetScriptExecutionOnComputeNodes())
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_execution_type"] = (s.GetTargetExecutionType())
				temp["timeout"] = (s.GetTimeout())
				temp["unsupported_hyper_flex_versions"] = (s.GetUnsupportedHyperFlexVersions())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				hyperflexHealthCheckDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexHealthCheckDefinitionResults))
	if err := d.Set("results", hyperflexHealthCheckDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexHealthCheckDefinitionResults[0]["moid"].(string))
	return de
}
