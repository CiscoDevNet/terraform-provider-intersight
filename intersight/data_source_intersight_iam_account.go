package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamAccountRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
				Description: "Name of the Intersight account. By default, name is same as the MoID of the account.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the account. To activate the Intersight account, claim a device to the account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamAccount().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamAccountRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamAccount{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamAccount object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamAccountList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamAccount: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamAccountList.GetCount()
	var i int32
	var iamAccountResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamAccountList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamAccount: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamAccountList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamAccount data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["app_registrations"] = flattenListIamAppRegistrationRelationship(s.GetAppRegistrations(), d)
				temp["class_id"] = (s.GetClassId())

				temp["domain_groups"] = flattenListIamDomainGroupRelationship(s.GetDomainGroups(), d)

				temp["end_point_roles"] = flattenListIamEndPointRoleRelationship(s.GetEndPointRoles(), d)

				temp["idpreferences"] = flattenListIamIdpReferenceRelationship(s.GetIdpreferences(), d)

				temp["idps"] = flattenListIamIdpRelationship(s.GetIdps(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["permissions"] = flattenListIamPermissionRelationship(s.GetPermissions(), d)

				temp["privilege_sets"] = flattenListIamPrivilegeSetRelationship(s.GetPrivilegeSets(), d)

				temp["privileges"] = flattenListIamPrivilegeRelationship(s.GetPrivileges(), d)

				temp["resource_limits"] = flattenMapIamResourceLimitsRelationship(s.GetResourceLimits(), d)

				temp["roles"] = flattenListIamRoleRelationship(s.GetRoles(), d)

				temp["security_holder"] = flattenMapIamSecurityHolderRelationship(s.GetSecurityHolder(), d)

				temp["session_limits"] = flattenMapIamSessionLimitsRelationship(s.GetSessionLimits(), d)
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iamAccountResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamAccountResults))
	if err := d.Set("results", iamAccountResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamAccountResults[0]["moid"].(string))
	return de
}
