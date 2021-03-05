package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkconfigPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkconfigPolicyRead,
		Schema: map[string]*schema.Schema{
			"alternate_ipv4dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"alternate_ipv6dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
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
			"dynamic_dns_domain": {
				Description: "The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_dynamic_dns": {
				Description: "If enabled, updates the resource records to the DNS from Cisco IMC.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv4dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6": {
				Description: "If enabled, allows to configure IPv6 properties.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.",
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
			"preferred_ipv4dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"preferred_ipv6dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceNetworkconfigPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceNetworkconfigPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkconfigPolicy{}
	if v, ok := d.GetOk("alternate_ipv4dns_server"); ok {
		x := (v.(string))
		o.SetAlternateIpv4dnsServer(x)
	}
	if v, ok := d.GetOk("alternate_ipv6dns_server"); ok {
		x := (v.(string))
		o.SetAlternateIpv6dnsServer(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("dynamic_dns_domain"); ok {
		x := (v.(string))
		o.SetDynamicDnsDomain(x)
	}
	if v, ok := d.GetOk("enable_dynamic_dns"); ok {
		x := (v.(bool))
		o.SetEnableDynamicDns(x)
	}
	if v, ok := d.GetOk("enable_ipv4dns_from_dhcp"); ok {
		x := (v.(bool))
		o.SetEnableIpv4dnsFromDhcp(x)
	}
	if v, ok := d.GetOk("enable_ipv6"); ok {
		x := (v.(bool))
		o.SetEnableIpv6(x)
	}
	if v, ok := d.GetOk("enable_ipv6dns_from_dhcp"); ok {
		x := (v.(bool))
		o.SetEnableIpv6dnsFromDhcp(x)
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
	if v, ok := d.GetOk("preferred_ipv4dns_server"); ok {
		x := (v.(string))
		o.SetPreferredIpv4dnsServer(x)
	}
	if v, ok := d.GetOk("preferred_ipv6dns_server"); ok {
		x := (v.(string))
		o.SetPreferredIpv6dnsServer(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NetworkconfigPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NetworkconfigApi.GetNetworkconfigPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NetworkconfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NetworkconfigPolicyList.GetCount()
	var i int32
	var networkconfigPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NetworkconfigApi.GetNetworkconfigPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NetworkconfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NetworkconfigPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NetworkconfigPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["alternate_ipv4dns_server"] = (s.GetAlternateIpv4dnsServer())
				temp["alternate_ipv6dns_server"] = (s.GetAlternateIpv6dnsServer())

				temp["appliance_account"] = flattenMapIamAccountRelationship(s.GetApplianceAccount(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["dynamic_dns_domain"] = (s.GetDynamicDnsDomain())
				temp["enable_dynamic_dns"] = (s.GetEnableDynamicDns())
				temp["enable_ipv4dns_from_dhcp"] = (s.GetEnableIpv4dnsFromDhcp())
				temp["enable_ipv6"] = (s.GetEnableIpv6())
				temp["enable_ipv6dns_from_dhcp"] = (s.GetEnableIpv6dnsFromDhcp())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["preferred_ipv4dns_server"] = (s.GetPreferredIpv4dnsServer())
				temp["preferred_ipv6dns_server"] = (s.GetPreferredIpv6dnsServer())

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				networkconfigPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(networkconfigPolicyResults))
	if err := d.Set("results", networkconfigPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(networkconfigPolicyResults[0]["moid"].(string))
	return de
}
