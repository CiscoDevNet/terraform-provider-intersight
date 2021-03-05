package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkflowWorkflowMeta() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowWorkflowMetaRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The description for the workflow.",
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
				Description: "The name given to the workflow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"retryable": {
				Description: "When true, this workflow can be retried for 2 weeks since the last modification of the workflow.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"schema_version": {
				Description: "The Conductor schema version that decides what attribute can be supported.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"src": {
				Description: "The src is workflow owner service.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "The type of workflow definition.\n* `SystemDefined` - System defined workflow definition.\n* `UserDefined` - User defined workflow definition.\n* `Dynamic` - Dynamically defined workflow definition.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "The version for the workflow so we can support multiple versions for the same workflow name.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"wait_on_duplicate": {
				Description: "Parameter decides if workflows will wait for a duplicate to finish before starting a new one.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"description": {
						Description: "The description for the workflow.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"input_parameters": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "The name given to the workflow.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"organization": {
						Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"output_parameters": {
						Description: "The workflow output parameters.",
						Type:        schema.TypeMap,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						}, Optional: true,
					},
					"retryable": {
						Description: "When true, this workflow can be retried for 2 weeks since the last modification of the workflow.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"schema_version": {
						Description: "The Conductor schema version that decides what attribute can be supported.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"src": {
						Description: "The src is workflow owner service.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"tasks": {
						Description: "The tasks contained inside of the workflow.",
						Type:        schema.TypeMap,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						}, Optional: true,
					},
					"type": {
						Description: "The type of workflow definition.\n* `SystemDefined` - System defined workflow definition.\n* `UserDefined` - User defined workflow definition.\n* `Dynamic` - Dynamically defined workflow definition.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version for the workflow so we can support multiple versions for the same workflow name.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"wait_on_duplicate": {
						Description: "Parameter decides if workflows will wait for a duplicate to finish before starting a new one.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceWorkflowWorkflowMetaRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowWorkflowMeta{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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
	if v, ok := d.GetOk("retryable"); ok {
		x := (v.(bool))
		o.SetRetryable(x)
	}
	if v, ok := d.GetOk("schema_version"); ok {
		x := int64(v.(int))
		o.SetSchemaVersion(x)
	}
	if v, ok := d.GetOk("src"); ok {
		x := (v.(string))
		o.SetSrc(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := int64(v.(int))
		o.SetVersion(x)
	}
	if v, ok := d.GetOk("wait_on_duplicate"); ok {
		x := (v.(bool))
		o.SetWaitOnDuplicate(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of WorkflowWorkflowMeta object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowMetaList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of WorkflowWorkflowMeta: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.WorkflowWorkflowMetaList.GetCount()
	var i int32
	var workflowWorkflowMetaResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowMetaList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowWorkflowMeta: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.WorkflowWorkflowMetaList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for WorkflowWorkflowMeta data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["input_parameters"] = (s.GetInputParameters())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["retryable"] = (s.GetRetryable())
				temp["schema_version"] = (s.GetSchemaVersion())
				temp["src"] = (s.GetSrc())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["nr_version"] = (s.GetVersion())
				temp["wait_on_duplicate"] = (s.GetWaitOnDuplicate())
				workflowWorkflowMetaResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(workflowWorkflowMetaResults))
	if err := d.Set("results", workflowWorkflowMetaResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(workflowWorkflowMetaResults[0]["moid"].(string))
	return de
}
