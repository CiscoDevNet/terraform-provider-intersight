# VnicEthAdapterPolicyRelationship

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
**AdvancedFilter** | Pointer to **bool** | Enables advanced filtering on the interface. | [optional] 
**ArfsSettings** | Pointer to [**VnicArfsSettings**](vnic.ArfsSettings.md) |  | [optional] 
**CompletionQueueSettings** | Pointer to [**VnicCompletionQueueSettings**](vnic.CompletionQueueSettings.md) |  | [optional] 
**InterruptScaling** | Pointer to **bool** | Enables Interrupt Scaling on the interface. | [optional] 
**InterruptSettings** | Pointer to [**VnicEthInterruptSettings**](vnic.EthInterruptSettings.md) |  | [optional] 
**NvgreSettings** | Pointer to [**VnicNvgreSettings**](vnic.NvgreSettings.md) |  | [optional] 
**RoceSettings** | Pointer to [**VnicRoceSettings**](vnic.RoceSettings.md) |  | [optional] 
**RssHashSettings** | Pointer to [**VnicRssHashSettings**](vnic.RssHashSettings.md) |  | [optional] 
**RssSettings** | Pointer to **bool** | Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. | [optional] 
**RxQueueSettings** | Pointer to [**VnicEthRxQueueSettings**](vnic.EthRxQueueSettings.md) |  | [optional] 
**TcpOffloadSettings** | Pointer to [**VnicTcpOffloadSettings**](vnic.TcpOffloadSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**VnicEthTxQueueSettings**](vnic.EthTxQueueSettings.md) |  | [optional] 
**UplinkFailbackTimeout** | Pointer to **int64** | Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC. | [optional] 
**VxlanSettings** | Pointer to [**VnicVxlanSettings**](vnic.VxlanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicEthAdapterPolicyRelationship

`func NewVnicEthAdapterPolicyRelationship(classId string, objectType string, ) *VnicEthAdapterPolicyRelationship`

NewVnicEthAdapterPolicyRelationship instantiates a new VnicEthAdapterPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthAdapterPolicyRelationshipWithDefaults

`func NewVnicEthAdapterPolicyRelationshipWithDefaults() *VnicEthAdapterPolicyRelationship`

NewVnicEthAdapterPolicyRelationshipWithDefaults instantiates a new VnicEthAdapterPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *VnicEthAdapterPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *VnicEthAdapterPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *VnicEthAdapterPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *VnicEthAdapterPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *VnicEthAdapterPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthAdapterPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthAdapterPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *VnicEthAdapterPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *VnicEthAdapterPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *VnicEthAdapterPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *VnicEthAdapterPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *VnicEthAdapterPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *VnicEthAdapterPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *VnicEthAdapterPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *VnicEthAdapterPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *VnicEthAdapterPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *VnicEthAdapterPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *VnicEthAdapterPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *VnicEthAdapterPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *VnicEthAdapterPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *VnicEthAdapterPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *VnicEthAdapterPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *VnicEthAdapterPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *VnicEthAdapterPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthAdapterPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthAdapterPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *VnicEthAdapterPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *VnicEthAdapterPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *VnicEthAdapterPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *VnicEthAdapterPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *VnicEthAdapterPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *VnicEthAdapterPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *VnicEthAdapterPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *VnicEthAdapterPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *VnicEthAdapterPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VnicEthAdapterPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VnicEthAdapterPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VnicEthAdapterPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *VnicEthAdapterPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *VnicEthAdapterPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *VnicEthAdapterPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *VnicEthAdapterPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *VnicEthAdapterPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *VnicEthAdapterPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *VnicEthAdapterPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *VnicEthAdapterPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *VnicEthAdapterPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *VnicEthAdapterPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *VnicEthAdapterPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VnicEthAdapterPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VnicEthAdapterPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VnicEthAdapterPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *VnicEthAdapterPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *VnicEthAdapterPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *VnicEthAdapterPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *VnicEthAdapterPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *VnicEthAdapterPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *VnicEthAdapterPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *VnicEthAdapterPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *VnicEthAdapterPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *VnicEthAdapterPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *VnicEthAdapterPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *VnicEthAdapterPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *VnicEthAdapterPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *VnicEthAdapterPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VnicEthAdapterPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VnicEthAdapterPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VnicEthAdapterPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VnicEthAdapterPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthAdapterPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthAdapterPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthAdapterPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdvancedFilter

`func (o *VnicEthAdapterPolicyRelationship) GetAdvancedFilter() bool`

GetAdvancedFilter returns the AdvancedFilter field if non-nil, zero value otherwise.

### GetAdvancedFilterOk

`func (o *VnicEthAdapterPolicyRelationship) GetAdvancedFilterOk() (*bool, bool)`

GetAdvancedFilterOk returns a tuple with the AdvancedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFilter

`func (o *VnicEthAdapterPolicyRelationship) SetAdvancedFilter(v bool)`

SetAdvancedFilter sets AdvancedFilter field to given value.

### HasAdvancedFilter

`func (o *VnicEthAdapterPolicyRelationship) HasAdvancedFilter() bool`

HasAdvancedFilter returns a boolean if a field has been set.

### GetArfsSettings

`func (o *VnicEthAdapterPolicyRelationship) GetArfsSettings() VnicArfsSettings`

GetArfsSettings returns the ArfsSettings field if non-nil, zero value otherwise.

### GetArfsSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetArfsSettingsOk() (*VnicArfsSettings, bool)`

GetArfsSettingsOk returns a tuple with the ArfsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArfsSettings

`func (o *VnicEthAdapterPolicyRelationship) SetArfsSettings(v VnicArfsSettings)`

SetArfsSettings sets ArfsSettings field to given value.

### HasArfsSettings

`func (o *VnicEthAdapterPolicyRelationship) HasArfsSettings() bool`

HasArfsSettings returns a boolean if a field has been set.

### GetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) GetCompletionQueueSettings() VnicCompletionQueueSettings`

GetCompletionQueueSettings returns the CompletionQueueSettings field if non-nil, zero value otherwise.

### GetCompletionQueueSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool)`

GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) SetCompletionQueueSettings(v VnicCompletionQueueSettings)`

SetCompletionQueueSettings sets CompletionQueueSettings field to given value.

### HasCompletionQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) HasCompletionQueueSettings() bool`

HasCompletionQueueSettings returns a boolean if a field has been set.

### GetInterruptScaling

`func (o *VnicEthAdapterPolicyRelationship) GetInterruptScaling() bool`

GetInterruptScaling returns the InterruptScaling field if non-nil, zero value otherwise.

### GetInterruptScalingOk

`func (o *VnicEthAdapterPolicyRelationship) GetInterruptScalingOk() (*bool, bool)`

GetInterruptScalingOk returns a tuple with the InterruptScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptScaling

`func (o *VnicEthAdapterPolicyRelationship) SetInterruptScaling(v bool)`

SetInterruptScaling sets InterruptScaling field to given value.

### HasInterruptScaling

`func (o *VnicEthAdapterPolicyRelationship) HasInterruptScaling() bool`

HasInterruptScaling returns a boolean if a field has been set.

### GetInterruptSettings

`func (o *VnicEthAdapterPolicyRelationship) GetInterruptSettings() VnicEthInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicEthAdapterPolicyRelationship) SetInterruptSettings(v VnicEthInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicEthAdapterPolicyRelationship) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### GetNvgreSettings

`func (o *VnicEthAdapterPolicyRelationship) GetNvgreSettings() VnicNvgreSettings`

GetNvgreSettings returns the NvgreSettings field if non-nil, zero value otherwise.

### GetNvgreSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetNvgreSettingsOk() (*VnicNvgreSettings, bool)`

GetNvgreSettingsOk returns a tuple with the NvgreSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvgreSettings

`func (o *VnicEthAdapterPolicyRelationship) SetNvgreSettings(v VnicNvgreSettings)`

SetNvgreSettings sets NvgreSettings field to given value.

### HasNvgreSettings

`func (o *VnicEthAdapterPolicyRelationship) HasNvgreSettings() bool`

HasNvgreSettings returns a boolean if a field has been set.

### GetRoceSettings

`func (o *VnicEthAdapterPolicyRelationship) GetRoceSettings() VnicRoceSettings`

GetRoceSettings returns the RoceSettings field if non-nil, zero value otherwise.

### GetRoceSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetRoceSettingsOk() (*VnicRoceSettings, bool)`

GetRoceSettingsOk returns a tuple with the RoceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoceSettings

`func (o *VnicEthAdapterPolicyRelationship) SetRoceSettings(v VnicRoceSettings)`

SetRoceSettings sets RoceSettings field to given value.

### HasRoceSettings

`func (o *VnicEthAdapterPolicyRelationship) HasRoceSettings() bool`

HasRoceSettings returns a boolean if a field has been set.

### GetRssHashSettings

`func (o *VnicEthAdapterPolicyRelationship) GetRssHashSettings() VnicRssHashSettings`

GetRssHashSettings returns the RssHashSettings field if non-nil, zero value otherwise.

### GetRssHashSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetRssHashSettingsOk() (*VnicRssHashSettings, bool)`

GetRssHashSettingsOk returns a tuple with the RssHashSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssHashSettings

`func (o *VnicEthAdapterPolicyRelationship) SetRssHashSettings(v VnicRssHashSettings)`

SetRssHashSettings sets RssHashSettings field to given value.

### HasRssHashSettings

`func (o *VnicEthAdapterPolicyRelationship) HasRssHashSettings() bool`

HasRssHashSettings returns a boolean if a field has been set.

### GetRssSettings

`func (o *VnicEthAdapterPolicyRelationship) GetRssSettings() bool`

GetRssSettings returns the RssSettings field if non-nil, zero value otherwise.

### GetRssSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetRssSettingsOk() (*bool, bool)`

GetRssSettingsOk returns a tuple with the RssSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSettings

`func (o *VnicEthAdapterPolicyRelationship) SetRssSettings(v bool)`

SetRssSettings sets RssSettings field to given value.

### HasRssSettings

`func (o *VnicEthAdapterPolicyRelationship) HasRssSettings() bool`

HasRssSettings returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) GetRxQueueSettings() VnicEthRxQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) SetRxQueueSettings(v VnicEthRxQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### GetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyRelationship) GetTcpOffloadSettings() VnicTcpOffloadSettings`

GetTcpOffloadSettings returns the TcpOffloadSettings field if non-nil, zero value otherwise.

### GetTcpOffloadSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool)`

GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyRelationship) SetTcpOffloadSettings(v VnicTcpOffloadSettings)`

SetTcpOffloadSettings sets TcpOffloadSettings field to given value.

### HasTcpOffloadSettings

`func (o *VnicEthAdapterPolicyRelationship) HasTcpOffloadSettings() bool`

HasTcpOffloadSettings returns a boolean if a field has been set.

### GetTxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) GetTxQueueSettings() VnicEthTxQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) SetTxQueueSettings(v VnicEthTxQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicEthAdapterPolicyRelationship) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### GetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyRelationship) GetUplinkFailbackTimeout() int64`

GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field if non-nil, zero value otherwise.

### GetUplinkFailbackTimeoutOk

`func (o *VnicEthAdapterPolicyRelationship) GetUplinkFailbackTimeoutOk() (*int64, bool)`

GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyRelationship) SetUplinkFailbackTimeout(v int64)`

SetUplinkFailbackTimeout sets UplinkFailbackTimeout field to given value.

### HasUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyRelationship) HasUplinkFailbackTimeout() bool`

HasUplinkFailbackTimeout returns a boolean if a field has been set.

### GetVxlanSettings

`func (o *VnicEthAdapterPolicyRelationship) GetVxlanSettings() VnicVxlanSettings`

GetVxlanSettings returns the VxlanSettings field if non-nil, zero value otherwise.

### GetVxlanSettingsOk

`func (o *VnicEthAdapterPolicyRelationship) GetVxlanSettingsOk() (*VnicVxlanSettings, bool)`

GetVxlanSettingsOk returns a tuple with the VxlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanSettings

`func (o *VnicEthAdapterPolicyRelationship) SetVxlanSettings(v VnicVxlanSettings)`

SetVxlanSettings sets VxlanSettings field to given value.

### HasVxlanSettings

`func (o *VnicEthAdapterPolicyRelationship) HasVxlanSettings() bool`

HasVxlanSettings returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicEthAdapterPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicEthAdapterPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicEthAdapterPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicEthAdapterPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


