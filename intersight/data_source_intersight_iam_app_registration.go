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

func dataSourceIamAppRegistration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamAppRegistrationRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_id": {
				Description: "A unique identifier for the OAuth2 client application.\nThe client ID is auto-generated when the AppRegistration object is created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"client_name": {
				Description: "App Registration name specified by user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_secret": {
				Description: "The OAuth2 client secret.\nThe value of this property is generated when grantType includes 'client-credentials'.\nOtherwise, no client-secret is generated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"client_type": {
				Description: "The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1.\n* `public` - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application.\n* `confidential` - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \"trusted\" end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \"client_credentials\" grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the application.",
				Type:        schema.TypeString,
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
			"renew_client_secret": {
				Description: "Set value to true to renew the client-secret. Applicable to client_credentials grant type.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"revocation_timestamp": {
				Description: "Used to perform revocation for tokens of AppRegistration.\nUpdated only internally is case Revoke property come from UI with value true.\nOn each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the\ncorresponding App Registration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"revoke": {
				Description: "Used to trigger update the revocationTimestamp value.\nIf UI sent updating request with the Revoke value is true, then update RevocationTimestamp.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceIamAppRegistration().Schema},
				Computed: true,
			}},
	}
}

func dataSourceIamAppRegistrationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IamAppRegistration{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("client_id"); ok {
		x := (v.(string))
		o.SetClientId(x)
	}
	if v, ok := d.GetOk("client_name"); ok {
		x := (v.(string))
		o.SetClientName(x)
	}
	if v, ok := d.GetOk("client_secret"); ok {
		x := (v.(string))
		o.SetClientSecret(x)
	}
	if v, ok := d.GetOk("client_type"); ok {
		x := (v.(string))
		o.SetClientType(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("renew_client_secret"); ok {
		x := (v.(bool))
		o.SetRenewClientSecret(x)
	}
	if v, ok := d.GetOk("revocation_timestamp"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetRevocationTimestamp(x)
	}
	if v, ok := d.GetOk("revoke"); ok {
		x := (v.(bool))
		o.SetRevoke(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IamAppRegistration object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.IamApi.GetIamAppRegistrationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of IamAppRegistration: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.IamAppRegistrationList.GetCount()
	var i int32
	var iamAppRegistrationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.IamApi.GetIamAppRegistrationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching IamAppRegistration: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.IamAppRegistrationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for IamAppRegistration data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["client_id"] = (s.GetClientId())
				temp["client_name"] = (s.GetClientName())
				temp["client_secret"] = (s.GetClientSecret())
				temp["client_type"] = (s.GetClientType())
				temp["description"] = (s.GetDescription())
				temp["grant_types"] = (s.GetGrantTypes())
				temp["moid"] = (s.GetMoid())

				temp["oauth_tokens"] = flattenListIamOAuthTokenRelationship(s.GetOauthTokens(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["permission"] = flattenMapIamPermissionRelationship(s.GetPermission(), d)
				temp["redirect_uris"] = (s.GetRedirectUris())
				temp["renew_client_secret"] = (s.GetRenewClientSecret())
				temp["response_types"] = (s.GetResponseTypes())

				temp["revocation_timestamp"] = (s.GetRevocationTimestamp()).String()
				temp["revoke"] = (s.GetRevoke())

				temp["roles"] = flattenListIamRoleRelationship(s.GetRoles(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["user"] = flattenMapIamUserRelationship(s.GetUser(), d)
				iamAppRegistrationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(iamAppRegistrationResults))
	if err := d.Set("results", iamAppRegistrationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(iamAppRegistrationResults[0]["moid"].(string))
	return de
}
