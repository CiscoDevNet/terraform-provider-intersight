package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApplianceRestore() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplianceRestoreCreate,
		Read:   resourceApplianceRestoreRead,
		Delete: resourceApplianceRestoreDelete,
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
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
				ForceNew:         true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"elapsed_time": {
				Description: "Elapsed time in seconds since the restore process has started.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"end_time": {
				Description: "End date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"filename": {
				Description: "Backup filename to backup or restore.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"messages": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}, ForceNew: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"password": {
				Description: "Password for authenticating with the file server.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"protocol": {
				Description: "Communication protocol used by the file server (e.g. scp or sftp).\n* `scp` - Secure Copy Protocol (SCP) to access the file server.\n* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "scp",
				ForceNew:    true,
			},
			"remote_host": {
				Description: "Hostname of the remote file server.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"remote_path": {
				Description: "File server directory to copy the file.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"remote_port": {
				Description: "Remote TCP port on the file server (e.g. 22 for scp).",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"start_time": {
				Description: "Start date and time of the restore process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"status": {
				Description: "Status of the restore managed object.\n* `Started` - Backup or restore process has started.\n* `Created` - Backup or restore is in created state.\n* `Failed` - Backup or restore process has failed.\n* `Completed` - Backup or restore process has completed.\n* `Copied` - Backup file has been copied.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
							ForceNew:         true,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ForceNew: true,
			},
			"username": {
				Description: "Username to authenticate the fileserver.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceApplianceRestoreCreate(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewApplianceRestoreWithDefaults()
	if v, ok := d.GetOk("account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("appliance.Restore")

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

	if v, ok := d.GetOkExists("is_password_set"); ok {
		x := v.(bool)
		o.SetIsPasswordSet(x)
	}

	if v, ok := d.GetOk("messages"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetMessages(x)
		}
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("appliance.Restore")

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

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	r := conn.ApiClient.ApplianceApi.CreateApplianceRestore(conn.ctx).ApplianceRestore(*o)
	result, _, err := r.Execute()
	if err != nil {
		return fmt.Errorf("Failed to invoke operation: %v", err)
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceApplianceRestoreRead(d, meta)
}

func resourceApplianceRestoreRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	r := conn.ApiClient.ApplianceApi.GetApplianceRestoreByMoid(conn.ctx, d.Id())
	s, _, err := r.Execute()

	if err != nil {
		return fmt.Errorf("error in unmarshaling model for read Error: %s", err.Error())
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Account: %+v", err)
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
	}

	if err := d.Set("elapsed_time", (s.GetElapsedTime())); err != nil {
		return fmt.Errorf("error occurred while setting property ElapsedTime: %+v", err)
	}

	if err := d.Set("end_time", (s.GetEndTime()).String()); err != nil {
		return fmt.Errorf("error occurred while setting property EndTime: %+v", err)
	}

	if err := d.Set("filename", (s.GetFilename())); err != nil {
		return fmt.Errorf("error occurred while setting property Filename: %+v", err)
	}

	if err := d.Set("is_password_set", (s.GetIsPasswordSet())); err != nil {
		return fmt.Errorf("error occurred while setting property IsPasswordSet: %+v", err)
	}

	if err := d.Set("messages", (s.GetMessages())); err != nil {
		return fmt.Errorf("error occurred while setting property Messages: %+v", err)
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return fmt.Errorf("error occurred while setting property Moid: %+v", err)
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
	}

	if err := d.Set("protocol", (s.GetProtocol())); err != nil {
		return fmt.Errorf("error occurred while setting property Protocol: %+v", err)
	}

	if err := d.Set("remote_host", (s.GetRemoteHost())); err != nil {
		return fmt.Errorf("error occurred while setting property RemoteHost: %+v", err)
	}

	if err := d.Set("remote_path", (s.GetRemotePath())); err != nil {
		return fmt.Errorf("error occurred while setting property RemotePath: %+v", err)
	}

	if err := d.Set("remote_port", (s.GetRemotePort())); err != nil {
		return fmt.Errorf("error occurred while setting property RemotePort: %+v", err)
	}

	if err := d.Set("start_time", (s.GetStartTime()).String()); err != nil {
		return fmt.Errorf("error occurred while setting property StartTime: %+v", err)
	}

	if err := d.Set("status", (s.GetStatus())); err != nil {
		return fmt.Errorf("error occurred while setting property Status: %+v", err)
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return fmt.Errorf("error occurred while setting property Tags: %+v", err)
	}

	if err := d.Set("username", (s.GetUsername())); err != nil {
		return fmt.Errorf("error occurred while setting property Username: %+v", err)
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return nil
}

func resourceApplianceRestoreDelete(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	p := conn.ApiClient.ApplianceApi.DeleteApplianceRestore(conn.ctx, d.Id())
	_, err := p.Execute()
	if err != nil {
		return fmt.Errorf("error occurred while deleting: %s", err.Error())
	}
	return err
}
