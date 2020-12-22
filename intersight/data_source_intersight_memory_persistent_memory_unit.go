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

func dataSourceMemoryPersistentMemoryUnit() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMemoryPersistentMemoryUnitRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"admin_state": {
				Description: "This represents the administrative state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"app_direct_capacity": {
				Description: "AppDirect capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"array_id": {
				Description: "This represents the memory array to which the memory unit belongs to.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bank": {
				Description: "This represents the memory bank of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"capacity": {
				Description: "This represents the memory capacity in MiB of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"clock": {
				Description: "This represents the clock of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"count_status": {
				Description: "Count status of the Persistent Memory Module on a server.",
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
			"firmware_version": {
				Description: "Firmware version of the firware running on the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"form_factor": {
				Description: "This represents the form factor of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"frozen_status": {
				Description: "Frozen status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"health_state": {
				Description: "Health state of the Persistent Memory Module on a server.",
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
			"latency": {
				Description: "This represents the latency of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Description: "This represents the location of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"lock_status": {
				Description: "Lock status of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_array": {
				Description: "A reference to a memoryArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"memory_capacity": {
				Description: "Memory capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"memory_id": {
				Description: "ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeInt,
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_power_state": {
				Description: "This represents the operational power state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"oper_state": {
				Description: "This represents the operational state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operability": {
				Description: "This represents the operability of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"persistent_memory_capacity": {
				Description: "Persistent Memory capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"presence": {
				Description: "This represents the presence state of the memory unit on a server.",
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
			"reserved_capacity": {
				Description: "Reserved capacity in GiB of the Persistent Memory Module on a server.",
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
			"security_status": {
				Description: "Security status of the Persistent Memory Module on a server.",
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
			"set": {
				Description: "This represents the set of the memory unit on a server.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"socket_id": {
				Description: "Socket ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"socket_memory_id": {
				Description: "Socket Memory ID of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"speed": {
				Description: "This represents the speed of the memory unit on a server.",
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
			"thermal": {
				Description: "This represents the thremal state of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_capacity": {
				Description: "Total capacity in GiB of the Persistent Memory Module on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "This represents the memory type of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uid": {
				Description: "UID of the Persistent Memory Module on a server.",
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
			"visibility": {
				Description: "This represents the visibility of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"width": {
				Description: "This represents the width of the memory unit on a server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceMemoryPersistentMemoryUnitRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.MemoryPersistentMemoryUnit{}
	if v, ok := d.GetOk("admin_state"); ok {
		x := (v.(string))
		o.SetAdminState(x)
	}
	if v, ok := d.GetOk("app_direct_capacity"); ok {
		x := (v.(string))
		o.SetAppDirectCapacity(x)
	}
	if v, ok := d.GetOk("array_id"); ok {
		x := int64(v.(int))
		o.SetArrayId(x)
	}
	if v, ok := d.GetOk("bank"); ok {
		x := int64(v.(int))
		o.SetBank(x)
	}
	if v, ok := d.GetOk("capacity"); ok {
		x := (v.(string))
		o.SetCapacity(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("clock"); ok {
		x := (v.(string))
		o.SetClock(x)
	}
	if v, ok := d.GetOk("count_status"); ok {
		x := (v.(string))
		o.SetCountStatus(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("firmware_version"); ok {
		x := (v.(string))
		o.SetFirmwareVersion(x)
	}
	if v, ok := d.GetOk("form_factor"); ok {
		x := (v.(string))
		o.SetFormFactor(x)
	}
	if v, ok := d.GetOk("frozen_status"); ok {
		x := (v.(string))
		o.SetFrozenStatus(x)
	}
	if v, ok := d.GetOk("health_state"); ok {
		x := (v.(string))
		o.SetHealthState(x)
	}
	if v, ok := d.GetOk("latency"); ok {
		x := (v.(string))
		o.SetLatency(x)
	}
	if v, ok := d.GetOk("location"); ok {
		x := (v.(string))
		o.SetLocation(x)
	}
	if v, ok := d.GetOk("lock_status"); ok {
		x := (v.(string))
		o.SetLockStatus(x)
	}
	if v, ok := d.GetOk("memory_capacity"); ok {
		x := (v.(string))
		o.SetMemoryCapacity(x)
	}
	if v, ok := d.GetOk("memory_id"); ok {
		x := int64(v.(int))
		o.SetMemoryId(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
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
	if v, ok := d.GetOk("persistent_memory_capacity"); ok {
		x := (v.(string))
		o.SetPersistentMemoryCapacity(x)
	}
	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}
	if v, ok := d.GetOk("reserved_capacity"); ok {
		x := (v.(string))
		o.SetReservedCapacity(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("security_status"); ok {
		x := (v.(string))
		o.SetSecurityStatus(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("set"); ok {
		x := int64(v.(int))
		o.SetSet(x)
	}
	if v, ok := d.GetOk("socket_id"); ok {
		x := (v.(string))
		o.SetSocketId(x)
	}
	if v, ok := d.GetOk("socket_memory_id"); ok {
		x := (v.(string))
		o.SetSocketMemoryId(x)
	}
	if v, ok := d.GetOk("speed"); ok {
		x := (v.(string))
		o.SetSpeed(x)
	}
	if v, ok := d.GetOk("thermal"); ok {
		x := (v.(string))
		o.SetThermal(x)
	}
	if v, ok := d.GetOk("total_capacity"); ok {
		x := (v.(string))
		o.SetTotalCapacity(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("uid"); ok {
		x := (v.(string))
		o.SetUid(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("visibility"); ok {
		x := (v.(string))
		o.SetVisibility(x)
	}
	if v, ok := d.GetOk("width"); ok {
		x := (v.(string))
		o.SetWidth(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of MemoryPersistentMemoryUnit object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.MemoryApi.GetMemoryPersistentMemoryUnitList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching MemoryPersistentMemoryUnit: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for MemoryPersistentMemoryUnit list: %s", err.Error())
	}
	var s = &models.MemoryPersistentMemoryUnitList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to MemoryPersistentMemoryUnit list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for MemoryPersistentMemoryUnit did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.MemoryPersistentMemoryUnit{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("admin_state", (s.GetAdminState())); err != nil {
				return diag.Errorf("error occurred while setting property AdminState: %s", err.Error())
			}
			if err := d.Set("app_direct_capacity", (s.GetAppDirectCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property AppDirectCapacity: %s", err.Error())
			}
			if err := d.Set("array_id", (s.GetArrayId())); err != nil {
				return diag.Errorf("error occurred while setting property ArrayId: %s", err.Error())
			}
			if err := d.Set("bank", (s.GetBank())); err != nil {
				return diag.Errorf("error occurred while setting property Bank: %s", err.Error())
			}
			if err := d.Set("capacity", (s.GetCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property Capacity: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("clock", (s.GetClock())); err != nil {
				return diag.Errorf("error occurred while setting property Clock: %s", err.Error())
			}
			if err := d.Set("count_status", (s.GetCountStatus())); err != nil {
				return diag.Errorf("error occurred while setting property CountStatus: %s", err.Error())
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceMoId: %s", err.Error())
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return diag.Errorf("error occurred while setting property Dn: %s", err.Error())
			}
			if err := d.Set("firmware_version", (s.GetFirmwareVersion())); err != nil {
				return diag.Errorf("error occurred while setting property FirmwareVersion: %s", err.Error())
			}
			if err := d.Set("form_factor", (s.GetFormFactor())); err != nil {
				return diag.Errorf("error occurred while setting property FormFactor: %s", err.Error())
			}
			if err := d.Set("frozen_status", (s.GetFrozenStatus())); err != nil {
				return diag.Errorf("error occurred while setting property FrozenStatus: %s", err.Error())
			}
			if err := d.Set("health_state", (s.GetHealthState())); err != nil {
				return diag.Errorf("error occurred while setting property HealthState: %s", err.Error())
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InventoryDeviceInfo: %s", err.Error())
			}
			if err := d.Set("latency", (s.GetLatency())); err != nil {
				return diag.Errorf("error occurred while setting property Latency: %s", err.Error())
			}
			if err := d.Set("location", (s.GetLocation())); err != nil {
				return diag.Errorf("error occurred while setting property Location: %s", err.Error())
			}
			if err := d.Set("lock_status", (s.GetLockStatus())); err != nil {
				return diag.Errorf("error occurred while setting property LockStatus: %s", err.Error())
			}

			if err := d.Set("memory_array", flattenMapMemoryArrayRelationship(s.GetMemoryArray(), d)); err != nil {
				return diag.Errorf("error occurred while setting property MemoryArray: %s", err.Error())
			}
			if err := d.Set("memory_capacity", (s.GetMemoryCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property MemoryCapacity: %s", err.Error())
			}
			if err := d.Set("memory_id", (s.GetMemoryId())); err != nil {
				return diag.Errorf("error occurred while setting property MemoryId: %s", err.Error())
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return diag.Errorf("error occurred while setting property Model: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("oper_power_state", (s.GetOperPowerState())); err != nil {
				return diag.Errorf("error occurred while setting property OperPowerState: %s", err.Error())
			}
			if err := d.Set("oper_state", (s.GetOperState())); err != nil {
				return diag.Errorf("error occurred while setting property OperState: %s", err.Error())
			}
			if err := d.Set("operability", (s.GetOperability())); err != nil {
				return diag.Errorf("error occurred while setting property Operability: %s", err.Error())
			}
			if err := d.Set("persistent_memory_capacity", (s.GetPersistentMemoryCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property PersistentMemoryCapacity: %s", err.Error())
			}
			if err := d.Set("presence", (s.GetPresence())); err != nil {
				return diag.Errorf("error occurred while setting property Presence: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("reserved_capacity", (s.GetReservedCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property ReservedCapacity: %s", err.Error())
			}
			if err := d.Set("revision", (s.GetRevision())); err != nil {
				return diag.Errorf("error occurred while setting property Revision: %s", err.Error())
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return diag.Errorf("error occurred while setting property Rn: %s", err.Error())
			}
			if err := d.Set("security_status", (s.GetSecurityStatus())); err != nil {
				return diag.Errorf("error occurred while setting property SecurityStatus: %s", err.Error())
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return diag.Errorf("error occurred while setting property Serial: %s", err.Error())
			}
			if err := d.Set("set", (s.GetSet())); err != nil {
				return diag.Errorf("error occurred while setting property Set: %s", err.Error())
			}
			if err := d.Set("socket_id", (s.GetSocketId())); err != nil {
				return diag.Errorf("error occurred while setting property SocketId: %s", err.Error())
			}
			if err := d.Set("socket_memory_id", (s.GetSocketMemoryId())); err != nil {
				return diag.Errorf("error occurred while setting property SocketMemoryId: %s", err.Error())
			}
			if err := d.Set("speed", (s.GetSpeed())); err != nil {
				return diag.Errorf("error occurred while setting property Speed: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("thermal", (s.GetThermal())); err != nil {
				return diag.Errorf("error occurred while setting property Thermal: %s", err.Error())
			}
			if err := d.Set("total_capacity", (s.GetTotalCapacity())); err != nil {
				return diag.Errorf("error occurred while setting property TotalCapacity: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("uid", (s.GetUid())); err != nil {
				return diag.Errorf("error occurred while setting property Uid: %s", err.Error())
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return diag.Errorf("error occurred while setting property Vendor: %s", err.Error())
			}
			if err := d.Set("visibility", (s.GetVisibility())); err != nil {
				return diag.Errorf("error occurred while setting property Visibility: %s", err.Error())
			}
			if err := d.Set("width", (s.GetWidth())); err != nil {
				return diag.Errorf("error occurred while setting property Width: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
