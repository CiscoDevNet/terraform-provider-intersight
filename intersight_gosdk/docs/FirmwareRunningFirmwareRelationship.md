# FirmwareRunningFirmwareRelationship

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
**Component** | Pointer to **string** | Kind of the firmware - boot-booloader/system/kernel. | [optional] [readonly] 
**PackageVersion** | Pointer to **string** | Package version which the firmware belongs to. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the firmware. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the firmware. | [optional] [readonly] 
**BiosUnit** | Pointer to [**BiosUnitRelationship**](bios.Unit.Relationship.md) |  | [optional] 
**GraphicsCard** | Pointer to [**GraphicsCardRelationship**](graphics.Card.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](network.Element.Relationship.md) | An array of relationships to networkElement resources. | [optional] 
**PciSwitch** | Pointer to [**PciSwitchRelationship**](pci.Switch.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) |  | [optional] 
**StoragePhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareRunningFirmwareRelationship

`func NewFirmwareRunningFirmwareRelationship(classId string, objectType string, ) *FirmwareRunningFirmwareRelationship`

NewFirmwareRunningFirmwareRelationship instantiates a new FirmwareRunningFirmwareRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareRunningFirmwareRelationshipWithDefaults

`func NewFirmwareRunningFirmwareRelationshipWithDefaults() *FirmwareRunningFirmwareRelationship`

NewFirmwareRunningFirmwareRelationshipWithDefaults instantiates a new FirmwareRunningFirmwareRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FirmwareRunningFirmwareRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FirmwareRunningFirmwareRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FirmwareRunningFirmwareRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FirmwareRunningFirmwareRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FirmwareRunningFirmwareRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareRunningFirmwareRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareRunningFirmwareRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FirmwareRunningFirmwareRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FirmwareRunningFirmwareRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FirmwareRunningFirmwareRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FirmwareRunningFirmwareRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FirmwareRunningFirmwareRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FirmwareRunningFirmwareRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FirmwareRunningFirmwareRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FirmwareRunningFirmwareRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FirmwareRunningFirmwareRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FirmwareRunningFirmwareRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FirmwareRunningFirmwareRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FirmwareRunningFirmwareRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FirmwareRunningFirmwareRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FirmwareRunningFirmwareRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FirmwareRunningFirmwareRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FirmwareRunningFirmwareRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FirmwareRunningFirmwareRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareRunningFirmwareRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareRunningFirmwareRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FirmwareRunningFirmwareRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FirmwareRunningFirmwareRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FirmwareRunningFirmwareRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FirmwareRunningFirmwareRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FirmwareRunningFirmwareRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FirmwareRunningFirmwareRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FirmwareRunningFirmwareRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FirmwareRunningFirmwareRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FirmwareRunningFirmwareRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FirmwareRunningFirmwareRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FirmwareRunningFirmwareRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FirmwareRunningFirmwareRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FirmwareRunningFirmwareRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FirmwareRunningFirmwareRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FirmwareRunningFirmwareRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FirmwareRunningFirmwareRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FirmwareRunningFirmwareRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FirmwareRunningFirmwareRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FirmwareRunningFirmwareRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FirmwareRunningFirmwareRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FirmwareRunningFirmwareRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FirmwareRunningFirmwareRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FirmwareRunningFirmwareRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FirmwareRunningFirmwareRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FirmwareRunningFirmwareRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FirmwareRunningFirmwareRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FirmwareRunningFirmwareRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FirmwareRunningFirmwareRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FirmwareRunningFirmwareRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FirmwareRunningFirmwareRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FirmwareRunningFirmwareRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FirmwareRunningFirmwareRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FirmwareRunningFirmwareRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FirmwareRunningFirmwareRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FirmwareRunningFirmwareRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FirmwareRunningFirmwareRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FirmwareRunningFirmwareRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FirmwareRunningFirmwareRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *FirmwareRunningFirmwareRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *FirmwareRunningFirmwareRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *FirmwareRunningFirmwareRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *FirmwareRunningFirmwareRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *FirmwareRunningFirmwareRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *FirmwareRunningFirmwareRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *FirmwareRunningFirmwareRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *FirmwareRunningFirmwareRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *FirmwareRunningFirmwareRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *FirmwareRunningFirmwareRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *FirmwareRunningFirmwareRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *FirmwareRunningFirmwareRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetComponent

`func (o *FirmwareRunningFirmwareRelationship) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FirmwareRunningFirmwareRelationship) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FirmwareRunningFirmwareRelationship) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FirmwareRunningFirmwareRelationship) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPackageVersion

`func (o *FirmwareRunningFirmwareRelationship) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *FirmwareRunningFirmwareRelationship) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *FirmwareRunningFirmwareRelationship) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *FirmwareRunningFirmwareRelationship) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetType

`func (o *FirmwareRunningFirmwareRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirmwareRunningFirmwareRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirmwareRunningFirmwareRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirmwareRunningFirmwareRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *FirmwareRunningFirmwareRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FirmwareRunningFirmwareRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FirmwareRunningFirmwareRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FirmwareRunningFirmwareRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBiosUnit

`func (o *FirmwareRunningFirmwareRelationship) GetBiosUnit() BiosUnitRelationship`

GetBiosUnit returns the BiosUnit field if non-nil, zero value otherwise.

### GetBiosUnitOk

`func (o *FirmwareRunningFirmwareRelationship) GetBiosUnitOk() (*BiosUnitRelationship, bool)`

GetBiosUnitOk returns a tuple with the BiosUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUnit

`func (o *FirmwareRunningFirmwareRelationship) SetBiosUnit(v BiosUnitRelationship)`

SetBiosUnit sets BiosUnit field to given value.

### HasBiosUnit

`func (o *FirmwareRunningFirmwareRelationship) HasBiosUnit() bool`

HasBiosUnit returns a boolean if a field has been set.

### GetGraphicsCard

`func (o *FirmwareRunningFirmwareRelationship) GetGraphicsCard() GraphicsCardRelationship`

GetGraphicsCard returns the GraphicsCard field if non-nil, zero value otherwise.

### GetGraphicsCardOk

`func (o *FirmwareRunningFirmwareRelationship) GetGraphicsCardOk() (*GraphicsCardRelationship, bool)`

GetGraphicsCardOk returns a tuple with the GraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCard

`func (o *FirmwareRunningFirmwareRelationship) SetGraphicsCard(v GraphicsCardRelationship)`

SetGraphicsCard sets GraphicsCard field to given value.

### HasGraphicsCard

`func (o *FirmwareRunningFirmwareRelationship) HasGraphicsCard() bool`

HasGraphicsCard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FirmwareRunningFirmwareRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FirmwareRunningFirmwareRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *FirmwareRunningFirmwareRelationship) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *FirmwareRunningFirmwareRelationship) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *FirmwareRunningFirmwareRelationship) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *FirmwareRunningFirmwareRelationship) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetNetworkElements

`func (o *FirmwareRunningFirmwareRelationship) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareRunningFirmwareRelationship) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareRunningFirmwareRelationship) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareRunningFirmwareRelationship) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareRunningFirmwareRelationship) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareRunningFirmwareRelationship) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetPciSwitch

`func (o *FirmwareRunningFirmwareRelationship) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *FirmwareRunningFirmwareRelationship) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *FirmwareRunningFirmwareRelationship) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *FirmwareRunningFirmwareRelationship) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FirmwareRunningFirmwareRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FirmwareRunningFirmwareRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FirmwareRunningFirmwareRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FirmwareRunningFirmwareRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *FirmwareRunningFirmwareRelationship) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *FirmwareRunningFirmwareRelationship) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *FirmwareRunningFirmwareRelationship) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *FirmwareRunningFirmwareRelationship) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *FirmwareRunningFirmwareRelationship) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *FirmwareRunningFirmwareRelationship) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *FirmwareRunningFirmwareRelationship) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *FirmwareRunningFirmwareRelationship) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.

### GetStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareRelationship) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship`

GetStoragePhysicalDisk returns the StoragePhysicalDisk field if non-nil, zero value otherwise.

### GetStoragePhysicalDiskOk

`func (o *FirmwareRunningFirmwareRelationship) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareRelationship) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship)`

SetStoragePhysicalDisk sets StoragePhysicalDisk field to given value.

### HasStoragePhysicalDisk

`func (o *FirmwareRunningFirmwareRelationship) HasStoragePhysicalDisk() bool`

HasStoragePhysicalDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


