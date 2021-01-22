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

func dataSourceHyperflexHxapVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHxapVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"affinity_selectors": {
				Type:     schema.TypeList,
				Optional: true,
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
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
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
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"age": {
				Description: "Denotes age or life time of the VM in nano seconds.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"anti_affinity_selectors": {
				Type:     schema.TypeList,
				Optional: true,
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
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
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
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"capacity": {
				Description: "Provisioned CPU and memory information for this virtual machine.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"cpu_cores": {
							Description: "The number of cpu cores on this hardware platform.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_speed": {
							Description: "Speed of cpu in MHz. Usually cpu speeds are reported for modern cpus in GHz but MHz makes it more precise.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"memory_size": {
							Description: "The amount of memory allocated (bytes) to this hardware platform.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster": {
				Description: "A reference to a hyperflexHxapCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"disks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"boot_order": {
							Description: "Boot order for this disk.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"bus": {
							Description: "Virtual machine brige name of interface.\n* `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration.\n* `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives.\n* `scsi` - SCSI (Small Computer System Interface) bus used..",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Disk name associated with virtual machine.",
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
						"type": {
							Description: "Type of the Disk (hdd, cdrom).\n* `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image.\n* `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"virtual_disk": {
							Description: "Virtual disk configuration.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_mode": {
										Description: "Access mode of the virtual disk.\n* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.\n* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.\n* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"capacity": {
										Description: "Total disk capacity represented in Gibibytes (GiB).",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"mode": {
										Description: "File mode of the disk, example - Filesystem, Block.\n* `Block` - It is a Block virtual disk.\n* `Filesystem` - It is a File system virtual disk.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the virtual disk.",
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
									"source_file_path": {
										Description: "Source file path associated with virtual machine disk.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"source_virtual_disk": {
										Description: "Source disk name from where the clone is done.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Description: "Status of virtual machine disk.",
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
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"download_percentage": {
													Description: "Percentage of download completed.",
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
												"state": {
													Description: "Current state of the virtual disk.\n* `Unknown` - No details available on the disk state.\n* `Succeeded` - Last operation on the disk has been successful.\n* `ImportInProgress` - Import operation on the disk is in progress.\n* `ImportFailed` - Import operation on the disk has failed.\n* `CloneInProgress` - Disk clone operation on the disk is in progress.\n* `CloneFailed` - Clone operation on the disk has failed.\n* `CloneScheduled` - Clone operation on the disk has been scheduled.\n* `ImportScheduled` - Import operation on the disk has been scheduled.\n* `Pending` - Submitted operation on the disk is currently pending.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"volume_handle": {
													Description: "Identity of the Volume associated with virtual machine disk.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"uuid": {
										Description: "UUID of the virtual disk.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"virtual_disk_reference": {
							Description: "Virtual disk reference name.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"failure_reason": {
				Description: "Reason of the failure when VM is in failed state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"graphic_console_url": {
				Description: "Graphical console URL of this VM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"guest_info": {
				Description: "Guest operating system details running on this machine.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"hostname": {
							Description: "Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_address": {
							Description: "Primary IP address of the guest os.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "The name of the guest running on this VM. This may not be the same as the hostname.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"operating_system": {
							Description: "The name of the guest OS running on this VM (Cent OS 4/5/6/7).",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"host": {
				Description: "A reference to a hyperflexHxapHost resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"hypervisor_type": {
				Description: "Type of hypervisor where the virtual machine is hosted for example ESXi.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"bridge": {
							Description: "Virtual machine brige name of interface.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_address": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"mac_address": {
							Description: "Virtual machine interface mac address.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"model": {
							Description: "Virtual machine interface model.",
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
					},
				},
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
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
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
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
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"memory_capacity": {
				Description: "The capacity and usage information for memory on this virtual machine.",
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
						"capacity": {
							Description: "The total memory capacity of the entity in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"free": {
							Description: "Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used": {
							Description: "Memory (bytes) that has been already used up, as a point-in-time snapshot.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "User-provided name to identify the virtual machine.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_count": {
				Description: "Number network interfaces associated with the virtual machine.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"power_state": {
				Description: "Power state of the virtual machine.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_capacity": {
				Description: "The capacity and usage information for CPU power on this virtual machine.",
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
						"capacity": {
							Description: "Total capacity of the entity in MHz.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"free": {
							Description: "Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"used": {
							Description: "Used CPU capacity of the entity in MHz, as a point-in-time snapshot.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"start_time": {
				Description: "Denotes the VM start timestamp.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Status of virtual machine.\n* `Unknown` - Virtual machine state is not available.\n* `Running` - Virtual machine is running normally.\n* `Stopped` - Virtual machine has been stopped.\n* `WaitForLaunch` - Virtual machine is wating to be launched.\n* `Paused` - Virtual machine is currently paused.\n* `ImportInProgress` - Virtual machine image is being imported into the platform.\n* `ImportFailed` - Virtual machine image import operation failed.\n* `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress.\n* `DiskCloneFailed` - Disk clone operation for the virtual machine failed.\n* `Processing` - Virtual machine is being created.\n* `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.\n* `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"up_time": {
				Description: "Denotes how long this VM has been running in nano seconds.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "The uuid of this virtual machine. The uuid is internally generated and not user specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceHyperflexHxapVirtualMachineRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHxapVirtualMachine{}
	if v, ok := d.GetOk("age"); ok {
		x := (v.(string))
		o.SetAge(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.SetFailureReason(x)
	}
	if v, ok := d.GetOk("graphic_console_url"); ok {
		x := (v.(string))
		o.SetGraphicConsoleUrl(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("network_count"); ok {
		x := int64(v.(int))
		o.SetNetworkCount(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("power_state"); ok {
		x := (v.(string))
		o.SetPowerState(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x := (v.(string))
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("up_time"); ok {
		x := (v.(string))
		o.SetUpTime(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHxapVirtualMachine object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapVirtualMachineList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexHxapVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexHxapVirtualMachine list: %s", err.Error())
	}
	var s = &models.HyperflexHxapVirtualMachineList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexHxapVirtualMachine list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexHxapVirtualMachine data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexHxapVirtualMachine data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexHxapVirtualMachine{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("affinity_selectors", flattenListInfraMetaData(s.GetAffinitySelectors(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AffinitySelectors: %s", err.Error())
			}
			if err := d.Set("age", (s.GetAge())); err != nil {
				return diag.Errorf("error occurred while setting property Age: %s", err.Error())
			}

			if err := d.Set("anti_affinity_selectors", flattenListInfraMetaData(s.GetAntiAffinitySelectors(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AntiAffinitySelectors: %s", err.Error())
			}

			if err := d.Set("capacity", flattenMapInfraHardwareInfo(s.GetCapacity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Capacity: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster", flattenMapHyperflexHxapClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}

			if err := d.Set("disks", flattenListHyperflexVmDisk(s.GetDisks(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Disks: %s", err.Error())
			}
			if err := d.Set("failure_reason", (s.GetFailureReason())); err != nil {
				return diag.Errorf("error occurred while setting property FailureReason: %s", err.Error())
			}
			if err := d.Set("graphic_console_url", (s.GetGraphicConsoleUrl())); err != nil {
				return diag.Errorf("error occurred while setting property GraphicConsoleUrl: %s", err.Error())
			}

			if err := d.Set("guest_info", flattenMapVirtualizationGuestInfo(s.GetGuestInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property GuestInfo: %s", err.Error())
			}

			if err := d.Set("host", flattenMapHyperflexHxapHostRelationship(s.GetHost(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Host: %s", err.Error())
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorType: %s", err.Error())
			}
			if err := d.Set("identity", (s.GetIdentity())); err != nil {
				return diag.Errorf("error occurred while setting property Identity: %s", err.Error())
			}

			if err := d.Set("interfaces", flattenListHyperflexVmInterface(s.GetInterfaces(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Interfaces: %s", err.Error())
			}
			if err := d.Set("ip_address", (s.GetIpAddress())); err != nil {
				return diag.Errorf("error occurred while setting property IpAddress: %s", err.Error())
			}

			if err := d.Set("labels", flattenListInfraMetaData(s.GetLabels(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Labels: %s", err.Error())
			}

			if err := d.Set("memory_capacity", flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property MemoryCapacity: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("network_count", (s.GetNetworkCount())); err != nil {
				return diag.Errorf("error occurred while setting property NetworkCount: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("power_state", (s.GetPowerState())); err != nil {
				return diag.Errorf("error occurred while setting property PowerState: %s", err.Error())
			}

			if err := d.Set("processor_capacity", flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ProcessorCapacity: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("start_time", (s.GetStartTime())); err != nil {
				return diag.Errorf("error occurred while setting property StartTime: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("up_time", (s.GetUpTime())); err != nil {
				return diag.Errorf("error occurred while setting property UpTime: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
