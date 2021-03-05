package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMemoryPersistentMemoryPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMemoryPersistentMemoryPolicyRead,
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
			"management_mode": {
				Description: "Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System.\n* `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy.\n* `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools.",
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
			"retain_namespaces": {
				Description: "Persistent Memory Namespaces to be retained or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceMemoryPersistentMemoryPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceMemoryPersistentMemoryPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.MemoryPersistentMemoryPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
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
	if v, ok := d.GetOk("retain_namespaces"); ok {
		x := (v.(bool))
		o.SetRetainNamespaces(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of MemoryPersistentMemoryPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.MemoryApi.GetMemoryPersistentMemoryPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of MemoryPersistentMemoryPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.MemoryPersistentMemoryPolicyList.GetCount()
	var i int32
	var memoryPersistentMemoryPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.MemoryApi.GetMemoryPersistentMemoryPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching MemoryPersistentMemoryPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.MemoryPersistentMemoryPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for MemoryPersistentMemoryPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["goals"] = flattenListMemoryPersistentMemoryGoal(s.GetGoals(), d)

				temp["local_security"] = flattenMapMemoryPersistentMemoryLocalSecurity(s.GetLocalSecurity(), d)

				temp["logical_namespaces"] = flattenListMemoryPersistentMemoryLogicalNamespace(s.GetLogicalNamespaces(), d)
				temp["management_mode"] = (s.GetManagementMode())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["retain_namespaces"] = (s.GetRetainNamespaces())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				memoryPersistentMemoryPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(memoryPersistentMemoryPolicyResults))
	if err := d.Set("results", memoryPersistentMemoryPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(memoryPersistentMemoryPolicyResults[0]["moid"].(string))
	return de
}
