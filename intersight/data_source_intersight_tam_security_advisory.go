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

func dataSourceTamSecurityAdvisory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTamSecurityAdvisoryRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"advisory_id": {
				Description: "Cisco generated identifier for the published security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"base_score": {
				Description: "CVSS version 3 base score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
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
			"date_published": {
				Description: "Date when the security advisory was first published by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"date_updated": {
				Description: "Date when the security advisory was last updated by Cisco.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Brief description of the advisory details.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"environmental_score": {
				Description: "CVSS version 3 environmental score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"external_url": {
				Description: "A link to an external URL describing security Advisory in more details.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"name": {
				Description: "A user defined name for the Intersight Advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"recommendation": {
				Description: "Recommended action to resolve the security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "Current state of the advisory.\n* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.\n* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.\n* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.\n* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"temporal_score": {
				Description: "CVSS version 3 temporal score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"nr_version": {
				Description: "Cisco assigned advisory version after latest revision.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"workaround": {
				Description: "Workarounds available for the advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceTamSecurityAdvisory().Schema},
				Computed: true,
			}},
	}
}

func dataSourceTamSecurityAdvisoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TamSecurityAdvisory{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("advisory_id"); ok {
		x := (v.(string))
		o.SetAdvisoryId(x)
	}
	if v, ok := d.GetOk("base_score"); ok {
		x := v.(float32)
		o.SetBaseScore(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("date_published"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetDatePublished(x)
	}
	if v, ok := d.GetOk("date_updated"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetDateUpdated(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("environmental_score"); ok {
		x := v.(float32)
		o.SetEnvironmentalScore(x)
	}
	if v, ok := d.GetOk("external_url"); ok {
		x := (v.(string))
		o.SetExternalUrl(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
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
	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.SetRecommendation(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("temporal_score"); ok {
		x := v.(float32)
		o.SetTemporalScore(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}
	if v, ok := d.GetOk("workaround"); ok {
		x := (v.(string))
		o.SetWorkaround(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of TamSecurityAdvisory object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.TamApi.GetTamSecurityAdvisoryList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of TamSecurityAdvisory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of TamSecurityAdvisory: %s", responseErr.Error())
	}
	count := countResponse.TamSecurityAdvisoryList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for TamSecurityAdvisory data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var tamSecurityAdvisoryResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.TamApi.GetTamSecurityAdvisoryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching TamSecurityAdvisory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching TamSecurityAdvisory: %s", responseErr.Error())
		}
		results := resMo.TamSecurityAdvisoryList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())

				temp["actions"] = flattenListTamAction(s.GetActions(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["advisory_id"] = (s.GetAdvisoryId())

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["api_data_sources"] = flattenListTamApiDataSource(s.GetApiDataSources(), d)
				temp["base_score"] = (s.GetBaseScore())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["cve_ids"] = (s.GetCveIds())

				temp["date_published"] = (s.GetDatePublished()).String()

				temp["date_updated"] = (s.GetDateUpdated()).String()
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["environmental_score"] = (s.GetEnvironmentalScore())
				temp["external_url"] = (s.GetExternalUrl())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["recommendation"] = (s.GetRecommendation())

				temp["severity"] = flattenMapTamSeverity(s.GetSeverity(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["state"] = (s.GetState())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["temporal_score"] = (s.GetTemporalScore())
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["workaround"] = (s.GetWorkaround())
				tamSecurityAdvisoryResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(tamSecurityAdvisoryResults))
	if err := d.Set("results", tamSecurityAdvisoryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(tamSecurityAdvisoryResults[0]["moid"].(string))
	return de
}
