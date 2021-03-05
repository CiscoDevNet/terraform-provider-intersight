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

func dataSourceApplianceRestore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceRestoreRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"elapsed_time": {
				Description: "Elapsed time in seconds since the restore process has started.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"end_time": {
				Description: "End date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
				Description: "Password for authenticating with the file server.",
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
			"start_time": {
				Description: "Start date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the restore managed object.\n* `Started` - Backup or restore process has started.\n* `Created` - Backup or restore is in created state.\n* `Failed` - Backup or restore process has failed.\n* `Completed` - Backup or restore process has completed.\n* `Copied` - Backup file has been copied.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Description: "Username to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceApplianceRestore().Schema},
				Computed: true,
			}},
	}
}

func dataSourceApplianceRestoreRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceRestore{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("elapsed_time"); ok {
		x := int64(v.(int))
		o.SetElapsedTime(x)
	}
	if v, ok := d.GetOk("end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetEndTime(x)
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
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
	if v, ok := d.GetOk("start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetStartTime(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceRestore object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceRestoreList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ApplianceRestore: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ApplianceRestoreList.GetCount()
	var i int32
	var applianceRestoreResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceRestoreList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ApplianceRestore: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ApplianceRestoreList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ApplianceRestore data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["elapsed_time"] = (s.GetElapsedTime())

				temp["end_time"] = (s.GetEndTime()).String()
				temp["filename"] = (s.GetFilename())
				temp["is_password_set"] = (s.GetIsPasswordSet())
				temp["messages"] = (s.GetMessages())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["protocol"] = (s.GetProtocol())
				temp["remote_host"] = (s.GetRemoteHost())
				temp["remote_path"] = (s.GetRemotePath())
				temp["remote_port"] = (s.GetRemotePort())

				temp["start_time"] = (s.GetStartTime()).String()
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["username"] = (s.GetUsername())
				applianceRestoreResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceRestoreResults))
	if err := d.Set("results", applianceRestoreResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceRestoreResults[0]["moid"].(string))
	return de
}
