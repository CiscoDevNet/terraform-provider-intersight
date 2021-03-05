package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageFlexUtilPhysicalDrive() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageFlexUtilPhysicalDriveRead,
		Schema: map[string]*schema.Schema{
			"block_size": {
				Description: "Block size of the FlexUtil Physical drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"capacity": {
				Description: "Capacity of the FlexUtil Physical drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller": {
				Description: "Type of the Physical Drive Controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drives_enabled": {
				Description: "The number of drives enabled in the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"health": {
				Description: "Health of the FlexUtil Physical drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"manufacturer_date": {
				Description: "Manufacturing date of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"manufacturer_id": {
				Description: "Manufacturer identity of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "This field identifies the model of the given component.",
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
			"oem_id": {
				Description: "The OEM Identifier of the FlexUtil physical drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partition_count": {
				Description: "The number of partitions present on the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pd_status": {
				Description: "Status of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"physical_drive": {
				Description: "The type of physical drive. Example - microSD.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_name": {
				Description: "Product name of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_revision": {
				Description: "Product revision of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"read_error_count": {
				Description: "Read error count of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"read_error_threshold": {
				Description: "Read error threshold for FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"revision": {
				Description: "This field identifies the revision of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"write_enabled": {
				Description: "Write access state of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"write_error_count": {
				Description: "Write error count of the FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"write_error_threshold": {
				Description: "Write error threshold for FlexUtil Physical Drive.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"block_size": {
						Description: "Block size of the FlexUtil Physical drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"capacity": {
						Description: "Capacity of the FlexUtil Physical drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"controller": {
						Description: "Type of the Physical Drive Controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"device_mo_id": {
						Description: "The database identifier of the registered device of an object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"dn": {
						Description: "The Distinguished Name unambiguously identifies an object in the system.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drives_enabled": {
						Description: "The number of drives enabled in the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"health": {
						Description: "Health of the FlexUtil Physical drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"inventory_device_info": {
						Description: "A reference to a inventoryDeviceInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"manufacturer_date": {
						Description: "Manufacturing date of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"manufacturer_id": {
						Description: "Manufacturer identity of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"model": {
						Description: "This field identifies the model of the given component.",
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
					"oem_id": {
						Description: "The OEM Identifier of the FlexUtil physical drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"partition_count": {
						Description: "The number of partitions present on the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"pd_status": {
						Description: "Status of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"physical_drive": {
						Description: "The type of physical drive. Example - microSD.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"product_name": {
						Description: "Product name of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"product_revision": {
						Description: "Product revision of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"read_error_count": {
						Description: "Read error count of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"read_error_threshold": {
						Description: "Read error threshold for FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"revision": {
						Description: "This field identifies the revision of the given component.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"rn": {
						Description: "The Relative Name uniquely identifies an object within a given context.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"serial": {
						Description: "This field identifies the serial of the given component.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"storage_flex_util_controller": {
						Description: "A reference to a storageFlexUtilController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"vendor": {
						Description: "This field identifies the vendor of the given component.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"write_enabled": {
						Description: "Write access state of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"write_error_count": {
						Description: "Write error count of the FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"write_error_threshold": {
						Description: "Write error threshold for FlexUtil Physical Drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageFlexUtilPhysicalDriveRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageFlexUtilPhysicalDrive{}
	if v, ok := d.GetOk("block_size"); ok {
		x := (v.(string))
		o.SetBlockSize(x)
	}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.SetCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("controller"); ok {
		x := (v.(string))
		o.SetController(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("drives_enabled"); ok {
		x := (v.(string))
		o.SetDrivesEnabled(x)
	}
	if v, ok := d.GetOk("health"); ok {
		x := (v.(string))
		o.SetHealth(x)
	}
	if v, ok := d.GetOk("manufacturer_date"); ok {
		x := (v.(string))
		o.SetManufacturerDate(x)
	}
	if v, ok := d.GetOk("manufacturer_id"); ok {
		x := (v.(string))
		o.SetManufacturerId(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oem_id"); ok {
		x := (v.(string))
		o.SetOemId(x)
	}
	if v, ok := d.GetOk("partition_count"); ok {
		x := (v.(string))
		o.SetPartitionCount(x)
	}
	if v, ok := d.GetOk("pd_status"); ok {
		x := (v.(string))
		o.SetPdStatus(x)
	}
	if v, ok := d.GetOk("physical_drive"); ok {
		x := (v.(string))
		o.SetPhysicalDrive(x)
	}
	if v, ok := d.GetOk("product_name"); ok {
		x := (v.(string))
		o.SetProductName(x)
	}
	if v, ok := d.GetOk("product_revision"); ok {
		x := (v.(string))
		o.SetProductRevision(x)
	}
	if v, ok := d.GetOk("read_error_count"); ok {
		x := (v.(string))
		o.SetReadErrorCount(x)
	}
	if v, ok := d.GetOk("read_error_threshold"); ok {
		x := (v.(string))
		o.SetReadErrorThreshold(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("write_enabled"); ok {
		x := (v.(string))
		o.SetWriteEnabled(x)
	}
	if v, ok := d.GetOk("write_error_count"); ok {
		x := (v.(string))
		o.SetWriteErrorCount(x)
	}
	if v, ok := d.GetOk("write_error_threshold"); ok {
		x := (v.(string))
		o.SetWriteErrorThreshold(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageFlexUtilPhysicalDrive object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageFlexUtilPhysicalDriveList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageFlexUtilPhysicalDrive: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageFlexUtilPhysicalDriveList.GetCount()
	var i int32
	var storageFlexUtilPhysicalDriveResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageFlexUtilPhysicalDriveList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageFlexUtilPhysicalDrive: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageFlexUtilPhysicalDriveList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageFlexUtilPhysicalDrive data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["block_size"] = (s.GetBlockSize())
				temp["capacity"] = (s.GetCapacity())
				temp["class_id"] = (s.GetClassId())
				temp["controller"] = (s.GetController())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())
				temp["drives_enabled"] = (s.GetDrivesEnabled())
				temp["health"] = (s.GetHealth())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["manufacturer_date"] = (s.GetManufacturerDate())
				temp["manufacturer_id"] = (s.GetManufacturerId())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["oem_id"] = (s.GetOemId())
				temp["partition_count"] = (s.GetPartitionCount())
				temp["pd_status"] = (s.GetPdStatus())
				temp["physical_drive"] = (s.GetPhysicalDrive())
				temp["product_name"] = (s.GetProductName())
				temp["product_revision"] = (s.GetProductRevision())
				temp["read_error_count"] = (s.GetReadErrorCount())
				temp["read_error_threshold"] = (s.GetReadErrorThreshold())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())
				temp["serial"] = (s.GetSerial())

				temp["storage_flex_util_controller"] = flattenMapStorageFlexUtilControllerRelationship(s.GetStorageFlexUtilController(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["write_enabled"] = (s.GetWriteEnabled())
				temp["write_error_count"] = (s.GetWriteErrorCount())
				temp["write_error_threshold"] = (s.GetWriteErrorThreshold())
				storageFlexUtilPhysicalDriveResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageFlexUtilPhysicalDriveResults))
	if err := d.Set("results", storageFlexUtilPhysicalDriveResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageFlexUtilPhysicalDriveResults[0]["moid"].(string))
	return de
}
