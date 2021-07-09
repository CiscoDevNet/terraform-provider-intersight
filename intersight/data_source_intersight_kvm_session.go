package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKvmSession() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKvmSessionRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_ip_address": {
				Description: "The user agent IP address from which the session is launched.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The time at which the session ended.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"one_time_password": {
				Description: "Temporary one-time password for vKVM access.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"role": {
				Description: "Role of the user who launched the session.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sso_supported": {
				Description: "Indicates if vKVM SSO is supported on the server.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "The status of the session.\n* `Active` - The session is currently active.\n* `Ended` - The session has ended normally.\n* `Terminated` - The session was terminated by an admin.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"target_name": {
				Description: "Name of target on which session is initiated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user_id_or_email": {
				Description: "User ID or E-mail Address of the user who launched the session.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Description: "Username used for vKVM access.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceKvmSession().Schema},
				Computed: true,
			}},
	}
}

func dataSourceKvmSessionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KvmSession{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("client_ip_address"); ok {
		x := (v.(string))
		o.SetClientIpAddress(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("one_time_password"); ok {
		x := (v.(string))
		o.SetOneTimePassword(x)
	}
	if v, ok := d.GetOk("role"); ok {
		x := (v.(string))
		o.SetRole(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("sso_supported"); ok {
		x := (v.(bool))
		o.SetSsoSupported(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("target_name"); ok {
		x := (v.(string))
		o.SetTargetName(x)
	}
	if v, ok := d.GetOk("user_id_or_email"); ok {
		x := (v.(string))
		o.SetUserIdOrEmail(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of KvmSession object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.KvmApi.GetKvmSessionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of KvmSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of KvmSession: %s", responseErr.Error())
	}
	count := countResponse.KvmSessionList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for KvmSession data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var kvmSessionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.KvmApi.GetKvmSessionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching KvmSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching KvmSession: %s", responseErr.Error())
		}
		results := resMo.KvmSessionList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())
				temp["client_ip_address"] = (s.GetClientIpAddress())

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["end_time"] = (s.GetEndTime()).String()

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["one_time_password"] = (s.GetOneTimePassword())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["role"] = (s.GetRole())

				temp["server"] = flattenMapComputePhysicalRelationship(s.GetServer(), d)

				temp["session"] = flattenMapSessionAbstractSessionRelationship(s.GetSession(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["sso_supported"] = (s.GetSsoSupported())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["target"] = flattenMapMoBaseMoRelationship(s.GetTarget(), d)
				temp["target_name"] = (s.GetTargetName())

				temp["tunnel"] = flattenMapKvmTunnelRelationship(s.GetTunnel(), d)

				temp["user"] = flattenMapIamUserRelationship(s.GetUser(), d)
				temp["user_id_or_email"] = (s.GetUserIdOrEmail())
				temp["username"] = (s.GetUsername())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				kvmSessionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(kvmSessionResults))
	if err := d.Set("results", kvmSessionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(kvmSessionResults[0]["moid"].(string))
	return de
}
