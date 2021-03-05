package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCapabilitySwitchCapability() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilitySwitchCapabilityRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_fcoe_vlan": {
				Description: "Default Fcoe VLAN associated with this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"dynamic_vifs_supported": {
				Description: "Dynamic VIFs support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"fan_modules_supported": {
				Description: "Fan Modules support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"fc_uplink_ports_auto_negotiation_supported": {
				Description: "Fc Uplink ports auto negotiation speed support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"locator_beacon_supported": {
				Description: "Locator Beacon LED support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_ports": {
				Description: "Maximum allowed physical ports on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_slots": {
				Description: "Maximum allowed physical slots on this switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "An unique identifer for a capability descriptor.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sereno_netflow_supported": {
				Description: "Sereno Adaptor with Netflow support on this switch.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"sku": {
				Description: "SKU information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"unified_rule": {
				Description: "The Slider rule for Unified ports on this switch.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCapabilitySwitchCapability().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCapabilitySwitchCapabilityRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilitySwitchCapability{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_fcoe_vlan"); ok {
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}
	if v, ok := d.GetOk("dynamic_vifs_supported"); ok {
		x := (v.(bool))
		o.SetDynamicVifsSupported(x)
	}
	if v, ok := d.GetOk("fan_modules_supported"); ok {
		x := (v.(bool))
		o.SetFanModulesSupported(x)
	}
	if v, ok := d.GetOk("fc_uplink_ports_auto_negotiation_supported"); ok {
		x := (v.(bool))
		o.SetFcUplinkPortsAutoNegotiationSupported(x)
	}
	if v, ok := d.GetOk("locator_beacon_supported"); ok {
		x := (v.(bool))
		o.SetLocatorBeaconSupported(x)
	}
	if v, ok := d.GetOk("max_ports"); ok {
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}
	if v, ok := d.GetOk("max_slots"); ok {
		x := int64(v.(int))
		o.SetMaxSlots(x)
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
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("sereno_netflow_supported"); ok {
		x := (v.(bool))
		o.SetSerenoNetflowSupported(x)
	}
	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}
	if v, ok := d.GetOk("unified_rule"); ok {
		x := (v.(string))
		o.SetUnifiedRule(x)
	}
	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilitySwitchCapability object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CapabilitySwitchCapabilityList.GetCount()
	var i int32
	var capabilitySwitchCapabilityResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CapabilitySwitchCapabilityList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CapabilitySwitchCapability data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["default_fcoe_vlan"] = (s.GetDefaultFcoeVlan())
				temp["dynamic_vifs_supported"] = (s.GetDynamicVifsSupported())
				temp["fan_modules_supported"] = (s.GetFanModulesSupported())

				temp["fc_end_host_mode_reserved_vsans"] = flattenListCapabilityPortRange(s.GetFcEndHostModeReservedVsans(), d)
				temp["fc_uplink_ports_auto_negotiation_supported"] = (s.GetFcUplinkPortsAutoNegotiationSupported())
				temp["locator_beacon_supported"] = (s.GetLocatorBeaconSupported())
				temp["max_ports"] = (s.GetMaxPorts())
				temp["max_slots"] = (s.GetMaxSlots())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["network_limits"] = flattenMapCapabilitySwitchNetworkLimits(s.GetNetworkLimits(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["pid"] = (s.GetPid())

				temp["ports_supporting100g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting100gSpeed(), d)

				temp["ports_supporting10g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting10gSpeed(), d)

				temp["ports_supporting1g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting1gSpeed(), d)

				temp["ports_supporting25g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting25gSpeed(), d)

				temp["ports_supporting40g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting40gSpeed(), d)

				temp["ports_supporting_breakout"] = flattenListCapabilityPortRange(s.GetPortsSupportingBreakout(), d)

				temp["ports_supporting_fcoe"] = flattenListCapabilityPortRange(s.GetPortsSupportingFcoe(), d)

				temp["ports_supporting_server_role"] = flattenListCapabilityPortRange(s.GetPortsSupportingServerRole(), d)

				temp["reserved_vsans"] = flattenListCapabilityPortRange(s.GetReservedVsans(), d)
				temp["sereno_netflow_supported"] = (s.GetSerenoNetflowSupported())
				temp["sku"] = (s.GetSku())

				temp["storage_limits"] = flattenMapCapabilitySwitchStorageLimits(s.GetStorageLimits(), d)

				temp["switching_mode_capabilities"] = flattenListCapabilitySwitchingModeCapability(s.GetSwitchingModeCapabilities(), d)

				temp["system_limits"] = flattenMapCapabilitySwitchSystemLimits(s.GetSystemLimits(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["unified_ports"] = flattenListCapabilityPortRange(s.GetUnifiedPorts(), d)
				temp["unified_rule"] = (s.GetUnifiedRule())
				temp["vid"] = (s.GetVid())
				capabilitySwitchCapabilityResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilitySwitchCapabilityResults))
	if err := d.Set("results", capabilitySwitchCapabilityResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilitySwitchCapabilityResults[0]["moid"].(string))
	return de
}
