package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStorageFlexUtilPhysicalDrive() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageFlexUtilPhysicalDriveRead,
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
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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

func dataSourceStorageFlexUtilPhysicalDriveRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewStorageFlexUtilPhysicalDriveWithDefaults()
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.StorageApi.GetStorageFlexUtilPhysicalDriveList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.StorageFlexUtilPhysicalDriveList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to StorageFlexUtilPhysicalDrive: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewStorageFlexUtilPhysicalDriveWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("block_size", (s.GetBlockSize())); err != nil {
				return fmt.Errorf("error occurred while setting property BlockSize: %+v", err)
			}
			if err := d.Set("capacity", (s.GetCapacity())); err != nil {
				return fmt.Errorf("error occurred while setting property Capacity: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("controller", (s.GetController())); err != nil {
				return fmt.Errorf("error occurred while setting property Controller: %+v", err)
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceMoId: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}
			if err := d.Set("drives_enabled", (s.GetDrivesEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property DrivesEnabled: %+v", err)
			}
			if err := d.Set("health", (s.GetHealth())); err != nil {
				return fmt.Errorf("error occurred while setting property Health: %+v", err)
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property InventoryDeviceInfo: %+v", err)
			}
			if err := d.Set("manufacturer_date", (s.GetManufacturerDate())); err != nil {
				return fmt.Errorf("error occurred while setting property ManufacturerDate: %+v", err)
			}
			if err := d.Set("manufacturer_id", (s.GetManufacturerId())); err != nil {
				return fmt.Errorf("error occurred while setting property ManufacturerId: %+v", err)
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return fmt.Errorf("error occurred while setting property Model: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("oem_id", (s.GetOemId())); err != nil {
				return fmt.Errorf("error occurred while setting property OemId: %+v", err)
			}
			if err := d.Set("partition_count", (s.GetPartitionCount())); err != nil {
				return fmt.Errorf("error occurred while setting property PartitionCount: %+v", err)
			}
			if err := d.Set("pd_status", (s.GetPdStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property PdStatus: %+v", err)
			}
			if err := d.Set("physical_drive", (s.GetPhysicalDrive())); err != nil {
				return fmt.Errorf("error occurred while setting property PhysicalDrive: %+v", err)
			}
			if err := d.Set("product_name", (s.GetProductName())); err != nil {
				return fmt.Errorf("error occurred while setting property ProductName: %+v", err)
			}
			if err := d.Set("product_revision", (s.GetProductRevision())); err != nil {
				return fmt.Errorf("error occurred while setting property ProductRevision: %+v", err)
			}
			if err := d.Set("read_error_count", (s.GetReadErrorCount())); err != nil {
				return fmt.Errorf("error occurred while setting property ReadErrorCount: %+v", err)
			}
			if err := d.Set("read_error_threshold", (s.GetReadErrorThreshold())); err != nil {
				return fmt.Errorf("error occurred while setting property ReadErrorThreshold: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return fmt.Errorf("error occurred while setting property Revision: %+v", err)
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return fmt.Errorf("error occurred while setting property Rn: %+v", err)
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return fmt.Errorf("error occurred while setting property Serial: %+v", err)
			}

			if err := d.Set("storage_flex_util_controller", flattenMapStorageFlexUtilControllerRelationship(s.GetStorageFlexUtilController(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property StorageFlexUtilController: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property Vendor: %+v", err)
			}
			if err := d.Set("write_enabled", (s.GetWriteEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property WriteEnabled: %+v", err)
			}
			if err := d.Set("write_error_count", (s.GetWriteErrorCount())); err != nil {
				return fmt.Errorf("error occurred while setting property WriteErrorCount: %+v", err)
			}
			if err := d.Set("write_error_threshold", (s.GetWriteErrorThreshold())); err != nil {
				return fmt.Errorf("error occurred while setting property WriteErrorThreshold: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
