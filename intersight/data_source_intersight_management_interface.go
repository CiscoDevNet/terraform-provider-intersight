package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceManagementInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceManagementInterfaceRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"gateway": {
				Description: "Default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_name": {
				Description: "Hostname configured for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_address": {
				Description: "IP address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Description: "IPv4 address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_gateway": {
				Description: "IPv4 default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_mask": {
				Description: "IPv4 Netmask for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ipv6_address": {
				Description: "IPv6 address of the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6_gateway": {
				Description: "IPv6 default gateway for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6_prefix": {
				Description: "IPv6 prefix for the interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"mac_address": {
				Description: "MAC address configured for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mask": {
				Description: "Netmask for the interface.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
			"switch_id": {
				Description: "Switch Id connected to the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"uem_conn_status": {
				Description: "The event channel connection status for the interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"virtual_host_name": {
				Description: "Virtual hostname configured for the interface in case of clustered environment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vlan_id": {
				Description: "VlanId configured for the interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"gateway": {
						Description: "Default gateway for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"host_name": {
						Description: "Hostname configured for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"ip_address": {
						Description: "IP address of the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ipv4_address": {
						Description: "IPv4 address of the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ipv4_gateway": {
						Description: "IPv4 default gateway for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ipv4_mask": {
						Description: "IPv4 Netmask for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"ipv6_address": {
						Description: "IPv6 address of the interface.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_gateway": {
						Description: "IPv6 default gateway for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ipv6_prefix": {
						Description: "IPv6 prefix for the interface.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"mac_address": {
						Description: "MAC address configured for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"management_controller": {
						Description: "A reference to a managementController resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"mask": {
						Description: "Netmask for the interface.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
					"rn": {
						Description: "The Relative Name uniquely identifies an object within a given context.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"switch_id": {
						Description: "Switch Id connected to the interface.",
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
					"uem_conn_status": {
						Description: "The event channel connection status for the interface.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"virtual_host_name": {
						Description: "Virtual hostname configured for the interface in case of clustered environment.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vlan_id": {
						Description: "VlanId configured for the interface.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceManagementInterfaceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ManagementInterface{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("gateway"); ok {
		x := (v.(string))
		o.SetGateway(x)
	}
	if v, ok := d.GetOk("host_name"); ok {
		x := (v.(string))
		o.SetHostName(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.SetIpv4Address(x)
	}
	if v, ok := d.GetOk("ipv4_gateway"); ok {
		x := (v.(string))
		o.SetIpv4Gateway(x)
	}
	if v, ok := d.GetOk("ipv4_mask"); ok {
		x := (v.(string))
		o.SetIpv4Mask(x)
	}
	if v, ok := d.GetOk("ipv6_address"); ok {
		x := (v.(string))
		o.SetIpv6Address(x)
	}
	if v, ok := d.GetOk("ipv6_gateway"); ok {
		x := (v.(string))
		o.SetIpv6Gateway(x)
	}
	if v, ok := d.GetOk("ipv6_prefix"); ok {
		x := int64(v.(int))
		o.SetIpv6Prefix(x)
	}
	if v, ok := d.GetOk("mac_address"); ok {
		x := (v.(string))
		o.SetMacAddress(x)
	}
	if v, ok := d.GetOk("mask"); ok {
		x := (v.(string))
		o.SetMask(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}
	if v, ok := d.GetOk("switch_id"); ok {
		x := (v.(string))
		o.SetSwitchId(x)
	}
	if v, ok := d.GetOk("uem_conn_status"); ok {
		x := (v.(string))
		o.SetUemConnStatus(x)
	}
	if v, ok := d.GetOk("virtual_host_name"); ok {
		x := (v.(string))
		o.SetVirtualHostName(x)
	}
	if v, ok := d.GetOk("vlan_id"); ok {
		x := int64(v.(int))
		o.SetVlanId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ManagementInterface object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ManagementApi.GetManagementInterfaceList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ManagementInterface: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ManagementInterfaceList.GetCount()
	var i int32
	var managementInterfaceResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ManagementApi.GetManagementInterfaceList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ManagementInterface: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ManagementInterfaceList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ManagementInterface data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())
				temp["gateway"] = (s.GetGateway())
				temp["host_name"] = (s.GetHostName())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)
				temp["ip_address"] = (s.GetIpAddress())
				temp["ipv4_address"] = (s.GetIpv4Address())
				temp["ipv4_gateway"] = (s.GetIpv4Gateway())
				temp["ipv4_mask"] = (s.GetIpv4Mask())
				temp["ipv6_address"] = (s.GetIpv6Address())
				temp["ipv6_gateway"] = (s.GetIpv6Gateway())
				temp["ipv6_prefix"] = (s.GetIpv6Prefix())
				temp["mac_address"] = (s.GetMacAddress())

				temp["management_controller"] = flattenMapManagementControllerRelationship(s.GetManagementController(), d)
				temp["mask"] = (s.GetMask())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["rn"] = (s.GetRn())
				temp["switch_id"] = (s.GetSwitchId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["uem_conn_status"] = (s.GetUemConnStatus())
				temp["virtual_host_name"] = (s.GetVirtualHostName())
				temp["vlan_id"] = (s.GetVlanId())
				managementInterfaceResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(managementInterfaceResults))
	if err := d.Set("results", managementInterfaceResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(managementInterfaceResults[0]["moid"].(string))
	return de
}
