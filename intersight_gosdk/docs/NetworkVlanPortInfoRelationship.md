# NetworkVlanPortInfoRelationship

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
**AccessVlanPortCount** | Pointer to **int64** | The number of available VLAN access ports on a Fabric Interconnect. | [optional] [readonly] 
**BorderVlanPortCount** | Pointer to **int64** | The number of available VLAN border ports on a Fabric Interconnect. | [optional] [readonly] 
**CompressedVlanPortCount** | Pointer to **string** | The number of compressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**UncompressedVlanPortCount** | Pointer to **string** | The number of uncompressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**VlanPortLimit** | Pointer to **int64** | The maximum number of VLAN ports allowed on a Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNetworkVlanPortInfoRelationship

`func NewNetworkVlanPortInfoRelationship(classId string, objectType string, ) *NetworkVlanPortInfoRelationship`

NewNetworkVlanPortInfoRelationship instantiates a new NetworkVlanPortInfoRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlanPortInfoRelationshipWithDefaults

`func NewNetworkVlanPortInfoRelationshipWithDefaults() *NetworkVlanPortInfoRelationship`

NewNetworkVlanPortInfoRelationshipWithDefaults instantiates a new NetworkVlanPortInfoRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *NetworkVlanPortInfoRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *NetworkVlanPortInfoRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *NetworkVlanPortInfoRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *NetworkVlanPortInfoRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *NetworkVlanPortInfoRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVlanPortInfoRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVlanPortInfoRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *NetworkVlanPortInfoRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NetworkVlanPortInfoRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NetworkVlanPortInfoRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NetworkVlanPortInfoRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *NetworkVlanPortInfoRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *NetworkVlanPortInfoRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *NetworkVlanPortInfoRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *NetworkVlanPortInfoRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *NetworkVlanPortInfoRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *NetworkVlanPortInfoRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *NetworkVlanPortInfoRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *NetworkVlanPortInfoRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *NetworkVlanPortInfoRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *NetworkVlanPortInfoRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *NetworkVlanPortInfoRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *NetworkVlanPortInfoRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *NetworkVlanPortInfoRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVlanPortInfoRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVlanPortInfoRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *NetworkVlanPortInfoRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *NetworkVlanPortInfoRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *NetworkVlanPortInfoRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *NetworkVlanPortInfoRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *NetworkVlanPortInfoRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *NetworkVlanPortInfoRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *NetworkVlanPortInfoRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *NetworkVlanPortInfoRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *NetworkVlanPortInfoRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkVlanPortInfoRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkVlanPortInfoRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworkVlanPortInfoRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *NetworkVlanPortInfoRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *NetworkVlanPortInfoRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *NetworkVlanPortInfoRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *NetworkVlanPortInfoRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *NetworkVlanPortInfoRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *NetworkVlanPortInfoRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *NetworkVlanPortInfoRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *NetworkVlanPortInfoRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *NetworkVlanPortInfoRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *NetworkVlanPortInfoRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *NetworkVlanPortInfoRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *NetworkVlanPortInfoRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *NetworkVlanPortInfoRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *NetworkVlanPortInfoRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *NetworkVlanPortInfoRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *NetworkVlanPortInfoRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *NetworkVlanPortInfoRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *NetworkVlanPortInfoRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *NetworkVlanPortInfoRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *NetworkVlanPortInfoRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *NetworkVlanPortInfoRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *NetworkVlanPortInfoRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *NetworkVlanPortInfoRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *NetworkVlanPortInfoRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *NetworkVlanPortInfoRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *NetworkVlanPortInfoRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *NetworkVlanPortInfoRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *NetworkVlanPortInfoRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *NetworkVlanPortInfoRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *NetworkVlanPortInfoRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *NetworkVlanPortInfoRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NetworkVlanPortInfoRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NetworkVlanPortInfoRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NetworkVlanPortInfoRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *NetworkVlanPortInfoRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *NetworkVlanPortInfoRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *NetworkVlanPortInfoRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *NetworkVlanPortInfoRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetAccessVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) GetAccessVlanPortCount() int64`

GetAccessVlanPortCount returns the AccessVlanPortCount field if non-nil, zero value otherwise.

### GetAccessVlanPortCountOk

`func (o *NetworkVlanPortInfoRelationship) GetAccessVlanPortCountOk() (*int64, bool)`

GetAccessVlanPortCountOk returns a tuple with the AccessVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) SetAccessVlanPortCount(v int64)`

SetAccessVlanPortCount sets AccessVlanPortCount field to given value.

### HasAccessVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) HasAccessVlanPortCount() bool`

HasAccessVlanPortCount returns a boolean if a field has been set.

### GetBorderVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) GetBorderVlanPortCount() int64`

GetBorderVlanPortCount returns the BorderVlanPortCount field if non-nil, zero value otherwise.

### GetBorderVlanPortCountOk

`func (o *NetworkVlanPortInfoRelationship) GetBorderVlanPortCountOk() (*int64, bool)`

GetBorderVlanPortCountOk returns a tuple with the BorderVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) SetBorderVlanPortCount(v int64)`

SetBorderVlanPortCount sets BorderVlanPortCount field to given value.

### HasBorderVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) HasBorderVlanPortCount() bool`

HasBorderVlanPortCount returns a boolean if a field has been set.

### GetCompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) GetCompressedVlanPortCount() string`

GetCompressedVlanPortCount returns the CompressedVlanPortCount field if non-nil, zero value otherwise.

### GetCompressedVlanPortCountOk

`func (o *NetworkVlanPortInfoRelationship) GetCompressedVlanPortCountOk() (*string, bool)`

GetCompressedVlanPortCountOk returns a tuple with the CompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) SetCompressedVlanPortCount(v string)`

SetCompressedVlanPortCount sets CompressedVlanPortCount field to given value.

### HasCompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) HasCompressedVlanPortCount() bool`

HasCompressedVlanPortCount returns a boolean if a field has been set.

### GetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) GetUncompressedVlanPortCount() string`

GetUncompressedVlanPortCount returns the UncompressedVlanPortCount field if non-nil, zero value otherwise.

### GetUncompressedVlanPortCountOk

`func (o *NetworkVlanPortInfoRelationship) GetUncompressedVlanPortCountOk() (*string, bool)`

GetUncompressedVlanPortCountOk returns a tuple with the UncompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) SetUncompressedVlanPortCount(v string)`

SetUncompressedVlanPortCount sets UncompressedVlanPortCount field to given value.

### HasUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoRelationship) HasUncompressedVlanPortCount() bool`

HasUncompressedVlanPortCount returns a boolean if a field has been set.

### GetVlanPortLimit

`func (o *NetworkVlanPortInfoRelationship) GetVlanPortLimit() int64`

GetVlanPortLimit returns the VlanPortLimit field if non-nil, zero value otherwise.

### GetVlanPortLimitOk

`func (o *NetworkVlanPortInfoRelationship) GetVlanPortLimitOk() (*int64, bool)`

GetVlanPortLimitOk returns a tuple with the VlanPortLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortLimit

`func (o *NetworkVlanPortInfoRelationship) SetVlanPortLimit(v int64)`

SetVlanPortLimit sets VlanPortLimit field to given value.

### HasVlanPortLimit

`func (o *NetworkVlanPortInfoRelationship) HasVlanPortLimit() bool`

HasVlanPortLimit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *NetworkVlanPortInfoRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkVlanPortInfoRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkVlanPortInfoRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkVlanPortInfoRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVlanPortInfoRelationship) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVlanPortInfoRelationship) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVlanPortInfoRelationship) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVlanPortInfoRelationship) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVlanPortInfoRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVlanPortInfoRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVlanPortInfoRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVlanPortInfoRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


