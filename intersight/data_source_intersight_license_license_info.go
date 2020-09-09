package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceLicenseLicenseInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseLicenseInfoRead,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"days_left": {
				Description: "The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "The date and time when the trial period expires.\nThe value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"enforce_mode": {
				Description: "The entitlement mode reported by Cisco Smart Software Manager.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"error_desc": {
				Description: "The detailed error message when there is any error related to this licensing entitlement.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"evaluation_period": {
				Description: "The default Trial or Grace period customer is entitled to.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"extra_evaluation": {
				Description: "The number of days the trial Trial or Grace period is extended.\nThe trial or grace period can be extended once.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"license_count": {
				Description: "The total number of devices claimed in the Intersight account.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"license_state": {
				Description: "The license state defined by Intersight.\nThe value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired.\n* `NotLicensed` - The license token is neither activated nor registered.\n* `GraceExpired` - The license grace period has expired.\n* `TrialPeriod` - The 90 days of trial period.\n* `OutOfCompliance` - The license is out of compliance.\n* `Compliance` - The license is in compliance.\n* `TrialExpired` - The trial period of 90 days has expired.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"license_type": {
				Description: "The name of the Intersight license entitlement.\nFor example, this property may be set to 'Essential'.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state.",
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
			"trial_admin": {
				Description: "The administrative state of the trial license.\nWhen the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period,\ni.e. licenseState is set to be TrialPeriod.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceLicenseLicenseInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewLicenseLicenseInfoWithDefaults()
	if v, ok := d.GetOk("active_admin"); ok {
		x := (v.(bool))
		o.SetActiveAdmin(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("days_left"); ok {
		x := int64(v.(int))
		o.SetDaysLeft(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("enforce_mode"); ok {
		x := (v.(string))
		o.SetEnforceMode(x)
	}
	if v, ok := d.GetOk("error_desc"); ok {
		x := (v.(string))
		o.SetErrorDesc(x)
	}
	if v, ok := d.GetOk("evaluation_period"); ok {
		x := int64(v.(int))
		o.SetEvaluationPeriod(x)
	}
	if v, ok := d.GetOk("extra_evaluation"); ok {
		x := int64(v.(int))
		o.SetExtraEvaluation(x)
	}
	if v, ok := d.GetOk("license_count"); ok {
		x := int64(v.(int))
		o.SetLicenseCount(x)
	}
	if v, ok := d.GetOk("license_state"); ok {
		x := (v.(string))
		o.SetLicenseState(x)
	}
	if v, ok := d.GetOk("license_type"); ok {
		x := (v.(string))
		o.SetLicenseType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("trial_admin"); ok {
		x := (v.(bool))
		o.SetTrialAdmin(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.LicenseApi.GetLicenseLicenseInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.LicenseLicenseInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to LicenseLicenseInfo: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewLicenseLicenseInfoWithDefaults()
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
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("days_left", (s.GetDaysLeft())); err != nil {
				return fmt.Errorf("error occurred while setting property DaysLeft: %+v", err)
			}

			if err := d.Set("end_time", (s.GetEndTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndTime: %+v", err)
			}
			if err := d.Set("enforce_mode", (s.GetEnforceMode())); err != nil {
				return fmt.Errorf("error occurred while setting property EnforceMode: %+v", err)
			}
			if err := d.Set("error_desc", (s.GetErrorDesc())); err != nil {
				return fmt.Errorf("error occurred while setting property ErrorDesc: %+v", err)
			}
			if err := d.Set("evaluation_period", (s.GetEvaluationPeriod())); err != nil {
				return fmt.Errorf("error occurred while setting property EvaluationPeriod: %+v", err)
			}
			if err := d.Set("extra_evaluation", (s.GetExtraEvaluation())); err != nil {
				return fmt.Errorf("error occurred while setting property ExtraEvaluation: %+v", err)
			}
			if err := d.Set("license_count", (s.GetLicenseCount())); err != nil {
				return fmt.Errorf("error occurred while setting property LicenseCount: %+v", err)
			}
			if err := d.Set("license_state", (s.GetLicenseState())); err != nil {
				return fmt.Errorf("error occurred while setting property LicenseState: %+v", err)
			}
			if err := d.Set("license_type", (s.GetLicenseType())); err != nil {
				return fmt.Errorf("error occurred while setting property LicenseType: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("start_time", (s.GetStartTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property StartTime: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("trial_admin", (s.GetTrialAdmin())); err != nil {
				return fmt.Errorf("error occurred while setting property TrialAdmin: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
