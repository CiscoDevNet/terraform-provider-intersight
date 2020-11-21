package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVirtualizationVmwareVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVirtualizationVmwareVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"annotation": {
				Description: "List of annotations provided to this VM by user. Can be long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_time": {
				Description: "Time when this VM booted up.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster": {
				Description: "A reference to a virtualizationVmwareCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"config_name": {
				Description: "The configuration name for this VM. This maybe the same as the guest hostname.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connection_state": {
				Description: "Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_hot_add_enabled": {
				Description: "Indicates if the capability to add CPUs to a running VM is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"cpu_shares": {
				Description: "Shows the relative importance of a VM and its CPU limits.",
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
						"cpu_limit": {
							Description: "Upper limit on CPU allocation (MHz).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_overhead_limit": {
							Description: "Amount of CPU for virtualization overhead.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_reservation": {
							Description: "Guaranteed minimum allocation of CPU resource (MHz).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"cpu_shares": {
							Description: "Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.",
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
			"cpu_socket_info": {
				Description: "Details of CPUs/sockets of this VM.",
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
						"cores_per_socket": {
							Description: "The number of core per CPU socket (value may be empty).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"num_cpus": {
							Description: "Number of CPUs allocated to this VM.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"num_sockets": {
							Description: "The number of CPU sockets allocated.",
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
			"custom_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"datacenter": {
				Description: "A reference to a virtualizationVmwareDatacenter resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"datastores": {
				Description: "An array of relationships to virtualizationVmwareDatastore resources.",
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
			"default_power_off_type": {
				Description: "Indicates how the VM will be powered off (soft, hard etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dhcp_enabled": {
				Description: "Shows if DHCP is used for IP/DNS on this VM.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"disk_commit_info": {
				Description: "Information about the virtual machine's disk commits, sharing and limits.",
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
						"committed_disk": {
							Description: "Disk committed in bytes on this virtual machine (disk space used up).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"un_committed_disk": {
							Description: "Total uncommitted disk space that is available for use (in bytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"unshared_disk": {
							Description: "Total unshared disk space (in bytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"dns_server_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"dns_suffix_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"folder": {
				Description: "The folder name associated with this VM.",
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
			"guest_state": {
				Description: "The state of the guest OS running on this VM. Could be running, not running etc.\n* `Unknown` - Indicates that the guest OS state cannot be determined.\n* `NotRunning` - Indicates that the guest OS is not running.\n* `Resetting` - Indicates that the guest OS is resetting.\n* `Running` - Indicates that the guest OS is running normally.\n* `ShuttingDown` - Indicates that the guest OS is shutting down.\n* `Standby` - Indicates that the guest OS is in standby mode.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host": {
				Description: "A reference to a virtualizationVmwareHost resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Description: "Type of hypervisor where the virtual machine is hosted for example ESXi.\n* `ESXi` - A Vmware ESXi hypervisor of any version.\n* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"instance_uuid": {
				Description: "UUID assigned by vCenter to every VM.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"is_template": {
				Description: "If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"mac_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"mem_shares": {
				Description: "Similar to CPU Shares but applicable to memory.",
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
						"mem_limit": {
							Description: "Limit on the memory sharing imposed (in Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_overhead_limit": {
							Description: "Limit on memory overhead imposed (in Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_reservation": {
							Description: "Similar to CPU reservations (Mbytes).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"mem_shares": {
							Description: "Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.",
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
			"memory_hot_add_enabled": {
				Description: "Adding memory to a running VM.",
				Type:        schema.TypeBool,
				Optional:    true,
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
				Description: "Indicates how many networks are used by this VM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
			"protected_vm": {
				Description: "Shows if this is a protected VM. VMs can be in protection groups.",
				Type:        schema.TypeBool,
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
			"remote_display_info": {
				Description: "Applies only when remoteDisplayvnc is enabled.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"remote_display_password": {
							Description: "The password used for remote access. It is stored base64 encoded.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_display_vnc_key": {
							Description: "The access key for the remote display, potentially a long string.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_display_vnc_port": {
							Description: "Applies only when remoteDisplayvnc is enabled.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"remote_display_vnc_enabled": {
				Description: "Shows if support for a remote VNC access is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"resource_pool": {
				Description: "Name of the resource pool to which this VM belongs (optional).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_pool_owner": {
				Description: "Who owns the resource pool.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_pool_parent": {
				Description: "The parent of the current resource pool to which this VM belongs.",
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
			"tool_running_status": {
				Description: "Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tools_version": {
				Description: "The version of the guest tools, usually not specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "The uuid of this virtual machine. The uuid is internally generated and not user specified.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_disk_count": {
				Description: "Shows the number of disks assigned to this VM.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vm_overall_status": {
				Description: "The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_path": {
				Description: "Example - [datastore3] VCSA-134/VCSA-134.vmx.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_version": {
				Description: "Information about the version of this VM (vmx-09, vmx-11 etc.).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vm_vnic_count": {
				Description: "How many vnics are present.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vnic_device_config_id": {
				Description: "Information related to the guest info's VNIC virtual device. It is a comma-separated list.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceVirtualizationVmwareVirtualMachineRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.VirtualizationVmwareVirtualMachine{}
	if v, ok := d.GetOk("annotation"); ok {
		x := (v.(string))
		o.SetAnnotation(x)
	}
	if v, ok := d.GetOk("boot_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetBootTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("config_name"); ok {
		x := (v.(string))
		o.SetConfigName(x)
	}
	if v, ok := d.GetOk("connection_state"); ok {
		x := (v.(string))
		o.SetConnectionState(x)
	}
	if v, ok := d.GetOk("cpu_hot_add_enabled"); ok {
		x := (v.(bool))
		o.SetCpuHotAddEnabled(x)
	}
	if v, ok := d.GetOk("default_power_off_type"); ok {
		x := (v.(string))
		o.SetDefaultPowerOffType(x)
	}
	if v, ok := d.GetOk("dhcp_enabled"); ok {
		x := (v.(bool))
		o.SetDhcpEnabled(x)
	}
	if v, ok := d.GetOk("folder"); ok {
		x := (v.(string))
		o.SetFolder(x)
	}
	if v, ok := d.GetOk("guest_state"); ok {
		x := (v.(string))
		o.SetGuestState(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("instance_uuid"); ok {
		x := (v.(string))
		o.SetInstanceUuid(x)
	}
	if v, ok := d.GetOk("is_template"); ok {
		x := (v.(bool))
		o.SetIsTemplate(x)
	}
	if v, ok := d.GetOk("memory_hot_add_enabled"); ok {
		x := (v.(bool))
		o.SetMemoryHotAddEnabled(x)
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
	if v, ok := d.GetOk("protected_vm"); ok {
		x := (v.(bool))
		o.SetProtectedVm(x)
	}
	if v, ok := d.GetOk("remote_display_vnc_enabled"); ok {
		x := (v.(bool))
		o.SetRemoteDisplayVncEnabled(x)
	}
	if v, ok := d.GetOk("resource_pool"); ok {
		x := (v.(string))
		o.SetResourcePool(x)
	}
	if v, ok := d.GetOk("resource_pool_owner"); ok {
		x := (v.(string))
		o.SetResourcePoolOwner(x)
	}
	if v, ok := d.GetOk("resource_pool_parent"); ok {
		x := (v.(string))
		o.SetResourcePoolParent(x)
	}
	if v, ok := d.GetOk("tool_running_status"); ok {
		x := (v.(string))
		o.SetToolRunningStatus(x)
	}
	if v, ok := d.GetOk("tools_version"); ok {
		x := (v.(string))
		o.SetToolsVersion(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("vm_disk_count"); ok {
		x := int64(v.(int))
		o.SetVmDiskCount(x)
	}
	if v, ok := d.GetOk("vm_overall_status"); ok {
		x := (v.(string))
		o.SetVmOverallStatus(x)
	}
	if v, ok := d.GetOk("vm_path"); ok {
		x := (v.(string))
		o.SetVmPath(x)
	}
	if v, ok := d.GetOk("vm_version"); ok {
		x := (v.(string))
		o.SetVmVersion(x)
	}
	if v, ok := d.GetOk("vm_vnic_count"); ok {
		x := int64(v.(int))
		o.SetVmVnicCount(x)
	}
	if v, ok := d.GetOk("vnic_device_config_id"); ok {
		x := (v.(string))
		o.SetVnicDeviceConfigId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareVirtualMachineList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.VirtualizationVmwareVirtualMachineList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to VirtualizationVmwareVirtualMachine: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.VirtualizationVmwareVirtualMachine{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("annotation", (s.GetAnnotation())); err != nil {
				return fmt.Errorf("error occurred while setting property Annotation: %+v", err)
			}

			if err := d.Set("boot_time", (s.GetBootTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property BootTime: %+v", err)
			}

			if err := d.Set("capacity", flattenMapInfraHardwareInfo(s.GetCapacity(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Capacity: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("cluster", flattenMapVirtualizationVmwareClusterRelationship(s.GetCluster(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Cluster: %+v", err)
			}
			if err := d.Set("config_name", (s.GetConfigName())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigName: %+v", err)
			}
			if err := d.Set("connection_state", (s.GetConnectionState())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionState: %+v", err)
			}
			if err := d.Set("cpu_hot_add_enabled", (s.GetCpuHotAddEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuHotAddEnabled: %+v", err)
			}

			if err := d.Set("cpu_shares", flattenMapVirtualizationVmwareVmCpuShareInfo(s.GetCpuShares(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CpuShares: %+v", err)
			}

			if err := d.Set("cpu_socket_info", flattenMapVirtualizationVmwareVmCpuSocketInfo(s.GetCpuSocketInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CpuSocketInfo: %+v", err)
			}
			if err := d.Set("custom_attributes", (s.GetCustomAttributes())); err != nil {
				return fmt.Errorf("error occurred while setting property CustomAttributes: %+v", err)
			}

			if err := d.Set("datacenter", flattenMapVirtualizationVmwareDatacenterRelationship(s.GetDatacenter(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Datacenter: %+v", err)
			}

			if err := d.Set("datastores", flattenListVirtualizationVmwareDatastoreRelationship(s.GetDatastores(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Datastores: %+v", err)
			}
			if err := d.Set("default_power_off_type", (s.GetDefaultPowerOffType())); err != nil {
				return fmt.Errorf("error occurred while setting property DefaultPowerOffType: %+v", err)
			}
			if err := d.Set("dhcp_enabled", (s.GetDhcpEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property DhcpEnabled: %+v", err)
			}

			if err := d.Set("disk_commit_info", flattenMapVirtualizationVmwareVmDiskCommitInfo(s.GetDiskCommitInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property DiskCommitInfo: %+v", err)
			}
			if err := d.Set("dns_server_list", (s.GetDnsServerList())); err != nil {
				return fmt.Errorf("error occurred while setting property DnsServerList: %+v", err)
			}
			if err := d.Set("dns_suffix_list", (s.GetDnsSuffixList())); err != nil {
				return fmt.Errorf("error occurred while setting property DnsSuffixList: %+v", err)
			}
			if err := d.Set("folder", (s.GetFolder())); err != nil {
				return fmt.Errorf("error occurred while setting property Folder: %+v", err)
			}

			if err := d.Set("guest_info", flattenMapVirtualizationGuestInfo(s.GetGuestInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property GuestInfo: %+v", err)
			}
			if err := d.Set("guest_state", (s.GetGuestState())); err != nil {
				return fmt.Errorf("error occurred while setting property GuestState: %+v", err)
			}

			if err := d.Set("host", flattenMapVirtualizationVmwareHostRelationship(s.GetHost(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Host: %+v", err)
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return fmt.Errorf("error occurred while setting property HypervisorType: %+v", err)
			}
			if err := d.Set("identity", (s.GetIdentity())); err != nil {
				return fmt.Errorf("error occurred while setting property Identity: %+v", err)
			}
			if err := d.Set("instance_uuid", (s.GetInstanceUuid())); err != nil {
				return fmt.Errorf("error occurred while setting property InstanceUuid: %+v", err)
			}
			if err := d.Set("ip_address", (s.GetIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property IpAddress: %+v", err)
			}
			if err := d.Set("is_template", (s.GetIsTemplate())); err != nil {
				return fmt.Errorf("error occurred while setting property IsTemplate: %+v", err)
			}
			if err := d.Set("mac_address", (s.GetMacAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property MacAddress: %+v", err)
			}

			if err := d.Set("mem_shares", flattenMapVirtualizationVmwareVmMemoryShareInfo(s.GetMemShares(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property MemShares: %+v", err)
			}

			if err := d.Set("memory_capacity", flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property MemoryCapacity: %+v", err)
			}
			if err := d.Set("memory_hot_add_enabled", (s.GetMemoryHotAddEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property MemoryHotAddEnabled: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("network_count", (s.GetNetworkCount())); err != nil {
				return fmt.Errorf("error occurred while setting property NetworkCount: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("port_groups", (s.GetPortGroups())); err != nil {
				return fmt.Errorf("error occurred while setting property PortGroups: %+v", err)
			}
			if err := d.Set("power_state", (s.GetPowerState())); err != nil {
				return fmt.Errorf("error occurred while setting property PowerState: %+v", err)
			}

			if err := d.Set("processor_capacity", flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property ProcessorCapacity: %+v", err)
			}
			if err := d.Set("protected_vm", (s.GetProtectedVm())); err != nil {
				return fmt.Errorf("error occurred while setting property ProtectedVm: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}

			if err := d.Set("remote_display_info", flattenMapVirtualizationVmwareRemoteDisplayInfo(s.GetRemoteDisplayInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RemoteDisplayInfo: %+v", err)
			}
			if err := d.Set("remote_display_vnc_enabled", (s.GetRemoteDisplayVncEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property RemoteDisplayVncEnabled: %+v", err)
			}
			if err := d.Set("resource_pool", (s.GetResourcePool())); err != nil {
				return fmt.Errorf("error occurred while setting property ResourcePool: %+v", err)
			}
			if err := d.Set("resource_pool_owner", (s.GetResourcePoolOwner())); err != nil {
				return fmt.Errorf("error occurred while setting property ResourcePoolOwner: %+v", err)
			}
			if err := d.Set("resource_pool_parent", (s.GetResourcePoolParent())); err != nil {
				return fmt.Errorf("error occurred while setting property ResourcePoolParent: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("tool_running_status", (s.GetToolRunningStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property ToolRunningStatus: %+v", err)
			}
			if err := d.Set("tools_version", (s.GetToolsVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property ToolsVersion: %+v", err)
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return fmt.Errorf("error occurred while setting property Uuid: %+v", err)
			}
			if err := d.Set("vm_disk_count", (s.GetVmDiskCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VmDiskCount: %+v", err)
			}
			if err := d.Set("vm_overall_status", (s.GetVmOverallStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property VmOverallStatus: %+v", err)
			}
			if err := d.Set("vm_path", (s.GetVmPath())); err != nil {
				return fmt.Errorf("error occurred while setting property VmPath: %+v", err)
			}
			if err := d.Set("vm_version", (s.GetVmVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property VmVersion: %+v", err)
			}
			if err := d.Set("vm_vnic_count", (s.GetVmVnicCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VmVnicCount: %+v", err)
			}
			if err := d.Set("vnic_device_config_id", (s.GetVnicDeviceConfigId())); err != nil {
				return fmt.Errorf("error occurred while setting property VnicDeviceConfigId: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
