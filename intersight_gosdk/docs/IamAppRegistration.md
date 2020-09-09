# IamAppRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | A unique identifier for the OAuth2 client application. The client ID is auto-generated when the AppRegistration object is created. | [optional] [readonly] 
**ClientName** | Pointer to **string** | App Registration name specified by user. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth2 client secret. The value of this property is generated when grantType includes &#39;client-credentials&#39;. Otherwise, no client-secret is generated. | [optional] 
**ClientType** | Pointer to **string** | The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1. * &#x60;public&#x60; - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application. * &#x60;confidential&#x60; - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \&quot;trusted\&quot; end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \&quot;client_credentials\&quot; grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession. | [optional] [default to "public"]
**Description** | Pointer to **string** | Description of the application. | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**RenewClientSecret** | Pointer to **bool** | Set value to true to renew the client-secret. Applicable to client_credentials grant type. | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**RevocationTimestamp** | Pointer to [**time.Time**](time.Time.md) | Used to perform revocation for tokens of AppRegistration. Updated only internally is case Revoke property come from UI with value true. On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the corresponding App Registration. | [optional] [readonly] 
**Revoke** | Pointer to **bool** | Used to trigger update the revocationTimestamp value. If UI sent updating request with the Revoke value is true, then update RevocationTimestamp. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**OauthTokens** | Pointer to [**[]IamOAuthTokenRelationship**](iam.OAuthToken.Relationship.md) | An array of relationships to iamOAuthToken resources. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](iam.Role.Relationship.md) | An array of relationships to iamRole resources. | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamAppRegistration

`func NewIamAppRegistration() *IamAppRegistration`

NewIamAppRegistration instantiates a new IamAppRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAppRegistrationWithDefaults

`func NewIamAppRegistrationWithDefaults() *IamAppRegistration`

NewIamAppRegistrationWithDefaults instantiates a new IamAppRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IamAppRegistration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamAppRegistration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamAppRegistration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamAppRegistration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *IamAppRegistration) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *IamAppRegistration) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *IamAppRegistration) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *IamAppRegistration) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientSecret

`func (o *IamAppRegistration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IamAppRegistration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IamAppRegistration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *IamAppRegistration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientType

`func (o *IamAppRegistration) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *IamAppRegistration) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *IamAppRegistration) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *IamAppRegistration) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDescription

`func (o *IamAppRegistration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamAppRegistration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamAppRegistration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamAppRegistration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrantTypes

`func (o *IamAppRegistration) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *IamAppRegistration) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *IamAppRegistration) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *IamAppRegistration) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *IamAppRegistration) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *IamAppRegistration) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *IamAppRegistration) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *IamAppRegistration) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRenewClientSecret

`func (o *IamAppRegistration) GetRenewClientSecret() bool`

GetRenewClientSecret returns the RenewClientSecret field if non-nil, zero value otherwise.

### GetRenewClientSecretOk

`func (o *IamAppRegistration) GetRenewClientSecretOk() (*bool, bool)`

GetRenewClientSecretOk returns a tuple with the RenewClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewClientSecret

`func (o *IamAppRegistration) SetRenewClientSecret(v bool)`

SetRenewClientSecret sets RenewClientSecret field to given value.

### HasRenewClientSecret

`func (o *IamAppRegistration) HasRenewClientSecret() bool`

HasRenewClientSecret returns a boolean if a field has been set.

### GetResponseTypes

`func (o *IamAppRegistration) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *IamAppRegistration) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *IamAppRegistration) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *IamAppRegistration) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRevocationTimestamp

`func (o *IamAppRegistration) GetRevocationTimestamp() time.Time`

GetRevocationTimestamp returns the RevocationTimestamp field if non-nil, zero value otherwise.

### GetRevocationTimestampOk

`func (o *IamAppRegistration) GetRevocationTimestampOk() (*time.Time, bool)`

GetRevocationTimestampOk returns a tuple with the RevocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimestamp

`func (o *IamAppRegistration) SetRevocationTimestamp(v time.Time)`

SetRevocationTimestamp sets RevocationTimestamp field to given value.

### HasRevocationTimestamp

`func (o *IamAppRegistration) HasRevocationTimestamp() bool`

HasRevocationTimestamp returns a boolean if a field has been set.

### GetRevoke

`func (o *IamAppRegistration) GetRevoke() bool`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *IamAppRegistration) GetRevokeOk() (*bool, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *IamAppRegistration) SetRevoke(v bool)`

SetRevoke sets Revoke field to given value.

### HasRevoke

`func (o *IamAppRegistration) HasRevoke() bool`

HasRevoke returns a boolean if a field has been set.

### GetAccount

`func (o *IamAppRegistration) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamAppRegistration) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamAppRegistration) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamAppRegistration) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOauthTokens

`func (o *IamAppRegistration) GetOauthTokens() []IamOAuthTokenRelationship`

GetOauthTokens returns the OauthTokens field if non-nil, zero value otherwise.

### GetOauthTokensOk

`func (o *IamAppRegistration) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool)`

GetOauthTokensOk returns a tuple with the OauthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokens

`func (o *IamAppRegistration) SetOauthTokens(v []IamOAuthTokenRelationship)`

SetOauthTokens sets OauthTokens field to given value.

### HasOauthTokens

`func (o *IamAppRegistration) HasOauthTokens() bool`

HasOauthTokens returns a boolean if a field has been set.

### SetOauthTokensNil

`func (o *IamAppRegistration) SetOauthTokensNil(b bool)`

 SetOauthTokensNil sets the value for OauthTokens to be an explicit nil

### UnsetOauthTokens
`func (o *IamAppRegistration) UnsetOauthTokens()`

UnsetOauthTokens ensures that no value is present for OauthTokens, not even an explicit nil
### GetPermission

`func (o *IamAppRegistration) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamAppRegistration) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamAppRegistration) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamAppRegistration) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRoles

`func (o *IamAppRegistration) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamAppRegistration) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamAppRegistration) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamAppRegistration) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamAppRegistration) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamAppRegistration) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetUser

`func (o *IamAppRegistration) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamAppRegistration) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamAppRegistration) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamAppRegistration) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


