package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryMsoContractDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryMsoContractDetailsRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"contract_name": {
				Description: "Name of contract in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployed_sites": {
				Description: "Site Ids to which this contract is deployed to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_local": {
				Description: "Is the contract local to the Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"reference": {
				Description: "Unique reference for the contract in the Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"schema_id": {
				Description: "Schema ID in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"schema_name": {
				Description: "Schema name this contract belongs to in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"template_name": {
				Description: "Template name this contract belongs to in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
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
					"contract_name": {
						Description: "Name of contract in Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"deployed_sites": {
						Description: "Site Ids to which this contract is deployed to.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_local": {
						Description: "Is the contract local to the Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"reference": {
						Description: "Unique reference for the contract in the Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
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
					},
					"schema_id": {
						Description: "Schema ID in Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"schema_name": {
						Description: "Schema name this contract belongs to in Multi-Site Orchestrator.",
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
					"template_name": {
						Description: "Template name this contract belongs to in Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryMsoContractDetailsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryMsoContractDetails{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("contract_name"); ok {
		x := (v.(string))
		o.SetContractName(x)
	}
	if v, ok := d.GetOk("deployed_sites"); ok {
		x := (v.(string))
		o.SetDeployedSites(x)
	}
	if v, ok := d.GetOk("is_local"); ok {
		x := (v.(string))
		o.SetIsLocal(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("reference"); ok {
		x := (v.(string))
		o.SetReference(x)
	}
	if v, ok := d.GetOk("schema_id"); ok {
		x := (v.(string))
		o.SetSchemaId(x)
	}
	if v, ok := d.GetOk("schema_name"); ok {
		x := (v.(string))
		o.SetSchemaName(x)
	}
	if v, ok := d.GetOk("template_name"); ok {
		x := (v.(string))
		o.SetTemplateName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryMsoContractDetails object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryMsoContractDetailsList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryMsoContractDetails: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryMsoContractDetailsList.GetCount()
	var i int32
	var niatelemetryMsoContractDetailsResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryMsoContractDetailsList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryMsoContractDetails: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryMsoContractDetailsList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryMsoContractDetails data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["contract_name"] = (s.GetContractName())
				temp["deployed_sites"] = (s.GetDeployedSites())
				temp["is_local"] = (s.GetIsLocal())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["reference"] = (s.GetReference())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["schema_id"] = (s.GetSchemaId())
				temp["schema_name"] = (s.GetSchemaName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["template_name"] = (s.GetTemplateName())
				niatelemetryMsoContractDetailsResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryMsoContractDetailsResults))
	if err := d.Set("results", niatelemetryMsoContractDetailsResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryMsoContractDetailsResults[0]["moid"].(string))
	return de
}
