package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceRemoteFileImport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceRemoteFileImportRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filename": {
				Description: "The name of the file to be imported.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": {
				Description: "The hostname of the machine where the file is located.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password for remote requiest.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"path": {
				Description: "The port that should be used for the remote request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"port": {
				Description: "The port that should be used for the remote request.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"protocol": {
				Description: "Protocol for the remote request.\n* `scp` - Secure Copy Protocol (SCP) to access the file server.\n* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"username": {
				Description: "The username for the remote request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceApplianceRemoteFileImport().Schema},
				Computed: true,
			}},
	}
}

func dataSourceApplianceRemoteFileImportRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceRemoteFileImport{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}
	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.SetHostname(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
	}
	if v, ok := d.GetOk("path"); ok {
		x := (v.(string))
		o.SetPath(x)
	}
	if v, ok := d.GetOk("port"); ok {
		x := int64(v.(int))
		o.SetPort(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceRemoteFileImport object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceRemoteFileImportList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceRemoteFileImport: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceRemoteFileImportList.GetCount()
	var i int32
	var applianceRemoteFileImportResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceRemoteFileImportList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceRemoteFileImport: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceRemoteFileImportList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceRemoteFileImport data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["filename"] = (s.GetFilename())
				temp["hostname"] = (s.GetHostname())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["path"] = (s.GetPath())
				temp["port"] = (s.GetPort())
				temp["protocol"] = (s.GetProtocol())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["username"] = (s.GetUsername())
				applianceRemoteFileImportResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceRemoteFileImportResults))
	if err := d.Set("results", applianceRemoteFileImportResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceRemoteFileImportResults[0]["moid"].(string))
	return de
}
