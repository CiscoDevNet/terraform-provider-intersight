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

func dataSourceVirtualizationVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualizationVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"action": {
				Description: "Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc).\n* `None` - A place holder for the default value.\n* `PowerState` - Power action is performed on the virtual machine.\n* `Migrate` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.\n* `Create` - The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation.\n* `Delete` - The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"action_info": {
				Description: "Details of an action performed on the virtul machine. Contains name of the action performed, status, failure reason message etc.",
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
						"failure_reason": {
							Description: "Description of reason for failure. Derived from the workflow failure message.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the Action performed on a resource like Virtual Machine, Disk etc.",
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
						"status": {
							Description: "Status of the Action like InProgress, Success, Failure etc.\n* `None` - A place holder for the default value.\n* `InProgress` - Previous action triggered on the resource is still running.\n* `Success` - Previous action triggered on the resource has completed successfully.\n* `Failure` - Previous action triggered on the resource has failed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cloud_init_config": {
				Description: "Cloud init configuration data for virtual machine.",
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
						"config_type": {
							Description: "Virtual machine cloud init configuration type.\n* `` - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected.\n* `NoCloudSource` - Allows the user to provide user-data to the instance without running a network service.\n* `CloudConfigDrive` - Allows the user to provide user-data and network-data from cloud.",
							Type:        schema.TypeString,
							Optional:    true,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
				Computed: true,
			},
			"cluster": {
				Description: "A reference to a virtualizationBaseCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"discovered": {
				Description: "Flag to indicate whether the configuration is created from inventory object.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"disk": {
				Type:     schema.TypeList,
				Optional: true,
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
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Virtual machine network bridge name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
						},
						"virtual_disk": {
							Description: "Virtual disk configuration.",
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
										Description: "Disk capacity to be provided with units example - 10Gi.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"mode": {
										Description: "File mode of the disk, example - Filesystem, Block.\n* `Block` - It is a Block virtual disk.\n* `Filesystem` - It is a File system virtual disk.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Name of the virtual disk.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
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
							Computed: true,
						},
						"virtual_disk_reference": {
							Description: "Name of the existing virtual disk to be attached to the Virtual Machine.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"guest_os": {
				Description: "Guest operating system running on virtual machine.\n* `linux` - A Linux operating system.\n* `windows` - A Windows operating system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host": {
				Description: "A reference to a virtualizationBaseHost resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"host_esxi": {
				Description: "Host where virtual machine is deployed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adaptor_type": {
							Description: "Virtual machine network adaptor type.\n* `Unknown` - The type of the network adaptor type is unknown.\n* `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC.\n* `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support.\n* `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.\n* `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features.",
							Type:        schema.TypeString,
							Optional:    true,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
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
						"mac_address": {
							Description: "Virtual machine network mac address.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"inventory": {
				Description: "A reference to a virtualizationBaseVirtualMachine resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"memory": {
				Description: "Virtual machine memory defined in mega bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Virtual machine name contains only letters, numbers, allowed special character and must be unique.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"power_state": {
				Description: "Expected power state of virtual machine (PowerOn, PowerOff, Restart).\n* `PowerOff` - The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected.\n* `PowerOn` - The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected.\n* `Suspend` - The virtual machine will be put into  a suspended state.\n* `ShutDownGuestOS` - The guest operating system is shut down gracefully.\n* `RestartGuestOS` - It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts.\n* `Reset` - Resets the virtual machine abruptly, with no consideration for work in progress.\n* `Restart` - The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up.\n* `Unknown` - Power state of the entity is unknown.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"provision_type": {
				Description: "Identifies the provision type to create a new virtual machine.\n* `OVA` - Deploy virtual machine using OVA/F file.\n* `Template` - Provision virtual machine using a template file.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"vm_config": {
				Description: "Virtual machine configuration to provision.",
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
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"workflow_info": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		},
	}
}

func dataSourceVirtualizationVirtualMachineRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVirtualMachine{}
	if v, ok := d.GetOk("action"); ok {
		x := (v.(string))
		o.SetAction(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_esxi"); ok {
		x := (v.(string))
		o.SetClusterEsxi(x)
	}
	if v, ok := d.GetOk("cpu"); ok {
		x := int64(v.(int))
		o.SetCpu(x)
	}
	if v, ok := d.GetOk("discovered"); ok {
		x := (v.(bool))
		o.SetDiscovered(x)
	}
	if v, ok := d.GetOk("guest_os"); ok {
		x := (v.(string))
		o.SetGuestOs(x)
	}
	if v, ok := d.GetOk("host_esxi"); ok {
		x := (v.(string))
		o.SetHostEsxi(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("memory"); ok {
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("power_state"); ok {
		x := (v.(string))
		o.SetPowerState(x)
	}
	if v, ok := d.GetOk("provision_type"); ok {
		x := (v.(string))
		o.SetProvisionType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VirtualizationVirtualMachine object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVirtualMachineList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching VirtualizationVirtualMachine: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for VirtualizationVirtualMachine list: %s", err.Error())
	}
	var s = &models.VirtualizationVirtualMachineList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to VirtualizationVirtualMachine list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for VirtualizationVirtualMachine data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for VirtualizationVirtualMachine data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.VirtualizationVirtualMachine{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("action", (s.GetAction())); err != nil {
				return diag.Errorf("error occurred while setting property Action: %s", err.Error())
			}

			if err := d.Set("action_info", flattenMapVirtualizationActionInfo(s.GetActionInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ActionInfo: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("affinity_selectors", flattenListInfraMetaData(s.GetAffinitySelectors(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AffinitySelectors: %s", err.Error())
			}

			if err := d.Set("anti_affinity_selectors", flattenListInfraMetaData(s.GetAntiAffinitySelectors(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AntiAffinitySelectors: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cloud_init_config", flattenMapVirtualizationCloudInitConfig(s.GetCloudInitConfig(), d)); err != nil {
				return diag.Errorf("error occurred while setting property CloudInitConfig: %s", err.Error())
			}

			if err := d.Set("cluster", flattenMapVirtualizationBaseClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}
			if err := d.Set("cluster_esxi", (s.GetClusterEsxi())); err != nil {
				return diag.Errorf("error occurred while setting property ClusterEsxi: %s", err.Error())
			}
			if err := d.Set("cpu", (s.GetCpu())); err != nil {
				return diag.Errorf("error occurred while setting property Cpu: %s", err.Error())
			}
			if err := d.Set("discovered", (s.GetDiscovered())); err != nil {
				return diag.Errorf("error occurred while setting property Discovered: %s", err.Error())
			}

			if err := d.Set("disk", flattenListVirtualizationVirtualMachineDisk(s.GetDisk(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Disk: %s", err.Error())
			}
			if err := d.Set("guest_os", (s.GetGuestOs())); err != nil {
				return diag.Errorf("error occurred while setting property GuestOs: %s", err.Error())
			}

			if err := d.Set("host", flattenMapVirtualizationBaseHostRelationship(s.GetHost(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Host: %s", err.Error())
			}
			if err := d.Set("host_esxi", (s.GetHostEsxi())); err != nil {
				return diag.Errorf("error occurred while setting property HostEsxi: %s", err.Error())
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorType: %s", err.Error())
			}

			if err := d.Set("interfaces", flattenListVirtualizationNetworkInterface(s.GetInterfaces(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Interfaces: %s", err.Error())
			}

			if err := d.Set("inventory", flattenMapVirtualizationBaseVirtualMachineRelationship(s.GetInventory(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Inventory: %s", err.Error())
			}

			if err := d.Set("labels", flattenListInfraMetaData(s.GetLabels(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Labels: %s", err.Error())
			}
			if err := d.Set("memory", (s.GetMemory())); err != nil {
				return diag.Errorf("error occurred while setting property Memory: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("power_state", (s.GetPowerState())); err != nil {
				return diag.Errorf("error occurred while setting property PowerState: %s", err.Error())
			}
			if err := d.Set("provision_type", (s.GetProvisionType())); err != nil {
				return diag.Errorf("error occurred while setting property ProvisionType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("vm_config", flattenMapVirtualizationBaseVmConfiguration(s.GetVmConfig(), d)); err != nil {
				return diag.Errorf("error occurred while setting property VmConfig: %s", err.Error())
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowInfo: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
