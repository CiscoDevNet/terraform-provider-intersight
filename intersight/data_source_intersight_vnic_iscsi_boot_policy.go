package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicIscsiBootPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicIscsiBootPolicyRead,
		Schema: map[string]*schema.Schema{
			"auto_targetvendor_name": {
				Description: "Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters.",
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
			"initiator_ip_source": {
				Description: "Source Type of Initiator IP Address - Auto/Static/Pool.\n* `DHCP` - The IP address is assigned using DHCP, if available.\n* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.\n* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"initiator_static_ip_v4_address": {
				Description: "Static IP address provided for iSCSI Initiator.",
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
			"target_source_type": {
				Description: "Source Type of Targets - Auto/Static.\n* `Static` - Type indicates that static target interface is assigned to iSCSI boot.\n* `Auto` - Type indicates that the system selects the target interface automatically during iSCSI boot.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicIscsiBootPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicIscsiBootPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicIscsiBootPolicy{}
	if v, ok := d.GetOk("auto_targetvendor_name"); ok {
		x := (v.(string))
		o.SetAutoTargetvendorName(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("initiator_ip_source"); ok {
		x := (v.(string))
		o.SetInitiatorIpSource(x)
	}
	if v, ok := d.GetOk("initiator_static_ip_v4_address"); ok {
		x := (v.(string))
		o.SetInitiatorStaticIpV4Address(x)
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
	if v, ok := d.GetOk("target_source_type"); ok {
		x := (v.(string))
		o.SetTargetSourceType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicIscsiBootPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicIscsiBootPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicIscsiBootPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicIscsiBootPolicyList.GetCount()
	var i int32
	var vnicIscsiBootPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicIscsiBootPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicIscsiBootPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicIscsiBootPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicIscsiBootPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["auto_targetvendor_name"] = (s.GetAutoTargetvendorName())

				temp["chap"] = flattenMapVnicIscsiAuthProfile(s.GetChap(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["initiator_ip_pool"] = flattenMapIppoolPoolRelationship(s.GetInitiatorIpPool(), d)
				temp["initiator_ip_source"] = (s.GetInitiatorIpSource())
				temp["initiator_static_ip_v4_address"] = (s.GetInitiatorStaticIpV4Address())

				temp["initiator_static_ip_v4_config"] = flattenMapIppoolIpV4Config(s.GetInitiatorStaticIpV4Config(), d)

				temp["iscsi_adapter_policy"] = flattenMapVnicIscsiAdapterPolicyRelationship(s.GetIscsiAdapterPolicy(), d)
				temp["moid"] = (s.GetMoid())

				temp["mutual_chap"] = flattenMapVnicIscsiAuthProfile(s.GetMutualChap(), d)
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["primary_target_policy"] = flattenMapVnicIscsiStaticTargetPolicyRelationship(s.GetPrimaryTargetPolicy(), d)

				temp["secondary_target_policy"] = flattenMapVnicIscsiStaticTargetPolicyRelationship(s.GetSecondaryTargetPolicy(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_source_type"] = (s.GetTargetSourceType())
				vnicIscsiBootPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicIscsiBootPolicyResults))
	if err := d.Set("results", vnicIscsiBootPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicIscsiBootPolicyResults[0]["moid"].(string))
	return de
}
