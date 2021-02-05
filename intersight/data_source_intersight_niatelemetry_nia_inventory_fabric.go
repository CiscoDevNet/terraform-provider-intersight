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

func dataSourceNiatelemetryNiaInventoryFabric() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryFabricRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
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
		},
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
	if v, ok := d.GetOk("spine_count"); ok {
		x := int64(v.(int))
		o.SetSpineCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventoryFabric object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryFabricList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryFabric: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NiatelemetryNiaInventoryFabric list: %s", err.Error())
	}
	var s = &models.NiatelemetryNiaInventoryFabricList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NiatelemetryNiaInventoryFabric list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NiatelemetryNiaInventoryFabric data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NiatelemetryNiaInventoryFabric data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiatelemetryNiaInventoryFabric{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("anycast_gw_mac", (s.GetAnycastGwMac())); err != nil {
				return diag.Errorf("error occurred while setting property AnycastGwMac: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("dcnmtracker_enabled", (s.GetDcnmtrackerEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property DcnmtrackerEnabled: %s", err.Error())
			}
			if err := d.Set("fabric_id", (s.GetFabricId())); err != nil {
				return diag.Errorf("error occurred while setting property FabricId: %s", err.Error())
			}
			if err := d.Set("fabric_name", (s.GetFabricName())); err != nil {
				return diag.Errorf("error occurred while setting property FabricName: %s", err.Error())
			}
			if err := d.Set("is_ngoam_enabled", (s.GetIsNgoamEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property IsNgoamEnabled: %s", err.Error())
			}
			if err := d.Set("is_scheduled_back_up_enabled", (s.GetIsScheduledBackUpEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property IsScheduledBackUpEnabled: %s", err.Error())
			}
			if err := d.Set("leaf_count", (s.GetLeafCount())); err != nil {
				return diag.Errorf("error occurred while setting property LeafCount: %s", err.Error())
			}

			if err := d.Set("logical_links", flattenListNiatelemetryLogicalLink(s.GetLogicalLinks(), d)); err != nil {
				return diag.Errorf("error occurred while setting property LogicalLinks: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("nxos_vrf_count", (s.GetNxosVrfCount())); err != nil {
				return diag.Errorf("error occurred while setting property NxosVrfCount: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("spine_count", (s.GetSpineCount())); err != nil {
				return diag.Errorf("error occurred while setting property SpineCount: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
