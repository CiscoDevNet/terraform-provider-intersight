package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesNetworkPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cni_type": {
				Description: "Supported CNI type. Currently we only support Calico.\n* `Calico` - Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin.\n* `Aci` - Cisco ACI Container Network Interface plugin.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
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
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod_network_cidr": {
				Description: "CIDR block to allocate pod network IP addresses from.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"service_cidr": {
				Description: "CIDR block to allocate cluster service IP addresses from.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKubernetesNetworkPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKubernetesNetworkPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesNetworkPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cni_type"); ok {
		x := (v.(string))
		o.SetCniType(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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
	if v, ok := d.GetOk("pod_network_cidr"); ok {
		x := (v.(string))
		o.SetPodNetworkCidr(x)
	}
	if v, ok := d.GetOk("service_cidr"); ok {
		x := (v.(string))
		o.SetServiceCidr(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KubernetesNetworkPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesNetworkPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of KubernetesNetworkPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.KubernetesNetworkPolicyList.GetCount()
	var i int32
	var kubernetesNetworkPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KubernetesApi.GetKubernetesNetworkPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesNetworkPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.KubernetesNetworkPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for KubernetesNetworkPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profiles"] = flattenListKubernetesClusterProfileRelationship(s.GetClusterProfiles(), d)

				temp["cni_config"] = flattenMapKubernetesCniConfig(s.GetCniConfig(), d)
				temp["cni_type"] = (s.GetCniType())
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["pod_network_cidr"] = (s.GetPodNetworkCidr())
				temp["service_cidr"] = (s.GetServiceCidr())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				kubernetesNetworkPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kubernetesNetworkPolicyResults))
	if err := d.Set("results", kubernetesNetworkPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kubernetesNetworkPolicyResults[0]["moid"].(string))
	return de
}
