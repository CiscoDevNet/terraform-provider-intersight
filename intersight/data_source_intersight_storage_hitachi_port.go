package intersight

import (
	"context"
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
				Description: "IPv4 address of Hitachi Port.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
						Description: "IPv4 address of Hitachi Port.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
				}},
				Computed: true,
			}},
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
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiPortList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageHitachiPort: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageHitachiPortList.GetCount()
	var i int32
	var storageHitachiPortResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiPortList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageHitachiPort: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageHitachiPortList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageHitachiPort data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array"] = flattenMapStorageHitachiArrayRelationship(s.GetArray(), d)
				temp["class_id"] = (s.GetClassId())
				temp["fabric_mode"] = (s.GetFabricMode())
				temp["ipv4_address"] = (s.GetIpv4Address())
				temp["ipv6_global_address"] = (s.GetIpv6GlobalAddress())
				temp["ipv6_link_local_address"] = (s.GetIpv6LinkLocalAddress())
				temp["iqn"] = (s.GetIqn())
				temp["is_ipv6_enable"] = (s.GetIsIpv6Enable())
				temp["loop_id"] = (s.GetLoopId())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["port_connection"] = (s.GetPortConnection())
				temp["port_lun_security"] = (s.GetPortLunSecurity())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["shortport_id"] = (s.GetShortportId())
				temp["speed"] = (s.GetSpeed())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["tcp_mtu"] = (s.GetTcpMtu())
				temp["type"] = (s.GetType())
				temp["wwn"] = (s.GetWwn())
				temp["wwnn"] = (s.GetWwnn())
				temp["wwpn"] = (s.GetWwpn())
				storageHitachiPortResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageHitachiPortResults))
	if err := d.Set("results", storageHitachiPortResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageHitachiPortResults[0]["moid"].(string))
	return de
}
