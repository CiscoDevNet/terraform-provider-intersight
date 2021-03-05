package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricEthNetworkControlPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricEthNetworkControlPolicyRead,
		Schema: map[string]*schema.Schema{
			"cdp_enabled": {
				Description: "Enables the CDP on an interface.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
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
			"forge_mac": {
				Description: "Determines if the MAC forging is allowed or denied on an interface.\n* `allow` - Allows mac forging on an interface.\n* `deny` - Denies mac forging on an interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mac_registration_mode": {
				Description: "Determines the MAC addresses that have to be registered with the switch.\n* `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN.\n* `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs.",
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
			"uplink_fail_action": {
				Description: "Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned.\n* `linkDown` - The vethernet will go down in case a suitable uplink is not pinned.\n* `warning` - The vethernet will remain up even if a suitable uplink is not pinned.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricEthNetworkControlPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricEthNetworkControlPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricEthNetworkControlPolicy{}
	if v, ok := d.GetOk("cdp_enabled"); ok {
		x := (v.(bool))
		o.SetCdpEnabled(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("forge_mac"); ok {
		x := (v.(string))
		o.SetForgeMac(x)
	}
	if v, ok := d.GetOk("mac_registration_mode"); ok {
		x := (v.(string))
		o.SetMacRegistrationMode(x)
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
	if v, ok := d.GetOk("uplink_fail_action"); ok {
		x := (v.(string))
		o.SetUplinkFailAction(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricEthNetworkControlPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricEthNetworkControlPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricEthNetworkControlPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricEthNetworkControlPolicyList.GetCount()
	var i int32
	var fabricEthNetworkControlPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricEthNetworkControlPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricEthNetworkControlPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricEthNetworkControlPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricEthNetworkControlPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["cdp_enabled"] = (s.GetCdpEnabled())
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["forge_mac"] = (s.GetForgeMac())

				temp["lldp_settings"] = flattenMapFabricLldpSettings(s.GetLldpSettings(), d)
				temp["mac_registration_mode"] = (s.GetMacRegistrationMode())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["network_policy"] = flattenListVnicEthNetworkPolicyRelationship(s.GetNetworkPolicy(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uplink_fail_action"] = (s.GetUplinkFailAction())
				fabricEthNetworkControlPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricEthNetworkControlPolicyResults))
	if err := d.Set("results", fabricEthNetworkControlPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricEthNetworkControlPolicyResults[0]["moid"].(string))
	return de
}
