# ProcessorUnitRelationship

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
**Architecture** | Pointer to **string** | The architecture of the installed processor. | [optional] [readonly] 
**NumCores** | Pointer to **int64** | The number of cores present in a given processor. | [optional] [readonly] 
**NumCoresEnabled** | Pointer to **string** | The number of enabled cores in the installed processor. | [optional] [readonly] 
**NumThreads** | Pointer to **string** | The maximum number of threads available in the installed processor. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The power state of the processor. | [optional] [readonly] 
**OperState** | Pointer to **string** | The health indicator of the processor, &#39;OK&#39; indicates the processor is operatinal. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of the central processing unit. | [optional] [readonly] 
**Presence** | Pointer to **string** | The valid values are &#39;equipped&#39; and &#39;absent&#39;. | [optional] [readonly] 
**ProcessorId** | Pointer to **int64** | The ID number of a given processor. | [optional] [readonly] 
**SocketDesignation** | Pointer to **string** | The socket ID of the installed processor. | [optional] [readonly] 
**Speed** | Pointer to **float32** | The maximum speed of the installed processor in GHz. | [optional] [readonly] 
**Stepping** | Pointer to **string** | The CPU stepping of the installed processor. | [optional] [readonly] 
**Thermal** | Pointer to **string** | The temperature of the processor in centigrade. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewProcessorUnitRelationship

`func NewProcessorUnitRelationship(classId string, objectType string, ) *ProcessorUnitRelationship`

NewProcessorUnitRelationship instantiates a new ProcessorUnitRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorUnitRelationshipWithDefaults

`func NewProcessorUnitRelationshipWithDefaults() *ProcessorUnitRelationship`

NewProcessorUnitRelationshipWithDefaults instantiates a new ProcessorUnitRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ProcessorUnitRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ProcessorUnitRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ProcessorUnitRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ProcessorUnitRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ProcessorUnitRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ProcessorUnitRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ProcessorUnitRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ProcessorUnitRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ProcessorUnitRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ProcessorUnitRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ProcessorUnitRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ProcessorUnitRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ProcessorUnitRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ProcessorUnitRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ProcessorUnitRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ProcessorUnitRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ProcessorUnitRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ProcessorUnitRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ProcessorUnitRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ProcessorUnitRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ProcessorUnitRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ProcessorUnitRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ProcessorUnitRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ProcessorUnitRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProcessorUnitRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProcessorUnitRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ProcessorUnitRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ProcessorUnitRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ProcessorUnitRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ProcessorUnitRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ProcessorUnitRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ProcessorUnitRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ProcessorUnitRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ProcessorUnitRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ProcessorUnitRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProcessorUnitRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProcessorUnitRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProcessorUnitRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ProcessorUnitRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ProcessorUnitRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ProcessorUnitRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ProcessorUnitRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ProcessorUnitRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ProcessorUnitRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ProcessorUnitRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ProcessorUnitRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ProcessorUnitRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ProcessorUnitRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ProcessorUnitRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ProcessorUnitRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ProcessorUnitRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ProcessorUnitRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ProcessorUnitRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ProcessorUnitRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ProcessorUnitRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ProcessorUnitRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ProcessorUnitRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ProcessorUnitRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ProcessorUnitRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ProcessorUnitRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ProcessorUnitRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ProcessorUnitRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ProcessorUnitRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ProcessorUnitRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ProcessorUnitRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ProcessorUnitRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ProcessorUnitRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ProcessorUnitRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ProcessorUnitRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ProcessorUnitRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ProcessorUnitRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ProcessorUnitRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ProcessorUnitRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ProcessorUnitRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ProcessorUnitRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ProcessorUnitRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *ProcessorUnitRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProcessorUnitRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProcessorUnitRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProcessorUnitRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ProcessorUnitRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ProcessorUnitRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ProcessorUnitRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ProcessorUnitRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *ProcessorUnitRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ProcessorUnitRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ProcessorUnitRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ProcessorUnitRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *ProcessorUnitRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ProcessorUnitRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ProcessorUnitRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ProcessorUnitRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetArchitecture

`func (o *ProcessorUnitRelationship) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ProcessorUnitRelationship) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ProcessorUnitRelationship) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ProcessorUnitRelationship) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetNumCores

`func (o *ProcessorUnitRelationship) GetNumCores() int64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *ProcessorUnitRelationship) GetNumCoresOk() (*int64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *ProcessorUnitRelationship) SetNumCores(v int64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *ProcessorUnitRelationship) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetNumCoresEnabled

`func (o *ProcessorUnitRelationship) GetNumCoresEnabled() string`

GetNumCoresEnabled returns the NumCoresEnabled field if non-nil, zero value otherwise.

### GetNumCoresEnabledOk

`func (o *ProcessorUnitRelationship) GetNumCoresEnabledOk() (*string, bool)`

GetNumCoresEnabledOk returns a tuple with the NumCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresEnabled

`func (o *ProcessorUnitRelationship) SetNumCoresEnabled(v string)`

SetNumCoresEnabled sets NumCoresEnabled field to given value.

### HasNumCoresEnabled

`func (o *ProcessorUnitRelationship) HasNumCoresEnabled() bool`

HasNumCoresEnabled returns a boolean if a field has been set.

### GetNumThreads

`func (o *ProcessorUnitRelationship) GetNumThreads() string`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ProcessorUnitRelationship) GetNumThreadsOk() (*string, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ProcessorUnitRelationship) SetNumThreads(v string)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ProcessorUnitRelationship) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ProcessorUnitRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ProcessorUnitRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ProcessorUnitRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ProcessorUnitRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *ProcessorUnitRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ProcessorUnitRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ProcessorUnitRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ProcessorUnitRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ProcessorUnitRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ProcessorUnitRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ProcessorUnitRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ProcessorUnitRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPresence

`func (o *ProcessorUnitRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ProcessorUnitRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ProcessorUnitRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ProcessorUnitRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProcessorId

`func (o *ProcessorUnitRelationship) GetProcessorId() int64`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *ProcessorUnitRelationship) GetProcessorIdOk() (*int64, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *ProcessorUnitRelationship) SetProcessorId(v int64)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *ProcessorUnitRelationship) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetSocketDesignation

`func (o *ProcessorUnitRelationship) GetSocketDesignation() string`

GetSocketDesignation returns the SocketDesignation field if non-nil, zero value otherwise.

### GetSocketDesignationOk

`func (o *ProcessorUnitRelationship) GetSocketDesignationOk() (*string, bool)`

GetSocketDesignationOk returns a tuple with the SocketDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketDesignation

`func (o *ProcessorUnitRelationship) SetSocketDesignation(v string)`

SetSocketDesignation sets SocketDesignation field to given value.

### HasSocketDesignation

`func (o *ProcessorUnitRelationship) HasSocketDesignation() bool`

HasSocketDesignation returns a boolean if a field has been set.

### GetSpeed

`func (o *ProcessorUnitRelationship) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ProcessorUnitRelationship) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ProcessorUnitRelationship) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ProcessorUnitRelationship) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStepping

`func (o *ProcessorUnitRelationship) GetStepping() string`

GetStepping returns the Stepping field if non-nil, zero value otherwise.

### GetSteppingOk

`func (o *ProcessorUnitRelationship) GetSteppingOk() (*string, bool)`

GetSteppingOk returns a tuple with the Stepping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepping

`func (o *ProcessorUnitRelationship) SetStepping(v string)`

SetStepping sets Stepping field to given value.

### HasStepping

`func (o *ProcessorUnitRelationship) HasStepping() bool`

HasStepping returns a boolean if a field has been set.

### GetThermal

`func (o *ProcessorUnitRelationship) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *ProcessorUnitRelationship) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *ProcessorUnitRelationship) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *ProcessorUnitRelationship) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetComputeBlade

`func (o *ProcessorUnitRelationship) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *ProcessorUnitRelationship) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *ProcessorUnitRelationship) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *ProcessorUnitRelationship) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeBoard

`func (o *ProcessorUnitRelationship) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *ProcessorUnitRelationship) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *ProcessorUnitRelationship) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *ProcessorUnitRelationship) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *ProcessorUnitRelationship) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *ProcessorUnitRelationship) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *ProcessorUnitRelationship) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *ProcessorUnitRelationship) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ProcessorUnitRelationship) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ProcessorUnitRelationship) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ProcessorUnitRelationship) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ProcessorUnitRelationship) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ProcessorUnitRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ProcessorUnitRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ProcessorUnitRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ProcessorUnitRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


