# AdapterHostFcInterfaceRelationship

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
**AdminState** | Pointer to **string** | Admin Configured State of Host Fibre Channel Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Fibre Channel Interface. | [optional] [readonly] 
**HostFcInterfaceId** | Pointer to **int64** | Identifier of Host Fibre Channel Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Fibre Channel Interface. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational State of Host Fibre Channel Interface. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability status of Host Fibre Channel Interface. | [optional] [readonly] 
**OriginalWwnn** | Pointer to **string** | The uniquely distinguishable factory default  World Wide Node Name of the Host. | [optional] [readonly] 
**OriginalWwpn** | Pointer to **string** | The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerPort Dn of Host Fibre Channel Interface. | [optional] [readonly] 
**Wwnn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Node Name of the Host. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterHostFcInterfaceRelationship

`func NewAdapterHostFcInterfaceRelationship(classId string, objectType string, ) *AdapterHostFcInterfaceRelationship`

NewAdapterHostFcInterfaceRelationship instantiates a new AdapterHostFcInterfaceRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostFcInterfaceRelationshipWithDefaults

`func NewAdapterHostFcInterfaceRelationshipWithDefaults() *AdapterHostFcInterfaceRelationship`

NewAdapterHostFcInterfaceRelationshipWithDefaults instantiates a new AdapterHostFcInterfaceRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *AdapterHostFcInterfaceRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AdapterHostFcInterfaceRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AdapterHostFcInterfaceRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AdapterHostFcInterfaceRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *AdapterHostFcInterfaceRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostFcInterfaceRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostFcInterfaceRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *AdapterHostFcInterfaceRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AdapterHostFcInterfaceRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AdapterHostFcInterfaceRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AdapterHostFcInterfaceRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *AdapterHostFcInterfaceRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *AdapterHostFcInterfaceRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *AdapterHostFcInterfaceRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *AdapterHostFcInterfaceRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *AdapterHostFcInterfaceRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *AdapterHostFcInterfaceRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *AdapterHostFcInterfaceRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *AdapterHostFcInterfaceRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *AdapterHostFcInterfaceRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *AdapterHostFcInterfaceRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *AdapterHostFcInterfaceRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *AdapterHostFcInterfaceRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *AdapterHostFcInterfaceRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostFcInterfaceRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostFcInterfaceRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *AdapterHostFcInterfaceRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AdapterHostFcInterfaceRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AdapterHostFcInterfaceRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AdapterHostFcInterfaceRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *AdapterHostFcInterfaceRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *AdapterHostFcInterfaceRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *AdapterHostFcInterfaceRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *AdapterHostFcInterfaceRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *AdapterHostFcInterfaceRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdapterHostFcInterfaceRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdapterHostFcInterfaceRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdapterHostFcInterfaceRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *AdapterHostFcInterfaceRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *AdapterHostFcInterfaceRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *AdapterHostFcInterfaceRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *AdapterHostFcInterfaceRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *AdapterHostFcInterfaceRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *AdapterHostFcInterfaceRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *AdapterHostFcInterfaceRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *AdapterHostFcInterfaceRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *AdapterHostFcInterfaceRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *AdapterHostFcInterfaceRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *AdapterHostFcInterfaceRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AdapterHostFcInterfaceRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AdapterHostFcInterfaceRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AdapterHostFcInterfaceRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *AdapterHostFcInterfaceRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *AdapterHostFcInterfaceRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *AdapterHostFcInterfaceRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *AdapterHostFcInterfaceRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *AdapterHostFcInterfaceRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *AdapterHostFcInterfaceRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *AdapterHostFcInterfaceRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *AdapterHostFcInterfaceRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *AdapterHostFcInterfaceRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *AdapterHostFcInterfaceRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *AdapterHostFcInterfaceRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *AdapterHostFcInterfaceRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *AdapterHostFcInterfaceRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *AdapterHostFcInterfaceRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *AdapterHostFcInterfaceRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *AdapterHostFcInterfaceRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *AdapterHostFcInterfaceRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *AdapterHostFcInterfaceRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *AdapterHostFcInterfaceRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *AdapterHostFcInterfaceRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *AdapterHostFcInterfaceRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *AdapterHostFcInterfaceRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *AdapterHostFcInterfaceRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *AdapterHostFcInterfaceRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *AdapterHostFcInterfaceRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AdapterHostFcInterfaceRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AdapterHostFcInterfaceRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AdapterHostFcInterfaceRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *AdapterHostFcInterfaceRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AdapterHostFcInterfaceRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AdapterHostFcInterfaceRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AdapterHostFcInterfaceRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *AdapterHostFcInterfaceRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AdapterHostFcInterfaceRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AdapterHostFcInterfaceRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AdapterHostFcInterfaceRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *AdapterHostFcInterfaceRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdapterHostFcInterfaceRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdapterHostFcInterfaceRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdapterHostFcInterfaceRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdminState

`func (o *AdapterHostFcInterfaceRelationship) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostFcInterfaceRelationship) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostFcInterfaceRelationship) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostFcInterfaceRelationship) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostFcInterfaceRelationship) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostFcInterfaceRelationship) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostFcInterfaceRelationship) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostFcInterfaceRelationship) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostFcInterfaceId

`func (o *AdapterHostFcInterfaceRelationship) GetHostFcInterfaceId() int64`

GetHostFcInterfaceId returns the HostFcInterfaceId field if non-nil, zero value otherwise.

### GetHostFcInterfaceIdOk

`func (o *AdapterHostFcInterfaceRelationship) GetHostFcInterfaceIdOk() (*int64, bool)`

GetHostFcInterfaceIdOk returns a tuple with the HostFcInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFcInterfaceId

`func (o *AdapterHostFcInterfaceRelationship) SetHostFcInterfaceId(v int64)`

SetHostFcInterfaceId sets HostFcInterfaceId field to given value.

### HasHostFcInterfaceId

`func (o *AdapterHostFcInterfaceRelationship) HasHostFcInterfaceId() bool`

HasHostFcInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostFcInterfaceRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostFcInterfaceRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostFcInterfaceRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostFcInterfaceRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterHostFcInterfaceRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostFcInterfaceRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostFcInterfaceRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostFcInterfaceRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostFcInterfaceRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostFcInterfaceRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostFcInterfaceRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostFcInterfaceRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalWwnn

`func (o *AdapterHostFcInterfaceRelationship) GetOriginalWwnn() string`

GetOriginalWwnn returns the OriginalWwnn field if non-nil, zero value otherwise.

### GetOriginalWwnnOk

`func (o *AdapterHostFcInterfaceRelationship) GetOriginalWwnnOk() (*string, bool)`

GetOriginalWwnnOk returns a tuple with the OriginalWwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwnn

`func (o *AdapterHostFcInterfaceRelationship) SetOriginalWwnn(v string)`

SetOriginalWwnn sets OriginalWwnn field to given value.

### HasOriginalWwnn

`func (o *AdapterHostFcInterfaceRelationship) HasOriginalWwnn() bool`

HasOriginalWwnn returns a boolean if a field has been set.

### GetOriginalWwpn

`func (o *AdapterHostFcInterfaceRelationship) GetOriginalWwpn() string`

GetOriginalWwpn returns the OriginalWwpn field if non-nil, zero value otherwise.

### GetOriginalWwpnOk

`func (o *AdapterHostFcInterfaceRelationship) GetOriginalWwpnOk() (*string, bool)`

GetOriginalWwpnOk returns a tuple with the OriginalWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwpn

`func (o *AdapterHostFcInterfaceRelationship) SetOriginalWwpn(v string)`

SetOriginalWwpn sets OriginalWwpn field to given value.

### HasOriginalWwpn

`func (o *AdapterHostFcInterfaceRelationship) HasOriginalWwpn() bool`

HasOriginalWwpn returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostFcInterfaceRelationship) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostFcInterfaceRelationship) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostFcInterfaceRelationship) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostFcInterfaceRelationship) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetWwnn

`func (o *AdapterHostFcInterfaceRelationship) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *AdapterHostFcInterfaceRelationship) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *AdapterHostFcInterfaceRelationship) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *AdapterHostFcInterfaceRelationship) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetWwpn

`func (o *AdapterHostFcInterfaceRelationship) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *AdapterHostFcInterfaceRelationship) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *AdapterHostFcInterfaceRelationship) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *AdapterHostFcInterfaceRelationship) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostFcInterfaceRelationship) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostFcInterfaceRelationship) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostFcInterfaceRelationship) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostFcInterfaceRelationship) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostFcInterfaceRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostFcInterfaceRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostFcInterfaceRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostFcInterfaceRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostFcInterfaceRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


