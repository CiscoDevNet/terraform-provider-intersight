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

func dataSourceSnmpPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSnmpPolicyRead,
		Schema: map[string]*schema.Schema{
			"access_community_string": {
				Description: "The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"community_access": {
				Description: "Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.\n* `Disabled` - Blocks access to the information in the inventory tables.\n* `Limited` - Partial access to read the information in the inventory tables.\n* `Full` - Full access to read the information in the inventory tables.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"engine_id": {
				Description: "User-defined unique identification of the static engine.",
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
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
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
			"snmp_port": {
				Description: "Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"snmp_traps": {
				Type:     schema.TypeList,
				Optional: true,
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
						"destination": {
							Description: "Address to which the SNMP trap information is sent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enabled": {
							Description: "Enables/disables the trap on the server If enabled, trap is active on the server.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Description: "Port used by the server to communicate with the trap destination. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269).",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"type": {
							Description: "Type of trap which decides whether to receive a notification when a trap is received at the destination.\n* `Trap` - Do not receive notifications when trap is sent to the destination.\n* `Inform` - Receive notifications when trap is sent to the destination. This option is valid only for V2 users.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"user": {
							Description: "SNMP user for the trap. Applicable only to SNMPv3.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_version": {
							Description: "SNMP version used for the trap.\n* `V3` - SNMP v3 trap version notifications.\n* `V2` - SNMP v2 trap version notifications.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"snmp_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"auth_password": {
							Description: "Authorization password for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"auth_type": {
							Description: "Authorization protocol for authenticating the user.\n* `NA` - Authentication protocol is not applicable.\n* `MD5` - MD5 protocol is used to authenticate SNMP user.\n* `SHA` - SHA protocol is used to authenticate SNMP user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_auth_password_set": {
							Description: "Indicates whether the value of the 'authPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_privacy_password_set": {
							Description: "Indicates whether the value of the 'privacyPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "SNMP username. Must have a minimum of 1 and and a maximum of 31 characters.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"privacy_password": {
							Description: "Privacy password for the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"privacy_type": {
							Description: "Privacy protocol for the user.\n* `NA` - Privacy protocol is not applicable.\n* `DES` - DES privacy protocol is used for SNMP user.\n* `AES` - AES privacy protocol is used for SNMP user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"security_level": {
							Description: "Security mechanism used for communication between agent and manager.\n* `AuthPriv` - The user requires both an authorization password and a privacy password.\n* `NoAuthNoPriv` - The user does not require an authorization or privacy password.\n* `AuthNoPriv` - The user requires an authorization password but not a privacy password.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"sys_contact": {
				Description: "Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sys_location": {
				Description: "Location of host on which the SNMP agent (server) runs.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"trap_community": {
				Description: "SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSnmpPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SnmpPolicy{}
	if v, ok := d.GetOk("access_community_string"); ok {
		x := (v.(string))
		o.SetAccessCommunityString(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("community_access"); ok {
		x := (v.(string))
		o.SetCommunityAccess(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(bool))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("engine_id"); ok {
		x := (v.(string))
		o.SetEngineId(x)
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
	if v, ok := d.GetOk("snmp_port"); ok {
		x := int64(v.(int))
		o.SetSnmpPort(x)
	}
	if v, ok := d.GetOk("sys_contact"); ok {
		x := (v.(string))
		o.SetSysContact(x)
	}
	if v, ok := d.GetOk("sys_location"); ok {
		x := (v.(string))
		o.SetSysLocation(x)
	}
	if v, ok := d.GetOk("trap_community"); ok {
		x := (v.(string))
		o.SetTrapCommunity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SnmpPolicy object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.SnmpApi.GetSnmpPolicyList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching SnmpPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for SnmpPolicy list: %s", err.Error())
	}
	var s = &models.SnmpPolicyList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to SnmpPolicy list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for SnmpPolicy data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for SnmpPolicy data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.SnmpPolicy{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("access_community_string", (s.GetAccessCommunityString())); err != nil {
				return diag.Errorf("error occurred while setting property AccessCommunityString: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("community_access", (s.GetCommunityAccess())); err != nil {
				return diag.Errorf("error occurred while setting property CommunityAccess: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("enabled", (s.GetEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property Enabled: %s", err.Error())
			}
			if err := d.Set("engine_id", (s.GetEngineId())); err != nil {
				return diag.Errorf("error occurred while setting property EngineId: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}

			if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Profiles: %s", err.Error())
			}
			if err := d.Set("snmp_port", (s.GetSnmpPort())); err != nil {
				return diag.Errorf("error occurred while setting property SnmpPort: %s", err.Error())
			}

			if err := d.Set("snmp_traps", flattenListSnmpTrap(s.GetSnmpTraps(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SnmpTraps: %s", err.Error())
			}

			if err := d.Set("snmp_users", flattenListSnmpUser(s.GetSnmpUsers(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SnmpUsers: %s", err.Error())
			}
			if err := d.Set("sys_contact", (s.GetSysContact())); err != nil {
				return diag.Errorf("error occurred while setting property SysContact: %s", err.Error())
			}
			if err := d.Set("sys_location", (s.GetSysLocation())); err != nil {
				return diag.Errorf("error occurred while setting property SysLocation: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("trap_community", (s.GetTrapCommunity())); err != nil {
				return diag.Errorf("error occurred while setting property TrapCommunity: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
