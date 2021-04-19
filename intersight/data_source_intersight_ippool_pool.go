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

func dataSourceIppoolPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIppoolPoolRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"assigned": {
				Description: "Number of IDs that are currently assigned.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"assignment_order": {
				Description: "Assignment order decides the order in which the next identifier is allocated.\n* `sequential` - Identifiers are assigned in a sequential order.\n* `default` - Assignment order is decided by the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
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
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Description: "Total number of identifiers in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v4_assigned": {
				Description: "Number of IPv4 addresses currently in use.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v4_size": {
				Description: "Number of IPv4 addresses in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v6_assigned": {
				Description: "Number of IPv6 addresses currently in use.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v6_size": {
				Description: "Number of IPv6 addresses in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIppoolPool().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIppoolPoolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IppoolPool{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("assigned"); ok {
		x := int64(v.(int))
		o.SetAssigned(x)
	}
	if v, ok := d.GetOk("assignment_order"); ok {
		x := (v.(string))
		o.SetAssignmentOrder(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("v4_assigned"); ok {
		x := int64(v.(int))
		o.SetV4Assigned(x)
	}
	if v, ok := d.GetOk("v4_size"); ok {
		x := int64(v.(int))
		o.SetV4Size(x)
	}
	if v, ok := d.GetOk("v6_assigned"); ok {
		x := int64(v.(int))
		o.SetV6Assigned(x)
	}
	if v, ok := d.GetOk("v6_size"); ok {
		x := int64(v.(int))
		o.SetV6Size(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IppoolPool object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IppoolApi.GetIppoolPoolList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of IppoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of IppoolPool: %s", responseErr.Error())
	}
	count := countResponse.IppoolPoolList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for IppoolPool data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var ippoolPoolResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IppoolApi.GetIppoolPoolList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching IppoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching IppoolPool: %s", responseErr.Error())
		}
		results := resMo.IppoolPoolList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["assigned"] = (s.GetAssigned())
				temp["assignment_order"] = (s.GetAssignmentOrder())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["ip_v4_blocks"] = flattenListIppoolIpV4Block(s.GetIpV4Blocks(), d)

				temp["ip_v4_config"] = flattenMapIppoolIpV4Config(s.GetIpV4Config(), d)

				temp["ip_v6_blocks"] = flattenListIppoolIpV6Block(s.GetIpV6Blocks(), d)

				temp["ip_v6_config"] = flattenMapIppoolIpV6Config(s.GetIpV6Config(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["shadow_pools"] = flattenListIppoolShadowPoolRelationship(s.GetShadowPools(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["size"] = (s.GetSize())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["v4_assigned"] = (s.GetV4Assigned())
				temp["v4_size"] = (s.GetV4Size())
				temp["v6_assigned"] = (s.GetV6Assigned())
				temp["v6_size"] = (s.GetV6Size())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				ippoolPoolResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(ippoolPoolResults))
	if err := d.Set("results", ippoolPoolResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(ippoolPoolResults[0]["moid"].(string))
	return de
}
