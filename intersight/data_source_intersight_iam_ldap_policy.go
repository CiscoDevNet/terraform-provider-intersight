package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamLdapPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamLdapPolicyRead,
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
			"enable_dns": {
				Description: "Enables DNS to access LDAP servers.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enabled": {
				Description: "LDAP server performs authentication.",
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
			"user_search_precedence": {
				Description: "Search precedence between local user database and LDAP user database.\n* `LocalUserDb` - Precedence is given to local user database while searching.\n* `LDAPUserDb` - Precedence is given to LADP user database while searching.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamLdapPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamLdapPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamLdapPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enable_dns"); ok {
		x := (v.(bool))
		o.SetEnableDns(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
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
	if v, ok := d.GetOk("user_search_precedence"); ok {
		x := (v.(string))
		o.SetUserSearchPrecedence(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamLdapPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamLdapPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamLdapPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamLdapPolicyList.GetCount()
	var i int32
	var iamLdapPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamLdapPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamLdapPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamLdapPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamLdapPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["appliance_account"] = flattenMapIamAccountRelationship(s.GetApplianceAccount(), d)

				temp["base_properties"] = flattenMapIamLdapBaseProperties(s.GetBaseProperties(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["dns_parameters"] = flattenMapIamLdapDnsParameters(s.GetDnsParameters(), d)
				temp["enable_dns"] = (s.GetEnableDns())
				temp["enabled"] = (s.GetEnabled())

				temp["groups"] = flattenListIamLdapGroupRelationship(s.GetGroups(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)

				temp["nr_providers"] = flattenListIamLdapProviderRelationship(s.GetProviders(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["user_search_precedence"] = (s.GetUserSearchPrecedence())
				iamLdapPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamLdapPolicyResults))
	if err := d.Set("results", iamLdapPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamLdapPolicyResults[0]["moid"].(string))
	return de
}
