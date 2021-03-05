package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiatelemetryNiaInventoryFabric() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryFabricRead,
		Schema: map[string]*schema.Schema{
			"anycast_gw_mac": {
				Description: "Returns the aycast gateway mac.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dcnmtracker_enabled": {
				Description: "Returns the value of the dcnmtrackerEnabled field.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"fabric_id": {
				Description: "Uniquely identifies a fabric.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fabric_name": {
				Description: "Returns the value of the Name of a fabric.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_ngoam_enabled": {
				Description: "Returns if ngoam is enabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_scheduled_back_up_enabled": {
				Description: "Returns if the scheduled backup is enabled.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"leaf_count": {
				Description: "Returns total number of leafs in the fabric.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nxos_vrf_count": {
				Description: "Returns the value of the nxosVrfCount field.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "Serial number of device being inventoried. The serial number is unique per device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name": {
				Description: "Name of fabric domain of the controller.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"spine_count": {
				Description: "Returns total number of spines in the fabric.",
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
					"anycast_gw_mac": {
						Description: "Returns the aycast gateway mac.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"dcnmtracker_enabled": {
						Description: "Returns the value of the dcnmtrackerEnabled field.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"fabric_id": {
						Description: "Uniquely identifies a fabric.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"fabric_name": {
						Description: "Returns the value of the Name of a fabric.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_ngoam_enabled": {
						Description: "Returns if ngoam is enabled.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_scheduled_back_up_enabled": {
						Description: "Returns if the scheduled backup is enabled.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"leaf_count": {
						Description: "Returns total number of leafs in the fabric.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"logical_links": {
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"db_id": {
									Description: "Return value of dbId attribute.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"is_present": {
									Description: "Return value of isPresent attribute.",
									Type:        schema.TypeBool,
									Optional:    true,
								},
								"link_addr1": {
									Description: "Return value of linkAddr1 attribute.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"link_addr2": {
									Description: "Return value of linkAddr2 attribute.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"link_state": {
									Description: "Return value of linkState attribute.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"link_type": {
									Description: "Return value of linkType attribute.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"uptime": {
									Description: "Return value of uptime attribute.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"nxos_vrf_count": {
						Description: "Returns the value of the nxosVrfCount field.",
						Type:        schema.TypeInt,
						Optional:    true,
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
					"serial": {
						Description: "Serial number of device being inventoried. The serial number is unique per device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"site_name": {
						Description: "Name of fabric domain of the controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"spine_count": {
						Description: "Returns total number of spines in the fabric.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiatelemetryNiaInventoryFabricRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventoryFabric{}
	if v, ok := d.GetOk("anycast_gw_mac"); ok {
		x := (v.(string))
		o.SetAnycastGwMac(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("dcnmtracker_enabled"); ok {
		x := (v.(bool))
		o.SetDcnmtrackerEnabled(x)
	}
	if v, ok := d.GetOk("fabric_id"); ok {
		x := (v.(string))
		o.SetFabricId(x)
	}
	if v, ok := d.GetOk("fabric_name"); ok {
		x := (v.(string))
		o.SetFabricName(x)
	}
	if v, ok := d.GetOk("is_ngoam_enabled"); ok {
		x := (v.(string))
		o.SetIsNgoamEnabled(x)
	}
	if v, ok := d.GetOk("is_scheduled_back_up_enabled"); ok {
		x := (v.(string))
		o.SetIsScheduledBackUpEnabled(x)
	}
	if v, ok := d.GetOk("leaf_count"); ok {
		x := int64(v.(int))
		o.SetLeafCount(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("nxos_vrf_count"); ok {
		x := int64(v.(int))
		o.SetNxosVrfCount(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("spine_count"); ok {
		x := int64(v.(int))
		o.SetSpineCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventoryFabric object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryFabricList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryFabric: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryNiaInventoryFabricList.GetCount()
	var i int32
	var niatelemetryNiaInventoryFabricResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryFabricList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryFabric: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryNiaInventoryFabricList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryNiaInventoryFabric data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["anycast_gw_mac"] = (s.GetAnycastGwMac())
				temp["class_id"] = (s.GetClassId())
				temp["dcnmtracker_enabled"] = (s.GetDcnmtrackerEnabled())
				temp["fabric_id"] = (s.GetFabricId())
				temp["fabric_name"] = (s.GetFabricName())
				temp["is_ngoam_enabled"] = (s.GetIsNgoamEnabled())
				temp["is_scheduled_back_up_enabled"] = (s.GetIsScheduledBackUpEnabled())
				temp["leaf_count"] = (s.GetLeafCount())

				temp["logical_links"] = flattenListNiatelemetryLogicalLink(s.GetLogicalLinks(), d)
				temp["moid"] = (s.GetMoid())
				temp["nxos_vrf_count"] = (s.GetNxosVrfCount())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial"] = (s.GetSerial())
				temp["site_name"] = (s.GetSiteName())
				temp["spine_count"] = (s.GetSpineCount())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				niatelemetryNiaInventoryFabricResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaInventoryFabricResults))
	if err := d.Set("results", niatelemetryNiaInventoryFabricResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaInventoryFabricResults[0]["moid"].(string))
	return de
}
