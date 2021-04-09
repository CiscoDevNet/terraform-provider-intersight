package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseLicenseInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLicenseLicenseInfoRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"active_admin": {
				Description: "The license administrative state.\nSet this property to 'true' to activate the license entitlements.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
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
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
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
			"mod_time": {
				Description: "The time when this managed object was last modified.",
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
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
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
			"trial_admin": {
				Description: "The administrative state of the trial license.\nWhen the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period,\ni.e. licenseState is set to be TrialPeriod.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceLicenseLicenseInfo().Schema},
				Computed: true,
			}},
	}
}

func dataSourceLicenseLicenseInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.LicenseLicenseInfo{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("active_admin"); ok {
		x := (v.(bool))
		o.SetActiveAdmin(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("days_left"); ok {
		x := int64(v.(int))
		o.SetDaysLeft(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
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
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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
		return diag.Errorf("json marshal of LicenseLicenseInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseLicenseInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of LicenseLicenseInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of LicenseLicenseInfo: %s", responseErr.Error())
	}
	count := countResponse.LicenseLicenseInfoList.GetCount()
	var i int32
	var licenseLicenseInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseLicenseInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching LicenseLicenseInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching LicenseLicenseInfo: %s", responseErr.Error())
		}
		results := resMo.LicenseLicenseInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for LicenseLicenseInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account_license_data"] = flattenMapLicenseAccountLicenseDataRelationship(s.GetAccountLicenseData(), d)
				temp["account_moid"] = (s.GetAccountMoid())
				temp["active_admin"] = (s.GetActiveAdmin())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["days_left"] = (s.GetDaysLeft())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["enforce_mode"] = (s.GetEnforceMode())
				temp["error_desc"] = (s.GetErrorDesc())
				temp["evaluation_period"] = (s.GetEvaluationPeriod())
				temp["extra_evaluation"] = (s.GetExtraEvaluation())
				temp["license_count"] = (s.GetLicenseCount())
				temp["license_state"] = (s.GetLicenseState())
				temp["license_type"] = (s.GetLicenseType())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["start_time"] = (s.GetStartTime()).String()

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["trial_admin"] = (s.GetTrialAdmin())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				licenseLicenseInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(licenseLicenseInfoResults))
	if err := d.Set("results", licenseLicenseInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(licenseLicenseInfoResults[0]["moid"].(string))
	return de
}
