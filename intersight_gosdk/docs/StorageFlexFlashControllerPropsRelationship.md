# StorageFlexFlashControllerPropsRelationship

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
**CardsManageable** | Pointer to **string** | Manageable card on the flex flash controller. | [optional] 
**ConfiguredMode** | Pointer to **string** | Mode configured on the flex flash controller. | [optional] 
**ControllerName** | Pointer to **string** | The current name of the flex flash controller. | [optional] 
**ControllerStatus** | Pointer to **string** | The current status of the flex flash controller. | [optional] 
**FwVersion** | Pointer to **string** | Firmware version of the flex flash controller. | [optional] 
**InternalState** | Pointer to **string** | Internal state of the flex flash controller. | [optional] 
**OperatingMode** | Pointer to **string** | Operating mode of flex flash controller. | [optional] 
**PhysicalDriveCount** | Pointer to **string** | Number of connected physical drives to a specific Flex flash controller. | [optional] 
**ProductName** | Pointer to **string** | Product name of the flex flash controller. | [optional] 
**StartupFwVersion** | Pointer to **string** | Startup firmware version of the Flex flash controller. | [optional] 
**VirtualDriveCount** | Pointer to **string** | Number of virtual drives for a specific Flex flash controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](storage.FlexFlashController.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashControllerPropsRelationship

`func NewStorageFlexFlashControllerPropsRelationship(classId string, objectType string, ) *StorageFlexFlashControllerPropsRelationship`

NewStorageFlexFlashControllerPropsRelationship instantiates a new StorageFlexFlashControllerPropsRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerPropsRelationshipWithDefaults

`func NewStorageFlexFlashControllerPropsRelationshipWithDefaults() *StorageFlexFlashControllerPropsRelationship`

NewStorageFlexFlashControllerPropsRelationshipWithDefaults instantiates a new StorageFlexFlashControllerPropsRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *StorageFlexFlashControllerPropsRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *StorageFlexFlashControllerPropsRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *StorageFlexFlashControllerPropsRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *StorageFlexFlashControllerPropsRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashControllerPropsRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *StorageFlexFlashControllerPropsRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *StorageFlexFlashControllerPropsRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *StorageFlexFlashControllerPropsRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *StorageFlexFlashControllerPropsRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *StorageFlexFlashControllerPropsRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *StorageFlexFlashControllerPropsRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *StorageFlexFlashControllerPropsRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *StorageFlexFlashControllerPropsRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *StorageFlexFlashControllerPropsRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *StorageFlexFlashControllerPropsRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *StorageFlexFlashControllerPropsRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *StorageFlexFlashControllerPropsRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *StorageFlexFlashControllerPropsRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashControllerPropsRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *StorageFlexFlashControllerPropsRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *StorageFlexFlashControllerPropsRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *StorageFlexFlashControllerPropsRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *StorageFlexFlashControllerPropsRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *StorageFlexFlashControllerPropsRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *StorageFlexFlashControllerPropsRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *StorageFlexFlashControllerPropsRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StorageFlexFlashControllerPropsRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StorageFlexFlashControllerPropsRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *StorageFlexFlashControllerPropsRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *StorageFlexFlashControllerPropsRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *StorageFlexFlashControllerPropsRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *StorageFlexFlashControllerPropsRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *StorageFlexFlashControllerPropsRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *StorageFlexFlashControllerPropsRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *StorageFlexFlashControllerPropsRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *StorageFlexFlashControllerPropsRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *StorageFlexFlashControllerPropsRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageFlexFlashControllerPropsRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageFlexFlashControllerPropsRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *StorageFlexFlashControllerPropsRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *StorageFlexFlashControllerPropsRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *StorageFlexFlashControllerPropsRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *StorageFlexFlashControllerPropsRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *StorageFlexFlashControllerPropsRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *StorageFlexFlashControllerPropsRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *StorageFlexFlashControllerPropsRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *StorageFlexFlashControllerPropsRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *StorageFlexFlashControllerPropsRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *StorageFlexFlashControllerPropsRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *StorageFlexFlashControllerPropsRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *StorageFlexFlashControllerPropsRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *StorageFlexFlashControllerPropsRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *StorageFlexFlashControllerPropsRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *StorageFlexFlashControllerPropsRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *StorageFlexFlashControllerPropsRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *StorageFlexFlashControllerPropsRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *StorageFlexFlashControllerPropsRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *StorageFlexFlashControllerPropsRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *StorageFlexFlashControllerPropsRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StorageFlexFlashControllerPropsRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StorageFlexFlashControllerPropsRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *StorageFlexFlashControllerPropsRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageFlexFlashControllerPropsRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StorageFlexFlashControllerPropsRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *StorageFlexFlashControllerPropsRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageFlexFlashControllerPropsRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageFlexFlashControllerPropsRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *StorageFlexFlashControllerPropsRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StorageFlexFlashControllerPropsRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StorageFlexFlashControllerPropsRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCardsManageable

`func (o *StorageFlexFlashControllerPropsRelationship) GetCardsManageable() string`

GetCardsManageable returns the CardsManageable field if non-nil, zero value otherwise.

### GetCardsManageableOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetCardsManageableOk() (*string, bool)`

GetCardsManageableOk returns a tuple with the CardsManageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsManageable

`func (o *StorageFlexFlashControllerPropsRelationship) SetCardsManageable(v string)`

SetCardsManageable sets CardsManageable field to given value.

### HasCardsManageable

`func (o *StorageFlexFlashControllerPropsRelationship) HasCardsManageable() bool`

HasCardsManageable returns a boolean if a field has been set.

### GetConfiguredMode

`func (o *StorageFlexFlashControllerPropsRelationship) GetConfiguredMode() string`

GetConfiguredMode returns the ConfiguredMode field if non-nil, zero value otherwise.

### GetConfiguredModeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetConfiguredModeOk() (*string, bool)`

GetConfiguredModeOk returns a tuple with the ConfiguredMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMode

`func (o *StorageFlexFlashControllerPropsRelationship) SetConfiguredMode(v string)`

SetConfiguredMode sets ConfiguredMode field to given value.

### HasConfiguredMode

`func (o *StorageFlexFlashControllerPropsRelationship) HasConfiguredMode() bool`

HasConfiguredMode returns a boolean if a field has been set.

### GetControllerName

`func (o *StorageFlexFlashControllerPropsRelationship) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *StorageFlexFlashControllerPropsRelationship) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *StorageFlexFlashControllerPropsRelationship) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageFlexFlashControllerPropsRelationship) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageFlexFlashControllerPropsRelationship) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageFlexFlashControllerPropsRelationship) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) GetFwVersion() string`

GetFwVersion returns the FwVersion field if non-nil, zero value otherwise.

### GetFwVersionOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetFwVersionOk() (*string, bool)`

GetFwVersionOk returns a tuple with the FwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) SetFwVersion(v string)`

SetFwVersion sets FwVersion field to given value.

### HasFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) HasFwVersion() bool`

HasFwVersion returns a boolean if a field has been set.

### GetInternalState

`func (o *StorageFlexFlashControllerPropsRelationship) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *StorageFlexFlashControllerPropsRelationship) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *StorageFlexFlashControllerPropsRelationship) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetOperatingMode

`func (o *StorageFlexFlashControllerPropsRelationship) GetOperatingMode() string`

GetOperatingMode returns the OperatingMode field if non-nil, zero value otherwise.

### GetOperatingModeOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetOperatingModeOk() (*string, bool)`

GetOperatingModeOk returns a tuple with the OperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMode

`func (o *StorageFlexFlashControllerPropsRelationship) SetOperatingMode(v string)`

SetOperatingMode sets OperatingMode field to given value.

### HasOperatingMode

`func (o *StorageFlexFlashControllerPropsRelationship) HasOperatingMode() bool`

HasOperatingMode returns a boolean if a field has been set.

### GetPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) GetPhysicalDriveCount() string`

GetPhysicalDriveCount returns the PhysicalDriveCount field if non-nil, zero value otherwise.

### GetPhysicalDriveCountOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetPhysicalDriveCountOk() (*string, bool)`

GetPhysicalDriveCountOk returns a tuple with the PhysicalDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) SetPhysicalDriveCount(v string)`

SetPhysicalDriveCount sets PhysicalDriveCount field to given value.

### HasPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) HasPhysicalDriveCount() bool`

HasPhysicalDriveCount returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexFlashControllerPropsRelationship) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexFlashControllerPropsRelationship) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexFlashControllerPropsRelationship) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStartupFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) GetStartupFwVersion() string`

GetStartupFwVersion returns the StartupFwVersion field if non-nil, zero value otherwise.

### GetStartupFwVersionOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetStartupFwVersionOk() (*string, bool)`

GetStartupFwVersionOk returns a tuple with the StartupFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) SetStartupFwVersion(v string)`

SetStartupFwVersion sets StartupFwVersion field to given value.

### HasStartupFwVersion

`func (o *StorageFlexFlashControllerPropsRelationship) HasStartupFwVersion() bool`

HasStartupFwVersion returns a boolean if a field has been set.

### GetVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) GetVirtualDriveCount() string`

GetVirtualDriveCount returns the VirtualDriveCount field if non-nil, zero value otherwise.

### GetVirtualDriveCountOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetVirtualDriveCountOk() (*string, bool)`

GetVirtualDriveCountOk returns a tuple with the VirtualDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) SetVirtualDriveCount(v string)`

SetVirtualDriveCount sets VirtualDriveCount field to given value.

### HasVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsRelationship) HasVirtualDriveCount() bool`

HasVirtualDriveCount returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerPropsRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerPropsRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerPropsRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsRelationship) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashControllerPropsRelationship) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsRelationship) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsRelationship) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


