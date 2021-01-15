package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStoragePureReplicationSchedule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStoragePureReplicationScheduleRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"array": {
				Description: "A reference to a storagePureArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"daily_limit": {
				Description: "Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage\nFlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired.\nIn case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"frequency": {
				Description: "Replication frequency. It is an interval at which replication is set to trigger.\nExamples:\n    PT2H, Snapshot is generated every 2 hours.\n    P30D, Snapshot is scheduled for every 30 days.\n    PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
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
				Description: "Replication schedule name.",
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
			"protection_group": {
				Description: "A reference to a storagePureProtectionGroup resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"replication_blackout_intervals": {
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
						"end": {
							Description: "The end time of day for replication blackout window.\nExample: 17:00:01 which is 17 hours, 0 minutes, 1 seconds.",
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
						"start": {
							Description: "The start time of day when replication blackout is active. When replication blackout is active, the storage array temporarily disables replication.\nExample: 15:04:03.123 which is 15 hours, 4 minutes, 3 seconds and 123 milliseconds.\nThe fractional seconds are written using the standard decimal notation which can be used for setting milliseconds and microseconds.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"replication_time": {
				Description: "Preferred time of day at which to replicate the snapshots on target array.\nIt is applicable only if the replication frequency is set for a day or more.\nFormat: hh:mm:ss\nExample: 15:00:00, Replication is set for 3:00 PM.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"retention_time": {
				Description: "Duration to keep the replicated snapshots on the targets.\nReplicated snapshots are deleted from target array once the retention period has elapsed.\nExamples:\nP30D, Snapshots are available for 30 days.\nPT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_expiry_time": {
				Description: "Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.",
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

func dataSourceStoragePureReplicationScheduleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StoragePureReplicationSchedule{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("daily_limit"); ok {
		x := int64(v.(int))
		o.SetDailyLimit(x)
	}
	if v, ok := d.GetOk("frequency"); ok {
		x := (v.(string))
		o.SetFrequency(x)
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
	if v, ok := d.GetOk("replication_time"); ok {
		x := (v.(string))
		o.SetReplicationTime(x)
	}
	if v, ok := d.GetOk("retention_time"); ok {
		x := (v.(string))
		o.SetRetentionTime(x)
	}
	if v, ok := d.GetOk("snapshot_expiry_time"); ok {
		x := (v.(string))
		o.SetSnapshotExpiryTime(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StoragePureReplicationSchedule object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStoragePureReplicationScheduleList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching StoragePureReplicationSchedule: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StoragePureReplicationSchedule list: %s", err.Error())
	}
	var s = &models.StoragePureReplicationScheduleList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StoragePureReplicationSchedule list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StoragePureReplicationSchedule data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StoragePureReplicationSchedule data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StoragePureReplicationSchedule{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("array", flattenMapStoragePureArrayRelationship(s.GetArray(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Array: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("daily_limit", (s.GetDailyLimit())); err != nil {
				return diag.Errorf("error occurred while setting property DailyLimit: %s", err.Error())
			}
			if err := d.Set("frequency", (s.GetFrequency())); err != nil {
				return diag.Errorf("error occurred while setting property Frequency: %s", err.Error())
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

			if err := d.Set("protection_group", flattenMapStoragePureProtectionGroupRelationship(s.GetProtectionGroup(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ProtectionGroup: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("replication_blackout_intervals", flattenListStoragePureReplicationBlackout(s.GetReplicationBlackoutIntervals(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ReplicationBlackoutIntervals: %s", err.Error())
			}
			if err := d.Set("replication_time", (s.GetReplicationTime())); err != nil {
				return diag.Errorf("error occurred while setting property ReplicationTime: %s", err.Error())
			}
			if err := d.Set("retention_time", (s.GetRetentionTime())); err != nil {
				return diag.Errorf("error occurred while setting property RetentionTime: %s", err.Error())
			}
			if err := d.Set("snapshot_expiry_time", (s.GetSnapshotExpiryTime())); err != nil {
				return diag.Errorf("error occurred while setting property SnapshotExpiryTime: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
