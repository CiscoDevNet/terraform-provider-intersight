# RecoveryConfigResultRelationship

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
**ConfigStage** | Pointer to **string** | The current running stage of the configuration or workflow. | [optional] 
**ConfigState** | Pointer to **string** | Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**ValidationState** | Pointer to **string** | Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**Var0OnDemandBackup** | Pointer to [**RecoveryOnDemandBackupRelationship**](recovery.OnDemandBackup.Relationship.md) |  | [optional] 
**BackupProfile** | Pointer to [**RecoveryBackupProfileRelationship**](recovery.BackupProfile.Relationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]RecoveryConfigResultEntryRelationship**](recovery.ConfigResultEntry.Relationship.md) | An array of relationships to recoveryConfigResultEntry resources. | [optional] 

## Methods

### NewRecoveryConfigResultRelationship

`func NewRecoveryConfigResultRelationship(classId string, objectType string, ) *RecoveryConfigResultRelationship`

NewRecoveryConfigResultRelationship instantiates a new RecoveryConfigResultRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryConfigResultRelationshipWithDefaults

`func NewRecoveryConfigResultRelationshipWithDefaults() *RecoveryConfigResultRelationship`

NewRecoveryConfigResultRelationshipWithDefaults instantiates a new RecoveryConfigResultRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *RecoveryConfigResultRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *RecoveryConfigResultRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *RecoveryConfigResultRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *RecoveryConfigResultRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *RecoveryConfigResultRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryConfigResultRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryConfigResultRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *RecoveryConfigResultRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RecoveryConfigResultRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RecoveryConfigResultRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RecoveryConfigResultRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *RecoveryConfigResultRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *RecoveryConfigResultRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *RecoveryConfigResultRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *RecoveryConfigResultRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *RecoveryConfigResultRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *RecoveryConfigResultRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *RecoveryConfigResultRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *RecoveryConfigResultRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *RecoveryConfigResultRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *RecoveryConfigResultRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *RecoveryConfigResultRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *RecoveryConfigResultRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *RecoveryConfigResultRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryConfigResultRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryConfigResultRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *RecoveryConfigResultRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *RecoveryConfigResultRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *RecoveryConfigResultRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *RecoveryConfigResultRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *RecoveryConfigResultRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *RecoveryConfigResultRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *RecoveryConfigResultRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *RecoveryConfigResultRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *RecoveryConfigResultRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecoveryConfigResultRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecoveryConfigResultRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecoveryConfigResultRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *RecoveryConfigResultRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *RecoveryConfigResultRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *RecoveryConfigResultRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *RecoveryConfigResultRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *RecoveryConfigResultRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *RecoveryConfigResultRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *RecoveryConfigResultRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *RecoveryConfigResultRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *RecoveryConfigResultRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *RecoveryConfigResultRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *RecoveryConfigResultRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RecoveryConfigResultRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RecoveryConfigResultRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RecoveryConfigResultRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *RecoveryConfigResultRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *RecoveryConfigResultRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *RecoveryConfigResultRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *RecoveryConfigResultRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *RecoveryConfigResultRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *RecoveryConfigResultRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *RecoveryConfigResultRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *RecoveryConfigResultRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *RecoveryConfigResultRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *RecoveryConfigResultRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *RecoveryConfigResultRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *RecoveryConfigResultRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetConfigStage

`func (o *RecoveryConfigResultRelationship) GetConfigStage() string`

GetConfigStage returns the ConfigStage field if non-nil, zero value otherwise.

### GetConfigStageOk

`func (o *RecoveryConfigResultRelationship) GetConfigStageOk() (*string, bool)`

GetConfigStageOk returns a tuple with the ConfigStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStage

`func (o *RecoveryConfigResultRelationship) SetConfigStage(v string)`

SetConfigStage sets ConfigStage field to given value.

### HasConfigStage

`func (o *RecoveryConfigResultRelationship) HasConfigStage() bool`

HasConfigStage returns a boolean if a field has been set.

### GetConfigState

`func (o *RecoveryConfigResultRelationship) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *RecoveryConfigResultRelationship) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *RecoveryConfigResultRelationship) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *RecoveryConfigResultRelationship) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetValidationState

`func (o *RecoveryConfigResultRelationship) GetValidationState() string`

GetValidationState returns the ValidationState field if non-nil, zero value otherwise.

### GetValidationStateOk

`func (o *RecoveryConfigResultRelationship) GetValidationStateOk() (*string, bool)`

GetValidationStateOk returns a tuple with the ValidationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationState

`func (o *RecoveryConfigResultRelationship) SetValidationState(v string)`

SetValidationState sets ValidationState field to given value.

### HasValidationState

`func (o *RecoveryConfigResultRelationship) HasValidationState() bool`

HasValidationState returns a boolean if a field has been set.

### GetVar0OnDemandBackup

`func (o *RecoveryConfigResultRelationship) GetVar0OnDemandBackup() RecoveryOnDemandBackupRelationship`

GetVar0OnDemandBackup returns the Var0OnDemandBackup field if non-nil, zero value otherwise.

### GetVar0OnDemandBackupOk

`func (o *RecoveryConfigResultRelationship) GetVar0OnDemandBackupOk() (*RecoveryOnDemandBackupRelationship, bool)`

GetVar0OnDemandBackupOk returns a tuple with the Var0OnDemandBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0OnDemandBackup

`func (o *RecoveryConfigResultRelationship) SetVar0OnDemandBackup(v RecoveryOnDemandBackupRelationship)`

SetVar0OnDemandBackup sets Var0OnDemandBackup field to given value.

### HasVar0OnDemandBackup

`func (o *RecoveryConfigResultRelationship) HasVar0OnDemandBackup() bool`

HasVar0OnDemandBackup returns a boolean if a field has been set.

### GetBackupProfile

`func (o *RecoveryConfigResultRelationship) GetBackupProfile() RecoveryBackupProfileRelationship`

GetBackupProfile returns the BackupProfile field if non-nil, zero value otherwise.

### GetBackupProfileOk

`func (o *RecoveryConfigResultRelationship) GetBackupProfileOk() (*RecoveryBackupProfileRelationship, bool)`

GetBackupProfileOk returns a tuple with the BackupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProfile

`func (o *RecoveryConfigResultRelationship) SetBackupProfile(v RecoveryBackupProfileRelationship)`

SetBackupProfile sets BackupProfile field to given value.

### HasBackupProfile

`func (o *RecoveryConfigResultRelationship) HasBackupProfile() bool`

HasBackupProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *RecoveryConfigResultRelationship) GetResultEntries() []RecoveryConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *RecoveryConfigResultRelationship) GetResultEntriesOk() (*[]RecoveryConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *RecoveryConfigResultRelationship) SetResultEntries(v []RecoveryConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *RecoveryConfigResultRelationship) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *RecoveryConfigResultRelationship) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *RecoveryConfigResultRelationship) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


