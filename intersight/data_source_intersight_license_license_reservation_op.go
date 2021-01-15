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

func dataSourceLicenseLicenseReservationOp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLicenseLicenseReservationOpRead,
		Schema: map[string]*schema.Schema{
			"account": {
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
			"auth_code": {
				Description: "Revervation code used to install the license.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auth_code_installed": {
				Description: "Flag to indicate whether authorization code is installed.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"confirm_code": {
				Description: "Confirm code used to complete the license update on smart license account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"generate_request_code": {
				Description: "Trigger the generation of request code for specific license reservation.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"generate_return_code": {
				Description: "Trigger the generation of return code for specific license reservation.",
				Type:        schema.TypeBool,
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
			"request_code": {
				Description: "Revervation code used to generate authorization code from CSSM.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"return_code": {
				Description: "Return code used to return the reserved license to smart license account.",
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

func dataSourceLicenseLicenseReservationOpRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.LicenseLicenseReservationOp{}
	if v, ok := d.GetOk("auth_code"); ok {
		x := (v.(string))
		o.SetAuthCode(x)
	}
	if v, ok := d.GetOk("auth_code_installed"); ok {
		x := (v.(bool))
		o.SetAuthCodeInstalled(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("confirm_code"); ok {
		x := (v.(string))
		o.SetConfirmCode(x)
	}
	if v, ok := d.GetOk("generate_request_code"); ok {
		x := (v.(bool))
		o.SetGenerateRequestCode(x)
	}
	if v, ok := d.GetOk("generate_return_code"); ok {
		x := (v.(bool))
		o.SetGenerateReturnCode(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("request_code"); ok {
		x := (v.(string))
		o.SetRequestCode(x)
	}
	if v, ok := d.GetOk("return_code"); ok {
		x := (v.(string))
		o.SetReturnCode(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of LicenseLicenseReservationOp object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseLicenseReservationOpList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for LicenseLicenseReservationOp list: %s", err.Error())
	}
	var s = &models.LicenseLicenseReservationOpList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to LicenseLicenseReservationOp list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for LicenseLicenseReservationOp data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for LicenseLicenseReservationOp data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.LicenseLicenseReservationOp{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Account: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("auth_code", (s.GetAuthCode())); err != nil {
				return diag.Errorf("error occurred while setting property AuthCode: %s", err.Error())
			}
			if err := d.Set("auth_code_installed", (s.GetAuthCodeInstalled())); err != nil {
				return diag.Errorf("error occurred while setting property AuthCodeInstalled: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("confirm_code", (s.GetConfirmCode())); err != nil {
				return diag.Errorf("error occurred while setting property ConfirmCode: %s", err.Error())
			}
			if err := d.Set("generate_request_code", (s.GetGenerateRequestCode())); err != nil {
				return diag.Errorf("error occurred while setting property GenerateRequestCode: %s", err.Error())
			}
			if err := d.Set("generate_return_code", (s.GetGenerateReturnCode())); err != nil {
				return diag.Errorf("error occurred while setting property GenerateReturnCode: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("request_code", (s.GetRequestCode())); err != nil {
				return diag.Errorf("error occurred while setting property RequestCode: %s", err.Error())
			}
			if err := d.Set("return_code", (s.GetReturnCode())); err != nil {
				return diag.Errorf("error occurred while setting property ReturnCode: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
