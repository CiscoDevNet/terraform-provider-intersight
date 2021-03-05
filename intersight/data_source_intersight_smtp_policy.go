package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSmtpPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSmtpPolicyRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "If enabled, controls the state of the SMTP client service on the managed device.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"min_severity": {
				Description: "Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level.\n* `critical` - Minimum severity to report is critical.\n* `condition` - Minimum severity to report is informational.\n* `warning` - Minimum severity to report is warning.\n* `minor` - Minimum severity to report is minor.\n* `major` - Minimum severity to report is major.",
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
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sender_email": {
				Description: "The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"smtp_port": {
				Description: "Port number used by the SMTP server for outgoing SMTP communication.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"smtp_server": {
				Description: "IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSmtpPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSmtpPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SmtpPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("min_severity"); ok {
		x := (v.(string))
		o.SetMinSeverity(x)
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
	if v, ok := d.GetOk("sender_email"); ok {
		x := (v.(string))
		o.SetSenderEmail(x)
	}
	if v, ok := d.GetOk("smtp_port"); ok {
		x := int64(v.(int))
		o.SetSmtpPort(x)
	}
	if v, ok := d.GetOk("smtp_server"); ok {
		x := (v.(string))
		o.SetSmtpServer(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SmtpPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SmtpApi.GetSmtpPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of SmtpPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.SmtpPolicyList.GetCount()
	var i int32
	var smtpPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SmtpApi.GetSmtpPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching SmtpPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.SmtpPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for SmtpPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["enabled"] = (s.GetEnabled())
				temp["min_severity"] = (s.GetMinSeverity())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["sender_email"] = (s.GetSenderEmail())
				temp["smtp_port"] = (s.GetSmtpPort())
				temp["smtp_recipients"] = (s.GetSmtpRecipients())
				temp["smtp_server"] = (s.GetSmtpServer())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				smtpPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(smtpPolicyResults))
	if err := d.Set("results", smtpPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(smtpPolicyResults[0]["moid"].(string))
	return de
}
