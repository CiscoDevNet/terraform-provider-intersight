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

func dataSourceUcsdBackupInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUcsdBackupInfoRead,
		Schema: map[string]*schema.Schema{
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceUcsdBackupInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
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
		return diag.Errorf("json marshal of UcsdBackupInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.UcsdApi.GetUcsdBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of UcsdBackupInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.UcsdBackupInfoList.GetCount()
	var i int32
	var ucsdBackupInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.UcsdApi.GetUcsdBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching UcsdBackupInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.UcsdBackupInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for UcsdBackupInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["backup_file_name"] = (s.GetBackupFileName())
				temp["backup_location"] = (s.GetBackupLocation())
				temp["backup_server_ip"] = (s.GetBackupServerIp())
				temp["backup_size"] = (s.GetBackupSize())
				temp["class_id"] = (s.GetClassId())

				temp["connectors"] = flattenListUcsdConnectorPack(s.GetConnectors(), d)
				temp["duration"] = (s.GetDuration())
				temp["failure_reason"] = (s.GetFailureReason())
				temp["is_purged"] = (s.GetIsPurged())

				temp["last_modified"] = (s.GetLastModified()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["percentage_completion"] = (s.GetPercentageCompletion())
				temp["product_version"] = (s.GetProductVersion())
				temp["protocol"] = (s.GetProtocol())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["stage_completion"] = (s.GetStageCompletion())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				ucsdBackupInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(ucsdBackupInfoResults))
	if err := d.Set("results", ucsdBackupInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(ucsdBackupInfoResults[0]["moid"].(string))
	return de
}
