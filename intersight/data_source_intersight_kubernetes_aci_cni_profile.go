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

func dataSourceKubernetesAciCniProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesAciCniProfileRead,
		Schema: map[string]*schema.Schema{
			"aaep_name": {
				Description: "Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy.",
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
			"cluster_aci_allocations": {
				Description: "An array of relationships to kubernetesAciCniTenantClusterAllocation resources.",
				Type:        schema.TypeList,
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
			"cluster_profiles": {
				Description: "An array of relationships to kubernetesClusterProfile resources.",
				Type:        schema.TypeList,
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
				Computed: true,
			},
			"pod_subnet_start": {
				Description: "Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"src_template": {
				Description: "A reference to a policyAbstractProfile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"svc_subnet_start": {
				Description: "Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
		},
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
	resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesAciCniProfileList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching KubernetesAciCniProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for KubernetesAciCniProfile list: %s", err.Error())
	}
	var s = &models.KubernetesAciCniProfileList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to KubernetesAciCniProfile list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for KubernetesAciCniProfile did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.KubernetesAciCniProfile{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("aaep_name", (s.GetAaepName())); err != nil {
				return diag.Errorf("error occurred while setting property AaepName: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster_aci_allocations", flattenListKubernetesAciCniTenantClusterAllocationRelationship(s.GetClusterAciAllocations(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ClusterAciAllocations: %s", err.Error())
			}

			if err := d.Set("cluster_profiles", flattenListKubernetesClusterProfileRelationship(s.GetClusterProfiles(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ClusterProfiles: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("ext_svc_dyn_subnet_start", (s.GetExtSvcDynSubnetStart())); err != nil {
				return diag.Errorf("error occurred while setting property ExtSvcDynSubnetStart: %s", err.Error())
			}
			if err := d.Set("ext_svc_static_subnet_start", (s.GetExtSvcStaticSubnetStart())); err != nil {
				return diag.Errorf("error occurred while setting property ExtSvcStaticSubnetStart: %s", err.Error())
			}
			if err := d.Set("infra_vlan_id", (s.GetInfraVlanId())); err != nil {
				return diag.Errorf("error occurred while setting property InfraVlanId: %s", err.Error())
			}
			if err := d.Set("l3_out_network_name", (s.GetL3OutNetworkName())); err != nil {
				return diag.Errorf("error occurred while setting property L3OutNetworkName: %s", err.Error())
			}
			if err := d.Set("l3_out_policy_name", (s.GetL3OutPolicyName())); err != nil {
				return diag.Errorf("error occurred while setting property L3OutPolicyName: %s", err.Error())
			}
			if err := d.Set("l3_out_tenant", (s.GetL3OutTenant())); err != nil {
				return diag.Errorf("error occurred while setting property L3OutTenant: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("nested_vmm_domain", (s.GetNestedVmmDomain())); err != nil {
				return diag.Errorf("error occurred while setting property NestedVmmDomain: %s", err.Error())
			}
			if err := d.Set("node_svc_subnet_start", (s.GetNodeSvcSubnetStart())); err != nil {
				return diag.Errorf("error occurred while setting property NodeSvcSubnetStart: %s", err.Error())
			}
			if err := d.Set("node_vlan_range_end", (s.GetNodeVlanRangeEnd())); err != nil {
				return diag.Errorf("error occurred while setting property NodeVlanRangeEnd: %s", err.Error())
			}
			if err := d.Set("node_vlan_range_start", (s.GetNodeVlanRangeStart())); err != nil {
				return diag.Errorf("error occurred while setting property NodeVlanRangeStart: %s", err.Error())
			}
			if err := d.Set("number_of_kubernetes_clusters", (s.GetNumberOfKubernetesClusters())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfKubernetesClusters: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("opflex_multicast_address_range", (s.GetOpflexMulticastAddressRange())); err != nil {
				return diag.Errorf("error occurred while setting property OpflexMulticastAddressRange: %s", err.Error())
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}
			if err := d.Set("pod_subnet_start", (s.GetPodSubnetStart())); err != nil {
				return diag.Errorf("error occurred while setting property PodSubnetStart: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("src_template", flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SrcTemplate: %s", err.Error())
			}
			if err := d.Set("svc_subnet_start", (s.GetSvcSubnetStart())); err != nil {
				return diag.Errorf("error occurred while setting property SvcSubnetStart: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("vrf", (s.GetVrf())); err != nil {
				return diag.Errorf("error occurred while setting property Vrf: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
