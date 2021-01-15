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

func dataSourceStorageHitachiPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageHitachiPoolRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"array": {
				Description: "A reference to a storageHitachiArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"blocking_mode_blockade": {
				Description: "Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"blocking_mode_full": {
				Description: "Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"depletion_threshold": {
				Description: "The depletion threshold set for the pool (%).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_shrinking": {
				Description: "Whether the pool is shrinking is output.",
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
			"monitoring_mode": {
				Description: "Performance monitoring execution mode (monitor mode).\n* `N/A` - Not available.\n* `Period mode` - Period mode.\n* `Continuous mode` - Continuous mode.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"monitoring_status": {
				Description: "Status of monitor information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Human readable name of the pool, limited to 64 characters.",
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
			"pool_action_mode": {
				Description: "Execution mode for the pool.\n* `N/A` - Not available.\n* `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator).\n* `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pool_id": {
				Description: "Object ID for the pool. Platforms that use a number should convert it to string.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"progress_of_replacing": {
				Description: "Displays the status of the tier relocation processing.",
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
			"status": {
				Description: "Human readable status of the pool, indicating the current health.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_utilization": {
				Description: "Storage utilization of the pool entity in a storage array.",
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
						"available": {
							Description: "Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"capacity_utilization": {
							Description: "Percentage of used capacity.",
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"free": {
							Description: "Unused space available for applications to consume, represented in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"total": {
							Description: "Total storage capacity, represented in bytes. It is set by the component manufacturer.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"used": {
							Description: "Used or consumed storage capacity, represented in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"total_reserved_capacity": {
				Description: "Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Human readable type of the pool, such as thin, tiered, active-flash, etc.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"warning_threshold": {
				Description: "The warning threshold set for the pool (%).",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceStorageHitachiPoolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageHitachiPool{}
	if v, ok := d.GetOk("blocking_mode_blockade"); ok {
		x := (v.(string))
		o.SetBlockingModeBlockade(x)
	}
	if v, ok := d.GetOk("blocking_mode_full"); ok {
		x := (v.(string))
		o.SetBlockingModeFull(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("depletion_threshold"); ok {
		x := (v.(string))
		o.SetDepletionThreshold(x)
	}
	if v, ok := d.GetOk("is_shrinking"); ok {
		x := (v.(bool))
		o.SetIsShrinking(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("monitoring_mode"); ok {
		x := (v.(string))
		o.SetMonitoringMode(x)
	}
	if v, ok := d.GetOk("monitoring_status"); ok {
		x := (v.(string))
		o.SetMonitoringStatus(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pool_action_mode"); ok {
		x := (v.(string))
		o.SetPoolActionMode(x)
	}
	if v, ok := d.GetOk("pool_id"); ok {
		x := (v.(string))
		o.SetPoolId(x)
	}
	if v, ok := d.GetOk("progress_of_replacing"); ok {
		x := (v.(string))
		o.SetProgressOfReplacing(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("total_reserved_capacity"); ok {
		x := int64(v.(int))
		o.SetTotalReservedCapacity(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("warning_threshold"); ok {
		x := int64(v.(int))
		o.SetWarningThreshold(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageHitachiPool object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiPoolList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching StorageHitachiPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageHitachiPool list: %s", err.Error())
	}
	var s = &models.StorageHitachiPoolList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageHitachiPool list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageHitachiPool data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageHitachiPool data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageHitachiPool{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("array", flattenMapStorageHitachiArrayRelationship(s.GetArray(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Array: %s", err.Error())
			}
			if err := d.Set("blocking_mode_blockade", (s.GetBlockingModeBlockade())); err != nil {
				return diag.Errorf("error occurred while setting property BlockingModeBlockade: %s", err.Error())
			}
			if err := d.Set("blocking_mode_full", (s.GetBlockingModeFull())); err != nil {
				return diag.Errorf("error occurred while setting property BlockingModeFull: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("depletion_threshold", (s.GetDepletionThreshold())); err != nil {
				return diag.Errorf("error occurred while setting property DepletionThreshold: %s", err.Error())
			}
			if err := d.Set("is_shrinking", (s.GetIsShrinking())); err != nil {
				return diag.Errorf("error occurred while setting property IsShrinking: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("monitoring_mode", (s.GetMonitoringMode())); err != nil {
				return diag.Errorf("error occurred while setting property MonitoringMode: %s", err.Error())
			}
			if err := d.Set("monitoring_status", (s.GetMonitoringStatus())); err != nil {
				return diag.Errorf("error occurred while setting property MonitoringStatus: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("pool_action_mode", (s.GetPoolActionMode())); err != nil {
				return diag.Errorf("error occurred while setting property PoolActionMode: %s", err.Error())
			}
			if err := d.Set("pool_id", (s.GetPoolId())); err != nil {
				return diag.Errorf("error occurred while setting property PoolId: %s", err.Error())
			}
			if err := d.Set("progress_of_replacing", (s.GetProgressOfReplacing())); err != nil {
				return diag.Errorf("error occurred while setting property ProgressOfReplacing: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("storage_utilization", flattenMapStorageBaseCapacity(s.GetStorageUtilization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageUtilization: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("total_reserved_capacity", (s.GetTotalReservedCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property TotalReservedCapacity: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("warning_threshold", (s.GetWarningThreshold())); err != nil {
				return diag.Errorf("error occurred while setting property WarningThreshold: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
