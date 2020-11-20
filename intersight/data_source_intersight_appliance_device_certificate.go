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

func dataSourceApplianceDeviceCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceDeviceCertificateRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ca_certificate": {
				Description: "The base64 encoded certificate in PEM format.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ca_certificate_expiry_time": {
				Description: "The expiry datetime of new ca certificate which need to be applied on device connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"certificate_renewal_expiry_time": {
				Description: "The date time allocated till cert renewal will be executed. This time used here will be based on cert\nrenewal plan.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"completed_phases": {
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
						"end_time": {
							Description: "End date of the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"failed": {
							Description: "Indicates if the cert renewal phase has failed or not.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"message": {
							Description: "Status message set during the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the cert renewal phase phase.\n* `Init` - New certificate detected, cleanup the old process if any running.\n* `ScheduleCertificateAddOperation` - Certificate Add Operation Schedulled.\n* `WaitForCertCollectionByEndpoint` - Monitor cert collection by endpoint.\n* `Success` - Certificate Renewal Task Success.\n* `Fail` - Certificate Renewal Task Fail.\n* `Cancel` - Certificate Renewal Task Cancel.",
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
						"start_time": {
							Description: "Start date of the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"configuration_mo_id": {
				Description: "The operation configuration MOId.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"current_phase": {
				Description: "Current phase of the Intersight Appliance's Certificate Renewal.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"end_time": {
							Description: "End date of the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"failed": {
							Description: "Indicates if the cert renewal phase has failed or not.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"message": {
							Description: "Status message set during the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the cert renewal phase phase.\n* `Init` - New certificate detected, cleanup the old process if any running.\n* `ScheduleCertificateAddOperation` - Certificate Add Operation Schedulled.\n* `WaitForCertCollectionByEndpoint` - Monitor cert collection by endpoint.\n* `Success` - Certificate Renewal Task Success.\n* `Fail` - Certificate Renewal Task Fail.\n* `Cancel` - Certificate Renewal Task Cancel.",
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
						"start_time": {
							Description: "Start date of the cert renewal phase.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"end_time": {
				Description: "End date of the certificate renewal.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_success_poll_time": {
				Description: "The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
			"start_time": {
				Description: "Start date of the certificate renewal.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"status": {
				Description: "The status of ca certificate renewal.",
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

func dataSourceApplianceDeviceCertificateRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewApplianceDeviceCertificateWithDefaults()
	if v, ok := d.GetOk("ca_certificate"); ok {
		x := (v.(string))
		o.SetCaCertificate(x)
	}
	if v, ok := d.GetOk("ca_certificate_expiry_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCaCertificateExpiryTime(x)
	}
	if v, ok := d.GetOk("certificate_renewal_expiry_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCertificateRenewalExpiryTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("configuration_mo_id"); ok {
		x := (v.(string))
		o.SetConfigurationMoId(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("last_success_poll_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastSuccessPollTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.ApplianceApi.GetApplianceDeviceCertificateList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.ApplianceDeviceCertificateList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to ApplianceDeviceCertificate: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewApplianceDeviceCertificateWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("ca_certificate", (s.GetCaCertificate())); err != nil {
				return fmt.Errorf("error occurred while setting property CaCertificate: %+v", err)
			}

			if err := d.Set("ca_certificate_expiry_time", (s.GetCaCertificateExpiryTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property CaCertificateExpiryTime: %+v", err)
			}

			if err := d.Set("certificate_renewal_expiry_time", (s.GetCertificateRenewalExpiryTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property CertificateRenewalExpiryTime: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("completed_phases", flattenListApplianceCertRenewalPhase(s.GetCompletedPhases(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CompletedPhases: %+v", err)
			}
			if err := d.Set("configuration_mo_id", (s.GetConfigurationMoId())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigurationMoId: %+v", err)
			}

			if err := d.Set("current_phase", flattenMapApplianceCertRenewalPhase(s.GetCurrentPhase(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property CurrentPhase: %+v", err)
			}

			if err := d.Set("end_time", (s.GetEndTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property EndTime: %+v", err)
			}

			if err := d.Set("last_success_poll_time", (s.GetLastSuccessPollTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastSuccessPollTime: %+v", err)
			}
			if err := d.Set("messages", (s.GetMessages())); err != nil {
				return fmt.Errorf("error occurred while setting property Messages: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("start_time", (s.GetStartTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property StartTime: %+v", err)
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
