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

func dataSourceApplianceUpgrade() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceUpgradeRead,
		Schema: map[string]*schema.Schema{
			"active": {
				Description: "Indicates if the software upgrade is active or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"auto_created": {
				Description: "Indicates that the request was automatically created by the system.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the software upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"elapsed_time": {
				Description: "Elapsed time in seconds during the software upgrade.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "End date of the software upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"error_code": {
				Description: "Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"fingerprint": {
				Description: "Software upgrade manifest's fingerprint.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_rolling_back": {
				Description: "Track if software upgrade is upgrading or rolling back.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"is_user_triggered": {
				Description: "Indicates if the upgrade is triggered by user or due to schedule.",
				Type:        schema.TypeBool,
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
			"rollback_needed": {
				Description: "Track if rollback is needed.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"rollback_status": {
				Description: "Status of the Intersight Appliance's software rollback status.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Status of the Intersight Appliance's software upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_phases": {
				Description: "TotalPhase represents the total number of the upgradePhases for one upgrade.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "Software upgrade manifest's version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account": {
					Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"active": {
						Description: "Indicates if the software upgrade is active or not.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"auto_created": {
						Description: "Indicates that the request was automatically created by the system.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"completed_phases": {
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
								"elapsed_time": {
									Description: "Elapsed time in seconds to complete the current upgrade phase.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"end_time": {
									Description: "End date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"failed": {
									Description: "Indicates if the upgrade phase has failed or not.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"message": {
									Description: "Status message set during the upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
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
								"start_time": {
									Description: "Start date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"current_phase": {
						Description: "Current phase of the Intersight Appliance's software upgrade.",
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
								"elapsed_time": {
									Description: "Elapsed time in seconds to complete the current upgrade phase.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"end_time": {
									Description: "End date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"failed": {
									Description: "Indicates if the upgrade phase has failed or not.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"message": {
									Description: "Status message set during the upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
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
								"start_time": {
									Description: "Start date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"description": {
						Description: "Description of the software upgrade.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"elapsed_time": {
						Description: "Elapsed time in seconds during the software upgrade.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"end_time": {
						Description: "End date of the software upgrade.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"error_code": {
						Description: "Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"fingerprint": {
						Description: "Software upgrade manifest's fingerprint.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"image_bundle": {
						Description: "A reference to a applianceImageBundle resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"is_rolling_back": {
						Description: "Track if software upgrade is upgrading or rolling back.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"is_user_triggered": {
						Description: "Indicates if the upgrade is triggered by user or due to schedule.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"messages": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
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
					"rollback_needed": {
						Description: "Track if rollback is needed.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"rollback_phases": {
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
								"elapsed_time": {
									Description: "Elapsed time in seconds to complete the current upgrade phase.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"end_time": {
									Description: "End date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"failed": {
									Description: "Indicates if the upgrade phase has failed or not.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"message": {
									Description: "Status message set during the upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the upgrade phase.\n* `init` - Upgrade service initialization phase.\n* `Prepare` - Upgrade service prepares folders and templated files.\n* `ServiceLoad` - Upgrade service loads the service images into the local docker cache.\n* `UiLoad` - Upgrade service loads the UI packages into the local cache.\n* `GenerateConfig` - Upgrade service generates the Kubernetes configuration files.\n* `DeployService` - Upgrade service deploys the Kubernetes services.\n* `Success` - Upgrade completed successfully.\n* `Fail` - Indicates that the upgrade process has failed.\n* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance.\n* `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.",
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
								"start_time": {
									Description: "Start date of the software upgrade phase.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"rollback_status": {
						Description: "Status of the Intersight Appliance's software rollback status.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"services": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"start_time": {
						Description: "Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of the Intersight Appliance's software upgrade.",
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
					"total_phases": {
						Description: "TotalPhase represents the total number of the upgradePhases for one upgrade.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"ui_packages": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"nr_version": {
						Description: "Software upgrade manifest's version.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceApplianceUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceUpgrade{}
	if v, ok := d.GetOk("active"); ok {
		x := (v.(bool))
		o.SetActive(x)
	}
	if v, ok := d.GetOk("auto_created"); ok {
		x := (v.(bool))
		o.SetAutoCreated(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("elapsed_time"); ok {
		x := int64(v.(int))
		o.SetElapsedTime(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("error_code"); ok {
		x := int64(v.(int))
		o.SetErrorCode(x)
	}
	if v, ok := d.GetOk("fingerprint"); ok {
		x := (v.(string))
		o.SetFingerprint(x)
	}
	if v, ok := d.GetOk("is_rolling_back"); ok {
		x := (v.(bool))
		o.SetIsRollingBack(x)
	}
	if v, ok := d.GetOk("is_user_triggered"); ok {
		x := (v.(bool))
		o.SetIsUserTriggered(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("rollback_needed"); ok {
		x := (v.(bool))
		o.SetRollbackNeeded(x)
	}
	if v, ok := d.GetOk("rollback_status"); ok {
		x := (v.(string))
		o.SetRollbackStatus(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("total_phases"); ok {
		x := int64(v.(int))
		o.SetTotalPhases(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceUpgradeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceUpgradeList.GetCount()
	var i int32
	var applianceUpgradeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceUpgradeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceUpgrade data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["active"] = (s.GetActive())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["auto_created"] = (s.GetAutoCreated())
				temp["class_id"] = (s.GetClassId())

				temp["completed_phases"] = flattenListOnpremUpgradePhase(s.GetCompletedPhases(), d)

				temp["current_phase"] = flattenMapOnpremUpgradePhase(s.GetCurrentPhase(), d)
				temp["description"] = (s.GetDescription())
				temp["elapsed_time"] = (s.GetElapsedTime())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["error_code"] = (s.GetErrorCode())
				temp["fingerprint"] = (s.GetFingerprint())

				temp["image_bundle"] = flattenMapApplianceImageBundleRelationship(s.GetImageBundle(), d)
				temp["is_rolling_back"] = (s.GetIsRollingBack())
				temp["is_user_triggered"] = (s.GetIsUserTriggered())
				temp["messages"] = (s.GetMessages())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["rollback_needed"] = (s.GetRollbackNeeded())

				temp["rollback_phases"] = flattenListOnpremUpgradePhase(s.GetRollbackPhases(), d)
				temp["rollback_status"] = (s.GetRollbackStatus())
				temp["services"] = (s.GetServices())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["total_phases"] = (s.GetTotalPhases())
				temp["ui_packages"] = (s.GetUiPackages())
				temp["nr_version"] = (s.GetVersion())
				applianceUpgradeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceUpgradeResults))
	if err := d.Set("results", applianceUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceUpgradeResults[0]["moid"].(string))
	return de
}
