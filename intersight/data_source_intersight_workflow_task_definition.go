package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkflowTaskDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowTaskDefinitionRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"default_version": {
				Description: "When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"description": {
				Description: "A user friendly description about task on what operations are done as part of the task execution and any other specific information about task input and output.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"label": {
				Description: "A user friendly short name to identify the task definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"license_entitlement": {
				Description: "License entitlement required to run this task. It is determined by license requirement of features.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.\n* `IWO-Essential` - IWO-Essential as a License type.\n* `IWO-Advantage` - IWO-Advantage as a License type.\n* `IWO-Premier` - IWO-Premier as a License type.",
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
				Description: "The name of the task definition. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), or an underscore (_). Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"secure_prop_access": {
				Description: "If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"nr_version": {
				Description: "The version of the task definition so we can support multiple versions of a task definition.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceWorkflowTaskDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceWorkflowTaskDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowTaskDefinition{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("default_version"); ok {
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("license_entitlement"); ok {
		x := (v.(string))
		o.SetLicenseEntitlement(x)
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
	if v, ok := d.GetOk("secure_prop_access"); ok {
		x := (v.(bool))
		o.SetSecurePropAccess(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := int64(v.(int))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowTaskDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowTaskDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of WorkflowTaskDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.WorkflowTaskDefinitionList.GetCount()
	var i int32
	var workflowTaskDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowTaskDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowTaskDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.WorkflowTaskDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for WorkflowTaskDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["catalog"] = flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["default_version"] = (s.GetDefaultVersion())
				temp["description"] = (s.GetDescription())

				temp["implemented_tasks"] = flattenListWorkflowTaskDefinitionRelationship(s.GetImplementedTasks(), d)

				temp["interface_task"] = flattenMapWorkflowTaskDefinitionRelationship(s.GetInterfaceTask(), d)

				temp["internal_properties"] = flattenMapWorkflowInternalProperties(s.GetInternalProperties(), d)
				temp["label"] = (s.GetLabel())
				temp["license_entitlement"] = (s.GetLicenseEntitlement())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["properties"] = flattenMapWorkflowProperties(s.GetProperties(), d)

				temp["rollback_tasks"] = flattenListWorkflowRollbackTask(s.GetRollbackTasks(), d)
				temp["secure_prop_access"] = (s.GetSecurePropAccess())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["task_metadata"] = flattenMapWorkflowTaskMetadataRelationship(s.GetTaskMetadata(), d)
				temp["nr_version"] = (s.GetVersion())
				workflowTaskDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowTaskDefinitionResults))
	if err := d.Set("results", workflowTaskDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowTaskDefinitionResults[0]["moid"].(string))
	return de
}
