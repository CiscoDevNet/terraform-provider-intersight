package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIaasDiagnosticMessages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaasDiagnosticMessagesRead,
		Schema: map[string]*schema.Schema{
			"category": {
				Description: "Message category of the alerts.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"guid": {
				Description: "Unique ID of UCS Director getting registerd with Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"item": {
				Description: "Message target type of the alerts.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_checked": {
				Description: "Last checked time of the alerts.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Description: "Detailed info about the alert.",
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
			"recommendation": {
				Description: "Recommended fix for the alert.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the given alert.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status_id": {
				Description: "Status Id of the given alert.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"category": {
						Description: "Message category of the alerts.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"guid": {
						Description: "Unique ID of UCS Director getting registerd with Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"item": {
						Description: "Message target type of the alerts.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"last_checked": {
						Description: "Last checked time of the alerts.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"message": {
						Description: "Detailed info about the alert.",
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
					"recommendation": {
						Description: "Recommended fix for the alert.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"status": {
						Description: "Status of the given alert.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"status_id": {
						Description: "Status Id of the given alert.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceIaasDiagnosticMessagesRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IaasDiagnosticMessages{}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("guid"); ok {
		x := (v.(string))
		o.SetGuid(x)
	}
	if v, ok := d.GetOk("item"); ok {
		x := (v.(string))
		o.SetItem(x)
	}
	if v, ok := d.GetOk("last_checked"); ok {
		x := (v.(string))
		o.SetLastChecked(x)
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.SetMessage(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.SetRecommendation(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("status_id"); ok {
		x := (v.(string))
		o.SetStatusId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IaasDiagnosticMessages object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IaasApi.GetIaasDiagnosticMessagesList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IaasDiagnosticMessages: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IaasDiagnosticMessagesList.GetCount()
	var i int32
	var iaasDiagnosticMessagesResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IaasApi.GetIaasDiagnosticMessagesList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IaasDiagnosticMessages: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IaasDiagnosticMessagesList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IaasDiagnosticMessages data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["category"] = (s.GetCategory())
				temp["class_id"] = (s.GetClassId())
				temp["guid"] = (s.GetGuid())
				temp["item"] = (s.GetItem())
				temp["last_checked"] = (s.GetLastChecked())
				temp["message"] = (s.GetMessage())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["recommendation"] = (s.GetRecommendation())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["status"] = (s.GetStatus())
				temp["status_id"] = (s.GetStatusId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iaasDiagnosticMessagesResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iaasDiagnosticMessagesResults))
	if err := d.Set("results", iaasDiagnosticMessagesResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iaasDiagnosticMessagesResults[0]["moid"].(string))
	return de
}
