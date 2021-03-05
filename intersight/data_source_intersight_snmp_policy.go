package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSnmpPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSnmpPolicyRead,
		Schema: map[string]*schema.Schema{
			"access_community_string": {
				Description: "The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"community_access": {
				Description: "Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.\n* `Disabled` - Blocks access to the information in the inventory tables.\n* `Limited` - Partial access to read the information in the inventory tables.\n* `Full` - Full access to read the information in the inventory tables.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"engine_id": {
				Description: "User-defined unique identification of the static engine.",
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
			"snmp_port": {
				Description: "Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"sys_contact": {
				Description: "Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sys_location": {
				Description: "Location of host on which the SNMP agent (server) runs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"trap_community": {
				Description: "SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSnmpPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSnmpPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SnmpPolicy{}
	if v, ok := d.GetOk("access_community_string"); ok {
		x := (v.(string))
		o.SetAccessCommunityString(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("community_access"); ok {
		x := (v.(string))
		o.SetCommunityAccess(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("engine_id"); ok {
		x := (v.(string))
		o.SetEngineId(x)
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
	if v, ok := d.GetOk("snmp_port"); ok {
		x := int64(v.(int))
		o.SetSnmpPort(x)
	}
	if v, ok := d.GetOk("sys_contact"); ok {
		x := (v.(string))
		o.SetSysContact(x)
	}
	if v, ok := d.GetOk("sys_location"); ok {
		x := (v.(string))
		o.SetSysLocation(x)
	}
	if v, ok := d.GetOk("trap_community"); ok {
		x := (v.(string))
		o.SetTrapCommunity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SnmpPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SnmpApi.GetSnmpPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SnmpPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SnmpPolicyList.GetCount()
	var i int32
	var snmpPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SnmpApi.GetSnmpPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SnmpPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SnmpPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SnmpPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["access_community_string"] = (s.GetAccessCommunityString())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["community_access"] = (s.GetCommunityAccess())
				temp["description"] = (s.GetDescription())
				temp["enabled"] = (s.GetEnabled())
				temp["engine_id"] = (s.GetEngineId())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["snmp_port"] = (s.GetSnmpPort())

				temp["snmp_traps"] = flattenListSnmpTrap(s.GetSnmpTraps(), d)

				temp["snmp_users"] = flattenListSnmpUser(s.GetSnmpUsers(), d)
				temp["sys_contact"] = (s.GetSysContact())
				temp["sys_location"] = (s.GetSysLocation())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["trap_community"] = (s.GetTrapCommunity())
				snmpPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(snmpPolicyResults))
	if err := d.Set("results", snmpPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(snmpPolicyResults[0]["moid"].(string))
	return de
}
