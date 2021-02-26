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

func dataSourceStorageHitachiVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageHitachiVolumeRead,
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
			"attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"clpr_id": {
				Description: "CLPR (Cache Logical Partition) number of this volume.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"data_reduction_mode": {
				Description: "Setting of the capacity saving function (dedupe and compression).\n* `N/A` - Not available.\n* `Compression` - The capacity saving function (compression) is enabled.\n* `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled.\n* `Disabled` - The capacity saving function (compression and deduplication) is disabled.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"data_reduction_status": {
				Description: "Status of the capacity saving function.\n* `N/A` - Not available.\n* `Enabled` - The capacity saving function is enabled.\n* `Disabled` - The capacity saving function is disabled.\n* `Enabling` - The capacity saving function is being enabled.\n* `Rehydrating` - The capacity saving function is being disabled.\n* `Deleting` - The volumes for which the capacity saving function is enabled are being deleted.\n* `Failed` - An attempt to enable the capacity saving function failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Short description about the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_type": {
				Description: "Code indicating the drive type of the drive belonging to the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"emulation_type": {
				Description: "The volume emulation type or the volume status information.\n* `N/A` - Not available.\n* `NOT DEFINED` - The volume is not implemented.\n* `DEFINING` - The volume is being created.\n* `REMOVING` - The volume is being removed.\n* `OPEN-V` - To be provided by Hitachi.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_full_allocation_enabled": {
				Description: "Whether pages are reserved by the FullAllocation functionality.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"label": {
				Description: "Label of the volume, as configured in the storage array.",
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
			"naa_id": {
				Description: "NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Named entity of the volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_of_paths": {
				Description: "Number of paths set for the volume.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parity_group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"parity_groups": {
				Description: "An array of relationships to storageHitachiParityGroup resources.",
				Type:        schema.TypeList,
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
			"pool": {
				Description: "A reference to a storageHitachiPool resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"pool_id": {
				Description: "ID of the pool with which the volume is associated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raid_level": {
				Description: "RAID level for the volume.\n* `N/A` - Not available.\n* `RAID1` - RAID1.\n* `RAID5` - RAID5.\n* `RAID6` - RAID6.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raid_type": {
				Description: "RAID type drive configuration.",
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
			"size": {
				Description: "User provisioned volume size. It is the size exposed to host.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status information of the volume.\n* `N/A` - Not available.\n* `NML` - The volume is in normal status.\n* `BLK` - The volume is blocked.\n* `BSY` - The volume status is being changed.\n* `Unknown` - The volume status is unknown (not supported).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_utilization": {
				Description: "Storage utilization of volume entity in storage array.",
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
		},
	}
}

func dataSourceStorageHitachiVolumeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageHitachiVolume{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("clpr_id"); ok {
		x := int64(v.(int))
		o.SetClprId(x)
	}
	if v, ok := d.GetOk("data_reduction_mode"); ok {
		x := (v.(string))
		o.SetDataReductionMode(x)
	}
	if v, ok := d.GetOk("data_reduction_status"); ok {
		x := (v.(string))
		o.SetDataReductionStatus(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("drive_type"); ok {
		x := (v.(string))
		o.SetDriveType(x)
	}
	if v, ok := d.GetOk("emulation_type"); ok {
		x := (v.(string))
		o.SetEmulationType(x)
	}
	if v, ok := d.GetOk("is_full_allocation_enabled"); ok {
		x := (v.(bool))
		o.SetIsFullAllocationEnabled(x)
	}
	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("naa_id"); ok {
		x := (v.(string))
		o.SetNaaId(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("num_of_paths"); ok {
		x := int64(v.(int))
		o.SetNumOfPaths(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pool_id"); ok {
		x := (v.(string))
		o.SetPoolId(x)
	}
	if v, ok := d.GetOk("raid_level"); ok {
		x := (v.(string))
		o.SetRaidLevel(x)
	}
	if v, ok := d.GetOk("raid_type"); ok {
		x := (v.(string))
		o.SetRaidType(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageHitachiVolume object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiVolumeList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching StorageHitachiVolume: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageHitachiVolume list: %s", err.Error())
	}
	var s = &models.StorageHitachiVolumeList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageHitachiVolume list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageHitachiVolume data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageHitachiVolume data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageHitachiVolume{}
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
			if err := d.Set("attributes", (s.GetAttributes())); err != nil {
				return diag.Errorf("error occurred while setting property Attributes: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("clpr_id", (s.GetClprId())); err != nil {
				return diag.Errorf("error occurred while setting property ClprId: %s", err.Error())
			}
			if err := d.Set("data_reduction_mode", (s.GetDataReductionMode())); err != nil {
				return diag.Errorf("error occurred while setting property DataReductionMode: %s", err.Error())
			}
			if err := d.Set("data_reduction_status", (s.GetDataReductionStatus())); err != nil {
				return diag.Errorf("error occurred while setting property DataReductionStatus: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("drive_type", (s.GetDriveType())); err != nil {
				return diag.Errorf("error occurred while setting property DriveType: %s", err.Error())
			}
			if err := d.Set("emulation_type", (s.GetEmulationType())); err != nil {
				return diag.Errorf("error occurred while setting property EmulationType: %s", err.Error())
			}
			if err := d.Set("is_full_allocation_enabled", (s.GetIsFullAllocationEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property IsFullAllocationEnabled: %s", err.Error())
			}
			if err := d.Set("label", (s.GetLabel())); err != nil {
				return diag.Errorf("error occurred while setting property Label: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("naa_id", (s.GetNaaId())); err != nil {
				return diag.Errorf("error occurred while setting property NaaId: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("num_of_paths", (s.GetNumOfPaths())); err != nil {
				return diag.Errorf("error occurred while setting property NumOfPaths: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("parity_group_ids", (s.GetParityGroupIds())); err != nil {
				return diag.Errorf("error occurred while setting property ParityGroupIds: %s", err.Error())
			}

			if err := d.Set("parity_groups", flattenListStorageHitachiParityGroupRelationship(s.GetParityGroups(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ParityGroups: %s", err.Error())
			}

			if err := d.Set("pool", flattenMapStorageHitachiPoolRelationship(s.GetPool(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Pool: %s", err.Error())
			}
			if err := d.Set("pool_id", (s.GetPoolId())); err != nil {
				return diag.Errorf("error occurred while setting property PoolId: %s", err.Error())
			}
			if err := d.Set("raid_level", (s.GetRaidLevel())); err != nil {
				return diag.Errorf("error occurred while setting property RaidLevel: %s", err.Error())
			}
			if err := d.Set("raid_type", (s.GetRaidType())); err != nil {
				return diag.Errorf("error occurred while setting property RaidType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return diag.Errorf("error occurred while setting property Size: %s", err.Error())
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
			d.SetId(s.GetMoid())
		}
	}
	return de
}
