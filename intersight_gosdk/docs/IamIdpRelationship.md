# IamIdpRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**DomainName** | Pointer to **string** | Email domain name of the user for this IdP. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. | [optional] 
**IdpEntityId** | Pointer to **string** | The Entity ID of the IdP. In SAML, the entity ID uniquely identifies the IdP or Service Provider. | [optional] [readonly] 
**Metadata** | Pointer to **string** | SAML metadata of the IdP. | [optional] 
**Name** | Pointer to **string** | The name of the Identity Provider, for example Cisco, Okta, or OneID. | [optional] 
**Type** | Pointer to **string** | Authentication protocol used by the IdP. * &#x60;saml&#x60; - Use SAML as the authentication protocol for sign-on. * &#x60;oidc&#x60; - Open ID connect to be used as an authentication protocol for sign-on. * &#x60;local&#x60; - The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP. | [optional] [default to "saml"]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**LdapPolicy** | Pointer to [**IamLdapPolicyRelationship**](iam.LdapPolicy.Relationship.md) |  | [optional] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 
**UserPreferences** | Pointer to [**[]IamUserPreferenceRelationship**](iam.UserPreference.Relationship.md) | An array of relationships to iamUserPreference resources. | [optional] [readonly] 
**Usergroups** | Pointer to [**[]IamUserGroupRelationship**](iam.UserGroup.Relationship.md) | An array of relationships to iamUserGroup resources. | [optional] 
**Users** | Pointer to [**[]IamUserRelationship**](iam.User.Relationship.md) | An array of relationships to iamUser resources. | [optional] 

## Methods

### NewIamIdpRelationship

`func NewIamIdpRelationship(classId string, objectType string, ) *IamIdpRelationship`

NewIamIdpRelationship instantiates a new IamIdpRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdpRelationshipWithDefaults

`func NewIamIdpRelationshipWithDefaults() *IamIdpRelationship`

NewIamIdpRelationshipWithDefaults instantiates a new IamIdpRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamIdpRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamIdpRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamIdpRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamIdpRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamIdpRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIdpRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIdpRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamIdpRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamIdpRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamIdpRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamIdpRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamIdpRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamIdpRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamIdpRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamIdpRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamIdpRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamIdpRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamIdpRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamIdpRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamIdpRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamIdpRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamIdpRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamIdpRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamIdpRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIdpRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIdpRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamIdpRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamIdpRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamIdpRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamIdpRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamIdpRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamIdpRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamIdpRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamIdpRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamIdpRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamIdpRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamIdpRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamIdpRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamIdpRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamIdpRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamIdpRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamIdpRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamIdpRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamIdpRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamIdpRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamIdpRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamIdpRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamIdpRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamIdpRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamIdpRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamIdpRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamIdpRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamIdpRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamIdpRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamIdpRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamIdpRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamIdpRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamIdpRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamIdpRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamIdpRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamIdpRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamIdpRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamIdpRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamIdpRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDomainName

`func (o *IamIdpRelationship) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamIdpRelationship) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamIdpRelationship) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamIdpRelationship) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamIdpRelationship) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamIdpRelationship) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamIdpRelationship) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamIdpRelationship) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamIdpRelationship) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamIdpRelationship) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamIdpRelationship) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamIdpRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamIdpRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamIdpRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamIdpRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamIdpRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IamIdpRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamIdpRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamIdpRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamIdpRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *IamIdpRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamIdpRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamIdpRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamIdpRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLdapPolicy

`func (o *IamIdpRelationship) GetLdapPolicy() IamLdapPolicyRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamIdpRelationship) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamIdpRelationship) SetLdapPolicy(v IamLdapPolicyRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamIdpRelationship) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.

### GetSystem

`func (o *IamIdpRelationship) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamIdpRelationship) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamIdpRelationship) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamIdpRelationship) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUserPreferences

`func (o *IamIdpRelationship) GetUserPreferences() []IamUserPreferenceRelationship`

GetUserPreferences returns the UserPreferences field if non-nil, zero value otherwise.

### GetUserPreferencesOk

`func (o *IamIdpRelationship) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool)`

GetUserPreferencesOk returns a tuple with the UserPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPreferences

`func (o *IamIdpRelationship) SetUserPreferences(v []IamUserPreferenceRelationship)`

SetUserPreferences sets UserPreferences field to given value.

### HasUserPreferences

`func (o *IamIdpRelationship) HasUserPreferences() bool`

HasUserPreferences returns a boolean if a field has been set.

### SetUserPreferencesNil

`func (o *IamIdpRelationship) SetUserPreferencesNil(b bool)`

 SetUserPreferencesNil sets the value for UserPreferences to be an explicit nil

### UnsetUserPreferences
`func (o *IamIdpRelationship) UnsetUserPreferences()`

UnsetUserPreferences ensures that no value is present for UserPreferences, not even an explicit nil
### GetUsergroups

`func (o *IamIdpRelationship) GetUsergroups() []IamUserGroupRelationship`

GetUsergroups returns the Usergroups field if non-nil, zero value otherwise.

### GetUsergroupsOk

`func (o *IamIdpRelationship) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUsergroupsOk returns a tuple with the Usergroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroups

`func (o *IamIdpRelationship) SetUsergroups(v []IamUserGroupRelationship)`

SetUsergroups sets Usergroups field to given value.

### HasUsergroups

`func (o *IamIdpRelationship) HasUsergroups() bool`

HasUsergroups returns a boolean if a field has been set.

### SetUsergroupsNil

`func (o *IamIdpRelationship) SetUsergroupsNil(b bool)`

 SetUsergroupsNil sets the value for Usergroups to be an explicit nil

### UnsetUsergroups
`func (o *IamIdpRelationship) UnsetUsergroups()`

UnsetUsergroups ensures that no value is present for Usergroups, not even an explicit nil
### GetUsers

`func (o *IamIdpRelationship) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamIdpRelationship) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamIdpRelationship) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamIdpRelationship) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamIdpRelationship) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamIdpRelationship) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


