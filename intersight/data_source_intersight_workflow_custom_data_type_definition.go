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

func dataSourceWorkflowCustomDataTypeDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowCustomDataTypeDefinitionRead,
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
			"composite_type": {
				Description: "When true this data type definition is a collection of type definitions to represent composite data like JSON.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "A human-friendly description of this custom data type indicating it's domain and usage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"label": {
				Description: "A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_).",
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
				Description: "The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceWorkflowCustomDataTypeDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceWorkflowCustomDataTypeDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowCustomDataTypeDefinition{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("composite_type"); ok {
		x := (v.(bool))
		o.SetCompositeType(x)
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
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowCustomDataTypeDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowCustomDataTypeDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of WorkflowCustomDataTypeDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of WorkflowCustomDataTypeDefinition: %s", responseErr.Error())
	}
	count := countResponse.WorkflowCustomDataTypeDefinitionList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for WorkflowCustomDataTypeDefinition data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var workflowCustomDataTypeDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowCustomDataTypeDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching WorkflowCustomDataTypeDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching WorkflowCustomDataTypeDefinition: %s", responseErr.Error())
		}
		results := resMo.WorkflowCustomDataTypeDefinitionList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["catalog"] = flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["composite_type"] = (s.GetCompositeType())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["label"] = (s.GetLabel())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parameter_set"] = flattenListWorkflowParameterSet(s.GetParameterSet(), d)

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["properties"] = flattenMapWorkflowCustomDataTypeProperties(s.GetProperties(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["type_definition"] = flattenListWorkflowBaseDataType(s.GetTypeDefinition(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				workflowCustomDataTypeDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowCustomDataTypeDefinitionResults))
	if err := d.Set("results", workflowCustomDataTypeDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowCustomDataTypeDefinitionResults[0]["moid"].(string))
	return de
}
