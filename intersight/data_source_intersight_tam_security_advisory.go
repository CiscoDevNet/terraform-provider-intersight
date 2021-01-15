package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTamSecurityAdvisory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTamSecurityAdvisoryRead,
		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"affected_object_type": {
							Description: "Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"alert_type": {
							Description: "Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).\n* `psirt` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).\n* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"identifiers": {
							Type:     schema.TypeList,
							Optional: true,
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
									"name": {
										Description: "Name of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"value": {
										Description: "Value of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"name": {
							Description: "Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operation_type": {
							Description: "Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.\n* `create` - Create an instance of AdvisoryInstance.\n* `remove` - Remove an instance of AdvisoryInstance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"queries": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"type": {
							Description: "Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.\n* `restApi` - Repesents the use of REST API for carrying out alert actions.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"advisory_id": {
				Description: "Cisco generated identifier for the published security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"api_data_sources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"mo_type": {
							Description: "Type of Intersight managed object used as data source.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Name is used to unique identify and refer a given data source in an alert definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"queries": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Description: "An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"query": {
										Description: "A SparkSQL query to be used on a given data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
							Computed: true,
						},
						"type": {
							Description: "Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).\n* `nxos` - Collector type for this data collection is NXOS.\n* `intersightApi` - Collector type for this data collection is Intersight APIs.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"base_score": {
				Description: "CVSS version 3 base score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cve_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"recommendation": {
				Description: "Recommended action to resolve the security advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"severity": {
				Description: "Severity level of the Intersight Advisory.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
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
		},
	}
}

func dataSourceTamSecurityAdvisoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TamSecurityAdvisory{}
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
	if v, ok := d.GetOk("environmental_score"); ok {
		x := v.(float32)
		o.SetEnvironmentalScore(x)
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
	resMo, _, responseErr := conn.ApiClient.TamApi.GetTamSecurityAdvisoryList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching TamSecurityAdvisory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for TamSecurityAdvisory list: %s", err.Error())
	}
	var s = &models.TamSecurityAdvisoryList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to TamSecurityAdvisory list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for TamSecurityAdvisory data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for TamSecurityAdvisory data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.TamSecurityAdvisory{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}

			if err := d.Set("actions", flattenListTamAction(s.GetActions(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Actions: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("advisory_id", (s.GetAdvisoryId())); err != nil {
				return diag.Errorf("error occurred while setting property AdvisoryId: %s", err.Error())
			}

			if err := d.Set("api_data_sources", flattenListTamApiDataSource(s.GetApiDataSources(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ApiDataSources: %s", err.Error())
			}
			if err := d.Set("base_score", (s.GetBaseScore())); err != nil {
				return diag.Errorf("error occurred while setting property BaseScore: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("cve_ids", (s.GetCveIds())); err != nil {
				return diag.Errorf("error occurred while setting property CveIds: %s", err.Error())
			}

			if err := d.Set("date_published", (s.GetDatePublished()).String()); err != nil {
				return diag.Errorf("error occurred while setting property DatePublished: %s", err.Error())
			}

			if err := d.Set("date_updated", (s.GetDateUpdated()).String()); err != nil {
				return diag.Errorf("error occurred while setting property DateUpdated: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("environmental_score", (s.GetEnvironmentalScore())); err != nil {
				return diag.Errorf("error occurred while setting property EnvironmentalScore: %s", err.Error())
			}
			if err := d.Set("external_url", (s.GetExternalUrl())); err != nil {
				return diag.Errorf("error occurred while setting property ExternalUrl: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}
			if err := d.Set("recommendation", (s.GetRecommendation())); err != nil {
				return diag.Errorf("error occurred while setting property Recommendation: %s", err.Error())
			}

			if err := d.Set("severity", flattenMapTamSeverity(s.GetSeverity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Severity: %s", err.Error())
			}
			if err := d.Set("state", (s.GetState())); err != nil {
				return diag.Errorf("error occurred while setting property State: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("temporal_score", (s.GetTemporalScore())); err != nil {
				return diag.Errorf("error occurred while setting property TemporalScore: %s", err.Error())
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return diag.Errorf("error occurred while setting property Version: %s", err.Error())
			}
			if err := d.Set("workaround", (s.GetWorkaround())); err != nil {
				return diag.Errorf("error occurred while setting property Workaround: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
