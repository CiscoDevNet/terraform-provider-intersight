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

func resourceTamSecurityAdvisory() *schema.Resource {
	return &schema.Resource{
		Create: resourceTamSecurityAdvisoryCreate,
		Read:   resourceTamSecurityAdvisoryRead,
		Update: resourceTamSecurityAdvisoryUpdate,
		Delete: resourceTamSecurityAdvisoryDelete,
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
							Default:     "psirt",
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the filter paramter.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"name": {
							Description: "Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operation_type": {
							Description: "Operation type for the alert action. An action is used to carry out the process of \"reacting\" to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.\n* `create` - Create an instance of AdvisoryInstance.\n* `remove` - Remove an instance of AdvisoryInstance.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "create",
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
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"type": {
							Description: "Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.\n* `restApi` - Repesents the use of REST API for carrying out alert actions.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "restApi",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
										Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"type": {
							Description: "Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).\n* `nxos` - Collector type for this data collection is NXOS.\n* `intersightApi` - Collector type for this data collection is Intersight APIs.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "nxos",
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"base_score": {
				Description: "CVSS version 3 base score for the security Advisory.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				ForceNew:    true,
			},
			"name": {
				Description: "A user defined name for the Intersight Advisory.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
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
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"state": {
				Description: "Current state of the advisory.\n* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.\n* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ready",
			},
			"status": {
				Description: "Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.\n* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.\n* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "interim",
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

func resourceTamSecurityAdvisoryCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewTamSecurityAdvisoryWithDefaults()
	if v, ok := d.GetOk("actions"); ok {
		x := make([]models.TamAction, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewTamActionWithDefaults()
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
			if v, ok := l["affected_object_type"]; ok {
				{
					x := (v.(string))
					o.SetAffectedObjectType(x)
				}
			}
			if v, ok := l["alert_type"]; ok {
				{
					x := (v.(string))
					o.SetAlertType(x)
				}
			}
			o.SetClassId("tam.Action")
			if v, ok := l["identifiers"]; ok {
				{
					x := make([]models.TamIdentifiers, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamIdentifiersWithDefaults()
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
						o.SetClassId("tam.Identifiers")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
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
						o.SetIdentifiers(x)
					}
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["operation_type"]; ok {
				{
					x := (v.(string))
					o.SetOperationType(x)
				}
			}
			if v, ok := l["queries"]; ok {
				{
					x := make([]models.TamQueryEntry, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamQueryEntryWithDefaults()
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
						o.SetClassId("tam.QueryEntry")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["priority"]; ok {
							{
								x := int64(v.(int))
								o.SetPriority(x)
							}
						}
						if v, ok := l["query"]; ok {
							{
								x := (v.(string))
								o.SetQuery(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetQueries(x)
					}
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetActions(x)
		}
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("advisory_id"); ok {
		x := (v.(string))
		o.SetAdvisoryId(x)
	}

	if v, ok := d.GetOk("api_data_sources"); ok {
		x := make([]models.TamApiDataSource, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewTamApiDataSourceWithDefaults()
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
			o.SetClassId("tam.ApiDataSource")
			if v, ok := l["mo_type"]; ok {
				{
					x := (v.(string))
					o.SetMoType(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["queries"]; ok {
				{
					x := make([]models.TamQueryEntry, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamQueryEntryWithDefaults()
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
						o.SetClassId("tam.QueryEntry")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["priority"]; ok {
							{
								x := int64(v.(int))
								o.SetPriority(x)
							}
						}
						if v, ok := l["query"]; ok {
							{
								x := (v.(string))
								o.SetQuery(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetQueries(x)
					}
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetApiDataSources(x)
		}
	}

	if v, ok := d.GetOk("base_score"); ok {
		x := v.(float32)
		o.SetBaseScore(x)
	}

	o.SetClassId("tam.SecurityAdvisory")

	if v, ok := d.GetOk("cve_ids"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetCveIds(x)
		}
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

	o.SetObjectType("tam.SecurityAdvisory")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("recommendation"); ok {
		x := (v.(string))
		o.SetRecommendation(x)
	}

	if v, ok := d.GetOk("severity"); ok {
		p := make([]models.TamSeverity, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewTamSeverityWithDefaults()
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
			o.SetClassId("tam.Severity")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSeverity(x)
		}
	}

	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}

	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
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

	r := conn.ApiClient.TamApi.CreateTamSecurityAdvisory(conn.ctx).TamSecurityAdvisory(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceTamSecurityAdvisoryRead(d, meta)
}

func resourceTamSecurityAdvisoryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.TamApi.GetTamSecurityAdvisoryByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("actions", flattenListTamAction(s.GetActions(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Actions: %+v", err)
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("advisory_id", (s.GetAdvisoryId())); err != nil {
		return fmt.Errorf("error occurred while setting property AdvisoryId: %+v", err)
	}

	if err := d.Set("api_data_sources", flattenListTamApiDataSource(s.GetApiDataSources(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property ApiDataSources: %+v", err)
	}

	if err := d.Set("base_score", (s.GetBaseScore())); err != nil {
		return fmt.Errorf("error occurred while setting property BaseScore: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("cve_ids", (s.GetCveIds())); err != nil {
		return fmt.Errorf("error occurred while setting property CveIds: %+v", err)
	}

	if err := d.Set("date_published", (s.GetDatePublished()).String()); err != nil {
		return fmt.Errorf("error occurred while setting property DatePublished: %+v", err)
	}

	if err := d.Set("date_updated", (s.GetDateUpdated()).String()); err != nil {
		return fmt.Errorf("error occurred while setting property DateUpdated: %+v", err)
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return fmt.Errorf("error occurred while setting property Description: %+v", err)
	}

	if err := d.Set("environmental_score", (s.GetEnvironmentalScore())); err != nil {
		return fmt.Errorf("error occurred while setting property EnvironmentalScore: %+v", err)
	}

	if err := d.Set("external_url", (s.GetExternalUrl())); err != nil {
		return fmt.Errorf("error occurred while setting property ExternalUrl: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return fmt.Errorf("error occurred while setting property Name: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Organization: %+v", err)
	}

	if err := d.Set("recommendation", (s.GetRecommendation())); err != nil {
		return fmt.Errorf("error occurred while setting property Recommendation: %+v", err)
	}

	if err := d.Set("severity", flattenMapTamSeverity(s.GetSeverity(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Severity: %+v", err)
	}

	if err := d.Set("state", (s.GetState())); err != nil {
		return fmt.Errorf("error occurred while setting property State: %+v", err)
	}

	if err := d.Set("status", (s.GetStatus())); err != nil {
		return fmt.Errorf("error occurred while setting property Status: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	if err := d.Set("temporal_score", (s.GetTemporalScore())); err != nil {
		return fmt.Errorf("error occurred while setting property TemporalScore: %+v", err)
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return fmt.Errorf("error occurred while setting property Version: %+v", err)
	}

	if err := d.Set("workaround", (s.GetWorkaround())); err != nil {
		return fmt.Errorf("error occurred while setting property Workaround: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceTamSecurityAdvisoryUpdate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewTamSecurityAdvisoryWithDefaults()
	if d.HasChange("actions") {
		v := d.Get("actions")
		x := make([]models.TamAction, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewTamActionWithDefaults()
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
			if v, ok := l["affected_object_type"]; ok {
				{
					x := (v.(string))
					o.SetAffectedObjectType(x)
				}
			}
			if v, ok := l["alert_type"]; ok {
				{
					x := (v.(string))
					o.SetAlertType(x)
				}
			}
			o.SetClassId("tam.Action")
			if v, ok := l["identifiers"]; ok {
				{
					x := make([]models.TamIdentifiers, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamIdentifiersWithDefaults()
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
						o.SetClassId("tam.Identifiers")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
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
						o.SetIdentifiers(x)
					}
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["operation_type"]; ok {
				{
					x := (v.(string))
					o.SetOperationType(x)
				}
			}
			if v, ok := l["queries"]; ok {
				{
					x := make([]models.TamQueryEntry, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamQueryEntryWithDefaults()
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
						o.SetClassId("tam.QueryEntry")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["priority"]; ok {
							{
								x := int64(v.(int))
								o.SetPriority(x)
							}
						}
						if v, ok := l["query"]; ok {
							{
								x := (v.(string))
								o.SetQuery(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetQueries(x)
					}
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetActions(x)
		}
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

	if d.HasChange("advisory_id") {
		v := d.Get("advisory_id")
		x := (v.(string))
		o.SetAdvisoryId(x)
	}

	if d.HasChange("api_data_sources") {
		v := d.Get("api_data_sources")
		x := make([]models.TamApiDataSource, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewTamApiDataSourceWithDefaults()
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
			o.SetClassId("tam.ApiDataSource")
			if v, ok := l["mo_type"]; ok {
				{
					x := (v.(string))
					o.SetMoType(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["queries"]; ok {
				{
					x := make([]models.TamQueryEntry, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewTamQueryEntryWithDefaults()
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
						o.SetClassId("tam.QueryEntry")
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["priority"]; ok {
							{
								x := int64(v.(int))
								o.SetPriority(x)
							}
						}
						if v, ok := l["query"]; ok {
							{
								x := (v.(string))
								o.SetQuery(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetQueries(x)
					}
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetApiDataSources(x)
		}
	}

	if d.HasChange("base_score") {
		v := d.Get("base_score")
		x := v.(float32)
		o.SetBaseScore(x)
	}

	o.SetClassId("tam.SecurityAdvisory")

	if d.HasChange("cve_ids") {
		v := d.Get("cve_ids")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetCveIds(x)
		}
	}

	if d.HasChange("date_published") {
		v := d.Get("date_published")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetDatePublished(x)
	}

	if d.HasChange("date_updated") {
		v := d.Get("date_updated")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetDateUpdated(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("environmental_score") {
		v := d.Get("environmental_score")
		x := v.(float32)
		o.SetEnvironmentalScore(x)
	}

	if d.HasChange("external_url") {
		v := d.Get("external_url")
		x := (v.(string))
		o.SetExternalUrl(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("tam.SecurityAdvisory")

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if d.HasChange("recommendation") {
		v := d.Get("recommendation")
		x := (v.(string))
		o.SetRecommendation(x)
	}

	if d.HasChange("severity") {
		v := d.Get("severity")
		p := make([]models.TamSeverity, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewTamSeverityWithDefaults()
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
			o.SetClassId("tam.Severity")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSeverity(x)
		}
	}

	if d.HasChange("state") {
		v := d.Get("state")
		x := (v.(string))
		o.SetState(x)
	}

	if d.HasChange("status") {
		v := d.Get("status")
		x := (v.(string))
		o.SetStatus(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
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

	if d.HasChange("temporal_score") {
		v := d.Get("temporal_score")
		x := v.(float32)
		o.SetTemporalScore(x)
	}

	if d.HasChange("nr_version") {
		v := d.Get("nr_version")
		x := (v.(string))
		o.SetVersion(x)
	}

	if d.HasChange("workaround") {
		v := d.Get("workaround")
		x := (v.(string))
		o.SetWorkaround(x)
	}

	r := conn.ApiClient.TamApi.UpdateTamSecurityAdvisory(conn.ctx, d.Id()).TamSecurityAdvisory(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while updating: %s", err.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceTamSecurityAdvisoryRead(d, meta)
}

func resourceTamSecurityAdvisoryDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	p := conn.ApiClient.TamApi.DeleteTamSecurityAdvisory(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
