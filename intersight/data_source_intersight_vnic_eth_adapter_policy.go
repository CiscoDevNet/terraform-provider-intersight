package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicEthAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicEthAdapterPolicyRead,
		Schema: map[string]*schema.Schema{
			"advanced_filter": {
				Description: "Enables advanced filtering on the interface.",
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
			"interrupt_scaling": {
				Description: "Enables Interrupt Scaling on the interface.",
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
			"rss_settings": {
				Description: "Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"uplink_failback_timeout": {
				Description: "Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicEthAdapterPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicEthAdapterPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicEthAdapterPolicy{}
	if v, ok := d.GetOk("advanced_filter"); ok {
		x := (v.(bool))
		o.SetAdvancedFilter(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("interrupt_scaling"); ok {
		x := (v.(bool))
		o.SetInterruptScaling(x)
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
	if v, ok := d.GetOk("rss_settings"); ok {
		x := (v.(bool))
		o.SetRssSettings(x)
	}
	if v, ok := d.GetOk("uplink_failback_timeout"); ok {
		x := int64(v.(int))
		o.SetUplinkFailbackTimeout(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicEthAdapterPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthAdapterPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicEthAdapterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicEthAdapterPolicyList.GetCount()
	var i int32
	var vnicEthAdapterPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthAdapterPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicEthAdapterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicEthAdapterPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicEthAdapterPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["advanced_filter"] = (s.GetAdvancedFilter())

				temp["arfs_settings"] = flattenMapVnicArfsSettings(s.GetArfsSettings(), d)
				temp["class_id"] = (s.GetClassId())

				temp["completion_queue_settings"] = flattenMapVnicCompletionQueueSettings(s.GetCompletionQueueSettings(), d)
				temp["description"] = (s.GetDescription())
				temp["interrupt_scaling"] = (s.GetInterruptScaling())

				temp["interrupt_settings"] = flattenMapVnicEthInterruptSettings(s.GetInterruptSettings(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["nvgre_settings"] = flattenMapVnicNvgreSettings(s.GetNvgreSettings(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["roce_settings"] = flattenMapVnicRoceSettings(s.GetRoceSettings(), d)

				temp["rss_hash_settings"] = flattenMapVnicRssHashSettings(s.GetRssHashSettings(), d)
				temp["rss_settings"] = (s.GetRssSettings())

				temp["rx_queue_settings"] = flattenMapVnicEthRxQueueSettings(s.GetRxQueueSettings(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tcp_offload_settings"] = flattenMapVnicTcpOffloadSettings(s.GetTcpOffloadSettings(), d)

				temp["tx_queue_settings"] = flattenMapVnicEthTxQueueSettings(s.GetTxQueueSettings(), d)
				temp["uplink_failback_timeout"] = (s.GetUplinkFailbackTimeout())

				temp["vxlan_settings"] = flattenMapVnicVxlanSettings(s.GetVxlanSettings(), d)
				vnicEthAdapterPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicEthAdapterPolicyResults))
	if err := d.Set("results", vnicEthAdapterPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicEthAdapterPolicyResults[0]["moid"].(string))
	return de
}
