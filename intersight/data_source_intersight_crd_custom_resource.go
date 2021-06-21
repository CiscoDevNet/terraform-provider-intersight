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

func dataSourceCrdCustomResource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCrdCustomResourceRead,
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
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dc_launcher": {
				Description: "Type of custom resource or 'kind' in K8s.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"image": {
				Description: "Docker image URL for public cloud DC.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Description: "A string property called name which is used as part of a uniqueness constraint. See 'identity' specification in this MO definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"namespace": {
				Description: "Namespace to launch the deployment associated with the custom resource.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Description: "Port used for public cloud DC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"target_id": {
				Description: "Target ID for public cloud DC.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"target_moid": {
				Description: "Target Moid for public cloud DC.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"target_type": {
				Description: "Target type for public cloud DC.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCrdCustomResource().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCrdCustomResourceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CrdCustomResource{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("dc_launcher"); ok {
		x := (v.(string))
		o.SetDcLauncher(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("image"); ok {
		x := (v.(string))
		o.SetImage(x)
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
	if v, ok := d.GetOk("namespace"); ok {
		x := (v.(string))
		o.SetNamespace(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("port"); ok {
		x := int64(v.(int))
		o.SetPort(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("target_id"); ok {
		x := (v.(string))
		o.SetTargetId(x)
	}
	if v, ok := d.GetOk("target_moid"); ok {
		x := (v.(string))
		o.SetTargetMoid(x)
	}
	if v, ok := d.GetOk("target_type"); ok {
		x := (v.(string))
		o.SetTargetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CrdCustomResource object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CrdApi.GetCrdCustomResourceList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of CrdCustomResource: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of CrdCustomResource: %s", responseErr.Error())
	}
	count := countResponse.CrdCustomResourceList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for CrdCustomResource data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var crdCustomResourceResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CrdApi.GetCrdCustomResourceList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching CrdCustomResource: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching CrdCustomResource: %s", responseErr.Error())
		}
		results := resMo.CrdCustomResourceList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["dc_launcher"] = (s.GetDcLauncher())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["image"] = (s.GetImage())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["namespace"] = (s.GetNamespace())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["port"] = (s.GetPort())

				temp["properties"] = flattenListCrdCustomResourceConfigProperty(s.GetProperties(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_id"] = (s.GetTargetId())
				temp["target_moid"] = (s.GetTargetMoid())
				temp["target_type"] = (s.GetTargetType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				crdCustomResourceResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(crdCustomResourceResults))
	if err := d.Set("results", crdCustomResourceResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(crdCustomResourceResults[0]["moid"].(string))
	return de
}
