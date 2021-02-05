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

func dataSourceStorageNetAppVolume() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageNetAppVolumeRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"array": {
				Description: "A reference to a storageNetAppCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"autosize_mode": {
				Description: "The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink.\n* `off` - The volume will not grow or shrink in size in response to the amount of used space.\n* `grow` - The volume will automatically grow when used space in the volume is above the grow threshold.\n* `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"created_time": {
				Description: "Storage container's creation time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"disk_pool": {
				Description: "An array of relationships to storageNetAppAggregate resources.",
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
			"export_policy_name": {
				Description: "Name of Export Policy.",
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
				Description: "Name of the storage container.",
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
			"snapshot_policy_name": {
				Description: "Name of Snapshot Policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_policy_uuid": {
				Description: "Name of Snapshot Policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_utilized_capacity": {
				Description: "The total space used by Snapshot copies in the volume represented in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "The current state of a NetApp volume.\n* `offline` - Read and write access to the volume is not allowed.\n* `online` - Read and write access to the volume is allowed.\n* `error` - Storage volume state of error type.\n* `mixed` - The constituents of a FlexGroup volume are not all in the same state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_utilization": {
				Description: "Storage utilization information of storage container.",
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
			"tenant": {
				Description: "A reference to a storageNetAppStorageVm resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"type": {
				Description: "NetApp volume type. The volume type can be Read-write or Data-protection, Load-sharing, or Data-cache.\n* `data-protection` - Prevents modification of the data on the Volume.\n* `read-write` - Data on the Volume can be modified.\n* `load-sharing` - Load Sharing.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "UUID of NetApp Volume.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceStorageNetAppVolumeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageNetAppVolume{}
	if v, ok := d.GetOk("autosize_mode"); ok {
		x := (v.(string))
		o.SetAutosizeMode(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("created_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreatedTime(x)
	}
	if v, ok := d.GetOk("export_policy_name"); ok {
		x := (v.(string))
		o.SetExportPolicyName(x)
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
	if v, ok := d.GetOk("snapshot_policy_name"); ok {
		x := (v.(string))
		o.SetSnapshotPolicyName(x)
	}
	if v, ok := d.GetOk("snapshot_policy_uuid"); ok {
		x := (v.(string))
		o.SetSnapshotPolicyUuid(x)
	}
	if v, ok := d.GetOk("snapshot_utilized_capacity"); ok {
		x := int64(v.(int))
		o.SetSnapshotUtilizedCapacity(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageNetAppVolume object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppVolumeList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching StorageNetAppVolume: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageNetAppVolume list: %s", err.Error())
	}
	var s = &models.StorageNetAppVolumeList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageNetAppVolume list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageNetAppVolume data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageNetAppVolume data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageNetAppVolume{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("array", flattenMapStorageNetAppClusterRelationship(s.GetArray(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Array: %s", err.Error())
			}
			if err := d.Set("autosize_mode", (s.GetAutosizeMode())); err != nil {
				return diag.Errorf("error occurred while setting property AutosizeMode: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("created_time", (s.GetCreatedTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property CreatedTime: %s", err.Error())
			}

			if err := d.Set("disk_pool", flattenListStorageNetAppAggregateRelationship(s.GetDiskPool(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DiskPool: %s", err.Error())
			}
			if err := d.Set("export_policy_name", (s.GetExportPolicyName())); err != nil {
				return diag.Errorf("error occurred while setting property ExportPolicyName: %s", err.Error())
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
			if err := d.Set("snapshot_policy_name", (s.GetSnapshotPolicyName())); err != nil {
				return diag.Errorf("error occurred while setting property SnapshotPolicyName: %s", err.Error())
			}
			if err := d.Set("snapshot_policy_uuid", (s.GetSnapshotPolicyUuid())); err != nil {
				return diag.Errorf("error occurred while setting property SnapshotPolicyUuid: %s", err.Error())
			}
			if err := d.Set("snapshot_utilized_capacity", (s.GetSnapshotUtilizedCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property SnapshotUtilizedCapacity: %s", err.Error())
			}
			if err := d.Set("state", (s.GetState())); err != nil {
				return diag.Errorf("error occurred while setting property State: %s", err.Error())
			}

			if err := d.Set("storage_utilization", flattenMapStorageBaseCapacity(s.GetStorageUtilization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageUtilization: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("tenant", flattenMapStorageNetAppStorageVmRelationship(s.GetTenant(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tenant: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
