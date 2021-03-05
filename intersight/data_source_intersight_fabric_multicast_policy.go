package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricMulticastPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricMulticastPolicyRead,
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
			"querier_ip_address": {
				Description: "Used to define the IGMP Querier IP address.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"querier_ip_address_peer": {
				Description: "Used to define the IGMP Querier IP address of the peer switch.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"querier_state": {
				Description: "Administrative state of the IGMP Querier for this VLAN.\n* `Disabled` - IGMP Querier Disabled State.\n* `Enabled` - IGMP Querier Enabled State.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snooping_state": {
				Description: "Administrative state of the IGMP Snooping for this VLAN.\n* `Enabled` - IGMP Snooping Enabled State.\n* `Disabled` - IGMP Snooping Disabled State.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricMulticastPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricMulticastPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricMulticastPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("querier_ip_address"); ok {
		x := (v.(string))
		o.SetQuerierIpAddress(x)
	}
	if v, ok := d.GetOk("querier_ip_address_peer"); ok {
		x := (v.(string))
		o.SetQuerierIpAddressPeer(x)
	}
	if v, ok := d.GetOk("querier_state"); ok {
		x := (v.(string))
		o.SetQuerierState(x)
	}
	if v, ok := d.GetOk("snooping_state"); ok {
		x := (v.(string))
		o.SetSnoopingState(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricMulticastPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricMulticastPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricMulticastPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricMulticastPolicyList.GetCount()
	var i int32
	var fabricMulticastPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricMulticastPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricMulticastPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricMulticastPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricMulticastPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["querier_ip_address"] = (s.GetQuerierIpAddress())
				temp["querier_ip_address_peer"] = (s.GetQuerierIpAddressPeer())
				temp["querier_state"] = (s.GetQuerierState())
				temp["snooping_state"] = (s.GetSnoopingState())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				fabricMulticastPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricMulticastPolicyResults))
	if err := d.Set("results", fabricMulticastPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricMulticastPolicyResults[0]["moid"].(string))
	return de
}
