# ComputePhysicalRelationship

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
**AdminPowerState** | Pointer to **string** | The desired power state of the server. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**ComputeAlarmSummary**](compute.AlarmSummary.md) |  | [optional] 
**AssetTag** | Pointer to **string** | The user defined asset tag assigned to the server. | [optional] [readonly] 
**AvailableMemory** | Pointer to **int64** | The amount of memory available on the server. | [optional] [readonly] 
**BiosPostComplete** | Pointer to **bool** | The BIOS POST completion status of the server. | [optional] 
**FaultSummary** | Pointer to **int64** | The fault summary for the server. | [optional] 
**KvmIpAddresses** | Pointer to [**[]ComputeIpAddress**](compute.IpAddress.md) |  | [optional] 
**ManagementMode** | Pointer to **string** | The management mode of the server. * &#x60;IntersightStandalone&#x60; - Intersight Standalone mode of operation. * &#x60;UCSM&#x60; - Unified Computing System Manager mode of operation. * &#x60;Intersight&#x60; - Intersight managed mode of operation. | [optional] [default to "IntersightStandalone"]
**MemorySpeed** | Pointer to **string** | The maximum memory speed in MHz available on the server. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | Management address of the server. | [optional] 
**NumAdaptors** | Pointer to **int64** | The total number of network adapters present on the server. | [optional] [readonly] 
**NumCpuCores** | Pointer to **int64** | The total number of CPU cores present on the server. | [optional] [readonly] 
**NumCpuCoresEnabled** | Pointer to **int64** | The total number of CPU cores enabled on the server. | [optional] [readonly] 
**NumCpus** | Pointer to **int64** | The total number of CPUs present on the server. | [optional] [readonly] 
**NumEthHostInterfaces** | Pointer to **int64** | The total number of vNICs which are visible to a host on the server. | [optional] [readonly] 
**NumFcHostInterfaces** | Pointer to **int64** | The total number of vHBAs which are visible to a host on the server. | [optional] [readonly] 
**NumThreads** | Pointer to **int64** | The total number of threads the server is capable of handling. | [optional] [readonly] 
**OperPowerState** | Pointer to **string** | The actual power state of the server. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the server. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability of the server. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the registered device - whether managed by UCSM or operating in standalone mode. | [optional] 
**Presence** | Pointer to **string** | Indicates if a server is present in a slot and is applicable for blade servers. | [optional] [readonly] 
**ServiceProfile** | Pointer to **string** | The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM. | [optional] [readonly] 
**TotalMemory** | Pointer to **int64** | The total memory available on the server. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the server. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The universally unique identity of the server. | [optional] [readonly] 
**MgmtIdentity** | Pointer to [**EquipmentPhysicalIdentityRelationship**](equipment.PhysicalIdentity.Relationship.md) |  | [optional] 

## Methods

### NewComputePhysicalRelationship

`func NewComputePhysicalRelationship(classId string, objectType string, ) *ComputePhysicalRelationship`

NewComputePhysicalRelationship instantiates a new ComputePhysicalRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePhysicalRelationshipWithDefaults

`func NewComputePhysicalRelationshipWithDefaults() *ComputePhysicalRelationship`

NewComputePhysicalRelationshipWithDefaults instantiates a new ComputePhysicalRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ComputePhysicalRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ComputePhysicalRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ComputePhysicalRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ComputePhysicalRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ComputePhysicalRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePhysicalRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePhysicalRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ComputePhysicalRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ComputePhysicalRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ComputePhysicalRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ComputePhysicalRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ComputePhysicalRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ComputePhysicalRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ComputePhysicalRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ComputePhysicalRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ComputePhysicalRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ComputePhysicalRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ComputePhysicalRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ComputePhysicalRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ComputePhysicalRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ComputePhysicalRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ComputePhysicalRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ComputePhysicalRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ComputePhysicalRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePhysicalRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePhysicalRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ComputePhysicalRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ComputePhysicalRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ComputePhysicalRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ComputePhysicalRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ComputePhysicalRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ComputePhysicalRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ComputePhysicalRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ComputePhysicalRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ComputePhysicalRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComputePhysicalRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComputePhysicalRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ComputePhysicalRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ComputePhysicalRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ComputePhysicalRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ComputePhysicalRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ComputePhysicalRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ComputePhysicalRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ComputePhysicalRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ComputePhysicalRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ComputePhysicalRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ComputePhysicalRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ComputePhysicalRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ComputePhysicalRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ComputePhysicalRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ComputePhysicalRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ComputePhysicalRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ComputePhysicalRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ComputePhysicalRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ComputePhysicalRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ComputePhysicalRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ComputePhysicalRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ComputePhysicalRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ComputePhysicalRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ComputePhysicalRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ComputePhysicalRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ComputePhysicalRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ComputePhysicalRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ComputePhysicalRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceMoId

`func (o *ComputePhysicalRelationship) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *ComputePhysicalRelationship) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *ComputePhysicalRelationship) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *ComputePhysicalRelationship) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetDn

`func (o *ComputePhysicalRelationship) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputePhysicalRelationship) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputePhysicalRelationship) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputePhysicalRelationship) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRn

`func (o *ComputePhysicalRelationship) GetRn() string`

GetRn returns the Rn field if non-nil, zero value otherwise.

### GetRnOk

`func (o *ComputePhysicalRelationship) GetRnOk() (*string, bool)`

GetRnOk returns a tuple with the Rn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRn

`func (o *ComputePhysicalRelationship) SetRn(v string)`

SetRn sets Rn field to given value.

### HasRn

`func (o *ComputePhysicalRelationship) HasRn() bool`

HasRn returns a boolean if a field has been set.

### GetModel

`func (o *ComputePhysicalRelationship) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputePhysicalRelationship) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputePhysicalRelationship) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputePhysicalRelationship) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ComputePhysicalRelationship) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputePhysicalRelationship) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputePhysicalRelationship) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputePhysicalRelationship) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerial

`func (o *ComputePhysicalRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ComputePhysicalRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ComputePhysicalRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ComputePhysicalRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *ComputePhysicalRelationship) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputePhysicalRelationship) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputePhysicalRelationship) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputePhysicalRelationship) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAdminPowerState

`func (o *ComputePhysicalRelationship) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *ComputePhysicalRelationship) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *ComputePhysicalRelationship) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *ComputePhysicalRelationship) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetAlarmSummary

`func (o *ComputePhysicalRelationship) GetAlarmSummary() ComputeAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ComputePhysicalRelationship) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ComputePhysicalRelationship) SetAlarmSummary(v ComputeAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ComputePhysicalRelationship) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputePhysicalRelationship) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputePhysicalRelationship) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputePhysicalRelationship) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputePhysicalRelationship) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetAvailableMemory

`func (o *ComputePhysicalRelationship) GetAvailableMemory() int64`

GetAvailableMemory returns the AvailableMemory field if non-nil, zero value otherwise.

### GetAvailableMemoryOk

`func (o *ComputePhysicalRelationship) GetAvailableMemoryOk() (*int64, bool)`

GetAvailableMemoryOk returns a tuple with the AvailableMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMemory

`func (o *ComputePhysicalRelationship) SetAvailableMemory(v int64)`

SetAvailableMemory sets AvailableMemory field to given value.

### HasAvailableMemory

`func (o *ComputePhysicalRelationship) HasAvailableMemory() bool`

HasAvailableMemory returns a boolean if a field has been set.

### GetBiosPostComplete

`func (o *ComputePhysicalRelationship) GetBiosPostComplete() bool`

GetBiosPostComplete returns the BiosPostComplete field if non-nil, zero value otherwise.

### GetBiosPostCompleteOk

`func (o *ComputePhysicalRelationship) GetBiosPostCompleteOk() (*bool, bool)`

GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosPostComplete

`func (o *ComputePhysicalRelationship) SetBiosPostComplete(v bool)`

SetBiosPostComplete sets BiosPostComplete field to given value.

### HasBiosPostComplete

`func (o *ComputePhysicalRelationship) HasBiosPostComplete() bool`

HasBiosPostComplete returns a boolean if a field has been set.

### GetFaultSummary

`func (o *ComputePhysicalRelationship) GetFaultSummary() int64`

GetFaultSummary returns the FaultSummary field if non-nil, zero value otherwise.

### GetFaultSummaryOk

`func (o *ComputePhysicalRelationship) GetFaultSummaryOk() (*int64, bool)`

GetFaultSummaryOk returns a tuple with the FaultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultSummary

`func (o *ComputePhysicalRelationship) SetFaultSummary(v int64)`

SetFaultSummary sets FaultSummary field to given value.

### HasFaultSummary

`func (o *ComputePhysicalRelationship) HasFaultSummary() bool`

HasFaultSummary returns a boolean if a field has been set.

### GetKvmIpAddresses

`func (o *ComputePhysicalRelationship) GetKvmIpAddresses() []ComputeIpAddress`

GetKvmIpAddresses returns the KvmIpAddresses field if non-nil, zero value otherwise.

### GetKvmIpAddressesOk

`func (o *ComputePhysicalRelationship) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool)`

GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmIpAddresses

`func (o *ComputePhysicalRelationship) SetKvmIpAddresses(v []ComputeIpAddress)`

SetKvmIpAddresses sets KvmIpAddresses field to given value.

### HasKvmIpAddresses

`func (o *ComputePhysicalRelationship) HasKvmIpAddresses() bool`

HasKvmIpAddresses returns a boolean if a field has been set.

### GetManagementMode

`func (o *ComputePhysicalRelationship) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *ComputePhysicalRelationship) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *ComputePhysicalRelationship) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *ComputePhysicalRelationship) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetMemorySpeed

`func (o *ComputePhysicalRelationship) GetMemorySpeed() string`

GetMemorySpeed returns the MemorySpeed field if non-nil, zero value otherwise.

### GetMemorySpeedOk

`func (o *ComputePhysicalRelationship) GetMemorySpeedOk() (*string, bool)`

GetMemorySpeedOk returns a tuple with the MemorySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySpeed

`func (o *ComputePhysicalRelationship) SetMemorySpeed(v string)`

SetMemorySpeed sets MemorySpeed field to given value.

### HasMemorySpeed

`func (o *ComputePhysicalRelationship) HasMemorySpeed() bool`

HasMemorySpeed returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *ComputePhysicalRelationship) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *ComputePhysicalRelationship) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *ComputePhysicalRelationship) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *ComputePhysicalRelationship) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNumAdaptors

`func (o *ComputePhysicalRelationship) GetNumAdaptors() int64`

GetNumAdaptors returns the NumAdaptors field if non-nil, zero value otherwise.

### GetNumAdaptorsOk

`func (o *ComputePhysicalRelationship) GetNumAdaptorsOk() (*int64, bool)`

GetNumAdaptorsOk returns a tuple with the NumAdaptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdaptors

`func (o *ComputePhysicalRelationship) SetNumAdaptors(v int64)`

SetNumAdaptors sets NumAdaptors field to given value.

### HasNumAdaptors

`func (o *ComputePhysicalRelationship) HasNumAdaptors() bool`

HasNumAdaptors returns a boolean if a field has been set.

### GetNumCpuCores

`func (o *ComputePhysicalRelationship) GetNumCpuCores() int64`

GetNumCpuCores returns the NumCpuCores field if non-nil, zero value otherwise.

### GetNumCpuCoresOk

`func (o *ComputePhysicalRelationship) GetNumCpuCoresOk() (*int64, bool)`

GetNumCpuCoresOk returns a tuple with the NumCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCores

`func (o *ComputePhysicalRelationship) SetNumCpuCores(v int64)`

SetNumCpuCores sets NumCpuCores field to given value.

### HasNumCpuCores

`func (o *ComputePhysicalRelationship) HasNumCpuCores() bool`

HasNumCpuCores returns a boolean if a field has been set.

### GetNumCpuCoresEnabled

`func (o *ComputePhysicalRelationship) GetNumCpuCoresEnabled() int64`

GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field if non-nil, zero value otherwise.

### GetNumCpuCoresEnabledOk

`func (o *ComputePhysicalRelationship) GetNumCpuCoresEnabledOk() (*int64, bool)`

GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpuCoresEnabled

`func (o *ComputePhysicalRelationship) SetNumCpuCoresEnabled(v int64)`

SetNumCpuCoresEnabled sets NumCpuCoresEnabled field to given value.

### HasNumCpuCoresEnabled

`func (o *ComputePhysicalRelationship) HasNumCpuCoresEnabled() bool`

HasNumCpuCoresEnabled returns a boolean if a field has been set.

### GetNumCpus

`func (o *ComputePhysicalRelationship) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *ComputePhysicalRelationship) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *ComputePhysicalRelationship) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *ComputePhysicalRelationship) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumEthHostInterfaces

`func (o *ComputePhysicalRelationship) GetNumEthHostInterfaces() int64`

GetNumEthHostInterfaces returns the NumEthHostInterfaces field if non-nil, zero value otherwise.

### GetNumEthHostInterfacesOk

`func (o *ComputePhysicalRelationship) GetNumEthHostInterfacesOk() (*int64, bool)`

GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEthHostInterfaces

`func (o *ComputePhysicalRelationship) SetNumEthHostInterfaces(v int64)`

SetNumEthHostInterfaces sets NumEthHostInterfaces field to given value.

### HasNumEthHostInterfaces

`func (o *ComputePhysicalRelationship) HasNumEthHostInterfaces() bool`

HasNumEthHostInterfaces returns a boolean if a field has been set.

### GetNumFcHostInterfaces

`func (o *ComputePhysicalRelationship) GetNumFcHostInterfaces() int64`

GetNumFcHostInterfaces returns the NumFcHostInterfaces field if non-nil, zero value otherwise.

### GetNumFcHostInterfacesOk

`func (o *ComputePhysicalRelationship) GetNumFcHostInterfacesOk() (*int64, bool)`

GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFcHostInterfaces

`func (o *ComputePhysicalRelationship) SetNumFcHostInterfaces(v int64)`

SetNumFcHostInterfaces sets NumFcHostInterfaces field to given value.

### HasNumFcHostInterfaces

`func (o *ComputePhysicalRelationship) HasNumFcHostInterfaces() bool`

HasNumFcHostInterfaces returns a boolean if a field has been set.

### GetNumThreads

`func (o *ComputePhysicalRelationship) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ComputePhysicalRelationship) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ComputePhysicalRelationship) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ComputePhysicalRelationship) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetOperPowerState

`func (o *ComputePhysicalRelationship) GetOperPowerState() string`

GetOperPowerState returns the OperPowerState field if non-nil, zero value otherwise.

### GetOperPowerStateOk

`func (o *ComputePhysicalRelationship) GetOperPowerStateOk() (*string, bool)`

GetOperPowerStateOk returns a tuple with the OperPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperPowerState

`func (o *ComputePhysicalRelationship) SetOperPowerState(v string)`

SetOperPowerState sets OperPowerState field to given value.

### HasOperPowerState

`func (o *ComputePhysicalRelationship) HasOperPowerState() bool`

HasOperPowerState returns a boolean if a field has been set.

### GetOperState

`func (o *ComputePhysicalRelationship) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *ComputePhysicalRelationship) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *ComputePhysicalRelationship) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *ComputePhysicalRelationship) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *ComputePhysicalRelationship) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *ComputePhysicalRelationship) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *ComputePhysicalRelationship) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *ComputePhysicalRelationship) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPlatformType

`func (o *ComputePhysicalRelationship) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ComputePhysicalRelationship) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ComputePhysicalRelationship) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ComputePhysicalRelationship) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPresence

`func (o *ComputePhysicalRelationship) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *ComputePhysicalRelationship) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *ComputePhysicalRelationship) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *ComputePhysicalRelationship) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetServiceProfile

`func (o *ComputePhysicalRelationship) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ComputePhysicalRelationship) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ComputePhysicalRelationship) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ComputePhysicalRelationship) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### GetTotalMemory

`func (o *ComputePhysicalRelationship) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *ComputePhysicalRelationship) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *ComputePhysicalRelationship) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *ComputePhysicalRelationship) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetUserLabel

`func (o *ComputePhysicalRelationship) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ComputePhysicalRelationship) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ComputePhysicalRelationship) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ComputePhysicalRelationship) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetUuid

`func (o *ComputePhysicalRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComputePhysicalRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComputePhysicalRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComputePhysicalRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMgmtIdentity

`func (o *ComputePhysicalRelationship) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship`

GetMgmtIdentity returns the MgmtIdentity field if non-nil, zero value otherwise.

### GetMgmtIdentityOk

`func (o *ComputePhysicalRelationship) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool)`

GetMgmtIdentityOk returns a tuple with the MgmtIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIdentity

`func (o *ComputePhysicalRelationship) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship)`

SetMgmtIdentity sets MgmtIdentity field to given value.

### HasMgmtIdentity

`func (o *ComputePhysicalRelationship) HasMgmtIdentity() bool`

HasMgmtIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


