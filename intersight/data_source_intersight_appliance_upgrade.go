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

func dataSourceApplianceUpgrade() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceUpgradeRead,
		Schema: map[string]*schema.Schema{
			"account": {
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
		},
	}
}

func dataSourceApplianceUpgradeRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewApplianceUpgradeWithDefaults()
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.ApplianceApi.GetApplianceUpgradeList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.ApplianceUpgradeList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to ApplianceUpgrade: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewApplianceUpgradeWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Account: %+v", err)
			}
			if err := d.Set("active", (s.GetActive())); err != nil {
				return fmt.Errorf("error occurred while setting property Active: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("auto_created", (s.GetAutoCreated())); err != nil {
				return fmt.Errorf("error occurred while setting property AutoCreated: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("completed_phases", flattenListOnpremUpgradePhase(s.GetCompletedPhases(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CompletedPhases: %+v", err)
			}

			if err := d.Set("current_phase", flattenMapOnpremUpgradePhase(s.GetCurrentPhase(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CurrentPhase: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}
			if err := d.Set("elapsed_time", (s.GetElapsedTime())); err != nil {
				return fmt.Errorf("error occurred while setting property ElapsedTime: %+v", err)
			}

			if err := d.Set("end_time", (s.GetEndTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndTime: %+v", err)
			}
			if err := d.Set("error_code", (s.GetErrorCode())); err != nil {
				return fmt.Errorf("error occurred while setting property ErrorCode: %+v", err)
			}
			if err := d.Set("fingerprint", (s.GetFingerprint())); err != nil {
				return fmt.Errorf("error occurred while setting property Fingerprint: %+v", err)
			}

			if err := d.Set("image_bundle", flattenMapApplianceImageBundleRelationship(s.GetImageBundle(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property ImageBundle: %+v", err)
			}
			if err := d.Set("is_rolling_back", (s.GetIsRollingBack())); err != nil {
				return fmt.Errorf("error occurred while setting property IsRollingBack: %+v", err)
			}
			if err := d.Set("is_user_triggered", (s.GetIsUserTriggered())); err != nil {
				return fmt.Errorf("error occurred while setting property IsUserTriggered: %+v", err)
			}
			if err := d.Set("messages", (s.GetMessages())); err != nil {
				return fmt.Errorf("error occurred while setting property Messages: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("rollback_needed", (s.GetRollbackNeeded())); err != nil {
				return fmt.Errorf("error occurred while setting property RollbackNeeded: %+v", err)
			}

			if err := d.Set("rollback_phases", flattenListOnpremUpgradePhase(s.GetRollbackPhases(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RollbackPhases: %+v", err)
			}
			if err := d.Set("rollback_status", (s.GetRollbackStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property RollbackStatus: %+v", err)
			}
			if err := d.Set("services", (s.GetServices())); err != nil {
				return fmt.Errorf("error occurred while setting property Services: %+v", err)
			}

			if err := d.Set("start_time", (s.GetStartTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property StartTime: %+v", err)
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("total_phases", (s.GetTotalPhases())); err != nil {
				return fmt.Errorf("error occurred while setting property TotalPhases: %+v", err)
			}
			if err := d.Set("ui_packages", (s.GetUiPackages())); err != nil {
				return fmt.Errorf("error occurred while setting property UiPackages: %+v", err)
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property Version: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
