# IamUserRelationship

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
**ClientIpAddress** | Pointer to **string** | IP address from which the user last logged in to Intersight. | [optional] [readonly] 
**Email** | Pointer to **string** | Email of the user. Users are added to Intersight using the email configured in the IdP. | [optional] 
**FirstName** | Pointer to **string** | First name of the user. This field is populated from the IdP attributes received after authentication. | [optional] [readonly] 
**LastLoginTime** | Pointer to [**time.Time**](time.Time.md) | Last successful login time for user. | [optional] [readonly] 
**LastName** | Pointer to **string** | Last name of the user. This field is populated from the IdP attributes received after authentication. | [optional] [readonly] 
**Name** | Pointer to **string** | Name as configured in the IdP. | [optional] [readonly] 
**UserIdOrEmail** | Pointer to **string** | UserID or email as configured in the IdP. | [optional] 
**UserType** | Pointer to **string** | Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic. | [optional] [readonly] 
**UserUniqueIdentifier** | Pointer to **string** | Unique id of the user used by the identity provider to store the user. | [optional] [readonly] 
**ApiKeys** | Pointer to [**[]IamApiKeyRelationship**](iam.ApiKey.Relationship.md) | An array of relationships to iamApiKey resources. | [optional] [readonly] 
**AppRegistrations** | Pointer to [**[]IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) | An array of relationships to iamAppRegistration resources. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](iam.Idp.Relationship.md) |  | [optional] 
**Idpreference** | Pointer to [**IamIdpReferenceRelationship**](iam.IdpReference.Relationship.md) |  | [optional] 
**LocalUserPassword** | Pointer to [**IamLocalUserPasswordRelationship**](iam.LocalUserPassword.Relationship.md) |  | [optional] 
**OauthTokens** | Pointer to [**[]IamOAuthTokenRelationship**](iam.OAuthToken.Relationship.md) | An array of relationships to iamOAuthToken resources. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](iam.Permission.Relationship.md) | An array of relationships to iamPermission resources. | [optional] 
**Sessions** | Pointer to [**[]IamSessionRelationship**](iam.Session.Relationship.md) | An array of relationships to iamSession resources. | [optional] [readonly] 

## Methods

### NewIamUserRelationship

`func NewIamUserRelationship(classId string, objectType string, ) *IamUserRelationship`

NewIamUserRelationship instantiates a new IamUserRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserRelationshipWithDefaults

`func NewIamUserRelationshipWithDefaults() *IamUserRelationship`

NewIamUserRelationshipWithDefaults instantiates a new IamUserRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamUserRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamUserRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamUserRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamUserRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamUserRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamUserRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamUserRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamUserRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamUserRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamUserRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamUserRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamUserRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamUserRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamUserRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamUserRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamUserRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamUserRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamUserRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamUserRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamUserRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamUserRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamUserRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamUserRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamUserRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamUserRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamUserRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamUserRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamUserRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamUserRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamUserRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamUserRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamUserRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamUserRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamUserRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamUserRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamUserRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamUserRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamUserRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamUserRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamUserRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamUserRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamUserRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamUserRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamUserRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamUserRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamUserRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamUserRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamUserRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamUserRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamUserRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamUserRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamUserRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamUserRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamUserRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamUserRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamUserRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamUserRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamUserRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamUserRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamUserRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamUserRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamUserRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamUserRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamUserRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetClientIpAddress

`func (o *IamUserRelationship) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamUserRelationship) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamUserRelationship) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamUserRelationship) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetEmail

`func (o *IamUserRelationship) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamUserRelationship) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamUserRelationship) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamUserRelationship) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *IamUserRelationship) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IamUserRelationship) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IamUserRelationship) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IamUserRelationship) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamUserRelationship) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamUserRelationship) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamUserRelationship) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamUserRelationship) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastName

`func (o *IamUserRelationship) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IamUserRelationship) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IamUserRelationship) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IamUserRelationship) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetName

`func (o *IamUserRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUserRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUserRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamUserRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *IamUserRelationship) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *IamUserRelationship) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *IamUserRelationship) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *IamUserRelationship) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetUserType

`func (o *IamUserRelationship) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *IamUserRelationship) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *IamUserRelationship) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *IamUserRelationship) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserUniqueIdentifier

`func (o *IamUserRelationship) GetUserUniqueIdentifier() string`

GetUserUniqueIdentifier returns the UserUniqueIdentifier field if non-nil, zero value otherwise.

### GetUserUniqueIdentifierOk

`func (o *IamUserRelationship) GetUserUniqueIdentifierOk() (*string, bool)`

GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUniqueIdentifier

`func (o *IamUserRelationship) SetUserUniqueIdentifier(v string)`

SetUserUniqueIdentifier sets UserUniqueIdentifier field to given value.

### HasUserUniqueIdentifier

`func (o *IamUserRelationship) HasUserUniqueIdentifier() bool`

HasUserUniqueIdentifier returns a boolean if a field has been set.

### GetApiKeys

`func (o *IamUserRelationship) GetApiKeys() []IamApiKeyRelationship`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *IamUserRelationship) GetApiKeysOk() (*[]IamApiKeyRelationship, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *IamUserRelationship) SetApiKeys(v []IamApiKeyRelationship)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *IamUserRelationship) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### SetApiKeysNil

`func (o *IamUserRelationship) SetApiKeysNil(b bool)`

 SetApiKeysNil sets the value for ApiKeys to be an explicit nil

### UnsetApiKeys
`func (o *IamUserRelationship) UnsetApiKeys()`

UnsetApiKeys ensures that no value is present for ApiKeys, not even an explicit nil
### GetAppRegistrations

`func (o *IamUserRelationship) GetAppRegistrations() []IamAppRegistrationRelationship`

GetAppRegistrations returns the AppRegistrations field if non-nil, zero value otherwise.

### GetAppRegistrationsOk

`func (o *IamUserRelationship) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool)`

GetAppRegistrationsOk returns a tuple with the AppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistrations

`func (o *IamUserRelationship) SetAppRegistrations(v []IamAppRegistrationRelationship)`

SetAppRegistrations sets AppRegistrations field to given value.

### HasAppRegistrations

`func (o *IamUserRelationship) HasAppRegistrations() bool`

HasAppRegistrations returns a boolean if a field has been set.

### SetAppRegistrationsNil

`func (o *IamUserRelationship) SetAppRegistrationsNil(b bool)`

 SetAppRegistrationsNil sets the value for AppRegistrations to be an explicit nil

### UnsetAppRegistrations
`func (o *IamUserRelationship) UnsetAppRegistrations()`

UnsetAppRegistrations ensures that no value is present for AppRegistrations, not even an explicit nil
### GetIdp

`func (o *IamUserRelationship) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamUserRelationship) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamUserRelationship) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamUserRelationship) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpreference

`func (o *IamUserRelationship) GetIdpreference() IamIdpReferenceRelationship`

GetIdpreference returns the Idpreference field if non-nil, zero value otherwise.

### GetIdpreferenceOk

`func (o *IamUserRelationship) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpreferenceOk returns a tuple with the Idpreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreference

`func (o *IamUserRelationship) SetIdpreference(v IamIdpReferenceRelationship)`

SetIdpreference sets Idpreference field to given value.

### HasIdpreference

`func (o *IamUserRelationship) HasIdpreference() bool`

HasIdpreference returns a boolean if a field has been set.

### GetLocalUserPassword

`func (o *IamUserRelationship) GetLocalUserPassword() IamLocalUserPasswordRelationship`

GetLocalUserPassword returns the LocalUserPassword field if non-nil, zero value otherwise.

### GetLocalUserPasswordOk

`func (o *IamUserRelationship) GetLocalUserPasswordOk() (*IamLocalUserPasswordRelationship, bool)`

GetLocalUserPasswordOk returns a tuple with the LocalUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserPassword

`func (o *IamUserRelationship) SetLocalUserPassword(v IamLocalUserPasswordRelationship)`

SetLocalUserPassword sets LocalUserPassword field to given value.

### HasLocalUserPassword

`func (o *IamUserRelationship) HasLocalUserPassword() bool`

HasLocalUserPassword returns a boolean if a field has been set.

### GetOauthTokens

`func (o *IamUserRelationship) GetOauthTokens() []IamOAuthTokenRelationship`

GetOauthTokens returns the OauthTokens field if non-nil, zero value otherwise.

### GetOauthTokensOk

`func (o *IamUserRelationship) GetOauthTokensOk() (*[]IamOAuthTokenRelationship, bool)`

GetOauthTokensOk returns a tuple with the OauthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokens

`func (o *IamUserRelationship) SetOauthTokens(v []IamOAuthTokenRelationship)`

SetOauthTokens sets OauthTokens field to given value.

### HasOauthTokens

`func (o *IamUserRelationship) HasOauthTokens() bool`

HasOauthTokens returns a boolean if a field has been set.

### SetOauthTokensNil

`func (o *IamUserRelationship) SetOauthTokensNil(b bool)`

 SetOauthTokensNil sets the value for OauthTokens to be an explicit nil

### UnsetOauthTokens
`func (o *IamUserRelationship) UnsetOauthTokens()`

UnsetOauthTokens ensures that no value is present for OauthTokens, not even an explicit nil
### GetPermissions

`func (o *IamUserRelationship) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamUserRelationship) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamUserRelationship) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamUserRelationship) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamUserRelationship) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamUserRelationship) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSessions

`func (o *IamUserRelationship) GetSessions() []IamSessionRelationship`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *IamUserRelationship) GetSessionsOk() (*[]IamSessionRelationship, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *IamUserRelationship) SetSessions(v []IamSessionRelationship)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *IamUserRelationship) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *IamUserRelationship) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *IamUserRelationship) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


