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

func dataSourceAssetSubscriptionDeviceContractInformation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAssetSubscriptionDeviceContractInformationRead,
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
			"device_contract_information": {
				Description: "A reference to a assetDeviceContractInformation resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"device_id": {
				Description: "Unique identifier of the Cisco device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_information": {
				Description: "Detailed device information including installation status and customer address.",
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
						"application_name": {
							Description: "Application name reported by Cisco Install Base.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"device_transactions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Description: "The action taken by Cisco Install Base on the device.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
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
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"status_description": {
										Description: "Description of status of Cisco device reported by Cisco Install Base.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"status_id": {
										Description: "Status ID of device reported by Cisco Install Base.\n* `0` - A default value to catch cases where device status is not correctly detected.\n* `10000` - Device is installed successfully.\n* `1010041` - Device is currently in Return Material Authorization process.\n* `10005` - Device is replaced successfully with another device.\n* `10007` - Device is returned succcessfuly.\n* `10009` - Device is terminated at customer's end.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"timestamp": {
										Description: "Timestamp field reported by Cisco Install Base.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"transaction_batch_id": {
										Description: "Transaction batch ID reported by Cisco Install Base.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"transaction_date": {
										Description: "Transaction Date reported by Cisco Install Base.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"transaction_sequence": {
										Description: "Transaction sequence reported by Cisco Install Base.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
								},
							},
							Computed: true,
						},
						"end_customer": {
							Description: "End customer information for the Cisco support contract purchased for the Cisco device.",
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
									"address": {
										Description: "Address as per the information provided by the user.",
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
												"address1": {
													Description: "Address Line one of the address information. example \"PO BOX 641570\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"address2": {
													Description: "Address Line two of the address information. example \"Cisco Systems\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"address3": {
													Description: "Address Line three of the address information. example \"Cisco Systems\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"city": {
													Description: "City in which the address resides. example \"San Jose\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"class_id": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"country": {
													Description: "Country in which the address resides. example \"US\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"county": {
													Description: "County in which the address resides. example \"Washington County\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"location": {
													Description: "Location in which the address resides. example \"14852\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"name": {
													Description: "Name of the user whose address is being populated.",
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
												"postal_code": {
													Description: "Postal Code in which the address resides. example \"95164-1570\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"province": {
													Description: "Province in which the address resides. example \"AB\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"state": {
													Description: "State in which the address resides. example \"CA\".",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"id": {
										Description: "Unique identifier for an end customer. This identifier is allocated by Cisco.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name as per the information provided by the user.",
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
								},
							},
						},
						"instance_id": {
							Description: "Instance number of the device. example \"917280220\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"item_type": {
							Description: "Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"mlb_offer_type": {
							Description: "Identifier for consumption based pricing. MLB refers to MultiLine Bundle.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"mlb_product_id": {
							Description: "Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"mlb_product_name": {
							Description: "Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle.",
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
						"old_device_id": {
							Description: "Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"old_device_status_description": {
							Description: "Description of status of old Cisco device, which got replaced by the new device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"old_device_status_id": {
							Description: "Status ID of old Cisco device, which got replaced by the new device.\n* `0` - A default value to catch cases where device status is not correctly detected.\n* `10000` - Device is installed successfully.\n* `1010041` - Device is currently in Return Material Authorization process.\n* `10005` - Device is replaced successfully with another device.\n* `10007` - Device is returned succcessfuly.\n* `10009` - Device is terminated at customer's end.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"old_instance_id": {
							Description: "Instance number of the old device, which got replaced by the new device.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"product_type": {
							Description: "Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"device_pid": {
				Description: "Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
			"subscription_ref_id": {
				Description: "Identifies the consumption-based subscription.",
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

func dataSourceAssetSubscriptionDeviceContractInformationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.AssetSubscriptionDeviceContractInformation{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("device_pid"); ok {
		x := (v.(string))
		o.SetDevicePid(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("subscription_ref_id"); ok {
		x := (v.(string))
		o.SetSubscriptionRefId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of AssetSubscriptionDeviceContractInformation object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.AssetApi.GetAssetSubscriptionDeviceContractInformationList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching AssetSubscriptionDeviceContractInformation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for AssetSubscriptionDeviceContractInformation list: %s", err.Error())
	}
	var s = &models.AssetSubscriptionDeviceContractInformationList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to AssetSubscriptionDeviceContractInformation list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for AssetSubscriptionDeviceContractInformation did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.AssetSubscriptionDeviceContractInformation{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("device_contract_information", flattenMapAssetDeviceContractInformationRelationship(s.GetDeviceContractInformation(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DeviceContractInformation: %s", err.Error())
			}
			if err := d.Set("device_id", (s.GetDeviceId())); err != nil {
				return diag.Errorf("error occurred while setting property DeviceId: %s", err.Error())
			}

			if err := d.Set("device_information", flattenMapAssetDeviceInformation(s.GetDeviceInformation(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DeviceInformation: %s", err.Error())
			}
			if err := d.Set("device_pid", (s.GetDevicePid())); err != nil {
				return diag.Errorf("error occurred while setting property DevicePid: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("subscription_ref_id", (s.GetSubscriptionRefId())); err != nil {
				return diag.Errorf("error occurred while setting property SubscriptionRefId: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
