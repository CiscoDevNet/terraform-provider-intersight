package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNiaapiApicSweol() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiApicSweolRead,
		Schema: map[string]*schema.Schema{
			"affected_versions": {
				Description: "String contains the Release versions affected by this notice, seperated by comma.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date": {
				Description: "Date time of this notice Announced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date_epoch": {
				Description: "Epoch time of this notice Announced.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bulletin_no": {
				Description: "The bulletinno of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The description of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date": {
				Description: "Date time of End of New service attachment .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date_epoch": {
				Description: "Epoch time of End of New service attachment .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_service_contract_renewal_date": {
				Description: "Date time of End of Renewal of service contract .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_service_contract_renewal_date_epoch": {
				Description: "Epoch time of End of Renewal of service contract .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sw_maintenance_date": {
				Description: "Date time of End of Maintenance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sw_maintenance_date_epoch": {
				Description: "Epoch time of End of Maintenance.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"headline": {
				Description: "The title of this software release end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support": {
				Description: "Date time of Last day of Support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support_epoch": {
				Description: "Epoch time of Last day of Support .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"last_ship_date": {
				Description: "Date time of Lastship Date.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_ship_date_epoch": {
				Description: "Epoch time of Lastship Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"migration_url": {
				Description: "The URL of this migration notice.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"software_eol_url": {
				Description: "Software end of life notice URL link to the notice webpage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"affected_versions": {
						Description: "String contains the Release versions affected by this notice, seperated by comma.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"announcement_date": {
						Description: "Date time of this notice Announced.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"announcement_date_epoch": {
						Description: "Epoch time of this notice Announced.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"bulletin_no": {
						Description: "The bulletinno of this software release end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"description": {
						Description: "The description of this software release end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_new_service_attachment_date": {
						Description: "Date time of End of New service attachment .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_new_service_attachment_date_epoch": {
						Description: "Epoch time of End of New service attachment .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_service_contract_renewal_date": {
						Description: "Date time of End of Renewal of service contract .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_service_contract_renewal_date_epoch": {
						Description: "Epoch time of End of Renewal of service contract .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_sw_maintenance_date": {
						Description: "Date time of End of Maintenance.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_sw_maintenance_date_epoch": {
						Description: "Epoch time of End of Maintenance.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"headline": {
						Description: "The title of this software release end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_dateof_support": {
						Description: "Date time of Last day of Support .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_dateof_support_epoch": {
						Description: "Epoch time of Last day of Support .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"last_ship_date": {
						Description: "Date time of Lastship Date.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_ship_date_epoch": {
						Description: "Epoch time of Lastship Date.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"migration_url": {
						Description: "The URL of this migration notice.",
						Type:        schema.TypeString,
						Optional:    true,
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
					"software_eol_url": {
						Description: "Software end of life notice URL link to the notice webpage.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceNiaapiApicSweolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiApicSweol{}
	if v, ok := d.GetOk("affected_versions"); ok {
		x := (v.(string))
		o.SetAffectedVersions(x)
	}
	if v, ok := d.GetOk("announcement_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetAnnouncementDate(x)
	}
	if v, ok := d.GetOk("announcement_date_epoch"); ok {
		x := int64(v.(int))
		o.SetAnnouncementDateEpoch(x)
	}
	if v, ok := d.GetOk("bulletin_no"); ok {
		x := (v.(string))
		o.SetBulletinNo(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofNewServiceAttachmentDate(x)
	}
	if v, ok := d.GetOk("endof_new_service_attachment_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofNewServiceAttachmentDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofServiceContractRenewalDate(x)
	}
	if v, ok := d.GetOk("endof_service_contract_renewal_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofServiceContractRenewalDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofSwMaintenanceDate(x)
	}
	if v, ok := d.GetOk("endof_sw_maintenance_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofSwMaintenanceDateEpoch(x)
	}
	if v, ok := d.GetOk("headline"); ok {
		x := (v.(string))
		o.SetHeadline(x)
	}
	if v, ok := d.GetOk("last_dateof_support"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastDateofSupport(x)
	}
	if v, ok := d.GetOk("last_dateof_support_epoch"); ok {
		x := int64(v.(int))
		o.SetLastDateofSupportEpoch(x)
	}
	if v, ok := d.GetOk("last_ship_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastShipDate(x)
	}
	if v, ok := d.GetOk("last_ship_date_epoch"); ok {
		x := int64(v.(int))
		o.SetLastShipDateEpoch(x)
	}
	if v, ok := d.GetOk("migration_url"); ok {
		x := (v.(string))
		o.SetMigrationUrl(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("software_eol_url"); ok {
		x := (v.(string))
		o.SetSoftwareEolUrl(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiApicSweol object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiApicSweolList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiaapiApicSweol: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiaapiApicSweolList.GetCount()
	var i int32
	var niaapiApicSweolResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiApicSweolList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiaapiApicSweol: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiaapiApicSweolList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiaapiApicSweol data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["affected_versions"] = (s.GetAffectedVersions())

				temp["announcement_date"] = (s.GetAnnouncementDate()).String()
				temp["announcement_date_epoch"] = (s.GetAnnouncementDateEpoch())
				temp["bulletin_no"] = (s.GetBulletinNo())
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["endof_new_service_attachment_date"] = (s.GetEndofNewServiceAttachmentDate()).String()
				temp["endof_new_service_attachment_date_epoch"] = (s.GetEndofNewServiceAttachmentDateEpoch())

				temp["endof_service_contract_renewal_date"] = (s.GetEndofServiceContractRenewalDate()).String()
				temp["endof_service_contract_renewal_date_epoch"] = (s.GetEndofServiceContractRenewalDateEpoch())

				temp["endof_sw_maintenance_date"] = (s.GetEndofSwMaintenanceDate()).String()
				temp["endof_sw_maintenance_date_epoch"] = (s.GetEndofSwMaintenanceDateEpoch())
				temp["headline"] = (s.GetHeadline())

				temp["last_dateof_support"] = (s.GetLastDateofSupport()).String()
				temp["last_dateof_support_epoch"] = (s.GetLastDateofSupportEpoch())

				temp["last_ship_date"] = (s.GetLastShipDate()).String()
				temp["last_ship_date_epoch"] = (s.GetLastShipDateEpoch())
				temp["migration_url"] = (s.GetMigrationUrl())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["software_eol_url"] = (s.GetSoftwareEolUrl())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				niaapiApicSweolResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niaapiApicSweolResults))
	if err := d.Set("results", niaapiApicSweolResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niaapiApicSweolResults[0]["moid"].(string))
	return de
}
