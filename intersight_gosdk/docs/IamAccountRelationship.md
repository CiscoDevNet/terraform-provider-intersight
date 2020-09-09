# IamAccountRelationship

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
**Name** | Pointer to **string** | Name of the Intersight account. By default, name is same as the MoID of the account. | [optional] 
**Status** | Pointer to **string** | Status of the account. To activate the Intersight account, claim a device to the account. | [optional] [readonly] 
**Var2LicenseReservationOp** | Pointer to [**LicenseLicenseReservationOpRelationship**](license.LicenseReservationOp.Relationship.md) |  | [optional] 
**AppRegistrations** | Pointer to [**[]IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) | An array of relationships to iamAppRegistration resources. | [optional] [readonly] 
**DomainGroups** | Pointer to [**[]IamDomainGroupRelationship**](iam.DomainGroup.Relationship.md) | An array of relationships to iamDomainGroup resources. | [optional] [readonly] 
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](iam.EndPointRole.Relationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Idpreferences** | Pointer to [**[]IamIdpReferenceRelationship**](iam.IdpReference.Relationship.md) | An array of relationships to iamIdpReference resources. | [optional] [readonly] 
**Idps** | Pointer to [**[]IamIdpRelationship**](iam.Idp.Relationship.md) | An array of relationships to iamIdp resources. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](iam.Permission.Relationship.md) | An array of relationships to iamPermission resources. | [optional] [readonly] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](iam.PrivilegeSet.Relationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](iam.Privilege.Relationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**ResourceLimits** | Pointer to [**IamResourceLimitsRelationship**](iam.ResourceLimits.Relationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](iam.Role.Relationship.md) | An array of relationships to iamRole resources. | [optional] [readonly] 
**SecurityHolder** | Pointer to [**IamSecurityHolderRelationship**](iam.SecurityHolder.Relationship.md) |  | [optional] 
**SessionLimits** | Pointer to [**IamSessionLimitsRelationship**](iam.SessionLimits.Relationship.md) |  | [optional] 

## Methods

### NewIamAccountRelationship

`func NewIamAccountRelationship(classId string, objectType string, ) *IamAccountRelationship`

NewIamAccountRelationship instantiates a new IamAccountRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountRelationshipWithDefaults

`func NewIamAccountRelationshipWithDefaults() *IamAccountRelationship`

NewIamAccountRelationshipWithDefaults instantiates a new IamAccountRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamAccountRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamAccountRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamAccountRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamAccountRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamAccountRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAccountRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAccountRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamAccountRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamAccountRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamAccountRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamAccountRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamAccountRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamAccountRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamAccountRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamAccountRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamAccountRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamAccountRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamAccountRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamAccountRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamAccountRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamAccountRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamAccountRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamAccountRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamAccountRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAccountRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAccountRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamAccountRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamAccountRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamAccountRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamAccountRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamAccountRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamAccountRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamAccountRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamAccountRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamAccountRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamAccountRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamAccountRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamAccountRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamAccountRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamAccountRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamAccountRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamAccountRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamAccountRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamAccountRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamAccountRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamAccountRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamAccountRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamAccountRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamAccountRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamAccountRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamAccountRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamAccountRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamAccountRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamAccountRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamAccountRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamAccountRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamAccountRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamAccountRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamAccountRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamAccountRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamAccountRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamAccountRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamAccountRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamAccountRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetName

`func (o *IamAccountRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamAccountRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamAccountRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamAccountRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IamAccountRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamAccountRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamAccountRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamAccountRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVar2LicenseReservationOp

`func (o *IamAccountRelationship) GetVar2LicenseReservationOp() LicenseLicenseReservationOpRelationship`

GetVar2LicenseReservationOp returns the Var2LicenseReservationOp field if non-nil, zero value otherwise.

### GetVar2LicenseReservationOpOk

`func (o *IamAccountRelationship) GetVar2LicenseReservationOpOk() (*LicenseLicenseReservationOpRelationship, bool)`

GetVar2LicenseReservationOpOk returns a tuple with the Var2LicenseReservationOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2LicenseReservationOp

`func (o *IamAccountRelationship) SetVar2LicenseReservationOp(v LicenseLicenseReservationOpRelationship)`

SetVar2LicenseReservationOp sets Var2LicenseReservationOp field to given value.

### HasVar2LicenseReservationOp

`func (o *IamAccountRelationship) HasVar2LicenseReservationOp() bool`

HasVar2LicenseReservationOp returns a boolean if a field has been set.

### GetAppRegistrations

`func (o *IamAccountRelationship) GetAppRegistrations() []IamAppRegistrationRelationship`

GetAppRegistrations returns the AppRegistrations field if non-nil, zero value otherwise.

### GetAppRegistrationsOk

`func (o *IamAccountRelationship) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool)`

GetAppRegistrationsOk returns a tuple with the AppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistrations

`func (o *IamAccountRelationship) SetAppRegistrations(v []IamAppRegistrationRelationship)`

SetAppRegistrations sets AppRegistrations field to given value.

### HasAppRegistrations

`func (o *IamAccountRelationship) HasAppRegistrations() bool`

HasAppRegistrations returns a boolean if a field has been set.

### SetAppRegistrationsNil

`func (o *IamAccountRelationship) SetAppRegistrationsNil(b bool)`

 SetAppRegistrationsNil sets the value for AppRegistrations to be an explicit nil

### UnsetAppRegistrations
`func (o *IamAccountRelationship) UnsetAppRegistrations()`

UnsetAppRegistrations ensures that no value is present for AppRegistrations, not even an explicit nil
### GetDomainGroups

`func (o *IamAccountRelationship) GetDomainGroups() []IamDomainGroupRelationship`

GetDomainGroups returns the DomainGroups field if non-nil, zero value otherwise.

### GetDomainGroupsOk

`func (o *IamAccountRelationship) GetDomainGroupsOk() (*[]IamDomainGroupRelationship, bool)`

GetDomainGroupsOk returns a tuple with the DomainGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroups

`func (o *IamAccountRelationship) SetDomainGroups(v []IamDomainGroupRelationship)`

SetDomainGroups sets DomainGroups field to given value.

### HasDomainGroups

`func (o *IamAccountRelationship) HasDomainGroups() bool`

HasDomainGroups returns a boolean if a field has been set.

### SetDomainGroupsNil

`func (o *IamAccountRelationship) SetDomainGroupsNil(b bool)`

 SetDomainGroupsNil sets the value for DomainGroups to be an explicit nil

### UnsetDomainGroups
`func (o *IamAccountRelationship) UnsetDomainGroups()`

UnsetDomainGroups ensures that no value is present for DomainGroups, not even an explicit nil
### GetEndPointRoles

`func (o *IamAccountRelationship) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamAccountRelationship) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamAccountRelationship) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamAccountRelationship) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### SetEndPointRolesNil

`func (o *IamAccountRelationship) SetEndPointRolesNil(b bool)`

 SetEndPointRolesNil sets the value for EndPointRoles to be an explicit nil

### UnsetEndPointRoles
`func (o *IamAccountRelationship) UnsetEndPointRoles()`

UnsetEndPointRoles ensures that no value is present for EndPointRoles, not even an explicit nil
### GetIdpreferences

`func (o *IamAccountRelationship) GetIdpreferences() []IamIdpReferenceRelationship`

GetIdpreferences returns the Idpreferences field if non-nil, zero value otherwise.

### GetIdpreferencesOk

`func (o *IamAccountRelationship) GetIdpreferencesOk() (*[]IamIdpReferenceRelationship, bool)`

GetIdpreferencesOk returns a tuple with the Idpreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreferences

`func (o *IamAccountRelationship) SetIdpreferences(v []IamIdpReferenceRelationship)`

SetIdpreferences sets Idpreferences field to given value.

### HasIdpreferences

`func (o *IamAccountRelationship) HasIdpreferences() bool`

HasIdpreferences returns a boolean if a field has been set.

### SetIdpreferencesNil

`func (o *IamAccountRelationship) SetIdpreferencesNil(b bool)`

 SetIdpreferencesNil sets the value for Idpreferences to be an explicit nil

### UnsetIdpreferences
`func (o *IamAccountRelationship) UnsetIdpreferences()`

UnsetIdpreferences ensures that no value is present for Idpreferences, not even an explicit nil
### GetIdps

`func (o *IamAccountRelationship) GetIdps() []IamIdpRelationship`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *IamAccountRelationship) GetIdpsOk() (*[]IamIdpRelationship, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *IamAccountRelationship) SetIdps(v []IamIdpRelationship)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *IamAccountRelationship) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### SetIdpsNil

`func (o *IamAccountRelationship) SetIdpsNil(b bool)`

 SetIdpsNil sets the value for Idps to be an explicit nil

### UnsetIdps
`func (o *IamAccountRelationship) UnsetIdps()`

UnsetIdps ensures that no value is present for Idps, not even an explicit nil
### GetPermissions

`func (o *IamAccountRelationship) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamAccountRelationship) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamAccountRelationship) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamAccountRelationship) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamAccountRelationship) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamAccountRelationship) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPrivilegeSets

`func (o *IamAccountRelationship) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamAccountRelationship) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamAccountRelationship) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamAccountRelationship) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamAccountRelationship) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamAccountRelationship) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetPrivileges

`func (o *IamAccountRelationship) GetPrivileges() []IamPrivilegeRelationship`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *IamAccountRelationship) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *IamAccountRelationship) SetPrivileges(v []IamPrivilegeRelationship)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *IamAccountRelationship) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### SetPrivilegesNil

`func (o *IamAccountRelationship) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *IamAccountRelationship) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil
### GetResourceLimits

`func (o *IamAccountRelationship) GetResourceLimits() IamResourceLimitsRelationship`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *IamAccountRelationship) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *IamAccountRelationship) SetResourceLimits(v IamResourceLimitsRelationship)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *IamAccountRelationship) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.

### GetRoles

`func (o *IamAccountRelationship) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamAccountRelationship) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamAccountRelationship) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamAccountRelationship) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamAccountRelationship) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamAccountRelationship) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSecurityHolder

`func (o *IamAccountRelationship) GetSecurityHolder() IamSecurityHolderRelationship`

GetSecurityHolder returns the SecurityHolder field if non-nil, zero value otherwise.

### GetSecurityHolderOk

`func (o *IamAccountRelationship) GetSecurityHolderOk() (*IamSecurityHolderRelationship, bool)`

GetSecurityHolderOk returns a tuple with the SecurityHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityHolder

`func (o *IamAccountRelationship) SetSecurityHolder(v IamSecurityHolderRelationship)`

SetSecurityHolder sets SecurityHolder field to given value.

### HasSecurityHolder

`func (o *IamAccountRelationship) HasSecurityHolder() bool`

HasSecurityHolder returns a boolean if a field has been set.

### GetSessionLimits

`func (o *IamAccountRelationship) GetSessionLimits() IamSessionLimitsRelationship`

GetSessionLimits returns the SessionLimits field if non-nil, zero value otherwise.

### GetSessionLimitsOk

`func (o *IamAccountRelationship) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool)`

GetSessionLimitsOk returns a tuple with the SessionLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimits

`func (o *IamAccountRelationship) SetSessionLimits(v IamSessionLimitsRelationship)`

SetSessionLimits sets SessionLimits field to given value.

### HasSessionLimits

`func (o *IamAccountRelationship) HasSessionLimits() bool`

HasSessionLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


