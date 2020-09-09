# IamIdpReferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. | [optional] [readonly] 
**IdpEntityId** | Pointer to **string** | Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. | [optional] [readonly] 
**MultiFactorAuthentication** | Pointer to **bool** | The flag represents if the second factor of authentication is required for Cisco IdP users. | [optional] 
**Name** | Pointer to **string** | Cisco IdP reference in an account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Idp** | Pointer to [**IamIdpRelationship**](iam.Idp.Relationship.md) |  | [optional] 
**UserPreferences** | Pointer to [**[]IamUserPreferenceRelationship**](iam.UserPreference.Relationship.md) | An array of relationships to iamUserPreference resources. | [optional] [readonly] 
**Usergroups** | Pointer to [**[]IamUserGroupRelationship**](iam.UserGroup.Relationship.md) | An array of relationships to iamUserGroup resources. | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](iam.User.Relationship.md) | An array of relationships to iamUser resources. | [optional] 

## Methods

### NewIamIdpReferenceAllOf

`func NewIamIdpReferenceAllOf() *IamIdpReferenceAllOf`

NewIamIdpReferenceAllOf instantiates a new IamIdpReferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdpReferenceAllOfWithDefaults

`func NewIamIdpReferenceAllOfWithDefaults() *IamIdpReferenceAllOf`

NewIamIdpReferenceAllOfWithDefaults instantiates a new IamIdpReferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *IamIdpReferenceAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamIdpReferenceAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamIdpReferenceAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamIdpReferenceAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamIdpReferenceAllOf) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamIdpReferenceAllOf) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamIdpReferenceAllOf) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamIdpReferenceAllOf) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetMultiFactorAuthentication

`func (o *IamIdpReferenceAllOf) GetMultiFactorAuthentication() bool`

GetMultiFactorAuthentication returns the MultiFactorAuthentication field if non-nil, zero value otherwise.

### GetMultiFactorAuthenticationOk

`func (o *IamIdpReferenceAllOf) GetMultiFactorAuthenticationOk() (*bool, bool)`

GetMultiFactorAuthenticationOk returns a tuple with the MultiFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthentication

`func (o *IamIdpReferenceAllOf) SetMultiFactorAuthentication(v bool)`

SetMultiFactorAuthentication sets MultiFactorAuthentication field to given value.

### HasMultiFactorAuthentication

`func (o *IamIdpReferenceAllOf) HasMultiFactorAuthentication() bool`

HasMultiFactorAuthentication returns a boolean if a field has been set.

### GetName

`func (o *IamIdpReferenceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamIdpReferenceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamIdpReferenceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamIdpReferenceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *IamIdpReferenceAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamIdpReferenceAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamIdpReferenceAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamIdpReferenceAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIdp

`func (o *IamIdpReferenceAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamIdpReferenceAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamIdpReferenceAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamIdpReferenceAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetUserPreferences

`func (o *IamIdpReferenceAllOf) GetUserPreferences() []IamUserPreferenceRelationship`

GetUserPreferences returns the UserPreferences field if non-nil, zero value otherwise.

### GetUserPreferencesOk

`func (o *IamIdpReferenceAllOf) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool)`

GetUserPreferencesOk returns a tuple with the UserPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPreferences

`func (o *IamIdpReferenceAllOf) SetUserPreferences(v []IamUserPreferenceRelationship)`

SetUserPreferences sets UserPreferences field to given value.

### HasUserPreferences

`func (o *IamIdpReferenceAllOf) HasUserPreferences() bool`

HasUserPreferences returns a boolean if a field has been set.

### SetUserPreferencesNil

`func (o *IamIdpReferenceAllOf) SetUserPreferencesNil(b bool)`

 SetUserPreferencesNil sets the value for UserPreferences to be an explicit nil

### UnsetUserPreferences
`func (o *IamIdpReferenceAllOf) UnsetUserPreferences()`

UnsetUserPreferences ensures that no value is present for UserPreferences, not even an explicit nil
### GetUsergroups

`func (o *IamIdpReferenceAllOf) GetUsergroups() []IamUserGroupRelationship`

GetUsergroups returns the Usergroups field if non-nil, zero value otherwise.

### GetUsergroupsOk

`func (o *IamIdpReferenceAllOf) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUsergroupsOk returns a tuple with the Usergroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroups

`func (o *IamIdpReferenceAllOf) SetUsergroups(v []IamUserGroupRelationship)`

SetUsergroups sets Usergroups field to given value.

### HasUsergroups

`func (o *IamIdpReferenceAllOf) HasUsergroups() bool`

HasUsergroups returns a boolean if a field has been set.

### SetUsergroupsNil

`func (o *IamIdpReferenceAllOf) SetUsergroupsNil(b bool)`

 SetUsergroupsNil sets the value for Usergroups to be an explicit nil

### UnsetUsergroups
`func (o *IamIdpReferenceAllOf) UnsetUsergroups()`

UnsetUsergroups ensures that no value is present for Usergroups, not even an explicit nil
### GetUsers

`func (o *IamIdpReferenceAllOf) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamIdpReferenceAllOf) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamIdpReferenceAllOf) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamIdpReferenceAllOf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamIdpReferenceAllOf) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamIdpReferenceAllOf) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


