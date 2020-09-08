package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIaasDeviceStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIaasDeviceStatusRead,
		Schema: map[string]*schema.Schema{
			"account_name": {
				Description: "The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"account_type": {
				Description: "The UCSD Infra Account type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"category": {
				Description: "The UCSD Infra Account category.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"claim_status": {
				Description: "Describes if the device is claimed in Intersight or not.\n* `Unknown` - If UCS Director managed account claim status information is unknown.\n* `Yes` - If UCS Director managed account is claimed in Intersight.\n* `No` - If UCS Director managed account is not claimed in Intersight.\n* `Not Applicable` - If UCS Director managed account is not capable of providing claim status information.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status": {
				Description: "Describes about the connection status between the UCSD and the actual end device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_model": {
				Description: "Describes about the device model.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_vendor": {
				Description: "Describes about the device vendor/manufacturer of the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_version": {
				Description: "Describes about the current firmware version running on the device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "A reference to a iaasUcsdInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"ip_address": {
				Description: "The IPAddress of the device.",
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
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod": {
				Description: "Describes about the pod to which this device belongs to in UCSD.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pod_type": {
				Description: "Describes about the podType of Pod to which this device belongs to in UCSD.",
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
		},
	}
}

func dataSourceIaasDeviceStatusRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewIaasDeviceStatusWithDefaults()
	if v, ok := d.GetOk("account_name"); ok {
		x := (v.(string))
		o.SetAccountName(x)
	}
	if v, ok := d.GetOk("account_type"); ok {
		x := (v.(string))
		o.SetAccountType(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("claim_status"); ok {
		x := (v.(string))
		o.SetClaimStatus(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}
	if v, ok := d.GetOk("device_model"); ok {
		x := (v.(string))
		o.SetDeviceModel(x)
	}
	if v, ok := d.GetOk("device_vendor"); ok {
		x := (v.(string))
		o.SetDeviceVendor(x)
	}
	if v, ok := d.GetOk("device_version"); ok {
		x := (v.(string))
		o.SetDeviceVersion(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pod"); ok {
		x := (v.(string))
		o.SetPod(x)
	}
	if v, ok := d.GetOk("pod_type"); ok {
		x := (v.(string))
		o.SetPodType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.IaasApi.GetIaasDeviceStatusList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.IaasDeviceStatusList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to IaasDeviceStatus: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewIaasDeviceStatusWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("account_name", (s.GetAccountName())); err != nil {
				return fmt.Errorf("error occurred while setting property AccountName: %+v", err)
			}
			if err := d.Set("account_type", (s.GetAccountType())); err != nil {
				return fmt.Errorf("error occurred while setting property AccountType: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("category", (s.GetCategory())); err != nil {
				return fmt.Errorf("error occurred while setting property Category: %+v", err)
			}
			if err := d.Set("claim_status", (s.GetClaimStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property ClaimStatus: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("connection_status", (s.GetConnectionStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property ConnectionStatus: %+v", err)
			}
			if err := d.Set("device_model", (s.GetDeviceModel())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceModel: %+v", err)
			}
			if err := d.Set("device_vendor", (s.GetDeviceVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceVendor: %+v", err)
			}
			if err := d.Set("device_version", (s.GetDeviceVersion())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceVersion: %+v", err)
			}

			if err := d.Set("guid", flattenMapIaasUcsdInfoRelationship(s.GetGuid(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Guid: %+v", err)
			}
			if err := d.Set("ip_address", (s.GetIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property IpAddress: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("pod", (s.GetPod())); err != nil {
				return fmt.Errorf("error occurred while setting property Pod: %+v", err)
			}
			if err := d.Set("pod_type", (s.GetPodType())); err != nil {
				return fmt.Errorf("error occurred while setting property PodType: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
