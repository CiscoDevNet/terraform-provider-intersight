package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamLocalUserPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamLocalUserPasswordPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"min_char_difference": {
				Description: "Minimum number of characters different from previous password.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_days_between_password_change": {
				Description: "Minimum Days allowed between password change.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_length_password": {
				Description: "Minimum length of password.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_lower_case": {
				Description: "Minimum number of required lower case characters.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_numeric": {
				Description: "Minimum number of required numeric characters.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_special_char": {
				Description: "Minimum number of required special characters.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_upper_case": {
				Description: "Minimum number of required upper case characters.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_previous_passwords_disallowed": {
				Description: "Number of previous passwords disallowed.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account": {
					Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
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
					Computed: true,
				},
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
					"min_char_difference": {
						Description: "Minimum number of characters different from previous password.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_days_between_password_change": {
						Description: "Minimum Days allowed between password change.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_length_password": {
						Description: "Minimum length of password.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_lower_case": {
						Description: "Minimum number of required lower case characters.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_numeric": {
						Description: "Minimum number of required numeric characters.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_special_char": {
						Description: "Minimum number of required special characters.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"min_upper_case": {
						Description: "Minimum number of required upper case characters.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"num_previous_passwords_disallowed": {
						Description: "Number of previous passwords disallowed.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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

func dataSourceIamLocalUserPasswordPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamLocalUserPasswordPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("min_char_difference"); ok {
		x := int64(v.(int))
		o.SetMinCharDifference(x)
	}
	if v, ok := d.GetOk("min_days_between_password_change"); ok {
		x := int64(v.(int))
		o.SetMinDaysBetweenPasswordChange(x)
	}
	if v, ok := d.GetOk("min_length_password"); ok {
		x := int64(v.(int))
		o.SetMinLengthPassword(x)
	}
	if v, ok := d.GetOk("min_lower_case"); ok {
		x := int64(v.(int))
		o.SetMinLowerCase(x)
	}
	if v, ok := d.GetOk("min_numeric"); ok {
		x := int64(v.(int))
		o.SetMinNumeric(x)
	}
	if v, ok := d.GetOk("min_special_char"); ok {
		x := int64(v.(int))
		o.SetMinSpecialChar(x)
	}
	if v, ok := d.GetOk("min_upper_case"); ok {
		x := int64(v.(int))
		o.SetMinUpperCase(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_previous_passwords_disallowed"); ok {
		x := int64(v.(int))
		o.SetNumPreviousPasswordsDisallowed(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamLocalUserPasswordPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamLocalUserPasswordPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamLocalUserPasswordPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamLocalUserPasswordPolicyList.GetCount()
	var i int32
	var iamLocalUserPasswordPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamLocalUserPasswordPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamLocalUserPasswordPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamLocalUserPasswordPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamLocalUserPasswordPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["min_char_difference"] = (s.GetMinCharDifference())
				temp["min_days_between_password_change"] = (s.GetMinDaysBetweenPasswordChange())
				temp["min_length_password"] = (s.GetMinLengthPassword())
				temp["min_lower_case"] = (s.GetMinLowerCase())
				temp["min_numeric"] = (s.GetMinNumeric())
				temp["min_special_char"] = (s.GetMinSpecialChar())
				temp["min_upper_case"] = (s.GetMinUpperCase())
				temp["moid"] = (s.GetMoid())
				temp["num_previous_passwords_disallowed"] = (s.GetNumPreviousPasswordsDisallowed())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iamLocalUserPasswordPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamLocalUserPasswordPolicyResults))
	if err := d.Set("results", iamLocalUserPasswordPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamLocalUserPasswordPolicyResults[0]["moid"].(string))
	return de
}
