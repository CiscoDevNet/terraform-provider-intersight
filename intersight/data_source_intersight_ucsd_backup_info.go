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

func dataSourceUcsdBackupInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUcsdBackupInfoRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"backup_file_name": {
				Description: "Auto generated backup File Name with combination of file prefix given an user input and the timestamp.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_location": {
				Description: "Backup location that contains the backup images for end device which can be used for restore operation.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_server_ip": {
				Description: "Backup server where backup images are maintained.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_size": {
				Description: "Size of the backup image in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connectors": {
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
						"connector_feature": {
							Description: "State of the connector pack whether it is enabled or disabled.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"dependency_names": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"downloaded_version": {
							Description: "Version of the connector pack that is last downloaded successfully to UCS Director.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the connector pack running on the UCS Director.",
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
						"services": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"state": {
							Description: "State of the connector pack whether it is enabled or disabled.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Version of the connector pack.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"duration": {
				Description: "Time taken for the backup to be completed.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"encryption_key": {
				Description: "The key used for encrypting the backup file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"failure_reason": {
				Description: "Reason for backup failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_purged": {
				Description: "Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"last_modified": {
				Description: "Last modified time when this backup record got updated.",
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
			"percentage_completion": {
				Description: "Backup current precentage completion status information.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"product_version": {
				Description: "The end device product version when the backup image was taken.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"protocol": {
				Description: "Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"stage_completion": {
				Description: "Backup current status stage information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "Start time of backup when it got initiated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Current status of Backup current.",
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
		},
	}
}

func dataSourceUcsdBackupInfoRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.UcsdBackupInfo{}
	if v, ok := d.GetOk("backup_file_name"); ok {
		x := (v.(string))
		o.SetBackupFileName(x)
	}
	if v, ok := d.GetOk("backup_location"); ok {
		x := (v.(string))
		o.SetBackupLocation(x)
	}
	if v, ok := d.GetOk("backup_server_ip"); ok {
		x := (v.(string))
		o.SetBackupServerIp(x)
	}
	if v, ok := d.GetOk("backup_size"); ok {
		x := int64(v.(int))
		o.SetBackupSize(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("duration"); ok {
		x := int64(v.(int))
		o.SetDuration(x)
	}
	if v, ok := d.GetOk("encryption_key"); ok {
		x := (v.(string))
		o.SetEncryptionKey(x)
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.SetFailureReason(x)
	}
	if v, ok := d.GetOk("is_purged"); ok {
		x := (v.(bool))
		o.SetIsPurged(x)
	}
	if v, ok := d.GetOk("last_modified"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastModified(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("percentage_completion"); ok {
		x := int64(v.(int))
		o.SetPercentageCompletion(x)
	}
	if v, ok := d.GetOk("product_version"); ok {
		x := (v.(string))
		o.SetProductVersion(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("stage_completion"); ok {
		x := (v.(string))
		o.SetStageCompletion(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.UcsdApi.GetUcsdBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.UcsdBackupInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to UcsdBackupInfo: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.UcsdBackupInfo{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("backup_file_name", (s.GetBackupFileName())); err != nil {
				return fmt.Errorf("error occurred while setting property BackupFileName: %+v", err)
			}
			if err := d.Set("backup_location", (s.GetBackupLocation())); err != nil {
				return fmt.Errorf("error occurred while setting property BackupLocation: %+v", err)
			}
			if err := d.Set("backup_server_ip", (s.GetBackupServerIp())); err != nil {
				return fmt.Errorf("error occurred while setting property BackupServerIp: %+v", err)
			}
			if err := d.Set("backup_size", (s.GetBackupSize())); err != nil {
				return fmt.Errorf("error occurred while setting property BackupSize: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("connectors", flattenListUcsdConnectorPack(s.GetConnectors(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Connectors: %+v", err)
			}
			if err := d.Set("duration", (s.GetDuration())); err != nil {
				return fmt.Errorf("error occurred while setting property Duration: %+v", err)
			}
			if err := d.Set("failure_reason", (s.GetFailureReason())); err != nil {
				return fmt.Errorf("error occurred while setting property FailureReason: %+v", err)
			}
			if err := d.Set("is_purged", (s.GetIsPurged())); err != nil {
				return fmt.Errorf("error occurred while setting property IsPurged: %+v", err)
			}

			if err := d.Set("last_modified", (s.GetLastModified()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastModified: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("percentage_completion", (s.GetPercentageCompletion())); err != nil {
				return fmt.Errorf("error occurred while setting property PercentageCompletion: %+v", err)
			}
			if err := d.Set("product_version", (s.GetProductVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property ProductVersion: %+v", err)
			}
			if err := d.Set("protocol", (s.GetProtocol())); err != nil {
				return fmt.Errorf("error occurred while setting property Protocol: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("stage_completion", (s.GetStageCompletion())); err != nil {
				return fmt.Errorf("error occurred while setting property StageCompletion: %+v", err)
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
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
