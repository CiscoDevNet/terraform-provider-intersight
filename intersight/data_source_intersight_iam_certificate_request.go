package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIamCertificateRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamCertificateRequestRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"email_address": {
				Description: "User input email address, an optional part of the subject of the certificate request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the certificate request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request": {
				Description: "Generated certificate signing request (CSR) in PEM format.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"self_signed": {
				Description: "Whether the user wants the generated CSR to be self-signed by the appliance.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamCertificateRequest().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamCertificateRequestRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamCertificateRequest{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("email_address"); ok {
		x := (v.(string))
		o.SetEmailAddress(x)
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
	if v, ok := d.GetOk("request"); ok {
		x := (v.(string))
		o.SetRequest(x)
	}
	if v, ok := d.GetOk("self_signed"); ok {
		x := (v.(bool))
		o.SetSelfSigned(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamCertificateRequest object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamCertificateRequestList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamCertificateRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamCertificateRequestList.GetCount()
	var i int32
	var iamCertificateRequestResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamCertificateRequestList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamCertificateRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamCertificateRequestList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamCertificateRequest data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["certificate"] = flattenMapIamCertificateRelationship(s.GetCertificate(), d)
				temp["class_id"] = (s.GetClassId())
				temp["email_address"] = (s.GetEmailAddress())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["private_key_spec"] = flattenMapIamPrivateKeySpecRelationship(s.GetPrivateKeySpec(), d)
				temp["request"] = (s.GetRequest())
				temp["self_signed"] = (s.GetSelfSigned())

				temp["subject"] = flattenMapPkixDistinguishedName(s.GetSubject(), d)

				temp["subject_alternate_name"] = flattenMapPkixSubjectAlternateName(s.GetSubjectAlternateName(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				iamCertificateRequestResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamCertificateRequestResults))
	if err := d.Set("results", iamCertificateRequestResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamCertificateRequestResults[0]["moid"].(string))
	return de
}
