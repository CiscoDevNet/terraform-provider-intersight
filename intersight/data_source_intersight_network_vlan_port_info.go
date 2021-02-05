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

func dataSourceNetworkVlanPortInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkVlanPortInfoRead,
		Schema: map[string]*schema.Schema{
			"access_vlan_port_count": {
				Description: "The number of available VLAN access ports on a Fabric Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"border_vlan_port_count": {
				Description: "The number of available VLAN border ports on a Fabric Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"compressed_optimization_sets_value": {
				Description: "The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"compressed_vlan_port_count": {
				Description: "The number of compressed VLAN ports on a Fabric Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"compressed_vlan_port_count_value": {
				Description: "The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.",
				Type:        schema.TypeInt,
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"network_element": {
				Description: "A reference to a networkElement resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"total_vlan_port_count": {
				Description: "The total number of VLAN ports on a Fabric Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"uncompressed_vlan_port_count": {
				Description: "The number of uncompressed VLAN ports on a Fabric Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uncompressed_vlan_port_count_value": {
				Description: "The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vlan_port_limit": {
				Description: "The maximum number of VLAN ports allowed on a Fabric Interconnect.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceNetworkVlanPortInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkVlanPortInfo{}
	if v, ok := d.GetOk("access_vlan_port_count"); ok {
		x := int64(v.(int))
		o.SetAccessVlanPortCount(x)
	}
	if v, ok := d.GetOk("border_vlan_port_count"); ok {
		x := int64(v.(int))
		o.SetBorderVlanPortCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("compressed_optimization_sets_value"); ok {
		x := int64(v.(int))
		o.SetCompressedOptimizationSetsValue(x)
	}
	if v, ok := d.GetOk("compressed_vlan_port_count"); ok {
		x := (v.(string))
		o.SetCompressedVlanPortCount(x)
	}
	if v, ok := d.GetOk("compressed_vlan_port_count_value"); ok {
		x := int64(v.(int))
		o.SetCompressedVlanPortCountValue(x)
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
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
	if v, ok := d.GetOk("total_vlan_port_count"); ok {
		x := int64(v.(int))
		o.SetTotalVlanPortCount(x)
	}
	if v, ok := d.GetOk("uncompressed_vlan_port_count"); ok {
		x := (v.(string))
		o.SetUncompressedVlanPortCount(x)
	}
	if v, ok := d.GetOk("uncompressed_vlan_port_count_value"); ok {
		x := int64(v.(int))
		o.SetUncompressedVlanPortCountValue(x)
	}
	if v, ok := d.GetOk("vlan_port_limit"); ok {
		x := int64(v.(int))
		o.SetVlanPortLimit(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NetworkVlanPortInfo object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NetworkApi.GetNetworkVlanPortInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching NetworkVlanPortInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NetworkVlanPortInfo list: %s", err.Error())
	}
	var s = &models.NetworkVlanPortInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NetworkVlanPortInfo list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NetworkVlanPortInfo data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NetworkVlanPortInfo data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NetworkVlanPortInfo{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("access_vlan_port_count", (s.GetAccessVlanPortCount())); err != nil {
				return diag.Errorf("error occurred while setting property AccessVlanPortCount: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("border_vlan_port_count", (s.GetBorderVlanPortCount())); err != nil {
				return diag.Errorf("error occurred while setting property BorderVlanPortCount: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("compressed_optimization_sets_value", (s.GetCompressedOptimizationSetsValue())); err != nil {
				return diag.Errorf("error occurred while setting property CompressedOptimizationSetsValue: %s", err.Error())
			}
			if err := d.Set("compressed_vlan_port_count", (s.GetCompressedVlanPortCount())); err != nil {
				return diag.Errorf("error occurred while setting property CompressedVlanPortCount: %s", err.Error())
			}
			if err := d.Set("compressed_vlan_port_count_value", (s.GetCompressedVlanPortCountValue())); err != nil {
				return diag.Errorf("error occurred while setting property CompressedVlanPortCountValue: %s", err.Error())
			}
			if err := d.Set("device_mo_id", (s.GetDeviceMoId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceMoId: %s", err.Error())
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return diag.Errorf("error occurred while setting property Dn: %s", err.Error())
			}

			if err := d.Set("inventory_device_info", flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InventoryDeviceInfo: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}

			if err := d.Set("network_element", flattenMapNetworkElementRelationship(s.GetNetworkElement(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NetworkElement: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("rn", (s.GetRn())); err != nil {
				return diag.Errorf("error occurred while setting property Rn: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("total_vlan_port_count", (s.GetTotalVlanPortCount())); err != nil {
				return diag.Errorf("error occurred while setting property TotalVlanPortCount: %s", err.Error())
			}
			if err := d.Set("uncompressed_vlan_port_count", (s.GetUncompressedVlanPortCount())); err != nil {
				return diag.Errorf("error occurred while setting property UncompressedVlanPortCount: %s", err.Error())
			}
			if err := d.Set("uncompressed_vlan_port_count_value", (s.GetUncompressedVlanPortCountValue())); err != nil {
				return diag.Errorf("error occurred while setting property UncompressedVlanPortCountValue: %s", err.Error())
			}
			if err := d.Set("vlan_port_limit", (s.GetVlanPortLimit())); err != nil {
				return diag.Errorf("error occurred while setting property VlanPortLimit: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
