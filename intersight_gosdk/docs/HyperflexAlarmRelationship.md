# HyperflexAlarmRelationship

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
**Acknowledged** | Pointer to **bool** |  | [optional] [readonly] 
**AcknowledgedBy** | Pointer to **string** |  | [optional] [readonly] 
**AcknowledgedTime** | Pointer to **int64** |  | [optional] [readonly] 
**AcknowledgedTimeAsUtc** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**EntityData** | Pointer to **string** |  | [optional] [readonly] 
**EntityName** | Pointer to **string** |  | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**EntityUuId** | Pointer to **string** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] [default to "UNKNOWN"]
**TriggeredTime** | Pointer to **int64** |  | [optional] [readonly] 
**TriggeredTimeAsUtc** | Pointer to **string** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexAlarmRelationship

`func NewHyperflexAlarmRelationship(classId string, objectType string, ) *HyperflexAlarmRelationship`

NewHyperflexAlarmRelationship instantiates a new HyperflexAlarmRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAlarmRelationshipWithDefaults

`func NewHyperflexAlarmRelationshipWithDefaults() *HyperflexAlarmRelationship`

NewHyperflexAlarmRelationshipWithDefaults instantiates a new HyperflexAlarmRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexAlarmRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexAlarmRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexAlarmRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexAlarmRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexAlarmRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAlarmRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAlarmRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexAlarmRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexAlarmRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexAlarmRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexAlarmRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexAlarmRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexAlarmRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexAlarmRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexAlarmRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexAlarmRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexAlarmRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexAlarmRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexAlarmRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexAlarmRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexAlarmRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexAlarmRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexAlarmRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexAlarmRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAlarmRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAlarmRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexAlarmRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexAlarmRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexAlarmRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexAlarmRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexAlarmRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexAlarmRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexAlarmRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexAlarmRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexAlarmRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexAlarmRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexAlarmRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexAlarmRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexAlarmRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexAlarmRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexAlarmRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexAlarmRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexAlarmRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexAlarmRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexAlarmRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexAlarmRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexAlarmRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexAlarmRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexAlarmRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexAlarmRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexAlarmRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexAlarmRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexAlarmRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexAlarmRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexAlarmRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexAlarmRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexAlarmRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexAlarmRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexAlarmRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexAlarmRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexAlarmRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexAlarmRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexAlarmRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexAlarmRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAcknowledged

`func (o *HyperflexAlarmRelationship) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *HyperflexAlarmRelationship) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *HyperflexAlarmRelationship) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *HyperflexAlarmRelationship) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetAcknowledgedBy

`func (o *HyperflexAlarmRelationship) GetAcknowledgedBy() string`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *HyperflexAlarmRelationship) GetAcknowledgedByOk() (*string, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *HyperflexAlarmRelationship) SetAcknowledgedBy(v string)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *HyperflexAlarmRelationship) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetAcknowledgedTime

`func (o *HyperflexAlarmRelationship) GetAcknowledgedTime() int64`

GetAcknowledgedTime returns the AcknowledgedTime field if non-nil, zero value otherwise.

### GetAcknowledgedTimeOk

`func (o *HyperflexAlarmRelationship) GetAcknowledgedTimeOk() (*int64, bool)`

GetAcknowledgedTimeOk returns a tuple with the AcknowledgedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTime

`func (o *HyperflexAlarmRelationship) SetAcknowledgedTime(v int64)`

SetAcknowledgedTime sets AcknowledgedTime field to given value.

### HasAcknowledgedTime

`func (o *HyperflexAlarmRelationship) HasAcknowledgedTime() bool`

HasAcknowledgedTime returns a boolean if a field has been set.

### GetAcknowledgedTimeAsUtc

`func (o *HyperflexAlarmRelationship) GetAcknowledgedTimeAsUtc() string`

GetAcknowledgedTimeAsUtc returns the AcknowledgedTimeAsUtc field if non-nil, zero value otherwise.

### GetAcknowledgedTimeAsUtcOk

`func (o *HyperflexAlarmRelationship) GetAcknowledgedTimeAsUtcOk() (*string, bool)`

GetAcknowledgedTimeAsUtcOk returns a tuple with the AcknowledgedTimeAsUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTimeAsUtc

`func (o *HyperflexAlarmRelationship) SetAcknowledgedTimeAsUtc(v string)`

SetAcknowledgedTimeAsUtc sets AcknowledgedTimeAsUtc field to given value.

### HasAcknowledgedTimeAsUtc

`func (o *HyperflexAlarmRelationship) HasAcknowledgedTimeAsUtc() bool`

HasAcknowledgedTimeAsUtc returns a boolean if a field has been set.

### GetDescription

`func (o *HyperflexAlarmRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexAlarmRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexAlarmRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexAlarmRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityData

`func (o *HyperflexAlarmRelationship) GetEntityData() string`

GetEntityData returns the EntityData field if non-nil, zero value otherwise.

### GetEntityDataOk

`func (o *HyperflexAlarmRelationship) GetEntityDataOk() (*string, bool)`

GetEntityDataOk returns a tuple with the EntityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityData

`func (o *HyperflexAlarmRelationship) SetEntityData(v string)`

SetEntityData sets EntityData field to given value.

### HasEntityData

`func (o *HyperflexAlarmRelationship) HasEntityData() bool`

HasEntityData returns a boolean if a field has been set.

### GetEntityName

`func (o *HyperflexAlarmRelationship) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *HyperflexAlarmRelationship) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *HyperflexAlarmRelationship) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *HyperflexAlarmRelationship) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *HyperflexAlarmRelationship) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HyperflexAlarmRelationship) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HyperflexAlarmRelationship) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *HyperflexAlarmRelationship) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityUuId

`func (o *HyperflexAlarmRelationship) GetEntityUuId() string`

GetEntityUuId returns the EntityUuId field if non-nil, zero value otherwise.

### GetEntityUuIdOk

`func (o *HyperflexAlarmRelationship) GetEntityUuIdOk() (*string, bool)`

GetEntityUuIdOk returns a tuple with the EntityUuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuId

`func (o *HyperflexAlarmRelationship) SetEntityUuId(v string)`

SetEntityUuId sets EntityUuId field to given value.

### HasEntityUuId

`func (o *HyperflexAlarmRelationship) HasEntityUuId() bool`

HasEntityUuId returns a boolean if a field has been set.

### GetMessage

`func (o *HyperflexAlarmRelationship) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HyperflexAlarmRelationship) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HyperflexAlarmRelationship) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HyperflexAlarmRelationship) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *HyperflexAlarmRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexAlarmRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexAlarmRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexAlarmRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexAlarmRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexAlarmRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexAlarmRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexAlarmRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggeredTime

`func (o *HyperflexAlarmRelationship) GetTriggeredTime() int64`

GetTriggeredTime returns the TriggeredTime field if non-nil, zero value otherwise.

### GetTriggeredTimeOk

`func (o *HyperflexAlarmRelationship) GetTriggeredTimeOk() (*int64, bool)`

GetTriggeredTimeOk returns a tuple with the TriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTime

`func (o *HyperflexAlarmRelationship) SetTriggeredTime(v int64)`

SetTriggeredTime sets TriggeredTime field to given value.

### HasTriggeredTime

`func (o *HyperflexAlarmRelationship) HasTriggeredTime() bool`

HasTriggeredTime returns a boolean if a field has been set.

### GetTriggeredTimeAsUtc

`func (o *HyperflexAlarmRelationship) GetTriggeredTimeAsUtc() string`

GetTriggeredTimeAsUtc returns the TriggeredTimeAsUtc field if non-nil, zero value otherwise.

### GetTriggeredTimeAsUtcOk

`func (o *HyperflexAlarmRelationship) GetTriggeredTimeAsUtcOk() (*string, bool)`

GetTriggeredTimeAsUtcOk returns a tuple with the TriggeredTimeAsUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTimeAsUtc

`func (o *HyperflexAlarmRelationship) SetTriggeredTimeAsUtc(v string)`

SetTriggeredTimeAsUtc sets TriggeredTimeAsUtc field to given value.

### HasTriggeredTimeAsUtc

`func (o *HyperflexAlarmRelationship) HasTriggeredTimeAsUtc() bool`

HasTriggeredTimeAsUtc returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexAlarmRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexAlarmRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexAlarmRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexAlarmRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexAlarmRelationship) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexAlarmRelationship) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexAlarmRelationship) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexAlarmRelationship) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


