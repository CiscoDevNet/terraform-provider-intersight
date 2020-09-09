package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiaapiDcnmSweol() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiaapiDcnmSweolRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
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
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
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
		},
	}
}

func dataSourceNiaapiDcnmSweolRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNiaapiDcnmSweolWithDefaults()
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiaapiApi.GetNiaapiDcnmSweolList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiaapiDcnmSweolList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiaapiDcnmSweol: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiaapiDcnmSweolWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("affected_versions", (s.GetAffectedVersions())); err != nil {
				return fmt.Errorf("error occurred while setting property AffectedVersions: %+v", err)
			}

			if err := d.Set("announcement_date", (s.GetAnnouncementDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property AnnouncementDate: %+v", err)
			}
			if err := d.Set("announcement_date_epoch", (s.GetAnnouncementDateEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property AnnouncementDateEpoch: %+v", err)
			}
			if err := d.Set("bulletin_no", (s.GetBulletinNo())); err != nil {
				return fmt.Errorf("error occurred while setting property BulletinNo: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}

			if err := d.Set("endof_new_service_attachment_date", (s.GetEndofNewServiceAttachmentDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndofNewServiceAttachmentDate: %+v", err)
			}
			if err := d.Set("endof_new_service_attachment_date_epoch", (s.GetEndofNewServiceAttachmentDateEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property EndofNewServiceAttachmentDateEpoch: %+v", err)
			}

			if err := d.Set("endof_service_contract_renewal_date", (s.GetEndofServiceContractRenewalDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndofServiceContractRenewalDate: %+v", err)
			}
			if err := d.Set("endof_service_contract_renewal_date_epoch", (s.GetEndofServiceContractRenewalDateEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property EndofServiceContractRenewalDateEpoch: %+v", err)
			}

			if err := d.Set("endof_sw_maintenance_date", (s.GetEndofSwMaintenanceDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndofSwMaintenanceDate: %+v", err)
			}
			if err := d.Set("endof_sw_maintenance_date_epoch", (s.GetEndofSwMaintenanceDateEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property EndofSwMaintenanceDateEpoch: %+v", err)
			}
			if err := d.Set("headline", (s.GetHeadline())); err != nil {
				return fmt.Errorf("error occurred while setting property Headline: %+v", err)
			}

			if err := d.Set("last_dateof_support", (s.GetLastDateofSupport()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastDateofSupport: %+v", err)
			}
			if err := d.Set("last_dateof_support_epoch", (s.GetLastDateofSupportEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property LastDateofSupportEpoch: %+v", err)
			}

			if err := d.Set("last_ship_date", (s.GetLastShipDate()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastShipDate: %+v", err)
			}
			if err := d.Set("last_ship_date_epoch", (s.GetLastShipDateEpoch())); err != nil {
				return fmt.Errorf("error occurred while setting property LastShipDateEpoch: %+v", err)
			}
			if err := d.Set("migration_url", (s.GetMigrationUrl())); err != nil {
				return fmt.Errorf("error occurred while setting property MigrationUrl: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("software_eol_url", (s.GetSoftwareEolUrl())); err != nil {
				return fmt.Errorf("error occurred while setting property SoftwareEolUrl: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
