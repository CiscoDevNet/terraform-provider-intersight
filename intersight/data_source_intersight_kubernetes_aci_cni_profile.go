package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesAciCniProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAciCniProfileRead,
		Schema: map[string]*schema.Schema{
			"aaep_name": {
				Description: "Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ext_svc_dyn_subnet_start": {
				Description: "Start of range of IP subnets for external services with dynamic IP allocation for use by Kubernetes clusters using this ACI CNI policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ext_svc_static_subnet_start": {
				Description: "Start of range of IP subnets for external services with static IP allocation for use by Kubernetes clusters using this ACI CNI policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"infra_vlan_id": {
				Description: "Value of ACI infrastructuere VLAN ID for the ACI fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"l3_out_network_name": {
				Description: "Name of ACI L3Out network to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"l3_out_policy_name": {
				Description: "Name of ACI L3Out policy to be used for all Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"l3_out_tenant": {
				Description: "Tenant in ACI used by this L3Out and Common VRF.",
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
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nested_vmm_domain": {
				Description: "VMM domain within which Kubernetes clusters using this policy are nested.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"node_svc_subnet_start": {
				Description: "Start of range of ACI Node Service IP subnets to use by Kubernetes clusters using this ACI CNI policy This is used for the service graph which is used for ACI PBR based load balancing.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"node_vlan_range_end": {
				Description: "Ending value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"node_vlan_range_start": {
				Description: "Starting value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_kubernetes_clusters": {
				Description: "Number of k8s clusters currently using this ACI CNI profile.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"opflex_multicast_address_range": {
				Description: "Range of IP Multicast addresses to be used by the Opflex protocol for Kubernetes clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pod_subnet_start": {
				Description: "Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"svc_subnet_start": {
				Description: "Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vrf": {
				Description: "VRF (Virtual Routing and Forwarding) domain to be used within ACI fabric by all k8s clusters using this policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesAciCniProfile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesAciCniProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAciCniProfile{}
	if v, ok := d.GetOk("aaep_name"); ok {
		x := (v.(string))
		o.SetAaepName(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("ext_svc_dyn_subnet_start"); ok {
		x := (v.(string))
		o.SetExtSvcDynSubnetStart(x)
	}
	if v, ok := d.GetOk("ext_svc_static_subnet_start"); ok {
		x := (v.(string))
		o.SetExtSvcStaticSubnetStart(x)
	}
	if v, ok := d.GetOk("infra_vlan_id"); ok {
		x := int64(v.(int))
		o.SetInfraVlanId(x)
	}
	if v, ok := d.GetOk("l3_out_network_name"); ok {
		x := (v.(string))
		o.SetL3OutNetworkName(x)
	}
	if v, ok := d.GetOk("l3_out_policy_name"); ok {
		x := (v.(string))
		o.SetL3OutPolicyName(x)
	}
	if v, ok := d.GetOk("l3_out_tenant"); ok {
		x := (v.(string))
		o.SetL3OutTenant(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("nested_vmm_domain"); ok {
		x := (v.(string))
		o.SetNestedVmmDomain(x)
	}
	if v, ok := d.GetOk("node_svc_subnet_start"); ok {
		x := (v.(string))
		o.SetNodeSvcSubnetStart(x)
	}
	if v, ok := d.GetOk("node_vlan_range_end"); ok {
		x := int64(v.(int))
		o.SetNodeVlanRangeEnd(x)
	}
	if v, ok := d.GetOk("node_vlan_range_start"); ok {
		x := int64(v.(int))
		o.SetNodeVlanRangeStart(x)
	}
	if v, ok := d.GetOk("number_of_kubernetes_clusters"); ok {
		x := int64(v.(int))
		o.SetNumberOfKubernetesClusters(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("opflex_multicast_address_range"); ok {
		x := (v.(string))
		o.SetOpflexMulticastAddressRange(x)
	}
	if v, ok := d.GetOk("pod_subnet_start"); ok {
		x := (v.(string))
		o.SetPodSubnetStart(x)
	}
	if v, ok := d.GetOk("svc_subnet_start"); ok {
		x := (v.(string))
		o.SetSvcSubnetStart(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("vrf"); ok {
		x := (v.(string))
		o.SetVrf(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesAciCniProfile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAciCniProfileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesAciCniProfileList.GetCount()
	var i int32
	var kubernetesAciCniProfileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAciCniProfileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesAciCniProfileList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesAciCniProfile data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["aaep_name"] = (s.GetAaepName())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_aci_allocations"] = flattenListKubernetesAciCniTenantClusterAllocationRelationship(s.GetClusterAciAllocations(), d)

				temp["cluster_profiles"] = flattenListKubernetesClusterProfileRelationship(s.GetClusterProfiles(), d)
				temp["description"] = (s.GetDescription())
				temp["ext_svc_dyn_subnet_start"] = (s.GetExtSvcDynSubnetStart())
				temp["ext_svc_static_subnet_start"] = (s.GetExtSvcStaticSubnetStart())
				temp["infra_vlan_id"] = (s.GetInfraVlanId())
				temp["l3_out_network_name"] = (s.GetL3OutNetworkName())
				temp["l3_out_policy_name"] = (s.GetL3OutPolicyName())
				temp["l3_out_tenant"] = (s.GetL3OutTenant())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["nested_vmm_domain"] = (s.GetNestedVmmDomain())
				temp["node_svc_subnet_start"] = (s.GetNodeSvcSubnetStart())
				temp["node_vlan_range_end"] = (s.GetNodeVlanRangeEnd())
				temp["node_vlan_range_start"] = (s.GetNodeVlanRangeStart())
				temp["number_of_kubernetes_clusters"] = (s.GetNumberOfKubernetesClusters())
				temp["object_type"] = (s.GetObjectType())
				temp["opflex_multicast_address_range"] = (s.GetOpflexMulticastAddressRange())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["pod_subnet_start"] = (s.GetPodSubnetStart())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["src_template"] = flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)
				temp["svc_subnet_start"] = (s.GetSvcSubnetStart())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["vrf"] = (s.GetVrf())
				kubernetesAciCniProfileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesAciCniProfileResults))
	if err := d.Set("results", kubernetesAciCniProfileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesAciCniProfileResults[0]["moid"].(string))
	return de
}
