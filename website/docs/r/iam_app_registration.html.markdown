---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_app_registration"
description: |-
        AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in
        https://tools.ietf.org/html/rfc7591#section-2.
        Registered client applications have a set of metadata values associated with their client identifier
        at the Intersight authorization server, including the list of valid redirection URIs or a display name.
        The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently
        invoke Intersight API on behalf of this AppRegistration.
        To register an OAuth2 application, the following information must be provided.
        1) Application name
        2) An icon for the application
        3) URL to the application's home page
        4) A short description of the application
        5) A list of redirect URLs
        When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.

---

# Resource: intersight_iam_app_registration
AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in
https://tools.ietf.org/html/rfc7591#section-2.
Registered client applications have a set of metadata values associated with their client identifier
at the Intersight authorization server, including the list of valid redirection URIs or a display name.
The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently
invoke Intersight API on behalf of this AppRegistration.
To register an OAuth2 application, the following information must be provided.
1) Application name
2) An icon for the application
3) URL to the application's home page
4) A short description of the application
5) A list of redirect URLs
When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) -(ReadOnly) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `admin_status`:(string) Used to trigger the enable or disable action on the App Registration. These actions change the status of an App Registration.* `enable` - Used to enable a disabled API key/App Registration. If the API key/App Registration is already expired, this action has no effect.* `disable` - Used to disable an active API key/App Registration. If the API key/App Registration is already expired, this action has no effect. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `client_application`:(HashMap) -(ReadOnly) A reference to a iamAbstractClientApplication resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `client_id`:(string)(ReadOnly) A unique identifier for the OAuth2 client application.The client ID is auto-generated when the AppRegistration object is created. 
* `client_name`:(string) App Registration name specified by user. 
* `client_secret`:(string) The OAuth2 client secret.The value of this property is generated when grantType includes 'client-credentials'.Otherwise, no client-secret is generated. 
* `client_type`:(string) The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1.* `public` - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application.* `confidential` - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \ trusted\  end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \ client_credentials\  grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the application. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `entity_id`:(string)(ReadOnly) EntityId holds the Id of the client application that is using this AppRegistration. 
* `expiry_date_time`:(string) The expiration date of the App Registration which is set at the time of its creation. Its value can only be assigned a date that falls within the range determined by the maximum expiration time configured at the account level. The expiry date can be edited to be earlier or later, provided it stays within the designated expiry period. This period is determined by adding the 'startTime' property of the App Registration to the maximum expiry time configured at the account level. 
* `grant_types`:
                (Array of schema.TypeString) -
* `is_never_expiring`:(bool) Used to mark the App Registration as a never-expiring App Registration. 
* `last_used_ip`:(string)(ReadOnly) The ip address from which the App Registration was last used. 
* `last_used_time`:(string)(ReadOnly) The time at which the App Registration was last used. It is updated every 24 hours. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oauth_tokens`:(Array)(ReadOnly) An array of relationships to iamOAuthToken resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `oper_status`:(string)(ReadOnly) The current status of the App Registration that dictates the validity of the app.* `enabled` - An API key/App Registration having enabled status can be used for API invocation.* `disabled` - An API key/App Registration having disabled status cannot be used for API invocation.* `expired` - An API key/App Registration having expired status cannot be used for API invocation as the expiration date has passed. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission`:(HashMap) -(ReadOnly) A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `redirect_uris`:
                (Array of schema.TypeString) -
* `renew_client_secret`:(bool) Set value to true to renew the client-secret. Applicable to client_credentials grant type. 
* `response_types`:
                (Array of schema.TypeString) -
* `revocation_timestamp`:(string)(ReadOnly) Used to perform revocation for tokens of AppRegistration.Updated only internally is case Revoke property come from UI with value true.On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of thecorresponding App Registration. 
* `revoke`:(bool) Used to trigger update the revocationTimestamp value.If UI sent updating request with the Revoke value is true, then update RevocationTimestamp. 
* `roles`:(Array) An array of relationships to iamRole resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `scope`:(HashMap) -(ReadOnly) Scope holds a collection of account Id, permission Id to which the current session is scoped to. 
This complex property has following sub-properties:
  + `account_access_control_id`:(string)(ReadOnly) Moid of the AccountAccessControl through which the access is given to switch scope. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `request_identifier`:(string)(ReadOnly) Stores the identifier of the issue for which user is trying to switch scope to another account. 
  + `switched_from_account`:(HashMap) -(ReadOnly) Permission for the Account from which user switched the scope. 
This complex property has following sub-properties:
    + `account_id`:(string)(ReadOnly) Moid of the Account to/from which user switched the scope. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `permission_id`:(string)(ReadOnly) Moid of the Permission for the Account to/from which user switched the scope. 
  + `switched_to_accounts`:(Array)
This complex property has following sub-properties:
    + `account_id`:(string)(ReadOnly) Moid of the Account to/from which user switched the scope. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `permission_id`:(string)(ReadOnly) Moid of the Permission for the Account to/from which user switched the scope. 
* `security_holder`:(HashMap) -(ReadOnly) A reference to a iamSecurityHolder resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `show_consent_screen`:(bool) Set to true if consent screen needs to be shown during the OAuth login process.Applicable only for public AppRegistrations, means only 'authorization_code' grantType.Note that consent screen will be shown on each login. 
* `start_time`:(string)(ReadOnly) The timestamp at which an expiry date was first set on this app registration. For expiring App Registrations, this field is same as the create time of the App Registration.For never-expiring App Registrations, this field is set initially to zero time value. If a never-expiry App Registration is later changed to have an expiration, the timestamp marking the start of this transition is recorded in this field. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `token_expiry`:(string)(ReadOnly) Defines the expiry time of the token generated via the AppRegistration.* `generic` - This sets the expiryTime to ten minutes from the token issuing time.* `longLived` - This sets the expiryTime to an year from the token issuing time.* `infinite` - This allows for a never-expiring token. Use with caution. 
* `user`:(HashMap) -(ReadOnly) A reference to a iamUser resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_iam_app_registration` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_app_registration.example 1234567890987654321abcde
``` 
