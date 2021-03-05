package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamIdp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamIdpRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_name": {
				Description: "Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_single_logout": {
				Description: "Setting that indicates whether 'Single Logout (SLO)' has been enabled for this IdP.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"idp_entity_id": {
				Description: "The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"metadata": {
				Description: "SAML metadata of the IdP.",
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
				Description: "The name of the Identity Provider, for example Cisco, Okta, or OneID.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Authentication protocol used by the IdP.\n* `saml` - Use SAML as the authentication protocol for sign-on.\n* `oidc` - Open ID connect to be used as an authentication protocol for sign-on.\n* `local` - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamIdp().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamIdpRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamIdp{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("domain_name"); ok {
		x := (v.(string))
		o.SetDomainName(x)
	}
	if v, ok := d.GetOk("enable_single_logout"); ok {
		x := (v.(bool))
		o.SetEnableSingleLogout(x)
	}
	if v, ok := d.GetOk("idp_entity_id"); ok {
		x := (v.(string))
		o.SetIdpEntityId(x)
	}
	if v, ok := d.GetOk("metadata"); ok {
		x := (v.(string))
		o.SetMetadata(x)
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
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamIdp object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamIdpList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamIdp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamIdpList.GetCount()
	var i int32
	var iamIdpResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamIdpList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamIdp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamIdpList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamIdp data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["domain_name"] = (s.GetDomainName())
				temp["enable_single_logout"] = (s.GetEnableSingleLogout())
				temp["idp_entity_id"] = (s.GetIdpEntityId())

				temp["ldap_policy"] = flattenMapIamLdapPolicyRelationship(s.GetLdapPolicy(), d)
				temp["metadata"] = (s.GetMetadata())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["system"] = flattenMapIamSystemRelationship(s.GetSystem(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())

				temp["user_preferences"] = flattenListIamUserPreferenceRelationship(s.GetUserPreferences(), d)

				temp["usergroups"] = flattenListIamUserGroupRelationship(s.GetUsergroups(), d)

				temp["users"] = flattenListIamUserRelationship(s.GetUsers(), d)
				iamIdpResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamIdpResults))
	if err := d.Set("results", iamIdpResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamIdpResults[0]["moid"].(string))
	return de
}
