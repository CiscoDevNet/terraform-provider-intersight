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

func dataSourceStorageHitachiPort() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageHitachiPortRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"array": {
				Description: "A reference to a storageHitachiArray resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fabric_mode": {
				Description: "Fabric mode of the port. true, Set. false, Not set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "IPv4 address.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv6_global_address": {
				Description: "IPv6 global address value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv6_link_local_address": {
				Description: "IPv6 link local address value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"iqn": {
				Description: "ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"is_ipv6_enable": {
				Description: "IPv6 mode.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"loop_id": {
				Description: "The value set for the port loop ID (AL_PA).",
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
				Description: "Name of the physical port available in storage array.",
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
			"port_connection": {
				Description: "Topology setting for the port.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port_lun_security": {
				Description: "LUN security setting for the port.",
				Type:        schema.TypeBool,
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
			"shortport_id": {
				Description: "Port ID (short) of the port.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"speed": {
				Description: "Operational speed of physical port measured in Gbps.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of storage array port.\n* `Unknown` - Component status is not available.\n* `Ok` - Component is healthy and no issues found.\n* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.\n* `Critical` - Not functioning or requiring immediate attention.\n* `Offline` - Component is installed, but powered off.\n* `Identifying` - Component is in initialization process.\n* `NotAvailable` - Component is not installed or not available.\n* `Updating` - Software update is in progress.\n* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported.",
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
			"tcp_mtu": {
				Description: "Value of MTU for iSCSI communication.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Description: "Port type - possible values are FC, FCoE and iSCSI.\n* `FC` - Port supports fibre channel protocol.\n* `iSCSI` - Port supports iSCSI protocol.\n* `FCoE` - Port supports fibre channel over ethernet protocol.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wwn": {
				Description: "World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon.\nExample: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wwnn": {
				Description: "World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wwpn": {
				Description: "World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceStorageHitachiPortRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageHitachiPort{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("fabric_mode"); ok {
		x := (v.(bool))
		o.SetFabricMode(x)
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.SetIpv4Address(x)
	}
	if v, ok := d.GetOk("ipv6_global_address"); ok {
		x := (v.(string))
		o.SetIpv6GlobalAddress(x)
	}
	if v, ok := d.GetOk("ipv6_link_local_address"); ok {
		x := (v.(string))
		o.SetIpv6LinkLocalAddress(x)
	}
	if v, ok := d.GetOk("iqn"); ok {
		x := (v.(string))
		o.SetIqn(x)
	}
	if v, ok := d.GetOk("is_ipv6_enable"); ok {
		x := (v.(bool))
		o.SetIsIpv6Enable(x)
	}
	if v, ok := d.GetOk("loop_id"); ok {
		x := (v.(string))
		o.SetLoopId(x)
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
	if v, ok := d.GetOk("port_connection"); ok {
		x := (v.(string))
		o.SetPortConnection(x)
	}
	if v, ok := d.GetOk("port_lun_security"); ok {
		x := (v.(bool))
		o.SetPortLunSecurity(x)
	}
	if v, ok := d.GetOk("shortport_id"); ok {
		x := (v.(string))
		o.SetShortportId(x)
	}
	if v, ok := d.GetOk("speed"); ok {
		x := int64(v.(int))
		o.SetSpeed(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("tcp_mtu"); ok {
		x := int64(v.(int))
		o.SetTcpMtu(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("wwn"); ok {
		x := (v.(string))
		o.SetWwn(x)
	}
	if v, ok := d.GetOk("wwnn"); ok {
		x := (v.(string))
		o.SetWwnn(x)
	}
	if v, ok := d.GetOk("wwpn"); ok {
		x := (v.(string))
		o.SetWwpn(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageHitachiPort object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiPortList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching StorageHitachiPort: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageHitachiPort list: %s", err.Error())
	}
	var s = &models.StorageHitachiPortList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageHitachiPort list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageHitachiPort data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageHitachiPort data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageHitachiPort{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("array", flattenMapStorageHitachiArrayRelationship(s.GetArray(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Array: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("fabric_mode", (s.GetFabricMode())); err != nil {
				return diag.Errorf("error occurred while setting property FabricMode: %s", err.Error())
			}
			if err := d.Set("ipv4_address", (s.GetIpv4Address())); err != nil {
				return diag.Errorf("error occurred while setting property Ipv4Address: %s", err.Error())
			}
			if err := d.Set("ipv6_global_address", (s.GetIpv6GlobalAddress())); err != nil {
				return diag.Errorf("error occurred while setting property Ipv6GlobalAddress: %s", err.Error())
			}
			if err := d.Set("ipv6_link_local_address", (s.GetIpv6LinkLocalAddress())); err != nil {
				return diag.Errorf("error occurred while setting property Ipv6LinkLocalAddress: %s", err.Error())
			}
			if err := d.Set("iqn", (s.GetIqn())); err != nil {
				return diag.Errorf("error occurred while setting property Iqn: %s", err.Error())
			}
			if err := d.Set("is_ipv6_enable", (s.GetIsIpv6Enable())); err != nil {
				return diag.Errorf("error occurred while setting property IsIpv6Enable: %s", err.Error())
			}
			if err := d.Set("loop_id", (s.GetLoopId())); err != nil {
				return diag.Errorf("error occurred while setting property LoopId: %s", err.Error())
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
			if err := d.Set("port_connection", (s.GetPortConnection())); err != nil {
				return diag.Errorf("error occurred while setting property PortConnection: %s", err.Error())
			}
			if err := d.Set("port_lun_security", (s.GetPortLunSecurity())); err != nil {
				return diag.Errorf("error occurred while setting property PortLunSecurity: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("shortport_id", (s.GetShortportId())); err != nil {
				return diag.Errorf("error occurred while setting property ShortportId: %s", err.Error())
			}
			if err := d.Set("speed", (s.GetSpeed())); err != nil {
				return diag.Errorf("error occurred while setting property Speed: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("tcp_mtu", (s.GetTcpMtu())); err != nil {
				return diag.Errorf("error occurred while setting property TcpMtu: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			if err := d.Set("wwn", (s.GetWwn())); err != nil {
				return diag.Errorf("error occurred while setting property Wwn: %s", err.Error())
			}
			if err := d.Set("wwnn", (s.GetWwnn())); err != nil {
				return diag.Errorf("error occurred while setting property Wwnn: %s", err.Error())
			}
			if err := d.Set("wwpn", (s.GetWwpn())); err != nil {
				return diag.Errorf("error occurred while setting property Wwpn: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
