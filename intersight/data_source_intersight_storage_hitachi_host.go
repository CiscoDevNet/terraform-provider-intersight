package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageHitachiHost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageHitachiHostRead,
		Schema: map[string]*schema.Schema{
			"authentication_mode": {
				Description: "Authentication mode for the iSCSI target.\n* `N/A` - Authentication Mode is not available.\n* `CHAP` - CHAP-authentication mode.\n* `NONE` - Authentication mode is not set.\n* `BOTH` - Comply with Host Setting.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Short description about the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_group_id": {
				Description: "ID of the host group.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_group_number": {
				Description: "Host group number for this host.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"is_chap_mutual": {
				Description: "Mutual CHAP setting that is Enabled or Disabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"iscsi_name": {
				Description: "IQN (iSCSI qualified name). Can be up to 255 characters long.",
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
				Description: "Name of the host in storage array.",
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
			"os_type": {
				Description: "Operating system running on the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port_id": {
				Description: "Port ID of the host group.",
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
			"type": {
				Description: "Host Group type, it can be FC or iSCSI.\n* `FC` - Port supports fibre channel protocol.\n* `iSCSI` - Port supports iSCSI protocol.\n* `FCoE` - Port supports fibre channel over ethernet protocol.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wwn": {
				Description: "World wide port name, 64 bit address represented in hexa decimal notation.",
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
					"authentication_mode": {
						Description: "Authentication mode for the iSCSI target.\n* `N/A` - Authentication Mode is not available.\n* `CHAP` - CHAP-authentication mode.\n* `NONE` - Authentication mode is not set.\n* `BOTH` - Comply with Host Setting.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"description": {
						Description: "Short description about the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"host_group_id": {
						Description: "ID of the host group.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"host_group_number": {
						Description: "Host group number for this host.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"host_mode_options": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeInt}},
					"initiators": {
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"iqn": {
									Description: "IQN (iSCSI qualified name). Can be up to 255 characters long and has the format iqn.yyyy-mm.naming-authority:unique name.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Name of the initiator represented in the storage array.",
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
								"type": {
									Description: "Initiator type, it can be FC or iSCSI.\n* `FC` - Fibre channel initiator type which contains WWN of an HBA on the host.\n* `iSCSI` - An iSCSI initiator type which contains the IQN (iSCSI Qualified Name) used by the host.\n* `Mixed` - Initiator type for systems using both FC and iSCSI connections.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"wwn": {
									Description: "World wide name, 128 bit address represented in hexadecimal notation. For example, 51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"is_chap_mutual": {
						Description: "Mutual CHAP setting that is Enabled or Disabled.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"iscsi_name": {
						Description: "IQN (iSCSI qualified name). Can be up to 255 characters long.",
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
						Description: "Name of the host in storage array.",
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
					"os_type": {
						Description: "Operating system running on the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"port_id": {
						Description: "Port ID of the host group.",
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
					"storage_utilization": {
						Description: "Storage utilization of host entity in storage array.",
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
								"available": {
									Description: "Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"capacity_utilization": {
									Description: "Percentage of used capacity.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Unused space available for applications to consume, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"total": {
									Description: "Total storage capacity, represented in bytes. It is set by the component manufacturer.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Used or consumed storage capacity, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
							},
						},
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
					"type": {
						Description: "Host Group type, it can be FC or iSCSI.\n* `FC` - Port supports fibre channel protocol.\n* `iSCSI` - Port supports iSCSI protocol.\n* `FCoE` - Port supports fibre channel over ethernet protocol.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"wwn": {
						Description: "World wide port name, 64 bit address represented in hexa decimal notation.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceStorageHitachiHostRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageHitachiHost{}
	if v, ok := d.GetOk("authentication_mode"); ok {
		x := (v.(string))
		o.SetAuthenticationMode(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("host_group_id"); ok {
		x := (v.(string))
		o.SetHostGroupId(x)
	}
	if v, ok := d.GetOk("host_group_number"); ok {
		x := int64(v.(int))
		o.SetHostGroupNumber(x)
	}
	if v, ok := d.GetOk("is_chap_mutual"); ok {
		x := (v.(bool))
		o.SetIsChapMutual(x)
	}
	if v, ok := d.GetOk("iscsi_name"); ok {
		x := (v.(string))
		o.SetIscsiName(x)
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
	if v, ok := d.GetOk("os_type"); ok {
		x := (v.(string))
		o.SetOsType(x)
	}
	if v, ok := d.GetOk("port_id"); ok {
		x := (v.(string))
		o.SetPortId(x)
	}
	if v, ok := d.GetOk("port_lun_security"); ok {
		x := (v.(bool))
		o.SetPortLunSecurity(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("wwn"); ok {
		x := (v.(string))
		o.SetWwn(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageHitachiHost object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiHostList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageHitachiHost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageHitachiHostList.GetCount()
	var i int32
	var storageHitachiHostResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageHitachiHostList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageHitachiHost: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageHitachiHostList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageHitachiHost data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["array"] = flattenMapStorageHitachiArrayRelationship(s.GetArray(), d)
				temp["authentication_mode"] = (s.GetAuthenticationMode())
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["host_group_id"] = (s.GetHostGroupId())
				temp["host_group_number"] = (s.GetHostGroupNumber())
				temp["host_mode_options"] = (s.GetHostModeOptions())

				temp["initiators"] = flattenListStorageBaseInitiator(s.GetInitiators(), d)
				temp["is_chap_mutual"] = (s.GetIsChapMutual())
				temp["iscsi_name"] = (s.GetIscsiName())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["os_type"] = (s.GetOsType())
				temp["port_id"] = (s.GetPortId())
				temp["port_lun_security"] = (s.GetPortLunSecurity())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["storage_utilization"] = flattenMapStorageBaseCapacity(s.GetStorageUtilization(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())
				temp["wwn"] = (s.GetWwn())
				storageHitachiHostResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageHitachiHostResults))
	if err := d.Set("results", storageHitachiHostResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageHitachiHostResults[0]["moid"].(string))
	return de
}
