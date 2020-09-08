package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceComputePhysicalSummary() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComputePhysicalSummaryRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"admin_power_state": {
				Description: "The desired power state of the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"alarm_summary": {
				Description: "The summary of alarm counts based on the severity types (Critical or Warning).",
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
						"critical": {
							Description: "The count of alarms that have severity type Critical.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"warning": {
							Description: "The count of alarms that have severity type Warning.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"asset_tag": {
				Description: "The user defined asset tag assigned to the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"available_memory": {
				Description: "The amount of memory available on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bios_post_complete": {
				Description: "The BIOS POST completion status of the server.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"chassis_id": {
				Description: "The id of the chassis that the blade is located in.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status": {
				Description: "Connectivity Status of RackUnit to Switch - A or B or AB.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cpu_capacity": {
				Description: "CPU Capacity = Number of CPU Sockets x Enabled Cores x Speed (GHz).",
				Type:        schema.TypeFloat,
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
			"equipment_chassis": {
				Description: "A reference to a equipmentChassis resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"fault_summary": {
				Description: "The fault summary for the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"firmware": {
				Description: "The firmware version of the Cisco Integrated Management Controller (CIMC) for this server.",
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
			"ipv4_address": {
				Description: "The IPv4 address configured on the management interface of the Integrated Management Controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"kvm_ip_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"address": {
							Description: "IP Address to be used for KVM.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"category": {
							Description: "Category of the Kvm IP Address.\n* `Equipment` - Ip Address assigned to an equipment.\n* `ServiceProfile` - Ip Address assigned to a Service Profile.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default_gateway": {
							Description: "Default gateway property of KVM IP Address.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"dn": {
							Description: "The distinguished name for this managed object.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"http_port": {
							Description: "HTTP port of an IP Address.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"https_port": {
							Description: "Secured HTTP port of an IP Address.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"kvm_port": {
							Description: "Port number on which the KVM is running and used for connecting to KVM console.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"kvm_vlan": {
							Description: "VLAN Identifier of Inband IP Address.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name to identify the KVM IP Address.\n* `Outband` - The user assigned Out of band IP Address.\n* `Inband` - The user assigned Inband IP Address.",
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
						"subnet": {
							Description: "Subnet detail of a KVM IP Address.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Description: "Type of the KVM IP Address.\n* `MgmtInterface` - Ip Address of a Management Interface.\n* `VnicIpV4StaticAddr` - Static Ipv4 Address of a Virtual Network Interface.\n* `VnicIpV4PooledAddr` - Ipv4 Address of a Virtual Network Interface from an address pool.\n* `VnicIpV4MgmtPooledAddr` - Ipv4 Address of a Virtual Network Interface from a Managed address pool.\n* `VnicIpV6StaticAddr` - Static Ipv6 Address of a Virtual Network Interface.\n* `VnicIpV6MgmtPooledAddr` - Ipv6 Address of a Virtual Network Interface from a Managed address pool.\n* `VnicIpV4ProfDerivedAddr` - Server Profile derived Ipv4 Address of a Virtual Network Interface.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"management_mode": {
				Description: "The management mode of the server.\n* `IntersightStandalone` - Intersight Standalone mode of operation.\n* `UCSM` - Unified Computing System Manager mode of operation.\n* `Intersight` - Intersight managed mode of operation.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_speed": {
				Description: "The maximum memory speed in MHz available on the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mgmt_ip_address": {
				Description: "Management address of the server.",
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
				Description: "The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC).\nWhen this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect.\nWhen this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_adaptors": {
				Description: "The total number of network adapters present on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpu_cores": {
				Description: "The total number of CPU cores present on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpu_cores_enabled": {
				Description: "The total number of CPU cores enabled on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_cpus": {
				Description: "The total number of CPUs present on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_eth_host_interfaces": {
				Description: "The total number of vNICs which are visible to a host on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_fc_host_interfaces": {
				Description: "The total number of vHBAs which are visible to a host on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"num_threads": {
				Description: "The total number of threads the server is capable of handling.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_power_state": {
				Description: "The actual power state of the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "The operational state of the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "The operability of the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "The platform type of the registered device - whether managed by UCSM or operating in standalone mode.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "Indicates if a server is present in a slot and is applicable for blade servers.",
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
			"scaled_mode": {
				Description: "The mode of the server that determines it is scaled.",
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
			"server_id": {
				Description: "RackUnit ID that uniquely identifies the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"service_profile": {
				Description: "The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"slot_id": {
				Description: "The slot number in the chassis that the blade is located in.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"source_object_type": {
				Description: "The source object type of this view MO.",
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
			"topology_scan_status": {
				Description: "To maintain the Topology workflow run status.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_memory": {
				Description: "The total memory available on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"user_label": {
				Description: "The user defined label assigned to the server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "The universally unique identity of the server.",
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
		},
	}
}

func dataSourceComputePhysicalSummaryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewComputePhysicalSummaryWithDefaults()
	if v, ok := d.GetOk("admin_power_state"); ok {
		x := (v.(string))
		o.SetAdminPowerState(x)
	}
	if v, ok := d.GetOk("asset_tag"); ok {
		x := (v.(string))
		o.SetAssetTag(x)
	}
	if v, ok := d.GetOk("available_memory"); ok {
		x := int64(v.(int))
		o.SetAvailableMemory(x)
	}
	if v, ok := d.GetOk("bios_post_complete"); ok {
		x := (v.(bool))
		o.SetBiosPostComplete(x)
	}
	if v, ok := d.GetOk("chassis_id"); ok {
		x := (v.(string))
		o.SetChassisId(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}
	if v, ok := d.GetOk("cpu_capacity"); ok {
		x := v.(float32)
		o.SetCpuCapacity(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fault_summary"); ok {
		x := int64(v.(int))
		o.SetFaultSummary(x)
	}
	if v, ok := d.GetOk("firmware"); ok {
		x := (v.(string))
		o.SetFirmware(x)
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.SetIpv4Address(x)
	}
	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
	}
	if v, ok := d.GetOk("memory_speed"); ok {
		x := (v.(string))
		o.SetMemorySpeed(x)
	}
	if v, ok := d.GetOk("mgmt_ip_address"); ok {
		x := (v.(string))
		o.SetMgmtIpAddress(x)
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
	if v, ok := d.GetOk("num_adaptors"); ok {
		x := int64(v.(int))
		o.SetNumAdaptors(x)
	}
	if v, ok := d.GetOk("num_cpu_cores"); ok {
		x := int64(v.(int))
		o.SetNumCpuCores(x)
	}
	if v, ok := d.GetOk("num_cpu_cores_enabled"); ok {
		x := int64(v.(int))
		o.SetNumCpuCoresEnabled(x)
	}
	if v, ok := d.GetOk("num_cpus"); ok {
		x := int64(v.(int))
		o.SetNumCpus(x)
	}
	if v, ok := d.GetOk("num_eth_host_interfaces"); ok {
		x := int64(v.(int))
		o.SetNumEthHostInterfaces(x)
	}
	if v, ok := d.GetOk("num_fc_host_interfaces"); ok {
		x := int64(v.(int))
		o.SetNumFcHostInterfaces(x)
	}
	if v, ok := d.GetOk("num_threads"); ok {
		x := int64(v.(int))
		o.SetNumThreads(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.SetOperPowerState(x)
	}
	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}
	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("scaled_mode"); ok {
		x := (v.(string))
		o.SetScaledMode(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("server_id"); ok {
		x := int64(v.(int))
		o.SetServerId(x)
	}
	if v, ok := d.GetOk("service_profile"); ok {
		x := (v.(string))
		o.SetServiceProfile(x)
	}
	if v, ok := d.GetOk("slot_id"); ok {
		x := int64(v.(int))
		o.SetSlotId(x)
	}
	if v, ok := d.GetOk("source_object_type"); ok {
		x := (v.(string))
		o.SetSourceObjectType(x)
	}
	if v, ok := d.GetOk("topology_scan_status"); ok {
		x := (v.(string))
		o.SetTopologyScanStatus(x)
	}
	if v, ok := d.GetOk("total_memory"); ok {
		x := int64(v.(int))
		o.SetTotalMemory(x)
	}
	if v, ok := d.GetOk("user_label"); ok {
		x := (v.(string))
		o.SetUserLabel(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.ComputeApi.GetComputePhysicalSummaryList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.ComputePhysicalSummaryList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to ComputePhysicalSummary: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewComputePhysicalSummaryWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("admin_power_state", (s.GetAdminPowerState())); err != nil {
				return fmt.Errorf("error occurred while setting property AdminPowerState: %+v", err)
			}

			if err := d.Set("alarm_summary", flattenMapComputeAlarmSummary(s.GetAlarmSummary(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property AlarmSummary: %+v", err)
			}
			if err := d.Set("asset_tag", (s.GetAssetTag())); err != nil {
				return fmt.Errorf("error occurred while setting property AssetTag: %+v", err)
			}
			if err := d.Set("available_memory", (s.GetAvailableMemory())); err != nil {
				return fmt.Errorf("error occurred while setting property AvailableMemory: %+v", err)
			}
			if err := d.Set("bios_post_complete", (s.GetBiosPostComplete())); err != nil {
				return fmt.Errorf("error occurred while setting property BiosPostComplete: %+v", err)
			}
			if err := d.Set("chassis_id", (s.GetChassisId())); err != nil {
				return fmt.Errorf("error occurred while setting property ChassisId: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("connection_status", (s.GetConnectionStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionStatus: %+v", err)
			}
			if err := d.Set("cpu_capacity", (s.GetCpuCapacity())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuCapacity: %+v", err)
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceMoId: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}

			if err := d.Set("equipment_chassis", flattenMapEquipmentChassisRelationship(s.GetEquipmentChassis(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property EquipmentChassis: %+v", err)
			}
			if err := d.Set("fault_summary", (s.GetFaultSummary())); err != nil {
				return fmt.Errorf("error occurred while setting property FaultSummary: %+v", err)
			}
			if err := d.Set("firmware", (s.GetFirmware())); err != nil {
				return fmt.Errorf("error occurred while setting property Firmware: %+v", err)
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property InventoryDeviceInfo: %+v", err)
			}
			if err := d.Set("ipv4_address", (s.GetIpv4Address())); err != nil {
				return fmt.Errorf("error occurred while setting property Ipv4Address: %+v", err)
			}

			if err := d.Set("kvm_ip_addresses", flattenListComputeIpAddress(s.GetKvmIpAddresses(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property KvmIpAddresses: %+v", err)
			}
			if err := d.Set("management_mode", (s.GetManagementMode())); err != nil {
				return fmt.Errorf("error occurred while setting property ManagementMode: %+v", err)
			}
			if err := d.Set("memory_speed", (s.GetMemorySpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property MemorySpeed: %+v", err)
			}
			if err := d.Set("mgmt_ip_address", (s.GetMgmtIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property MgmtIpAddress: %+v", err)
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return fmt.Errorf("error occurred while setting property Model: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("num_adaptors", (s.GetNumAdaptors())); err != nil {
				return fmt.Errorf("error occurred while setting property NumAdaptors: %+v", err)
			}
			if err := d.Set("num_cpu_cores", (s.GetNumCpuCores())); err != nil {
				return fmt.Errorf("error occurred while setting property NumCpuCores: %+v", err)
			}
			if err := d.Set("num_cpu_cores_enabled", (s.GetNumCpuCoresEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property NumCpuCoresEnabled: %+v", err)
			}
			if err := d.Set("num_cpus", (s.GetNumCpus())); err != nil {
				return fmt.Errorf("error occurred while setting property NumCpus: %+v", err)
			}
			if err := d.Set("num_eth_host_interfaces", (s.GetNumEthHostInterfaces())); err != nil {
				return fmt.Errorf("error occurred while setting property NumEthHostInterfaces: %+v", err)
			}
			if err := d.Set("num_fc_host_interfaces", (s.GetNumFcHostInterfaces())); err != nil {
				return fmt.Errorf("error occurred while setting property NumFcHostInterfaces: %+v", err)
			}
			if err := d.Set("num_threads", (s.GetNumThreads())); err != nil {
				return fmt.Errorf("error occurred while setting property NumThreads: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("oper_power_state", (s.GetOperPowerState())); err != nil {
				return fmt.Errorf("error occurred while setting property OperPowerState: %+v", err)
			}
			if err := d.Set("oper_state", (s.GetOperState())); err != nil {
				return fmt.Errorf("error occurred while setting property OperState: %+v", err)
			}
			if err := d.Set("operability", (s.GetOperability())); err != nil {
				return fmt.Errorf("error occurred while setting property Operability: %+v", err)
			}
			if err := d.Set("platform_type", (s.GetPlatformType())); err != nil {
				return fmt.Errorf("error occurred while setting property PlatformType: %+v", err)
			}
			if err := d.Set("presence", (s.GetPresence())); err != nil {
				return fmt.Errorf("error occurred while setting property Presence: %+v", err)
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
			if err := d.Set("scaled_mode", (s.GetScaledMode())); err != nil {
				return fmt.Errorf("error occurred while setting property ScaledMode: %+v", err)
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return fmt.Errorf("error occurred while setting property Serial: %+v", err)
			}
			if err := d.Set("server_id", (s.GetServerId())); err != nil {
				return fmt.Errorf("error occurred while setting property ServerId: %+v", err)
			}
			if err := d.Set("service_profile", (s.GetServiceProfile())); err != nil {
				return fmt.Errorf("error occurred while setting property ServiceProfile: %+v", err)
			}
			if err := d.Set("slot_id", (s.GetSlotId())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotId: %+v", err)
			}
			if err := d.Set("source_object_type", (s.GetSourceObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property SourceObjectType: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("topology_scan_status", (s.GetTopologyScanStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property TopologyScanStatus: %+v", err)
			}
			if err := d.Set("total_memory", (s.GetTotalMemory())); err != nil {
				return fmt.Errorf("error occurred while setting property TotalMemory: %+v", err)
			}
			if err := d.Set("user_label", (s.GetUserLabel())); err != nil {
				return fmt.Errorf("error occurred while setting property UserLabel: %+v", err)
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return fmt.Errorf("error occurred while setting property Uuid: %+v", err)
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property Vendor: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
