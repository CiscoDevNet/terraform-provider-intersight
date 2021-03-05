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

func dataSourceNiaapiApicHweol() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiaapiApicHweolRead,
		Schema: map[string]*schema.Schema{
			"affected_pids": {
				Description: "String contains the PID of hardwares affected by this notice, seperated by comma.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date": {
				Description: "When this notice is announced.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"announcement_date_epoch": {
				Description: "Epoch time of Announcement Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bulletin_no": {
				Description: "The bulletinno of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The description of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date": {
				Description: "Date time of end of new services attachment  .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_new_service_attachment_date_epoch": {
				Description: "Epoch time of New service attachment Date .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_routine_failure_analysis_date": {
				Description: "Date time of end of routinefailure analysis.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_routine_failure_analysis_date_epoch": {
				Description: "Epoch time of End of Routine Failure Analysis Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sale_date": {
				Description: "When this hardware will end sale.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sale_date_epoch": {
				Description: "Epoch time of End of Sale Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_security_support": {
				Description: "Date time of end of security support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_security_support_epoch": {
				Description: "Epoch time of End of Security Support Date .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_service_contract_renewal_date": {
				Description: "Date time of end of service contract renew .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_service_contract_renewal_date_epoch": {
				Description: "Epoch time of End of Renewal service contract.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"endof_sw_maintenance_date": {
				Description: "The date of end of maintainance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"endof_sw_maintenance_date_epoch": {
				Description: "Epoch time of End of maintenance Date.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"hardware_eol_url": {
				Description: "Hardware end of sale URL link to the notice webpage.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"headline": {
				Description: "The title of this hardware end of life notice.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support": {
				Description: "Date time of end of support .",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_dateof_support_epoch": {
				Description: "Epoch time of last date of support .",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"last_ship_date": {
				Description: "Date time of Lastship Date.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_ship_date_epoch": {
				Description: "Epoch time of last ship Date.",
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
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"affected_pids": {
						Description: "String contains the PID of hardwares affected by this notice, seperated by comma.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"announcement_date": {
						Description: "When this notice is announced.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"announcement_date_epoch": {
						Description: "Epoch time of Announcement Date.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"bulletin_no": {
						Description: "The bulletinno of this hardware end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"description": {
						Description: "The description of this hardware end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_new_service_attachment_date": {
						Description: "Date time of end of new services attachment  .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_new_service_attachment_date_epoch": {
						Description: "Epoch time of New service attachment Date .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_routine_failure_analysis_date": {
						Description: "Date time of end of routinefailure analysis.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_routine_failure_analysis_date_epoch": {
						Description: "Epoch time of End of Routine Failure Analysis Date.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_sale_date": {
						Description: "When this hardware will end sale.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_sale_date_epoch": {
						Description: "Epoch time of End of Sale Date.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_security_support": {
						Description: "Date time of end of security support .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_security_support_epoch": {
						Description: "Epoch time of End of Security Support Date .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_service_contract_renewal_date": {
						Description: "Date time of end of service contract renew .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_service_contract_renewal_date_epoch": {
						Description: "Epoch time of End of Renewal service contract.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"endof_sw_maintenance_date": {
						Description: "The date of end of maintainance.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"endof_sw_maintenance_date_epoch": {
						Description: "Epoch time of End of maintenance Date.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"hardware_eol_url": {
						Description: "Hardware end of sale URL link to the notice webpage.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"headline": {
						Description: "The title of this hardware end of life notice.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_dateof_support": {
						Description: "Date time of end of support .",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_dateof_support_epoch": {
						Description: "Epoch time of last date of support .",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"last_ship_date": {
						Description: "Date time of Lastship Date.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"last_ship_date_epoch": {
						Description: "Epoch time of last ship Date.",
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

func dataSourceNiaapiApicHweolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiaapiApicHweol{}
	if v, ok := d.GetOk("affected_pids"); ok {
		x := (v.(string))
		o.SetAffectedPids(x)
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
	if v, ok := d.GetOk("endof_routine_failure_analysis_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofRoutineFailureAnalysisDate(x)
	}
	if v, ok := d.GetOk("endof_routine_failure_analysis_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofRoutineFailureAnalysisDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_sale_date"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofSaleDate(x)
	}
	if v, ok := d.GetOk("endof_sale_date_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofSaleDateEpoch(x)
	}
	if v, ok := d.GetOk("endof_security_support"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndofSecuritySupport(x)
	}
	if v, ok := d.GetOk("endof_security_support_epoch"); ok {
		x := int64(v.(int))
		o.SetEndofSecuritySupportEpoch(x)
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
	if v, ok := d.GetOk("hardware_eol_url"); ok {
		x := (v.(string))
		o.SetHardwareEolUrl(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiaapiApicHweol object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiApicHweolList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiaapiApicHweol: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiaapiApicHweolList.GetCount()
	var i int32
	var niaapiApicHweolResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiaapiApi.GetNiaapiApicHweolList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiaapiApicHweol: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiaapiApicHweolList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiaapiApicHweol data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["affected_pids"] = (s.GetAffectedPids())

				temp["announcement_date"] = (s.GetAnnouncementDate()).String()
				temp["announcement_date_epoch"] = (s.GetAnnouncementDateEpoch())
				temp["bulletin_no"] = (s.GetBulletinNo())
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["endof_new_service_attachment_date"] = (s.GetEndofNewServiceAttachmentDate()).String()
				temp["endof_new_service_attachment_date_epoch"] = (s.GetEndofNewServiceAttachmentDateEpoch())

				temp["endof_routine_failure_analysis_date"] = (s.GetEndofRoutineFailureAnalysisDate()).String()
				temp["endof_routine_failure_analysis_date_epoch"] = (s.GetEndofRoutineFailureAnalysisDateEpoch())

				temp["endof_sale_date"] = (s.GetEndofSaleDate()).String()
				temp["endof_sale_date_epoch"] = (s.GetEndofSaleDateEpoch())

				temp["endof_security_support"] = (s.GetEndofSecuritySupport()).String()
				temp["endof_security_support_epoch"] = (s.GetEndofSecuritySupportEpoch())

				temp["endof_service_contract_renewal_date"] = (s.GetEndofServiceContractRenewalDate()).String()
				temp["endof_service_contract_renewal_date_epoch"] = (s.GetEndofServiceContractRenewalDateEpoch())

				temp["endof_sw_maintenance_date"] = (s.GetEndofSwMaintenanceDate()).String()
				temp["endof_sw_maintenance_date_epoch"] = (s.GetEndofSwMaintenanceDateEpoch())
				temp["hardware_eol_url"] = (s.GetHardwareEolUrl())
				temp["headline"] = (s.GetHeadline())

				temp["last_dateof_support"] = (s.GetLastDateofSupport()).String()
				temp["last_dateof_support_epoch"] = (s.GetLastDateofSupportEpoch())

				temp["last_ship_date"] = (s.GetLastShipDate()).String()
				temp["last_ship_date_epoch"] = (s.GetLastShipDateEpoch())
				temp["migration_url"] = (s.GetMigrationUrl())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				niaapiApicHweolResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niaapiApicHweolResults))
	if err := d.Set("results", niaapiApicHweolResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niaapiApicHweolResults[0]["moid"].(string))
	return de
}
