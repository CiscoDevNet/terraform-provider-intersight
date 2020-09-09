# StoragePureProtectionGroupRelationship

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
**Name** | Pointer to **string** | Name of the protection Group. | [optional] 
**Prefix** | Pointer to **string** | Prefix used for all generated snapshots from the protection group. | [optional] 
**ReplicationEnabled** | Pointer to **bool** | Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. | [optional] 
**SnapshotEnabled** | Pointer to **bool** | Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. | [optional] 
**Size** | Pointer to **int64** | Overall size of all snapshots in the protection group, represented in bytes. | [optional] 
**Source** | Pointer to **string** | Name of PureStorage array name on which the protection group is created. | [optional] [readonly] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**HostGroups** | Pointer to [**[]StoragePureHostGroupRelationship**](storage.PureHostGroup.Relationship.md) | An array of relationships to storagePureHostGroup resources. | [optional] [readonly] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](storage.PureHost.Relationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Volumes** | Pointer to [**[]StoragePureVolumeRelationship**](storage.PureVolume.Relationship.md) | An array of relationships to storagePureVolume resources. | [optional] [readonly] 

## Methods

### NewStoragePureProtectionGroupRelationship

`func NewStoragePureProtectionGroupRelationship(classId string, objectType string, ) *StoragePureProtectionGroupRelationship`

NewStoragePureProtectionGroupRelationship instantiates a new StoragePureProtectionGroupRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureProtectionGroupRelationshipWithDefaults

`func NewStoragePureProtectionGroupRelationshipWithDefaults() *StoragePureProtectionGroupRelationship`

NewStoragePureProtectionGroupRelationshipWithDefaults instantiates a new StoragePureProtectionGroupRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StoragePureProtectionGroupRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StoragePureProtectionGroupRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StoragePureProtectionGroupRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StoragePureProtectionGroupRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StoragePureProtectionGroupRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureProtectionGroupRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureProtectionGroupRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StoragePureProtectionGroupRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StoragePureProtectionGroupRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StoragePureProtectionGroupRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StoragePureProtectionGroupRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StoragePureProtectionGroupRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StoragePureProtectionGroupRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StoragePureProtectionGroupRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StoragePureProtectionGroupRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StoragePureProtectionGroupRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StoragePureProtectionGroupRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StoragePureProtectionGroupRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StoragePureProtectionGroupRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StoragePureProtectionGroupRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StoragePureProtectionGroupRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StoragePureProtectionGroupRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StoragePureProtectionGroupRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StoragePureProtectionGroupRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureProtectionGroupRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureProtectionGroupRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StoragePureProtectionGroupRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StoragePureProtectionGroupRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StoragePureProtectionGroupRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StoragePureProtectionGroupRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StoragePureProtectionGroupRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StoragePureProtectionGroupRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StoragePureProtectionGroupRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StoragePureProtectionGroupRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StoragePureProtectionGroupRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoragePureProtectionGroupRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoragePureProtectionGroupRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoragePureProtectionGroupRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StoragePureProtectionGroupRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StoragePureProtectionGroupRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StoragePureProtectionGroupRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StoragePureProtectionGroupRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StoragePureProtectionGroupRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StoragePureProtectionGroupRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StoragePureProtectionGroupRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StoragePureProtectionGroupRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StoragePureProtectionGroupRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StoragePureProtectionGroupRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StoragePureProtectionGroupRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StoragePureProtectionGroupRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StoragePureProtectionGroupRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StoragePureProtectionGroupRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StoragePureProtectionGroupRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StoragePureProtectionGroupRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StoragePureProtectionGroupRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StoragePureProtectionGroupRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StoragePureProtectionGroupRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StoragePureProtectionGroupRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StoragePureProtectionGroupRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StoragePureProtectionGroupRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StoragePureProtectionGroupRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StoragePureProtectionGroupRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StoragePureProtectionGroupRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StoragePureProtectionGroupRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetName

`func (o *StoragePureProtectionGroupRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureProtectionGroupRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureProtectionGroupRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureProtectionGroupRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *StoragePureProtectionGroupRelationship) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *StoragePureProtectionGroupRelationship) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *StoragePureProtectionGroupRelationship) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *StoragePureProtectionGroupRelationship) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReplicationEnabled

`func (o *StoragePureProtectionGroupRelationship) GetReplicationEnabled() bool`

GetReplicationEnabled returns the ReplicationEnabled field if non-nil, zero value otherwise.

### GetReplicationEnabledOk

`func (o *StoragePureProtectionGroupRelationship) GetReplicationEnabledOk() (*bool, bool)`

GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabled

`func (o *StoragePureProtectionGroupRelationship) SetReplicationEnabled(v bool)`

SetReplicationEnabled sets ReplicationEnabled field to given value.

### HasReplicationEnabled

`func (o *StoragePureProtectionGroupRelationship) HasReplicationEnabled() bool`

HasReplicationEnabled returns a boolean if a field has been set.

### GetSnapshotEnabled

`func (o *StoragePureProtectionGroupRelationship) GetSnapshotEnabled() bool`

GetSnapshotEnabled returns the SnapshotEnabled field if non-nil, zero value otherwise.

### GetSnapshotEnabledOk

`func (o *StoragePureProtectionGroupRelationship) GetSnapshotEnabledOk() (*bool, bool)`

GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnabled

`func (o *StoragePureProtectionGroupRelationship) SetSnapshotEnabled(v bool)`

SetSnapshotEnabled sets SnapshotEnabled field to given value.

### HasSnapshotEnabled

`func (o *StoragePureProtectionGroupRelationship) HasSnapshotEnabled() bool`

HasSnapshotEnabled returns a boolean if a field has been set.

### GetSize

`func (o *StoragePureProtectionGroupRelationship) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StoragePureProtectionGroupRelationship) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StoragePureProtectionGroupRelationship) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StoragePureProtectionGroupRelationship) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSource

`func (o *StoragePureProtectionGroupRelationship) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StoragePureProtectionGroupRelationship) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StoragePureProtectionGroupRelationship) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StoragePureProtectionGroupRelationship) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *StoragePureProtectionGroupRelationship) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *StoragePureProtectionGroupRelationship) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *StoragePureProtectionGroupRelationship) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *StoragePureProtectionGroupRelationship) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureProtectionGroupRelationship) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureProtectionGroupRelationship) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureProtectionGroupRelationship) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureProtectionGroupRelationship) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHostGroups

`func (o *StoragePureProtectionGroupRelationship) GetHostGroups() []StoragePureHostGroupRelationship`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *StoragePureProtectionGroupRelationship) GetHostGroupsOk() (*[]StoragePureHostGroupRelationship, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *StoragePureProtectionGroupRelationship) SetHostGroups(v []StoragePureHostGroupRelationship)`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *StoragePureProtectionGroupRelationship) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### SetHostGroupsNil

`func (o *StoragePureProtectionGroupRelationship) SetHostGroupsNil(b bool)`

 SetHostGroupsNil sets the value for HostGroups to be an explicit nil

### UnsetHostGroups
`func (o *StoragePureProtectionGroupRelationship) UnsetHostGroups()`

UnsetHostGroups ensures that no value is present for HostGroups, not even an explicit nil
### GetHosts

`func (o *StoragePureProtectionGroupRelationship) GetHosts() []StoragePureHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StoragePureProtectionGroupRelationship) GetHostsOk() (*[]StoragePureHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StoragePureProtectionGroupRelationship) SetHosts(v []StoragePureHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StoragePureProtectionGroupRelationship) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *StoragePureProtectionGroupRelationship) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *StoragePureProtectionGroupRelationship) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureProtectionGroupRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureProtectionGroupRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureProtectionGroupRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureProtectionGroupRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVolumes

`func (o *StoragePureProtectionGroupRelationship) GetVolumes() []StoragePureVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StoragePureProtectionGroupRelationship) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StoragePureProtectionGroupRelationship) SetVolumes(v []StoragePureVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StoragePureProtectionGroupRelationship) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *StoragePureProtectionGroupRelationship) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *StoragePureProtectionGroupRelationship) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


