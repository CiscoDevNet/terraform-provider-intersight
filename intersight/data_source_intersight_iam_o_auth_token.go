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

func dataSourceIamOAuthToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIamOAuthTokenRead,
		Schema: map[string]*schema.Schema{
			"access_expiration_time": {
				Description: "Expiration time for the JWT token to which it can be used for api calls.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"app_registration": {
				Description: "A reference to a iamAppRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
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
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_id": {
				Description: "The identifier of the registered application to which the token belongs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_ip_address": {
				Description: "The user agent IP address from which the auth token is launched.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"client_name": {
				Description: "The name of the registered application to which the token belongs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"expiration_time": {
				Description: "Expiration time for the JWT token to which it can be refreshed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_login_client": {
				Description: "The client address from which last login is initiated.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_login_time": {
				Description: "The last login time for user.",
				Type:        schema.TypeString,
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
			"permission": {
				Description: "A reference to a iamPermission resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
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
			"token_id": {
				Description: "Token identifier. Not the Access Token itself.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Description: "A reference to a iamUser resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
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
			},
			"user_meta": {
				Description: "User Device meta information.",
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
						"device_model": {
							Description: "Parsed device model from raw User-Agent.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"user_agent": {
							Description: "The value of the \"User-Agent\" HTTP header, as sent by the HTTP client when it initiate a session to Intersight. This can be used to identify the client operating system, browser type and browser version.\nExample - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36\nIt is set when User successfully passed OAuth login flow and receives Access Token.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func dataSourceIamOAuthTokenRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.IamOAuthToken{}
	if v, ok := d.GetOk("access_expiration_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetAccessExpirationTime(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("client_id"); ok {
		x := (v.(string))
		o.SetClientId(x)
	}
	if v, ok := d.GetOk("client_ip_address"); ok {
		x := (v.(string))
		o.SetClientIpAddress(x)
	}
	if v, ok := d.GetOk("client_name"); ok {
		x := (v.(string))
		o.SetClientName(x)
	}
	if v, ok := d.GetOk("expiration_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetExpirationTime(x)
	}
	if v, ok := d.GetOk("last_login_client"); ok {
		x := (v.(string))
		o.SetLastLoginClient(x)
	}
	if v, ok := d.GetOk("last_login_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetLastLoginTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("token_id"); ok {
		x := (v.(string))
		o.SetTokenId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.IamApi.GetIamOAuthTokenList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.IamOAuthTokenList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to IamOAuthToken: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.IamOAuthToken{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("access_expiration_time", (s.GetAccessExpirationTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property AccessExpirationTime: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}

			if err := d.Set("app_registration", flattenMapIamAppRegistrationRelationship(s.GetAppRegistration(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property AppRegistration: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("client_id", (s.GetClientId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClientId: %+v", err)
			}
			if err := d.Set("client_ip_address", (s.GetClientIpAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property ClientIpAddress: %+v", err)
			}
			if err := d.Set("client_name", (s.GetClientName())); err != nil {
				return fmt.Errorf("error occurred while setting property ClientName: %+v", err)
			}

			if err := d.Set("expiration_time", (s.GetExpirationTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property ExpirationTime: %+v", err)
			}
			if err := d.Set("last_login_client", (s.GetLastLoginClient())); err != nil {
				return fmt.Errorf("error occurred while setting property LastLoginClient: %+v", err)
			}

			if err := d.Set("last_login_time", (s.GetLastLoginTime()).String()); err != nil {
				return fmt.Errorf("error occurred while setting property LastLoginTime: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("permission", flattenMapIamPermissionRelationship(s.GetPermission(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Permission: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("token_id", (s.GetTokenId())); err != nil {
				return fmt.Errorf("error occurred while setting property TokenId: %+v", err)
			}

			if err := d.Set("user", flattenMapIamUserRelationship(s.GetUser(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property User: %+v", err)
			}

			if err := d.Set("user_meta", flattenMapIamClientMeta(s.GetUserMeta(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property UserMeta: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
