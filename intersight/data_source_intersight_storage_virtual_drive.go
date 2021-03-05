package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageVirtualDrive() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageVirtualDriveRead,
		Schema: map[string]*schema.Schema{
			"access_policy": {
				Description: "The access policy of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"actual_write_cache_policy": {
				Description: "The current write cache policy of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"available_size": {
				Description: "Available storage capacity of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"block_size": {
				Description: "Block size of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"bootable": {
				Description: "The virtual drive bootable state.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_state": {
				Description: "The configuration state of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"configured_write_cache_policy": {
				Description: "The requested write cache policy of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_protocol": {
				Description: "The connection protocol of the virtual drive.",
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
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_cache": {
				Description: "The state of the drive cache of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_security": {
				Description: "The driveSecurity state of the Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_state": {
				Description: "The state of the Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"io_policy": {
				Description: "The Input/Output Policy defined on the Virtual drive.",
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
			"name": {
				Description: "The name of the Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_blocks": {
				Description: "Number of Blocks on the Physical Disk.",
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
			"oper_state": {
				Description: "The current operational state of Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "The current operability state of Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"physical_block_size": {
				Description: "The block size of the the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "The presence status of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"read_policy": {
				Description: "The read-ahead cache mode of the virtual drive.",
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
			"security_flags": {
				Description: "The security flags set for this virtual drive.",
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
			"size": {
				Description: "The size of the virtual drive in MB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"strip_size": {
				Description: "The strip size is the portion of a stripe that resides on a single drive in the drive group, this is measured in KB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "The raid type of the virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "The uuid of the virtual drive.",
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
			"vendor_uuid": {
				Description: "The UUID value of the vendor.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"virtual_drive_id": {
				Description: "The identifier for this Virtual drive.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"access_policy": {
					Description: "The access policy of the virtual drive.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"actual_write_cache_policy": {
						Description: "The current write cache policy of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"available_size": {
						Description: "Available storage capacity of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"block_size": {
						Description: "Block size of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"bootable": {
						Description: "The virtual drive bootable state.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"config_state": {
						Description: "The configuration state of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"configured_write_cache_policy": {
						Description: "The requested write cache policy of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"connection_protocol": {
						Description: "The connection protocol of the virtual drive.",
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
					"disk_group": {
						Description: "A reference to a storageDiskGroup resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
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
					"dn": {
						Description: "The Distinguished Name unambiguously identifies an object in the system.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drive_cache": {
						Description: "The state of the drive cache of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drive_security": {
						Description: "The driveSecurity state of the Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drive_state": {
						Description: "The state of the Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
					"io_policy": {
						Description: "The Input/Output Policy defined on the Virtual drive.",
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
					"name": {
						Description: "The name of the Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"num_blocks": {
						Description: "Number of Blocks on the Physical Disk.",
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
					"oper_state": {
						Description: "The current operational state of Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"operability": {
						Description: "The current operability state of Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"physical_block_size": {
						Description: "The block size of the the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"physical_disk_usages": {
						Description: "An array of relationships to storagePhysicalDiskUsage resources.",
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
					"presence": {
						Description: "The presence status of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"read_policy": {
						Description: "The read-ahead cache mode of the virtual drive.",
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
					"security_flags": {
						Description: "The security flags set for this virtual drive.",
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
					"size": {
						Description: "The size of the virtual drive in MB.",
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
					"storage_virtual_drive_container": {
						Description: "A reference to a storageVirtualDriveContainer resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
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
					"strip_size": {
						Description: "The strip size is the portion of a stripe that resides on a single drive in the drive group, this is measured in KB.",
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
					"type": {
						Description: "The raid type of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"uuid": {
						Description: "The uuid of the virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"vd_member_eps": {
						Description: "An array of relationships to storageVdMemberEp resources.",
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
					"vendor": {
						Description: "This field identifies the vendor of the given component.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"vendor_uuid": {
						Description: "The UUID value of the vendor.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"virtual_drive_extension": {
						Description: "A reference to a storageVirtualDriveExtension resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"virtual_drive_id": {
						Description: "The identifier for this Virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageVirtualDriveRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageVirtualDrive{}
	if v, ok := d.GetOk("access_policy"); ok {
		x := (v.(string))
		o.SetAccessPolicy(x)
	}
	if v, ok := d.GetOk("actual_write_cache_policy"); ok {
		x := (v.(string))
		o.SetActualWriteCachePolicy(x)
	}
	if v, ok := d.GetOk("available_size"); ok {
		x := (v.(string))
		o.SetAvailableSize(x)
	}
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
	if v, ok := d.GetOk("config_state"); ok {
		x := (v.(string))
		o.SetConfigState(x)
	}
	if v, ok := d.GetOk("configured_write_cache_policy"); ok {
		x := (v.(string))
		o.SetConfiguredWriteCachePolicy(x)
	}
	if v, ok := d.GetOk("connection_protocol"); ok {
		x := (v.(string))
		o.SetConnectionProtocol(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("drive_cache"); ok {
		x := (v.(string))
		o.SetDriveCache(x)
	}
	if v, ok := d.GetOk("drive_security"); ok {
		x := (v.(string))
		o.SetDriveSecurity(x)
	}
	if v, ok := d.GetOk("drive_state"); ok {
		x := (v.(string))
		o.SetDriveState(x)
	}
	if v, ok := d.GetOk("io_policy"); ok {
		x := (v.(string))
		o.SetIoPolicy(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("num_blocks"); ok {
		x := (v.(string))
		o.SetNumBlocks(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("physical_block_size"); ok {
		x := (v.(string))
		o.SetPhysicalBlockSize(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("read_policy"); ok {
		x := (v.(string))
		o.SetReadPolicy(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("security_flags"); ok {
		x := (v.(string))
		o.SetSecurityFlags(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := (v.(string))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("strip_size"); ok {
		x := (v.(string))
		o.SetStripSize(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("vendor_uuid"); ok {
		x := (v.(string))
		o.SetVendorUuid(x)
	}
	if v, ok := d.GetOk("virtual_drive_id"); ok {
		x := (v.(string))
		o.SetVirtualDriveId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageVirtualDrive object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageVirtualDriveList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageVirtualDrive: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageVirtualDriveList.GetCount()
	var i int32
	var storageVirtualDriveResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageVirtualDriveList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageVirtualDrive: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageVirtualDriveList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageVirtualDrive data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["access_policy"] = (s.GetAccessPolicy())
				temp["actual_write_cache_policy"] = (s.GetActualWriteCachePolicy())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["available_size"] = (s.GetAvailableSize())
				temp["block_size"] = (s.GetBlockSize())
				temp["bootable"] = (s.GetBootable())
				temp["class_id"] = (s.GetClassId())
				temp["config_state"] = (s.GetConfigState())
				temp["configured_write_cache_policy"] = (s.GetConfiguredWriteCachePolicy())
				temp["connection_protocol"] = (s.GetConnectionProtocol())
				temp["device_mo_id"] = (s.GetDeviceMoId())

				temp["disk_group"] = flattenMapStorageDiskGroupRelationship(s.GetDiskGroup(), d)
				temp["dn"] = (s.GetDn())
				temp["drive_cache"] = (s.GetDriveCache())
				temp["drive_security"] = (s.GetDriveSecurity())
				temp["drive_state"] = (s.GetDriveState())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["io_policy"] = (s.GetIoPolicy())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["num_blocks"] = (s.GetNumBlocks())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_state"] = (s.GetOperState())
				temp["operability"] = (s.GetOperability())
				temp["physical_block_size"] = (s.GetPhysicalBlockSize())

				temp["physical_disk_usages"] = flattenListStoragePhysicalDiskUsageRelationship(s.GetPhysicalDiskUsages(), d)
				temp["presence"] = (s.GetPresence())
				temp["read_policy"] = (s.GetReadPolicy())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())
				temp["security_flags"] = (s.GetSecurityFlags())
				temp["serial"] = (s.GetSerial())
				temp["size"] = (s.GetSize())

				temp["storage_controller"] = flattenMapStorageControllerRelationship(s.GetStorageController(), d)

				temp["storage_virtual_drive_container"] = flattenMapStorageVirtualDriveContainerRelationship(s.GetStorageVirtualDriveContainer(), d)
				temp["strip_size"] = (s.GetStripSize())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["uuid"] = (s.GetUuid())

				temp["vd_member_eps"] = flattenListStorageVdMemberEpRelationship(s.GetVdMemberEps(), d)
				temp["vendor"] = (s.GetVendor())
				temp["vendor_uuid"] = (s.GetVendorUuid())

				temp["virtual_drive_extension"] = flattenMapStorageVirtualDriveExtensionRelationship(s.GetVirtualDriveExtension(), d)
				temp["virtual_drive_id"] = (s.GetVirtualDriveId())
				storageVirtualDriveResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageVirtualDriveResults))
	if err := d.Set("results", storageVirtualDriveResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageVirtualDriveResults[0]["moid"].(string))
	return de
}
