# IamOAuthTokenRelationship

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
**AccessExpirationTime** | Pointer to [**time.Time**](time.Time.md) | Expiration time for the JWT token to which it can be used for api calls. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The identifier of the registered application to which the token belongs. | [optional] 
**ClientIpAddress** | Pointer to **string** | The user agent IP address from which the auth token is launched. | [optional] [readonly] 
**ClientName** | Pointer to **string** | The name of the registered application to which the token belongs. | [optional] 
**ExpirationTime** | Pointer to [**time.Time**](time.Time.md) | Expiration time for the JWT token to which it can be refreshed. | [optional] [readonly] 
**LastLoginClient** | Pointer to **string** | The client address from which last login is initiated. | [optional] [readonly] 
**LastLoginTime** | Pointer to [**time.Time**](time.Time.md) | The last login time for user. | [optional] [readonly] 
**TokenId** | Pointer to **string** | Token identifier. Not the Access Token itself. | [optional] [readonly] 
**UserMeta** | Pointer to [**IamClientMeta**](iam.ClientMeta.md) |  | [optional] 
**AppRegistration** | Pointer to [**IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamOAuthTokenRelationship

`func NewIamOAuthTokenRelationship(classId string, objectType string, ) *IamOAuthTokenRelationship`

NewIamOAuthTokenRelationship instantiates a new IamOAuthTokenRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamOAuthTokenRelationshipWithDefaults

`func NewIamOAuthTokenRelationshipWithDefaults() *IamOAuthTokenRelationship`

NewIamOAuthTokenRelationshipWithDefaults instantiates a new IamOAuthTokenRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamOAuthTokenRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamOAuthTokenRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamOAuthTokenRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamOAuthTokenRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamOAuthTokenRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamOAuthTokenRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamOAuthTokenRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamOAuthTokenRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamOAuthTokenRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamOAuthTokenRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamOAuthTokenRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamOAuthTokenRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamOAuthTokenRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamOAuthTokenRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamOAuthTokenRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamOAuthTokenRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamOAuthTokenRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamOAuthTokenRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamOAuthTokenRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamOAuthTokenRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamOAuthTokenRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamOAuthTokenRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamOAuthTokenRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamOAuthTokenRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamOAuthTokenRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamOAuthTokenRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamOAuthTokenRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamOAuthTokenRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamOAuthTokenRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamOAuthTokenRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamOAuthTokenRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamOAuthTokenRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamOAuthTokenRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamOAuthTokenRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamOAuthTokenRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamOAuthTokenRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamOAuthTokenRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamOAuthTokenRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamOAuthTokenRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamOAuthTokenRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamOAuthTokenRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamOAuthTokenRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamOAuthTokenRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamOAuthTokenRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamOAuthTokenRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamOAuthTokenRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamOAuthTokenRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamOAuthTokenRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamOAuthTokenRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamOAuthTokenRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamOAuthTokenRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamOAuthTokenRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamOAuthTokenRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamOAuthTokenRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamOAuthTokenRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamOAuthTokenRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamOAuthTokenRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamOAuthTokenRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamOAuthTokenRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamOAuthTokenRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamOAuthTokenRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamOAuthTokenRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamOAuthTokenRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamOAuthTokenRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAccessExpirationTime

`func (o *IamOAuthTokenRelationship) GetAccessExpirationTime() time.Time`

GetAccessExpirationTime returns the AccessExpirationTime field if non-nil, zero value otherwise.

### GetAccessExpirationTimeOk

`func (o *IamOAuthTokenRelationship) GetAccessExpirationTimeOk() (*time.Time, bool)`

GetAccessExpirationTimeOk returns a tuple with the AccessExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpirationTime

`func (o *IamOAuthTokenRelationship) SetAccessExpirationTime(v time.Time)`

SetAccessExpirationTime sets AccessExpirationTime field to given value.

### HasAccessExpirationTime

`func (o *IamOAuthTokenRelationship) HasAccessExpirationTime() bool`

HasAccessExpirationTime returns a boolean if a field has been set.

### GetClientId

`func (o *IamOAuthTokenRelationship) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IamOAuthTokenRelationship) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IamOAuthTokenRelationship) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IamOAuthTokenRelationship) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIpAddress

`func (o *IamOAuthTokenRelationship) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *IamOAuthTokenRelationship) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *IamOAuthTokenRelationship) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *IamOAuthTokenRelationship) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetClientName

`func (o *IamOAuthTokenRelationship) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *IamOAuthTokenRelationship) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *IamOAuthTokenRelationship) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *IamOAuthTokenRelationship) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetExpirationTime

`func (o *IamOAuthTokenRelationship) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *IamOAuthTokenRelationship) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *IamOAuthTokenRelationship) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *IamOAuthTokenRelationship) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastLoginClient

`func (o *IamOAuthTokenRelationship) GetLastLoginClient() string`

GetLastLoginClient returns the LastLoginClient field if non-nil, zero value otherwise.

### GetLastLoginClientOk

`func (o *IamOAuthTokenRelationship) GetLastLoginClientOk() (*string, bool)`

GetLastLoginClientOk returns a tuple with the LastLoginClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginClient

`func (o *IamOAuthTokenRelationship) SetLastLoginClient(v string)`

SetLastLoginClient sets LastLoginClient field to given value.

### HasLastLoginClient

`func (o *IamOAuthTokenRelationship) HasLastLoginClient() bool`

HasLastLoginClient returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *IamOAuthTokenRelationship) GetLastLoginTime() time.Time`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *IamOAuthTokenRelationship) GetLastLoginTimeOk() (*time.Time, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *IamOAuthTokenRelationship) SetLastLoginTime(v time.Time)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *IamOAuthTokenRelationship) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetTokenId

`func (o *IamOAuthTokenRelationship) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *IamOAuthTokenRelationship) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *IamOAuthTokenRelationship) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *IamOAuthTokenRelationship) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUserMeta

`func (o *IamOAuthTokenRelationship) GetUserMeta() IamClientMeta`

GetUserMeta returns the UserMeta field if non-nil, zero value otherwise.

### GetUserMetaOk

`func (o *IamOAuthTokenRelationship) GetUserMetaOk() (*IamClientMeta, bool)`

GetUserMetaOk returns a tuple with the UserMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMeta

`func (o *IamOAuthTokenRelationship) SetUserMeta(v IamClientMeta)`

SetUserMeta sets UserMeta field to given value.

### HasUserMeta

`func (o *IamOAuthTokenRelationship) HasUserMeta() bool`

HasUserMeta returns a boolean if a field has been set.

### GetAppRegistration

`func (o *IamOAuthTokenRelationship) GetAppRegistration() IamAppRegistrationRelationship`

GetAppRegistration returns the AppRegistration field if non-nil, zero value otherwise.

### GetAppRegistrationOk

`func (o *IamOAuthTokenRelationship) GetAppRegistrationOk() (*IamAppRegistrationRelationship, bool)`

GetAppRegistrationOk returns a tuple with the AppRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistration

`func (o *IamOAuthTokenRelationship) SetAppRegistration(v IamAppRegistrationRelationship)`

SetAppRegistration sets AppRegistration field to given value.

### HasAppRegistration

`func (o *IamOAuthTokenRelationship) HasAppRegistration() bool`

HasAppRegistration returns a boolean if a field has been set.

### GetPermission

`func (o *IamOAuthTokenRelationship) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamOAuthTokenRelationship) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamOAuthTokenRelationship) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamOAuthTokenRelationship) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamOAuthTokenRelationship) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamOAuthTokenRelationship) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamOAuthTokenRelationship) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamOAuthTokenRelationship) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


