package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTamAdvisoryDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTamAdvisoryDefinitionRead,
		Schema: map[string]*schema.Schema{
			"advisory_id": {
				Description: "Cisco generated identifier for the published security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"external_url": {
				Description: "A link to an external URL describing security Advisory in more details.",
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
			"state": {
				Description: "Current state of the advisory.\n* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.\n* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "The type (field notice, security advisory etc.) of Intersight advisory.\n* `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).\n* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).",
				Type:        schema.TypeString,
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
				Elem:     &schema.Resource{Schema: resourceTamAdvisoryDefinition().Schema},
				Computed: true,
			}},
	}
}

func dataSourceTamAdvisoryDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TamAdvisoryDefinition{}
	if v, ok := d.GetOk("advisory_id"); ok {
		x := (v.(string))
		o.SetAdvisoryId(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("external_url"); ok {
		x := (v.(string))
		o.SetExternalUrl(x)
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
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
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
		return diag.Errorf("json marshal of TamAdvisoryDefinition object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.TamApi.GetTamAdvisoryDefinitionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of TamAdvisoryDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.TamAdvisoryDefinitionList.GetCount()
	var i int32
	var tamAdvisoryDefinitionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.TamApi.GetTamAdvisoryDefinitionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching TamAdvisoryDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.TamAdvisoryDefinitionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for TamAdvisoryDefinition data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["actions"] = flattenListTamAction(s.GetActions(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["advisory_details"] = flattenMapTamBaseAdvisoryDetails(s.GetAdvisoryDetails(), d)
				temp["advisory_id"] = (s.GetAdvisoryId())

				temp["api_data_sources"] = flattenListTamApiDataSource(s.GetApiDataSources(), d)
				temp["class_id"] = (s.GetClassId())

				temp["date_published"] = (s.GetDatePublished()).String()

				temp["date_updated"] = (s.GetDateUpdated()).String()
				temp["description"] = (s.GetDescription())
				temp["external_url"] = (s.GetExternalUrl())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["recommendation"] = (s.GetRecommendation())

				temp["severity"] = flattenMapTamSeverity(s.GetSeverity(), d)
				temp["state"] = (s.GetState())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["nr_version"] = (s.GetVersion())
				temp["workaround"] = (s.GetWorkaround())
				tamAdvisoryDefinitionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(tamAdvisoryDefinitionResults))
	if err := d.Set("results", tamAdvisoryDefinitionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(tamAdvisoryDefinitionResults[0]["moid"].(string))
	return de
}
