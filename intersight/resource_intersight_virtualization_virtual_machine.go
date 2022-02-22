package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVirtualizationVirtualMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVirtualizationVirtualMachineCreate,
		ReadContext:   resourceVirtualizationVirtualMachineRead,
		UpdateContext: resourceVirtualizationVirtualMachineUpdate,
		DeleteContext: resourceVirtualizationVirtualMachineDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CustomizeTagDiff,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"action": {
				Description: "Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc).\n* `None` - A place holder for the default value.\n* `PowerState` - Power action is performed on the virtual machine.\n* `Migrate` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.\n* `Create` - The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation.\n* `Delete` - The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "None",
			},
			"action_info": {
				Description: "Details of an action performed on the virtual machine. Contains name of the action performed, status, failure reason message etc.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "virtualization.ActionInfo",
						},
						"failure_reason": {
							Description: "Description of reason for failure. Derived from the workflow failure message.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"name": {
							Description: "Name of the Action performed on a resource like Virtual Machine, Disk etc.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.ActionInfo",
						},
						"status": {
							Description: "Status of the Action like InProgress, Success, Failure etc.\n* `None` - A place holder for the default value.\n* `InProgress` - Previous action triggered on the resource is still running.\n* `Success` - Previous action triggered on the resource has completed successfully.\n* `Failure` - Previous action triggered on the resource has failed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"affinity_selectors": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "infra.MetaData",
						},
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "infra.MetaData",
						},
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"ancestors": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"anti_affinity_selectors": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "infra.MetaData",
						},
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "infra.MetaData",
						},
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "virtualization.VirtualMachine",
			},
			"cloud_init_config": {
				Description: "Cloud init configuration data for virtual machine.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "virtualization.CloudInitConfig",
						},
						"config_type": {
							Description: "Virtual machine cloud init configuration type.\n* `` - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected.\n* `NoCloudSource` - Allows the user to provide user-data to the instance without running a network service.\n* `CloudConfigDrive` - Allows the user to provide user-data and network-data from cloud.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
						},
						"network_data": {
							Description: "Network configuration data for a virtual machine.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"network_data_base64_encoded": {
							Description: "Set to true, if the cloud init network data is in base64 format.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.CloudInitConfig",
						},
						"user_data": {
							Description: "User configuration data for a virtual machine such as adding user, group etc.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"user_data_base64_encoded": {
							Description: "Set to true, if the cloud init user data is in base64 format.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
			"cluster": {
				Description: "A reference to a virtualizationBaseCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
				ForceNew: true,
			},
			"cluster_esxi": {
				Description: "Cluster where virtual machine is deployed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu": {
				Description: "Number of vCPUs allocated to virtual machine.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"discovered": {
				Description: "Flag to indicate whether the configuration is created from inventory object.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"disk": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"bus": {
							Description: "Disk bus name given for a virtual machine.\n* `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration.\n* `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives.\n* `scsi` - SCSI (Small Computer System Interface) bus used..",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtio",
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.VirtualMachineDisk",
						},
						"name": {
							Description: "Virtual machine network bridge name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.VirtualMachineDisk",
						},
						"order": {
							Description: "Priority order of the disk.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"type": {
							Description: "Disk type hdd or cdrom for a virtual machine.\n* `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image.\n* `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "hdd",
						},
						"virtual_disk": {
							Description: "Virtual disk configuration.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"capacity": {
										Description: "Disk capacity to be provided with units example - 10Gi.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "virtualization.VirtualDiskConfig",
									},
									"mode": {
										Description: "File mode of the disk, example - Filesystem, Block.\n* `Block` - It is a Block virtual disk.\n* `Filesystem` - It is a File system virtual disk.\n* `` - Disk mode is either unknown or not supported.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "Block",
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "virtualization.VirtualDiskConfig",
									},
									"source_certs": {
										Description: "Base64 encoded CA certificates of the https source to check against.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"source_disk_to_clone": {
										Description: "Source disk name from where the clone is done.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"source_file_path": {
										Description: "Disk image source for the virtual machine.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"virtual_disk_reference": {
							Description: "Name of the existing virtual disk to be attached to the Virtual Machine.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"force_delete": {
				Description: "Normally any virtual machine that is still powered on cannot be deleted. The expected sequence from a user is to first power off the virtual machine and then invoke the delete operation. However, in special circumstances, the owner of the virtual machine may know very well that the virtual machine is no longer needed and just wants to dispose it off. In such situations a delete operation of a virtual machine object is accepted only when this forceDelete attribute is set to true. Under normal circumstances (forceDelete is false), delete operation first confirms that the virtual machine is powered off and then proceeds to delete the virtual machine.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"guest_os": {
				Description: "Guest operating system running on virtual machine.\n* `linux` - A Linux operating system.\n* `windows` - A Windows operating system.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "linux",
			},
			"host": {
				Description: "A reference to a virtualizationBaseHost resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
				ForceNew: true,
			},
			"host_esxi": {
				Description: "Host where virtual machine is deployed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.\n* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"interfaces": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adaptor_type": {
							Description: "Virtual machine network adaptor type.\n* `Unknown` - The type of the network adaptor type is unknown.\n* `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC.\n* `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support.\n* `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.\n* `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features.\n* `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later.\n* `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets.\n* `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length.\n* `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance.\n* `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware.\n* `` - Default network adaptor type supported by the hypervisor.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Unknown",
						},
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"bridge": {
							Description: "Virtual machine network bridge name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.NetworkInterface",
						},
						"connect_at_power_on": {
							Description: "Connect the adaptor at virtual machine power on.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"direct_path_io": {
							Description: "Enable the direct path I/O.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"ip_forwarding_enabled": {
							Description: "Set to true, if IP forwarding is enabled on the NIC.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"ipv6_address": {
							Description: "Set to true, if IPv6 address should be allocated for the NIC.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"mac_address": {
							Description: "Virtual machine network mac address.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Name of the network interface. This may be different from guest operating system assigned.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"network_id": {
							Description: "Identity of the network to which this network interface belongs.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nic_id": {
							Description: "Identity of the network interface.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "virtualization.NetworkInterface",
						},
						"order": {
							Description: "Order of the NIC attachment to the VM.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"private_ip_allocation_mode": {
							Description: "Allocation mode for NIC addresses e.g. DHCP or static.\n* `DHCP` - Dynamic IP address allocation using DHCP protocol.\n* `STATIC_IP` - Assign fixed / static IPs to resources for use.\n* `IPAM_CALLOUT` - Use callout scripts to query cloud IP allocation tools to assign network parameters.\n* `PREALLOCATE_IP` - Allows the cloud infrastructure IP allocation to be dynamically provided before the server boots up.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "DHCP",
						},
						"public_ip_allocate": {
							Description: "Set to true, if public IP should be allocated for the NIC.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"security_groups": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type: schema.TypeString}},
						"static_ip_address": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "virtualization.IpAddressInfo",
									},
									"gateway_ip": {
										Description: "IP address of the device on network which forwards local traffic to other networks.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"ip_address": {
										Description: "An IP address is a 32-bit number. It uniquely identifies a host in given network.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "virtualization.IpAddressInfo",
									},
									"subnet_mask": {
										Description: "A 32 bit number which helps to identify the host and rest of the network.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"subnet_id": {
							Description: "Subnet identifier for the NIC.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"inventory": {
				Description: "A reference to a virtualizationBaseVirtualMachine resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"labels": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "infra.MetaData",
						},
						"name": {
							Description: "Name of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "infra.MetaData",
						},
						"value": {
							Description: "Value of the meta property which identifies a specific resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"memory": {
				Description: "Virtual machine memory in mebi bytes (one mebibyte, 1MiB, is 1048576 bytes, and 1KiB is 1024 bytes). Input must be a whole number and scientific notation is not acceptable. For example, enter 1730 and not 1.73e03.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Virtual machine name that is unique. Hypervisors enforce platform specific limits and character sets. The name length limit, both min and max, vary among hypervisors. Therefore, the basic limits are set here and proper enforcement is done elsewhere.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "virtualization.VirtualMachine",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"power_state": {
				Description: "Expected power state of virtual machine (PowerOn, PowerOff, Restart).\n* `PowerOff` - The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected.\n* `PowerOn` - The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected.\n* `Suspend` - The virtual machine will be put into  a suspended state.\n* `ShutDownGuestOS` - The guest operating system is shut down gracefully.\n* `RestartGuestOS` - It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts.\n* `Reset` - Resets the virtual machine abruptly, with no consideration for work in progress.\n* `Restart` - The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up.\n* `Unknown` - Power state of the entity is unknown.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PowerOff",
			},
			"provision_type": {
				Description: "Identifies the provision type to create a new virtual machine.\n* `OVA` - Deploy virtual machine using OVA/F file.\n* `Template` - Provision virtual machine using a template file.\n* `Discovered` - A virtual machine was 'discovered' and not created from Intersight. No provisioning information is available.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "OVA",
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
				ForceNew: true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"timestamp": {
							Description: "The time this versioned Managed Object was created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"vm_config": {
				Description: "Virtual machine configuration to provision.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"workflow_info": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"wait_for_completion": {
				Description: "This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state.",
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
			}},
	}
}

func resourceVirtualizationVirtualMachineCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewVirtualizationVirtualMachineWithDefaults()

	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("affinity_selectors"); ok {
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewInfraMetaDataWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetAffinitySelectors(x)
		}
	}

	if v, ok := d.GetOk("anti_affinity_selectors"); ok {
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewInfraMetaDataWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetAntiAffinitySelectors(x)
		}
	}

	o.SetClassId("virtualization.VirtualMachine")

	if v, ok := d.GetOk("cloud_init_config"); ok {
		p := make([]models.VirtualizationCloudInitConfig, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewVirtualizationCloudInitConfigWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["config_type"]; ok {
				{
					x := (v.(string))
					o.SetConfigType(x)
				}
			}
			if v, ok := l["network_data"]; ok {
				{
					x := (v.(string))
					o.SetNetworkData(x)
				}
			}
			if v, ok := l["network_data_base64_encoded"]; ok {
				{
					x := (v.(bool))
					o.SetNetworkDataBase64Encoded(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["user_data"]; ok {
				{
					x := (v.(string))
					o.SetUserData(x)
				}
			}
			if v, ok := l["user_data_base64_encoded"]; ok {
				{
					x := (v.(bool))
					o.SetUserDataBase64Encoded(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCloudInitConfig(x)
		}
	}

	if v, ok := d.GetOk("cluster"); ok {
		p := make([]models.VirtualizationBaseClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVirtualizationBaseClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCluster(x)
		}
	}

	if v, ok := d.GetOk("cluster_esxi"); ok {
		x := (v.(string))
		o.SetClusterEsxi(x)
	}

	if v, ok := d.GetOkExists("cpu"); ok {
		x := int64(v.(int))
		o.SetCpu(x)
	}

	if v, ok := d.GetOk("disk"); ok {
		x := make([]models.VirtualizationVirtualMachineDisk, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewVirtualizationVirtualMachineDiskWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["bus"]; ok {
				{
					x := (v.(string))
					o.SetBus(x)
				}
			}
			o.SetClassId("virtualization.VirtualMachineDisk")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["order"]; ok {
				{
					x := int64(v.(int))
					o.SetOrder(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			if v, ok := l["virtual_disk"]; ok {
				{
					p := make([]models.VirtualizationVirtualDiskConfig, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewVirtualizationVirtualDiskConfigWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["capacity"]; ok {
							{
								x := (v.(string))
								o.SetCapacity(x)
							}
						}
						o.SetClassId("")
						if v, ok := l["mode"]; ok {
							{
								x := (v.(string))
								o.SetMode(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["source_certs"]; ok {
							{
								x := (v.(string))
								o.SetSourceCerts(x)
							}
						}
						if v, ok := l["source_disk_to_clone"]; ok {
							{
								x := (v.(string))
								o.SetSourceDiskToClone(x)
							}
						}
						if v, ok := l["source_file_path"]; ok {
							{
								x := (v.(string))
								o.SetSourceFilePath(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetVirtualDisk(x)
					}
				}
			}
			if v, ok := l["virtual_disk_reference"]; ok {
				{
					x := (v.(string))
					o.SetVirtualDiskReference(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetDisk(x)
		}
	}

	if v, ok := d.GetOkExists("force_delete"); ok {
		x := (v.(bool))
		o.SetForceDelete(x)
	}

	if v, ok := d.GetOk("guest_os"); ok {
		x := (v.(string))
		o.SetGuestOs(x)
	}

	if v, ok := d.GetOk("host"); ok {
		p := make([]models.VirtualizationBaseHostRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVirtualizationBaseHostRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetHost(x)
		}
	}

	if v, ok := d.GetOk("host_esxi"); ok {
		x := (v.(string))
		o.SetHostEsxi(x)
	}

	if v, ok := d.GetOk("interfaces"); ok {
		x := make([]models.VirtualizationNetworkInterface, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewVirtualizationNetworkInterfaceWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["adaptor_type"]; ok {
				{
					x := (v.(string))
					o.SetAdaptorType(x)
				}
			}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["bridge"]; ok {
				{
					x := (v.(string))
					o.SetBridge(x)
				}
			}
			o.SetClassId("virtualization.NetworkInterface")
			if v, ok := l["connect_at_power_on"]; ok {
				{
					x := (v.(bool))
					o.SetConnectAtPowerOn(x)
				}
			}
			if v, ok := l["direct_path_io"]; ok {
				{
					x := (v.(bool))
					o.SetDirectPathIo(x)
				}
			}
			if v, ok := l["ip_forwarding_enabled"]; ok {
				{
					x := (v.(bool))
					o.SetIpForwardingEnabled(x)
				}
			}
			if v, ok := l["ipv6_address"]; ok {
				{
					x := (v.(bool))
					o.SetIpv6Address(x)
				}
			}
			if v, ok := l["mac_address"]; ok {
				{
					x := (v.(string))
					o.SetMacAddress(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["network_id"]; ok {
				{
					x := (v.(string))
					o.SetNetworkId(x)
				}
			}
			if v, ok := l["nic_id"]; ok {
				{
					x := (v.(string))
					o.SetNicId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["order"]; ok {
				{
					x := int64(v.(int))
					o.SetOrder(x)
				}
			}
			if v, ok := l["private_ip_allocation_mode"]; ok {
				{
					x := (v.(string))
					o.SetPrivateIpAllocationMode(x)
				}
			}
			if v, ok := l["public_ip_allocate"]; ok {
				{
					x := (v.(bool))
					o.SetPublicIpAllocate(x)
				}
			}
			if v, ok := l["security_groups"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSecurityGroups(x)
					}
				}
			}
			if v, ok := l["static_ip_address"]; ok {
				{
					x := make([]models.VirtualizationIpAddressInfo, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewVirtualizationIpAddressInfoWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("virtualization.IpAddressInfo")
						if v, ok := l["gateway_ip"]; ok {
							{
								x := (v.(string))
								o.SetGatewayIp(x)
							}
						}
						if v, ok := l["ip_address"]; ok {
							{
								x := (v.(string))
								o.SetIpAddress(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["subnet_mask"]; ok {
							{
								x := (v.(string))
								o.SetSubnetMask(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetStaticIpAddress(x)
					}
				}
			}
			if v, ok := l["subnet_id"]; ok {
				{
					x := (v.(string))
					o.SetSubnetId(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetInterfaces(x)
		}
	}

	if v, ok := d.GetOk("labels"); ok {
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewInfraMetaDataWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetLabels(x)
		}
	}

	if v, ok := d.GetOkExists("memory"); ok {
		x := int64(v.(int))
		o.SetMemory(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("virtualization.VirtualMachine")

	if v, ok := d.GetOk("power_state"); ok {
		x := (v.(string))
		o.SetPowerState(x)
	}

	if v, ok := d.GetOk("provision_type"); ok {
		x := (v.(string))
		o.SetProvisionType(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("vm_config"); ok {
		p := make([]models.VirtualizationBaseVmConfiguration, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewVirtualizationBaseVmConfigurationWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmConfig(x)
		}
	}

	r := conn.ApiClient.VirtualizationApi.CreateVirtualizationVirtualMachine(conn.ctx).VirtualizationVirtualMachine(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating VirtualizationVirtualMachine: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	var waitForCompletion bool
	if v, ok := d.GetOk("wait_for_completion"); ok {
		waitForCompletion = v.(bool)
	}
	// Check for Workflow Status
	if waitForCompletion {
		for i := 0; i < 4; i += 1 {
			resultMo, _, responseErr = conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineByMoid(conn.ctx, resultMo.GetMoid()).Execute()
			if responseErr != nil {
				errorType := fmt.Sprintf("%T", responseErr)
				if strings.Contains(errorType, "GenericOpenAPIError") {
					responseErr := responseErr.(models.GenericOpenAPIError)
					return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
				}
				return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s", responseErr.Error())
			}
			if _, ok := resultMo.GetWorkflowInfoOk(); ok {
				log.Println("Workflow details found")
				break
			}
		}
		resultMo, _, responseErr = conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineByMoid(conn.ctx, resultMo.GetMoid()).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s", responseErr.Error())
		}
		var runningWorkflows []models.WorkflowWorkflowInfoRelationship
		if _, ok := resultMo.GetWorkflowInfoOk(); ok {
			runningWorkflows = append(runningWorkflows, resultMo.GetWorkflowInfo())
		}
		for _, w := range runningWorkflows {
			warning, err := checkWorkflowStatus(conn, w)
			if err != nil {
				errorType := fmt.Sprintf("%T", err)
				if strings.Contains(errorType, "GenericOpenAPIError") {
					err := err.(models.GenericOpenAPIError)
					return diag.Errorf("failed while fetching workflow information in VirtualizationVirtualMachine: %s Response from endpoint: %s", err.Error(), string(err.Body()))
				}
				return diag.Errorf("failed while fetching workflow information in VirtualizationVirtualMachine: %s", err.Error())
			}
			if len(warning) > 0 {
				de = append(de, diag.Diagnostic{Severity: diag.Warning, Summary: warning})
			}
		}
	}
	return append(de, resourceVirtualizationVirtualMachineRead(c, d, meta)...)
}

func resourceVirtualizationVirtualMachineRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	r := conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "VirtualizationVirtualMachine object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("action", (s.GetAction())); err != nil {
		return diag.Errorf("error occurred while setting property Action in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("action_info", flattenMapVirtualizationActionInfo(s.GetActionInfo(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ActionInfo in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("affinity_selectors", flattenListInfraMetaData(s.GetAffinitySelectors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property AffinitySelectors in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("anti_affinity_selectors", flattenListInfraMetaData(s.GetAntiAffinitySelectors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property AntiAffinitySelectors in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("cloud_init_config", flattenMapVirtualizationCloudInitConfig(s.GetCloudInitConfig(), d)); err != nil {
		return diag.Errorf("error occurred while setting property CloudInitConfig in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("cluster", flattenMapVirtualizationBaseClusterRelationship(s.GetCluster(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Cluster in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("cluster_esxi", (s.GetClusterEsxi())); err != nil {
		return diag.Errorf("error occurred while setting property ClusterEsxi in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("cpu", (s.GetCpu())); err != nil {
		return diag.Errorf("error occurred while setting property Cpu in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("discovered", (s.GetDiscovered())); err != nil {
		return diag.Errorf("error occurred while setting property Discovered in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("disk", flattenListVirtualizationVirtualMachineDisk(s.GetDisk(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Disk in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("force_delete", (s.GetForceDelete())); err != nil {
		return diag.Errorf("error occurred while setting property ForceDelete in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("guest_os", (s.GetGuestOs())); err != nil {
		return diag.Errorf("error occurred while setting property GuestOs in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("host", flattenMapVirtualizationBaseHostRelationship(s.GetHost(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Host in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("host_esxi", (s.GetHostEsxi())); err != nil {
		return diag.Errorf("error occurred while setting property HostEsxi in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
		return diag.Errorf("error occurred while setting property HypervisorType in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("interfaces", flattenListVirtualizationNetworkInterface(s.GetInterfaces(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Interfaces in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("inventory", flattenMapVirtualizationBaseVirtualMachineRelationship(s.GetInventory(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Inventory in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("labels", flattenListInfraMetaData(s.GetLabels(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Labels in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("memory", (s.GetMemory())); err != nil {
		return diag.Errorf("error occurred while setting property Memory in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("power_state", (s.GetPowerState())); err != nil {
		return diag.Errorf("error occurred while setting property PowerState in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("provision_type", (s.GetProvisionType())); err != nil {
		return diag.Errorf("error occurred while setting property ProvisionType in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
		return diag.Errorf("error occurred while setting property RegisteredDevice in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("vm_config", flattenMapVirtualizationBaseVmConfiguration(s.GetVmConfig(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VmConfig in VirtualizationVirtualMachine object: %s", err.Error())
	}

	if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)); err != nil {
		return diag.Errorf("error occurred while setting property WorkflowInfo in VirtualizationVirtualMachine object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceVirtualizationVirtualMachineUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVirtualMachine{}

	if d.HasChange("action") {
		v := d.Get("action")
		x := (v.(string))
		o.SetAction(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("affinity_selectors") {
		v := d.Get("affinity_selectors")
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.InfraMetaData{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetAffinitySelectors(x)
	}

	if d.HasChange("anti_affinity_selectors") {
		v := d.Get("anti_affinity_selectors")
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.InfraMetaData{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetAntiAffinitySelectors(x)
	}

	o.SetClassId("virtualization.VirtualMachine")

	if d.HasChange("cloud_init_config") {
		v := d.Get("cloud_init_config")
		p := make([]models.VirtualizationCloudInitConfig, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VirtualizationCloudInitConfig{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["config_type"]; ok {
				{
					x := (v.(string))
					o.SetConfigType(x)
				}
			}
			if v, ok := l["network_data"]; ok {
				{
					x := (v.(string))
					o.SetNetworkData(x)
				}
			}
			if v, ok := l["network_data_base64_encoded"]; ok {
				{
					x := (v.(bool))
					o.SetNetworkDataBase64Encoded(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["user_data"]; ok {
				{
					x := (v.(string))
					o.SetUserData(x)
				}
			}
			if v, ok := l["user_data_base64_encoded"]; ok {
				{
					x := (v.(bool))
					o.SetUserDataBase64Encoded(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCloudInitConfig(x)
		}
	}

	if d.HasChange("cluster") {
		v := d.Get("cluster")
		p := make([]models.VirtualizationBaseClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVirtualizationBaseClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCluster(x)
		}
	}

	if d.HasChange("cluster_esxi") {
		v := d.Get("cluster_esxi")
		x := (v.(string))
		o.SetClusterEsxi(x)
	}

	if d.HasChange("cpu") {
		v := d.Get("cpu")
		x := int64(v.(int))
		o.SetCpu(x)
	}

	if d.HasChange("disk") {
		v := d.Get("disk")
		x := make([]models.VirtualizationVirtualMachineDisk, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.VirtualizationVirtualMachineDisk{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["bus"]; ok {
				{
					x := (v.(string))
					o.SetBus(x)
				}
			}
			o.SetClassId("virtualization.VirtualMachineDisk")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["order"]; ok {
				{
					x := int64(v.(int))
					o.SetOrder(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					x := (v.(string))
					o.SetType(x)
				}
			}
			if v, ok := l["virtual_disk"]; ok {
				{
					p := make([]models.VirtualizationVirtualDiskConfig, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewVirtualizationVirtualDiskConfigWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						if v, ok := l["capacity"]; ok {
							{
								x := (v.(string))
								o.SetCapacity(x)
							}
						}
						o.SetClassId("")
						if v, ok := l["mode"]; ok {
							{
								x := (v.(string))
								o.SetMode(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["source_certs"]; ok {
							{
								x := (v.(string))
								o.SetSourceCerts(x)
							}
						}
						if v, ok := l["source_disk_to_clone"]; ok {
							{
								x := (v.(string))
								o.SetSourceDiskToClone(x)
							}
						}
						if v, ok := l["source_file_path"]; ok {
							{
								x := (v.(string))
								o.SetSourceFilePath(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetVirtualDisk(x)
					}
				}
			}
			if v, ok := l["virtual_disk_reference"]; ok {
				{
					x := (v.(string))
					o.SetVirtualDiskReference(x)
				}
			}
			x = append(x, *o)
		}
		o.SetDisk(x)
	}

	if d.HasChange("force_delete") {
		v := d.Get("force_delete")
		x := (v.(bool))
		o.SetForceDelete(x)
	}

	if d.HasChange("guest_os") {
		v := d.Get("guest_os")
		x := (v.(string))
		o.SetGuestOs(x)
	}

	if d.HasChange("host") {
		v := d.Get("host")
		p := make([]models.VirtualizationBaseHostRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVirtualizationBaseHostRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetHost(x)
		}
	}

	if d.HasChange("host_esxi") {
		v := d.Get("host_esxi")
		x := (v.(string))
		o.SetHostEsxi(x)
	}

	if d.HasChange("interfaces") {
		v := d.Get("interfaces")
		x := make([]models.VirtualizationNetworkInterface, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.VirtualizationNetworkInterface{}
			l := s[i].(map[string]interface{})
			if v, ok := l["adaptor_type"]; ok {
				{
					x := (v.(string))
					o.SetAdaptorType(x)
				}
			}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["bridge"]; ok {
				{
					x := (v.(string))
					o.SetBridge(x)
				}
			}
			o.SetClassId("virtualization.NetworkInterface")
			if v, ok := l["connect_at_power_on"]; ok {
				{
					x := (v.(bool))
					o.SetConnectAtPowerOn(x)
				}
			}
			if v, ok := l["direct_path_io"]; ok {
				{
					x := (v.(bool))
					o.SetDirectPathIo(x)
				}
			}
			if v, ok := l["ip_forwarding_enabled"]; ok {
				{
					x := (v.(bool))
					o.SetIpForwardingEnabled(x)
				}
			}
			if v, ok := l["ipv6_address"]; ok {
				{
					x := (v.(bool))
					o.SetIpv6Address(x)
				}
			}
			if v, ok := l["mac_address"]; ok {
				{
					x := (v.(string))
					o.SetMacAddress(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["network_id"]; ok {
				{
					x := (v.(string))
					o.SetNetworkId(x)
				}
			}
			if v, ok := l["nic_id"]; ok {
				{
					x := (v.(string))
					o.SetNicId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["order"]; ok {
				{
					x := int64(v.(int))
					o.SetOrder(x)
				}
			}
			if v, ok := l["private_ip_allocation_mode"]; ok {
				{
					x := (v.(string))
					o.SetPrivateIpAllocationMode(x)
				}
			}
			if v, ok := l["public_ip_allocate"]; ok {
				{
					x := (v.(bool))
					o.SetPublicIpAllocate(x)
				}
			}
			if v, ok := l["security_groups"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSecurityGroups(x)
					}
				}
			}
			if v, ok := l["static_ip_address"]; ok {
				{
					x := make([]models.VirtualizationIpAddressInfo, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewVirtualizationIpAddressInfoWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("virtualization.IpAddressInfo")
						if v, ok := l["gateway_ip"]; ok {
							{
								x := (v.(string))
								o.SetGatewayIp(x)
							}
						}
						if v, ok := l["ip_address"]; ok {
							{
								x := (v.(string))
								o.SetIpAddress(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["subnet_mask"]; ok {
							{
								x := (v.(string))
								o.SetSubnetMask(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetStaticIpAddress(x)
					}
				}
			}
			if v, ok := l["subnet_id"]; ok {
				{
					x := (v.(string))
					o.SetSubnetId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetInterfaces(x)
	}

	if d.HasChange("labels") {
		v := d.Get("labels")
		x := make([]models.InfraMetaData, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.InfraMetaData{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("infra.MetaData")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetLabels(x)
	}

	if d.HasChange("memory") {
		v := d.Get("memory")
		x := int64(v.(int))
		o.SetMemory(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("virtualization.VirtualMachine")

	if d.HasChange("power_state") {
		v := d.Get("power_state")
		x := (v.(string))
		o.SetPowerState(x)
	}

	if d.HasChange("provision_type") {
		v := d.Get("provision_type")
		x := (v.(string))
		o.SetProvisionType(x)
	}

	if d.HasChange("registered_device") {
		v := d.Get("registered_device")
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if d.HasChange("vm_config") {
		v := d.Get("vm_config")
		p := make([]models.VirtualizationBaseVmConfiguration, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VirtualizationBaseVmConfiguration{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmConfig(x)
		}
	}

	r := conn.ApiClient.VirtualizationApi.UpdateVirtualizationVirtualMachine(conn.ctx, d.Id()).VirtualizationVirtualMachine(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating VirtualizationVirtualMachine: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	var waitForCompletion bool
	if v, ok := d.GetOk("wait_for_completion"); ok {
		waitForCompletion = v.(bool)
	}
	// Check for Workflow Status
	if waitForCompletion {
		for i := 0; i < 4; i += 1 {
			result, _, responseErr = conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineByMoid(conn.ctx, result.GetMoid()).Execute()
			if responseErr != nil {
				errorType := fmt.Sprintf("%T", responseErr)
				if strings.Contains(errorType, "GenericOpenAPIError") {
					responseErr := responseErr.(models.GenericOpenAPIError)
					return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
				}
				return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s", responseErr.Error())
			}
			if _, ok := result.GetWorkflowInfoOk(); ok {
				log.Println("Workflow details found")
				break
			}
		}
		result, _, responseErr = conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineByMoid(conn.ctx, result.GetMoid()).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s", responseErr.Error())
		}
		var runningWorkflows []models.WorkflowWorkflowInfoRelationship
		if _, ok := result.GetWorkflowInfoOk(); ok {
			runningWorkflows = append(runningWorkflows, result.GetWorkflowInfo())
		}
		for _, w := range runningWorkflows {
			warning, err := checkWorkflowStatus(conn, w)
			if err != nil {
				errorType := fmt.Sprintf("%T", err)
				if strings.Contains(errorType, "GenericOpenAPIError") {
					err := err.(models.GenericOpenAPIError)
					return diag.Errorf("failed while fetching workflow information in VirtualizationVirtualMachine: %s Response from endpoint: %s", err.Error(), string(err.Body()))
				}
				return diag.Errorf("failed while fetching workflow information in VirtualizationVirtualMachine: %s", err.Error())
			}
			if len(warning) > 0 {
				de = append(de, diag.Diagnostic{Severity: diag.Warning, Summary: warning})
			}
		}
	}
	return append(de, resourceVirtualizationVirtualMachineRead(c, d, meta)...)
}

func resourceVirtualizationVirtualMachineDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.VirtualizationApi.DeleteVirtualizationVirtualMachine(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "VirtualizationVirtualMachineDelete: VirtualizationVirtualMachine object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting VirtualizationVirtualMachine object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting VirtualizationVirtualMachine object: %s", deleteErr.Error())
	}
	return de
}
