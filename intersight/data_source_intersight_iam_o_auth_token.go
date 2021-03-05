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

func dataSourceIamOAuthToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamOAuthTokenRead,
		Schema: map[string]*schema.Schema{
			"access_expiration_time": {
				Description: "Expiration time for the JWT token to which it can be used for api calls.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"token_id": {
				Description: "Token identifier. Not the Access Token itself.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"access_expiration_time": {
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceIamOAuthTokenRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
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
		return diag.Errorf("json marshal of IamOAuthToken object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamOAuthTokenList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamOAuthToken: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamOAuthTokenList.GetCount()
	var i int32
	var iamOAuthTokenResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamOAuthTokenList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamOAuthToken: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamOAuthTokenList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamOAuthToken data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["access_expiration_time"] = (s.GetAccessExpirationTime()).String()
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["app_registration"] = flattenMapIamAppRegistrationRelationship(s.GetAppRegistration(), d)
				temp["class_id"] = (s.GetClassId())
				temp["client_id"] = (s.GetClientId())
				temp["client_ip_address"] = (s.GetClientIpAddress())
				temp["client_name"] = (s.GetClientName())

				temp["expiration_time"] = (s.GetExpirationTime()).String()
				temp["last_login_client"] = (s.GetLastLoginClient())

				temp["last_login_time"] = (s.GetLastLoginTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["permission"] = flattenMapIamPermissionRelationship(s.GetPermission(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["token_id"] = (s.GetTokenId())

				temp["user"] = flattenMapIamUserRelationship(s.GetUser(), d)

				temp["user_meta"] = flattenMapIamClientMeta(s.GetUserMeta(), d)
				iamOAuthTokenResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamOAuthTokenResults))
	if err := d.Set("results", iamOAuthTokenResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamOAuthTokenResults[0]["moid"].(string))
	return de
}
