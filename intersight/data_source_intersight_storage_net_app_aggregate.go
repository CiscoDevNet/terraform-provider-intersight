package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageNetAppAggregate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageNetAppAggregateRead,
		Schema: map[string]*schema.Schema{
			"aggregate_type": {
				Description: "Storage disk type for NetApp aggregate.\n* `HDD` - Hard Disk Drive.\n* `Hybrid` - Solid State Hard Disk Drive.\n* `Hybrid (Flash Pool)` - SSHD in a flash pool.\n* `SSD` - Solid State Disk.\n* `SSD (FabricPool)` - SSD in a flash pool.\n* `VMDisk (SDS)` - Storage disk with Hard disk drive.\n* `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives.\n* `LUN (FlexArray)` - LUN as a disk.\n* `Not Mapped` - Storage disk is not mapped.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
			"raid_size": {
				Description: "Size of the RAID group represented by number of disks in the group.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"raid_type": {
				Description: "The RAID configuration type.\n* `Unknown` - Default unknown RAID type.\n* `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information.\n* `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.\n* `RAID4` - RAID4 stripes block level data and dedicates a disk to parity.\n* `RAID5` - RAID5  distributes striping and parity at a block level.\n* `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping.\n* `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.\n* `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.\n* `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Description: "Current state of the NetApp aggregate.\n* `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server.\n* `Online` - Aggregate is ready and available.\n* `Onlining` - Transitioning online.\n* `Offline` - Aggregate is unavailable.\n* `Offlining` - Transitioning offline.\n* `Relocating` - The aggregate is being relocated.\n* `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed.\n* `Failed` - The aggregate cannot be brought online.\n* `Inconsistent` - The aggregate has been marked corrupted; contact technical support.\n* `Unmounted` - The aggregate is not mounted.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Human readable status of the pool, indicating the current health.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Human readable type of the pool, such as thin, tiered, active-flash, etc.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "Uuid of  NetApp Aggregate.",
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
					"aggregate_type": {
						Description: "Storage disk type for NetApp aggregate.\n* `HDD` - Hard Disk Drive.\n* `Hybrid` - Solid State Hard Disk Drive.\n* `Hybrid (Flash Pool)` - SSHD in a flash pool.\n* `SSD` - Solid State Disk.\n* `SSD (FabricPool)` - SSD in a flash pool.\n* `VMDisk (SDS)` - Storage disk with Hard disk drive.\n* `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives.\n* `LUN (FlexArray)` - LUN as a disk.\n* `Not Mapped` - Storage disk is not mapped.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"array_controller": {
						Description: "A reference to a storageNetAppNode resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
					"raid_size": {
						Description: "Size of the RAID group represented by number of disks in the group.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"raid_type": {
						Description: "The RAID configuration type.\n* `Unknown` - Default unknown RAID type.\n* `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information.\n* `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.\n* `RAID4` - RAID4 stripes block level data and dedicates a disk to parity.\n* `RAID5` - RAID5  distributes striping and parity at a block level.\n* `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping.\n* `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.\n* `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.\n* `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"state": {
						Description: "Current state of the NetApp aggregate.\n* `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server.\n* `Online` - Aggregate is ready and available.\n* `Onlining` - Transitioning online.\n* `Offline` - Aggregate is unavailable.\n* `Offlining` - Transitioning offline.\n* `Relocating` - The aggregate is being relocated.\n* `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed.\n* `Failed` - The aggregate cannot be brought online.\n* `Inconsistent` - The aggregate has been marked corrupted; contact technical support.\n* `Unmounted` - The aggregate is not mounted.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"type": {
						Description: "Human readable type of the pool, such as thin, tiered, active-flash, etc.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"uuid": {
						Description: "Uuid of  NetApp Aggregate.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageNetAppAggregateRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageNetAppAggregate{}
	if v, ok := d.GetOk("aggregate_type"); ok {
		x := (v.(string))
		o.SetAggregateType(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("pool_id"); ok {
		x := (v.(string))
		o.SetPoolId(x)
	}
	if v, ok := d.GetOk("raid_size"); ok {
		x := int64(v.(int))
		o.SetRaidSize(x)
	}
	if v, ok := d.GetOk("raid_type"); ok {
		x := (v.(string))
		o.SetRaidType(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
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
		return diag.Errorf("json marshal of StorageNetAppAggregate object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppAggregateList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageNetAppAggregate: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageNetAppAggregateList.GetCount()
	var i int32
	var storageNetAppAggregateResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppAggregateList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageNetAppAggregate: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageNetAppAggregateList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageNetAppAggregate data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["aggregate_type"] = (s.GetAggregateType())

				temp["array_controller"] = flattenMapStorageNetAppNodeRelationship(s.GetArrayController(), d)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["pool_id"] = (s.GetPoolId())
				temp["raid_size"] = (s.GetRaidSize())
				temp["raid_type"] = (s.GetRaidType())
				temp["state"] = (s.GetState())
				temp["status"] = (s.GetStatus())

				temp["storage_utilization"] = flattenMapStorageBaseCapacity(s.GetStorageUtilization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["uuid"] = (s.GetUuid())
				storageNetAppAggregateResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageNetAppAggregateResults))
	if err := d.Set("results", storageNetAppAggregateResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageNetAppAggregateResults[0]["moid"].(string))
	return de
}
