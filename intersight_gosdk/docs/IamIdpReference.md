# IamIdpReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.IdpReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.IdpReference"]
**DomainName** | Pointer to **string** | The email domain name for this IdP of the user. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. | [optional] [readonly] 
**IdpEntityId** | Pointer to **string** | Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP/Service Provider. | [optional] [readonly] 
**MultiFactorAuthentication** | Pointer to **bool** | The flag represents if the second factor of authentication is required for Cisco IdP users. | [optional] [default to false]
**Name** | Pointer to **string** | Cisco IdP reference in an account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**UserPreferences** | Pointer to [**[]IamUserPreferenceRelationship**](IamUserPreferenceRelationship.md) | An array of relationships to iamUserPreference resources. | [optional] [readonly] 
**Usergroups** | Pointer to [**[]IamUserGroupRelationship**](IamUserGroupRelationship.md) | An array of relationships to iamUserGroup resources. | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](IamUserRelationship.md) | An array of relationships to iamUser resources. | [optional] 

## Methods

### NewIamIdpReference

`func NewIamIdpReference(classId string, objectType string, ) *IamIdpReference`

NewIamIdpReference instantiates a new IamIdpReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdpReferenceWithDefaults

`func NewIamIdpReferenceWithDefaults() *IamIdpReference`

NewIamIdpReferenceWithDefaults instantiates a new IamIdpReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIdpReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIdpReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIdpReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIdpReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIdpReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIdpReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomainName

`func (o *IamIdpReference) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamIdpReference) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamIdpReference) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamIdpReference) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamIdpReference) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamIdpReference) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamIdpReference) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamIdpReference) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetMultiFactorAuthentication

`func (o *IamIdpReference) GetMultiFactorAuthentication() bool`

GetMultiFactorAuthentication returns the MultiFactorAuthentication field if non-nil, zero value otherwise.

### GetMultiFactorAuthenticationOk

`func (o *IamIdpReference) GetMultiFactorAuthenticationOk() (*bool, bool)`

GetMultiFactorAuthenticationOk returns a tuple with the MultiFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthentication

`func (o *IamIdpReference) SetMultiFactorAuthentication(v bool)`

SetMultiFactorAuthentication sets MultiFactorAuthentication field to given value.

### HasMultiFactorAuthentication

`func (o *IamIdpReference) HasMultiFactorAuthentication() bool`

HasMultiFactorAuthentication returns a boolean if a field has been set.

### GetName

`func (o *IamIdpReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamIdpReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamIdpReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamIdpReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *IamIdpReference) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamIdpReference) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamIdpReference) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamIdpReference) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIdp

`func (o *IamIdpReference) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamIdpReference) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamIdpReference) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamIdpReference) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetUserPreferences

`func (o *IamIdpReference) GetUserPreferences() []IamUserPreferenceRelationship`

GetUserPreferences returns the UserPreferences field if non-nil, zero value otherwise.

### GetUserPreferencesOk

`func (o *IamIdpReference) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool)`

GetUserPreferencesOk returns a tuple with the UserPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPreferences

`func (o *IamIdpReference) SetUserPreferences(v []IamUserPreferenceRelationship)`

SetUserPreferences sets UserPreferences field to given value.

### HasUserPreferences

`func (o *IamIdpReference) HasUserPreferences() bool`

HasUserPreferences returns a boolean if a field has been set.

### SetUserPreferencesNil

`func (o *IamIdpReference) SetUserPreferencesNil(b bool)`

 SetUserPreferencesNil sets the value for UserPreferences to be an explicit nil

### UnsetUserPreferences
`func (o *IamIdpReference) UnsetUserPreferences()`

UnsetUserPreferences ensures that no value is present for UserPreferences, not even an explicit nil
### GetUsergroups

`func (o *IamIdpReference) GetUsergroups() []IamUserGroupRelationship`

GetUsergroups returns the Usergroups field if non-nil, zero value otherwise.

### GetUsergroupsOk

`func (o *IamIdpReference) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUsergroupsOk returns a tuple with the Usergroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroups

`func (o *IamIdpReference) SetUsergroups(v []IamUserGroupRelationship)`

SetUsergroups sets Usergroups field to given value.

### HasUsergroups

`func (o *IamIdpReference) HasUsergroups() bool`

HasUsergroups returns a boolean if a field has been set.

### SetUsergroupsNil

`func (o *IamIdpReference) SetUsergroupsNil(b bool)`

 SetUsergroupsNil sets the value for Usergroups to be an explicit nil

### UnsetUsergroups
`func (o *IamIdpReference) UnsetUsergroups()`

UnsetUsergroups ensures that no value is present for Usergroups, not even an explicit nil
### GetUsers

`func (o *IamIdpReference) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamIdpReference) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamIdpReference) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamIdpReference) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamIdpReference) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamIdpReference) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


