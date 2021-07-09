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

func dataSourceVmrcConsole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVmrcConsoleRead,
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
			"client_url": {
				Description: "The multiplexer URL for the client to connect on.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVmrcConsole().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVmrcConsoleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VmrcConsole{}
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
	if v, ok := d.GetOk("client_url"); ok {
		x := (v.(string))
		o.SetClientUrl(x)
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
	if v, ok := d.GetOk("role"); ok {
		x := (v.(string))
		o.SetRole(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VmrcConsole object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VmrcApi.GetVmrcConsoleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of VmrcConsole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of VmrcConsole: %s", responseErr.Error())
	}
	count := countResponse.VmrcConsoleList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for VmrcConsole data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var vmrcConsoleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VmrcApi.GetVmrcConsoleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching VmrcConsole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching VmrcConsole: %s", responseErr.Error())
		}
		results := resMo.VmrcConsoleList.GetResults()
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
				temp["client_url"] = (s.GetClientUrl())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["end_time"] = (s.GetEndTime()).String()

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["role"] = (s.GetRole())

				temp["session"] = flattenMapSessionAbstractSessionRelationship(s.GetSession(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["target"] = flattenMapMoBaseMoRelationship(s.GetTarget(), d)
				temp["target_name"] = (s.GetTargetName())

				temp["user"] = flattenMapIamUserRelationship(s.GetUser(), d)
				temp["user_id_or_email"] = (s.GetUserIdOrEmail())

				temp["vcenter"] = flattenMapVirtualizationVmwareVcenterRelationship(s.GetVcenter(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["virtual_machine"] = flattenMapVirtualizationVmwareVirtualMachineRelationship(s.GetVirtualMachine(), d)
				vmrcConsoleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vmrcConsoleResults))
	if err := d.Set("results", vmrcConsoleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vmrcConsoleResults[0]["moid"].(string))
	return de
}
