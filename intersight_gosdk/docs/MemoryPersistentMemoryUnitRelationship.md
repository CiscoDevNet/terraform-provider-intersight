# MemoryPersistentMemoryUnitRelationship

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
**Model** | Pointer to **string** | This field identifies the model of the given component. | [optional] [readonly] 
**Revision** | Pointer to **string** | This field identifies the revision of the given component. | [optional] [readonly] 
**Serial** | Pointer to **string** | This field identifies the serial of the given component. | [optional] [readonly] 
**Vendor** | Pointer to **string** | This field identifies the vendor of the given component. | [optional] [readonly] 
**AdminState** | Pointer to **string** | This represents the administrative state of the memory unit on a server. | [optional] [readonly] 
**ArrayId** | Pointer to **int64** | This represents the memory array to which the memory unit belongs to. | [optional] [readonly] 
**Bank** | Pointer to **int64** | This represents the memory bank of the memory unit on a server. | [optional] [readonly] 
**Capacity** | Pointer to **string** | This represents the memory capacity in MiB of the memory unit on a server. | [optional] [readonly] 
**Clock** | Pointer to **string** | This represents the clock of the memory unit on a server. | [optional] [readonly] 
**FormFactor** | Pointer to **string** | This represents the form factor of the memory unit on a server. | [optional] [readonly] 
**Latency** | Pointer to **string** | This represents the latency of the memory unit on a server. | [optional] [readonly] 
**Location** | Pointer to **string** | This represents the location of the memory unit on a server. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | This represents the operational power state of the memory unit on a server. | [optional] [readonly] 
**OperState** | Pointer to **string** | This represents the operational state of the memory unit on a server. | [optional] [readonly] 
**Operability** | Pointer to **string** | This represents the operability of the memory unit on a server. | [optional] [readonly] 
**Presence** | Pointer to **string** | This represents the presence state of the memory unit on a server. | [optional] [readonly] 
**Set** | Pointer to **int64** | This represents the set of the memory unit on a server. | [optional] [readonly] 
**Speed** | Pointer to **string** | This represents the speed of the memory unit on a server. | [optional] [readonly] 
**Thermal** | Pointer to **string** | This represents the thremal state of the memory unit on a server. | [optional] [readonly] 
**Type** | Pointer to **string** | This represents the memory type of the memory unit on a server. | [optional] [readonly] 
**Visibility** | Pointer to **string** | This represents the visibility of the memory unit on a server. | [optional] [readonly] 
**Width** | Pointer to **string** | This represents the width of the memory unit on a server. | [optional] [readonly] 
**AppDirectCapacity** | Pointer to **string** | AppDirect capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**CountStatus** | Pointer to **string** | Count status of the Persistent Memory Module on a server. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Firmware version of the firware running on the Persistent Memory Module on a server. | [optional] [readonly] 
**FrozenStatus** | Pointer to **string** | Frozen status of the Persistent Memory Module on a server. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Module on a server. | [optional] [readonly] 
**LockStatus** | Pointer to **string** | Lock status of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryCapacity** | Pointer to **string** | Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryId** | Pointer to **int64** | ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**PersistentMemoryCapacity** | Pointer to **string** | Persistent Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**ReservedCapacity** | Pointer to **string** | Reserved capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**SecurityStatus** | Pointer to **string** | Security status of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**Uid** | Pointer to **string** | UID of the Persistent Memory Module on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryArray** | Pointer to [**MemoryArrayRelationship**](memory.Array.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryUnitRelationship

`func NewMemoryPersistentMemoryUnitRelationship(classId string, objectType string, ) *MemoryPersistentMemoryUnitRelationship`

NewMemoryPersistentMemoryUnitRelationship instantiates a new MemoryPersistentMemoryUnitRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryUnitRelationshipWithDefaults

`func NewMemoryPersistentMemoryUnitRelationshipWithDefaults() *MemoryPersistentMemoryUnitRelationship`

NewMemoryPersistentMemoryUnitRelationshipWithDefaults instantiates a new MemoryPersistentMemoryUnitRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *MemoryPersistentMemoryUnitRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *MemoryPersistentMemoryUnitRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *MemoryPersistentMemoryUnitRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *MemoryPersistentMemoryUnitRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryUnitRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *MemoryPersistentMemoryUnitRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MemoryPersistentMemoryUnitRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MemoryPersistentMemoryUnitRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *MemoryPersistentMemoryUnitRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *MemoryPersistentMemoryUnitRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *MemoryPersistentMemoryUnitRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *MemoryPersistentMemoryUnitRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *MemoryPersistentMemoryUnitRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *MemoryPersistentMemoryUnitRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *MemoryPersistentMemoryUnitRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *MemoryPersistentMemoryUnitRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *MemoryPersistentMemoryUnitRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *MemoryPersistentMemoryUnitRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryUnitRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *MemoryPersistentMemoryUnitRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MemoryPersistentMemoryUnitRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MemoryPersistentMemoryUnitRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *MemoryPersistentMemoryUnitRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *MemoryPersistentMemoryUnitRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *MemoryPersistentMemoryUnitRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *MemoryPersistentMemoryUnitRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemoryPersistentMemoryUnitRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemoryPersistentMemoryUnitRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *MemoryPersistentMemoryUnitRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *MemoryPersistentMemoryUnitRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *MemoryPersistentMemoryUnitRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *MemoryPersistentMemoryUnitRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *MemoryPersistentMemoryUnitRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *MemoryPersistentMemoryUnitRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *MemoryPersistentMemoryUnitRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *MemoryPersistentMemoryUnitRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *MemoryPersistentMemoryUnitRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MemoryPersistentMemoryUnitRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MemoryPersistentMemoryUnitRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *MemoryPersistentMemoryUnitRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *MemoryPersistentMemoryUnitRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *MemoryPersistentMemoryUnitRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *MemoryPersistentMemoryUnitRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *MemoryPersistentMemoryUnitRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *MemoryPersistentMemoryUnitRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *MemoryPersistentMemoryUnitRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *MemoryPersistentMemoryUnitRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *MemoryPersistentMemoryUnitRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *MemoryPersistentMemoryUnitRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *MemoryPersistentMemoryUnitRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *MemoryPersistentMemoryUnitRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *MemoryPersistentMemoryUnitRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *MemoryPersistentMemoryUnitRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *MemoryPersistentMemoryUnitRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *MemoryPersistentMemoryUnitRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *MemoryPersistentMemoryUnitRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *MemoryPersistentMemoryUnitRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *MemoryPersistentMemoryUnitRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *MemoryPersistentMemoryUnitRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MemoryPersistentMemoryUnitRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MemoryPersistentMemoryUnitRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *MemoryPersistentMemoryUnitRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *MemoryPersistentMemoryUnitRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *MemoryPersistentMemoryUnitRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *MemoryPersistentMemoryUnitRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *MemoryPersistentMemoryUnitRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *MemoryPersistentMemoryUnitRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *MemoryPersistentMemoryUnitRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MemoryPersistentMemoryUnitRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MemoryPersistentMemoryUnitRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdminState

`func (o *MemoryPersistentMemoryUnitRelationship) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *MemoryPersistentMemoryUnitRelationship) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *MemoryPersistentMemoryUnitRelationship) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetArrayId

`func (o *MemoryPersistentMemoryUnitRelationship) GetArrayId() int64`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetArrayIdOk() (*int64, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *MemoryPersistentMemoryUnitRelationship) SetArrayId(v int64)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *MemoryPersistentMemoryUnitRelationship) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetBank

`func (o *MemoryPersistentMemoryUnitRelationship) GetBank() int64`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetBankOk() (*int64, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *MemoryPersistentMemoryUnitRelationship) SetBank(v int64)`

SetBank sets Bank field to given value.

### HasBank

`func (o *MemoryPersistentMemoryUnitRelationship) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClock

`func (o *MemoryPersistentMemoryUnitRelationship) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *MemoryPersistentMemoryUnitRelationship) SetClock(v string)`

SetClock sets Clock field to given value.

### HasClock

`func (o *MemoryPersistentMemoryUnitRelationship) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetFormFactor

`func (o *MemoryPersistentMemoryUnitRelationship) GetFormFactor() string`

GetFormFactor returns the FormFactor field if non-nil, zero value otherwise.

### GetFormFactorOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetFormFactorOk() (*string, bool)`

GetFormFactorOk returns a tuple with the FormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFactor

`func (o *MemoryPersistentMemoryUnitRelationship) SetFormFactor(v string)`

SetFormFactor sets FormFactor field to given value.

### HasFormFactor

`func (o *MemoryPersistentMemoryUnitRelationship) HasFormFactor() bool`

HasFormFactor returns a boolean if a field has been set.

### GetLatency

`func (o *MemoryPersistentMemoryUnitRelationship) GetLatency() string`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetLatencyOk() (*string, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MemoryPersistentMemoryUnitRelationship) SetLatency(v string)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MemoryPersistentMemoryUnitRelationship) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLocation

`func (o *MemoryPersistentMemoryUnitRelationship) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MemoryPersistentMemoryUnitRelationship) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MemoryPersistentMemoryUnitRelationship) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOperPowerState

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *MemoryPersistentMemoryUnitRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *MemoryPersistentMemoryUnitRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *MemoryPersistentMemoryUnitRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *MemoryPersistentMemoryUnitRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *MemoryPersistentMemoryUnitRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *MemoryPersistentMemoryUnitRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPresence

`func (o *MemoryPersistentMemoryUnitRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *MemoryPersistentMemoryUnitRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *MemoryPersistentMemoryUnitRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSet

`func (o *MemoryPersistentMemoryUnitRelationship) GetSet() int64`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSetOk() (*int64, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *MemoryPersistentMemoryUnitRelationship) SetSet(v int64)`

SetSet sets Set field to given value.

### HasSet

`func (o *MemoryPersistentMemoryUnitRelationship) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetSpeed

`func (o *MemoryPersistentMemoryUnitRelationship) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MemoryPersistentMemoryUnitRelationship) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MemoryPersistentMemoryUnitRelationship) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetThermal

`func (o *MemoryPersistentMemoryUnitRelationship) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *MemoryPersistentMemoryUnitRelationship) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *MemoryPersistentMemoryUnitRelationship) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetType

`func (o *MemoryPersistentMemoryUnitRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemoryPersistentMemoryUnitRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemoryPersistentMemoryUnitRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *MemoryPersistentMemoryUnitRelationship) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MemoryPersistentMemoryUnitRelationship) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MemoryPersistentMemoryUnitRelationship) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWidth

`func (o *MemoryPersistentMemoryUnitRelationship) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MemoryPersistentMemoryUnitRelationship) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MemoryPersistentMemoryUnitRelationship) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetAppDirectCapacity() string`

GetAppDirectCapacity returns the AppDirectCapacity field if non-nil, zero value otherwise.

### GetAppDirectCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetAppDirectCapacityOk() (*string, bool)`

GetAppDirectCapacityOk returns a tuple with the AppDirectCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetAppDirectCapacity(v string)`

SetAppDirectCapacity sets AppDirectCapacity field to given value.

### HasAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasAppDirectCapacity() bool`

HasAppDirectCapacity returns a boolean if a field has been set.

### GetCountStatus

`func (o *MemoryPersistentMemoryUnitRelationship) GetCountStatus() string`

GetCountStatus returns the CountStatus field if non-nil, zero value otherwise.

### GetCountStatusOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetCountStatusOk() (*string, bool)`

GetCountStatusOk returns a tuple with the CountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountStatus

`func (o *MemoryPersistentMemoryUnitRelationship) SetCountStatus(v string)`

SetCountStatus sets CountStatus field to given value.

### HasCountStatus

`func (o *MemoryPersistentMemoryUnitRelationship) HasCountStatus() bool`

HasCountStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *MemoryPersistentMemoryUnitRelationship) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *MemoryPersistentMemoryUnitRelationship) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *MemoryPersistentMemoryUnitRelationship) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFrozenStatus

`func (o *MemoryPersistentMemoryUnitRelationship) GetFrozenStatus() string`

GetFrozenStatus returns the FrozenStatus field if non-nil, zero value otherwise.

### GetFrozenStatusOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetFrozenStatusOk() (*string, bool)`

GetFrozenStatusOk returns a tuple with the FrozenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenStatus

`func (o *MemoryPersistentMemoryUnitRelationship) SetFrozenStatus(v string)`

SetFrozenStatus sets FrozenStatus field to given value.

### HasFrozenStatus

`func (o *MemoryPersistentMemoryUnitRelationship) HasFrozenStatus() bool`

HasFrozenStatus returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryUnitRelationship) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryUnitRelationship) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryUnitRelationship) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetLockStatus

`func (o *MemoryPersistentMemoryUnitRelationship) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *MemoryPersistentMemoryUnitRelationship) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *MemoryPersistentMemoryUnitRelationship) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryCapacity() string`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryCapacityOk() (*string, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetMemoryCapacity(v string)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryId() int64`

GetMemoryId returns the MemoryId field if non-nil, zero value otherwise.

### GetMemoryIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryIdOk() (*int64, bool)`

GetMemoryIdOk returns a tuple with the MemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) SetMemoryId(v int64)`

SetMemoryId sets MemoryId field to given value.

### HasMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) HasMemoryId() bool`

HasMemoryId returns a boolean if a field has been set.

### GetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetPersistentMemoryCapacity() string`

GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field if non-nil, zero value otherwise.

### GetPersistentMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetPersistentMemoryCapacityOk() (*string, bool)`

GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetPersistentMemoryCapacity(v string)`

SetPersistentMemoryCapacity sets PersistentMemoryCapacity field to given value.

### HasPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasPersistentMemoryCapacity() bool`

HasPersistentMemoryCapacity returns a boolean if a field has been set.

### GetReservedCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetReservedCapacity() string`

GetReservedCapacity returns the ReservedCapacity field if non-nil, zero value otherwise.

### GetReservedCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetReservedCapacityOk() (*string, bool)`

GetReservedCapacityOk returns a tuple with the ReservedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetReservedCapacity(v string)`

SetReservedCapacity sets ReservedCapacity field to given value.

### HasReservedCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasReservedCapacity() bool`

HasReservedCapacity returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *MemoryPersistentMemoryUnitRelationship) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *MemoryPersistentMemoryUnitRelationship) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *MemoryPersistentMemoryUnitRelationship) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryUnitRelationship) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryUnitRelationship) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryUnitRelationship) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryUnitRelationship) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryUnitRelationship) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUid

`func (o *MemoryPersistentMemoryUnitRelationship) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MemoryPersistentMemoryUnitRelationship) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MemoryPersistentMemoryUnitRelationship) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryArray

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryArray() MemoryArrayRelationship`

GetMemoryArray returns the MemoryArray field if non-nil, zero value otherwise.

### GetMemoryArrayOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetMemoryArrayOk() (*MemoryArrayRelationship, bool)`

GetMemoryArrayOk returns a tuple with the MemoryArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArray

`func (o *MemoryPersistentMemoryUnitRelationship) SetMemoryArray(v MemoryArrayRelationship)`

SetMemoryArray sets MemoryArray field to given value.

### HasMemoryArray

`func (o *MemoryPersistentMemoryUnitRelationship) HasMemoryArray() bool`

HasMemoryArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryPersistentMemoryUnitRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryUnitRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryUnitRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryUnitRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


