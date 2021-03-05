package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceBackupPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceBackupPolicyRead,
		Schema: map[string]*schema.Schema{
			"backup_time": {
				Description: "The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filename": {
				Description: "Backup filename to backup or restore.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"manual_backup": {
				Description: "Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored.",
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
			"password": {
				Description: "Password to authenticate the file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"protocol": {
				Description: "Communication protocol used by the file server (e.g. scp or sftp).\n* `scp` - Secure Copy Protocol (SCP) to access the file server.\n* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_host": {
				Description: "Hostname of the remote file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_path": {
				Description: "File server directory to copy the file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"remote_port": {
				Description: "Remote TCP port on the file server (e.g. 22 for scp).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"username": {
				Description: "Username to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceApplianceBackupPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceApplianceBackupPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceBackupPolicy{}
	if v, ok := d.GetOk("backup_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetBackupTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("manual_backup"); ok {
		x := (v.(bool))
		o.SetManualBackup(x)
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
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("remote_host"); ok {
		x := (v.(string))
		o.SetRemoteHost(x)
	}
	if v, ok := d.GetOk("remote_path"); ok {
		x := (v.(string))
		o.SetRemotePath(x)
	}
	if v, ok := d.GetOk("remote_port"); ok {
		x := int64(v.(int))
		o.SetRemotePort(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceBackupPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceBackupPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceBackupPolicyList.GetCount()
	var i int32
	var applianceBackupPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceBackupPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceBackupPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceBackupPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["backup_time"] = (s.GetBackupTime()).String()
				temp["class_id"] = (s.GetClassId())
				temp["filename"] = (s.GetFilename())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["manual_backup"] = (s.GetManualBackup())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["protocol"] = (s.GetProtocol())
				temp["remote_host"] = (s.GetRemoteHost())
				temp["remote_path"] = (s.GetRemotePath())
				temp["remote_port"] = (s.GetRemotePort())

				temp["schedule"] = flattenMapOnpremSchedule(s.GetSchedule(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["username"] = (s.GetUsername())
				applianceBackupPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceBackupPolicyResults))
	if err := d.Set("results", applianceBackupPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceBackupPolicyResults[0]["moid"].(string))
	return de
}
