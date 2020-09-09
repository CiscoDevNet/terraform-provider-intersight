# RecoveryBackupConfigPolicyRelationship

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
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**Name** | Pointer to **string** | Name of the concrete policy. | [optional] 
**FileNamePrefix** | Pointer to **string** | The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**LocationType** | Pointer to **string** | Specifies whether the backup will be stored locally or remotely. * &#x60;Network Share&#x60; - The backup is stored remotely on a separate server. * &#x60;Local Storage&#x60; - The backup is stored locally on the endpoint. | [optional] [default to "Network Share"]
**Password** | Pointer to **string** | Password of Backup server. | [optional] 
**Path** | Pointer to **string** | The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. | [optional] 
**Protocol** | Pointer to **string** | Protocol for transferring the backup image to the network share location. * &#x60;SCP&#x60; - Secure Copy Protocol (SCP) to access the file server. * &#x60;SFTP&#x60; - SSH File Transfer Protocol (SFTP) to access file server. * &#x60;FTP&#x60; - File Transfer Protocol (FTP) to access file server. | [optional] [default to "SCP"]
**RetentionCount** | Pointer to **int64** | Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. | [optional] 
**UserName** | Pointer to **string** | Backup server user name. | [optional] 
**BackupProfiles** | Pointer to [**[]RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) | An array of relationships to recoveryBackupProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewRecoveryBackupConfigPolicyRelationship

`func NewRecoveryBackupConfigPolicyRelationship(classId string, objectType string, ) *RecoveryBackupConfigPolicyRelationship`

NewRecoveryBackupConfigPolicyRelationship instantiates a new RecoveryBackupConfigPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupConfigPolicyRelationshipWithDefaults

`func NewRecoveryBackupConfigPolicyRelationshipWithDefaults() *RecoveryBackupConfigPolicyRelationship`

NewRecoveryBackupConfigPolicyRelationshipWithDefaults instantiates a new RecoveryBackupConfigPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *RecoveryBackupConfigPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *RecoveryBackupConfigPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *RecoveryBackupConfigPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *RecoveryBackupConfigPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryBackupConfigPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *RecoveryBackupConfigPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RecoveryBackupConfigPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RecoveryBackupConfigPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *RecoveryBackupConfigPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *RecoveryBackupConfigPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *RecoveryBackupConfigPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *RecoveryBackupConfigPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *RecoveryBackupConfigPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *RecoveryBackupConfigPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *RecoveryBackupConfigPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *RecoveryBackupConfigPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *RecoveryBackupConfigPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *RecoveryBackupConfigPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryBackupConfigPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *RecoveryBackupConfigPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *RecoveryBackupConfigPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *RecoveryBackupConfigPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *RecoveryBackupConfigPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *RecoveryBackupConfigPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *RecoveryBackupConfigPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *RecoveryBackupConfigPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecoveryBackupConfigPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecoveryBackupConfigPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *RecoveryBackupConfigPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *RecoveryBackupConfigPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *RecoveryBackupConfigPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *RecoveryBackupConfigPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *RecoveryBackupConfigPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *RecoveryBackupConfigPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *RecoveryBackupConfigPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *RecoveryBackupConfigPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *RecoveryBackupConfigPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RecoveryBackupConfigPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RecoveryBackupConfigPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *RecoveryBackupConfigPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *RecoveryBackupConfigPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *RecoveryBackupConfigPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *RecoveryBackupConfigPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *RecoveryBackupConfigPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *RecoveryBackupConfigPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *RecoveryBackupConfigPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *RecoveryBackupConfigPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *RecoveryBackupConfigPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *RecoveryBackupConfigPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *RecoveryBackupConfigPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecoveryBackupConfigPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecoveryBackupConfigPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *RecoveryBackupConfigPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryBackupConfigPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecoveryBackupConfigPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileNamePrefix

`func (o *RecoveryBackupConfigPolicyRelationship) GetFileNamePrefix() string`

GetFileNamePrefix returns the FileNamePrefix field if non-nil, zero value otherwise.

### GetFileNamePrefixOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetFileNamePrefixOk() (*string, bool)`

GetFileNamePrefixOk returns a tuple with the FileNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePrefix

`func (o *RecoveryBackupConfigPolicyRelationship) SetFileNamePrefix(v string)`

SetFileNamePrefix sets FileNamePrefix field to given value.

### HasFileNamePrefix

`func (o *RecoveryBackupConfigPolicyRelationship) HasFileNamePrefix() bool`

HasFileNamePrefix returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *RecoveryBackupConfigPolicyRelationship) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *RecoveryBackupConfigPolicyRelationship) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *RecoveryBackupConfigPolicyRelationship) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocationType

`func (o *RecoveryBackupConfigPolicyRelationship) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *RecoveryBackupConfigPolicyRelationship) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *RecoveryBackupConfigPolicyRelationship) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetPassword

`func (o *RecoveryBackupConfigPolicyRelationship) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RecoveryBackupConfigPolicyRelationship) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RecoveryBackupConfigPolicyRelationship) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *RecoveryBackupConfigPolicyRelationship) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RecoveryBackupConfigPolicyRelationship) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RecoveryBackupConfigPolicyRelationship) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *RecoveryBackupConfigPolicyRelationship) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecoveryBackupConfigPolicyRelationship) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RecoveryBackupConfigPolicyRelationship) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRetentionCount

`func (o *RecoveryBackupConfigPolicyRelationship) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *RecoveryBackupConfigPolicyRelationship) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *RecoveryBackupConfigPolicyRelationship) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetUserName

`func (o *RecoveryBackupConfigPolicyRelationship) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RecoveryBackupConfigPolicyRelationship) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RecoveryBackupConfigPolicyRelationship) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetBackupProfiles

`func (o *RecoveryBackupConfigPolicyRelationship) GetBackupProfiles() []RecoveryBackupProfileRelationship`

GetBackupProfiles returns the BackupProfiles field if non-nil, zero value otherwise.

### GetBackupProfilesOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool)`

GetBackupProfilesOk returns a tuple with the BackupProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfiles

`func (o *RecoveryBackupConfigPolicyRelationship) SetBackupProfiles(v []RecoveryBackupProfileRelationship)`

SetBackupProfiles sets BackupProfiles field to given value.

### HasBackupProfiles

`func (o *RecoveryBackupConfigPolicyRelationship) HasBackupProfiles() bool`

HasBackupProfiles returns a boolean if a field has been set.

### SetBackupProfilesNil

`func (o *RecoveryBackupConfigPolicyRelationship) SetBackupProfilesNil(b bool)`

 SetBackupProfilesNil sets the value for BackupProfiles to be an explicit nil

### UnsetBackupProfiles
`func (o *RecoveryBackupConfigPolicyRelationship) UnsetBackupProfiles()`

UnsetBackupProfiles ensures that no value is present for BackupProfiles, not even an explicit nil
### GetOrganization

`func (o *RecoveryBackupConfigPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RecoveryBackupConfigPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RecoveryBackupConfigPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RecoveryBackupConfigPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


