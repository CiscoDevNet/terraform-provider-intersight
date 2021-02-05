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

func dataSourceStorageFlexUtilPhysicalDrive() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageFlexUtilPhysicalDriveRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
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
		},
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
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageFlexUtilPhysicalDriveList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching StorageFlexUtilPhysicalDrive: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageFlexUtilPhysicalDrive list: %s", err.Error())
	}
	var s = &models.StorageFlexUtilPhysicalDriveList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageFlexUtilPhysicalDrive list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageFlexUtilPhysicalDrive data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageFlexUtilPhysicalDrive data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageFlexUtilPhysicalDrive{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("block_size", (s.GetBlockSize())); err != nil {
				return diag.Errorf("error occurred while setting property BlockSize: %s", err.Error())
			}
			if err := d.Set("capacity", (s.GetCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property Capacity: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("controller", (s.GetController())); err != nil {
				return diag.Errorf("error occurred while setting property Controller: %s", err.Error())
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceMoId: %s", err.Error())
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return diag.Errorf("error occurred while setting property Dn: %s", err.Error())
			}
			if err := d.Set("drives_enabled", (s.GetDrivesEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property DrivesEnabled: %s", err.Error())
			}
			if err := d.Set("health", (s.GetHealth())); err != nil {
				return diag.Errorf("error occurred while setting property Health: %s", err.Error())
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InventoryDeviceInfo: %s", err.Error())
			}
			if err := d.Set("manufacturer_date", (s.GetManufacturerDate())); err != nil {
				return diag.Errorf("error occurred while setting property ManufacturerDate: %s", err.Error())
			}
			if err := d.Set("manufacturer_id", (s.GetManufacturerId())); err != nil {
				return diag.Errorf("error occurred while setting property ManufacturerId: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("oem_id", (s.GetOemId())); err != nil {
				return diag.Errorf("error occurred while setting property OemId: %s", err.Error())
			}
			if err := d.Set("partition_count", (s.GetPartitionCount())); err != nil {
				return diag.Errorf("error occurred while setting property PartitionCount: %s", err.Error())
			}
			if err := d.Set("pd_status", (s.GetPdStatus())); err != nil {
				return diag.Errorf("error occurred while setting property PdStatus: %s", err.Error())
			}
			if err := d.Set("physical_drive", (s.GetPhysicalDrive())); err != nil {
				return diag.Errorf("error occurred while setting property PhysicalDrive: %s", err.Error())
			}
			if err := d.Set("product_name", (s.GetProductName())); err != nil {
				return diag.Errorf("error occurred while setting property ProductName: %s", err.Error())
			}
			if err := d.Set("product_revision", (s.GetProductRevision())); err != nil {
				return diag.Errorf("error occurred while setting property ProductRevision: %s", err.Error())
			}
			if err := d.Set("read_error_count", (s.GetReadErrorCount())); err != nil {
				return diag.Errorf("error occurred while setting property ReadErrorCount: %s", err.Error())
			}
			if err := d.Set("read_error_threshold", (s.GetReadErrorThreshold())); err != nil {
				return diag.Errorf("error occurred while setting property ReadErrorThreshold: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return diag.Errorf("error occurred while setting property Revision: %s", err.Error())
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return diag.Errorf("error occurred while setting property Rn: %s", err.Error())
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return diag.Errorf("error occurred while setting property Serial: %s", err.Error())
			}

			if err := d.Set("storage_flex_util_controller", flattenMapStorageFlexUtilControllerRelationship(s.GetStorageFlexUtilController(), d)); err != nil {
				return diag.Errorf("error occurred while setting property StorageFlexUtilController: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return diag.Errorf("error occurred while setting property Vendor: %s", err.Error())
			}
			if err := d.Set("write_enabled", (s.GetWriteEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property WriteEnabled: %s", err.Error())
			}
			if err := d.Set("write_error_count", (s.GetWriteErrorCount())); err != nil {
				return diag.Errorf("error occurred while setting property WriteErrorCount: %s", err.Error())
			}
			if err := d.Set("write_error_threshold", (s.GetWriteErrorThreshold())); err != nil {
				return diag.Errorf("error occurred while setting property WriteErrorThreshold: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
