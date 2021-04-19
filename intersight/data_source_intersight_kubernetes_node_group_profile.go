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

func dataSourceKubernetesNodeGroupProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesNodeGroupProfileRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"action": {
				Description: "User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"currentsize": {
				Description: "Current number of nodes in this node group at any given point in time.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"desiredsize": {
				Description: "Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"maxsize": {
				Description: "Maximum number of nodes desired in this node group.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"minsize": {
				Description: "Minimum number of nodes desired in this node group.",
				Type:        schema.TypeInt,
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
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"node_type": {
				Description: "The node type ControlPlane, Worker or ControlPlaneWorker.\n* `Worker` - Node will be marked as a worker node.\n* `ControlPlane` - Node will be marked as a control plane node.\n* `ControlPlaneWorker` - Node will be both a controle plane and a worker.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesNodeGroupProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesNodeGroupProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesNodeGroupProfile{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("currentsize"); ok {
		x := int64(v.(int))
		o.SetCurrentsize(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("desiredsize"); ok {
		x := int64(v.(int))
		o.SetDesiredsize(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("maxsize"); ok {
		x := int64(v.(int))
		o.SetMaxsize(x)
	}
	if v, ok := d.GetOk("minsize"); ok {
		x := int64(v.(int))
		o.SetMinsize(x)
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
	if v, ok := d.GetOk("node_type"); ok {
		x := (v.(string))
		o.SetNodeType(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesNodeGroupProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesNodeGroupProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of KubernetesNodeGroupProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of KubernetesNodeGroupProfile: %s", responseErr.Error())
	}
	count := countResponse.KubernetesNodeGroupProfileList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for KubernetesNodeGroupProfile data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var kubernetesNodeGroupProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesNodeGroupProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching KubernetesNodeGroupProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching KubernetesNodeGroupProfile: %s", responseErr.Error())
		}
		results := resMo.KubernetesNodeGroupProfileList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profile"] = flattenMapKubernetesClusterProfileRelationship(s.GetClusterProfile(), d)

				temp["config_context"] = flattenMapPolicyConfigContext(s.GetConfigContext(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["currentsize"] = (s.GetCurrentsize())
				temp["description"] = (s.GetDescription())
				temp["desiredsize"] = (s.GetDesiredsize())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["infra_provider"] = flattenMapKubernetesBaseInfrastructureProviderRelationship(s.GetInfraProvider(), d)

				temp["ip_pools"] = flattenListIppoolPoolRelationship(s.GetIpPools(), d)

				temp["kubernetes_version"] = flattenMapKubernetesVersionPolicyRelationship(s.GetKubernetesVersion(), d)

				temp["labels"] = flattenListKubernetesNodeGroupLabel(s.GetLabels(), d)
				temp["maxsize"] = (s.GetMaxsize())
				temp["minsize"] = (s.GetMinsize())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["node_type"] = (s.GetNodeType())

				temp["nodes"] = flattenListKubernetesNodeProfileRelationship(s.GetNodes(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["policy_bucket"] = flattenListPolicyAbstractPolicyRelationship(s.GetPolicyBucket(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["taints"] = flattenListKubernetesNodeGroupTaint(s.GetTaints(), d)
				temp["type"] = (s.GetType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				kubernetesNodeGroupProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesNodeGroupProfileResults))
	if err := d.Set("results", kubernetesNodeGroupProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesNodeGroupProfileResults[0]["moid"].(string))
	return de
}
