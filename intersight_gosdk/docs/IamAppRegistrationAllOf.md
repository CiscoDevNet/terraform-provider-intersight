# IamAppRegistrationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.AppRegistration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.AppRegistration"]
**AdminStatus** | Pointer to **string** | Used to trigger the enable or disable action on the App Registration. These actions change the status of an App Registration. * &#x60;enable&#x60; - Used to enable a disabled API key/App Registration. If the API key/App Registration is already expired, this action has no effect. * &#x60;disable&#x60; - Used to disable an active API key/App Registration. If the API key/App Registration is already expired, this action has no effect. | [optional] [default to "enable"]
**ClientId** | Pointer to **string** | A unique identifier for the OAuth2 client application. The client ID is auto-generated when the AppRegistration object is created. | [optional] [readonly] 
**ClientName** | Pointer to **string** | App Registration name specified by user. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth2 client secret. The value of this property is generated when grantType includes &#39;client-credentials&#39;. Otherwise, no client-secret is generated. | [optional] 
**ClientType** | Pointer to **string** | The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1. * &#x60;public&#x60; - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application. * &#x60;confidential&#x60; - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \&quot;trusted\&quot; end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \&quot;client_credentials\&quot; grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession. | [optional] [default to "public"]
**Description** | Pointer to **string** | Description of the application. | [optional] 
**ExpiryDateTime** | Pointer to **time.Time** | The expiration date of the App Registration which is set at the time of its creation. Its value can only be assigned a date that falls within the range determined by the maximum expiration time configured at the account level. The expiry date can be edited to be earlier or later, provided it stays within the designated expiry period. This period is determined by adding the &#39;startTime&#39; property of the App Registration to the maximum expiry time configured at the account level. | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**IsNeverExpiring** | Pointer to **bool** | Used to mark the App Registration as a never-expiring App Registration. | [optional] [default to false]
**LastUsedIp** | Pointer to **string** | The ip address from which the App Registration was last used. | [optional] [readonly] 
**LastUsedTime** | Pointer to **time.Time** | The time at which the App Registration was last used. It is updated every 24 hours. | [optional] [readonly] 
**OperStatus** | Pointer to **string** | The current status of the App Registration that dictates the validity of the app. * &#x60;enabled&#x60; - An API key/App Registration having enabled status can be used for API invocation. * &#x60;disabled&#x60; - An API key/App Registration having disabled status cannot be used for API invocation. * &#x60;expired&#x60; - An API key/App Registration having expired status cannot be used for API invocation as the expiration date has passed. | [optional] [readonly] [default to "enabled"]
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RenewClientSecret** | Pointer to **bool** | Set value to true to renew the client-secret. Applicable to client_credentials grant type. | [optional] [default to false]
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**RevocationTimestamp** | Pointer to **time.Time** | Used to perform revocation for tokens of AppRegistration. Updated only internally is case Revoke property come from UI with value true. On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the corresponding App Registration. | [optional] [readonly] 
**Revoke** | Pointer to **bool** | Used to trigger update the revocationTimestamp value. If UI sent updating request with the Revoke value is true, then update RevocationTimestamp. | [optional] [default to false]
**ShowConsentScreen** | Pointer to **bool** | Set to true if consent screen needs to be shown during the OAuth login process. Applicable only for public AppRegistrations, means only &#39;authorization_code&#39; grantType. Note that consent screen will be shown on each login. | [optional] [default to false]
**StartTime** | Pointer to **time.Time** | The timestamp at which an expiry date was first set on this app registration.  For expiring App Registrations, this field is same as the create time of the App Registration. For never-expiring App Registrations, this field is set initially to zero time value. If a never-expiry App Registration is later changed to have an expiration, the timestamp marking the start of this transition is recorded in this field. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**OauthTokens** | Pointer to [**[]IamOAuthTokenRelationship**](IamOAuthTokenRelationship.md) | An array of relationships to iamOAuthToken resources. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamAppRegistrationAllOf

`func NewIamAppRegistrationAllOf(classId string, objectType string, ) *IamAppRegistrationAllOf`

NewIamAppRegistrationAllOf instantiates a new IamAppRegistrationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAppRegistrationAllOfWithDefaults

`func NewIamAppRegistrationAllOfWithDefaults() *IamAppRegistrationAllOf`

NewIamAppRegistrationAllOfWithDefaults instantiates a new IamAppRegistrationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAppRegistrationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAppRegistrationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAppRegistrationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAppRegistrationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAppRegistrationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAppRegistrationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *IamAppRegistrationAllOf) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *IamAppRegistrationAllOf) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *IamAppRegistrationAllOf) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *IamAppRegistrationAllOf) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetClientId

`func (o *IamAppRegistrationAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamAppRegistrationAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamAppRegistrationAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamAppRegistrationAllOf) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *IamAppRegistrationAllOf) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *IamAppRegistrationAllOf) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *IamAppRegistrationAllOf) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *IamAppRegistrationAllOf) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamAppRegistrationAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamAppRegistrationAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamAppRegistrationAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamAppRegistrationAllOf) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientType

`func (o *IamAppRegistrationAllOf) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *IamAppRegistrationAllOf) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *IamAppRegistrationAllOf) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *IamAppRegistrationAllOf) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDescription

`func (o *IamAppRegistrationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamAppRegistrationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamAppRegistrationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamAppRegistrationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *IamAppRegistrationAllOf) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *IamAppRegistrationAllOf) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *IamAppRegistrationAllOf) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *IamAppRegistrationAllOf) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetGrantTypes

`func (o *IamAppRegistrationAllOf) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *IamAppRegistrationAllOf) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *IamAppRegistrationAllOf) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *IamAppRegistrationAllOf) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### SetGrantTypesNil

`func (o *IamAppRegistrationAllOf) SetGrantTypesNil(b bool)`

 SetGrantTypesNil sets the value for GrantTypes to be an explicit nil

### UnsetGrantTypes
`func (o *IamAppRegistrationAllOf) UnsetGrantTypes()`

UnsetGrantTypes ensures that no value is present for GrantTypes, not even an explicit nil
### GetIsNeverExpiring

`func (o *IamAppRegistrationAllOf) GetIsNeverExpiring() bool`

GetIsNeverExpiring returns the IsNeverExpiring field if non-nil, zero value otherwise.

### GetIsNeverExpiringOk

`func (o *IamAppRegistrationAllOf) GetIsNeverExpiringOk() (*bool, bool)`

GetIsNeverExpiringOk returns a tuple with the IsNeverExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeverExpiring

`func (o *IamAppRegistrationAllOf) SetIsNeverExpiring(v bool)`

SetIsNeverExpiring sets IsNeverExpiring field to given value.

### HasIsNeverExpiring

`func (o *IamAppRegistrationAllOf) HasIsNeverExpiring() bool`

HasIsNeverExpiring returns a boolean if a field has been set.

### GetLastUsedIp

`func (o *IamAppRegistrationAllOf) GetLastUsedIp() string`

GetLastUsedIp returns the LastUsedIp field if non-nil, zero value otherwise.

### GetLastUsedIpOk

`func (o *IamAppRegistrationAllOf) GetLastUsedIpOk() (*string, bool)`

GetLastUsedIpOk returns a tuple with the LastUsedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIp

`func (o *IamAppRegistrationAllOf) SetLastUsedIp(v string)`

SetLastUsedIp sets LastUsedIp field to given value.

### HasLastUsedIp

`func (o *IamAppRegistrationAllOf) HasLastUsedIp() bool`

HasLastUsedIp returns a boolean if a field has been set.

### GetLastUsedTime

`func (o *IamAppRegistrationAllOf) GetLastUsedTime() time.Time`

GetLastUsedTime returns the LastUsedTime field if non-nil, zero value otherwise.

### GetLastUsedTimeOk

`func (o *IamAppRegistrationAllOf) GetLastUsedTimeOk() (*time.Time, bool)`

GetLastUsedTimeOk returns a tuple with the LastUsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedTime

`func (o *IamAppRegistrationAllOf) SetLastUsedTime(v time.Time)`

SetLastUsedTime sets LastUsedTime field to given value.

### HasLastUsedTime

`func (o *IamAppRegistrationAllOf) HasLastUsedTime() bool`

HasLastUsedTime returns a boolean if a field has been set.

### GetOperStatus

`func (o *IamAppRegistrationAllOf) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *IamAppRegistrationAllOf) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *IamAppRegistrationAllOf) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *IamAppRegistrationAllOf) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IamAppRegistrationAllOf) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IamAppRegistrationAllOf) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IamAppRegistrationAllOf) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IamAppRegistrationAllOf) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### SetRedirectUrisNil

`func (o *IamAppRegistrationAllOf) SetRedirectUrisNil(b bool)`

 SetRedirectUrisNil sets the value for RedirectUris to be an explicit nil

### UnsetRedirectUris
`func (o *IamAppRegistrationAllOf) UnsetRedirectUris()`

UnsetRedirectUris ensures that no value is present for RedirectUris, not even an explicit nil
### GetRenewClientSecret

`func (o *IamAppRegistrationAllOf) GetRenewClientSecret() bool`

GetRenewClientSecret returns the RenewClientSecret field if non-nil, zero value otherwise.

### GetRenewClientSecretOk

`func (o *IamAppRegistrationAllOf) GetRenewClientSecretOk() (*bool, bool)`

GetRenewClientSecretOk returns a tuple with the RenewClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewClientSecret

`func (o *IamAppRegistrationAllOf) SetRenewClientSecret(v bool)`

SetRenewClientSecret sets RenewClientSecret field to given value.

### HasRenewClientSecret

`func (o *IamAppRegistrationAllOf) HasRenewClientSecret() bool`

HasRenewClientSecret returns a boolean if a field has been set.

### GetResponseTypes

`func (o *IamAppRegistrationAllOf) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *IamAppRegistrationAllOf) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *IamAppRegistrationAllOf) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *IamAppRegistrationAllOf) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### SetResponseTypesNil

`func (o *IamAppRegistrationAllOf) SetResponseTypesNil(b bool)`

 SetResponseTypesNil sets the value for ResponseTypes to be an explicit nil

### UnsetResponseTypes
`func (o *IamAppRegistrationAllOf) UnsetResponseTypes()`

UnsetResponseTypes ensures that no value is present for ResponseTypes, not even an explicit nil
### GetRevocationTimestamp

`func (o *IamAppRegistrationAllOf) GetRevocationTimestamp() time.Time`

GetRevocationTimestamp returns the RevocationTimestamp field if non-nil, zero value otherwise.

### GetRevocationTimestampOk

`func (o *IamAppRegistrationAllOf) GetRevocationTimestampOk() (*time.Time, bool)`

GetRevocationTimestampOk returns a tuple with the RevocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimestamp

`func (o *IamAppRegistrationAllOf) SetRevocationTimestamp(v time.Time)`

SetRevocationTimestamp sets RevocationTimestamp field to given value.

### HasRevocationTimestamp

`func (o *IamAppRegistrationAllOf) HasRevocationTimestamp() bool`

HasRevocationTimestamp returns a boolean if a field has been set.

### GetRevoke

`func (o *IamAppRegistrationAllOf) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *IamAppRegistrationAllOf) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *IamAppRegistrationAllOf) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.

### HasRevoke

`func (o *IamAppRegistrationAllOf) HasRevoke() bool`

HasRevoke returns a boolean if a field has been set.

### GetShowConsentScreen

`func (o *IamAppRegistrationAllOf) GetShowConsentScreen() bool`

GetShowConsentScreen returns the ShowConsentScreen field if non-nil, zero value otherwise.

### GetShowConsentScreenOk

`func (o *IamAppRegistrationAllOf) GetShowConsentScreenOk() (*bool, bool)`

GetShowConsentScreenOk returns a tuple with the ShowConsentScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConsentScreen

`func (o *IamAppRegistrationAllOf) SetShowConsentScreen(v bool)`

SetShowConsentScreen sets ShowConsentScreen field to given value.

### HasShowConsentScreen

`func (o *IamAppRegistrationAllOf) HasShowConsentScreen() bool`

HasShowConsentScreen returns a boolean if a field has been set.

### GetStartTime

`func (o *IamAppRegistrationAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IamAppRegistrationAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IamAppRegistrationAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IamAppRegistrationAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAccount

`func (o *IamAppRegistrationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamAppRegistrationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamAppRegistrationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamAppRegistrationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOauthTokens

`func (o *IamAppRegistrationAllOf) GetOauthTokens() []IamOAuthTokenRelationship`

GetOauthTokens returns the OauthTokens field if non-nil, zero value otherwise.

### GetOauthTokensOk

`func (o *IamAppRegistrationAllOf) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool)`

GetOauthTokensOk returns a tuple with the OauthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokens

`func (o *IamAppRegistrationAllOf) SetOauthTokens(v []IamOAuthTokenRelationship)`

SetOauthTokens sets OauthTokens field to given value.

### HasOauthTokens

`func (o *IamAppRegistrationAllOf) HasOauthTokens() bool`

HasOauthTokens returns a boolean if a field has been set.

### SetOauthTokensNil

`func (o *IamAppRegistrationAllOf) SetOauthTokensNil(b bool)`

 SetOauthTokensNil sets the value for OauthTokens to be an explicit nil

### UnsetOauthTokens
`func (o *IamAppRegistrationAllOf) UnsetOauthTokens()`

UnsetOauthTokens ensures that no value is present for OauthTokens, not even an explicit nil
### GetPermission

`func (o *IamAppRegistrationAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamAppRegistrationAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamAppRegistrationAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamAppRegistrationAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoles

`func (o *IamAppRegistrationAllOf) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamAppRegistrationAllOf) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamAppRegistrationAllOf) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamAppRegistrationAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamAppRegistrationAllOf) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamAppRegistrationAllOf) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetUser

`func (o *IamAppRegistrationAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamAppRegistrationAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamAppRegistrationAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamAppRegistrationAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


