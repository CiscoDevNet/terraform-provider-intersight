package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryEpg() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryEpgRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"azure_pack_count": {
				Description: "Azure Pack NAT with ASA feature usage.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "Dn value for the End Point Groups present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"epg_delimiter_count": {
				Description: "EPG Delimiter scale where the delimiter value is present.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fc_npv_count": {
				Description: "Number of ports with FC path attribute of type FC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fcoe_count": {
				Description: "Number of FCoE per End Point Group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fv_rs_dom_att_count": {
				Description: "FvRsDomAtt scale per End Point Group with VMware configured.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"intra_epg_dvs_bm_count": {
				Description: "Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"intra_epg_hyperv": {
				Description: "Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_attr_based": {
				Description: "Gets the state of End Point Groups with isAttrBasedEPg value as configured.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"microsegmentation": {
				Description: "Gets the state of End Point Groups where microsegmentation is present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"microsoft_useg_count": {
				Description: "FvRsDomAtt scale per End Point Group with Microsoft configured.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name value for the End Point Groups present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orchsl_dev_vip_cfg_count": {
				Description: "Simplified Service Graph Integration with Windows Azure Pack count scale.",
				Type:        schema.TypeInt,
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
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
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
			"site_name": {
				Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
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
			"useg_hyperv_count": {
				Description: "Logical Operators for attribute based microsegmentation in a hypervisor.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}

func dataSourceNiatelemetryEpgRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNiatelemetryEpgWithDefaults()
	if v, ok := d.GetOk("azure_pack_count"); ok {
		x := int64(v.(int))
		o.SetAzurePackCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("epg_delimiter_count"); ok {
		x := int64(v.(int))
		o.SetEpgDelimiterCount(x)
	}
	if v, ok := d.GetOk("fc_npv_count"); ok {
		x := int64(v.(int))
		o.SetFcNpvCount(x)
	}
	if v, ok := d.GetOk("fcoe_count"); ok {
		x := (v.(string))
		o.SetFcoeCount(x)
	}
	if v, ok := d.GetOk("fv_rs_dom_att_count"); ok {
		x := int64(v.(int))
		o.SetFvRsDomAttCount(x)
	}
	if v, ok := d.GetOk("intra_epg_dvs_bm_count"); ok {
		x := int64(v.(int))
		o.SetIntraEpgDvsBmCount(x)
	}
	if v, ok := d.GetOk("intra_epg_hyperv"); ok {
		x := (v.(string))
		o.SetIntraEpgHyperv(x)
	}
	if v, ok := d.GetOk("is_attr_based"); ok {
		x := (v.(string))
		o.SetIsAttrBased(x)
	}
	if v, ok := d.GetOk("microsegmentation"); ok {
		x := (v.(string))
		o.SetMicrosegmentation(x)
	}
	if v, ok := d.GetOk("microsoft_useg_count"); ok {
		x := int64(v.(int))
		o.SetMicrosoftUsegCount(x)
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
	if v, ok := d.GetOk("orchsl_dev_vip_cfg_count"); ok {
		x := int64(v.(int))
		o.SetOrchslDevVipCfgCount(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("useg_hyperv_count"); ok {
		x := int64(v.(int))
		o.SetUsegHypervCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiatelemetryApi.GetNiatelemetryEpgList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiatelemetryEpgList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiatelemetryEpg: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiatelemetryEpgWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("azure_pack_count", (s.GetAzurePackCount())); err != nil {
				return fmt.Errorf("error occurred while setting property AzurePackCount: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}
			if err := d.Set("epg_delimiter_count", (s.GetEpgDelimiterCount())); err != nil {
				return fmt.Errorf("error occurred while setting property EpgDelimiterCount: %+v", err)
			}
			if err := d.Set("fc_npv_count", (s.GetFcNpvCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FcNpvCount: %+v", err)
			}
			if err := d.Set("fcoe_count", (s.GetFcoeCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FcoeCount: %+v", err)
			}
			if err := d.Set("fv_rs_dom_att_count", (s.GetFvRsDomAttCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FvRsDomAttCount: %+v", err)
			}
			if err := d.Set("intra_epg_dvs_bm_count", (s.GetIntraEpgDvsBmCount())); err != nil {
				return fmt.Errorf("error occurred while setting property IntraEpgDvsBmCount: %+v", err)
			}
			if err := d.Set("intra_epg_hyperv", (s.GetIntraEpgHyperv())); err != nil {
				return fmt.Errorf("error occurred while setting property IntraEpgHyperv: %+v", err)
			}
			if err := d.Set("is_attr_based", (s.GetIsAttrBased())); err != nil {
				return fmt.Errorf("error occurred while setting property IsAttrBased: %+v", err)
			}
			if err := d.Set("microsegmentation", (s.GetMicrosegmentation())); err != nil {
				return fmt.Errorf("error occurred while setting property Microsegmentation: %+v", err)
			}
			if err := d.Set("microsoft_useg_count", (s.GetMicrosoftUsegCount())); err != nil {
				return fmt.Errorf("error occurred while setting property MicrosoftUsegCount: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("orchsl_dev_vip_cfg_count", (s.GetOrchslDevVipCfgCount())); err != nil {
				return fmt.Errorf("error occurred while setting property OrchslDevVipCfgCount: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("site_name", (s.GetSiteName())); err != nil {
				return fmt.Errorf("error occurred while setting property SiteName: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("useg_hyperv_count", (s.GetUsegHypervCount())); err != nil {
				return fmt.Errorf("error occurred while setting property UsegHypervCount: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
