package intersight

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseLicenseInfo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLicenseLicenseInfoCreate,
		ReadContext:   resourceLicenseLicenseInfoRead,
		UpdateContext: resourceLicenseLicenseInfoUpdate,
		DeleteContext: resourceLicenseLicenseInfoDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
				Description: "The name of the Intersight license entitlement.\nFor example, this property may be set to 'Essential'.\n* `Base` - Base as a License type. It is default license type.\n* `Essential` - Essential as a License type.\n* `Standard` - Standard as a License type.\n* `Advantage` - Advantage as a License type.\n* `Premier` - Premier as a License type.\n* `IWO-Essential` - IWO-Essential as a License type.\n* `IWO-Advantage` - IWO-Advantage as a License type.\n* `IWO-Premier` - IWO-Premier as a License type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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

func resourceLicenseLicenseInfoCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewLicenseLicenseInfoWithDefaults()
	if v, ok := d.GetOk("account_license_data"); ok {
		p := make([]models.LicenseAccountLicenseDataRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsLicenseAccountLicenseDataRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccountLicenseData(x)
		}
	}

	if v, ok := d.GetOkExists("active_admin"); ok {
		x := v.(bool)
		o.SetActiveAdmin(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("license.LicenseInfo")

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

	o.SetObjectType("license.LicenseInfo")

	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOkExists("trial_admin"); ok {
		x := v.(bool)
		o.SetTrialAdmin(x)
	}

	r := conn.ApiClient.LicenseApi.CreateLicenseLicenseInfo(conn.ctx).LicenseLicenseInfo(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating LicenseLicenseInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceLicenseLicenseInfoRead(c, d, meta)
}

func resourceLicenseLicenseInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.LicenseApi.GetLicenseLicenseInfoByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "LicenseLicenseInfo object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching LicenseLicenseInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("account_license_data", flattenMapLicenseAccountLicenseDataRelationship(s.GetAccountLicenseData(), d)); err != nil {
		return diag.Errorf("error occurred while setting property AccountLicenseData in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("active_admin", (s.GetActiveAdmin())); err != nil {
		return diag.Errorf("error occurred while setting property ActiveAdmin in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("days_left", (s.GetDaysLeft())); err != nil {
		return diag.Errorf("error occurred while setting property DaysLeft in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("end_time", (s.GetEndTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property EndTime in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("enforce_mode", (s.GetEnforceMode())); err != nil {
		return diag.Errorf("error occurred while setting property EnforceMode in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("error_desc", (s.GetErrorDesc())); err != nil {
		return diag.Errorf("error occurred while setting property ErrorDesc in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("evaluation_period", (s.GetEvaluationPeriod())); err != nil {
		return diag.Errorf("error occurred while setting property EvaluationPeriod in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("extra_evaluation", (s.GetExtraEvaluation())); err != nil {
		return diag.Errorf("error occurred while setting property ExtraEvaluation in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("license_count", (s.GetLicenseCount())); err != nil {
		return diag.Errorf("error occurred while setting property LicenseCount in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("license_state", (s.GetLicenseState())); err != nil {
		return diag.Errorf("error occurred while setting property LicenseState in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("license_type", (s.GetLicenseType())); err != nil {
		return diag.Errorf("error occurred while setting property LicenseType in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("start_time", (s.GetStartTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property StartTime in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in LicenseLicenseInfo object: %s", err.Error())
	}

	if err := d.Set("trial_admin", (s.GetTrialAdmin())); err != nil {
		return diag.Errorf("error occurred while setting property TrialAdmin in LicenseLicenseInfo object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceLicenseLicenseInfoUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.LicenseLicenseInfo{}
	if d.HasChange("account_license_data") {
		v := d.Get("account_license_data")
		p := make([]models.LicenseAccountLicenseDataRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsLicenseAccountLicenseDataRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccountLicenseData(x)
		}
	}

	if d.HasChange("active_admin") {
		v := d.Get("active_admin")
		x := (v.(bool))
		o.SetActiveAdmin(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("license.LicenseInfo")

	if d.HasChange("days_left") {
		v := d.Get("days_left")
		x := int64(v.(int))
		o.SetDaysLeft(x)
	}

	if d.HasChange("end_time") {
		v := d.Get("end_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}

	if d.HasChange("enforce_mode") {
		v := d.Get("enforce_mode")
		x := (v.(string))
		o.SetEnforceMode(x)
	}

	if d.HasChange("error_desc") {
		v := d.Get("error_desc")
		x := (v.(string))
		o.SetErrorDesc(x)
	}

	if d.HasChange("evaluation_period") {
		v := d.Get("evaluation_period")
		x := int64(v.(int))
		o.SetEvaluationPeriod(x)
	}

	if d.HasChange("extra_evaluation") {
		v := d.Get("extra_evaluation")
		x := int64(v.(int))
		o.SetExtraEvaluation(x)
	}

	if d.HasChange("license_count") {
		v := d.Get("license_count")
		x := int64(v.(int))
		o.SetLicenseCount(x)
	}

	if d.HasChange("license_state") {
		v := d.Get("license_state")
		x := (v.(string))
		o.SetLicenseState(x)
	}

	if d.HasChange("license_type") {
		v := d.Get("license_type")
		x := (v.(string))
		o.SetLicenseType(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("license.LicenseInfo")

	if d.HasChange("start_time") {
		v := d.Get("start_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if d.HasChange("trial_admin") {
		v := d.Get("trial_admin")
		x := (v.(bool))
		o.SetTrialAdmin(x)
	}

	r := conn.ApiClient.LicenseApi.UpdateLicenseLicenseInfo(conn.ctx, d.Id()).LicenseLicenseInfo(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating LicenseLicenseInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceLicenseLicenseInfoRead(c, d, meta)
}

func resourceLicenseLicenseInfoDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	var warning = diag.Diagnostic{Severity: diag.Warning, Summary: "LicenseLicenseInfo does not allow delete functionality"}
	de = append(de, warning)
	return de
}
