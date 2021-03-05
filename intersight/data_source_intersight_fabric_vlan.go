package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricVlan() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricVlanRead,
		Schema: map[string]*schema.Schema{
			"auto_allow_on_uplinks": {
				Description: "Used to determine whether this VLAN will be allowed on all uplink ports and PCs in this FI.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_native": {
				Description: "Used to define whether this VLAN is to be classified as 'native' for traffic in this FI.",
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
				Description: "The 'name' used to identify this VLAN.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vlan_id": {
				Description: "The identifier for this Virtual LAN.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricVlan().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricVlanRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricVlan{}
	if v, ok := d.GetOk("auto_allow_on_uplinks"); ok {
		x := (v.(bool))
		o.SetAutoAllowOnUplinks(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("is_native"); ok {
		x := (v.(bool))
		o.SetIsNative(x)
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
	if v, ok := d.GetOk("vlan_id"); ok {
		x := int64(v.(int))
		o.SetVlanId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricVlan object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricVlanList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricVlan: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricVlanList.GetCount()
	var i int32
	var fabricVlanResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricVlanList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricVlan: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricVlanList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricVlan data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["auto_allow_on_uplinks"] = (s.GetAutoAllowOnUplinks())
				temp["class_id"] = (s.GetClassId())

				temp["eth_network_policy"] = flattenMapFabricEthNetworkPolicyRelationship(s.GetEthNetworkPolicy(), d)
				temp["is_native"] = (s.GetIsNative())
				temp["moid"] = (s.GetMoid())

				temp["multicast_policy"] = flattenMapFabricMulticastPolicyRelationship(s.GetMulticastPolicy(), d)
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vlan_id"] = (s.GetVlanId())
				fabricVlanResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricVlanResults))
	if err := d.Set("results", fabricVlanResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricVlanResults[0]["moid"].(string))
	return de
}
