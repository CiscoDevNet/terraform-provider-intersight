package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceLicenseCustomerOp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseCustomerOpRead,
		Schema: map[string]*schema.Schema{
			"account_license_data": {
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
		},
	}
}

func dataSourceLicenseCustomerOpRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewLicenseCustomerOpWithDefaults()
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.LicenseApi.GetLicenseCustomerOpList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.LicenseCustomerOpList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to LicenseCustomerOp: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewLicenseCustomerOpWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("account_license_data", flattenMapLicenseAccountLicenseDataRelationship(s.GetAccountLicenseData(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property AccountLicenseData: %+v", err)
			}
			if err := d.Set("active_admin", (s.GetActiveAdmin())); err != nil {
				return fmt.Errorf("error occurred while setting property ActiveAdmin: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("all_devices_to_default_tier", (s.GetAllDevicesToDefaultTier())); err != nil {
				return fmt.Errorf("error occurred while setting property AllDevicesToDefaultTier: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("deregister_device", (s.GetDeregisterDevice())); err != nil {
				return fmt.Errorf("error occurred while setting property DeregisterDevice: %+v", err)
			}
			if err := d.Set("enable_trial", (s.GetEnableTrial())); err != nil {
				return fmt.Errorf("error occurred while setting property EnableTrial: %+v", err)
			}
			if err := d.Set("evaluation_period", (s.GetEvaluationPeriod())); err != nil {
				return fmt.Errorf("error occurred while setting property EvaluationPeriod: %+v", err)
			}
			if err := d.Set("extra_evaluation", (s.GetExtraEvaluation())); err != nil {
				return fmt.Errorf("error occurred while setting property ExtraEvaluation: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("renew_authorization", (s.GetRenewAuthorization())); err != nil {
				return fmt.Errorf("error occurred while setting property RenewAuthorization: %+v", err)
			}
			if err := d.Set("renew_id_certificate", (s.GetRenewIdCertificate())); err != nil {
				return fmt.Errorf("error occurred while setting property RenewIdCertificate: %+v", err)
			}
			if err := d.Set("show_agent_tech_support", (s.GetShowAgentTechSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property ShowAgentTechSupport: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
