package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStoragePhysicalDisk() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStoragePhysicalDiskRead,
		Schema: map[string]*schema.Schema{
			"block_size": {
				Description: "The block size of the physical disk in bytes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"bootable": {
				Description: "This field identifies the disk drive as bootable if set to true.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"configuration_checkpoint": {
				Description: "The current configuration checkpoint of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"configuration_state": {
				Description: "The current configuration state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Description: "The database identifier of the registered device of an object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"discovered_path": {
				Description: "The discovered path of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"disk_id": {
				Description: "This field identifies the ID assigned to physical disks.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"disk_state": {
				Description: "This field identifies the health of the disk.",
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
			"drive_firmware": {
				Description: "This field identifies the disk firmware running in the disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"drive_state": {
				Description: "The drive state as reported by the controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"fde_capable": {
				Description: "Full-Disk Encryption capability parameter of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hot_spare_type": {
				Description: "Type of hotspare configured on the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"link_speed": {
				Description: "The speed of the link between the drive and the controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"link_state": {
				Description: "The current link state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"num_blocks": {
				Description: "The number of blocks present on the physical disk.",
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
			"oper_power_state": {
				Description: "Operational power of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_qualifier_reason": {
				Description: "For certain states, indicates the reason why the operState is in that state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "This field identifies the disk operability of the disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"physical_block_size": {
				Description: "The block size of the installed physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "This field identifies the Product ID for physicalDisk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "The presence state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Description: "This field identifies the disk protocol used for communication.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"raw_size": {
				Description: "The raw size of the physical disk in MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"secured": {
				Description: "This field identifies whether the disk is encrypted.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial": {
				Description: "This field identifies the serial of the given component.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"size": {
				Description: "The size of the physical disk in MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"thermal": {
				Description: "Thermal state of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "This field identifies the type of the physical disk.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"variant_type": {
				Description: "The variant type of the physical disk.",
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
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"block_size": {
						Description: "The block size of the physical disk in bytes.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"bootable": {
						Description: "This field identifies the disk drive as bootable if set to true.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configuration_checkpoint": {
						Description: "The current configuration checkpoint of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"configuration_state": {
						Description: "The current configuration state of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_mo_id": {
						Description: "The database identifier of the registered device of an object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"discovered_path": {
						Description: "The discovered path of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"disk_id": {
						Description: "This field identifies the ID assigned to physical disks.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"disk_state": {
						Description: "This field identifies the health of the disk.",
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
					"drive_firmware": {
						Description: "This field identifies the disk firmware running in the disk.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"drive_state": {
						Description: "The drive state as reported by the controller.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"fde_capable": {
						Description: "Full-Disk Encryption capability parameter of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hot_spare_type": {
						Description: "Type of hotspare configured on the physical disk.",
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
					"link_speed": {
						Description: "The speed of the link between the drive and the controller.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"link_state": {
						Description: "The current link state of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"locator_led": {
						Description: "A reference to a equipmentLocatorLed resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"num_blocks": {
						Description: "The number of blocks present on the physical disk.",
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
					"oper_power_state": {
						Description: "Operational power of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"oper_qualifier_reason": {
						Description: "For certain states, indicates the reason why the operState is in that state.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"operability": {
						Description: "This field identifies the disk operability of the disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"physical_block_size": {
						Description: "The block size of the installed physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"physical_disk_extensions": {
						Description: "An array of relationships to storagePhysicalDiskExtension resources.",
						Type:        schema.TypeList,
						Optional:    true,
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
						Computed: true,
					},
					"pid": {
						Description: "This field identifies the Product ID for physicalDisk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"presence": {
						Description: "The presence state of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"protocol": {
						Description: "This field identifies the disk protocol used for communication.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"raw_size": {
						Description: "The raw size of the physical disk in MB.",
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
					"running_firmware": {
						Description: "An array of relationships to firmwareRunningFirmware resources.",
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
					"sas_ports": {
						Description: "An array of relationships to storageSasPort resources.",
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
					"secured": {
						Description: "This field identifies whether the disk is encrypted.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"serial": {
						Description: "This field identifies the serial of the given component.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"size": {
						Description: "The size of the physical disk in MB.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"storage_controller": {
						Description: "A reference to a storageController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"storage_enclosure": {
						Description: "A reference to a storageEnclosure resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"thermal": {
						Description: "Thermal state of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"type": {
						Description: "This field identifies the type of the physical disk.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"variant_type": {
						Description: "The variant type of the physical disk.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceStoragePhysicalDiskRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StoragePhysicalDisk{}
	if v, ok := d.GetOk("block_size"); ok {
		x := (v.(string))
		o.SetBlockSize(x)
	}
	if v, ok := d.GetOk("bootable"); ok {
		x := (v.(string))
		o.SetBootable(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("configuration_checkpoint"); ok {
		x := (v.(string))
		o.SetConfigurationCheckpoint(x)
	}
	if v, ok := d.GetOk("configuration_state"); ok {
		x := (v.(string))
		o.SetConfigurationState(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("discovered_path"); ok {
		x := (v.(string))
		o.SetDiscoveredPath(x)
	}
	if v, ok := d.GetOk("disk_id"); ok {
		x := (v.(string))
		o.SetDiskId(x)
	}
	if v, ok := d.GetOk("disk_state"); ok {
		x := (v.(string))
		o.SetDiskState(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("drive_firmware"); ok {
		x := (v.(string))
		o.SetDriveFirmware(x)
	}
	if v, ok := d.GetOk("drive_state"); ok {
		x := (v.(string))
		o.SetDriveState(x)
	}
	if v, ok := d.GetOk("fde_capable"); ok {
		x := (v.(string))
		o.SetFdeCapable(x)
	}
	if v, ok := d.GetOk("hot_spare_type"); ok {
		x := (v.(string))
		o.SetHotSpareType(x)
	}
	if v, ok := d.GetOk("link_speed"); ok {
		x := (v.(string))
		o.SetLinkSpeed(x)
	}
	if v, ok := d.GetOk("link_state"); ok {
		x := (v.(string))
		o.SetLinkState(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_blocks"); ok {
		x := (v.(string))
		o.SetNumBlocks(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.SetOperPowerState(x)
	}
	if v, ok := d.GetOk("oper_qualifier_reason"); ok {
		x := (v.(string))
		o.SetOperQualifierReason(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("physical_block_size"); ok {
		x := (v.(string))
		o.SetPhysicalBlockSize(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("raw_size"); ok {
		x := (v.(string))
		o.SetRawSize(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("secured"); ok {
		x := (v.(string))
		o.SetSecured(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := (v.(string))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.SetThermal(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("variant_type"); ok {
		x := (v.(string))
		o.SetVariantType(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StoragePhysicalDisk object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStoragePhysicalDiskList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StoragePhysicalDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StoragePhysicalDiskList.GetCount()
	var i int32
	var storagePhysicalDiskResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStoragePhysicalDiskList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StoragePhysicalDisk: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StoragePhysicalDiskList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StoragePhysicalDisk data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["block_size"] = (s.GetBlockSize())
				temp["bootable"] = (s.GetBootable())
				temp["class_id"] = (s.GetClassId())
				temp["configuration_checkpoint"] = (s.GetConfigurationCheckpoint())
				temp["configuration_state"] = (s.GetConfigurationState())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["discovered_path"] = (s.GetDiscoveredPath())
				temp["disk_id"] = (s.GetDiskId())
				temp["disk_state"] = (s.GetDiskState())
				temp["dn"] = (s.GetDn())
				temp["drive_firmware"] = (s.GetDriveFirmware())
				temp["drive_state"] = (s.GetDriveState())
				temp["fde_capable"] = (s.GetFdeCapable())
				temp["hot_spare_type"] = (s.GetHotSpareType())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["link_speed"] = (s.GetLinkSpeed())
				temp["link_state"] = (s.GetLinkState())

				temp["locator_led"] = flattenMapEquipmentLocatorLedRelationship(s.GetLocatorLed(), d)
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["num_blocks"] = (s.GetNumBlocks())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_power_state"] = (s.GetOperPowerState())
				temp["oper_qualifier_reason"] = (s.GetOperQualifierReason())
				temp["operability"] = (s.GetOperability())
				temp["physical_block_size"] = (s.GetPhysicalBlockSize())

				temp["physical_disk_extensions"] = flattenListStoragePhysicalDiskExtensionRelationship(s.GetPhysicalDiskExtensions(), d)
				temp["pid"] = (s.GetPid())
				temp["presence"] = (s.GetPresence())
				temp["protocol"] = (s.GetProtocol())
				temp["raw_size"] = (s.GetRawSize())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())

				temp["running_firmware"] = flattenListFirmwareRunningFirmwareRelationship(s.GetRunningFirmware(), d)

				temp["sas_ports"] = flattenListStorageSasPortRelationship(s.GetSasPorts(), d)
				temp["secured"] = (s.GetSecured())
				temp["serial"] = (s.GetSerial())
				temp["size"] = (s.GetSize())

				temp["storage_controller"] = flattenMapStorageControllerRelationship(s.GetStorageController(), d)

				temp["storage_enclosure"] = flattenMapStorageEnclosureRelationship(s.GetStorageEnclosure(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["thermal"] = (s.GetThermal())
				temp["type"] = (s.GetType())
				temp["variant_type"] = (s.GetVariantType())
				temp["vendor"] = (s.GetVendor())
				storagePhysicalDiskResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storagePhysicalDiskResults))
	if err := d.Set("results", storagePhysicalDiskResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storagePhysicalDiskResults[0]["moid"].(string))
	return de
}
