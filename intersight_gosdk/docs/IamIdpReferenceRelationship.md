# IamIdpReferenceRelationship

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

### NewIamIdpReferenceRelationship

`func NewIamIdpReferenceRelationship(classId string, objectType string, ) *IamIdpReferenceRelationship`

NewIamIdpReferenceRelationship instantiates a new IamIdpReferenceRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIdpReferenceRelationshipWithDefaults

`func NewIamIdpReferenceRelationshipWithDefaults() *IamIdpReferenceRelationship`

NewIamIdpReferenceRelationshipWithDefaults instantiates a new IamIdpReferenceRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamIdpReferenceRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamIdpReferenceRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamIdpReferenceRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamIdpReferenceRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamIdpReferenceRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIdpReferenceRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIdpReferenceRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamIdpReferenceRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamIdpReferenceRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamIdpReferenceRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamIdpReferenceRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamIdpReferenceRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamIdpReferenceRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamIdpReferenceRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamIdpReferenceRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamIdpReferenceRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamIdpReferenceRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamIdpReferenceRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamIdpReferenceRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamIdpReferenceRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamIdpReferenceRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamIdpReferenceRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamIdpReferenceRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamIdpReferenceRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIdpReferenceRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIdpReferenceRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamIdpReferenceRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamIdpReferenceRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamIdpReferenceRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamIdpReferenceRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamIdpReferenceRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamIdpReferenceRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamIdpReferenceRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamIdpReferenceRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamIdpReferenceRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamIdpReferenceRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamIdpReferenceRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamIdpReferenceRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamIdpReferenceRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamIdpReferenceRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamIdpReferenceRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamIdpReferenceRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamIdpReferenceRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamIdpReferenceRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamIdpReferenceRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamIdpReferenceRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamIdpReferenceRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamIdpReferenceRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamIdpReferenceRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamIdpReferenceRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamIdpReferenceRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamIdpReferenceRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamIdpReferenceRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamIdpReferenceRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamIdpReferenceRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamIdpReferenceRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamIdpReferenceRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamIdpReferenceRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamIdpReferenceRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamIdpReferenceRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamIdpReferenceRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamIdpReferenceRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamIdpReferenceRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamIdpReferenceRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDomainName

`func (o *IamIdpReferenceRelationship) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamIdpReferenceRelationship) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamIdpReferenceRelationship) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamIdpReferenceRelationship) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamIdpReferenceRelationship) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamIdpReferenceRelationship) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamIdpReferenceRelationship) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamIdpReferenceRelationship) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetMultiFactorAuthentication

`func (o *IamIdpReferenceRelationship) GetMultiFactorAuthentication() bool`

GetMultiFactorAuthentication returns the MultiFactorAuthentication field if non-nil, zero value otherwise.

### GetMultiFactorAuthenticationOk

`func (o *IamIdpReferenceRelationship) GetMultiFactorAuthenticationOk() (*bool, bool)`

GetMultiFactorAuthenticationOk returns a tuple with the MultiFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthentication

`func (o *IamIdpReferenceRelationship) SetMultiFactorAuthentication(v bool)`

SetMultiFactorAuthentication sets MultiFactorAuthentication field to given value.

### HasMultiFactorAuthentication

`func (o *IamIdpReferenceRelationship) HasMultiFactorAuthentication() bool`

HasMultiFactorAuthentication returns a boolean if a field has been set.

### GetName

`func (o *IamIdpReferenceRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamIdpReferenceRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamIdpReferenceRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamIdpReferenceRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *IamIdpReferenceRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamIdpReferenceRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamIdpReferenceRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamIdpReferenceRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIdp

`func (o *IamIdpReferenceRelationship) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamIdpReferenceRelationship) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamIdpReferenceRelationship) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamIdpReferenceRelationship) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetUserPreferences

`func (o *IamIdpReferenceRelationship) GetUserPreferences() []IamUserPreferenceRelationship`

GetUserPreferences returns the UserPreferences field if non-nil, zero value otherwise.

### GetUserPreferencesOk

`func (o *IamIdpReferenceRelationship) GetUserPreferencesOk() (*[]IamUserPreferenceRelationship, bool)`

GetUserPreferencesOk returns a tuple with the UserPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPreferences

`func (o *IamIdpReferenceRelationship) SetUserPreferences(v []IamUserPreferenceRelationship)`

SetUserPreferences sets UserPreferences field to given value.

### HasUserPreferences

`func (o *IamIdpReferenceRelationship) HasUserPreferences() bool`

HasUserPreferences returns a boolean if a field has been set.

### SetUserPreferencesNil

`func (o *IamIdpReferenceRelationship) SetUserPreferencesNil(b bool)`

 SetUserPreferencesNil sets the value for UserPreferences to be an explicit nil

### UnsetUserPreferences
`func (o *IamIdpReferenceRelationship) UnsetUserPreferences()`

UnsetUserPreferences ensures that no value is present for UserPreferences, not even an explicit nil
### GetUsergroups

`func (o *IamIdpReferenceRelationship) GetUsergroups() []IamUserGroupRelationship`

GetUsergroups returns the Usergroups field if non-nil, zero value otherwise.

### GetUsergroupsOk

`func (o *IamIdpReferenceRelationship) GetUsergroupsOk() (*[]IamUserGroupRelationship, bool)`

GetUsergroupsOk returns a tuple with the Usergroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsergroups

`func (o *IamIdpReferenceRelationship) SetUsergroups(v []IamUserGroupRelationship)`

SetUsergroups sets Usergroups field to given value.

### HasUsergroups

`func (o *IamIdpReferenceRelationship) HasUsergroups() bool`

HasUsergroups returns a boolean if a field has been set.

### SetUsergroupsNil

`func (o *IamIdpReferenceRelationship) SetUsergroupsNil(b bool)`

 SetUsergroupsNil sets the value for Usergroups to be an explicit nil

### UnsetUsergroups
`func (o *IamIdpReferenceRelationship) UnsetUsergroups()`

UnsetUsergroups ensures that no value is present for Usergroups, not even an explicit nil
### GetUsers

`func (o *IamIdpReferenceRelationship) GetUsers() []IamUserRelationship`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamIdpReferenceRelationship) GetUsersOk() (*[]IamUserRelationship, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamIdpReferenceRelationship) SetUsers(v []IamUserRelationship)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamIdpReferenceRelationship) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IamIdpReferenceRelationship) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IamIdpReferenceRelationship) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


