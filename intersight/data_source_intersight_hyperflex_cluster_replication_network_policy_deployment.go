package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexClusterReplicationNetworkPolicyDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterReplicationNetworkPolicyDeploymentRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_uuid": {
				Description: "Uuid of the HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description from corresponding ClusterReplicationNetworkPolicy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"discovered": {
				Description: "True if record created by discovery on HyperFlex cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name from corresponding ClusterReplicationNetworkPolicy.",
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
			"policy_moid": {
				Description: "Deployed network policy moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"profile_moid": {
				Description: "Deployed cluster profile moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"replication_bandwidth_mbps": {
				Description: "Bandwidth for the Replication network in Mbps.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"replication_mtu": {
				Description: "MTU for the Replication network.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"request_id": {
				Description: "Unique request ID allowing retry of the same logical request following a transient communication failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexClusterReplicationNetworkPolicyDeployment().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterReplicationNetworkPolicyDeploymentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterReplicationNetworkPolicyDeployment{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_uuid"); ok {
		x := (v.(string))
		o.SetClusterUuid(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("discovered"); ok {
		x := (v.(bool))
		o.SetDiscovered(x)
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
	if v, ok := d.GetOk("policy_moid"); ok {
		x := (v.(string))
		o.SetPolicyMoid(x)
	}
	if v, ok := d.GetOk("profile_moid"); ok {
		x := (v.(string))
		o.SetProfileMoid(x)
	}
	if v, ok := d.GetOk("replication_bandwidth_mbps"); ok {
		x := int64(v.(int))
		o.SetReplicationBandwidthMbps(x)
	}
	if v, ok := d.GetOk("replication_mtu"); ok {
		x := int64(v.(int))
		o.SetReplicationMtu(x)
	}
	if v, ok := d.GetOk("request_id"); ok {
		x := (v.(string))
		o.SetRequestId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexClusterReplicationNetworkPolicyDeployment object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterReplicationNetworkPolicyDeploymentList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexClusterReplicationNetworkPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexClusterReplicationNetworkPolicyDeploymentList.GetCount()
	var i int32
	var hyperflexClusterReplicationNetworkPolicyDeploymentResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterReplicationNetworkPolicyDeploymentList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexClusterReplicationNetworkPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexClusterReplicationNetworkPolicyDeploymentList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexClusterReplicationNetworkPolicyDeployment data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapHyperflexClusterRelationship(s.GetCluster(), d)
				temp["cluster_uuid"] = (s.GetClusterUuid())
				temp["description"] = (s.GetDescription())
				temp["discovered"] = (s.GetDiscovered())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["policy_moid"] = (s.GetPolicyMoid())
				temp["profile_moid"] = (s.GetProfileMoid())
				temp["replication_bandwidth_mbps"] = (s.GetReplicationBandwidthMbps())

				temp["replication_ipranges"] = flattenListHyperflexIpAddrRange(s.GetReplicationIpranges(), d)
				temp["replication_mtu"] = (s.GetReplicationMtu())

				temp["replication_vlan"] = flattenMapHyperflexNamedVlan(s.GetReplicationVlan(), d)
				temp["request_id"] = (s.GetRequestId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				hyperflexClusterReplicationNetworkPolicyDeploymentResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterReplicationNetworkPolicyDeploymentResults))
	if err := d.Set("results", hyperflexClusterReplicationNetworkPolicyDeploymentResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterReplicationNetworkPolicyDeploymentResults[0]["moid"].(string))
	return de
}
