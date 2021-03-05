package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIpmioverlanPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpmioverlanPolicyRead,
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
				Description: "State of the IPMI Over LAN service on the endpoint.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"encryption_key": {
				Description: "The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_encryption_key_set": {
				Description: "Indicates whether the value of the 'encryptionKey' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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
			"privilege": {
				Description: "The highest privilege level that can be assigned to an IPMI session on a server.\n* `admin` - Privilege to perform all actions available through IPMI.\n* `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks.\n* `read-only` - Privilege to view information throught IPMI but restriction on making any changes.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIpmioverlanPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIpmioverlanPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IpmioverlanPolicy{}
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
	if v, ok := d.GetOk("encryption_key"); ok {
		x := (v.(string))
		o.SetEncryptionKey(x)
	}
	if v, ok := d.GetOk("is_encryption_key_set"); ok {
		x := (v.(bool))
		o.SetIsEncryptionKeySet(x)
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
	if v, ok := d.GetOk("privilege"); ok {
		x := (v.(string))
		o.SetPrivilege(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IpmioverlanPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IpmioverlanApi.GetIpmioverlanPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IpmioverlanPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IpmioverlanPolicyList.GetCount()
	var i int32
	var ipmioverlanPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IpmioverlanApi.GetIpmioverlanPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IpmioverlanPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IpmioverlanPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IpmioverlanPolicy data source did not return results. Please change your search criteria and try again")
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
				temp["is_encryption_key_set"] = (s.GetIsEncryptionKeySet())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["privilege"] = (s.GetPrivilege())

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				ipmioverlanPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(ipmioverlanPolicyResults))
	if err := d.Set("results", ipmioverlanPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(ipmioverlanPolicyResults[0]["moid"].(string))
	return de
}
