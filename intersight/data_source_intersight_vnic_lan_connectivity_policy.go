package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicLanConnectivityPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicLanConnectivityPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"iqn_allocation_type": {
				Description: "Allocation Type of iSCSI Qualified Name - Static/Pool/None.\n* `None` - Type indicates that there is no IQN associated to an interface.\n* `Static` - Type represents that static IQN is associated to an interface.\n* `Pool` - Type indicates that IQN value is sourced from an associated pool.",
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
			"placement_mode": {
				Description: "The mode used for placement of vNICs on network adapters. It can either be Auto or Custom.\n* `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.\n* `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"static_iqn_name": {
				Description: "User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"target_platform": {
				Description: "The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.\n* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.\n* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicLanConnectivityPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicLanConnectivityPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicLanConnectivityPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("iqn_allocation_type"); ok {
		x := (v.(string))
		o.SetIqnAllocationType(x)
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
	if v, ok := d.GetOk("placement_mode"); ok {
		x := (v.(string))
		o.SetPlacementMode(x)
	}
	if v, ok := d.GetOk("static_iqn_name"); ok {
		x := (v.(string))
		o.SetStaticIqnName(x)
	}
	if v, ok := d.GetOk("target_platform"); ok {
		x := (v.(string))
		o.SetTargetPlatform(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicLanConnectivityPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicLanConnectivityPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicLanConnectivityPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicLanConnectivityPolicyList.GetCount()
	var i int32
	var vnicLanConnectivityPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicLanConnectivityPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicLanConnectivityPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicLanConnectivityPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicLanConnectivityPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["eth_ifs"] = flattenListVnicEthIfRelationship(s.GetEthIfs(), d)
				temp["iqn_allocation_type"] = (s.GetIqnAllocationType())

				temp["iqn_pool"] = flattenMapIqnpoolPoolRelationship(s.GetIqnPool(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["placement_mode"] = (s.GetPlacementMode())

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["static_iqn_name"] = (s.GetStaticIqnName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_platform"] = (s.GetTargetPlatform())
				vnicLanConnectivityPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicLanConnectivityPolicyResults))
	if err := d.Set("results", vnicLanConnectivityPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicLanConnectivityPolicyResults[0]["moid"].(string))
	return de
}
