# VnicFcAdapterPolicyRelationship

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
**ErrorDetectionTimeout** | Pointer to **int64** | Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. | [optional] 
**ErrorRecoverySettings** | Pointer to [**VnicFcErrorRecoverySettings**](vnic.FcErrorRecoverySettings.md) |  | [optional] 
**FlogiSettings** | Pointer to [**VnicFlogiSettings**](vnic.FlogiSettings.md) |  | [optional] 
**InterruptSettings** | Pointer to [**VnicFcInterruptSettings**](vnic.FcInterruptSettings.md) |  | [optional] 
**IoThrottleCount** | Pointer to **int64** | The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. | [optional] 
**LunCount** | Pointer to **int64** | The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. | [optional] 
**LunQueueDepth** | Pointer to **int64** | The number of commands that the HBA can send and receive in a single transmission per LUN. | [optional] 
**PlogiSettings** | Pointer to [**VnicPlogiSettings**](vnic.PlogiSettings.md) |  | [optional] 
**ResourceAllocationTimeout** | Pointer to **int64** | Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. | [optional] 
**RxQueueSettings** | Pointer to [**VnicFcQueueSettings**](vnic.FcQueueSettings.md) |  | [optional] 
**ScsiQueueSettings** | Pointer to [**VnicScsiQueueSettings**](vnic.ScsiQueueSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**VnicFcQueueSettings**](vnic.FcQueueSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicFcAdapterPolicyRelationship

`func NewVnicFcAdapterPolicyRelationship(classId string, objectType string, ) *VnicFcAdapterPolicyRelationship`

NewVnicFcAdapterPolicyRelationship instantiates a new VnicFcAdapterPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcAdapterPolicyRelationshipWithDefaults

`func NewVnicFcAdapterPolicyRelationshipWithDefaults() *VnicFcAdapterPolicyRelationship`

NewVnicFcAdapterPolicyRelationshipWithDefaults instantiates a new VnicFcAdapterPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VnicFcAdapterPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VnicFcAdapterPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VnicFcAdapterPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VnicFcAdapterPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VnicFcAdapterPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcAdapterPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcAdapterPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VnicFcAdapterPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VnicFcAdapterPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VnicFcAdapterPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VnicFcAdapterPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VnicFcAdapterPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VnicFcAdapterPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VnicFcAdapterPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VnicFcAdapterPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VnicFcAdapterPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VnicFcAdapterPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VnicFcAdapterPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VnicFcAdapterPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VnicFcAdapterPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VnicFcAdapterPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VnicFcAdapterPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VnicFcAdapterPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VnicFcAdapterPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcAdapterPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcAdapterPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VnicFcAdapterPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VnicFcAdapterPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VnicFcAdapterPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VnicFcAdapterPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VnicFcAdapterPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VnicFcAdapterPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VnicFcAdapterPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VnicFcAdapterPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VnicFcAdapterPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VnicFcAdapterPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VnicFcAdapterPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VnicFcAdapterPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VnicFcAdapterPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VnicFcAdapterPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VnicFcAdapterPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VnicFcAdapterPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VnicFcAdapterPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VnicFcAdapterPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VnicFcAdapterPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VnicFcAdapterPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VnicFcAdapterPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VnicFcAdapterPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VnicFcAdapterPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VnicFcAdapterPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VnicFcAdapterPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VnicFcAdapterPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VnicFcAdapterPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VnicFcAdapterPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VnicFcAdapterPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VnicFcAdapterPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VnicFcAdapterPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VnicFcAdapterPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VnicFcAdapterPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VnicFcAdapterPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VnicFcAdapterPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VnicFcAdapterPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VnicFcAdapterPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VnicFcAdapterPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *VnicFcAdapterPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VnicFcAdapterPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VnicFcAdapterPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VnicFcAdapterPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VnicFcAdapterPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicFcAdapterPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicFcAdapterPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicFcAdapterPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyRelationship) GetErrorDetectionTimeout() int64`

GetErrorDetectionTimeout returns the ErrorDetectionTimeout field if non-nil, zero value otherwise.

### GetErrorDetectionTimeoutOk

`func (o *VnicFcAdapterPolicyRelationship) GetErrorDetectionTimeoutOk() (*int64, bool)`

GetErrorDetectionTimeoutOk returns a tuple with the ErrorDetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyRelationship) SetErrorDetectionTimeout(v int64)`

SetErrorDetectionTimeout sets ErrorDetectionTimeout field to given value.

### HasErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyRelationship) HasErrorDetectionTimeout() bool`

HasErrorDetectionTimeout returns a boolean if a field has been set.

### GetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyRelationship) GetErrorRecoverySettings() VnicFcErrorRecoverySettings`

GetErrorRecoverySettings returns the ErrorRecoverySettings field if non-nil, zero value otherwise.

### GetErrorRecoverySettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetErrorRecoverySettingsOk() (*VnicFcErrorRecoverySettings, bool)`

GetErrorRecoverySettingsOk returns a tuple with the ErrorRecoverySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyRelationship) SetErrorRecoverySettings(v VnicFcErrorRecoverySettings)`

SetErrorRecoverySettings sets ErrorRecoverySettings field to given value.

### HasErrorRecoverySettings

`func (o *VnicFcAdapterPolicyRelationship) HasErrorRecoverySettings() bool`

HasErrorRecoverySettings returns a boolean if a field has been set.

### GetFlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) GetFlogiSettings() VnicFlogiSettings`

GetFlogiSettings returns the FlogiSettings field if non-nil, zero value otherwise.

### GetFlogiSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetFlogiSettingsOk() (*VnicFlogiSettings, bool)`

GetFlogiSettingsOk returns a tuple with the FlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) SetFlogiSettings(v VnicFlogiSettings)`

SetFlogiSettings sets FlogiSettings field to given value.

### HasFlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) HasFlogiSettings() bool`

HasFlogiSettings returns a boolean if a field has been set.

### GetInterruptSettings

`func (o *VnicFcAdapterPolicyRelationship) GetInterruptSettings() VnicFcInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetInterruptSettingsOk() (*VnicFcInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicFcAdapterPolicyRelationship) SetInterruptSettings(v VnicFcInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicFcAdapterPolicyRelationship) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### GetIoThrottleCount

`func (o *VnicFcAdapterPolicyRelationship) GetIoThrottleCount() int64`

GetIoThrottleCount returns the IoThrottleCount field if non-nil, zero value otherwise.

### GetIoThrottleCountOk

`func (o *VnicFcAdapterPolicyRelationship) GetIoThrottleCountOk() (*int64, bool)`

GetIoThrottleCountOk returns a tuple with the IoThrottleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoThrottleCount

`func (o *VnicFcAdapterPolicyRelationship) SetIoThrottleCount(v int64)`

SetIoThrottleCount sets IoThrottleCount field to given value.

### HasIoThrottleCount

`func (o *VnicFcAdapterPolicyRelationship) HasIoThrottleCount() bool`

HasIoThrottleCount returns a boolean if a field has been set.

### GetLunCount

`func (o *VnicFcAdapterPolicyRelationship) GetLunCount() int64`

GetLunCount returns the LunCount field if non-nil, zero value otherwise.

### GetLunCountOk

`func (o *VnicFcAdapterPolicyRelationship) GetLunCountOk() (*int64, bool)`

GetLunCountOk returns a tuple with the LunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunCount

`func (o *VnicFcAdapterPolicyRelationship) SetLunCount(v int64)`

SetLunCount sets LunCount field to given value.

### HasLunCount

`func (o *VnicFcAdapterPolicyRelationship) HasLunCount() bool`

HasLunCount returns a boolean if a field has been set.

### GetLunQueueDepth

`func (o *VnicFcAdapterPolicyRelationship) GetLunQueueDepth() int64`

GetLunQueueDepth returns the LunQueueDepth field if non-nil, zero value otherwise.

### GetLunQueueDepthOk

`func (o *VnicFcAdapterPolicyRelationship) GetLunQueueDepthOk() (*int64, bool)`

GetLunQueueDepthOk returns a tuple with the LunQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunQueueDepth

`func (o *VnicFcAdapterPolicyRelationship) SetLunQueueDepth(v int64)`

SetLunQueueDepth sets LunQueueDepth field to given value.

### HasLunQueueDepth

`func (o *VnicFcAdapterPolicyRelationship) HasLunQueueDepth() bool`

HasLunQueueDepth returns a boolean if a field has been set.

### GetPlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) GetPlogiSettings() VnicPlogiSettings`

GetPlogiSettings returns the PlogiSettings field if non-nil, zero value otherwise.

### GetPlogiSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetPlogiSettingsOk() (*VnicPlogiSettings, bool)`

GetPlogiSettingsOk returns a tuple with the PlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) SetPlogiSettings(v VnicPlogiSettings)`

SetPlogiSettings sets PlogiSettings field to given value.

### HasPlogiSettings

`func (o *VnicFcAdapterPolicyRelationship) HasPlogiSettings() bool`

HasPlogiSettings returns a boolean if a field has been set.

### GetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyRelationship) GetResourceAllocationTimeout() int64`

GetResourceAllocationTimeout returns the ResourceAllocationTimeout field if non-nil, zero value otherwise.

### GetResourceAllocationTimeoutOk

`func (o *VnicFcAdapterPolicyRelationship) GetResourceAllocationTimeoutOk() (*int64, bool)`

GetResourceAllocationTimeoutOk returns a tuple with the ResourceAllocationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyRelationship) SetResourceAllocationTimeout(v int64)`

SetResourceAllocationTimeout sets ResourceAllocationTimeout field to given value.

### HasResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyRelationship) HasResourceAllocationTimeout() bool`

HasResourceAllocationTimeout returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) GetRxQueueSettings() VnicFcQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetRxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) SetRxQueueSettings(v VnicFcQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### GetScsiQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) GetScsiQueueSettings() VnicScsiQueueSettings`

GetScsiQueueSettings returns the ScsiQueueSettings field if non-nil, zero value otherwise.

### GetScsiQueueSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetScsiQueueSettingsOk() (*VnicScsiQueueSettings, bool)`

GetScsiQueueSettingsOk returns a tuple with the ScsiQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) SetScsiQueueSettings(v VnicScsiQueueSettings)`

SetScsiQueueSettings sets ScsiQueueSettings field to given value.

### HasScsiQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) HasScsiQueueSettings() bool`

HasScsiQueueSettings returns a boolean if a field has been set.

### GetTxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) GetTxQueueSettings() VnicFcQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicFcAdapterPolicyRelationship) GetTxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) SetTxQueueSettings(v VnicFcQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicFcAdapterPolicyRelationship) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicFcAdapterPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicFcAdapterPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicFcAdapterPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicFcAdapterPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


