package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRecoveryBackupConfigPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRecoveryBackupConfigPolicyRead,
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
			"file_name_prefix": {
				Description: "The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"location_type": {
				Description: "Specifies whether the backup will be stored locally or remotely.\n* `Network Share` - The backup is stored remotely on a separate server.\n* `Local Storage` - The backup is stored locally on the endpoint.",
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
			"password": {
				Description: "Password of Backup server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"path": {
				Description: "The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"protocol": {
				Description: "Protocol for transferring the backup image to the network share location.\n* `SCP` - Secure Copy Protocol (SCP) to access the file server.\n* `SFTP` - SSH File Transfer Protocol (SFTP) to access file server.\n* `FTP` - File Transfer Protocol (FTP) to access file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"retention_count": {
				Description: "Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"user_name": {
				Description: "Username for the backup server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceRecoveryBackupConfigPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceRecoveryBackupConfigPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.RecoveryBackupConfigPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("file_name_prefix"); ok {
		x := (v.(string))
		o.SetFileNamePrefix(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("location_type"); ok {
		x := (v.(string))
		o.SetLocationType(x)
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
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
	}
	if v, ok := d.GetOk("path"); ok {
		x := (v.(string))
		o.SetPath(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("retention_count"); ok {
		x := int64(v.(int))
		o.SetRetentionCount(x)
	}
	if v, ok := d.GetOk("user_name"); ok {
		x := (v.(string))
		o.SetUserName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of RecoveryBackupConfigPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.RecoveryApi.GetRecoveryBackupConfigPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of RecoveryBackupConfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.RecoveryBackupConfigPolicyList.GetCount()
	var i int32
	var recoveryBackupConfigPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.RecoveryApi.GetRecoveryBackupConfigPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching RecoveryBackupConfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.RecoveryBackupConfigPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for RecoveryBackupConfigPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["backup_profiles"] = flattenListRecoveryBackupProfileRelationship(s.GetBackupProfiles(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["file_name_prefix"] = (s.GetFileNamePrefix())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["location_type"] = (s.GetLocationType())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["path"] = (s.GetPath())
				temp["protocol"] = (s.GetProtocol())
				temp["retention_count"] = (s.GetRetentionCount())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["user_name"] = (s.GetUserName())
				recoveryBackupConfigPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(recoveryBackupConfigPolicyResults))
	if err := d.Set("results", recoveryBackupConfigPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(recoveryBackupConfigPolicyResults[0]["moid"].(string))
	return de
}
