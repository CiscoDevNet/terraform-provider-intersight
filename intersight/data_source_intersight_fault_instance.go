package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFaultInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFaultInstanceRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Description: "The user acknowledgement state of the fault.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_dn": {
				Description: "The Distinguished Name of the Managed object which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_id": {
				Description: "Managed object Id which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_type": {
				Description: "Managed object type which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ancestor_mo_id": {
				Description: "Object Id of the parent of the Managed object which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ancestor_mo_type": {
				Description: "Object type of the parent of the Managed object which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"code": {
				Description: "Numerical fault code of the fault found.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"creation_time": {
				Description: "The time of creation of the fault instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Detailed message of the fault.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_transition_time": {
				Description: "Last transition time of the fault.",
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
			"num_occurrences": {
				Description: "The number of times this fault has occured.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"original_severity": {
				Description: "Current Severity of the fault found.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"previous_severity": {
				Description: "The Severity of the fault prior to user update.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rule": {
				Description: "The rule that is responsible for generation of the fault.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Description: "Severity of the fault found.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"acknowledged": {
					Description: "The user acknowledgement state of the fault.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"affected_dn": {
						Description: "The Distinguished Name of the Managed object which was affected.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"affected_mo_id": {
						Description: "Managed object Id which was affected.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"affected_mo_type": {
						Description: "Managed object type which was affected.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ancestor_mo_id": {
						Description: "Object Id of the parent of the Managed object which was affected.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ancestor_mo_type": {
						Description: "Object type of the parent of the Managed object which was affected.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"code": {
						Description: "Numerical fault code of the fault found.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"creation_time": {
						Description: "The time of creation of the fault instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"description": {
						Description: "Detailed message of the fault.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_mo_id": {
						Description: "The database identifier of the registered device of an object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"dn": {
						Description: "The Distinguished Name unambiguously identifies an object in the system.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"inventory_device_info": {
						Description: "A reference to a inventoryDeviceInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"last_transition_time": {
						Description: "Last transition time of the fault.",
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
					"num_occurrences": {
						Description: "The number of times this fault has occured.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"original_severity": {
						Description: "Current Severity of the fault found.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"previous_severity": {
						Description: "The Severity of the fault prior to user update.",
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
					"rn": {
						Description: "The Relative Name uniquely identifies an object within a given context.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"rule": {
						Description: "The rule that is responsible for generation of the fault.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"severity": {
						Description: "Severity of the fault found.",
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

func dataSourceFaultInstanceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FaultInstance{}
	if v, ok := d.GetOk("acknowledged"); ok {
		x := (v.(string))
		o.SetAcknowledged(x)
	}
	if v, ok := d.GetOk("affected_dn"); ok {
		x := (v.(string))
		o.SetAffectedDn(x)
	}
	if v, ok := d.GetOk("affected_mo_id"); ok {
		x := (v.(string))
		o.SetAffectedMoId(x)
	}
	if v, ok := d.GetOk("affected_mo_type"); ok {
		x := (v.(string))
		o.SetAffectedMoType(x)
	}
	if v, ok := d.GetOk("ancestor_mo_id"); ok {
		x := (v.(string))
		o.SetAncestorMoId(x)
	}
	if v, ok := d.GetOk("ancestor_mo_type"); ok {
		x := (v.(string))
		o.SetAncestorMoType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("code"); ok {
		x := (v.(string))
		o.SetCode(x)
	}
	if v, ok := d.GetOk("creation_time"); ok {
		x := (v.(string))
		o.SetCreationTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("last_transition_time"); ok {
		x := (v.(string))
		o.SetLastTransitionTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_occurrences"); ok {
		x := int64(v.(int))
		o.SetNumOccurrences(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("original_severity"); ok {
		x := (v.(string))
		o.SetOriginalSeverity(x)
	}
	if v, ok := d.GetOk("previous_severity"); ok {
		x := (v.(string))
		o.SetPreviousSeverity(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("rule"); ok {
		x := (v.(string))
		o.SetRule(x)
	}
	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.SetSeverity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FaultInstance object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FaultApi.GetFaultInstanceList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FaultInstance: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FaultInstanceList.GetCount()
	var i int32
	var faultInstanceResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FaultApi.GetFaultInstanceList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FaultInstance: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FaultInstanceList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FaultInstance data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["acknowledged"] = (s.GetAcknowledged())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["affected_dn"] = (s.GetAffectedDn())
				temp["affected_mo_id"] = (s.GetAffectedMoId())
				temp["affected_mo_type"] = (s.GetAffectedMoType())
				temp["ancestor_mo_id"] = (s.GetAncestorMoId())
				temp["ancestor_mo_type"] = (s.GetAncestorMoType())
				temp["class_id"] = (s.GetClassId())
				temp["code"] = (s.GetCode())
				temp["creation_time"] = (s.GetCreationTime())
				temp["description"] = (s.GetDescription())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["last_transition_time"] = (s.GetLastTransitionTime())
				temp["moid"] = (s.GetMoid())
				temp["num_occurrences"] = (s.GetNumOccurrences())
				temp["object_type"] = (s.GetObjectType())
				temp["original_severity"] = (s.GetOriginalSeverity())
				temp["previous_severity"] = (s.GetPreviousSeverity())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["rn"] = (s.GetRn())
				temp["rule"] = (s.GetRule())
				temp["severity"] = (s.GetSeverity())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				faultInstanceResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(faultInstanceResults))
	if err := d.Set("results", faultInstanceResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(faultInstanceResults[0]["moid"].(string))
	return de
}
