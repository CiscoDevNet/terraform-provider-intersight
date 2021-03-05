package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseCustomerOp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLicenseCustomerOpRead,
		Schema: map[string]*schema.Schema{
			"active_admin": {
				Description: "The license administrative state.\nSet this property to 'true' to activate the license entitlements.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"all_devices_to_default_tier": {
				Description: "Move all licensed devices to default license tier.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deregister_device": {
				Description: "Trigger de-registration/disable.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_trial": {
				Description: "Enable trial for Intersight licensing.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"evaluation_period": {
				Description: "The default Trial or Grace period customer is entitled to.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"extra_evaluation": {
				Description: "The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.",
				Type:        schema.TypeInt,
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
			"renew_authorization": {
				Description: "Trigger renew authorization.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"renew_id_certificate": {
				Description: "Trigger renew registration.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"show_agent_tech_support": {
				Description: "Trigger show tech support feature.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_license_data": {
					Description: "A reference to a licenseAccountLicenseData resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"active_admin": {
						Description: "The license administrative state.\nSet this property to 'true' to activate the license entitlements.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"all_devices_to_default_tier": {
						Description: "Move all licensed devices to default license tier.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"deregister_device": {
						Description: "Trigger de-registration/disable.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"enable_trial": {
						Description: "Enable trial for Intersight licensing.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"evaluation_period": {
						Description: "The default Trial or Grace period customer is entitled to.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"extra_evaluation": {
						Description: "The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.",
						Type:        schema.TypeInt,
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
					"renew_authorization": {
						Description: "Trigger renew authorization.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"renew_id_certificate": {
						Description: "Trigger renew registration.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"show_agent_tech_support": {
						Description: "Trigger show tech support feature.",
						Type:        schema.TypeBool,
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceLicenseCustomerOpRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.LicenseCustomerOp{}
	if v, ok := d.GetOk("active_admin"); ok {
		x := (v.(bool))
		o.SetActiveAdmin(x)
	}
	if v, ok := d.GetOk("all_devices_to_default_tier"); ok {
		x := (v.(bool))
		o.SetAllDevicesToDefaultTier(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("deregister_device"); ok {
		x := (v.(bool))
		o.SetDeregisterDevice(x)
	}
	if v, ok := d.GetOk("enable_trial"); ok {
		x := (v.(bool))
		o.SetEnableTrial(x)
	}
	if v, ok := d.GetOk("evaluation_period"); ok {
		x := int64(v.(int))
		o.SetEvaluationPeriod(x)
	}
	if v, ok := d.GetOk("extra_evaluation"); ok {
		x := int64(v.(int))
		o.SetExtraEvaluation(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("renew_authorization"); ok {
		x := (v.(bool))
		o.SetRenewAuthorization(x)
	}
	if v, ok := d.GetOk("renew_id_certificate"); ok {
		x := (v.(bool))
		o.SetRenewIdCertificate(x)
	}
	if v, ok := d.GetOk("show_agent_tech_support"); ok {
		x := (v.(bool))
		o.SetShowAgentTechSupport(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of LicenseCustomerOp object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseCustomerOpList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of LicenseCustomerOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.LicenseCustomerOpList.GetCount()
	var i int32
	var licenseCustomerOpResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseCustomerOpList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching LicenseCustomerOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.LicenseCustomerOpList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for LicenseCustomerOp data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account_license_data"] = flattenMapLicenseAccountLicenseDataRelationship(s.GetAccountLicenseData(), d)
				temp["active_admin"] = (s.GetActiveAdmin())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["all_devices_to_default_tier"] = (s.GetAllDevicesToDefaultTier())
				temp["class_id"] = (s.GetClassId())
				temp["deregister_device"] = (s.GetDeregisterDevice())
				temp["enable_trial"] = (s.GetEnableTrial())
				temp["evaluation_period"] = (s.GetEvaluationPeriod())
				temp["extra_evaluation"] = (s.GetExtraEvaluation())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["renew_authorization"] = (s.GetRenewAuthorization())
				temp["renew_id_certificate"] = (s.GetRenewIdCertificate())
				temp["show_agent_tech_support"] = (s.GetShowAgentTechSupport())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				licenseCustomerOpResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(licenseCustomerOpResults))
	if err := d.Set("results", licenseCustomerOpResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(licenseCustomerOpResults[0]["moid"].(string))
	return de
}
