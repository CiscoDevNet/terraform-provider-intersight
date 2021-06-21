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

func dataSourceHyperflexClusterProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterProfileRead,
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
			"data_ip_address": {
				Description: "The storage data IP address for the HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"host_name_prefix": {
				Description: "The node name prefix that is used to automatically generate the default hostname for each server.\nA dash (-) will be appended to the prefix followed by the node number to form a hostname.\nThis default naming scheme can be manually overridden in the node configuration.\nThe maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must\nstart with an alphanumeric character.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_control_ip_address": {
				Description: "The hypervisor control virtual IP address for the HyperFlex compute cluster that is used for node/pod management.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "The hypervisor type for the HyperFlex cluster.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mac_address_prefix": {
				Description: "The MAC address prefix in the form of 00:25:B5:XX.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mgmt_ip_address": {
				Description: "The management IP address for the HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mgmt_platform": {
				Description: "The management platform for the HyperFlex cluster.\n* `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.\n* `EDGE` - The host servers used in the cluster deployment are standalone severs.",
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
			"replication": {
				Description: "The number of copies of each data block written.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_cluster_auxiliary_ip": {
				Description: "The auxiliary storage IP address for the HyperFlex cluster. For two node clusters, this is the IP address of the auxiliary ZK controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"storage_type": {
				Description: "The storage type used for the HyperFlex cluster (HyperFlex Storage or 3rd party).\n* `HyperFlexDp` - The type of storage is HyperFlex Data Platform.\n* `ThirdParty` - The type of storage is 3rd Party Storage (PureStorage, etc..).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted values are instance or template.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"wwxn_prefix": {
				Description: "The WWxN prefix in the form of 20:00:00:25:B5:XX.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexClusterProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterProfile{}
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
	if v, ok := d.GetOk("data_ip_address"); ok {
		x := (v.(string))
		o.SetDataIpAddress(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("host_name_prefix"); ok {
		x := (v.(string))
		o.SetHostNamePrefix(x)
	}
	if v, ok := d.GetOk("hypervisor_control_ip_address"); ok {
		x := (v.(string))
		o.SetHypervisorControlIpAddress(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("mac_address_prefix"); ok {
		x := (v.(string))
		o.SetMacAddressPrefix(x)
	}
	if v, ok := d.GetOk("mgmt_ip_address"); ok {
		x := (v.(string))
		o.SetMgmtIpAddress(x)
	}
	if v, ok := d.GetOk("mgmt_platform"); ok {
		x := (v.(string))
		o.SetMgmtPlatform(x)
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
	if v, ok := d.GetOk("replication"); ok {
		x := int64(v.(int))
		o.SetReplication(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("storage_cluster_auxiliary_ip"); ok {
		x := (v.(string))
		o.SetStorageClusterAuxiliaryIp(x)
	}
	if v, ok := d.GetOk("storage_type"); ok {
		x := (v.(string))
		o.SetStorageType(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("wwxn_prefix"); ok {
		x := (v.(string))
		o.SetWwxnPrefix(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexClusterProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexClusterProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexClusterProfile: %s", responseErr.Error())
	}
	count := countResponse.HyperflexClusterProfileList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexClusterProfile data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexClusterProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexClusterProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexClusterProfile: %s", responseErr.Error())
		}
		results := resMo.HyperflexClusterProfileList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["action"] = (s.GetAction())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["associated_cluster"] = flattenMapHyperflexClusterRelationship(s.GetAssociatedCluster(), d)

				temp["associated_compute_cluster"] = flattenMapHyperflexHxapClusterRelationship(s.GetAssociatedComputeCluster(), d)

				temp["auto_support"] = flattenMapHyperflexAutoSupportPolicyRelationship(s.GetAutoSupport(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_network"] = flattenMapHyperflexClusterNetworkPolicyRelationship(s.GetClusterNetwork(), d)

				temp["cluster_storage"] = flattenMapHyperflexClusterStoragePolicyRelationship(s.GetClusterStorage(), d)

				temp["config_context"] = flattenMapPolicyConfigContext(s.GetConfigContext(), d)

				temp["config_result"] = flattenMapHyperflexConfigResultRelationship(s.GetConfigResult(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["data_ip_address"] = (s.GetDataIpAddress())
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["ext_fc_storage"] = flattenMapHyperflexExtFcStoragePolicyRelationship(s.GetExtFcStorage(), d)

				temp["ext_iscsi_storage"] = flattenMapHyperflexExtIscsiStoragePolicyRelationship(s.GetExtIscsiStorage(), d)
				temp["host_name_prefix"] = (s.GetHostNamePrefix())

				temp["httpproxypolicy"] = flattenMapCommHttpProxyPolicyRelationship(s.GetHttpproxypolicy(), d)
				temp["hypervisor_control_ip_address"] = (s.GetHypervisorControlIpAddress())
				temp["hypervisor_type"] = (s.GetHypervisorType())

				temp["local_credential"] = flattenMapHyperflexLocalCredentialPolicyRelationship(s.GetLocalCredential(), d)
				temp["mac_address_prefix"] = (s.GetMacAddressPrefix())
				temp["mgmt_ip_address"] = (s.GetMgmtIpAddress())
				temp["mgmt_platform"] = (s.GetMgmtPlatform())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["node_config"] = flattenMapHyperflexNodeConfigPolicyRelationship(s.GetNodeConfig(), d)

				temp["node_profile_config"] = flattenListHyperflexNodeProfileRelationship(s.GetNodeProfileConfig(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["policy_bucket"] = flattenListPolicyAbstractPolicyRelationship(s.GetPolicyBucket(), d)

				temp["proxy_setting"] = flattenMapHyperflexProxySettingPolicyRelationship(s.GetProxySetting(), d)
				temp["replication"] = (s.GetReplication())

				temp["running_workflows"] = flattenListWorkflowWorkflowInfoRelationship(s.GetRunningWorkflows(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["software_version"] = flattenMapHyperflexSoftwareVersionPolicyRelationship(s.GetSoftwareVersion(), d)

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)
				temp["storage_cluster_auxiliary_ip"] = (s.GetStorageClusterAuxiliaryIp())

				temp["storage_data_vlan"] = flattenMapHyperflexNamedVlan(s.GetStorageDataVlan(), d)
				temp["storage_type"] = (s.GetStorageType())

				temp["sys_config"] = flattenMapHyperflexSysConfigPolicyRelationship(s.GetSysConfig(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())

				temp["ucsm_config"] = flattenMapHyperflexUcsmConfigPolicyRelationship(s.GetUcsmConfig(), d)

				temp["vcenter_config"] = flattenMapHyperflexVcenterConfigPolicyRelationship(s.GetVcenterConfig(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["wwxn_prefix"] = (s.GetWwxnPrefix())
				hyperflexClusterProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterProfileResults))
	if err := d.Set("results", hyperflexClusterProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterProfileResults[0]["moid"].(string))
	return de
}
