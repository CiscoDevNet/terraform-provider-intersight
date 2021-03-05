package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageFlexFlashControllerProps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageFlexFlashControllerPropsRead,
		Schema: map[string]*schema.Schema{
			"cards_manageable": {
				Description: "Manageable card on the flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"configured_mode": {
				Description: "Mode configured on the flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller_name": {
				Description: "The current name of the flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"controller_status": {
				Description: "The current status of the flex flash controller.",
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
			"fw_version": {
				Description: "Firmware version of the flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal_state": {
				Description: "Internal state of the flex flash controller.",
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
			"operating_mode": {
				Description: "Operating mode of flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"physical_drive_count": {
				Description: "Number of connected physical drives to a specific Flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_name": {
				Description: "Product name of the flex flash controller.",
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
			"startup_fw_version": {
				Description: "Startup firmware version of the Flex flash controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "This field identifies the vendor of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"virtual_drive_count": {
				Description: "Number of virtual drives for a specific Flex flash controller.",
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
					"cards_manageable": {
						Description: "Manageable card on the flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configured_mode": {
						Description: "Mode configured on the flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"controller_name": {
						Description: "The current name of the flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"controller_status": {
						Description: "The current status of the flex flash controller.",
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
					"fw_version": {
						Description: "Firmware version of the flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"internal_state": {
						Description: "Internal state of the flex flash controller.",
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
					"operating_mode": {
						Description: "Operating mode of flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"physical_drive_count": {
						Description: "Number of connected physical drives to a specific Flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"product_name": {
						Description: "Product name of the flex flash controller.",
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
					"startup_fw_version": {
						Description: "Startup firmware version of the Flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"storage_flex_flash_controller": {
						Description: "A reference to a storageFlexFlashController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"virtual_drive_count": {
						Description: "Number of virtual drives for a specific Flex flash controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageFlexFlashControllerPropsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageFlexFlashControllerProps{}
	if v, ok := d.GetOk("cards_manageable"); ok {
		x := (v.(string))
		o.SetCardsManageable(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("configured_mode"); ok {
		x := (v.(string))
		o.SetConfiguredMode(x)
	}
	if v, ok := d.GetOk("controller_name"); ok {
		x := (v.(string))
		o.SetControllerName(x)
	}
	if v, ok := d.GetOk("controller_status"); ok {
		x := (v.(string))
		o.SetControllerStatus(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fw_version"); ok {
		x := (v.(string))
		o.SetFwVersion(x)
	}
	if v, ok := d.GetOk("internal_state"); ok {
		x := (v.(string))
		o.SetInternalState(x)
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
	if v, ok := d.GetOk("operating_mode"); ok {
		x := (v.(string))
		o.SetOperatingMode(x)
	}
	if v, ok := d.GetOk("physical_drive_count"); ok {
		x := (v.(string))
		o.SetPhysicalDriveCount(x)
	}
	if v, ok := d.GetOk("product_name"); ok {
		x := (v.(string))
		o.SetProductName(x)
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
	if v, ok := d.GetOk("startup_fw_version"); ok {
		x := (v.(string))
		o.SetStartupFwVersion(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("virtual_drive_count"); ok {
		x := (v.(string))
		o.SetVirtualDriveCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageFlexFlashControllerProps object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageFlexFlashControllerPropsList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageFlexFlashControllerProps: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageFlexFlashControllerPropsList.GetCount()
	var i int32
	var storageFlexFlashControllerPropsResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageFlexFlashControllerPropsList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageFlexFlashControllerProps: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageFlexFlashControllerPropsList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageFlexFlashControllerProps data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["cards_manageable"] = (s.GetCardsManageable())
				temp["class_id"] = (s.GetClassId())
				temp["configured_mode"] = (s.GetConfiguredMode())
				temp["controller_name"] = (s.GetControllerName())
				temp["controller_status"] = (s.GetControllerStatus())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())
				temp["fw_version"] = (s.GetFwVersion())
				temp["internal_state"] = (s.GetInternalState())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["operating_mode"] = (s.GetOperatingMode())
				temp["physical_drive_count"] = (s.GetPhysicalDriveCount())
				temp["product_name"] = (s.GetProductName())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())
				temp["serial"] = (s.GetSerial())
				temp["startup_fw_version"] = (s.GetStartupFwVersion())

				temp["storage_flex_flash_controller"] = flattenMapStorageFlexFlashControllerRelationship(s.GetStorageFlexFlashController(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["virtual_drive_count"] = (s.GetVirtualDriveCount())
				storageFlexFlashControllerPropsResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageFlexFlashControllerPropsResults))
	if err := d.Set("results", storageFlexFlashControllerPropsResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageFlexFlashControllerPropsResults[0]["moid"].(string))
	return de
}
