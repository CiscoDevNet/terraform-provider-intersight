# IamIdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Idp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Idp"]
**DomainName** | Pointer to **string** | Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. | [optional] 
**DomainNames** | Pointer to **[]string** |  | [optional] 
**EnableSingleLogout** | Pointer to **bool** | Setting that indicates whether &#39;Single Logout (SLO)&#39; has been enabled for this IdP. | [optional] 
**IdpEntityId** | Pointer to **string** | The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider. | [optional] [readonly] 
**Metadata** | Pointer to **string** | SAML metadata of the IdP. | [optional] 
**Name** | Pointer to **string** | The name of the Identity Provider, for example Cisco, Okta, or OneID. | [optional] 
**SkipWarning** | Pointer to **bool** | When users attempt the Account URL login with an unverified Domain Name, they get a warning stating that they are logging in using an unverified Domain Name. Enable the slider if you do not wish to see the warning message. | [optional] 
**Type** | Pointer to **string** | Authentication protocol used by the IdP. * &#x60;saml&#x60; - Use SAML as the authentication protocol for sign-on. * &#x60;oidc&#x60; - Open ID connect to be used as an authentication protocol for sign-on. * &#x60;local&#x60; - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP. | [optional] [default to "saml"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**LdapMeta** | Pointer to [**NullableIamLdapMetaRelationship**](IamLdapMetaRelationship.md) |  | [optional] 
**LdapPolicy** | Pointer to [**NullableIamLdapPolicyRelationship**](IamLdapPolicyRelationship.md) |  | [optional] 
**System** | Pointer to [**NullableIamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 
**TestIdp** | Pointer to [**NullableIamTestIdpConfigurationRelationship**](IamTestIdpConfigurationRelationship.md) |  | [optional] 
**UserPreferences** | Pointer to [**[]IamUserPreferenceRelationship**](IamUserPreferenceRelationship.md) | An array of relationships to iamUserPreference resources. | [optional] [readonly] 
**UserSettings** | Pointer to [**[]IamUserSettingRelationship**](IamUserSettingRelationship.md) | An array of relationships to iamUserSetting resources. | [optional] [readonly] 
**Usergroups** | Pointer to [**[]IamUserGroupRelationship**](IamUserGroupRelationship.md) | An array of relationships to iamUserGroup resources. | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](IamUserRelationship.md) | An array of relationships to iamUser resources. | [optional] 

## Methods

### NewIamIdp

`func NewIamIdp(classId string, objectType string, ) *IamIdp`

NewIamIdp instantiates a new IamIdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdpWithDefaults

`func NewIamIdpWithDefaults() *IamIdp`

NewIamIdpWithDefaults instantiates a new IamIdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIdp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIdp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIdp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIdp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIdp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIdp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomainName

`func (o *IamIdp) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamIdp) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamIdp) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamIdp) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNames

`func (o *IamIdp) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *IamIdp) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *IamIdp) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *IamIdp) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *IamIdp) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *IamIdp) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetEnableSingleLogout

`func (o *IamIdp) GetEnableSingleLogout() bool`

GetEnableSingleLogout returns the EnableSingleLogout field if non-nil, zero value otherwise.

### GetEnableSingleLogoutOk

`func (o *IamIdp) GetEnableSingleLogoutOk() (*bool, bool)`

GetEnableSingleLogoutOk returns a tuple with the EnableSingleLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleLogout

`func (o *IamIdp) SetEnableSingleLogout(v bool)`

SetEnableSingleLogout sets EnableSingleLogout field to given value.

### HasEnableSingleLogout

`func (o *IamIdp) HasEnableSingleLogout() bool`

HasEnableSingleLogout returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamIdp) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamIdp) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamIdp) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamIdp) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamIdp) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamIdp) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamIdp) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamIdp) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamIdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamIdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamIdp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamIdp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipWarning

`func (o *IamIdp) GetSkipWarning() bool`

GetSkipWarning returns the SkipWarning field if non-nil, zero value otherwise.

### GetSkipWarningOk

`func (o *IamIdp) GetSkipWarningOk() (*bool, bool)`

GetSkipWarningOk returns a tuple with the SkipWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWarning

`func (o *IamIdp) SetSkipWarning(v bool)`

SetSkipWarning sets SkipWarning field to given value.

### HasSkipWarning

`func (o *IamIdp) HasSkipWarning() bool`

HasSkipWarning returns a boolean if a field has been set.

### GetType

`func (o *IamIdp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamIdp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamIdp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamIdp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *IamIdp) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamIdp) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamIdp) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamIdp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamIdp) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamIdp) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetLdapMeta

`func (o *IamIdp) GetLdapMeta() IamLdapMetaRelationship`

GetLdapMeta returns the LdapMeta field if non-nil, zero value otherwise.

### GetLdapMetaOk

`func (o *IamIdp) GetLdapMetaOk() (*IamLdapMetaRelationship, bool)`

GetLdapMetaOk returns a tuple with the LdapMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapMeta

`func (o *IamIdp) SetLdapMeta(v IamLdapMetaRelationship)`

SetLdapMeta sets LdapMeta field to given value.

### HasLdapMeta

`func (o *IamIdp) HasLdapMeta() bool`

HasLdapMeta returns a boolean if a field has been set.

### SetLdapMetaNil

`func (o *IamIdp) SetLdapMetaNil(b bool)`

 SetLdapMetaNil sets the value for LdapMeta to be an explicit nil

### UnsetLdapMeta
`func (o *IamIdp) UnsetLdapMeta()`

UnsetLdapMeta ensures that no value is present for LdapMeta, not even an explicit nil
### GetLdapPolicy

`func (o *IamIdp) GetLdapPolicy() IamLdapPolicyRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamIdp) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamIdp) SetLdapPolicy(v IamLdapPolicyRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamIdp) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.

### SetLdapPolicyNil

`func (o *IamIdp) SetLdapPolicyNil(b bool)`

 SetLdapPolicyNil sets the value for LdapPolicy to be an explicit nil

### UnsetLdapPolicy
`func (o *IamIdp) UnsetLdapPolicy()`

UnsetLdapPolicy ensures that no value is present for LdapPolicy, not even an explicit nil
### GetSystem

`func (o *IamIdp) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamIdp) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamIdp) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamIdp) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *IamIdp) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *IamIdp) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTestIdp

`func (o *IamIdp) GetTestIdp() IamTestIdpConfigurationRelationship`

GetTestIdp returns the TestIdp field if non-nil, zero value otherwise.

### GetTestIdpOk

`func (o *IamIdp) GetTestIdpOk() (*IamTestIdpConfigurationRelationship, bool)`

GetTestIdpOk returns a tuple with the TestIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestIdp

`func (o *IamIdp) SetTestIdp(v IamTestIdpConfigurationRelationship)`

SetTestIdp sets TestIdp field to given value.

### HasTestIdp

`func (o *IamIdp) HasTestIdp() bool`

HasTestIdp returns a boolean if a field has been set.

### SetTestIdpNil

`func (o *IamIdp) SetTestIdpNil(b bool)`

 SetTestIdpNil sets the value for TestIdp to be an explicit nil

### UnsetTestIdp
`func (o *IamIdp) UnsetTestIdp()`

UnsetTestIdp ensures that no value is present for TestIdp, not even an explicit nil
### GetUserPreferences

`func (o *IamIdp) GetUserPreferences() []IamUserPreferenceRelationship`

GetUserPreferences returns the UserPreferences field if non-nil, zero value otherwise.

### GetUserPreferencesOk

`func (o *IamIdp) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool)`

GetUserPreferencesOk returns a tuple with the UserPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPreferences

`func (o *IamIdp) SetUserPreferences(v []IamUserPreferenceRelationship)`

SetUserPreferences sets UserPreferences field to given value.

### HasUserPreferences

`func (o *IamIdp) HasUserPreferences() bool`

HasUserPreferences returns a boolean if a field has been set.

### SetUserPreferencesNil

`func (o *IamIdp) SetUserPreferencesNil(b bool)`

 SetUserPreferencesNil sets the value for UserPreferences to be an explicit nil

### UnsetUserPreferences
`func (o *IamIdp) UnsetUserPreferences()`

UnsetUserPreferences ensures that no value is present for UserPreferences, not even an explicit nil
### GetUserSettings

`func (o *IamIdp) GetUserSettings() []IamUserSettingRelationship`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *IamIdp) GetUserSettingsOk() (*[]IamUserSettingRelationship, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *IamIdp) SetUserSettings(v []IamUserSettingRelationship)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *IamIdp) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.

### SetUserSettingsNil

`func (o *IamIdp) SetUserSettingsNil(b bool)`

 SetUserSettingsNil sets the value for UserSettings to be an explicit nil

### UnsetUserSettings
`func (o *IamIdp) UnsetUserSettings()`

UnsetUserSettings ensures that no value is present for UserSettings, not even an explicit nil
### GetUsergroups

`func (o *IamIdp) GetUsergroups() []IamUserGroupRelationship`

GetUsergroups returns the Usergroups field if non-nil, zero value otherwise.

### GetUsergroupsOk

`func (o *IamIdp) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUsergroupsOk returns a tuple with the Usergroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroups

`func (o *IamIdp) SetUsergroups(v []IamUserGroupRelationship)`

SetUsergroups sets Usergroups field to given value.

### HasUsergroups

`func (o *IamIdp) HasUsergroups() bool`

HasUsergroups returns a boolean if a field has been set.

### SetUsergroupsNil

`func (o *IamIdp) SetUsergroupsNil(b bool)`

 SetUsergroupsNil sets the value for Usergroups to be an explicit nil

### UnsetUsergroups
`func (o *IamIdp) UnsetUsergroups()`

UnsetUsergroups ensures that no value is present for Usergroups, not even an explicit nil
### GetUsers

`func (o *IamIdp) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamIdp) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamIdp) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamIdp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamIdp) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamIdp) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


