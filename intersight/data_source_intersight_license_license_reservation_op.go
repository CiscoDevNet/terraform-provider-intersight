package intersight

import (
	"context"
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceLicenseLicenseReservationOp().Schema},
				Computed: true,
			}},
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
	countResponse, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseLicenseReservationOpList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.LicenseLicenseReservationOpList.GetCount()
	var i int32
	var licenseLicenseReservationOpResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.LicenseApi.GetLicenseLicenseReservationOpList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.LicenseLicenseReservationOpList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for LicenseLicenseReservationOp data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["auth_code"] = (s.GetAuthCode())
				temp["auth_code_installed"] = (s.GetAuthCodeInstalled())
				temp["class_id"] = (s.GetClassId())
				temp["confirm_code"] = (s.GetConfirmCode())
				temp["generate_request_code"] = (s.GetGenerateRequestCode())
				temp["generate_return_code"] = (s.GetGenerateReturnCode())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["request_code"] = (s.GetRequestCode())
				temp["return_code"] = (s.GetReturnCode())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				licenseLicenseReservationOpResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(licenseLicenseReservationOpResults))
	if err := d.Set("results", licenseLicenseReservationOpResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(licenseLicenseReservationOpResults[0]["moid"].(string))
	return de
}
