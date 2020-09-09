# AdapterUnitRelationship

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
**AdapterId** | Pointer to **string** | Unique Identifier of an adapter Unit within a Rack Interface. | [optional] [readonly] 
**BaseMacAddress** | Pointer to **string** | Original Base Mac address of an adapter unit. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connectivity Status of adapter - A or B or AB. | [optional] [readonly] 
**Integrated** | Pointer to **string** | Cisco Integrated adapter or other type. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of an adapter unit. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of an adapter unit. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number of an adapter unit. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | PCIe slot of the adapter in the server. | [optional] [readonly] 
**Power** | Pointer to **string** | Power state of an adapter unit. | [optional] [readonly] 
**Presence** | Pointer to **string** | Adapter Unit presence or absence. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of an adapter unit. | [optional] [readonly] 
**Vid** | Pointer to **string** | Virtual Id of the adapter in the server. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**Controller** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**ExtEthIfs** | Pointer to [**[]AdapterExtEthInterfaceRelationship**](adapter.ExtEthInterface.Relationship.md) | An array of relationships to adapterExtEthInterface resources. | [optional] [readonly] 
**HostEthIfs** | Pointer to [**[]AdapterHostEthInterfaceRelationship**](adapter.HostEthInterface.Relationship.md) | An array of relationships to adapterHostEthInterface resources. | [optional] [readonly] 
**HostFcIfs** | Pointer to [**[]AdapterHostFcInterfaceRelationship**](adapter.HostFcInterface.Relationship.md) | An array of relationships to adapterHostFcInterface resources. | [optional] [readonly] 
**HostIscsiIfs** | Pointer to [**[]AdapterHostIscsiInterfaceRelationship**](adapter.HostIscsiInterface.Relationship.md) | An array of relationships to adapterHostIscsiInterface resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterUnitRelationship

`func NewAdapterUnitRelationship(classId string, objectType string, ) *AdapterUnitRelationship`

NewAdapterUnitRelationship instantiates a new AdapterUnitRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterUnitRelationshipWithDefaults

`func NewAdapterUnitRelationshipWithDefaults() *AdapterUnitRelationship`

NewAdapterUnitRelationshipWithDefaults instantiates a new AdapterUnitRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *AdapterUnitRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *AdapterUnitRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *AdapterUnitRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *AdapterUnitRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *AdapterUnitRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterUnitRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterUnitRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *AdapterUnitRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AdapterUnitRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AdapterUnitRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AdapterUnitRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *AdapterUnitRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *AdapterUnitRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *AdapterUnitRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *AdapterUnitRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *AdapterUnitRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *AdapterUnitRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *AdapterUnitRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *AdapterUnitRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *AdapterUnitRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *AdapterUnitRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *AdapterUnitRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *AdapterUnitRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *AdapterUnitRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterUnitRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterUnitRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *AdapterUnitRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AdapterUnitRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AdapterUnitRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *AdapterUnitRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *AdapterUnitRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *AdapterUnitRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *AdapterUnitRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *AdapterUnitRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *AdapterUnitRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdapterUnitRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdapterUnitRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdapterUnitRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *AdapterUnitRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *AdapterUnitRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *AdapterUnitRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *AdapterUnitRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *AdapterUnitRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *AdapterUnitRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *AdapterUnitRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *AdapterUnitRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *AdapterUnitRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *AdapterUnitRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *AdapterUnitRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AdapterUnitRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AdapterUnitRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AdapterUnitRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *AdapterUnitRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *AdapterUnitRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *AdapterUnitRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *AdapterUnitRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *AdapterUnitRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *AdapterUnitRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *AdapterUnitRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *AdapterUnitRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *AdapterUnitRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *AdapterUnitRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *AdapterUnitRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *AdapterUnitRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *AdapterUnitRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *AdapterUnitRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *AdapterUnitRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *AdapterUnitRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *AdapterUnitRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *AdapterUnitRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *AdapterUnitRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *AdapterUnitRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *AdapterUnitRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *AdapterUnitRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *AdapterUnitRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *AdapterUnitRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *AdapterUnitRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AdapterUnitRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AdapterUnitRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AdapterUnitRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *AdapterUnitRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AdapterUnitRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AdapterUnitRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AdapterUnitRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *AdapterUnitRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AdapterUnitRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AdapterUnitRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AdapterUnitRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *AdapterUnitRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdapterUnitRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdapterUnitRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdapterUnitRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdapterId

`func (o *AdapterUnitRelationship) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *AdapterUnitRelationship) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *AdapterUnitRelationship) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *AdapterUnitRelationship) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetBaseMacAddress

`func (o *AdapterUnitRelationship) GetBaseMacAddress() string`

GetBaseMacAddress returns the BaseMacAddress field if non-nil, zero value otherwise.

### GetBaseMacAddressOk

`func (o *AdapterUnitRelationship) GetBaseMacAddressOk() (*string, bool)`

GetBaseMacAddressOk returns a tuple with the BaseMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMacAddress

`func (o *AdapterUnitRelationship) SetBaseMacAddress(v string)`

SetBaseMacAddress sets BaseMacAddress field to given value.

### HasBaseMacAddress

`func (o *AdapterUnitRelationship) HasBaseMacAddress() bool`

HasBaseMacAddress returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AdapterUnitRelationship) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AdapterUnitRelationship) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AdapterUnitRelationship) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AdapterUnitRelationship) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetIntegrated

`func (o *AdapterUnitRelationship) GetIntegrated() string`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *AdapterUnitRelationship) GetIntegratedOk() (*string, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *AdapterUnitRelationship) SetIntegrated(v string)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *AdapterUnitRelationship) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterUnitRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterUnitRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterUnitRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterUnitRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterUnitRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterUnitRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterUnitRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterUnitRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPartNumber

`func (o *AdapterUnitRelationship) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *AdapterUnitRelationship) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *AdapterUnitRelationship) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *AdapterUnitRelationship) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *AdapterUnitRelationship) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *AdapterUnitRelationship) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *AdapterUnitRelationship) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *AdapterUnitRelationship) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPower

`func (o *AdapterUnitRelationship) GetPower() string`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *AdapterUnitRelationship) GetPowerOk() (*string, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *AdapterUnitRelationship) SetPower(v string)`

SetPower sets Power field to given value.

### HasPower

`func (o *AdapterUnitRelationship) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPresence

`func (o *AdapterUnitRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *AdapterUnitRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *AdapterUnitRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *AdapterUnitRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetThermal

`func (o *AdapterUnitRelationship) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *AdapterUnitRelationship) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *AdapterUnitRelationship) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *AdapterUnitRelationship) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetVid

`func (o *AdapterUnitRelationship) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *AdapterUnitRelationship) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *AdapterUnitRelationship) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *AdapterUnitRelationship) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetComputeBlade

`func (o *AdapterUnitRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *AdapterUnitRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *AdapterUnitRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *AdapterUnitRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *AdapterUnitRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *AdapterUnitRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *AdapterUnitRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *AdapterUnitRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetController

`func (o *AdapterUnitRelationship) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AdapterUnitRelationship) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AdapterUnitRelationship) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *AdapterUnitRelationship) HasController() bool`

HasController returns a boolean if a field has been set.

### GetExtEthIfs

`func (o *AdapterUnitRelationship) GetExtEthIfs() []AdapterExtEthInterfaceRelationship`

GetExtEthIfs returns the ExtEthIfs field if non-nil, zero value otherwise.

### GetExtEthIfsOk

`func (o *AdapterUnitRelationship) GetExtEthIfsOk() (*[]AdapterExtEthInterfaceRelationship, bool)`

GetExtEthIfsOk returns a tuple with the ExtEthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtEthIfs

`func (o *AdapterUnitRelationship) SetExtEthIfs(v []AdapterExtEthInterfaceRelationship)`

SetExtEthIfs sets ExtEthIfs field to given value.

### HasExtEthIfs

`func (o *AdapterUnitRelationship) HasExtEthIfs() bool`

HasExtEthIfs returns a boolean if a field has been set.

### SetExtEthIfsNil

`func (o *AdapterUnitRelationship) SetExtEthIfsNil(b bool)`

 SetExtEthIfsNil sets the value for ExtEthIfs to be an explicit nil

### UnsetExtEthIfs
`func (o *AdapterUnitRelationship) UnsetExtEthIfs()`

UnsetExtEthIfs ensures that no value is present for ExtEthIfs, not even an explicit nil
### GetHostEthIfs

`func (o *AdapterUnitRelationship) GetHostEthIfs() []AdapterHostEthInterfaceRelationship`

GetHostEthIfs returns the HostEthIfs field if non-nil, zero value otherwise.

### GetHostEthIfsOk

`func (o *AdapterUnitRelationship) GetHostEthIfsOk() (*[]AdapterHostEthInterfaceRelationship, bool)`

GetHostEthIfsOk returns a tuple with the HostEthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthIfs

`func (o *AdapterUnitRelationship) SetHostEthIfs(v []AdapterHostEthInterfaceRelationship)`

SetHostEthIfs sets HostEthIfs field to given value.

### HasHostEthIfs

`func (o *AdapterUnitRelationship) HasHostEthIfs() bool`

HasHostEthIfs returns a boolean if a field has been set.

### SetHostEthIfsNil

`func (o *AdapterUnitRelationship) SetHostEthIfsNil(b bool)`

 SetHostEthIfsNil sets the value for HostEthIfs to be an explicit nil

### UnsetHostEthIfs
`func (o *AdapterUnitRelationship) UnsetHostEthIfs()`

UnsetHostEthIfs ensures that no value is present for HostEthIfs, not even an explicit nil
### GetHostFcIfs

`func (o *AdapterUnitRelationship) GetHostFcIfs() []AdapterHostFcInterfaceRelationship`

GetHostFcIfs returns the HostFcIfs field if non-nil, zero value otherwise.

### GetHostFcIfsOk

`func (o *AdapterUnitRelationship) GetHostFcIfsOk() (*[]AdapterHostFcInterfaceRelationship, bool)`

GetHostFcIfsOk returns a tuple with the HostFcIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFcIfs

`func (o *AdapterUnitRelationship) SetHostFcIfs(v []AdapterHostFcInterfaceRelationship)`

SetHostFcIfs sets HostFcIfs field to given value.

### HasHostFcIfs

`func (o *AdapterUnitRelationship) HasHostFcIfs() bool`

HasHostFcIfs returns a boolean if a field has been set.

### SetHostFcIfsNil

`func (o *AdapterUnitRelationship) SetHostFcIfsNil(b bool)`

 SetHostFcIfsNil sets the value for HostFcIfs to be an explicit nil

### UnsetHostFcIfs
`func (o *AdapterUnitRelationship) UnsetHostFcIfs()`

UnsetHostFcIfs ensures that no value is present for HostFcIfs, not even an explicit nil
### GetHostIscsiIfs

`func (o *AdapterUnitRelationship) GetHostIscsiIfs() []AdapterHostIscsiInterfaceRelationship`

GetHostIscsiIfs returns the HostIscsiIfs field if non-nil, zero value otherwise.

### GetHostIscsiIfsOk

`func (o *AdapterUnitRelationship) GetHostIscsiIfsOk() (*[]AdapterHostIscsiInterfaceRelationship, bool)`

GetHostIscsiIfsOk returns a tuple with the HostIscsiIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIscsiIfs

`func (o *AdapterUnitRelationship) SetHostIscsiIfs(v []AdapterHostIscsiInterfaceRelationship)`

SetHostIscsiIfs sets HostIscsiIfs field to given value.

### HasHostIscsiIfs

`func (o *AdapterUnitRelationship) HasHostIscsiIfs() bool`

HasHostIscsiIfs returns a boolean if a field has been set.

### SetHostIscsiIfsNil

`func (o *AdapterUnitRelationship) SetHostIscsiIfsNil(b bool)`

 SetHostIscsiIfsNil sets the value for HostIscsiIfs to be an explicit nil

### UnsetHostIscsiIfs
`func (o *AdapterUnitRelationship) UnsetHostIscsiIfs()`

UnsetHostIscsiIfs ensures that no value is present for HostIscsiIfs, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *AdapterUnitRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterUnitRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterUnitRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterUnitRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterUnitRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterUnitRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterUnitRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterUnitRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


