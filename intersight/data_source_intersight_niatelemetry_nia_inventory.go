package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryNiaInventory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryNiaInventoryRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bgp_peer_count": {
				Description: "Number of BGP peers on a node.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu": {
				Description: "CPU usage of device being inventoried. This determines the percentage of CPU resources used.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"crash_reset_logs": {
				Description: "Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_name": {
				Description: "Name of device being inventoried. The name the user assigns to the device is inventoried here.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_type": {
				Description: "Type of device being inventoried. This determines whether the device is a controller, leaf or spine.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disk": {
				Description: "Disk Usage of device being inventoried. This determines the amount of disk usage.",
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
						"free": {
							Description: "The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"name": {
							Description: "Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"total": {
							Description: "The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"used": {
							Description: "The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"dn": {
				Description: "Dn for the inventories present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fex_count": {
				Description: "Number of fabric extendors utilized.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"infra_wi_node_count": {
				Description: "Number of appliances as physical device that are wired into the cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ip_address": {
				Description: "The IP address of the device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"license_state": {
				Description: "A reference to a niatelemetryNiaLicenseState resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"log_in_time": {
				Description: "Last log in time device being inventoried. This determines the last login time on the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"log_out_time": {
				Description: "Last log out time of device being inventoried. This determines the last logout time on the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mac_sec_count": {
				Description: "Number of Macsec configured interfaces on a TOR.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"mac_sec_fab_count": {
				Description: "Number of Macsec configured interfaces on a Spine.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"macsec_total_count": {
				Description: "Number of total Macsec configured interfaces for all nodes.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"memory": {
				Description: "Memory usage of device being inventoried. This determines the percentage of memory resources used.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"node_id": {
				Description: "The ID of the device being inventoried.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"record_type": {
				Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_version": {
				Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
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
			"route_prefix_count": {
				Description: "Total nuumber of v4 and v6 routes per node.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"route_prefix_v4_count": {
				Description: "Number of v4 routes per node.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"route_prefix_v6_count": {
				Description: "Number of v6 routes per node.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"serial": {
				Description: "Serial number of device being invetoried. The serial number is unique per device and will be used as the key.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_download": {
				Description: "Last software downloaded of device being inventoried. This determines if software download API was used.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"system_up_time": {
				Description: "The amount of time that the device being inventoried been up.",
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
			"nr_version": {
				Description: "Software version of device being inventoried. The various software version values for each device are available on cisco.com.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceNiatelemetryNiaInventoryRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.NiatelemetryNiaInventory{}
	if v, ok := d.GetOk("bgp_peer_count"); ok {
		x := int64(v.(int))
		o.SetBgpPeerCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cpu"); ok {
		x := v.(float32)
		o.SetCpu(x)
	}
	if v, ok := d.GetOk("crash_reset_logs"); ok {
		x := (v.(string))
		o.SetCrashResetLogs(x)
	}
	if v, ok := d.GetOk("device_name"); ok {
		x := (v.(string))
		o.SetDeviceName(x)
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fex_count"); ok {
		x := int64(v.(int))
		o.SetFexCount(x)
	}
	if v, ok := d.GetOk("infra_wi_node_count"); ok {
		x := int64(v.(int))
		o.SetInfraWiNodeCount(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("log_in_time"); ok {
		x := (v.(string))
		o.SetLogInTime(x)
	}
	if v, ok := d.GetOk("log_out_time"); ok {
		x := (v.(string))
		o.SetLogOutTime(x)
	}
	if v, ok := d.GetOk("mac_sec_count"); ok {
		x := int64(v.(int))
		o.SetMacSecCount(x)
	}
	if v, ok := d.GetOk("mac_sec_fab_count"); ok {
		x := int64(v.(int))
		o.SetMacSecFabCount(x)
	}
	if v, ok := d.GetOk("macsec_total_count"); ok {
		x := int64(v.(int))
		o.SetMacsecTotalCount(x)
	}
	if v, ok := d.GetOk("memory"); ok {
		x := int64(v.(int))
		o.SetMemory(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("node_id"); ok {
		x := (v.(string))
		o.SetNodeId(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("route_prefix_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixCount(x)
	}
	if v, ok := d.GetOk("route_prefix_v4_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV4Count(x)
	}
	if v, ok := d.GetOk("route_prefix_v6_count"); ok {
		x := int64(v.(int))
		o.SetRoutePrefixV6Count(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("software_download"); ok {
		x := (v.(string))
		o.SetSoftwareDownload(x)
	}
	if v, ok := d.GetOk("system_up_time"); ok {
		x := (v.(string))
		o.SetSystemUpTime(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiatelemetryNiaInventoryList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiatelemetryNiaInventory: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiatelemetryNiaInventory{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("bgp_peer_count", (s.GetBgpPeerCount())); err != nil {
				return fmt.Errorf("error occurred while setting property BgpPeerCount: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("cpu", (s.GetCpu())); err != nil {
				return fmt.Errorf("error occurred while setting property Cpu: %+v", err)
			}
			if err := d.Set("crash_reset_logs", (s.GetCrashResetLogs())); err != nil {
				return fmt.Errorf("error occurred while setting property CrashResetLogs: %+v", err)
			}
			if err := d.Set("device_name", (s.GetDeviceName())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceName: %+v", err)
			}
			if err := d.Set("device_type", (s.GetDeviceType())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceType: %+v", err)
			}

			if err := d.Set("disk", flattenMapNiatelemetryDiskinfo(s.GetDisk(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Disk: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}
			if err := d.Set("fex_count", (s.GetFexCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FexCount: %+v", err)
			}
			if err := d.Set("infra_wi_node_count", (s.GetInfraWiNodeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property InfraWiNodeCount: %+v", err)
			}
			if err := d.Set("ip_address", (s.GetIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property IpAddress: %+v", err)
			}

			if err := d.Set("license_state", flattenMapNiatelemetryNiaLicenseStateRelationship(s.GetLicenseState(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property LicenseState: %+v", err)
			}
			if err := d.Set("log_in_time", (s.GetLogInTime())); err != nil {
				return fmt.Errorf("error occurred while setting property LogInTime: %+v", err)
			}
			if err := d.Set("log_out_time", (s.GetLogOutTime())); err != nil {
				return fmt.Errorf("error occurred while setting property LogOutTime: %+v", err)
			}
			if err := d.Set("mac_sec_count", (s.GetMacSecCount())); err != nil {
				return fmt.Errorf("error occurred while setting property MacSecCount: %+v", err)
			}
			if err := d.Set("mac_sec_fab_count", (s.GetMacSecFabCount())); err != nil {
				return fmt.Errorf("error occurred while setting property MacSecFabCount: %+v", err)
			}
			if err := d.Set("macsec_total_count", (s.GetMacsecTotalCount())); err != nil {
				return fmt.Errorf("error occurred while setting property MacsecTotalCount: %+v", err)
			}
			if err := d.Set("memory", (s.GetMemory())); err != nil {
				return fmt.Errorf("error occurred while setting property Memory: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("node_id", (s.GetNodeId())); err != nil {
				return fmt.Errorf("error occurred while setting property NodeId: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("record_type", (s.GetRecordType())); err != nil {
				return fmt.Errorf("error occurred while setting property RecordType: %+v", err)
			}
			if err := d.Set("record_version", (s.GetRecordVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property RecordVersion: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("route_prefix_count", (s.GetRoutePrefixCount())); err != nil {
				return fmt.Errorf("error occurred while setting property RoutePrefixCount: %+v", err)
			}
			if err := d.Set("route_prefix_v4_count", (s.GetRoutePrefixV4Count())); err != nil {
				return fmt.Errorf("error occurred while setting property RoutePrefixV4Count: %+v", err)
			}
			if err := d.Set("route_prefix_v6_count", (s.GetRoutePrefixV6Count())); err != nil {
				return fmt.Errorf("error occurred while setting property RoutePrefixV6Count: %+v", err)
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return fmt.Errorf("error occurred while setting property Serial: %+v", err)
			}
			if err := d.Set("software_download", (s.GetSoftwareDownload())); err != nil {
				return fmt.Errorf("error occurred while setting property SoftwareDownload: %+v", err)
			}
			if err := d.Set("system_up_time", (s.GetSystemUpTime())); err != nil {
				return fmt.Errorf("error occurred while setting property SystemUpTime: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property Version: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
