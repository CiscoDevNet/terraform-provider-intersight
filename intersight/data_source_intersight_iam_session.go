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

func dataSourceIamSession() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamSessionRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_ip_address": {
				Description: "The user agent IP address from which the session is launched.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"expiration": {
				Description: "Expiration time for the session.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"idle_time_expiration": {
				Description: "Idle time expiration for the session.",
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
			"session_id": {
				Description: "Session token shared with the user agent which is used to identify the user session when API requests are received to perform authorization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_permissions": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"account_identifier": {
								Description: "MOID of the account which a user can select after authentication.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"account_name": {
								Description: "Name of the account which a user can select after authentication.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"account_status": {
								Description: "Status of the account. Account remains inactive until a device is claimed to the account.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
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
							"object_type": {
								Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
								Type:        schema.TypeString,
								Optional:    true,
								Computed:    true,
							},
							"permissions": {
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
										"object_type": {
											Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
											Type:        schema.TypeString,
											Optional:    true,
											Computed:    true,
										},
										"permission_identifier": {
											Description: "MOID of the permission which user has access to.",
											Type:        schema.TypeString,
											Optional:    true,
											Computed:    true,
										},
										"permission_name": {
											Description: "Name of the permission which user has access to.",
											Type:        schema.TypeString,
											Optional:    true,
											Computed:    true,
										},
									},
								},
								Computed: true,
							},
						},
					},
					Computed: true,
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
					"client_ip_address": {
						Description: "The user agent IP address from which the session is launched.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"expiration": {
						Description: "Expiration time for the session.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"idle_time_expiration": {
						Description: "Idle time expiration for the session.",
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
					"session_id": {
						Description: "Session token shared with the user agent which is used to identify the user session when API requests are received to perform authorization.",
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
				}},
				Computed: true,
			}},
	}
}

func dataSourceIamSessionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamSession{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("client_ip_address"); ok {
		x := (v.(string))
		o.SetClientIpAddress(x)
	}
	if v, ok := d.GetOk("expiration"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetExpiration(x)
	}
	if v, ok := d.GetOk("idle_time_expiration"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetIdleTimeExpiration(x)
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
	if v, ok := d.GetOk("session_id"); ok {
		x := (v.(string))
		o.SetSessionId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamSession object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamSessionList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamSessionList.GetCount()
	var i int32
	var iamSessionResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamSessionList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamSession: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamSessionList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamSession data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account_permissions"] = flattenListIamAccountPermissions(s.GetAccountPermissions(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["client_ip_address"] = (s.GetClientIpAddress())

				temp["expiration"] = (s.GetExpiration()).String()

				temp["idle_time_expiration"] = (s.GetIdleTimeExpiration()).String()
				temp["last_login_client"] = (s.GetLastLoginClient())

				temp["last_login_time"] = (s.GetLastLoginTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["permission"] = flattenMapIamPermissionRelationship(s.GetPermission(), d)
				temp["session_id"] = (s.GetSessionId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["user"] = flattenMapIamUserRelationship(s.GetUser(), d)
				iamSessionResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamSessionResults))
	if err := d.Set("results", iamSessionResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamSessionResults[0]["moid"].(string))
	return de
}
