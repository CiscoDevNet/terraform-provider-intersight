package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFirmwareUpgradeStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirmwareUpgradeStatusRead,
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
			"download_error": {
				Description: "The error message from the endpoint during the download.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_message": {
				Description: "The message from the endpoint during the download.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_percentage": {
				Description: "The percentage of the image downloaded in the endpoint.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_progress": {
				Description: "The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_retries": {
				Description: "The number of retries the plugin attempted before succeeding or failing the download.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_stage": {
				Description: "The image download stages. Example:downloading, flashing.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ep_power_status": {
				Description: "The server power status after the upgrade request is submitted in the endpoint.\n* `none` - Server power status is none.\n* `powered on` - Server power status is powered on.\n* `powered off` - Server power status is powered off.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"overall_error": {
				Description: "The reason for the operation failure.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"overall_percentage": {
				Description: "The overall percentage of the operation.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"overallstatus": {
				Description: "The overall status of the operation.\n* `none` - Upgrade stage is no upgrade stage.\n* `started` - Upgrade stage is started.\n* `prepare initiating` - Upgrade configuration is being prepared.\n* `prepare initiated` - Upgrade configuration is initiated.\n* `prepared` - Upgrade configuration is prepared.\n* `download initiating` - Upgrade stage is download initiating.\n* `download initiated` - Upgrade stage is download initiated.\n* `downloading` - Upgrade stage is downloading.\n* `downloaded` - Upgrade stage is downloaded.\n* `upgrade initiating on fabric A` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint.\n* `upgrade initiated on fabric A` - Upgrade stage is in upgrade initiated when the upgrade has started in endpoint.\n* `upgrading fabric A` - Upgrade stage is in upgrading when the upgrade requires reboot to complete.\n* `rebooting fabric A` - Upgrade is in rebooting when the endpoint is being rebooted.\n* `upgraded fabric A` - Upgrade stage is in upgraded when the corresponding endpoint has completed.\n* `upgrade initiating on fabric B` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint.\n* `upgrade initiated on fabric B` - Upgrade stage is in upgrade initiated when upgrade has started in endpoint.\n* `upgrading fabric B` - Upgrade stage is in upgrading when the upgrade requires reboot to complete.\n* `rebooting fabric B` - Upgrade is in rebooting when the endpoint is being rebooted.\n* `upgraded fabric B` - Upgrade stage is in upgraded when the corresponding endpoint has completed.\n* `upgrade initiating` - Upgrade stage is upgrade initiating.\n* `upgrade initiated` - Upgrade stage is upgrade initiated.\n* `upgrading` - Upgrade stage is upgrading.\n* `oob images staging` - Out-of-band component images staging.\n* `oob images staged` - Out-of-band component images staged.\n* `rebooting` - Upgrade is rebooting the endpoint.\n* `upgraded` - Upgrade stage is upgraded.\n* `success` - Upgrade stage is success.\n* `failed` - Upgrade stage is upgrade failed.\n* `terminated` - Upgrade stage is terminated.\n* `pending` - Upgrade stage is pending.\n* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.\n* `Caching` - The image will be cached into Intersight Appliance or an endpoint cache.\n* `Cached` - The image has been cached into the Intersight Appliance or endpoint cache.\n* `CachingFailed` - The image caching into the Intersight Appliance failed or endpoint cache.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pending_type": {
				Description: "Pending reason for the upgrade waiting.\n* `none` - Upgrade pending reason is none.\n* `pending for next reboot` - Upgrade pending reason is pending for next reboot.",
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
			"upgrade": {
				Description: "A reference to a firmwareUpgradeBase resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"workflow": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		},
	}
}

func dataSourceFirmwareUpgradeStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.FirmwareUpgradeStatus{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("download_error"); ok {
		x := (v.(string))
		o.SetDownloadError(x)
	}
	if v, ok := d.GetOk("download_message"); ok {
		x := (v.(string))
		o.SetDownloadMessage(x)
	}
	if v, ok := d.GetOk("download_percentage"); ok {
		x := int64(v.(int))
		o.SetDownloadPercentage(x)
	}
	if v, ok := d.GetOk("download_progress"); ok {
		x := int64(v.(int))
		o.SetDownloadProgress(x)
	}
	if v, ok := d.GetOk("download_retries"); ok {
		x := int64(v.(int))
		o.SetDownloadRetries(x)
	}
	if v, ok := d.GetOk("download_stage"); ok {
		x := (v.(string))
		o.SetDownloadStage(x)
	}
	if v, ok := d.GetOk("ep_power_status"); ok {
		x := (v.(string))
		o.SetEpPowerStatus(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("overall_error"); ok {
		x := (v.(string))
		o.SetOverallError(x)
	}
	if v, ok := d.GetOk("overall_percentage"); ok {
		x := int64(v.(int))
		o.SetOverallPercentage(x)
	}
	if v, ok := d.GetOk("overallstatus"); ok {
		x := (v.(string))
		o.SetOverallstatus(x)
	}
	if v, ok := d.GetOk("pending_type"); ok {
		x := (v.(string))
		o.SetPendingType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.FirmwareApi.GetFirmwareUpgradeStatusList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.FirmwareUpgradeStatusList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to FirmwareUpgradeStatus: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.FirmwareUpgradeStatus{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("download_error", (s.GetDownloadError())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadError: %+v", err)
			}
			if err := d.Set("download_message", (s.GetDownloadMessage())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadMessage: %+v", err)
			}
			if err := d.Set("download_percentage", (s.GetDownloadPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadPercentage: %+v", err)
			}
			if err := d.Set("download_progress", (s.GetDownloadProgress())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadProgress: %+v", err)
			}
			if err := d.Set("download_retries", (s.GetDownloadRetries())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadRetries: %+v", err)
			}
			if err := d.Set("download_stage", (s.GetDownloadStage())); err != nil {
				return fmt.Errorf("error occurred while setting property DownloadStage: %+v", err)
			}
			if err := d.Set("ep_power_status", (s.GetEpPowerStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property EpPowerStatus: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("overall_error", (s.GetOverallError())); err != nil {
				return fmt.Errorf("error occurred while setting property OverallError: %+v", err)
			}
			if err := d.Set("overall_percentage", (s.GetOverallPercentage())); err != nil {
				return fmt.Errorf("error occurred while setting property OverallPercentage: %+v", err)
			}
			if err := d.Set("overallstatus", (s.GetOverallstatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Overallstatus: %+v", err)
			}
			if err := d.Set("pending_type", (s.GetPendingType())); err != nil {
				return fmt.Errorf("error occurred while setting property PendingType: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}

			if err := d.Set("upgrade", flattenMapFirmwareUpgradeBaseRelationship(s.GetUpgrade(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Upgrade: %+v", err)
			}

			if err := d.Set("workflow", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflow(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Workflow: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
