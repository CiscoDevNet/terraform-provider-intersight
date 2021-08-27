# IamAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Account"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Account"]
**Name** | Pointer to **string** | Name of the Intersight account. By default, name is same as the MoID of the account. | [optional] 
**Status** | Pointer to **string** | Status of the account. To activate the Intersight account, claim a device to the account. | [optional] [readonly] 
**AppRegistrations** | Pointer to [**[]IamAppRegistrationRelationship**](IamAppRegistrationRelationship.md) | An array of relationships to iamAppRegistration resources. | [optional] [readonly] 
**DomainGroups** | Pointer to [**[]IamDomainGroupRelationship**](IamDomainGroupRelationship.md) | An array of relationships to iamDomainGroup resources. | [optional] [readonly] 
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Idpreferences** | Pointer to [**[]IamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) | An array of relationships to iamIdpReference resources. | [optional] [readonly] 
**Idps** | Pointer to [**[]IamIdpRelationship**](IamIdpRelationship.md) | An array of relationships to iamIdp resources. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](IamPermissionRelationship.md) | An array of relationships to iamPermission resources. | [optional] [readonly] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](IamPrivilegeRelationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**ResourceLimits** | Pointer to [**IamResourceLimitsRelationship**](IamResourceLimitsRelationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] [readonly] 
**SecurityHolder** | Pointer to [**IamSecurityHolderRelationship**](IamSecurityHolderRelationship.md) |  | [optional] 
**SessionLimits** | Pointer to [**IamSessionLimitsRelationship**](IamSessionLimitsRelationship.md) |  | [optional] 

## Methods

### NewIamAccount

`func NewIamAccount(classId string, objectType string, ) *IamAccount`

NewIamAccount instantiates a new IamAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountWithDefaults

`func NewIamAccountWithDefaults() *IamAccount`

NewIamAccountWithDefaults instantiates a new IamAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAccount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAccount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAccount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAccount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAccount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAccount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *IamAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IamAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppRegistrations

`func (o *IamAccount) GetAppRegistrations() []IamAppRegistrationRelationship`

GetAppRegistrations returns the AppRegistrations field if non-nil, zero value otherwise.

### GetAppRegistrationsOk

`func (o *IamAccount) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool)`

GetAppRegistrationsOk returns a tuple with the AppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistrations

`func (o *IamAccount) SetAppRegistrations(v []IamAppRegistrationRelationship)`

SetAppRegistrations sets AppRegistrations field to given value.

### HasAppRegistrations

`func (o *IamAccount) HasAppRegistrations() bool`

HasAppRegistrations returns a boolean if a field has been set.

### SetAppRegistrationsNil

`func (o *IamAccount) SetAppRegistrationsNil(b bool)`

 SetAppRegistrationsNil sets the value for AppRegistrations to be an explicit nil

### UnsetAppRegistrations
`func (o *IamAccount) UnsetAppRegistrations()`

UnsetAppRegistrations ensures that no value is present for AppRegistrations, not even an explicit nil
### GetDomainGroups

`func (o *IamAccount) GetDomainGroups() []IamDomainGroupRelationship`

GetDomainGroups returns the DomainGroups field if non-nil, zero value otherwise.

### GetDomainGroupsOk

`func (o *IamAccount) GetDomainGroupsOk() (*[]IamDomainGroupRelationship, bool)`

GetDomainGroupsOk returns a tuple with the DomainGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroups

`func (o *IamAccount) SetDomainGroups(v []IamDomainGroupRelationship)`

SetDomainGroups sets DomainGroups field to given value.

### HasDomainGroups

`func (o *IamAccount) HasDomainGroups() bool`

HasDomainGroups returns a boolean if a field has been set.

### SetDomainGroupsNil

`func (o *IamAccount) SetDomainGroupsNil(b bool)`

 SetDomainGroupsNil sets the value for DomainGroups to be an explicit nil

### UnsetDomainGroups
`func (o *IamAccount) UnsetDomainGroups()`

UnsetDomainGroups ensures that no value is present for DomainGroups, not even an explicit nil
### GetEndPointRoles

`func (o *IamAccount) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamAccount) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamAccount) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamAccount) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### SetEndPointRolesNil

`func (o *IamAccount) SetEndPointRolesNil(b bool)`

 SetEndPointRolesNil sets the value for EndPointRoles to be an explicit nil

### UnsetEndPointRoles
`func (o *IamAccount) UnsetEndPointRoles()`

UnsetEndPointRoles ensures that no value is present for EndPointRoles, not even an explicit nil
### GetIdpreferences

`func (o *IamAccount) GetIdpreferences() []IamIdpReferenceRelationship`

GetIdpreferences returns the Idpreferences field if non-nil, zero value otherwise.

### GetIdpreferencesOk

`func (o *IamAccount) GetIdpreferencesOk() (*[]IamIdpReferenceRelationship, bool)`

GetIdpreferencesOk returns a tuple with the Idpreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreferences

`func (o *IamAccount) SetIdpreferences(v []IamIdpReferenceRelationship)`

SetIdpreferences sets Idpreferences field to given value.

### HasIdpreferences

`func (o *IamAccount) HasIdpreferences() bool`

HasIdpreferences returns a boolean if a field has been set.

### SetIdpreferencesNil

`func (o *IamAccount) SetIdpreferencesNil(b bool)`

 SetIdpreferencesNil sets the value for Idpreferences to be an explicit nil

### UnsetIdpreferences
`func (o *IamAccount) UnsetIdpreferences()`

UnsetIdpreferences ensures that no value is present for Idpreferences, not even an explicit nil
### GetIdps

`func (o *IamAccount) GetIdps() []IamIdpRelationship`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *IamAccount) GetIdpsOk() (*[]IamIdpRelationship, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *IamAccount) SetIdps(v []IamIdpRelationship)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *IamAccount) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### SetIdpsNil

`func (o *IamAccount) SetIdpsNil(b bool)`

 SetIdpsNil sets the value for Idps to be an explicit nil

### UnsetIdps
`func (o *IamAccount) UnsetIdps()`

UnsetIdps ensures that no value is present for Idps, not even an explicit nil
### GetPermissions

`func (o *IamAccount) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamAccount) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamAccount) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamAccount) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *IamAccount) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *IamAccount) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPrivilegeSets

`func (o *IamAccount) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamAccount) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamAccount) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamAccount) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamAccount) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamAccount) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetPrivileges

`func (o *IamAccount) GetPrivileges() []IamPrivilegeRelationship`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *IamAccount) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *IamAccount) SetPrivileges(v []IamPrivilegeRelationship)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *IamAccount) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### SetPrivilegesNil

`func (o *IamAccount) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *IamAccount) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil
### GetResourceLimits

`func (o *IamAccount) GetResourceLimits() IamResourceLimitsRelationship`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *IamAccount) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *IamAccount) SetResourceLimits(v IamResourceLimitsRelationship)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *IamAccount) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.

### GetRoles

`func (o *IamAccount) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamAccount) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamAccount) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamAccount) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamAccount) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSecurityHolder

`func (o *IamAccount) GetSecurityHolder() IamSecurityHolderRelationship`

GetSecurityHolder returns the SecurityHolder field if non-nil, zero value otherwise.

### GetSecurityHolderOk

`func (o *IamAccount) GetSecurityHolderOk() (*IamSecurityHolderRelationship, bool)`

GetSecurityHolderOk returns a tuple with the SecurityHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityHolder

`func (o *IamAccount) SetSecurityHolder(v IamSecurityHolderRelationship)`

SetSecurityHolder sets SecurityHolder field to given value.

### HasSecurityHolder

`func (o *IamAccount) HasSecurityHolder() bool`

HasSecurityHolder returns a boolean if a field has been set.

### GetSessionLimits

`func (o *IamAccount) GetSessionLimits() IamSessionLimitsRelationship`

GetSessionLimits returns the SessionLimits field if non-nil, zero value otherwise.

### GetSessionLimitsOk

`func (o *IamAccount) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool)`

GetSessionLimitsOk returns a tuple with the SessionLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimits

`func (o *IamAccount) SetSessionLimits(v IamSessionLimitsRelationship)`

SetSessionLimits sets SessionLimits field to given value.

### HasSessionLimits

`func (o *IamAccount) HasSessionLimits() bool`

HasSessionLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


