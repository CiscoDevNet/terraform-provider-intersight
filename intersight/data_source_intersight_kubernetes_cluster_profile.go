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

func dataSourceKubernetesClusterProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesClusterProfileRead,
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
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"managed_mode": {
				Description: "Management mode for the cluster. In some cases Intersight kubernetes service is not required\nto provision and manage the management entities and endpoints (for e.g. EKS). In most other cases\nit will be required to provision and manage these entities and endpoints.\n* `Provided` - Cluster management entities and endpoints are provided by the infrastructure platform.\n* `Managed` - Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service.",
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
				Description: "Name of the profile instance or profile template.",
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
			"status": {
				Description: "Status of the Kubernetes cluster and its nodes.\n* `Undeployed` - The cluster is undeployed.\n* `Configuring` - The cluster is being configured.\n* `Deploying` - The cluster is being deployed.\n* `Undeploying` - The cluster is being undeployed.\n* `DeployFailedTerminal` - The cluster deployment failed terminally and can not be recovered.\n* `DeployFailed` - The cluster deployment failed.\n* `Upgrading` - The cluster is being upgraded.\n* `Deleting` - The cluster is being deleted.\n* `DeleteFailed` - The cluster delete failed.\n* `Ready` - The cluster is ready for use.\n* `Active` - The cluster is being active.\n* `Shutdown` - All the nodes in the cluster are powered off.\n* `Terminated` - The cluster is terminated.\n* `Deployed` - The cluster is deployed. The cluster may not yet be ready for use.\n* `UndeployFailed` - The cluster undeploy action failed.\n* `NotReady` - The cluster is created and some nodes are not ready.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted values are instance or template.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesClusterProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesClusterProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesClusterProfile{}
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
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("managed_mode"); ok {
		x := (v.(string))
		o.SetManagedMode(x)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesClusterProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesClusterProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of KubernetesClusterProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of KubernetesClusterProfile: %s", responseErr.Error())
	}
	count := countResponse.KubernetesClusterProfileList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for KubernetesClusterProfile data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var kubernetesClusterProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesClusterProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching KubernetesClusterProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching KubernetesClusterProfile: %s", responseErr.Error())
		}
		results := resMo.KubernetesClusterProfileList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())

				temp["aci_cni_profile"] = flattenMapKubernetesAciCniProfileRelationship(s.GetAciCniProfile(), d)
				temp["action"] = (s.GetAction())

				temp["action_info"] = flattenMapKubernetesActionInfo(s.GetActionInfo(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["associated_cluster"] = flattenMapKubernetesClusterRelationship(s.GetAssociatedCluster(), d)

				temp["cert_config"] = flattenMapKubernetesClusterCertificateConfiguration(s.GetCertConfig(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_ip_pools"] = flattenListIppoolPoolRelationship(s.GetClusterIpPools(), d)

				temp["config_context"] = flattenMapPolicyConfigContext(s.GetConfigContext(), d)

				temp["container_runtime_config"] = flattenMapKubernetesContainerRuntimePolicyRelationship(s.GetContainerRuntimeConfig(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["essential_addons"] = flattenListKubernetesEssentialAddon(s.GetEssentialAddons(), d)

				temp["kube_config"] = flattenMapKubernetesConfiguration(s.GetKubeConfig(), d)

				temp["loadbalancer_ip_leases"] = flattenListIppoolIpLeaseRelationship(s.GetLoadbalancerIpLeases(), d)
				temp["managed_mode"] = (s.GetManagedMode())

				temp["management_config"] = flattenMapKubernetesClusterManagementConfig(s.GetManagementConfig(), d)

				temp["master_vip_lease"] = flattenMapIppoolIpLeaseRelationship(s.GetMasterVipLease(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["net_config"] = flattenMapKubernetesNetworkPolicyRelationship(s.GetNetConfig(), d)

				temp["node_groups"] = flattenListKubernetesNodeGroupProfileRelationship(s.GetNodeGroups(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["policy_bucket"] = flattenListPolicyAbstractPolicyRelationship(s.GetPolicyBucket(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)
				temp["status"] = (s.GetStatus())

				temp["sys_config"] = flattenMapKubernetesSysConfigPolicyRelationship(s.GetSysConfig(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["trusted_registries"] = flattenMapKubernetesTrustedRegistriesPolicyRelationship(s.GetTrustedRegistries(), d)
				temp["type"] = (s.GetType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["workflow_info"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)
				kubernetesClusterProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesClusterProfileResults))
	if err := d.Set("results", kubernetesClusterProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesClusterProfileResults[0]["moid"].(string))
	return de
}
