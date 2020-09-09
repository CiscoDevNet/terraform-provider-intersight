# EquipmentSharedIoModuleRelationship

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
**ConfigState** | Pointer to **string** | This field identifies the configuration state for this SIOM Unit. | [optional] [readonly] 
**Discovery** | Pointer to **string** | This field identifies the discovery state of SIOM. | [optional] [readonly] 
**MacOfSharedIomAside** | Pointer to **string** | This field identifies the MAC of IOM-A side. | [optional] [readonly] 
**MacOfSharedIomBside** | Pointer to **string** | This field identifies the MAC of IOM-B side. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the SIOM operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this SIOM Unit. | [optional] [readonly] 
**Reachability** | Pointer to **string** | This field identifies the reachability to FI-A and B side. | [optional] [readonly] 
**UsrLbl** | Pointer to **string** | User label configured for the SIOM. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the vendor id for this SIOM Unit. | [optional] [readonly] 
**Controller** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**EquipmentSystemIoControllerRelationship**](equipment.SystemIoController.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroups** | Pointer to [**[]PortGroupRelationship**](port.Group.Relationship.md) | An array of relationships to portGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentSharedIoModuleRelationship

`func NewEquipmentSharedIoModuleRelationship(classId string, objectType string, ) *EquipmentSharedIoModuleRelationship`

NewEquipmentSharedIoModuleRelationship instantiates a new EquipmentSharedIoModuleRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSharedIoModuleRelationshipWithDefaults

`func NewEquipmentSharedIoModuleRelationshipWithDefaults() *EquipmentSharedIoModuleRelationship`

NewEquipmentSharedIoModuleRelationshipWithDefaults instantiates a new EquipmentSharedIoModuleRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *EquipmentSharedIoModuleRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *EquipmentSharedIoModuleRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *EquipmentSharedIoModuleRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *EquipmentSharedIoModuleRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *EquipmentSharedIoModuleRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSharedIoModuleRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSharedIoModuleRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *EquipmentSharedIoModuleRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EquipmentSharedIoModuleRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EquipmentSharedIoModuleRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EquipmentSharedIoModuleRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *EquipmentSharedIoModuleRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *EquipmentSharedIoModuleRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *EquipmentSharedIoModuleRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *EquipmentSharedIoModuleRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *EquipmentSharedIoModuleRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *EquipmentSharedIoModuleRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *EquipmentSharedIoModuleRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *EquipmentSharedIoModuleRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *EquipmentSharedIoModuleRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *EquipmentSharedIoModuleRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *EquipmentSharedIoModuleRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *EquipmentSharedIoModuleRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *EquipmentSharedIoModuleRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSharedIoModuleRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSharedIoModuleRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *EquipmentSharedIoModuleRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *EquipmentSharedIoModuleRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *EquipmentSharedIoModuleRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *EquipmentSharedIoModuleRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *EquipmentSharedIoModuleRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *EquipmentSharedIoModuleRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *EquipmentSharedIoModuleRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *EquipmentSharedIoModuleRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *EquipmentSharedIoModuleRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EquipmentSharedIoModuleRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EquipmentSharedIoModuleRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EquipmentSharedIoModuleRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *EquipmentSharedIoModuleRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *EquipmentSharedIoModuleRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *EquipmentSharedIoModuleRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *EquipmentSharedIoModuleRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *EquipmentSharedIoModuleRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *EquipmentSharedIoModuleRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *EquipmentSharedIoModuleRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *EquipmentSharedIoModuleRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *EquipmentSharedIoModuleRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *EquipmentSharedIoModuleRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *EquipmentSharedIoModuleRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EquipmentSharedIoModuleRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EquipmentSharedIoModuleRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EquipmentSharedIoModuleRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *EquipmentSharedIoModuleRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *EquipmentSharedIoModuleRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *EquipmentSharedIoModuleRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *EquipmentSharedIoModuleRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *EquipmentSharedIoModuleRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *EquipmentSharedIoModuleRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *EquipmentSharedIoModuleRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *EquipmentSharedIoModuleRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *EquipmentSharedIoModuleRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *EquipmentSharedIoModuleRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *EquipmentSharedIoModuleRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *EquipmentSharedIoModuleRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *EquipmentSharedIoModuleRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EquipmentSharedIoModuleRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EquipmentSharedIoModuleRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EquipmentSharedIoModuleRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *EquipmentSharedIoModuleRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *EquipmentSharedIoModuleRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *EquipmentSharedIoModuleRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *EquipmentSharedIoModuleRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *EquipmentSharedIoModuleRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *EquipmentSharedIoModuleRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *EquipmentSharedIoModuleRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *EquipmentSharedIoModuleRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentSharedIoModuleRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentSharedIoModuleRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentSharedIoModuleRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentSharedIoModuleRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *EquipmentSharedIoModuleRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EquipmentSharedIoModuleRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EquipmentSharedIoModuleRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EquipmentSharedIoModuleRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentSharedIoModuleRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentSharedIoModuleRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentSharedIoModuleRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentSharedIoModuleRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentSharedIoModuleRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentSharedIoModuleRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentSharedIoModuleRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentSharedIoModuleRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetConfigState

`func (o *EquipmentSharedIoModuleRelationship) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentSharedIoModuleRelationship) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentSharedIoModuleRelationship) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentSharedIoModuleRelationship) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDiscovery

`func (o *EquipmentSharedIoModuleRelationship) GetDiscovery() string`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *EquipmentSharedIoModuleRelationship) GetDiscoveryOk() (*string, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *EquipmentSharedIoModuleRelationship) SetDiscovery(v string)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *EquipmentSharedIoModuleRelationship) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleRelationship) GetMacOfSharedIomAside() string`

GetMacOfSharedIomAside returns the MacOfSharedIomAside field if non-nil, zero value otherwise.

### GetMacOfSharedIomAsideOk

`func (o *EquipmentSharedIoModuleRelationship) GetMacOfSharedIomAsideOk() (*string, bool)`

GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleRelationship) SetMacOfSharedIomAside(v string)`

SetMacOfSharedIomAside sets MacOfSharedIomAside field to given value.

### HasMacOfSharedIomAside

`func (o *EquipmentSharedIoModuleRelationship) HasMacOfSharedIomAside() bool`

HasMacOfSharedIomAside returns a boolean if a field has been set.

### GetMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleRelationship) GetMacOfSharedIomBside() string`

GetMacOfSharedIomBside returns the MacOfSharedIomBside field if non-nil, zero value otherwise.

### GetMacOfSharedIomBsideOk

`func (o *EquipmentSharedIoModuleRelationship) GetMacOfSharedIomBsideOk() (*string, bool)`

GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleRelationship) SetMacOfSharedIomBside(v string)`

SetMacOfSharedIomBside sets MacOfSharedIomBside field to given value.

### HasMacOfSharedIomBside

`func (o *EquipmentSharedIoModuleRelationship) HasMacOfSharedIomBside() bool`

HasMacOfSharedIomBside returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentSharedIoModuleRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSharedIoModuleRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSharedIoModuleRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSharedIoModuleRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSharedIoModuleRelationship) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSharedIoModuleRelationship) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSharedIoModuleRelationship) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSharedIoModuleRelationship) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetReachability

`func (o *EquipmentSharedIoModuleRelationship) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *EquipmentSharedIoModuleRelationship) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *EquipmentSharedIoModuleRelationship) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *EquipmentSharedIoModuleRelationship) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetUsrLbl

`func (o *EquipmentSharedIoModuleRelationship) GetUsrLbl() string`

GetUsrLbl returns the UsrLbl field if non-nil, zero value otherwise.

### GetUsrLblOk

`func (o *EquipmentSharedIoModuleRelationship) GetUsrLblOk() (*string, bool)`

GetUsrLblOk returns a tuple with the UsrLbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsrLbl

`func (o *EquipmentSharedIoModuleRelationship) SetUsrLbl(v string)`

SetUsrLbl sets UsrLbl field to given value.

### HasUsrLbl

`func (o *EquipmentSharedIoModuleRelationship) HasUsrLbl() bool`

HasUsrLbl returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentSharedIoModuleRelationship) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentSharedIoModuleRelationship) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentSharedIoModuleRelationship) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentSharedIoModuleRelationship) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetController

`func (o *EquipmentSharedIoModuleRelationship) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *EquipmentSharedIoModuleRelationship) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *EquipmentSharedIoModuleRelationship) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *EquipmentSharedIoModuleRelationship) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleRelationship) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship`

GetEquipmentSystemIoController returns the EquipmentSystemIoController field if non-nil, zero value otherwise.

### GetEquipmentSystemIoControllerOk

`func (o *EquipmentSharedIoModuleRelationship) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool)`

GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleRelationship) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship)`

SetEquipmentSystemIoController sets EquipmentSystemIoController field to given value.

### HasEquipmentSystemIoController

`func (o *EquipmentSharedIoModuleRelationship) HasEquipmentSystemIoController() bool`

HasEquipmentSystemIoController returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSharedIoModuleRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSharedIoModuleRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroups

`func (o *EquipmentSharedIoModuleRelationship) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSharedIoModuleRelationship) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSharedIoModuleRelationship) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSharedIoModuleRelationship) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSharedIoModuleRelationship) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSharedIoModuleRelationship) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSharedIoModuleRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSharedIoModuleRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSharedIoModuleRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSharedIoModuleRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


