package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexHxapVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHxapVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"age": {
				Description: "Denotes age or life time of the VM in nano seconds.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"failure_reason": {
				Description: "Reason of the failure when VM is in failed state.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"mod_time": {
				Description: "The time when this managed object was last modified.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"power_state": {
				Description: "Power state of the virtual machine.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.\n* `` - The entity's power state is not available.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"start_time": {
				Description: "Denotes the VM start timestamp.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Status of virtual machine.\n* `Unknown` - Virtual machine state is not available.\n* `Running` - Virtual machine is running normally.\n* `Stopped` - Virtual machine has been stopped.\n* `WaitForLaunch` - Virtual machine is wating to be launched.\n* `Paused` - Virtual machine is currently paused.\n* `ImportInProgress` - Virtual machine image is being imported into the platform.\n* `ImportFailed` - Virtual machine image import operation failed.\n* `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress.\n* `DiskCloneFailed` - Disk clone operation for the virtual machine failed.\n* `Processing` - Virtual machine is being created.\n* `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.\n* `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation.\n* `` - Virtual machine status is not available.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "The uuid of this virtual machine. The uuid is internally generated and not user specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_moid": {
					Description: "The Account ID for this managed object.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
					"ancestors": {
						Description: "An array of relationships to moBaseMo resources.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					"create_time": {
						Description: "The time when this managed object was created.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
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
									Description: "Virtual machine bridge name of interface.\n* `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration.\n* `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives.\n* `scsi` - SCSI (Small Computer System Interface) bus used..",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
												Description: "Access mode of the virtual disk.\n* `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine.\n* `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines.\n* `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.\n* `` - Unknown disk access mode.",
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
												Description: "Disk capacity represented in bytes.",
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
												Description: "File mode of the disk, example - Filesystem, Block.\n* `Block` - It is a Block virtual disk.\n* `Filesystem` - It is a File system virtual disk.\n* `` - Disk mode is either unknown or not supported.",
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
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
															Description: "Current state of the virtual disk.\n* `Unknown` - No details available on the disk state.\n* `Succeeded` - Last operation on the disk has been successful.\n* `ImportInProgress` - Import operation on the disk is in progress.\n* `ImportFailed` - Import operation on the disk has failed.\n* `CloneInProgress` - Disk clone operation on the disk is in progress.\n* `CloneFailed` - Clone operation on the disk has failed.\n* `CloneScheduled` - Clone operation on the disk has been scheduled.\n* `ImportScheduled` - Import operation on the disk has been scheduled.\n* `Pending` - Submitted operation on the disk is currently pending.\n* `` - Disk state is not available.",
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
														"volume_name": {
															Description: "Name of the Volume associated with virtual machine disk.",
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
					"domain_group_moid": {
						Description: "The DomainGroup ID for this managed object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"failure_reason": {
						Description: "Reason of the failure when VM is in failed state.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
									Description: "Virtual machine interface model.\n* `Unknown` - The type of the network adaptor type is unknown.\n* `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC.\n* `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support.\n* `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.\n* `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features.\n* `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later.\n* `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets.\n* `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length.\n* `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance.\n* `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware.\n* `` - Default network adaptor type supported by the hypervisor.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the virtual machine interface.",
									Type:        schema.TypeString,
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					"mod_time": {
						Description: "The time when this managed object was last modified.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"owners": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"parent": {
						Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"permission_resources": {
						Description: "An array of relationships to moBaseMo resources.",
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
					"power_state": {
						Description: "Power state of the virtual machine.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.\n* `` - The entity's power state is not available.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					"shared_scope": {
						Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"start_time": {
						Description: "Denotes the VM start timestamp.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "Status of virtual machine.\n* `Unknown` - Virtual machine state is not available.\n* `Running` - Virtual machine is running normally.\n* `Stopped` - Virtual machine has been stopped.\n* `WaitForLaunch` - Virtual machine is wating to be launched.\n* `Paused` - Virtual machine is currently paused.\n* `ImportInProgress` - Virtual machine image is being imported into the platform.\n* `ImportFailed` - Virtual machine image import operation failed.\n* `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress.\n* `DiskCloneFailed` - Disk clone operation for the virtual machine failed.\n* `Processing` - Virtual machine is being created.\n* `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.\n* `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation.\n* `` - Virtual machine status is not available.",
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
					"uuid": {
						Description: "The uuid of this virtual machine. The uuid is internally generated and not user specified.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_context": {
						Description: "The versioning info for this managed object.",
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
								"interested_mos": {
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
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"ref_mo": {
									Description: "A reference to the original Managed Object.",
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
								"timestamp": {
									Description: "The time this versioned Managed Object was created.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"nr_version": {
									Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"version_type": {
									Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexHxapVirtualMachineRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHxapVirtualMachine{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("age"); ok {
		x := (v.(string))
		o.SetAge(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("failure_reason"); ok {
		x := (v.(string))
		o.SetFailureReason(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
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
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x := (v.(string))
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHxapVirtualMachine object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapVirtualMachineList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexHxapVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexHxapVirtualMachine: %s", responseErr.Error())
	}
	count := countResponse.HyperflexHxapVirtualMachineList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexHxapVirtualMachine data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexHxapVirtualMachineResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapVirtualMachineList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexHxapVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexHxapVirtualMachine: %s", responseErr.Error())
		}
		results := resMo.HyperflexHxapVirtualMachineList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["affinity_selectors"] = flattenListInfraMetaData(s.GetAffinitySelectors(), d)
				temp["age"] = (s.GetAge())

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["anti_affinity_selectors"] = flattenListInfraMetaData(s.GetAntiAffinitySelectors(), d)

				temp["capacity"] = flattenMapInfraHardwareInfo(s.GetCapacity(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapHyperflexHxapClusterRelationship(s.GetCluster(), d)

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["disks"] = flattenListHyperflexVmDisk(s.GetDisks(), d)
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["failure_reason"] = (s.GetFailureReason())

				temp["guest_info"] = flattenMapVirtualizationGuestInfo(s.GetGuestInfo(), d)

				temp["host"] = flattenMapHyperflexHxapHostRelationship(s.GetHost(), d)
				temp["hypervisor_type"] = (s.GetHypervisorType())
				temp["identity"] = (s.GetIdentity())

				temp["interfaces"] = flattenListHyperflexVmInterface(s.GetInterfaces(), d)
				temp["ip_address"] = (s.GetIpAddress())

				temp["labels"] = flattenListInfraMetaData(s.GetLabels(), d)

				temp["memory_capacity"] = flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["network_count"] = (s.GetNetworkCount())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["power_state"] = (s.GetPowerState())

				temp["processor_capacity"] = flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["start_time"] = (s.GetStartTime())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uuid"] = (s.GetUuid())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				hyperflexHxapVirtualMachineResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexHxapVirtualMachineResults))
	if err := d.Set("results", hyperflexHxapVirtualMachineResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexHxapVirtualMachineResults[0]["moid"].(string))
	return de
}
