package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualizationVmwareHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualizationVmwareHostRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"boot_time": {
				Description: "The time when this host booted up.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster": {
				Description: "A reference to a virtualizationVmwareCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"connection_state": {
				Description: "Indicates if the host is connected to the vCenter. Values are connected, not connected.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_info": {
				Description: "Details about the CPUs installed on this host are represented here.",
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
						"cores": {
							Description: "Number of cores per CPU, as reported by the manufacturer.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"description": {
							Description: "The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"sockets": {
							Description: "Number of CPU sockets available.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"speed": {
							Description: "Speed of the CPUs in Hertz. For example, 2593749663.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"vendor": {
							Description: "Manufacturer of the CPU . For example, Intel.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
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
			"hardware_info": {
				Description: "The hardware details of this host. It includes capacity, manufacturer, and model information.",
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
			"hw_power_state": {
				Description: "Is the host Powered-up or Powered-down.\n* `Unknown` - The entity's power state is unknown.\n* `PoweredOn` - The entity is powered on.\n* `PoweredOff` - The entity is powered down.\n* `StandBy` - The entity is in standby mode.\n* `Paused` - The entity is in pause state.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Identifies the broad type of the underlying hypervisor.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"maintenance_mode": {
				Description: "Is this host in maintenance mode. Set to true or false.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"memory_capacity": {
				Description: "The memory capacity and usage information on this host.",
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
			"model": {
				Description: "Commercial model information about this hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_adapter_count": {
				Description: "The count of all network adapters attached to this host.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"processor_capacity": {
				Description: "The capacity and usage information for CPU power on this host.",
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
			"product_info": {
				Description: "Details of this product, such as vendor, model, etc. are represented here.",
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
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"product_name": {
							Description: "Commercial product name. For example, VMware ESXi.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"product_type": {
							Description: "Product name provided by the vendor. For example, embeddedEsx.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"product_vendor": {
							Description: "Commercial vendor name. For example, VMware Inc.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_version": {
							Description: "Hypervisor version running on the system.",
							Type:        schema.TypeString,
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
			"resource_consumed": {
				Description: "Snapshot of resources (CPU, memory, etc.) consumed by this host.",
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
						"cpu_consumed": {
							Description: "The amount of CPU consumed in Hz.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"memory_consumed": {
							Description: "Memory consumed by this host in bytes.",
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
			"serial": {
				Description: "Serial number of this host (internally generated).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "Host health status, as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"storage_adapter_count": {
				Description: "The count of all storage adapters attached to this host.",
				Type:        schema.TypeInt,
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
				Description: "The uptime of the host, stored as Duration (from w3c).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uuid": {
				Description: "Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vcenter_host_id": {
				Description: "The identity of this host within vCenter (optional).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "Commercial vendor details of this hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceVirtualizationVmwareHostRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VirtualizationVmwareHost{}
	if v, ok := d.GetOk("boot_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetBootTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_state"); ok {
		x := (v.(string))
		o.SetConnectionState(x)
	}
	if v, ok := d.GetOk("hw_power_state"); ok {
		x := (v.(string))
		o.SetHwPowerState(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("maintenance_mode"); ok {
		x := (v.(bool))
		o.SetMaintenanceMode(x)
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
	if v, ok := d.GetOk("network_adapter_count"); ok {
		x := int64(v.(int))
		o.SetNetworkAdapterCount(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("storage_adapter_count"); ok {
		x := int64(v.(int))
		o.SetStorageAdapterCount(x)
	}
	if v, ok := d.GetOk("up_time"); ok {
		x := (v.(string))
		o.SetUpTime(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}
	if v, ok := d.GetOk("vcenter_host_id"); ok {
		x := (v.(string))
		o.SetVcenterHostId(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VirtualizationVmwareHost object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.VirtualizationApi.GetVirtualizationVmwareHostList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching VirtualizationVmwareHost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for VirtualizationVmwareHost list: %s", err.Error())
	}
	var s = &models.VirtualizationVmwareHostList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to VirtualizationVmwareHost list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for VirtualizationVmwareHost data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for VirtualizationVmwareHost data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.VirtualizationVmwareHost{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("boot_time", (s.GetBootTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property BootTime: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster", flattenMapVirtualizationVmwareClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}
			if err := d.Set("connection_state", (s.GetConnectionState())); err != nil {
				return diag.Errorf("error occurred while setting property ConnectionState: %s", err.Error())
			}

			if err := d.Set("cpu_info", flattenMapVirtualizationCpuInfo(s.GetCpuInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property CpuInfo: %s", err.Error())
			}

			if err := d.Set("datacenter", flattenMapVirtualizationVmwareDatacenterRelationship(s.GetDatacenter(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Datacenter: %s", err.Error())
			}

			if err := d.Set("datastores", flattenListVirtualizationVmwareDatastoreRelationship(s.GetDatastores(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Datastores: %s", err.Error())
			}

			if err := d.Set("hardware_info", flattenMapInfraHardwareInfo(s.GetHardwareInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property HardwareInfo: %s", err.Error())
			}
			if err := d.Set("hw_power_state", (s.GetHwPowerState())); err != nil {
				return diag.Errorf("error occurred while setting property HwPowerState: %s", err.Error())
			}
			if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorType: %s", err.Error())
			}
			if err := d.Set("identity", (s.GetIdentity())); err != nil {
				return diag.Errorf("error occurred while setting property Identity: %s", err.Error())
			}
			if err := d.Set("maintenance_mode", (s.GetMaintenanceMode())); err != nil {
				return diag.Errorf("error occurred while setting property MaintenanceMode: %s", err.Error())
			}

			if err := d.Set("memory_capacity", flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property MemoryCapacity: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("network_adapter_count", (s.GetNetworkAdapterCount())); err != nil {
				return diag.Errorf("error occurred while setting property NetworkAdapterCount: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("processor_capacity", flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ProcessorCapacity: %s", err.Error())
			}

			if err := d.Set("product_info", flattenMapVirtualizationProductInfo(s.GetProductInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ProductInfo: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("resource_consumed", flattenMapVirtualizationVmwareResourceConsumption(s.GetResourceConsumed(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ResourceConsumed: %s", err.Error())
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return diag.Errorf("error occurred while setting property Serial: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}
			if err := d.Set("storage_adapter_count", (s.GetStorageAdapterCount())); err != nil {
				return diag.Errorf("error occurred while setting property StorageAdapterCount: %s", err.Error())
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
			if err := d.Set("vcenter_host_id", (s.GetVcenterHostId())); err != nil {
				return diag.Errorf("error occurred while setting property VcenterHostId: %s", err.Error())
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return diag.Errorf("error occurred while setting property Vendor: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
