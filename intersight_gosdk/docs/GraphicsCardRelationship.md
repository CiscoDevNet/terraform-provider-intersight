# GraphicsCardRelationship

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
**CardId** | Pointer to **int64** | The id of the graphics card. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device id of the graphics card. | [optional] [readonly] 
**ExpanderSlot** | Pointer to **string** | The expander slot information of the card. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | The firmware version of the graphics card. | [optional] [readonly] 
**Mode** | Pointer to **string** | The current mode of the graphics card. | [optional] [readonly] 
**NumGpus** | Pointer to **string** | The number of controllers under each card. | [optional] 
**OperState** | Pointer to **string** | The current operational state of the graphics card. | [optional] [readonly] 
**PciAddress** | Pointer to **string** | The PCI address of the graphics card. | [optional] [readonly] 
**PciAddressList** | Pointer to **string** | This list contains the PCI address of all controllers for corresponding card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name of the graphics card. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**GraphicsControllers** | Pointer to [**[]GraphicsControllerRelationship**](graphics.Controller.Relationship.md) | An array of relationships to graphicsController resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewGraphicsCardRelationship

`func NewGraphicsCardRelationship(classId string, objectType string, ) *GraphicsCardRelationship`

NewGraphicsCardRelationship instantiates a new GraphicsCardRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphicsCardRelationshipWithDefaults

`func NewGraphicsCardRelationshipWithDefaults() *GraphicsCardRelationship`

NewGraphicsCardRelationshipWithDefaults instantiates a new GraphicsCardRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *GraphicsCardRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *GraphicsCardRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *GraphicsCardRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *GraphicsCardRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *GraphicsCardRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *GraphicsCardRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *GraphicsCardRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *GraphicsCardRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GraphicsCardRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GraphicsCardRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GraphicsCardRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *GraphicsCardRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *GraphicsCardRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *GraphicsCardRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *GraphicsCardRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *GraphicsCardRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *GraphicsCardRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *GraphicsCardRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *GraphicsCardRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *GraphicsCardRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *GraphicsCardRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *GraphicsCardRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *GraphicsCardRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *GraphicsCardRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GraphicsCardRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GraphicsCardRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *GraphicsCardRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *GraphicsCardRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *GraphicsCardRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *GraphicsCardRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *GraphicsCardRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *GraphicsCardRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *GraphicsCardRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *GraphicsCardRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *GraphicsCardRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GraphicsCardRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GraphicsCardRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GraphicsCardRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *GraphicsCardRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *GraphicsCardRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *GraphicsCardRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *GraphicsCardRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *GraphicsCardRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *GraphicsCardRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *GraphicsCardRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *GraphicsCardRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *GraphicsCardRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *GraphicsCardRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *GraphicsCardRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GraphicsCardRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GraphicsCardRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GraphicsCardRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *GraphicsCardRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *GraphicsCardRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *GraphicsCardRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *GraphicsCardRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *GraphicsCardRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *GraphicsCardRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *GraphicsCardRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *GraphicsCardRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *GraphicsCardRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *GraphicsCardRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *GraphicsCardRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *GraphicsCardRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *GraphicsCardRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *GraphicsCardRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *GraphicsCardRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *GraphicsCardRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *GraphicsCardRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *GraphicsCardRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *GraphicsCardRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *GraphicsCardRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *GraphicsCardRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *GraphicsCardRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *GraphicsCardRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *GraphicsCardRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *GraphicsCardRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GraphicsCardRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GraphicsCardRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GraphicsCardRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *GraphicsCardRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GraphicsCardRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GraphicsCardRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GraphicsCardRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *GraphicsCardRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GraphicsCardRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GraphicsCardRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GraphicsCardRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *GraphicsCardRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GraphicsCardRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GraphicsCardRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GraphicsCardRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCardId

`func (o *GraphicsCardRelationship) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GraphicsCardRelationship) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GraphicsCardRelationship) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *GraphicsCardRelationship) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetDeviceId

`func (o *GraphicsCardRelationship) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GraphicsCardRelationship) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GraphicsCardRelationship) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GraphicsCardRelationship) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpanderSlot

`func (o *GraphicsCardRelationship) GetExpanderSlot() string`

GetExpanderSlot returns the ExpanderSlot field if non-nil, zero value otherwise.

### GetExpanderSlotOk

`func (o *GraphicsCardRelationship) GetExpanderSlotOk() (*string, bool)`

GetExpanderSlotOk returns a tuple with the ExpanderSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderSlot

`func (o *GraphicsCardRelationship) SetExpanderSlot(v string)`

SetExpanderSlot sets ExpanderSlot field to given value.

### HasExpanderSlot

`func (o *GraphicsCardRelationship) HasExpanderSlot() bool`

HasExpanderSlot returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GraphicsCardRelationship) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GraphicsCardRelationship) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GraphicsCardRelationship) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GraphicsCardRelationship) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetMode

`func (o *GraphicsCardRelationship) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GraphicsCardRelationship) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GraphicsCardRelationship) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GraphicsCardRelationship) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumGpus

`func (o *GraphicsCardRelationship) GetNumGpus() string`

GetNumGpus returns the NumGpus field if non-nil, zero value otherwise.

### GetNumGpusOk

`func (o *GraphicsCardRelationship) GetNumGpusOk() (*string, bool)`

GetNumGpusOk returns a tuple with the NumGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGpus

`func (o *GraphicsCardRelationship) SetNumGpus(v string)`

SetNumGpus sets NumGpus field to given value.

### HasNumGpus

`func (o *GraphicsCardRelationship) HasNumGpus() bool`

HasNumGpus returns a boolean if a field has been set.

### GetOperState

`func (o *GraphicsCardRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *GraphicsCardRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *GraphicsCardRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *GraphicsCardRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPciAddress

`func (o *GraphicsCardRelationship) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *GraphicsCardRelationship) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *GraphicsCardRelationship) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *GraphicsCardRelationship) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPciAddressList

`func (o *GraphicsCardRelationship) GetPciAddressList() string`

GetPciAddressList returns the PciAddressList field if non-nil, zero value otherwise.

### GetPciAddressListOk

`func (o *GraphicsCardRelationship) GetPciAddressListOk() (*string, bool)`

GetPciAddressListOk returns a tuple with the PciAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddressList

`func (o *GraphicsCardRelationship) SetPciAddressList(v string)`

SetPciAddressList sets PciAddressList field to given value.

### HasPciAddressList

`func (o *GraphicsCardRelationship) HasPciAddressList() bool`

HasPciAddressList returns a boolean if a field has been set.

### GetPciSlot

`func (o *GraphicsCardRelationship) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *GraphicsCardRelationship) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *GraphicsCardRelationship) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *GraphicsCardRelationship) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetComputeBlade

`func (o *GraphicsCardRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *GraphicsCardRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *GraphicsCardRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *GraphicsCardRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *GraphicsCardRelationship) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *GraphicsCardRelationship) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *GraphicsCardRelationship) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *GraphicsCardRelationship) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *GraphicsCardRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *GraphicsCardRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *GraphicsCardRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *GraphicsCardRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetGraphicsControllers

`func (o *GraphicsCardRelationship) GetGraphicsControllers() []GraphicsControllerRelationship`

GetGraphicsControllers returns the GraphicsControllers field if non-nil, zero value otherwise.

### GetGraphicsControllersOk

`func (o *GraphicsCardRelationship) GetGraphicsControllersOk() (*[]GraphicsControllerRelationship, bool)`

GetGraphicsControllersOk returns a tuple with the GraphicsControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsControllers

`func (o *GraphicsCardRelationship) SetGraphicsControllers(v []GraphicsControllerRelationship)`

SetGraphicsControllers sets GraphicsControllers field to given value.

### HasGraphicsControllers

`func (o *GraphicsCardRelationship) HasGraphicsControllers() bool`

HasGraphicsControllers returns a boolean if a field has been set.

### SetGraphicsControllersNil

`func (o *GraphicsCardRelationship) SetGraphicsControllersNil(b bool)`

 SetGraphicsControllersNil sets the value for GraphicsControllers to be an explicit nil

### UnsetGraphicsControllers
`func (o *GraphicsCardRelationship) UnsetGraphicsControllers()`

UnsetGraphicsControllers ensures that no value is present for GraphicsControllers, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *GraphicsCardRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *GraphicsCardRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *GraphicsCardRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *GraphicsCardRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *GraphicsCardRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *GraphicsCardRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *GraphicsCardRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *GraphicsCardRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *GraphicsCardRelationship) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *GraphicsCardRelationship) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *GraphicsCardRelationship) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *GraphicsCardRelationship) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *GraphicsCardRelationship) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *GraphicsCardRelationship) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


