# MemoryPersistentMemoryConfigResultRelationship

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
**DeviceMoId** | Pointer to **string** | The database identifier of the registered device of an object. | [optional] [readonly] 
**Dn** | Pointer to **string** | The Distinguished Name unambiguously identifies an object in the system. | [optional] [readonly] 
**Rn** | Pointer to **string** | The Relative Name uniquely identifies an object within a given context. | [optional] [readonly] 
**ConfigErrorDesc** | Pointer to **string** | Error in the result of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigResult** | Pointer to **string** | Result of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigSequenceNo** | Pointer to **int64** | Sequence number of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | State of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**PersistentMemoryNamespaceConfigResults** | Pointer to [**[]MemoryPersistentMemoryNamespaceConfigResultRelationship**](memory.PersistentMemoryNamespaceConfigResult.Relationship.md) | An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryConfigResultRelationship

`func NewMemoryPersistentMemoryConfigResultRelationship(classId string, objectType string, ) *MemoryPersistentMemoryConfigResultRelationship`

NewMemoryPersistentMemoryConfigResultRelationship instantiates a new MemoryPersistentMemoryConfigResultRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryConfigResultRelationshipWithDefaults

`func NewMemoryPersistentMemoryConfigResultRelationshipWithDefaults() *MemoryPersistentMemoryConfigResultRelationship`

NewMemoryPersistentMemoryConfigResultRelationshipWithDefaults instantiates a new MemoryPersistentMemoryConfigResultRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *MemoryPersistentMemoryConfigResultRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *MemoryPersistentMemoryConfigResultRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *MemoryPersistentMemoryConfigResultRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigErrorDesc() string`

GetConfigErrorDesc returns the ConfigErrorDesc field if non-nil, zero value otherwise.

### GetConfigErrorDescOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigErrorDescOk() (*string, bool)`

GetConfigErrorDescOk returns a tuple with the ConfigErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetConfigErrorDesc(v string)`

SetConfigErrorDesc sets ConfigErrorDesc field to given value.

### HasConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasConfigErrorDesc() bool`

HasConfigErrorDesc returns a boolean if a field has been set.

### GetConfigResult

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigResult() string`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigResultOk() (*string, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetConfigResult(v string)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigSequenceNo() int64`

GetConfigSequenceNo returns the ConfigSequenceNo field if non-nil, zero value otherwise.

### GetConfigSequenceNoOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigSequenceNoOk() (*int64, bool)`

GetConfigSequenceNoOk returns a tuple with the ConfigSequenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetConfigSequenceNo(v int64)`

SetConfigSequenceNo sets ConfigSequenceNo field to given value.

### HasConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasConfigSequenceNo() bool`

HasConfigSequenceNo returns a boolean if a field has been set.

### GetConfigState

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetPersistentMemoryNamespaceConfigResults() []MemoryPersistentMemoryNamespaceConfigResultRelationship`

GetPersistentMemoryNamespaceConfigResults returns the PersistentMemoryNamespaceConfigResults field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespaceConfigResultsOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetPersistentMemoryNamespaceConfigResultsOk() (*[]MemoryPersistentMemoryNamespaceConfigResultRelationship, bool)`

GetPersistentMemoryNamespaceConfigResultsOk returns a tuple with the PersistentMemoryNamespaceConfigResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetPersistentMemoryNamespaceConfigResults(v []MemoryPersistentMemoryNamespaceConfigResultRelationship)`

SetPersistentMemoryNamespaceConfigResults sets PersistentMemoryNamespaceConfigResults field to given value.

### HasPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasPersistentMemoryNamespaceConfigResults() bool`

HasPersistentMemoryNamespaceConfigResults returns a boolean if a field has been set.

### SetPersistentMemoryNamespaceConfigResultsNil

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetPersistentMemoryNamespaceConfigResultsNil(b bool)`

 SetPersistentMemoryNamespaceConfigResultsNil sets the value for PersistentMemoryNamespaceConfigResults to be an explicit nil

### UnsetPersistentMemoryNamespaceConfigResults
`func (o *MemoryPersistentMemoryConfigResultRelationship) UnsetPersistentMemoryNamespaceConfigResults()`

UnsetPersistentMemoryNamespaceConfigResults ensures that no value is present for PersistentMemoryNamespaceConfigResults, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryConfigResultRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


